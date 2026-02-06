// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./User.sol";
import "./DroneObjectFactory.sol";

/// @title GridObjectFactory - 网格对象工厂
/// @notice 管理 10m×10m 网格的防控状态和维护记录
contract GridObjectFactory {

    enum GridStatus {
        Normal,             // 正常
        PsyllidRisk,        // 木虱风险
        MarkedSuspect,      // 已标记待确认
        ConfirmedInfected,  // 确认感染
        Removed,            // 已移除
        Replanted,          // 已补种
        Quarantine,         // 隔离观察
        Unknown             // 未知
    }

    enum GridTerrainType {
        Valley,     // 山谷
        Flat,       // 平地
        Ridge,      // 山脊
        Cliff       // 悬崖
    }

    enum CitrusDiseaseType {
        None,
        HuangLongBing,      // 黄龙病
        Anthracnose,        // 炭疽病
        Canker,             // 溃疡病
        Other
    }

    struct GridInfo {
        string gridCode;
        address farmId;
        GridTerrainType terrainType;
        GridStatus status;
        CitrusDiseaseType diseaseType;
        uint32 lastStatusUpdateTime;
        bool exists;
    }

    // 维护记录优化：使用映射替代数组
    struct MaintenanceRecord {
        uint32 timestamp;
        uint8 operationTypeCode;  // 1=生物防治, 2=施肥, 3=修剪, 4=诱捕带
        string detail;            // 简短描述
        bytes32 evidenceIPFSHash; // IPFS hash (32字节)
        address operator;
    }

    mapping(address => GridInfo) public grids;
    mapping(address => mapping(uint32 => MaintenanceRecord)) public maintenanceRecords;
    mapping(address => uint32) public maintenanceRecordCount;

    User public immutable userContract;
    DroneObjectFactory public immutable droneFactory;
    
    // 授权的调用者（FarmFactory）
    mapping(address => bool) public authorizedCallers;

    // ─── 事件 ────────────────────────────────────────────────
    event GridRegistered(
        address indexed farmId, 
        address indexed gridId, 
        string gridCode
    );
    event GridStatusUpdated(
        address indexed gridId, 
        GridStatus indexed oldStatus, 
        GridStatus indexed newStatus
    );
    event GridDiseaseTypeUpdated(
        address indexed gridId, 
        CitrusDiseaseType oldType, 
        CitrusDiseaseType newType
    );
    event MaintenanceRecorded(
        address indexed gridId, 
        uint32 indexed recordId, 
        uint8 operationType
    );
    event CallerAuthorized(address indexed caller);
    event CallerRevoked(address indexed caller);

    // ─── 修饰符 ──────────────────────────────────────────────
    modifier onlyAdmin() {
        require(
            userContract.getUserRole(msg.sender) == User.UserRole.Admin,
            "Only admin allowed"
        );
        _;
    }

    modifier onlyAuthorizedCaller() {
        require(
            authorizedCallers[msg.sender] || 
            userContract.getUserRole(msg.sender) == User.UserRole.Admin,
            "Unauthorized caller"
        );
        _;
    }

    modifier onlyFarmerOrDrone() {
        bool isValidUser = userContract.userExists(msg.sender) && 
                          userContract.getUserRole(msg.sender) == User.UserRole.Farmer;
        bool isValidDrone = droneFactory.droneExists(msg.sender);
        
        require(isValidUser || isValidDrone, "Unauthorized");
        _;
    }

    // ─── 构造函数 ────────────────────────────────────────────
    constructor(address _userContract, address _droneFactory) {
        require(_userContract != address(0), "Invalid user contract");
        require(_droneFactory != address(0), "Invalid drone factory");
        
        userContract = User(_userContract);
        droneFactory = DroneObjectFactory(_droneFactory);
    }

    // ─── 授权管理 ────────────────────────────────────────────
    
    /// @notice 授权调用者（如 FarmFactory）
    function authorizeCaller(address _caller) external onlyAdmin {
        require(_caller != address(0), "Invalid caller");
        authorizedCallers[_caller] = true;
        emit CallerAuthorized(_caller);
    }

    function revokeCaller(address _caller) external onlyAdmin {
        authorizedCallers[_caller] = false;
        emit CallerRevoked(_caller);
    }

    // ─── 核心函数 ────────────────────────────────────────────
    
    /// @notice 注册网格（由授权调用者如 FarmFactory 调用）
    function registerGrid(
        address gridId,
        string calldata gridCode,
        address farmId,
        GridTerrainType terrainType
    ) external onlyAuthorizedCaller {
        require(gridId != address(0), "Invalid gridId");
        require(bytes(gridCode).length > 0, "Empty gridCode");
        require(farmId != address(0), "Invalid farmId");
        require(!grids[gridId].exists, "Grid already registered");

        grids[gridId] = GridInfo({
            gridCode: gridCode,
            farmId: farmId,
            terrainType: terrainType,
            status: GridStatus.Normal,
            diseaseType: CitrusDiseaseType.None,
            lastStatusUpdateTime: uint32(block.timestamp),
            exists: true
        });

        emit GridRegistered(farmId, gridId, gridCode);
    }

    /// @notice 更新网格状态（授权调用者或农民/无人机）
    function updateGridStatus(
        address gridId,
        GridStatus newStatus
    ) external onlyAuthorizedCaller {
        require(grids[gridId].exists, "Grid not found");
        
        GridStatus oldStatus = grids[gridId].status;
        grids[gridId].status = newStatus;
        grids[gridId].lastStatusUpdateTime = uint32(block.timestamp);

        emit GridStatusUpdated(gridId, oldStatus, newStatus);
    }

    /// @notice 更新病害类型
    function updateDiseaseType(
        address gridId,
        CitrusDiseaseType newType
    ) external onlyAuthorizedCaller {
        require(grids[gridId].exists, "Grid not found");
        
        CitrusDiseaseType oldType = grids[gridId].diseaseType;
        grids[gridId].diseaseType = newType;
        grids[gridId].lastStatusUpdateTime = uint32(block.timestamp);

        emit GridDiseaseTypeUpdated(gridId, oldType, newType);
    }

    /// @notice 添加维护记录（优化版）
    function addMaintenanceRecord(
        address gridId,
        uint8 operationTypeCode,
        string calldata detail,
        bytes32 evidenceIPFSHash
    ) external onlyFarmerOrDrone {
        require(grids[gridId].exists, "Grid not found");
        require(operationTypeCode > 0 && operationTypeCode <= 4, "Invalid type");
        
        uint32 recordId = maintenanceRecordCount[gridId];
        
        maintenanceRecords[gridId][recordId] = MaintenanceRecord({
            timestamp: uint32(block.timestamp),
            operationTypeCode: operationTypeCode,
            detail: detail,
            evidenceIPFSHash: evidenceIPFSHash,
            operator: msg.sender
        });
        
        maintenanceRecordCount[gridId]++;
        
        emit MaintenanceRecorded(gridId, recordId, operationTypeCode);
    }

    // ─── 查询函数 ────────────────────────────────────────────
    
    function getGrid(address gridId) external view returns (GridInfo memory) {
        require(grids[gridId].exists, "Grid not found");
        return grids[gridId];
    }

    function gridExists(address gridId) external view returns (bool) {
        return grids[gridId].exists;
    }

    function getMaintenanceRecord(
        address gridId, 
        uint32 recordId
    ) external view returns (MaintenanceRecord memory) {
        require(recordId < maintenanceRecordCount[gridId], "Record not found");
        return maintenanceRecords[gridId][recordId];
    }

    /// @notice 获取最近N条维护记录
    function getRecentMaintenanceRecords(
        address gridId,
        uint32 count
    ) external view returns (MaintenanceRecord[] memory) {
        uint32 total = maintenanceRecordCount[gridId];
        uint32 actualCount = count > total ? total : count;
        
        MaintenanceRecord[] memory records = new MaintenanceRecord[](actualCount);
        
        for (uint32 i = 0; i < actualCount; i++) {
            records[i] = maintenanceRecords[gridId][total - actualCount + i];
        }
        
        return records;
    }
}

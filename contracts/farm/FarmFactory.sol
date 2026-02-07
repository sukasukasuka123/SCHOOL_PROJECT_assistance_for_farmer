// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./GridObjectFactory.sol";
import "./DroneObjectFactory.sol";
import "./User.sol";

/// @notice 统一管理农场及其网格的权限与操作入口
contract FarmFactory {

    struct Farm {
        string name;
        address owner;
        uint32 createdTime;
        bool exists;
    }

    // 农场信息
    mapping(address => Farm) public farms;
    mapping(address => mapping(string => address)) public farmGridCodeToId;
    mapping(address => address[]) public farmGridList;

    GridObjectFactory public immutable gridFactory;
    User public immutable userContract;
    DroneObjectFactory public immutable droneFactory;
    
    // FarmRecordManager 合约地址
    address public recordManager;

    // ─── 事件 ────────────────────────────────────────────────
    event FarmCreated(address indexed farmId, string name, address indexed owner);
    event GridAddedToFarm(address indexed farmId, address indexed gridId, string gridCode);
    event RecordManagerSet(address indexed recordManager);

    // ─── 修饰符 ──────────────────────────────────────────────
    modifier onlyFarmer() {
        require(userContract.getUserRole(msg.sender) == User.UserRole.Farmer, "Only farmer allowed");
        _;
    }

    modifier onlyFarmOwner(address farmId) {
        require(farms[farmId].exists, "Farm not found");
        require(farms[farmId].owner == msg.sender, "Not farm owner");
        _;
    }

    modifier onlyAdmin() {
        require(
            userContract.getUserRole(msg.sender) == User.UserRole.Admin,
            "Only admin allowed"
        );
        _;
    }

    modifier onlyRecordManager() {
        require(msg.sender == recordManager, "Only record manager allowed");
        _;
    }

    // ─── 构造函数 ────────────────────────────────────────────
    constructor(address _gridFactory, address _userContract, address _droneFactory) {
        require(_gridFactory != address(0), "Invalid grid factory");
        require(_userContract != address(0), "Invalid user contract");
        require(_droneFactory != address(0), "Invalid drone factory");
        
        gridFactory = GridObjectFactory(_gridFactory);
        userContract = User(_userContract);
        droneFactory = DroneObjectFactory(_droneFactory);
    }

    // ─── RecordManager 设置 ──────────────────────────────────
    
    /// @notice 设置 RecordManager 合约地址（仅部署时调用一次）
    function setRecordManager(address _recordManager) external onlyAdmin {
        require(_recordManager != address(0), "Invalid record manager");
        require(recordManager == address(0), "Record manager already set");
        recordManager = _recordManager;
        emit RecordManagerSet(_recordManager);
    }

    // ─── 农场管理 ────────────────────────────────────────────
    
    function createFarm(address farmId, string calldata name) external onlyFarmer {
        require(farmId != address(0), "Invalid farmId");
        require(bytes(name).length > 0, "Name required");
        require(!farms[farmId].exists, "Farm already exists");

        farms[farmId] = Farm({
            name: name,
            owner: msg.sender,
            createdTime: uint32(block.timestamp),
            exists: true
        });

        emit FarmCreated(farmId, name, msg.sender);
    }

    function addGridToFarm(
        address farmId,
        address gridId,
        string calldata gridCode,
        GridObjectFactory.GridTerrainType terrainType
    ) external onlyFarmOwner(farmId) {
        require(gridId != address(0), "Invalid gridId");
        require(farmGridCodeToId[farmId][gridCode] == address(0), "Grid code already used");

        gridFactory.registerGrid(gridId, gridCode, farmId, terrainType);

        farmGridCodeToId[farmId][gridCode] = gridId;
        farmGridList[farmId].push(gridId);

        emit GridAddedToFarm(farmId, gridId, gridCode);
    }

    function updateGridStatus(
        address farmId,
        string calldata gridCode,
        GridObjectFactory.GridStatus newStatus
    ) external onlyFarmOwner(farmId) {
        address gridId = farmGridCodeToId[farmId][gridCode];
        require(gridId != address(0), "Grid not found");
        gridFactory.updateGridStatus(gridId, newStatus);
    }

    function updateGridDiseaseType(
        address farmId,
        string calldata gridCode,
        GridObjectFactory.CitrusDiseaseType newType
    ) external onlyFarmOwner(farmId) {
        address gridId = farmGridCodeToId[farmId][gridCode];
        require(gridId != address(0), "Grid not found");
        gridFactory.updateDiseaseType(gridId, newType);
    }

    function addGridMaintenance(
        address farmId,
        string calldata gridCode,
        uint8 operationTypeCode,
        string calldata detail,
        bytes32 evidenceIPFSHash
    ) external onlyFarmOwner(farmId) {
        address gridId = farmGridCodeToId[farmId][gridCode];
        require(gridId != address(0), "Grid not found");
        gridFactory.addMaintenanceRecord(gridId, operationTypeCode, detail, evidenceIPFSHash);
    }

    // ─── RecordManager 回调函数 ──────────────────────────────
    
    /// @notice 供 RecordManager 调用以更新网格状态
    function updateGridStatusFromRecord(
        address farmId,
        string calldata gridCode,
        GridObjectFactory.GridStatus newStatus
    ) external onlyRecordManager {
        address gridId = farmGridCodeToId[farmId][gridCode];
        require(gridId != address(0), "Grid not found");
        gridFactory.updateGridStatus(gridId, newStatus);
    }

    /// @notice 供 RecordManager 调用以更新病害类型
    function updateGridDiseaseTypeFromRecord(
        address farmId,
        string calldata gridCode,
        GridObjectFactory.CitrusDiseaseType newType
    ) external onlyRecordManager {
        address gridId = farmGridCodeToId[farmId][gridCode];
        require(gridId != address(0), "Grid not found");
        gridFactory.updateDiseaseType(gridId, newType);
    }

    // ─── 查询函数 ────────────────────────────────────────────
    
    function getFarm(address farmId) external view returns (Farm memory) {
        require(farms[farmId].exists, "Farm not found");
        return farms[farmId];
    }

    function farmExists(address farmId) public view returns (bool) {
        return farms[farmId].exists;
    }

    function isFarmOwner(address farmId, address account) external view returns (bool) {
        return farms[farmId].exists && farms[farmId].owner == account;
    }

    function getGridIdByCode(address farmId, string calldata gridCode) external view returns (address) {
        return farmGridCodeToId[farmId][gridCode];
    }

    function getFarmGrids(address farmId) external view returns (address[] memory) {
        return farmGridList[farmId];
    }

    function getGridInfo(address farmId, string calldata gridCode) 
        external view returns (GridObjectFactory.GridInfo memory) 
    {
        address gridId = farmGridCodeToId[farmId][gridCode];
        require(gridId != address(0), "Grid not found");
        return gridFactory.getGrid(gridId);
    }

    function gridExistsInFarm(address farmId, string calldata gridCode) external view returns (bool) {
        return farmGridCodeToId[farmId][gridCode] != address(0);
    }
}

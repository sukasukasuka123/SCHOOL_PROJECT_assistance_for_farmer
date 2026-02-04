// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./GridObjectFactory.sol";
import "./DroneObjectFactory.sol";

interface IGridObjectFactory {
    function registerGrid(
        address gridId,
        string calldata gridCode,
        address farmId,
        GridObjectFactory.GridTerrainType terrainType
    ) external;

    function updateGridStatus(address gridId, GridObjectFactory.GridStatus newStatus) external;
    function updateDiseaseType(address gridId, GridObjectFactory.CitrusDiseaseType newType) external;
    function addMaintenanceRecord(
        address gridId,
        uint32 recordId,
        string calldata operationType,
        string calldata detail,
        string calldata evidenceIPFS
    ) external;

    function getGrid(address gridId) external view returns (GridObjectFactory.GridInfo memory);
}

/// @title FarmFactory - 农场工厂，统一管理农场及其下属网格的权限与操作入口
contract FarmFactory {

    struct Farm {
        address farmId;
        string name;
        address owner;
        uint256 gridCount;
    }

    // farmId → Farm 信息
    mapping(address => Farm) public farms;

    // farmId → gridCode → gridId（快速反查）
    mapping(address => mapping(string => address)) public farmGridCodeToId;

    // farmId → 所有 gridId 列表（便于前端分页/展示）
    mapping(address => address[]) public farmGridList;

    IGridObjectFactory public immutable gridFactory;

    // ─── 修饰符 ───────────────────────────────────────────────

    modifier onlyFarmOwner(address farmId) {
        require(farms[farmId].owner == msg.sender, "Not farm owner");
        _;
    }

    // ─── 事件 ─────────────────────────────────────────────────

    event FarmCreated(address indexed farmId, string name, address owner);
    event GridAddedToFarm(address indexed farmId, address indexed gridId, string gridCode);

    constructor(address _gridFactory) {
        require(_gridFactory != address(0), "Invalid grid factory");
        gridFactory = IGridObjectFactory(_gridFactory);
    }

    function createFarm(address farmId, string calldata name) external {
        require(farmId != address(0), "Invalid farmId");
        require(bytes(name).length > 0, "Name required");
        require(farms[farmId].farmId == address(0), "Farm already exists");

        farms[farmId] = Farm({
            farmId: farmId,
            name: name,
            owner: msg.sender,
            gridCount: 0
        });

        emit FarmCreated(farmId, name, msg.sender);
    }

    function addGridToFarm(
        address farmId,
        address gridId,
        string calldata gridCode,
        GridObjectFactory.GridTerrainType terrainType
    ) external onlyFarmOwner(farmId) {
        require(farms[farmId].farmId != address(0), "Farm not found");
        require(farmGridCodeToId[farmId][gridCode] == address(0), "Grid code already used");

        gridFactory.registerGrid(gridId, gridCode, farmId, terrainType);

        farmGridCodeToId[farmId][gridCode] = gridId;
        farmGridList[farmId].push(gridId);
        farms[farmId].gridCount++;

        emit GridAddedToFarm(farmId, gridId, gridCode);
    }

    // ─── 代理操作（果农通过 Farm 调用，统一权限控制） ──────

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
        uint32 recordId,
        string calldata operationType,
        string calldata detail,
        string calldata evidenceIPFS
    ) external onlyFarmOwner(farmId) {
        address gridId = farmGridCodeToId[farmId][gridCode];
        require(gridId != address(0), "Grid not found");

        gridFactory.addMaintenanceRecord(gridId, recordId, operationType, detail, evidenceIPFS);
    }

    // ─── 查询 ─────────────────────────────────────────────────

    function getFarm(address farmId) external view returns (Farm memory) {
        return farms[farmId];
    }

    function getGridIdByCode(address farmId, string calldata gridCode) external view returns (address) {
        return farmGridCodeToId[farmId][gridCode];
    }

    function getFarmGrids(address farmId) external view returns (address[] memory) {
        return farmGridList[farmId];
    }

    function getGridInfo(address farmId, string calldata gridCode) external view returns (GridObjectFactory.GridInfo memory) {
        address gridId = farmGridCodeToId[farmId][gridCode];
        require(gridId != address(0), "Grid not found");
        return gridFactory.getGrid(gridId);
    }
//———————————————————————————————————————————————————————————— 
//——————————————业务结构体————————————————————————————————————
//———————————————————————————————————————————————————————————— 
    // 1. 风险区检测记录（侦察机发现木虱高风险区）
    struct RiskDetectionRecord {
        uint32 recordId;                // 记录唯一ID
        uint32 timestamp;               // 检测时间
        string gridCode;                // 网格编号
        bytes32 gpsCenter;              // 网格中心坐标（简化存储为 bytes32）
        uint8 riskLevel;                // 风险等级 0-100
        uint16 ndviValue;               // NDVI × 100
        uint16[3] spectralData;         // 680/730/800 nm 数据
        uint8 psyllidDensity;           // 木虱密度（只/板）
        string evidenceImageIPFS;       // 航拍证据
        address operatorDroneId;        // 执行侦察的无人机地址
    }

    // 2. 标记操作记录（荧光标记 + 隔离带）
    struct MarkingOperationLog {
        uint32 operationId;
        uint32 linkedDetectionId;       // 关联的检测记录ID
        uint32 timestamp;
        string gridCode;
        uint8 markingPointCount;        // 标记点数
        bytes32[] markingGPS;           // 标记点坐标数组
        string markerType;              // 标记剂类型描述
        uint16 markerDosagePerPoint;    // 单点用量 ml
        string isolationPesticide;      // 隔离带农药
        uint16 isolationDosage;         // 隔离带总用量 ml
        address operatorDroneId;        // 执行标记的无人机
        string operationImageIPFS;      // 标记完成照片
    }

    // 3. 病株移除记录
    struct DiseaseTreeRemovalRecord {
        uint32 removalId;
        uint32 linkedMarkingId;         // 关联的标记操作ID
        uint32 timestamp;
        string gridCode;
        uint8 confirmedDiseaseCount;    // 确认病株数
        uint8 falsePositiveCount;       // 误判数
        bytes32[] removalGPS;           // 移除位置
        string removalEvidenceIPFS;     // 照片证据
        address farmerAddress;          // 执行移除的果农地址
        bool replantScheduled;          // 是否已安排补种
    }

    // 4. 补种记录
    struct ReplantingRecord {
        uint32 replantId;
        uint32 linkedRemovalId;         // 关联的移除记录ID
        uint32 timestamp;
        string gridCode;
        uint8 seedlingCount;
        bytes32[] replantGPS;
        string seedlingSource;          // 苗木来源
        string quarantineCertHash;      // 检疫证明哈希
        string seedlingVariety;         // 品种
        string replantImageIPFS;
        address farmerAddress;
    }

    // 存储映射（以 farmId 为维度）

    mapping(address => RiskDetectionRecord[]) public riskDetections;        // farmId => records
    mapping(address => MarkingOperationLog[]) public markingOperations;
    mapping(address => DiseaseTreeRemovalRecord[]) public removals;
    mapping(address => ReplantingRecord[]) public replantings;

    function getRiskDetections(address farmId) external view returns (RiskDetectionRecord[] memory) {
        return riskDetections[farmId];
    }

    function getMarkingOperations(address farmId) external view returns (MarkingOperationLog[] memory) {
        return markingOperations[farmId];
    }

    function getRemovals(address farmId) external view returns (DiseaseTreeRemovalRecord[] memory) {
        return removals[farmId];
    }

    function getReplantings(address farmId) external view returns (ReplantingRecord[] memory) {
        return replantings[farmId];
    }

    event RiskDetected(address indexed farmId, string gridCode, uint32 recordId);
    event GridMarked(address indexed farmId, string gridCode, uint32 operationId);
    event TreesRemoved(address indexed farmId, string gridCode, uint32 removalId);
    event TreesReplanted(address indexed farmId, string gridCode, uint32 replantId);

    // 修改这几个函数的签名，使用 calldata 版本的结构体

function addRiskDetection(
    address farmId,
    RiskDetectionRecord calldata record
) external onlyFarmOwner(farmId) {
    address gridId = farmGridCodeToId[farmId][record.gridCode];
    require(gridId != address(0), "Grid not found");

    // 强制覆盖链上生成字段
    RiskDetectionRecord memory newRecord = record;
    newRecord.timestamp = uint32(block.timestamp);

    riskDetections[farmId].push(newRecord);

    if (record.riskLevel >= 50) {
        gridFactory.updateGridStatus(gridId, GridObjectFactory.GridStatus.PsyllidRisk);
    }

    emit RiskDetected(farmId, record.gridCode, record.recordId);
}

function addMarkingOperation(
    address farmId,
    MarkingOperationLog calldata record
) external onlyFarmOwner(farmId) {
    address gridId = farmGridCodeToId[farmId][record.gridCode];
    require(gridId != address(0), "Grid not found");

    MarkingOperationLog memory newRecord = record;
    newRecord.timestamp = uint32(block.timestamp);

    markingOperations[farmId].push(newRecord);

    gridFactory.updateGridStatus(gridId, GridObjectFactory.GridStatus.MarkedSuspect);

    emit GridMarked(farmId, record.gridCode, record.operationId);
}

function addTreeRemoval(
    address farmId,
    DiseaseTreeRemovalRecord calldata record
) external onlyFarmOwner(farmId) {
    address gridId = farmGridCodeToId[farmId][record.gridCode];
    require(gridId != address(0), "Grid not found");

    DiseaseTreeRemovalRecord memory newRecord = record;
    newRecord.timestamp = uint32(block.timestamp);
    newRecord.farmerAddress = msg.sender;

    removals[farmId].push(newRecord);

    if (record.confirmedDiseaseCount > 0) {
        gridFactory.updateDiseaseType(gridId, GridObjectFactory.CitrusDiseaseType.HuangLongBing);
        gridFactory.updateGridStatus(gridId, GridObjectFactory.GridStatus.Removed);
    }

    emit TreesRemoved(farmId, record.gridCode, record.removalId);
}

function addReplanting(
    address farmId,
    ReplantingRecord calldata record
) external onlyFarmOwner(farmId) {
    address gridId = farmGridCodeToId[farmId][record.gridCode];
    require(gridId != address(0), "Grid not found");

    ReplantingRecord memory newRecord = record;
    newRecord.timestamp = uint32(block.timestamp);
    newRecord.farmerAddress = msg.sender;

    replantings[farmId].push(newRecord);

    gridFactory.updateGridStatus(gridId, GridObjectFactory.GridStatus.Replanted);

    emit TreesReplanted(farmId, record.gridCode, record.replantId);
}
}
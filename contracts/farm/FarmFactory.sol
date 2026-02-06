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

    // 业务记录优化：使用映射 + 计数器
    struct RiskDetectionRecord {
        uint32 timestamp;
        bytes32 gridCodeHash;
        bytes32 gpsCenter;
        uint8 riskLevel;
        uint16 ndviValue;
        uint16[3] spectralData;
        uint8 psyllidDensity;
        bytes32 evidenceIPFSHash;
        address operatorDrone;
    }

    struct MarkingOperationLog {
        uint32 timestamp;
        uint32 linkedDetectionId;
        bytes32 gridCodeHash;
        uint8 markingPointCount;
        uint8 markerTypeCode;
        uint16 markerDosagePerPoint;
        uint8 isolationPesticideCode;
        uint16 isolationDosage;
        address operatorDrone;
        bytes32 operationIPFSHash;
    }

    struct DiseaseTreeRemovalRecord {
        uint32 timestamp;
        uint32 linkedMarkingId;
        bytes32 gridCodeHash;
        uint8 confirmedDiseaseCount;
        uint8 falsePositiveCount;
        bytes32 removalEvidenceIPFSHash;
        address farmer;
        bool replantScheduled;
    }

    struct ReplantingRecord {
        uint32 timestamp;
        uint32 linkedRemovalId;
        bytes32 gridCodeHash;
        uint8 seedlingCount;
        uint8 seedlingVarietyCode;
        bytes32 quarantineCertHash;
        bytes32 replantIPFSHash;
        address farmer;
    }

    // 农场信息
    mapping(address => Farm) public farms;
    mapping(address => mapping(string => address)) public farmGridCodeToId;
    mapping(address => address[]) public farmGridList;

    // 业务记录（使用映射存储）
    mapping(address => mapping(uint32 => RiskDetectionRecord)) public riskDetections;
    mapping(address => uint32) public riskDetectionCount;
    
    mapping(address => mapping(uint32 => MarkingOperationLog)) public markingOperations;
    mapping(address => uint32) public markingOperationCount;
    
    mapping(address => mapping(uint32 => DiseaseTreeRemovalRecord)) public removals;
    mapping(address => uint32) public removalCount;
    
    mapping(address => mapping(uint32 => ReplantingRecord)) public replantings;
    mapping(address => uint32) public replantingCount;

    GridObjectFactory public immutable gridFactory;
    User public immutable userContract;
    DroneObjectFactory public immutable droneFactory;

    // ─── 事件 ────────────────────────────────────────────────
    event FarmCreated(address indexed farmId, string name, address indexed owner);
    event GridAddedToFarm(address indexed farmId, address indexed gridId, string gridCode);
    event RiskDetected(address indexed farmId, uint32 indexed recordId, uint8 riskLevel);
    event GridMarked(address indexed farmId, uint32 indexed operationId);
    event TreesRemoved(address indexed farmId, uint32 indexed removalId);
    event TreesReplanted(address indexed farmId, uint32 indexed replantId);

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

    modifier onlyScoutDrone() {
        require(droneFactory.droneExists(msg.sender), "Not a drone");
        require(droneFactory.getDroneType(msg.sender) == DroneObjectFactory.DroneType.Scout, "Only scout drone allowed");
        _;
    }

    modifier onlyMarkerDrone() {
        require(droneFactory.droneExists(msg.sender), "Not a drone");
        require(droneFactory.getDroneType(msg.sender) == DroneObjectFactory.DroneType.Marker, "Only marker drone allowed");
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

    // ─── 查询函数 ────────────────────────────────────────────
    
    function getFarm(address farmId) external view returns (Farm memory) {
        require(farms[farmId].exists, "Farm not found");
        return farms[farmId];
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

    // ─── 业务函数 ────────────────────────────────────────────

    function addRiskDetection(
        address farmId,
        string calldata gridCode,
        bytes32 gpsCenter,
        uint8 riskLevel,
        uint16 ndviValue,
        uint16[3] calldata spectralData,
        uint8 psyllidDensity,
        bytes32 evidenceIPFSHash
    ) external onlyScoutDrone {
        require(farms[farmId].exists, "Farm not found");
        address gridId = farmGridCodeToId[farmId][gridCode];
        require(gridId != address(0), "Grid not found");

        uint32 recordId = riskDetectionCount[farmId];
        
        riskDetections[farmId][recordId] = RiskDetectionRecord({
            timestamp: uint32(block.timestamp),
            gridCodeHash: keccak256(bytes(gridCode)),
            gpsCenter: gpsCenter,
            riskLevel: riskLevel,
            ndviValue: ndviValue,
            spectralData: spectralData,
            psyllidDensity: psyllidDensity,
            evidenceIPFSHash: evidenceIPFSHash,
            operatorDrone: msg.sender
        });
        
        riskDetectionCount[farmId]++;

        if (riskLevel >= 50) {
            gridFactory.updateGridStatus(gridId, GridObjectFactory.GridStatus.PsyllidRisk);
        }

        emit RiskDetected(farmId, recordId, riskLevel);
    }

    function addMarkingOperation(
        address farmId,
        string calldata gridCode,
        uint32 linkedDetectionId,
        uint8 markingPointCount,
        uint8 markerTypeCode,
        uint16 markerDosagePerPoint,
        uint8 isolationPesticideCode,
        uint16 isolationDosage,
        bytes32 operationIPFSHash
    ) external onlyMarkerDrone {
        require(farms[farmId].exists, "Farm not found");
        address gridId = farmGridCodeToId[farmId][gridCode];
        require(gridId != address(0), "Grid not found");

        uint32 operationId = markingOperationCount[farmId];
        
        markingOperations[farmId][operationId] = MarkingOperationLog({
            timestamp: uint32(block.timestamp),
            linkedDetectionId: linkedDetectionId,
            gridCodeHash: keccak256(bytes(gridCode)),
            markingPointCount: markingPointCount,
            markerTypeCode: markerTypeCode,
            markerDosagePerPoint: markerDosagePerPoint,
            isolationPesticideCode: isolationPesticideCode,
            isolationDosage: isolationDosage,
            operatorDrone: msg.sender,
            operationIPFSHash: operationIPFSHash
        });
        
        markingOperationCount[farmId]++;
        gridFactory.updateGridStatus(gridId, GridObjectFactory.GridStatus.MarkedSuspect);

        emit GridMarked(farmId, operationId);
    }

    function addTreeRemoval(
        address farmId,
        string calldata gridCode,
        uint32 linkedMarkingId,
        uint8 confirmedDiseaseCount,
        uint8 falsePositiveCount,
        bytes32 removalEvidenceIPFSHash,
        bool replantScheduled
    ) external onlyFarmOwner(farmId) {
        address gridId = farmGridCodeToId[farmId][gridCode];
        require(gridId != address(0), "Grid not found");

        uint32 removalId = removalCount[farmId];
        
        removals[farmId][removalId] = DiseaseTreeRemovalRecord({
            timestamp: uint32(block.timestamp),
            linkedMarkingId: linkedMarkingId,
            gridCodeHash: keccak256(bytes(gridCode)),
            confirmedDiseaseCount: confirmedDiseaseCount,
            falsePositiveCount: falsePositiveCount,
            removalEvidenceIPFSHash: removalEvidenceIPFSHash,
            farmer: msg.sender,
            replantScheduled: replantScheduled
        });
        
        removalCount[farmId]++;

        if (confirmedDiseaseCount > 0) {
            gridFactory.updateDiseaseType(gridId, GridObjectFactory.CitrusDiseaseType.HuangLongBing);
            gridFactory.updateGridStatus(gridId, GridObjectFactory.GridStatus.Removed);
        }

        emit TreesRemoved(farmId, removalId);
    }

    function addReplanting(
        address farmId,
        string calldata gridCode,
        uint32 linkedRemovalId,
        uint8 seedlingCount,
        uint8 seedlingVarietyCode,
        bytes32 quarantineCertHash,
        bytes32 replantIPFSHash
    ) external onlyFarmOwner(farmId) {
        address gridId = farmGridCodeToId[farmId][gridCode];
        require(gridId != address(0), "Grid not found");

        uint32 replantId = replantingCount[farmId];
        
        replantings[farmId][replantId] = ReplantingRecord({
            timestamp: uint32(block.timestamp),
            linkedRemovalId: linkedRemovalId,
            gridCodeHash: keccak256(bytes(gridCode)),
            seedlingCount: seedlingCount,
            seedlingVarietyCode: seedlingVarietyCode,
            quarantineCertHash: quarantineCertHash,
            replantIPFSHash: replantIPFSHash,
            farmer: msg.sender
        });
        
        replantingCount[farmId]++;
        gridFactory.updateGridStatus(gridId, GridObjectFactory.GridStatus.Replanted);

        emit TreesReplanted(farmId, replantId);
    }

    // ─── 业务记录查询 ────────────────────────────────────────

    function getRiskDetection(address farmId, uint32 recordId) 
        external view returns (RiskDetectionRecord memory) 
    {
        require(recordId < riskDetectionCount[farmId], "Record not found");
        return riskDetections[farmId][recordId];
    }

    function getRecentRiskDetections(address farmId, uint32 count) 
        external view returns (RiskDetectionRecord[] memory) 
    {
        uint32 total = riskDetectionCount[farmId];
        uint32 actualCount = count > total ? total : count;
        
        RiskDetectionRecord[] memory records = new RiskDetectionRecord[](actualCount);
        for (uint32 i = 0; i < actualCount; i++) {
            records[i] = riskDetections[farmId][total - actualCount + i];
        }
        return records;
    }

    function getMarkingOperation(address farmId, uint32 operationId) 
        external view returns (MarkingOperationLog memory) 
    {
        require(operationId < markingOperationCount[farmId], "Operation not found");
        return markingOperations[farmId][operationId];
    }

    function getRecentMarkingOperations(address farmId, uint32 count) 
        external view returns (MarkingOperationLog[] memory) 
    {
        uint32 total = markingOperationCount[farmId];
        uint32 actualCount = count > total ? total : count;
        
        MarkingOperationLog[] memory records = new MarkingOperationLog[](actualCount);
        for (uint32 i = 0; i < actualCount; i++) {
            records[i] = markingOperations[farmId][total - actualCount + i];
        }
        return records;
    }

    function getRemoval(address farmId, uint32 removalId) 
        external view returns (DiseaseTreeRemovalRecord memory) 
    {
        require(removalId < removalCount[farmId], "Removal not found");
        return removals[farmId][removalId];
    }

    function getRecentRemovals(address farmId, uint32 count) 
        external view returns (DiseaseTreeRemovalRecord[] memory) 
    {
        uint32 total = removalCount[farmId];
        uint32 actualCount = count > total ? total : count;
        
        DiseaseTreeRemovalRecord[] memory records = new DiseaseTreeRemovalRecord[](actualCount);
        for (uint32 i = 0; i < actualCount; i++) {
            records[i] = removals[farmId][total - actualCount + i];
        }
        return records;
    }

    function getReplanting(address farmId, uint32 replantId) 
        external view returns (ReplantingRecord memory) 
    {
        require(replantId < replantingCount[farmId], "Replanting not found");
        return replantings[farmId][replantId];
    }

    function getRecentReplantings(address farmId, uint32 count) 
        external view returns (ReplantingRecord[] memory) 
    {
        uint32 total = replantingCount[farmId];
        uint32 actualCount = count > total ? total : count;
        
        ReplantingRecord[] memory records = new ReplantingRecord[](actualCount);
        for (uint32 i = 0; i < actualCount; i++) {
            records[i] = replantings[farmId][total - actualCount + i];
        }
        return records;
    }

    function getRecordCounts(address farmId) external view returns (
        uint32 risks,
        uint32 markings,
        uint32 removalsCount,
        uint32 replantingsCount
    ) {
        return (
            riskDetectionCount[farmId],
            markingOperationCount[farmId],
            removalCount[farmId],
            replantingCount[farmId]
        );
    }
}

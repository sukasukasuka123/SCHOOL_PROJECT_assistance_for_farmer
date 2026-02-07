// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./GridObjectFactory.sol";
import "./DroneObjectFactory.sol";
import "./User.sol";

/// @notice FarmFactory 的接口，用于回调
interface IFarmFactory {
    function farmExists(address farmId) external view returns (bool);
    function isFarmOwner(address farmId, address account) external view returns (bool);
    function gridExistsInFarm(address farmId, string calldata gridCode) external view returns (bool);
    function updateGridStatusFromRecord(address farmId, string calldata gridCode, GridObjectFactory.GridStatus newStatus) external;
    function updateGridDiseaseTypeFromRecord(address farmId, string calldata gridCode, GridObjectFactory.CitrusDiseaseType newType) external;
}

/// @notice 管理农场的所有业务记录（风险检测、标记、移除、补种）
contract FarmRecordManager {

    // 业务记录结构体
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

    // 业务记录存储
    mapping(address => mapping(uint32 => RiskDetectionRecord)) public riskDetections;
    mapping(address => uint32) public riskDetectionCount;
    
    mapping(address => mapping(uint32 => MarkingOperationLog)) public markingOperations;
    mapping(address => uint32) public markingOperationCount;
    
    mapping(address => mapping(uint32 => DiseaseTreeRemovalRecord)) public removals;
    mapping(address => uint32) public removalCount;
    
    mapping(address => mapping(uint32 => ReplantingRecord)) public replantings;
    mapping(address => uint32) public replantingCount;

    IFarmFactory public immutable farmFactory;
    User public immutable userContract;
    DroneObjectFactory public immutable droneFactory;

    // ─── 事件 ────────────────────────────────────────────────
    event RiskDetected(address indexed farmId, uint32 indexed recordId, uint8 riskLevel);
    event GridMarked(address indexed farmId, uint32 indexed operationId);
    event TreesRemoved(address indexed farmId, uint32 indexed removalId);
    event TreesReplanted(address indexed farmId, uint32 indexed replantId);

    // ─── 修饰符 ──────────────────────────────────────────────
    modifier onlyFarmOwner(address farmId) {
        require(farmFactory.farmExists(farmId), "Farm not found");
        require(farmFactory.isFarmOwner(farmId, msg.sender), "Not farm owner");
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
    constructor(address _farmFactory, address _userContract, address _droneFactory) {
        require(_farmFactory != address(0), "Invalid farm factory");
        require(_userContract != address(0), "Invalid user contract");
        require(_droneFactory != address(0), "Invalid drone factory");
        
        farmFactory = IFarmFactory(_farmFactory);
        userContract = User(_userContract);
        droneFactory = DroneObjectFactory(_droneFactory);
    }

    // ─── 业务函数 ────────────────────────────────────────────

    /// @notice 添加风险检测记录（由侦察无人机调用）
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
        require(farmFactory.farmExists(farmId), "Farm not found");
        require(farmFactory.gridExistsInFarm(farmId, gridCode), "Grid not found");

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

        // 回调 FarmFactory 更新网格状态
        if (riskLevel >= 50) {
            farmFactory.updateGridStatusFromRecord(farmId, gridCode, GridObjectFactory.GridStatus.PsyllidRisk);
        }

        emit RiskDetected(farmId, recordId, riskLevel);
    }

    /// @notice 添加标记操作记录（由标记无人机调用）
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
        require(farmFactory.farmExists(farmId), "Farm not found");
        require(farmFactory.gridExistsInFarm(farmId, gridCode), "Grid not found");

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
        
        // 回调 FarmFactory 更新网格状态
        farmFactory.updateGridStatusFromRecord(farmId, gridCode, GridObjectFactory.GridStatus.MarkedSuspect);

        emit GridMarked(farmId, operationId);
    }

    /// @notice 添加病树移除记录（由农场主调用）
    function addTreeRemoval(
        address farmId,
        string calldata gridCode,
        uint32 linkedMarkingId,
        uint8 confirmedDiseaseCount,
        uint8 falsePositiveCount,
        bytes32 removalEvidenceIPFSHash,
        bool replantScheduled
    ) external onlyFarmOwner(farmId) {
        require(farmFactory.gridExistsInFarm(farmId, gridCode), "Grid not found");

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

        // 回调 FarmFactory 更新网格状态和病害类型
        if (confirmedDiseaseCount > 0) {
            farmFactory.updateGridDiseaseTypeFromRecord(farmId, gridCode, GridObjectFactory.CitrusDiseaseType.HuangLongBing);
            farmFactory.updateGridStatusFromRecord(farmId, gridCode, GridObjectFactory.GridStatus.Removed);
        }

        emit TreesRemoved(farmId, removalId);
    }

    /// @notice 添加补种记录（由农场主调用）
    function addReplanting(
        address farmId,
        string calldata gridCode,
        uint32 linkedRemovalId,
        uint8 seedlingCount,
        uint8 seedlingVarietyCode,
        bytes32 quarantineCertHash,
        bytes32 replantIPFSHash
    ) external onlyFarmOwner(farmId) {
        require(farmFactory.gridExistsInFarm(farmId, gridCode), "Grid not found");

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
        
        // 回调 FarmFactory 更新网格状态
        farmFactory.updateGridStatusFromRecord(farmId, gridCode, GridObjectFactory.GridStatus.Replanted);

        emit TreesReplanted(farmId, replantId);
    }

    // ─── 查询函数 ────────────────────────────────────────────

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

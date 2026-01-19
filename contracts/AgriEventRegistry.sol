// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./DroneIdentity.sol";

contract AgriEventRegistry {
    DroneIdentity public immutable identityRegistry;

    uint32 public eventCounter;

    struct Detection {
        uint32 id;
        bytes32 gps;
        uint8 severity;
        uint16[3] spectral; // 680/730/800nm
        string ipfsHash;
    }

    struct Operation {
        uint32 detectionId;
        string substance;
        uint16 dosage;
        uint64 timestamp;
    }

    mapping(uint32 => Detection) public detections;
    mapping(uint32 => Operation) public operations;

    constructor(address _identityRegistry) {
        identityRegistry = DroneIdentity(_identityRegistry);
    }

    // ── 权限修饰符 ──
    modifier onlyDrone() {
        require(
            identityRegistry.hasRole(msg.sender, identityRegistry.ROLE_DRONE()),
            "Only Drone"
        );
        _;
    }

    modifier onlyFarmer() {
        require(
            identityRegistry.hasRole(msg.sender, identityRegistry.ROLE_FARMER()),
            "Only Farmer"
        );
        _;
    }

    // 无人机记录检测结果
    function recordDetection(
        bytes32 gps,
        uint8 severity,
        uint16[3] memory spectral,
        string calldata ipfsHash
    ) external onlyDrone {
        eventCounter++;
        detections[eventCounter] = Detection({
            id: eventCounter,
            gps: gps,
            severity: severity,
            spectral: spectral,
            ipfsHash: ipfsHash
        });
    }

    // 果农/植保记录实际用药/防治
    function recordOperation(
        uint32 detectionId,
        string calldata substance,
        uint16 dosage
    ) external onlyFarmer {
        require(detections[detectionId].id != 0, "Detection not found");
        require(operations[detectionId].detectionId == 0, "Already operated");

        operations[detectionId] = Operation({
            detectionId: detectionId,
            substance: substance,
            dosage: dosage,
            timestamp: uint64(block.timestamp)
        });
    }

    // 查询接口（公开）
    function getDetection(uint32 id) external view returns (Detection memory) {
        return detections[id];
    }

    function getOperation(uint32 id) external view returns (Operation memory) {
        return operations[id];
    }
}

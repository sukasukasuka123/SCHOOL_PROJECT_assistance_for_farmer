// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./DroneIdentity.sol";
import "./AgriEventRegistry.sol";

contract EmergencyManager {
    AgriEventRegistry public immutable registry;
    DroneIdentity public immutable identityRegistry;

    struct Alert {
        string level;
        bool isActive;
        bool humanHandled;
        string feedback;
        uint64 createTime;
    }

    mapping(uint32 => Alert) public alerts; // key = detectionId

    constructor(address _registry, address _identityRegistry) {
        registry = AgriEventRegistry(_registry);
        identityRegistry = DroneIdentity(_identityRegistry);
    }

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

    // 无人机触发警报
    function triggerAlert(uint32 detectionId, string calldata level) external onlyDrone {
        require(registry.getDetection(detectionId).id != 0, "Detection not exist");
        require(!alerts[detectionId].isActive, "Alert already exists");

        alerts[detectionId] = Alert({
            level: level,
            isActive: true,
            humanHandled: false,
            feedback: "",
            createTime: uint64(block.timestamp)
        });
    }

    // 果农处理/确认警报
    function resolveAlert(uint32 detectionId, string calldata feedback) external onlyFarmer {
        Alert storage alert = alerts[detectionId];
        require(alert.isActive, "No active alert");
        require(!alert.humanHandled, "Already resolved");

        alert.humanHandled = true;
        alert.isActive = false;
        alert.feedback = feedback;
    }

    function getAlert(uint32 detectionId) external view returns (Alert memory) {
        return alerts[detectionId];
    }
}
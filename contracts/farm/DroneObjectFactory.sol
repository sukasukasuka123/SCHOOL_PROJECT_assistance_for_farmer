// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./User.sol";

/// @notice 管理侦察机和标记机的注册、状态和维护记录
contract DroneObjectFactory {
    
    enum DroneType { Scout, Marker, Other }
    enum DroneState { Used, Free, Maintenance }
    
    struct Drone {
        DroneType droneType;
        string firmwareVersion;
        DroneState state;
        uint32 registeredTime;
        bool exists;
    }

    // 无人机信息映射
    mapping(address => Drone) public drones;
    
    // 维护日志优化：使用映射 + 计数器，避免大数组
    struct MaintenanceLog {
        uint32 timestamp;
        uint8 maintenanceTypeCode;  // 编码：1=传感器校准, 2=固件升级, 3=其他
        bool sensorCalibrated;
        string firmwareVersion;
        address technician;
    }
    
    mapping(address => mapping(uint32 => MaintenanceLog)) public maintenanceLogs;
    mapping(address => uint32) public maintenanceLogCount;

    User public immutable userContract;

    // ─── 事件 ────────────────────────────────────────────────
    event DroneRegistered(
        address indexed droneId, 
        DroneType indexed droneType,
        uint32 timestamp
    );
    event DroneStateUpdated(
        address indexed droneId, 
        DroneState indexed newState
    );
    event MaintenanceLogAdded(
        address indexed droneId, 
        uint32 indexed logId,
        uint8 maintenanceType
    );
    event FirmwareUpdated(
        address indexed droneId, 
        string newVersion
    );

    // ─── 修饰符 ──────────────────────────────────────────────
    modifier onlyTechnician() {
        require(
            userContract.getUserRole(msg.sender) == User.UserRole.Technician ||
            userContract.getUserRole(msg.sender) == User.UserRole.Admin,
            "Only technician or admin allowed"
        );
        _;
    }

    // ─── 构造函数 ────────────────────────────────────────────
    constructor(address _userContract) {
        require(_userContract != address(0), "Invalid user contract");
        userContract = User(_userContract);
    }

    // ─── 核心函数 ────────────────────────────────────────────
    
    /// @notice 注册无人机
    function registerDrone(
        address _droneId,
        DroneType _droneType,
        string calldata _firmwareVersion,
        DroneState _droneState
    ) external onlyTechnician {
        require(_droneId != address(0), "Invalid drone ID");
        require(!drones[_droneId].exists, "Drone already registered");
        
        drones[_droneId] = Drone({
            droneType: _droneType,
            firmwareVersion: _firmwareVersion,
            state: _droneState,
            registeredTime: uint32(block.timestamp),
            exists: true
        });

        emit DroneRegistered(_droneId, _droneType, uint32(block.timestamp));
    }

    /// @notice 更新无人机状态
    function updateDroneState(
        address _droneId, 
        DroneState _newState
    ) external onlyTechnician {
        require(drones[_droneId].exists, "Drone not found");
        drones[_droneId].state = _newState;
        emit DroneStateUpdated(_droneId, _newState);
    }

    /// @notice 更新固件版本
    function updateFirmwareVersion(
        address _droneId, 
        string calldata _newFirmwareVersion
    ) external onlyTechnician {
        require(drones[_droneId].exists, "Drone not found");
        drones[_droneId].firmwareVersion = _newFirmwareVersion;
        emit FirmwareUpdated(_droneId, _newFirmwareVersion);
    }

    /// @notice 添加维护日志（优化版）
    function addMaintenanceLog(
        address _droneId,
        uint8 _maintenanceTypeCode,
        bool _sensorCalibrated,
        string calldata _firmwareVersion
    ) external onlyTechnician {
        require(drones[_droneId].exists, "Drone not found");
        require(_maintenanceTypeCode > 0 && _maintenanceTypeCode <= 3, "Invalid type");
        
        uint32 logId = maintenanceLogCount[_droneId];
        
        maintenanceLogs[_droneId][logId] = MaintenanceLog({
            timestamp: uint32(block.timestamp),
            maintenanceTypeCode: _maintenanceTypeCode,
            sensorCalibrated: _sensorCalibrated,
            firmwareVersion: _firmwareVersion,
            technician: msg.sender
        });
        
        maintenanceLogCount[_droneId]++;
        
        emit MaintenanceLogAdded(_droneId, logId, _maintenanceTypeCode);
    }

    // ─── 查询函数 ────────────────────────────────────────────
    
    function getDrone(address _droneId) external view returns (Drone memory) {
        require(drones[_droneId].exists, "Drone not found");
        return drones[_droneId];
    }

    function getDroneType(address _droneId) external view returns (DroneType) {
        require(drones[_droneId].exists, "Drone not found");
        return drones[_droneId].droneType;
    }

    function droneExists(address _droneId) external view returns (bool) {
        return drones[_droneId].exists;
    }

    /// @notice 获取维护日志（按ID）
    function getMaintenanceLog(
        address _droneId, 
        uint32 _logId
    ) external view returns (MaintenanceLog memory) {
        require(_logId < maintenanceLogCount[_droneId], "Log not found");
        return maintenanceLogs[_droneId][_logId];
    }

    /// @notice 获取最近N条维护日志（避免返回全部）
    function getRecentMaintenanceLogs(
        address _droneId,
        uint32 _count
    ) external view returns (MaintenanceLog[] memory) {
        uint32 total = maintenanceLogCount[_droneId];
        uint32 count = _count > total ? total : _count;
        
        MaintenanceLog[] memory logs = new MaintenanceLog[](count);
        
        for (uint32 i = 0; i < count; i++) {
            logs[i] = maintenanceLogs[_droneId][total - count + i];
        }
        
        return logs;
    }
}

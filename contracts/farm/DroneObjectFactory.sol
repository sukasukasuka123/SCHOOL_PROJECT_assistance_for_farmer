// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract DroneObjectFactory {
    enum DroneType { Scout, Marker, Other }
    enum state  { Used, Free, Maintenance } // Used - 使用中, Free - 空闲, Maintenance - 维护中
    struct Drone{
        address droneId;               // 无人机设备地址(即唯一的无人机ID)  
        DroneType droneTypeValue;              // 设备类型："侦察机"/"标记机"
        string firmwareVersion;        // 固件版本号
        state droneState;              // 无人机状态
    }

    mapping(address => Drone) public drones;

    function registerDrone(
        address _droneId,
        DroneType _droneType,
        string calldata _firmwareVersion,
        state _droneState
    ) external {
        require(_droneId != address(0), "Invalid drone ID");
        require(drones[_droneId].droneId == address(0), "Drone already registered");
        drones[_droneId] = Drone({
            droneId: _droneId,
            droneTypeValue: DroneType(_droneType),
            firmwareVersion: _firmwareVersion,
            droneState: _droneState
        });
    }

    function getDrone(address _droneId) external view returns (Drone memory) {
        return drones[_droneId];
    }

    function updateDroneState(address _droneId, state _newState) external {
        drones[_droneId].droneState = _newState;
    }
    // ────────────────────────────────────────────────
    //                  维护日志部分
    // ────────────────────────────────────────────────

    struct DroneMaintenanceLog {    
        uint32 logId;                  // 日志ID    
        address droneId;               // 无人机设备地址    
        uint32 timestamp;              // 维护时间    
        string droneType;              // 设备类型："侦察机"/"标记机"    
        string maintenanceType;        // 维护类型："传感器校准"/"固件升级"    
        bool sensorCalibrated;         // 传感器校准状态    
        string firmwareVersion;        // 固件版本号    
        address technicianAddress;     // 技术员地址    
    }    
    mapping(address => DroneMaintenanceLog[]) public droneMaintenanceLogs;

    function addDroneMaintenanceLog(
        address _droneId,
        uint32 _logId,
        uint32 _timestamp,
        string calldata _droneType,
        string calldata _maintenanceType,
        bool _sensorCalibrated,
        string calldata _firmwareVersion,
        address _technicianAddress
    ) external {
        droneMaintenanceLogs[_droneId].push(DroneMaintenanceLog({
            logId: _logId,
            droneId: _droneId,
            timestamp: _timestamp,
            droneType: _droneType,
            maintenanceType: _maintenanceType,
            sensorCalibrated: _sensorCalibrated,
            firmwareVersion: _firmwareVersion,
            technicianAddress: _technicianAddress
        }));
    }
    function getDroneMaintenanceLogs(address _droneId) external view returns (DroneMaintenanceLog[] memory) {
        return droneMaintenanceLogs[_droneId];
    }
    function getDroneMaintenanceLogCount(address _droneId) external view returns (uint256) {
        return droneMaintenanceLogs[_droneId].length;
    }
    function getDroneMaintenanceLogByIndex(address _droneId, uint256 index) external view returns (DroneMaintenanceLog memory) {
        require(index < droneMaintenanceLogs[_droneId].length, "Index out of bounds");
        return droneMaintenanceLogs[_droneId][index];
    }
    function getDronMaintenanceLogById(address _droneId, uint32 _logId) external view returns (DroneMaintenanceLog memory) {
        DroneMaintenanceLog[] memory logs = droneMaintenanceLogs[_droneId];
        for (uint256 i = 0; i < logs.length; i++) {
            if (logs[i].logId == _logId) {
                return logs[i];
            }
        }
        revert("Log ID not found");
    }
    function getDronMaintenanceLogInPage(address _droneId, uint256 page, uint256 pageSize) external view returns (DroneMaintenanceLog[] memory) {
        DroneMaintenanceLog[] memory logs = droneMaintenanceLogs[_droneId];
        uint256 start = page * pageSize;
        uint256 end = start + pageSize;
        if (end > logs.length) {
            end = logs.length;
        }
        require(start < end, "Invalid page or page size");

        DroneMaintenanceLog[] memory pagedLogs = new DroneMaintenanceLog[](end - start);
        for (uint256 i = start; i < end; i++) {
            pagedLogs[i - start] = logs[i];
        }
        return pagedLogs;
    }
    function updateFirmwareVersion(address _droneId, string calldata _newFirmwareVersion) external {
        drones[_droneId].firmwareVersion = _newFirmwareVersion;
    }
    function updateDroneType(address _droneId, DroneType _newDroneType) external {
        drones[_droneId].droneTypeValue = _newDroneType;
    }
}
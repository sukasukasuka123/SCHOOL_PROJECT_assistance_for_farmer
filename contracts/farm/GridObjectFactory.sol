// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @title GridObjectFactory - 网格对象工厂，管理果园 10m×10m 网格的基本信息与防控状态
contract GridObjectFactory {

    // 网格防控状态（根据产品故事全流程设计）
    enum GridStatus {
        Normal,             // 正常，无风险
        PsyllidRisk,        // 木虱风险（侦察发现高密度，但未确认病株）
        MarkedSuspect,      // 已荧光标记，待人工确认
        ConfirmedInfected,  // 人工确认黄龙病阳性（待移除）
        Removed,            // 病株已移除（等待补种或已补种）
        Replanted,          // 已补种健康苗
        Quarantine,         // 区域隔离观察（片状爆发后）
        Unknown             // 数据缺失或未巡检
    }

    enum GridTerrainType {
        Valley,     // 山谷
        Flat,       // 平地
        Ridge,      // 山脊
        Cliff       // 悬崖/陡坡
    }

    // 黄龙病及其他主要病害类型（可扩展）
    enum CitrusDiseaseType {
        None,
        HuangLongBing,      // 黄龙病
        Anthracnose,        // 炭疽病
        Canker,             // 溃疡病
        Other
    }

    struct GridInfo {
        address gridId;             // 网格唯一标识（建议用 CREATE2 或专用地址）
        string gridCode;            // 网格编码，如 "B3"、"C1"
        address farmId;             // 所属农场
        GridTerrainType terrainType;
        GridStatus status;      // 当前防控状态
        CitrusDiseaseType diseaseType; // 当前主要病害（通常为 HuangLongBing）
        uint32 lastStatusUpdateTime;   // 状态最后变更时间戳
    }

    // 网格信息表
    mapping(address => GridInfo) public grids;

    // 网格维护/养护日志（生物防治、施肥、修剪等常规操作）
    struct GridMaintenanceRecord {
        uint32 recordId;            // 日志唯一ID（可由调用者生成或自增）
        uint32 timestamp;
        string operationType;       // 如：生物防治、人工施肥、修剪、诱捕带设置
        string detail;              // 详细描述，如 "释放捕食螨 500 只"
        string evidenceIPFS;        // 现场照片/视频 IPFS CID
        address operator;           // 操作人（果农或技术员）
    }

    mapping(address => GridMaintenanceRecord[]) public maintenanceRecords;

    // ─── 事件 ────────────────────────────────────────────────

    event FarmGridRegistered(address indexed farmId, address indexed gridId, string gridCode);
    event GridStatusUpdated(address indexed gridId, GridStatus oldStatus, GridStatus newStatus);
    event GridDiseaseTypeUpdated(address indexed gridId, CitrusDiseaseType oldType, CitrusDiseaseType newType);
    event GridMaintenanceRecorded(address indexed gridId, uint32 recordId, string operationType, address operator);

    // ─── 核心修改函数 ───────────────────────────────────────

    /// @notice 注册一个新的网格（通常由 FarmFactory 调用）
    function registerGrid(
        address gridId,
        string calldata gridCode,
        address farmId,
        GridTerrainType terrainType
    ) external {
        require(gridId != address(0), "Invalid gridId");
        require(bytes(gridCode).length > 0, "Empty gridCode");
        require(farmId != address(0), "Invalid farmId");
        require(grids[gridId].gridId == address(0), "Grid already registered");

        grids[gridId] = GridInfo({
            gridId: gridId,
            gridCode: gridCode,
            farmId: farmId,
            terrainType: terrainType,
            status: GridStatus.Normal,
            diseaseType: CitrusDiseaseType.None,
            lastStatusUpdateTime: uint32(block.timestamp)
        });

        emit FarmGridRegistered(farmId, gridId, gridCode);
    }

    /// @notice 更新网格防控状态（侦察、标记、移除、补种等关键节点调用）
    function updateGridStatus(
        address gridId,
        GridStatus newStatus
    ) external {
        GridInfo storage info = grids[gridId];
        require(info.gridId != address(0), "Grid not found");

        GridStatus old = info.status;
        info.status = newStatus;
        info.lastStatusUpdateTime = uint32(block.timestamp);

        emit GridStatusUpdated(gridId, old, newStatus);
    }

    /// @notice 更新病害类型（通常在人工确认后调用）
    function updateDiseaseType(
        address gridId,
        CitrusDiseaseType newType
    ) external {
        GridInfo storage info = grids[gridId];
        require(info.gridId != address(0), "Grid not found");

        CitrusDiseaseType old = info.diseaseType;
        info.diseaseType = newType;
        info.lastStatusUpdateTime = uint32(block.timestamp);

        emit GridDiseaseTypeUpdated(gridId, old, newType);
    }

    /// @notice 添加日常养护记录（生物防治、施肥等）
    function addMaintenanceRecord(
        address gridId,
        uint32 recordId,
        string calldata operationType,
        string calldata detail,
        string calldata evidenceIPFS
    ) external {
        require(grids[gridId].gridId != address(0), "Grid not found");
        require(bytes(operationType).length > 0, "Operation type required");

        maintenanceRecords[gridId].push(GridMaintenanceRecord({
            recordId: recordId,
            timestamp: uint32(block.timestamp),
            operationType: operationType,
            detail: detail,
            evidenceIPFS: evidenceIPFS,
            operator: msg.sender
        }));

        emit GridMaintenanceRecorded(gridId, recordId, operationType, msg.sender);
    }

    // ─── 查询函数 ─────────────────────────────────────────────

    function getGrid(address gridId) external view returns (GridInfo memory) {
        require(grids[gridId].gridId != address(0), "Grid not found");
        return grids[gridId];
    }

    function getMaintenanceRecords(address gridId) external view returns (GridMaintenanceRecord[] memory) {
        return maintenanceRecords[gridId];
    }

    function getMaintenanceRecordCount(address gridId) external view returns (uint256) {
        return maintenanceRecords[gridId].length;
    }

    function getMaintenanceRecordByIndex(address gridId, uint256 index) external view returns (GridMaintenanceRecord memory) {
        require(index < maintenanceRecords[gridId].length, "Index out of bounds");
        return maintenanceRecords[gridId][index];
    }
}
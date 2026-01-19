// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./DroneIdentity.sol";
import "./AgriEventRegistry.sol";

contract ProductTraceability {
    AgriEventRegistry public immutable registry;
    DroneIdentity public immutable identityRegistry;

    struct Batch {
        string zone;
        uint32[] detectionIds;     // 建议关联检测记录
        uint16 biologicalCount;
        uint16 chemicalCount;
        uint64 createTime;
    }

    mapping(uint32 => Batch) public batches; // key = batchId (业务系统生成)

    constructor(address _registry, address _identityRegistry) {
        registry = AgriEventRegistry(_registry);
        identityRegistry = DroneIdentity(_identityRegistry);
    }

    modifier onlyFarmer() {
        require(
            identityRegistry.hasRole(msg.sender, identityRegistry.ROLE_FARMER()),
            "Only Farmer"
        );
        _;
    }

    // 创建采摘批次（通常在采摘/包装阶段由果农/合作社操作）
    function createHarvestBatch(
        uint32 batchId,
        string calldata zone,
        uint32[] calldata detectionIds,
        uint16 biologicalCount,
        uint16 chemicalCount
    ) external onlyFarmer {
        require(batches[batchId].createTime == 0, "Batch already exists");

        // 可选：基础校验
        for (uint256 i = 0; i < detectionIds.length; i++) {
            require(
                registry.getDetection(detectionIds[i]).id != 0,
                "Invalid detection id"
            );
        }

        batches[batchId] = Batch({
            zone: zone,
            detectionIds: detectionIds,
            biologicalCount: biologicalCount,
            chemicalCount: chemicalCount,
            createTime: uint64(block.timestamp)
        });
    }

    // 消费者/公开查询接口
    function getTraceabilityBasic(uint32 batchId)
        external
        view
        returns (
            string memory zone,
            uint16 bioCount,
            uint16 chemCount
        )
    {
        Batch memory b = batches[batchId];
        require(b.createTime != 0, "Batch not found");
        return (b.zone, b.biologicalCount, b.chemicalCount);
    }

    // 如需更详细数据，可调用此接口（视前端需求）
    function getBatchDetail(uint32 batchId) external view returns (Batch memory) {
        return batches[batchId];
    }
}

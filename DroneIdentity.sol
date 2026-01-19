// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title DroneIdentity
 * @notice 基于角色的访问控制合约，管理无人机和果农身份
 */
contract DroneIdentity {
    // ───────────────────────────────────────────────
    // 角色常量
    // ───────────────────────────────────────────────
    bytes32 public constant ROLE_ADMIN  = keccak256("ROLE_ADMIN");
    bytes32 public constant ROLE_DRONE  = keccak256("ROLE_DRONE");
    bytes32 public constant ROLE_FARMER = keccak256("ROLE_FARMER");

    // ───────────────────────────────────────────────
    // 存储
    // ───────────────────────────────────────────────
    mapping(bytes32 => mapping(address => bool)) private _roles;
    
    // 维护管理员地址集合，用于快速计数
    address[] private _adminList;
    mapping(address => uint256) private _adminIndex; // 地址 => 数组索引+1 (0表示不存在)

    // ───────────────────────────────────────────────
    // 事件
    // ───────────────────────────────────────────────
    event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender);
    event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender);

    // ───────────────────────────────────────────────
    // 构造函数
    // ───────────────────────────────────────────────
    constructor() {
        _grantRole(ROLE_ADMIN, msg.sender);
    }

    // ───────────────────────────────────────────────
    // 修饰符
    // ───────────────────────────────────────────────
    modifier onlyAdmin() {
        require(_hasRole(msg.sender, ROLE_ADMIN), "Caller is not admin");
        _;
    }

    // ───────────────────────────────────────────────
    // 核心查询接口（保持兼容）
    // ───────────────────────────────────────────────
    function hasRole(address account, bytes32 role) external view returns (bool) {
        return _hasRole(account, role);
    }

    function isDrone(address account) external view returns (bool) {
        return _hasRole(account, ROLE_DRONE);
    }

    function isFarmer(address account) external view returns (bool) {
        return _hasRole(account, ROLE_FARMER);
    }

    function isAdmin(address account) external view returns (bool) {
        return _hasRole(account, ROLE_ADMIN);
    }

    // ───────────────────────────────────────────────
    // 角色管理函数
    // ───────────────────────────────────────────────
    /**
     * @notice 授予角色（单次）
     */
    function grantRole(address account, bytes32 role) external onlyAdmin {
        require(account != address(0), "Invalid account address");
        _grantRole(role, account);
    }

    /**
     * @notice 撤销角色（单次）
     */
    function revokeRole(address account, bytes32 role) external onlyAdmin {
        require(account != address(0), "Invalid account address");

        // 保护：不能撤销最后一个管理员
        if (role == ROLE_ADMIN) {
            require(_countActiveAdmins() > 1, "Cannot revoke the last admin");
        }

        _revokeRole(role, account);
    }

    /**
     * @notice 用户主动放弃自己的角色
     */
    function renounceRole(bytes32 role) external {
        require(_hasRole(msg.sender, role), "Caller does not have this role");

        // 保护：管理员不能通过 renounce 方式放弃最后权限
        if (role == ROLE_ADMIN) {
            require(_countActiveAdmins() > 1, "Cannot renounce the last admin");
        }

        _revokeRole(role, msg.sender);
    }

    /**
     * @notice 批量授予同一角色
     */
    function grantRoleBatch(address[] calldata accounts, bytes32 role) external onlyAdmin {
        uint256 len = accounts.length;
        for (uint256 i = 0; i < len; ) {
            address account = accounts[i];
            require(account != address(0), "Invalid address in batch");
            _grantRole(role, account);
            unchecked { ++i; }
        }
    }

    /**
     * @notice 获取所有活跃管理员地址
     */
    function getAdmins() external view returns (address[] memory) {
        uint256 activeCount = _countActiveAdmins();
        address[] memory admins = new address[](activeCount);
        uint256 index = 0;
        
        uint256 len = _adminList.length;
        for (uint256 i = 0; i < len; ) {
            address admin = _adminList[i];
            if (_hasRole(admin, ROLE_ADMIN)) {
                admins[index] = admin;
                unchecked { ++index; }
            }
            unchecked { ++i; }
        }
        
        return admins;
    }

    // ───────────────────────────────────────────────
    // 内部函数
    // ───────────────────────────────────────────────
    function _hasRole(address account, bytes32 role) internal view returns (bool) {
        return _roles[role][account];
    }

    function _grantRole(bytes32 role, address account) internal {
        if (!_roles[role][account]) {
            _roles[role][account] = true;
            emit RoleGranted(role, account, msg.sender);
            
            // 如果是管理员角色，添加到管理员列表
            if (role == ROLE_ADMIN && _adminIndex[account] == 0) {
                _adminList.push(account);
                _adminIndex[account] = _adminList.length; // 存储索引+1
            }
        }
    }

    function _revokeRole(bytes32 role, address account) internal {
        if (_roles[role][account]) {
            _roles[role][account] = false;
            emit RoleRevoked(role, account, msg.sender);
            
            // 注意：我们不从 _adminList 中删除地址以节省 gas
            // _adminIndex 保留，getAdmins() 会过滤非活跃管理员
        }
    }

    /**
     * @dev 计算当前活跃管理员数量
     * @notice O(n) 操作，n 为曾经被授予过管理员的地址数量
     */
    function _countActiveAdmins() internal view returns (uint256 count) {
        uint256 len = _adminList.length;
        for (uint256 i = 0; i < len; ) {
            if (_hasRole(_adminList[i], ROLE_ADMIN)) {
                unchecked { ++count; }
            }
            unchecked { ++i; }
        }
    }
}
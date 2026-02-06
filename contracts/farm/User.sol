// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @notice 管理系统中的用户角色和权限
contract User {
    enum UserRole {
        Admin,      // 管理员
        Technician, // 技术人员
        Farmer,     // 果农
        Guest       // 访客
    }

    struct UserInfo {
        string name;            // 用户姓名
        UserRole role;          // 用户角色
        string contactInfo;     // 联系方式
        uint32 registeredTime;  // 注册时间戳
        bool exists;            // 是否存在（避免使用 address(0) 判断）
    }

    mapping(address => UserInfo) public users;
    
    // 管理员列表（用于多管理员场景）
    mapping(address => bool) public isAdmin;
    uint256 public adminCount;

    // ─── 事件 ────────────────────────────────────────────────
    event UserRegistered(
        address indexed userId, 
        string name, 
        UserRole indexed role, 
        uint32 timestamp
    );
    event UserRoleUpdated(
        address indexed userId, 
        UserRole indexed newRole, 
        uint32 timestamp
    );
    event AdminAdded(address indexed newAdmin);
    event AdminRemoved(address indexed removedAdmin);

    // ─── 修饰符 ──────────────────────────────────────────────
    modifier onlyAdmin() {
        require(isAdmin[msg.sender], "Only admin allowed");
        _;
    }

    // ─── 构造函数 ────────────────────────────────────────────
    constructor() {
        // 部署者自动成为管理员
        isAdmin[msg.sender] = true;
        adminCount = 1;
        users[msg.sender] = UserInfo({
            name: "System Admin",
            role: UserRole.Admin,
            contactInfo: "",
            registeredTime: uint32(block.timestamp),
            exists: true
        });
    }

    // ─── 核心函数 ────────────────────────────────────────────
    
    /// @notice 注册新用户
    function registerUser(
        address _userId, 
        string calldata _name, 
        UserRole _role, 
        string calldata _contactInfo
    ) external onlyAdmin {
        require(_userId != address(0), "Invalid address");
        require(!users[_userId].exists, "User already registered");
        
        users[_userId] = UserInfo({
            name: _name,
            role: _role,
            contactInfo: _contactInfo,
            registeredTime: uint32(block.timestamp),
            exists: true
        });

        if (_role == UserRole.Admin) {
            isAdmin[_userId] = true;
            adminCount++;
        }

        emit UserRegistered(_userId, _name, _role, uint32(block.timestamp));
    }

    /// @notice 更新用户角色
    function updateUserRole(
        address _userId, 
        UserRole _newRole
    ) external onlyAdmin {
        require(users[_userId].exists, "User not registered");
        
        UserRole oldRole = users[_userId].role;
        users[_userId].role = _newRole;

        // 处理管理员权限变更
        if (oldRole == UserRole.Admin && _newRole != UserRole.Admin) {
            require(adminCount > 1, "Cannot remove last admin");
            isAdmin[_userId] = false;
            adminCount--;
        } else if (oldRole != UserRole.Admin && _newRole == UserRole.Admin) {
            isAdmin[_userId] = true;
            adminCount++;
        }

        emit UserRoleUpdated(_userId, _newRole, uint32(block.timestamp));
    }

    // ─── 查询函数 ────────────────────────────────────────────
    
    function getUserRole(address _userId) external view returns (UserRole) {
        require(users[_userId].exists, "User not registered");
        return users[_userId].role;
    }

    function getUserInfo(address _userId) external view returns (UserInfo memory) {
        require(users[_userId].exists, "User not registered");
        return users[_userId];
    }

    function userExists(address _userId) external view returns (bool) {
        return users[_userId].exists;
    }
}
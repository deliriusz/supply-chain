// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title Authorization
 * @author Rafał Kalinowski
 * @notice This is a helper contract used for authorization and verification of elevated permissions roles
 * used for Firmex management
 * @dev We could as well use openzeppelin TimelockController, but it would be simply boring.
 * hence it's really simple role based auth check
 */
contract Authorization is Ownable {
    mapping(address => bytes4) public _accountRoleAssigned;
    mapping(address => uint256) public _accoutRoleAssignmentTime;

    uint256 public immutable ROLE_ASSIGNMENT_DELAY_SECS;
    bytes4 public constant ROLE_ADMIN = bytes4(keccak256("ROLE_ADMIN"));
    bytes4 public constant ROLE_DASHBOARD_VIEWER =
        bytes4(keccak256("ROLE_DASHBOARD_VIEWER"));

    modifier hasRole(address _who, bytes4 _role) {
        require(
            _accountRoleAssigned[_who] == _role,
            "Authorization: the address does not have a required role."
        );
        _;
    }

    modifier isRoleActive(address _who) {
        require(
            block.timestamp - _accoutRoleAssignmentTime[_who] >
                ROLE_ASSIGNMENT_DELAY_SECS,
            "Authorization: the address does not have an active role assigned yet."
        );
        _;
    }

    event RoleAssigned(address indexed to, bytes4 role);
    event RoleRevoked(address indexed who);

    constructor(uint256 roleAssignmentDelaySecs) {
        ROLE_ASSIGNMENT_DELAY_SECS = roleAssignmentDelaySecs;
        _assignAdminRole();
    }

    function assignRole(address _to, bytes4 _role)
        external
        hasRole(msg.sender, ROLE_ADMIN)
        isRoleActive(msg.sender)
    {
        if (_accountRoleAssigned[_to] != _role) {
            _accoutRoleAssignmentTime[_to] =
                block.timestamp +
                ROLE_ASSIGNMENT_DELAY_SECS;
            _accountRoleAssigned[_to] = _role;

            emit RoleAssigned(_to, _role);
        }
    }

    function _assignAdminRole() private onlyOwner {
        _accoutRoleAssignmentTime[msg.sender] = block.timestamp;
        _accountRoleAssigned[msg.sender] = ROLE_ADMIN;

        emit RoleAssigned(msg.sender, ROLE_ADMIN);
    }

    function revokeRole(address _to)
        external
        hasRole(msg.sender, ROLE_ADMIN)
        isRoleActive(msg.sender)
    {
        _accountRoleAssigned[_to] = bytes4(0);

        emit RoleRevoked(_to);
    }

    function getUserRole(address _who) public view returns (bytes4) {
        if (
            block.timestamp - _accoutRoleAssignmentTime[_who] <
            ROLE_ASSIGNMENT_DELAY_SECS
        ) {
            return bytes4(0);
        }
        return _accountRoleAssigned[_who];
    }
}

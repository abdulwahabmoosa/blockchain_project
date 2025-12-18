/*
ApprovalService is a simple admin-controlled whitelist contract. 
It lets the admin mark which wallet addresses are allowed to participate in the system (for example, who can buy or receive property tokens). 
The admin (assigned in the constructor) can approve or revoke any user, and the contract stores that in a mapping(address → bool). 
Other contracts (like your ERC20 PropertyToken) call check(address) to verify whether a user is approved before allowing transfers. 
Events are emitted whenever a user is approved or revoked, making the system transparent and easy to track. Essentially, 
this contract acts as the access gatekeeper for your entire platform.
*/
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract ApprovalService is AccessControl {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    mapping(address => bool) public isApproved;

    event Approved(address indexed user, address indexed admin);
    event Revoked(address indexed user, address indexed admin);
    // Emitted when a user is explicitly rejected (on-chain mirror of backend rejection state)
    event UserRejected(address indexed user);

    constructor(address initialAdmin) {
        // Grant admin role to the specified wallet
        _grantRole(ADMIN_ROLE, initialAdmin);
        _grantRole(DEFAULT_ADMIN_ROLE, initialAdmin); // for role management
    }

    modifier onlyAdmin() {
        require(hasRole(ADMIN_ROLE, msg.sender), "AS: not admin");
        _;
    }

    function approve(address user) external onlyAdmin {
        isApproved[user] = true;
        emit Approved(user, msg.sender);
    }

    function revoke(address user) external onlyAdmin {
        isApproved[user] = false;
        emit Revoked(user, msg.sender);
    }

    /**
     * @dev Explicitly mark a user as rejected on-chain.
     * This is a thin wrapper around {revoke} so existing logic remains unchanged,
     * while providing a dedicated event for off-chain listeners.
     */
    function rejectUser(address user) external onlyAdmin {
        isApproved[user] = false;
        emit Revoked(user, msg.sender);
        emit UserRejected(user);
    }

    // Optional: add or remove admins
    function grantAdmin(address user) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _grantRole(ADMIN_ROLE, user);
    }

    function revokeAdmin(address user) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _revokeRole(ADMIN_ROLE, user);
    }

    function check(address user) external view returns (bool) {
        return isApproved[user];
    }
}

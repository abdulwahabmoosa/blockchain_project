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
    bytes32 public constant ADMIN_ROLE = DEFAULT_ADMIN_ROLE;
    mapping(address => bool) public isApproved;

    event Approved(address indexed user);
    event Revoked(address indexed user);

    constructor(address admin) {
        _setupRole(ADMIN_ROLE, admin);
    }

    modifier onlyAdmin() {
        require(hasRole(ADMIN_ROLE, msg.sender), "AS: not admin");
        _;
    }

    function approve(address user) external onlyAdmin {
        isApproved[user] = true;
        emit Approved(user);
    }

    function revoke(address user) external onlyAdmin {
        isApproved[user] = false;
        emit Revoked(user);
    }

    function check(address user) external view returns (bool) {
        return isApproved[user];
    }
}

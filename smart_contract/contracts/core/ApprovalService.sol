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

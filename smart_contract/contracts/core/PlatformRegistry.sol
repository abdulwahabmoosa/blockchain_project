/*
PlatformRegistry is the control center for your entire system. 
It stores the official addresses of the three core contracts: the PropertyFactory, ApprovalService, 
and RevenueDistribution. Only an admin wallet (defined at deployment) is allowed to 
set or update these addresses. This ensures that every part of 
the platform always knows which contracts to interact with and prevents unauthorized 
changes. When the admin updates a contract address, an 
event is emitted so the frontend/backend can track the change. 
Other contracts simply read from the registry using getter 
functions to know where the core services are located. 
In short, the registry acts as the platform’s “address book + admin panel,” keeping everything coordinated and secure.*/






// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract PlatformRegistry is AccessControl {
    bytes32 public constant ADMIN_ROLE = DEFAULT_ADMIN_ROLE;

    address public propertyFactory;
    address public approvalService;
    address public revenueDistribution;

    event PropertyFactorySet(address indexed factory);
    event ApprovalServiceSet(address indexed approval);
    event RevenueDistributionSet(address indexed revenue);

    constructor(address admin) {
        _setupRole(ADMIN_ROLE, admin);
    }

    modifier onlyAdmin() {
        require(hasRole(ADMIN_ROLE, msg.sender), "PR: not admin");
        _;
    }

    function setPropertyFactory(address _factory) external onlyAdmin {
        propertyFactory = _factory;
        emit PropertyFactorySet(_factory);
    }

    function setApprovalService(address _approval) external onlyAdmin {
        approvalService = _approval;
        emit ApprovalServiceSet(_approval);
    }

    function setRevenueDistribution(address _revenue) external onlyAdmin {
        revenueDistribution = _revenue;
        emit RevenueDistributionSet(_revenue);
    }

    function getPropertyFactory() external view returns (address) {
        return propertyFactory;
    }

    function getApprovalService() external view returns (address) {
        return approvalService;
    }

    function getRevenueDistribution() external view returns (address) {
        return revenueDistribution;
    }
}

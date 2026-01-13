// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IApprovalService {
    function isApproved(address user) external view returns (bool);
}

interface IPropertyAsset {
    enum Status { Active, Paused, Disputed, Closed }

    function getStatus() external view returns (Status);
    function propertyToken() external view returns (address);
    function assetId() external view returns (uint256);
    function propertyDataHash() external view returns (string memory);
}

interface IPropertyToken {
    function snapshotNow() external returns (uint256);
    function balanceOfAt(address user, uint256 snapshotId) external view returns (uint256);
    function totalSupplyAt(uint256 snapshotId) external view returns (uint256);
}


interface IRevenueDistribution {
    function depositRevenue(address token, address stablecoin, uint256 amount) external;
    function claimRevenue(uint256 distributionId) external;
}


interface IPlatformRegistry {
    function getPropertyFactory() external view returns (address);
    function getApprovalService() external view returns (address);
    function getRevenueDistribution() external view returns (address);
}
// Usage: construct interface instances and call functions from within a contract or function scope; remove top-level statements.

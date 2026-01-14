/*
The PropertyFactory is responsible for creating a complete tokenized property on the blockchain.
When the admin (who has the CREATOR_ROLE) calls createProperty,
the factory automatically generates two linked smart contracts:
a PropertyAsset (ERC-721 NFT representing the property itself)
and a PropertyToken (ERC-20 token representing the fractional ownership shares).
It gives the property a unique ID, stores the metadata hash, and sets the property's
valuation. It then deploys the ERC20 token using the approval service from the registry,
links the token back to the asset, and finally mints all fractional tokens to the property owner.
 In short, this contract is the "machine" that assembles a fully functional tokenized property — creating the deed,
 creating the shares, connecting them together, and registering the new property in one transactionn.

*/



// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "../assets/PropertyAsset.sol";
import "../assets/PropertyToken.sol";
import "./PlatformRegistry.sol";

// PropertyFactory - creates tokenized properties on blockchain
contract PropertyFactory is AccessControl {
    // role for creating properties
    bytes32 public constant CREATOR_ROLE = keccak256("CREATOR_ROLE");
    // reference to platform registry for contract addresses
    PlatformRegistry public registry;
    // counter for unique asset IDs
    uint256 private _nextAssetId;

    // event when property is created
    event PropertyRegistered(
        address indexed owner,
        address propertyAsset,
        address propertyToken,
        string propertyDataHash,
        uint256 valuation
    );

    // constructor - setup roles and registry
    constructor(address registryAddr, address admin) {
        registry = PlatformRegistry(registryAddr);
        _setupRole(DEFAULT_ADMIN_ROLE, admin);
        _setupRole(CREATOR_ROLE, admin);
    }

    // createProperty - main function to create tokenized property
    // deploys asset contract, token contract, links them, mints tokens
    function createProperty(
        address owner,
        string memory assetName,
        string memory assetSymbol,
        string memory propertyDataHash,
        uint256 valuation,
        uint256 tokenSupply,
        string memory tokenName,
        string memory tokenSymbol
    ) external {
        require(hasRole(CREATOR_ROLE, msg.sender), "PF: not creator");

        uint256 assetId = ++_nextAssetId;

        // deploy asset
        PropertyAsset asset = new PropertyAsset(
            assetName,
            assetSymbol,
            owner,
            assetId,
            propertyDataHash,
            valuation,
            //msg.senderaddress(this)
            address(this) //factory becomes manager

        );

        // approval service address from registry
        address approval = registry.getApprovalService();

        // deploy token
        PropertyToken token = new PropertyToken(
            tokenName,
            tokenSymbol,
            address(this), 
            address(asset),
            approval
        );

        // link token and asset
        asset.setPropertyToken(address(token));

        // mint tokens to owner
        token.mint(owner, tokenSupply);

        // Grant SNAPSHOT_ROLE to RevenueDistribution so it can create snapshots for revenue distribution
        // This allows RevenueDistribution to call snapshotNow() on this token
        // Get RevenueDistribution address from registry
        address revenueDistributionAddr = registry.getRevenueDistribution();
        if (revenueDistributionAddr != address(0)) {
            bytes32 snapshotRole = keccak256("SNAPSHOT_ROLE");
            token.grantRole(snapshotRole, revenueDistributionAddr);
        }

        emit PropertyRegistered(owner, address(asset), address(token), propertyDataHash, valuation);
    }

    /**
     * Grant SNAPSHOT_ROLE to RevenueDistribution on an existing PropertyToken
     * This is needed for properties created before this auto-granting was added
     */
    function grantSnapshotRoleToRevenue(address tokenAddress) external {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "PF: not admin");
        
        address revenueDistributionAddr = registry.getRevenueDistribution();
        require(revenueDistributionAddr != address(0), "PF: revenue distribution not set in registry");
        
        PropertyToken token = PropertyToken(tokenAddress);
        bytes32 snapshotRole = keccak256("SNAPSHOT_ROLE");
        token.grantRole(snapshotRole, revenueDistributionAddr);
    }
}

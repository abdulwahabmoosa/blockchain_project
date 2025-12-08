/*
The PropertyFactory is responsible for creating a complete tokenized property on the blockchain. 
When the admin (who has the CREATOR_ROLE) calls createProperty, 
the factory automatically generates two linked smart contracts: 
a PropertyAsset (ERC-721 NFT representing the property itself) 
and a PropertyToken (ERC-20 token representing the fractional ownership shares). 
It gives the property a unique ID, stores the metadata hash, and sets the property’s 
valuation. It then deploys the ERC20 token using the approval service from the registry, 
links the token back to the asset, and finally mints all fractional tokens to the property owner.
 In short, this contract is the “machine” that assembles a fully functional tokenized property — creating the deed, 
 creating the shares, connecting them together, and registering the new property in one transaction.

*/



// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "../assets/PropertyAsset.sol";
import "../assets/PropertyToken.sol";
import "./PlatformRegistry.sol";

contract PropertyFactory is AccessControl {
    bytes32 public constant CREATOR_ROLE = keccak256("CREATOR_ROLE");
    PlatformRegistry public registry;
    uint256 private _nextAssetId;

    event PropertyRegistered(
        address indexed owner,
        address propertyAsset,
        address propertyToken,
        string propertyDataHash,
        uint256 valuation
    );

    constructor(address registryAddr, address admin) {
        registry = PlatformRegistry(registryAddr);
        _setupRole(DEFAULT_ADMIN_ROLE, admin);
        _setupRole(CREATOR_ROLE, admin);
    }

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

        emit PropertyRegistered(owner, address(asset), address(token), propertyDataHash, valuation);
    }
}

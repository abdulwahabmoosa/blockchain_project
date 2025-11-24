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
            msg.sender
        );

        // approval service address from registry
        address approval = registry.getApprovalService();

        // deploy token
        PropertyToken token = new PropertyToken(
            tokenName,
            tokenSymbol,
            msg.sender,
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

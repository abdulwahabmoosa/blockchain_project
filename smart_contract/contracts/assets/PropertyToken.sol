
/*
  PropertyToken is the ERC20 token that represents fractional ownership of a real estate property.
 It can only be minted by addresses with the MINTER_ROLE, usually the PropertyFactory when the property is created.
The token includes snapshot functionality, which allows the platform to freeze ownership balances at
  a specific moment in time so that rental revenue can later be distributed fairly to whoever held tokens at that snapshot.
 Transfers of these tokens are restricted: they only work when the entire token contract is not paused,
  the underlying property (represented by an ERC721 PropertyAsset) is in an Active status,
  and the receiving wallet has been approved by the ApprovalService. This ensures compliance-like behavior without KYC.
  The admin (DEFAULT_ADMIN_ROLE) can pause/unpause transfers and trigger snapshots if needed. Before every transfer,
  the token checks three things: the contract isn't paused, the property is active, and the recipient is approved.
  In short: this contract creates, controls, and safely restricts the movement of fractional tokens while enabling
  revenue distribution through snapshotss.

  */
  //Snapshots freeze ERC20 token ownership at the moment rental income is deposited, 
  //so revenue can be split exactly based on who owned tokens at that time — unaffected by later transfers.


// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Snapshot.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/security/Pausable.sol";


//Interfaces let your contract talk to another contract
//without importing its full codee.
//They define what can be called, not how it works.
interface IPropertyAsset {
    enum Status { Active, Paused, Disputed, Closed }
    function getStatus() external view returns (Status);
}

interface IApprovalService {
    function isApproved(address user) external view returns (bool);
}

contract PropertyToken is ERC20,ERC20Snapshot, AccessControl, Pausable {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant SNAPSHOT_ROLE = keccak256("SNAPSHOT_ROLE");

    address public propertyAsset;
    address public approvalService;

    event PropertyAssetLinked(address asset);
    event ApprovalServiceLinked(address approval);

    constructor(
        string memory name_,
        string memory symbol_,
        address admin_,
        address _propertyAsset,
        address _approvalService
    ) ERC20(name_, symbol_) {
        _setupRole(DEFAULT_ADMIN_ROLE, admin_);
        _setupRole(MINTER_ROLE, admin_);
        _setupRole(SNAPSHOT_ROLE, admin_);
        propertyAsset = _propertyAsset;
        approvalService = _approvalService;
        emit PropertyAssetLinked(_propertyAsset);
        emit ApprovalServiceLinked(_approvalService);
    }

    function mint(address to, uint256 amount) external {
        require(hasRole(MINTER_ROLE, msg.sender), "PT: not minter");
        _mint(to, amount);
    }

    function snapshotNow() external returns (uint256) {
        require(hasRole(SNAPSHOT_ROLE, msg.sender), "PT: not snapshot role");
        return _snapshot();
    }

    function pause() external {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "PT: not admin");
        _pause();
    }

    function unpause() external {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "PT: not admin");
        _unpause();
    }

    function _beforeTokenTransfer(address from, address to, uint256 amount)
        internal 
        override(ERC20,ERC20Snapshot)
    {
        super._beforeTokenTransfer(from, to, amount);
        require(!paused(), "PT: token paused");

        //allow mint/burn If a property is de-tokenized or liquidated, the platform destroys all tokens so they no longer exist.
        if (from == address(0) || to == address(0)) {
            return;
        }

        // check property asset status
        IPropertyAsset asset = IPropertyAsset(propertyAsset);
        require(asset.getStatus() == IPropertyAsset.Status.Active, "PT: property not active");

        // check approval service
        IApprovalService approval = IApprovalService(approvalService);
        require(approval.isApproved(to), "PT: recipient not approved");
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract PropertyAsset is ERC721, AccessControl {
    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER_ROLE");

    enum Status { Active, Paused, Disputed, Closed }

    uint256 public assetId;
    string public propertyDataHash; // IPFS CID or SHA-256 hex
    uint256 public valuation; //the real-world price 
    address public propertyToken; // linked ERC20
    Status public status;

    event PropertyStatusChanged(Status newStatus);
    event PropertyTokenLinked(address token);
    event PropertyApproved(uint256 assetId);
    event PropertyRejected(uint256 assetId);

    constructor(
        string memory name_,
        string memory symbol_,
        address owner_,
        uint256 _assetId,
        string memory _dataHash,
        uint256 _valuation,
        address manager
    ) ERC721(name_, symbol_) {
        assetId = _assetId;
        propertyDataHash = _dataHash;
        valuation = _valuation;
        status = Status.Active;
        _grantRole(MANAGER_ROLE, manager);
        _mint(owner_, assetId);
    }

    modifier onlyManager() {
        require(hasRole(MANAGER_ROLE, msg.sender), "PA: not manager");
        _;
    }

    function setPropertyToken(address token) external onlyManager {
        propertyToken = token;
        emit PropertyTokenLinked(token);
    }

    function setStatus(Status newStatus) external onlyManager {
        status = newStatus;
        emit PropertyStatusChanged(newStatus);
    }

    /**
     * @dev Convenience helper for explicitly approving a property on-chain.
     * Mirrors backend approval logic by setting the status to Active and emitting a dedicated event.
     */
    function approveProperty() external onlyManager {
        status = Status.Active;
        emit PropertyStatusChanged(Status.Active);
        emit PropertyApproved(assetId);
    }

    /**
     * @dev Convenience helper for explicitly rejecting a property on-chain.
     * Mirrors backend rejection logic by setting the status to Closed and emitting a dedicated event.
     */
    function rejectProperty() external onlyManager {
        status = Status.Closed;
        emit PropertyStatusChanged(Status.Closed);
        emit PropertyRejected(assetId);
    }

    function getStatus() external view returns (Status) {
        return status;
    }


//Lets tools know this contract supports ERC721 + AccessControl interfaces.
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, AccessControl)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}

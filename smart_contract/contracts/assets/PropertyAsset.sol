// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract PropertyAsset is ERC721, AccessControl {
    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER_ROLE");

    enum Status { Active, Paused, Disputed, Closed }

    uint256 public assetId;
    string public propertyDataHash; // IPFS CID or SHA-256 hex
    uint256 public valuation;
    address public propertyToken; // linked ERC20
    Status public status;

    event PropertyStatusChanged(Status newStatus);
    event PropertyTokenLinked(address token);

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

    function getStatus() external view returns (Status) {
        return status;
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, AccessControl)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}

/*
This contract is responsible for paying rental income (revenue) to token holders in a fair and automated way.
 When the admin (who has the DISTRIBUTOR_ROLE) deposits 
 revenue for a property, the contract first asks the property’s ERC20
  token to create a snapshot, which freezes everyone’s token balances at
   that exact moment. Then the contract stores a new “distribution record” containing 
   the amount deposited, which token it belongs to, and the snapshot ID. Later, 
   any investor who held tokens at snapshot time can call claimRevenue() — the contract checks their balance
    at the snapshot, calculates their share of the revenue based on total supply at that time, 
    ensures they haven’t already claimed, and transfers their portion of stablecoin to them.
     It prevents double claiming and protects against reentrancy attacks. In short: 
     Admin deposits income → snapshot taken → each investor claims their share safely.

*/



// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "../assets/PropertyToken.sol";

contract RevenueDistribution is AccessControl, ReentrancyGuard {
    using SafeERC20 for IERC20;

    bytes32 public constant DISTRIBUTOR_ROLE = keccak256("DISTRIBUTOR_ROLE");

    struct Distribution {
        address token; // property token
        uint256 snapshotId;
        uint256 totalAmount;
        address stablecoin;
        bool withdrawn; // optional
    }

    Distribution[] public distributions;
    mapping(uint256 => mapping(address => bool)) public claimed;

    event RevenueDeposited(uint256 indexed distributionId, address indexed token, uint256 amount, uint256 snapshotId, address stablecoin);
    event RevenueClaimed(uint256 indexed distributionId, address indexed claimant, uint256 amount);

    constructor(address admin) {
        _setupRole(DEFAULT_ADMIN_ROLE, admin);
        _setupRole(DISTRIBUTOR_ROLE, admin);
    }

    function depositRevenue(address tokenAddress, address stablecoin, uint256 amount) external {
        require(hasRole(DISTRIBUTOR_ROLE, msg.sender), "RD: not distributor");
        PropertyToken token = PropertyToken(tokenAddress);

        // create snapshot (token has snapshot role configured)
        uint256 snap = token.snapshotNow();

        // transfer stablecoin into contract
        IERC20(stablecoin).safeTransferFrom(msg.sender, address(this), amount);

        distributions.push(Distribution({
            token: tokenAddress,
            snapshotId: snap,
            totalAmount: amount,
            stablecoin: stablecoin,
            withdrawn: false
        }));

        uint256 id = distributions.length - 1;
        emit RevenueDeposited(id, tokenAddress, amount, snap, stablecoin);
    }

    function claimRevenue(uint256 distributionId) external nonReentrant {
        require(distributionId < distributions.length, "RD: invalid id");
        require(!claimed[distributionId][msg.sender], "RD: already claimed");

        Distribution memory d = distributions[distributionId];
        PropertyToken token = PropertyToken(d.token);

        uint256 holderBalance = token.balanceOfAt(msg.sender, d.snapshotId);
        uint256 totalSupplyAt = token.totalSupplyAt(d.snapshotId);
        require(totalSupplyAt > 0, "RD: zero supply");

        uint256 share = (d.totalAmount * holderBalance) / totalSupplyAt;
        claimed[distributionId][msg.sender] = true;

        IERC20(d.stablecoin).safeTransfer(msg.sender, share);
        emit RevenueClaimed(distributionId, msg.sender, share);
    }

    function distributionsCount() external view returns (uint256) {
        return distributions.length;
    }
}

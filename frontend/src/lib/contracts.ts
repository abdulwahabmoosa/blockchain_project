import { ethers } from "ethers";
import type { BrowserProvider, JsonRpcSigner } from "ethers";
import PropertyTokenABI from "../deployments/sepolia/PropertyToken.abi.json";
import RevenueDistributionABI from "../deployments/sepolia/RevenueDistribution.abi.json";
import ApprovalServiceABI from "../deployments/sepolia/ApprovalService.abi.json";
import { getContractAddresses } from "./wallet";

/**
 * Contracts Interaction Utility
 * 
 * This module provides functions to interact with the smart contracts directly.
 * These functions use ethers.js to read data and send transactions to the blockchain.
 */

const addresses = getContractAddresses();

/**
 * Get token balance for a specific address
 * 
 * PropertyToken is an ERC20 token, so we use the standard balanceOf function.
 * 
 * @param tokenAddress - The PropertyToken contract address
 * @param walletAddress - The wallet address to check balance for
 * @param provider - The ethers provider instance
 * @returns Promise with the token balance as a BigInt
 */
export const getTokenBalance = async (
  tokenAddress: string,
  walletAddress: string,
  provider: BrowserProvider
): Promise<bigint> => {
  // Create a contract instance (read-only, no signer needed)
  const tokenContract = new ethers.Contract(
    tokenAddress,
    PropertyTokenABI,
    provider
  );

  // Call the balanceOf function (standard ERC20)
  const balance = await tokenContract.balanceOf(walletAddress);
  return balance;
};

/**
 * Transfer tokens to another address
 * 
 * This function transfers PropertyToken (ERC20) tokens from the connected wallet
 * to another address. The recipient must be approved by the ApprovalService.
 * 
 * @param tokenAddress - The PropertyToken contract address
 * @param toAddress - The recipient wallet address
 * @param amount - The amount of tokens to transfer (as string or BigInt)
 * @param signer - The ethers signer instance (must be connected wallet)
 * @returns Promise with the transaction response
 */
export const transferTokens = async (
  tokenAddress: string,
  toAddress: string,
  amount: string | bigint,
  signer: JsonRpcSigner
): Promise<ethers.ContractTransactionResponse> => {
  // Create a contract instance with signer (for sending transactions)
  const tokenContract = new ethers.Contract(
    tokenAddress,
    PropertyTokenABI,
    signer
  );

  // Convert amount to BigInt if it's a string
  const amountBigInt = typeof amount === "string" ? BigInt(amount) : amount;

  // Call the transfer function (standard ERC20)
  // This will trigger MetaMask to pop up asking for confirmation
  const tx = await tokenContract.transfer(toAddress, amountBigInt);
  return tx;
};

/**
 * Check if a wallet address is approved in the ApprovalService
 *
 * The ApprovalService maintains a whitelist of approved users.
 * Only approved users can receive PropertyToken transfers.
 *
 * @param walletAddress - The wallet address to check
 * @param provider - The ethers provider instance
 * @returns Promise with boolean indicating if the address is approved
 */
export const checkApprovalStatus = async (
  walletAddress: string,
  provider: BrowserProvider | ethers.AbstractProvider
): Promise<boolean> => {
  console.log('🔧 checkApprovalStatus v2.0 - NEW CODE LOADED'); // Unique marker to verify new code
  const approvalServiceAddress = addresses.ApprovalService;

  // Validate inputs
  if (!walletAddress || !ethers.isAddress(walletAddress)) {
    throw new Error(`Invalid wallet address: ${walletAddress}`);
  }

  // Check if contract is deployed (not a zero address)
  if (!approvalServiceAddress || approvalServiceAddress === "0x0000000000000000000000000000000000000000") {
    console.log(`⚠️ ApprovalService contract not deployed locally. Assuming user is not approved.`);
    return false; // Assume not approved if contract not deployed
  }

  if (!ethers.isAddress(approvalServiceAddress)) {
    throw new Error(`Invalid approval service address: ${approvalServiceAddress}`);
  }

  console.log(`🔍 Checking approval for ${walletAddress} on contract ${approvalServiceAddress}`);
  console.log(`📋 Contract details:`, {
    approvalServiceAddress,
    walletAddress
  });

  // First, check if contract exists at the address
  try {
    const contractCode = await provider.getCode(approvalServiceAddress);
    if (!contractCode || contractCode === "0x" || contractCode === "0x0") {
      console.log(`⚠️ No contract deployed at ${approvalServiceAddress}. This may indicate:`);
      console.log(`   - Contracts not deployed to this network`);
      console.log(`   - Wrong network selected in wallet`);
      console.log(`   - RPC endpoint issue`);
      console.log(`   Returning false - caller should check persisted approvals as fallback.`);
      return false;
    }
    console.log(`✅ Contract exists at ${approvalServiceAddress} (bytecode length: ${contractCode.length})`);
  } catch (codeCheckError: any) {
    console.warn(`⚠️ Failed to check contract code at ${approvalServiceAddress}:`, codeCheckError);
    console.warn(`   This may be an RPC issue. Caller should check persisted approvals as fallback.`);
    // Continue anyway - might still be able to call
  }

  // Create a contract instance (read-only)
  const approvalContract = new ethers.Contract(
    approvalServiceAddress,
    ApprovalServiceABI,
    provider
  );

  // Call the check function with robust retry logic
  let lastError: any = null;

  for (let attempt = 1; attempt <= 3; attempt++) {
    try {
      console.log(`🔄 Approval check attempt ${attempt} for ${walletAddress}`);
      console.log(`📡 Calling contract.check(${walletAddress})`);

      // Add timeout to individual call
      const checkPromise = approvalContract.check(walletAddress);
      const timeoutPromise = new Promise((_, reject) =>
        setTimeout(() => reject(new Error(`Contract call timeout (attempt ${attempt})`)), attempt === 1 ? 5000 : 3000)
      );

      const isApproved = await Promise.race([checkPromise, timeoutPromise]);
      console.log(`📋 Raw contract response for ${walletAddress}:`, isApproved);

      // Validate the result is boolean
      if (typeof isApproved !== 'boolean') {
        throw new Error(`Invalid response type from contract: ${typeof isApproved}`);
      }

      console.log(`✅ Approval check successful for ${walletAddress}: ${isApproved}`);
      return isApproved;

    } catch (error: any) {
      // IMMEDIATELY check for BAD_DATA errors - don't retry these
      // Check error code, message, and string representation
      const errorCode = error?.code || error?.info?.error?.code || error?.info?.code;
      const errorMessage = String(error?.message || '');
      const errorString = String(error || '');
      
      // DEBUG: Log error structure to understand what we're dealing with
      console.log('🔍 ERROR DEBUG:', {
        errorCode,
        errorMessage,
        hasCodeBadData: errorCode === 'BAD_DATA',
        messageHasDecode: errorMessage.includes('could not decode result data'),
        messageHasBadData: errorMessage.includes('BAD_DATA'),
        errorStringHasBadData: errorString.includes('BAD_DATA')
      });
      
      const isBadDataError = 
        errorCode === 'BAD_DATA' || 
        errorCode === 'CALL_EXCEPTION' ||
        errorMessage.includes('could not decode result data') ||
        errorMessage.includes('BAD_DATA') ||
        errorMessage.includes('value="0x"') ||
        errorString.includes('BAD_DATA') ||
        errorString.includes('could not decode') ||
        errorString.includes('code=BAD_DATA');
      
      if (isBadDataError) {
        console.log(`✅ DETECTED BAD_DATA ERROR - Returning false immediately`);
        console.log(`⚠️ Contract call failed with BAD_DATA/CALL_EXCEPTION (contract may not exist or method not available at ${approvalServiceAddress}). Assuming user is not approved.`);
        return false;
      }
      
      console.warn(`⚠️ Approval check attempt ${attempt} failed for ${walletAddress}:`, error);
      
      // Log detailed error info for debugging (only if not BAD_DATA)
      const errorDetails = {
        code: error.code,
        message: error.message,
        reason: error.reason,
        info: error.info,
        infoCode: error.info?.error?.code,
        errorString: String(error)
      };
      console.warn(`📋 Error details:`, errorDetails);
      lastError = error;

      // Wait before retry (exponential backoff) - only for other types of errors
      if (attempt < 3) {
        await new Promise(resolve => setTimeout(resolve, 1000 * attempt));
      }
    }
  }

  console.error(`❌ All approval check attempts failed for ${walletAddress}:`, lastError);
  throw new Error(`Failed to check approval status after 3 attempts: ${lastError?.message || 'Unknown error'}`);
};


/**
 * Get the total number of revenue distributions
 * 
 * @param provider - The ethers provider instance
 * @returns Promise with the count of distributions
 */
export const getDistributionsCount = async (
  provider: BrowserProvider
): Promise<bigint> => {
  const revenueDistributionAddress = addresses.RevenueDistribution;

  const revenueContract = new ethers.Contract(
    revenueDistributionAddress,
    RevenueDistributionABI,
    provider
  );

  const count = await revenueContract.distributionsCount();
  return count;
};

/**
 * Get distribution details
 * 
 * Returns information about a specific revenue distribution including:
 * - token address
 * - snapshot ID
 * - total amount
 * - stablecoin address
 * 
 * @param distributionId - The ID of the distribution (index in the array)
 * @param provider - The ethers provider instance
 * @returns Promise with distribution details
 */
export const getDistribution = async (
  distributionId: number,
  provider: BrowserProvider
): Promise<{
  token: string;
  snapshotId: bigint;
  totalAmount: bigint;
  stablecoin: string;
  withdrawn: boolean;
}> => {
  const revenueDistributionAddress = addresses.RevenueDistribution;

  const revenueContract = new ethers.Contract(
    revenueDistributionAddress,
    RevenueDistributionABI,
    provider
  );

  // Access the distributions array (public, so we can call it as a function)
  const distribution = await revenueContract.distributions(distributionId);

  return {
    token: distribution.token,
    snapshotId: distribution.snapshotId,
    totalAmount: distribution.totalAmount,
    stablecoin: distribution.stablecoin,
    withdrawn: distribution.withdrawn,
  };
};

/**
 * Check if a user has already claimed a specific distribution
 * 
 * @param distributionId - The ID of the distribution
 * @param walletAddress - The wallet address to check
 * @param provider - The ethers provider instance
 * @returns Promise with boolean indicating if already claimed
 */
export const checkDistributionClaimed = async (
  distributionId: number,
  walletAddress: string,
  provider: BrowserProvider
): Promise<boolean> => {
  const revenueDistributionAddress = addresses.RevenueDistribution;

  const revenueContract = new ethers.Contract(
    revenueDistributionAddress,
    RevenueDistributionABI,
    provider
  );

  // Access the claimed mapping (public, so we can call it as a function)
  const claimed = await revenueContract.claimed(distributionId, walletAddress);
  return claimed;
};

/**
 * Claim revenue from a distribution
 * 
 * This function allows a token holder to claim their share of revenue
 * from a distribution. The user must have held tokens at the snapshot time.
 * 
 * @param distributionId - The ID of the distribution to claim from
 * @param signer - The ethers signer instance (must be connected wallet)
 * @returns Promise with the transaction response
 */
export const claimRevenue = async (
  distributionId: number,
  signer: JsonRpcSigner
): Promise<ethers.ContractTransactionResponse> => {
  const revenueDistributionAddress = addresses.RevenueDistribution;

  // Create a contract instance with signer (for sending transactions)
  const revenueContract = new ethers.Contract(
    revenueDistributionAddress,
    RevenueDistributionABI,
    signer
  );

  // Call the claimRevenue function
  // This will trigger MetaMask to pop up asking for confirmation
  const tx = await revenueContract.claimRevenue(distributionId);
  return tx;
};

/**
 * Get token balance at a specific snapshot
 * 
 * Used for calculating how much revenue a user can claim from a distribution.
 * 
 * @param tokenAddress - The PropertyToken contract address
 * @param walletAddress - The wallet address to check
 * @param snapshotId - The snapshot ID to check balance at
 * @param provider - The ethers provider instance
 * @returns Promise with the token balance at snapshot
 */
export const getTokenBalanceAtSnapshot = async (
  tokenAddress: string,
  walletAddress: string,
  snapshotId: bigint,
  provider: BrowserProvider
): Promise<bigint> => {
  const tokenContract = new ethers.Contract(
    tokenAddress,
    PropertyTokenABI,
    provider
  );

  // Call balanceOfAt function (from ERC20Snapshot)
  const balance = await tokenContract.balanceOfAt(walletAddress, snapshotId);
  return balance;
};

/**
 * Get total token supply at a specific snapshot
 * 
 * Used for calculating revenue shares.
 * 
 * @param tokenAddress - The PropertyToken contract address
 * @param snapshotId - The snapshot ID
 * @param provider - The ethers provider instance
 * @returns Promise with the total supply at snapshot
 */
export const getTotalSupplyAtSnapshot = async (
  tokenAddress: string,
  snapshotId: bigint,
  provider: BrowserProvider
): Promise<bigint> => {
  const tokenContract = new ethers.Contract(
    tokenAddress,
    PropertyTokenABI,
    provider
  );

  // Call totalSupplyAt function (from ERC20Snapshot)
  const supply = await tokenContract.totalSupplyAt(snapshotId);
  return supply;
};

/**
 * Calculate claimable revenue amount for a user
 * 
 * This helper function calculates how much revenue a user can claim
 * from a distribution based on their token balance at the snapshot.
 * 
 * Formula: (userBalance / totalSupply) * totalAmount
 * 
 * @param userBalance - User's token balance at snapshot
 * @param totalSupply - Total token supply at snapshot
 * @param totalAmount - Total revenue amount in distribution
 * @returns The claimable amount
 */
export const calculateClaimableAmount = (
  userBalance: bigint,
  totalSupply: bigint,
  totalAmount: bigint
): bigint => {
  if (totalSupply === 0n) {
    return 0n;
  }
  // Calculate: (userBalance * totalAmount) / totalSupply
  return (userBalance * totalAmount) / totalSupply;
};

/**
 * Get all claimable distributions for a user
 * 
 * This function checks all distributions and determines which ones
 * the user can claim (has balance at snapshot time and hasn't claimed yet).
 * 
 * @param tokenAddress - The PropertyToken contract address
 * @param walletAddress - The wallet address to check for
 * @param provider - The ethers provider instance
 * @returns Promise with array of claimable distribution IDs and amounts
 */
export const getClaimableDistributions = async (
  tokenAddress: string,
  walletAddress: string,
  provider: BrowserProvider
): Promise<
  Array<{
    distributionId: number;
    amount: bigint;
    stablecoin: string;
  }>
> => {
  const count = await getDistributionsCount(provider);
  const claimable: Array<{
    distributionId: number;
    amount: bigint;
    stablecoin: string;
  }> = [];

  // Iterate through all distributions
  for (let i = 0; i < Number(count); i++) {
    const distribution = await getDistribution(i, provider);

    // Check if this distribution is for the specified token
    if (distribution.token.toLowerCase() !== tokenAddress.toLowerCase()) {
      continue;
    }

    // Check if user has already claimed
    const claimed = await checkDistributionClaimed(i, walletAddress, provider);
    if (claimed) {
      continue;
    }

    // Get user's balance at snapshot
    const userBalance = await getTokenBalanceAtSnapshot(
      tokenAddress,
      walletAddress,
      distribution.snapshotId,
      provider
    );

    // If user has no balance at snapshot, skip
    if (userBalance === 0n) {
      continue;
    }

    // Get total supply at snapshot
    const totalSupply = await getTotalSupplyAtSnapshot(
      tokenAddress,
      distribution.snapshotId,
      provider
    );

    // Calculate claimable amount
    const claimableAmount = calculateClaimableAmount(
      userBalance,
      totalSupply,
      distribution.totalAmount
    );

    claimable.push({
      distributionId: i,
      amount: claimableAmount,
      stablecoin: distribution.stablecoin,
    });
  }

  return claimable;
};








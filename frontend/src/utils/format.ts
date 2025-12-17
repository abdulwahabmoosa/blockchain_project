/**
 * Format BigInt token balance to readable string
 */
export const formatTokenBalance = (balance: bigint, decimals: number = 18): string => {
  const divisor = BigInt(10 ** decimals);
  const wholePart = balance / divisor;
  const fractionalPart = balance % divisor;
  
  if (fractionalPart === 0n) {
    return wholePart.toString();
  }
  
  const fractionalStr = fractionalPart.toString().padStart(decimals, "0");
  const trimmedFractional = fractionalStr.replace(/0+$/, "");
  
  if (trimmedFractional.length === 0) {
    return wholePart.toString();
  }
  
  return `${wholePart}.${trimmedFractional}`;
};

/**
 * Format address to shortened version
 */
export const formatAddress = (address: string): string => {
  if (!address) return "";
  return `${address.substring(0, 6)}...${address.substring(38)}`;
};





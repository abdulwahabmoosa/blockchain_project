const ethers = require('ethers');

async function checkUserBalance() {
  try {
    const provider = new ethers.JsonRpcProvider('https://ethereum-sepolia-rpc.publicnode.com');

    // User's wallet address from the dashboard
    const userAddress = '0x2b4F...0509'; // This needs to be the full address
    console.log('User address (truncated):', userAddress);

    // First, let's check the user's ETH balance
    const balance = await provider.getBalance(userAddress);
    console.log('ETH Balance:', ethers.formatEther(balance));

    // Now let's try to check some known property token addresses
    // We need to know what property tokens exist

    // PropertyToken ABI - minimal for balanceOf
    const tokenAbi = [
      "function balanceOf(address) view returns (uint256)",
      "function name() view returns (string)",
      "function symbol() view returns (string)",
      "function totalSupply() view returns (uint256)"
    ];

    // Example token addresses - replace with actual ones from your deployment
    // These would come from the database/backend
    const testTokenAddresses = [
      // Add some known token addresses here if you have them
    ];

    if (testTokenAddresses.length === 0) {
      console.log('No token addresses to check. Need to get them from the database or factory contract.');
      return;
    }

    for (const tokenAddr of testTokenAddresses) {
      try {
        console.log(`\nChecking token: ${tokenAddr}`);
        const token = new ethers.Contract(tokenAddr, tokenAbi, provider);

        const name = await token.name();
        const symbol = await token.symbol();
        const totalSupply = await token.totalSupply();
        const userBalance = await token.balanceOf(userAddress);

        console.log(`Token: ${name} (${symbol})`);
        console.log(`Total Supply: ${ethers.formatEther(totalSupply)}`);
        console.log(`User Balance: ${ethers.formatEther(userBalance)}`);
      } catch (err) {
        console.error(`Error checking token ${tokenAddr}:`, err.message);
      }
    }

  } catch (error) {
    console.error('Error checking user balance:', error);
  }
}

// For now, let's just check if we can get the full user address
// Since the user showed 0x2b4F...0509, let's assume it's a valid address
// But we need the full address to check balances

console.log('Please provide the full wallet address to check balances.');
console.log('From the dashboard, it shows: 0x2b4F...0509');
console.log('We need the complete address like: 0x2b4Fxxxxxxxxxxxxxxxxxxxxxxxxx0509');

// checkUserBalance();
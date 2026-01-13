const ethers = require('ethers');

async function checkEvents() {
  const provider = new ethers.providers.JsonRpcProvider('https://ethereum-sepolia-rpc.publicnode.com');
  const factoryAddress = '0x9Aa0a2Ee2Cb91587484b0F7F2BE992E1F4484F1e';

  const factory = new ethers.Contract(factoryAddress, [
    'event PropertyRegistered(address indexed owner, address propertyAsset, address propertyToken, string propertyDataHash, uint256 valuation)'
  ], provider);

  console.log('Checking PropertyRegistered events on PropertyFactory...');
  console.log('Factory Address:', factoryAddress);

  try {
    const currentBlock = await provider.getBlockNumber();
    console.log('Current block:', currentBlock);

    // Query last 10000 blocks
    const fromBlock = Math.max(0, currentBlock - 10000);
    console.log('Querying from block', fromBlock, 'to', currentBlock);

    const events = await factory.queryFilter('PropertyRegistered', fromBlock, currentBlock);
    console.log('Found', events.length, 'PropertyRegistered events');
    console.log('');

    if (events.length === 0) {
      console.log('No PropertyRegistered events found in the last 10000 blocks.');
      return;
    }

    // Show events in reverse chronological order (newest first)
    events.reverse().forEach((event, i) => {
      console.log(`=== PROPERTY ${events.length - i} (Latest) ===`);
      console.log('Owner Address:', event.args.owner);
      console.log('Property Asset Contract:', event.args.propertyAsset);
      console.log('Property Token Contract:', event.args.propertyToken);
      console.log('IPFS Metadata Hash:', event.args.propertyDataHash);
      console.log('Property Valuation:', ethers.utils.formatEther(event.args.valuation), 'ETH');
      console.log('Block Number:', event.blockNumber);
      console.log('Transaction Hash:', event.transactionHash);
      console.log('Transaction Link: https://sepolia.etherscan.io/tx/' + event.transactionHash);
      console.log('');
    });

  } catch (error) {
    console.error('Error querying blockchain:', error.message);
  }
}

checkEvents().catch(console.error);


const { ethers } = require('hardhat');

async function checkRegistry() {
  const provider = new ethers.providers.JsonRpcProvider(process.env.SEPOLIA_RPC);
  const signer = new ethers.Wallet(process.env.PRIVATE_KEY || process.env.PRIVATE_KEY_2, provider);

  const registryAddress = '0x113FdC450314d59eAbEEfA2ff9F232DD7b08ad38';
  const factoryAddress = '0x31531Cea335B72E05493924af6d97Ef3C89C49bd';
  const approvalAddress = '0xa44254E9dC27069a56582C36521722DC0A3ff39E';
  const revenueAddress = '0xD7ebec79B2e8D218F59A3777fa615eae762Bb7f6';

  const Registry = await ethers.getContractFactory('PlatformRegistry');
  const registry = Registry.attach(registryAddress);

  console.log('Checking registry addresses...');

  const storedFactory = await registry.getPropertyFactory();
  console.log('Stored PropertyFactory:', storedFactory);
  console.log('Expected PropertyFactory:', factoryAddress);
  console.log('Factory address matches:', storedFactory.toLowerCase() === factoryAddress.toLowerCase());

  const storedApproval = await registry.getApprovalService();
  console.log('Stored ApprovalService:', storedApproval);
  console.log('Expected ApprovalService:', approvalAddress);
  console.log('Approval address matches:', storedApproval.toLowerCase() === approvalAddress.toLowerCase());

  const storedRevenue = await registry.getRevenueDistribution();
  console.log('Stored RevenueDistribution:', storedRevenue);
  console.log('Expected RevenueDistribution:', revenueAddress);
  console.log('Revenue address matches:', storedRevenue.toLowerCase() === revenueAddress.toLowerCase());
}

checkRegistry().catch(console.error);



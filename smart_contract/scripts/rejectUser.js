const hre = require('hardhat');

async function main() {
  const [admin] = await hre.ethers.getSigners();

  const registryAddr = process.env.REGISTRY;
  const targetWallet = process.env.REJECT_WALLET;

  if (!registryAddr) {
    console.log('Set REGISTRY env var first.');
    return;
  }

  if (!targetWallet) {
    console.log('Set REJECT_WALLET env var to the user wallet to reject.');
    return;
  }

  const Registry = await hre.ethers.getContractAt(
    'PlatformRegistry',
    registryAddr
  );
  const approvalAddr = await Registry.getApprovalService();

  const Approval = await hre.ethers.getContractAt(
    'ApprovalService',
    approvalAddr
  );

  console.log('Rejecting user on-chain:', targetWallet);
  const tx = await Approval.connect(admin).rejectUser(targetWallet);
  console.log('Transaction submitted:', tx.hash);
  await tx.wait();
  console.log('User rejected on-chain:', targetWallet);
}

main()
  .then(() => process.exit(0))
  .catch((err) => {
    console.error(err);
    process.exit(1);
  });



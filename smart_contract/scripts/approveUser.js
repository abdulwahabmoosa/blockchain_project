const hre = require('hardhat');

async function main() {
  const [admin, investor] = await hre.ethers.getSigners();

  const registryAddr = process.env.REGISTRY;
  if (!registryAddr) {
    console.log('Set REGISTRY env var first.');
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

  await Approval.connect(admin).approve(investor.address);

  console.log('Approved investor:', investor.address);
}

// main().catch(err => {
//   console.error(err);
//   process.exitCode = 1;
// });
main()
  .then(() => process.exit(0))
  .catch((err) => {
    console.error(err);
    process.exit(1);
  });

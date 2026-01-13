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

  // Approve the admin address first
  const adminAddress = "0x98e5a749e25c56e19c28008505df75aff4988049";
  await Approval.connect(admin).approve(adminAddress);
  console.log('Approved admin:', adminAddress);

  // Also approve the investor if provided
  if (investor) {
    await Approval.connect(admin).approve(investor.address);
    console.log('Approved investor:', investor.address);
  }
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

import { FirmexProductNFT } from '../src/types/FirmexProductNFT'
import { ethers } from 'hardhat'

async function main() {
  const [deployer] = await ethers.getSigners();

  console.log("Deploying contracts with the account:", deployer.address);

  console.log("Account balance: ", (await deployer.getBalance()).toString());

  const FirmexProductNFT = await ethers.getContractFactory("FirmexProductNFT");
  const firmexProductNFT = await FirmexProductNFT.deploy("http://localhost:3000");

  console.log("FirmexProductNFT address:", firmexProductNFT.address);

  const ProductFactory = await ethers.getContractFactory("ProductFactory");
  const productFactory = await ProductFactory.deploy(firmexProductNFT.address);

  console.log("ProductFactory address:", productFactory.address);

  const deployedNft = (await ethers.getContractAt("FirmexProductNFT", firmexProductNFT.address, deployer)) as FirmexProductNFT;
  await deployedNft.transferOwnership(productFactory.address)
    .then(transaction => {
      return transaction.wait(1)
    }
    )
    .then(receipt => {
      if (receipt.status !== 1) {
        console.error("Transfering ownership of deployed nft failed")
        console.error(JSON.stringify(receipt))
      }
    })

  console.log("Ownership of product nft changed properly. New owner address: " + await deployedNft.owner())
  console.log("done")
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
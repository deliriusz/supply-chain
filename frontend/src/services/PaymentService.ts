import { ethers } from 'ethers'
import { ProductFactory as ProductFactoryContract } from '../types/ProductFactory';
import ProductFactory from '../../artifacts/contracts/ProductFactory.sol/ProductFactory.json';
import { provider } from "../web3"

const processPayment = async (productId: number, amount: number, unit = 'gwei'): Promise<any> => {
   const amountInWei = ethers.utils.parseUnits(`${amount}`, unit)

   await provider.send("eth_requestAccounts", []);
   const signer = provider.getSigner();
   const signerAddress = await signer.getAddress()

   const deployedContract = new ethers.Contract(process.env.FIRMEX_PRODUCT_LIBRARY_CONTRACT_ADDRESS!, ProductFactory.abi, signer) as ProductFactoryContract

   deployedContract.sellProduct(productId, signerAddress, { value: amountInWei, from: signerAddress })
   // const deployedNft = (await ethers.getContractAt("ProductFactory", `${process.env.FIRMEX_PRODUCT_LIBRARY_CONTRACT_ADDRESS}`, userSigner)) as ProductFactory;
   // const deployedNft = (ethers.getContractAt("ProductFactory", `${process.env.FIRMEX_PRODUCT_LIBRARY_CONTRACT_ADDRESS}`, deployer)) as Promise<ProductFactory>;
   // deployedNft.sellProduct(productId, userSigner.address, { value: amountInWei, from: deployedNft.address })
   // return web3.
}

export { processPayment }
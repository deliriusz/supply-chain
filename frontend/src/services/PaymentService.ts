import { ethers } from 'ethers'
import { ProductFactory as ProductFactoryContract } from '../types/ProductFactory';
import ProductFactory from '../artifacts/src/contracts/ProductFactory.sol/ProductFactory.json';
import { getContract, getSigner } from './EthereumHelpers';

const processPayment = async (productId: number, amount: number, unit = 'gwei'): Promise<boolean> => {
   const amountInWei = ethers.utils.parseUnits(`${amount}`, unit)

   const signer = await getSigner()
   const signerAddress = await signer.getAddress()

   const deployedContract = await getContract<ProductFactoryContract>(process.env.REACT_APP_FIRMEX_PRODUCT_LIBRARY_CONTRACT_ADDRESS!, ProductFactory.abi, signer)

   return deployedContract.sellProduct(
      productId,
      signerAddress,
      {
         value: amountInWei,
         from: signerAddress
      }).then(transaction => {
         return transaction.wait(1)
      }).then(receipt => {
         return receipt.status === 1
      })
}

export { processPayment }
import { Provider } from "@ethersproject/providers"
import { ContractInterface, ethers, Signer } from "ethers"
import { provider } from "../web3";


const getContract = async <T>(addressOrName: string, contractInterface: ContractInterface, signerOrProvider: Signer | Provider): Promise<T> => {
   return new ethers.Contract(addressOrName, contractInterface, signerOrProvider) as unknown as T
}

const getSigner = async (): Promise<Signer> => {
   await provider.send("eth_requestAccounts", []);
   return provider.getSigner();
}

export { getContract, getSigner }
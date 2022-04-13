import { ethers } from "ethers";
import Web3 from "web3"

let web3: Web3 | undefined;
let ethereum = (window as any).ethereum
let provider: ethers.providers.Web3Provider

if (typeof ethereum !== "undefined") {
   web3 = new Web3(ethereum)
   provider = new ethers.providers.Web3Provider(ethereum, "any")
   ethereum.enable()
}

export default web3
export { provider }
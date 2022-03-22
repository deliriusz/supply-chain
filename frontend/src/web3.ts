import Web3 from "web3"

let web3: Web3 | undefined;
let ethereum = (window as any).ethereum

if (typeof ethereum !== "undefined") {
   web3 = new Web3(ethereum)
   ethereum.enable()
}

export default web3;
/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import { Provider, TransactionRequest } from "@ethersproject/providers";
import type {
  ProductFactory,
  ProductFactoryInterface,
} from "../ProductFactory";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "nftAddress",
        type: "address",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "previousOwner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "OwnershipTransferred",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "string",
        name: "_name",
        type: "string",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "_initialPrice",
        type: "uint256",
      },
      {
        indexed: true,
        internalType: "uint256",
        name: "_extId",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "string",
        name: "_nftUri",
        type: "string",
      },
    ],
    name: "ProductCreated",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "uint256",
        name: "_forPrice",
        type: "uint256",
      },
      {
        indexed: true,
        internalType: "uint256",
        name: "_extId",
        type: "uint256",
      },
      {
        indexed: true,
        internalType: "address",
        name: "_to",
        type: "address",
      },
    ],
    name: "ProductSold",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "_extId",
        type: "uint256",
      },
      {
        internalType: "enum ProductLib.LifecycleState",
        name: "_newState",
        type: "uint8",
      },
    ],
    name: "changeProductState",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "_name",
        type: "string",
      },
      {
        internalType: "uint256",
        name: "_initialPrice",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "_extId",
        type: "uint256",
      },
      {
        internalType: "string",
        name: "_nftUri",
        type: "string",
      },
    ],
    name: "create",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "getNftAddress",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "_extId",
        type: "uint256",
      },
    ],
    name: "getProduct",
    outputs: [
      {
        components: [
          {
            internalType: "enum ProductLib.LifecycleState",
            name: "state",
            type: "uint8",
          },
          {
            internalType: "string",
            name: "name",
            type: "string",
          },
          {
            internalType: "uint256",
            name: "initialPrice",
            type: "uint256",
          },
        ],
        internalType: "struct ProductLib.Product",
        name: "",
        type: "tuple",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
      {
        internalType: "address",
        name: "",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
      {
        internalType: "bytes",
        name: "",
        type: "bytes",
      },
    ],
    name: "onERC721Received",
    outputs: [
      {
        internalType: "bytes4",
        name: "",
        type: "bytes4",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "owner",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "renounceOwnership",
    outputs: [],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "_extId",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "_to",
        type: "address",
      },
    ],
    name: "sellProduct",
    outputs: [],
    stateMutability: "payable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    name: "transferOwnership",
    outputs: [],
    stateMutability: "view",
    type: "function",
  },
];

const _bytecode =
  "0x60806040523480156200001157600080fd5b5060405162001b5f38038062001b5f833981810160405281019062000037919062000182565b620000576200004b6200009f60201b60201c565b620000a760201b60201c565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000207565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000815190506200017c81620001ed565b92915050565b6000602082840312156200019b576200019a620001e8565b5b6000620001ab848285016200016b565b91505092915050565b6000620001c182620001c8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620001f881620001b4565b81146200020457600080fd5b50565b61194880620002176000396000f3fe6080604052600436106100865760003560e01c80638da5cb5b116100595780638da5cb5b14610124578063a59e2f0e1461014f578063b9db15b414610178578063be9a71bd146101b5578063f2fde38b146101e057610086565b80630b5691401461008b578063150b7a02146100a75780635a886d79146100e4578063715018a61461010d575b600080fd5b6100a560048036038101906100a09190610ffb565b610209565b005b3480156100b357600080fd5b506100ce60048036038101906100c99190610e9f565b610503565b6040516100db9190611364565b60405180910390f35b3480156100f057600080fd5b5061010b60048036038101906101069190610f27565b610573565b005b34801561011957600080fd5b5061012261085f565b005b34801561013057600080fd5b506101396108e0565b60405161014691906112d2565b60405180910390f35b34801561015b57600080fd5b506101766004803603810190610171919061103b565b610909565b005b34801561018457600080fd5b5061019f600480360381019061019a9190610fce565b610a9f565b6040516101ac919061147f565b60405180910390f35b3480156101c157600080fd5b506101ca610bfe565b6040516101d791906112d2565b60405180910390f35b3480156101ec57600080fd5b5061020760048036038101906102029190610e72565b610c28565b005b816000600160008381526020019081526020016000206002015411610263576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025a9061143f565b60405180910390fd5b6000600160008581526020019081526020016000209050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156102ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e1906113bf565b60405180910390fd5b80600201543414610330576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610327906113ff565b60405180910390fd5b600160038111156103445761034361160c565b5b8160000160009054906101000a900460ff1660038111156103685761036761160c565b5b146103a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039f9061145f565b60405180910390fd5b60028160000160006101000a81548160ff021916908360038111156103d0576103cf61160c565b5b0217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342842e0e3085876040518463ffffffff1660e01b8152600401610434939291906112ed565b600060405180830381600087803b15801561044e57600080fd5b505af1158015610462573d6000803e3d6000fd5b5050505061046e6108e0565b73ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f193505050501580156104b3573d6000803e3d6000fd5b508273ffffffffffffffffffffffffffffffffffffffff168482600201547f8e0ba5781a71057cc312a55c76a829d419224994c1502a2560c2a6cb2c498ecb60405160405180910390a450505050565b600083600060016000838152602001908152602001600020600201541161055f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105569061143f565b60405180910390fd5b63150b7a0260e01b91505095945050505050565b61057b610ca9565b73ffffffffffffffffffffffffffffffffffffffff166105996108e0565b73ffffffffffffffffffffffffffffffffffffffff16146105ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e69061141f565b60405180910390fd5b6000600160008581526020019081526020016000206002015414610648576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063f906113df565b60405180910390fd5b6000841161068b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106829061137f565b60405180910390fd5b60006040518060600160405280600060038111156106ac576106ab61160c565b5b815260200188888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152602001868152509050806001600086815260200190815260200160002060008201518160000160006101000a81548160ff021916908360038111156107405761073f61160c565b5b02179055506020820151816001019080519060200190610761929190610cb1565b5060408201518160020155905050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c772745a308585886040518563ffffffff1660e01b81526004016107d09493929190611324565b600060405180830381600087803b1580156107ea57600080fd5b505af11580156107fe573d6000803e3d6000fd5b505050508387876040516108139291906112b9565b60405180910390207fae7a6a88e2960ae3dad5ea83186c61a418b547693aa235642ca288c47d72c11a87868660405161084e939291906114a1565b60405180910390a350505050505050565b610867610ca9565b73ffffffffffffffffffffffffffffffffffffffff166108856108e0565b73ffffffffffffffffffffffffffffffffffffffff16146108db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d29061141f565b60405180910390fd5b600080fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610911610ca9565b73ffffffffffffffffffffffffffffffffffffffff1661092f6108e0565b73ffffffffffffffffffffffffffffffffffffffff1614610985576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097c9061141f565b60405180910390fd5b8160006001600083815260200190815260200160002060020154116109df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d69061143f565b60405180910390fd5b6000600160008581526020019081526020016000209050826003811115610a0957610a0861160c565b5b8160000160009054906101000a900460ff166003811115610a2d57610a2c61160c565b5b10610a6d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a649061139f565b60405180910390fd5b828160000160006101000a81548160ff02191690836003811115610a9457610a9361160c565b5b021790555050505050565b610aa7610d37565b816000600160008381526020019081526020016000206002015411610b01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af89061143f565b60405180910390fd5b600160008481526020019081526020016000206040518060600160405290816000820160009054906101000a900460ff166003811115610b4457610b4361160c565b5b6003811115610b5657610b5561160c565b5b8152602001600182018054610b6a906115da565b80601f0160208091040260200160405190810160405280929190818152602001828054610b96906115da565b8015610be35780601f10610bb857610100808354040283529160200191610be3565b820191906000526020600020905b815481529060010190602001808311610bc657829003601f168201915b50505050508152602001600282015481525050915050919050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610c30610ca9565b73ffffffffffffffffffffffffffffffffffffffff16610c4e6108e0565b73ffffffffffffffffffffffffffffffffffffffff1614610ca4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9b9061141f565b60405180910390fd5b600080fd5b600033905090565b828054610cbd906115da565b90600052602060002090601f016020900481019282610cdf5760008555610d26565b82601f10610cf857805160ff1916838001178555610d26565b82800160010185558215610d26579182015b82811115610d25578251825591602001919060010190610d0a565b5b509050610d339190610d6a565b5090565b604051806060016040528060006003811115610d5657610d5561160c565b5b815260200160608152602001600081525090565b5b80821115610d83576000816000905550600101610d6b565b5090565b600081359050610d96816118d4565b92915050565b60008083601f840112610db257610db161166f565b5b8235905067ffffffffffffffff811115610dcf57610dce61166a565b5b602083019150836001820283011115610deb57610dea611674565b5b9250929050565b600081359050610e01816118eb565b92915050565b60008083601f840112610e1d57610e1c61166f565b5b8235905067ffffffffffffffff811115610e3a57610e3961166a565b5b602083019150836001820283011115610e5657610e55611674565b5b9250929050565b600081359050610e6c816118fb565b92915050565b600060208284031215610e8857610e8761167e565b5b6000610e9684828501610d87565b91505092915050565b600080600080600060808688031215610ebb57610eba61167e565b5b6000610ec988828901610d87565b9550506020610eda88828901610d87565b9450506040610eeb88828901610e5d565b935050606086013567ffffffffffffffff811115610f0c57610f0b611679565b5b610f1888828901610d9c565b92509250509295509295909350565b60008060008060008060808789031215610f4457610f4361167e565b5b600087013567ffffffffffffffff811115610f6257610f61611679565b5b610f6e89828a01610e07565b96509650506020610f8189828a01610e5d565b9450506040610f9289828a01610e5d565b935050606087013567ffffffffffffffff811115610fb357610fb2611679565b5b610fbf89828a01610e07565b92509250509295509295509295565b600060208284031215610fe457610fe361167e565b5b6000610ff284828501610e5d565b91505092915050565b600080604083850312156110125761101161167e565b5b600061102085828601610e5d565b925050602061103185828601610d87565b9150509250929050565b600080604083850312156110525761105161167e565b5b600061106085828601610e5d565b925050602061107185828601610df2565b9150509250929050565b6110848161150b565b82525050565b6110938161151d565b82525050565b6110a281611586565b82525050565b60006110b483856114ef565b93506110c1838584611598565b6110ca83611683565b840190509392505050565b60006110e18385611500565b93506110ee838584611598565b82840190509392505050565b6000611105826114d3565b61110f81856114de565b935061111f8185602086016115a7565b61112881611683565b840191505092915050565b60006111406024836114ef565b915061114b82611694565b604082019050919050565b60006111636026836114ef565b915061116e826116e3565b604082019050919050565b6000611186601f836114ef565b915061119182611732565b602082019050919050565b60006111a96024836114ef565b91506111b48261175b565b604082019050919050565b60006111cc6039836114ef565b91506111d7826117aa565b604082019050919050565b60006111ef6020836114ef565b91506111fa826117f9565b602082019050919050565b60006112126024836114ef565b915061121d82611822565b604082019050919050565b6000611235602b836114ef565b915061124082611871565b604082019050919050565b60006060830160008301516112636000860182611099565b506020830151848203602086015261127b82826110fa565b9150506040830151611290604086018261129b565b508091505092915050565b6112a48161157c565b82525050565b6112b38161157c565b82525050565b60006112c68284866110d5565b91508190509392505050565b60006020820190506112e7600083018461107b565b92915050565b6000606082019050611302600083018661107b565b61130f602083018561107b565b61131c60408301846112aa565b949350505050565b6000606082019050611339600083018761107b565b818103602083015261134c8185876110a8565b905061135b60408301846112aa565b95945050505050565b6000602082019050611379600083018461108a565b92915050565b6000602082019050818103600083015261139881611133565b9050919050565b600060208201905081810360008301526113b881611156565b9050919050565b600060208201905081810360008301526113d881611179565b9050919050565b600060208201905081810360008301526113f88161119c565b9050919050565b60006020820190508181036000830152611418816111bf565b9050919050565b60006020820190508181036000830152611438816111e2565b9050919050565b6000602082019050818103600083015261145881611205565b9050919050565b6000602082019050818103600083015261147881611228565b9050919050565b60006020820190508181036000830152611499818461124b565b905092915050565b60006040820190506114b660008301866112aa565b81810360208301526114c98184866110a8565b9050949350505050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60006115168261155c565b9050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6000819050611557826118c0565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061159182611549565b9050919050565b82818337600083830152505050565b60005b838110156115c55780820151818401526020810190506115aa565b838111156115d4576000848401525b50505050565b600060028204905060018216806115f257607f821691505b602082108114156116065761160561163b565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f496e697469616c205072696365206d757374206265206772656174657220746860008201527f616e203000000000000000000000000000000000000000000000000000000000602082015250565b7f4e65772073746174652073686f756c6420626520626967676572207468616e2060008201527f61637475616c0000000000000000000000000000000000000000000000000000602082015250565b7f64657374696e6174696f6e20616464726573732063616e6e6f74206265203000600082015250565b7f50726f64756374207769746820676976656e20696420616c726561647920657860008201527f6973747300000000000000000000000000000000000000000000000000000000602082015250565b7f56616c75652073656e7420666f72207468652070726f6475637420646f65732060008201527f6e6f74206d617463682070726f64756374277320707269636500000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f50726f64756374207769746820676976656e20696420646f6573206e6f74206560008201527f7869737400000000000000000000000000000000000000000000000000000000602082015250565b7f50726f647563742073686f756c64206265206372656174656420616e64206e6f60008201527f7420796574207061796564000000000000000000000000000000000000000000602082015250565b600481106118d1576118d061160c565b5b50565b6118dd8161150b565b81146118e857600080fd5b50565b600481106118f857600080fd5b50565b6119048161157c565b811461190f57600080fd5b5056fea264697066735822122083c236a71b19bd5fbf03cc52544006863f1cd7f8d46de038f80bcb8739b5361564736f6c63430008070033";

export class ProductFactory__factory extends ContractFactory {
  constructor(
    ...args: [signer: Signer] | ConstructorParameters<typeof ContractFactory>
  ) {
    if (args.length === 1) {
      super(_abi, _bytecode, args[0]);
    } else {
      super(...args);
    }
  }

  deploy(
    nftAddress: string,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ProductFactory> {
    return super.deploy(nftAddress, overrides || {}) as Promise<ProductFactory>;
  }
  getDeployTransaction(
    nftAddress: string,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(nftAddress, overrides || {});
  }
  attach(address: string): ProductFactory {
    return super.attach(address) as ProductFactory;
  }
  connect(signer: Signer): ProductFactory__factory {
    return super.connect(signer) as ProductFactory__factory;
  }
  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ProductFactoryInterface {
    return new utils.Interface(_abi) as ProductFactoryInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): ProductFactory {
    return new Contract(address, _abi, signerOrProvider) as ProductFactory;
  }
}

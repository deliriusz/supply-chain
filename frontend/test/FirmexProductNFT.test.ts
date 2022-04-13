import { ethers, waffle } from 'hardhat'
import { Signer } from 'ethers'
import chai, { expect } from 'chai'
import { FirmexProductNFT } from '../src/types/FirmexProductNFT'
import FirmexProductNFTArtifact from '../src/artifacts/src/contracts/FirmexProductNFT.sol/FirmexProductNFT.json'
import chaiAsPromised from 'chai-as-promised'
import { MockProvider, solidity } from 'ethereum-waffle'

chai.use(solidity) // solidiity matchers, e.g. expect().to.be.revertedWith("message")
chai.use(chaiAsPromised) //eventually

describe('FirmexProductNFT contract', () => {
   let accounts: Signer[];
   let owner: Signer;
   let signer1: Signer;
   let signer2: Signer;
   let signer3: Signer;
   let ownerAddress: string;
   let signer1Address: string;
   let signer2Address: string;
   let signer3Address: string;
   let firmexProductNFT: FirmexProductNFT

   beforeEach(async () => {
      accounts = new MockProvider({ ganacheOptions: { gasLimit: 100000000 } }).getWallets();
      [owner, signer1, signer2, signer3] = accounts;
      ownerAddress = await owner.getAddress()
      signer1Address = await signer1.getAddress()
      signer2Address = await signer2.getAddress()
      signer3Address = await signer3.getAddress()

      firmexProductNFT = (await waffle.deployContract(owner, FirmexProductNFTArtifact, ["http://example.com/"])) as FirmexProductNFT
   })

   it('should be named correctly', async () => {
      expect(firmexProductNFT.name()).eventually.eq('FirmexProductNFT')
      expect(firmexProductNFT.symbol()).eventually.eq('FXP')
   })

   it('can mint NFTs', async () => {
      await firmexProductNFT.safeMint(signer1Address, 'externalidx1', 1)
      await firmexProductNFT.safeMint(signer2Address, 'externalidx2', 2)
      await firmexProductNFT.safeMint(signer2Address, 'externalidx3', 3)

      expect(firmexProductNFT.balanceOf(signer1Address)).eventually.eq(1)
      expect(firmexProductNFT.balanceOf(signer2Address)).eventually.eq(2)
      expect(firmexProductNFT.balanceOf(signer3Address)).eventually.eq(0)
   })

   it('can send token', async () => {

   })

   it('returns expected URL for a token', async () => {
      await firmexProductNFT.safeMint(signer2Address, 'externalidx3', 3)

      expect("http://example.com/externalidx3").to.be.eq(await firmexProductNFT.tokenURI(3))
      await expect(firmexProductNFT.tokenURI(1)).to.be.revertedWith('ERC721URIStorage: URI query for nonexistent token')
   })
})

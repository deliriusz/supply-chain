import { ethers, waffle } from 'hardhat'
import { Signer } from 'ethers'
import chai, { expect } from 'chai'
import { FirmexProductNFT } from '../src/types/FirmexProductNFT'
import FirmexProductNFTArtifact from '../src/artifacts/src/contracts/FirmexProductNFT.sol/FirmexProductNFT.json'
import { ProductFactory } from '../src/types/ProductFactory'
import ProductFactoryArtifact from '../src/artifacts/src/contracts/ProductFactory.sol/ProductFactory.json'
import chaiAsPromised from 'chai-as-promised'
import { solidity, MockProvider } from 'ethereum-waffle'

chai.use(solidity) // solidiity matchers, e.g. expect().to.be.revertedWith("message")
chai.use(chaiAsPromised) //eventually

describe('ProductFactory contract', () => {
   let accounts: Signer[];
   let owner: Signer;
   let signer1: Signer;
   let signer2: Signer;
   let signer3: Signer;
   let ownerAddress: string;
   let signer1Address: string;
   let signer2Address: string;
   let signer3Address: string;
   let firmexProductNFT: FirmexProductNFT;
   let productFactory: ProductFactory;

   beforeEach(async () => {
      accounts = new MockProvider({ ganacheOptions: { gasLimit: 100000000 } }).getWallets();
      [owner, signer1, signer2, signer3] = accounts;
      ownerAddress = await owner.getAddress()
      signer1Address = await signer1.getAddress()
      signer2Address = await signer2.getAddress()
      signer3Address = await signer3.getAddress()

      firmexProductNFT = (await waffle.deployContract(owner, FirmexProductNFTArtifact)) as FirmexProductNFT
      productFactory = (await waffle.deployContract(owner, ProductFactoryArtifact, [firmexProductNFT.address])) as ProductFactory

      await firmexProductNFT.transferOwnership(productFactory.address);

      await productFactory.create('productX', 2000, 100, 'uriX');
   })

   it('transfers ownership correctly', async () => {
      expect(productFactory.address).to.be.eq(await firmexProductNFT.owner());
   })

   it('can create nft', async () => {
      await productFactory.create('product1', 1000, 1, 'uri1');
      await expect('safeMint').to.be.calledOnContractWith(firmexProductNFT, [productFactory.address, 'uri1', 1]);
      expect(productFactory.address).to.be.eq(await firmexProductNFT.ownerOf(1));
   })

   it('fails when non-owner tries to create a product', async () => {
      await expect(productFactory.connect(signer1).create('product1', 1000, 1, 'uri1')).to.be.revertedWith('Ownable: caller is not the owner');
   })

   it('fails when other signer tries to mint nft', async () => {
      await expect(firmexProductNFT.connect(owner).safeMint(signer1Address, 'externalidx1', 1)).to.be.revertedWith('Ownable: caller is not the owner');
      await expect(firmexProductNFT.connect(signer1).safeMint(signer1Address, 'externalidx1', 1)).to.be.revertedWith('Ownable: caller is not the owner');
   })

   it('fails when price is 0', async () => {
      await expect(productFactory.create('product1', 0, 1, 'uri1')).to.be.revertedWith('Initial Price must be greater than 0');
   })

   it('emits event on successful product creation', async () => {
      expect(productFactory.create('product1', 1000, 1, 'uri1')).to.emit(productFactory, 'ProductCreated').withArgs('product1', 1000, 1, 'uri1');
   })

   it('can change state of product', async () => {
      const product = await productFactory.getProduct(100);
      expect(0).to.be.eq(product.state);
      await productFactory.changeProductState(100, 1);
      expect(1).to.be.eq((await productFactory.getProduct(100)).state);
   })

   it('reverts when previous state is passed', async () => {
      await productFactory.changeProductState(100, 1);
      expect(1).to.be.eq((await productFactory.getProduct(100)).state);
      expect(productFactory.changeProductState(100, 0)).to.be.revertedWith('New state should be bigger than actual');
   })

   it('can get existing product', async () => {
      const product = await productFactory.getProduct(100);
      expect(product.state).to.be.eq(0);
      expect(product.name).to.be.eq('productX');
      expect(product.initialPrice).to.be.eq(2000);
   })

   it('reverts when getting non-existing product', async () => {
      expect(productFactory.getProduct(999)).to.be.revertedWith('Product with given id does not exist');
   })

   it('reverts when trying to renounce ownership', async () => {
      expect(productFactory.renounceOwnership()).to.be.reverted;
   })

   it('reverts when trying to change ownership', async () => {
      expect(productFactory.transferOwnership(signer1Address)).to.be.reverted;
   })

   it('fails when selling product with different price than expected', async () => {
      await productFactory.changeProductState(100, 1);
      expect(productFactory.sellProduct(100, signer1Address, { value: 1 })).to.be.revertedWith('Value sent for the product does not match product\'s price');
      expect(productFactory.sellProduct(100, signer1Address, { value: 9000 })).to.be.revertedWith('Value sent for the product does not match product\'s price');
   })

   it('fails when selling product which is not is state Created', async () => {
      expect(productFactory.sellProduct(100, signer1Address, { value: 2000 })).to.be.revertedWith('Product should be created and not yet payed');
      await productFactory.changeProductState(100, 3);
      expect(productFactory.sellProduct(100, signer1Address, { value: 2000 })).to.be.revertedWith('Product should be created and not yet payed');
   })

   it('fails when selling product which is already sold', async () => {
      await productFactory.changeProductState(100, 1);
      await productFactory.sellProduct(100, signer1Address, { value: 2000 });
      expect(productFactory.sellProduct(100, signer2Address, { value: 2000 })).to.be.revertedWith('Product should be created and not yet payed');
   })

   it('sells a product sucessfully when all requirements are met', async () => {
      const x = await productFactory.changeProductState(100, 1);
      await expect(productFactory.sellProduct(100, signer1Address, { value: 2000 })).to.emit(productFactory, 'ProductSold').withArgs(2000, 100, signer1Address);
      // function below won't work, as function is inherited from parent, and Waffle checks for function name, which does not exist there
      // expect('safeTransferFrom').to.be.calledOnContractWith(firmexProductNFT, [ownerAddress, signer1Address, 100]);

      expect(await firmexProductNFT.ownerOf(100)).to.be.eq(signer1Address);
   })
})

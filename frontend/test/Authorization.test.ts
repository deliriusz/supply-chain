import { ethers, waffle } from 'hardhat'
import { Signer } from 'ethers'
import chai, { expect } from 'chai'
import { Authorization } from '../src/types/Authorization'
import AuthorizationContract from '../src/artifacts/src/contracts/Authorization.sol/Authorization.json'
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
   let authorization: Authorization;

   let ROLE_ASSIGNMENT_DELAY_SECS = 30;
   let ROLE_ADMIN: string;
   let ROLE_DASHBOARD_VIEWER: string;

   beforeEach(async () => {
      accounts = new MockProvider({ ganacheOptions: { gasLimit: 100000000 } }).getWallets();
      [owner, signer1, signer2, signer3] = accounts;
      ownerAddress = await owner.getAddress()
      signer1Address = await signer1.getAddress()
      signer2Address = await signer2.getAddress()
      signer3Address = await signer3.getAddress()

      authorization = (await waffle.deployContract(owner, AuthorizationContract, [ROLE_ASSIGNMENT_DELAY_SECS])) as Authorization

      ROLE_ADMIN = await authorization.ROLE_ADMIN()
      ROLE_DASHBOARD_VIEWER = await authorization.ROLE_DASHBOARD_VIEWER()
   })

   it('assigns admin role instantly after deployment', async () => {
      expect(ownerAddress).to.be.eq(await authorization.owner());

      expect(await authorization.ROLE_ADMIN()).to.be.eq(await authorization.getUserRole(ownerAddress));
   })

   it('assigns roles to different users', async () => {
      // await expect('safeMint').to.be.calledOnContractWith(authorization, [productFactory.address, 'uri1', 1]);

      expect(await authorization.ROLE_ADMIN()).to.be.eq(await authorization.getUserRole(ownerAddress));
   })

   it('fails when non admin tries to change access', async () => {
      // await expect(productFactory.connect(signer1).create('product1', 1000, 1, 'uri1')).to.be.revertedWith('Ownable: caller is not the owner');
   })

   it('does nothing when the same role is reassigned to an account', async () => {
   })

   it('returns assigned role only if required amount of time passed', async () => {
   })

   it('emits an event on role assignment', async () => {
      expect(authorization.assignRole(signer1Address, ROLE_ADMIN)).to.emit(authorization, 'productcreated').withArgs(signer1Address, ROLE_ADMIN);
   })

   it('emits an event on role revocation', async () => {
      // expect(productfactory.create('product1', 1000, 1, 'uri1')).to.emit(productfactory, 'productcreated').withargs('product1', 1000, 1, 'uri1');
   })
})

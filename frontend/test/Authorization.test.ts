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

   let ROLE_ASSIGNMENT_DELAY_SECS = 5;
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

      console.log("USER ROLE=" + await authorization.getUserRole(ownerAddress))

      expect(ROLE_ADMIN).to.be.eq(await authorization.getUserRole(ownerAddress));
   })

   it('assigns roles to different users', async () => {
      await authorization.assignRole(signer1Address, ROLE_DASHBOARD_VIEWER);
      await expect('assignRole').to.be.calledOnContractWith(authorization, [signer1Address, ROLE_DASHBOARD_VIEWER]);

      await authorization.assignRole(signer2Address, ROLE_DASHBOARD_VIEWER);
      await expect('assignRole').to.be.calledOnContractWith(authorization, [signer2Address, ROLE_DASHBOARD_VIEWER]);

      setTimeout(async () => {
         expect(await authorization.getUserRole(signer1Address)).to.be.eq(ROLE_DASHBOARD_VIEWER)
         expect(await authorization.getUserRole(signer2Address)).to.be.eq(ROLE_DASHBOARD_VIEWER)
         console.log("INSIDE SETTIMEOUT")
      }, ROLE_ASSIGNMENT_DELAY_SECS * 1000)
   })

   it('fails when non admin tries to change access', async () => {
      // await expect(productFactory.connect(signer1).create('product1', 1000, 1, 'uri1')).to.be.revertedWith('Ownable: caller is not the owner');
   })

   it('does nothing when the same role is reassigned to an account', async () => {
   })

   it('returns assigned role only if required amount of time passed', async () => {
   })

   it('fails when trying to assign not existing role', async () => {
   })

   it('emits an event on role assignment', async () => {
      expect(authorization.assignRole(signer1Address, ROLE_ADMIN)).to.emit(authorization, 'RoleAssigned').withArgs(signer1Address, ROLE_ADMIN);
   })

   it('emits an event on role revocation', async () => {
      expect(authorization.revokeRole(signer1Address)).to.emit(authorization, 'RoleRevoked').withArgs(signer1Address);
   })
})

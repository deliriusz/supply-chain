import { waffle } from 'hardhat'
import { Signer } from 'ethers'
import chai, { expect } from 'chai'
import { Authorization } from '../src/types/Authorization'
import AuthorizationContract from '../src/artifacts/src/contracts/Authorization.sol/Authorization.json'
import chaiAsPromised from 'chai-as-promised'
import { solidity, MockProvider } from 'ethereum-waffle'

chai.use(solidity) // solidiity matchers, e.g. expect().to.be.revertedWith("message")
chai.use(chaiAsPromised) //eventually

describe('Authorization contract', () => {
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
   const ROLE_UNAUTHORIZED = '0x00000000'
   let provider: MockProvider;

   beforeEach(async () => {
      provider = new MockProvider({ ganacheOptions: { gasLimit: 100000000 } })

      accounts = provider.getWallets();
      [owner, signer1, signer2, signer3] = accounts
      ownerAddress = await owner.getAddress()
      signer1Address = await signer1.getAddress()
      signer2Address = await signer2.getAddress()
      signer3Address = await signer3.getAddress()

      authorization = (await waffle.deployContract(owner, AuthorizationContract, [ROLE_ASSIGNMENT_DELAY_SECS])) as Authorization

      ROLE_ADMIN = await authorization.ROLE_ADMIN()
      ROLE_DASHBOARD_VIEWER = await authorization.ROLE_DASHBOARD_VIEWER()
   })

   it('assigns admin role instantly after deployment', async () => {
      expect(ownerAddress).to.be.eq(await authorization.owner())

      expect(ROLE_ADMIN).to.be.eq(await authorization.getUserRole(ownerAddress))
   })

   it('delays role assignment for required time', async () => {
      await authorization.assignRole(signer1Address, ROLE_ADMIN)
      await expect('assignRole').to.be.calledOnContractWith(authorization, [signer1Address, ROLE_ADMIN])

      await expect(await authorization.getUserRole(signer1Address)).to.be.equal(ROLE_UNAUTHORIZED)
      // await expect(await authorization.connect(signer1Address)
      //    .assignRole(signer2Address, ROLE_ADMIN)).to.be.revertedWith('Authorization: the address does not have an active role assigned yet.')

      const blockNumBefore = await provider.getBlockNumber();
      const blockBefore = await provider.getBlock(blockNumBefore);
      const timestampBefore = blockBefore.timestamp;

      await provider.send('evm_increaseTime', [ROLE_ASSIGNMENT_DELAY_SECS]);
      await provider.send('evm_mine', []);

      const blockNumAfter = await provider.getBlockNumber();
      const blockAfter = await provider.getBlock(blockNumAfter);
      const timestampAfter = blockAfter.timestamp;

      expect(blockNumAfter).to.be.equal(blockNumBefore + 1);
      //testing this on npx ganache node will fail by few seconds as the time passes by constantly
      expect(timestampAfter).to.be.equal(timestampBefore + ROLE_ASSIGNMENT_DELAY_SECS)

      expect(await authorization.getUserRole(signer1Address)).to.be.eq(ROLE_ADMIN)
   })

   it('assigns roles to different users', async () => {
      await authorization.assignRole(signer1Address, ROLE_DASHBOARD_VIEWER);
      await expect('assignRole').to.be.calledOnContractWith(authorization, [signer1Address, ROLE_DASHBOARD_VIEWER])

      await authorization.assignRole(signer2Address, ROLE_DASHBOARD_VIEWER);
      await expect('assignRole').to.be.calledOnContractWith(authorization, [signer2Address, ROLE_DASHBOARD_VIEWER])

      await provider.send('evm_increaseTime', [ROLE_ASSIGNMENT_DELAY_SECS]);
      await provider.send('evm_mine', []);

      expect(await authorization.getUserRole(signer1Address)).to.be.eq(ROLE_DASHBOARD_VIEWER)
      expect(await authorization.getUserRole(signer2Address)).to.be.eq(ROLE_DASHBOARD_VIEWER)
   })

   it('fails when non admin tries to change access', async () => {
      await expect(authorization.connect(signer1)
         .assignRole(signer3Address, ROLE_ADMIN))
         .to.be.revertedWith('Authorization: the address does not have a required role.');
   })

   it('does nothing when the same role is reassigned to an account', async () => {
      await authorization.assignRole(signer1Address, ROLE_DASHBOARD_VIEWER);

      const roleActivationTime = await authorization._accoutRoleAssignmentTime(signer1Address)

      await provider.send('evm_increaseTime', [ROLE_ASSIGNMENT_DELAY_SECS]);
      await provider.send('evm_mine', []);

      expect(await authorization.getUserRole(signer1Address)).to.be.eq(ROLE_DASHBOARD_VIEWER)

      await authorization.assignRole(signer1Address, ROLE_DASHBOARD_VIEWER);
      await provider.send('evm_increaseTime', [ROLE_ASSIGNMENT_DELAY_SECS]);
      await provider.send('evm_mine', []);

      expect(await authorization.getUserRole(signer1Address)).to.be.eq(ROLE_DASHBOARD_VIEWER)
      expect(await authorization._accoutRoleAssignmentTime(signer1Address)).to.be.eq(roleActivationTime)
   })

   it('emits an event on role assignment', async () => {
      expect(authorization.assignRole(signer1Address, ROLE_ADMIN)).to.emit(authorization, 'RoleAssigned').withArgs(signer1Address, ROLE_ADMIN);
   })

   it('emits an event on role revocation', async () => {
      expect(authorization.revokeRole(signer1Address)).to.emit(authorization, 'RoleRevoked').withArgs(signer1Address);
   })
})

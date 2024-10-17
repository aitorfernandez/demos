const { expect } = require('chai')
const { expectRevert } = require('@openzeppelin/test-helpers')

const FiatToken = artifacts.require('FiatToken')

contract('FiatToken', ([owner, alice, bob, carol, dan, ...rest]) => {
  let token

  beforeEach(async () => {
    token = await FiatToken.new(alice, {
      from: owner,
    })
  })

  it('supply', async () => {
    expect((await token.getSupply()).toNumber()).to.equal(0)
  })

  it('setMasterMinter', async () => {
    const signers = await ethers.getSigners()

    await token.setMasterMinter(bob)

    expect(await token.masterMinter()).to.equal(bob)
  })

  it('non owner cannot set masterMinter', async () => {
    await expectRevert(
      token.setMasterMinter(carol, { from: dan }),
      'Ownable: caller is not the owner',
    )
  })
})

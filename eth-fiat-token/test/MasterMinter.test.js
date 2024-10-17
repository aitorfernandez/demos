const { expect } = require('chai')
const { expectRevert } = require('@openzeppelin/test-helpers')

const FiatToken = artifacts.require('FiatToken')
const MasterMinter = artifacts.require('MasterMinter')

contract('MasterMinter', ([owner, mintOwner, controller, minter, alice, bob, ...rest]) => {
  let token, masterMinter

  beforeEach(async () => {
    token = await FiatToken.new(alice, {
      from: owner,
    })
    masterMinter = await MasterMinter.new(token.address, {
      from: mintOwner,
    })
    await token.setMasterMinter(masterMinter.address)
  })

  it("should mint through mint controller", async () => {
    const signers = await ethers.getSigners()
    const amount = 999;

    await masterMinter.configureController(controller, minter, {
      from: mintOwner,
    })

    await masterMinter.configureMinter(amount, {
      // Controller manages minter
      from: controller,
    })

    await token.mint(alice, amount, {
      from: minter,
    })

    expect((await token.getSupply()).toNumber()).to.equal(amount)
  })
})

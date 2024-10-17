const { ethers } = require('hardhat')

async function main() {
  const throwaway = '0x0000000000000000000000000000000000000001'

  // const accounts = await ethers.getSigners()

  const FiatToken = await ethers.getContractFactory('FiatToken')
  const MasterMinter = await ethers.getContractFactory('MasterMinter')

  const token = await FiatToken.deploy(throwaway)
  const masterMinter = await MasterMinter.deploy(token.address)

  console.log(`token address ${token.address}`)
  console.log(`masterMinter address ${masterMinter.address}`)

  await token.setMasterMinter(masterMinter.address)

  console.log(await token.getSupply())

  console.log(`token owner ${await token.owner()}`)
  console.log(`masterMinter owner ${await masterMinter.owner()}`)
}

main()

require('@nomiclabs/hardhat-waffle')
require('@nomiclabs/hardhat-etherscan')
require('dotenv').config()

const privateKey1 = '0xdf57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e'

module.exports = {
  solidity: '0.8.4',
  networks: {
    localhost: {
      accounts: [privateKey1],
    },
    rinkeby: {
      url: process.env.RINKEBY_URL,
      accounts: [`${process.env.RINKEBY_PRIVATE_KEY}`],
    }
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_KEY
  }
}

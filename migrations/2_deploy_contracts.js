var EthUSD = artifacts.require("./EthUSD.sol");
var MockEtherPriceOracle = artifacts.require("./MockEtherPriceOracle.sol");
const priceOraclesByNework = {
  mainnet: "0xf5c600cda3b7289b2863872b23084527fb4c6107",
  ropsten: "0x9f7f9f38825012623cae7982ead15cc0571adcde",
  rinkeby: "0xdcd37d397ab4eeb86ab8aed69fdc551f7c0bc77b",
  kovan: "0x9e803018893dc7ec24c19c8779a4444ba49c44e4",
};

module.exports = function(deployer, network) {
  if (network == "development") {
    deployer.deploy(MockEtherPriceOracle).then(() => {
      return deployer.deploy(EthUSD, MockEtherPriceOracle.address);
    });
  } else {
    deployer.deploy(EthUSD, priceOraclesByNework[network]);
  }
};

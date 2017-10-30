require('dotenv').config();
const _ = require("lodash");
const WalletProvider = require("truffle-wallet-provider");
const Wallet = require('ethereumjs-wallet');

const networks = ["rinkeby", "kovan", "ropsten", "mainnet"];

const infuraNetworks = _.fromPairs(_.compact(networks.map((network) => {
  var envVarName = `${network.toUpperCase()}_PRIVATE_KEY`
  var privateKeyHex = process.env[envVarName];

  if(privateKeyHex) {
    var privateKey = new Buffer(process.env[envVarName], "hex")
    var wallet = Wallet.fromPrivateKey(privateKey);
    var provider = new WalletProvider(wallet, `https://${network}.infura.io/`);

    return [
      network,
      {
        host: "localhost",
        port: 8545,
        network_id: "*",
        gas: 4612388,
        provider,
      }
    ];
  }
})));

module.exports = {
  networks: {
    development: {
      host: "localhost",
      port: 8545,
      network_id: "*"
    },
    ...infuraNetworks,
  }
};

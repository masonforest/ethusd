__Disclaimer__

_This is a Proof of Concept only. Please don't deposit more than you're willing
to loose._

EthUSD
====

EthUSD is an Ether backed stablecoin pegged to the USD.

Users can issue EthUSD by paying Ether into the contract. At a later date they
can withdraw Ether. If the price of Ether doubles and they withdraw they'll
receive half as much Ether as they deposited (The same amount in USD). The remaining ether is stored in the contract to cover losses when the price
falls.

When a user has deposited funds and the price falls two different things can happen. If the contract has the funds to pay the user it will do so (Eg. If you bought in 2 Ether and the price halved you'd get back 4 Ethers). If the contract doesn't have enough to fully cover the losses it will pay out proportionally to what it has left (eg If you bought in 2 Ethers and the price halved but the contract only had 2 Ethers it'd give you back 2 Ethers).

EthUSD is simpler than other stablecoins because it doesn't try to solve the bad collateral problem.
If the price of the collateral (Ether) starts to fall the stablecoin will fall in value at at the same rate of depreciation.


EtherUSD is deployed at the following addresses:

|  Network  |   Address    |
------------|---------------
|Main Network | [0x914558f943d117660cc576f567240fbb433698c8](https://etherscan.io/address/0x914558f943d117660cc576f567240fbb433698c8) |
Ropsten  | [0x7a406e27b9f4feb58bf23e4fe36497a74882e2f8](https://ropsten.etherscan.io/address/0x7a406e27b9f4feb58bf23e4fe36497a74882e2f8) |
Rinkeby  | [0x39c688b428c51fdccace92517bec295440e55343](https://rinkeby.etherscan.io/address/0x39c688b428c51fdccace92517bec295440e55343) |
Kovan    | [0x33fb4e08ecd73e6d851c4234d87680eb7b641868](https://kovan.etherscan.io/address/0x33fb4e08ecd73e6d851c4234d87680eb7b641868) |

Usage
-----

To issue tokens yourself:

1. Install [MetaMask](https://metamask.io/) or another Ethereum enabled browser.

2. Visit https://wallet.ethereum.org/
3. Click on "Contracts"
4. Click on "Watch Token"
5. Enter the contract address found in the table below.
6. Click "Watch Contract"
7. Enter the contract address found in the table above.
8. Copy and paste the [ABI](https://raw.githubusercontent.com/masonforest/ethusd/mf-convert-to-truffle/EthUSD.json).
9. Click on "EthUSD"
10. Under "Select function" choose "issue"
11. Enter the amount of EthUSD you'd like to issue
12. Click on "Execute"
13. When the confirmation dialog appears click on "Submit"

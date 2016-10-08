import "EthUSD.sol";

contract TestEthUSD is EthUSD {
  function TestEthUSD(string _name, string _symbol) EthUSD(_name,_symbol) {
  }


  function setExchangeRate(uint256 _amount){
    exchangeRate = _amount;
  }

  function setBalance(address _address, uint256 _amount){
    balances[_address] = _amount;
  }

   // these functions fail on the golang backend simulator so they are
   // overridden
  function usingOraclize(){}
  modifier oraclizeAPI {}
  function oraclize_setProof(byte proofP) oraclizeAPI internal {}
}

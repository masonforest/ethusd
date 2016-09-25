import "EthUSD.sol";

contract TestEthUSD is EthUSD {

	string public name;

  function TestEthUSD(string _name){
		name = _name;
  }

	function usingOraclize(){
    // noop
    // this function fails on the golang simulator so it's
    // overridden
	}

  function setExchangeRate(uint256 _amount){
    exchangeRate = _amount;
  }
}

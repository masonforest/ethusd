//import "ERC20Token.sol";
import "oraclizeApi.sol";

contract EthUSD is usingOraclize {
	string public name;
  uint public exchangeRate;
	event Purchase(address indexed from, uint256 value);
	mapping (address => uint256) balances;

  function EthUSD(){
		oraclize_setProof(proofType_TLSNotary | proofStorage_IPFS);
		updateExchangeRate(0);
  }


	function balanceOf(address _owner) constant returns (uint256 balance) {
		return balances[_owner];
   }

	function purchase() returns (bool success) {
		if (msg.value > 0) {
				balances[msg.sender] += msg.value * exchangeRate;
				return true;
		} else {
			return false;
		}
	}

  function __callback(bytes32 myid, string result, bytes proof) {
    if (msg.sender != oraclize_cbAddress()) throw;
    exchangeRate = parseInt(result, 2); // save it as $ cents
  }

  function updateExchangeRate(uint delay){
    oraclize_query(delay, "URL",
      "https://poloniex.com/public?command=returnTicker).USDT_ETH.last");
  }
}

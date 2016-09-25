//import "ERC20Token.sol";
import "oraclizeApi.sol";

contract EthUSD is usingOraclize {
  uint public exchangeRate;
	bool public called;
	mapping (address => uint256) balances;

  function EthUSD(){
		//oraclize_setProof(proofType_TLSNotary | proofStorage_IPFS);
    exchangeRate = 100;
	//	updateExchangeRate(0);
  }

	function buy() returns (bool success) {
		if (msg.value > 0) {
				balances[msg.sender] += msg.value;
				return true;
		} else {
			return false;
		}
	}

  //function __callback(bytes32 myid, string result, bytes proof) {
  //  if (msg.sender != oraclize_cbAddress()) throw;
  //  exchangeRate = parseInt(result, 2); // save it as $ cents
	//	called = true;
  //}

  //function updateExchangeRate(uint delay){
  //  oraclize_query(delay, "URL",
  //    "https://poloniex.com/public?command=returnTicker).USDT_ETH.last");
  //}
}

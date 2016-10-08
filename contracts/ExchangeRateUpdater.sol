import "oraclizeApi.sol";

contract ExchangeRateUpdater is usingOraclize {
  uint public exchangeRate;

  function ExchangeRateUpdater(){
    oraclize_setProof(proofType_TLSNotary | proofStorage_IPFS);
    updateExchangeRate(0);
  }

  function __callback(bytes32 myid, string result, bytes proof) {
    if (msg.sender != oraclize_cbAddress()) throw;
    exchangeRate = parseInt(result, 2); // save it as $ cents
  }

  function updateExchangeRate(uint delay){
    oraclize_query(delay, "URL", "json(https://poloniex.com/public?command=returnTicker).USDT_ETH.last");
  }
}

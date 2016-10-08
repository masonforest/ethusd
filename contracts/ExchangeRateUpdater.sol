import "oraclizeApi.sol";

contract ExchangeRateUpdater is usingOraclize {
  uint public exchangeRateCap;
  uint public exchangeRate;

  function ExchangeRateUpdater(){
    oraclize_setProof(proofType_TLSNotary | proofStorage_IPFS);
    updateExchangeRate(0);
  }

  function __callback(bytes32 myid, string result, bytes proof) {
    if (msg.sender != oraclize_cbAddress()) throw;

    // If exchangeRate has not yet been set this is
    // the first run and we set should exchangeRateCap aswell
    if (exchangeRate == 0 ) {
      exchangeRate = parseInt(result, 2);
      exchangeRateCap = exchangeRate;
    } else {
      exchangeRate = parseInt(result, 2);
    }
  }

  function updateExchangeRate(uint delay){
    oraclize_query(delay, "URL", "json(https://poloniex.com/public?command=returnTicker).USDT_ETH.last");
  }
}

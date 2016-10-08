import "ExchangeRateUpdater.sol";
import "StandardToken.sol";

contract EthUSD is ExchangeRateUpdater, StandardToken {
  string public name;
  string public symbol;
  mapping (address => uint256) public balances;

  event Purchase(address indexed from, uint256 value);
  event Withdraw(address indexed from, uint256 value);

  function EthUSD(string _name, string _symbol) {
    name = _name;
    symbol = _symbol;
  }

  function balanceOf(address _owner) constant returns (uint256 balance) {
    return balances[_owner];
  }

  function purchase() returns (bool success) {
    if (msg.value > 0) {
      Purchase(msg.sender, msg.value);
      uint amount = (msg.value * cappedExhangeRate()) / 1 ether;
      balances[msg.sender] += amount;
      Purchase(msg.sender, amount);
      return true;
    } else {
      return false;
    }
  }

  function withdraw(uint amountInCents) returns (bool success) {
    if(balances[msg.sender] > amountInCents){
      balances[msg.sender] -= amountInCents;
      uint amountInWei = amountInCents / exchangeRate;
      msg.sender.send(amountInWei);
      Purchase(msg.sender, amountInWei);
      return true;
    } else {
      return false;
    }
  }

  function cappedExhangeRate() private returns (uint256 rate) {
    if (exchangeRate > exchangeRateCap) {
      return exchangeRateCap;
    } else {
      return exchangeRate;
    }
  }
}

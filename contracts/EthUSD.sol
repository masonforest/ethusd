import "ExchangeRateUpdater.sol";

contract EthUSD is ExchangeRateUpdater {
	string public name;
	mapping (address => uint256) balances;

	event Purchase(address indexed from, uint256 value);
	event Withdraw(address indexed from, uint256 value);

	function balanceOf(address _owner) constant returns (uint256 balance) {
		return balances[_owner];
  }

	function purchase() returns (bool success) {
		if (msg.value > 0) {
				uint amount = msg.value * exchangeRate;
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
}

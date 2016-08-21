contract EthUSD {
    string public name;
    string public symbol;
    uint8 public decimals;

    mapping (address => uint256) public balanceOf;
    mapping (address => mapping (address => uint)) public allowance;
    mapping (address => mapping (address => uint)) public spentAllowance;

    event Transfer(address indexed from, address indexed to, uint256 value);

    function EthUSD(uint256 initialSupply, string tokenName) {
        balanceOf[msg.sender] = initialSupply;
        name = tokenName;
    }

    function transfer(address _to, uint256 _value) {
        if (balanceOf[msg.sender] < _value) throw;
        if (balanceOf[_to] + _value < balanceOf[_to]) throw;
        balanceOf[msg.sender] -= _value;
        balanceOf[_to] += _value;
        Transfer(msg.sender, _to, _value);
    }

    function () {
        throw;
    }
}

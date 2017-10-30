pragma solidity ^0.4.15;

import "./EtherPriceOracleInterface.sol";

contract MockEtherPriceOracle is EtherPriceOracleInterface {
    uint public price_;

    function setPrice(uint price) public {
      price_ = price;
    }
    function price() public constant returns (uint){
      return price_;
    }
}

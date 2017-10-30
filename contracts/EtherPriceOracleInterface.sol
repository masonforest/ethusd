pragma solidity ^0.4.15;

contract EtherPriceOracleInterface {
  function price() public constant returns (uint);
}

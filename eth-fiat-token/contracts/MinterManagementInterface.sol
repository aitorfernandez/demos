pragma solidity ^0.8.0;

interface MinterManagementInterface {
  function configureMinter(address _minter, uint256 _allowance)
    external returns (bool);
}

pragma solidity ^0.8.0;

import "./Controller.sol";
import "./MinterManagementInterface.sol";


contract MinterManager is Controller {
  MinterManagementInterface internal minterManager;

  function configureMinter(uint256 _allowance)
    public onlyController returns (bool)
  {
    address minter = controllers[msg.sender];

    return minterManager.configureMinter(minter, _allowance);
  }
}

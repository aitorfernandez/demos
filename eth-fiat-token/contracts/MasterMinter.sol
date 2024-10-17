pragma solidity ^0.8.0;

import "./MinterManager.sol";


contract MasterMinter is MinterManager {
  constructor(address _minterManager) public {
    minterManager = MinterManagementInterface(_minterManager);
  }
}

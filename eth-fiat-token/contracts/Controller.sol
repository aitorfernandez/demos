pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";


contract Controller is Ownable {
  mapping(address => address) internal controllers;

  modifier onlyController() {
    require(
      controllers[msg.sender] != address(0),
      "The value of controllers[msg.sender] must be non-zero"
    );
    _;
  }

  function configureController(address _controller, address _minter)
    public onlyOwner
  {
    controllers[_controller] = _minter;
  }
}

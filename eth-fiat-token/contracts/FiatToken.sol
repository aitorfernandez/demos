pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract FiatToken is ERC20, Ownable {
  uint256 internal _totalSupply = 0;
  address public masterMinter;

  mapping(address => bool) internal _minters;
  mapping(address => uint256) internal _minterAllowed;

  modifier onlyMasterMinter() {
    require(msg.sender == masterMinter, "[FiatToken] caller is not the masterMinter");
    _;
  }

  modifier onlyMinters() {
    require(_minters[msg.sender], "[FiatToken] Caller is not a minter");
    _;
  }

  constructor(address _masterMinter) ERC20("AIT", "AIT") {
    masterMinter = _masterMinter;
  }

  function mint(address _to, uint256 _amount)
    external onlyMinters
  {
    require(_amount > 0, "[FiatToken] mint amount not greater than 0");
    _totalSupply = _totalSupply + _amount;
  }

  function getSupply()
    external view returns (uint256)
  {
    return _totalSupply;
  }

  function addMinter(address _minter)
    external onlyMasterMinter
  {
    _minters[_minter] = true;
  }

  function setMasterMinter(address _address)
    external onlyOwner
  {
    masterMinter = _address;
  }

  function configureMinter(address _minter, uint256 _allowance)
    external onlyMasterMinter returns (bool)
  {
    _minters[_minter] = true;
    _minterAllowed[_minter] = _allowance;
    return true;
  }
}

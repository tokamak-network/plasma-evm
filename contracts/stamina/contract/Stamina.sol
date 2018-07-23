pragma solidity ^0.4.24;

contract Stamina {
  /**
   * Internal States
   */
  // delegatee of `from` account
  // `from` => `delegatee`
  mapping (address => address) _delegatee;

  // Stamina balance of delegatee
  // `delegatee` => `balance`
  mapping (address => uint) _stamina;

  // total deposit of delegatee
  // `delegatee` => `total deposit`
  mapping (address => uint) _total_deposit;

  // deposit of delegatee
  // `depositor` => `delegatee` => `deposit`
  mapping (address => mapping (address => uint)) _deposit;

  uint public t = 0xdead;

  bool public initialized;

  /**
   * Public States
   */
  uint public MIN_DEPOSIT;

  /**
   * Modifiers
   */
  modifier onlyChain() {
    // TODO: uncomment below
    // require(msg.sender == address(0));
    _;
  }

  /**
   * Events
   */
  event Deposited(address indexed depositor, address indexed delegatee, uint amount);
  event Withdrawn(address indexed depositor, address indexed delegatee, uint amount);
  event DelegateeChanged(address delegater, address oldDelegatee, address newDelegatee);

  /**
   * Init
   */
  function init(uint minDeposit) external {
    require(!initialized);

    MIN_DEPOSIT = minDeposit;

    initialized = true;
  }

  /**
   * Getters
   */
  function getDelegatee(address delegater) public view returns (address) {
    return _delegatee[delegater];
  }

  function getStamina(address addr) public view returns (uint) {
    return _stamina[addr];
  }

  function getTotalDeposit(address delegatee) public view returns (uint) {
    return _total_deposit[delegatee];
  }

  function getDeposit(address depositor, address delegatee) public view returns (uint) {
    return _deposit[depositor][delegatee];
  }

  /**
   * Setters and External functions
   */
  /// @notice set `msg.sender` as delegatee of `delegater`
  function setDelegatee(address delegater) external returns (bool) {
    address oldDelegatee = _delegatee[delegater];

    _delegatee[delegater] = msg.sender;

    emit DelegateeChanged(delegater, oldDelegatee, msg.sender);
    return true;
  }

  /// @notice deposit Ether to delegatee
  function deposit(address delegatee) external payable returns (bool) {
    require(msg.value >= MIN_DEPOSIT);

    uint dTotalDeposit = _total_deposit[delegatee];
    uint fDeposit = _deposit[msg.sender][delegatee];

    require(dTotalDeposit + msg.value > dTotalDeposit);
    require(fDeposit + msg.value > fDeposit);

    _total_deposit[delegatee] = dTotalDeposit + msg.value;
    _deposit[msg.sender][delegatee] = fDeposit + msg.value;

    emit Deposited(msg.sender, delegatee, msg.value);
    return true;
  }

  /// @notice request to withdraw Ether from delegatee. it store Ether to Escrow contract.
  ///         later `withdrawPayments` transfers Ether from Escrow to the depositor
  function withdraw(address delegatee, uint amount) external returns (bool) {
    uint dTotalDeposit = _total_deposit[delegatee];
    uint fDeposit = _deposit[msg.sender][delegatee];

    require(dTotalDeposit - amount < dTotalDeposit);
    require(fDeposit - amount < fDeposit);

    _total_deposit[delegatee] = dTotalDeposit - amount;
    _deposit[msg.sender][delegatee] = fDeposit - amount;

    msg.sender.transfer(amount);

    emit Withdrawn(msg.sender, delegatee, amount);
    return true;
  }

  /// @notice reset stamina up to total deposit of delegatee
  function resetStamina(address delegatee) external onlyChain {
    _stamina[delegatee] = _total_deposit[delegatee];
  }

  /// @notice add stamina of delegatee. The upper bound of stamina is total deposit of delegatee.
  function addStamina(address delegatee, uint amount) external onlyChain returns (bool) {
    uint dTotalDeposit = _total_deposit[delegatee];
    uint dBalance = _stamina[delegatee];

    require(dBalance + amount > dBalance);
    uint targetBalance = dBalance + amount;

    if (targetBalance > dTotalDeposit) _stamina[delegatee] = dTotalDeposit;
    else _stamina[delegatee] = targetBalance;

    return true;
  }

  /// @notice subtracte stamina of delegatee.
  function subtractStamina(address delegatee, uint amount) external onlyChain returns (bool) {
    uint dBalance = _stamina[delegatee];

    require(dBalance - amount < dBalance);
    _stamina[delegatee] = dBalance - amount;
    return true;
  }
}

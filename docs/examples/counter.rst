.. _example-counter:

*******
Counter
*******

`BaseCounter <https://github.com/Onther-Tech/requestable-contract-examples/blob/master/counter/solidity/BaseCounter.sol>`_
~~~~~~~~~~~~

BaseCounter is a simple base counter contract. Other requestable counter contracts inherit this base contract and implement requestable interface.

::

  pragma solidity ^0.4.24;


  contract BaseCounter {
    uint n;

    event Counted(uint _n);

    function count() external {
      n++;
      emit Counted(n);
    }

    function getCount() external view returns (uint) {
      return n;
    }
  }


`SimpleCounter <https://github.com/Onther-Tech/requestable-contract-examples/blob/master/counter/solidity/SimpleCounter.sol>`_
~~~~~~~~~~~~

When there is a counter contract that can be increased by anyone, the most intuitive enter is to reduce the number in the root chain and increase the value by that much in the child chain. Exit, in contrast, reduces the value first in the child chain and increase the value in the root chain. SimpleCounter that implements this is illustrated as follows.

.. image:: /res/simple-counter.png

A yellow box means that the ``counter()`` has increased the status variable ``n`` by 1, a red box means entering the request changes ``n``, and a green box means exiting the request changes ``n``.


::

  pragma solidity ^0.4.24;

  import {SimpleDecode} from "../lib/SimpleDecode.sol";
  import {RequestableI} from "../lib/RequestableI.sol";
  import {BaseCounter} from "./BaseCounter.sol";
  import {SafeMath} from "openzeppelin-solidity/contracts/math/SafeMath.sol";


  /// @notice A request can decrease `n`. Is it right to decrease the count?
  contract SimpleCounter is BaseCounter, RequestableI {
    // SimpleDecode library to decode trieValue.
    using SimpleDecode for bytes;
    using SafeMath for *;

    // trie key for state variable `n`.
    bytes32 constant public TRIE_KEY_N = 0x00;

    // address of RootChain contract.
    address public rootchain;

    mapping (uint => bool) appliedRequests;

    constructor(address _rootchain) {
      rootchain = _rootchain;
    }

    function applyRequestInRootChain(
      bool isExit,
      uint256 requestId,
      address requestor,
      bytes32 trieKey,
      bytes trieValue
    ) external returns (bool success) {
      require(!appliedRequests[requestId]);
      require(msg.sender == rootchain);

      // only accept request for `n`.
      require(trieKey == TRIE_KEY_N);

      if (isExit) {
        n = n.add(trieValue.toUint());
      } else {
        n = n.sub(trieValue.toUint());
      }

      appliedRequests[requestId] = true;
    }

    function applyRequestInChildChain(
      bool isExit,
      uint256 requestId,
      address requestor,
      bytes32 trieKey,
      bytes trieValue
    ) external returns (bool success) {
      require(!appliedRequests[requestId]);
      require(msg.sender == address(0));

      // only accept request for `n`.
      require(trieKey == TRIE_KEY_N);

      if (isExit) {
        n = n.sub(trieValue.toUint());
      } else {
        n = n.add(trieValue.toUint());
      }

      appliedRequests[requestId] = true;
    }
  }

Let's read the code one by one.

.. TODO: describe source code.

However, SimpleCounter may decrease with variable ``n`` due to enter and exit. If this is not desired, you can implement counter contract as below.


`FreezableCounter <https://github.com/Onther-Tech/requestable-contract-examples/blob/master/counter/solidity/FreezableCounter.sol>`_
~~~~~~~~~~~~

Enter and exit can be applied after freezing the contracts in each chain. FreezableCounter can be avoided if the number decreases through the request method after freezing.

.. image:: /res/freezable-counter.png


::

  pragma solidity ^0.4.24;

  ...


  /// @notice Both contract may be frozen at the same time. Is it right?
  contract FreezableCounter is BaseCounter, RequestableI {
    ...

    // freeze counter before make request.
    bool public frozen;

    constructor(address _rootchain) {
      rootchain = _rootchain;

      // Counter in child chain is frozen at first.
      if (_rootchain == address(0)) {
        frozen = true;
      }
    }

    function freeze() external returns (bool success) {
      frozen = true;
      return true;
    }

    function applyRequestInRootChain(
      bool isExit,
      uint256 requestId,
      address requestor,
      bytes32 trieKey,
      bytes trieValue
    ) external returns (bool success) {
      ...
      require(frozen);

      ...

      if (isExit) {
        frozen = false;
        n = trieValue.toUint();
      } else {
        require(n == trieValue.toUint());
      }

      ...
    }

    function applyRequestInChildChain(
      bool isExit,
      uint256 requestId,
      address requestor,
      bytes32 trieKey,
      bytes trieValue
    ) external returns (bool success) {
      ...
      require(frozen);

      ...

      if (isExit) {
        require(n == trieValue.toUint());
      } else {
        n = trieValue.toUint();
        frozen = false;
      }

      ...
    }
  }


However, the challenge period exists until exit is applied in root chain, for this freeze counter, all counters in each chain are frozen before the end of this challenge period. The enter is relatively short, but both are frozen. Therefore, to prevent this, the state variable used for enter and the state variable used for exit must be different.


`TrackableCounter <https://github.com/Onther-Tech/requestable-contract-examples/blob/master/counter/solidity/TrackableCounter.sol>`_
~~~~~~~~~~~~

TrackableCounter checks whether enter and exit is possible through a separate state variable ``requestableN`` in enter in the root chain and exit in child chain, reduces the value, and increases ``n`` in exit in the root chain and enter in the child chain. Both operations can prevent the reduction of ``n`` and apply only the correct enter and exit.

.. image:: /res/trackable-counter.png

::

  pragma solidity ^0.4.24;

  ...

  contract TrackableCounter is BaseCounter, RequestableI {
    ...

    // previous count before enter request in root chain and exit request in child chain.
    uint public requestableN;

    ...

    /// @dev override BaseCounter.count function.
    function count() external {
      requestableN++;
      n++;
      emit Counted(n);
    }

    function applyRequestInRootChain(
      bool isExit,
      uint256 requestId,
      address requestor,
      bytes32 trieKey,
      bytes trieValue
    ) external returns (bool success) {
      ...

      uint _n = trieValue.toUint()
      if (isExit) {
        n = n.add(_n);
      } else {
        requestableN = requestableN.sub(_n);
      }

      ...
    }

    function applyRequestInChildChain(
      bool isExit,
      uint256 requestId,
      address requestor,
      bytes32 trieKey,
      bytes trieValue
    ) external returns (bool success) {
      ...

      if (isExit) {
        requestableN = requestableN.sub(_n);
      } else {
        n = n.add(_n);
      }

      ...
    }
  }
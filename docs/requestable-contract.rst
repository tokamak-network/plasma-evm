.. _Requestable_Contract:
===============================
Requestable Contract
===============================

Requestable is a interface to be able to adapt Plasma EVM. Any contract implementiong ``requestable`` can accept enter and exit reqeust from RootChain contract.


.. _Request:

Request
=======

**Request** is an entity that makes users to interact with contracts. If user creates request, it is recorded in :ref:`RootChain<RootChain_Manager_Contract>` manager contract. A request is appiled in child chain as a **request transaction**. The sender of request transaction is ``0x00`` and it is included in **request block** to change the state of child chain. And it is enforced by RootChain contract to mine specific request block.


------------

Before go further, it is recommended to see how ``RootChain`` contract handle request :ref:`here<How_To_Handle_Request>`.


.. _Enter_and_Exit:
Enter and Exit
==============

**Etner** is "moving something from root chain to child chain". **Exit** is "moving something from child chain to root chain". The most intuitive example is token transfer. Depositing ERC20 to child chain is enter, and withdrawing it from child chain is exit.

`Enter request` is applied in root chain, then applied in child chain through request transaction. If applying in root chain is invalid, it **MUST** be reverted to prevent invalid enter request from being created.

`Exit request` is applied in child chain through request transaction, then applied in root chain. If the request is invalid, anyone can challenge on the invalid exit **with transaction receipt as proof**. If exit request is not challenged, anyone can `finalize` the reqeust and apply it to the requestable contract in root chain.

.. _Apply_Request_Functions:
``ApplyRequestIn*Chain`` Functions
==============================

If user wants to ``enter`` or ``exit``, he sends a transaction to ``RootChain`` contract to make  **enter request** or **exit request**. ``RootChain.startEnter()`` and ``RootChain.startExit()`` make user to create enter or exit request.

To accept those requests, contracts must implement **Requestable** interface.

See more how those requests are applied in child chain :ref:`here <How_to_Handle_Request>`.

::

  interface RequestableI {
    /// @notice Apply exit or enter request to requestable contract
    ///         deployed in root chain.
    function applyRequestInRootChain(
      bool isExit,
      uint256 requestId,
      address requestor,
      bytes32 trieKey,
      bytes trieValue
    ) external returns (bool success);

    /// @notice Apply exit or enter request to requestable contract
    ///         deployed in child chain.
    function applyRequestInChildChain(
      bool isExit,
      uint256 requestId,
      address requestor,
      bytes32 trieKey,
      bytes trieValue
    ) external returns (bool success);
  }


``applyRequestIn*Chain`` functions have common parameters.

- ``isExit``: ``true`` if the request is exit.
- ``requestId``: Identifier for the request. ``RootChain`` contract assigns it.
- ``requestor``: Address who made the request.
- ``trieKey``: Identifier for request type. ``trieKey`` tells the contract **what state variable should be changed** for this request
- ``trieValue``: Value of the request. ``trieValue`` tells the contract **how state should be changed**.


See more :ref:`examples <Requestable_Contract_Examples>`.
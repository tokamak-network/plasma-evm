.. _RootChain_Manager_Contract:

===============================
RootChain Manager Contract
===============================

**RootChain** contract is a manager contract for Plasma EVM. Operator submits plasma block data to this contract. Users also send transaction to make :ref:`request<Request>` for a requestable contract.

.. - Creat Request. ``startEnter``, ``startExit``
.. - Apply Request
.. - Convert request to request transaction
.. - Request Block
.. - Data Availability

The Units (?)
=========

Plasma blocks are abstracted in many way depending on each contexts. ``Block`` is the smallest unit that operator commits to RootChain contract. ``Epoch`` includes many blocks. ``Cycle`` includes many epochs.

.. _3_Types_of_Block:
3 Types of Block
~~~~~~~~~~~~~~~~

In Plasma EVM, There are many types of block; Non-Request Block, Request Block, Escape Block. **Non-Request Block** (NRB) is a regular block that users send transaction. It is same as block of Ethereum, Bitcoin, and other blockchain. **Request Block** (RB) is a block to apply :ref:`request<Request>` in child chain. Users create request and operator mines request block to apply the request in child chain. It is enforced what transactions must be included in request block by RootChain contract calculating transactions root. **Escape Block** (EB) is kind of request block to provide user to deal with block withholding attack by operator. For escape block, see continuous rebase.

NRB data can be withheld, but RB and EB cannot because anyone can get data from RootChain contract to mine RB and EB.


.. _Epoch:
Epoch
~~~~~

Epoch is a period of blocks. There are also **Non-Request Epoch** (NRE), **Request Epoch** (RE) and **Escape Epoch** (EE) for each block type. Epoch is required to determine when other type of block should be placed. Also a single request block cannot take all requests due to block gas limit and block size of child chain (and block gas limit of root chain to computate transactions root of request block).

The length of epoch is the number of blocks in the epoch. So RE and EE can have length of 0 if no RB or EB exists.

.. note::
  The length of NRE is a constant that is provided when RootChain contract is deployed.

.. _Cycle:
Cycle
~~~~~

**Continuous Rebase** requires many steps to commit a single block to root chain. And Pre-commit and Commit step incldue fixed number of NRBs. **Cycle** covers Pre-commit, DA check, Commit, Challenge steps and the length of cycle is ``the number of NREs + the number of OREs + 1 (the number of EE)``

.. note::
  The length of cycle is a constant that is provided when RootChain contract is deployed. But the number of blocks in a cycle may change depending on the number of requests.


.. _Block_Mining:
Block Mining
============

In 1 cycle, child chain block is mined as below:

- In pre-commit step, the first epoch is NRE. operator mines NRB with transactions from users.
- If users create enter requests and exit requests, they are reserved to be included in ORE after next NRE. For example, requests created in ``NRE#1`` or ``ORE#2`` are applied in ``ORE#4``. If no request created, ORE is empty.
- Pairs of ``NRE`` - ``ORE`` are repeated ``(size(cycle)) - 1 / 2`` times.
- In DA check step, user can create escape requests and undo requests.
- In Commit step, the first epoch is EE. Operator mines escape blocks if any request created in DA step. The parent of first block is the last block of previous cycle's last block.
- Operator rebases pre-committed block and commit them.


.. _How_to_Handle_Request:
How to Handle Request
=====================


Addresses of Requestable Contracts in Both Chain
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


:ref:`Reqeustable contracts<Requestable_Contract>` must be deployed in both chain and 2 addresses must be mapped in RootChain contract. Then user can make :ref:`requests<Request>` for requestable contracts.


Create Enter Request
~~~~~~~~~~~~~~~~~~~~

User can send transaction to RootChain contract to create enter request.

1. User send transaction to RootChain contract to call ``RootChain.startEnter()``.
2. RootChain contract apply the request to the corresponding requestable contract. Those happens in root chain.
3. If step 2 is not reverted, RootChain contract record the request.
4. In request epoch, operator mines request block with request transactions. See how request is converted into reqeust transaction :ref:`here<Apply_Request_in_Child_Chain>`.


::

  function startEnter(address _to,bytes32 _trieKey,bytes _trieValue)

``to`` is the address of target reqeustable contract in root chain. ``trieKey`` and ``trieValue`` is parameters for the request.


Create Enter Request
~~~~~~~~~~~~~~~~~~~~

User also can send transaction to RootChain contract to create exit request.

1. User send transaction to RootChain contract to call ``RootChain.startExit()``.
2. Unlike enter request, exit request is immediately recorded and mined in reuqest block with reqeust transactions. See how request is converted into reqeust transaction :ref:`here<Apply_Request_in_Child_Chain>`.
3. After challenge period for the requst block, challenge period for exit request starts. If the request transaction in step 2 is reverted, anyone can challenge on this by calling ``RootChain.challengeExit()`` with the transaction inclusion proof and receipt data.
4. If there is no successful challenge, User finalize the request by calling ``RootChain.finalizeRequest()``. In the function, RootChain contract apply the request to the corresponding requestable contract in root chain.

::

  function startExit(address _to,bytes32 _trieKey,bytes _trieValue)

Pamateres are same as ``startEnter``.


.. _Apply_Request_in_Child_Chain:
Apply Request in Child Chain
~~~~~~~~~~~~~~~~~~~~~~~~~~~~


A request has four important fields, ``requestor`` is a address who made the request, ``to`` is a address of requestable contract deployed in root chain, ``trieKey`` is a identifier for request type, and ``trieValue`` is the value of request.

When a request is transformed into **request transaction**, the transaction has those fields as follow:

- ``msg.sender``: it is always ``0x00``. It prevents other from creating request transaction because nobody know the private key of address ``0x00``. Due to this, signature of request transaction is zero , ``v = r = s = 0``.
- ``msg.to``: requestable contract **deployed in child chain**. RootChain contract must know it.
- ``msg.value``: it is always ``0``.
- ``msg.data``: To invoke message-call in transaction, this field must contain function signature and parameters for ``applyRequestInChildChain`` function. RootChain contract always knows what bytes should be in this field.

When the current epoch is RE, operator mines request block with request transactions to transit state of child chian. RootChain contract enforces operator to include what request transactions should be in the request block by calculating transactions root of the block.

Those request transactions are applied to requestable contract by :ref:`apply request functions<Apply_Request_Functions>`
======
Design
======

In this journey, we describe abstractly how Plama EVM is designed from Requetable to Continuous Rebase. See more details in technical paper.

Principle
===============

The fundamental pricinple is making plasma chain client by forking current Ethereum client (go-ethereum, pyethereum, parity). It provides many benefits like secure P2P communicartion, robustness, key management, and other things. But the most important thing is that other programming languages and tools for Ethereum (Solidity, Truffle, and so on) can be reused. It is wasteful to invent the wheel from scratch.

So the basic concept of Plasma EVM, Requestable, comes out. If we have 2 Ethereum networks and if they have to interoprate (like transfer `ERC20` or `(P)ETH` to another chain) in a truestless way, they need kind of **receipts** saying "Something is changed in chain A. So corresponding change is required in chain B". (e.g. transfer ERC20 from chain A to chain B by burning in chain A and minting in chain B).

Plasma EVM manager contract in root chain (a.k.a ``RootChain`` contract) needs a common interface for contracts that are compatible with Plasma EVM. We define those contracts as :doc:`Requestable contract <requestable-contract>`.

To accept the change in both chains, same requestable contracts must be deployed in both chain.

Deal with Data Availability
===========================

Exit Game
~~~~~~~~~

UTXO-based plasma proposals have a mechanism to protect assets in plasma chain. Plasma MVP uses confirmation model and exit game. Plasma Cash removes confirmation by introducing Sparse Merkle Tree to describe plasma chain state. Even if operator withholds block data, any user can exit their assets through exit game with proofs. They are absolutely secure models. Plasma EVM also needs those mechamism.

But Plasma EVM is account-based and storage of contract account is same as Ethereum, patricia trie. So we cannot adopt exit game. Instead, we developped user-activated fork.

User-activated Fork
~~~~~~~~~~~~~~~~~~~

Basic concept is that user can make special kind of request, ``exit request for URB``. And another type of block, user-submitted request block (``URB``), can be submitted to ``RootChain`` contract by any user if he can afford it. User-submitted request block has the last finalized block as parent.

So any user can make a new fork without fear of data unavailability. If they notice unavailability, then will make a new fork before unavailable block is finalized. But **to give higher priority on user-submitted request block than operator-submitted block**, the block number of user-submitted block has to be less than those of operator-submitted block.

We call this updating priority by **rebase**, named after git's rebase. Any operator-submitted blocks are placed after user-submitted block by rebasing. Rebase requires many computations in plasma chain and transactions in root chain, we considered "making fork even when data is available" as a attack.

But there is no concrete structural solution to prevent this attack. Our first naive approach is to charge higher cost to user, but this naive and finantial solution cannot be considerred as a soultion.

So we set rebase as the default behavior.

Continous Rebase
~~~~~~~~~~~~~~~~

In this model, rebase is considered as a default behavior. URB is renamed ``escape block`` and exit request for URB as ``escape request``.

Below 4 stages are processed sequentially in one **cycle**.

- In **pre-commit** stage, operator commits only transactions root to ``RootChain`` contract. If operator stop committing in this stage, **halting condition** is fulfilled. Halting condition forces to proceed the next stage.
- In **DA check** stage, if users notice unavailability for the root, they make a escape request. User also
- In **rebase** stage, operator mines and commits escape blocks whose parent is last block of previous cycle.
- In **commit** stage, operator commits other merkle roots of pre-committed blocks.

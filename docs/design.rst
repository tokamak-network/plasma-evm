======
Design
======

Here we discus about how Plama EVM is designed from Requetable to Rebase.


***********************
Requestable
***********************

The first pricinple is making plasma chain client by forking current Ethereum client (go-ethereum, pyethereum, parity). It provides many benefits like secure P2P communicartion, robustness, key management, and other things. But the most important thing is that other programming languages and tools for Ethereum (Solidity, Truffle, and so on) can be reused. It is wasteful to invent the wheel from scratch.

So the basic concept of Plasma EVM, Requestable, comes out. If we have 2 Ethereum networks and if they have to interoprate (like transfer `ERC20` or `(P)ETH` to another chain) in a truestless way, they need kind of receipts saying "Something is changed in chain A. So corresponding change is required in chain B". (e.g. transfer ERC20 from chain A to chain B by burning in chain A and minting in chain B)

But receipt-based state change requires tracing all previous changes in aother chain.
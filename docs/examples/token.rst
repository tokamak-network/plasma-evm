.. _example-token:

*****
Token
*****


Requestable Simple Token
~~~~~~~~~~~~~~~~~~~~~~~~

Requestable Simple Token is a simple implementation of Requestable Token for testing enter/exit process.

Test at Ropsten
===============
Imagine Alice wants to send a quantity of Token to Bob through Tokamak plasma chain, i.e. Alice is the sender and the Bob is the recipient.

Alice enter her balance to Faraday and send entered token to Bob.

If Bob want to use his token at rootchain, then exit his token to rootchain.


To test above scenario at `Ropsten`, you can use `RequestableSimpleTokenWithoutOwnership`.
Its contract address at Ropsten is ``0x6B27C38e3376C4E8B29cFbB3986f00676267D489`` and contract address at Faraday is
``0x1d93d7bd7d820ac7691109ace371e42d5004e1c1``. And `RootChain` contract address at Ropsten is ``0x3122546c1544FD0F910A423A8c80fdCD48d742Fd``.

The scenario will work as follows:

1. Alice mint ``RequestableSimpleTokenWithoutOwnership`` at the Ropsten.
2. Alice get her trieKey by using ``RequestableSimpleTokenWithoutOwnership.getBalanceTrieKey()`` and trieValue.
3. Alice call ``RootChain.startEnter()`` method to start entering process to Faraday.
4. After entering process is finished, you can check entered token balance by using ``RequestableSimpleTokenWithoutOwnership.balances()`` at Faraday.
5. Alice send her token balance token balance to Bob at Tokamak plasma chain.
6. If Bob wants to use his token at rootchain, then start exit process to rootchain by using ``RootChain.startExit()``.

::

    const Web3 = require('web3');

    const web3r = new Web3(new Web3.providers.HttpProvider(httpProviderUrlRoot)); // rpc endpoint of Ropsten
    const web3c = new Web3(new Web3.providers.HttpProvider(httpProviderUrlChild)); // rpc endpoint of Faraday

    const abiPathToken = path.join(__dirname, 'build', 'contracts', 'RequestableSimpleTokenWithoutOwnership.json');
    const abiToken = JSON.parse(fs.readFileSync(abiPath).toString()).abi;

    const abiPathRC = path.join(__dirname, 'build', 'contracts', 'RootChain.json');
    const abiRC = JSON.parse(fs.readFileSync(abiPathRC).toString()).abi;

    const rootchain = await web3r.eth.contract(abiRC).at('0x3122546c1544FD0F910A423A8c80fdCD48d742Fd');
    const tokenR = await web3r.eth.contract(abiToken).at('0x6B27C38e3376C4E8B29cFbB3986f00676267D489');
    const tokenC = await web3c.eth.contract(abiToken).at('0x1d93d7bd7d820ac7691109ace371e42d5004e1c1');

    const Alice = web3.eth.accounts[0];
    const Bob = web3.eth.accounts[1];

    async function scenario() {
        await tokenR.mint(Alice, 1000000, {from: Alice});

        const trieKeyAlice = await rootchain.getBalanceTrieKey(Alice);
        const trieValue = '0x0000000000000000000000000000000000000000000000000000000000000010' // 16 token
        await rootchain.startEnter('0x6B27C38e3376C4E8B29cFbB3986f00676267D489', trieKeyAlice, trieValue, {from: Alice})

        await tokenC.transfer(Alice, Bob, 16, {from: Alice})

        const trieKeyBob = await rootchain.getBalanceTrieKey(Bob);
        await rootchain.startExit('0x6B27C38e3376C4E8B29cFbB3986f00676267D489', trieKeyBob, trieValue, {from: Bob});

        await rootchain.finalizeRequests(n,{from: Bob});
    }


Requestable Simple Token Without Ownership
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


Mint
====
.. code-block:: javascript

    RequestableSimpleTokenWithoutOwnership.mint(address, amoutOfToken, {from: address})

``RequestableSimpleTokenWithoutOwnership`` has no ownable. So, evey body can mint this token fo exter/exit test.

Parameters
^^^^^^^^^^
1. ``String`` - The address to get the token.
2. ``Number`` - The amount of token.
3. ``String`` - The account that executes transaction.

Returns transaction hash, receipt and logs.

Example
^^^^^^^
.. code-block:: javascript

    RequestableSimpleTokenWithoutOwnership.mint('0x3cd9f729c8d882b851f8c70fb36d22b391a288cd', 1000000000, {from: '0x3cd9f729c8d882b851f8c70fb36d22b391a288cd'})

getBalanceTrieKey
-----------------
.. code-block:: javascript

    RequestableSimpleTokenWithoutOwnership.getBalanceTrieKey(addressWhoWantToDoRequest)

After minting ``RequestableSimpleTokenWithoutOwnership``, you must obtain the balance trieKey by using `getBalanceTrieKey` for the request.

Parameters
^^^^^^^^^^
1. ``String`` - The address to get the trieKey.

Example
^^^^^^^
.. code-block:: javascript

    RequestableSimpleTokenWithoutOwnership.getBalanceTrieKey('0x3cd9f729c8d882b851f8c70fb36d22b391a288cd')
    > '0xa874b0d77bbfe62ea7f8dea1dae821c054257558feac31e5646198857f7f1ba4'

startEnter
==========
.. code-block:: javascript

    RootChain.startEnter(tokenAddress, trieKeyOfAccountWhoWantEnter, trieValue, {from: address})


Parameters
^^^^^^^^^^
1. ``String`` - Token contract address that deployed at Ropsten.
2. ``String`` - .
3. ``String`` - .
4. ``String`` - The account that executes transaction.

Example
^^^^^^^
.. code-block:: javascript

    RootChain.startEnter('0x2f415f51fd16900fc1f92943b1f9a07f1b7eea14', '0xa874b0d77bbfe62ea7f8dea1dae821c054257558feac31e5646198857f7f1ba4', '0x0000000000000000000000000000000000000000000000000000000000000010', {from: '0x3cd9f729c8d882b851f8c70fb36d22b391a288cd'})

Returns ``transaction hash``, ``receipt`` and ``logs``.


startExit
=========
.. code-block:: javascript

    RootChain.startExit(tokenAddress, trieKeyOfAccountWhoWantExit, trieValue, {from: addressWhoExecuteExit, value: exitPrice})


Parameters
^^^^^^^^^^
1. ``String`` - Token contract address that deployed at Ropsten.
2. ``String`` - .
3. ``String`` - .
4. ``String`` - The account that executes transaction.
5. ``Number`` - An amount of Ether that used to pay for exit process.

Example
^^^^^^^
.. code-block:: javascript

    RootChain.startEnter('0x2f415f51fd16900fc1f92943b1f9a07f1b7eea14', '0xa874b0d77bbfe62ea7f8dea1dae821c054257558feac31e5646198857f7f1ba4', '0x0000000000000000000000000000000000000000000000000000000000000010', {from: '0x3cd9f729c8d882b851f8c70fb36d22b391a288cd'})

Returns ``transaction hash``, ``receipt`` and ``logs``.


finalizeRequest
===============
.. code-block:: javascript

    RootChain.finalizeRequests(numberOfRequestToFinalize)

To apply request at the Ropsten, you must finalize all of the requests.

Parameter
^^^^^^^^^

1. ``Number`` - Number of requests to finalize.

Example
^^^^^^^
.. code-block:: javascript

    RootChain.finalizeRequests(3)
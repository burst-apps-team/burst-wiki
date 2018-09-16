| Transaction and Fee Types |
|---------------------------|
| **Status**                |
||

Table Of Contents
-----------------

\_\_TOC\_\_

Introduction
------------

As of block height 500.000 the Burst blockchain supports, on top of ordinary transactions (one sender to one recipient), sending Burst from one account to multiple recipients in a single transaction. The first Pre-Dymaxion hard fork enabled using amounts below 1 Burst as transaction amounts and [transaction fees.](slot-based-transaction-fees.md)

Ordinary Transactions
---------------------

Ordinary Burst transactions are one-to-one transactions i.e. one sender sends Burst to one recipient. They can be issued through the local or web wallet.

To issue an ordinary transaction, the sender will specify the Burst account of the recipient, the amount of the transaction and the transaction fee. ![](1_Ordinary_transaction_-_basic.png "fig:1_Ordinary_transaction_-_basic.png") Optionally, users can add an arbitrary message to each outgoing transaction, which can be encrypted, in case of which the content of the message will be visible only to the recipient.

The transaction will be issued by clicking the “Send BURST” button.

#### Advanced options

Users can issue an ordinary transaction with a number of advanced options, which enable the conditional execution of the transaction and offline transaction signing. ![Advanced ordinary transaction options](3_Transaction_advanced_options.png "fig:Advanced ordinary transaction options")

-   Deadline sets the duration of transaction validity if it remains unconfirmed. The default and maximum allowed deadline is 24 hours. In case the transaction isn't confirmed within the configured deadline, it will be deleted from the mempool of the blockchain. In case the transaction expires in such a way, users will have to issue the transaction again in order to execute it.

<!-- -->

-   Referenced transaction hash represents a condition that has to be met before the transaction that is being issued can be confirmed. This mechanism works as follows: a transaction with hash *txhash<sub>1</sub>* has been issued. The user is currently creating transaction *tx<sub>2</sub>*. If the *txhash<sub>1</sub>* is provided as the “References Transaction Hash”, transaction *tx<sub>2</sub>* will be executed only after the transaction with *txhash<sub>1</sub>* has been confirmed.
-   The “Do Not Broadcast” option prevents the signed unconfirmed transaction to be broadcast to the network. Once the transaction has been signed offline, the user can broadcast it at a later time. When the “Do Not Broadcast” option is checked, the user can see the raw transaction details, as shown in the image below: ![Raw transaction details](4_Raw_transaction_details.png "fig:Raw transaction details")To broadcast the signed transaction later, the user has to save the raw transaction details in a separate file.
-   “Add Note to Self?” - this option will allow the entry of an encrypted note for the transaction.

#### Broadcast offline signed transactions

To broadcast a transaction that has been previously signed offline, users should access “Transaction Operations”. ![Transaction operations](5_Transaction_operations.png "fig:Transaction operations") The “Advanced Transaction Operations” then allows for the signed transaction bytes to be broadcast to the network: ![Broadcast transaction](6_Broadcast_transaction.png "fig:Broadcast transaction")

Multi-out Transactions
----------------------

Multi-out transactions allow the sender to send Burst to up to 64 unique recipient accounts as a single transaction i.e. with considerably lower fee compared to sending the same amount through 64 ordinary transactions.

Same as with ordinary transactions, multi-out transactions can be issued through local or web Burst wallets. ![Multi out](2_Multi_out.png "fig:Multi out")In case the amount to be sent is same for all recipients, the “Same amount” option should be checked. The number of recipients in Multi-out Same transactions can be up to 128.

#### Multi-out Transactions Lookup - Multi-out Reverse

Due to design optimizations, recipients of Multi-out and Multi-out Same transactions are able to see the correct account balance in their wallets, but the details of Multi-out and Multi-out Same transactions (the sender) is not shown in the wallet.

To see the details of the Multi-out transactions, recipients should use the block explorer, available at <https://explore.burst.cryptoguru.org/>, insert their Burst account into the search box and access the “Multiout Reverse” tab, where all Multi-out transaction details are shown: ![Multi-out Reverse](7_Multiout_reverse.png "fig:Multi-out Reverse")

Fee Types
---------

Unconfirmed transactions can be issued with arbitrary fees, with the lowest possible fee being 735.000 Plancks (0.00735000 Burst). Depending on the fee amount, unconfirmed transactions will be execute sooner or later i.e. they will be picked up into a block on the Burst blockchain.

[The Burst Reference Software](burst-reference-software.md) (BRS aka wallet) is equipped with a fee suggestion tool, which suggests the amount of fee to be used for transaction, based on the transaction load on the chain in the last 10 blocks. The fee suggestion tool offers three fee amounts:

-   Standard: 50% probability the transaction will be included in the next block
-   Cheap: 50% probability the transaction will be included in the next 10 blocks
-   Priority: 90% probability the transaction will be included in the next block, 99% probability the transaction will be included in the next two blocks.

The fee suggestion tool can also be used using [The Burst API](the-burst-api-suggest-fee.md).

**The documentation here refers to the version or later.**

Description
---------------------------------------------------------------------------------------------------------

The Burst API allows interaction with Burst nodes using HTTP requests to port 8125. Most HTTP requests can use either the GET or POST methods, but some API calls accept only the POST method for security reasons. Responses are returned as JSON objects.

Each API call is documented below, with definitions given for HTTP request parameters and JSON response fields, followed by an example:

-   The JSON response fields are each followed by one of *S* for string, *A* for array, *O* for object, *N* for number or *B* for boolean.
-   In the examples, the Burst node is represented as *localhost* and requests and responses are formatted for easy reading; line breaks and spaces are not actually used except in some parameter values. All requests are in URL format which implies the HTTP GET method. When GET is allowed, the URL can be entered into a browser URL field but proper URL encoding is usually required (e.g., spaces in a parameter value must be replaced by *+* or *%20*). Otherwise, the URL should be used as a guide to preparing an HTTP POST request using cURL, for example.

All API calls can be viewed and tested on the TestNet at <https://wallet.dev.burst-test.net/test>. For specific API calls, use the GET url <https://wallet.dev.burst-test.net/burst?requestType>=*specificRequestType*.

General Notes
-----------------------------------------------------------------------------------------------------------

### Genesis Block

Many API requests make reference to the .

### Flexible Account IDs

All API requests that require an account ID accept either an account number or the corresponding [Reed-Solomon address](rs-address-format.md).

### Quantity Units BURST, NQT and QNT

The Burst system has a currency BURST used to quantify value in the system. Like all currencies, BURST circulates in the system, moving from one user to another by payments and purchases. Also, a small fee is required for every transaction, including those in which no BURST is transfered, such as simple messages. This fee goes to the owner of the node that forges (generates) the new block containing the transaction that is accepted onto the blockchain.

Yet internally, the currency is still stored in integer form in units of NQT or NxtQuant, where 1 BURST = 10<sup>8</sup> NQT. All parameters and fields in the API involving a quantity of BURST are denominated in units of NQT, for example *feeNQT*. The only exception is the field *effectiveBalanceNXT*, used in forging calculations.

Other assets can be created within Burst using [Issue Asset](#issue-asset). The issuer must specify the number of decimal places to use in quantifying the asset, and the amount of the asset to create in generic units of QNT or Quant, distinct from NQT. Quantities of assets are stored internally as integers in units of QNT, and assets are priced in NQT per QNT.

| Decimal (BURST) | Canonical Name  | Alternate Name | NQT         |
|-----------------|-----------------|----------------|-------------|
| **1**.00000000  | BURST           | Burst          | 100,000,000 |
| 0.0**1**000000  | cBURST          | Bessie         | 1,000,000   |
| 0.00**1**00000  | mBURST          | -              | 100,000     |
| 0.0000**1**000  | -               | Maybel         | 1,000       |
| 0.00000**1**00  | uBURST          | -              | 100         |
| 0.0000000**1**  | -               | Planck         | 1           |
| 0.12345678      | digit reference | -              | -           |

### Creating Unsigned Transactions

All API requests that create a new transaction will accept either a *secretPhrase* or a *publicKey* parameter:

-   If *secretPhrase* is supplied, a transaction is created, signed at the server, and broadcast by the server as usual.
-   If only a *publicKey* parameter is supplied as a 64-digit (32-byte) hex string, the transaction will be prepared by the server and returned in the JSON response as *transactionJSON* without a signature. This JSON object along with *secretPhrase* can then be supplied to [Sign Transaction](#sign-transaction) as *unsignedTransactionJSON* and the resulting signed *transactionJSON* can then be supplied to [Broadcast Transaction](#broadcast-transaction). This sequence allows for offline signing of transactions so that *secretPhrase* never needs to be exposed.
-   *unsignedTransactionBytes* may be used instead of unsigned *transactionJSON* when there is no encrypted message. Messages to be encrypted require the *secretPhrase* for encryption and so cannot be included in *unsignedTransactionBytes*.

### Escrow Operations

All API requests that create a new transaction will accept an optional *referencedTransactionFullHash* parameter which creates a chained transaction, meaning that the new transaction cannot be confirmed unless the referenced transaction is also confirmed. This feature allows a simple way of transaction escrow:

-   Alice prepares and signs a transaction A, but doesn't broadcast it by setting the *broadcast* parameter to *false*. She sends to Bob the *unsignedTransactionBytes*, the *fullHash* of the transaction, and the *signatureHash*. All of those are included in the JSON returned by the API request. (Warning: make sure not to send the signed *transactionBytes*, or the *signature* itself, as then Bob can just broadcast transaction A himself).
-   Bob prepares, signs and broadcasts transaction B, setting the *referencedTransactionFullHash* parameter to the *fullHash* of A provided by Alice. He can verify that this hash indeed belongs to the transaction he expects from Alice, by using [Calculate Full Hash](#calculate-full-hash), which takes *unsignedTransactionBytes* and *signatureHash* (both of which Alice has also sent to Bob). He can also use [Parse Transaction](#parse-transaction) to decode the unsigned bytes and inspect all transaction fields.
-   Transaction B is accepted in the unconfirmed transaction pool, but as long as A is still missing, B will not be confirmed, i.e. will not be included in the blockchain. If A is never submitted, B will eventually expire -- so Bob should make sure to set a long enough deadline, such as the maximum of 32767 minutes.
-   Once in the unconfirmed transactions pool, Bob has no way of recalling B back. So now Alice can safely submit her transaction A, by just broadcasting the *signedTransactionBytes* she got in the first step. Transaction A will get included in the blockchain first, and in the next block Bob's transaction B will also be included.

Note that while the above scheme is good enough for a simple escrow, the blockchain does not enforce that if A is included, B will also be included. It may happen due to forks and blockchain reorganization, that B never gets a chance to be included and expires unconfirmed, while A still remains in the blockchain. However, it is not practically possible for Bob to intentionally cause such chain of events and to prevent B from being confirmed.

### Prunable Data

Prunable data can be removed from the blockchain without affecting the integrity of the blockchain. When a transaction containing prunable data is created, only the 32-byte sha256 hash of the prunable data is included in the *transactionBytes*, not the prunable data itself. The non-prunable signed *transactionBytes* are used to verify the signature and to generate the transaction's *fullHash* and ID; when the prunable part of the transaction is removed at a later time, none of these operations are affected.

Prunable data has a predetermined minimum lifetime of two weeks (24 hours on the [Testnet](testnet.md)) from the timestamp of the transaction. Transactions and blocks received from peer nodes are not accepted if prunable data is missing before this time has elapsed. After this time has elapsed, prunable data is no longer included with transactions and blocks transmitted to peer nodes, and is no longer included in the transaction JSON returned by general-purpose API calls such as Get Transaction; the only way to retrieve it, if still available, is through special-purpose API calls such as Get Prunable Message.

Expired prunable data remains stored in the blockchain until removed at the same time derived tables are trimmed, which occurs automatically every 1000 blocks by default.

Prunable data can be preserved on a node beyond the predetermined minimum lifetime by setting the *nxt.maxPrunableLifetime* property to a larger value than two weeks or to *-1* to preserve it indefinitely. To force the node to include such preserved prunable data when transactions and blocks are transmitted to peer nodes, set the *nxt.includeExpiredPrunables* property to *true*, thus making it an archival node.

Currently, there is only one variety of prunable data in the Burst system: prunable [Arbitrary Messages](arbitrary-messages.md). It has a maximum prunable data length of 42 kilobytes.

### Properties Files

The behavior of some API calls is affected by property settings loaded from files in the *brs/conf* directory during Burst server intialization. This directory contains the *brs-default.properties* and *logging-default.properties* files, both of which contain default property settings along with documentation. A few of the property settings can be obtained while the server is running through the [Get Blockchain Status](#get-blockchain) and [Get State](#get-state) calls.

It is recommended not to modify default properties files because they can be overwritten during software updates. Instead, properties in the default files can be overridden by including them in optional *brs.properties* and *logging.properties* files in the same directory. For example, a *brs.properties* file can be created with the content:

`DEV.TestNet = yes`

This causes the Burst server to connect to the [TestNet](testnet.md) instead of the MainNet.

Create Transaction
------------------

The following API calls create Burst transactions using HTTP POST requests. Follow the links for examples and HTTP POST request parameters specific to each call. Refer to the sections below for [HTTP POST request parameters](#create-transaction-request) and [JSON response fields](#create-transaction-response) common to all calls that create transactions.

-   [Send Money](#send-money)
-   [Set Account Info](#set-account-info)
-   [Buy / Sell Alias](#buy--2f-sell-alias)
-   [Set Alias](#set-alias)
-   [Send Message](#send-message)
-   [Cancel Order](#cancel-order)
-   [Issue Asset](#issue-asset)
-   [Place Order](#place-order)
-   [Transfer Asset](#transfer-asset)
-   [DGS Delisting](#dgs-delisting)
-   [DGS Delivery](#dgs-delivery)
-   [DGS Feedback](#dgs-feedback)
-   [DGS Listing](#dgs-listing)
-   [DGS Purchase](#dgs-purchase)
-   [DGS Quantity Change](#dgs-quantity-change)
-   [DGS Refund](#dgs-refund)
-   [Lease Balance](#lease-balance)

### Create Transaction Request

The following HTTP POST parameters are common to all API calls that create transactions:

For `feeNQT`, please refer to the following “rules”:

-   minimum 1000 BURST for [Issue Asset](#issue-asset), unless singleton asset is issued, for which the fee is 1 BURST
-   2 BURST in base fee for [Set Alias](#set-alias), with 2 BURST additional fee for each 32 chars of name plus URI total length, after the first 32 chars
-   1 BURST for the first 32 bytes of a unencrypted non-prunable [message](#send-message), 1 BURST for each additional 32 bytes
-   2 BURST for the first 32 bytes of an encrypted non-prunable [message](#send-message), 1 BURST for each additional 32 bytes. The length is measured excluding the nonce and the 16 byte AES initialization vector.
-   1 BURST for the first 1024 bytes of a prunable [message](#send-message), 0.1 BURST for each additional 1024 prunable bytes
-   1 BURST for [Set Account Info](#set-account-info), including 32 chars, with 2 BURST additional fee for each 32 chars
-   2 BURST for [DGS Listing](#dgs-listing), including 32 chars of name plust description. With 2 BURST additional fee for each 32 chars.
-   1 BURST for [DGS Delivery](#dgs-delivery), including 32 bytes of encrypted goods data (AES initialization bytes and nonce excluded). With 2 BURST additional fee for each 32 bytes.
-   2 BURST for transactions that make use of referencedTransactionFullHash property when creating a new transaction.
-   Dynamic tx fee otherwise, where 1 BURST = 100000000 NQT

| Tx no | Tx fees |
|-------|---------|
| 1     | 0.00735 |
| 100   | 0.73500 |
| 255   | 1.87425 |
| 510   | 3.74850 |
| 765   | 5.62275 |
| 1020  | 7.49700 |

**Note:** An optional arbitrary message can be appended to any transaction by adding message-related parameters as in [Send Message](#send-message).

### Create Transaction Response

The following JSON response fields are common to all API calls that create transactions:

Account Operations
------------------

### Get Account

Get account information given an account ID.

**Response:**

**Example:** Refer to [Get Account](the-burst-api-examples.md#get-account) example.

### Get Account Block Ids

Get the block IDs of all blocks forged (generated) by an account in reverse block height order.

**Response:**

### Get Account Blocks

Get all blocks forged (generated) by an account in reverse block height order.

**Response:**

### Get Account Id

Get an account ID given a secret passphrase or public key. POST only.

**Response:**

**Example:** Refer to [Get Account Id](the-burst-api-examples.md#get-account-id) example.

### Get Account Lessors

Get the lessors to an account.

**Note:** If table trimming is enabled (default), the *height* must be within 1440 blocks of the last block.

**Response:**

-   *lessorRS* (S)
-   *lessor* (S)
-   *guaranteedBalanceNQT* (S)

**Example:** Refer to [Get Account Lessors](the-burst-api-examples.md#get-account-lessors) example.

### Get Account Public Key

Get the public key associated with an account ID.

**Response:**

**Example:** Refer to [Get Account Public Key](the-burst-api-examples.md#get-account-public-key) example.

### Get Account Transaction Ids

Get the transaction IDs associated with an account in reverse block timestamp order. *This call only returns non-phased transactions as of [Version 1.5.7e](burst-software-change-log-version-1-5-7e.md) and is deprecated, to be removed in version 1.6. Use [Get Blockchain Transactions](#get-blockchain-transactions) instead.*

**Note:** Refer to [Get Constants](#get-constants) for definitions of types and subtypes

**Response:**

**Example:** Refer to [Get Account Transaction Ids](the-burst-api-examples.md#get-account-transaction-ids) example.

### Get Account Transactions

Get the transactions associated with an account in reverse block timestamp order. *This call only returns non-phased transactions as of [Version 1.5.7e](burst-software-change-log-version-1-5-7e.md) and is depricated, to be removed in version 1.6. Use [Get Blockchain Transactions](#get-blockchain-transactions) instead.*

**Note:** Refer to [Get Constants](#get-constants) for definitions of types and subtypes

**Response:**

**Example:** Refer to [Get Account Transactions](the-burst-api-examples.md#get-account-transactions) example.

### Get Balance

Get the balance of an account.

**Request:**

**Response:**

**Example:** Refer to [Get Balance](the-burst-api-examples.md#get-balance) example.

### Get Guaranteed Balance

Get the balance of an account that is confirmed at least a specified number of times.

**Response:**

**Example:** Refer to [Get Guaranteed Balance](the-burst-api-examples.md#get-guaranteed-balance) example.

### Get Unconfirmed Transaction Ids

Get a list of unconfirmed transaction IDs associated with an account.

**Request:**

**Response:**

**Example:** Refer to [Get Unconfirmed Transaction Ids](the-burst-api-examples.md#get-unconfirmed-transaction-ids) example.

### Get Unconfirmed Transactions

Get a list of unconfirmed transactions associated with an account.

**Response:**

**Example:** Refer to [Get Unconfirmed Transactions](the-burst-api-examples.md#get-unconfirmed-transactions) example.

### Send Money

Send BURST to an account. POST only. Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

**Response:** Refer to [Create Transaction Response](#create-transaction-response).

**Example:** Refer to [Send Money](the-burst-api-examples.md#send-money) example.

### Send Money Multi

Send individual amounts of BURST to up to 64 recipients. POST only. Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

**Response:** Refer to [Create Transaction Response](#create-transaction-response).

**Example:** To do

### Send Money Multi Same

Send the same amount of BURST to up to 128 recipients. POST only. Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

**Response:** Refer to [Create Transaction Response](#create-transaction-response).

**Example: To do**

### Set Account Info

Set account information. POST only. Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

**Response:** Refer to [Create Transaction Response](#create-transaction-response).

**Example:** Refer to [Set Account Info](the-burst-api-examples.md#set-account-info) example.

Alias Operations
----------------

### Buy / Sell Alias

Buy or sell an alias. POST only. Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

**Note**: An alias can be transferred rather than sold by setting *priceNQT* to zero. A pending sale can be canceled by selling again to self for a price of zero.

**Response:** Refer to [Create Transaction Response](#create-transaction-response).

**Example:** Refer to [Buy / Sell Alias](the-burst-api-examples.md#buy---sell-alias) example.

### Set Alias

Create and/or assign an alias. POST only. Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

**Response:** Refer to [Create Transaction Response](#create-transaction-response). The transaction ID is also the alias ID.

**Example:** Refer to [Set Alias](the-burst-api-examples.md#set-alias) example.

### Get Alias

Get information about a given alias.

**Response:**

**Example:** Refer to [Get Alias](the-burst-api-examples.md#get-alias) example.

### Get Aliases

Get information on aliases owned by a given account in alias name order.

**Response:**

**Example:** Refer to [Get Aliases](the-burst-api-examples.md#get-aliases) example.

Arbitrary Message System Operations
-----------------------------------

### Decrypt From

Decrypt an AES-encrypted message.

**Request:**

-   *requestType* is *decryptFrom*
-   *secretPhrase* is the secret passphrase of the recipient
-   *account* is the account ID of the recipient
-   *data* is AES-encrypted data
-   *nonce* is the unique nonce associated with the encrypted data
-   *decryptedMessageIsText* is *false* if the decrypted message is a hex string, otherwise the decrypted message is text (optional)
-   *uncompressDecryptedMessage* is *false* to prevent gzip uncompression after decryption (optional)

**Response:**

-   *decryptedMessage* (S) is the decrypted message
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Decrypt From](the-burst-api-examples.md#decrypt-from) example.

### Encrypt To

Encrypt a message using AES without sending it.

**Request:**

-   *requestType* is *encryptTo*
-   *secretPhrase* is the secret passphrase of the sender
-   *recipient* is the account ID of the recipient
-   *messageToEncrypt* is either UTF-8 text or a string of hex digits to be compressed and converted into a 1000 byte maximum bytecode then encrypted using AES
-   *messageToEncryptIsText* is *false* if the message to encrypt is a hex string, otherwise the message to encrypt is text (optional)
-   *compressMessageToEncrypt* is *false* to prevent gzip compression before encryption (optional)

**Response:**

-   *data* (S) is the AES-encrypted data
-   *nonce* (S) is a 32-byte pseudorandom nonce
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Encrypt To](the-burst-api-examples.md#encrypt-to) example.

### Read Message

Get a message given a transaction ID.

**Request:**

-   *requestType* is *readMessage*
-   *transaction* is the transaction ID of the message
-   *secretPhrase* is the secret passphrase of the account that received the message (optional)
-   *sharedKey* is the shared key used to decrypt the message (optional) (see [Get Shared Key](#get-shared-key))
-   *retrieve* is *true* to retrieve pruned data from archival nodes (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *messageIsPrunable* (B) is *true* if there is a plain message and it is prunable, *false* if it is not prunable
-   *message* (S) is the plain message, if applicable
-   *encryptedMessageIsPrunable* (B) is *true* if there is an encrypted message and it is prunable, *false* if it is not prunable
-   *decryptedMessage* (S) is the decrypted message, if applicable and only if the provided *secretPhrase* belongs to either the sender or receiver of the transaction
-   *decryptedMessageToSelf* (S) is the decrypted message sent to self, if applicable and only if the provided *secretPhrase* belongs to the sender of transaction
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Read Message](the-burst-api-examples.md#read-message) example.

### Send Message

Create an Arbitrary Message transaction. POST only.

**Request:** Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

-   *requestType* is *sendMessage*
-   *recipient* is the account ID of the recipient (optional)
-   *recipientPublicKey* is the public key of the receiving account (optional, enhances security of a new account)
-   *message* is either UTF-8 text or a string of hex digits (perhaps previously encoded using an arbitrary algorithm) to be converted into a bytecode with a maximum length of one kilobyte, 42 kilobytes if prunable (optional)
-   *messageIsText* is *false* if the message is a hex string, otherwise the message is text (optional)
-   *messageIsPrunable* is *true* if the message is prunable (optional)
-   *messageToEncrypt* is either UTF-8 text or a string of hex digits to be compressed (unless *compressMessageToEncrypt* is *false*) and converted into a bytecode with a maximum length of one kilobyte, 42 kilobytes if prunable, then encrypted using AES (optional)
-   *messageToEncryptIsText* is *false* if the message to encrypt is a hex string, otherwise the message to encrypt is text (optional)
-   *encryptedMessageData* is already encrypted data which overrides *messageToEncrypt* if provided (optional)
-   *encryptedMessageNonce* is a unique 32-byte number which cannot be reused (optional unless *encryptedMessageData* is provided)
-   *encryptedMessageIsPrunable* is *true* if the encrypted message is prunable (optional)
-   *compressMessageToEncrypt* is *false* to prevent gzip compression before encryption (optional)
-   *messageToEncryptToSelf* is either UTF-8 text or a string of hex digits to be compressed (unless *compressMessageToEncryptToSelf* is false) and converted into a one kilobyte maximum bytecode then encrypted with AES, then sent to the sending account (optional)
-   *messageToEncryptToSelfIsText* is *false* if the message to self-encrypt is a hex string, otherwise the message to encrypt is text (optional)
-   *encryptToSelfMessageData* is already encrypted data which overrides messageToEncryptToSelf if provided (optional)
-   *encryptToSelfMessageNonce* is a unique 32-byte number which cannot be reused (optional unless encryptToSelfMessageData is provided)
-   *compressMessageToEncryptToSelf* is *false* to prevent gzip compression before encryption (optional)

**Note:** Any combination (including none or all) of the three options plain *message*, *messageToEncrypt*, and *messageToEncryptToSelf* will be included in the transaction. However, one and only one prunable message may be included in a single transaction if there is not already a message of the same type (either plain or encrypted).

**Note:** The *encryptedMessageData-encryptedMessageNonce* pair or the *encryptToSelfMessageData-encryptToSelfMessageNonce* pair can be the output of [Encrypt To](#encrypt-to)

**Response:** Refer to [Create Transaction Response](#create-transaction-response).

**Example:** Refer to [Send Message](the-burst-api-examples.md#send-message) example.

Asset Exchange Operations
-------------------------

### Cancel Order

Cancel an existing asset order. POST only.

**Request:** Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

-   *requestType* is either *cancelBidOrder* or *cancelAskOrder*
-   *order* is the order ID of the order being canceled

**Response:** Refer to [Create Transaction Response](#create-transaction-response).

**Example:** Refer to [Cancel Order](the-burst-api-examples.md#cancel-order) example.

#### Cancel Ask Order

Refer to [Cancel Order](#cancel-order).

#### Cancel Bid Order

Refer to [Cancel Order](#cancel-order).

### Get Account Current Order Ids

Get current asset order IDs given an account ID in reverse block height order.

**Request:**

-   *requestType* is either *getAccountCurrentBidOrderIds* or *getAccountCurrentAskOrderIds*
-   *account* is the account ID
-   *asset* is an asset ID filter (optional)
-   *firstIndex* is a zero-based index to the first order ID to retrieve (optional)
-   *lastIndex* is a zero-based index to the last order ID to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *bidOrderIds* or *askOrderIds* (A) is an array of order IDs
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Current Order Ids](the-burst-api-examples.md#get-account-current-order-ids) example.

#### Get Account Current Ask Order Ids

Refer to [Get Account Current Order Ids](#get-account-current-order-ids).

#### Get Account Current Bid Order Ids

Refer to [Get Account Current Order Ids](#get-account-current-order-ids).

### Get Account Current Orders

Get current asset orders given an account ID in reverse block height order.

**Request:**

-   *requestType* is either *getAccountCurrentBidOrders* or *getAccountCurrentAskOrders*
-   *account* is the account ID
-   *asset* is an asset ID filter (optional)
-   *firstIndex* is a zero-based index to the first order to retrieve (optional)
-   *lastIndex* is a zero-based index to the last order to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *bidOrders* or *askOrders* (A) is an array of order objects (refer to [Get Order](#get-order) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Current Orders](the-burst-api-examples.md#get-account-current-orders) example.

#### Get Account Current Ask Orders

Refer to [Get Account Current Orders](#get-account-current-orders).

#### Get Account Current Bid Orders

Refer to [Get Account Current Orders](#get-account-current-orders).

### Get All Assets

Get all assets in the exchange in reverse block height of creation order.

**Request:**

-   *requestType* is *getAllAssets*
-   *firstIndex* is a zero-based index to the first asset to retrieve (optional)
-   *lastIndex* is a zero-based index to the last asset to retrieve (optional)
-   *includeCounts* is *true* if the fields beginning with *numberOf...* are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *assets* (A) is an array of asset objects (refer to [Get Asset](#get-asset))
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get All Assets](the-burst-api-examples.md#get-all-assets) example.

### Get All Open Orders

Get all open bid/ask orders in reverse block height order.

**Request:**

-   *requestType* is either *getAllOpenBidOrders* or *getAllOpenAskOrders*
-   *firstIndex* is a zero-based index to the first order to retrieve (optional)
-   *lastIndex* is a zero-based index to the last order to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *openOrders* (A) is an array of order objects (refer to [Get Order](#get-order) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get All Open Orders](the-burst-api-examples.md#get-all-open-orders) example.

#### Get All Open Ask Orders

Refer to [Get All Open Orders](#get-all-open-orders).

#### Get All Open Bid Orders

Refer to [Get All Open Orders](#get-all-open-orders).

### Get All Trades

Get all trades since a given timestamp in reverse block height order.

**Request:**

-   *requestType* is *getAllTrades*
-   *timestamp* is the timestamp (in seconds since the genesis block) to begin retrieving trades (optional, default 0)
-   *firstIndex* is a zero-based index to the first trade to retrieve (optional)
-   *lastIndex* is a zero-based index to the last trade to retrieve (optional)
-   *includeAssetInfo* is *true* if asset information is to be included in the result (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** If *timestamp* is omitted or zero, and no index is given, all trades in the entire blockchain will be retrieved, which may timeout or crash your system.

**Response:**

-   *trades* (A) is an array of trade objects (refer to [Get Trades](#get-trades))
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get All Trades](the-burst-api-examples.md#get-all-trades) example.

### Get Asset

Get asset information given an asset ID.

**Request:**

-   *requestType* is *getAsset*
-   *asset* is the asset ID
-   *includeCounts* is *true* if the fields beginning with *numberOf...* are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *account* (S) is the number of the account that issued the asset
-   *accountRS* (S) is the Reed-Solomon address of the account that issued the asset
-   *name* (S) is the asset name
-   *description* (S) is the asset description
-   *quantityQNT* (S) is the total asset quantity (in QNT) in existence
-   *asset* (N) is the asset ID
-   *decimals* (N) is the number of decimal places used by the asset
-   *numberOfAccounts* (N) is the number of accounts that own the asset
-   *numberOfTrades* (N) is the number of trades of this asset
-   *numberOfTransfers* (N) is the number of transfers of this asset
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Asset](the-burst-api-examples.md#get-asset) example.

### Get Asset Accounts

Get the accounts that own an asset given the asset ID in reverse quantity order.

**Request:**

-   *requestType* is *getAssetAccounts*
-   *asset* is the asset ID
-   *height* is the height of the blockchain to determine the accounts (optional, default is last block)
-   *firstIndex* is a zero-based index to the first account to retrieve (optional)
-   *lastIndex* is a zero-based index to the last account to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** If table trimming is enabled (default), the height must be within 1440 blocks of the last block.

**Response:**

-   *accountAssets* (A) is an array of asset objects with the following fields for each asset:
    -   *quantityQNT* (S) is the quantity (in QNT) of the asset
    -   *accountRS* (S) is the Reed-Solomon address of the account that owns the asset
    -   *unconfirmedQuantityQNT* (S) is the unconfirmed quantity (in QNT) of the asset
    -   *asset* (S) is the asset ID
    -   *account* (S) is the number of the account that owns the asset
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Asset Accounts](the-burst-api-examples.md#get-asset-accounts) example.

### Get Asset Ids

Get the IDs of all assets in the exchange in reverse block height of creation order.

**Request:**

-   *requestType* is *getAssetIds*
-   *firstIndex* is a zero-based index to the first asset ID to retrieve (optional)
-   *lastIndex* is a zero-based index to the last asset ID to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *assets* (A) is an array of asset IDs
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Asset Ids](the-burst-api-examples.md#get-asset-ids) example.

### Get Asset Transfers

Get transfers associated with a given asset and/or account in reverse block height order (or in the expected order of execution for expected transfers).

**Request:**

-   *requestType* is either *getAssetTransfers* or *getExpectedAssetTransfers*, where expected transfers are from the unconfirmed transactions pool or are phased transactions scheduled to finish in the next block
-   *asset* is the asset ID (optional)
-   *account* is the account ID (optional if *asset* provided)
-   *timestamp* is the earliest transfer (in seconds since the genesis block) to retrieve (optional, does not apply to expected transfers)
-   *firstIndex* is a zero-based index to the first transfer to retrieve (optional, does not apply to expected transfers)
-   *lastIndex* is a zero-based index to the last transfer to retrieve (optional, does not apply to expected transfers)
-   *includeAssetInfo* is *true* if the *decimals* and *name* fields are to be included (optional, does not apply to expected transfers)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *transfers* (A) is an array of transfer objects with the following fields for each transfer:
    -   *quantityQNT* (S) is the quantity (in QNT) of the asset traded
    -   *senderRS* (S) is the Reed-Solomon address of the sender
    -   *assetTransfer* (S) is the transaction ID of the asset transfer
    -   *sender* (S) is the account number of the sender
    -   *recipientRS* (S) is the Reed-Solomon address of the recipient
    -   *decimals* (N) is the number of decimal places used by the asset (if *includeAssetInfo* is *true*)
    -   *recipient* (S) is the account number of the recipient
    -   *name* (S) is the name of the asset (if *includeAssetInfo* is *true*)
    -   *asset* (S) is the asset ID
    -   *height* (N) is the height of the transfer block
    -   *timestamp* (N) is the timestamp (in seconds since the genesis block) of the transfer block, does not apply to an expected transfer
    -   *phased* (B) is *true* if the transaction is phased, *false* otherwise, applies only to an expected transfer
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Asset Transfers](the-burst-api-examples.md#get-asset-transfers) example.

#### Get Expected Asset Transfers

Refer to [Get Asset Transfers](#get-asset-transfers).

### Get Assets

Get asset information given multiple asset IDs

**Request:**

-   *requestType* is *getAssets*
-   *assets* is one the multiple asset IDs
-   *assets* is one the multiple asset IDs

<!-- -->

-   *includeCounts* is *true* if the fields beginning with *numberOf...* are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *assets* (A) is an array of asset objects (refer to [Get Asset](#get-asset))
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Assets](the-burst-api-examples.md#get-assets) example.

### Get Assets By Issuer

Get asset information given multiple creation account IDs in reverse block height of creation order.

**Request:**

-   *requestType* is *getAssetsByIssuer*
-   *account* is one of the multiple account IDs
-   *account* is one of the multiple account IDs

<!-- -->

-   *firstIndex* is a zero-based index to the first asset to retrieve (optional)
-   *lastIndex* is a zero-based index to the last asset to retrieve (optional)
-   *includeCounts* is *true* if the fields beginning with *numberOf...* are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *assets* (A) is an array of asset objects (refer to [Get Asset](#get-asset))
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Assets By Issuer](the-burst-api-examples.md#get-assets-by-issuer) example.

### Get Order

Get a bid/ask order given an order ID.

**Request:**

-   *requestType* is either *getBidOrder* or *getAskOrder*
-   *order* is the Order ID
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *account* (S) is the account number associated with the order
-   *accountRS* (S) is the Reed-Solomon address of the account
-   *asset* (S) is the ID of the asset being ordered
-   *quantityQNT* (S) is the order quantity (in QNT)
-   *priceNQT* (S) is the order price (in NQT)
-   *height* (N) is the block height of the order transaction
-   *transactionHeight* (N) is the transaction height
-   *transactionIndex* (N) is a zero-based index giving the order of the transaction in its block
-   *order* (S) is the ID of the order
-   *type* (S) is the type of order (*bid* or *ask*)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Order](the-burst-api-examples.md#get-order) example.

#### Get Ask Order

Refer to [Get Order](#get-order).

#### Get Bid Order

Refer to [Get Order](#get-order).

### Get Order Ids

Get bid/ask order IDs given an asset ID, in order of decreasing bid price or increasing ask price.

**Request:**

-   *requestType* is either *getBidOrderIds* or *getAskOrderIds*
-   *asset* is the asset ID
-   *firstIndex* is a zero-based index to the first order ID to retrieve (optional)
-   *lastIndex* is a zero-based index to the last order ID to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *bidOrderIds* or *askOrderIds* (A) is an array of order IDs
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Order Ids](the-burst-api-examples.md#get-order-ids) example.

#### Get Ask Order Ids

Refer to [Get Order Ids](#get-order-ids).

#### Get Bid Order Ids

Refer to [Get Order Ids](#get-order-ids).

### Get Orders

Get bid/ask orders given an asset ID, in order of decreasing bid price or increasing ask price (if *sortByPrice* is *true* for expected orders, otherwise in the expected order of execution).

**Request:**

-   *requestType* is one of *getBidOrders*, *getAskOrders*, where expected orders are from the unconfirmed transactions pool or are phased transactions scheduled to finish in the next block
-   *asset* is the asset ID
-   *sortByPrice* is *true* to sort by price (optional, applies only to expected orders, which are returned in expected order of execution by default)
-   *showExpectedCancellations* is *true* to include orders that are expected to be cancelled in the next block, based on the content of the unconfirmed transactions pool and the phased transactions expected to finish in the next block (optional, does not apply to expected orders)
-   *firstIndex* is a zero-based index to the first order to retrieve (optional, does not apply to expected orders)
-   *lastIndex* is a zero-based index to the last order to retrieve (optional, does not apply to expected orders)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *bidOrders* or *askOrders* (A) is an array of order objects (refer to [Get Order](#get-order) for details) with the following additional field only for an expected order:
    -   *phased* (B) is *true* if the order is phased, *false* otherwise
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Orders](the-burst-api-examples.md#get-orders) example.

#### Get Ask Orders

Refer to [Get Orders](#get-orders).

#### Get Bid Orders

Refer to [Get Orders](#get-orders).

### Get Trades

Get trades associated with a given asset and/or account in reverse block height order.

**Request:**

-   *requestType* is *getTrades*
-   *asset* is the asset ID (optional)
-   *account* is the account ID (optional if *asset* provided)
-   *firstIndex* is a zero-based index to the first trade to retrieve (optional)
-   *lastIndex* is a zero-based index to the last trade to retrieve (optional)
-   *timestamp* is the earliest block (in seconds since the genesis block) to retrieve (optional)
-   *includeAssetInfo* is *true* if the *decimals* and *name* fields are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *trades* (A) is an array of trade objects with the following fields for each trade:
    -   *seller* (S) is the account number of the seller
    -   *quantityQNT* (S) is the quantity (in QNT) of the asset traded
    -   *bidOrder* (S) is the bid order ID
    -   *sellerRS* (S) is the Reed-Solomon address of the seller
    -   *buyer* (S) is the account number of the buyer
    -   *priceNQT* (S) is the trade price (in NQT, the ask price for a buy or the bid price for a sell)
    -   *askOrder* (S) is the ask order ID
    -   *buyerRS* (S) is the Reed-Solomon address of the buyer
    -   *decimals* (N) is the number of decimal places used by the asset
    -   *name* (S) is the name of the asset (if *includeAssetInfo* is *true*)
    -   *block* (S) is the block ID of the trade (if *includeAssetInfo* is *true*)
    -   *asset* (S) is the asset ID
    -   *askOrderHeight* (N) is the block height of the ask order
    -   *bidOrderHeight* (N) is the block height of the bid order
    -   *tradeType* (S) is the trade type (*sell* or *buy*, where *buy* implies that the bid occurred after the ask, or if in the same block, has a greater order ID)
    -   *timestamp* (N) is the timestamp (in seconds since the genesis block) of the trade block
    -   *height* (N) is the height of the trade block
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Trades](the-burst-api-examples.md#get-trades) example.

### Issue Asset

Create an asset on the exchange. POST only.

**Request:** Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

-   *requestType* is *issueAsset*
-   *name* is the name of the asset
-   *description* is a url-encoded description of the asset in UTF-8 with a maximum length of 1000 bytes (optional)
-   *quantityQNT* is the total amount (in QNT) of the asset in existence
-   *decimals* is the number of decimal places used by the asset (optional, zero default)

**Response:** Refer to [Create Transaction Response](#create-transaction-response). The transaction ID is also the asset ID.

**Example:** Refer to [Issue Asset](the-burst-api-examples.md#issue-asset) example.

### Place Order

Place an asset order. POST only.

**Request:** Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

-   *requestType* is either *placeBidOrder* or *placeAskOrder*
-   *asset* is the asset ID of the asset being ordered
-   *quantityQNT* is the amount (in QNT) of the asset being ordered
-   *priceNQT* is the bid/ask price (in NQT)

**Response:** Refer to [Create Transaction Response](#create-transaction-response). The transaction ID is also the order ID.

**Example:** Refer to [Place Order](the-burst-api-examples.md#place-order) example.

#### Place Ask Order

Refer to [Place Order](#place-order).

#### Place Bid Order

Refer to [Place Order](#place-order).

### Transfer Asset

Transfer a quantity of an asset from one account to another. POST only.

**Request:** Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

-   *requestType* is *transferAsset*
-   *recipient* is the recipient account ID
-   *recipientPublicKey* is the public key of the recipient account (optional, enhances security of a new account)
-   *asset* is the ID of the asset being transferred
-   *quantityQNT* is the amount (in QNT) of the asset being transferred

**Response:** Refer to [Create Transaction Response](#create-transaction-response). The transaction ID is also the transfered asset ID.

**Example:** Refer to [Transfer Asset](the-burst-api-examples.md#transfer-asset) example.

Block Operations
----------------

### Get Block

Get a block object given a block ID or block height.

**Request:**

-   *requestType* is *getBlock*
-   *block* is the block ID (optional)
-   *height* is the block height (optional if *block* provided)
-   *timestamp* is the timestamp (in seconds since the genesis block) of the block (optional if *height* provided)
-   *includeTransactions* is *true* to include transaction details (optional)
-   *includeExecutedPhased* is *true* to include approved and executed phased transactions (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** *block* overrides *height* which overrides *timestamp*.

**Response:**

-   *previousBlockHash* (S) is the 32-byte hash of the previous block
-   *payloadLength* (N) is the length (in bytes) of all transactions included in the block
-   *totalAmountNQT* (S) is the total amount (in NQT) of the transactions in the block
-   *generationSignature* (S) is the 32-byte generation signature of the generating account
-   *generator* (S) is the generating account number
-   *generatorPublicKey* (S) is the 32-byte public key of the generating account
-   *baseTarget* (S) is the base target for the next block generation
-   *payloadHash* (S) is the 32-byte hash of the payload (all transactions)
-   *generatorRS* (S) is the Reed-Solomon address of the generating account
-   *nextBlock* (S) is the next block ID
-   *numberOfTransactions* (N) is the number of transactions in the block
-   *blockSignature* (S) is the 64-byte block signature
-   *transactions* (A) is an array of transaction IDs or transaction objects (if *includeTransactions* provided, refer to [Get Transaction](#get-transaction) for details)
-   *executedPhasedTransactions* (A) is an array of transaction IDs or transaction objects (if *includeExecutedPhased* provided, refer to [Get Transaction](#get-transaction) for details)
-   *version* (N) is the block version
-   *totalFeeNQT* (S) is the total fee (in NQT) of the transactions in the block
-   *previousBlock* (S) is the previous block ID
-   *cumulativeDifficulty* (S) is the cumulative difficulty for the next block generation
-   *block* (S) is the block ID
-   *height* (N) is the zero-based block height
-   *timestamp* (N) is the timestamp (in seconds since the genesis block) of the block
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Block](the-burst-api-examples.md#get-block) example.

### Get Block Id

Get a block ID given a block height.

**Request:**

-   *requestType* is *getBlockId*
-   *height* is the block height
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *block* (S) is the block ID
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Block Id](the-burst-api-examples.md#get-block-id) example.

### Get Blocks

Get blocks from the blockchain in reverse block height order.

**Request:**

-   *requestType* is *getBlocks*
-   *timestamp* is the earliest block (in seconds since the genesis block) to retrieve (optional)
-   *firstIndex* is first block to retrieve (optional, default is zero or the last block on the blockchain)
-   *lastIndex* is the last block to retrieve (optional, default is *firstIndex* + 99)
-   *includeTransactions* is *true* to include transaction details (optional)
-   *includeExecutedPhased* is *true* to include approved and executed phased transactions (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *blocks* (A) is an array of blocks retrieved (refer to [Get Block](#get-block) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Blocks](the-burst-api-examples.md#get-blocks) example.

### Get EC Block

Get Economic Cluster block data.

**Request:**

-   *requestType* is *getECBlock*
-   *timestamp* is the timestamp (in seconds since the genesis block) of the EC block (optional, default (or zero) is the current timestamp)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** If *timestamp* is more than 15 seconds before the timestamp of the last block on the blockchain, *errorCode* 4 is returned.

**Response:**

-   *ecBlockHeight* (N) is the EC block height
-   *ecBlockId* (S) is the EC block ID
-   *timestamp* (N) is the timestamp (in seconds since the genesis block) of the EC block
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get EC Block](the-burst-api-examples.md#get-ec-block) example.

Digital Goods Store Operations
------------------------------

In the [Burst client interface](burst-client-interface.md), the Digital Goods Store (DGS) is referred to as [Marketplace](marketplace.md).

### DGS Delisting

Delist a listed product. POST only.

**Request:** Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

-   *requestType* is *dgsDelisting*
-   *goods* is the goods ID

**Response:** Refer to [Create Transaction Response](#create-transaction-response).

**Example:** Refer to [DGS Delisting](the-burst-api-examples.md#dgs-delisting) example.

### DGS Delivery

Deliver a product. POST only.

**Request:** Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

-   *requestType* is *dgsDelivery*
-   *purchase* is the purchase order ID
-   *discountNQT* is a discount (in NQT) off the selling price (optional, default is zero)
-   *goodsToEncrypt* is the product, a text or a hex string to be encrypted (optional if *goodsData* provided)
-   *goodsIsText* is *false* if *goodsToEncrypt* is a hex string (optional)
-   *goodsData* is AES-encrypted (using [Encrypt To](#encrypt-to)) *goodsToEncrypt*, up to 1000 bytes long (required only if *secretPhrase* is omitted)
-   *goodsNonce* is the unique nonce associated with the encrypted data (required only if *secretPhrase* is omitted)

**Note:** If the encrypted goods data is longer than 1000 bytes, use a prunable encrypted message to deliver the goods.

**Response:** Refer to [Create Transaction Response](#create-transaction-response).

**Example:** Refer to [DGS Delivery](the-burst-api-examples.md#dgs-delivery) example.

### DGS Feedback

Give feedback about a purchased product after delivery. POST only.

**Request:** Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

-   *requestType* is *dgsFeedback*
-   *purchase* is the purchase order ID
-   *message* is unencrypted (public) feedback text up to 1000 bytes

**Note**: The unencrypted *message* parameter is used for public feedback, but in addition or instead, an encrypted message can be used for private feedback to the seller and/or an encrypted message can be sent to self (buyer) although the current [BRS client](burst-client-interface.md) does not recognize non-public feedback messages.

**Response:** Refer to [Create Transaction Response](#create-transaction-response).

**Example:** Refer to [DGS Feedback](the-burst-api-examples.md#dgs-feedback) example.

### DGS Listing

List a product in the DGS by creating a listing transaction. POST only.

**Request:** Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

-   *requestType* is *dgsListing*
-   *name* is the name of the product up to 100 characters in length
-   *description* is a description of the product up to 1000 characters in length
-   *tags* are up to three comma separated keywords describing the product up to 100 characters in length (optional)
-   *quantity* is the quantity of the product for sale
-   *priceNQT* is the price (in NQT) of the product

**Response:** Refer to [Create Transaction Response](#create-transaction-response). The transaction ID is also the goods ID.

**Example:** Refer to [DGS Listing](the-burst-api-examples.md#dgs-listing) example.

### DGS Price Change

Change the price of a listed product. POST only.

**Request:** Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

-   *requestType* is *dgsPriceChange*
-   *goods* is the goods ID of the product
-   *priceNQT* is the new price of the product

**Response:** Refer to [Create Transaction Response](#create-transaction-response).

**Example:** Refer to [DGS Price Change](the-burst-api-examples.md#dgs-price-change) example.

### DGS Purchase

Purchase a product for sale. POST only.

**Request:** Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

-   *requestType* is *dgsPurchase*
-   *goods* is the goods ID of the product
-   *priceNQT* is the price of the product
-   *quantity* is the quantity to be purchased
-   *deliveryDeadlineTimestamp* is the timestamp (in seconds since the genesis block) by which delivery of the product must occur

**Response:** Refer to [Create Transaction Response](#create-transaction-response). The transaction ID is also the purchase order ID.

**Example:** Refer to [DGS Purchase](the-burst-api-examples.md#dgs-purchase) example.

### DGS Quantity Change

Change the quantity of a listed product. POST only.

**Request:** Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

-   *requestType* is *dgsQuantityChange*
-   *goods* is the goods ID of the product
-   *deltaQuantity* is the change in the quantity of the product for sale (use negative numbers for a decrease in quantity)

**Response:** Refer to [Create Transaction Response](#create-transaction-response).

**Example:** Refer to [DGS Quantity Change](the-burst-api-examples.md#dgs-quantity-change) example.

### DGS Refund

Refund a purchase. POST only.

**Request:** Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

-   *requestType* is *dgsRefund*
-   *purchase* is the purchase order ID
-   *refundNQT* is the amount (in NQT) of the refund

**Response:** Refer to [Create Transaction Response](#create-transaction-response).

**Example:** Refer to [DGS Refund](the-burst-api-examples.md#dgs-refund) example.

### Get DGS Good

Get a DGS product given a goods ID.

**Request:**

-   *requestType* is *getDGSGood*
-   *goods* is the goods ID of the product
-   *includeCounts* is *true* if the fields beginning with numberOf... are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *seller* (S) is the seller's account ID
-   *quantity* (N) is the quantity of the product remaining for sale
-   *goods* (S) is the ID of the product
-   *description* (S) is the description of the product
-   *sellerRS* (S) is the Reed-Solomon address of the seller's account
-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *delisted* (B) is *true* if the product has been delisted, *false* otherwise
-   *parsedTags* (A) is an array of up to three tag strings, parsed from the *tags* field
-   *tags* (S) is the comma separated list of tags provided by the seller when the listing was created
-   *priceNQT* (S) is the current price of the product
-   *numberOfPublicFeedbacks* (N) is the number of public feedbacks given for the product
-   *name* (S) is the name of the product
-   *numberOfPurchases* (N) is the number of purchases of the product
-   *timestamp* (N) is the timestamp (in seconds since the genesis block) of the creation of the product listing
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)

**Example:** Refer to [Get DGS Good](the-burst-api-examples.md#get-dgs-good) example.

### Get DGS Goods

Get DGS products for sale in reverse chronological listing creation order unless a seller is given, then in product name order.

**Request:**

-   *requestType* is *getDGSGoods*
-   *seller* is the account ID of the product seller (optional)
-   *firstIndex* is a zero-based index to the first product to retrieve (optional)
-   *lastIndex* is a zero-based index to the last product to retrieve (optional)
-   *inStockOnly* is *false* if out-of-stock products (zero quantity) are to be retrieved (optional)
-   *hideDelisted* is *true* if delisted products are to be omitted (optional)
-   *includeCounts* is *true* if the fields beginning with *numberOf...* are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** If none of the optional parameters are specified, all in-stock products in the blockchain are retrieved at once, which may take a long time.

**Response:**

-   *goods* (A) is an array of goods (refer to [Get DGS Good](#get-dgs-good) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get DGS Goods](the-burst-api-examples.md#get-dgs-goods) example.

### Get DGS Pending Purchases

Get pending purchase orders given a seller ID in reverse chronological order.

**Request:**

-   *requestType* is *getDGSPendingPurchases*
-   *seller* is the account ID of the seller
-   *firstIndex* is a zero-based index to the first purchase order to retrieve (optional)
-   *lastIndex* is a zero-based index to the last purchase order to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *purchases* (A) is an array of pending purchase orders (refer to [Get DGS Purchase](#get-dgs-purchase) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get DGS Pending Purchases](the-burst-api-examples.md#get-dgs-pending-purchases) example.

### Get DGS Purchase

Get a purchase order given a purchase order ID.

**Request:**

-   *requestType* is *getDGSPurchase*
-   *purchase* is the purchase order ID
-   *sharedKey* is the shared key used to decrypt the message (optional) (see [Get Shared Key](#get-shared-key))
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *seller* (S) is the account number of the seller
-   *quantity* (N) is the quantity of the product to be purchased
-   *feedbackNotes* (A) is an array of AES-encrypted objects, each with *data* (S) and *nonce* (S) fields, in reverse chronological order, if applicable
-   *publicFeedbacks* (A) is an array of feedback strings in reverse chronological order if applicable
-   *pending* (B) is *true* if the *deliveryDeadline* has not passed, *false* otherwise
-   *purchase* (S) is the purchase order ID
-   *goods* (S) is the ID of the product
-   *sellerRS* (S) is the Reed-Solomon address of the seller
-   *buyer* (S) is the account number of the buyer
-   *priceNQT* (S) is the price (in NQT) of the product
-   *deliveryDeadlineTimestamp* (N) is the timestamp (in seconds since the genesis block) by which the product must be delivered
-   *goodsIsText* (B) is *false* if the message is a hex string, otherwise the message is text (optional)
-   *buyerRS* (S) is the Reed-Solomon address of the buyer
-   *refundNQT* (S) is the amount (in NQT) refunded, if applicable
-   *name* (S) is the name of the product
-   *goodsData* (O) is an object with the two fields *data* (S) (the encrypted product hex string) and *nonce* (S), if the product has been delivered
-   *timestamp* (N) is the timestamp (in seconds since the genesis block) of the purchase order
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get DGS Purchase](the-burst-api-examples.md#get-dgs-purchase) example.

### Get DGS Purchases

Get purchase orders given a seller and/or buyer ID in reverse chronological order.

**Request:**

-   *requestType* is *getDGSPurchases*
-   *seller* is the account ID of the seller (optional)
-   *buyer* is the account ID of the buyer (optional if *seller* provided)
-   *firstIndex* is a zero-based index to the purchase order to retrieve (optional)
-   *lastIndex* is a zero-based index to the purchase order to retrieve (optional)
-   *withPublicFeedbacksOnly* is *true* if purchase orders without public feedback are to be omitted (optional)
-   *completed* is *true* if only completed purchase orders are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *purchases* (A) is an array of purchase orders (refer to [Get DGS Purchase](#get-dgs-purchase) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get DGS Purchases](the-burst-api-examples.md#get-dgs-purchases) example.

Forging Operations
------------------

### Lease Balance

[Lease](account-leasing.md) the entire guaranteed balance of BURST to another account, after 1440 confirmations. POST only.

**Request:** Refer to [Create Transaction Request](#create-transaction-request) for common parameters.

-   *requestType* is *leaseBalance*
-   *period* is the lease period (in number of blocks, 1440 minimum)
-   *recipient* is the lessee (recipient) account
-   *recipientPublicKey* is the public key of the lessee (recipient) account (optional, enhances security of a new account)

**Response:** Refer to [Create Transaction Response](#create-transaction-response).

**Example:** Refer to [Lease Balance](the-burst-api-examples.md#lease-balance) example.

Hallmark Operations
-------------------

### Decode Hallmark

Decode a node hallmark.

**Request:**

-   *requestType* is *decodeHallmark*
-   *hallmark* is the hallmark value

**Response:**

-   *valid* (B) is *true* if *host* is less than 100 characters, *weight* &gt; 0 and the embedded signature is verified
-   *weight* (N) is the weight assigned to the hallmark
-   *host* (S) is the IP address or domain name associated with the hallmark
-   *account* (S) is the account number associated with the hallmark
-   *accountRS* (S) is the Reed-Solomon address of the account
-   *date* (S) is the date the hallmark was created, in YYYY-MM-DD format
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Decode Hallmark](the-burst-api-examples.md#decode-hallmark) example.

### Mark Host

Generates a node hallmark. POST only.

**Request:**

-   *requestType* is *markHost*
-   *secretPhrase* is the secret passphrase for the account that will be hallmarked on the node
-   *host* is the IP address or domain name of the node
-   *weight* is the weight to assign to the node
-   *date* is the current date in YYYY-MM-DD format

**Note:** Refer to [Create Hallmark](how-to-createhallmark.md) for details.

**Response:**

-   *hallmark* (S) is the hallmark hex string
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** Refer to [Create Hallmark](how-to-createhallmark.md) for instructions for applying the hallmark to a public node.

**Example:** Refer to [Mark Host](the-burst-api-examples.md#mark-host) example.

#### Generate Hallmark

Refer to [Mark Host](#mark-host).

Networking Operations
---------------------

### Get My Info

Get hostname and address of the requesting node.

**Request:**

-   *requestType* is *getMyInfo*

**Response:**

-   *host* (S) is the node hostname
-   *address* (S) is the node address
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get My Info](the-burst-api-examples.md#get-my-info) example.

### Get Peer

Get information about a given peer.

**Request:**

-   *requestType* is *getPeer*
-   *peer* is the IP address or domain name of the peer (plus optional port)

**Response:**

-   *hallmark* (S) is the hex string of the peer's hallmark, if it is defined
-   *downloadedVolume* (N) is the number of bytes downloaded by the peer
-   *address* (S) the IP address or DNS name of the peer
-   *weight* (N) is the peer's weight value
-   *uploadedVolume* (N) is the number of bytes uploaded by the peer
-   *version* (S) is the version of the software running on the peer
-   *platform* (S) is a string representing the peer's platform
-   *lastUpdated* (N) is the timestamp (in seconds since the genesis block) of the last peer status update
-   *blacklisted* (B) is *true* if the peer is blacklisted
-   *services* (A) is an array of strings with the services the node provides
-   *blacklistingCause* (S) is the cause of blacklisting (if *blacklisted* is *true*)
-   *announcedAddress* (S) is the name that the peer announced to the network (could be a DNS name, IP address, or any other string)
-   *application* (S) is the name of the software application, typically *BRS*
-   *state* (N) defines the state of the peer: 0 for NON\_CONNECTED, 1 for CONNECTED, or 2 for DISCONNECTED
-   *shareAddress* (B) is *true* if the address is allowed to be shared with other peers
-   *inbound* (B) is *true* if the peer has made a request to this node
-   *inboundWebSocket* (B) is *true* if an inbound websocket has been established from this node
-   *outboundWebSocket* (B) is *true* if an outbound websocket has been established to this node
-   *lastConnectAttempt* (B) is the timestamp (in seconds since the genesis block) of the last connection attempt to the peer
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Peer](the-burst-api-examples.md#get-peer) example.

### Get Peers

Get a list of peer IP addresses.

**Request:**

-   *requestType* is *getPeers*
-   *active* is *true* for active (not NON\_CONNECTED) peers only (optional, if *true* overrides *state*)
-   *state* is the state of the peers, one of *NON\_CONNECTED*, *CONNECTED*, or *DISCONNECTED* (optional)
-   *includePeerInfo* is *true* to include peer detail as in [Get Peer](#get-peer)
-   *service* to filter on a specific service

**Note:** If neither *active* nor *state* is specified, all known peers are retrieved.

**Response:**

-   *peers* (A) is an array of peer addresses
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Peers](the-burst-api-examples.md#get-peers) example.

Server Information Operations
-----------------------------

### Get Blockchain Status

Get the blockchain status.

**Request:**

-   *requestType* is *getBlockchainStatus*

**Response:**

-   *currentMinRollbackHeight* (N) is the current minimum rollback height
-   *numberOfBlocks* (N) is the number of blocks in the blockchain (height + 1)
-   *isTestnet* (B) is *true* if the node is connected to testnet, *false* otherwise
-   *includeExpiredPrunable* (B) is the value of the *nxt.includeExpiredPrunable* property
-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *version* (S) is the application version
-   *maxRollback* (N) is the value of the *nxt.maxRollback* property
-   *lastBlock* (S) is the last block ID on the blockchain
-   *application* (S) is application name, typically *BRS*
-   *isScanning* (B) is *true* if the blockchain is being scanned by the application, *false* otherwise
-   *isDownloading* (B) is *true* if a download is in progress, *false* otherwise; *true* when a batch of more than 10 blocks at once has been downloaded from a peer, reset to *false* when an attempt to download more blocks from a peer does not result in any new blocks
-   *cumulativeDifficulty* (S) is the cumulative difficulty
-   *lastBlockchainFeederHeight* (N) is the height of the last blockchain of greatest cumulative difficulty obtained from a peer
-   *maxPrunableLifetime* (N) is the maximum prunable lifetime (in seconds)
-   *time* (N) is the current timestamp (in seconds since the genesis block)
-   *lastBlockchainFeeder* (S) is the address or announced address of the peer providing the last blockchain of greatest cumulative difficulty
-   *blockchainState* (S) Current state of this node's blockchain (UP\_TO\_DATE or DOWNLOADING)

**Example:** Refer to [Get Blockchain Status](the-burst-api-examples.md#get-blockchain-status) example.

### Get Constants

Get all defined constants.

**Request:**

-   *requestType* is *getConstants*

**Response:**

-   *maxBlockPayloadLength* (N) is the maximum block payload length (in bytes)
-   *maxArbitraryMessageLength* (N) is the maximum length (in bytes) of an arbitrary message
-   *maxPrunableMessageLength* (N) is the maximum length (in bytes) of a prunable message
-   *maxTaggedDataDataLength* (N) is the maximum length (in bytes) of tagged data
-   *maxPhasingDuration* (N) is the maximum allowed phasing duration in block height
-   *epochBeginning* (N) is the time in milliseconds when genesis block was created
-   *genesisAccountId* (S) is the genesis account number
-   *genesisBlockId* (S) is the genesis block ID
-   *transactionTypes* (A) is an array of defined transaction types and subtypes (refer to the example below)
-   *transactionSubTypes* (A) is an array of defined transaction subtypes and subtypes (refer to the example below)
-   *peerStates* (A) is an array of defined peer states (refer to the example below)
-   *currencyTypes* (A) is an array of defined currency types (refer to the example below)
-   *disabledAPIs* (A) is an array of configured disabled apis (refer to the example below)
-   *apiTags* (A) is an array of defined api tags (refer to the example below)
-   *disabledAPITags* (A) is an array of configured disabled api tags (refer to the example below)
-   *votingModels* (A) is an array of defined voting models (refer to the example below)
-   *holdingTypes* (A) is an array of defined holding types (refer to the example below)
-   *minBalanceModels* (A) is an array of defined minimum balance models (refer to the example below)
-   *shufflingStages* (A) is an array of defined shuffling stages (refer to the example below)
-   *shufflingParticipantStates* (A) is an array of defined shuffling participant states (refer to the example below)
-   *hashAlgorithms* (A) is an array of defined hash algorithms (refer to the example below)
-   *mintingHashAlgorithms* (A) is an array of defined minting hash algorithms (refer to the example below)
-   *phasingHashAlgorithms* (A) is an array of defined phasing hash algorithms (refer to the example below)
-   *requestTypes* (A) is an array of decined request types (refer to the example below)

**Example:** Refer to [Get Constants](the-burst-api-examples.md#get-constants) example.

### Get State

Get the state of the server node and network.

**Request:**

-   *requestType* is *getState*
-   *includeCounts* is *true* if the fields beginning with *numberOf...* are to be included (optional); password protected like the [Debug Operations](#debug-operations) if *true*.

**Response:**

-   *numberOfPeers* (N) is the number of known peers on the network
-   *numberOfGoods* (N) is the number of DGS goods in the blockchain
-   *numberOfPolls* (N) is the number of polls in the blockchain
-   *numberOfUnlockedAccounts* (N) is the number of unlocked accounts on this node
-   *numberOfTransfers* (N) is the number of AE transfers in the blockchain
-   *includeExpiredPrunable* (B) is the value of the *nxt.includeExpiredPrunable* property
-   *numberOfOrders* (N) is the number of AE orders in the blockchain
-   *numberOfTransactions* (N) is the number of transactions in the blockchain
-   *maxMemory* (N) is the maximum amount of memory the node may use (in Bytes)
-   *maxRollback* (N) is the value of the *nxt.maxRollback* property
-   *numberOfOffers* (N) is the number of buy currency offers in the blockchain
-   *isDownloading* (B) is *true* if a download is in progress, *false* otherwise; *true* when a batch of more than 10 blocks at once has been downloaded from a peer, reset to *false* when an attempt to download more blocks from a peer does not result in any new blocks
-   *isScanning* (B) is *true* if this node is scanning the blockchain, *false* otherwise
-   *cumulativeDifficulty* (S) is the current cumulative forging difficulty
-   *numberOfCurrencies* (N) is the number of currencies in the blockchain
-   *numberOfAssets* (N) is the number of AE assets in the blockchain
-   *numberOfPrunableMessages* (N) is the number of prunable messages in the blockchain
-   *freeMemory* (N) is the amount of free memory on this node (in Bytes)
-   *peerPort* (N) is the port used for connecting to peers
-   *numberOfVotes* (N) is the number of votes in the blockchain
-   *availableProcessors* (N) is the number of processors on this node
-   *numberOfTaggedData* (N) is the number of tagged data in the blockchain
-   *numberOfActiveAccountLeases* (N) is the number of active account leases in the blockchain
-   *numberOfAccountLeases* (N) is the total number of account leases including scheduled leases (first scheduled lease only) in the blockchain
-   *numberOfAccounts* (N) is the number of accounts in the blockchain
-   *numberOfDataTags* (N) is the number of data tags in the blockchain
-   *needsAdminPassword* (B) is *true* if the *nxt.disableAdminPassword* property is *false*
-   *currentMinRollbackHeight* (N) is the current minimum rollback height
-   *numberOfBlocks* (N) is the number of blocks (height + 1) in the blockchain
-   *isTestnet* (B) is *true* if the node is connected to testnet, *false* otherwise
-   *numberOfCurrencyTransfers* (N) is the number of currency transfers in the blockchain
-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *numberOfPhasedTransactions* (N) is the number of phased transactions in the blockchain
-   *version* (S) is the software version on this node
-   *numberOfBidOrders* (N) is the number of AE bid orders in the blockchain
-   *lastBlock* (S) is the last block address
-   *totalMemory* (N) is the amount of memory this node is using (in Bytes)
-   *application* (S) is the name of the software running on this node (typically *BRS*)
-   *numberOfAliases* (N) is the number of aliases in the blockchain
-   *numberOfActivePeers* (N) is the number of active peers on the network
-   *lastBlockchainFeederHeight* (N) is the height of the last blockchain feeder
-   *maxPrunableLifetime* (N) is the maximum prunable lifetime (in seconds)
-   *numberOfExchanges* (N) is the number of currency exchanges in the blockchain
-   *numberOfTrades* (N) is the number of AE trades in the blockchain
-   *numberOfPurchases* (N) is the number of DGS purchases in the blockchain
-   *numberOfTags* (N) is the number of DGS tags in the blockchain
-   *time* (N) is the current node time (in seconds since the genesis block)
-   *numberOfAskOrders* (N) is the number of AE ask orders in the blockchain
-   *lastBlockchainFeeder* (S) is the announced name of the feeder of the last blockchain
-   *isOffline* (B) is *true* if this node is connected to other peers, *false* otherwise

**Note:** AE is Asset Exchange, DGS is Digital Goods Store

**Example:** Refer to [Get State](the-burst-api-examples.md#get-state) example.

### Get Time

Get the current time.

**Request:**

-   *requestType* is *getTime*

**Response:**

-   *time* (N) is the current time (in seconds since the genesis block).
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Time](the-burst-api-examples.md#get-time) example.

Token Operations
----------------

### Decode Token

Validate a token without requiring the transmission of a secret passphrase.

**Request:**

-   *requestType* is *decodeToken*
-   *website* is the signed text, typically an authorized URL
-   *token* is the token generated by [Generate Token](#generate-token)

**Response:**

-   *account* (S) is the account number that generated the token
-   *accountRS* (S) is the Reed-Solomon address of the account
-   *timestamp* (N) is the time (in seconds since the genesis block) that the token was created
-   *valid* (B) is *true* if token is valid, *false* otherwise
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** Since *token* contains the token generator's public key and digital signature, *website* can be validated as authorized by the owner of the public key, and the public key determines the account ID.

**Example:** Refer to [Decode Token](the-burst-api-examples.md#decode-token) example.

### Generate Token

Generate a token. POST only.

**Request:**

-   *requestType* is *generateToken*
-   *secretPhrase* is the passphrase of the account generating the token
-   *website* is a web site URL for which authorization should be granted, or general text to be digitally signed

**Note:** *website* is typically a URL (with the leading <http://> unnecessary) that an account owner signs with his *secretPhrase* (private key) to bind the account to the URL, but *website* can be any text that the owner wishes to sign.

**Response:**

-   *token* (S) is a 160 character string representing the 100-byte token which consists of a 32-byte public key, a 4-byte timestamp, and a 64-byte signature
-   *account* (S) is the account number corresponding to *secretPhrase*
-   *accountRS* (S) is the Reed-Solomon address of the account
-   *timestamp* (N) is the time (in seconds since the genesis block) that the token was generated
-   *valid* (B) is *true* if token is valid, *false* otherwise
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** Since *token* contains the token generator's public key and signature, the *website* can be validated as authorized by the owner of the public key using [Decode Token](#decode-token).

**Example:** Refer to [Generate Token](the-burst-api-examples.md#generate-token) example.

Transaction Operations
----------------------

### Broadcast Transaction

Broadcast a transaction to the network. POST only.

**Request:**

-   *requestType* is *broadcastTransaction*
-   *transactionBytes* is the bytecode of a signed transaction (optional)
-   *transactionJSON* is the transaction object (optional if *transactionBytes* provided)
-   *prunableAttachmentJSON* is the attachment object embedded in *transactionJSON* containing a prunable message (required if *transactionJSON* not provided because *transactionBytes* never includes prunable data)

**Response:**

-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *fullHash* (S) is the full hash of the signed transaction
-   *transaction* (S) is the transaction ID

**Example:** Refer to [Broadcast Transaction](the-burst-api-examples.md#broadcast-transaction) example.

### Calculate Full Hash

Calculate the full hash of a transaction.

**Request:**

-   *requestType* is *calculateFullHash*
-   *unsignedTransactionJSON* is the unsigned transaction JSON object (optional)
-   *unsignedTransactionBytes* are the unsigned bytes of a transaction (optional if *unsignedTransactionJSON* is provided)
-   *signatureHash* is a SHA-256 hash of the transaction signature

**Response:**

-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *fullHash* (S) is the full hash of the signed transaction

**Example:** Refer to [Calculate Full Hash](the-burst-api-examples.md#calculate-full-hash) example.

### Get Transaction

Get a transaction object given a transaction ID.

**Request:**

-   *requestType* is *getTransaction*
-   *transaction* is the transaction ID (optional)
-   *fullHash* is the full hash of the transaction (optional if transaction ID is provided)
-   *includePhasingResult* is *true* to retrieve execution status of each phased transaction (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *sender* (S) is the account ID of the sender
-   *senderRS* (S) is the Reed-Solomon address of the sender
-   *feeNQT* (S) is the fee (in NQT) of the transaction
-   *amountNQT* (S) is the amount (in NQT) of the transaction
-   *transactionIndex* (N) is a zero-based index giving the order of the transaction in its block
-   *timestamp* (N) is the time (in seconds since the genesis block) of the transaction
-   *referencedTransactionFullHash* (S) is the full hash of a transaction referenced by this one, omitted if no previous transaction is referenced
-   *confirmations* (N) is the number of transaction confirmations
-   *subtype* (N) is the transaction subtype (refer to [Get Constants](#get-constants) for a current list of subtypes)
-   *block* (S) is the ID of the block containing the transaction
-   *blockTimestamp* (N) is the timestamp (in seconds since the genesis block) of the block
-   *height* (N) is the height of the block in the blockchain
-   *senderPublicKey* (S) is the public key of the sending account for the transaction
-   *type* (N) is the transaction type (refer to [Get Constants](#get-constants) for a current list of types)
-   *deadline* (N) is the deadline (in minutes) for the transaction to be confirmed
-   *signature* (S) is the digital signature of the transaction
-   *recipient* (S) is the account number of the recipient, if applicable
-   *recipientRS* (S) is the Reed-Solomon address of the recipient, if applicable
-   *fullHash* (S) is the full hash of the signed transaction
-   *signatureHash* (S) is a SHA-256 hash of the transaction signature
-   *approved* (B) is a boolean indicating if the transaction is approved (only included when *includePhasingResult* is true and the transaction is phased)
-   *result* (S) is a string containing the result of the transaction (only included when *includePhasingResult* is true and the transaction is phased)
-   *executionHeight* (N) is the height the transaction was executed (only included when *includePhasingResult* is true and the transaction is phased)
-   *transaction* (S) is the transaction ID
-   *version* (N) is the transaction version number
-   *phased* (B) is *true* if the transaction is phased, *false* otherwise
-   *ecBlockId* (N) is the economic clustering block ID
-   *ecBlockHeight* (N) is the economic clustering block height
-   *attachment* (O) is an object containing any additional data needed for the transaction, if applicable
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** The *block*, *blockTimestamp* and *confirmations* fields are omitted for unconfirmed transactions. Double-spending transactions are not retrieved.

**Example:** Refer to [Get Transaction](the-burst-api-examples.md#get-transaction) example.

### Get Transaction Bytes

Get the bytecode of a transaction.

**Request:**

-   *requestType* is *getTransactionBytes*
-   *transaction* is a transaction ID
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *confirmations* (N) is the number of transaction confirmations
-   *transactionBytes* (S) are the bytes contained in the transaction
-   *unsignedTransactionBytes* (S) are the unsigned bytes contained in the transaction
-   *prunableAttachmentJSON* (O) is the prunable attachment JSON object, if applicable, because *transactionBytes* never includes prunable data
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Transaction Bytes](the-burst-api-examples.md#get-transaction-bytes) example.

### Parse Transaction

Get a transaction object given a (signed or unsigned) transaction bytecode, or re-parse a transaction object. Verify the signature.

**Request:**

-   *requestType* is *parseTransaction*
-   *transactionBytes* is the signed or unsigned bytecode of the transaction (optional)
-   *transactionJSON* is the transaction object (optional if transactionBytes is included)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:** Refer to [Get Transaction](#get-transaction) for additional fields.

-   *verify* (B) is *true* if the signature is verified, *false* otherwise

**Example:** Refer to [Parse Transaction](the-burst-api-examples.md#parse-transaction) example.

### Sign Transaction

Calculates the full hash, signature hash, and transaction ID of an unsigned transaction.

**Request:**

-   *requestType* is *signTransaction*
-   *unsignedTransactionJSON* is the *transactionJSON* field of the transaction, without a signature subfield
-   *unsignedTransactionBytes* is the *unsignedTransactionBytes* field of the transaction (optional, if *unsignedTransactionJSON* provided)
-   *prunableAttachmentJSON* is a prunable attachment JSON object, if applicable, because *unsignedTransactionBytes* never includes prunable data (optional if the transaction has already been broadcast and the prunable message can still be retrieved from the database)
-   *secretPhrase* is the secret passphrase of the signing account
-   *validate* is *false* to skip validation of the transaction bytes being signed (useful on nodes where the full blockchain is not downloaded)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *signatureHash* (S) is a SHA-256 hash of the transaction signature, used in escrow transactions
-   *verify* (B) is *true* the signature is verified, *false* if not
-   *transactionJSON* (O) is the signed transaction JSON object
-   *transactionBytes* (S) are the signed transaction bytes
-   *fullHash* (S) is the full hash of the signed transaction
-   *prunableAttachmentJSON* (O) is the prunable attachment JSON object, if applicable, because *transactionBytes* never includes prunable data
-   *transaction* (S) is the transaction ID, derived from the *fullHash*
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Sign Transaction](the-burst-api-examples.md#sign-transaction) example.

Utilities
---------

### Long Convert

Converts an ID to the signed long integer representation used internally.

**Request:**

-   *requestType* is *longConvert*
-   *id* is a numerical ID, in decimal form but equivalent to an 8-byte unsigned integer as produced by SHA-256 hashing

**Response:**

-   *stringId* (S) is the numerical ID
-   *longId* (S) is the signed long integer (8-bytes) representation of the ID used internally, returned as a string
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** Java does not support unsigned integers, so any unsigned ID (such as a block ID) visible in the [BRS client](burst-client-interface.md) is represented internally as a signed integer.

**Example:** Refer to [Long Convert](the-burst-api-examples.md#long-convert) example.

### RS Convert

Get both the Reed-Solomon account address and the account number given an account ID.

**Request:**

-   *requestType* is *rsConvert*
-   *account* is an account ID (either RS address or number)

**Response:**

-   *accountRS* (S) is the Reed-Solomon address of the account
-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *account* (S) is the account number


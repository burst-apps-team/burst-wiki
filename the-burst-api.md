**Caution! Most of this text is still based on the original Nxt API documentation. While Burst and Nxt share a fair amount of similarities, Burst starts to differ as it's being developed further. The changes may or may not be already documented on this page. Also, the documentation here refers to the version or later**

Description
-----------

The Burst API allows interaction with Burst nodes using HTTP requests to port 8123. Most HTTP requests can use either the GET or POST methods, but some API calls accept only the POST method for security reasons. Responses are returned as JSON objects.

Each API call is documented below, with definitions given for HTTP request parameters and JSON response fields, followed by an example:

-   The JSON response fields are each followed by one of *S* for string, *A* for array, *O* for object, *N* for number or *B* for boolean.
-   In the examples, the Burst node is represented as *localhost* and requests and responses are formatted for easy reading; line breaks and spaces are not actually used except in some parameter values. All requests are in URL format which implies the HTTP GET method. When GET is allowed, the URL can be entered into a browser URL field but proper URL encoding is usually required (e.g., spaces in a parameter value must be replaced by *+* or *%20*). Otherwise, the URL should be used as a guide to preparing an HTTP POST request using cURL, for example.

All API calls can be viewed and tested on the TestNet at <https://wallet.dev.burst-test.net/test>. For specific API calls, use <https://wallet.dev.burst-test.net/test?requestType>=*specificRequestType*.

Table Of Contents
-----------------

\_\_TOC\_\_

General Notes
-------------

### Genesis Block

Many API requests make reference to the .

### Flexible Account IDs

All API requests that require an account ID accept either an account number or the corresponding [Reed-Solomon address](rs-address-format.md).

### Quantity Units BURST, NQT and QNT

The Burst system has a currency BURST used to quantify value in the system. Like all currencies, BURST circulates in the system, moving from one user to another by payments and purchases. Also, a small fee is required for every transaction, including those in which no BURST is transfered, such as simple messages. This fee goes to the owner of the node that forges (generates) the new block containing the transaction that is accepted onto the blockchain.

Yet internally, the currency is still stored in integer form in units of NQT or NxtQuant, where 1 BURST = 10<sup>8</sup> NQT. All parameters and fields in the API involving a quantity of BURST are denominated in units of NQT, for example *feeNQT*. The only exception is the field *effectiveBalanceNXT*, used in forging calculations.

Other assets can be created within Burst using [Issue Asset](the-burst-api-issue-asset.md). The issuer must specify the number of decimal places to use in quantifying the asset, and the amount of the asset to create in generic units of QNT or Quant, distinct from NQT. Quantities of assets are stored internally as integers in units of QNT, and assets are priced in NQT per QNT.

### Creating Unsigned Transactions

All API requests that create a new transaction will accept either a *secretPhrase* or a *publicKey* parameter:

-   If *secretPhrase* is supplied, a transaction is created, signed at the server, and broadcast by the server as usual.
-   If only a *publicKey* parameter is supplied as a 64-digit (32-byte) hex string, the transaction will be prepared by the server and returned in the JSON response as *transactionJSON* without a signature. This JSON object along with *secretPhrase* can then be supplied to [Sign Transaction](the-burst-api-sign-transaction.md) as *unsignedTransactionJSON* and the resulting signed *transactionJSON* can then be supplied to [Broadcast Transaction](the-burst-api-broadcast-transaction.md). This sequence allows for offline signing of transactions so that *secretPhrase* never needs to be exposed.
-   *unsignedTransactionBytes* may be used instead of unsigned *transactionJSON* when there is no encrypted message. Messages to be encrypted require the *secretPhrase* for encryption and so cannot be included in *unsignedTransactionBytes*.

### Escrow Operations

All API requests that create a new transaction will accept an optional *referencedTransactionFullHash* parameter which creates a chained transaction, meaning that the new transaction cannot be confirmed unless the referenced transaction is also confirmed. This feature allows a simple way of transaction escrow:

-   Alice prepares and signs a transaction A, but doesn't broadcast it by setting the *broadcast* parameter to *false*. She sends to Bob the *unsignedTransactionBytes*, the *fullHash* of the transaction, and the *signatureHash*. All of those are included in the JSON returned by the API request. (Warning: make sure not to send the signed *transactionBytes*, or the *signature* itself, as then Bob can just broadcast transaction A himself).
-   Bob prepares, signs and broadcasts transaction B, setting the *referencedTransactionFullHash* parameter to the *fullHash* of A provided by Alice. He can verify that this hash indeed belongs to the transaction he expects from Alice, by using [Calculate Full Hash](the-burst-api-calculate-full-hash.md), which takes *unsignedTransactionBytes* and *signatureHash* (both of which Alice has also sent to Bob). He can also use [Parse Transaction](the-burst-api-parse-transaction.md) to decode the unsigned bytes and inspect all transaction fields.
-   Transaction B is accepted in the unconfirmed transaction pool, but as long as A is still missing, B will not be confirmed, i.e. will not be included in the blockchain. If A is never submitted, B will eventually expire -- so Bob should make sure to set a long enough deadline, such as the maximum of 32767 minutes.
-   Once in the unconfirmed transactions pool, Bob has no way of recalling B back. So now Alice can safely submit her transaction A, by just broadcasting the *signedTransactionBytes* she got in the first step. Transaction A will get included in the blockchain first, and in the next block Bob's transaction B will also be included.

Note that while the above scheme is good enough for a simple escrow, the blockchain does not enforce that if A is included, B will also be included. It may happen due to forks and blockchain reorganization, that B never gets a chance to be included and expires unconfirmed, while A still remains in the blockchain. However, it is not practically possible for Bob to intentionally cause such chain of events and to prevent B from being confirmed.

### Prunable Data

Prunable data can be removed from the blockchain without affecting the integrity of the blockchain. When a transaction containing prunable data is created, only the 32-byte sha256 hash of the prunable data is included in the *transactionBytes*, not the prunable data itself. The non-prunable signed *transactionBytes* are used to verify the signature and to generate the transaction's *fullHash* and ID; when the prunable part of the transaction is removed at a later time, none of these operations are affected.

Prunable data has a predetermined minimum lifetime of two weeks (24 hours on the [Testnet](testnet.md)) from the timestamp of the transaction. Transactions and blocks received from peer nodes are not accepted if prunable data is missing before this time has elapsed. After this time has elapsed, prunable data is no longer included with transactions and blocks transmitted to peer nodes, and is no longer included in the transaction JSON returned by general-purpose API calls such as Get Transaction; the only way to retrieve it, if still available, is through special-purpose API calls such as Get Prunable Message.

Expired prunable data remains stored in the blockchain until removed at the same time derived tables are trimmed, which occurs automatically every 1000 blocks by default. Use [Trim Derived Tables](the-burst-api-trim-derived-tables.md) to remove expired prunable data immediately.

Prunable data can be preserved on a node beyond the predetermined minimum lifetime by setting the *nxt.maxPrunableLifetime* property to a larger value than two weeks or to *-1* to preserve it indefinitely. To force the node to include such preserved prunable data when transactions and blocks are transmitted to peer nodes, set the *nxt.includeExpiredPrunables* property to *true*, thus making it an archival node.

Currently, there are two varieties of prunable data in the Burst system: prunable [Arbitrary Messages](the-burst-api-arbitrary-message-system-operations.md) and [Tagged Data](the-burst-api-tagged-data-operations.md). Both varities have a maximum prunable data length of 42 kilobytes.

### Properties Files

The behavior of some API calls is affected by property settings loaded from files in the *nxt/conf* directory during Burst server intialization. This directory contains the *nxt-default.properties* and *logging-default.properties* files, both of which contain default property settings along with documentation. A few of the property settings can be obtained while the server is running through the [Get Blockchain Status](the-burst-api-get-blockchain-status.md) and [Get State](the-burst-api-get-state.md) calls.

It is recommended not to modify default properties files because they can be overwritten during software updates. Instead, properties in the default files can be overridden by including them in optional *nxt.properties* and *logging.properties* files in the same directory. For example, a *nxt.properties* file can be created with the contents:

nxt.isTestnet=true

This causes the Burst server to connect to the [Testnet](testnet.md) instead of the Mainnet.

### Admin Password

Some API functions take an adminPassword parameter, which must match the nxt.adminPassword property unless the nxt.disableAdminPassword property is set to true or the API server only listens on the localhost interface (when the nxt.apiServerHost property is set to 127.0.0.1).

All [Debug Operations](the-burst-api-debug-operations.md) require adminPassword since they require some kind of privilege. On some functions adminPassword is used so you can override maximum allowed value for lastIndex parameter, which is set to 100 by default by the nxt.maxAPIRecords property. Giving you the option to retrieve more than objects per request.

### Roaming and Light Client Modes

The remote node to use when in roaming and light client modes is selected randomly, but can be changed manually in the UI, or using the new [set API Proxy Peer](the-burst-api-set-api-proxy-peer.md) API, or forced to a specific peer using the *nxt.forceAPIProxyServerURL* property.

Remote nodes can be blacklisted from the UI, or using the [Blacklist API Proxy Peer](the-burst-api-blacklist-api-proxy-peer.md) API. This blacklisting is independent from peer blacklisting. The API proxy blacklisting period can be set using the *nxt.apiProxyBlacklistingPeriod* property (default 1800000 milliseconds).

API requests that require sending the secret phrase, shared key, or admin password to the server, for features like forging, shuffling, or running a funding monitor, are disabled when in roaming or light client mode.

Create Transaction
------------------

The following API calls create Burst transactions using HTTP POST requests. Follow the links for examples and HTTP POST request parameters specific to each call. Refer to the sections below for [HTTP POST request parameters](the-burst-api-create-transaction-request.md) and [JSON response fields](the-burst-api-create-transaction-response.md) common to all calls that create transactions. Calls in *italics* are phasing-safe (refer to [Get Constants](the-burst-api-get-constants.md) and [Create Phasing Poll](the-burst-api-create-phasing-poll.md))

-   *[Send Money](the-burst-api-send-money.md)*
-   *[Set Account Information](the-burst-api-set-account-information.md)*
-   *[Set Account Property](the-burst-api-set-account-property.md)*
-   [Buy / Sell Alias](the-burst-api-buy--2f-sell-alias.md)
-   [Delete Alias](the-burst-api-delete-alias.md)
-   [Set Alias](the-burst-api-set-alias.md)
-   [Send Message](the-burst-api-send-message.md)
-   *[Cancel Order](the-burst-api-cancel-order.md)*
-   *[Dividend Payment](the-burst-api-dividend-payment.md)*
-   *[Issue Asset](the-burst-api-issue-asset.md)*
-   *[Place Order](the-burst-api-place-order.md)*
-   *[Transfer Asset](the-burst-api-transfer-asset.md)*
-   *[DGS Delisting](the-burst-api-dgs-delisting.md)*
-   [DGS Delivery](the-burst-api-dgs-delivery.md)
-   [DGS Feedback](the-burst-api-dgs-feedback.md)
-   *[DGS Listing](the-burst-api-dgs-listing.md)*
-   [DGS PriceChange](the-burst-api-dgs-pricechange.md)
-   [DGS Purchase](the-burst-api-dgs-purchase.md)
-   [DGS Quantity Change](the-burst-api-dgs-quantity-change.md)
-   [DGS Refund](the-burst-api-dgs-refund.md)
-   *[Lease Balance](the-burst-api-lease-balance.md)*
-   [Currency Buy / Sell](the-burst-api-currency-buy--2f-sell.md)
-   [Currency Mint](the-burst-api-currency-mint.md)
-   [Currency Reserve Claim](the-burst-api-currency-reserve-claim.md)
-   [Currency Reserve Increase](the-burst-api-currency-reserve-increase.md)
-   [Delete Currency](the-burst-api-delete-currency.md)
-   [Issue Currency](the-burst-api-issue-currency.md)
-   [Publish Exchange Offer](the-burst-api-publish-exchange-offer.md)
-   [Transfer Currency](the-burst-api-transfer-currency.md)
-   [Create Poll](the-burst-api-create-poll.md)
-   [Cast Vote](the-burst-api-cast-vote.md)
-   *[Approve Transaction](the-burst-api-approve-transaction.md)*
-   [Extend Tagged Data](the-burst-api-extend-tagged-data.md)
-   [Upload Tagged Data](the-burst-api-upload-tagged-data.md)

### Create Transaction Request

The following HTTP POST parameters are common to all API calls that create transactions:

-   *secretPhrase* is the secret passphrase of the account (optional, but transaction neither signed nor broadcast if omitted)
-   *publicKey* is the public key of the account (optional if *secretPhrase* provided)
-   *feeNQT* is the fee (in NQT) for the transaction:
    -   minimum 1000 NXT for [Issue Asset](the-burst-api-issue-asset.md), unless singleton asset is issued, for which the fee is 1 NXT
    -   2 NXT in base fee for [Set Alias](the-burst-api-set-alias.md), with 2 NXT additional fee for each 32 chars of name plus URI total length, after the first 32 chars
    -   \[25000, 1000, 40\] NXT for \[3-letter, 4-letter, 5-letter\] [Issue Currency](the-burst-api-issue-currency.md)
    -   40 NXT for re-issue of any currency
    -   10 NXT for a [Create Poll](the-burst-api-create-poll.md), including 20 options and total size of poll name plus poll description plus total option length not exceeding 320 chars. For each option above 20, an additional fee of 1 NXT, and for each 32 chars after 320, an additional fee of 2 NXT.
    -   \[2, 21\] NXT for a \[basic, required-minimum-balance\] [Create Phasing Poll](the-burst-api-create-phasing-poll.md). 1 NXT will be added for each option (answer) beyond 20, and 1 NXT for each 32 bytes of hashedSecret or linkedFullHash fields.
    -   1 NXT for the first 32 bytes of a unencrypted non-prunable [message](the-burst-api-send-message.md), 1 NXT for each additional 32 bytes
    -   2 NXT for the first 32 bytes of an encrypted non-prunable [message](the-burst-api-send-message.md), 1 NXT for each additional 32 bytes. The length is measured excluding the nonce and the 16 byte AES initialization vector.
    -   1 NXT for the first 1024 bytes of a prunable [message](the-burst-api-send-message.md), 0.1 NXT for each additional 1024 prunable bytes
    -   1 NXT for [Set Account Info](the-burst-api-set-account-info.md), including 32 chars, with 2 NXT additional fee for each 32 chars
    -   2 NXT for [DGS Listing](the-burst-api-dgs-listing.md), including 32 chars of name plust description. With 2 NXT additional fee for each 32 chars.
    -   1 NXT for [DGS Delivery](the-burst-api-dgs-delivery.md), including 32 bytes of encrypted goods data (AES initialization bytes and nonce excluded). With 2 NXT additional fee for each 32 bytes.
    -   2 NXT for transactions that make use of referencedTransactionFullHash property when creating a new transaction.
    -   1 NXT otherwise, where 1 NXT = 100000000 NQT
-   *deadline* is the deadline (in minutes) for the transaction to be confirmed, 32767 minutes maximum
-   *referencedTransactionFullHash* creates a chained transaction, meaning that the current transaction cannot be confirmed unless the referenced transaction is also confirmed (optional)
-   *broadcast* is set to *false* to prevent broadcasting the transaction to the network (optional)

**Note:** An optional arbitrary message can be appended to any transaction by adding message-related parameters as in [Send Message](the-burst-api-send-message.md).

**Note:** Any phasing-safe transaction (refer to [Create Transaction](the-burst-api-create-transaction.md)) can have its execution deferred either conditionally or unconditionally (refer to [Create Phasing Poll](the-burst-api-create-phasing-poll.md)).

### Create Transaction Response

The following JSON response fields are common to all API calls that create transactions:

-   *signatureHash* (S) is a SHA-256 hash of the transaction signature
-   *unsignedTransactionBytes* (S) are the unsigned transaction bytes
-   *transactionJSON* (O) is a transaction object (refer to [Get Transaction](the-burst-api-get-transaction.md) for details)
-   *broadcasted* (B) is *true* if the transaction was broadcast, *false* otherwise
-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *transactionBytes* (S) are the signed transaction bytes
-   *fullHash* (S) is the full hash of the signed transaction
-   *transaction* (S) is the ID of the newly created transaction

Account Operations
------------------

### Delete Account Property

Deletes an account property. POST only.

-   *property* is the name of the property
-   *recipient* is the account where a property should be removed (optional)
-   *setter* is the account who did set the property (optional)

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Delete Account Property](the-burst-api-examples-delete-account-property.md) example.

### Get Account

Get account information given an account ID.

-   *account* is the account ID
-   *includeLessors* is *true* to include *lessors*, *lessorsRS* and *lessorsInfo* (optional)
-   *includeAssets* is *true* to include *assetBalances* and *unconfirmedAssetBalances* (optional)
-   *includeCurrencies* is *true* to include *accountCurrencies* (optional)
-   *includeEffectiveBalance* is *true* to include *effectiveBalanceNXT* and *guaranteedBalanceNQT* (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *unconfirmedBalanceNQT* (S) is *balanceNQT* less unconfirmed outgoing transactions, the balance displayed in the client
-   *effectiveBalanceNXT* (N) is the balance (in NXT) of the account available for forging: the unleased guaranteedBalance of this account plus the leased guaranteedBalance of all lessors to this account
-   *lessorsInfo* (A) is an array of lessor objects including the fields:
    -   *currentHeightTo* (S)
    -   *nextHeightFrom* (S)
    -   *effectiveBalanceNXT* (S)
    -   *nextLesseeRS* (S)
    -   *currentLesseeRS* (S)
    -   *currentHeightFrom* (S)
    -   *nextHeightTo* (S)
-   *lessors* (A) is an array of lessor account IDs
-   *currentLessee* (S) is the account number of the lessee, if applicable
-   *currentLeasingHeightTo* (N) is the block height when the lease completes, if applicable
-   *forgedBalanceNQT* (S) is the balance (in NQT) that the account has forged
-   *balanceNQT* (S) is the minimally confirmed basic balance (in NQT) of the account
-   *publicKey* (S) is the public key of the account
-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *assetBalances* (A) is an array of asset objects including the fields *balanceQNT* (S) and *asset* (S) ID
-   *guaranteedBalanceNQT* (S) is the balance (in NQT) of the account with at least 1440 confirmations
-   *unconfirmedAssetBalances* (A) is an array of asset objects including the fields *unconfirmedBalanceQNT* (S) and *asset* (S) ID
-   *currentLesseeRS* (S) is the Reed-Solomon address of the lessee account
-   *accountRS* (S) is the Reed-Solomon address of the account
-   *lessorsRS* (A) is an array of Reed-Solomon lessor account addresses
-   *accountCurrencies* (A) is an array of currency objects (refer to [Get Account Currencies](the-burst-api-get-account-currencies.md) for details)
-   *name* (S) is the name associated with the account, if applicable
-   *description* (S) is the description of the account, if applicable
-   *account* (S) is the account number
-   *currentLeasingHeightFrom* (N) is the block height when the lease starts, if applicable
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)

**Example:** Refer to [Get Account](the-burst-api-examples-get-account.md) example.

### Get Account Block Count

Get the number of blocks forged by an account.

-   *account* is an account ID
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *numberOfBlocks* (N) is the number of blocks forged by the account
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Block Count](the-burst-api-examples-get-account-block-count.md) example.

### Get Account Block Ids

Get the block IDs of all blocks forged (generated) by an account in reverse block height order.

**Response:**

### Get Account Blocks

Get all blocks forged (generated) by an account in reverse block height order.

**Response:**

### Get Account Id

Get an account ID given a secret passphrase or public key. POST only.

**Request:**

-   *requestType* is *getAccountId*
-   *secretPhrase* is the secret passphrase of the account (optional)
-   *publicKey* is the public key of the account (optional if *secretPhrase* provided)

**Response:**

-   *accountRS* (S) is the Reed-Solomon address of the account
-   *publicKey* (S) is the public key of the account
-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *account* (S) is the account number

**Example:** Refer to [Get Account Id](the-burst-api-examples-get-account-id.md) example.

### Get Account Ledger

Get multiple account ledger entries.

**Request:**

-   *requestType* is *getAccountLedger*
-   *account* is the account ID (optional)
-   *firstIndex* is a zero-based index to the first block to retrieve (optional)
-   *lastIndex* is a zero-based index to the last block to retrieve (optional)
-   *event* is the event ID (optional)
-   *eventType* is a string representing the event type (optional)
-   *holdingType* is a string representing the holding type (optional)
-   *holding* is the holding ID (optional)
-   *includeTransactions* is true to retrieve transaction details, otherwise only transaction IDs are retrieved (optional)
-   *includeHoldingInfo* is true to retrieve asset or currency info (optional) with each ledger entry. The default is false.
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *entries* (A) is an array of ledger objects including the fields:
    -   *change* (S) is the change in the balance for the holding identified by 'holdingType'
    -   *eventType* (S) is the event type causing the account change
    -   *ledgerId* (S) is the ledger entry ID
    -   *holding* (S) is the item identifier for an asset or currency balance
    -   *isTransactionEvent* (B) is true if the event is associated with a transaction and false if it is associated with a block.
    -   *balance* (S) is the balance for the holding identified by 'holdingType'
    -   *holdingType* (S) is the item being changed (account balance, asset balance or currency balance)
    -   *accountRS* (S) is the Reed-Solomon address of the account
    -   *block* (S) is the block ID that created the ledger entry
    -   *event* (S) is the block or transaction associated with the event
    -   *account* (S) is the account number
    -   *height* (N) is the the block height associated with the event
    -   *timestamp* (N) is the the block timestamp associated with the event
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Ledger](the-burst-api-examples-get-account-ledger.md) example.

### Get Account Ledger Entry

Get a specific account ledger entry.

**Request:**

-   *requestType* is *getAccountLedgerEntry*
-   *ledgerId* is the ledger ID
-   *includeTransactions* is true to retrieve transaction details, otherwise only transaction IDs are retrieved (optional)
-   *includeHoldingInfo* is *true* to retrieve asset or currency info (optional) with the ledger entry. The default is false.

**Response:**

-   *change* (S) is the change in the balance for the holding identified by 'holdingType'
-   *eventType* (S) is the event type causing the account change
-   *ledgerId* (S) is the ledger entry ID
-   *holding* (S) is the item identifier for an asset or currency balance
-   *isTransactionEvent* (B) is true if the event is associated with a transaction and false if it is associated with a block.
-   *balance* (S) is the balance for the holding identified by 'holdingType'
-   *holdingType* (S) is the item being changed (account balance, asset balance or currency balance)
-   *accountRS* (S) is the Reed-Solomon address of the account
-   *block* (S) is the block ID that created the ledger entry
-   *event* (S) is the block or transaction associated with the event
-   *account* (S) is the account number
-   *height* (N) is the the block height associated with the event
-   *timestamp* (N) is the the block timestamp associated with the event
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Ledger Entry](the-burst-api-examples-get-account-ledger-entry.md) example.

### Get Account Lessors

Get the lessors to an account.

**Request:**

-   *requestType* is *getAccountLessors*
-   *account* is the account ID
-   *height* is the height of the blockchain to determine the lessors (optional, default is last block)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** If table trimming is enabled (default), the *height* must be within 1440 blocks of the last block.

**Response:**

-   *accountRS* (S) is the Reed-Solomon address of the account
-   *account* (S) is the account number
-   *height* (N) is the height of the blockchain
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *lessors* (A) is an array of lessor objects including the fields:
    -   *lessorRS* (S)
    -   *lessor* (S)
    -   *guaranteedBalanceNQT* (S)

**Example:** Refer to [Get Account Lessors](the-burst-api-examples-get-account-lessors.md) example.

### Get Account Properties

Get the Account properties for a specific account or setter.

**Request:**

-   *requestType* is *getAccountProperties*
-   *recipient* is the account ID to which the property is attached to
-   *setter* is the account ID who set the property (optional if *recipient* provided)
-   *property* is the property key (optional)
-   *firstIndex* is a zero-based index to the first block to retrieve (optional)
-   *lastIndex* is a zero-based index to the last block to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *setterRS*: (S) is the Reed-Solomon address of the setter account (only if setter is defined in the request)
-   *recipientRS*: (S) is the Reed-Solomon address of the recipient account (only if recipient is defined in the request)
-   *recipient*: (S) is the account number of the recipient account (only if recipient is defined in the request)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *setter*: (S) is the account number of the setter account (only if setter is defined in the request)
-   *properties*: (A) is an array of properties including the fields:
    -   *setterRS*: (S) is the Reed-Solomon address of the setter account (only if setter is omitted in the request)
    -   *recipientRS*: (S) is the Reed-Solomon address of the recipient account (only if recipient is omitted in the request)
    -   *recipient*: (S) is the account number of the recipient account (only if recipient is omitted in the request)
    -   *property*: (S) property name
    -   *setter*: (S) is the account number of the setter account (only if setter is omitted in the request)
    -   *value*: (S) property value

**Example:** Refer to [Get Account Properties](the-burst-api-examples-get-account-properties.md) example.

### Get Account Public Key

Get the public key associated with an account ID.

**Request:**

-   *requestType* is *getAccountPublicKey*
-   *account* is the account ID
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *publicKey* (S) is the 32-byte public key associated with the account, returned as a hex string
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Public Key](the-burst-api-examples-get-account-public-key.md) example.

### Get Account Transaction Ids

Get the transaction IDs associated with an account in reverse block timestamp order. *This call only returns non-phased transactions as of [Version 1.5.7e](burst-software-change-log-version-1-5-7e.md) and is deprecated, to be removed in version 1.6. Use [Get Blockchain Transactions](the-burst-api-get-blockchain-transactions.md) instead.*

**Request:**

-   *requestType* is *getAccountTransactionIds*
-   *account* is the account ID
-   *timestamp* is the earliest block (in seconds since the genesis block) to retrieve (optional)
-   *type* is the type of transactions to retrieve (optional)
-   *subtype* is the subtype of transactions to retrieve (optional)
-   *firstIndex* is a zero-based index to the first transaction ID to retrieve (optional)
-   *lastIndex* is a zero-based index to the last transaction ID to retrieve (optional)
-   *numberOfConfirmations* is the required number of confirmations per transaction (optional)
-   *withMessage* is *true* to retrieve only only transactions having a message attachment, either non-encrypted or decryptable by the account (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** Refer to [Get Constants](the-burst-api-get-constants.md) for definitions of types and subtypes

**Response:**

-   *transactionIds* (A) is an array of transaction IDs
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Transaction Ids](the-burst-api-examples-get-account-transaction-ids.md) example.

### Get Account Transactions

Get the transactions associated with an account in reverse block timestamp order. *This call only returns non-phased transactions as of [Version 1.5.7e](burst-software-change-log-version-1-5-7e.md) and is depricated, to be removed in version 1.6. Use [Get Blockchain Transactions](the-burst-api-get-blockchain-transactions.md) instead.*

**Request:**

-   *requestType* is *getAccountTransactions*
-   *account* is the account ID
-   *timestamp* is the earliest block (in seconds since the genesis block) to retrieve (optional)
-   *type* is the type of transactions to retrieve (optional)
-   *subtype* is the subtype of transactions to retrieve (optional)
-   *firstIndex* is a zero-based index to the first transaction to retrieve (optional)
-   *lastIndex* is a zero-based index to the last transaction to retrieve (optional)
-   *numberOfConfirmations* is the required number of confirmations per transaction (optional)
-   *withMessage* is *true* to retrieve only only transactions having a message attachment, either non-encrypted or decryptable by the account (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** Refer to [Get Constants](the-burst-api-get-constants.md) for definitions of types and subtypes

**Response:**

-   *transactions* (A) is an array of transactions (refer to [Get Transaction](the-burst-api-get-transaction.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Transactions](the-burst-api-examples-get-account-transactions.md) example.

### Get Balance

Get the balance of an account.

**Request:**

-   *requestType* is *getBalance*
-   *account* is an account ID
-   *includeEffectiveBalance* is *true* to include *effectiveBalanceNXT* and *guaranteedBalanceNQT* (optional)
-   *height* is the height to retrieve account balance for, if still available (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *unconfirmedBalanceNQT* (S) is *balanceNQT* less unconfirmed outgoing transactions, the balance displayed in the client
-   *guaranteedBalanceNQT* (S) is the balance (in NQT) of the account with at least 1440 confirmations
-   *effectiveBalanceNXT* (N) is the balance (in NXT) of the account available for forging: the unleased guaranteedBalance of this account plus the leased guaranteedBalance of all lessors to this account
-   *forgedBalanceNQT* (S) is the balance (in NQT) that the account has forged
-   *balanceNQT* (S) is the minimally confirmed basic balance (in NQT) of the account
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Balance](the-burst-api-examples-get-balance.md) example.

### Get Blockchain Transactions

Get the transactions associated with an account in reverse block timestamp order.

**Request:**

-   *requestType* is *getBlockchainTransactions*
-   *account* is the account ID
-   *timestamp* is the earliest block (in seconds since the genesis block) to retrieve (optional)
-   *type* is the type of transactions to retrieve (optional)
-   *subtype* is the subtype of transactions to retrieve (optional)
-   *firstIndex* is a zero-based index to the first transaction to retrieve (optional)
-   *lastIndex* is a zero-based index to the last transaction to retrieve (optional)
-   *numberOfConfirmations* is the required number of confirmations per transaction (optional)
-   *withMessage* is *true* to retrieve only only transactions having a message attachment, either non-encrypted or decryptable by the account (optional)
-   *phasedOnly* is *true* to retrieve only phased transactions (optional unless *nonPhasedOnly* provided)
-   *nonPhasedOnly* is *true* to retrieve only nonphased transactions (optional unless *phasedOnly* provided)
-   *includeExpiredPrunable* is *true* to retrieve pruned data if available (optional)
-   *includePhasingResult* is *true* to retrieve execution status of each phased transaction (optional)
-   *executedOnly* is *true* to exclude phased transactions that are not yet executed (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** Refer to [Get Constants](the-burst-api-get-constants.md) for definitions of types and subtypes

**Response:**

-   *transactions* (A) is an array of transactions (refer to [Get Transaction](the-burst-api-get-transaction.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Blockchain Transactions](the-burst-api-examples-get-blockchain-transactions.md) example.

### Get Guaranteed Balance

Get the balance of an account that is confirmed at least a specified number of times.

**Request:**

-   *requestType* is *getGuaranteedBalance*
-   *account* is an account ID
-   *numberOfConfirmations* is the minimum number of confirmations for a transaction to be included in the guaranteed balance (optional, if omitted or zero then minimally confirmed transactions are included)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *guaranteedBalanceNQT* (S) is the balance (in NQT) of the account with at least *numberOfConfirmations* confirmations
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Guaranteed Balance](the-burst-api-examples-get-guaranteed-balance.md) example.

### Get Unconfirmed Transaction Ids

Get a list of unconfirmed transaction IDs associated with an account.

**Request:**

-   *requestType* is *getUnconfirmedTransactionIds*
-   *account* is one account ID (optional)
-   *account* is one account ID (optional)

⋮

-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)
-   *firstIndex* is a zero-based index to the first transaction ID to retrieve (optional)
-   *lastIndex* is a zero-based index to the last transaction ID to retrieve (optional)

**Response:**

-   *unconfirmedTransactionIds* (A) is an array of unconfirmed transaction IDs
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Unconfirmed Transaction Ids](the-burst-api-examples-get-unconfirmed-transaction-ids.md) example.

### Get Unconfirmed Transactions

Get a list of unconfirmed transactions associated with an account.

**Request:**

-   *requestType* is *getUnconfirmedTransactions*
-   *account* is one account ID (optional)
-   *account* is one account ID (optional)

⋮

-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)
-   *firstIndex* is a zero-based index to the first unconfirmed transaction to retrieve (optional)
-   *lastIndex* is a zero-based index to the last unconfirmed transaction to retrieve (optional)

**Response:**

-   *unconfirmedTransactions* (A) is an array of unconfirmed transactions (refer to [Get Transaction](the-burst-api-get-transaction.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Unconfirmed Transactions](the-burst-api-examples-get-unconfirmed-transactions.md) example.

### Search Accounts

Get accounts having a name or description that match a given query in reverse relevance order.

**Request:**

-   *requestType* is *searchAccounts*
-   *query* is a full text query on the account fields *name* (S) and *description* (S) in the [standard Lucene syntax](http://lucene.apache.org/core/2_9_4/queryparsersyntax.html#Overview)
-   *firstIndex* is a zero-based index to the first account to retrieve (optional)
-   *lastIndex* is a zero-based index to the last account to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *accounts* (A) is an array of account objects with the following fields:
    -   *account* (S) is the account number
    -   *accountRS* (S) is the Reed-Solomon address of the account
    -   *name* (S) is the name of the account
    -   *description* (S) is the description of the account (if applicable)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Search Accounts](the-burst-api-examples-search-accounts.md) example.

### Send Money

Send NXT to an account. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *sendMoney*
-   *amountNQT* is the amount (in NQT) in the transaction
-   *recipient* is the account ID of the recipient
-   *recipientPublicKey* is the public key of the receiving account (optional, enhances security of a new account)

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Send Money](the-burst-api-examples-send-money.md) example.

#### Send NXT

Refer to [Send Money](the-burst-api-send-money.md).

### Set Account Info

Set account information. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *setAccountInfo*
-   *name* is a name to associate with the account (optional)
-   *description* is a description to associate with the account (optional)
-   *messagePatternRegex* is a regular expression pattern following the java.util.regex.Pattern specification; incoming transactions are only accepted if they contain a plain text message which matches this pattern (disabled indefinitely due to a security issue)
-   *messagePatternFlags* is a bitmask of java.util.regex.Pattern flags, such as 2 (Pattern.CASE\_INSENSITIVE)

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Set Account Info](the-burst-api-examples-set-account-info.md) example.

### Set Account Property

Set account property. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *setAccountProperty*
-   *recipient* is the account ID of the recipient.
-   *property* is the property name.
-   *value* is the property value.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Set Account Property](the-burst-api-examples-set-account-property.md) example.

### Start Funding Monitor

Starts a funding monitor that will transfer NXT, ASSET or CURRENCY from the funding account to a recipient account when the amount held by the recipient account drops below the threshold. The transfer will not be done until the current block height is greater than equal to the block height of the last transfer plus the interval. The funding account is identified by the secret phrase.

The recipient accounts are identified by the specified account property (see [Set Account Property](the-burst-api-set-account-property.md)). Each account that has this property set by the funding account will be monitored for changes. The property value can be omitted or it can consist of a JSON string containing one or more values in the format: {“amount”:long,“threshold”:long,“interval”:integer} . POST only.

**Request:**

-   *requestType* is *startFundingMonitor*
-   *property* is the name of the account property
-   *amount* is the amount to fund the recipient account with (in NQT or QNT)
-   *threshold* is the threshold
-   *interval* is the the number of blocks to wait after before funding the recipient
-   *secretPhrase* is the secret phrase of the funding account
-   *holdingType* is a string representing the holding type (optional)
-   *holding* is the holding ID (optional)

**Response:**

-   *started* (B) is *true* if the monitor has been started
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Start Funding Monitor](the-burst-api-examples-start-funding-monitor.md) example.

### Stop Funding Monitor

Stop a funding monitor. When the secret phrase is specified, a single monitor will be stopped. The monitor is identified by the secret phrase, holding and account property. The administrator password is not required and will be ignored.

When the administrator password is specified, a single monitor can be stopped by specifying the funding account, holding and account property. If no account is specified, all monitors will be stopped.

The holding type and account property name must be specified when the secret phrase or account is specified. Holding type codes are listed in getConstants. In addition, the holding identifier must be specified when the holding type is ASSET or CURRENCY. POST only.

**Request:**

-   *requestType* is *stopFundingMonitor*
-   *secretPhrase* is the secret phrase of the funding account, used to stop a single monitor. (optional)
-   *adminPassword* is the admin password, used to stop a single monitor or all monitors (optional if *secretPhrase* is provided)
-   *property* is the name of the account property (optional)
-   *holdingType* is a string representing the holding type (optional)
-   *holding* is the holding ID (optional)
-   *account* is the account ID (optional)

**Response:**

-   *stopped* (N) is the number of the monitors that have been stopped
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Stop Funding Monitor](the-burst-api-examples-stop-funding-monitor.md) example.

Account Control Operations
--------------------------

### Get All Phasing Only Controls

Retrieve all accounts subject to phasing control with their respective restrictions.

**Request:**

-   *requestType* is *getAllPhasingOnlyControls*
-   *firstIndex* is a zero-based index to the first block ID to retrieve (optional)
-   *lastIndex* is a zero-based index to the last block ID to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *phasingOnlyControls* (A) is an array with phasing only controls objects (Refer to [Get Phasing Only Control](the-burst-api-get-phasing-only-control.md) for details)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get All Phasing Only Controls](the-burst-api-examples-get-all-phasing-only-controls.md) example.

### Get Phasing Only Control

Retrieve phasing control with their respective restrictions for a specific account.

**Request:**

-   *requestType* is *getPhasingOnlyControl*
-   *account* is the account ID
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *account* (S) is the account number
-   *accountRS* (S) is the Reed-Solomon address of the account
-   *quorum* (S) is the minimum number of votes needed to approve the transaction
-   *whitelist* (A) is an array with the whitelisted accounts including the fields:
    -   *whitelisted* (S) is the account number
    -   *whitelistedRS* (S) is the Reed-Solomon address of the account
-   *maxFees* (S) is the maximum fees the account can spend per block
-   *minDuration* (N) is the minimum duration of the phasing period
-   *maxDuration* (N) is the maximum duration of the phasing period
-   *votingModel* (N) is an integer code for the method of approval
-   *minBalance* (S) is the minimum balance (in NQT or QNT) required for voting
-   *minBalanceModel* (N) is the minimum balance model
-   *holding* (S) is the asset or currency ID (only included if holding != 0)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Phasing Only Control](the-burst-api-examples-get-phasing-only-control.md) example.

### Set Phasing Only Control

Sets (or removes) phasing control for a specific account. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *setPhasingOnlyControl*
-   *controlVotingModel* is the voting model or -1 to remove phasing control
-   *controlQuorum* is the expected quorum (optional)
-   *controlMinBalance* is the expected minimum balance (optional)
-   *controlMinBalanceModel* is the expected minimum balance model (optional)
-   *controlHolding* is the holding ID (optional)
-   *controlWhitelisted* is the whitelisted accounts (optional, multiple values)
-   *controlWhitelisted* is the whitelisted accounts (optional, multiple values)

⋮

-   *controlMaxFees* is the maximum allowed accumulated total fees for not yet finished phased transactions (optional)
-   *controlMinDuration* is the minimum duration in block height (optional)
-   *controlMaxDuration* is the maximum phasing duration in block height (optional)

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Set Phasing Only Control](the-burst-api-examples-set-phasing-only-control.md) example.

Alias Operations
----------------

### Buy / Sell Alias

Buy or sell an alias. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is either *buyAlias* or *sellAlias*
-   *alias* is the ID of the alias (optional)
-   *aliasName* is the alias name (optional if *alias* provided)
-   *priceNQT* is the asking price (in NQT) of the alias (*sellAlias* only)
-   *amountNQT* is the amount (in NQT) offered for an alias for sale (*buyAlias* only)
-   *recipient* is the account ID of the recipient (only required for *sellAlias* and only if there is a designated recipient)
-   *recipientPublicKey* is the public key of the recipient account (only applicable if *recipient* provided; optional, enhances security of a new account)

**Note**: An alias can be transferred rather than sold by setting *priceNQT* to zero. A pending sale can be canceled by selling again to self for a price of zero.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Buy / Sell Alias](the-burst-api-examples-buy---sell-alias.md) example.

#### Buy Alias

Refer to [Buy / Sell Alias](the-burst-api-buy--2f-sell-alias.md).

#### Sell Alias

Refer to [Buy / Sell Alias](the-burst-api-buy--2f-sell-alias.md).

### Set Alias

Create and/or assign an alias. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *setAlias*
-   *aliasName* is the alias name
-   *aliasURI* is the alias URI (e.g. <http://www.google.com/>)

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md). The transaction ID is also the alias ID.

**Example:** Refer to [Set Alias](the-burst-api-examples-set-alias.md) example.

### Delete Alias

Delete an alias given an alias ID or name. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *deleteAlias*
-   *alias* is the alias ID (optional)
-   *aliasName* is the alias name to be deleted (optional if *alias* provided)

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Delete Alias](the-burst-api-examples-delete-alias.md) example.

### Get Alias

Get information about a given alias

**Request:**

-   *requestType* is *getAlias*
-   *alias* is the alias ID (optional)
-   *aliasName* is the name of the alias (optional if *alias* provided)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *timestamp* (N) is the time (in seconds since the genesis block) when the alias was created or last transferred
-   *aliasName* (S) is the name of the alias
-   *account* (S) is the number of the account that owns the alias
-   *accountRS* (S) is the Reed-Solomon address of the account that owns the alias
-   *aliasURI* (S) is what the alias points to, in URI format
-   *alias* (S) is the alias ID
-   *priceNQT* (S) is the asking price (in NQT) of the alias if it is for sale
-   *buyer* (S) is the account number of the buyer if the alias is for sale and a buyer is specified
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Alias](the-burst-api-examples-get-alias.md) example.

### Get Alias Count

Get the number of aliases owned by an account given the account ID.

**Request:**

-   *requestType* is *getAliasCount*
-   *account* is the account ID
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *numberOfAliases* (N) is the number of aliases owned by the account
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Alias Count](the-burst-api-examples-get-alias-count.md) example.

### Get Aliases

Get information on aliases owned by a given account in alias name order.

**Request:**

-   *requestType* is *getAliases*
-   *account* is the ID of the account that owns the aliases
-   *timestamp* is the earliest creation time (in seconds since the genesis block) of the aliases (optional)
-   *firstIndex* is a zero-based index to the first alias to retrieve (optional)
-   *lastIndex* is a zero-based index to the last alias to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *aliases* (A) is an array of alias objects (refer to [Get Alias](the-burst-api-get-alias.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Aliases](the-burst-api-examples-get-aliases.md) example.

### Get Aliases Like

Get all aliases starting with a given prefix in alias name order.

**Request:**

-   *requestType* is *getAliasesLike*
-   *aliasPrefix* is the prefix (at least 2 characters long) of the *aliasName*
-   *firstIndex* is a zero-based index to the first alias to retrieve (optional)
-   *lastIndex* is a zero-based index to the last alias to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *aliases* (A) is an array of alias objects (refer to [Get Alias](the-burst-api-get-alias.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Aliases Like](the-burst-api-examples-get-aliases-like.md) example.

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

**Example:** Refer to [Decrypt From](the-burst-api-examples-decrypt-from.md) example.

### Download Prunable Message

Downloadins a prunable message attachments directly as binary data. An optional secretPhrase parameter is supported, to allow decryption and downloading of the encrypted part of the message instead of the plain text part.

**Request:**

-   *requestType* is *downloadPrunableMessage*
-   *transaction* is the transaction ID
-   *secretPhrase* is the secret passphrase used to decrypt the encrypted part of the message (optional)
-   *sharedKey* is the shared key used to decrypt the message (optional) (see [Get Shared Key](the-burst-api-get-shared-key.md))
-   *retrieve* is *true* to retrieve the message from achival node if needed (optional)
-   *requireBlock* is theblock ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:** The prunable data as a binary file.

**Example:** Refer to [Download Prunable Message](the-burst-api-examples-download-prunable-message.md) example.

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

**Example:** Refer to [Encrypt To](the-burst-api-examples-encrypt-to.md) example.

### Get All Prunable Messages

Get all available prunable messages in reverse block timestamp order.

**Request:**

-   *requestType* is *getAllPrunableMessages*
-   *timestamp* is the earliest message (in seconds since the genesis block) to retrieve (optional)
-   *firstIndex* is a zero-based index to the first prunable message to retrieve (optional)
-   *lastIndex* is a zero-based index to the last prunable message to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *prunableMessages* (A) is an array of prunable messages (refer to [Get Prunable Message](the-burst-api-get-prunable-message.md))
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millsec)

**Example:** Refer to [Get All Prunable Messages](the-burst-api-examples-get-all-prunable-messages.md) example.

### Get Prunable Message

Get the prunable message given a transaction ID, optionally decrypting it if encrypted and if a secretPhrase is provided, if it is still available.

**Request:**

-   *requestType* is *getPrunableMessage*
-   *transaction* is the transaction ID
-   *secretPhrase* is the secret phrase needed for decryption (optional)
-   *sharedKey* is the shared key used to decrypt the message (optional) (see [Get Shared Key](the-burst-api-get-shared-key.md))
-   *retrieve* is *true* to retrieve pruned data from other nodes if not available (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *sender* (S) is the sender account number
-   *senderRS* (S) is the Reed-Solomon address of the sender account
-   *recipient* (S) is the recipient account number
-   *recipientRS* (S) is the Reed-Solomon address of the recipient account
-   *message* (S) is the plain message text
-   *messageIsText* (B) is *true* if the plain message is text, *false* if it is a hex string
-   *encryptedMessage* (O) is the encrypted message object, containing *data* (S) and *nonce* (S) fields
-   *encryptedMessageIsText* (B) is *true* if the encrypted message is text, *false* if it is a hex string
-   *decryptedMessage* (S) is the decrypted and decompressed (if necessary) encrypted message, if applicable and if *secretPhrase* is provided
-   *isCompressed* (B) is *true* if the encrypted message was compressed before encryption, if applicable
-   *transaction* (S) is the transaction ID
-   *transactionTimestamp* (N) is the transaction timestamp (in seconds since the genesis block)
-   *blockTimestamp* (N) is the block timestamp (in seconds since the genesis block)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millsec)

**Example:** Refer to [Get Prunable Message](the-burst-api-examples-get-prunable-message.md) example.

### Get Prunable Messages

Get all still-available prunable messages given an account id, optionally limiting to only those with another account as recipient or sender, in reverse chronological order.

**Request:**

-   *requestType* is *getPrunableMessages*
-   *account* is the account ID
-   *otherAccount* is another account ID, either sender or recipient, to limit the response
-   *secretPhrase* is the secret phrase needed for decryption (optional)
-   *timestamp* is the earliest prunable message (in seconds since the genesis block) to retrieve (optional)
-   *firstIndex* is a zero-based index to the first prunable message to retrieve (optional)
-   *lastIndex* is a zero-based index to the last prunable message to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *prunableMessages* (A) is an array of prunable message objects (refer to [Get Prunable Message](the-burst-api-get-prunable-message.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millsec)

**Example:** Refer to [Get Prunable Messages](the-burst-api-examples-get-prunable-messages.md) example.

### Get Shared Key

Get the one-time shared key used for encryption of messages.

**Request:**

-   *requestType* is *getSharedKey*
-   *account* is the recipient account ID
-   *secretPhrase* is the secret phrase of the sender
-   *nonce* is the 32-byte pseudorandom nonce

**Response:**

-   *sharedKey* (S) is shared key as a hexadecimal string
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Shared Key](the-burst-api-examples-get-shared-key.md) example.

### Read Message

Get a message given a transaction ID.

**Request:**

-   *requestType* is *readMessage*
-   *transaction* is the transaction ID of the message
-   *secretPhrase* is the secret passphrase of the account that received the message (optional)
-   *sharedKey* is the shared key used to decrypt the message (optional) (see [Get Shared Key](the-burst-api-get-shared-key.md))
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

**Example:** Refer to [Read Message](the-burst-api-examples-read-message.md) example.

### Send Message

Create an Arbitrary Message transaction. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

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

**Note:** The *encryptedMessageData-encryptedMessageNonce* pair or the *encryptToSelfMessageData-encryptToSelfMessageNonce* pair can be the output of [Encrypt To](the-burst-api-encrypt-to.md)

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Send Message](the-burst-api-examples-send-message.md) example.

### Verify Prunable Message

Verify that a prunable message obtained from any source, when hashed, matches the hash of the original prunable message.

**Request:** Refer to [Send Message](the-burst-api-send-message.md) for more details about the following request parameters.

-   *requestType* is *verifyPrunableMessage*
-   *message* is the plain message, if applicable (optional)
-   *messageIsText* is *false* if the provided plain *message* is a hex string (optional)
-   *encryptedMessageData* is the data part of the encrypted data-nonce pair (optional if *message* provided)
-   *encryptedMessageNonce* is the nonce part of the encrypted data-nonce pair (required if *encryptedMessageData* provided)
-   *messageToEncryptIsText* is *false* if the encrypted message was a hex string before encryption (optional)
-   *compressMessageToEncrypt* is *false* if the encrypted message was not compressed before encryption (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** The hash is computed from the message itself plus its associated flag(s) *isText* and *isCompressed* (encrypted only); therefore the flag(s) must be provided for verification if different from the default(s). The original *encryptedMessageData-encryptedMessageNonce* pair used to compute the original hash must be provided again to recompute the hash for verification of a prunable encrypted message.

**Response:**

-   *version.PrunablePlainMessage* or *version.PrunableEncryptedMessage* (N) is *1*, the version number
-   *verify* (B) is *true* if the original hash matches the hash computed from the provided values
-   *message* (S) or *encryptedMessage* (O) is the prunable plain message or the prunable encrypted message object containing the fields:
    -   *data* (S)
    -   *nonce* (B)
    -   *isText* (B)
    -   *isCompressed* (B)
-   *messageIsText* (B) is *true* if the plain message is text, *false* if it is a hex string, if applicable
-   *messageHash* or *encryptedMessageHash* (S) is the hash
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millsec)

**Example:** Refer to [Verify Prunable Message](the-burst-api-examples-verify-prunable-message.md) example.

Asset Exchange Operations
-------------------------

### Cancel Order

Cancel an existing asset order. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is either *cancelBidOrder* or *cancelAskOrder*
-   *order* is the order ID of the order being canceled

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Cancel Order](the-burst-api-examples-cancel-order.md) example.

#### Cancel Ask Order

Refer to [Cancel Order](the-burst-api-cancel-order.md).

#### Cancel Bid Order

Refer to [Cancel Order](the-burst-api-cancel-order.md).

### Delete Asset Shares

Permanently deletes a specified quantity of owned asset shares.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *deleteAssetShares*
-   *asset* is the asset ID
-   *quantityQNT* is the quantity (in QNT) of the asset to be deleted

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Delete Asset Shares](the-burst-api-examples-delete-asset-shares.md) example.

### Dividend Payment

Pay dividend to all shareholders of an asset. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *dividendPayment*
-   *asset* is the asset ID
-   *height* is the blockchain height at which asset holders shares will be counted (must be less than 1440 blocks in the past)
-   *amountNQTPerQNT* is dividend amount (in NQT per QNT of the asset)

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Dividend Payment](the-burst-api-examples-dividend-payment.md) example.

### Get Account Asset Count

Get the number of assets owned by an account given the account ID.

**Request:**

-   *requestType* is *getAccountAssetCount*
-   *account* is the account ID
-   *height* is the height of the blockchain to determine the asset count (optional, default is last block)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** If table trimming is enabled (default), the height must be within 1440 blocks of the last block.

**Response:**

-   *numberOfAssets* (N) is the number of assets owned by the account
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Asset Count](the-burst-api-examples-get-account-asset-count.md) example.

### Get Account Assets

Get the assets owned by a given account in reverse quantity order.

**Request:**

-   *requestType* is *getAccountAssets*
-   *account* is the account ID
-   *asset* is an asset ID filter (optional)
-   *height* is the blockchain height at which to retrieve balances (optional, default is the last block in the blockchain)
-   *includeAssetInfo* is *true* if asset information is to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** If table trimming is enabled (default), the height must be within 1440 blocks of the last block.

**Response:**

-   *accountAssets* (A) is an array of asset objects (unless the *asset* parameter is specified) with the following fields for each asset:
    -   *quantityQNT* (S) is the quantity (in QNT) of the asset
    -   *unconfirmedQuantityQNT* (S) is the unconfirmed quantity (in QNT) of the asset
    -   *decimals* (N) is the number of decimal places used by the asset (if *includeAssetInfo* is *true*)
    -   *name* (S) is the asset name (if *includeAssetInfo* is *true*)
    -   *asset* (S) is the asset ID
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Assets](the-burst-api-examples-get-account-assets.md) example.

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

**Example:** Refer to [Get Account Current Order Ids](the-burst-api-examples-get-account-current-order-ids.md) example.

#### Get Account Current Ask Order Ids

Refer to [Get Account Current Order Ids](the-burst-api-get-account-current-order-ids.md).

#### Get Account Current Bid Order Ids

Refer to [Get Account Current Order Ids](the-burst-api-get-account-current-order-ids.md).

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

-   *bidOrders* or *askOrders* (A) is an array of order objects (refer to [Get Order](the-burst-api-get-order.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Current Orders](the-burst-api-examples-get-account-current-orders.md) example.

#### Get Account Current Ask Orders

Refer to [Get Account Current Orders](the-burst-api-get-account-current-orders.md).

#### Get Account Current Bid Orders

Refer to [Get Account Current Orders](the-burst-api-get-account-current-orders.md).

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

-   *assets* (A) is an array of asset objects (refer to [Get Asset](the-burst-api-get-asset.md))
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get All Assets](the-burst-api-examples-get-all-assets.md) example.

### Get All Open Orders

Get all open bid/ask orders in reverse block height order.

**Request:**

-   *requestType* is either *getAllOpenBidOrders* or *getAllOpenAskOrders*
-   *firstIndex* is a zero-based index to the first order to retrieve (optional)
-   *lastIndex* is a zero-based index to the last order to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *openOrders* (A) is an array of order objects (refer to [Get Order](the-burst-api-get-order.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get All Open Orders](the-burst-api-examples-get-all-open-orders.md) example.

#### Get All Open Ask Orders

Refer to [Get All Open Orders](the-burst-api-get-all-open-orders.md).

#### Get All Open Bid Orders

Refer to [Get All Open Orders](the-burst-api-get-all-open-orders.md).

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

-   *trades* (A) is an array of trade objects (refer to [Get Trades](the-burst-api-get-trades.md))
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get All Trades](the-burst-api-examples-get-all-trades.md) example.

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

**Example:** Refer to [Get Asset](the-burst-api-examples-get-asset.md) example.

### Get Asset Account Count

Get the number of accounts that own an asset given the asset ID.

**Request:**

-   *requestType* is *getAssetAccountCount*
-   *asset* is the asset ID
-   *height* is the height of the blockchain to determine the account count (optional, default is last block)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** If table trimming is enabled (default), the height must be within 1440 blocks of the last block.

**Response:**

-   *numberOfAccounts* (N) is the number of accounts that own the asset
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Asset Account Count](the-burst-api-examples-get-asset-account-count.md) example.

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

**Example:** Refer to [Get Asset Accounts](the-burst-api-examples-get-asset-accounts.md) example.

### Get Asset Deletes

Get asset deletions for a specific asset or account.

**Request:**

-   *requestType* is *getAssetDeletes*
-   *asset* is the asset ID (optional if account is provided)
-   *account* is the account ID (optional if asset is provided)
-   *firstIndex* is a zero-based index to the first phased transaction to retrieve (optional)
-   *lastIndex* is a zero-based index to the last phased transaction to retrieve (optional)
-   *timestamp* is the earliest deletion (in seconds since the genesis block) to retrieve (optional)
-   *includeAssetInfo* is *true* if asset information is to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *deletes* (A) is an array of asset delete objects with following properties:
    -   *quantityQNT* (S) is the number of shares that was deleted
    -   *assetDelete* (S) is the transaction ID
    -   *account* (S) is the account ID
    -   *accountRS* (S) is the account Reed Solomon address
    -   *asset* (S) is the asset ID
    -   *height* (N) is the block height of the delete
    -   *timestamp* (N) is the block timestamp of the delete
    -   *decimals* (N) is the number of decimal places used by the asset (if *includeAssetInfo* is *true*)
    -   *name* (S) is the asset name (if *includeAssetInfo* is *true*)

**Example:** Refer to [Get Asset Deletes](the-burst-api-examples-get-asset-deletes.md) example.

### Get Asset Dividends

Get the dividend payment history for a specific asset.

**Request:**

-   *requestType* is *getAssetDividends*
-   *asset* is the asset ID
-   *firstIndex* is a zero-based index to the first dividend payment to retrieve (optional)
-   *lastIndex* is a zero-based index to the last dividend payment to retrieve (optional)
-   *timestamp* is the earliest dividend payment (in seconds since the genesis block) to retrieve (optional)
-   *adminPassword* is a string with the admin password (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *dividends* (A) is an array of dividend transactions with the following properties:
    -   *assetDividend* (S) is the dividend payment transaction ID
    -   *numberOfAccounts* (N) is the number of accounts that received a dividend
    -   *amountNQTPerQNT* (S) is the amount of NXT (in NQT) paid per quantity (in QNT) of the asset
    -   *totalDividend* (S) is the total amount of NXT (in NQT) sent in the dividend payment
    -   *dividendHeight* (N) is the block height of the dividend calculation
    -   *asset* (S) is the asset ID
    -   *height* (N) is the block height of the dividend payment
    -   *timestamp* (N) is the block timestamp of the dividend payment

**Example:** Refer to [Get Asset Dividends](the-burst-api-examples-get-asset-dividends.md) example.

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

**Example:** Refer to [Get Asset Ids](the-burst-api-examples-get-asset-ids.md) example.

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

**Example:** Refer to [Get Asset Transfers](the-burst-api-examples-get-asset-transfers.md) example.

#### Get Expected Asset Transfers

Refer to [Get Asset Transfers](the-burst-api-get-asset-transfers.md).

### Get Assets

Get asset information given multiple asset IDs

**Request:**

-   *requestType* is *getAssets*
-   *assets* is one the multiple asset IDs
-   *assets* is one the multiple asset IDs

⋮

-   *includeCounts* is *true* if the fields beginning with *numberOf...* are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *assets* (A) is an array of asset objects (refer to [Get Asset](the-burst-api-get-asset.md))
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Assets](the-burst-api-examples-get-assets.md) example.

### Get Assets By Issuer

Get asset information given multiple creation account IDs in reverse block height of creation order.

**Request:**

-   *requestType* is *getAssetsByIssuer*
-   *account* is one of the multiple account IDs
-   *account* is one of the multiple account IDs

⋮

-   *firstIndex* is a zero-based index to the first asset to retrieve (optional)
-   *lastIndex* is a zero-based index to the last asset to retrieve (optional)
-   *includeCounts* is *true* if the fields beginning with *numberOf...* are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *assets* (A) is an array of asset objects (refer to [Get Asset](the-burst-api-get-asset.md))
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Assets By Issuer](the-burst-api-examples-get-assets-by-issuer.md) example.

### Get Expected Asset Deletes

Gets asset deletes which are expected to be executed in the next block.

**Request:**

-   *requestType* is either *getExpectedAssetDeletes*
-   *asset* is the asset ID (optional)
-   *account* is the account ID (optional)
-   *firstIndex* is a zero-based index to the first phased transaction to retrieve (optional)
-   *lastIndex* is a zero-based index to the last phased transaction to retrieve (optional)
-   *timestamp* is the earliest deletion (in seconds since the genesis block) to retrieve (optional)
-   *includeAssetInfo* is *true* if asset information is to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *deletes* (A) is an array of expected asset delete objects with following properties:
    -   *assetDelete* (S) is the transaction ID
    -   *asset* (S) is the asset ID
    -   *account* (S) is the account ID
    -   *accountRS* (S) is the account Reed Solomon address
    -   *quantityQNT* (S) is the number of shares that will be deleted
    -   *decimals* (N) is the number of decimal places used by the asset (if *includeAssetInfo* is *true*)
    -   *name* (S) is the asset name (if *includeAssetInfo* is *true*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Expected Asset Deletes](the-burst-api-examples-get-expected-asset-deletes.md) example.

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

**Example:** Refer to [Get Order](the-burst-api-examples-get-order.md) example.

#### Get Ask Order

Refer to [Get Order](the-burst-api-get-order.md).

#### Get Bid Order

Refer to [Get Order](the-burst-api-get-order.md).

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

**Example:** Refer to [Get Order Ids](the-burst-api-examples-get-order-ids.md) example.

#### Get Ask Order Ids

Refer to [Get Order Ids](the-burst-api-get-order-ids.md).

#### Get Bid Order Ids

Refer to [Get Order Ids](the-burst-api-get-order-ids.md).

### Get Orders

Get bid/ask orders given an asset ID, in order of decreasing bid price or increasing ask price (if *sortByPrice* is *true* for expected orders, otherwise in the expected order of execution).

**Request:**

-   *requestType* is one of *getBidOrders*, *getAskOrders*, *getExpectedBidOrders* or *getExpectedAskOrders*, where expected orders are from the unconfirmed transactions pool or are phased transactions scheduled to finish in the next block
-   *asset* is the asset ID
-   *sortByPrice* is *true* to sort by price (optional, applies only to expected orders, which are returned in expected order of execution by default)
-   *showExpectedCancellations* is *true* to include orders that are expected to be cancelled in the next block, based on the content of the unconfirmed transactions pool and the phased transactions expected to finish in the next block (optional, does not apply to expected orders)
-   *firstIndex* is a zero-based index to the first order to retrieve (optional, does not apply to expected orders)
-   *lastIndex* is a zero-based index to the last order to retrieve (optional, does not apply to expected orders)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *bidOrders* or *askOrders* (A) is an array of order objects (refer to [Get Order](the-burst-api-get-order.md) for details) with the following additional field only for an expected order:
    -   *phased* (B) is *true* if the order is phased, *false* otherwise
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Orders](the-burst-api-examples-get-orders.md) example.

#### Get Ask Orders

Refer to [Get Orders](the-burst-api-get-orders.md).

#### Get Bid Orders

Refer to [Get Orders](the-burst-api-get-orders.md).

#### Get Expected Ask Orders

Refer to [Get Orders](the-burst-api-get-orders.md).

#### Get Expected Bid Orders

Refer to [Get Orders](the-burst-api-get-orders.md).

### Get Expected Order Cancellations

Get all expected order cancellations in the order in which they are expected to be executed.

**Request:**

-   *requestType* is *getExpectedOrderCancellations*, where expected cancellations are from the unconfirmed transactions pool or are phased transactions scheduled to finish in the next block
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *orderCancellations* (A) is an array of order cancellation objects with the following fields for each transfer:
    -   *account* (S) is the cancelling account number
    -   *accountRS* (S) is the Reed-Solomon address of the account
    -   *order* (S) is the ID of the order to be cancelled
    -   *height* (N) is the block height of the order cancellation transaction
    -   *phased* (B) is *true* if the order cancellation transaction is phased
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Expected Order Cancellations](the-burst-api-examples-get-expected-order-cancellations.md) example.

### Get Last Trades

Get the last trade of each of multiple assets.

**Request:**

-   *requestType* is *getLastTrades*
-   *assets* is one of multiple asset IDs
-   *assets* is one of multiple asset IDs

⋮

-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *trades* (A) is an array of trade objects (refer to [Get Trades](the-burst-api-get-trades.md) without *name* and *decimals* for details)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Last Trades](the-burst-api-examples-get-last-trades.md) example.

### Get Order Trades

Get all trades that were executed as a result of a given *askOrder* and/or *bidOrder* in reverse block height order.

**Request:**

-   *requestType* is *getOrderTrades*
-   *askOrder* is an ask order ID (optional)
-   *bidOrder* is a bid order ID (optional if *askOrder* provided)
-   *firstIndex* is a zero-based index to the first trade to retrieve (optional)
-   *lastIndex* is a zero-based index to the last trade to retrieve (optional)
-   *includeAssetInfo* is *true* if the *decimals* and *name* fields are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:** Refer to [Get Trades](the-burst-api-get-trades.md).

**Example:** Refer to [Get Order Trades](the-burst-api-examples-get-order-trades.md) example.

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

**Example:** Refer to [Get Trades](the-burst-api-examples-get-trades.md) example.

### Issue Asset

Create an asset on the exchange. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *issueAsset*
-   *name* is the name of the asset
-   *description* is a url-encoded description of the asset in UTF-8 with a maximum length of 1000 bytes (optional)
-   *quantityQNT* is the total amount (in QNT) of the asset in existence
-   *decimals* is the number of decimal places used by the asset (optional, zero default)

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md). The transaction ID is also the asset ID.

**Example:** Refer to [Issue Asset](the-burst-api-examples-issue-asset.md) example.

### Place Order

Place an asset order. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is either *placeBidOrder* or *placeAskOrder*
-   *asset* is the asset ID of the asset being ordered
-   *quantityQNT* is the amount (in QNT) of the asset being ordered
-   *priceNQT* is the bid/ask price (in NQT)

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md). The transaction ID is also the order ID.

**Example:** Refer to [Place Order](the-burst-api-examples-place-order.md) example.

#### Place Ask Order

Refer to [Place Order](the-burst-api-place-order.md).

#### Place Bid Order

Refer to [Place Order](the-burst-api-place-order.md).

### Search Assets

Get assets having a name or description that match a given query in reverse relevance order.

**Request:**

-   *requestType* is *searchAssets*
-   *query* is a full text query on the asset fields *name* (S) and *description* (S) in the [standard Lucene syntax](http://lucene.apache.org/core/2_9_4/queryparsersyntax.html#Overview)
-   *firstIndex* is a zero-based index to the first asset to retrieve (optional)
-   *lastIndex* is a zero-based index to the last asset to retrieve (optional)
-   *includeCounts* is *true* if the fields beginning with *numberOf...* are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *assets* (A) is an array of asset objects (refer to [Get Asset](the-burst-api-get-asset.md))
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Search Assets](the-burst-api-examples-search-assets.md) example.

### Transfer Asset

Transfer a quantity of an asset from one account to another. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *transferAsset*
-   *recipient* is the recipient account ID
-   *recipientPublicKey* is the public key of the recipient account (optional, enhances security of a new account)
-   *asset* is the ID of the asset being transferred
-   *quantityQNT* is the amount (in QNT) of the asset being transferred

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md). The transaction ID is also the transfered asset ID.

**Example:** Refer to [Transfer Asset](the-burst-api-examples-transfer-asset.md) example.

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
-   *transactions* (A) is an array of transaction IDs or transaction objects (if *includeTransactions* provided, refer to [Get Transaction](the-burst-api-get-transaction.md) for details)
-   *executedPhasedTransactions* (A) is an array of transaction IDs or transaction objects (if *includeExecutedPhased* provided, refer to [Get Transaction](the-burst-api-get-transaction.md) for details)
-   *version* (N) is the block version
-   *totalFeeNQT* (S) is the total fee (in NQT) of the transactions in the block
-   *previousBlock* (S) is the previous block ID
-   *cumulativeDifficulty* (S) is the cumulative difficulty for the next block generation
-   *block* (S) is the block ID
-   *height* (N) is the zero-based block height
-   *timestamp* (N) is the timestamp (in seconds since the genesis block) of the block
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Block](the-burst-api-examples-get-block.md) example.

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

**Example:** Refer to [Get Block Id](the-burst-api-examples-get-block-id.md) example.

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

-   *blocks* (A) is an array of blocks retrieved (refer to [Get Block](the-burst-api-get-block.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Blocks](the-burst-api-examples-get-blocks.md) example.

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

**Example:** Refer to [Get EC Block](the-burst-api-examples-get-ec-block.md) example.

Digital Goods Store Operations
------------------------------

In the [Burst client interface](burst-client-interface.md), the Digital Goods Store (DGS) is referred to as [Marketplace](marketplace.md).

### DGS Delisting

Delist a listed product. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *dgsDelisting*
-   *goods* is the goods ID

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [DGS Delisting](the-burst-api-examples-dgs-delisting.md) example.

### DGS Delivery

Deliver a product. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *dgsDelivery*
-   *purchase* is the purchase order ID
-   *discountNQT* is a discount (in NQT) off the selling price (optional, default is zero)
-   *goodsToEncrypt* is the product, a text or a hex string to be encrypted (optional if *goodsData* provided)
-   *goodsIsText* is *false* if *goodsToEncrypt* is a hex string (optional)
-   *goodsData* is AES-encrypted (using [Encrypt To](the-burst-api-encrypt-to.md)) *goodsToEncrypt*, up to 1000 bytes long (required only if *secretPhrase* is omitted)
-   *goodsNonce* is the unique nonce associated with the encrypted data (required only if *secretPhrase* is omitted)

**Note:** If the encrypted goods data is longer than 1000 bytes, use a prunable encrypted message to deliver the goods.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [DGS Delivery](the-burst-api-examples-dgs-delivery.md) example.

### DGS Feedback

Give feedback about a purchased product after delivery. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *dgsFeedback*
-   *purchase* is the purchase order ID
-   *message* is unencrypted (public) feedback text up to 1000 bytes

**Note**: The unencrypted *message* parameter is used for public feedback, but in addition or instead, an encrypted message can be used for private feedback to the seller and/or an encrypted message can be sent to self (buyer) although the current [BRS client](burst-client-interface.md) does not recognize non-public feedback messages.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [DGS Feedback](the-burst-api-examples-dgs-feedback.md) example.

### DGS Listing

List a product in the DGS by creating a listing transaction. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *dgsListing*
-   *name* is the name of the product up to 100 characters in length
-   *description* is a description of the product up to 1000 characters in length
-   *tags* are up to three comma separated keywords describing the product up to 100 characters in length (optional)
-   *quantity* is the quantity of the product for sale
-   *priceNQT* is the price (in NQT) of the product

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md). The transaction ID is also the goods ID.

**Example:** Refer to [DGS Listing](the-burst-api-examples-dgs-listing.md) example.

### DGS Price Change

Change the price of a listed product. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *dgsPriceChange*
-   *goods* is the goods ID of the product
-   *priceNQT* is the new price of the product

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [DGS Price Change](the-burst-api-examples-dgs-price-change.md) example.

### DGS Purchase

Purchase a product for sale. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *dgsPurchase*
-   *goods* is the goods ID of the product
-   *priceNQT* is the price of the product
-   *quantity* is the quantity to be purchased
-   *deliveryDeadlineTimestamp* is the timestamp (in seconds since the genesis block) by which delivery of the product must occur

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md). The transaction ID is also the purchase order ID.

**Example:** Refer to [DGS Purchase](the-burst-api-examples-dgs-purchase.md) example.

### DGS Quantity Change

Change the quantity of a listed product. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *dgsQuantityChange*
-   *goods* is the goods ID of the product
-   *deltaQuantity* is the change in the quantity of the product for sale (use negative numbers for a decrease in quantity)

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [DGS Quantity Change](the-burst-api-examples-dgs-quantity-change.md) example.

### DGS Refund

Refund a purchase. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *dgsRefund*
-   *purchase* is the purchase order ID
-   *refundNQT* is the amount (in NQT) of the refund

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [DGS Refund](the-burst-api-examples-dgs-refund.md) example.

### Get DGS Expired Purchases

Get purchase orders which have expired without being delivered, given a seller ID, in reverse chronological order.

**Request:**

-   *requestType* is *getDGSExpiredPurchases*
-   *seller* is the account ID of the product seller
-   *firstIndex* is a zero-based index to the first purchase order to retrieve (optional)
-   *lastIndex* is a zero-based index to the last purchase order to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *purchases* (A) is an array of purchase orders (refer to [Get DGS Purchase](the-burst-api-get-dgs-purchase.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get DGS Expired Purchases](the-burst-api-examples-get-dgs-expired-purchases.md) example.

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

**Example:** Refer to [Get DGS Good](the-burst-api-examples-get-dgs-good.md) example.

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

-   *goods* (A) is an array of goods (refer to [Get DGS Good](the-burst-api-get-dgs-good.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get DGS Goods](the-burst-api-examples-get-dgs-goods.md) example.

### Get DGS Goods Count

Get the number of products for sale by a given seller or all sellers.

**Request:**

-   *requestType* is *getDGSGoodsCount*
-   *seller* is the account ID of the seller (optional, default is all sellers combined)
-   *inStockOnly* is *false* if out-of-stock (zero quantity) products are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *numberOfGoods* (N) is the number of goods for sale by the *seller*
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** The *numberOfGoods* field refers to the number of distinct products for sale, regardless of the quantity of each.

**Example:** Refer to [Get DGS Goods Count](the-burst-api-examples-get-dgs-goods-count.md) example.

### Get DGS Goods Purchase Count

Get the number of completed purchase orders given a goods ID.

**Request:**

-   *requestType* is *getDGSGoodsPurchaseCount*
-   *goods* is the goods ID
-   *withPublicFeedbacksOnly* is *true* if purchase orders without public feedback are to be omitted (optional)
-   *completed* is *true* if only completed purchase orders are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *numberOfPurchases* (N) is the number of completed purchase orders
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get DGS Goods Purchase Count](the-burst-api-examples-get-dgs-goods-purchase-count.md) example.

### Get DGS Goods Purchases

Get completed purchase orders given a goods ID and optionally a buyer ID in reverse chronological order.

**Request:**

-   *requestType* is *getDGSGoodsPurchases*
-   *goods* is the goods ID
-   *buyer* is a buyer ID (optional)
-   *firstIndex* is a zero-based index to the first purchase order to retrieve (optional)
-   *lastIndex* is a zero-based index to the last purchase order to retrieve (optional)
-   *withPublicFeedbacksOnly* is *true* if purchase orders without public feedback are to be omitted (optional)
-   *completed* is *true* if only completed purchase orders are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *purchases* (A) is an array of purchase orders (refer to [Get DGS Purchase](the-burst-api-get-dgs-purchase.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get DGS Goods Purchases](the-burst-api-examples-get-dgs-goods-purchases.md) example.

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

-   *purchases* (A) is an array of pending purchase orders (refer to [Get DGS Purchase](the-burst-api-get-dgs-purchase.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get DGS Pending Purchases](the-burst-api-examples-get-dgs-pending-purchases.md) example.

### Get DGS Purchase

Get a purchase order given a purchase order ID.

**Request:**

-   *requestType* is *getDGSPurchase*
-   *purchase* is the purchase order ID
-   *sharedKey* is the shared key used to decrypt the message (optional) (see [Get Shared Key](the-burst-api-get-shared-key.md))
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

**Example:** Refer to [Get DGS Purchase](the-burst-api-examples-get-dgs-purchase.md) example.

### Get DGS Purchase Count

Get the number of purchase orders given a seller and/or buyer ID, or all orders combined.

**Request:**

-   *requestType* is *getDGSPurchaseCount*
-   *seller* is the account ID of the seller (optional, default is all sellers)
-   *buyer* is the account ID of the buyer (optional, default is all buyers)
-   *withPublicFeedbacksOnly* is *true* if purchase orders without public feedback are to be omitted (optional)
-   *completed* is *true* if only completed purchase orders are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *numberOfPurchases* (N) is the number of purchase orders associated with the seller and/or buyer
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get DGS Purchase Count](the-burst-api-examples-get-dgs-purchase-count.md) example.

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

-   *purchases* (A) is an array of purchase orders (refer to [Get DGS Purchase](the-burst-api-get-dgs-purchase.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get DGS Purchases](the-burst-api-examples-get-dgs-purchases.md) example.

### Get DGS Tag Count

Get the number of tags used by all sellers.

**Request:**

-   *requestType* is *getDGSTagCount*
-   *inStockOnly* is *false* if tags with no associated in-stock products are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *numberOfTags* (N) is the number of tags
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get DGS Tag Count](the-burst-api-examples-get-dgs-tag-count.md) example.

### Get DGS Tags

Get tags used by all sellers in reverse *inStockCount*, reverse *totalCount*, *tag* order.

**Request:**

-   *requestType* is *getDGSTags*
-   *inStockOnly* is *false* if out-of-stock tags are to be retrieved (optional)
-   *firstIndex* is a zero-based index to the first tag to retrieve (optional)
-   *lastIndex* is a zero-based index to the last tag to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *tags* (A) is an array of tag objects with the following fields for each tag:
    -   *inStockCount* (N) is the number of products available for sale as tagged
    -   *tag* (S) is the tag word
    -   *totalCount* (N) is the total number of products as tagged
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** The *...Count* fields refer to the number of distinct products tagged, regardless of the quantity of each.

**Example:** Refer to [Get DGS Tags](the-burst-api-examples-get-dgs-tags.md) example.

### Get DGS Tags Like

Get all tags starting with a given prefix (at least 2 characters long) in reverse *inStockCount*, reverse *totalCount*, *tag* order.

**Request:**

-   *requestType* is *getDGSTagsLike*
-   *tagPrefix* is the prefix (at least 2 characters long) of the *tag*
-   *inStockOnly* is *false* if out-of-stock tags are to be retrieved (optional)
-   *firstIndex* is a zero-based index to the first tag to retrieve (optional)
-   *lastIndex* is a zero-based index to the last tag to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *tags* (A) is an array of tag objects with the following fields for each tag:
    -   *inStockCount* (N) is the number of products available for sale as tagged
    -   *tag* (S) is the tag word
    -   *totalCount* (N) is the total number of products as tagged
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** The *...Count* fields refer to the number of distinct products tagged, regardless of the quantity of each.

**Example:** Refer to [Get DGS Tags Like](the-burst-api-examples-get-dgs-tags-like.md) example.

### Search DGS Goods

Get product listings that have a name or description that match a given query in reverse relevance order, then name order (given a seller), then reverse chronological order.

**Request:**

-   *requestType* is *searchDGSGoods*
-   *query* is a full text query on the goods fields *name* and *description* in the [standard Lucene syntax](http://lucene.apache.org/core/2_9_4/queryparsersyntax.html#Overview) (optional)
-   *tag* is a query on the good field *tags* in the [standard Lucene syntax](http://lucene.apache.org/core/2_9_4/queryparsersyntax.html#Overview) (optional)
-   *seller* is the account ID of the product seller (optional)
-   *firstIndex* is a zero-based index to the first product to retrieve (optional)
-   *lastIndex* is a zero-based index to the last product to retrieve (optional)
-   *inStockOnly* is *false* if out-of-stock products (zero quantity) are to be retrieved (optional)
-   *hideDelisted* is *true* if delisted products are to be omitted (optional)
-   *includeCounts* is *true* if the fields beginning with *numberOf...* are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *goods* (A) is an array of goods objects (refer to [Get DGS Good](the-burst-api-get-dgs-good.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Search DGS Goods](the-burst-api-examples-search-dgs-goods.md) example.

Forging Operations
------------------

### Start / Stop / Get Forging

Start or stop forging with an account, or check to see if an account is forging. POST only.

**Request:**

-   *requestType* is either *startForging*, *stopForging* or *getForging*
-   *secretPhrase* is the secret passphrase of the account (optional for *stopForging* and *getForging* if password protected like the [Debug Operations](the-burst-api-debug-operations.md))

**Response:**

-   *deadline* (N) is the estimated time (in seconds since the last block) until the account will forge a block (*startForging* and *getForging* only)
-   *hitTime* (N) is the estimated time (in seconds since the genesis block) when the account will forge a block (*startForging* and *getForging* only)
-   *remaining* (N) is the deadline less the elapsed time since the last block (*getForging* only)
-   *foundAndStopped* (B) is *true* if forging was stopped, *false* if forging was already stopped (*stopForging* only)
-   *account* (S) is the account number (*getForging* only)
-   *accountRS* (S) is the Reed-Solomon address of the account (*getForging* only)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** A *getForging* request returns *errorCode* 5 if the account is not forging. If the account has a zero *effectiveBalance*, forging can be started but *deadline*, *remainingTime* and *hitTime* will be set to zero.

**Example:** Refer to [Start / Stop / Get Forging](the-burst-api-examples-start---stop---get-forging.md) example.

#### Get Forging

Refer to [Start / Stop / Get Forging](the-burst-api-start--2f-stop--2f-get-forging.md).

#### Start Forging

Refer to [Start / Stop / Get Forging](the-burst-api-start--2f-stop--2f-get-forging.md).

#### Stop Forging

Refer to [Start / Stop / Get Forging](the-burst-api-start--2f-stop--2f-get-forging.md).

### Lease Balance

[Lease](account-leasing.md) the entire guaranteed balance of NXT to another account, after 1440 confirmations. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *leaseBalance*
-   *period* is the lease period (in number of blocks, 1440 minimum)
-   *recipient* is the lessee (recipient) account
-   *recipientPublicKey* is the public key of the lessee (recipient) account (optional, enhances security of a new account)

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Lease Balance](the-burst-api-examples-lease-balance.md) example.

### Get Next Block Generators

Returns the next block generators ordered by hit time. The list of currently active forgers is first initialized using the block generators with at least 2 blocks generated within the previous 10,000 blocks, excluding accounts without a public key. The list is updated as new blocks are processed. The results are not 100% correct since previously active generators may no longer be running and new generators won't be known until they generate a block.

**Request:**

-   *requestType* is *getNextBlockGenerators*
-   *limit* (N) is the number of next block generators to display.

**Response:**

-   *activeCount* (N) is the number of active forging accounts
-   *lastBlock* (S) is the last block ID on the blockchain
-   *generators* (A) is an array containing the number of next block generators requested
    -   *effectiveBalanceNXT* (N) is the balance (in NXT) of the account available for forging: the unleased guaranteedBalance of this account plus the leased guaranteedBalance of all lessors to this account
    -   *accountRS* (S) is the Reed-Solomon address of the account
    -   *deadline* (N) is the estimated time (in seconds since the last block) until the account will forge a block
    -   *account* (S) is the account number
    -   *hitTime* (N) is the estimated time (in seconds since the genesis block) when the account will forge a block
-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *timestamp* (N) is the timestamp (in seconds since the genesis block) when the request was executed
-   *height* (N) is the height of the blockchain

**Example:** Refer to [Get Next Block Generators](the-burst-api-examples-get-next-block-generators.md) example.

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

**Example:** Refer to [Decode Hallmark](the-burst-api-examples-decode-hallmark.md) example.

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

**Example:** Refer to [Mark Host](the-burst-api-examples-mark-host.md) example.

#### Generate Hallmark

Refer to [Mark Host](the-burst-api-mark-host.md).

Monetary System Operations
--------------------------

General Monetary System documentation is available [here](https://bitbucket.org/JeanLucPicard/nxt/issue/205/monetary-system-documentation). Documentation on the MintWorker tool for currency minting is available [here](https://bitbucket.org/JeanLucPicard/nxt/issue/207/mint-worker-utility).

### Can Delete Currency

Determine if a currency can be deleted.

**Request:**

-   *requestType* is *canDeleteCurrency*
-   *account* is the account ID
-   *currency* is the currency ID
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *canDelete* (B) is *true* if the currency can be deleted, *false* otherwise
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** a currency can be deleted only when all units of the currency are held by *account*. A reservable currency that has not yet been issued can be deleted by the issuer. A mintable currency that has not completed minting cannot be deleted by a non-issuer.

**Example:** Refer to [Can Delete Currency](the-burst-api-examples-can-delete-currency.md) example.

### Currency Buy / Sell

Make an exchange request to buy or sell an exchangeable currency. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *currencyBuy* or *currencySell*
-   *currency* is the currency ID
-   *rateNQT* is the exchange rate (in NQT per QNT)
-   *units* is the amount of the currency to buy or sell (in QNT)

**Note:** An exchange request is immediately executed once accepted onto the blockchain based only on currently available offers (refer to [Publish Exchange Offer](the-burst-api-publish-exchange-offer.md)). The request then expires, regardless of the amount of currency exchanged; the request may be completely filled, partially filled, or expire without any exchange if no matching offers are found.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Currency Buy / Sell](the-burst-api-examples-currency-buy---sell-.md) example.

#### Currency Buy

Refer to [Currency Buy / Sell](the-burst-api-currency-buy--2f-sell.md).

#### Currency Sell

Refer to [Currency Buy / Sell](the-burst-api-currency-buy--2f-sell.md).

### Currency Mint

Submit a valid computed nonce to the blockchain in return for newly minted currency. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *currencyMint*
-   *currency* is the mintable currency ID
-   *nonce* is the computed nonce
-   *units* is the amount (in QNT) of currency to mint
-   *counter* (N) is the counter associated with the minting account

**Note:** The hash of *nonce* must be less than *targetBytes* provided by [Get Minting Target](the-burst-api-get-minting-target.md) for given *units* and *counter*. *counter* must be increased with each submission.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Currency Mint](the-burst-api-examples-currency-mint.md) example.

### Currency Reserve Claim

Claim currency reserve. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *currencyReserveClaim*
-   *currency* is the currency ID
-   *units* is the amount (in QNT) of reserve currency to claim

**Note:** Holders of a claimable currency may claim the locked NQT backing their units, thus reducing the supply of the currency.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Currency Reserve Claim](the-burst-api-examples-currency-reserve-claim.md) example.

### Currency Reserve Increase

Increase the currency reserve of an unissued reservable currency. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *currencyReserveIncrease*
-   *currency* is the currency ID
-   *amountPerUnitNQT* is the additional amount (in NQT per QNT of *reserveSupply*) to reserve (refer to [Issue Currency](the-burst-api-issue-currency.md))

**Note:** An additional *amountPerUnitNQT* \* *reserveSupply* NQT (beyond what has previously been reserved) will be locked until the *issuanceHeight* is reached. Upon issuance, if the currency is claimable the NQT will remain locked until claimed; otherwise the NQT will transfer to the issuing account. Also upon issuance, a portion of the *reserveSupply* QNT will be transfered to each reserving account in proportion to the NQT that was contributed.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Currency Reserve Increase](the-burst-api-examples-currency-reserve-increase.md) example.

### Delete Currency

Delete a deletable currency (refer to [Can Delete Currency](the-burst-api-can-delete-currency.md)). POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *deleteCurrency*
-   *currency* is the currency ID

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Delete Currency](the-burst-api-examples-delete-currency.md) example.

### Get Account Currencies

Get the currencies issued by a given account.

**Request:**

-   *requestType* is *getAccountCurrencies*
-   *account* is the account ID
-   *currency* is a currency ID filter (optional)
-   *height* is the blockchain height at which the response applies (optional, default is the current height)
-   *includeCurrencyInfo* is *true* if several currency information properties is to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *accountCurrencies* (A) is an array of currency objects with the following fields:
    -   *code* (S) is the currency code
    -   *unconfirmedUnits* (S) is the amount of unconfirmed currency units (in QNT)
    -   *decimals* (N) is the number of currency decimal places
    -   *name* (S) is the currency name
    -   *currency* (S) is the currency ID
    -   *units* (S) is the amount of currency (in QNT)
    -   *issuanceHeight* (N) is the blockchain height of issuance for a reservable currency
    -   *type* (N) is the currency type bitmask (refer to [Get Currency](the-burst-api-get-currency.md))
    -   *issuerAccountRS* (S) is the Reed-Solomon address of the issuer account
    -   *issuerAccount* (S) is the issuer account ID
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Currencies](the-burst-api-examples-get-account-currencies.md) example.

### Get Account Currency Count

Get the number of currencies issued by a given account.

**Request:**

-   *requestType* is *getAccountCurrencyCount*
-   *account* is the account ID
-   *height* is the blockchain height at which the response applies (optional, default is the current height)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *numberOfCurrencies* (N) is the number of currencies issued
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Currency Count](the-burst-api-examples-get-account-currency-count.md) example.

### Get Account Exchange Requests

Get the exchange requests associated with a given account and/or currency in reverse chronological order (or in expected order of execution for expected requests).

**Request:**

-   *requestType* is either *getAccountExchangeRequests* or *getExpectedExchangeRequests*, where expected requests are from the unconfirmed transactions pool or are phased transactions scheduled to finish in the next block
-   *account* is the account ID (optional for expected requests)
-   *currency* is the currency ID (optional for expected requests if *account* provided)
-   *includeCurrencyInfo* is *true* to include the response fields *code*, *decimals*, *name*, *issuanceHeight*, *type*, *timestamp*, *issuerAccountRS* and *issuerAccount* (optional, applies only to expected requests)
-   *firstIndex* is a zero-based index to the first request to retrieve (optional, does not apply to expected requests)
-   *lastIndex* is a zero-based index to the last request to retrieve (optional, does not apply to expected requests)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *exchangeRequests* (A) is an array of requests with the following fields for each request:
    -   *code* (S) is a currency code
    -   *subtype* (N) is *5* for buy and *6* for sell
    -   *decimals* (N) is the number of decimal places
    -   *name* (S) is the currency name
    -   *units* (S) is the number of currency units to buy or sell (in QNT)
    -   *issuanceHeight* (N) is the blockchain height of issuance for a reservable currency, zero otherwise
    -   *type* (N) is the currency type bitmask (refer to [Get Currency](the-burst-api-get-currency.md))
    -   *transaction* (S) is the transaction ID
    -   *timestamp* (N) is the timestamp (in seconds since the genesis block) of the block when the request was executed
    -   *rateNQT* (S) is the exchange rate (in NQT per QNT)
    -   *issuerAccountRS* (S) is the Reed-Solomon address of the issuer account
    -   *issuerAccount* (S) is the issuer account ID
    -   *phased* (B) is true if the transaction is phased (applies only to expected requests)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Notes:** Even expired, unfilled requests will be retrieved. As of [Nxt Software Version 1.5.13](nxt-software-change-log-version-1-5-13.md) phased exchange requests that have not yet finished, or have not been approved, are now excluded if *requestType* is *getAccountExchangeRequests*.

**Example:** Refer to [Get Account Exchange Requests](the-burst-api-examples-get-account-exchange-requests.md) example.

#### Get Expected Exchange Requests

Refer to [Get Account Exchange Requests](the-burst-api-get-account-exchange-requests.md).

### Get Funding Monitor

Get a funding monitor.

**Request:**

-   *requestType* is *getFundingMonitor*
-   *secretPhrase* is the secret phrase of the funding account, used to get a single monitor. (optional)
-   *adminPassword* is the admin password, used to get a single monitor or all monitors (optional if *secretPhrase* is provided)
-   *includeMonitoredAccounts* is *true* to include account info of the monitored accounts (optional)
-   *property* is the name of the account property (optional)
-   *holdingType* is a string representing the holding type (optional)
-   *holding* is the holding ID (optional)
-   *account* is the account ID (optional)

**Response:**

-   *monitors* (A) is an array of monitor objects including the following fields:
    -   *holding* (S) is the holding ID
    -   *amount* (S) is the amount to fund the monitored accounts with (NQT or QNT)
    -   *monitoredAccounts* (A) is an array of monitored account objects including the following fields (only if *includeMonitoredAccounts* is *true*):
        -   *amount* (S) is the amount to fund the monitored accounts with. Overrides amount in parent object.
        -   *account* (S) is the monitored account ID
        -   *accountRS* (S) is the monitored account Reed Solomon address
        -   *threshold* (S) is the threshold. Overrides threshold in parent object.
        -   *interval* (N) is the interval. Overrides interval in parent object.
    -   *holdingType* (N) is the holding type
    -   *account* (S) is the monitoring account ID
    -   *accountRS* (S) is the Reed Solomon address of the monitoring account
    -   *property* (S) is the account property
    -   *threshold* (S) is the threshold
    -   *interval* (N) is the interval
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Funding Monitor](the-burst-api-examples-get-funding-monitor.md) example.

### Get All Currencies

Get all currencies in reverse creation order.

**Request:**

-   *requestType* is *getAllCurrencies*
-   *firstIndex* is a zero-based index to the first currency to retrieve (optional)
-   *lastIndex* is a zero-based index to the last currency to retrieve (optional)
-   *includeCounts* is *true* to include *numberOf...* fields (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *currencies* (A) is an array of currency objects (refer to [Get Currency](the-burst-api-get-currency.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get All Currencies](the-burst-api-examples-get-all-currencies.md) example.

### Get All Exchanges

Get all currency exchanges in reverse chronological order.

**Request:**

-   *requestType* is *getAllExchanges*
-   *timestamp* is the earliest timestamp to retrieve (optional)
-   *firstIndex* is a zero-based index to the first exchange to retrieve (optional)
-   *lastIndex* is a zero-based index to the last exchange to retrieve (optional)
-   *includeCurrencyInfo* is *true* to include some currency details (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *exchanges* (A) is an array of exchange objects (refer to [Get Exchanges](the-burst-api-get-exchanges.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get All Exchanges](the-burst-api-examples-get-all-exchanges.md) example.

### Get Available To Buy

Calculates the rate required in order to completely fill an exchange request.

**Request:**

-   *requestType* is *getAvailableToBuy* or *getAvailableToSell*
-   *currency* is the currency ID
-   *units* is the number of units to buy
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *amountNQT* (S) is the total amount needed to fill the exchange request
-   *units* (S) is the number of units
-   *rateNQT* (S) is the rate for the currency units
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Available To Buy](the-burst-api-examples-get-available-to-buy.md) example.

#### Get Available To Sell

Refer to [Get Available To Buy](the-burst-api-get-available-to-buy.md).

### Get Buy / Sell Offers

Get currency buy or sell offers given a currency ID and/or an account ID in order of rate (if *sortByRate* is *true* for expected offers, otherwise in the expected order of execution).

**Request:**

-   *requestType* is one of *getBuyOffers*, *getSellOffers*, *getExpectedBuyOffers* or *getExpectedSellOffers*, where expected offers are from the unconfirmed transactions pool or are phased transactions scheduled to finish in the next block
-   *currency* is the currency ID (optional)
-   *account* is the account ID (optional if *currency* provided)
-   *availableOnly* is *true* to include only offers with non-zero supply and limit, but is ignored when both *currency* and *account* are given (optional, does not apply to expected offers)
-   *sortByRate* is *true* to sort by rate (optional, applies only to expected offers, which are returned in expected order of execution by default)
-   *firstIndex* is a zero-based index to the first offer to retrieve (optional, does not apply to expected offers)
-   *lastIndex* is a zero-based index to the last offer to retrieve (optional, does not apply to expected offers)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *offers* (A) is an array of buy or sell offer objects (refer to [Get Offer](the-burst-api-get-offer.md) for details) with the following additional field only for an expected offer:
    -   *phased* (B) is *true* if the offer is phased, *false* otherwise
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Buy / Sell Offers](the-burst-api-examples-get-buy---sell-offers.md) example.

#### Get Buy Offers

Refer to [Get Buy / Sell Offers](the-burst-api-get-buy--2f-sell-offers.md).

#### Get Sell Offers

Refer to [Get Buy / Sell Offers](the-burst-api-get-buy--2f-sell-offers.md).

#### Get Expected Buy Offers

Refer to [Get Buy / Sell Offers](the-burst-api-get-buy--2f-sell-offers.md).

#### Get Expected Sell Offers

Refer to [Get Buy / Sell Offers](the-burst-api-get-buy--2f-sell-offers.md).

### Get Currencies

Get currencies given multiple currency IDs.

**Request:**

-   *requestType* is *getCurrencies*
-   *currencies* is one of multiple currency IDs
-   *currencies* is one of multiple currency IDs

⋮

-   *includeCounts* is *true* to include *numberOf...* fields (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *currencies* (A) is an array of currency objects (refer to [Get Currency](the-burst-api-get-currency.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Currencies](the-burst-api-examples-get-currencies.md) example.

### Get Currencies By Issuer

Get currencies issued by multiple accounts in reverse creation order.

**Request:**

-   *requestType* is *getCurrenciesByIssuer*
-   *account* is one of multiple account IDs
-   *account* is one of multiple account IDs

⋮

-   *firstIndex* is a zero-based index to the first currency to retrieve (optional)
-   *lastIndex* is a zero-based index to the last currency to retrieve (optional)
-   *includeCounts* is *true* if *numberOf...* fields are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *currencies* (A) is an array of arrays of currency objects, where the outer array is indexed by the given account IDs (refer to [Get Currency](the-burst-api-get-currency.md) for details about the currency objects)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Currencies By Issuer](the-burst-api-examples-get-currencies-by-issuer.md) example.

### Get Currency

Get the details of a currency.

**Request:**

-   *requestType* is *getCurrency*
-   *currency* is the currency ID (optional)
-   *code* is the currency code (optional if *currency* provided)
-   *includeCounts* is *true* if *numberOf...* fields are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *initialSupply* (S) is the initial currency supply (in QNT)
-   *currentReservePerUnitNQT* (S) is the minimum currency reserve (in NQT per QNT)
-   *types* (A) is an array of currency types, one or more of:
    -   *EXCHANGEABLE*
    -   *CONTROLLABLE*
    -   *RESERVABLE*
    -   *CLAIMABLE*
    -   *MINTABLE*
    -   *NON\_SHUFFLEABLE*
-   *code* (S) is the currency code
-   *creationHeight* (N) is the blockchain height of the currency creation
-   *minDifficulty* (N) is the minimum difficulty for a mintable currency
-   *numberOfTransfers* (N) is the number of currency transfers
-   *description* (S) is the currency description
-   *minReservePerUnitNQT* (S) is the minimum currency reserve (in NQT per QNT) for a reservable currency
-   *currentSupply* (S) is the current currency supply (in QNT)
-   *issuanceHeight* (N) is the blockchain height of the currency issuance for a reservable currency
-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *type* (N) is the currency type bitmask, from least to most significant bit: exchangeable, controllable, reservable, claimable, mintable, non-shuffleable
-   *reserveSupply* (S) is the reserve currency supply (in NQT) for a reservable currency
-   *maxDifficulty* (N) is the maximum difficulty for a mintable currency
-   *accountRS* (S) is the Reed-Solomon address of the issuing account
-   *decimals* (N) is the number of decimal places used by the currency
-   *name* (S) is the name of the currency
-   *numberOfExchanges* (N) is the number of currency exchanges
-   *currency* (S) is the currency ID
-   *maxSupply* (S) is the maximum currency supply (in QNT)
-   *account* (S) is the account ID of the currency issuer
-   *algorithm* (N) is the algorithm number for a mintable currency: 2 for SHA256, 3 for SHA3, 5 for Scrypt, 25 for Keccak25
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)

**Example:** Refer to [Get Currency](the-burst-api-examples-get-currency.md) example.

### Get Currency Account Count

Get the number of accounts that hold a given currency.

**Request:**

-   *requestType* is *getCurrencyAccountCount*
-   *currency* is the currency ID
-   *height* is the blockchain height at which the response applies (optional, default is the current height)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *numberOfAccounts* (N) is the number of accounts that hold the currency
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Currency Account Count](the-burst-api-examples-get-currency-account-count.md) example.

### Get Currency Accounts

Get the accounts that hold a given currency in reverse units order.

**Request:**

-   *requestType* is *getCurrencyAccounts*
-   *currency* is the currency ID
-   *height* is the blockchain height at which the response applies (optional, default is current height)
-   *firstIndex* is a zero-based index to the first account to retrieve (optional)
-   *lastIndex* is a zero-based index to the last account to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *accountCurrencies* (A) is an array of account objects with the following fields:
    -   *unconfirmedUnits* (S) is the amount of unconfirmed currency units (in QNT)
    -   *accountRS* (S) is the Reed-Solomon address of the account
    -   *currency* (S) is the currency ID
    -   *units* (S) is the amount of currency (in QNT)
    -   *account* (S) is the account number
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Currency Accounts](the-burst-api-examples-get-currency-accounts.md) example.

### Get Currency Founders

Get a reservable currency's founders.

**Request:**

-   *requestType* is *getCurrencyFounders*
-   *currency* is the currency ID
-   *account* is an account ID (optional)
-   *firstIndex* is a zero-based index to the first founder to retrieve (optional)
-   *lastIndex* is a zero-based index to the last founder to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *founders* (A) is an array of founder objects each of which has the following fields:
    -   *accountRS* (S) is the Reed-Solomon address of the founding account
    -   *amountPerUnitNQT* (S) is the amount (in NQT per QNT of *reserveSupply*) reserved by the founder
    -   *currency* (S) is the currency ID
    -   *account* (S) is the founding account number
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Currency Founders](the-burst-api-examples-get-currency-founders.md) example.

### Get Currency Ids

Get all currency IDs in reverse chronological creation order.

**Request:**

-   *requestType* is *getCurrencyIds*
-   *firstIndex* is a zero-based index to the first currency to retrieve (optional)
-   *lastIndex* is a zero-based index to the last currency to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *currencyIds*(A) is an array of currency IDs
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Currency Ids](the-burst-api-examples-get-currency-ids.md) example.

### Get Currency Transfers

Get currency transfers given a currency ID and/or an account ID in reverse block height order (or in expected order of execution for expected transfers).

**Request:**

-   *requestType* is either *getCurrencyTransfers* or *getExpectedCurrencyTransfers*, where expected transfers are from the unconfirmed transactions pool or are phased transactions scheduled to finish in the next block
-   *currency* is the currency ID (optional)
-   *account* is the account ID (optional if *currency* provided)
-   *timestamp* is the earliest transfer (in seconds since the genesis block) to retrieve (optional, does not apply to expected transfers)
-   *firstIndex* is a zero-based index to the first transfer to retrieve (optional, does not apply to expected transfers)
-   *lastIndex* is a zero-based index to the last transfer to retrieve (optional, does not apply to expected transfers)
-   *includeCurrencyInfo* is *true* to include some currency fields (optional, does not apply to expected transfers)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *transfers* (A) is an array of transfer objects with the following fields for each transfer:
    -   *code* (S) is the currency code
    -   *units* (S) is the amount (in QNT) of the transfer
    -   *issuanceHeight* (N) is the blockchain height of the currency issuance for a reservable currency
    -   *type* (N) is the currency type bitmask (refer to [Get Currency](the-burst-api-get-currency.md) for details)
    -   *issuerAccountRS* (S) is the Reed-Solomon address of the issuer account
    -   *transfer* (S) is the transfer ID
    -   *senderRS* (S) is the Reed-Solomon address of the sender account
    -   *sender* (S) is the account number of the sender account
    -   *recipientRS* (S) is the Reed-Solomon address of the recipient account
    -   *decimals* (N) is the number of decimal places used by the currency
    -   *recipient* (S) is the account number of the recipient account
    -   *name* (S) is the currency name
    -   *currency* (S) is the currency ID
    -   *issuerAccount* (S) is the issuer account ID
    -   *height* (N) is the blockchain height of the transfer
    -   *timestamp* (N) is the timestamp (in seconds since the genesis block) of the transfer block, does not apply to an expected transfer
    -   *phased* (B) is *true* if the transaction is phased, *false* otherwise, applies only to an expected transfer
    -   *issuerAccountRS* (S) is the Reed-Solomon address of the issuer account
    -   *issuerAccount* (S) is the issuer account ID
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Currency Transfers](the-burst-api-examples-get-currency-transfers.md) example.

#### Get Expected Currency Transfers

Refer to [Get Currency Transfers](the-burst-api-get-currency-transfers.md).

### Get Exchanges

Get currency exchanges given a currency and/or an account in reverse chronological order.

**Request:**

-   *requestType* is *getExchanges*
-   *currency* is a currency ID (optional)
-   *account* is an account ID (optional if *currency* provided)
-   *firstIndex* is a zero-based index to the first currency exchange to retrieve (optional)
-   *lastIndex* is a zero-based index to the last currency exchange to retrieve (optional)
-   *timestamp* is the earliest block (in seconds since the genesis block) to retrieve (optional)
-   *includeCurrencyInfo* is *true* to include several currency-related fields (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *exchanges* (A) is an array of exchange objects with the following fields for each exchange:
    -   *seller* (S) is the seller account number
    -   *code* (S) is the currency code
    -   *sellerRS* (S) is the Reed-Solomon address of the seller account
    -   *units* (S) is the amount of currency exchanged (in QNT)
    -   *issuanceHeight* (N) is the blockchain height of currency issuance for a reservable currency
    -   *type* (N) is the currency type bitmask (refer to [Get Currency](the-burst-api-get-currency.md) for details)
    -   *rateNQT* (S) is the currency exchange rate (in NQT per QNT)
    -   *buyer* (S) is the account number of the buyer
    -   *offer* (S) is the offer ID
    -   *buyerRS* (S) is the Reed-Solomon address of the buyer account
    -   *decimals* (N) is the number of decimal places used by the currency
    -   *name* (S) is the currency name
    -   *currency* (S) is the currency ID
    -   *block* (S) is the block ID of the block containing the exchange transaction
    -   *transaction* (S) is the transaction ID of the exchange
    -   *timestamp* (N) is the timestamp (in seconds since the genesis block) of the exchange
    -   *height* is the blockchain height of the block containing the exchange transaction
    -   *issuerAccountRS* (S) is the Reed-Solomon address of the issuer account
    -   *issuerAccount* (S) is the issuer account ID
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Exchanges](the-burst-api-examples-get-exchanges.md) example.

### Get Exchanges By Exchange Request

Get currency exchanges given an exchange request transaction ID in reverse chronological order.

**Request:**

-   *requestType* is *getExchangesByExchangeRequest*
-   *transaction* is the transaction ID of the exchange request
-   *includeCurrencyInfo* is *true* to include several currency-related fields (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *exchanges* (A) is an array of exchange objects (refer to [Get Exchanges](the-burst-api-get-exchanges.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Exchanges By Exchange Request](the-burst-api-examples-get-exchanges-by-exchange-request.md) example.

### Get Exchanges By Offer

Get currency exchanges given a currency offer ID in reverse chronological order.

**Request:**

-   *requestType* is *getExchangesByOffer*
-   *offer* (S) is a currency offer ID
-   *includeCurrencyInfo* is *true* to include several currency-related fields (optional)
-   *firstIndex* is a zero-based index to the first currency exchange to retrieve (optional)
-   *lastIndex* is a zero-based index to the last currency exchange to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *exchanges* (A) is an array of exchange objects (refer to [Get Exchanges](the-burst-api-get-exchanges.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Exchanges By Offer](the-burst-api-examples-get-exchanges-by-offer.md) example.

### Get Last Exchanges

Get the last exchange of each of multiple currencies.

**Request:**

-   *requestType* is *getLastExchanges*
-   *currencies* is one of multiple currency IDs
-   *currencies* is one of multiple currency IDs

⋮

-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *exchanges* (A) is an array of exchange objects (refer to [Get Exchanges](the-burst-api-get-exchanges.md) without *name*, *decimals*, *code*, *issuanceHeight*, *type*, *issuerAccountRS* and *issuerAccount* for details)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Last Exchanges](the-burst-api-examples-get-last-exchanges.md) example.

### Get Minting Target

Get the current minting target of a mintable currency.

**Request:**

-   *requestType* is *getMintingTarget*
-   *currency* is the mintable currency ID
-   *account* is the minting account ID
-   *units* is the amount (in QNT) of currency to mint
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** *units* cannot be greater than 1/10000 of the *maxSupply* (refer to [Issue Currency](the-burst-api-issue-currency.md)). Increasing *units* decreases *targetBytes*.

**Response:**

-   *difficulty* (S) is the current difficulty, an estimate of the number of hashes needed to meet the target
-   *targetBytes* (S) is the 32-byte target (little endian), which must equal or exceed the computed hash of the *nonce*
-   *currency* (S) is the currency ID
-   *counter* (N) is the counter associated with the minting account, the value previously submitted to [Currency Mint](the-burst-api-currency-mint.md)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** If a *nonce* is found such that its hash is less than the target, it can be submitted to the blockchain along with *counter* + 1 using [Currency Mint](the-burst-api-currency-mint.md), which results in *units* NQT being credited to the minting account. *difficulty* is inversely related to the target, and so increases exponentially as *maxSupply* is approached because the target is defined as (2<sup>exp</sup>-1)/*units*, where exp decreases linearly from 256-*minDifficulty* to 256-*maxDifficulty*. (Refer to [Issue Currency](the-burst-api-issue-currency.md) for *maxSupply*, *minDifficulty* and *maxDifficulty*.)

**Example:** Refer to [Get Minting Target](the-burst-api-examples-get-minting-target.md) example.

### Get Offer

Get offer details given an offer ID.

**Request:**

-   *requestType* is *getOffer*
-   *offer* is the offer ID
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *sellOffer* and *buyOffer* (O) are objects with the following fields:
    -   *offer* (S) is the offer ID
    -   *expirationHeight* (N) is the blockchain height of offer expiration
    -   *accountRS* (S) is the Reed-Solomon address of the offering account
    -   *limit* (S) is the cumulative limit of currency buys or sells
    -   *currency* (S) is the currency ID
    -   *supply* (S) is the current currency supply
    -   *account* (S) is the offering account number
    -   *height* (N) is the blockchain height of offer creation
    -   *rateNQT* (S) is the currency exchange rate (in NQT per QNT)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Offer](the-burst-api-examples-get-offer.md) example.

### Issue Currency

Issue a new currency or re-issue an existing currency with different properties. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *issueCurrency*
-   *name* is the currency name, 3 to 10 characters and longer than the currency code
-   *code* is the uppercase currency code with the following fee structure: three letters 25000 NXT, four letters 1000 NXT, five letters 40 NXT, re-issue 40 NXT
-   *description* is the currency description
-   *type* is the currency type bitmask, from least to most significant bit: exchangeable, controllable, reservable, claimable, mintable, non-shuffleable
-   *initialSupply* is the initial currency supply (in QNT) (must match *maxSupply* unless mintable or claimable, must be zero for claimable)
-   *reserveSupply* is the reserve currency supply (in QNT) (must match *maxSupply*)
-   *maxSupply* is the maximum currency supply (in QNT)
-   *issuanceHeight* is the blockchain height at which a reservable currency is issued if the reserve minimum is met
-   *minReservePerUnitNQT* is the minimum currency reserve (in NQT per QNT of *reserveSupply*) for issuance of a reservable currency
-   *minDifficulty* is the minimum difficulty (minimum *1*) for a mintable currency
-   *maxDifficulty* is the maximum difficulty (maximum *255* and greater than *minDifficulty*) for a mintable currency
-   *ruleset* is for future use, always set to zero
-   *algorithm* is an algorithm code for a mintable currency: *2* for SHA256, *3* for SHA3, *5* for Scrypt, *25* for Keccak25
-   *decimals* is the number of decimal places used by the currency (optional, default zero)

**Notes:** Reservable requires exchangeable and/or claimable, as does controllable; but mintable requires exchangeable. Claimable requires reservable, non-mintable and zero *initialSupply*.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md). The transaction ID is also the currency ID.

**Example:** Refer to [Issue Currency](the-burst-api-examples-issue-currency.md) example.

### Publish Exchange Offer

Publish an exchange offer for an exchangeable currency. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *publishExchangeOffer*
-   *currency* is the currency ID
-   *buyRateNQT* is the offered buy rate (in NQT per QNT)
-   *sellRateNQT* is the offered sell rate (in NQT per QNT)
-   *totalBuyLimit* is the cumulative limit (in QNT) of currency buys
-   *totalSellLimit* is the cumulative limit (in QNT) of currency sells
-   *initialBuySupply* is the initial amount (in QNT) of currency offered to buy, cannot exceed *totalBuyLimit*
-   *initialSellSupply* is the initial amount (in QNT) of currency offered to sell, cannot exceed *totalSellLimit*
-   *expirationHeight* is the blockchain height for expiration of the offer

**Notes:** Each time currency is bought in response to an exchange request to sell currency (refer to [Currency Sell](the-burst-api-currency-buy--2f-sell.md)), *totalBuyLimit* is reduced and the supply of currency offered to sell increases by the amount bought. When *totalBuyLimit* becomes zero, the buy offer is withdrawn. These same notes apply if *buy* and *sell* are interchanged. Only the most recent offer associated with an account is valid, even if an earlier offer by that account has not yet expired or reached its limits.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md). The transaction ID is also the offer ID.

**Example:** Refer to [Publish Exchange Offer](the-burst-api-examples-publish-exchange-offer.md) example.

### Search Currencies

Get currencies having a code that matches a given query in reverse relevance order.

**Request:**

-   *requestType* is *searchCurrencies*
-   *query* is a full text query on the currency field *code* in the [standard Lucene syntax](http://lucene.apache.org/core/2_9_4/queryparsersyntax.html#Overview)
-   *firstIndex* is a zero-based index to the first currency to retrieve (optional)
-   *lastIndex* is a zero-based index to the last currency to retrieve (optional)
-   *includeCounts* is *true* if the fields beginning with *numberOf...* are to be included (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *currencies* (A) is an array of currency objects (refer to [Get Currency](the-burst-api-get-currency.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Search Currencies](the-burst-api-examples-search-currencies.md) example.

### Transfer Currency

Transfer currency to a given recipient. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *transferCurrency*
-   *recipient* is the account ID of the transfer recipient
-   *currency* is the currency ID
-   *units* is the amount (in QNT) of the transfer

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Transfer Currency](the-burst-api-examples-transfer-currency.md) example.

Networking Operations
---------------------

### Add Peer

Add a peer to the list of known peers and attempt to connect to it. Password protected like the [Debug Operations](the-burst-api-debug-operations.md). POST only.

**Request:**

-   *requestType* is *addPeer*
-   *peer* is the IP address or domain name of the peer (plus optional port)

**Response:** refer to [Get Peer](the-burst-api-get-peer.md)

-   *isNewlyAdded* is *true* if the peer was not already known, omitted otherwise

**Example:** Refer to [Add Peer](the-burst-api-examples-add-peer.md) example.

### Blacklist API Proxy Peer

Blacklist a remote node from the UI, so it won't be used when in roaming and light client modes. POST only.

**Request:**

-   *requestType* is *blacklistAPIProxyPeer*
-   *peer* is the IP address or domain name of the peer (plus optional port)
-   *adminPassword* is a string with the admin password (optional)

**Response:**

-   *done* (B) is *true* if the peer is blacklisted
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Blacklist API Proxy Peer](the-burst-api-examples-blacklist-api-proxy-peer.md) example.

### Blacklist Peer

Blacklist a peer for the default blacklisting period. Password protected like the [Debug Operations](the-burst-api-debug-operations.md). POST only.

**Request:**

-   *requestType* is *blacklistPeer*
-   *peer* is the IP address or domain name of the peer (plus optional port)

**Response:**

-   *done* (B) is *true* if the peer is blacklisted
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Blacklist Peer](the-burst-api-examples-blacklist-peer.md) example.

### Get Inbound Peers

Get all peers that have sent a request to this peer in the last 30 minutes.

**Request:**

-   *requestType* is *getInboundPeers*
-   *includePeerInfo* is *true* to include peer information, otherwise include only the address (optional)

**Response:**

-   *peers* (A) is an array of peer addresses or peer objects (refer to [Get Peer](the-burst-api-get-peer.md) for details) depending on *includePeerInfo*
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Inbound Peers](the-burst-api-examples-get-inbound-peers.md) example.

### Get My Info

Get hostname and address of the requesting node.

**Request:**

-   *requestType* is *getMyInfo*

**Response:**

-   *host* (S) is the node hostname
-   *address* (S) is the node address
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get My Info](the-burst-api-examples-get-my-info.md) example.

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

**Example:** Refer to [Get Peer](the-burst-api-examples-get-peer.md) example.

### Get Peers

Get a list of peer IP addresses.

**Request:**

-   *requestType* is *getPeers*
-   *active* is *true* for active (not NON\_CONNECTED) peers only (optional, if *true* overrides *state*)
-   *state* is the state of the peers, one of *NON\_CONNECTED*, *CONNECTED*, or *DISCONNECTED* (optional)
-   *includePeerInfo* is *true* to include peer detail as in [Get Peer](the-burst-api-get-peer.md)
-   *service* to filter on a specific service

**Note:** If neither *active* nor *state* is specified, all known peers are retrieved.

**Response:**

-   *peers* (A) is an array of peer addresses
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Peers](the-burst-api-examples-get-peers.md) example.

### Set API Proxy Peer

Set the remote node to use when in roaming and light client modes. POST only.

**Request:**

-   *requestType* is *setAPIProxyPeer*
-   *peer* is the IP address or domain name of the peer (plus optional port)
-   *adminPassword* is a string with the admin password (optional)

**Response:**

-   *downloadedVolume* (N) is the number of bytes downloaded by the peer
-   *address* (S) the IP address or DNS name of the peer
-   *weight* (N) is the peer's weight value
-   *uploadedVolume* (N) is the number of bytes uploaded by the peer
-   *version* (S) is the version of the software running on the peer
-   *platform* (S) is a string representing the peer's platform
-   *blockchainState* (S) is a string describing the state of the blockchain in the peer
-   *lastUpdated* (N) is the timestamp (in seconds since the genesis block) of the last peer status update
-   *blacklisted* (B) is *true* if the peer is blacklisted
-   *services* (A) is an array of strings with the services the node provides
-   *apiPort* (N) is the API access port of the peer
-   *apiSSLPort* (N) is the SSL API access port of the peer
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

**Example:** Refer to [Set API Proxy Peer](the-burst-api-examples-set-api-proxy-peer.md) example.

Phasing Operations
------------------

### Approve Transaction

Approve (vote for) a phased transaction. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *approveTransaction*
-   *transactionFullHash* is the full hash of the transaction to be approved (may be used up to 10 times per API request)
-   *revealedSecret* is the secret phrase (required only for *phasingVotingModel 5*, refer to [Create Phasing Poll](the-burst-api-create-phasing-poll.md))
-   *revealedSecretIsText* is a way of specifying whether revealedSecret is text or binary.

**Note:** This transaction will be accepted in the blockchain only if all phased transactions it is voting for are already in it.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Approve Transaction](the-burst-api-examples-approve-transaction.md) example.

### Create Phasing Poll

Create a phased transaction with conditional deferred execution based on the result of a vote, on a list of linked transactions or on the revelation of a secret; or simply with unconditional deferred execution. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is any type from the [Create Transaction](the-burst-api-create-transaction.md) list which is phasing-safe, indicated with italics such as *Send Money*
-   *phased* is *true* to create a phased transaction (optional but required for all of the following parameters, which are all optional for unphased transactions)
-   *phasingFinishHeight* is the block height at which the poll will finish; the transaction will be executed at this block height only if all conditions (if any) have been met, otherwise the transaction will never be executed
-   *phasingVotingModel* is an integer code for the method of approval:
    -   *-1* for *None*
    -   *0* for *Vote By Account*
    -   *1* for *Vote By Account Balance*
    -   *2* for *Vote By Asset Balance*
    -   *3* for *Vote By Currency Balance*
    -   *4* for *By Linked Transactions*
    -   *5* for *By Secret*
-   *phasingQuorum* is the number of “votes” needed for transaction approval (required if *phasingVotingModel* &gt;= *0*, default *0*):
    -   *0* for voting model *-1*
    -   the number of accounts for model *0*
    -   total NQT for model *1*
    -   total QNT for models *2* and *3*
    -   the number of transactions for model *4*
    -   *1* for model *5*
-   *phasingMinBalance* is the minimum balance (in NQT or QNT) required for voting (optional, default *0*)
-   *phasingMinBalanceModel* is (required if *phasingMinBalance* &gt; *0*, must match *phasingVotingModel* when *phasingVotingModel* = *1*, *2* or *3*):
    -   *1* for NXT balance
    -   *2* for an asset balance
    -   *3* for a currency balance
-   *phasingHolding* is the asset or currency ID (required if *phasingMinBalanceModel* = *2* or *3*)
-   *phasingWhitelisted* is the account ID of an account allowed to vote for the transaction; once used, *only* whitelisted accounts are allowed to vote (optional, may be used up to ten times per API request)
-   *phasingLinkedFullHash* is the full hash of a transaction that must be in the blockchain at the *phasingFinishHeight*; transactions already in the blockchain before the acceptance of the phased transaction can also be linked, as long as they are not more than 60 days old, or themselves phased transactions (required only for voting model *4*, may be used up to ten times per API request)
-   *phasingHashedSecret* is the hash of a secret phrase (up to 100 bytes long) required for approval (required only for voting model *5*)
-   *phasingHashedSecretAlgorithm* is the hash function used: *2* for SHA256, *6* for RIPEMD160 and *62* for SHA256 followed by RIPEMD160, according to [Get Constants](the-burst-api-get-constants.md) (required for a *phasingHashedSecret*)

**Note:** When a balance affects the poll result, the result depends only on the balance (including pending outgoing phased transfers) computed just prior to the finish height.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Create Phasing Poll](the-burst-api-examples-create-phasing-poll.md) example.

### Get Account Phased Transaction Count

Get the number of pending phased transactions associated with an account given the account ID.

**Request:**

-   *requestType* is *getAccountPhasedTransactionCount*
-   *account* is the account ID
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *numberOfPhasedTransactions* (N) is the number of pending phased transactions
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Phased Transaction Count](the-burst-api-examples-get-account-phased-transaction-count.md) example.

### Get Account Phased Transactions

Get pending phased transactions associated with an account given the account ID in reverse chronological creation order.

**Request:**

-   *requestType* is *getAccountPhasedTransactions*
-   *account* is the account ID
-   *firstIndex* is a zero-based index to the first phased transaction to retrieve (optional)
-   *lastIndex* is a zero-based index to the last phased transaction to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:** Refer to [Get Transaction](the-burst-api-get-transaction.md) for details.

**Example:** Refer to [Get Account Phased Transactions](the-burst-api-examples-get-account-phased-transactions.md) example.

### Get Asset Phased Transactions

Get pending phased transactions based on an asset in reverse chronological creation order. These transactions can be considered transaction approval requests.

**Request:**

-   *requestType* is *getAssetPhasedTransactions*
-   *asset* is the asset ID
-   *account* is an account ID of the account that created the phased transactions (optional)
-   *withoutWhitelist* is *true* to omit phased transactions that include a whitelist (optional)
-   *firstIndex* is a zero-based index to the first phased transaction to retrieve (optional)
-   *lastIndex* is a zero-based index to the last phased transaction to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:** Refer to [Get Transaction](the-burst-api-get-transaction.md) for details.

**Example:** Refer to [Get Asset Phased Transactions](the-burst-api-examples-get-asset-phased-transactions.md) example.

### Get Currency Phased Transactions

Get pending phased transactions based on a currency in reverse chronological creation order. These transactions can be considered transaction approval requests.

**Request:**

-   *requestType* is *getCurrencyPhasedTransactions*
-   *currency* is the currency ID
-   *account* is an account ID of the account that created the phased transactions (optional)
-   *withoutWhitelist* is *true* to omit phased transactions that include a whitelist (optional)
-   *firstIndex* is a zero-based index to the first phased transaction to retrieve (optional)
-   *lastIndex* is a zero-based index to the last phased transaction to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:** Refer to [Get Transaction](the-burst-api-get-transaction.md) for details.

**Example:** Refer to [Get Currency Phased Transactions](the-burst-api-examples-get-currency-phased-transactions.md) example.

### Get Linked Phased Transactions

Gets the phased transactions with by-transaction voting model for a given linkedFullHash, regardless of their phasing status (pending, approved or rejected). Since the corresponding table is trimmed after finish height however, the result will not include those transactions that finished before the last trimming height.

**Request:**

-   *requestType* is *getLinkedPhasedTransactions*
-   *linkedFullHash* is the full hash of the transaction transaction
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *transactions* (A) is an array of transactions (refer to [Get Transaction](the-burst-api-get-transaction.md) for details)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Linked Phased Transactions](the-burst-api-examples-get-linked-phased-transactions.md) example.

### Get Phasing Poll

Get the details of a phasing poll.

**Request:**

-   *requestType* is *getPhasingPoll*
-   *transaction* is the transaction ID of the phasing poll
-   *countVotes* is *true* to compute the poll *result* while the votes are still available (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *transaction* (S) is the transaction ID of the phasing poll
-   *account* (S) is the number of the account that created the phasing poll
-   *accountRS* (S) is the Reed-Solomon address of the account that created the phasing poll
-   *finishHeight* (N) is the block height at which the poll finished or will finish
-   *votingModel* (N) is the voting model (refer to [Create Transaction Request](the-burst-api-create-transaction-request.md))
-   *quorum* (S) is the minimum number of votes needed to approve the poll
-   *transactionFullHash* (S) is the full hash of the phasing poll transaction
-   *finished* (B) is *true* if the poll is finished, *false* otherwise (omitted if *finished* is *false*)
-   *result* (S) is the sum of the *result* of each account that approved (voted for) the transaction; an account's *result* is *1* if the voting model is *0*, *4* or *5*; it is the NQT, asset QNT or currency QNT balance of the account if the voting model is *1*, *2* or *3* respectively; however, the *result* is *0* if *minBalance* is not met
-   *approved* (B) is *true* if the poll was approved, *false* otherwise
-   *minBalance* (S) is the required minimum balance of voting accounts to be eligible to vote
-   *minBalanceModel* (N) is the minimum balance model (refer to [Create Transaction Request](the-burst-api-create-transaction-request.md))
-   *hashedSecret* (S) is the hash of a secret that must be included in each approval (vote) transaction for the approval to be accepted (refer to [Create Transaction Request](the-burst-api-create-transaction-request.md))
-   *linkedFullHashes* (A) is an array of full hashes of linked transactions (omitted if *votingModel* != *4*)
-   *whitelist* (A) is an array of whitelist objects containing the following two fields (omitted if *votingModel* != *5*):
    -   *whitelisted* (S) is the number of the whitelisted account
    -   *whitelistedRS* (S) is the Reed-Solomon address of the whitelisted account
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Phasing Poll](the-burst-api-examples-get-phasing-poll.md) example.

### Get Phasing Poll Vote

Get a cast phasing poll vote given a phased transaction ID and an account ID of a voter, if it is still available.

**Request:**

-   *requestType* is *getPhasingPollVote*
-   *transaction* is the phased transaction ID
-   *account* is the account ID of a voter in the poll
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *voter* (S) is the account ID of the voter in the poll
-   *voterRS* (S) is the Reed-Solomon address of the *voter*
-   *transaction* (S) is the phased transaction ID
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Phasing Poll Vote](the-burst-api-examples-get-phasing-poll-vote.md) example.

### Get Phasing Poll Votes

Get all cast phasing poll votes in a phasing poll given a phased transaction ID, if they are still available.

**Request:**

-   *requestType* is *getPhasingPollVotes*
-   *transaction* is the phased transaction ID
-   *firstIndex* is a zero-based index to the first vote to retrieve (optional)
-   *lastIndex* is a zero-based index to the last vote to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:** Refer to [Get Phasing Poll Vote](the-burst-api-get-phasing-poll-vote.md)

**Example:** Refer to [Get Phasing Poll Votes](the-burst-api-examples-get-phasing-poll-votes.md) example.

### Get Phasing Polls

Get phasing poll details given multiple phased transaction IDs.

**Request:**

-   *requestType* is *getPhasingPolls*
-   *transaction* is one of the multiple phased transaction IDs
-   *transaction* is one of the multiple phased transaction IDs

⋮

-   *countVotes* is *true* to compute the poll *result* while the votes are still available (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:** Refer to [Get Phasing Poll](the-burst-api-get-phasing-poll.md).

**Example:** Refer to [Get Phasing Polls](the-burst-api-examples-get-phasing-polls.md) example.

### Get Voter Phased Transactions

Get pending phased transactions which include a whitelist in reverse chronological creation order. These transactions can be considered transaction approval requests.

**Request:**

-   *requestType* is *getVoterPhasedTransactions*
-   *account* is a whitelisted account ID included in the phased transactions
-   *firstIndex* is a zero-based index to the first phased transaction to retrieve (optional)
-   *lastIndex* is a zero-based index to the last phased transaction to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:** Refer to [Get Transaction](the-burst-api-get-transaction.md) for details.

**Example:** Refer to [Get Voter Phased Transactions](the-burst-api-examples-get-voter-phased-transactions.md) example.

Server Information Operations
-----------------------------

### Event Register

Create, modify or remove an event listener which can report server events via [Event Wait](the-burst-api-event-wait.md). POST only.

**Request:**

-   *requestType* is *eventRegister*
-   *event* is one of multiple server events from the following list of event names: (optional, default is all possible events)
    -   Block.BLOCK\_GENERATED
    -   Block.BLOCK\_POPPED
    -   Block.BLOCK\_PUSHED
    -   Peer.ADD\_INBOUND
    -   Peer.ADDED\_ACTIVE\_PEER
    -   Peer.BLACKLIST
    -   Peer.CHANGED\_ACTIVE\_PEER
    -   Peer.DEACTIVATE
    -   Peer.NEW\_PEER
    -   Peer.REMOVE
    -   Peer.REMOVE\_INBOUND
    -   Peer.UNBLACKLIST
    -   Transaction.ADDED\_CONFIRMED\_TRANSACTIONS
    -   Transaction.ADDED\_UNCONFIRMED\_TRANSACTIONS
    -   Transaction.REJECT\_PHASED\_TRANSACTION
    -   Transaction.RELEASE\_PHASED\_TRANSACTION
    -   Transaction.REMOVE\_UNCONFIRMED\_TRANSACTIONS
-   *event* is one of multiple server events (optional)

⋮

-   *add* is *true* to add events to an existing listener (optional, omit if *remove* is *true*)
-   *remove* is *true* to remove events from an existing listener (optional, omit if *add* is *true*)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** To create a new event listener, omit both *add* and *remove*. To remove an existing event listener, set *remove* to *true* and omit *event*; all registered events will be removed, any outstanding [Event Wait](the-burst-api-event-wait.md) will be completed and the listener will be deactivated.

**Note:** An event listener is automatically deactivated whenever all registered events are removed or if [Event Wait](the-burst-api-event-wait.md) is not called frequently enough, according to the *nxt.apiEventTimeout* property. The timeout is not precise; the removal process runs every *nxt.apiEventTimeout* / 2 seconds, so that many extra seconds may elapse before removal; the first [Event Wait](the-burst-api-event-wait.md) call should be made immediately after registration to avoid deactivation.

**Note:** Each API user (with a unique address) can create only one event listener. When a new one is created, it will replace an existing one. The maximum number of unique users is controlled by the *nxt.maxEventUsers* property.

**Response:**

-   *registered* is *true* if the operation completed successfully
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Event Register](the-burst-api-examples-event-register.md) example.

### Event Wait

Wait for events registered with [Event Register](the-burst-api-event-register.md). POST only.

**Request:**

-   *requestType* is *eventWait*
-   *timeout* is the amount of time (in seconds) to wait for an event before the call returns (optional, default and maximum is the nxt.apiEventTimeout property)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Notes:** The call returns immediately if one or more events have occurred since the last call; multiple events are all returned together. If a new call is made before the last one returns, the *timeout* timer resets to the new value. Event registration expires if wait calls are not made frequently enough, according to the *nxt.apiEventTimeout* property.

**Response:**

-   *events* (A) is an array of event objects each of which has the following fields:
    -   *name* (S) is the name of the event (refer to [Event Register](the-burst-api-event-register.md) for the list of event names)
    -   *ids* (A) is an array of identifiers, depending on the type of event:
        -   block string identifier (S) for a block event
        -   peer network address (S) for a peer event
        -   transaction string identifier (S) for a transaction event
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Event Wait](the-burst-api-examples-event-wait.md) example.

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

**Example:** Refer to [Get Blockchain Status](the-burst-api-examples-get-blockchain-status.md) example.

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

**Example:** Refer to [Get Constants](the-burst-api-examples-get-constants.md) example.

### Get Plugins

Get a list of all installed plugins on the server.

**Request:**

-   *requestType* is *getPlugins*

**Response:**

-   *plugins* (A) is an array of plugin names (S)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Plugins](the-burst-api-examples-get-plugins.md) example.

### Get State

Get the state of the server node and network.

**Request:**

-   *requestType* is *getState*
-   *includeCounts* is *true* if the fields beginning with *numberOf...* are to be included (optional); password protected like the [Debug Operations](the-burst-api-debug-operations.md) if *true*.

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

**Example:** Refer to [Get State](the-burst-api-examples-get-state.md) example.

### Get Time

Get the current time.

**Request:**

-   *requestType* is *getTime*

**Response:**

-   *time* (N) is the current time (in seconds since the genesis block).
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Time](the-burst-api-examples-get-time.md) example.

Shuffling Operations
--------------------

Coin shuffling can be used to perform mixing of NXT, MS currencies (unless created as non-shuffleable), or AE assets. Any account can create a new shuffling, specifying the holding to be shuffled, the shuffle amount, number of participants required, and registration deadline.

### Get Account Shufflings

Retrieves info about shufflings for a specific account.

**Request:**

-   *requestType* is *getAccountShufflings*
-   *account* is the account ID
-   *includeFinished* is *true* to include completed shufflings (optional)
-   *includeHoldingInfo* is *true* to include holding info (optional)
-   *firstIndex* is a zero-based index to the first tagged data to retrieve (optional)
-   *lastIndex* is a zero-based index to the last tagged data to retrieve (optional)
-   *adminPassword* is a string with the admin password (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *shufflings* (A) is an array containing the shuffling object (refer to [Get Shuffling](the-burst-api-get-shuffling.md))

**Example:** Refer to [Get Account Shufflings](the-burst-api-examples-get-account-shufflings.md) example.

### Get All Shufflings

Retrieves info about all shufflings.

**Request:**

-   *requestType* is *getAllShufflings*
-   *includeFinished* is *true* to include completed shufflings (optional)
-   *includeHoldingInfo* is *true* to include holding info (optional)
-   *finishedOnly* is *true* to omit not yet finished shufflings (optional)
-   *firstIndex* is a zero-based index to the first tagged data to retrieve (optional)
-   *lastIndex* is a zero-based index to the last tagged data to retrieve (optional)
-   *adminPassword* is a string with the admin password (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *shufflings* (A) is an array containing the shuffling object (refer to [Get Shuffling](the-burst-api-get-shuffling.md))

**Example:** Refer to [Get All Shufflings](the-burst-api-examples-get-all-shufflings.md) example.

### Get Assigned Shufflings

Retrieves info about a shufflings that are currently assigned to a specific account.

**Request:**

-   *requestType* is *getAssignedShufflings*
-   *account* is the account ID
-   *includeHoldingInfo* is *true* to include holding info (optional)
-   *firstIndex* is a zero-based index to the first tagged data to retrieve (optional)
-   *lastIndex* is a zero-based index to the last tagged data to retrieve (optional)
-   *adminPassword* is a string with the admin password (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *shufflings* (A) is an array containing the shuffling object (refer to [Get Shuffling](the-burst-api-get-shuffling.md))

**Example:** Refer to [Get Assigned Shufflings](the-burst-api-examples-get-assigned-shufflings.md) example.

### Get Holding Shufflings

Retrieves info about shufflings for a specific holding and/or stage.

**Request:**

-   *requestType* is *getHoldingShufflings*
-   *holding* is the holding ID (optional)
-   *stage* is the stage of the shuffling (See [Get Constants](the-burst-api-get-constants.md) for type definitions) (optional)
-   *includeFinished* is *true* to include completed shufflings (optional)
-   *firstIndex* is a zero-based index to the first tagged data to retrieve (optional)
-   *lastIndex* is a zero-based index to the last tagged data to retrieve (optional)
-   *adminPassword* is a string with the admin password (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *shufflings* (A) is an array containing the shuffling object (refer to [Get Shuffling](the-burst-api-get-shuffling.md))

**Example:** Refer to [Get Holding Shufflings](the-burst-api-examples-get-holding-shufflings.md) example.

### Get Shufflers

Retrieves info about active shufflers on the current node.

**Request:**

-   *requestType* is *getShufflers*
-   *account* is account ID (optional)
-   *shufflingFullHash* is shuffling full hash (optional)
-   *secretPhrase* is secret phrase of the account doing the shuffling (required if adminPassword is not provided)
-   *adminPassword* is the admin password (required if secretPhrase is not provided)
-   *includeParticipantState* to include each shuffling participant's state (optional)

**Response:**

-   *shufflers* (A) is an array containing all currently running shuffling processes on the node.
    -   *account* (S) is account ID
    -   *accountRS* (S) is the account Reed Solomon address
    -   *recipient* (S) is the recipient account ID to where the funds will be sent once the shuffling is completed
    -   *recipientRS* (S) is the recipient account Reed Solomon address to where the funds will be sent once the shuffling is completed
    -   *shuffling* (S) is the shuffling ID
    -   *shufflingFullHash* (S) is the shuffling full hash
    -   *participantState* (N) is the state for the participant (For more info, see shufflingParticipantStates array in [Get Constants](the-burst-api-get-constants.md))
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Shufflers](the-burst-api-examples-get-shufflers.md) example.

### Get Shuffling

Retrieves info about a shuffling.

**Request:**

-   *requestType* is *getShuffling*
-   *shuffling* is the shuffling ID
-   *includeHoldingInfo* is *true* to include holding info (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *blocksRemaining* (N) number of blocks remaining until current stage is finished.
-   *amount* (S) the amount of *holdingType* to be shuffled
-   *shufflingStateHash* (S) state hash of the shuffling
-   *shufflingFullHash* (S) the full hash of the shuffling
-   *issuer* (S) is the issuer account ID
-   *issuerRS* (S) is the Reed-Solomon address of the issuer account
-   *assignee* (S) is the current assignee account ID
-   *assigneeRS* (S) is the Reed-Solomon address of the current assignee account
-   *shuffling* (S) is the shuffling ID
-   *holding* (S) is the asset or currency ID
-   *holdingType* (N) is the holding type (See [Get Constants](the-burst-api-get-constants.md) for type definitions)
-   *stage* (N) is the current stage of the shuffling (See [Get Constants](the-burst-api-get-constants.md) for type definitions)
-   *participantCount* (N) is the number of participants required to start the shuffling
-   *registrantCount* (N) is the number of registrered participants
-   *recipientPublicKeys* (A) is an array of recipient public keys
-   *holdingInfo* (A) is an array containing the asset or currency info (if *includeHoldingInfo* is *true* and holdingType is not NXT)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Shuffling](the-burst-api-examples-get-shuffling.md) example.

### Get Shuffling Participants

Retrieves info about participants in a shuffling.

**Request:**

-   *requestType* is *getShufflingParticipants*
-   *shuffling* is the shuffling ID
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *participants* (A)
    -   *shuffling* (S) is the shuffling ID
    -   *account* (S) is the account ID
    -   *accountRS* (S) is the account Reed Solomong address
    -   *state* (N) is the state for the participant (For more info, see shufflingParticipantStates array in [Get Constants](the-burst-api-get-constants.md))
    -   *nextAccount* (S) is the account ID of the next account in turn
    -   *nextAccountRS* (S) is the account Reed Solomon address of the next account in turn
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Shuffling Participants](the-burst-api-examples-get-shuffling-participants.md) example.

### Shuffling Cancel

Cancels a shuffling

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *shufflingCancel*
-   *shuffling* is the shuffling ID
-   *shufflingStateHash* is the state hash of the shuffling
-   *cancellingAccount* is the account ID (optional)

**Response** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

### Shuffling Create

Creates a new shuffling.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *shufflingCreate*
-   *holding* is the holding id (optional if holdingType is 0)
-   *holdingType* is the holding type (See [Get Constants](the-burst-api-get-constants.md) for type definitions) (optional)
-   *amount* is the amount of the holding to shuffle
-   *participantCount* is the number of participants
-   *registrationPeriod* is the number of blocks the participants have to register until the shuffle is cancelled

**Response** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Shuffling Create](the-burst-api-examples-shuffling-create.md) example.

### Shuffling Process

Manually process the shuffling for a specific participant. Note that the shuffling must be in processing stage and the *secretPhrase* must match the current shuffling assignee.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *shufflingProcess*
-   *shuffling* is the shuffling ID
-   *recipientSecretPhrase* is the secret phrase of the recipient account (optional if *recipientPublicKey* is provided)
-   *recipientPublicKey* is the public key of the recipient account (optional if *recipientSecretPhrase* is provided)

**Response** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Shuffling Process](the-burst-api-examples-shuffling-process.md) example.

### Shuffling Register

Registers a new participant to an existing shuffling. The shuffling must be in stage registration.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *shufflingRegister*
-   *shufflingFullHash* is the full hash of the shuffling to register

**Response** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Shuffling Register](the-burst-api-examples-shuffling-register.md) example.

### Shuffling Verify

Sends a verification that an account's recipient public key is found within a shuffling.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *shufflingVerify*
-   *shuffling* is the shuffling ID
-   *shufflingStateHash* is the current state hash of the shuffle

**Response** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Shuffling Verify](the-burst-api-examples-shuffling-verify.md) example.

### Start Shuffler

Starts a automated Shuffler. Once started, the Shuffler monitors the blockchain state for transactions relevant to the specified shuffle, and automatically submits the required transactions on behalf of the user, performing shuffle processing, verification, or cancellation as needed.

**Request:**

-   *requestType* is *startShuffler*
-   *secretPhrase* is the secret phrase of the account entering the shuffling
-   *shufflingFullHash* the full hash of the shuffling
-   *recipientSecretPhrase* the secret phrase of the recipient account (optional if *recipientPublicKey* is present)
-   *recipientPublicKey* the public key of the recipient account (optional if *recipientSecretPhrase* is present)

**Response**

-   *shuffling* (S) is the shuffling ID
-   *shufflingFullHash* (S) is the full hash of the shuffling
-   *account* (S) is the account ID
-   *accountRS* (S) is the account Reed Solomong address
-   *recipient* (S) is the account ID of the recipient account
-   *recipientRS* (S) is the account Reed Solomon address of the recipient account
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Start Shuffler](the-burst-api-examples-start-shuffler.md) example.

### Stop Shuffler

Stops a previously started automated Shuffler.

**Request:**

-   *requestType* is *stopShuffler*
-   *account* is the account ID (optional if *shufflingFullHash* or *secretPhrase* or *adminPassword* is provided)
-   *shufflingFullHash* the full hash of the shuffling (optional if *account* or *adminPassword* is provided)
-   *secretPhrase* is the secret phrase of the account entering the shuffling (optional if *adminPassword* is provided)
-   *adminPassword* is the admin password (optional if *secretPhrase* is provided)

**Response**

-   *stoppedShuffler* (B) means the specified shuffler was stopped
-   *stoppedAllShufflers* (B) means all shufflers on the node was stopped (only if *adminPassword* is provided in request)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Stop Shuffler](the-burst-api-examples-stop-shuffler.md) example.

Tagged Data Operations
----------------------

Tagged data are similar to [prunable](the-burst-api-prunable-data.md) plain messages without a recipient, but with additional searchable metadata fields.

### Download Tagged Data

Download tagged data as a file if it is still available.

**Request:**

-   *requestType* is *downloadTaggedData*
-   *transaction* is the transaction ID of the tagged data
-   *retrieve* is *true* to retrieve pruned data from other nodes if not available (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:** The downloaded data file named *nxt*, unless there is an error in the request.

**Example:** Refer to [Download Tagged Data](the-burst-api-examples-download-tagged-data.md) example.

### Extend Tagged Data

Extend the expiration time of already uploaded tagged data. POST only.

**Request:**

-   *requestType* is *extendTaggedData*
-   *transaction* is the transaction ID of the tagged data
-   *data* is the tagged data (optional)
-   *file* is the pathname of a data file to upload (optional if *data* provided)
-   *filename*
-   *name*
-   *description*
-   *tags*
-   *type*
-   *channel*
-   *isText*

**Note:** The *data* and metadata (*filename*, *name*, *description*, *tags*, *type*, *channel* and *isText*) parameters can be omitted if the tagged data has not yet expired; otherwise, the tagged data can be restored to the blockchain only if these fields have exactly the same values as when the data was uploaded (refer to [Upload Tagged Data](the-burst-api-upload-tagged-data.md)).

**Note:** Anyone can submit an extension, not only the original uploader. Each extend transaction increases the expiration deadline by two weeks (24 hours on Testnet). Extending an existing tagged data from another account does not change the original submitter account ID by which it is indexed and searchable.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Extend Tagged Data](the-burst-api-examples-extend-tagged-data.md) example.

### Get Account Tagged Data

Get all available tagged data uploaded by a given account in reverse chronological order.

**Request:**

-   *requestType* is *getAccountTaggedData*
-   *account* is the account ID
-   *includeData* is *true* to include *data* (optional)
-   *firstIndex* is a zero-based index to the first tagged data to retrieve (optional)
-   *lastIndex* is a zero-based index to the last tagged data to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *data* (A) is an array of tagged data objects (refer to [Get Tagged Data](the-burst-api-get-tagged-data.md) with *hash* omitted for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Account Tagged Data](the-burst-api-examples-get-account-tagged-data.md) example.

### Get All Tagged Data

Get all available tagged data in reverse chronological order.

**Request:**

-   *requestType* is *getAllTaggedData*
-   *includeData* is *true* to include *data* (optional)
-   *firstIndex* is a zero-based index to the first tagged data to retrieve (optional)
-   *lastIndex* is a zero-based index to the last tagged data to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *data* (A) is an array of tagged data objects (refer to [Get Tagged Data](the-burst-api-get-tagged-data.md) with *hash* omitted for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get All Tagged Data](the-burst-api-examples-get-all-tagged-data.md) example.

### Get Channel Tagged Data

Get available tagged data by channel, optionally filtered by account, in reverse chronological order.

**Request:**

-   *requestType* is *getChannelTaggedData*
-   *channel* is the channel string
-   *account* is an account ID (optional)
-   *includeData* is *true* to include *data* (optional)
-   *firstIndex* is a zero-based index to the first tagged data to retrieve (optional)
-   *lastIndex* is a zero-based index to the last tagged data to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *data* (A) is an array of tagged data objects (refer to [Get Tagged Data](the-burst-api-get-tagged-data.md) with *hash* omitted for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Channel Tagged Data](the-burst-api-examples-get-channel-tagged-data.md) example.

### Get Data Tag Count

Get the total number of distinct available data tags.

**Request:**

-   *requestType* is *getDataTagCount*
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *numberOfDataTags* (N) is the total number of distinct data tags
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Data Tag Count](the-burst-api-examples-get-data-tag-count.md) example.

### Get Data Tags

Get the distinct tags of all available tagged data, with the number of uses of each tag, in order of number of uses, then alphabetical order.

**Request:**

-   *requestType* is *getDataTags*
-   *firstIndex* is a zero-based index to the first tag to retrieve (optional)
-   *lastIndex* is a zero-based index to the last tag to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *tags* (A) is an array of tag objects including the fields:
    -   *tag* (S) is a tag word
    -   *count* (N) is the number of uses of *tag* among all tagged data
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Data Tags](the-burst-api-examples-get-data-tags.md) example.

### Get Data Tags Like

Prefix search of available data tags, return in alphabetical order.

**Request:**

-   *requestType* is *getDataTagsLike*
-   *tagPrefix* is the prefix to search for (2 character minimum) among all data tags
-   *firstIndex* is a zero-based index to the first tag to retrieve (optional)
-   *lastIndex* is a zero-based index to the last tag to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *tags* (A) is an array of tag objects including the fields:
    -   *tag* (S) is a tag word
    -   *count* (N) is the number of uses of *tag* among all tagged data
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Data Tags Like](the-burst-api-examples-get-data-tags-like.md) example.

### Get Tagged Data

Get available tagged data given a transaction ID.

**Request:**

-   *requestType* is *getTaggedData*
-   *transaction* is the transaction ID
-   *includeData* is *true* to include *data* (optional)
-   *retrieve* is *true* to retrieve pruned data from other nodes if not available (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *data* (S) is the tagged data
-   *hash* (S) is the hash of the tagged data
-   *filename* (S) is the metadata *filename* field
-   *name* (S) is the metadata *name* field
-   *description* (S) is the metadata *description* field
-   *tags* (S) is the metadata *tags* field
-   *parsedTags* (A) is an array of tag words (S) parsed from *tags*
-   *type* (S) is the metadata *type* field
-   *channel* (S) is the metadata *channel* field
-   *isText* (B) is the metadata *isText* field
-   *account* (S) is the number of the account that originally uploaded the tagged data
-   *accountRS* (S) is the Reed-Solomon address of the uploading account
-   *transaction* (S) is the transaction ID
-   *transactionTimestamp* (N) is the transaction timestamp (in seconds since the genesis block)
-   *blockTimestamp* (N) is the block timestamp (in seconds since the genesis block)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** Refer to [Upload Tagged Data](the-burst-api-upload-tagged-data.md) for details about the *data* and metadata (*filename*, *name*, *description*, *tags*, *type*, *channel* and *isText*) fields.

**Example:** Refer to [Get Tagged Data](the-burst-api-examples-get-tagged-data.md) example.

### Get Tagged Data Extend Transactions

Retrieves all tagged data extend transactions for a given tagged data upload transaction.

**Request:**

-   *requestType* is *getTaggedDataExtendTransactions*
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *extendTransactions* (A) is an array of transactions
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Tagged Data Extend Transactions](the-burst-api-examples-get-tagged-data-extend-transactions.md) example.

### Search Tagged Data

Full text search on available tagged data name, description and tags; optionally filtered by tag, channel or uploading account; return in reverse relevance order.

**Request:**

-   *requestType* is *searchTaggedData*
-   *query* is a full text query on the metadata fields *name* (S), *description* (S) and *tags* (S) in the [standard Lucene syntax](http://lucene.apache.org/core/2_9_4/queryparsersyntax.html#Overview)
-   *tag* is a word in the *tags* string (optional)
-   *channel* is a channel string (optional)
-   *account* is an account ID (optional)
-   *includeData* is *true* to include *data* (optional)
-   *firstIndex* is a zero-based index to the first tagged data to retrieve (optional)
-   *lastIndex* is a zero-based index to the last tagged data to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *data* (A) is an array of tagged data objects (refer to [Get Tagged Data](the-burst-api-get-tagged-data.md) with *hash* omitted for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Search Tagged Data](the-burst-api-examples-search-tagged-data.md) example.

### Upload Tagged Data

Upload and broadcast new tagged data. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *uploadTaggedData*
-   *data* is the data (optional)
-   *file* is the pathname of a data file to upload (optional if *data* provided)
-   *filename* is a filename to associate with *data* (optional if *file* uploaded in which case the uploaded filename is always used)
-   *name* is the name or title of *data* (optional if *file* uploaded in which case the uploaded filename is used, but *name* takes precedence if provided)
-   *description* is a description of *data* (optional)
-   *tags* is a list of up to 5 words from 3 to 20 characters long and separated by spaces and/or commas, describing the actual content of *data*; for example: *audio,mp3,classical* (optional)
-   *type* is the mime type of *data* such as *torrent*, *pdf*, *doc*, *image*, etc. (optional)
-   *channel* is a data feed label such as *pirate bay torrents*, *wikileaks*, etc. (optional)
-   *isText* is *false* if *data* is a hex string (optional)

**Note:** The maximum length of *data* plus all associated metadata is 42 kilobytes. The maximum length of *description* is 1000 bytes. The maximum length of the other metadata (*name*, *tags*, *type*, *channel* and *filename*) is 100 bytes each.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Upload Tagged Data](the-burst-api-examples-upload-tagged-data.md) example.

### Verify Tagged Data

Verify expired tagged data downloaded from another node, against the hash in the blockchain.

**Request:**

-   *requestType* is *verifyTaggedData*
-   *transaction* is the transaction ID of the tagged data
-   *data* is the tagged data (optional)
-   *file* is the pathname of a data file to upload (optional if *data* provided)
-   *filename*
-   *name*
-   *description*
-   *tags*
-   *type*
-   *channel*
-   *isText*
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** The *data* and metadata (*filename*, *name*, *description*, *tags*, *type*, *channel* and *isText*) must have exactly the same values as when the data was uploaded (refer to [Upload Tagged Data](the-burst-api-upload-tagged-data.md)).

**Response:**

-   *verify* (B) is *true* if the hash of the provided *data* and metadata matches the hash in the blockchain
-   *hash* (S) is the hash of the tagged data
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *version.TaggedDataUpload* (N) is *1*, the version number

**Note:** This call returns an error if there is a hash mismatch.

**Example:** Refer to [Verify Tagged Data](the-burst-api-examples-verify-tagged-data.md) example.

Token Operations
----------------

### Decode File Token

Validate a file token without requiring the transmission of a secret passphrase. POST only.

**Request:**

-   *requestType* is *decodeFileToken*
-   *file* is the path to the file that was signed
-   *token* is the token of the *file*, as generated by [Generate File Token](the-burst-api-generate-file-token.md)

**Response:**

-   *account* (S) is the account number that generated the token
-   *accountRS* (S) is the Reed-Solomon address of the account
-   *timestamp* (N) is the time (in seconds since the genesis block) that the token was generated
-   *valid* (B) is *true* if token is valid, *false* otherwise
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** Since *token* contains the token generator's public key and digital signature, *file* can be validated as signed by the owner of the public key, and the public key determines the account ID.

**Example:** Refer to [Decode File Token](the-burst-api-examples-decode-file-token.md) example.

### Decode Token

Validate a token without requiring the transmission of a secret passphrase.

**Request:**

-   *requestType* is *decodeToken*
-   *website* is the signed text, typically an authorized URL
-   *token* is the token generated by [Generate Token](the-burst-api-generate-token.md)

**Response:**

-   *account* (S) is the account number that generated the token
-   *accountRS* (S) is the Reed-Solomon address of the account
-   *timestamp* (N) is the time (in seconds since the genesis block) that the token was created
-   *valid* (B) is *true* if token is valid, *false* otherwise
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** Since *token* contains the token generator's public key and digital signature, *website* can be validated as authorized by the owner of the public key, and the public key determines the account ID.

**Example:** Refer to [Decode Token](the-burst-api-examples-decode-token.md) example.

### Generate File Token

Generate a file token. POST only.

**Request:**

-   *requestType* is *generateFileToken*
-   *secretPhrase* is the passphrase of the account generating the token
-   *file* is the path to the file to be signed

**Response:**

-   *token* (S) is a 160 character string representing the 100-byte token which consists of a 32-byte public key, a 4-byte timestamp, and a 64-byte digital signature
-   *account* (S) is the account number corresponding to *secretPhrase*
-   *accountRS* (S) is the Reed-Solomon address of the account
-   *timestamp* (N) is the time (in seconds since the genesis block) that the token was generated
-   *valid* (B) is true if token is valid, false otherwise
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** Since *token* contains the token generator's public key and digital signature, the *file* can be validated as digitally signed by the owner of the public key using [Decode File Token](the-burst-api-decode-file-token.md).

**Example:** Refer to [Generate File Token](the-burst-api-examples-generate-file-token.md) example.

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

**Note:** Since *token* contains the token generator's public key and signature, the *website* can be validated as authorized by the owner of the public key using [Decode Token](the-burst-api-decode-token.md).

**Example:** Refer to [Generate Token](the-burst-api-examples-generate-token.md) example.

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

**Example:** Refer to [Broadcast Transaction](the-burst-api-examples-broadcast-transaction.md) example.

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

**Example:** Refer to [Calculate Full Hash](the-burst-api-examples-calculate-full-hash.md) example.

### Get Expected Transactions

Returns the non-phased unconfirmed transactions expected to be included in the next block (only), plus the phased transactions scheduled to finish in that block (whether approved or not).

**Request:**

-   *requestType* is *getExpectedTransactions*
-   *account* is one account ID (optional)
-   *account* is one account ID (optional)

⋮

-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *expectedTransactions* (A) is an array of expected transactions (refer to [Get Transaction](the-burst-api-get-transaction.md) for details)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Expected Transactions](the-burst-api-examples-get-expected-transactions.md) example.

### Get Referencing Transactions

Gets the transactions referencing a given transaction id.

**Request:**

-   *requestType* is *getReferencingTransactions*
-   *transaction* is one transaction ID
-   *firstIndex* is a zero-based index to the first block ID to retrieve (optional)
-   *lastIndex* is a zero-based index to the last block ID to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *transactions* (A) is an array of transactions (refer to [Get Transaction](the-burst-api-get-transaction.md) for details)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Referencing Transactions](the-burst-api-examples-get-referencing-transactions.md) example.

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
-   *subtype* (N) is the transaction subtype (refer to [Get Constants](the-burst-api-get-constants.md) for a current list of subtypes)
-   *block* (S) is the ID of the block containing the transaction
-   *blockTimestamp* (N) is the timestamp (in seconds since the genesis block) of the block
-   *height* (N) is the height of the block in the blockchain
-   *senderPublicKey* (S) is the public key of the sending account for the transaction
-   *type* (N) is the transaction type (refer to [Get Constants](the-burst-api-get-constants.md) for a current list of types)
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

**Example:** Refer to [Get Transaction](the-burst-api-examples-get-transaction.md) example.

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

**Example:** Refer to [Get Transaction Bytes](the-burst-api-examples-get-transaction-bytes.md) example.

### Parse Transaction

Get a transaction object given a (signed or unsigned) transaction bytecode, or re-parse a transaction object. Verify the signature.

**Request:**

-   *requestType* is *parseTransaction*
-   *transactionBytes* is the signed or unsigned bytecode of the transaction (optional)
-   *transactionJSON* is the transaction object (optional if transactionBytes is included)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:** Refer to [Get Transaction](the-burst-api-get-transaction.md) for additional fields.

-   *verify* (B) is *true* if the signature is verified, *false* otherwise

**Example:** Refer to [Parse Transaction](the-burst-api-examples-parse-transaction.md) example.

### Retrieve Pruned Transaction

Force retrieval of the prunable data for a given transaction, even if past the configured nxt.maxPrunableLifetime.

**Request:**

-   *requestType* is *retrievePrunedTransaction*
-   *transaction* is transaction ID

**Response:** Refer to [Get Transaction](the-burst-api-get-transaction.md) for fields.

**Example:** Refer to [Retrieve Pruned Transaction](the-burst-api-examples-retrieve-pruned-transaction.md) example.

### Send Transaction

It broadcasts a transaction to the network without validating it, without re-broadcasting it and without adding it locally as unconfirmed transaction. Specially intended for roaming or light clients to send transactions to remote peers. POST only.

**Request:**

-   *requestType* is *sendTransaction*
-   *transactionBytes* is the bytecode of a signed transaction (optional)
-   *transactionJSON* is the transaction object (optional if *transactionBytes* provided)
-   *prunableAttachmentJSON* is the attachment object embedded in *transactionJSON* containing a prunable message (required if *transactionJSON* not provided because *transactionBytes* never includes prunable data)
-   *adminPassword* is a string with the admin password (optional)

**Response:**

-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *fullHash* (S) is the full hash of the signed transaction
-   *transaction* (S) is the transaction ID

**Example:** Refer to [Send Transaction](the-burst-api-examples-send-transaction.md) example.

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

**Example:** Refer to [Sign Transaction](the-burst-api-examples-sign-transaction.md) example.

Voting System Operations
------------------------

### Cast Vote

Cast a vote on a poll. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *castVote*
-   *poll* is the poll ID
-   *vote00* is an integer within the allowed range to vote for option (answer) 0 (optional if *minNumberOfOptions* met, default is *-128*)
-   *vote01* is an integer within the allowed range to vote for option (answer) 1 (optional if *minNumberOfOptions* met, default is *-128*)
-   *vote02* is an integer within the allowed range to vote for option (answer) 2 (optional if *minNumberOfOptions* met, default is *-128*)

⋮

**Note:** The allowed vote values are integers between *minRangeValue* and *maxRangeValue*, inclusive. This range, along with the minimum and maximum number of options that can and must be voted on are specified when the poll is created. Refer to [Create Poll](the-burst-api-create-poll.md).

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md).

**Example:** Refer to [Cast Vote](the-burst-api-examples-cast-vote.md) example.

### Create Poll

Create a new poll. POST only.

**Request:** Refer to [Create Transaction Request](the-burst-api-create-transaction-request.md) for common parameters.

-   *requestType* is *createPoll*
-   *name* is the name of the poll
-   *description* is the description of the poll, or the question to be answered
-   *finishHeight* is the block height when the poll is completed
-   *votingModel* is *0* for *One Vote Per Account*, *1* for *Vote By NXT Balance*, *2* for *Vote By Asset Balance* and *3* for *Vote By Currency Balance*
-   *minNumberOfOptions* is the minimum number of options (answers) that must be voted for
-   *maxNumberOfOptions* is the maximum number of options (answers) that can be voted for
-   *minRangeValue* is the minimum integer value for an option (answer) (&gt;= *0*)
-   *maxRangeValue* is the maximum integer value for an option (answer) (&gt;= *minRangeValue*)
-   *minBalance* is the minimum balance (in NQT or QNT) required for voting (optional, default 0)
-   *minBalanceModel* is (required if *minBalance* &gt; *0*, must match *votingModel* when *votingModel* &gt; *0*)
    -   *1* for NXT balance
    -   *2* for an asset balance
    -   *3* for a currency balance
-   *holding* is the asset or currency ID (required if *minBalanceModel* &gt; *1*)
-   *option00* is the name of option (answer) 0
-   *option01* is the name of option (answer) 1 (optional)
-   *option02* is the name of option (answer) 2 (optional)

⋮

**Notes:** Up to 100 options (answers) can be specified, but there is an extra fee for each option beyond 20. Unlike the API, the [BRS client](burst-client-interface.md) treats a vote of *0* as a nonvote not contributing to the number answers (options) chosen. The BRS client uses checkboxes for selecting answers when *minRangeValue* = 0 and *maxRangeValue* = 1; otherwise sliding controls are used to select answers from the allowed range.

**Note:** When a balance affects the poll result, the result depends only on the balance (including pending outgoing phased transfers) computed just prior to the finish height.

**Response:** Refer to [Create Transaction Response](the-burst-api-create-transaction-response.md). The transaction ID is also the poll ID.

**Example:** Refer to [Create Poll](the-burst-api-examples-create-poll.md) example.

### Get Poll

Get the details of a poll.

**Request:**

-   *requestType* is *getPoll*
-   *poll* is the poll ID
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *poll* (S) is the poll ID
-   *account* (S) is the account number of the poll creator
-   *accountRS* (S) is the Reed-Solomon address of the account
-   *name* (S) is the name of the poll
-   *description* (S) is the description of the poll, or the question to be answered
-   *finishHeight* (N) is the block height when the poll is completed
-   *finished* (B) is *true* if the poll is completed, *false* otherwise
-   *votingModel* (N) is *0* for *One Vote Per Account*, *1* for *Vote By NXT Balance*, *2* for *Vote By Asset Balance* and *3* for *Vote By Currency Balance*
-   *minNumberOfOptions* (N) is the minimum number of options (answers) that must be voted for
-   *maxNumberOfOptions* (N) is the maximum number of options (answers) that can be voted for
-   *minBalance* (S) is the minimum balance (in NQT or QNT) required for voting
-   *minBalanceModel* (N) is *1* for NXT balance, *2* for an asset balance, *3* for a currency balance when *minBalance* &gt; 0
-   *holding* is the asset or currency ID when minBalanceModel &gt; 1
-   *options* (A) is the array of options (answers)
-   *minRangeValue* (N) is the minimum integer value for an option (answer)
-   *maxRangeValue* (N) is the maximum integer value for an option (answer)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Poll](the-burst-api-examples-get-poll.md) example.

### Get Poll Result

Get the result of a poll.

**Request:**

-   *requestType* is *getPollResult*
-   *poll* is the poll ID
-   *votingModel* (optional, default null)
-   *holding* (optional, default null)
-   *minBalance* (optional, default *0*)
-   *minBalanceModel* (required if *minBalance* &gt; *0*, must match *votingModel* when *votingModel* &gt; *0*)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Note:** The *votingModel*, *holding*, *minBalance* and *minBalanceModel* parameters are optional and default to the original values specified when the poll was created (refer to [Create Poll](the-burst-api-create-poll.md)). The original values can only be overridden while the votes are still stored in the database, until 1441 blocks after the poll is completed. If *votingModel* is specified, *holding*, *minBalance* and *minBalanceModel* or the defaults specified above apply, otherwise they are ignored.

**Response:**

-   *poll* (S) is the poll ID
-   *votingModel* (N) is the votingModel used to calculate *results* (refer to Note above)
-   *minBalanceModel* (N) is the minBalanceModel used to calculate *results* (refer to Note above)
-   *minBalance* (S) is the minBalance used to calculate *results* (refer to Note above)
-   *holding* (S) is the asset or currency ID if the voting model uses an asset or currency balance to determine *weight*, if applicable (refer to Note above)
-   *decimals* (N) is the number decimal places used by the asset or currency, if applicable
-   *finished* (B) is *true* if the poll is complete, *false* otherwise
-   *options* (A) is the array of options (answers) of the poll
-   *results* (A) is an array of result objects with the following fields for each result:
    -   *weight* (S) is the sum of the *weight* of each account that voted for the corresponding option (answer); an account's *weight* is *1* if the voting model is *0*, otherwise it is the NQT, asset QNT or currency QNT balance of the account if the voting model is *1*, *2* or *3* respectively; however, the *weight* is *0* if *minBalance* is not met
    -   *result* (S) is the sum over each account that voted for the corresponding option (answer) of: the product of the account's *weight* and the *rangeValue* selected when the vote was cast.
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Poll Result](the-burst-api-examples-get-poll-result.md) example.

### Get Poll Vote

Get a poll vote given a poll ID and an account ID.

**Request:**

-   *requestType* is *getPollVote*
-   *poll* is the poll ID
-   *account* is the account ID
-   *includeWeights* is *true* to calculate and return the weight assigned to each vote (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *votes* (A) is an array of votes, the range values (S) corresponding to each option (answer) with empty strings for non-votes
-   *voter* (S) is the account number of the voter
-   *voterRS* (S) is the Reed-Solomon address of the voter
-   *transaction* (S) is the transaction ID of the vote
-   *weight* (S) is the weight assigned to each vote (applies if *includeWeights* is *true*)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** Votes are deleted from the database 1441 blocks after the poll is finished. Only aggregate results are kept (refer to [Get Poll Result](the-burst-api-get-poll-result.md)).

**Example:** Refer to [Get Poll Vote](the-burst-api-examples-get-poll-vote.md) example.

### Get Poll Votes

Get all votes on a poll in reverse chronological order.

**Request:**

-   *requestType* is *getPollVotes*
-   *poll* is the poll ID
-   *includeWeights* is *true* to calculate and return the weight assigned to each vote (optional)
-   *firstIndex* is a zero-based index to the first vote to retrieve (optional)
-   *lastIndex* is a zero-based index to the last vote to retrieve (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *votes* (A) is an array of vote objects (refer to [Get Poll Vote](the-burst-api-get-poll-vote.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** Votes are deleted from the database 1441 blocks after the poll is finished. Only aggregate results are kept (refer to [Get Poll Result](the-burst-api-get-poll-result.md)).

**Example:** Refer to [Get Poll Votes](the-burst-api-examples-get-poll-votes.md) example.

### Get Polls

Get poll details in reverse creation order.

**Request:**

-   *requestType* is *getPolls*
-   *account* is a creation account ID filter (optional)
-   *timestamp* is the earliest poll (in seconds since the genesis block) to retrieve (optional)
-   *firstIndex* is a zero-based index to the first poll to retrieve (optional)
-   *lastIndex* is a zero-based index to the last poll to retrieve (optional)
-   *includeFinished* is *true* to include completed polls (optional)
-   *finishedOnly* is *true* to exclude not yet executed, phased transactions (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *polls* (A) is an array of polls (refer to [Get Poll](the-burst-api-get-poll.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Polls](the-burst-api-examples-get-polls.md) example.

### Search Polls

Search for poll details given a name/description query string.

**Request:**

-   *requestType* is *searchPolls*
-   *query* is a full text query on the poll fields *name* (S) and *description* (S) in the [standard Lucene syntax](http://lucene.apache.org/core/2_9_4/queryparsersyntax.html#Overview) (optional)
-   *firstIndex* is a zero-based index to the first poll to retrieve (optional)
-   *lastIndex* is a zero-based index to the last poll to retrieve (optional)
-   *includeFinished* is *true* to include completed polls (optional)
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *polls* (A) is an array of polls (refer to [Get Poll](the-burst-api-get-poll.md) for details)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Search Polls](the-burst-api-examples-search-polls.md) example.

Utilities
---------

### Decode QR Code

Decodes a base64-encoded jpeg to a UTF-8 string. POST only.

**Request:**

-   *requestType* is *decodeQRCode*
-   *qrCodeBase64* is a base64-encoded jpeg string to be decoded

**Response**

-   *qrCodeData* (S) is a UTF-8 string containing the decoded data from the base64 string
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Decode QR Code](the-burst-api-examples-decode-qr-code.md) example.

### Detect Mime Type

Gets the mime type of uploaded file or data.

**Request:**

-   *requestType* is *detectMimeType*
-   *data* is the data (optional)
-   *file* is the pathname of a data file to upload (optional if *data* provided)
-   *filename* is a filename to associate with *data* (optional if *file* uploaded in which case the uploaded filename is always used)
-   *isText* is *false* if *data* is a hex string (optional)

**Response**

-   *type* (S) is the mime type
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Detect Mime Type](the-burst-api-examples-detect-mime-type.md) example.

### Encode QR Code

Encodes a UTF-8 string to a base64-encoded jpeg. POST only.

**Request:**

-   *requestType* is *encodeQRCode*
-   *qrCodeData* is a UTF-8 text string to be encoded
-   *width* is the width of the output image (optional)
-   *height* is the height of the output image (optional)

**Response**

-   *qrCodeBase64* (S) is a base64 string encoding a jpeg image of the QR code
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Encode QR Code](the-burst-api-examples-encode-qr-code.md) example.

### Full Hash To Id

Converts a full hash to an ID.

**Request:**

-   *requestType* is *fullHashToId*
-   *fullHash* is the full hash 64-digit (32-byte) hex string

**Response:**

-   *stringId* (S) is the ID corresponding to the hash, in the form of an decimal string
-   *longId* (S) is the signed long integer (8-bytes) representation of the ID used internally, returned as a string
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Full Hash To Id](the-burst-api-examples-full-hash-to-id.md) example.

### Hash

Calculates the hash of a secret for use in [phased transactions](the-burst-api-create-phasing-poll.md) with voting model 5 (*Vote By Secret*).

**Request:**

-   *requestType* is *hash*
-   *hashAlgorithm* is the hash function used: *2* for SHA256, *3* for SHA3, *5* for SCRYPT, *6* for RIPEMD160, *25* for Keccack25 and *62* for SHA256 followed by RIPEMD160, according to [Get Constants](the-burst-api-get-constants.md)
-   *secret* is a secret phrase in text form or hex string form
-   *secretIsText* is *true* if *secret* is text, *false* if it is a hex string (optional)

**Note:** *secret* is converted from a hex string to a byte array, which is what the hash algorithm expects, unless *secretIsText* is *true*, in which case *secret* is first converted from text to a UTF-8 hex string as by [Hex Convert](the-burst-api-hex-convert.md).

**Response:**

-   *hash* (S) is the hash of the secret, in the form of a hex string
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Hash](the-burst-api-examples-hash.md) example.

### Hex Convert

Converts a text string into a UTF-8 hex string and if the text input is already a hex string, also into text.

**Request:**

-   *requestType* is *hexConvert*
-   *string* is a text string, possibly a hex string

**Response:**

-   *binary* (S) is the converted UTF-8 hex string
-   *text* (S) is a text string converted from *string* if it is a valid UTF-8 hex string
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Hex Convert](the-burst-api-examples-hex-convert.md) example.

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

**Example:** Refer to [Long Convert](the-burst-api-examples-long-convert.md) example.

### RS Convert

Get both the Reed-Solomon account address and the account number given an account ID.

**Request:**

-   *requestType* is *rsConvert*
-   *account* is an account ID (either RS address or number)

**Response:**

-   *accountRS* (S) is the Reed-Solomon address of the account
-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *account* (S) is the account number

**Example:** Refer to [RS Convert](the-burst-api-examples-rs-convert.md) example.

Debug Operations
----------------

All debug utilities require an *adminPassword* request parameter. See [Admin Password](the-burst-api-admin-password.md) for more info.

### Clear Unconfirmed Transactions

Empties the unconfirmed transaction pool. POST only.

**Request:**

-   *requestType* is *clearUnconfirmedTransactions*

**Response:**

-   *done* (B) is *true* if the operation completed successfully
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Clear Unconfirmed Transactions](the-burst-api-examples-clear-unconfirmed-transactions.md) example.

### Dump Peers

Get all active peers, optionally of a certain version or a minimum weight.

**Request:**

-   *requestType* is *dumpPeers*
-   *version* is a version filter such as *1.5.11* (optional)
-   *weight* is a minimum weight filter such as 1000 (optional)
-   *connect* is *true* to force a connection attempt to each known peer first (optional); password protected like the [Debug Operations](the-burst-api-debug-operations.md) if *true*

**Response:**

-   *peers* (S) is a string of peer IP addresses or DNS names, separated by semicolons
-   *count* (N) is the number of peers in the *peers* string.
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Dump Peers](the-burst-api-examples-dump-peers.md) example.

### Full Reset

Deletes the entire blockchain. POST only.

**Request:**

-   *requestType* is *fullReset*

**Response:**

-   *done* (B) is *true* if the operation completed successfully
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Note:** After successful completion of the reset, a new blockchain will automatically begin downloading.

**Example:** Refer to [Full Reset](the-burst-api-examples-full-reset.md) example.

### Get All Broadcasted Transactions

Get unconfirmed transactions broadcasted from this node but not yet received back from a peer, if transaction rebroadcasting is enabled.

**Request:**

-   *requestType* is *GetAllBroadcastedTransactions*
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *transactions* (A) is an array of broadcasted unconfirmed transactions not yet received back from a peer (S)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get All Broadcasted Transactions](the-burst-api-examples-get-all-broadcasted-transactions.md) example.

### Get All Waiting Transactions

Get unconfirmed transactions temporarily kept in memory during transaction processing.

**Request:**

-   *requestType* is *getAllWaitingTransactions*
-   *requireBlock* is the block ID of a block that must be present in the blockchain during execution (optional)
-   *requireLastBlock* is the block ID of a block that must be last in the blockchain during execution (optional)

**Response:**

-   *transactions* (A) is an array of unconfirmed transactions temporarily kept in memory (S)
-   *lastBlock* (S) is the last block ID on the blockchain (applies if *requireBlock* is provided but not *requireLastBlock*)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get All Waiting Transactions](the-burst-api-examples-get-all-waiting-transactions.md) example.

### Get Log

Get up to 100 of the most recent log messages from a memory buffer.

**Request:**

-   *requestType* is *getLog*
-   *count* is the number of messages to return (optional, default 100)

**Response:**

-   *messages* (A) is an array of log messages (S)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Get Log](the-burst-api-examples-get-log.md) example.

### Get Stack Traces

Get the stack traces of the currently running threads in reverse *id* order.

**Request:**

-   *requestType* is *getStackTraces*
-   *depth* is the maximum trace depth to retrieve (optional)

**Response:**

-   *requestProcessingTime* (N) is the API request processing time (in millisec)
-   *locks* (A) is an array of lock objects
-   *threads* (A) is an array of thread objects with the following fields:
    -   *trace* (A) is an array of traces (S)
    -   *name* (S) is the thread name
    -   *id* (N) is the thread ID
    -   *state* (S) is the thread state, one of *WAITING*, *TIMED\_WAITING* and *RUNNABLE*
    -   *locks* (A) is an array of lock objects with the following fields, if not empty:
        -   *trace* (S)
        -   *depth* (N)
        -   *name* (S)
        -   *hash* (N)

**Example:** Refer to [Get Stack Traces](the-burst-api-examples-get-stack-traces.md) example.

### Lucene Reindex

Forces a rebuild of the Lucene search index. POST only.

**Request:**

-   *requestType* is *luceneReindex*

**Response:**

-   *done* (B) is *true* if the operation completed successfully
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Lucene Reindex](the-burst-api-examples-lucene-reindex.md) example.

### Pop Off

Removes specified number of blocks (and associated transactions) from the top of the blockchain. POST only.

**Request:**

-   *requestType* is *popOff*
-   *numBlocks* is the number of blocks to pop off the blockchain (optional)
-   *height* is the new height of the blockchain after popping (optional if *numBlocks* provided)

**Note:** If table trimming is enabled (default), at most 1440 blocks can be popped off without triggering a full rescan.

**Response:**

-   *blocks* (A) is an array of the blocks popped off (refer to [Get Block](the-burst-api-get-block.md) for details)
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Pop Off](the-burst-api-examples-pop-off.md) example.

### Rebroadcast Unconfirmed Transactions

Rebroadcast transactions in the unconfirmed pool to peers, until received back or found in the blockchain. Rebroadcasting can be disabled by setting the *nxt.enableTransactionRebroadcasting* property to *false*. POST only.

**Request:**

-   *requestType* is *RebroadcastUnconfirmedTransactions*

**Response:**

-   *done* (B) is *true* if the operation completed successfully
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Rebroadcast Unconfirmed Transactions](the-burst-api-examples-rebroadcast-unconfirmed-transactions.md) example.

### Requeue Unconfirmed Transactions

Requeue unconfirmed transactions. POST only.

**Request:**

-   *requestType* is *requeueUnconfirmedTransactions*

**Response:**

-   *done* (B) is *true* if the operation completed successfully
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Requeue Unconfirmed Transactions](the-burst-api-examples-requeue-unconfirmed-transactions.md) example.

### Retrieve Pruned Data

Initiates a task of requesting and restoring missing prunable data. POST only.

**Request:**

-   *requestType* is *retrievePrunedData*

**Response:**

-   *done* (B) is *true* if the operation completed successfully
-   *numberOfPrunedData* (N) is the number of pruned data available pruned data transactions
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Retrieve Pruned Data](the-burst-api-examples-retrieve-pruned-data.md) example.

### Scan

Scans the top of the blockchain. POST only.

**Request:**

-   *requestType* is *scan*
-   *numBlocks* is the number of blocks to scan at the top of the blockchain (optional)
-   *height* is the height above which blockchain is to be scanned (optional if *numBlocks* provided)
-   *validate* is *true* if signatures are to be re-verified and blocks and transactions re-validated (optional)

**Note:** The derived object tables are rolled back and rebuilt by rescanning the existing blockchain. A request to rescan more than 1440 blocks when table trimming is enabled will do a full rescan starting from height 0. Rescan status is saved in the database, so that if a rescan is interrupted or fails it will resume on restart.

**Response:**

-   *scanTime* (N) is the scan time
-   *done* (B) is *true* if the operation completed successfully
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Scan](the-burst-api-examples-scan.md) example.

### Set Logging

Sets the log level and optionally specifies communication events to be logged, without restarting the server. POST only.

**Request:**

-   *requestType* is *setLogging*
-   *logLevel* is one of *ERROR*, *WARN*, *INFO* or *DEBUG* with each level in the list including all of the previous levels (optional, default is *INFO*)
-   *communicationEvent* is one of multiple communication (HTTP) events to be logged, from the list: *EXCEPTION*, *HTTP-ERROR*, *HTTP-OK* (optional)
-   *communicationEvent* is one of multiple communication events (optional)

⋮

**Note:** The initial log level is set by the *nxt.level* logging property, currently *FINE* (equivalent to *DEBUG*).

**Response:**

-   *loggingUpdated* (B) is *true* if the operation completed successfully

**Example:** Refer to [Set Logging](the-burst-api-examples-set-logging.md) example.

### Shutdown

Shutdown the server. POST only.

**Request:**

-   *requestType* is *shutdown*
-   *scan* is *true* to truncate the derived tables and schedule a full rescan with validation on the next start (optional)

**Response:**

-   *shutdown* (B) is *true* if the operation completed successfully
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Shutdown](the-burst-api-examples-shutdown.md) example.

### Trim Derived Tables

Trigger a derived tables trim, and a prunable tables pruning. POST only.

**Request:**

-   *requestType* is *trimDerivedTables*

**Response:**

-   *done* (B) is *true* if the operation completed successfully
-   *requestProcessingTime* (N) is the API request processing time (in millisec)

**Example:** Refer to [Trim Derived Tables](the-burst-api-examples-trim-derived-tables.md) example.

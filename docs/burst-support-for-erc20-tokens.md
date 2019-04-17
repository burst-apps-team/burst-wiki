Introduction
------------

ERC20 is a popular Ethereum standard for token management, explained here <https://github.com/ethereum/EIPs/issues/20>

The [BURST asset exchange](asset-exchange.md) already supports most of the ERC20 operations out of the box without the need to issue a smart contract.

This document explains how to implement the ERC20 APIs using the equivalent BURST APIs.

Issuing Tokens
--------------

Each ERC20 token is represented as a BURST asset.

Issue the asset using the [issueAsset](the-burst-api-issue-asset.md) API.

Each asset is identified by a unique asset id set to the transaction id of the issueAsset transaction.

For more details, please take a look at [The Burst API](the-burst-api.md) and [The Burst API Examples](the-burst-api-examples.md)

Implementing the ERC20 APIs
---------------------------

**totalSupply**

To get the total token supply invoke the getAsset API and supply the asset id. The total supply is provided by the `initialQuantityQNT` attribute, the existing supply, after possible share deletes, is specified by the `quantityQNT` attribute. Quantity values are always provided as integer values, the number of decimal positions to apply to these quantities is specified by the “decimals” attribute.

**balanceOf**

To get the account balance of another account with a given address, invoke the [getAssetAccounts](the-burst-api-get-asset-accounts.md) API and provide the BURST address as the `account` parameter and the asset id as the `asset` parameter.

Here is an [example](the-burst-api-examples-get-asset-accounts.md).

The `unconfirmedQuantityQNT` attribute in the response represents the current account balance. The `quantityQNT` attribute represents the quantity available for use at the moment after taking into account balance locked by open asset orders.

**transfer**

To send tokens from your address to another address use the [transferAsset](the-burst-api-transfer-asset.md) API, specify the recipient address, the asset id and the quantity to transfer as well as your account passphrase, the transaction fee and transaction deadline.

The [transferAsset](the-burst-api-transfer-asset.md) transaction is recorded on the BURST blockchain. You can later track token transfers using the [Get Asset Transfers](the-burst-api-get-asset-transfers.md) API

| Offline Transaction Signing |
|-----------------------------|
| **Status**                  |
| **API Version**             |
| **BRS Version**             |

Offline Transaction Signing
---------------------------

**Offline Transaction Signing \[offline device\]**

The term “offline transaction” refers to the practice of keeping the private keys on an offline device (not connected to the internet), and signing on individuals transactions. The signature is then copy-pasted from this device into a connected device, and broadcast into the network. Assuming the offline computer is malware-free, then this practice is virtually risk-free of theft.

**Online Transaction Signing \[local device\]**

In addition to signing your transactions from a offline device, the signing can also be done on an online device but still performed locally. Assuming the computer is malware-free, this is the most convenient option while still keeping your private keys secret. For example the [BRS (Burst Reference Software)](burst-reference-software.md) wallet uses this form of signing for its wallet interface through locally run javascript.

**Online Transaction Signing \[server side\]**

Just don't do it. Although its possible, it would only be considered “safe” to do this using a local host. Especially if you are developing/distributing software; don't present online signing as an option to your clients. You will make them a potential target for malicious actions.

------------------------------------------------------------------------

**Implementing Transaction Signing**

Any transaction needs to be signed before it can be broadcast and accepted into the memory pool. You use one of the [API functions](the-burst-api-create-transaction.md) to request `transactionBytes` from a node. This API call returns a JSON containing the `transactionBytes`. These bytes represent the transaction you want to make.

Important is to use the `publicKey` argument and not the `secretPhrase` for the `transactionBytes` request. And set the `broadcast` argument to false to prevent broadcasting the transaction to the network.

Now to sign the transactionBytes locally. Refer to these sources below to include the signing functions in your code.

**Code sources**

[Javascript / Burst Reference Software](https://github.com/PoC-Consortium/burstcoin/tree/master/html/ui/js/crypto)

[Java / Burst Reference Software](https://github.com/PoC-Consortium/burstcoin/tree/master/src/brs/crypto)

[Swift/Obj-C / BurstKit](https://github.com/aprock/BurstKit/tree/master/BurstKit/Crypto)

[C++ / CloudBurst](https://github.com/CurbShifter/CloudBurstDAPP/tree/master/Source/burst/crypto)

Pseudo code:

``
`function signTX(unsignedTransactionBytes) `
`{`
`    myBytes = unsignedTransactionBytes // keep a copy `
`    signature = crypto.sign(unsignedTransactionBytes, passPhrase) // make the signature`
`    myBytes.copy(96, signature); // copy the signature over the unsignedTransactionBytes with a offset of 96 bytes`
`    return myBytes`
`}`

**Note:** add signature and TransactionBytes length

Compare and verify your implementation with the requestType: [Sign Transaction](the-burst-api-sign-transaction.md).

The transaction is now signed and can now be broadcast (through POST only): [Broadcast the transaction](the-burst-api-broadcast-transaction.md) ----

[BRS API Transaction Operations](the-burst-api-transaction-operations.md)

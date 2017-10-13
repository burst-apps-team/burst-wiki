Table of contents
-----------------

1. Introduction

2. Algorithms and Acronyms

3. Generating a nonce.

4. Plot structure

5. Staggers and filenames.

Introduction
------------

The term plotting is a name for dedicating storage space to be used for calculations in the Burstcoin network. A Plot is a file containing pre-computed hashes that can be used to forge blocks for Burstcoin blockchain. The plots are later used by mining software and can be thought of as the miner’s hash rate.

Algorithms and acronyms
-----------------------

Before we dive deep into how plotting works we need to get familiar with all different terms used in the procedure.

**Shabal**

Shabal is the name of the crypto/hash function used in Burstcoin. Shabal is a rather heavy and slow crypto in relation to many other like i.e. SHA256. Because of this it makes it a good crypto for Proof of capacity coins like Burstcoin. This is because we store the precomputed hashes while it is still fast enough to do smaller live verifications. Burst uses the 256bit version of Shabal also known as Shabal256.

**Hash / Digest**

A hash or digest in this context is a 32Byte (256bit) long result of the Shabal256 Crypto.

**Nonce**

When generating a plot file, you generate something that is called nonces. Each nonce contains 256Kilobyte of data that can be used by miners to calculate Deadlines. Each nonce will have its own individual number. This number can range between 0-18446744073709551615. the number is also used as a seed when creating the nonce. Because of this each nonce has its own unique set of data. One plot file can contain many nonces.

**Scoop**

Each nonce is sorted into 4096 different places of data. These places are called scoop numbers. Each scoop contains 64byte of data which holds 2 hashes. Each of these hashes are xored with a final hash (we get to final hash in generating a nonce chapter).

**Account ID**

When you create your plot file it will be bound to a specific Burst account. The numeric account ID is used when you create your nonces. Because of this all miners have different plot files even if they use the same nonce numbers.

Generating a nonce
------------------

The first step in creating a nonce is to make the first seed. The seed is a 16byte long value containing the account id that we will be generating a nonce for and the nonce number. When this is done we start to feed the Shabal256 function to get our first hash.

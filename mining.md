| Mining     |
|------------|
| **Status** |
||

Table of Contents
-----------------

\_\_TOC\_\_

Introduction to Burst Mining
----------------------------

The process of mining Burst is the foundation for adding blocks to the Burst blockchain and securing transactions within the Burst network. In preparation for Mining, plotting software is used to precompute hashes and store them in plot files. During the mining process, Mining software reads quickly through the plot files to locate data to be submitted to the network in a competitive process whereby each miner provides their best cryptographic solution in an attempt to win the right to forge the related block of transactions.

Burst uses the Shabal256 cryptographic hash function as its main function. Shabal256 is slow and heavy in comparison to other cryptographic hash functions such as the SHA256. This feature makes it ideal for precomputation of hashes, yet it is still fast enough for small live verifications.

Lexicon
-------

**Plot files**

In order to mine Burst, users have to open a Burst account and activate it by creating an outgoing transaction, as described in this article. The next step is creating plot files needed for Burst mining. As plot files are bound to the Burst account ID, it is impossible for users with different account IDs to have identical plot files.

Plot files are made up of hashes. A hash is a result of of computing data using a cryptographic hash function. The length of a hash is 32 Bytes (265 bit). A nonce is a group of 8912 hashes, making its size 256 KiB. Each nonce is referenced by a number. The number of the nonce can take values between 0 and 264-1 (18446744073709551615). Nonces are sorted into 4096 scoops, therefore each scoop contains two hashes.

**Reward recipient assignment**

Before they start mining, users have to decide whether they'll join a mining pool or mine alone (“solo mine”). This is done through a transaction called “Reward recipient assignment” which has to be executed for both types of mining. The reward recipient assignment transaction, when used in context of pool mining, enables the pool to use the deadlines submitted by the miner and sign newly forged blocks with miner's account. This transaction also assigns the block rewards from blocks forged by a specific miner to the pool for distribution, in accordance with pool reward distribution policies. Note that forged blocks are always signed by an account belonging to the miner who submitted the deadline used to forge the block.

<img src="Reward_recipient.png" title="Reward_recipient.png" alt="Reward_recipient.png" width="1100" height="343" />

*Reward recipient assignment form*

After the reward recipient assignment transaction has been issued by the user, it will take 4 blocks before the pool or account for solo miners set as the reward recipient will start to accept and confirm deadlines submitted by the miner.

**Pool vs. solo mining**

The decision whether to join a Burst mining pool or mine as a solo miner should be based upon the total size of one's plot files, which ultimately determine the probability of forging a block and the potential revenue. There is no strict rule when it comes to making this decision, but one should have in mind that with solo mining, the block reward and the fees are earned only when a block is forged, while with pool mining the earnings are distributed among pool miners, who “pool” their mining resources with the intent to increase the chance of forging blocks. There are different types of pool, depending on the payout distribution:

-   0-100 mining revenue distribution implies that all block rewards and fees from mined blocks are distributed among pool miners (0% for the block forger, 100% for the pool). This type of pool is best suited for plots smaller than 40TB,

<!-- -->

-   20-80 pools distribute 20% of the revenue to the block forger and the remaining 80% to the pool. These pools are best suited for plot size from 30 to 80TB,

<!-- -->

-   50-50 mining revenue distribution is suitable for plots of 60 to 200TB in size. In 50-50 pools, block forgers receive half of the block reward and fees for forged blocks, while the other half is passed on to the pool and distributed among pool miners,

<!-- -->

-   80-20 pools are suited for plots sizes between 150 and 250TB, as they assign 80% of the block reward and fees to the block forger, and the remaining 20% to the pool,

<!-- -->

-   100-0 pools, which are basically solo mining pools are suited for plot sizes of 150TB or above, and all mining revenue is kept for the block forger.

The part of the mining revenue assigned to the pool for distribution to all pool miners is often referred to as “historical share”.

Note that the plot sizes indicated above are merely an informative guide, as there's no technical obstacle that would prevent anyone from e.g. solo mining with a 6TB plot, or joining a 0-100 pool with a plot of 200TB. Mining pools may or may not charge fees from miners. It is up to pool administrators to set fees, if any, as well as to determine the minimum payout and payout delay. Miners should get familiar with pool policies before they decide which pool to join.

**Effective plot size**

In pools that use the PoCC pool mining software, the effective plot size is the parameter used to determine the share of each miner in the pool mining revenue. The effective plot size is calculated based on the best deadlines confirmed by the pool for the last 360 blocks that were submitted by the miner, and it usually oscillates around the actual, physical size of the plot. The effective plot size can be optimized by adjusting the target deadlines sent to the pool. More information on the method of calculation the historical shares for pool miners is available at pool websites. Miners are advised to read and understand the method in order to optimize their mining performance and revenue.

**Mining revenue**

The mining revenue consists of the block reward and of transaction fees from the forged block. The block reward decreases each 10400 blocks. Depending on the plot size, the mining calculator provided in the Burst explorer can give an estimate regarding the expected mining revenue. The mining calculator is available at this link: <https://explore.burst.cryptoguru.org/tool/calculate> and will give an estimation using the current block reward and base target.

Technical information about mining
----------------------------------

**Mining and forging a block**

The entities involved in the process of mining and forging a block are the Burst wallet (locally installed, a web wallet, pool wallet) and the mining software (miner i.e. the software able to compute deadlines from plot files).

The mining process starts with the miner requesting mining information from the wallet. Prior to sending the mining information, the wallet will create the new generation signature by running the previous generation signature and the previous block generator through the Shabal256 hash function (1). The new generation signature is passed on to the miner, together with the base target value and the next block height (2). In the next step, the miner will produce the generation hash for the next block by running the Shabal256 function on the generation signature and block height received from the wallet (3). The generation hash is used as the argument of the modulo 4096 function in order to get the scoop number which will be used to process the plot files (4).

![](Intro_to_mining_burst.png "Intro_to_mining_burst.png")

*Process of mining and forging a block*

After the scoop number has been calculated, it is used to read all scoops from all nonces in all plot files. The processing is done individually for all nonces, by running them with the new generation signature through the Shabal256 hash function. The result of this is a hash referred to as target. The target is divided by the base target obtained from the wallet in step (2). The first 8 bytes of the result of the division is the deadline (5). To prevent the so-called “nonce spamming” the miner will check if the latest found deadline is lower than the lowest one found so far and proceed with it until a lower deadline has been found. Pools often set a maximum deadline limit (i.e. the highest deadline value that is accepted by the pool) and deadlines that exceed this limit will be discarded by the pool for historic share calculation. The miner will submit the deadline to the wallet, together with the numeric account ID bound to the plot file, and the nonce number that contains the scoop data used to generate the deadline. In case of mining solo, the information passed from the miner to the wallet will also include the secret passphrase of the account bound to the plot file (with pool mining, the passphrase of the pool account is used)(6). After the wallet has received the information from the miner, it will create the nonce to be able to find and verify the deadline submitted by the miner. If the deadline is verified, the wallet will wait for the deadline to expire (7) and check if the new valid block has already been announced on the network (8). If the miner submits new information (a new deadline), the wallet will create the nonce to check the validity of the deadline and if the newly submitted deadline is lower than the previous one. If this is the case, the wallet will use the lower deadline value (i.e. wait for it to expire). In case the new block hasn't been announced, and the deadline has expired, the wallet will forge a new block (9), otherwise it will discard the information received from the miner, since it is no longer valid. The the wallet starts to forge a block, it will collect unconfirmed transactions received by users or the network. After checking the validity of transactions (e.g. signatures, timestamps etc), it will try to fit as many as possible into a block or until all transactions have been processed. The constraints for forging a block are the maximum block payload, which cannot exceed 179,520 bytes (176 kB) and the maximum number of transactions in a block (the theoretical maximal number of transactions in a single block is 19.200, for more information on the slot-based transaction structure, please read this Burstcoin.ist article: <https://www.burstcoin.ist/2018/05/04/1st-hard-fork-explained-changes-in-transaction-dynamics/>). Note that transaction are not stored within the block, but separately.

**Contents of a block**

The Burst block explorer, as well as the wallet can be used to view block contents and information.

![](Block_contents.png "Block_contents.png")

*Block information available in the Burst block explorer*

Block version number refers to the block format, which determines what a block can contain and in which manner,

List of transaction Ids included in the block,

Payload hash is the Sha256 hash of all the data in the block payload,

Timestamp of the block forging derived from the birth of the blockchain (11th of August 2014, at 02:00:00),

Total amount i.e. the sum of all transactions contained in the block,

Total amount of transaction fees, which will be given to the block forger on tp of the block reward,

Payload length,

Public key of the account which forged the block,

Generation signature that was used to forge the block,

Previous block hash, the Sha256 hash of the contents of the previous block,

Previous block ID, which is the first 8 bytes of the previous block hash converted to a number,

Cumulative difficulty used to prevent “Nothing at Stake” problems during potential forks, calculated as: previous cumulative difficulty + 18446744073709551616/base target,

Base target used when the block was forged,

Block height,

Nonce number used to forge the block,

AT – payload bytes of the AT, in case AT was added to the block,

Block signature is a 64 byte hash generated with the forger's private key and block contents.

![](Block_details_1.png "Block_details_1.png")

![](Block_details_2.png "Block_details_2.png")

*Block details available in the Burst wallet*

Once a wallet forges a block, it will be announced to the network. The wallet connects to peers and send the block for verification and validation.

*Based on article by Quibus <http://files.getburst.net/Burstcoin-technical-information-about-mining-and-blockforging.pdf>*

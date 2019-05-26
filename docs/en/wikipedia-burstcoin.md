<!--This needs repair, many links targets nowhere. shefas 2019.05.26 -->
**Burstcoin** (also called **Burst**) is a digital [cryptocurrency](cryptocurrency.md) and [payment system](payment-system.md) based on the [blockchain](blockchain.md) technology. Burstcoin was introduced on the bitcointalk.org forum on 10 August 2014 as an [Nxt](nxt.md)-based currency. Burstcoins are [mined](cryptocurrency-mining.md) using an algorithm called [proof-of-capacity](proof-of-capacity.md) (PoC)[1] in which miners use computer storage instead of the more common energy-expensive method [proof-of-work](proof-of-work.md) (PoW) which involves permanent computational operations.

Burst was officially introduced on the 10th of August 2014 with the goal of solving other cryptocurrencies’ biggest problems : lack of decentralization, waste of energy, unfair releases, supremacy of big miners and corporations. It is the first and only cryptocurrency secured by the Proof-of-Capacity algorithm.

The critical difference between Proof of Work and Burst’s Proof of Capacity is that instead of needing ever more expensive, power hungry processors and graphic cards, it uses inexpensive, low-power hard drives. Proof of Capacity is inherently more secure, and trimmed versions of the blockchain are easier and more secure as well. The other biggest advantage that Proof of Capacity has going for it is its mining, working like a built-in coin faucet in the form of hard drives that allow anyone to earn free coins in exchange for providing extra security for the network.

The energy requirement for Burstcoin mining is minimal compared to most other cryptocurrencies, making Burstcoin one of the most [energy efficient](efficient-energy-use.md) within the field of proof-based cryptocurrencies.[2] Storage space is required however: as of March 2018, the estimated network size approached 341,000 [terabytes](terabyte.md).[3].

The Nxt blockchain platform allows for development flexibility, ensuring developers freedom to create their own [applications](application-software.md). In this sense, Burstcoin can be considered as a next-generation cryptographic application project (often called 'cryptocurrency 2.0') in contrast to the first generation cryptocurrencies like [bitcoin](bitcoin.md).[4]

Burstcoin is currently being developed by the PoCC (Proof of Capacity Consortium) since December 27th 2017. The PoC Consortium developers released a new white paper titled *The Burst Dymaxion*.[5] It describes the implementation of a tangle-based [lightning network](lightning-network.md) allowing for arbitrary scalability and anonymous transactions. The Dymaxion operates with payment channels opened using the [directed acyclic graph](directed-acyclic-graph.md) (DAG) technology on top of the Burst blockchain functionning with the Proof-of-Capacity protocol.

History
-------

### Origin to community takeover

Burstcoin was released to the public on 10 August 2014 on bitcointalk.org by the original developer who is known under the alias “Burstcoin”. Their real identity is still unknown today. The coin was launched without an [initial coin offering](initial-coin-offering.md) (ICO) or premine.[6] The genesis block was published on August 11th 2014.[7] Approximately one year later, the main developer “Burstcoin” disappeared without any explanation.

Being an [open source](open-source.md) project, other members of community then took over the development of Burstcoin: on January 11th 2016, a new forum thread was created on Bitcointalk.org by a senior community member.

On July 22nd 2017 an attacker spammed the Burst network with messages, causing wallets to crash and splitting the network into multiple forks.[8] The network remained unstable for several days. On August 11th 2017, in the aftermath of the spam attack, a new development team called the PoC Consortium[9] introduced itself on Bitcointalk.org. They have been recognized by the former development team as the designated successors for the Burst Reference Software.[10]

### Innovations

Burstcoin is the first cryptocurrency to use the proof-of-capacity algorithm. Proof-of-capacity was successfully implemented by the original developer, going by the name “Burstcoin” on bitcointalk.org forums.

Burstcoin was the first cryptocurrency to implement working, “[Turing complete](turing-completeness.md)” [smart contracts](smart-contracts.md)[11] in a live environment in the form of *Automated Transactions* (AT), this occurred before both [Ethereum](ethereum.md) and [Counterparty](Counterparty_(technology) "wikilink").[12][13][14] An application of these smart contracts was shown in the form of the world's first decentralized lottery.[15][16][17] It became the first ever program to run on top of a blockchain in a trustless decentralized manner. Other use cases of the Automated Transactions include decentralized [crowdfunding](crowdfunding.md).[18]

A more recent innovation by Burstcoin and Qora is the *Atomic cross-chain transactions* (ACCT),[19] this allows for full decentralized trading between two cryptocurrencies without the need for any third-party, namely an [online exchange](digital-currency-exchange.md). Cross-chain transactions have been successfully made between the blockchain of Burstcoin and Qora.[20][21]

On December 27th 2017, the PoC Consortium developers released a new white paper titled *The Burst Dymaxion*.[22] It describes the implementation of a tangle-based [lightning network](lightning-network.md) allowing for arbitrary scalability and anonymous transactions. The Dymaxion operates with payment channels opened using the [directed acyclic graph](directed-acyclic-graph.md) (DAG) technology on top of the Burst blockchain functionning with the Proof-of-Capacity protocol.

Design
------

### Blockchain

The Burstcoin *blockchain* is a public [ledger](ledger.md) that records every transaction. It is fully distributed and works without a central trusted authority: the blockchain is maintained by a network of computers known as [nodes](Node_(networking) "wikilink") running the Burstcoin software.

### Ownership

Ownership of Burstcoins implies that a user can spend Burstcoins linked with a specific address. For this to occur, a payer must [digitally sign](digital-signature.md) the transaction using the associated [private key](public-key-cryptography.md). Without knowledge of the private key, the transaction cannot be signed and Burstcoins cannot be spent. The network verifies the signature using the [public key](public-key-cryptography.md). If the private key is lost, the Burstcoin network will not recognize any other evidence of ownership; the coins are then unusable and are effectively lost.

### Transactions

A summary of a Burstcoin transaction is as follows: The sender details the parameters for the required transaction type (sending money, creating an alias, transmitting a message, issuing an asset or an order for an asset). All values for the transaction inputs are bounds checked for validity. If the transaction is found to be valid then the [public key](public-key-cryptography.md) for the generating account is computed using the supplied secret passphrase. A new transaction is created, with a type and sub-type value set to match the type of transaction being created. All specified parameters are included in the Transaction object. A unique transaction ID is generated with the creation of the object. The transaction is [digitally signed](digital-signature.md) using the sending account's [private key](public-key-cryptography.md). The encrypted transaction data is placed within a message instructing network peers to process the transaction. The transaction is then broadcast to all peers on the network. Burstcoin transactions are based on the [Nxt](nxt.md) code base. A detailed explanation of the transaction process can be found on its wiki page.[23]

### Mining (proof-of-capacity)

The mining process is based on a proof-of-capacity (PoC) consensus algorithm. In order to mine Burstcoins each miner first computes a large [data set](data-set.md) which is then saved to a computer storage medium. These data sets are known as plots. For each new block in the blockchain each miner will read through a tiny [subset](subset.md) (1/4096th - approximately 0.024%) of their own saved plots and return a result as a time interval in seconds known as a deadline, which is the time elapsed since last block creation after which this miner would be allowed to create a new block. The miner with the lowest deadline wins the block and is then rewarded with the transaction fees and the decreasing block reward of Burstcoins.

The computational resources for mining are limited to the time it takes the miner to read the plot files stored on mass storage. Once this is achieved, no other [computational resources](computational-resources.md) are needed until the next block, making Burstcoin highly energy efficient.[24] Miners compete on the size of their plot files, rather than on equipment speed as is the case with other cryptocurrencies.

Proof-of-capacity is also claimed to be [ASIC](application-specific-integrated-circuit.md)-proof. Burstcoin's proof-of-capacity algorithm is based on precomputed [proof-of-work](proof-of-work-system.md), so theoretically one could compute the proofs in real time. However, it is currently impossible to efficiently compute a significant amount of work during a 4-minute time interval. Inspecting the precomputed work stored on the [hard drive](hard-drive.md) is both faster and more energy efficient than any conceivable ASIC device could achieve providing in real-time.

### Mining pools

Given that it can take a long time to find the smallest deadline, some miners collectively mine in what is known as a mining [pool](Pooling_(resource_management) "wikilink").[25] Mining pools allow miners to have a more evenly distributed Burstcoin income: the reward for each block won by the pool is distributed between the miners of that pool (sometimes with the largest share going to the miner that found it).

If a pool has a “deadline limit”, it means the reward is shared only among miners whose submitted deadlines for the winning block came under the pool's deadline limit. This tends to skew payment distribution in favor of the larger miners in the pool (who have a higher probability of meeting that deadline limit for the winning block) while still distributing the proceeds to more participants than occurs with solo mining, where one miner wins 100% of the block reward.

To reduce transaction fees, most pools also have a minimum payment threshold, meaning the pool must owe you at least that much before it pays you.

Usage
-----

The Burstcoin [cryptocurrency wallet](cryptocurrency-wallet.md) is implemented as a [Java](Java_(programming_language) "wikilink")-based [Web server](web-server.md), which users can either run themselves (as is done by the Windows client), or, for convenience, use on somebody else's server machine (for example one hosted by a Burstcoin [mining pool](mining-pool.md)), if they trust said machine not to be copying their passphrase via [keystroke logging](keystroke-logging.md) (mining pools have an obvious incentive to maintain their reputation by not compromising their convenience wallet servers).

The passphrase (typically 12 dictionary words) is all that is needed to transfer funds out of an account, and any wallet server can convert such a passphrase into a public address (a string of alphanumeric characters beginning `BURST-`) which is used for receiving Burstcoin, and an Account ID (an integer) which is needed to generate the “plot” files for mining (for example with the `wplotgenerator` tool).

Features
--------

The core feature set of Burst is based on the [Nxt](nxt.md) platform, which allows the adding of external services to be built on top of the blockchain.

### Android wallet

The first [Android](Android_(operating_system) "wikilink") version of the Burst wallet was released in 2016[26] and was able to mine on Android devices. This app is now considered “outdated” by burst-coin.org[27] which instead recommends the POC Consortium wallet[28] or the BurstNation wallet.[29] These newer apps cannot mine on Android but do allow users to create wallets and to send and receive Burstcoin on any Android device.

### Asset Exchange

![Screenshot of the Burstcoin Asset Exchange window](Burst_Asset_Exchange_997x787.jpg "fig:Screenshot of the Burstcoin Asset Exchange window") The Burst Asset Exchange is a [peer-to-peer](peer-to-peer.md) exchange platform integrated into the Burst wallet. It functions primarily as a secure decentralized [trading platform](trading-platform.md) for Burst assets. The popularity of the asset exchange is based upon the absence of any third party, allowing improved efficiency and reduced costs. A burst asset is basically a token to represent anything the asset issuer deems to be of value so that it can be traded, common examples of such assets include shares in the following: mining pools, retirement funds, crypto mining rigs, crypto gambling sites and silver investments.

### Automated transactions (Smart Contracts)

![Screenshot of First Smart Contract using Burstcoin](Burstcoin_first_smart_contract.jpg "fig:Screenshot of First Smart Contract using Burstcoin") [Smart Contracts](smart-contract.md) are self-executing contractual states, stored on the blockchain.[30] In brief an Automated Transaction is a “[Turing complete](turing-completeness.md)” set of byte code instructions which will be executed by a byte code interpreter built into its host. An AT supporting host platform automatically supports various applications ranging from games of chance to automated crowdfunding and ensuring that “long term savings” are not lost forever.[31]

### Crowdfunding

The [crowdfunding](crowdfunding.md) feature allows users within the Burst community to raise funds in Burst for project creators in a decentralized way. Funds are refunded if the project target is not met.

### Escrow service

The Burstcoin Wallet has an inbuilt [escrow](escrow.md) service, which allows a quantity of Burstcoins to be held by a third party on behalf of transacting parties.

### Marketplace

The Burstcoin Wallet includes a completely decentralized [marketplace](marketplace.md) where Burstcoin users can view other users items for sale by referencing their account id. It will display all items for sale for the designated account holder.[32]

Notes
-----

References
----------

[1] 

[2] 

[3] <https://explore.burst.cryptoguru.org/chart/supply/network_size>

[4] 

[5] 

[6] 

[7] 

[8] 

[9] 

[10] 

[11] 

[12] 

[13] 

[14] 

[15] 

[16] 

[17] 

[18] 

[19] 

[20] 

[21] 

[22] 

[23] [How Tx Processing Works](https://nxtwiki.org/wiki/How_Tx_Processing_Works)

[24] 

[25] 

[26] 

[27] [Download wallet](https://www.burst-coin.org/download-wallet) “Burst-team Android wallet (outdated)”

[28] [POC Consortium Android Wallet](https://play.google.com/store/apps/details?id=org.icewave.burstcoinwallet) on Google Play Store

[29] [BurstNation Android Wallet](https://play.google.com/store/apps/details?id=com.official.bnwallet) on Google Play Store

[30] 

[31] 

[32] 

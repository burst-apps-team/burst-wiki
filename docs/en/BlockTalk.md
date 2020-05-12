Introduction
------------

[BlockTalk](https://github.com/burst-apps-team/blocktalk) is a platform for writing smart contracts to run on the Burstcoin blockchain.

[![BlockTalk](http://img.youtube.com/vi/zEdb-9teaOY/0.jpg)](http://www.youtube.com/watch?v=zEdb-9teaOY)

Burstcoin was the first blockchain to implement a turing-complete [smart contract](https://www.burst-coin.org/smart-contracts)
system in the form of [*Automated Transactions* (AT)](automated-transaction).
However, before BlockTalk, the creation and deployment of smart contracts required writing
(assembler-like) bytecode and testing on-chain, making the development of contracts cumbersome.

BlockTalk allows the user to write, debug, and deploy Burst smart contracts relying only on Java.
You can use a simple text editor or your preferred IDE.
BlockTalk consists of the following key components:
 - **[Contract.java](https://github.com/burst-apps-team/blocktalk/blob/master/src/main/java/bt/Contract.java)**: a Java abstract class defining the basic API available for contracts
 - **Emulator**: an emulated blockchain and respective UI
 - **Compiler**: a system to convert Java bytecode into Burst AT bytecode that can run on the Burst blockchain 

[![Simple hello world contract](http://img.youtube.com/vi/XcN5WxqjjGw/0.jpg)](https://www.youtube.com/watch?v=XcN5WxqjjGw "BlockTalk sample application")

## Sample Contracts
Take a look on the [samples source folder](https://github.com/burst-apps-team/blocktalk/blob/master/src/main/java/bt/sample).

### Vision

The vision of BlockTalk is to have a state-of-the-art and easy to deploy framework for smart contracts. The project decision was to make use of Java as the first high-level language to implement this framework. Although the project is already underway, anyone who like to join the project is welcome and should send a contact request to frank_the_tank on discord.

### Smart contract upgrades

After the SODIUM hard-fork (BRS version 2.5 and above) much more powerfull smart contracts can run on the Burst blockchain, for more information please see [CIP20](https://github.com/burst-apps-team/CIPs/blob/activate-cip20/cip-0020.md).


Burst NFTs?
------------

NFTs ([non-fungible-tokens](https://en.wikipedia.org/wiki/Non-fungible_token)) are basically new and unique representations of goods or assets that take the form of digital tokens. Through the use of cryptography, NFTs can prove the authenticity, as well as ownership of such assets and goods.

Let's say that there is a virtual artwork that is tokenized. With that in mind, whoever holds the tokens, acts as the direct owner of the piece of art itself. Obviously, NFTs are quite unique, and no other item can replace them. The possible implications of this are many, and they might even create an entirely new class of digital assets in the future. It is even possible that real-world items and assets might be this closely tied to digital tokens, all with the goal of securing the ownership of such items. The NFTs were brought to the mainstream together with CryptoKitties. This is a concept that gained a lot of popularity near the end of the previous year. Some of them were so popular, that they were priced and sold for hundreds of thousands of dollars.

With BlockTalk one can easily create NFTs that can not only be transfered or sold but also auctioned. Check out the source code of one implementation [here](https://github.com/burst-apps-team/blocktalk/blob/master/src/main/java/bt/sample/NFT2.java). 

References
----------

1. [https://www.burstcoin.ist/2019/04/06/blocktalk-the-new-smart-way-for-burst-ats/](https://www.burstcoin.ist/2019/04/06/blocktalk-the-new-smart-way-for-burst-ats/)

2. [https://www.reddit.com/r/burstcoin/comments/bi26xi/blocktalk_introduce_nftnonfungibletoken/](https://www.reddit.com/r/burstcoin/comments/bi26xi/blocktalk_introduce_nftnonfungibletoken/)

3. [https://github.com/burst-apps-team/CIPs/blob/activate-cip20/cip-0020.md](https://github.com/burst-apps-team/CIPs/blob/activate-cip20/cip-0020.md)

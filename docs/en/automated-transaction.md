Introduction
------------

Automated Transactions are how Burst implements smart contracts.

The formalism Burst implements for Automated Transactions (AT) is a technology created by [CIYAM](http://ciyam.org/at/) Developers which provides “Turing complete smart contracts”.

Burstcoin was actually the first to implement a turing-complete blockchain based smart contract system. However, up to recently the creation (or programming) of smart contracts required writing (assembler-like) bytecode and testing directly on-chain, making the development of contracts cumbersome.

More recently the [BlockTalk](https://github.com/burst-apps-team/blocktalk) platform was created. It allows the user to write, debug, and deploy Burst smart contracts relying only on Java. You can use a simple text editor or your preferred IDE. BlockTalk consists of the following key components:

 - Contract.java: a Java abstract class defining the basic API available for contracts
 - Emulator: an emulated blockchain and respective UI
 - Compiler: a system to convert Java bytecode into Burst AT bytecode that can run on the Burst blockchain

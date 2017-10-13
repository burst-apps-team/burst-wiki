Proof of Capacity PoC (also called Proof of Space) is, in general, a mechanism to ensure that a party has dedicated a particular (significant) amount of storage space towards a specific goal. Burst uses PoC in the process of transaction verification for the validation and generation of blocks. Here, [miners](mining.md) use a given amount of disk space to fill with [plots](plots.md). The more disk space is plotted, the higher the chances to generate (win) a block. PoC is thus in direct contrast to other consenus algorithms like Proof of Work (PoW) and Proof of Stake (PoS) where block generators need to proof the performance of some difficult computations or the ownership of a significant amount of currency, respectively.

Advantages of PoC
=================

In contrast to PoW, PoC is considered to be more eco friendly [1]. While PoW requires expensive computations each time a block is generated, PoC puts most of the required computational costs up-front in the plotting process. Reading the plots in order to proof the designated disk space investment is much cheaper. Additionaly, users can mine using their existing hardware, instead of having to to invest in special purpose hardware like e.g. ASICS for Bitcoin mining.

In contrast to PoS, PoC allows a fairer distribution of coins as they are generated over time and assigned to active miners. In a PoS coin like NXT, the number of coins is fixed from the beginning and has to be distributed *somehow* initially. This leads to an unequal distribution of wealth to a few selected early adopters, e.g. via airdrops.

Further Reading
===============

-   [Proofs of Space](https://eprint.iacr.org/2013/796.pdf) paper by Dziembowski, Faust, Kolmogorov and Pietrzak
-   [SpaceMint](https://eprint.iacr.org/2015/528.pdf) paper by Sunoo Park, Kwon, Fuchsbauer, Pietrzak et al.

References
==========

<references />

[1] <https://www.burstcoin.ist/how-burst-puts-bitcoin-to-shame/>

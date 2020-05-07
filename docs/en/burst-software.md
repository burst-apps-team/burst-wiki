Introduction
------------

The Burst ecosystem uses lots of software. A majority of these are used either to plot, mine, or simply to transfer BURST. A majority of users will want a local wallet, and there are plenty to choose from!

Wallets
-------

### QBundle

The easiest choice for beginners using Windows is to download Qbundle. QBundle is a convenient launcher which installs the BRS local wallet, plotting software, and mining software with a few clicks. 

* Main Developer : harry1453

* Download link : <https://github.com/burst-apps-team/Qbundle/releases>

* Platform : Windows

* Installation Guide : [QBundle Installation and User Guide](qbundle.md)

### Phoenix

According to [CIP-18](https://github.com/burst-apps-team/CIPs/blob/master/cip-0018.md) this is the new Cross-Platform Burst Wallet. A multi-platform application architecture is made possible through a few key front-end technologies: React Native, Angular, and Electron. By using these technologies in conjunction, Burst wallet developers will enjoy a modern application development workflow without sacrificing the quality of the end-product. 

* Main Developer : Burst Apps Team (BAT)

* Download link : <https://github.com/burst-apps-team/phoenix/releases>

* Platform : Windows, MacOS, Linux, iOS, Android

### Burst Reference Software (BRS)

The [Burst Reference Software (BRS)](burst-reference-software.md) is the main Burst wallet. It allows connection to the Burst network. The Burst Reference Software is the wallet/software included in QBundle. This wallet version is developed and maintained by the Burst Apps Team (BAT) and supports a multitude of database backends.

* Main Developer : Burst Apps Team (BAT)

* Download link : <https://github.com/burst-apps-team/burstcoin/releases/latest>

* Platform : Windows, Linux, macOS, Docker

* Installation Guide : [Windows Installation](brs-windows-installation.md), [Linux Installation](brs-linux-installation.md), [macOS Installation](brs-macos-installation.md), [Docker Installation](brs-docker-installation.md)

### BurstHotWallet

A light-weight version of the current BURST wallet for easy access to send and receive BURST. 

* Main Developer : CurbShifter

* Download link : <https://github.com/CurbShifter/BurstHotWallet/releases>

* Platform : Windows, MacOS

### Easy2Burst \[ARCHIVED\]

Easy2Burst was a Burst Wallet written in Rust, as the successor of QBundle. It is no longer being developed [<sup>\[1\]</sup>](https://github.com/HeosSacer/Easy2Burst/issues/2#issuecomment-531654134)

* Main Developer : HeosSacer

* Github repo : <https://github.com/HeosSacer/Easy2Burst>

* Platform : Windows, Linux, MacOS

### Aspera \[ARCHIVED\]

Aspera was a Burst Wallet written in go. This wallet opens up a whole new dimension for Burst with state-of-the-art technologies (partial blocks validation, new front-end, etc.). It is no longer being developed since the PoCC exit of Burstcoin.

* Main Developer : PoC-Consortium

* Github repo : <https://github.com/PoC-Consortium/Aspera>

* Platform : Windows, Linux, MacOS

Plotting Software
-----------------

A large collection of plotting software has been developed for specific needs, specific hardware, and specific operating systems.

### ENGRAVER

Engraver is a BURST plotter that natively generates optimized plot files. It generates PoC2 files and is intended to work on most platforms. Being the fastest plotter there is... it features: direct and async i/o, SIMD support (sse2, avx, avx2, avx512f), and gpu support.

* Main Developer : PoC-Consortium

* Download link : <https://github.com/PoC-Consortium/engraver/releases>

* Platform : Linux, Unix, Windows, MacOS (32&64bit)

* Installation guide : [Compile from Source](engraver-how-to-compile-from-source.md) 

### TurboPlotter 9000

TurboPlotter 9000 is a CPU/GPU PoC 1 + PoC 2 plotter and plot integrity checker.

* Main Developer : Blackpawn

* Download link : <https://blackpawn.com/tp/>

* Platform : Windows, Linux, MacOS

### XPlotterGui

XplotterGui is a GUI version of Xplotter with SSD Cache, File Merging and POC2 Conversion features.

* Main Developer : JohnnyFFM

* Download link : <https://github.com/JohnnyFFM/XPlotterGui/releases/latest>

* Platform : Windows

### XPlotter modded

XPlotter is the default plotter that comes with QBundle and provides CPU PoC 1 plotting for BURST whith AVX support. This modded version supports PoC2 plotting.

* Main Developer : JohnnyFFM

* Download link : <https://github.com/JohnnyFFM/XPlotter/releases/latest>

* Platform : Windows

### gpuPlotGenerator

gpuPlotGenerator is a GPU PoC 1 plotter.

* Main Developer : bhamon

* Download link : <https://github.com/bhamon/gpuPlotGenerator/releases/latest>

* Platform : Windows, Linux, MacOS

Mining Software
----------------

The following software is used to mine plot files.

### Scavenger

Scavenger is a Burst miner developed by PoC Consortium and written in Rust. It features direct io, avx512f, avx2, avx, sse and opencl.

* Main Developer : PoC Consortium

* Download link : <https://github.com/PoC-Consortium/scavenger/releases/latest>

* Platform : Windows, MacOS, Linux, Unix (64 bit)

### Gyra

Gyra is a front-end tool that creates formally valid and logically consistent configurations for the Scavenger Proof-of-Capacity miner. The configuration file used by Scavenger and created by Gyra is a .yaml file, which users can download in their browser and save it in the root directory where the Scavenger executable is placed.

* Main Developer : feyd27

* Source Code : <https://github.com/feyd27/gyra>

* Download link : <https://gyra-for-scavenger.firebaseapp.com/config>

* Announcement : <https://www.reddit.com/r/burst/comments/cuzect/gyra_for_scavenger_a_submission_for_the_2019/>

* Platform : Cross Platform

### foxy-miner

A scavenger / conqueror wrapper for collision free multi-mining.

* Main Developer : felix

* Download link : <https://github.com/felixbrucker/foxy-miner/releases>

* Platform : Windows, MacOS, Linux

### Blagominer quetzalcoatl

A modded version based off of andzno1's version of Blagominer. This version features major improvements and bug fixes.  

* Main Developer : quetzalcoatl

* Download link : <https://github.com/quetzalcoatl/blagominer/releases>

* Platform : Windows

### Blagominer andzno1

A modded version based off of JohnnyFFM's version of Blagominer. This version features collision free dual mining of Burstcoin and Bitcoin HD (configurable via priority) and tracking (and displaying) of possibly corrupted plot files.  

* Main Developer : andzno1

* Download link : <https://github.com/andzno1/blagominer/releases>

* Platform : Windows

### Blagominer JohnnyFFM

This is a modified version of Blagominer. This modification of Blagominer builds upon modifications previously made by Quibus. 

* Main Developer : JohnnyFFM

* Download link : <https://github.com/JohnnyFFM/miner-burst/releases/latest>

* Platform : Windows

### Burstcoin-jminer

Burstcoin-jminer is a GPU assisted Proof of Capacity (PoC) Miner for Burstcoin.

* Main Developer : de-luxe

* Download link : <https://github.com/de-luxe/burstcoin-jminer/releases/latest>

* Platform : Windows, Linux, MacOS

### CreepMiner

The creepMiner is a client application for mining Burst on a pool or solo. It supports CPU mining (SSE2/SSE4/AVX/AVX2) or GPU mining (OpenCL, CUDA).

* Main Developer : Creepsky

* Download link : <https://github.com/Creepsky/creepMiner/releases/latest>

* Platform : Windows, MacOS, Linux x86\_64 or ARM7

### The bencher

A PoW miner for PoC. Its main purpose is to benchmark hashing devices for PoC and to make different hardware comparable. It also has an educational aspect as it shows how inefficient it is to mine a PoC coin without storage (a BIS GPU can emulate around 1TB).

* Main Developer : JohnnyFFM

* Download link : <https://github.com/PoC-Consortium/bencher/releases>

* Platform : Windows, MacOS, Linux, Unix (64 bit)

Proxy Software
-----------------

### Archon

A collision free, multi-chain proof of capacity mining proxy.

* Main Developer : Aya

* Download link : <https://github.com/Bloodreaver/Archon/releases>

* Platform : Windows, Linux

### foxy-proxy

A BHD, LHD, DISC, BOOM and BURST proxy which supports solo and pool mining upstreams 

* Main Developer : felix

* Download link : <https://github.com/felixbrucker/foxy-proxy/releases>

* Platform : Windows, Linux

dApps
-----

### CloudBurst

CloudBurst is a decentralized cloud storage application to upload & download files from the Burstcoin blockchain.

* Main Developer : CurbShifter

* Download link : <https://github.com/CurbShifter/CloudBurstDAPP/releases/latest>

* Platform : Windows, MacOS

### BurstCoupon

BurstCoupon allows you to create and claim password protected coupons as extension for the Burstcoin blockchain.

* Main Developer : CurbShifter

* Download link : <https://github.com/CurbShifter/BurstCoupon/releases/latest>

* Platform : Windows, MacOS

Libraries - Development tools
-----------------------------

### BurstLib

BurstLib is a cross platform dynamic library to make developing applications and tools compatible with the Burst blockchain easier.

* Main Developer : Curbshifter

* Download link : <https://github.com/CurbShifter/BurstLib/releases/latest>

* Platform : Windows, MacOS, Linux, Unix

### @burstjs

@burstjs is a BURST JavaScript library to make it easier for developers to build apps using Angular/React/Node.js/JavaScript. The library is broken down into four packages: @burstjs/core, @burstjs/crypto, @burstjs/http, and @burstjs/util.

* Main Developer : blankey1337 (BAT) and Ohager (BAT)

* Download link : <https://burst-apps-team.github.io/phoenix/index.html>

* Platform : Cross Platform

### Burstkit4j

Burstkit4j is a BURST Java Development Framework.

* Main Developer : harry1453 (BAT)

* Download link : <https://github.com/burst-apps-team/burstkit4j/releases>

* Platform : Cross Platform

Pool Software 
--------------

### Nogrod

Nogrod is the all-new PoCC pool software written in GO. It provides SSE4 + AVX2 support, a fair share system based on estimated capacity, grpc api, multiple wallets as backends using the wallet API, can talk directly to wallet database, and dynamic payout thresholds/intervals based on messages on the blockchain. 

* Main Developer : PoC-Consortium

* Download link : <https://github.com/PoC-Consortium/Nogrod>

* Platform : Linux

### BurstNeon Burst-Pool

This is a pool software developed by [BurstNeon](http://burstneon.com).

* Main Developer : BurstNeon

* Download link : <https://github.com/BurstNeon/burst-pool/releases>

* Platform : Linux

Other tools
-----------

### TurboSwizzler

TurboSwizzler is an application to convert PoC 1 plots to the PoC 2 format. Currently supports drive to drive mode with plot merging.

* Main Developer : Blackpawn

* Download link : <https://blackpawn.com/tp/>

* Platform : Windows

### Poc1to2Converter

Poc1to2Converter converts PoC1 plots to POC2 plots.

* Main Developer : JohnnyFFM

* Download link : <https://github.com/JohnnyFFM/Poc1to2Converter/releases/latest>

* Platform : Windows

### PoC-Consortium Poc1to2 Converters

This Poc1to2 Converter for Linux converts PoC1 plots to POC2 plots. There is a rust binary converter (recommended) or a perl script version.

* Main Developer : PoC-Consortium

* Download link : <https://github.com/PoC-Consortium/Utilities>

* Platform : Linux

### Burstcoin Explorer Mobile App

Burstcoin explorer is a block explorer for the Burstcoin network

Features include:

-   Search for and view blocks, accounts and transactions
-   Keep a list of accounts you are watching
-   Keep an eye on the current Burst price, both in BTC and USD
-   Observe the current status of the Burstcoin Network
-   Countdown upcoming events such as the Pre-Dymaxion Hard Fork

* Main Developer : harry1453

* Download link : <https://github.com/harry1453/burstcoin-explorer-android/releases/latest> | <https://play.google.com/store/apps/details?id=com.harrysoft.burstcoinexplorer>

* Platform : Android

### PaperBurst

PaperBurst is a paper wallet generator for Burstcoin.

* Main Developer : Umbrellacorp03

* Download link : <https://github.com/umbrellacorp03/PaperBurst/releases/latest>

* Platform : Windows

### BlockTalk

BlockTalk allows the user to write, debug, and deploy Burst smart contracts relying only on Java. You can use a simple text editor or your preferred IDE.

* Main Developer : jjos

* Download link : <https://github.com/jjos2372/blocktalk>

* Platform : Cross Platform

* More Information: [BlockTalk](BlockTalk.md)

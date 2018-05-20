| Burst Software |
|----------------|
| **Status**     |

Table Of Contents
-----------------

\_\_TOC\_\_

Introduction
============

The Burst ecosystem uses lots of software. Each of these softwares are used either to plot, to mine, or simply to transfer BURST.

Wallets
-------

### [QBundle](qbundle.md) <img src="Stable.png" title="fig:Stable.png" alt="Stable.png" width="75" height="75" /> <img src="Local.png" title="fig:Local.png" alt="Local.png" width="75" height="75" /> <img src="Dymaxion_compatible.png" title="fig:Dymaxion_compatible.png" alt="Dymaxion_compatible.png" width="135" height="135" /> <img src="PoC_1.png" title="fig:PoC_1.png" alt="PoC_1.png" width="75" height="75" /> <img src="PoC_2.png" title="fig:PoC_2.png" alt="PoC_2.png" width="75" height="75" />

The easiest choice for beginners on Windows is to download Qbundle, a launcher allowing you to install the Burst local wallet in a few clicks. If you are a first timer you can simply follow the wizard walking you through the set up. Qbundle also includes plotting and mining software to start mining easily. You can find a download link below or on your right.

Download link : https://github.com/PoC-Consortium/Qbundle/releases/tag/v2.0.2

### Burstcoin mobile wallet <img src="Stable.png" title="fig:Stable.png" alt="Stable.png" width="75" height="75" /> <img src="Online.png" title="fig:Online.png" alt="Online.png" width="69" height="69" /> <img src="Dymaxion_compatible.png" title="fig:Dymaxion_compatible.png" alt="Dymaxion_compatible.png" width="135" height="135" /> <img src="PoC_1.png" title="fig:PoC_1.png" alt="PoC_1.png" width="75" height="75" /> <img src="PoC_2.png" title="fig:PoC_2.png" alt="PoC_2.png" width="75" height="75" />

The mobile wallet for Android and iOS which provides the following features :

-   Watch only addresses
-   Currency conversion
-   Client-side encryption and decryption
-   Encrypted / Unencrypted transaction messages
-   Contact book
-   QR code support
-   Secure and easy passphrase generation
-   Support for over 15 languages
-   Support for over 30 currencies

Download link : https://play.google.com/store/apps/details?id=org.icewave.burstcoinwallet

Plotting Softwares
------------------

There are several plotting softwares that are designed for specific needs and specific hardwares.

### XPlotter <img src="Stable.png" title="fig:Stable.png" alt="Stable.png" width="83" height="83" /> <img src="Dymaxion_incompatible.png" title="fig:Dymaxion_incompatible.png" alt="Dymaxion_incompatible.png" width="160" height="160" /> <img src="PoC_1.png" title="fig:PoC_1.png" alt="PoC_1.png" width="75" height="75" />

XPlotter is the default plotter that comes with QBundle and provides CPU PoC 1 plotting for BURST whith AVX support.

Download link : https://github.com/Blagodarenko/XPlotter/releases

### TurboPlotter 9000 <img src="Stable.png" title="fig:Stable.png" alt="Stable.png" width="83" height="83" /> <img src="Dymaxion_compatible.png" title="fig:Dymaxion_compatible.png" alt="Dymaxion_compatible.png" width="135" height="135" /> <img src="PoC_1.png" title="fig:PoC_1.png" alt="PoC_1.png" width="75" height="75" /> <img src="PoC_2.png" title="fig:PoC_2.png" alt="PoC_2.png" width="75" height="75" />

TurboPlotter 9000 is a CPU/GPU PoC 1 + PoC 2 plotter and plot integrity checker.

Download link : https://blackpawn.com/tp/

### gpuPlotGenerator <img src="Stable.png" title="fig:Stable.png" alt="Stable.png" width="83" height="83" /> <img src="Dymaxion_incompatible.png" title="fig:Dymaxion_incompatible.png" alt="Dymaxion_incompatible.png" width="160" height="160" /> <img src="PoC_1.png" title="fig:PoC_1.png" alt="PoC_1.png" width="75" height="75" />

gpuPlotGenerator is a GPU PoC 1 plotter.

Download link : https://github.com/bhamon/gpuPlotGenerator/releases

Mining Softwares
----------------

The following softwares are used to mine plot files.

### Blagominer <img src="Stable.png" title="fig:Stable.png" alt="Stable.png" width="83" height="83" /> <img src="Dymaxion_incompatible.png" title="fig:Dymaxion_incompatible.png" alt="Dymaxion_incompatible.png" width="160" height="160" /> <img src="PoC_1.png" title="fig:PoC_1.png" alt="PoC_1.png" width="75" height="75" />

Blagominer is the default miner that comes with QBundle. It does support AVX/AVX2.

Download link : https://github.com/Blagodarenko/miner-burst/releases

### CreepMiner <img src="Stable.png" title="fig:Stable.png" alt="Stable.png" width="83" height="83" /> <img src="Dymaxion_incompatible.png" title="fig:Dymaxion_incompatible.png" alt="Dymaxion_incompatible.png" width="160" height="160" /> <img src="PoC_1.png" title="fig:PoC_1.png" alt="PoC_1.png" width="75" height="75" />

The creepMiner is a client application for mining Burst on a pool or solo. It does support CPU mining (SSE2/SSE4/AVX/AVX2) or GPU mining (OpenCL, CUDA).

Download link : https://github.com/Creepsky/creepMiner/releases

### Mining

-   [plot software (aka Plotter)](plot-software.md)
-   [mining software](mining-software.md)

Pool Software 
--------------

### GoBurstPool <img src="Stable.png" title="fig:Stable.png" alt="Stable.png" width="83" height="83" /> <img src="Dymaxion_compatible.png" title="fig:Dymaxion_compatible.png" alt="Dymaxion_compatible.png" width="135" height="135" />

GoBurstPool is the all-new PoCC pool software written in GO. It provides SSE4 + AVX2 support, fair share system based on estimated capacity, grpc api, multiple wallets as backends using the wallet API, can talk directly to wallet database, and finaly dynamic payout thresholds/intervals based on messages on the blockchain.

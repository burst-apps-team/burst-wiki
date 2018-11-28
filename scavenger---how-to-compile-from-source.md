<languages></languages>

| Scavenger - How to compile from source |
|----------------------------------------|
| **Status**                             |
| **Credits**                            |

Description
-----------

Scavenger is a burst miner developped by the PoC Consortiun and written in Rust. If features direct io, avx 512f, avx2, avx, sse and opencl support. At the time of writing this is the preffered miner of choice for both CPU and GPU mining.

The following steps are to install Scavenger 1.6.4 by compiling from source on a fresh installation of **Debian Linux version 9 “Stretch”**. Please be sure to check Git Repository for the latest release before proceeding.

**Note:** Commands to be entered into the shell will be formatted in bold text, e.g.:

**`date`**

example command outputs will not be bold, e.g.:

`Wed 28 Nov 12:02:18 GMT 2018`

Prerequisites
-------------

Before you begin the installation, the following packages need to be installed:

-   curl
-   git
-   build-essential

Change to root user:

**`su`**

Enter root password

Install curl, git and build essential packages:

**`apt`` ``install`` ``curl`` ``git`` ``build-essential`` ``-y`**

Once the above packages have completed installation, its time to install Rust.

**Note:** The features used when compiling Scavenger are not available on the stable release channel of rust, therefore the nightly channel needs to be installed.

Installing Rust (nightly)
-------------------------

To install Rust, run the following command:

**`curl`` `[`https://sh.rustup.rs`](https://sh.rustup.rs)` ``-sSf`` ``|`` ``sh`**

You will be presented with three options (see below):

`Current installation options:`
`  default host triple: x86_64-unknown-linux-gnu`
`    default toolchain: stable`
` modify PATH variable: yes`
`1) Proceed with installation (default)`
`2) Customize installation`
`3) Cancel installation`
`>`

To install the nightly channel chose **option 2**, this will ask you additional questions, answer as per below:

`Q: Default host triple?`
`A: Press `**`enter`**` to leave unchanged`

`Q: Default toolchain? (stable/beta/nightly/none)`
`A: `**`nightly`**

`Q: Modify PATH variable? (y/n)`
`A: `**`y`**

You will be presented with a summary of the current installation options as per your answers provided above:

`Current installation options:`
`  default host triple: x86_64-unknown-linux-gnu`
`    default toolchain: nightly`
` modify PATH variable: yes`
`1) Proceed with installation (default)`
`2) Customize installation`
`3) Cancel installation`

Chose **option 1** to install

Once installation is complete, follow the instructions to configure your current terminal sessions' environment variable so that you can run Rust related commands:

**`source`` ``$HOME/.cargo/env`**

Switching from Rust (stable) to Rust (nightly)
----------------------------------------------

If you chose to or already have the stable channel of rust installed you can switch channels by using the following command:

**`rustup`` ``default`` ``nightly`**

To confirmthe instalation of the correct channel of Rust, type the following:

**`rustc`` ``--version`**

This should give you the output below, note the “*nightly*” present (version correct at time of writing):

`rustc 1.32.0-nightly (6bfb46e4a 2018-11-26)`

Obtaining the source files
--------------------------

To download the source files from the Git repository, run the following command, (please make sure to check the exact path to the latest version as this will change as later versions are released) :

**`wget`` `[`https://github.com/PoC-Consortium/scavenger/archive/1.6.4.zip`](https://github.com/PoC-Consortium/scavenger/archive/1.6.4.zip)**

This will download the file 1.6.4.zip to your current directory.

Extract the contents of the zip file:

**`unzip`` ``1.6.4.zip`**

This will create a new directory named scavenger-1.6.4.

Compiling the source files
--------------------------

Navigate to the scavenger-1.6.4 directory

**`cd`` ``scavenger-1.6.4`**

At this point it is worth deciding what you are going to use to do the mining, CPU/GPU and also what instruction sets are available on your device. See the **[README.md](https://github.com/PoC-Consortium/scavenger/tree/1.6.4#scavenger---burstminer-in-rust)** file located in the scavenger-1.6.4 directory for an overview of features to use.

In this example, a CPU with the SSE2 instruction set will be used therefore the simd feature option is required (the same would work for AVX, AVX2 etc). Run the following command to compile the code:

**`cargo`` ``build`` ``--release`` ``--features=simd`**

Post compiling steps
--------------------

Copy the scavenger executable to the scavenger-1.6.4 directory (two levels up in this case). This placed it in the same directory as the config.yaml file which is required for it to run.

**`cp`` ``scavenger`` ``../../scavenger`**

Navigate back to your scavenger-1.6.4 directory from the release directory

**`cd`` ``../..`**

At this point you can run scavenger and it will make use of the **test\_data** file to simulate mining on the dev.burst-test.net

Once you know it will run you will need to configure the config.yaml file to your setup. The config file itself is well commented and the GitHub Wiki has further information if required:

<https://github.com/PoC-Consortium/scavenger/wiki>

Additional notes
----------------

If the above steps were followed verbatim, all the scavenger files will be owned by root. It is not recommended to run as root so we will change the owner to a normal user, “burst” in this case. From the directory a level up from your scavenger-1.6.4 directory run:

**`chown`` ``-R`` ``burst:burst`` ``scavenger-1.6.4`**

This will recursively change ownership to the user “burst” for all files and subdirectories within the scavenger-1.6.4 directory.

to logout of root type:

**`exit`**

You are now back at the shell as your normal user (burst). You can now run scavenger by issuing the following command from the scavenger-1.6.4 directory:

`./scavenger`

Example output from running scavenger

<img src="Run-scav.PNG" title="Run-scav.PNG" alt="Run-scav.PNG" width="709" height="364" />

Links
-----

**Software**

<https://github.com/PoC-Consortium/scavenger/releases/latest>

**GitHub Documentation**

<https://github.com/PoC-Consortium/scavenger/wiki>

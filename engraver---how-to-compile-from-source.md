<languages></languages>

| Engraver - How to compile from source |
|---------------------------------------|
| **Status**                            |
| **Credits**                           |

Description
-----------

Engraver is the reference Burstcoin (or Burst) cross-platform plotting software, endorsed by the PoC Consortium, created to further improve and optimize Burst plot file creation.

The following steps are to install Engraver 2.2.0 by compiling from source on a fresh installation of **Debian Linux version 9 “Stretch”**. Please be sure to check Git Repository for the latest release before proceeding.

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

Installing Rust (stable)
------------------------

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

Chose **option 1** to install

Once installation is complete, follow the instructions to configure your current terminal sessions' environment variable so that you can run Rust related commands:

**`source`` ``$HOME/.cargo/env`**

To confirm the instalation of the correct channel of Rust, type the following:

**`rustc`` ``--version`**

This should give you the output below:

`rustc 1.31.0 (abe02cefd 2018-12-04)`

Obtaining the source files
--------------------------

To download the source files from the Git repository, run the following command, (please make sure to check the exact path to the latest version) :

**`git`` ``clone`` `[`https://github.com/PoC-Consortium/engraver`](https://github.com/PoC-Consortium/engraver)**

This will download and extract the files to the engraver sub directory (from where your present working directory is)

Compiling the source files
--------------------------

Navigate to the engraver directory:

**`cd`` ``engraver`**

build the application:

**`cargo`` ``build`` ``--release`**

Running Engraver
----------------

Once the above has completed, navigate to the executable:

**`cd`` ``/target/release`**

From here you can run engraver and begin plotting. Example command:

**`./engraver`` ``--n`` ``1000`` ``--id`` ``12345678912345678912`` ``--path`` ``/home/burst/Burst`` ``--sn`` ``0`**

To understand the complete range of flags/options available for the engraver application, please read the [Wik](https://github.com/PoC-Consortium/engraver/wiki)i page on GitHub

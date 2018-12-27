| Burst Reference Software - MacOS Installation |
|-----------------------------------------------|
| **Status**                                    |

Table Of Contents
-----------------

\_\_TOC\_\_

Installation of MacOS Burst Wallet
----------------------------------

This tutorial will walk you through the steps required to run a Burst wallet locally on MacOS.

### Quick Start

Note: A simple bash script is available to complete all of the steps for you if you prefer using a script. Simply download a .zip of the code located **[here](https://github.com/beatsbears/macos_burst)** and follow the directions in the README file. If you prefer to execute each step without using a script, begin the process with the dependencies section.

### **Dependencies**

There are a few dependencies that you should check on prior to starting the install process.

1.  OSX 10.10 or higher
2.  Sufficient disk space for the full blockchain (~8Gb or so)
3.  The ability to leave your computer up and running for several hours in order to allow the full blockchain to synchronize.

### **Steps**

#### 1. Install Homebrew

[Homebrew](https://brew.sh/) (brew) is an excellent package manager for MacOS. To install, open a new terminal and run the following command. Note: You may be asked for your password in order to complete.

     $ /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

If you already have brew installed, it is a good practice to make sure everything is up to date by entering

    $ brew update

#### 2. Install MariaDB

MariaDB is an open-source fork of the MySQL relational database. It is used to store the blockchain and other wallet information. Install MariaDB by typing the following.

     $ brew install mariadb

#### 3. Start MariaDB and create a user for the wallet

    $ brew services start mariadb

Once MariaDB is running, log into MariaDB to create a new wallet database and add a user to manage it. Note: Your MariaDB server will likely not have a root password set. Consult google for instructions for setting a new root password.

    $ mysql -u root -p -h localhost

Once you have successfully logged into MariaDB, execute the following SQL commands.

    CREATE DATABASE burstwallet;
     
    CREATE USER 'burstwallet'@'localhost' IDENTIFIED BY '<YOUR PASSWORD>';
     
    GRANT ALL PRIVILEGES ON burstwallet.* TO 'burstwallet'@'localhost';

Assuming your queries have executed successfully you can exit the MariaDB console by typing `\q` and hitting Enter.

#### 4. Installing the Java 8 SDK

This step was a bit confusing as brew would always like to give you the latest version of all packages. I’m not entirely sure if the latest version of Java will work without issue, but I didn’t want to risk it. First you need to install brew cask.

    $ brew tap caskroom/cask

Next you can use brew cask to install a specific version of a package. So with cask installed you would run.

    $ brew cask install java8

#### 5. Download and setup the wallet

You can download the wallet package from PoC Consortium [here](https://github.com/PoC-Consortium/burstcoin/releases). Once you’ve unzipped the release (burstcoin-1.3.6.zip) you should open the directory and open the `nxt.properties` file in the `conf` directory in a text editor. At the bottom of the file add the following lines.

    nxt.dbUrl=jdbc:mariadb://localhost:3306/burstwallet

    nxt.dbUsername=burstwallet

    nxt.dbPassword=<YOUR PASSWORD>

These values are just telling the Java program where it can find and access the MariaDB database you set up in step 3.

#### 6. Start the wallet

Finally it’s time to start the wallet. In your open terminal window make sure you’re in the same directory as the burst.sh startup script. If you run `ls` in your terminal window you should see a response like this. <img src="1_jlRqRGWNhSOKcNEjWLvNLQ.png" title="fig:1_jlRqRGWNhSOKcNEjWLvNLQ.png" alt="1_jlRqRGWNhSOKcNEjWLvNLQ.png" width="649" height="649" />[1]

In your terminal window you should use the following command to make sure you have permission to execute the startup script.

`chmod +x burst.sh`

Finally you can launch the wallet script. Note: You need to leave this terminal window up and running while the wallet synchronizes the full block chain which could take several hours or days depending on your personal setup and internet connection.

`./burst.sh`

You should see a lot of output flying by pretty quickly as the wallet starts up.

#### 7. Creating your account and signing in.

Assuming you see no obvious errors while the wallet is starting up. You should check to see if it’s successfully running by entering [`http://localhost:8125/index.html`](http://localhost:8125/index.html) in your browser. You should see a page that looks like this. <img src="1_hkcu_dhZ5JUqrjbjiIcUrQ.png" title="fig:1_hkcu_dhZ5JUqrjbjiIcUrQ.png" alt="1_hkcu_dhZ5JUqrjbjiIcUrQ.png" width="880" height="880" />[2]

Click the ‘`New? Create Your Account!’` button. The wallet will generate a passphrase of 12 words. *WRITE THESE WORDS DOWN AND DO NOT SHARE THEM*. This passphrase is your **private key** and will be how you access your wallet and Burst funds. Please review the section on [securing your Burst](https://burstwiki.org/wiki/Secure_Your_Burst). Follow the confirmation step on the next page and click `Next` .

You should now see the wallet dashboard which looks like this. <img src="1_tlOAFgZwd5ayXb3YtHatNA.png" title="fig:1_tlOAFgZwd5ayXb3YtHatNA.png" alt="1_tlOAFgZwd5ayXb3YtHatNA.png" width="880" height="880" />[3]

Congratulations! You’re up and running.

#### 8. Wait for the blockchain to download

Now that you have your wallet running locally and have created an account, you’ll need to wait for the blockchain to synchronize before you see your up to date burst balance. This may take quite a while depending on your connection speed. Make sure you take note of your Account Id listed at the top and on the left side of the wallet. This address is how you will send/receive Burst in the future.

#### 9. References

10. From BRS Main
=================

MacOS Installation
------------------

In order to run the wallet locally, we need a database to store the blochchain information. The following guide shows the setup of MariaDB as database and the installation of the Burst Reference Software on MacOS.

### Prerequisites

-   [Java](http://www.oracle.com/technetwork/java/javase/downloads/jre9-downloads-3848532.html)
-   [Homebrew](https://brew.sh/)

### Setup MariaDB

MariaDB is used as database to store the blockchain. The following steps will guide you through the installation of MariaDB on MacOS.

First check that Xcode is up to date by typing the following and pressing enter.

`xcode-select --install `

After installing Xcode, install MariaDB by typing the following and pressing enter.

`brew install mariadb `

Once MariaDB is finished installing, start MariaDB with the following.

`brew services start mariadb `

### Database configuration

As soon as MariaDB is successfully installed, we can log into the MariaDB interface and setup database like explained in [Setup MariaDB](brs-macos-installation-setup-mariadb.md)

### Setup Wallet

1.  Now, download the lateste release from the Github release section
2.  Find the folder you downloaded from github and double click it. Archive utility will now extract it to a new folder.
3.  Rename the extracted folder to `burstcoin`.
4.  Open the `conf` folder inside the extracted folder
5.  Right click the file `brs-default.properties` and select “Open With → TextEdit”.
6.  Find the following lines and enter `burstwallet` as database user and your password (assigned beforehand) as database password brs.dbUsername=burstwallet brs.dbPassword=<your password>
7.  Move that folder to your home folder. This may appear as your username.
8.  Back in terminal, navigate to the `burstcoin` folder. cd ~/burstcoin
9.  Make `burst.sh` executable. chmod +x burst.sh
10. Now execute `burst.sh`. This is how you will start the wallet in the future. To execute this command you must be inside the burtscoin folder in the terminal. ./burst.sh
11. The wallet should now be running. In your web browser navigate to http://localhost:8125/ to access the wallet interface.

How to run a full node on MacOS?
--------------------------------

### Why Run a Full Node ?

Running a full node is one of the most important ways to help support the Burst network. This allows other peers to connect to your wallet and synchronize the blockchain if they are not using an imported DB. When light clients are released, Full Nodes maintain a copy of the blockchain, in the decentralized manner that crypto currencies are designed to run, while light clients will not. Below are the steps for setting up a full node.

### Network Settings

-   Once you have a local wallet installed and running you need to forward port 8123 to allow other peers to connect to your wallet.
-   Now you will set the machine to have a static IP address. This is so that if the machine restarts, it will not change IP addresses and negate the port forwarding rule we will set up later.

### Steps on MacOS

Firstly, if you have the latest BRS version, UPnP should be enabled and you already run a full node. Otherwise, following these steps :

1.  Open “System Preferences”
2.  Select “Network”
3.  Select “Advanced”
4.  Select “TCP/IP”
5.  Note down your “IPv4 Address”
6.  Note down the “Router” address, this is your default gateway.
7.  Leave “Configure IPv6” set to “Automatic”.
8.  Change “Configure IPv4” to “Manual”.
9.  Now in the “IPv4 Address” field, enter the IPv4 address you copied down before switching to manual.
10. Make sure “Subnet Mask” is still 255.255.255.0
11. Make sure “Router” still matches the router number you copied down before switching to manual.
12. Click “Ok”
13. Click “Apply”
14. Close “System Preferences”

### Router Settings

Now open your network router settings by entering your default gateway IP in the URL field of your browser. Login to your router.

-   If you have never used these settings before, most ISPs have default username/password combinations. A quick google search will help you find it.
-   If you do not remember your username/password combo, you will need to factory reset your router. Be mindful, this will force you to re-setup your entire network.

1.  Find the “Port Forwarding” option. Note, “Port Forwarding” and “Port Triggering” are **NOT** the same. “Port Forwarding” is generally under the “Advanced” tab.
2.  Select “Add Service”, “Add Rule”, or anything along those lines.
3.  Set “Common Service” as “Other”.
4.  Set service name to “Burstcoin”.
5.  Set “Service Type” to “TCP/UDP”.
6.  Set IPv4 Address to your machines address which is the address you just made static.
7.  Disregard the IPv6 Address field.
8.  Set your “Start Port” and “End Port” to “8123”.&gt;
9.  Save the new service.

The local wallet must remain running for your full node to be accessible to the network.

References
==========

<https://www.reddit.com/r/burstcoin/comments/7lrdc1/guide_to_getting_the_poc_wallet_running_on_a_mac/>

<https://medium.com/@aclaytonscott/burst-part-2-macos-wallet-setup-tutorial-2822bb029f54>

<references />

[1] <https://medium.com/@aclaytonscott/burst-part-2-macos-wallet-setup-tutorial-2822bb029f54>

[2] 

[3] 

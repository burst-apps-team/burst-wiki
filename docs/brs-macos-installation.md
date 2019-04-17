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

Assuming these commands have executed successfully, exit the MariaDB console by typing `\q` and hitting Enter.

#### 4. Install Java 8 SDK

Brew gives you the latest version of all packages, but the latest version of Java may not work without issue. You can install the latest verified version (Java 8) by following these steps. Install brew cask.

    $ brew tap caskroom/cask

Use brew cask to install a specific version of a package.

    $ brew cask install java8

#### 5. Download and setup the wallet

The wallet package from PoC Consortium is located [here](https://github.com/PoC-Consortium/burstcoin/releases). Unzip the release (burstcoin-1.3.6.zip) and open the `nxt.properties` file from the `conf` directory in a text editor. Add the following lines to the end of the file

    nxt.dbUrl=jdbc:mariadb://localhost:3306/burstwallet

    nxt.dbUsername=burstwallet

    nxt.dbPassword=<YOUR PASSWORD>

These values tell Java where to find the MariaDB database.

#### 6. Start the wallet

In your open terminal window verify that you are in the same directory as the burst.sh startup script. If you run `ls` in your terminal window you should see a response similar to the following: <img src="1_jlRqRGWNhSOKcNEjWLvNLQ.png" title="fig:1_jlRqRGWNhSOKcNEjWLvNLQ.png" alt="1_jlRqRGWNhSOKcNEjWLvNLQ.png" width="649" height="649" />[1]

In your terminal window, use the following command to make sure you have permission to execute the startup script.

`chmod +x burst.sh`

Launch the wallet script. Note: You will need to leave this terminal window running until the full blockchain has synchronized. This could take several hours or days depending on your processor speed and and internet connection.

`./burst.sh`

You should see a lot of output flying by pretty quickly as the wallet starts up.

#### 7. Creating your account and signing in.

Assuming you see no obvious errors while the wallet is starting up, check to see if it is running successfully by entering [`http://localhost:8125/index.html`](http://localhost:8125/index.html) in your browser. You should see a page that looks like this: <img src="1_hkcu_dhZ5JUqrjbjiIcUrQ.png" title="fig:1_hkcu_dhZ5JUqrjbjiIcUrQ.png" alt="1_hkcu_dhZ5JUqrjbjiIcUrQ.png" width="880" height="880" />[2]

Click the `New? Create Your Account!` button. The wallet will generate a passphrase of 12 words. *WRITE THESE WORDS DOWN AND DO NOT SHARE THEM*. This passphrase is your **private key** and will be how you access your wallet and Burst funds. Please review the section on [securing your Burst](https://burstwiki.org/wiki/Secure_Your_Burst). Follow the confirmation step on the next page and click `Next` .

You should now see the wallet dashboard. <img src="1_tlOAFgZwd5ayXb3YtHatNA.png" title="fig:1_tlOAFgZwd5ayXb3YtHatNA.png" alt="1_tlOAFgZwd5ayXb3YtHatNA.png" width="880" height="880" />[3]

If so, congratulations! You are up and running.

#### 8. Wait for the blockchain to download

Now that you have your wallet running locally and have created an account, you will need to wait for the blockchain to fully synchronize before you can see your up-to-date Burst balance. This may take a few hours or a few days depending on your processor and internet connection speed. Please note: Your Account Id is listed in the top left corner of your wallet. This address is necessary for sending and receiving Burst.

#### 9. References

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

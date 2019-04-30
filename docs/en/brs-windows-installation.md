Introduction
------------

This guide is for installing the latest BRS Core wallet from [https://www.burst-coin.org/download-wallet](https://www.burst-coin.org/download-wallet) on Windows PCs. If you would like a simpler wallet installation... please download [QBundle](Qbundle.md).

Installation of BRS only (with MariaDB)
---------------------------------------

1. Download the latest version of [BRS (zip)](https://github.com/burst-apps-team/burstcoin/releases) and extract it to where ever you like. In the conf directory, copy brs-default.properties into a new file named brs.properties.

2. Have the latest version of Java installed on your computer : <https://www.java.com/en/download/>.

3. Download and install MariaDB : <https://downloads.mariadb.org/>.

The MariaDb installation will ask to setup a password for the root user. Add this password to the `brs.properties` file created above in the following section:

`DB.Url=jdbc:mariadb://localhost:3306/`**`burstdatabase`**
`DB.Username=`**`root`**
`DB.Password=`**`YOUR_PASSWORD`**

### Setup MariaDB

The MariaDB installation will also install HeidiSQL, a gui tool to administer MariaDb. Use it to connect to the newly created MariaDb server and create a new DB called 'burstwallet'.

MariaDB knowledge base : <https://mariadb.com/kb/en/library/documentation/>

HeidiSQL documentation : <https://www.heidisql.com/help.php>

You can also directly connect to mysql and create database trough console :

`mysql -u root -p`

Now we create the database `burstwallet` for the blockchain.

`CREATE DATABASE burstwallet; `

In addition, we create a user which is used by the wallet to access the database. Replace <your password> with a password of your choice.

`CREATE USER 'burstwallet'@'localhost' IDENTIFIED BY '`<your password>`'; `

Finally we, grant this user all privileges for the database `burstwallet`.

`GRANT ALL PRIVILEGES ON burstwallet.* TO 'burstwallet'@'localhost';`

Finally, now that the database is created and the `brs.properties` configured, you only have to run `burst.cmd` in the burstcoin folder.

How to run a full node on Windows?
----------------------------------

### Why Run a Full Node ?

Running a full node is one of the most important ways to help support the Burst network. This allows other peers to connect to your wallet and synchronize the blockchain if they are not using an imported DB. When light clients are released, Full Nodes maintain a copy of the blockchain, in the decentralized manner that crypto currencies are designed to run, while light clients will not. Below are the steps for setting up a full node.

### Steps on Windows

Firstly, if you have the latest BRS version, UPnP should be enabled and you already run a full node. Otherwise, following these steps :

1.  Once you have a local wallet installed and running you need to forward port 8123 to allow other peers to connect to your wallet.
2.  We now need to get some IP addresses.
    -   Open “Command Prompt”.
    -   Type “ipconfig”
    -   Find and record your “Default Gateway” and “IPv4 Address”
3.  Now you will set the machine to have a static IP address. This is so that if the machine restarts, it will not change IP addresses and negate the port forwarding rule we will set up later.
    -   Open “Control Panel”.
    -   Open “Network and Internet”.
    -   Open “Network and Charing Center”.
    -   In the left menu, open “Change Adapter Settings”.
    -   Right click your preferred method of internet connection.
    -   Select “Properties”.
    -   Uncheck “Internet Protocol Version 6 (TCP/IPv6)”.
    -   Click “Internet Protocol Version 4 (TCP/IPv4)”.
    -   Open “Properties”
    -   Select “Use the following IP address”
    -   In “IP address” type the IPv4 number that you got in Command Prompt.
    -   In “Subnet Mask” type “255.255.255.0”.
    -   In “Default Gateway” type the default gateway address you got in Command Prompt.
    -   Select “Use the following DNS server addresses”
    -   In “Preferred DNS server” type “8.8.8.8”
    -   In “Alternate DNS server” type “8.8.4.4”
    -   Check “Validate settings upon exit”
    -   Click “OK”
    -   If you receive a warning about “Multiple default gateways...” select “Yes”.
    -   Click “Close”
4.  Now open your network router settings by entering your default gateway IP in the URL field of your browser.
5.  Login to your router.
    -   If you have never used these settings before, most ISPs have default username/password combinations. A quick google search will help you find it.
    -   If you do not remember your username/password combo, you will need to factory reset your router. Be mindful, this will force you to re-setup your entire network.
6.  Find the “Port Forwarding” option. Note, “Port Forwarding” and “Port Triggering” are **NOT** the same. “Port Forwarding” is generally under the “Advanced” tab.
    -   Select “Add Service”, “Add Rule”, or anything along those lines.
    -   Set “Common Service” as “Other”.
    -   Set service name to “Burstcoin”.
    -   Set “Service Type” to “TCP/UDP”.
    -   Set IPv4 Address to your machines address which is the address you just made static.
    -   Disregard the IPv6 Address field.
    -   Set your “Start Port” and “End Port” to “8123”.
    -   Save the new service.
7.  Finally, you must reserve the IP address for your specific machine
    -   Navigate to “Connected Devices” (or something of the like).
    -   Find and click-on the machine you are running the node from.
    -   Select “Edit”
    -   Change configuration to “Reserved IP”
    -   Verify the IP is the same one you set in port forwarding.
    -   Select Save.
8.  You are now finished and can close out network settings.

The local wallet must remain running for your full node to be accessible to the network. After a few hours to a few days the [network observer](https://explore.burst.cryptoguru.org/tool/observe) will update and you can check that your set up worked. First find [your external IP](https://www.myexternalip.com/) address and then search your external IP address and port in the network observer (ex. 193.182.13.162:8123).

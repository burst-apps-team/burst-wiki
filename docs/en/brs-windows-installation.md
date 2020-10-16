Introduction
------------

This guide is for installing the latest BRS node from [burst-apps-team releases](https://github.com/burst-apps-team/burstcoin/releases) on Windows PCs. If you would like you can also check [QBundle](Qbundle.md).

Installation of BRS with H2 (default, simpler)
----------------------------------------------

1. Download the latest version of [BRS (zip)](https://github.com/burst-apps-team/burstcoin/releases) and extract it to where ever you like. In the conf directory, copy brs-default.properties into a new file named brs.properties.

2. Make sure you have a 64 bit Java installed on your computer (Java 8 recommended) : <https://www.java.com/en/download/>.

3. Run the burst.exe file.

Installation of BRS with MariaDB (optional, more complex)
---------------------------------------------------------

For those having issues with the H2 database backend, check if you have a 64 bits Java version 8 installed. If that does not solve your problem, use MariaDB as your backend.

1. Download the latest version of [BRS (zip)](https://github.com/burst-apps-team/burstcoin/releases) and extract it to where ever you like. In the conf directory, copy brs-default.properties into a new file named brs.properties.

2. Have the latest version of Java installed on your computer : <https://www.java.com/en/download/>.

3. Download and install MariaDB : <https://downloads.mariadb.org/>.

The MariaDb installation will ask to setup a password for the root user. Add this **`YOUR_MARIADB_PASSWORD`** password to the `brs.properties` file created above in the following section:

`DB.Url=jdbc:mariadb://localhost:3306/`**`burstwallet`**<br>
`DB.Username=`**`root`**<br>
`DB.Password=`**`YOUR_MARIADB_PASSWORD`**<br>

### Setup MariaDB

The MariaDB installation will also install HeidiSQL, a gui tool to administer MariaDb. Use it to connect to the newly created MariaDb server and create a new DB called 'burstwallet'.

MariaDB knowledge base : <https://mariadb.com/kb/en/library/documentation/>

HeidiSQL documentation : <https://www.heidisql.com/help.php>

You can also directly connect to mysql and create database trough console :

`mysql -u root -p`

Now we create the database `burstwallet` for the blockchain.

`CREATE DATABASE burstwallet; `

In addition, we create a user which is used by the wallet to access the database. Replace **`YOUR_PASSWORD`** with a password of your choice.

`CREATE USER 'burstwallet'@'localhost' IDENTIFIED BY '`**`YOUR_PASSWORD`**`'; `

Finally we, grant this user all privileges for the database `burstwallet`.

`GRANT ALL PRIVILEGES ON burstwallet.* TO 'burstwallet'@'localhost';`

Finally, now that the database is created and the `brs.properties` configured, you only have to run `burst.cmd` in the burstcoin folder.

### Setup Portable MariaDB

1.  Download latest ***mariadb-10.5.6-winx64.zip*** from https://downloads.mariadb.org/mariadb/+releases/

2.  Create ***burstcoin-2.5.3\mariadb*** folder

3.  Unzip ***mariadb-10.5.6-winx64.zip*** to ***burstcoin-2.5.3\mariadb*** folder

4.  Double click ***burstcoin-2.5.3\mariadb\bin\mariadb-install-db.exe***

5.  Open CMD as admin

6.  Run CMD command: `start burstcoin-2.5.3\mariadb\bin\mariadbd.exe --no-defaults`

7.  Run CMD command: `burstcoin-2.5.3\mariadb\bin\mysql.exe -u root -p`

8.  Hit enter. MariaDB has no root password by default.

9.  Run DB command: `CREATE DATABASE burstwallet;`

10. Run DB command: `CREATE USER 'burstwallet'@'localhost' IDENTIFIED BY 'burstwallet';`

11. Run DB command: `GRANT ALL PRIVILEGES ON burstwallet.* TO 'burstwallet'@'localhost';`

12. Edit ***burstcoin-2.5.3\conf\brs-default.properties*** file

>`DB.Url=jdbc:mariadb://localhost:3306/burstwallet`<br>
>`DB.Username=burstwallet`<br>
>`DB.Password=burstwallet`<br>

>`#DB.Url=jdbc:h2:file:./burst_db/burst.h2;DB_CLOSE_ON_EXIT=FALSE`<br>
>`#DB.Username=`<br>
>`#DB.Password=`<br>

13. Now your Burst blockchain data will store inside `burstcoin-2.5.3\mariadb\data` folder in a portable way

14. Create ***burstcoin-2.5.3\startburst.bat*** file for running BRS

>`start .\mariadb\bin\mariadbd.exe --no-defaults` <br>
>`timeout 10` <br>
>`start .\burst.jar` <br>

12. Create ***burstcoin-2.5.3\stopburst.bat*** file for shutdown BRS and MariaDB

>`taskkill /f /fi "windowtitle eq Burst Reference Software version v2.5.3"` <br>
>`timeout 10` <br>
>`taskkill /f /fi "windowtitle eq Burst Reference Software version v2.5.3 (STOPPED)"` <br>
>`timeout 10` <br>
>`start .\mariadb\bin\mariadb-admin.exe -u root shutdown` <br>

13. Run ***startburst.bat***

14. Open http://localhost:8125/index.html in browser

15. Enjoy

16. Run ***stopburst.bat*** to shutdown

Make sure you have a full node
------------------------------

### Why Run a Full Node ?

Running a full node is one of the most important ways to help support the Burst network. This allows other peers to connect to your node and synchronize the blockchain if they are not using an imported DB. When light clients are released, Full Nodes maintain a copy of the blockchain, in the decentralized manner that crypto currencies are designed to run, while light clients will not. Below are the steps for setting up a full node.

### Check if your node is acting as a full node

if you have the latest BRS version, UPnP should be enabled and you probably already run a full node.
BRS must remain running for your full node to be accessible to the network. After a few hours the [explorer network status](https://explorer.burstcoin.network/?action=network_status) will update and you can check that your set up worked. First find [your external IP](https://www.myexternalip.com/) address and then search your external IP address in the [explorer network status](https://explorer.burstcoin.network/?action=network_status) (ex. 193.182.13.162:8123).

### In case the UPnP didn't work

If the automatic UPnP setup didn't work for you, use the following steps :

1.  Once you have your node running you need to forward port 8123 to allow other peers to connect to your wallet.
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

After a few hours check again if the [explorer network status](https://explorer.burstcoin.network/?action=network_status) lists your node.

### Configure IPv6 full node

#### To be able to run BRS node through IPv6 protocol

- Your network should be IPv6 compatible which means your ISP allocate some dynamic or static IPv6 address to your host.

- You need to allowe IPv6 protocol through network settings.

Change ***burstcoin-2.5.3\conf\brs-default.properties*** file

>`API.allowed = *`<br>
>`P2P.Listen = [::]`<br>
>`API.Listen = [::]`<br>
>`API.V2.Listen = [::]`<br>

For NDS-A you need to configure

>`P2P.shareMyAddress = yes`<br>
>`P2P.myAddress = [YOUR_IPv6_PUBLIC_ADDRESS]`<br>
>`P2P.myPlatform = YOUR_BURST_ADDRESS`<br>
>`P2P.Port = 8123`

After the right configuration you are able to connect your BRS node fom anywhere typing ***http://[YOUR_IPv6_PUBLIC_ADDRESS]:8125/index.html*** to yourt browser.

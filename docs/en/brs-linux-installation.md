Linux Install Video
-----------------
<iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/dwWGejP0ZC8" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

Debian 9 “Stretch” installation
-------------------------------

The following steps are to install BRS Core Wallet on a fresh installation of **Debian Linux version 9 “Stretch”**. Please be sure to check Git Repository for the latest release before proceeding.

**Note:** Commands to be entered into the shell will be formatted in bold text, e.g.:

**`date`**

example command outputs will not be bold, e.g.:

`Wed 28 Nov 12:02:18 GMT 2018`

### Prerequisites

Before you begin the installation, the following packages need to be installed:

-   java 8 jdk
-   curl
-   dirmngr

Make sure your installation is up to date

**`apt update && apt upgrade`**

Run the following command to install the above packages:

**`apt install curl default-jdk dirmngr -y`**

### Configuring Java environment path

Find the path to the jdk

**`update-alternatives --config javac`**

If only one version of Java is installed, this will give you a similar output to the following:

`There is only one alternative in link group javac (providing /usr/bin/javac): /usr/lib/jvm/java-8-openjdk-amd64/bin/javac`
`Nothing to configure.`

Copy the path from the last output and edit the environment file:

**`nano /etc/environment`**

Add the following line to the file:

**`JAVA_HOME="/usr/lib/jvm/java-8-openjdk-amd64/bin/javac"`** (taken from the path above)

To make the above change effective in your current shell, run the following:

**`source /etc/environment`**

Now test that the JAVA\_HOME variable is part of PATH

**`echo $JAVA_HOME`**

This will output the path you entered into the environment file. Now confirm the version of java installed

**`javac -version`**

### Installing MariaDB

Add the MariaDB signing key

**`apt-key adv --recv-keys --keyserver hkp://keyserver.ubuntu.com:80 0xF1656F24C74CD1D8`**

Grab the and run MariaDB setup script

**`curl -sS https://downloads.mariadb.com/MariaDB/mariadb_repo_setup | bash`**

Update your sources list:

**`apt update`**

Install MariaDB:

**`apt install mariadb-server`**

During installation you will be prompted to enter a password for the root user. Enter and remember this as we will need it later.

Create a database(**brs\_master**), user (**brs\_user**) and assign a password to the user by running the following command:

**```echo "CREATE DATABASE brs_master; 
CREATE USER 'brs_user'@'localhost' IDENTIFIED BY 'yourpassword';
GRANT ALL PRIVILEGES ON brs_master.* TO 'brs_user'@'localhost';" | mysql -uroot -p
mysql -uroot -p < init-mysql.sql```**

This will prompt you for the password you entered for the root user when installing mariaDB

### Configuring the BRS Wallet

Grab the BRSWallet zip from [GitHub](https://github.com/burst-apps-team/burstcoin/releases) (always be sure to check link to the latest version. (2.4.2 was latest at time of writing)

`wget https://github.com/burst-apps-team/burstcoin/releases/download/v2.4.2/burstcoin-2.4.2.zip`

Create a directory to unzip to:

**`mkdir BRSWallet`**

Unzip the burstcoin zip file to the newly created directory:

**`unzip burstcoin-2.4.2.zip -d BRSWallet/`**

Navigate to the conf subdirectory within the newly created directory:

**`cd BRSWallet/conf`**

Create a brs.properties file inside the conf folder:

**`nano brs.properties`**

Add the following lines:

**`DB.Url=jdbc:mariadb://localhost:3306/brs_master`**
**`DB.Username=brs_user`**
**`DB.Password=yourpassword`**

The **burst.sh** file is not yet executable. We need to change the permissions to make it so. Drop back to the BRSWallet directory **`cd`` ``..`** and issue the following:

**`chmod 750 burst.sh`**

Now when we list the contents of the directory we can see that the **burst.sh** file has the correct permissions:

**`ls -l`**

```
total 1399340
-rw-r--r-- 1 root root 1409751247 Nov  5 12:29 brs.mariadb.zip
-rw-r--r-- 1 root root        507 Nov 18 22:17 burst.cmd
drwxr-xr-x 2 root root       4096 Nov 18 22:16 burst_db
-rw-r--r-- 1 root root   23015701 Nov 18 22:17 burst.jar
-rwxr-x--- 1 root root      17831 Nov 18 22:17 burst.sh
-rw-r--r-- 1 root root         46 Nov 18 22:17 Burst_Wallet.url
drwxr-xr-x 2 root root       4096 Nov 29 15:44 conf
-rw-r--r-- 1 root root      29288 Nov 18 22:17 genscoop.cl
drwxr-xr-x 3 root root       4096 Nov 18 22:17 html
-rw-r--r-- 1 root root      23232 Nov 18 22:17 init-mysql.sql
drwxr-xr-x 3 root root       4096 Nov 18 22:17 lib
-rw-r--r-- 1 root root      35149 Nov 18 22:17 LICENSE.txt
-rw-r--r-- 1 root root       4574 Nov 18 22:17 README.md
```

### Importing the database

The quickest way to get up an running is to bootstrap the database by importing a relatively up to date dump of the database. Your wallet will synchronise from this point in time automatically. If you prefer to synchronise from scratch this will take much longer to complete.

**`./burst.sh import mariadb`**

If you already have a/part of a brs\_master database present, you will be asked if you want to remove the current database, download and import a new one. Answer y or n.

`Do you want to remove the current databases, download and import new one? `**`y`**

If you created the **brs.properties** file as per above you will be asked the following question:

`Detected a previous database config(conf/brs.properties). Do you want to use the config?`

Answer: **`y`**

Answer the questions when prompted. If you are running the wallet locally on the same device that you have your terminal open to and you used the same database (**brs\_master**) and username (**brs\_user**) above when creating the database, you can simply press enter for the first three questions. For the fourth question, enter the password that you input when creating your database and user.

**`./burst.sh`` ``import`` ``mariadb`**
`Do you want to remove the current databases, download and import new one? `**`y`**
`Detected a previous database config(conf/brs.properties). Do you want to use the config? `**`y`**
`Could not find Database Password`
`Please enter your connection details`
`Host     (localhost) :`
`Database (brs_master):`
`Username (brs_user)  :`
`Password empty       : `**<Enter your password here>**
`--2018-11-29 15:44:47--  `[`https://download.cryptoguru.org/burst/wallet/brs.mariadb.zip`](https://download.cryptoguru.org/burst/wallet/brs.mariadb.zip)
`Resolving download.cryptoguru.org (download.cryptoguru.org)... 195.201.117.34`
`Connecting to download.cryptoguru.org (download.cryptoguru.org)|195.201.117.34|:443... connected.`
`HTTP request sent, awaiting response... 200 OK`
`Length: 1409751247 (1.3G) [application/zip]`
`Saving to: ‘brs.mariadb.zip’`
`brs.mariadb.zip                 100%[====================================================>]   1.31G  1.87MB/s    in 12m 4s`
`2018-11-29 15:56:52 (1.86 MB/s) - ‘brs.mariadb.zip’ saved [1409751247/1409751247]`
`Archive:  brs.mariadb.zip`
` inflating: brs.mariadb.sql`
`[+] Import successful - please remove brs.mariadb.zip`

Remove the download DB dump zip file to reclaim space and avoid an outdated zip file in future if you require a DB bootstrap

**`rm brs.mariadb.zip`**

### Running the Wallet

We need to change owner from root to a normal user as all the files/directories are now owned by root (see below)

```
total 1399340
-rw-r--r-- 1 root root 1409751247 Nov  5 12:29 brs.mariadb.zip
-rw-r--r-- 1 root root        507 Nov 18 22:17 burst.cmd
drwxr-xr-x 2 root root       4096 Nov 18 22:16 burst_db
-rw-r--r-- 1 root root   23015701 Nov 18 22:17 burst.jar
-rwxr-x--- 1 root root      17831 Nov 18 22:17 burst.sh
-rw-r--r-- 1 root root         46 Nov 18 22:17 Burst_Wallet.url
drwxr-xr-x 2 root root       4096 Nov 29 15:44 conf
-rw-r--r-- 1 root root      29288 Nov 18 22:17 genscoop.cl
drwxr-xr-x 3 root root       4096 Nov 18 22:17 html
-rw-r--r-- 1 root root      23232 Nov 18 22:17 init-mysql.sql
drwxr-xr-x 3 root root       4096 Nov 18 22:17 lib
-rw-r--r-- 1 root root      35149 Nov 18 22:17 LICENSE.txt
-rw-r--r-- 1 root root       4574 Nov 18 22:17 README.md
```

Navigate to the BRSWallet parent directory

**`cd ..`**

Change ownership of all files and directores within the BRSWallet directory:

**`chown -R burst:burst BRSWallet/`**

Now the ownership has ben changed to the user burst:

```
total 1399344
-rw-r--r-- 1 burst burst       2506 Nov 29 16:59 brs.log
-rw-r--r-- 1 burst burst 1409751247 Nov  5 12:29 brs.mariadb.zip
-rw-r--r-- 1 burst burst        507 Nov 18 22:17 burst.cmd
drwxr-xr-x 2 burst burst       4096 Nov 18 22:16 burst_db
-rw-r--r-- 1 burst burst   23015701 Nov 18 22:17 burst.jar
-rwxr-x--- 1 burst burst      17831 Nov 18 22:17 burst.sh
-rw-r--r-- 1 burst burst         46 Nov 18 22:17 Burst_Wallet.url
drwxr-xr-x 2 burst burst       4096 Nov 30 10:39 conf
-rw-r--r-- 1 burst burst      29288 Nov 18 22:17 genscoop.cl
drwxr-xr-x 3 burst burst       4096 Nov 18 22:17 html
-rw-r--r-- 1 burst burst      23232 Nov 18 22:17 init-mysql.sql
drwxr-xr-x 3 burst burst       4096 Nov 18 22:17 lib
-rw-r--r-- 1 burst burst      35149 Nov 18 22:17 LICENSE.txt
-rw-r--r-- 1 burst burst       4574 Nov 18 22:17 README.md
```

At this stage you can reboot the machine to ensure that the environmental variables are applied to every user session going forward. Following the reboot, log in as the your normal user (burst in this case) and run the wallet:

**`./burst.sh`**

You will see the following output:

```
Nov 29, 2018 4:28:33 PM brs.Burst loadProperties
INFO: Initializing Burst Reference Software (BRS) version 2.3.0
Nov 29, 2018 4:28:33 PM org.ehcache.core.EhcacheManager createCache
INFO: Cache 'account' created in EhcacheManager.
[INFO] 2018-11-29 16:28:33 brs.db.sql.Db - Using mariadb Backend
[INFO] 2018-11-29 16:28:33 brs.db.mariadb.MariadbDbVersion - Database update may take a while if needed, current db version 177...
```

**Note:** If closing the terminal or pressing Ctrl-c, the wallet will shutdown:

`[INFO] 2018-11-29 16:53:12 brs.Burst - Shutting down...`

### Configure BRS to run as a service/daemon

Closing the terminal or pressing Ctrl-c will result to a shutdown of the wallet. To avoid this, BRS should be run as a service. It is particularly helpful if you run BRS on a linux system using SSH.

First, if systemd isn't installed on your system, run the following command:

**`sudo apt install systemd`**

Then, you will have to create a new service file for BRS:

**`sudo nano /etc/systemd/system/brs.service`**

Add the following configuration (replace description, Exec, WorkingDirectory and User according to your setup):

```
[Unit]
Description=BRS Wallet
After=network.target
Wants=network.target
[Service]
ExecStart=/usr/bin/java -cp /home/burst/burst.jar:/home/burst/conf brs.Burst
WorkingDirectory=/home/burst/
User=brs-mariadb
Restart=always
RestartSec=90
[Install]
WantedBy=multi-user.target
```

After that, enable the service:

**`sudo systemctl enable brs.service`**

To start the service, type:

**`systemctl start brs.service`**

To get the status of the service:

**`systemctl status brs.service`**

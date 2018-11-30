| Burst Reference Software - Linux Installation |
|-----------------------------------------------|
| **Status**                                    |
| **Credits**                                   |

Table Of Contents
-----------------

\_\_TOC\_\_

### Debian & Ubuntu

Download the latest version of BRS here (zip) : <https://github.com/PoC-Consortium/burstcoin/releases>.

If you are running BRS for the first time,

-   install Java (https://www.java.com/en/download/help/linux\_install.xml)
-   install MariaDB (https://www.linode.com/docs/databases/mariadb/mariadb-setup-debian/\#install-mariadb)
-   run `./burst.sh help`

If you want to import the blockchain now, run `./burst.sh import mariadb`.

### Upgrading your wallet config from 1.3.6cg

`./burst.sh upgrade`

will take the old `nxt-default.properties`/`nxt.properties` files and create `brs-default.properties.converted`/`brs.properties.converted` files in the conf directory. This should give you a headstart with the new option naming system.

------------------------------------------------------------------------

### Step-by-Step Guide

The following steps are to install Scavenger 1.6.4 by compiling from source on a fresh installation of **Debian Linux version 9 “Stretch”**. Please be sure to check Git Repository for the latest release before proceeding.

**Note:** Commands to be entered into the shell will be formatted in bold text, e.g.:

**`date`**

example command outputs will not be bold, e.g.:

`Wed 28 Nov 12:02:18 GMT 2018`

Make sure your installation is up to date

`apt update && apt upgrade`
` `
` Prerequisites`
` `
` java 9 jdk`
` curl`
` dirmgr`
` apt install curl default-jdk dirmngr`
` `
` Configuring Java environment path`
`   `
`    update-alternatives --config javac`
`    vi /etc/environment`
`  `
`  add line JAVA_HOME=`“`/usr/lib/jvm/java-8-openjdk-amd64/bin/javac`”
`  `
`    source /etc/environment`
`    echo $JAVA_HOME`
`    javac -version`
`  `
`  `
`  Installing MariaDB`
`  `
`  `
`  Add the MariaDB signing key`
`  `
`    apt-key adv --recv-keys --keyserver hkp://keyserver.ubuntu.com:80 0xF1656F24C74CD1D8`
`  `
`  Grab the and run MariaDB setup script `
`    curl -sS `[`https://downloads.mariadb.com/MariaDB/mariadb_repo_setup`](https://downloads.mariadb.com/MariaDB/mariadb_repo_setup)` | bash`
`  `
`    update your sources list`
`  apt update`
`  `
`  Install MariaDB`
`    apt install mariadb-server`
`  `
`  During installation you will be prompted to enter a password for the root user. Enter and remember this as we will need it later.`
`  `
`  Create a database(brs_master), user (brs_user) and assign a password to the user by running the following command:`
`     echo "CREATE DATABASE brs_master; `
`     CREATE USER 'brs_user'@'localhost' IDENTIFIED BY 'yourpassword';`
`     GRANT ALL PRIVILEGES ON brs_master.* TO 'brs_user'@'localhost';" | mysql -uroot -p`
`     mysql -uroot < init-mysql.sql`

This will prompt you for the password you entered for the root user when installing mariaDB

`  Grab the BRSWallet zip from GitHub (always be sure to check link to the latest version. (2.2.5 was latest at time of writing)`
`     wget `[`https://github.com/PoC-Consortium/burstcoin/releases/download/2.2.5/burstcoin-2.2.5.zip`](https://github.com/PoC-Consortium/burstcoin/releases/download/2.2.5/burstcoin-2.2.5.zip)
`  `
`  create a directory to unzip to`
`    mkdir BRSWallet`
`  `
`  unzip the burstcoin zip file to the newly created directory`
`    unzip burstcoin-2.2.5.zip -d BRSWallet/`
`  `
`  navigate to the new directory`
`    cd BRSWallet/`
`  `
`  navigate to the conf directory `
`     cd conf`
`  `
`  make a copy of brs-default.properties and name it brs.properties`
`     cp brs-default.properties brs.properties`
`     `
`  edit brs.properties`
`     vi brs.properties`
`     `
`     *notes to be added for which line to edit*`
`  `
`  Note here how the burst.sh file is not executable. `
`  ls -l `
`  We need to change the permissions to make it so`
`  53  ls -l`
`  `
`  `
`    chmod 750 burst.sh`
`  `
`  now the burst.sh file has the correct permissions`
`    ls -l`
`  `

Importing the database

The quickest way to get up an running is to bootstrap the database by importing a relatively up to date dump of the database. Your wallet will synchronise from this point in time automatically. If you prefer to synchronise from scratch this will take much longer to complete.

`    ./burst.sh import mariadb`
`  `
`  if you already have a/part of a brs_master database present, you will be asked if you want to remove the current database, download and import a new one. Answer y or n.`
`  Do you want to remove the current databases, download and import new one? y`

If you created the brs.properties file as per above you will be asked the following question.

`  Detected a previous database config(conf/brs.properties). Do you want to use the config? `

Answer y

`  Answer the questions when prompted. If you are running the wallet locally on the same device that you have your terminal open to and you used the same database (brs_master) and username (brs_user) above when creating the database, you can simply press enter for the first three questions. For the fourth question, enter the password that you input when creating your database and user.`
`  `
`     `
`  `

./burst.sh import mariadb Do you want to remove the current databases, download and import new one? y Detected a previous database config(conf/brs.properties). Do you want to use the config? y Could not find Database Password

Please enter your connection details Host (localhost) : Database (brs\_master): Username (brs\_user) : Password empty : --2018-11-29 15:44:47-- <https://download.cryptoguru.org/burst/wallet/brs.mariadb.zip> Resolving download.cryptoguru.org (download.cryptoguru.org)... 195.201.117.34 Connecting to download.cryptoguru.org (download.cryptoguru.org)|195.201.117.34|:443... connected. HTTP request sent, awaiting response... 200 OK Length: 1409751247 (1.3G) \[application/zip\] Saving to: ‘brs.mariadb.zip’

brs.mariadb.zip 100%\[====================================================&gt;\] 1.31G 1.87MB/s in 12m 4s

2018-11-29 15:56:52 (1.86 MB/s) - ‘brs.mariadb.zip’ saved \[1409751247/1409751247\]

Archive: brs.mariadb.zip

` inflating: brs.mariadb.sql`

\[+\] Import successful - please remove brs.mariadb.zip root@BurstWallet:/home/burst/BRSWallet\#

` remove the download DB dump zip file to reclaim space and avoid an outdated zip file in future if you require a DB bootstrap`
` `
` rm brs.mariadb.zip`
` `
` `
` we need to change owner from root to a normal user`
` `
`  ls -l`
`  `

total 1399340 -rw-r--r-- 1 root root 1409751247 Nov 5 12:29 brs.mariadb.zip -rw-r--r-- 1 root root 507 Nov 18 22:17 burst.cmd drwxr-xr-x 2 root root 4096 Nov 18 22:16 burst\_db -rw-r--r-- 1 root root 23015701 Nov 18 22:17 burst.jar -rwxr-x--- 1 root root 17831 Nov 18 22:17 burst.sh -rw-r--r-- 1 root root 46 Nov 18 22:17 Burst\_Wallet.url drwxr-xr-x 2 root root 4096 Nov 29 15:44 conf -rw-r--r-- 1 root root 29288 Nov 18 22:17 genscoop.cl drwxr-xr-x 3 root root 4096 Nov 18 22:17 html -rw-r--r-- 1 root root 23232 Nov 18 22:17 init-mysql.sql drwxr-xr-x 3 root root 4096 Nov 18 22:17 lib -rw-r--r-- 1 root root 35149 Nov 18 22:17 LICENSE.txt -rw-r--r-- 1 root root 4574 Nov 18 22:17 README.md

`     cd ..`
`     chown -R burst:burst BRSWallet/`

logout from root

`     exit`

`run wallet  `

`     ./burst.sh`

You will see the following output

Nov 29, 2018 4:28:33 PM brs.Burst loadProperties INFO: Initializing Burst Reference Software (BRS) version 2.2.5 Nov 29, 2018 4:28:33 PM org.ehcache.core.EhcacheManager createCache INFO: Cache 'account' created in EhcacheManager. \[INFO\] 2018-11-29 16:28:33 brs.db.sql.Db - Using mariadb Backend \[INFO\] 2018-11-29 16:28:33 brs.db.mariadb.MariadbDbVersion - Database update may take a while if needed, current db version 177...

If closing the terminal or pressing Ctrl-c, the wallet will shutdown

\[INFO\] 2018-11-29 16:53:12 brs.Burst - Shutting down...

To avoid this, the burst scrips should be run as a Daemon (steps below) \*to be written completely\*

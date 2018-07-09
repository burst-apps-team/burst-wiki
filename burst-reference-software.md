| Burst Refereence Software |
|---------------------------|
| **Status**                |

Table Of Contents
-----------------

\_\_TOC\_\_

Introduction
------------

The Burst Reference Software (BRS) is the main burst wallet. It allows connection to the Burst network.

Windows Installation
--------------------

All the details regarding the installation of BRS can be found here: [QBundle](qbundle.md). Indeed, QBundle is a launcher allowing you to install the Burst local wallet in a few clicks.

Linux Installation
------------------

Packages are provided for Debian (stretch, jessie) and Ubuntu (zesty, xenial)

### Debian

#### **Packages**

<https://github.com/PoC-Consortium/burstcoin>

#### Quick Install

Copy & Paste this line into your terminal:

`\wget -q -O - http://package.cryptoguru.org/debian/init.sh | bash `

and now you can install packages like:

`apt-get install burstcoincg`

#### Shell Script

The Quick- Install downloads a shell script and executes it directly. If you're really concerned about the argument that it may contain nefarious activities within, you can easily review it before you run it.

`wget http://package.cryptoguru.org/debian/init.sh less init.sh bash init.sh `

and now you can install packages like:

`apt-get install burstcoincg `

#### Add to sources

-   **stretch** Add the following stuff to your sources.list deb https://package.cryptoguru.org/debian/stretch stretch main sometimes you will need to play with package pining
-   **jessie** Add the following stuff to your sources.list deb https://package.cryptoguru.org/debian/jessie jessie main sometimes you will need to play with package pining and you may have to add `deb http://ftp.debian.org/debian jessie-backports main`

### Ubuntu

#### Packages

-   burstcoincg - Burstcoin wallet powered by CryptoGuru based on MariaDB.
-   creepminer - The creepMiner is a client application for mining Burst on a pool or solo.

#### Quick Install

Copy & Paste this line into your terminal:

`\wget -q -O - http://package.cryptoguru.org/ubuntu/init.sh | bash `

and now you can install packages like:

`apt-get install burstcoincg `

Info

The wallet is started automatically after installation.

#### Shell Script

The Quick- Install downloads a shell script and executes it directly. If you're really concerned about the argument that it may contain nefarious activities within, you can easily review it before you run it.

`wget http://package.cryptoguru.org/ubuntu/init.sh less init.sh bash init.sh `

and now you can install packages like:

`apt-get install burstcoincg `

#### Add to sources

-   **zesty** Add the following stuff to your sources.list deb https://package.cryptoguru.org/ubuntu/zesty zesty main sometimes you will need to play with package pining
-   **xenial** Add the following stuff to your sources.list deb https://package.cryptoguru.org/ubuntu/xenial xenial main sometimes you will need to play with package pining

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

### Database configuration¶

As soon as MariaDB is successfully installed, we can log into the MariaDB interface.

`mysql -u root -p -h localhost`

Now we create the database `burstwallet` for the blockchain.

`CREATE DATABASE burstwallet; `

In addition, we create a user which is used by the wallet to access the database. Replace <your password> with a password of your choice.

`CREATE USER 'burstwallet'@'localhost' IDENTIFIED BY '`<your password>`'; `

Finally we, grant this user all privileges for the database `burstwallet`.

`GRANT ALL PRIVILEGES ON burstwallet.* TO 'burstwallet'@'localhost';`

### Setup Wallet¶

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

Docker Installation
-------------------

The Burst Reference Software can be run with Docker and Docker-Compose. Since Docker supports all major platforms, this can be an easy-to-setup alternative to a platform-dependent installation of the Burst Reference Software.

### Prerequisites

[Docker CE](https://docs.docker.com/engine/installation/)

[Docker Compose](https://docs.docker.com/compose/install/)

### Database

Currently, Docker images for BRS version 2.2 with MariaDB or H2 Database are available at Docker Hub. You have to choose one of them, which is used to store the blockchain.

### MariaDB

The following `docker-compose.yml` file can be used to run the BRS with MariaDB database in containers.

    version: '3'

    services:
      burstcoin:
        image: burstcoin/core:2.2-mariadb
        restart: always
        depends_on:
         - mariadb
        ports:
         - 8123:8123
         - 8125:8125
      mariadb:
        image: mariadb:10
        environment:
         - MYSQL_ROOT_PASSWORD=burst
         - MYSQL_DATABASE=burst
        command: mysqld --character_set_server=utf8mb4
        volumes:
         - ./burst_db:/var/lib/mysql

Save the file and run it.

`docker-compose up -d `

With the `-d` flag, both containers are started as background processes. A `burst_db` folder will be created in the directory of the `docker-compose.yml` file, if it was not created before. This folder is mounted to the MariaDB storage and holds the `burst` database containing the blockchain. The containers can be stopped and removed at any time via `docker-compose down`. With the next start, the `burst_db` folder will be mounted again and your previous blockchain status will be loaded.

### H2

Alternatively, H2 can be used as database to store the blockchain. H2 is an embedded database, therefore one does not have to run it in a separate container. Simply run the following command.

`docker run -p 8123:8123 -p 8125:8125 -v `“`$(pwd)`”`/burst_db:/etc/burstcoin/burst_db -d burstcoin/core:2.2-h2 `

“`$(pwd)`”`/burst_db` is the path to the folder which is mounted to the H2 storage. If it does not exist a new `burst_db` folder is created in the current directory.

### Custom configuration

In order to use a custom config - `brs.properties` file, you can simply mount a folder containing the `brs-default.properties` and the `brs.properties` to the `/etc/burstcoin/conf` mount point.

**MariaDB**

    version: '3'

    services:
      burstcoin:
        image: burstcoin/core:2.2-mariadb
        restart: always
        depends_on:
         - mariadb
        ports:
         - 8123:8123
         - 8125:8125
        volumes:
         - ./conf:/etc/burstcoin/conf
      mariadb:
        image: mariadb:10
        environment:
         - MYSQL_ROOT_PASSWORD=burst
         - MYSQL_DATABASE=burst
        command: mysqld --character_set_server=utf8mb4
        volumes:
         - ./burst_db:/var/lib/mysql

**H2**

`docker run -p 8123:8123 -p 8125:8125 -v `“`$(pwd)`”`/burst_db:/etc/burstcoin/burst_db -v `“`$(pwd)`”`/conf:/etc/burstcoin/conf -d burstcoin/core:2.2-h2`

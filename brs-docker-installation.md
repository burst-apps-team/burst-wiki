| Burst Reference Software - Docker Installation |
|------------------------------------------------|
| **Status**                                     |

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

    <!--T:29-->
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

    <!--T:33-->
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

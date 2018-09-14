| Burst Reference Software - Linux Installation |
|-----------------------------------------------|
| **Status**                                    |

Table Of Contents
-----------------

\_\_TOC\_\_

### Debian & Ubuntu

Download the latest version of BRS here (zip) : <https://github.com/PoC-Consortium/burstcoin/releases>.

If you are running BRS for the first time,

-   install Java (https://www.java.com/en/download/help/linux\_install.xml)
-   install MariaDB (https://www.linode.com/docs/databases/mariadb/mariadb-setup-debian/\#install-mariadb)
-   run `burst.sh help`

If you want to import the blockchain now, run `burst.sh import mariadb`.

### Upgrading your wallet config from 1.3.6cg

`burst.sh upgrade`

will take the old `nxt-default.properties`/`nxt.properties` files and create `brs-default.properties.converted`/`brs.properties.converted` files in the conf directory. This should give you a headstart with the new option naming system.

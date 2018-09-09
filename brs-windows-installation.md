| Burst Reference Software - Windows Installation |
|-------------------------------------------------|
| **Status**                                      |

Table Of Contents
-----------------

\_\_TOC\_\_

### Installation with QBundle

All the details regarding the installation of BRS can be found here: [QBundle](qbundle.md). Indeed, QBundle is a launcher allowing you to install the Burst local wallet in a few clicks.

### Installation of BRS only (with MariaDB)

Firstly, you will have to download the latest version of BRS here (zip) : <https://github.com/PoC-Consortium/burstcoin/releases>. In the conf directory, copy brs-default.properties into a new file named brs.properties.

Secondly, be sure to have the latest version of Java installed on your computer : <https://java.com/de/download/>.

You will also have to download and install MariaDB : <https://downloads.mariadb.org/>.

The MariaDb installation will ask to setup a password for the root user. Add this password to the `brs.properties` file created above in the following section:

`DB.Url=jdbc:mariadb://localhost:3306/brs_master`
`DB.Username=root`
`DB.Password=YOUR_PASSWORD`

#### Setup MariaDB

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

Linux Install Video
-----------------
<iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/dwWGejP0ZC8" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

Ubuntu/Debian installation
--------------------------

The following steps are to install BRS node on any recent installation of **Unbuntu or other Debian based distros**. Please be sure to check Git Repository for the latest release before proceeding.

**Note:** Commands to be entered into the shell will be formatted in bold text, e.g.:

**`date`**

example command outputs will not be bold, e.g.:

`Wed 28 Nov 12:02:18 GMT 2018`

### Installing BRS

Make sure your installation is up to date

**`apt update && apt upgrade`**

Run the following command to install java:

**`apt install default-jre-headless -y`**

1. Have the latest version of Java installed on your computer with the following command:
**`apt install default-jre-headless -y`**

2. Download the latest version of [BRS (zip)](https://github.com/burst-apps-team/burstcoin/releases) and extract it to where ever you like. Optionally, in the conf directory, copy `brs-default.properties` into a new file named `brs.properties` and edit it to suit your needs.

3. Inside the folder you extracted the zip contents run the following:
**`java -jar burst.jar`**

This will show you a GUI with logs and percentage of blockchain download. If you don't want the GUI to be shown run as `java -jar burst.jar --headless`.

### Installing MariaDB (optional)

By default, a file based database backend is used (store on folder `burst_db`).
If you want to use MariaDB instead, in addition to the above steps, use the following procedure (stop BRS from running before going forward).

Install MariaDB:

**`apt install mariadb-server`**

During installation you will be prompted to enter a password for the root user. Enter and remember this as we will need it later.

Create a database(**brs\_master**), user (**brs\_user**) and assign a password to the user by running the following command:

**```echo "CREATE DATABASE brs_master; 
CREATE USER 'brs_user'@'localhost' IDENTIFIED BY 'yourpassword';
GRANT ALL PRIVILEGES ON brs_master.* TO 'brs_user'@'localhost';" | mysql -uroot -p
mysql -uroot -p < init-mysql.sql```**

This will prompt you for the password you entered for the root user when installing mariaDB

#### Configuring BRS to use MariaDB

Navigate to the conf subdirectory within your BRS directory:

**`cd BRSWallet/conf`**

Create a brs.properties file inside the conf folder:

**`nano brs.properties`**

Add the following lines:

**`DB.Url=jdbc:mariadb://localhost:3306/brs_master`**
**`DB.Username=brs_user`**
**`DB.Password=yourpassword`**

Run BRS again:

**`java -jar burst.jar`**

You will see somethign similar to the following output:

```
Nov 29, 2018 4:28:33 PM brs.Burst loadProperties
INFO: Initializing Burst Reference Software (BRS) version 2.5.0
Nov 29, 2018 4:28:33 PM org.ehcache.core.EhcacheManager createCache
INFO: Cache 'account' created in EhcacheManager.
[INFO] 2018-11-29 16:28:33 brs.db.sql.Db - Using mariadb Backend
```

**Note:** If closing the terminal or pressing Ctrl-c, the node will shutdown:

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
ExecStart=/usr/bin/java -jar /home/burst/burst.jar
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

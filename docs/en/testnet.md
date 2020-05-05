# TestNet

## Description

Burst has a permanent TestNet running. Its purpose is to provide a sandbox for both developers and users to test Burst features and to experiment with features without remorse. Note: NEVER USE YOUR REAL PASSPHRASE ON TESTNET.

## How to set up a TestNet node

To run a node on the TestNet, you will need to install the latest released version of the Burst Reference Software (BRS) that can be found on [GitHub](https://github.com/burst-apps-team/burstcoin/releases).

Edit/add the following in conf/brs.properties (Don't forget to set your address, platform, and database settings because if any parts are missing, you will most likely encounter issues):

```properties
# Please change the following 2 lines!
P2P.myAddress=
P2P.myPlatform=TestNet Node

# This needs to be true to be a full node
P2P.shareMyAddress=true

# The rest of the config is needed for testnet. Please do not change
API.Listen = 0.0.0.0
API.allowed = *

DEV.API.Port = 6876
# This is a new port; please expose it.
DEV.API.V2.Port = 6878

# For H2
DEV.DB.Url=jdbc:h2:file:./burst_testnet;DB_CLOSE_ON_EXIT=FALSE
# For MariaDB
#DEV.DB.Url=jdbc:mariadb://localhost:3306/burst-testnet
DEV.DB.Username=
DEV.DB.Password=

DEV.TestNet = yes
DEV.Offline = no

DEV.P2P.BootstrapPeers = 144.217.93.166; testnet.getburst.net; octalsburstnode.ddns.net; 3.16.150.48; 75.100.126.230; testddns.gotdns.com; aya.onthewifi.com; 212.37.175.94; burst-node-test.duckdns.org; test-burst.megash.it; happyfarmer.internet-box.ch;
DEV.P2P.rebroadcastTo = 144.217.93.166; testnet.getburst.net; octalsburstnode.ddns.net; 3.16.150.48; 75.100.126.230; testddns.gotdns.com; aya.onthewifi.com; 212.37.175.94; burst-node-test.duckdns.org; test-burst.megash.it; happyfarmer.internet-box.ch;

P2P.savePeers=true
P2P.usePeersDb=true
P2P.getMorePeers=true

DEV.digitalGoodsStore.startBlock = 0
DEV.automatedTransactions.startBlock = 0
DEV.atFixBlock2.startBlock = 0
DEV.atFixBlock3.startBlock = 0
DEV.atFixBlock4.startBlock = 0
DEV.preDymaxion.startBlock = 0
DEV.poc2.startBlock = 0
DEV.rewardRecipient.startBlock = 6500
```
Keep in mind there may be scenarios where you would like to have your own private test net, say in case of some future revisions that might require a fork for adding to the current chain. In this case, you could simply change the value of DEV.P2P.BootstrapPeers to the value of your localhost.

If you would like access to another TestNet peer without having to install your own, you can use any of the ones listed [here](https://burstwiki.org/en/testnet/#testnet-public-nodes). If you want to use your own desired peer, you can just change the ip address accordingly.

You can get TestNet Burst coins at this [faucet](http://burstcoin.cc:7777/).

If for any reason you need to reset your peer, follow these instructions:

If using the H2 database backend (which is the default), just remove all files under the burst_db folder. 

If using mariaDB database backend, perform these steps:

Stop the node and execute the following in a terminal:
````
mysql -u root
````
Once in the mysql shell, execute the following commands (assuming $yourdatabase is your database's name and $youruser is your user).
````
DROP DATABASE $yourdatabase;
CREATE DATABASE $yourdatabase;
FLUSH PRIVILEGES;
````
Start node again, wait until sync is complete.

## Current TestNet Facilities

* [Faucet](http://burstcoin.cc:7777/)

* [Mining Pool](http://75.100.126.230:8124/)

* [Explorer 1](http://explorer.testnet.burst.devtrue.net/)

* [Explorer 2](https://testnet.explorer.burstcoin.network/)

## Testnet Public Nodes

You can access the following BAT operated TestNet nodes

<http://testnet.getburst.net:6876/index.html>[1]

<https://testnet-2.burst-alliance.org:6876/index.html>

Other community run public nodes:

<http://testnet.burstcoin.network:6876/index.html>

## Have Questions?

Join our TestNet community on Discord here <https://discord.gg/CQEUe3e> and go to the \#testnet channel. Stay informed about current updates being deployed on TestNet.

## References

1. <https://www.burstcoin.ist/2019/02/10/weekly-burst-report-69/>


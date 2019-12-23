# TestNet

## Description

Burst has a permanent TestNet running. Its purpose is to provide a sandbox for both developers and users to test Burst features and to experiment with features without remorse. To talk about the testnet, visit us in <https://discord.gg/UnwXKB> in the channel \#testnet

Burst has [AT-Testnet](http://at-testnet.burst-alliance.org:6876/) with Mock-mining active, 1 minute blocks and [CIP-20](https://github.com/burst-apps-team/CIPs/blob/master/cip-0020.md) enabled.

## How to set up a TestNet node ?

To run a node on the new TestNet, you will need to install the latest released version of the wallet. You can find instructions on setting up the wallet from the Burstcoin [website](https://burst-coin.org). Keep in mind there may be scenarios where you would like to have your own private test net, say in case of some future revisions that might require a fork for adding to the current chain. In this case you could simply change the value for DEV.P2P.BootstrapPeers, the new value would be your localhost. Keep in mind if you are missing any parts of this configuration you will most likely encounter issues that are resolved by the values below. Again this is for testing purposes and DO NOT USE YOUR REAL PASSPHRASE.

We do recommend you to check in frequently in [discord](https://discord.gg/9S3eUBy) (`#testnet` channel) so that you are aware of updates currently being deployed to test net and actively want to participate.

Current TestNet facilities:

* [Faucet](http://burstcoin.cc:7777/)

* [Mining Pool](http://75.100.126.230:8124/)

* [Explorer](http://explorer.testnet.burst.devtrue.net/)

Current AT-TestNet facilities:

* [Faucet](http://at-testnet.burst-alliance.org:7777/)

Edit/add the following in conf/brs.properties (Don't forget to set your address, platform, and database settings):

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

If you just need testnet wallet for testing please use this [link](http://3.16.150.48:6876/index.html#). If you want to use your own desired peer, you can just change ip addresses.

You can get the testnet coins at this [faucet](http://burstcoin.cc:7777/)

If for any reason you need to reset your peer, a fast easy way to do it is as follows:

Stop node, in a terminal execute:
````
mysql -u root
````
Once in the mysql shell execute the following commands, assumption is made $yourdatabase name is your database's name and $youruser is well your user.
````
DROP DATABASE $yourdatabase;
CREATE DATABASE $yourdatabase;
FLUSH PRIVILEGES;
````
Start node again, wait till sync is complete.

If you hit any snags, let me know. Also feel free to update this via PR. Thanks.

Originally adapted from the following: https://burstwiki.org/wiki/Testnet

Thank you PoCC and BurstWiki contributors!

## Testnet Online Wallet

You can access the BAT's operated TestNet wallet via <http://testnet.getburst.net:6876/index.html>[2]

Community run public nodes:

<https://octalsburstnode.ddns.net:6876/index.html>

<https://test-burst.megash.it/index.html>

<http://testnet.burstcoin.network:6876/index.html>



## Testnet Faucet

-   <http://burstcoin.cc:7777/>[3]

## References

1. <https://github.com/burst-apps-team/BurstcoinTestNetHow-To>
2. <https://www.burstcoin.ist/2019/02/10/weekly-burst-report-69/>
3. <https://www.reddit.com/r/burstcoin/comments/aemjfn/testnet/>

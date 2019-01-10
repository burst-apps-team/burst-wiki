| TestNet    |
|------------|
| **Status** |
||

Description
-----------

Burst has a permanent TestNet running. Its purpose is to provide a sandbox for both developers and users to test Burst features and to experiment with features without remorse[1].

You can access the PoCC-operated online TestNet-wallet via <https://wallet.dev.burst-test.net/>

To talk about the testnet, visit us in <https://discord.gg/UnwXKB> in the channel \#testnet

How to set up a TestNet node ?
------------------------------

Before launching [BRS](burst-software-burst-reference-software--28brs-29.md), edit `brs.properties` and change the following lines

`DEV.DB.Url = jdbc:mariadb://localhost:3777/yourbursttestdatabase`
`DEV.DB.Username = root`
`DEV.DB.Password = yourpassword`

`#### TestNet Settings ####`
`# Use testnet, leave set to false unless you are really testing.`
`# Never unlock your real accounts on testnet! Use separate accounts for testing only.`
`# When using testnet, all custom port settings will be ignored,`
`# and hardcoded ports of 6874 (peer networking), 6875 (UI) and 6876 (API) will be used.`
`DEV.Offline = no`
`DEV.TestNet = yes`
`# Time Acceleration in Offline/TestNet configurations (1 = normal time, 2 = twice as fast ...)`
`DEV.TimeWarp = 1`
`# Force winning with every deadline`
`DEV.mockMining = off`
`DEV.preDymaxion.startBlock = 0`
`DEV.poc2.startBlock = 0`
`DEV.digitalGoodsStore.startBlock = 0`
`DEV.automatedTransactions.startBlock = 0`
`DEV.atFixBlock2.startBlock = 0`
`DEV.atFixBlock3.startBlock = 0`
`DEV.atFixBlock4.startBlock = 0`
`DEV.P2P.NumBootstrapConnections = 1`
`BRS.allowedBotHosts = *`
`# Only use this Hosts if you really want to use the testnet!!!`
`DEV.P2P.BootstrapPeers = 3.16.150.48; aya.onthewifi.com; testnet.burst.fun; testddns.gotdns.com; test-burst.megash.it`
`DEV.P2P.rebroadcastTo = 3.16.150.48; aya.onthewifi.com; testnet.burst.fun; testddns.gotdns.com; test-burst.megash.it`
`# Allow the wallet to discover new peers, store them in the DB and reuse them after restart.`
`P2P.savePeers=true`
`P2P.usePeersDb=true`
`P2P.getMorePeers=true`
`#### Personal Customization #### `
`P2P.shareMyAddress=true`
`# our examples are aya.onthewifi.com OR testnet.burst.fun OR testddns.gotdns.com OR test-burst.megash.it`
`P2P.myAddress=fqdn.to.your.wallet`
`# Change to your needs, this will be displayed in Peer Overview, modify this to whatever you want.`
`P2P.myPlatform=Official burst.megash.it TestNET Node`
`# Enable/Disable API requests used for blockchain and database manipulation.`
`#API.Debug=false`

`#### Additional Stuff ####`
`# Enter a version. Upon exit, print a list of peers having this version.`
`DEV.dumpPeersVersion =`
`# Force re-validation of blocks and transaction at start.`
`DEV.forceValidate = off`
`# Force re-build of derived objects tables at start.`
`DEV.forceScan = off`

obviously, you should use a different DB for TestNet than you use for MainNet.

You can access your TestNet-wallet via <http://127.0.0.1:6876/index.html> .

test-net facuet
---------------

... is still missing.

References
----------

<references />

[1] <https://www.burstcoin.ist/2018/04/30/weekly-burst-report-34/>

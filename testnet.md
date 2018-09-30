| TestNet    |
|------------|
| **Status** |
||

Description
-----------

Burst has a permanent TestNet running. Its purpose is to provide a sandbox for both developers and users to test Burst features and to experiment with features without remorse[1].

You can access the PoCC-operated online TestNet-wallet via <https://wallet.dev.burst-test.net/>

How to set up a TestNet node ?
------------------------------

Before launching [BRS](burst-reference-software.md), edit `brs.properties` and change the following lines

`DEV.TestNet = yes`
`DEV.P2P.BootstrapPeers = wallet.dev.burst-test.net`

`DEV.DB.Url =jdbc:mariadb://localhost:3777/yourbursttestdatabase`
`DEV.DB.Username =root`
`DEV.DB.Password =yourpassword`

obviously, you should use a different DB for TestNet than you use for MainNet.

Self-service test-net facuet
----------------------------

1.  Visit: <https://wallet.dev.burst-test.net/> and create a wallet on the testnet. Record the Numeric Account ID (this can be found by clicking “more info” under the account balance).
2.  Visit the self-service test-net facuet and send yourself some test-burst to your new test-net account using the information below:

`Faucet: BURST-PHPV-FVLQ-73XA-HFHJF`
`Passphrase: `“`I`` ``will`` ``take`` ``only`` ``500`` ``Burst`”

References
----------

<references />

[1] <https://www.burstcoin.ist/2018/04/30/weekly-burst-report-34/>

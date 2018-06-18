| The Burst API Examples |
|------------------------|
| **Status**             |
| **API Version**        |
| **BRS Version**        |

Description <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />
---------------------------------------------------------------------------------------------------------

Examples of Burst API calls are collected on this page, individually linked from the main [Burst API](the-burst-api.md) page. The organization and ordering is the same for both pages so that the section numbers in the table of contents are identical.

The preliminary sections preceding the examples simply link back to the main page. For example: [The Burst API Description](the-burst-api-description.md).

A lot of API calls can be viewed and tested on the MainNet at <https://wallet.burst.cryptoguru.org:8125/burst>?. For specific API calls, use the GET url https://wallet.burst.cryptoguru.org:8125/burst?=*specificRequestType*.

Table Of Contents
-----------------

\_\_TOC\_\_

General Notes <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />
-----------------------------------------------------------------------------------------------------------

[The Burst API General Notes](the-burst-api-general-notes.md)

Create Transaction <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />
----------------------------------------------------------------------------------------------------------------

[The Burst API Create Transaction](the-burst-api-create-transaction.md)

Account Operations <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />
----------------------------------------------------------------------------------------------------------------

### Get Account <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "unconfirmedBalanceNQT": "100000000000",
    "guaranteedBalanceNQT": "100000000000",
    "unconfirmedAssetBalances": [
        {
            "unconfirmedBalanceQNT": "120100",
            "asset": "3702027329806229573"
        }
    ],
    "effectiveBalanceNXT": "100000000000",
    "accountRS": "BURST-GBFG-HVQ4-8AMM-GPCWR",
    "name": "Umbrellacorp03",
    "forgedBalanceNQT": "0",
    "balanceNQT": "100000000000",
    "publicKey": "f22d8aa787eddbf69caf6f5960f5972a4b73247eb3a9479ddddeda40224aca60",
    "requestProcessingTime": 1,
    "assetBalances": [
        {
            "balanceQNT": "120100",
            "asset": "3702027329806229573"
        }
    ],
    "account": "17001464071916561838"
}
```

### Get Account Block Ids <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response similar to:**

``` json
{"blockIds":["18426055195962363092","7915836588136735869","10232042957216858995","6879436477641441477","10946006553575734351","11461420847451026714"],"requestProcessingTime":2}
```

### Get Account Blocks <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response similar to:**

``` json
{
    "blocks": [
        {
            "previousBlockHash": "9bb982b31c196a59231197da7abc82aeecea8d21017eb0c12fbabed8a74ec71c",
            "payloadLength": 1936,
            "totalAmountNQT": "106611838192",
            "generationSignature": "dedb23e56e9cc7d17d0f8d09bde18890c950089a2d7253889e8197967a9cacce",
            "generator": "15922713504207360763",
            "generatorPublicKey": "177d9ea55714741b002665dd9b7a1bc49834281c49162282735fadbb654f525d",
            "baseTarget": "62600",
            "payloadHash": "9c2e60b58de75577d4f088abe14111756f602fc0498c8a1cd97555514b3794f3",
            "generatorRS": "BURST-JNRV-L9MB-QU9G-FR8YT",
            "blockReward": "1101",
            "nextBlock": "7294788635189399942",
            "scoopNum": 3287,
            "numberOfTransactions": 11,
            "blockSignature": "51cfdb787b44c6ef65fd44bcc1efe0d1880a1f36573b1976849a34514eebfb0c657092b04ab329147b08c14b2d36fb3b769ac7d348dd8b42c706feba1c31fe0d",
            "transactions": [
                "9974453735926447276",
                "10292157220435803480",
                "12118294371901969706",
                "14492253203769113366",
                "16186302445201505070",
                "16411262478294799982",
                "1509465029747109302",
                "3032891187575274734",
                "3133520208744662995",
                "4199249628605348221",
                "4284107851893128481"
            ],
            "nonce": "11577401",
            "version": 3,
            "totalFeeNQT": "1100000000",
            "previousBlock": "6442989827968383387",
            "block": "12993921455729618779",
            "height": 467931,
            "timestamp": 113186379
        }
    ],
    "requestProcessingTime": 3
}
```

### Get Account Id <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "accountRS": "BURST-L6FM-89WK-VK8P-FCRBB",
    "publicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
    "requestProcessingTime": 0,
    "account": "15323192282528158131"
}
```

### Get Account Lessors <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "lessors": [],
    "accountRS": "BURST-GBFG-HVQ4-8AMM-GPCWR",
    "requestProcessingTime": 0,
    "account": "17001464071916561838",
    "height": 496781
}
```

### Get Account Public Key <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "publicKey": "f22d8aa787eddbf69caf6f5960f5972a4b73247eb3a9479ddddeda40224aca60",
    "requestProcessingTime": 1
}
```

### Get Account Transaction Ids <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "transactionIds": [
        "14471919803527301514",
        "8408429517094397948",
        "7677510357080940908",
        "15605878519502379168",
        "13241821260511921007",
        "16748761036604486700",
        "6521866371443385678"
    ],
    "requestProcessingTime": 7
}
```

### Get Account Transactions <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "requestProcessingTime": 7,
    "transactions": [
        {
            "senderPublicKey": "2ebd275f257b5e73c58c1a5e13efbcba8a579829a3c336e79409f32a3bcf8379",
            "signature": "e96b97ee10365498641177ec07378b252b9bc6e727dc67130c4ee32e38cede076b24c0f3c2137a138234c154656b8cee45b913e2842bced41f8d8c195670eeb5",
            "feeNQT": "100000000",
            "type": 2,
            "confirmations": 85,
            "fullHash": "8a892fc36296d6c8540dab8d52458500e400483b091224db949fe8d36ce8aac5",
            "version": 1,
            "ecBlockId": "6134220061079224356",
            "signatureHash": "0811bab0ed422d3da9fe837e57bd38aa9bbe223f9bcc34109a650eaadc21ce5e",
            "attachment": {
                "version.AssetTransfer": 1,
                "quantityQNT": "10000",
                "asset": "3702027329806229573"
            },
            "senderRS": "BURST-CMTM-GH7K-58PB-BFP2U",
            "subtype": 1,
            "amountNQT": "0",
            "sender": "11313795926748122931",
            "recipientRS": "BURST-GBFG-HVQ4-8AMM-GPCWR",
            "recipient": "17001464071916561838",
            "ecBlockHeight": 496688,
            "block": "9331719154915300577",
            "blockTimestamp": 120142856,
            "deadline": 1440,
            "transaction": "14471919803527301514",
            "timestamp": 120142756,
            "height": 496701
        },
        {
            "senderPublicKey": "f22d8aa787eddbf69caf6f5960f5972a4b73247eb3a9479ddddeda40224aca60",
            "signature": "e74d1213e17fdd41f6337eb3cf4c501089931b651ec4b3a563cf6b113ab58b0cfd2516f7e9d64cbbdcf87e8e35ae75caddef7811d729b4a53338a4c75bb85977",
            "feeNQT": "100000000",
            "type": 0,
            "confirmations": 356,
            "fullHash": "fc27529518beb074357b59f0efa477e18240c9f44a02b89fd32ccdb0b85b135a",
            "version": 1,
            "ecBlockId": "6116962714576197152",
            "signatureHash": "f10eaaaf1836f07fe6958293cc79a91fb57dd8dfec9780ea73e0ee82de18c292",
            "senderRS": "BURST-GBFG-HVQ4-8AMM-GPCWR",
            "subtype": 0,
            "amountNQT": "53329697341",
            "sender": "17001464071916561838",
            "recipientRS": "BURST-NQAW-GQ84-CA6C-7B3LS",
            "recipient": "6649610434652035356",
            "ecBlockHeight": 496419,
            "block": "5936178255841968699",
            "blockTimestamp": 120078307,
            "deadline": 1440,
            "transaction": "8408429517094397948",
            "timestamp": 120078236,
            "height": 496430
        }
    ]
}
```

### Get Balance <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "unconfirmedBalanceNQT": "100000000000",
    "guaranteedBalanceNQT": "100000000000",
    "effectiveBalanceNXT": "100000000000",
    "forgedBalanceNQT": "0",
    "balanceNQT": "100000000000",
    "requestProcessingTime": 0
}
```

### Get Guaranteed Balance <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "guaranteedBalanceNQT": "100000000000",
    "requestProcessingTime": 0
}
```

### Get Unconfirmed Transaction Ids <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "requestProcessingTime": 0,
    "unconfirmedTransactionIds": []
}
```

### Get Unconfirmed Transactions <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "unconfirmedTransactions": [
        {
            "senderPublicKey": "a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12",
            "signature": "5316c88b00f767d2bbe0cac07ae4f84b1fe75f002e24e9ce1c1a6b81d6821f09474a4db22db8461be85ac03f7c7df37fa007bd1aa72c34fc4be0e64b605bb2d6",
            "feeNQT": "100000000",
            "type": 1,
            "fullHash": "480ae3eba6a4e445c692eacae605bfaecd060e91005f96b09d57647d3c5fcf25",
            "version": 1,
            "ecBlockId": "3371057862366354193",
            "signatureHash": "d156bf1cdf9c8fce254a9ef3a7d47ed2f16a35ad94c8c625ec48a83fbda1b615",
            "attachment": {
                "name": "API-Examples2",
                "description": "",
                "version.AccountInfo": 1
            },
            "senderRS": "BURST-FRDJ-UPLH-MY9A-GUKQP",
            "subtype": 5,
            "amountNQT": "0",
            "sender": "16922903237994405232",
            "ecBlockHeight": 498894,
            "deadline": 1440,
            "transaction": "5036331320136108616",
            "timestamp": 120675576,
            "height": 498903
        },
        {
            "senderPublicKey": "25cc2bb30ee7665737c9721090313c85176e485cd9a15495a0f3abc359d8d632",
            "signature": "194bd83c8473d5bfef972e50efde4699db401d9097dcb68eaa898042d5022c0a1e64d551e20d6f49402d34f930317577f90152949485f8515d8b89e5bbbf45c0",
            "feeNQT": "100000000",
            "type": 0,
            "fullHash": "0c53db6986e2ba38119354158128e60f23a749fec0acf893ff3ca2de22176df0",
            "version": 1,
            "ecBlockId": "3371057862366354193",
            "signatureHash": "3b1b88980766b041165a55c622a6eea4b08bb628eab95e3016c19838fda6c112",
            "senderRS": "BURST-HKML-NRG6-VBRA-2F8PS",
            "subtype": 0,
            "amountNQT": "10013275142",
            "sender": "888561138747819634",
            "recipientRS": "BURST-WXWK-MD2A-KXJL-HR27T",
            "recipient": "18200197947533981585",
            "ecBlockHeight": 498894,
            "deadline": 1440,
            "transaction": "4087828678721622796",
            "timestamp": 120675589,
            "height": 498903
        },
    ],
    "requestProcessingTime": 1
}
```

### Send Money <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "signatureHash": "b6bb90bac0d529d9ebc6771089f389c976db2a67d76867f7c9d2e1b90e00358a",
    "unsignedTransactionBytes": "001092812d07180057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93cae2527ec7e55f1eb00e1f5050000000000e1f505*",
    "transactionJSON": {
        "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
        "signature": "3b709a56a9a2308b2364cdb260cc8b08126b8ed13cfe891aa817c28760b0e90f108680db998938fd4747390a0ed9b3b3e25674b849e583dce011155f95ecaa3a",
        "feeNQT": "100000000",
        "type": 0,
        "fullHash": "808d5c32b12f4d4b963404c19523b6391ddf7a04a96ec4a495703aeead76c6ff",
        "version": 1,
        "ecBlockId": "16622227543717857480",
        "signatureHash": "b6bb90bac0d529d9ebc6771089f389c976db2a67d76867f7c9d2e1b90e00358a",
        "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
        "subtype": 0,
        "amountNQT": "100000000",
        "sender": "15323192282528158131",
        "recipientRS": "BURST-GBFG-HVQ4-8AMM-GPCWR",
        "recipient": "17001464071916561838",
        "ecBlockHeight": 497841,
        "deadline": 24,
        "transaction": "5426045564151958912",
        "timestamp": 120422802,
        "height": 2147483647
    },
    "broadcasted": true,
    "requestProcessingTime": 4,
    "transactionBytes": "001092812d07180057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93cae2527ec7e55f1eb00e1f5050000000000e1f505*",
    "fullHash": "808d5c32b12f4d4b963404c19523b6391ddf7a04a96ec4a495703aeead76c6ff",
    "transaction": "5426045564151958912"
}
```

### Set Account Info <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "signatureHash": "00e5694c213c978045fc32125092a1f3b16e6c5bf4d1d24c51005466be29f14b",
    "unsignedTransactionBytes": "0115c1832d07180057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c0000000000000000000000000000000000e1f505*",
    "transactionJSON": {
        "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
        "signature": "9b2634d7f439bf0c71341511a3efb202f744ec7f1b60673d07965bb297b4060d27bed164bf2611b738da1df34cd15a44458bfed2f6e78c94707af11da46f5044",
        "feeNQT": "100000000",
        "type": 1,
        "fullHash": "7ace7428ff55c4c5963d828f0cd210be4c988b482b2930b73c80c7fe55c19af5",
        "version": 1,
        "ecBlockId": "4725767517890678156",
        "signatureHash": "00e5694c213c978045fc32125092a1f3b16e6c5bf4d1d24c51005466be29f14b",
        "attachment": {
            "name": "API-Examples",
            "description": "",
            "version.AccountInfo": 1
        },
        "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
        "subtype": 5,
        "amountNQT": "0",
        "sender": "15323192282528158131",
        "ecBlockHeight": 497844,
        "deadline": 24,
        "transaction": "14250609675290857082",
        "timestamp": 120423361,
        "height": 2147483647
    },
    "broadcasted": true,
    "requestProcessingTime": 9,
    "transactionBytes": "0115c1832d07180057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c0000000000000000000000000000000000e1f505*",
    "fullHash": "7ace7428ff55c4c5963d828f0cd210be4c988b482b2930b73c80c7fe55c19af5",
    "transaction": "14250609675290857082"
}
```

Alias Operations <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />
--------------------------------------------------------------------------------------------------------------

### Buy / Sell Alias <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "signatureHash": "952eea5e0eb67f1ec78abfab60edcde55f477ae8dcf8f620077783532217c329",
    "unsignedTransactionBytes": "0116a16031071800a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12ae2527ec7e55f1eb000000000000000000e1f505*",
    "transactionJSON": {
        "senderPublicKey": "a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12",
        "signature": "6ce34c1f9a73ce7c375b9a0eeec1942f9932ecd6e34f8b5c9358a501007cf709fe26a2f205be7180af0cca753f89c648b6a1481bce7ed1e8b91dcbb8a0e7f0ce",
        "feeNQT": "100000000",
        "type": 1,
        "fullHash": "2d923132c2cd045c7c8d901b261ab237dbac7ebd50d3f7c549e601c8cc0f1cf9",
        "version": 1,
        "ecBlockId": "6797965034232082546",
        "signatureHash": "952eea5e0eb67f1ec78abfab60edcde55f477ae8dcf8f620077783532217c329",
        "attachment": {
            "alias": "AliasTestSell",
            "priceNQT": "1",
            "version.AliasSell": 1
        },
        "senderRS": "BURST-FRDJ-UPLH-MY9A-GUKQP",
        "subtype": 6,
        "amountNQT": "0",
        "sender": "16922903237994405232",
        "recipientRS": "BURST-GBFG-HVQ4-8AMM-GPCWR",
        "recipient": "17001464071916561838",
        "ecBlockHeight": 498897,
        "deadline": 24,
        "transaction": "6630650785345671725",
        "timestamp": 120676513,
        "height": 2147483647
    },
    "broadcasted": true,
    "requestProcessingTime": 7,
    "transactionBytes": "0116a16031071800a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12ae2527ec7e55f1eb000000000000000000e1f505*",
    "fullHash": "2d923132c2cd045c7c8d901b261ab237dbac7ebd50d3f7c549e601c8cc0f1cf9",
    "transaction": "6630650785345671725"
}
```

### Set Alias <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "signatureHash": "c59f26e0b40c306f5ecf0c222a51a83578a9f04165ff27d5f16a4f9f00fc63f7",
    "unsignedTransactionBytes": "0111a2852d07180057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c0000000000000000000000000000000000e1f505*",
    "transactionJSON": {
        "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
        "signature": "c27bd8a00b7f8025f7d5054d855d019f939388b1bb52781d9e7d901a81da1a04c726069306f881c2676ae8445cb2bf943465ce32d7daa9fae97a02f597db50b1",
        "feeNQT": "100000000",
        "type": 1,
        "fullHash": "ebe26aea6b9996886b245485f6a5005cbd9c4e2584211cf0c5fc16548db7ba29",
        "version": 1,
        "ecBlockId": "4323691795607025950",
        "signatureHash": "c59f26e0b40c306f5ecf0c222a51a83578a9f04165ff27d5f16a4f9f00fc63f7",
        "attachment": {
            "alias": "ApiExample",
            "version.AliasAssignment": 1,
            "uri": "acct:burst-l6fm-89wk-vk8p-fcrbb@burst"
        },
        "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
        "subtype": 1,
        "amountNQT": "0",
        "sender": "15323192282528158131",
        "ecBlockHeight": 497845,
        "deadline": 24,
        "transaction": "9842222724438221547",
        "timestamp": 120423842,
        "height": 2147483647
    },
    "broadcasted": true,
    "requestProcessingTime": 11,
    "transactionBytes": "0111a2852d07180057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c0000000000000000000000000000000000e1f505*",
    "transaction": "9842222724438221547"
}
```

### Get Alias <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "aliasURI": "acct:burst-frdj-uplh-my9a-gukqp@burst",
    "aliasName": "AliasTestSell",
    "accountRS": "BURST-FRDJ-UPLH-MY9A-GUKQP",
    "alias": "5747297803058313219",
    "requestProcessingTime": 1,
    "account": "16922903237994405232",
    "timestamp": 120675864
}
```

### Get Aliases <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "aliases": [
        {
            "aliasURI": "acct:burst-frdj-uplh-my9a-gukqp@burst",
            "aliasName": "AliasTestBuy",
            "accountRS": "BURST-FRDJ-UPLH-MY9A-GUKQP",
            "alias": "17197436179114898263",
            "account": "16922903237994405232",
            "timestamp": 120676320
        },
        {
            "aliasURI": "acct:burst-frdj-uplh-my9a-gukqp@burst",
            "aliasName": "AliasTestSell",
            "accountRS": "BURST-FRDJ-UPLH-MY9A-GUKQP",
            "alias": "5747297803058313219",
            "account": "16922903237994405232",
            "timestamp": 120675864
        }
    ],
    "requestProcessingTime": 2
}
```

Arbitrary Message System Operations <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />
---------------------------------------------------------------------------------------------------------------------------------

### Encrypt To <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "data": "d5a1958d12ce96ce30dbce5b6c8ead7ecbc0f59d857dc8e8fbeec10ae440e0e74e9120fef3b0fa586d4c63fde0f289340e709b30ae528e3c2d740b11e3ae3fdb5e5d5c63f724cf16157c75dabec31eaf",
    "requestProcessingTime": 34,
    "nonce": "7cefa6f66d5b71604e2ef56a18319b3f48a38e8aa5cf610369b294f1d40e0f8e"
}
```

### Decrypt From <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "decryptedMessage": "This is a message encrypted using \"encryptTo\".",
    "requestProcessingTime": 2
}
```

### Send Message <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "signatureHash": "2bcbafeab7a0bae40337fe34adea84110b1f770a33841c95f3fe9e19dde41bae",
    "unsignedTransactionBytes": "01104ec93f071800a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12b31119f931eaa6d4000000000000000018370b*",
    "transactionJSON": {
        "senderPublicKey": "a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12",
        "signature": "dc2503584a48e30ac62d62848f58461e0e9ff55070008743c24f380c24a9ef05525c70b5d40962566f3f4de2018277ba7956eb09d0aec84219784de7f3b76f6a",
        "feeNQT": "735000",
        "type": 1,
        "fullHash": "0f37d045bc7d4f2bd85cb565a5c4e575464ac387b986f80fb8c31635cf03923e",
        "version": 1,
        "ecBlockId": "1212249281481197658",
        "signatureHash": "2bcbafeab7a0bae40337fe34adea84110b1f770a33841c95f3fe9e19dde41bae",
        "attachment": {
            "version.Message": 1,
            "messageIsText": true,
            "message": "This is a sendMessage API example"
        },
        "senderRS": "BURST-FRDJ-UPLH-MY9A-GUKQP",
        "subtype": 0,
        "amountNQT": "0",
        "sender": "16922903237994405232",
        "recipientRS": "BURST-L6FM-89WK-VK8P-FCRBB",
        "recipient": "15323192282528158131",
        "ecBlockHeight": 502787,
        "deadline": 24,
        "transaction": "3120851314369640207",
        "timestamp": 121620814,
        "height": 2147483647
    },
    "broadcasted": true,
    "requestProcessingTime": 11,
    "transactionBytes": "01104ec93f071800a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12b31119f931eaa6d4000000000000000018370b*",
    "fullHash": "0f37d045bc7d4f2bd85cb565a5c4e575464ac387b986f80fb8c31635cf03923e",
    "transaction": "3120851314369640207"
}
```

### Read Message <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "requestProcessingTime": 0,
    "message": "This is a sendMessage API example"
}
```

Asset Exchange Operations
-------------------------

### Cancel Order

**Request:**

    http://localhost:8125/burst?
      requestType=cancelBidOrder&
      secretPhrase=IWontTellYou&
      order=17185236428295897167&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "20264f33a06331f6a8d7c4362d0525aee25e4ef991653f14bbfb1b2beebba433",
     "unsignedTransactionBytes": "02153c6ed6013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "30cf47942179f5969369699b548f5a3517ef7ff71b337db630ea5f8b6e740...",
      "feeNQT": "100000000",
      "type": 2,
      "fullHash": "725f441d9f50a9b2e02d780098a827b1015ec902199becd493bfa73a4843ae89",
      "version": 1,
      "ecBlockId": "196207598250363138",
      "signatureHash": "20264f33a06331f6a8d7c4362d0525aee25e4ef991653f14bbfb1b2beebba433",
      "attachment": {
       "version.BidOrderCancellation": 1,
       "order": "17185236428295897167"
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 5,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "ecBlockHeight": 285586,
      "deadline": 60,
      "transaction": "12873909654136315762",
      "timestamp": 30830140,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 7640,
     "transactionBytes": "02153c6ed6013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "fullHash": "725f441d9f50a9b2e02d780098a827b1015ec902199becd493bfa73a4843ae89",
     "transaction": "12873909654136315762"
    }

<small>*Verified 16-Nov-14*</small>

### Get Account Current Order Ids

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountCurrentBidOrderIds&
      account=BURST-L6FM-89WK-VK8P-FCRBB&
      asset=17554243582654188572

**Response:**

    {
     "bidOrderIds": [
      "17185236428295897167"
     ],
     "requestProcessingTime": 4
    }

<small>*Verified 16-Nov-14*</small>

### Get Account Current Orders

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountCurrentBidOrders&
      account=BURST-L6FM-89WK-VK8P-FCRBB

**Response:**

    {
     "bidOrders": [
      {
       "quantityQNT": "1000000",
       "priceNQT": "100",
       "accountRS": "BURST-L6FM-89WK-VK8P-FCRBB",
       "asset": "17554243582654188572",
       "type": "bid",
       "account": "15323192282528158131",
       "order": "17185236428295897167",
       "height": 285549
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 16-Nov-14*</small>

### Get All Assets

**Request:**

    http://localhost:8125/burst?
      requestType=getAllAssets&
      firstIndex=13&
      lastIndex=13

**Response:**

    {
     "assets": [
      {
       "quantityQNT": "2100000000000000",
       "numberOfAccounts": 31,
       "accountRS": "BURST-3TKA-UH62-478B-DQU6K",
       "decimals": 8,
       "numberOfTransfers": 63,
       "name": "mgwBTC",
       "description": "Production Multigateway BTC (mgwBTC) is backed 100% by...",
       "numberOfTrades": 9,
       "asset": "17554243582654188572",
       "account": "13300069592148796968"
      }
     ],
     "requestProcessingTime": 13
    }

<small>*Verified 18-Nov-14*</small>

### Get All Open Orders

**Request:**

    http://localhost:8125/burst?
      requestType=getAllOpenBidOrders&
      firstIndex=123&
      lastIndex=123

**Response:**

    {
     "requestProcessingTime": 5631,
     "openOrders": [
      {
       "quantityQNT": "1000000",
       "priceNQT": "101",
       "accountRS": "BURST-L6FM-89WK-VK8P-FCRBB",
       "asset": "17554243582654188572",
       "type": "bid",
       "account": "15323192282528158131",
       "order": "12743274869785967304",
       "height": 285577
      }
     ]
    }

<small>*Verified 17-Nov-14*</small>

### Get All Trades

**Request:**

    http://localhost:8125/burst?
      requestType=getAllTrades&
      lastIndex=0

**Response:**

    {
     "trades": [
      {
       "seller": "14968762166783718535",
       "quantityQNT": "10000",
       "bidOrder": "1166717226538227076",
       "sellerRS": "BURST-8F69-W9Z9-8M6Y-ETXGZ",
       "buyer": "202478233571806601",
       "priceNQT": "19796",
       "askOrder": "16222071953599729591",
       "buyerRS": "BURST-RMEB-W7TE-28EM-2SUM7",
       "decimals": 8,
       "name": "mgwBTC",
       "block": "8807797247643599359",
       "asset": "4551058913252105307",
       "askOrderHeight": 285768,
       "bidOrderHeight": 286453,
       "tradeType": "buy",
       "timestamp": 30920039,
       "height": 286453
      }
     ],
     "requestProcessingTime": 25840
    }

<small>*Verified 17-Nov-14*</small>

### Get Asset

**Request:**

    http://localhost:8125/burst?
      requestType=getAsset&
      asset=17554243582654188572

**Response:**

    {
     "quantityQNT": "2100000000000000",
     "numberOfAccounts": 31,
     "accountRS": "BURST-3TKA-UH62-478B-DQU6K",
     "decimals": 8,
     "numberOfTransfers": 63,
     "name": "mgwBTC",
     "description": "Production Multigateway BTC (mgwBTC) is backed 100% by...",
     "numberOfTrades": 9,
     "requestProcessingTime": 11,
     "asset": "17554243582654188572",
     "account": "13300069592148796968"
    }

<small>*Verified 18-Nov-14*</small>

### Get Asset Accounts

**Request:**

    http://localhost:8125/burst?
      requestType=getAssetAccounts&
      asset=5539238107226883203

**Response:**

    {
     "accountAssets": [
      {
       "quantityQNT": "100000000",
       "accountRS": "BURST-JTE5-HB7Y-QPS5-B58MZ",
       "unconfirmedQuantityQNT": "100000000",
       "asset": "5539238107226883203",
       "account": "11514793277306463619"
      }
     ],
     "requestProcessingTime": 19
    }

<small>*Verified 19-Nov-14*</small>

### Get Asset Ids

**Request:**

    http://localhost:8125/burst?
      requestType=getAssetIds&
      firstIndex=15&
      lastIndex=15

**Response:**

    {
     "assetIds": [
      "17554243582654188572"
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 19-Nov-14*</small>

### Get Asset Transfers

**Request:**

    http://localhost:8125/burst?
      requestType=getAssetTransfers&
      asset=17554243582654188572&
      lastIndex=0

**Response:**

    {
     "transfers": [
      {
       "quantityQNT": "500000",
       "senderRS": "BURST-9K6X-4ZQS-PCQN-465T4",
       "assetTransfer": "15255934090738722602",
       "sender": "2406158154854548637",
       "recipientRS": "BURST-YMEM-ERX3-BFUZ-4MQ4P",
       "decimals": 8,
       "recipient": "3065494931320556947",
       "name": "mgwBTC",
       "asset": "17554243582654188572",
       "height": 287648,
       "timestamp": 31042663
      }
     ],
     "requestProcessingTime": 4
    }

<small>*Verified 19-Nov-14*</small>

### Get Assets

**Request:**

    http://localhost:8125/burst?
      requestType=getAssets&
      assets=17554243582654188572

**Response:**

    {
     "assets": [
      {
       "quantityQNT": "2100000000000000",
       "numberOfAccounts": 31,
       "accountRS": "BURST-3TKA-UH62-478B-DQU6K",
       "decimals": 8,
       "numberOfTransfers": 63,
       "name": "mgwBTC",
       "description": "Production Multigateway BTC (mgwBTC) is backed 100% by...",
       "numberOfTrades": 9,
       "asset": "17554243582654188572",
       "account": "13300069592148796968"
      }
     ],
     "requestProcessingTime": 15
    }

<small>*Verified 18-Nov-14*</small>

### Get Assets By Issuer

**Request:**

    http://localhost:8125/burst?
      requestType=getAssetsByIssuer&
      account=BURST-DE2F-W76R-GL25-HMFPR&
      lastIndex=0

**Response:**

    {
     "assets": [
      [
       {
        "quantityQNT": "10000000000",
        "numberOfAccounts": 222,
        "accountRS": "BURST-DE2F-W76R-GL25-HMFPR",
        "decimals": 4,
        "numberOfTransfers": 278,
        "name": "NXTprivacy",
        "description": "NXTprivacy will contain various privacy related projects...",
        "numberOfTrades": 456,
        "asset": "17911762572811467637",
        "account": "18146608053740744717"
       }
      ]
     ],
     "requestProcessingTime": 9
    }

<small>*Verified 19-Nov-14*</small>

### Get Order

**Request:**

    http://localhost:8125/burst?
      requestType=getAskOrder&
      order=6044046093672850641

**Response:**

    {
     "quantityQNT": "100",
     "priceNQT": "100000000",
     "transactionHeight": 346634,
     "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
     "transactionIndex": 0,
     "requestProcessingTime": 0,
     "asset": "17091401215301664836",
     "type": "ask",
     "account": "15295723609781267838",
     "order": "6044046093672850641",
     "height": 346634
    }

<small>*Verified 10-Jul-15*</small>

### Get Order Ids

**Request:**

    http://localhost:8125/burst?
      requestType=getBidOrderIds&
      asset=17554243582654188572

**Response:**

    {
     "bidOrderIds": [
      "17972270381487138621",
      "8331653287549483600",
      "16386956089071870421",
      "12743274869785967304",
      "3409888667133338290"
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 17-Nov-14*</small>

### Get Orders

**Request:**

    http://localhost:8125/burst?
      requestType=getBidOrders&
      asset=17554243582654188572&
      firstIndex=3&
      lastIndex=3

**Response:**

    {
     "bidOrders": [
      {
       "quantityQNT": "1000000",
       "priceNQT": "101",
       "accountRS": "BURST-L6FM-89WK-VK8P-FCRBB",
       "asset": "17554243582654188572",
       "type": "bid",
       "account": "15323192282528158131",
       "order": "12743274869785967304",
       "height": 285577
      }
     ],
     "requestProcessingTime": 2
    }

<small>*Verified 17-Nov-14*</small>

### Get Trades

**Request:**

    http://localhost:8125/burst?
      requestType=getTrades&
      asset=17554243582654188572&
      lastIndex=0

**Response:**

    {
     "trades": [
      {
       "seller": "4012743767778395236",
       "quantityQNT": "922082",
       "bidOrder": "18332182738291742411",
       "sellerRS": "BURST-J356-8B4K-L4DK-533EH",
       "buyer": "17013046603665206934",
       "priceNQT": "19607",
       "askOrder": "5860848661439768841",
       "buyerRS": "BURST-4VNQ-RWZC-4WWQ-GVM8S",
       "decimals": 8,
       "name": "mgwBTC",
       "block": "15627545821108097361",
       "asset": "17554243582654188572",
       "askOrderHeight": 285964,
       "bidOrderHeight": 285821,
       "tradeType": "sell",
       "timestamp": 30869573,
       "height": 285964
      }
     ],
     "requestProcessingTime": 0
    }

<small>*Verified 17-Nov-14*</small>

### Issue Asset

**Request:**

    http://localhost:8125/burst?
      requestType=issueAsset&
      publicKey=57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c&
      name=SecretCoin&
      description=This+is+SecretCoin&
      quantityQNT=100&
      deadline=60&
      feeNQT=100000000000&
      broad=false

**Response:**

    {
     "unsignedTransactionBytes": "021095e5da013c0073080c6a224062660184f10ebb7fb431d459364a12403...",
     "transactionJSON": {
      "senderPublicKey": "73080c6a224062660184f10ebb7fb431d459364a12403320c7f601f9d75cc547",
      "feeNQT": "100000000000",
      "type": 2,
      "version": 1,
      "ecBlockId": "1564408139943737911",
      "attachment": {
       "name": "SecretCoin",
       "description": "This+is+SecretCoin",
       "quantityQNT": "100",
       "version.AssetIssuance": 1,
       "decimals": 0
      },
      "senderRS": "BURST-4VNQ-RWZC-4WWQ-GVM8S",
      "subtype": 0,
      "amountNQT": "0",
      "sender": "17013046603665206934",
      "ecBlockHeight": 288402,
      "deadline": 60,
      "timestamp": 31122837,
      "height": 2147483647
     },
     "broadcasted": false,
     "requestProcessingTime": 2
    }

<small>*Verified 19-Nov-14*</small>

### Place Order

**Request:**

    http://localhost:8125/burst?
      requestType=placeBidOrder&
      secretPhrase=IWontTellYou&
      asset=17554243582654188572&
      quantityQNT=1000000&
      priceNQT=100&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "44028b4ddb46e7d4383331425b79019bb0f004f88ede12a5aa66f05c23a75f03",
     "unsignedTransactionBytes": "02135a5ed6013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "65713e80be99a927bfa7bb2e9a7b6fbd2f17c226fb956494c68a6d90a8127...",
      "feeNQT": "100000000",
      "type": 2,
      "fullHash": "4f00aef17a397eee25027e834ca765660e4e3f3f1b162468bdac67b315aeb812",
      "version": 1,
      "ecBlockId": "14593256906948324209",
      "signatureHash": "44028b4ddb46e7d4383331425b79019bb0f004f88ede12a5aa66f05c23a75f03",
      "attachment": {
       "quantityQNT": "1000000",
       "priceNQT": "100",
       "asset": "17554243582654188572",
       "version.BidOrderPlacement": 1
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 3,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "ecBlockHeight": 285545,
      "deadline": 60,
      "transaction": "17185236428295897167",
      "timestamp": 30826074,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 8729,
     "transactionBytes": "02135a5ed6013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143...",
     "fullHash": "4f00aef17a397eee25027e834ca765660e4e3f3f1b162468bdac67b315aeb812",
     "transaction": "17185236428295897167"
    }

<small>*Verified 16-Nov-14*</small>

### Transfer Asset

**Request:**

    http://localhost:8125/burst?
      requestType=transferAsset&
      secretPhrase=IWontTellYou&
      recipient=BURST-4VNQ-RWZC-4WWQ-GVM8S&
      asset=17554243582654188572&
      quantityQNT=1000&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "f622557588bc82942984286e431c978e687783b32db6a68a1c554b2e11349751",
     "unsignedTransactionBytes": "0211c9ebda013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "33d0fccf1f19d18b4ca97d2636cf62f7a801c07628d1bbb1d6084dc3ca658d...",
      "feeNQT": "100000000",
      "type": 2,
      "fullHash": "c8f25b15b48fb5efd3341fb369627d00f8fefb59dc18016ba0c482b6de7cad59",
      "version": 1,
      "ecBlockId": "3925493493266246517",
      "signatureHash": "f622557588bc82942984286e431c978e687783b32db6a68a1c554b2e11349751",
      "attachment": {
       "version.AssetTransfer": 1,
       "quantityQNT": "1000",
       "asset": "17554243582654188572"
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 1,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "recipientRS": "BURST-4VNQ-RWZC-4WWQ-GVM8S",
      "recipient": "17013046603665206934",
      "ecBlockHeight": 288416,
      "deadline": 60,
      "transaction": "17272869949464638152",
      "timestamp": 31124425,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 4968,
     "transactionBytes": "0211c9ebda013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473b...",
     "fullHash": "c8f25b15b48fb5efd3341fb369627d00f8fefb59dc18016ba0c482b6de7cad59",
     "transaction": "17272869949464638152"
    }

<small>*Verified 19-Nov-14*</small>

Block Operations
----------------

### Get Block

**Request:**

    http://localhost:8125/burst?
      requestType=getBlock&
      block=8455642159445842600

**Response:**

    {
     "previousBlockHash": "c0574d7a7b8497373dbead497c2dc7f60fdcfc8c5a9fcb48f7d373acc9bbb099",
     "payloadLength": 1189,
     "totalAmountNQT": "100000000",
     "generationSignature": "c5098d37267bc71134fc8572a87b4af8727a2e5139d60fe8833fab98af22244a",
     "generator": "11693867635361772359",
     "generatorPublicKey": "fbb72a280228af5c8c74c7c754a290e1539f839553c00bc560cac7bfdb324a7c",
     "baseTarget": "1530224444",
     "payloadHash": "633992be640a593ba04e31ca4028deed70bbf47cff333e2a0372a4e2a4aba205",
     "generatorRS": "BURST-TWU9-P3E4-HCDM-CQ9L6",
     "nextBlock": "5937170741469897491",
     "requestProcessingTime": 175,
     "numberOfTransactions": 6,
     "blockSignature": "ff65a82e385c135cf9bd5be0861e9e5d3d3174fbd993e5b7f57935ec4...",
     "transactions": [
      "15184285173972564233",
      "15200280108574630445",
      "15200507403046301754",
      "15900338016714606285",
      "17881859777840687131",
      "18361738217269620028"
     ],
     "version": 3,
     "totalFeeNQT": "600000000",
     "previousBlock": "4005816059437078464",
     "cumulativeDifficulty": "10229109959119715",
     "block": "8455642159445842600",
     "height": 275730,
     "timestamp": 29797208
    }

<small>*Verified 15-Dec-14*</small>

### Get Block Id

**Request:**

    http://localhost:8125/burst?
      requestType=getBlockId&
      height=0

**Response:**

    {
     "block": "2680262203532249785",
     "requestProcessingTime": 1
    }

<small>*Verified 11-Nov-14*</small>

### Get Blocks

**Request:**

    http://localhost:8125/burst?
      requestType=getBlocks&
      lastIndex=1

**Response:**

    {
     "blocks": [
      {
       "previousBlockHash": "f88c75a36317e1795348330cb9a944f33153b517ebdf05d9f3f9a606e997618d",
       "payloadLength": 981,
       "totalAmountNQT": "0",
       "generationSignature": "02f7462b62270c0028c379d838d3a192cf0b782995f3bb1929a5378d26e7e8a9",
       "generator": "2218289317977832095",
       "generatorPublicKey": "98ccf5d5173b13e4c9eab2631372f61ce8ba506db559d73b285073a689872e75",
       "baseTarget": "151761236",
       "payloadHash": "82f8c8ca1f8f2252172ed1e9836d5228432fcb18aecfa7d55a119efcd242321c",
       "generatorRS": "BURST-TGNZ-E8VK-69EX-3L9LX",
       "numberOfTransactions": 5,
       "blockSignature": "76a3f7f966256c4985262fb4622190b2b9a19b900f6ce443ab7d581e3176c...",
       "transactions": [
        "10545999940082849455",
        "16751328983055099280",
        "1257496316971695605",
        "2048859884870801838",
        "5105170273384355243"
       ],
       "version": 3,
       "totalFeeNQT": "500000000",
       "previousBlock": "8782326465060769016",
       "block": "7299310714263322546",
       "height": 281683,
       "timestamp": 30416681
      },
      {
       "previousBlockHash": "edc2d65d24883b9b32c46da4eec3792a69a5ff9a9b1e629c4e7e0224432c87c9",
       "payloadLength": 0,
       "totalAmountNQT": "0",
       "generationSignature": "8c4944c7a9ef5700d1b89660fdf83bcbd3dbbfbe191fef005d0f248a573816fe",
       "generator": "15766845356521829337",
       "generatorPublicKey": "11636697faf4ade736cedf6c528bec0142353c4d93fce05cdb818c49e0390422",
       "baseTarget": "303522472",
       "payloadHash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
       "generatorRS": "BURST-2EYT-C522-VL6G-FQSMP",
       "nextBlock": "7299310714263322546",
       "numberOfTransactions": 0,
       "blockSignature": "90464d183ed01cb0930ef527e746eaa092d1558697c369cc9246add5f6eec...",
       "transactions": [],
       "version": 3,
       "totalFeeNQT": "0",
       "previousBlock": "11185683789279314669",
       "block": "8782326465060769016",
       "height": 281682,
       "timestamp": 30416671
      }
     ],
     "requestProcessingTime": 2
    }

<small>*Verified 11-Nov-14*</small>

### Get EC Block

**Request:**

    http://localhost:8125/burst?
      requestType=getECBlock

**Response:**

    {
     "ecBlockHeight": 281777,
     "requestProcessingTime": 2,
     "ecBlockId": "6565813579609649593",
     "timestamp": 30427868
    }

<small>*Verified 11-Nov-14*</small>

Digital Goods Store Operations
------------------------------

### DGS Delisting

**Request:**

    http://localhost:8125/burst?
      requestType=dgsDelisting&
      goods=11813734897437346473
      secretPhrase=IWontTellYou&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "87dd2aa00690acaa5da4e8cd9db6ab94b5ed8bf0f8fe2c2ba6a39c57d3073105",
     "unsignedTransactionBytes": "03111852e2013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "39ee52204d0a32c0bbd0e4bc6a6fb52299cd7d16ab74e9a040f857d3b1c3e...",
      "feeNQT": "100000000",
      "type": 3,
      "fullHash": "40e64c357e240f9b7ca5780e757e34d48a58ae93eeb19e62d020d5719b43e2b6",
      "version": 1,
      "ecBlockId": "5345754442518111082",
      "signatureHash": "87dd2aa00690acaa5da4e8cd9db6ab94b5ed8bf0f8fe2c2ba6a39c57d3073105",
      "attachment": {
       "version.DigitalGoodsDelisting": 1,
       "goods": "11813734897437346473"
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 1,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "ecBlockHeight": 292903,
      "deadline": 60,
      "transaction": "11173189325008201280",
      "timestamp": 31609368,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 5188,
     "transactionBytes": "03111852e2013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "fullHash": "40e64c357e240f9b7ca5780e757e34d48a58ae93eeb19e62d020d5719b43e2b6",
     "transaction": "11173189325008201280"
    }

<small>*Verified 25-Nov-14*</small>

### DGS Delivery

**Request:**

    http://localhost:8125/burst?
      requestType=dgsDelivery&
      purchase=3723760852542296589&
      goodsToEncrypt=Download Code.&
      secretPhrase=IWontTellYou&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "f3f1f1bf30c7a6f5a7776aa23502dfe5240efe0d5c016760f8e8a152c4eb1b1b",
     "unsignedTransactionBytes": "0315b20fe1013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "b96c42161ce38177168de5a8906ce747a5415017df5b16cdce1c015cb6e1f10b...",
      "feeNQT": "100000000",
      "type": 3,
      "fullHash": "349926ea025b627117c05fbe3ff298e97a8efdb07f7ae552b5a91e7112c5b82b",
      "version": 1,
      "ecBlockId": "15813950788992376142",
      "signatureHash": "f3f1f1bf30c7a6f5a7776aa23502dfe5240efe0d5c016760f8e8a152c4eb1b1b",
      "attachment": {
       "goodsIsText": true,
       "discountNQT": "0",
       "purchase": "3723760852542296589",
       "goodsData": "aef91bfe543844964b0dadc51c6fe4c624b6b0c72c23c5c61f7e626f51e15...",
       "version.DigitalGoodsDelivery": 1,
       "goodsNonce": "4cb474ec53a590c6fd20d050c429d009ddefd286b2782eb3a0782ae45c2bbbbd"
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 5,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "recipientRS": "BURST-6GMG-FC5F-YSX6-8CVEL",
      "recipient": "7580519603555678830",
      "ecBlockHeight": 292069,
      "deadline": 60,
      "transaction": "8170192742079961396",
      "timestamp": 31526834,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 8832,
     "transactionBytes": "0315b20fe1013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "fullHash": "349926ea025b627117c05fbe3ff298e97a8efdb07f7ae552b5a91e7112c5b82b",
     "transaction": "8170192742079961396"
    }

<small>*Verified 24-Nov-14*</small>

### DGS Feedback

**Request:**

    http://localhost:8125/burst?
      requestType=dgsFeedback&
      purchase=10234639413366748292&
      secretPhrase=IWontTellYou&
      feeNQT=100000000&
      deadline=60&
      message=Thank You!

**Response:**

    {
     "signatureHash": "bf31e4e3b2037bfe223efd72519e3a2d25cd1e0dc1a62e3eeefe59179e16a7a4",
     "unsignedTransactionBytes": "0316a463e2013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "a3ee76ce4bec0328fe5cfc1a2adfc229a8c9959a81b7912664de51156987c...",
      "feeNQT": "100000000",
      "type": 3,
      "fullHash": "491d94b9031b7a7a32d9ab5d0491b81364941eb080ee9029de58cb6ffa0161f6",
      "version": 1,
      "ecBlockId": "5484311742753527844",
      "signatureHash": "bf31e4e3b2037bfe223efd72519e3a2d25cd1e0dc1a62e3eeefe59179e16a7a4",
      "attachment": {
       "purchase": "10234639413366748292",
       "version.Message": 1,
       "messageIsText": true,
       "version.DigitalGoodsFeedback": 1,
       "message": "Thank you!"
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 6,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "recipientRS": "BURST-6GMG-FC5F-YSX6-8CVEL",
      "recipient": "7580519603555678830",
      "ecBlockHeight": 292938,
      "deadline": 60,
      "transaction": "8825396122598251849",
      "timestamp": 31613860,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 9343,
     "transactionBytes": "0316a463e2013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "fullHash": "491d94b9031b7a7a32d9ab5d0491b81364941eb080ee9029de58cb6ffa0161f6",
     "transaction": "8825396122598251849"
    }

<small>*Verified 25-Nov-14*</small>

### DGS Listing

**Request:**

    http://localhost:8125/burst?
      requestType=dgsListing&
      secretPhrase=IWontTellYou&
      name=Test Product&
      description=Testing the DGS.
      tags=test, product, tag, extra&
      quantity=3&
      priceNQT=100000000&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "565bc0a6140ae1331cd5db009fbd9da164d8802330939ef40204a9bc343b3149",
     "unsignedTransactionBytes": "031092aedf013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "465ed3dfef9a02fc97fc18a6f83bb9f07c285aef41ff78c957d59cda2972ba...",
      "feeNQT": "100000000",
      "type": 3,
      "fullHash": "a98a63204cd1f2a3304cee79776854f290f0472883c576056fec16a23efa90df",
      "version": 1,
      "ecBlockId": "4480409615309425420",
      "signatureHash": "565bc0a6140ae1331cd5db009fbd9da164d8802330939ef40204a9bc343b3149",
      "attachment": {
       "priceNQT": "100000000",
       "quantity": 3,
       "name": "Test Product",
       "description": "Testing the DGS.",
       "version.DigitalGoodsListing": 1,
       "tags": "test, product, tag, extra"
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 0,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "ecBlockHeight": 291240,
      "deadline": 60,
      "transaction": "11813734897437346473",
      "timestamp": 31436434,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 8958,
     "transactionBytes": "031092aedf013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473b...",
     "fullHash": "a98a63204cd1f2a3304cee79776854f290f0472883c576056fec16a23efa90df",
     "transaction": "11813734897437346473"
    }

<small>*Verified 23-Nov-14*</small>

### DGS Price Change

**Request:**

    http://localhost:8125/burst?
      requestType=dgsPriceChange&
      goods=11813734897437346473&
      priceNQT=200000000&
      secretPhrase=IWontTellYou&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "aad88476935eda8ca3a9190163b83106e6ca95733ddfa274d29a0378d773cc3c",
     "unsignedTransactionBytes": "0312cebbdf013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "533f9c330100cab8a7a0f9375b49a09f6432b0e029660ab715a728ac75e6c2...",
      "feeNQT": "100000000",
      "type": 3,
      "fullHash": "1964e4fe11f4a00a96c79e274f9da21f64d7fe10ed77a154975b92d3c65f1287",
      "version": 1,
      "ecBlockId": "14411471768450948944",
      "signatureHash": "aad88476935eda8ca3a9190163b83106e6ca95733ddfa274d29a0378d773cc3c",
      "attachment": {
       "goods": "11813734897437346473",
       "priceNQT": "200000000",
       "version.DigitalGoodsPriceChange": 1
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 2,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "ecBlockHeight": 291259,
      "deadline": 60,
      "transaction": "765880294780986393",
      "timestamp": 31439822,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 6444,
     "transactionBytes": "0312cebbdf013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "fullHash": "1964e4fe11f4a00a96c79e274f9da21f64d7fe10ed77a154975b92d3c65f1287",
     "transaction": "765880294780986393"
    }

<small>*Verified 23-Nov-14*</small>

### DGS Purchase

**Request:**

    http://localhost:8125/burst?
      requestType=dgsPurchase&
      goods=1587116104511359906&
      quantity=1&
      deliveryDeadlineTimestamp=31800000&
      secretPhrase=IWontTellYou&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "03e112d8707ae33ec37ff4405f31920c0e67c58439e33033c248b3eaca81d2c3",
     "unsignedTransactionBytes": "0314805be2013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "a2739ef63024d9d3c61bebbb2692e7a2092666e13fc380c2e8b29c3a3fc24...",
      "feeNQT": "100000000",
      "type": 3,
      "fullHash": "84f4ef2d52be088e011e5b7857fbab88665918e0df02102a8333aee53c3bb88b",
      "version": 1,
      "ecBlockId": "3186563001195424357",
      "signatureHash": "03e112d8707ae33ec37ff4405f31920c0e67c58439e33033c248b3eaca81d2c3",
      "attachment": {
       "goods": "1587116104511359906",
       "priceNQT": "100000000",
       "quantity": 1,
       "deliveryDeadlineTimestamp": 31800000,
       "version.DigitalGoodsPurchase": 1
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 4,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "recipientRS": "BURST-6GMG-FC5F-YSX6-8CVEL",
      "recipient": "7580519603555678830",
      "ecBlockHeight": 292925,
      "deadline": 60,
      "transaction": "10234639413366748292",
      "timestamp": 31611776,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 5536,
     "transactionBytes": "0314805be2013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "fullHash": "84f4ef2d52be088e011e5b7857fbab88665918e0df02102a8333aee53c3bb88b",
     "transaction": "10234639413366748292"
    }

<small>*Verified 25-Nov-14*</small>

### DGS Quantity Change

**Request:**

    http://localhost:8125/burst?
      requestType=dgsQuantityChange&
      goods=11813734897437346473&
      deltaQuantity=-1&
      secretPhrase=IWontTellYou&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "5c60fc94452980f88423c3a16a74d49209d246d0f7454024c87cc76ea3221500",
     "unsignedTransactionBytes": "031334b7df013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "7eaf8e48a780f51aa7e01c01091c9388b72a1bf075dddbd3594d55ad169272...",
      "feeNQT": "100000000",
      "type": 3,
      "fullHash": "f3cd9046bba3706dd6f63cf387f23c7c29532ff15e3423f57900613c844a646f",
      "version": 1,
      "ecBlockId": "14589429051005044326",
      "signatureHash": "5c60fc94452980f88423c3a16a74d49209d246d0f7454024c87cc76ea3221500",
      "attachment": {
       "goods": "11813734897437346473",
       "version.DigitalGoodsQuantityChange": 1,
       "deltaQuantity": -1
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 3,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "ecBlockHeight": 291249,
      "deadline": 60,
      "transaction": "7885982972263845363",
      "timestamp": 31438644,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 7474,
     "transactionBytes": "031334b7df013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "fullHash": "f3cd9046bba3706dd6f63cf387f23c7c29532ff15e3423f57900613c844a646f",
     "transaction": "7885982972263845363"
    }

<small>*Verified 23-Nov-14*</small>

### DGS Refund

**Request:**

    http://localhost:8125/burst?
      requestType=dgsRefund&
      purchase=3723760852542296589&
      refundNQT=100000000&
      secretPhrase=IWontTellYou&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "dafa980759db6bd6e5f11aa7bdc1cff1bb201be750328c397870f1758cef975c",
     "unsignedTransactionBytes": "0317944ee2013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "43a90c1f0dde1bca84623e69b0abe0284883600d7bfb5286b3489a9ee417...",
      "feeNQT": "100000000",
      "type": 3,
      "fullHash": "64c2af2811da3306f4671002ed1f12655fca1937a0b68164da2a71f3ee63adce",
      "version": 1,
      "ecBlockId": "11418000967717599433",
      "signatureHash": "dafa980759db6bd6e5f11aa7bdc1cff1bb201be750328c397870f1758cef975c",
      "attachment": {
       "purchase": "3723760852542296589",
       "version.DigitalGoodsRefund": 1,
       "refundNQT": "100000000"
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 7,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "recipientRS": "BURST-6GMG-FC5F-YSX6-8CVEL",
      "recipient": "7580519603555678830",
      "ecBlockHeight": 292899,
      "deadline": 60,
      "transaction": "446940555271717476",
      "timestamp": 31608468,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 10509,
     "transactionBytes": "0317944ee2013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "fullHash": "64c2af2811da3306f4671002ed1f12655fca1937a0b68164da2a71f3ee63adce",
     "transaction": "446940555271717476"
    }

<small>*Verified 25-Nov-14*</small>

### Get DGS Good

**Request:**

    http://localhost:8125/burst?
      requestType=getDGSGood&
      goods=11813734897437346473

**Response:**

    {
     "seller": "15323192282528158131",
     "quantity": 3,
     "goods": "11813734897437346473",
     "description": "Testing the DGS.",
     "sellerRS": "BURST-L6FM-89WK-VK8P-FCRBB",
     "requestProcessingTime": 1,
     "delisted": false,
     "parsedTags": [
      "test",
      "product",
      "tag"
     ],
     "tags": "test, product, tag, extra",
     "priceNQT": "100000000",
     "numberOfPublicFeedbacks": 0,
     "name": "Test Product",
     "numberOfPurchases": 0,
     "timestamp": 31436434
    }

<small>*Verified 23-Nov-14*</small>

### Get DGS Goods

**Request:**

    http://localhost:8125/burst?
      requestType=getDGSGoods&
      seller=BURST-L6FM-89WK-VK8P-FCRBB

**Response:**

    {
     "goods": [
      {
       "seller": "15323192282528158131",
       "quantity": 2,
       "goods": "11813734897437346473",
       "description": "Testing the DGS.",
       "sellerRS": "BURST-L6FM-89WK-VK8P-FCRBB",
       "delisted": false,
       "parsedTags": [
        "test",
        "product",
        "tag"
       ],
       "tags": "test, product, tag, extra",
       "priceNQT": "200000000",
       "numberOfPublicFeedbacks": 0,
       "name": "Test Product",
       "numberOfPurchases": 0,
       "timestamp": 31436434
      }
     ],
     "requestProcessingTime": 46
    }

<small>*Verified 23-Nov-14*</small>

### Get DGS Goods Purchases

**Request:**

    http://localhost:8125/burst?
      requestType=getDGSGoodsPurchases&
      goods=11813734897437346473

**Response:**

    {
     "purchases": [
      {
       "seller": "15323192282528158131",
       "quantity": 2,
       "pending": false,
       "purchase": "3723760852542296589",
       "goods": "11813734897437346473",
       "sellerRS": "BURST-L6FM-89WK-VK8P-FCRBB",
       "buyer": "7580519603555678830",
       "priceNQT": "200000000",
       "deliveryDeadlineTimestamp": 31600000,
       "goodsIsText": false,
       "buyerRS": "BURST-6GMG-FC5F-YSX6-8CVEL",
       "name": "Test Product",
       "goodsData": {
        "data": "aef91bfe543844964b0dadc51c6fe4c624b6b0c72c23c5c61f7e626f51e15...",
        "nonce": "4cb474ec53a590c6fd20d050c429d009ddefd286b2782eb3a0782ae45c2bbbbd"
       },
       "timestamp": 31520720
      }
     ],
     "requestProcessingTime": 305
    }

<small>*Verified 15-Dec-14*</small>

### Get DGS Pending Purchases

**Request:**

    http://localhost:8125/burst?
      requestType=getDGSPendingPurchases&
      seller=BURST-L6FM-89WK-VK8P-FCRBB

**Response:**

    {
     "purchases": [
      {
       "seller": "15323192282528158131",
       "priceNQT": "200000000",
       "quantity": 2,
       "deliveryDeadlineTimestamp": 31600000,
       "buyerRS": "BURST-6GMG-FC5F-YSX6-8CVEL",
       "pending": true,
       "purchase": "3723760852542296589",
       "name": "Test Product",
       "goods": "11813734897437346473",
       "sellerRS": "BURST-L6FM-89WK-VK8P-FCRBB",
       "buyer": "7580519603555678830",
       "timestamp": 31520720
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 24-Nov-14*</small>

### Get DGS Purchase

**Request:**

    http://localhost:8125/burst?
      requestType=getDGSPurchase&
      purchase=10234639413366748292

**Response:**

    {
     "seller": "7580519603555678830",
     "quantity": 1,
     "feedbackNotes": [
      {
       "data": "7086a82f4da0708d4eaa9b16d5fc5a25c556596ea29d957d0a1dddd0a482c...",
       "nonce": "c521481ce67f7778c41c6716806047d4ea641005392cd7e5ce8d20c49623dad8"
      }
     ],
     "publicFeedbacks": [
      "Thank you again!",
      "Thank you!"
     ],
     "pending": false,
     "purchase": "10234639413366748292",
     "goods": "1587116104511359906",
     "sellerRS": "BURST-6GMG-FC5F-YSX6-8CVEL",
     "requestProcessingTime": 1,
     "buyer": "15323192282528158131",
     "priceNQT": "100000000",
     "deliveryDeadlineTimestamp": 31800000,
     "goodsIsText": false,
     "buyerRS": "BURST-L6FM-89WK-VK8P-FCRBB",
     "discountNQT": "100000000",
     "name": "Test Product",
     "goodsData": {
      "data": "5f4868022381aa9532614a7aae1600e59e84c80571add107dabca891df97e7...",
      "nonce": "c017b846de4375741ebc9f3bff894270d218ff4090a66dfd505770ccdc2f54bd"
     },
     "timestamp": 31611776
    }

<small>*Verified 24-Nov-14*</small>

### Get DGS Purchases

**Request:**

    http://localhost:8125/burst?
      requestType=getDGSPurchases&
      seller=BURST-L6FM-89WK-VK8P-FCRBB&
      lastIndex=0

**Response:**

    {
     "purchases": [
      {
       "seller": "15323192282528158131",
       "priceNQT": "200000000",
       "quantity": 2,
       "deliveryDeadlineTimestamp": 31600000,
       "buyerRS": "BURST-6GMG-FC5F-YSX6-8CVEL",
       "pending": true,
       "purchase": "3723760852542296589",
       "name": "Test Product",
       "goods": "11813734897437346473",
       "sellerRS": "BURST-L6FM-89WK-VK8P-FCRBB",
       "buyer": "7580519603555678830",
       "timestamp": 31520720
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 15-Dec-14*</small>

Forging Operations
------------------

### Lease Balance

**Request:**

    http://localhost:8125/burst?
      requestType=leaseBalance&
      period=1440&
      recipient=BURST-4VNQ-RWZC-4WWQ-GVM8S&
      secretPhrase=IWontTellYou&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "462590bb85b615ddd34d59a1ffdc452a9baee2088044c1b9eb44298e49158f35",
     "unsignedTransactionBytes": "04107977ce013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473b...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "4a89702535246fd55bfc4b4691dc4266daa3100e00cdf0caaed57a5ad750da075a...",
      "feeNQT": "100000000",
      "type": 4,
      "fullHash": "251bcd86057c09bb2d055bbeeb9b67ccae861f75d4aada21b3b79a13db9712e1",
      "version": 1,
      "ecBlockId": "6565389899781382679",
      "signatureHash": "462590bb85b615ddd34d59a1ffdc452a9baee2088044c1b9eb44298e49158f35",
      "attachment": {
       "version.EffectiveBalanceLeasing": 1,
       "period": 1440
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 0,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "recipientRS": "BURST-4VNQ-RWZC-4WWQ-GVM8S",
      "recipient": "17013046603665206934",
      "ecBlockHeight": 280672,
      "deadline": 60,
      "transaction": "13477439723061189413",
      "timestamp": 30308217,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 7340,
     "transactionBytes": "04107977ce013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473b...",
     "fullHash": "251bcd86057c09bb2d055bbeeb9b67ccae861f75d4aada21b3b79a13db9712e1",
     "transaction": "13477439723061189413"
    }

<small>*Verified 10-Nov-14*</small>

Hallmark Operations
-------------------

### Decode Hallmark

**Request:**

    http://localhost:8125/burst?
      requestType=decodeHallmark&
      hallmark=827ed8cf83bbd36419002759d960a1bec952f2209db8ed3be958...

**Response:**

    {
     "date": "2013-12-10",
     "valid": true,
     "accountRS": "BURST-FEJ5-H4YB-QD3D-2L3W2",
     "host": "nxttyclub.info",
     "weight": 100,
     "requestProcessingTime": 2,
     "account": "31580691533050371"
    }

<small>*Verified 22-Nov-14*</small>

### Mark Host

**Request:**

    http://localhost:8125/burst?
      requestType=markHost&
      secretPhrase=IWontTellYou&
      host=iwonttellyou.com&
      weight=1&
      date=2014-11-22

**Response:**

    {
     "hallmark": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c...",
     "requestProcessingTime": 5
    }

<small>*Verified 22-Nov-14*</small>

Networking Operations
---------------------

### Get My Info

**Request:**

    http://localhost:8125/burst?
      requestType=getMyInfo

**Response:**

    {
     "address": "127.0.0.1",
     "host": "127.0.0.1",
     "requestProcessingTime": 1
    }

<small>*Verified 20-Nov-14*</small>

### Get Peer

**Request:**

    http://localhost:8125/burst?
      requestType=getPeer&
      peer=nxtx.ru

**Response:**

    {
     "hallmark": "5f107c33097fc8241f55b9eb23615fe80ed3f87c87dc8c23ed6a08ab9273417a07006...",
     "downloadedVolume": 352,
     "address": "188.226.174.169",
     "inbound": false,
     "weight": 53,
     "uploadedVolume": 2165,
     "requestProcessingTime": 30,
     "version": "1.5.11",
     "platform": "VPS",
     "inboundWebSocket": false,
     "lastUpdated": 48540348,
     "blacklisted": false,
     "announcedAddress": "nxtx.ru",
     "application": "BRS",
     "port": 7874,
     "outboundWebSocket": true,
     "lastConnectAttempt": 48540348,
     "state": 1,
     "shareAddress": true
    }

<small>*Verified 9-Jun-15*</small>

### Get Peers

**Request:**

    http://localhost:8125/burst?
      requestType=getPeers&
      state=DISCONNECTED

**Response:**

    {
     "peers": [
      "198.50.146.93",
      "213.46.57.77"
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 15-Dec-14*</small>

Server Information Operations
-----------------------------

### Get Blockchain Status

**Request:**

    http://localhost:8125/burst?
      requestType=getBlockchainStatus

**Response:**

    {
     "currentMinRollbackHeight": 445745,
     "numberOfBlocks": 446546,
     "isTestnet": false,
     "includeExpiredPrunable": false,
     "requestProcessingTime": 1,
     "version": "1.5.11",
     "maxRollback": 800,
     "lastBlock": "2164693711802180410",
     "application": "BRS",
     "isScanning": false,
     "isDownloading": false,
     "cumulativeDifficulty": "17966254519242206",
     "lastBlockchainFeederHeight": 446597,
     "maxPrunableLifetime": 1209600,
     "time": 48539846,
     "lastBlockchainFeeder": "84.253.125.186"
    }

<small>*Verified 9-Jun-15*</small>

### Get Constants

**Request:**

    http://localhost:8125/burst?
      requestType=getConstants

**Response:**

    {
      "transactionSubTypes":{
        "ReserveClaim":{
          "isPhasable":true,
          "subtype":2,
          "mustHaveRecipient":false,
          "name":"ReserveClaim",
          "canHaveRecipient":false,
          "type":5,
          "isPhasingSafe":false
        },
        "AssetIssuance":{
          "isPhasable":true,
          "subtype":0,
          "mustHaveRecipient":false,
          "name":"AssetIssuance",
          "canHaveRecipient":false,
          "type":2,
          "isPhasingSafe":true
        },
        "PhasingVoteCasting":{
          "isPhasable":true,
          "subtype":9,
          "mustHaveRecipient":false,
          "name":"PhasingVoteCasting",
          "canHaveRecipient":false,
          "type":1,
          "isPhasingSafe":true
        },
        "CurrencyIssuance":{
          "isPhasable":true,
          "subtype":0,
          "mustHaveRecipient":false,
          "name":"CurrencyIssuance",
          "canHaveRecipient":false,
          "type":5,
          "isPhasingSafe":false
        },
        "ShufflingRecipients":{
          "isPhasable":false,
          "subtype":3,
          "mustHaveRecipient":false,
          "name":"ShufflingRecipients",
          "canHaveRecipient":false,
          "type":7,
          "isPhasingSafe":false
        },
        "PublishExchangeOffer":{
          "isPhasable":true,
          "subtype":4,
          "mustHaveRecipient":false,
          "name":"PublishExchangeOffer",
          "canHaveRecipient":false,
          "type":5,
          "isPhasingSafe":false
        },
        "TaggedDataUpload":{
          "isPhasable":false,
          "subtype":0,
          "mustHaveRecipient":false,
          "name":"TaggedDataUpload",
          "canHaveRecipient":false,
          "type":6,
          "isPhasingSafe":false
        },
        "ReserveIncrease":{
          "isPhasable":true,
          "subtype":1,
          "mustHaveRecipient":false,
          "name":"ReserveIncrease",
          "canHaveRecipient":false,
          "type":5,
          "isPhasingSafe":false
        },
        "ArbitraryMessage":{
          "isPhasable":true,
          "subtype":0,
          "mustHaveRecipient":false,
          "name":"ArbitraryMessage",
          "canHaveRecipient":true,
          "type":1,
          "isPhasingSafe":false
        },
        "DigitalGoodsPriceChange":{
          "isPhasable":true,
          "subtype":2,
          "mustHaveRecipient":false,
          "name":"DigitalGoodsPriceChange",
          "canHaveRecipient":false,
          "type":3,
          "isPhasingSafe":false
        },
        "DigitalGoodsFeedback":{
          "isPhasable":true,
          "subtype":6,
          "mustHaveRecipient":true,
          "name":"DigitalGoodsFeedback",
          "canHaveRecipient":true,
          "type":3,
          "isPhasingSafe":false
        },
        "ExchangeBuy":{
          "isPhasable":true,
          "subtype":5,
          "mustHaveRecipient":false,
          "name":"ExchangeBuy",
          "canHaveRecipient":false,
          "type":5,
          "isPhasingSafe":false
        },
        "VoteCasting":{
          "isPhasable":true,
          "subtype":3,
          "mustHaveRecipient":false,
          "name":"VoteCasting",
          "canHaveRecipient":false,
          "type":1,
          "isPhasingSafe":false
        },
        "SetPhasingOnly":{
          "isPhasable":true,
          "subtype":1,
          "mustHaveRecipient":false,
          "name":"SetPhasingOnly",
          "canHaveRecipient":false,
          "type":4,
          "isPhasingSafe":false
        },
        "CurrencyMinting":{
          "isPhasable":true,
          "subtype":7,
          "mustHaveRecipient":false,
          "name":"CurrencyMinting",
          "canHaveRecipient":false,
          "type":5,
          "isPhasingSafe":false
        },
        "AliasAssignment":{
          "isPhasable":true,
          "subtype":1,
          "mustHaveRecipient":false,
          "name":"AliasAssignment",
          "canHaveRecipient":false,
          "type":1,
          "isPhasingSafe":false
        },
        "AliasDelete":{
          "isPhasable":true,
          "subtype":8,
          "mustHaveRecipient":false,
          "name":"AliasDelete",
          "canHaveRecipient":false,
          "type":1,
          "isPhasingSafe":false
        },
        "BidOrderCancellation":{
          "isPhasable":true,
          "subtype":5,
          "mustHaveRecipient":false,
          "name":"BidOrderCancellation",
          "canHaveRecipient":false,
          "type":2,
          "isPhasingSafe":true
        },
        "AssetDelete":{
          "isPhasable":true,
          "subtype":7,
          "mustHaveRecipient":false,
          "name":"AssetDelete",
          "canHaveRecipient":false,
          "type":2,
          "isPhasingSafe":true
        },
        "ShufflingProcessing":{
          "isPhasable":false,
          "subtype":2,
          "mustHaveRecipient":false,
          "name":"ShufflingProcessing",
          "canHaveRecipient":false,
          "type":7,
          "isPhasingSafe":false
        },
        "DigitalGoodsListing":{
          "isPhasable":true,
          "subtype":0,
          "mustHaveRecipient":false,
          "name":"DigitalGoodsListing",
          "canHaveRecipient":false,
          "type":3,
          "isPhasingSafe":true
        },
        "AskOrderCancellation":{
          "isPhasable":true,
          "subtype":4,
          "mustHaveRecipient":false,
          "name":"AskOrderCancellation",
          "canHaveRecipient":false,
          "type":2,
          "isPhasingSafe":true
        },
        "DigitalGoodsPurchase":{
          "isPhasable":true,
          "subtype":4,
          "mustHaveRecipient":true,
          "name":"DigitalGoodsPurchase",
          "canHaveRecipient":true,
          "type":3,
          "isPhasingSafe":false
        },
        "AccountInfo":{
          "isPhasable":true,
          "subtype":5,
          "mustHaveRecipient":false,
          "name":"AccountInfo",
          "canHaveRecipient":false,
          "type":1,
          "isPhasingSafe":true
        },
        "CurrencyTransfer":{
          "isPhasable":true,
          "subtype":3,
          "mustHaveRecipient":true,
          "name":"CurrencyTransfer",
          "canHaveRecipient":true,
          "type":5,
          "isPhasingSafe":false
        },
        "HubAnnouncement":{
          "isPhasable":true,
          "subtype":4,
          "mustHaveRecipient":false,
          "name":"HubAnnouncement",
          "canHaveRecipient":false,
          "type":1,
          "isPhasingSafe":true
        },
        "DigitalGoodsQuantityChange":{
          "isPhasable":true,
          "subtype":3,
          "mustHaveRecipient":false,
          "name":"DigitalGoodsQuantityChange",
          "canHaveRecipient":false,
          "type":3,
          "isPhasingSafe":false
        },
        "DigitalGoodsRefund":{
          "isPhasable":true,
          "subtype":7,
          "mustHaveRecipient":true,
          "name":"DigitalGoodsRefund",
          "canHaveRecipient":true,
          "type":3,
          "isPhasingSafe":false
        },
        "AssetTransfer":{
          "isPhasable":true,
          "subtype":1,
          "mustHaveRecipient":true,
          "name":"AssetTransfer",
          "canHaveRecipient":true,
          "type":2,
          "isPhasingSafe":true
        },
        "AliasSell":{
          "isPhasable":true,
          "subtype":6,
          "mustHaveRecipient":false,
          "name":"AliasSell",
          "canHaveRecipient":true,
          "type":1,
          "isPhasingSafe":false
        },
        "AccountProperty":{
          "isPhasable":true,
          "subtype":10,
          "mustHaveRecipient":true,
          "name":"AccountProperty",
          "canHaveRecipient":true,
          "type":1,
          "isPhasingSafe":true
        },
        "AccountPropertyDelete":{
          "isPhasable":true,
          "subtype":11,
          "mustHaveRecipient":true,
          "name":"AccountPropertyDelete",
          "canHaveRecipient":true,
          "type":1,
          "isPhasingSafe":true
        },
        "AliasBuy":{
          "isPhasable":true,
          "subtype":7,
          "mustHaveRecipient":true,
          "name":"AliasBuy",
          "canHaveRecipient":true,
          "type":1,
          "isPhasingSafe":false
        },
        "EffectiveBalanceLeasing":{
          "isPhasable":true,
          "subtype":0,
          "mustHaveRecipient":true,
          "name":"EffectiveBalanceLeasing",
          "canHaveRecipient":true,
          "type":4,
          "isPhasingSafe":true
        },
        "ShufflingRegistration":{
          "isPhasable":true,
          "subtype":1,
          "mustHaveRecipient":false,
          "name":"ShufflingRegistration",
          "canHaveRecipient":false,
          "type":7,
          "isPhasingSafe":false
        },
        "ShufflingVerification":{
          "isPhasable":false,
          "subtype":4,
          "mustHaveRecipient":false,
          "name":"ShufflingVerification",
          "canHaveRecipient":false,
          "type":7,
          "isPhasingSafe":false
        },
        "DividendPayment":{
          "isPhasable":true,
          "subtype":6,
          "mustHaveRecipient":false,
          "name":"DividendPayment",
          "canHaveRecipient":false,
          "type":2,
          "isPhasingSafe":true
        },
        "TaggedDataExtend":{
          "isPhasable":false,
          "subtype":1,
          "mustHaveRecipient":false,
          "name":"TaggedDataExtend",
          "canHaveRecipient":false,
          "type":6,
          "isPhasingSafe":false
        },
        "ShufflingCreation":{
          "isPhasable":true,
          "subtype":0,
          "mustHaveRecipient":false,
          "name":"ShufflingCreation",
          "canHaveRecipient":false,
          "type":7,
          "isPhasingSafe":false
        },
        "ShufflingCancellation":{
          "isPhasable":false,
          "subtype":5,
          "mustHaveRecipient":false,
          "name":"ShufflingCancellation",
          "canHaveRecipient":false,
          "type":7,
          "isPhasingSafe":false
        },
        "BidOrderPlacement":{
          "isPhasable":true,
          "subtype":3,
          "mustHaveRecipient":false,
          "name":"BidOrderPlacement",
          "canHaveRecipient":false,
          "type":2,
          "isPhasingSafe":true
        },
        "ExchangeSell":{
          "isPhasable":true,
          "subtype":6,
          "mustHaveRecipient":false,
          "name":"ExchangeSell",
          "canHaveRecipient":false,
          "type":5,
          "isPhasingSafe":false
        },
        "DigitalGoodsDelisting":{
          "isPhasable":true,
          "subtype":1,
          "mustHaveRecipient":false,
          "name":"DigitalGoodsDelisting",
          "canHaveRecipient":false,
          "type":3,
          "isPhasingSafe":true
        },
        "AskOrderPlacement":{
          "isPhasable":true,
          "subtype":2,
          "mustHaveRecipient":false,
          "name":"AskOrderPlacement",
          "canHaveRecipient":false,
          "type":2,
          "isPhasingSafe":true
        },
        "DigitalGoodsDelivery":{
          "isPhasable":true,
          "subtype":5,
          "mustHaveRecipient":true,
          "name":"DigitalGoodsDelivery",
          "canHaveRecipient":true,
          "type":3,
          "isPhasingSafe":false
        },
        "PollCreation":{
          "isPhasable":true,
          "subtype":2,
          "mustHaveRecipient":false,
          "name":"PollCreation",
          "canHaveRecipient":false,
          "type":1,
          "isPhasingSafe":false
        },
        "OrdinaryPayment":{
          "isPhasable":true,
          "subtype":0,
          "mustHaveRecipient":true,
          "name":"OrdinaryPayment",
          "canHaveRecipient":true,
          "type":0,
          "isPhasingSafe":true
        },
        "CurrencyDeletion":{
          "isPhasable":true,
          "subtype":8,
          "mustHaveRecipient":false,
          "name":"CurrencyDeletion",
          "canHaveRecipient":false,
          "type":5,
          "isPhasingSafe":false
        }
      },
      "genesisAccountId":"1739068987193023818",
      "genesisBlockId":"2680262203532249785",
      "transactionTypes":{
        "0":{
          "subtypes":{
            "0":{
              "isPhasable":true,
              "subtype":0,
              "mustHaveRecipient":true,
              "name":"OrdinaryPayment",
              "canHaveRecipient":true,
              "type":0,
              "isPhasingSafe":true
            }
          }
        },
        "1":{
          "subtypes":{
            "0":{
              "isPhasable":true,
              "subtype":0,
              "mustHaveRecipient":false,
              "name":"ArbitraryMessage",
              "canHaveRecipient":true,
              "type":1,
              "isPhasingSafe":false
            },
            "1":{
              "isPhasable":true,
              "subtype":1,
              "mustHaveRecipient":false,
              "name":"AliasAssignment",
              "canHaveRecipient":false,
              "type":1,
              "isPhasingSafe":false
            },
            "2":{
              "isPhasable":true,
              "subtype":2,
              "mustHaveRecipient":false,
              "name":"PollCreation",
              "canHaveRecipient":false,
              "type":1,
              "isPhasingSafe":false
            },
            "3":{
              "isPhasable":true,
              "subtype":3,
              "mustHaveRecipient":false,
              "name":"VoteCasting",
              "canHaveRecipient":false,
              "type":1,
              "isPhasingSafe":false
            },
            "4":{
              "isPhasable":true,
              "subtype":4,
              "mustHaveRecipient":false,
              "name":"HubAnnouncement",
              "canHaveRecipient":false,
              "type":1,
              "isPhasingSafe":true
            },
            "5":{
              "isPhasable":true,
              "subtype":5,
              "mustHaveRecipient":false,
              "name":"AccountInfo",
              "canHaveRecipient":false,
              "type":1,
              "isPhasingSafe":true
            },
            "6":{
              "isPhasable":true,
              "subtype":6,
              "mustHaveRecipient":false,
              "name":"AliasSell",
              "canHaveRecipient":true,
              "type":1,
              "isPhasingSafe":false
            },
            "7":{
              "isPhasable":true,
              "subtype":7,
              "mustHaveRecipient":true,
              "name":"AliasBuy",
              "canHaveRecipient":true,
              "type":1,
              "isPhasingSafe":false
            },
            "8":{
              "isPhasable":true,
              "subtype":8,
              "mustHaveRecipient":false,
              "name":"AliasDelete",
              "canHaveRecipient":false,
              "type":1,
              "isPhasingSafe":false
            },
            "9":{
              "isPhasable":true,
              "subtype":9,
              "mustHaveRecipient":false,
              "name":"PhasingVoteCasting",
              "canHaveRecipient":false,
              "type":1,
              "isPhasingSafe":true
            },
            "10":{
              "isPhasable":true,
              "subtype":10,
              "mustHaveRecipient":true,
              "name":"AccountProperty",
              "canHaveRecipient":true,
              "type":1,
              "isPhasingSafe":true
            },
            "11":{
              "isPhasable":true,
              "subtype":11,
              "mustHaveRecipient":true,
              "name":"AccountPropertyDelete",
              "canHaveRecipient":true,
              "type":1,
              "isPhasingSafe":true
            }
          }
        },
        "2":{
          "subtypes":{
            "0":{
              "isPhasable":true,
              "subtype":0,
              "mustHaveRecipient":false,
              "name":"AssetIssuance",
              "canHaveRecipient":false,
              "type":2,
              "isPhasingSafe":true
            },
            "1":{
              "isPhasable":true,
              "subtype":1,
              "mustHaveRecipient":true,
              "name":"AssetTransfer",
              "canHaveRecipient":true,
              "type":2,
              "isPhasingSafe":true
            },
            "2":{
              "isPhasable":true,
              "subtype":2,
              "mustHaveRecipient":false,
              "name":"AskOrderPlacement",
              "canHaveRecipient":false,
              "type":2,
              "isPhasingSafe":true
            },
            "3":{
              "isPhasable":true,
              "subtype":3,
              "mustHaveRecipient":false,
              "name":"BidOrderPlacement",
              "canHaveRecipient":false,
              "type":2,
              "isPhasingSafe":true
            },
            "4":{
              "isPhasable":true,
              "subtype":4,
              "mustHaveRecipient":false,
              "name":"AskOrderCancellation",
              "canHaveRecipient":false,
              "type":2,
              "isPhasingSafe":true
            },
            "5":{
              "isPhasable":true,
              "subtype":5,
              "mustHaveRecipient":false,
              "name":"BidOrderCancellation",
              "canHaveRecipient":false,
              "type":2,
              "isPhasingSafe":true
            },
            "6":{
              "isPhasable":true,
              "subtype":6,
              "mustHaveRecipient":false,
              "name":"DividendPayment",
              "canHaveRecipient":false,
              "type":2,
              "isPhasingSafe":true
            },
            "7":{
              "isPhasable":true,
              "subtype":7,
              "mustHaveRecipient":false,
              "name":"AssetDelete",
              "canHaveRecipient":false,
              "type":2,
              "isPhasingSafe":true
            }
          }
        },
        "3":{
          "subtypes":{
            "0":{
              "isPhasable":true,
              "subtype":0,
              "mustHaveRecipient":false,
              "name":"DigitalGoodsListing",
              "canHaveRecipient":false,
              "type":3,
              "isPhasingSafe":true
            },
            "1":{
              "isPhasable":true,
              "subtype":1,
              "mustHaveRecipient":false,
              "name":"DigitalGoodsDelisting",
              "canHaveRecipient":false,
              "type":3,
              "isPhasingSafe":true
            },
            "2":{
              "isPhasable":true,
              "subtype":2,
              "mustHaveRecipient":false,
              "name":"DigitalGoodsPriceChange",
              "canHaveRecipient":false,
              "type":3,
              "isPhasingSafe":false
            },
            "3":{
              "isPhasable":true,
              "subtype":3,
              "mustHaveRecipient":false,
              "name":"DigitalGoodsQuantityChange",
              "canHaveRecipient":false,
              "type":3,
              "isPhasingSafe":false
            },
            "4":{
              "isPhasable":true,
              "subtype":4,
              "mustHaveRecipient":true,
              "name":"DigitalGoodsPurchase",
              "canHaveRecipient":true,
              "type":3,
              "isPhasingSafe":false
            },
            "5":{
              "isPhasable":true,
              "subtype":5,
              "mustHaveRecipient":true,
              "name":"DigitalGoodsDelivery",
              "canHaveRecipient":true,
              "type":3,
              "isPhasingSafe":false
            },
            "6":{
              "isPhasable":true,
              "subtype":6,
              "mustHaveRecipient":true,
              "name":"DigitalGoodsFeedback",
              "canHaveRecipient":true,
              "type":3,
              "isPhasingSafe":false
            },
            "7":{
              "isPhasable":true,
              "subtype":7,
              "mustHaveRecipient":true,
              "name":"DigitalGoodsRefund",
              "canHaveRecipient":true,
              "type":3,
              "isPhasingSafe":false
            }
          }
        },
        "4":{
          "subtypes":{
            "0":{
              "isPhasable":true,
              "subtype":0,
              "mustHaveRecipient":true,
              "name":"EffectiveBalanceLeasing",
              "canHaveRecipient":true,
              "type":4,
              "isPhasingSafe":true
            },
            "1":{
              "isPhasable":true,
              "subtype":1,
              "mustHaveRecipient":false,
              "name":"SetPhasingOnly",
              "canHaveRecipient":false,
              "type":4,
              "isPhasingSafe":false
            }
          }
        },
        "5":{
          "subtypes":{
            "0":{
              "isPhasable":true,
              "subtype":0,
              "mustHaveRecipient":false,
              "name":"CurrencyIssuance",
              "canHaveRecipient":false,
              "type":5,
              "isPhasingSafe":false
            },
            "1":{
              "isPhasable":true,
              "subtype":1,
              "mustHaveRecipient":false,
              "name":"ReserveIncrease",
              "canHaveRecipient":false,
              "type":5,
              "isPhasingSafe":false
            },
            "2":{
              "isPhasable":true,
              "subtype":2,
              "mustHaveRecipient":false,
              "name":"ReserveClaim",
              "canHaveRecipient":false,
              "type":5,
              "isPhasingSafe":false
            },
            "3":{
              "isPhasable":true,
              "subtype":3,
              "mustHaveRecipient":true,
              "name":"CurrencyTransfer",
              "canHaveRecipient":true,
              "type":5,
              "isPhasingSafe":false
            },
            "4":{
              "isPhasable":true,
              "subtype":4,
              "mustHaveRecipient":false,
              "name":"PublishExchangeOffer",
              "canHaveRecipient":false,
              "type":5,
              "isPhasingSafe":false
            },
            "5":{
              "isPhasable":true,
              "subtype":5,
              "mustHaveRecipient":false,
              "name":"ExchangeBuy",
              "canHaveRecipient":false,
              "type":5,
              "isPhasingSafe":false
            },
            "6":{
              "isPhasable":true,
              "subtype":6,
              "mustHaveRecipient":false,
              "name":"ExchangeSell",
              "canHaveRecipient":false,
              "type":5,
              "isPhasingSafe":false
            },
            "7":{
              "isPhasable":true,
              "subtype":7,
              "mustHaveRecipient":false,
              "name":"CurrencyMinting",
              "canHaveRecipient":false,
              "type":5,
              "isPhasingSafe":false
            },
            "8":{
              "isPhasable":true,
              "subtype":8,
              "mustHaveRecipient":false,
              "name":"CurrencyDeletion",
              "canHaveRecipient":false,
              "type":5,
              "isPhasingSafe":false
            }
          }
        },
        "6":{
          "subtypes":{
            "0":{
              "isPhasable":false,
              "subtype":0,
              "mustHaveRecipient":false,
              "name":"TaggedDataUpload",
              "canHaveRecipient":false,
              "type":6,
              "isPhasingSafe":false
            },
            "1":{
              "isPhasable":false,
              "subtype":1,
              "mustHaveRecipient":false,
              "name":"TaggedDataExtend",
              "canHaveRecipient":false,
              "type":6,
              "isPhasingSafe":false
            }
          }
        },
        "7":{
          "subtypes":{
            "0":{
              "isPhasable":true,
              "subtype":0,
              "mustHaveRecipient":false,
              "name":"ShufflingCreation",
              "canHaveRecipient":false,
              "type":7,
              "isPhasingSafe":false
            },
            "1":{
              "isPhasable":true,
              "subtype":1,
              "mustHaveRecipient":false,
              "name":"ShufflingRegistration",
              "canHaveRecipient":false,
              "type":7,
              "isPhasingSafe":false
            },
            "2":{
              "isPhasable":false,
              "subtype":2,
              "mustHaveRecipient":false,
              "name":"ShufflingProcessing",
              "canHaveRecipient":false,
              "type":7,
              "isPhasingSafe":false
            },
            "3":{
              "isPhasable":false,
              "subtype":3,
              "mustHaveRecipient":false,
              "name":"ShufflingRecipients",
              "canHaveRecipient":false,
              "type":7,
              "isPhasingSafe":false
            },
            "4":{
              "isPhasable":false,
              "subtype":4,
              "mustHaveRecipient":false,
              "name":"ShufflingVerification",
              "canHaveRecipient":false,
              "type":7,
              "isPhasingSafe":false
            },
            "5":{
              "isPhasable":false,
              "subtype":5,
              "mustHaveRecipient":false,
              "name":"ShufflingCancellation",
              "canHaveRecipient":false,
              "type":7,
              "isPhasingSafe":false
            }
          }
        }
      },
      "votingModels":{
        "NQT":1,
        "CURRENCY":3,
        "ACCOUNT":0,
        "ASSET":2,
        "TRANSACTION":4,
        "NONE":-1,
        "HASH":5
      },
      "holdingTypes":{
        "CURRENCY":2,
        "ASSET":1,
        "NXT":0
      },
      "maxPrunableMessageLength":43008,
      "shufflingParticipantStates":{
        "CANCELLED":3,
        "REGISTERED":0,
        "PROCESSED":1,
        "VERIFIED":2
      },
      "currencyTypes":{
        "EXCHANGEABLE":1,
        "CLAIMABLE":8,
        "MINTABLE":16,
        "CONTROLLABLE":2,
        "RESERVABLE":4,
        "NON_SHUFFLEABLE":32
      },
      "disabledAPITags":[
        "Tagged Data"
      ],
      "maxBlockPayloadLength":44880,
      "maxPhasingDuration":20160,
      "mintingHashAlgorithms":{
        "SHA256":2,
        "SHA3":3,
        "SCRYPT":5,
        "Keccak25":25
      },
      "peerStates":{
        "DISCONNECTED":2,
        "NON_CONNECTED":0,
        "CONNECTED":1
      },
      "epochBeginning":1385294400000,
      "maxArbitraryMessageLength":160,
      "apiTags":{
        "NETWORK":{
          "name":"Networking",
          "enabled":true
        },
        "FORGING":{
          "name":"Forging",
          "enabled":true
        },
        "AE":{
          "name":"Asset Exchange",
          "enabled":true
        },
        "UTILS":{
          "name":"Utils",
          "enabled":true
        },
        "ACCOUNTS":{
          "name":"Accounts",
          "enabled":true
        },
        "SEARCH":{
          "name":"Search",
          "enabled":true
        },
        "MS":{
          "name":"Monetary System",
          "enabled":true
        },
        "TOKENS":{
          "name":"Tokens",
          "enabled":true
        },
        "INFO":{
          "name":"Server Info",
          "enabled":true
        },
        "TRANSACTIONS":{
          "name":"Transactions",
          "enabled":true
        },
        "DEBUG":{
          "name":"Debug",
          "enabled":true
        },
        "CREATE_TRANSACTION":{
          "name":"Create Transaction",
          "enabled":true
        },
        "DGS":{
          "name":"Digital Goods Store",
          "enabled":true
        },
        "SHUFFLING":{
          "name":"Shuffling",
          "enabled":true
        },
        "PHASING":{
          "name":"Phasing",
          "enabled":true
        },
        "BLOCKS":{
          "name":"Blocks",
          "enabled":true
        },
        "DATA":{
          "name":"Tagged Data",
          "enabled":false
        },
        "ALIASES":{
          "name":"Aliases",
          "enabled":true
        },
        "MESSAGES":{
          "name":"Messages",
          "enabled":true
        },
        "ADDONS":{
          "name":"Add-ons",
          "enabled":true
        },
        "VS":{
          "name":"Voting System",
          "enabled":true
        },
        "ACCOUNT_CONTROL":{
          "name":"Account Control",
          "enabled":true
        }
      },
      "maxTaggedDataDataLength":43008,
      "minBalanceModels":{
        "NQT":1,
        "CURRENCY":3,
        "ASSET":2,
        "NONE":0
      },
      "phasingHashAlgorithms":{
        "SHA256":2,
        "RIPEMD160":6,
        "RIPEMD160_SHA256":62
      },
      "shufflingStages":{
        "CANCELLED":4,
        "DONE":5,
        "PROCESSING":1,
        "BLAME":3,
        "REGISTRATION":0,
        "VERIFICATION":2
      },
      "hashAlgorithms":{
        "SHA256":2,
        "SHA3":3,
        "SCRYPT":5,
        "RIPEMD160":6,
        "Keccak25":25,
        "RIPEMD160_SHA256":62
      },
      "disabledAPIs":[
        [

        ]
      ],
      "requestTypes":{
        "getLastExchanges":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "startFundingMonitor":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getExpectedAskOrders":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAccountPublicKey":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "detectMimeType":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "fileParameter":"file",
          "requirePost":true,
          "enabled":false
        },
        "getBlocks":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAssetsByIssuer":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getExchangesByOffer":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAllOpenBidOrders":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "dgsPurchase":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getAccountBlockCount":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "deleteAlias":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "decodeFileToken":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "fileParameter":"file",
          "requirePost":true,
          "enabled":true
        },
        "getPlugins":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getFundingMonitor":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getDataTagsLike":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":false
        },
        "getPolls":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "downloadTaggedData":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":false
        },
        "getDataTags":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":false
        },
        "getPollVote":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAssetDeletes":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "addPeer":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "getSharedKey":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "decodeToken":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "popOff":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "getAccountPhasedTransactions":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAvailableToBuy":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getExpectedAssetDeletes":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "startForging":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getTaggedDataExtendTransactions":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":false
        },
        "getAssetAccounts":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getCurrencyFounders":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "currencyBuy":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "decodeQRCode":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "getAllExchanges":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getCurrencyTransfers":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getExpectedOrderCancellations":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "eventRegister":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "scan":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "hexConvert":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":false,
          "enabled":true
        },
        "getPhasingOnlyControl":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getDGSTagCount":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getOffer":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "encodeQRCode":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "getChannelTaggedData":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":false
        },
        "getAvailableToSell":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "cancelBidOrder":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "shufflingCancel":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getAccount":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getPeer":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAccountCurrentAskOrderIds":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getUnconfirmedTransactionIds":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAccountShufflings":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getExpectedSellOffers":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "dgsPriceChange":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getAliasesLike":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "dgsListing":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "fileParameter":"messageFile",
          "requirePost":true,
          "enabled":true
        },
        "getBidOrder":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "sendMessage":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getAllBroadcastedTransactions":{
          "allowRequiredBlockParameters":true,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":false,
          "enabled":true
        },
        "placeBidOrder":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getAccountBlocks":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getShuffling":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAccountCurrencies":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getExpectedTransactions":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAccountCurrentBidOrderIds":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAllPhasingOnlyControls":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "dgsRefund":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getAssetIds":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getTaggedData":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":false
        },
        "searchAccounts":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAccountLedger":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAccountAssets":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "deleteAccountProperty":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getBlockchainTransactions":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "sendMoney":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "extendTaggedData":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "fileParameter":"file",
          "requirePost":true,
          "enabled":false
        },
        "getMyInfo":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAccountTaggedData":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":false
        },
        "getAllTrades":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getStackTraces":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":false,
          "enabled":true
        },
        "rsConvert":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":false,
          "enabled":true
        },
        "searchTaggedData":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":false
        },
        "getAllTaggedData":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":false
        },
        "getDGSPendingPurchases":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getECBlock":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "generateFileToken":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "fileParameter":"file",
          "requirePost":true,
          "enabled":true
        },
        "searchDGSGoods":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAccountPhasedTransactionCount":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getCurrencyAccounts":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "shufflingCreate":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getAlias":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getPhasingPolls":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "markHost":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "canDeleteCurrency":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getPhasingPollVote":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "stopFundingMonitor":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getTime":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "buyAlias":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "searchPolls":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "eventWait":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "castVote":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getMintingTarget":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "generateToken":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "longConvert":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":false,
          "enabled":true
        },
        "getBlockId":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getLastTrades":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getExpectedBidOrders":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getBidOrderIds":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getBlockchainStatus":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getConstants":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":false,
          "enabled":true
        },
        "getTransaction":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getBlock":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "verifyTaggedData":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "fileParameter":"file",
          "requirePost":false,
          "enabled":false
        },
        "getExchangesByExchangeRequest":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getPrunableMessage":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "dividendPayment":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "broadcastTransaction":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "currencySell":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "blacklistPeer":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "dgsDelivery":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "setAccountProperty":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getShufflers":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getDGSGoodsPurchaseCount":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAssignedShufflings":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getGuaranteedBalance":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "fullHashToId":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":false,
          "enabled":true
        },
        "getExpectedBuyOffers":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAskOrders":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "stopForging":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getAccountExchangeRequests":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "downloadPrunableMessage":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAsset":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "clearUnconfirmedTransactions":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "getHoldingShufflings":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAssetPhasedTransactions":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAccountCurrentBidOrders":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "dgsQuantityChange":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getExpectedCurrencyTransfers":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "cancelAskOrder":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "searchAssets":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getDataTagCount":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":false
        },
        "dgsDelisting":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "deleteCurrency":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getAssetTransfers":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getBalance":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getCurrencyPhasedTransactions":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "setPhasingOnlyControl":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getCurrencies":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getDGSGoods":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "currencyReserveIncrease":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "deleteAssetShares":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "setLogging":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "getAliasCount":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getTransactionBytes":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "retrievePrunedTransaction":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getExpectedAssetTransfers":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAllAssets":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getInboundPeers":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "hash":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":false,
          "enabled":true
        },
        "createPoll":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "verifyPrunableMessage":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getDGSPurchase":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getReferencingTransactions":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getForging":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "readMessage":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "luceneReindex":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "fullReset":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "getAccountBlockIds":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getPollResult":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getDGSPurchaseCount":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAllWaitingTransactions":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":false,
          "enabled":true
        },
        "decryptFrom":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAccountAssetCount":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAssets":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getCurrenciesByIssuer":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getPeers":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAllShufflings":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "placeAskOrder":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "rebroadcastUnconfirmedTransactions":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "getAllCurrencies":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "setAccountInfo":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getDGSGood":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAskOrderIds":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAccountCurrencyCount":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "decodeHallmark":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAskOrder":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getExpectedExchangeRequests":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getCurrencyIds":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "shufflingProcess":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "requeueUnconfirmedTransactions":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "signTransaction":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":false,
          "enabled":true
        },
        "getAliases":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "trimDerivedTables":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "getSellOffers":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getLog":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":false,
          "enabled":true
        },
        "getAccountLedgerEntry":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "transferAsset":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "stopShuffler":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "publishExchangeOffer":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getLinkedPhasedTransactions":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "approveTransaction":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getDGSTagsLike":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "parseTransaction":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":false,
          "enabled":true
        },
        "getCurrency":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getBidOrders":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getDGSGoodsCount":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getCurrencyAccountCount":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getDGSPurchases":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getShufflingParticipants":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAccountLessors":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "startShuffler":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getPoll":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getVoterPhasedTransactions":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "transferCurrency":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "leaseBalance":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "setAlias":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "shutdown":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "getDGSExpiredPurchases":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "searchCurrencies":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "shufflingRegister":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "currencyReserveClaim":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getPollVotes":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAccountCurrentAskOrders":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getDGSTags":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getOrderTrades":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "sellAlias":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "dumpPeers":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":true,
          "enabled":true
        },
        "getAllOpenAskOrders":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAllPrunableMessages":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "dgsFeedback":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getPhasingPoll":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "shufflingVerify":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getDGSGoodsPurchases":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getAssetAccountCount":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getPhasingPollVotes":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "retrievePrunedData":{
          "allowRequiredBlockParameters":false,
          "requirePassword":true,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getUnconfirmedTransactions":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "encryptTo":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getBuyOffers":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getState":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "issueCurrency":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getAccountId":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":false,
          "enabled":true
        },
        "issueAsset":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "getTrades":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getPrunableMessages":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "calculateFullHash":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":false,
          "requirePost":false,
          "enabled":true
        },
        "currencyMint":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":true,
          "enabled":true
        },
        "uploadTaggedData":{
          "allowRequiredBlockParameters":false,
          "requirePassword":false,
          "requireBlockchain":true,
          "fileParameter":"file",
          "requirePost":true,
          "enabled":false
        },
        "getAccountProperties":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        },
        "getExchanges":{
          "allowRequiredBlockParameters":true,
          "requirePassword":false,
          "requireBlockchain":true,
          "requirePost":false,
          "enabled":true
        }
      }
    }

<small>*Verified 9-Jun-16*</small>

### Get State

**Request:**

    http://localhost:8125/burst?
      requestType=getState

**Response:**

    {
     "numberOfPolls": 19,
     "numberOfTransfers": 124216,
     "maxMemory": 900726784,
     "numberOfOffers": 90,
     "isDownloading": false,
     "cumulativeDifficulty": "17966254519242206",
     "numberOfCurrencies": 1833,
     "freeMemory": 36399744,
     "peerPort": 7874,
     "availableProcessors": 4,
     "numberOfTaggedData": 1,
     "numberOfActiveAccountLeases": 216,
     "currentMinRollbackHeight": 445745,
     "requestProcessingTime": 125533,
     "version": "1.5.11",
     "numberOfAliases": 142415,
     "numberOfActivePeers": 10,
     "maxPrunableLifetime": 1209600,
     "numberOfExchanges": 1750,
     "numberOfPurchases": 770,
     "numberOfAskOrders": 3536,
     "lastBlockchainFeeder": "84.253.125.186",
     "numberOfPeers": 289,
     "numberOfGoods": 990,
     "numberOfUnlockedAccounts": 0,
     "includeExpiredPrunable": false,
     "numberOfOrders": 4781,
     "numberOfTransactions": 1325666,
     "maxRollback": 800,
     "isScanning": false,
     "numberOfAssets": 539,
     "numberOfPrunableMessages": 45,
     "numberOfVotes": 106,
     "numberOfAccounts": 111301,
     "numberOfDataTags": 3,
     "needsAdminPassword": false,
     "numberOfBlocks": 446546,
     "isTestnet": false,
     "numberOfCurrencyTransfers": 719,
     "numberOfPhasedTransactions": 5,
     "numberOfAccountLeases": 231,
     "numberOfBidOrders": 1245,
     "lastBlock": "2164693711802180410",
     "totalMemory": 235929600,
     "application": "BRS",
     "lastBlockchainFeederHeight": 446597,
     "numberOfTrades": 102458,
     "numberOfTags": 723,
     "isOffline": false,
     "time": 48539770
    }

<small>*Verified 9-Jun-15*</small>

### Get Time

**Request:**

    http://localhost:8125/burst?
      requestType=getTime

**Response:**

    {
     "time": 31184078,
     "requestProcessingTime": 1
    }

<small>*Verified 20-Nov-14*</small>

Token Operations
----------------

### Decode Token

**Request:**

    http://localhost:8125/burst?
      requestType=decodeToken&
      website=test&
      token=u8q9ps0gdoo2bl158p4llpar583ld0cgejat9qnrgrgde4l5ut8bgn...

**Response:**

    {
     "valid": true,
     "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
     "requestProcessingTime": 2,
     "account": "15295723609781267838",
     "timestamp": 49762488
    }

<small>*Verified 22-Nov-14*</small>

### Generate Token

**Request:**

    http://localhost:8125/burst?
      requestType=generateToken&
      secretPhrase=secretPhrase&
      website=test

**Response:**

    {
     "valid": true,
     "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
     "requestProcessingTime": 4,
     "account": "15295723609781267838",
     "timestamp": 49762488,
     "token": "u8q9ps0gdoo2bl158p4llpar583ld0cgejat9qnrgrgde4l5ut8bgn..."
    }

<small>*Verified 23-Jun-15*</small>

Transaction Operations
----------------------

### Broadcast Transaction

**Request:**

    http://localhost:8125/burst?
      requestType=broadcastTransaction&
      transactionBytes=001046aac6013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143...

**Response:**

    {
     "requestProcessingTime": 4,
     "fullHash": "3a304584f20cf3d2cbbdd9698ff9a166427005ab98fbe9ca4ad6253651ee81f1",
     "transaction": "15200507403046301754"
    }

**Note**: If the transaction has already been broadcast, the following INFO notice appears in the console output and log file: *Transaction 15200507403046301754 already in blockchain (or unconfirmed pool), will not broadcast again*.

<small>*Verified 6-Nov-14*</small>

### Calculate Full Hash

**Request:**

    http://localhost:8125/burst?
      requestType=calculateFullHash&
      unsignedTransactionBytes=001046aac6013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f2...&
      signatureHash=b35eae7d2f01639810d37694138aa0a86fbbf8a9bf58c2be4f2a5b8f0f30b3f7

**Response:**

    {
     "requestProcessingTime": 1,
     "fullHash": "3a304584f20cf3d2cbbdd9698ff9a166427005ab98fbe9ca4ad6253651ee81f1"
    }

<small>*Verified 6-Nov-14*</small>

### Get Transaction

**Request:**

    http://localhost:8125/burst?
      requestType=getTransaction&
      transaction=15200507403046301754

**Response:**

    {
     "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
     "signature": "5f0378b7390ff5a815eadd1354de533eef682f139362b153576e2207320a6...",
     "feeNQT": "100000000",
     "transactionIndex": 2,
     "requestProcessingTime": 2842,
     "type": 0,
     "confirmations": 849,
     "fullHash": "3a304584f20cf3d2cbbdd9698ff9a166427005ab98fbe9ca4ad6253651ee81f1",
     "version": 1,
     "ecBlockId": "17321329645912574173",
     "signatureHash": "b35eae7d2f01639810d37694138aa0a86fbbf8a9bf58c2be4f2a5b8f0f30b3f7",
     "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
     "subtype": 0,
     "amountNQT": "100000000",
     "sender": "15323192282528158131",
     "recipientRS": "BURST-4VNQ-RWZC-4WWQ-GVM8S",
     "recipient": "17013046603665206934",
     "ecBlockHeight": 275727,
     "block": "8455642159445842600",
     "blockTimestamp": 29797208,
     "deadline": 60,
     "transaction": "15200507403046301754",
     "timestamp": 29796934,
     "height": 275730
    }

<small>*Verified 31-Dec-14*</small>

### Get Transaction Bytes

**Request:**

    http://localhost:8125/burst?
      requestType=getTransactionBytes&
      transaction=15200507403046301754

**Response:**

    {
     "unsignedTransactionBytes": "001046aac6013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473...",
     "requestProcessingTime": 66,
     "confirmations": 1019,
     "transactionBytes": "001046aac6013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473..."
    }

<small>*Verified 5-Nov-14*</small>

### Parse Transaction

**Request:**

    http://localhost:8125/burst?
      requestType=parseTransaction&
      transactionBytes=001046aac6013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143...

**Response:**

    {
     "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
     "signature": "5f0378b7390ff5a815eadd1354de533eef682f139362b153576e2207320a6...",
     "feeNQT": "100000000",
     "requestProcessingTime": 2,
     "type": 0,
     "fullHash": "3a304584f20cf3d2cbbdd9698ff9a166427005ab98fbe9ca4ad6253651ee81f1",
     "version": 1,
     "ecBlockId": "17321329645912574173",
     "signatureHash": "b35eae7d2f01639810d37694138aa0a86fbbf8a9bf58c2be4f2a5b8f0f30b3f7",
     "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
     "subtype": 0,
     "amountNQT": "100000000",
     "sender": "15323192282528158131",
     "recipientRS": "BURST-4VNQ-RWZC-4WWQ-GVM8S",
     "recipient": "17013046603665206934",
     "ecBlockHeight": 275727,
     "verify": true,
     "deadline": 60,
     "transaction": "15200507403046301754",
     "timestamp": 29796934,
     "height": 2147483647
    }

<small>*Verified 5-Nov-14*</small>

### Sign Transaction

**Request:**

    http://localhost:8125/burst?
      requestType=signTransaction&
      unsignedTransactionBytes=00100cfb3c03a00510f09c34f225d425306e5be55a494690...&
      secretPhrase=SecretPhrase

**Response:**

    {
     "signatureHash": "3f4830bf28f105d5075f5e084c36e4582a156e713abfe0bc6ee51cbf8b20f2d2",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "24513e93029688100c5a04183ffddc49812fd7a137a15fb3a2545aa9a2bcb5004a3...",
      "feeNQT": "100000000",
      "type": 0,
      "fullHash": "c34af8f1509e3be79c4562e24125ff2a8f026871fdd1a0366ad315bf8fab76b9",
      "version": 1,
      "phased": false,
      "ecBlockId": "15869644242181198665",
      "signatureHash": "3f4830bf28f105d5075f5e084c36e4582a156e713abfe0bc6ee51cbf8b20f2d2",
      "attachment": {
       "version.OrdinaryPayment": 0
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 0,
      "amountNQT": "1000000000",
      "sender": "15295723609781267838",
      "recipientRS": "BURST-BMUV-8QQR-47VK-CR7F3",
      "recipient": "11580081983047651163",
      "ecBlockHeight": 382777,
      "deadline": 1440,
      "transaction": "16662085316881435331",
      "timestamp": 54328076,
      "height": 2147483647
     },
     "verify": true,
     "requestProcessingTime": 5,
     "transactionBytes": "00100cfb3c03a00510f09c34f225d425306e5be55a4946908156072afbead4...",
     "fullHash": "c34af8f1509e3be79c4562e24125ff2a8f026871fdd1a0366ad315bf8fab76b9",
     "transaction": "16662085316881435331"
    }

<small>*Verified 15-Aug-15*</small>

Utilities
---------

### Long Convert

**Request:**

    http://localhost:8125/burst?
      requestType=longConvert&
      id=15323192282528158131

**Response:**

    {
     "stringId": "15323192282528158131",
     "requestProcessingTime": 0,
     "longId": "-3123551791181393485"
    }

<small>*Verified 12-Nov-14*</small>

### RS Convert

**Request:**

    http://localhost:8125/burst?
      requestType=rsConvert&
      account=BURST-L6FM-89WK-VK8P-FCRBB

**Response:**

    {
     "accountRS": "BURST-L6FM-89WK-VK8P-FCRBB",
     "requestProcessingTime": 1,
     "account": "15323192282528158131"
    }

<small>*Verified 7-Nov-14*</small>

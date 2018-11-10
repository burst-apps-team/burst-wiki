<languages/>

<translate>

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

### Get Subscription <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "senderRS": "BURST-2RN8-FSU8-P64Q-5AL9C",
    "sender": "3827576371473833606",
    "amountNQT": "500000000",
    "recipientRS": "BURST-JDMH-EZQC-UWY4-EXXFU",
    "recipient": "14787496155544039023",
    "id": "13721874590196751209",
    "timeNext": 128778762,
    "requestProcessingTime": 0,
    "frequency": 3600
}
```

### Get Subscriptions To Account <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "subscriptions": [
        {
            "senderRS": "BURST-2RN8-FSU8-P64Q-5AL9C",
            "sender": "3827576371473833606",
            "amountNQT": "500000000",
            "recipientRS": "BURST-JDMH-EZQC-UWY4-EXXFU",
            "recipient": "14787496155544039023",
            "id": "13721874590196751209",
            "timeNext": 128778762,
            "frequency": 3600
        }
    ],
    "requestProcessingTime": 2
}
```

### Get Account Subscriptions <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "subscriptions": [
        {
            "senderRS": "BURST-2RN8-FSU8-P64Q-5AL9C",
            "sender": "3827576371473833606",
            "amountNQT": "500000000",
            "recipientRS": "BURST-JDMH-EZQC-UWY4-EXXFU",
            "recipient": "14787496155544039023",
            "id": "13721874590196751209",
            "timeNext": 128778762,
            "frequency": 3600
        }
    ],
    "requestProcessingTime": 2
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

### Send Money Subscription <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "signatureHash": "c750d47f99a1dc5b01f3850d8c888b324fd49f1fe77c97fd8bccf339a82f3739",
    "unsignedTransactionBytes": "1513059bfe071800ad6cfaf61afba1b6542a5f3ef2a8f072a0e417b00d41cfcf86060cf780af1247d9f92fa97a86ce7200c2eb0b*",
    "transactionJSON": {
        "senderPublicKey": "ad6cfaf61afba1b6542a5f3ef2a8f072a0e417b00d41cfcf86060cf780af1247",
        "signature": "541e44ec8ea4642148dc6dd212076c16618580db47726f1522e2fda6805a680022cc3ef643f5cbefd0863796c9bcb8ee12cc611c38bc7a5bf01dbd2bc0777c29",
        "feeNQT": "10000000",
        "type": 21,
        "fullHash": "06491cddb3159f0ff9ff44e27b3ee2ed5a67756f9efe70dcbec83843c5bd2191",
        "version": 1,
        "ecBlockId": "3529501934685685596",
        "signatureHash": "c750d47f99a1dc5b01f3850d8c888b324fd49f1fe77c97fd8bccf339a82f3739",
        "attachment": {
            "version.SubscriptionSubscribe": 1,
            "frequency": 3600
        },
        "senderRS": "BURST-MNAM-VYK3-V3KC-DTACV",
        "subtype": 3,
        "amountNQT": "200000000",
        "sender": "13666482992520483091",
        "recipientRS": "BURST-ZYGT-HCNL-PU3A-98NM7",
        "recipient": "8272697426908805593",
        "ecBlockHeight": 554516,
        "deadline": 24,
        "transaction": "1125642294118861062",
        "timestamp": 134126341,
        "height": 2147483647
    },
    "broadcasted": true,
    "requestProcessingTime": 10,
    "transactionBytes": "1513059bfe071800ad6cfaf61afba1b6542a5f3ef2a8f072a0e417b00d41cfcf86060cf780af1247d9f92fa97a86ce7200c2eb0b*",
    "fullHash": "06491cddb3159f0ff9ff44e27b3ee2ed5a67756f9efe70dcbec83843c5bd2191",
    "transaction": "1125642294118861062"
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

Asset Exchange Operations <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />
-----------------------------------------------------------------------------------------------------------------------

### Cancel Order <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "signatureHash": "ea56c641c52e20cfeb2413dc49fc24a00eba23ac4aa444fdb059ad0aa66fdb0f",
    "unsignedTransactionBytes": "0214c0d347070100a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f120000000000000000000000000000000000e1f505*",
    "transactionJSON": {
        "senderPublicKey": "a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12",
        "signature": "67f67f75a4ed60109239719e29a11915e4e05addd7d851d75766ef4e64b7b7093ff55615144ffb26d2b0de4c8656330ecd1117b3d477219ed09a1098cb77cd81",
        "feeNQT": "100000000",
        "type": 2,
        "fullHash": "1454e33bfca5e67186d69ce027e408b3fc52dce6497d4f1d667598e52cae4a3e",
        "version": 1,
        "ecBlockId": "16853437834039716545",
        "signatureHash": "ea56c641c52e20cfeb2413dc49fc24a00eba23ac4aa444fdb059ad0aa66fdb0f",
        "attachment": {
            "version.AskOrderCancellation": 1,
            "order": "15076843998317839249"
        },
        "senderRS": "BURST-FRDJ-UPLH-MY9A-GUKQP",
        "subtype": 4,
        "amountNQT": "0",
        "sender": "16922903237994405232",
        "ecBlockHeight": 504969,
        "deadline": 1,
        "transaction": "8207429873684403220",
        "timestamp": 122147776,
        "height": 2147483647
    },
    "broadcasted": true,
    "requestProcessingTime": 7,
    "transactionBytes": "0214c0d347070100a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f120000000000000000000000000000000000e1f505*",
    "fullHash": "1454e33bfca5e67186d69ce027e408b3fc52dce6497d4f1d667598e52cae4a3e",
    "transaction": "8207429873684403220"
}
```

### Get Account Current Order Ids <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "requestProcessingTime": 2,
    "askOrderIds": [
        "4875748854378397140",
        "109196327778941945"
    ]
}
```

### Get Account Current Orders <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "askOrders": [
        {
            "quantityQNT": "100000",
            "priceNQT": "195000",
            "accountRS": "BURST-MBZS-2BCT-45QV-APCZB",
            "asset": "3702027329806229573",
            "type": "ask",
            "account": "9582909050628712440",
            "order": "4875748854378397140",
            "height": 495658
        },
        {
            "quantityQNT": "260100",
            "priceNQT": "200000",
            "accountRS": "BURST-MBZS-2BCT-45QV-APCZB",
            "asset": "3702027329806229573",
            "type": "ask",
            "account": "9582909050628712440",
            "order": "109196327778941945",
            "height": 494427
        }
    ],
    "requestProcessingTime": 1
}
```

### Get All Assets <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "assets": [
        {
            "quantityQNT": "10000000000",
            "numberOfAccounts": 77,
            "accountRS": "BURST-8E8K-WQ2F-ZDZ5-FQWHX",
            "decimals": 8,
            "numberOfTransfers": 26,
            "name": "BCPT",
            "description": "Burst Cryptoport Pool Token, is an asset that backed by profit from MiningPool at http://burst-pool.cryptoport.io, i will pay dividend from 50% fund i got from pool fee (monthly). dividend will be proportional to BCPT owned",
            "numberOfTrades": 510,
            "asset": "12791182347560578640",
            "account": "16050713509424738513"
        },
        {
            "quantityQNT": "100000000",
            "numberOfAccounts": 26,
            "accountRS": "BURST-H2ZW-3H4D-RJBS-FCVGV",
            "decimals": 2,
            "numberOfTransfers": 36,
            "name": "BurstFund",
            "description": "This asset will act as a stabilizer fund for Burst. This asset aims to assist in maintaining the value of Burst, providing market correction in the event of market manipulation. Each Quantity will be representative of 10 Burst at an initial market price of 0.00001000 each (or determined by asset market). This assets main focus will be ensuring Burst does not become another outdated and dead NXT clone...",
            "numberOfTrades": 2,
            "asset": "673530795527425458",
            "account": "15977480701804512252"
        }
    ],
    "requestProcessingTime": 3
}
```

### Get All Open Orders <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "requestProcessingTime": 0,
    "openOrders": [
        {
            "quantityQNT": "10",
            "priceNQT": "10000000000000",
            "accountRS": "BURST-359Q-QH73-4N5P-FP54C",
            "asset": "15295227971848272658",
            "type": "ask",
            "account": "15350648744942013686",
            "order": "14034527401109329159",
            "height": 17901
        },
        {
            "quantityQNT": "1",
            "priceNQT": "10000000000000",
            "accountRS": "BURST-Q4TR-YKRW-6RAN-EDBC8",
            "asset": "11375670541237055652",
            "type": "ask",
            "account": "14062819640288676663",
            "order": "18379469307992717843",
            "height": 11997
        },
        {
            "quantityQNT": "4392",
            "priceNQT": "100000000000",
            "accountRS": "BURST-SKL3-ACW8-DBN6-5M8VM",
            "asset": "11375670541237055652",
            "type": "ask",
            "account": "4173943238181013057",
            "order": "756014781951608408",
            "height": 11257
        }
    ]
}
```

### Get All Trades <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "trades": [
        {
            "seller": "2695253629713716937",
            "quantityQNT": "7700",
            "bidOrder": "17489219850997945774",
            "sellerRS": "BURST-Q7QB-WPPC-6VE2-4QVTC",
            "buyer": "14676337193484961173",
            "priceNQT": "100000",
            "askOrder": "7125231493760146086",
            "buyerRS": "BURST-EKEP-XEYS-6YPW-EL8DR",
            "decimals": 2,
            "name": "CryptoMaps",
            "block": "10810208340839229954",
            "asset": "3702027329806229573",
            "askOrderHeight": 504189,
            "tradeType": "sell",
            "timestamp": 122140342,
            "height": 504949
        },
        {
            "seller": "6905334832585552856",
            "quantityQNT": "1000",
            "bidOrder": "14942960390822608034",
            "sellerRS": "BURST-BEYS-G8VQ-6SEK-7F7PZ",
            "buyer": "17808054289263456125",
            "priceNQT": "66500000",
            "askOrder": "3461704332885328085",
            "buyerRS": "BURST-QVVX-2TL8-WZ2C-HDRAG",
            "decimals": 0,
            "name": "BTFGPool",
            "block": "9914674296156980645",
            "asset": "9036920395530551012",
            "askOrderHeight": 504920,
            "tradeType": "buy",
            "timestamp": 122132789,
            "height": 504920
        }
    ],
    "requestProcessingTime": 447
}
```

### Get Asset <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "quantityQNT": "20000000000",
    "numberOfAccounts": 9210,
    "accountRS": "BURST-CMAP-ME5N-TFKP-6BCER",
    "decimals": 2,
    "numberOfTransfers": 11869,
    "name": "CryptoMaps",
    "description": "The official CryptoMaps Token, see more details at https://token.cryptomaps.me",
    "numberOfTrades": 523,
    "requestProcessingTime": 0,
    "asset": "3702027329806229573",
    "account": "5454221553913122069"
}
```

### Get Asset Accounts <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "accountAssets": [
        {
            "quantityQNT": "5000000000",
            "accountRS": "BURST-CMAP-AY84-2TV8-6F2QK",
            "unconfirmedQuantityQNT": "5000000000",
            "asset": "3702027329806229573",
            "account": "5248959966645406997"
        },
        {
            "quantityQNT": "5000000000",
            "accountRS": "BURST-CMSP-HRMP-ZBTM-9P4J7",
            "unconfirmedQuantityQNT": "5000000000",
            "asset": "3702027329806229573",
            "account": "8268702915092631317"
        }
    ],
    "requestProcessingTime": 22
}
```

### Get Asset Ids <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "assetIds": [
        "2644409077762286513",
        "5533434524898779728",
        "7756017130240677072",
        "2663432644302202784",
        "11700625361170592721",
        "9789600218215198873"
    ],
    "requestProcessingTime": 1
}
```

### Get Asset Transfers <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "transfers": [
        {
            "quantityQNT": "10000",
            "senderRS": "BURST-RBHF-BY3C-V949-37G3N",
            "assetTransfer": "16436527819322024045",
            "sender": "1875121765155055085",
            "recipientRS": "BURST-59HP-HQ5R-7MGU-AG4LC",
            "decimals": 2,
            "recipient": "9604012506417831413",
            "name": "CryptoMaps",
            "asset": "3702027329806229573",
            "height": 504029,
            "timestamp": 121917645
        },
        {
            "quantityQNT": "5500",
            "senderRS": "BURST-PVZT-KWEH-ZTJE-7Z7C9",
            "assetTransfer": "9907383924348587793",
            "sender": "6028278722555736057",
            "recipientRS": "BURST-RKW5-ZVFT-JU69-6X4JC",
            "decimals": 2,
            "recipient": "4990091736314005379",
            "name": "CryptoMaps",
            "asset": "3702027329806229573",
            "height": 500631,
            "timestamp": 121093713
        }
    ],
    "requestProcessingTime": 2
}
```

### Get Assets <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "assets": [
        {
            "quantityQNT": "20000000000",
            "numberOfAccounts": 9210,
            "accountRS": "BURST-CMAP-ME5N-TFKP-6BCER",
            "decimals": 2,
            "numberOfTransfers": 11869,
            "name": "CryptoMaps",
            "description": "The official CryptoMaps Token, see more details at https://token.cryptomaps.me",
            "numberOfTrades": 523,
            "asset": "3702027329806229573",
            "account": "5454221553913122069"
        },
        {
            "quantityQNT": "200000000000",
            "numberOfAccounts": 1,
            "accountRS": "BURST-SPPT-TD44-HUVK-FSZS4",
            "decimals": 2,
            "numberOfTransfers": 0,
            "name": "SpotPoint",
            "description": "The official SpotPoint Token, see more details at https://spotpoint.me",
            "numberOfTrades": 0,
            "asset": "9306257191450064346",
            "account": "15088176726379615929"
        }
    ],
    "requestProcessingTime": 0
}
```

### Get Assets By Issuer <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "assets": [
        [
            {
                "quantityQNT": "100000000",
                "numberOfAccounts": 19,
                "accountRS": "BURST-KQL8-645F-3WU8-3GMHB",
                "decimals": 2,
                "numberOfTransfers": 0,
                "name": "BoB1Pos",
                "description": "Brotherhood of Blockchain www.bob-invest.com proudly present the first investment opportunity on Burst asset. This asset will invest in Prove of Stake coins and pay out dividend in Burst. The payout ratio is 70% to asset holder 20% equipment, internet, electricity, maintenance and admin. 10% reinvest in coins.This asset will make a monthly payment every 15th of the month, first payment is 15/6/2018.",
                "numberOfTrades": 35,
                "asset": "15511449532155436028",
                "account": "1494753212313950790"
            }
        ],
        [
            {
                "quantityQNT": "80000000000000",
                "numberOfAccounts": 6,
                "accountRS": "BURST-8CPS-Q628-PU9W-7H9R7",
                "decimals": 6,
                "numberOfTransfers": 3,
                "name": "GMK",
                "description": "GMK Coin",
                "numberOfTrades": 2,
                "asset": "8141917335139373659",
                "account": "5970910749481183928"
            }
        ]
    ],
    "requestProcessingTime": 3
}
```

### Get Order <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "quantityQNT": "10",
    "priceNQT": "10000000000000",
    "accountRS": "BURST-359Q-QH73-4N5P-FP54C",
    "requestProcessingTime": 1,
    "asset": "15295227971848272658",
    "type": "ask",
    "account": "15350648744942013686",
    "order": "14034527401109329159",
    "height": 17901
}
```

### Get Order Ids <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "requestProcessingTime": 3,
    "askOrderIds": [
        "9230225288402446558",
        "9243318318548315528",
        "9399454941412352140",
        "9436352279829246389",
        "9443733332658954573",
        "9545675672764030312"
    ]
}
```

### Get Orders <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "askOrders": [
        {
            "quantityQNT": "250",
            "priceNQT": "100000000",
            "accountRS": "BURST-CDQV-WN87-YB2L-FTEPP",
            "asset": "11700625361170592721",
            "type": "ask",
            "account": "15768678879341129435",
            "order": "9230225288402446558",
            "height": 501552
        },
        {
            "quantityQNT": "50000",
            "priceNQT": "100000000",
            "accountRS": "BURST-CDQV-WN87-YB2L-FTEPP",
            "asset": "11700625361170592721",
            "type": "ask",
            "account": "15768678879341129435",
            "order": "9243318318548315528",
            "height": 501552
        },
        {
            "quantityQNT": "50",
            "priceNQT": "100000000",
            "accountRS": "BURST-CDQV-WN87-YB2L-FTEPP",
            "asset": "11700625361170592721",
            "type": "ask",
            "account": "15768678879341129435",
            "order": "9399454941412352140",
            "height": 501552
        }
    ],
    "requestProcessingTime": 1
}
```

### Get Trades <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "trades": [
        {
            "seller": "2695253629713716937",
            "quantityQNT": "7700",
            "bidOrder": "17489219850997945774",
            "sellerRS": "BURST-Q7QB-WPPC-6VE2-4QVTC",
            "buyer": "14676337193484961173",
            "priceNQT": "100000",
            "askOrder": "7125231493760146086",
            "buyerRS": "BURST-EKEP-XEYS-6YPW-EL8DR",
            "block": "10810208340839229954",
            "asset": "3702027329806229573",
            "askOrderHeight": 504189,
            "tradeType": "sell",
            "timestamp": 122140342,
            "height": 504949
        },
        {
            "seller": "13248005565246971693",
            "quantityQNT": "100000",
            "bidOrder": "6312065606850043247",
            "sellerRS": "BURST-59TF-VVKH-X6SQ-DYLQH",
            "buyer": "9582909050628712440",
            "priceNQT": "120000",
            "askOrder": "17472059594274895773",
            "buyerRS": "BURST-MBZS-2BCT-45QV-APCZB",
            "block": "12097007668561592790",
            "asset": "3702027329806229573",
            "askOrderHeight": 504899,
            "tradeType": "buy",
            "timestamp": 122129038,
            "height": 504899
        },
        {
            "seller": "12065311156306389227",
            "quantityQNT": "7700",
            "bidOrder": "17489219850997945774",
            "sellerRS": "BURST-8C9D-LBN5-X9FG-CS6WG",
            "buyer": "14676337193484961173",
            "priceNQT": "100000",
            "askOrder": "2519849914836375271",
            "buyerRS": "BURST-EKEP-XEYS-6YPW-EL8DR",
            "block": "16551915903483185103",
            "asset": "3702027329806229573",
            "askOrderHeight": 504189,
            "tradeType": "sell",
            "timestamp": 122075678,
            "height": 504681
        }
    ],
    "requestProcessingTime": 1
}
```

### Issue Asset <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Note:** This is a fake example.

**Response:**

``` json
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
```

### Place Order <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "signatureHash": "3358f546ce7136ca478692102d0b526f80b4ca83eb78e90d15813b9636b654e5",
    "unsignedTransactionBytes": "0213c6cc47070100a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f120000000000000000000000000000000000e1f505*",
    "transactionJSON": {
        "senderPublicKey": "a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12",
        "signature": "48d14f308d26380e4dfa26b296c8728eef04af8704acff60f4d8bef28a126701ea29878753d2106a16228cfc3829e8f3d6a2e5baf739dbd0460fd45cff1241e0",
        "feeNQT": "100000000",
        "type": 2,
        "fullHash": "229f52765195110e0e096b81b91f04186be14ada7e29949841ee30e37b655775",
        "version": 1,
        "ecBlockId": "6656096466494440810",
        "signatureHash": "3358f546ce7136ca478692102d0b526f80b4ca83eb78e90d15813b9636b654e5",
        "attachment": {
            "quantityQNT": "1",
            "priceNQT": "100000000",
            "asset": "3702027329806229573",
            "version.BidOrderPlacement": 1
        },
        "senderRS": "BURST-FRDJ-UPLH-MY9A-GUKQP",
        "subtype": 3,
        "amountNQT": "0",
        "sender": "16922903237994405232",
        "ecBlockHeight": 504963,
        "deadline": 1,
        "transaction": "1013755568245088034",
        "timestamp": 122145990,
        "height": 2147483647
    },
    "broadcasted": false,
    "requestProcessingTime": 3,
    "transactionBytes": "0213c6cc47070100a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f120000000000000000000000000000000000e1f505*",
    "fullHash": "229f52765195110e0e096b81b91f04186be14ada7e29949841ee30e37b655775",
    "transaction": "1013755568245088034"
}
```

### Transfer Asset <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "signatureHash": "59456e26081926b27c38ca46d574e4a6373953b932e8368f131f7bc86ffc3aa9",
    "unsignedTransactionBytes": "021143d147070100a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12ae2527ec7e55f1eb000000000000000000e1f505*",
    "transactionJSON": {
        "senderPublicKey": "a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12",
        "signature": "b5f33599a52e5e4e56ad5c316cd6f6ced8ebe9e9931222aaee30407570ede60031a068b6a89714fca851b486e8d8c2509ef46b3e0e1ef44c819cbff5b0bb78c4",
        "feeNQT": "100000000",
        "type": 2,
        "fullHash": "3b7d9f5b998ce795a3300ed1e6df1ba8fe467b137e2f2c662c9149a5b0f878d9",
        "version": 1,
        "ecBlockId": "6974740076508916530",
        "signatureHash": "59456e26081926b27c38ca46d574e4a6373953b932e8368f131f7bc86ffc3aa9",
        "attachment": {
            "version.AssetTransfer": 1,
            "quantityQNT": "1",
            "asset": "3702027329806229573"
        },
        "senderRS": "BURST-FRDJ-UPLH-MY9A-GUKQP",
        "subtype": 1,
        "amountNQT": "0",
        "sender": "16922903237994405232",
        "recipientRS": "BURST-GBFG-HVQ4-8AMM-GPCWR",
        "recipient": "17001464071916561838",
        "ecBlockHeight": 504968,
        "deadline": 1,
        "transaction": "10801756821566487867",
        "timestamp": 122147139,
        "height": 2147483647
    },
    "broadcasted": true,
    "requestProcessingTime": 5,
    "transactionBytes": "021143d147070100a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12ae2527ec7e55f1eb000000000000000000e1f505*",
    "fullHash": "3b7d9f5b998ce795a3300ed1e6df1ba8fe467b137e2f2c662c9149a5b0f878d9",
    "transaction": "10801756821566487867"
}
```

Block Operations <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />
--------------------------------------------------------------------------------------------------------------

### Get Block <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "previousBlockHash": "98338ec9e496040f505dace85c3ab331e6ad257f24d6e9417877a4a22b449ab3",
    "payloadLength": 214,
    "totalAmountNQT": "0",
    "generationSignature": "be12abb71bfe2dcde77cbd5872e64855929d934d04a491cb08950b9e509f37eb",
    "generator": "16286104011619463406",
    "generatorPublicKey": "ba5464139a83fcdb600ba141b567430ec5ebd0e409d246cdd33c3597855e0a4e",
    "baseTarget": "116169",
    "payloadHash": "20585e4ab9b6ac07f8903a509b1494dab99b931bedf3e3ad11c5a556a82b7c3e",
    "generatorRS": "BURST-AW9G-BYKU-AZWL-GXY26",
    "blockReward": "944",
    "nextBlock": "10431007917907627236",
    "requestProcessingTime": 0,
    "scoopNum": 3035,
    "numberOfTransactions": 1,
    "blockSignature": "5c40fb39d21ec55a77718e9de25b40f008b5a940434208b191bf010322c453047a3fe6418b0ffa62c131a31d59142ec9b336815210c96a8a37709768a6d8e6c3",
    "transactions": [
        "3120851314369640207"
    ],
    "nonce": "77469254",
    "version": 3,
    "totalFeeNQT": "735000",
    "previousBlock": "1082155719854011288",
    "block": "2298247278137763160",
    "height": 502797,
    "timestamp": 121620850
}
```

### Get Block Id <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "block": "9782158559918006093",
    "requestProcessingTime": 1
}
```

### Get Blocks <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "blocks": [
        {
            "previousBlockHash": "e4cc2b7c6f62c290d1405ef236b846d316df4cc187c9554b18dcb54861f4ade2",
            "payloadLength": 1234,
            "totalAmountNQT": "132933935798",
            "generationSignature": "53f43986acef5899f08e388bde1bc6b935499c91badace2247076aa984fea577",
            "generator": "12019899115821911181",
            "generatorPublicKey": "f6b3ee829223351281340617eae3bdc3082ae9db6899395befddabc4aae37161",
            "baseTarget": "115429",
            "payloadHash": "2e71185d5f9c8cce3f2cc87dec5df9643dc0b6562f78f49b553610bc4d9af5a3",
            "generatorRS": "BURST-KM6F-KZ7B-46SF-C4UMF",
            "blockReward": "944",
            "scoopNum": 1265,
            "numberOfTransactions": 7,
            "blockSignature": "c5f9495278e685e90d4d42108c0a25b0a62f0006bdbb70f8df90d17d49a4670c09d13d63947fe69904340748348001277308de67891ee3bec177524629bc7383",
            "transactions": [
                "14343547955476128075",
                "14751316321929102651",
                "18403710335612237831",
                "2492947289038528624",
                "3664660987329932102",
                "5094936331677798562",
                "5165719427461695861"
            ],
            "nonce": "518857973",
            "version": 3,
            "totalFeeNQT": "510735000",
            "previousBlock": "10431007917907627236",
            "block": "3526547904120473437",
            "height": 502799,
            "timestamp": 121621514
        },
        {
            "previousBlockHash": "58e5680bb903e51f3dfe2221d5f291e5e7750f360fe78f7f5fee3c568867c7f2",
            "payloadLength": 938,
            "totalAmountNQT": "139674768871",
            "generationSignature": "142e073b4a84e95284b8eae679f1cc87d82f418897ffe0858d9bd1adda33d220",
            "generator": "12414125682074447698",
            "generatorPublicKey": "4f04a259ddc9c60e49bf3ef0ca5cb53c00ae97dedc1fcb59af2e48a9894ea745",
            "baseTarget": "116986",
            "payloadHash": "3d5153f350a07fb6d0b9a4cd07f2d4ef38b0bf312df2091eda05edf2ca4f1283",
            "generatorRS": "BURST-MVUL-K5UK-RRTY-CPYKS",
            "blockReward": "944",
            "nextBlock": "3526547904120473437",
            "scoopNum": 3885,
            "numberOfTransactions": 5,
            "blockSignature": "300fd5eace546b846213894c0afa9e68755fa300472a1101ca36d524d3433108445f7154d6b9151c3202272e516b7b9c180751ee65de949cdc7a459409a8b613",
            "transactions": [
                "13447447233023586787",
                "14140393734437081046",
                "3451268376629033899",
                "4144139998017169860",
                "5800257696655274819"
            ],
            "nonce": "77540123",
            "version": 3,
            "totalFeeNQT": "400735000",
            "previousBlock": "2298247278137763160",
            "block": "10431007917907627236",
            "height": 502798,
            "timestamp": 121620994
        }
    ],
    "requestProcessingTime": 1
}
```

### Get EC Block <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "ecBlockHeight": 502788,
    "requestProcessingTime": 1,
    "ecBlockId": "15769175919943831738",
    "timestamp": 121621703
}
```

Digital Goods Store Operations
------------------------------

### DGS Delisting <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
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
```

### DGS Delivery <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
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
```

### Feedback <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
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
```

### DGS Listing <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
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
```

### DGS Price Change <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
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
```

### DGS Purchase <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
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
```

### DGS Quantity Change <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json

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
```

### DGS Refund <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json

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
```

### Get DGS Good <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "seller": "16432157717192673510",
    "priceNQT": "750000000000",
    "quantity": 1000,
    "name": "100 TB Plot File",
    "goods": "490136863013671329",
    "description": "48-100h delivery Time\r\nSpecify how big you want your file\r\nMin size 1 TB\r\nMax size 10 TB\r\nDelivery is a Torrent Download",
    "sellerRS": "BURST-F498-E9P9-ACDE-G2Q4A",
    "requestProcessingTime": 1,
    "delisted": false,
    "tags": "plotter",
    "timestamp": 121088459
}
```

### Get DGS Goods <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "seller": "16432157717192673510",
    "priceNQT": "750000000000",
    "quantity": 1000,
    "name": "100 TB Plot File",
    "goods": "490136863013671329",
    "description": "48-100h delivery Time\r\nSpecify how big you want your file\r\nMin size 1 TB\r\nMax size 10 TB\r\nDelivery is a Torrent Download",
    "sellerRS": "BURST-F498-E9P9-ACDE-G2Q4A",
    "requestProcessingTime": 1,
    "delisted": false,
    "tags": "plotter",
    "timestamp": 121088459
}
```

### Get DGS Goods Purchases <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "goods": [
        {
            "seller": "869755421284789028",
            "priceNQT": "100000000",
            "quantity": 9996,
            "name": "Donate - Help CompSci Students",
            "goods": "8877051912510361584",
            "description": "My claass of compsci students is working on burst related projects as well as a mining rig.  Help me help them learn more, do more.  Please donate to their cause.",
            "sellerRS": "BURST-EUT6-YC9L-N63Y-2XH6S",
            "delisted": false,
            "tags": "school",
            "timestamp": 117044802
        },
        {
            "seller": "869755421284789028",
            "priceNQT": "1000000000",
            "quantity": 999999,
            "name": "Hard drives for high school students",
            "goods": "4170413457250059275",
            "description": "The high school computer science class is woefully funded.  Please help and donate..",
            "sellerRS": "BURST-EUT6-YC9L-N63Y-2XH6S",
            "delisted": false,
            "tags": "",
            "timestamp": 122140901
        }
    ],
    "requestProcessingTime": 1
}
```

### Get DGS Pending Purchases <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "purchases": [
        {
            "seller": "869755421284789028",
            "priceNQT": "100000000",
            "quantity": 1,
            "deliveryDeadlineTimestamp": 142820935,
            "buyerRS": "BURST-8J3W-BBNN-BJMG-HCPG3",
            "pending": true,
            "purchase": "18028864805999806823",
            "name": "Donate - Help CompSci Students",
            "goods": "8877051912510361584",
            "sellerRS": "BURST-EUT6-YC9L-N63Y-2XH6S",
            "buyer": "17346364151341203516",
            "timestamp": 120220136
        },
        {
            "seller": "869755421284789028",
            "note": {
                "data": "530c56998edacb24b664d9be3d9db3be52cb89bc10efdfbb1c97a5bb415bff2a1737ecf758202b57171145296bb3db2e",
                "nonce": "6b59813a09abe6809d87a34a0d7eb4ad7b97f43a28ae8d26f4811f8fff864712"
            },
            "quantity": 1,
            "pending": true,
            "purchase": "9754082484071034298",
            "goods": "8877051912510361584",
            "sellerRS": "BURST-EUT6-YC9L-N63Y-2XH6S",
            "buyer": "11786437085044315193",
            "priceNQT": "100000000",
            "deliveryDeadlineTimestamp": 139660265,
            "buyerRS": "BURST-X73T-4MLS-LV98-CQG69",
            "name": "Donate - Help CompSci Students",
            "timestamp": 117059465
        }
    ],
    "requestProcessingTime": 2
}
```

### Get DGS Purchase <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "seller": "869755421284789028",
    "quantity": 1,
    "pending": true,
    "purchase": "18028864805999806823",
    "goods": "8877051912510361584",
    "sellerRS": "BURST-EUT6-YC9L-N63Y-2XH6S",
    "requestProcessingTime": 1,
    "buyer": "17346364151341203516",
    "priceNQT": "100000000",
    "deliveryDeadlineTimestamp": 142820935,
    "buyerRS": "BURST-8J3W-BBNN-BJMG-HCPG3",
    "name": "Donate - Help CompSci Students",
    "timestamp": 120220136
}
```

### Get DGS Purchases <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "purchases": [
        {
            "seller": "869755421284789028",
            "priceNQT": "100000000",
            "quantity": 1,
            "deliveryDeadlineTimestamp": 142820935,
            "buyerRS": "BURST-8J3W-BBNN-BJMG-HCPG3",
            "pending": true,
            "purchase": "18028864805999806823",
            "name": "Donate - Help CompSci Students",
            "goods": "8877051912510361584",
            "sellerRS": "BURST-EUT6-YC9L-N63Y-2XH6S",
            "buyer": "17346364151341203516",
            "timestamp": 120220136
        },
        {
            "seller": "869755421284789028",
            "note": {
                "data": "530c56998edacb24b664d9be3d9db3be52cb89bc10efdfbb1c97a5bb415bff2a1737ecf758202b57171145296bb3db2e",
                "nonce": "6b59813a09abe6809d87a34a0d7eb4ad7b97f43a28ae8d26f4811f8fff864712"
            },
            "quantity": 1,
            "pending": true,
            "purchase": "9754082484071034298",
            "goods": "8877051912510361584",
            "sellerRS": "BURST-EUT6-YC9L-N63Y-2XH6S",
            "buyer": "11786437085044315193",
            "priceNQT": "100000000",
            "deliveryDeadlineTimestamp": 139660265,
            "buyerRS": "BURST-X73T-4MLS-LV98-CQG69",
            "name": "Donate - Help CompSci Students",
            "timestamp": 117059465
        }
    ],
    "requestProcessingTime": 2
}
```

Networking Operations <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />
-------------------------------------------------------------------------------------------------------------------

### Get My Info <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "address": "127.0.0.1",
    "host": "127.0.0.1",
    "requestProcessingTime": 0
}
```

### Get Peer <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "lastUpdated": 121622908,
    "downloadedVolume": 5557831,
    "blacklisted": false,
    "announcedAddress": "107.150.6.121",
    "application": "BRS",
    "uploadedVolume": 11482803,
    "state": 1,
    "requestProcessingTime": 0,
    "version": "2.2.1",
    "platform": "PC",
    "shareAddress": true
}
```

### Get Peers <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "peers": [
        "107.150.6.121",
        "37.205.11.73",
        "79.113.222.145",
        "74.141.142.82",
        "71.172.254.124",
        "89.133.236.99",
        "97.84.139.244"
    ],
    "requestProcessingTime": 0
}
```

Server Information Operations <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />
---------------------------------------------------------------------------------------------------------------------------

### Get Blockchain Status <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "lastBlock": "10431007917907627236",
    "application": "BRS",
    "isScanning": false,
    "cumulativeDifficulty": "29365339954194811259",
    "lastBlockchainFeederHeight": 501733,
    "numberOfBlocks": 502799,
    "time": 121621131,
    "requestProcessingTime": 0,
    "version": "2.2.1",
    "lastBlockchainFeeder": "86.26.25.217"
}
```

### Get Constants <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "maxBlockPayloadLength": 44880,
    "genesisAccountId": "0",
    "genesisBlockId": "3444294670862540038",
    "transactionTypes": [
        {
            "description": "Payment",
            "value": 0,
            "subtypes": [
                {
                    "description": "Ordinary payment",
                    "value": 0
                }
            ]
        },
        {
            "description": "Messaging",
            "value": 1,
            "subtypes": [
                {
                    "description": "Arbitrary message",
                    "value": 0
                },
                {
                    "description": "Alias assignment",
                    "value": 1
                },
                {
                    "description": "Alias sell",
                    "value": 6
                },
                {
                    "description": "Alias buy",
                    "value": 7
                },
                {
                    "description": "Account info",
                    "value": 5
                }
            ]
        },
        {
            "description": "Colored coins",
            "value": 2,
            "subtypes": [
                {
                    "description": "Asset issuance",
                    "value": 0
                },
                {
                    "description": "Asset transfer",
                    "value": 1
                },
                {
                    "description": "Ask order placement",
                    "value": 2
                },
                {
                    "description": "Bid order placement",
                    "value": 3
                },
                {
                    "description": "Ask order cancellation",
                    "value": 4
                },
                {
                    "description": "Bid order cancellation",
                    "value": 5
                }
            ]
        },
        {
            "description": "Digital goods",
            "value": 3,
            "subtypes": [
                {
                    "description": "Listing",
                    "value": 0
                },
                {
                    "description": "Delisting",
                    "value": 1
                },
                {
                    "description": "Price change",
                    "value": 2
                },
                {
                    "description": "Quantity change",
                    "value": 3
                },
                {
                    "description": "Purchase",
                    "value": 4
                },
                {
                    "description": "Delivery",
                    "value": 5
                },
                {
                    "description": "Feedback",
                    "value": 6
                },
                {
                    "description": "Refund",
                    "value": 7
                }
            ]
        },
        {
            "description": "Account Control",
            "value": 4,
            "subtypes": [
                {
                    "description": "Effective balance leasing",
                    "value": 0
                }
            ]
        }
    ],
    "peerStates": [
        {
            "description": "Non-connected",
            "value": 0
        },
        {
            "description": "Connected",
            "value": 1
        },
        {
            "description": "Disconnected",
            "value": 2
        }
    ],
    "maxArbitraryMessageLength": 1000,
    "requestTypes": {}
}
```

### Get State <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "numberOfPeers": 2748,
    "numberOfUnlockedAccounts": 0,
    "numberOfTransfers": 56779,
    "numberOfOrders": 2974,
    "numberOfTransactions": 7528488,
    "maxMemory": 1431830528,
    "isScanning": false,
    "cumulativeDifficulty": "29365339954194811259",
    "numberOfAssets": 371,
    "freeMemory": 617885288,
    "availableProcessors": 2,
    "totalEffectiveBalanceNXT": 1961368807,
    "numberOfAccounts": 147147,
    "numberOfBlocks": 502799,
    "requestProcessingTime": 6189,
    "version": "2.2.1",
    "numberOfBidOrders": 397,
    "lastBlock": "10431007917907627236",
    "totalMemory": 1167589376,
    "application": "BRS",
    "numberOfAliases": 17008,
    "lastBlockchainFeederHeight": 502798,
    "numberOfTrades": 109663,
    "time": 121621202,
    "numberOfAskOrders": 2577,
    "lastBlockchainFeeder": "169.55.184.235"
}
```

### Get Time <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "time": 121621260,
    "requestProcessingTime": 0
}
```

Transaction Operations
----------------------

### Broadcast Transaction <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "requestProcessingTime": 4,
    "fullHash": "d663f7fe8bf215906b29838ade5d860b2229188acea7827a737d823253b488db",
    "transaction": "10382471199064548310"
}
```

**Note**: If the transaction has already been broadcast, the following INFO notice appears in the console output and log file: *Transaction 15200507403046301754 already in blockchain (or unconfirmed pool), will not broadcast again*.

### Calculate Full Hash <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "requestProcessingTime": 0,
    "fullHash": "a8035211b0fb38415509ca451feee44787598462e7bc608affb8e7b2a1f81d05"
}
```

### Get Transaction <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "senderPublicKey": "a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12",
    "signature": "dc2503584a48e30ac62d62848f58461e0e9ff55070008743c24f380c24a9ef05525c70b5d40962566f3f4de2018277ba7956eb09d0aec84219784de7f3b76f6a",
    "feeNQT": "735000",
    "requestProcessingTime": 0,
    "type": 1,
    "confirmations": 7,
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
    "block": "2298247278137763160",
    "blockTimestamp": 121620850,
    "deadline": 24,
    "transaction": "3120851314369640207",
    "timestamp": 121620814,
    "height": 502797
}
```

### Get Transaction Bytes <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "unsignedTransactionBytes": "01104ec93f071800a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12b31119f931eaa6d4000000000000000018370b*",
    "requestProcessingTime": 0,
    "confirmations": 7,
    "transactionBytes": "01104ec93f071800a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12b31119f931eaa6d4000000000000000018370b*"
}
```

### Parse Transaction <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "senderPublicKey": "a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12",
    "signature": "dc2503584a48e30ac62d62848f58461e0e9ff55070008743c24f380c24a9ef05525c70b5d40962566f3f4de2018277ba7956eb09d0aec84219784de7f3b76f6a",
    "feeNQT": "735000",
    "requestProcessingTime": 238,
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
    "verify": true,
    "deadline": 24,
    "transaction": "3120851314369640207",
    "timestamp": 121620814,
    "height": 2147483647
}
```

### Sign Transaction <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "signatureHash": "2bcbafeab7a0bae40337fe34adea84110b1f770a33841c95f3fe9e19dde41bae",
    "verify": true,
    "requestProcessingTime": 2,
    "transactionBytes": "01104ec93f071800a61325eec9e83d7cac55544b8eca8ea8034559bafb5834b8a5d3b6d4efb85f12b31119f931eaa6d4000000000000000018370b*",
    "fullHash": "0f37d045bc7d4f2bd85cb565a5c4e575464ac387b986f80fb8c31635cf03923e",
    "transaction": "3120851314369640207"
}
```

Utilities <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />
-------------------------------------------------------------------------------------------------------

### Long Convert <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "stringId": "16922903237994405232",
    "requestProcessingTime": 0,
    "longId": "-1523840835715146384"
}
```

### RS Convert <img src="Verified.png" title="fig:Verified.png" alt="Verified.png" width="35" height="35" />

**Response:**

``` json
{
    "accountRS": "BURST-FRDJ-UPLH-MY9A-GUKQP",
    "requestProcessingTime": 0,
    "account": "16922903237994405232"
}
```

</translate>

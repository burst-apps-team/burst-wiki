Description
-----------

Examples of Burst API calls are collected on this page, individually linked from the main [Burst API](the-burst-api.md) page. The organization and ordering is the same for both pages so that the section numbers in the table of contents are identical. The preliminary sections preceding the examples simply link back to the main page. For example:

[The Burst API Description](the-burst-api-description.md)

Table Of Contents
-----------------

\_\_TOC\_\_

General Notes
-------------

[The Burst API General Notes](the-burst-api-general-notes.md)

Create Transaction
------------------

[The Burst API Create Transaction](the-burst-api-create-transaction.md)

Account Operations
------------------

### Delete Account Property

**Request:**

    http://localhost:8125/burst?
      requestType=deleteAccountProperty&
      recipient=BURST-7A48-47JL-T7LD-D5FS3&
      property=testkey1&
      secretPhrase=iWontTellYou&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
      "signatureHash": "4ff58a03d056ee8a3fee89766bf8e4acd008c2147216...",
      "transactionJSON": {
        "senderPublicKey": "373522bcd8904f4707472e590cbb67976d40e7af...",
        "signature": "26ed697fc82f3b15e6d2c972eff5b195445314aa4bacc8...",
        "feeNQT": "100000000",
        "type": 1,
        "fullHash": "33f7edaec1034153f8e28a996b13b2b2665d0d0a3e4a194...",
        "version": 1,
        "phased": false,
        "ecBlockId": "10023643060833833497",
        "signatureHash": "4ff58a03d056ee8a3fee89766bf8e4acd008c21472...",
        "attachment": {
            "property": "940296349549404868",
            "version.AccountPropertyDelete": 1
        },
        "senderRS": "BURST-7A48-47JL-T7LD-D5FS3",
        "subtype": 11,
        "amountNQT": "0",
        "sender": "12745647715474645062",
        "recipientRS": "BURST-7A48-47JL-T7LD-D5FS3",
        "recipient": "12745647715474645062",
        "ecBlockHeight": 754255,
        "deadline": 60,
        "transaction": "5999080309032613683",
        "timestamp": 80189128,
        "height": 2147483647
      },
      "unsignedTransactionBytes": "011bc896c7043c00373522bcd8904f4707472e590cbb67976d40e7af39650ea11c...",
      "broadcasted": false,
      "requestProcessingTime": 3,
      "transactionBytes": "011bc896c7043c00373522bcd8904f4707472e590cbb67976d40e7af39650ea11cb2be5734...",
      "fullHash": "33f7edaec1034153f8e28a996b13b2b2665d0d0a3e4a1942718aa480c6097cf6",
      "transaction": "5999080309032613683"
    }

<small>*Verified 9-Jun-16*</small>

### Get Account

**Request:**

    http://localhost:8125/burst?
      requestType=getAccount&
      account=BURST-4VNQ-RWZC-4WWQ-GVM8S

**Response:**

    {
     "unconfirmedBalanceNQT": "2501162882344",
     "effectiveBalanceNXT": 13983,
     "lessorsInfo": [
      {
       "currentHeightTo": "341420",
       "nextHeightFrom": "341420",
       "effectiveBalanceNXT": "544525",
       "nextLesseeRS": "BURST-7WVC-W7TJ-REQ2-4VDJD",
       "currentLesseeRS": "BURST-7WVC-W7TJ-REQ2-4VDJD",
       "currentHeightFrom": "308653",
       "nextHeightTo": "374187"
      }
     ],
     "currentLessee": "7114946486381367146",
     "currentLeasingHeightTo": 281179,
     "forgedBalanceNQT": "0",
     "balanceNQT": "2501162882344",
     "publicKey": "73080c6a224062660184f10ebb7fb431d4593...",
     "requestProcessingTime": 2,
     "assetBalances": [
      {
       "balanceQNT": "96651298",
       "asset": "4551058913252105307"
      }
     ],
     "guaranteedBalanceNQT": "1398383666344",
     "unconfirmedAssetBalances": [
      {
       "unconfirmedBalanceQNT": "96651298",
       "asset": "4551058913252105307"
      }
     ],
     "currentLesseeRS": "BURST-TMVC-69YC-SJB4-8YCH7",
     "accountRS": "BURST-4VNQ-RWZC-4WWQ-GVM8S",
     "name": "mystical",
     "account": "17013046603665206934",
     "currentLeasingHeightFrom": 279739
    }

<small>*Verified 7-Nov-14*</small>

### Get Account Block Count

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountBlockCount&
      account=7114946486381367146

**Response:**

    {
     "numberOfBlocks": 460,
     "requestProcessingTime": 70
    }

<small>*Verified 13-Nov-14*</small>

### Get Account Block Ids

**Request:**

<https://wallet.burst.cryptoguru.org:8125/burst?requestType=getAccountBlockIds&account=17274946210831421354&lastIndex=5>

**Response similar to:**

    {"blockIds":["18426055195962363092","7915836588136735869","10232042957216858995","6879436477641441477","10946006553575734351","11461420847451026714"],"requestProcessingTime":2}

<small>*Verified 11:28, 6 October 2017 (CEST)*</small>

### Get Account Blocks

**Request:**

<https://wallet.burst.cryptoguru.org:8125/burst?requestType=getAccountBlocks&account=17274946210831421354&lastIndex=0>

**Response similar to:**

    {"blocks":
    [{"previousBlockHash":"eba007c3393566ddad48cbf83604e9f6ad54dccf7722a1ca4a320988b1c6b3ed","payloadLength":578,"totalAmountNQT":"62450679293",
    "generationSignature":"3660dda874e1f4b4679a86226f4addbeb6a7816141de6d706260be002c5afd60","generator":"17274946210831421354",
    "generatorPublicKey":"2a689d7f067f6c5350c6df51ea5df1376ed9f60e72da8facf2240292a6c3c568","baseTarget":"107404",
    "payloadHash":"668b1467fbe7004c3ca9a22f9d31832da53c8b6fc34bdc958ec970bd674d313f","generatorRS":"BURST-EMXC-3PFB-J8HQ-GJ9HZ","blockReward":"1423","nextBlock":"13927268942427876599","scoopNum":2919,"numberOfTransactions":3,
    "blockSignature":"a44e8ae2f3ddc905ff43b7c0e9893cc2b100236ae2ba452e995b4865ccdcb90d308d712247fc8d764f3d5a9d211cc6a5f4b774ec4dbba36a705d74fac5a18106","transactions":["13877924564730410774","14965698104908109603","8929558255268593139"],"nonce":"76308248","version":3,
    "totalFeeNQT":"300000000","previousBlock":"15953497252208025835","block":"18426055195962363092",
    "height":411465,"timestamp":99544242}],"requestProcessingTime":3}

<small>*Verified 11:49, 6 October 2017 (CEST)*</small>

### Get Account Id

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountId&
      secretPhrase=IWontTellYou

**Response:**

    {
     "accountRS": "BURST-L6FM-89WK-VK8P-FCRBB",
     "publicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
     "requestProcessingTime": 2,
     "account": "15323192282528158131"
    }

<small>*Verified 7-Nov-14*</small>

### Get Account Ledger

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountLedger&
      account=BURST-7A48-47JL-T7LD-D5FS3

**Response:**

    {
      "entries": [
        {
          "change": "100",
          "eventType": "CURRENCY_OFFER_EXPIRED",
          "ledgerId": "532246",
          "holding": "6112509426732269765",
          "isTransactionEvent": true,
          "balance": "90000",
          "holdingType": "UNCONFIRMED_CURRENCY_BALANCE",
          "accountRS": "BURST-7A48-47JL-T7LD-D5FS3",
          "block": "2303344830040052747",
          "event": "13236557417702245931",
          "account": "12745647715474645062",
          "height": 736757,
          "timestamp": 78995635
        }
      ],
        "requestProcessingTime": 1
    }

<small>*Verified 9-Jun-16*</small>

### Get Account Ledger Entry

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountLedgerEntry&
      ledgerId=532246

**Response:**

    {
      "change": "100",
      "eventType": "CURRENCY_OFFER_EXPIRED",
      "requestProcessingTime": 1,
      "ledgerId": "532246",
      "holding": "6112509426732269765",
      "isTransactionEvent": true,
      "balance": "90000",
      "holdingType": "UNCONFIRMED_CURRENCY_BALANCE",
      "accountRS": "BURST-7A48-47JL-T7LD-D5FS3",
      "block": "2303344830040052747",
      "event": "13236557417702245931",
      "account": "12745647715474645062",
      "height": 736757,
      "timestamp": 78995635
    }

<small>*Verified 9-Jun-16*</small>

### Get Account Lessors

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountLessors&
      account=7114946486381367146&
      height=282497

**Response:**

    {
     "lessors": [
      {
       "guaranteedBalanceNQT": "2643314085738687",
       "lessorRS": "BURST-MRBN-8DFH-PFMK-A4DBM",
       "lessor": "9918441724915080500"
      }
     ],
     "accountRS": "BURST-TMVC-69YC-SJB4-8YCH7",
     "requestProcessingTime": 1,
     "account": "7114946486381367146",
     "height": 282497
    }

<small>*Verified 13-Nov-14*</small>

### Get Account Properties

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountProperties&
      recipient=BURST-7A48-47JL-T7LD-D5FS3

**Response:**

    {
      "recipientRS": "BURST-7A48-47JL-T7LD-D5FS3",
      "recipient": "12745647715474645062",
      "requestProcessingTime": 0,
      "properties": [
        {
          "setterRS": "BURST-7A48-47JL-T7LD-D5FS3",
          "property": "testkey1",
          "setter": "12745647715474645062",
          "value": "testvalue1"
        }
      ]
    }

<small>*Verified 9-Jun-16*</small>

### Get Account Public Key

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountPublicKey&
      account=BURST-L6FM-89WK-VK8P-FCRBB

**Response:**

    {
     "publicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d...",
     "requestProcessingTime": 36
    }

<small>*Verified 5-Nov-14*</small>

### Get Account Transaction Ids

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountTransactionIds&
      account=BURST-L6FM-89WK-VK8P-FCRBB

**Response:**

    {
     "transactionIds": [
      "15200507403046301754",
      "10900022216391397990"
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 6-Nov-14*</small>

### Get Account Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountTransactions&
      account=BURST-L6FM-89WK-VK8P-FCRBB

**Response:**

    {
     "requestProcessingTime": 1,
     "transactions": [
      {
       "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
       "signature": "5f0378b7390ff5a815eadd1354de533eef682f139362b153576e2207320a6...",
       "feeNQT": "100000000",
       "transactionIndex": 2,
       "type": 0,
       "confirmations": 1704,
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
      },
      {
       "senderPublicKey": "73080c6a224062660184f10ebb7fb431d459364a12403320c7f601f9d75cc547",
       "signature": "7f4a5b70e3f91dd1e7a089c7985bb08f7035666dbfe3e857e706f08ad93f6...",
       "feeNQT": "100000000",
       "transactionIndex": 0,
       "type": 0,
       "confirmations": 1706,
       "fullHash": "6612e07b74a84497b02d5bafea020391dcefadc157dc1cbd56611c66dc11f974",
       "version": 1,
       "ecBlockId": "4218793004869721496",
       "signatureHash": "0fc3d917e37111752004ac13a280ea121799388ff7aaf611f22f3ce93f1df5e0",
       "attachment": {
        "version.PublicKeyAnnouncement": 1,
        "recipientPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c"
       },
       "senderRS": "BURST-4VNQ-RWZC-4WWQ-GVM8S",
       "subtype": 0,
       "amountNQT": "200000000",
       "sender": "17013046603665206934",
       "recipientRS": "BURST-L6FM-89WK-VK8P-FCRBB",
       "recipient": "15323192282528158131",
       "ecBlockHeight": 275723,
       "block": "14241452309033661857",
       "blockTimestamp": 29796841,
       "deadline": 60,
       "transaction": "10900022216391397990",
       "timestamp": 29796542,
       "height": 275728
      }
     ]
    }

<small>*Verified 31-Dec-14*</small>

### Get Balance

**Request:**

    http://localhost:8125/burst?
      requestType=getBalance&
      account=7114946486381367146

**Response:**

    {
     "unconfirmedBalanceNQT": "9246231058415",
     "guaranteedBalanceNQT": "9242231058415",
     "effectiveBalanceNXT": 92422,
     "forgedBalanceNQT": "260560000000",
     "balanceNQT": "9246231058415",
     "requestProcessingTime": 1
    }

<small>*Verified 8-Nov-14*</small>

### Get Blockchain Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=getBlockchainTransactions&
      account=BURST-4VDY-LNVT-LMAY-FMCKA&
      lastIndex=0

**Response:**

    {
     "requestProcessingTime": 2,
     "transactions": [
      {
       "signature": "0bc2045c2e4291e9595702fc6a9e805f11b65a88a867d515e44b980ef72b440a2...",
       "transactionIndex": 0,
       "type": 0,
       "phased": false,
       "ecBlockId": "441034190304176853",
       "signatureHash": "31f92d5612115e174748c7a261cd0412e00028639301aae0f0c1ddfc7618b7e7",
       "attachment": {
        "version.OrdinaryPayment": 0
       },
       "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "subtype": 0,
       "amountNQT": "1400000000",
       "recipientRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "block": "6752144003309284467",
       "blockTimestamp": 46567698,
       "deadline": 1440,
       "timestamp": 46567612,
       "height": 291611,
       "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
       "feeNQT": "100000000",
       "confirmations": 2046,
       "fullHash": "a26a2a36086e5d13f069dd9da06ce4e6b0418e9a299bec0cda39bfa04a2ca5e3",
       "version": 1,
       "sender": "15295723609781267838",
       "recipient": "11580081983047651163",
       "ecBlockHeight": 291600,
       "transaction": "1395392441102264994"
      }
     ]
    }

<small>*Verified 19-May-15*</small>

### Get Guaranteed Balance

**Request:**

    http://localhost:8125/burst?
      requestType=getGuaranteedBalance&
      account=7114946486381367146&
      numberOfConfirmations=1440

**Response:**

    {
     "guaranteedBalanceNQT": "9242231058415",
     "requestProcessingTime": 0
    }

<small>*Verified 8-Nov-14*</small>

### Get Unconfirmed Transaction Ids

**Request:**

    http://localhost:8125/burst?
      requestType=getUnconfirmedTransactionIds&
      account=BURST-L6FM-89WK-VK8P-FCRBB

**Response:**

    {
     "requestProcessingTime": 1,
     "unconfirmedTransactionIds": []
    }

<small>*Verified 6-Nov-14*</small>

### Get Unconfirmed Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=getUnconfirmedTransactions&
      account=BURST-L6FM-89WK-VK8P-FCRBB

**Response:**

    {
     "unconfirmedTransactions": [],
     "requestProcessingTime": 1
    }

<small>*Verified 6-Nov-14*</small>

### Search Accounts

**Request:**

    http://localhost:8125/burst?
      requestType=searchAccounts&
      query=testnet AND tyler

**Response:**

    {
     "accounts": [
      {
       "accountRS": "BURST-7C4U-3Z9K-GZM8-CU8EJ",
       "name": "Tyler Jordan",
       "description": "testnet account",
       "account": "12119426358687475802"
      }
     ],
     "requestProcessingTime": 16
    }

<small>*Verified 29-Apr-15*</small>

### Send Money

**Request:**

    http://localhost:8125/burst?
      requestType=sendMoney&
      secretPhrase=IWontTellYou&
      recipient=BURST-4VNQ-RWZC-4WWQ-GVM8S&
      amountNQT=100000000&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "b35eae7d2f01639810d37694138aa0a86fbbf8a9bf58c2be4f2a5b8f0f30b3f7",
     "unsignedTransactionBytes": "001046aac6013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "5f0378b7390ff5a815eadd1354de533eef682f139362b153576e2207320a6...",
      "feeNQT": "100000000",
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
      "deadline": 60,
      "transaction": "15200507403046301754",
      "timestamp": 29796934,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 8475,
     "transactionBytes": "001046aac6013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143...",
     "fullHash": "3a304584f20cf3d2cbbdd9698ff9a166427005ab98fbe9ca4ad6253651ee81f1",
     "transaction": "15200507403046301754"
    }

<small>*Verified 4-Nov-14*</small>

### Set Account Info

**Request:**

    http://localhost:8125/burst?
      requestType=setAccountInfo&
      secretPhrase=IWontTellYou&
      name=iwonttellyou
      description=example account
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "15917aafd59ad9cece7dfc127ab256711d1c58a8ed1a0dc7334949ca826d8a32",
     "unsignedTransactionBytes": "0115dfeecb013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "3f5167a3a23677c85aba7fbbc8bc31cddf540a632abebee4a80fe08ba92b9a0...",
      "feeNQT": "100000000",
      "type": 1,
      "fullHash": "2d31c26aa2b0ae4cf233cc4035c555bca0c579bdcef24bc9819132dc2ce5b2e5",
      "version": 1,
      "ecBlockId": "17558522603047297060",
      "signatureHash": "15917aafd59ad9cece7dfc127ab256711d1c58a8ed1a0dc7334949ca826d8a32",
      "attachment": {
       "name": "iwonttellyou",
       "description": "example account",
       "version.AccountInfo": 1
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 5,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "ecBlockHeight": 279080,
      "deadline": 60,
      "transaction": "5525548004452479277",
      "timestamp": 30142175,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 8553,
     "transactionBytes": "0115dfeecb013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "fullHash": "2d31c26aa2b0ae4cf233cc4035c555bca0c579bdcef24bc9819132dc2ce5b2e5",
     "transaction": "5525548004452479277"
    }

<small>*Verified 8-Nov-14*</small>

### Set Account Property

**Request:**

    http://localhost:8125/burst?requestType=setAccountProperty&secretPhrase=
    iwonttellyou&property=country&value=switzerland&recipient=
    BURST-EZQ9-35P5-XZ8C-4LW87&feeNQT=100000000&deadline=60

The request is shown above in URL format for consistency. The actual request must be an HTTP POST request with a multipart content type. For example, the corresponding cURL command is as follows:

    [cURL command?]

**Response:**

    {
        "signatureHash": "34a690abc9f7fe4749da49d996e875e961aa0d38d6b80cdca9a01a7614004089",
        "transactionJSON": {
            "senderPublicKey": "b0db10704a831f1a0fd028e947784811e88fbe0fde25dc1f68209a2d9f93be13",
            "signature": "f5fd4b881c5625c778ffe38bced33d44231eda406e95d9068dea69407ea3370dc873f52a123a0cbe590ecf4adf2a3c6238dc26c73ee63aeb7ac4331958b53a1e",
            "feeNQT": "100000000",
            "type": 1,
            "fullHash": "d94baf764f0b7e47dbbe30c5aef1694fef6dfd2ecc7c86e29eefee7e6d1b2b39",
            "version": 1,
            "phased": false,
            "ecBlockId": "1026208230615777523",
            "signatureHash": "34a690abc9f7fe4749da49d996e875e961aa0d38d6b80cdca9a01a7614004089",
            "attachment": {
                "property": "\"switzerland\"",
                "value": "\"country\"",
                "version.AccountProperty": 1
            },
            "senderRS": "BURST-EZQ9-35P5-XZ8C-4LW87",
            "subtype": 10,
            "amountNQT": "0",
            "sender": "2493747385666535111",
            "recipientRS": "BURST-EZQ9-35P5-XZ8C-4LW87",
            "recipient": "2493747385666535111",
            "ecBlockHeight": 684730,
            "deadline": 360,
            "transaction": "5151567459679947737",
            "timestamp": 72072027,
            "height": 2147483647
        },
        "unsignedTransactionBytes": "011a5bbb4b046801b0db10704a831f1a0fd028e947784811e88fbe0fde25dc1f68209a2d9f93be13c77e36ea08929b22000000000000000000e1f5050000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ba720a00f3806fdbf2d23d0e010d22737769747a65726c616e64220922636f756e74727922",
        "broadcasted": true,
        "requestProcessingTime": 83,
        "transactionBytes": "011a5bbb4b046801b0db10704a831f1a0fd028e947784811e88fbe0fde25dc1f68209a2d9f93be13c77e36ea08929b22000000000000000000e1f505000000000000000000000000000000000000000000000000000000000000000000000000f5fd4b881c5625c778ffe38bced33d44231eda406e95d9068dea69407ea3370dc873f52a123a0cbe590ecf4adf2a3c6238dc26c73ee63aeb7ac4331958b53a1e00000000ba720a00f3806fdbf2d23d0e010d22737769747a65726c616e64220922636f756e74727922",
        "fullHash": "d94baf764f0b7e47dbbe30c5aef1694fef6dfd2ecc7c86e29eefee7e6d1b2b39",
        "transaction": "5151567459679947737"

<small>*Verified \[Date?\]*</small>

### Start Funding Monitor

**Request:**

    http://localhost:8125/burst?
     requestType=startFundingMonitor&
     property=funding&
     amount=1000000000&
     threshold=15000000000&
     interval=10&
     secretPhrase=IWontTellYou

**Response:**

    {
     "started": true,
     "requestProcessingTime": 5
    }

<small>*Verified 17-May-16*</small>

### Stop Funding Monitor

**Request:**

    http://localhost:8125/burst?
     requestType=stopFundingMonitor&
     property=funding&
     secretPhrase=IWontTellYou

**Response:**

    {
     "stopped": 1,
     "requestProcessingTime": 5
    }

<small>*Verified 17-May-16*</small>

Account Control Operations
--------------------------

### Get All Phasing Only Controls

**Request:**

    http://localhost:8125/burst?
      requestType=getAllPhasingOnlyControls&
      firstIndex=0&
      lastIndex=1

**Response:**

    {
     "phasingOnlyControls": [
      {
       "minDuration": 0,
       "votingModel": 0,
       "minBalance": "0",
       "accountRS": "BURST-AVGK-SKJZ-583G-A689A",
       "quorum": "2",
       "maxFees": "0",
       "whitelist": [
        {
         "whitelistedRS": "BURST-EVHD-5FLM-3NMQ-G46NR",
         "whitelisted": "16992224448242675179"
        },
        {
         "whitelistedRS": "BURST-XK4R-7VJU-6EQG-7R335",
         "whitelisted": "5873880488492319831"
        },
        {
         "whitelistedRS": "BURST-SZKV-J8TH-GSM9-9LKV6",
         "whitelisted": "8245583500397018683"
        }
       ],
       "minBalanceModel": 0,
       "account": "9519700060090428881",
       "maxDuration": 0
      },
      {
       "minDuration": 10,
       "votingModel": 0,
       "minBalance": "0",
       "accountRS": "BURST-VSTJ-MHZK-A6N3-CHRRH",
       "quorum": "1",
       "maxFees": "1000000000",
       "whitelist": [
        {
         "whitelistedRS": "BURST-D5A2-7CCE-G35P-BB324",
         "whitelisted": "10448396398360890624"
        },
        {
         "whitelistedRS": "BURST-KGPY-DF8U-HSF9-GQ69A",
         "whitelisted": "16437178058884561598"
        }
       ],
       "minBalanceModel": 0,
       "account": "12096369102442849072",
       "maxDuration": 100
      }
     ],
     "requestProcessingTime": 0
    }

<small>*Verified 14-Nov-16*</small>

### Get Phasing Only Control

**Request:**

    http://localhost:8125/burst?
      requestType=getPhasingOnlyControl&
      account=BURST-AVGK-SKJZ-583G-A689A

**Response:**

    {
     "minDuration": 0,
     "votingModel": 0,
     "minBalance": "0",
     "accountRS": "BURST-AVGK-SKJZ-583G-A689A",
     "quorum": "2",
     "maxFees": "0",
     "whitelist": [
      {
       "whitelistedRS": "BURST-EVHD-5FLM-3NMQ-G46NR",
       "whitelisted": "16992224448242675179"
      },
      {
       "whitelistedRS": "BURST-XK4R-7VJU-6EQG-7R335",
       "whitelisted": "5873880488492319831"
      },
      {
       "whitelistedRS": "BURST-SZKV-J8TH-GSM9-9LKV6",
       "whitelisted": "8245583500397018683"
      }
     ],
     "requestProcessingTime": 1,
     "minBalanceModel": 0,
     "account": "9519700060090428881",
     "maxDuration": 0
    }

<small>*Verified 14-Nov-16*</small>

### Set Phasing Only Control

**Request:**

    http://localhost:8125/burst?
     requestType=setPhasingOnlyControl&
     controlVotingModel=0&
     controlQuorum=1&
     controlWhitelisted=BURST-5MYN-AP7M-NKMH-CRQJZ&
     secretPhrase=IWontTellYou&
     feeNQT:100000000&
     deadline=60

**Response:**

    {
     "signatureHash": "bbb24087ab3639f508da77413a63ed8c45431ec9b7f0413fd574faea0de7e70a",
     "transactionJSON": {
      "senderPublicKey": "373522bcd8904f4707472e590cbb67976d40e7af39650ea11cb2be5734cdf30c",
      "signature": "eacaf3dc98c0e2ddc1b7b546209cecf092f16bb9e2fe9fd09760fb5815573f063...",
      "feeNQT": "100000000",
      "type": 4,
      "fullHash": "b6b7ac7a29f9f638244f1025c39199e8b5f8f7c5415f77ef0ca95c0f206e1e2d",
      "version": 1,
      "phased": false,
      "ecBlockId": "12092908062633562059",
      "signatureHash": "bbb24087ab3639f508da77413a63ed8c45431ec9b7f0413fd574faea0de7e70a",
      "attachment": {
       "version.SetPhasingOnly": 1,
       "controlMaxFees": "0",
       "controlMinDuration": 0,
       "controlMaxDuration": 0,
       "phasingControlParams": {
        "phasingHolding": "0",
        "phasingQuorum": 1,
        "phasingWhitelist": [
         "12664921794733526996"
        ],
        "phasingMinBalance": 0,
        "phasingMinBalanceModel": 0,
        "phasingVotingModel": 0
       }
      },
      "senderRS": "BURST-7A48-47JL-T7LD-D5FS3",
      "subtype": 1,
      "amountNQT": "0",
      "sender": "12745647715474645062",
      "ecBlockHeight": 767270,
      "deadline": 60,
      "transaction": "4104742066941900726",
      "timestamp": 80779288,
      "height": 2147483647
     },
     "unsignedTransactionBytes": "04111898d0043c00373522bcd8904f4707472e590cbb67976d40e7af39650ea...",
     "broadcasted": true,
     "requestProcessingTime": 5,
     "transactionBytes": "04111898d0043c00373522bcd8904f4707472e590cbb67976d40e7af39650ea1...",
     "fullHash": "b6b7ac7a29f9f638244f1025c39199e8b5f8f7c5415f77ef0ca95c0f206e1e2d",
     "transaction": "4104742066941900726"
    }

<small>*Verified 14-Nov-16*</small>

Alias Operations
----------------

### Buy / Sell Alias

**Request:**

    http://localhost:8125/burst?
      requestType=sellAlias&
      secretPhrase=IWontTellYou&
      aliasName=nextus&
      priceNQT=5&
      recipient=BURST-4VNQ-RWZC-4WWQ-GVM8S&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "d6f026cd8a883b5b6ff78a7d0121e4847eb6744b02757427de6d7ca0bf304226",
     "unsignedTransactionBytes": "01166e01d4013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "2041afc1687da2693ba092e33a84f665ad461e3b6762c18af61778261bc7e...",
      "feeNQT": "100000000",
      "type": 1,
      "fullHash": "14a3eeb17cd4082db287259a768d32065d4cf5397ed6053fffa25e92a8a66ac7",
      "version": 1,
      "ecBlockId": "1612829598027150491",
      "signatureHash": "d6f026cd8a883b5b6ff78a7d0121e4847eb6744b02757427de6d7ca0bf304226",
      "attachment": {
       "alias": "nextus",
       "priceNQT": "5",
       "version.AliasSell": 1
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 6,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "recipientRS": "BURST-4VNQ-RWZC-4WWQ-GVM8S",
      "recipient": "17013046603665206934",
      "ecBlockHeight": 284050,
      "deadline": 60,
      "transaction": "3245077163546682132",
      "timestamp": 30671214,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 8515,
     "transactionBytes": "01166e01d4013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143...",
     "fullHash": "14a3eeb17cd4082db287259a768d32065d4cf5397ed6053fffa25e92a8a66ac7",
     "transaction": "3245077163546682132"
    }

<small>*Verified 14-Nov-14*</small>

### Set Alias

**Request:**

    http://localhost:8125/burst?
      requestType=setAlias&
      secretPhrase=IWontTellYou&
      aliasName=iwonttellyou&
      aliasURI=acct:nxt-l6fm-89wk-vk8p-fcrbb@nxt&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "0dc7e07acef27fe86686cfabe2d1bd57c0c038f9465c3fe3d10f67932a97af10",
     "unsignedTransactionBytes": "011135d0d3013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "d92aefbab189b6e61f890c34b06a47e14a1a2b2ce868f77a4591d2067c51d...",
      "feeNQT": "100000000",
      "type": 1,
      "fullHash": "53917acbf44109391609a9bb57832c0d5903301e3d2bd6ffcf45cf893480f5a1",
      "version": 1,
      "ecBlockId": "4181883296304410027",
      "signatureHash": "0dc7e07acef27fe86686cfabe2d1bd57c0c038f9465c3fe3d10f67932a97af10",
      "attachment": {
       "alias": "iwonttellyou",
       "version.AliasAssignment": 1,
       "uri": "acct:nxt-l6fm-89wk-vk8p-fcrbb@nxt"
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 1,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "ecBlockHeight": 283939,
      "deadline": 60,
      "transaction": "4109888654593921363",
      "timestamp": 30658613,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 8104,
     "transactionBytes": "011135d0d3013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473...",
     "fullHash": "53917acbf44109391609a9bb57832c0d5903301e3d2bd6ffcf45cf893480f5a1",
     "transaction": "4109888654593921363"
    }

<small>*Verified 14-Nov-14*</small>

### Delete Alias

**Request:**

    https://localhost:7876/nxt?
      requestType=deleteAlias&
      aliasName=mystical

**Response:**

    {
     "signatureHash": "a6e68daed99c1015dd12546c042466612b52a9f5193d8513f7f12684aba5bf1d",
     "unsignedTransactionBytes": "011809e60c023c0010f09c34f225d425306e5be55a49469081...",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "a5088bc6469e528d884e7fd3c49afeefc8656dd59c9fa5ffeab2a17b465f6d03e77...",
      "feeNQT": "100000000",
      "type": 1,
      "fullHash": "63afc769e677b6210617ff7a9f5be2d7fe1aea7e46ccad968017d28df578fabf",
      "version": 1,
      "ecBlockId": "17895923487075501156",
      "signatureHash": "a6e68daed99c1015dd12546c042466612b52a9f5193d8513f7f12684aba5bf1d",
      "attachment": {
       "version.AliasDelete": 1,
       "alias": "mystical"
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 8,
      "amountNQT": "0",
      "sender": "15295723609781267838",
      "ecBlockHeight": 161137,
      "deadline": 60,
      "transaction": "2429260880513838947",
      "timestamp": 34399753,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 52,
     "transactionBytes": "011809e60c023c0010f09c34f225d425306e5be55a4946908156072...",
     "transaction": "2429260880513838947"
    }

<small>*Verified 27-Nov-14*</small>

### Get Alias

**Request:**

    https://localhost:7876/nxt?
      requestType=getAlias&
      alias=15515279700680480368

**Response:**

    {
     "aliasURI": "http:\/\/google.com",
     "aliasName": "google",
     "accountRS": "BURST-FLVS-VRBV-LDPD-6DZ9W",
     "alias": "15515279700680480368",
     "requestProcessingTime": 1,
     "account": "5629477397208681336",
     "timestamp": 2409343
    }

<small>*Verified 14-Nov-14*</small>

### Get Alias Count

**Request:**

    http://localhost:8125/burst?
      requestType=getAliasCount&
      account=BURST-FLVS-VRBV-LDPD-6DZ9W

**Response:**

    {
     "numberOfAliases": 200,
     "requestProcessingTime": 2
    }

<small>*Verified 21-Nov-14*</small>

### Get Aliases

**Request:**

    https://localhost:7876/nxt?
      requestType=getAliases&
      account=5629477397208681336&
      lastIndex=1

**Response:**

    {
     "aliases": [
      {
       "aliasURI": "",
       "aliasName": "101",
       "accountRS": "BURST-FLVS-VRBV-LDPD-6DZ9W",
       "alias": "8952438483248557843",
       "account": "5629477397208681336",
       "timestamp": 2409893
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 14-Nov-14*</small>

### Get Aliases Like

**Request:**

    https://localhost:7876/nxt?
      requestType=getAliasesLike&
      aliasPrefix=mysteri

**Response:**

    {
     "aliases": [
      {
       "aliasURI": "",
       "aliasName": "mysteries",
       "accountRS": "BURST-9DZL-XFE2-EWE5-HVVAY",
       "alias": "13234331415538245332",
       "account": "18384674354580664306",
       "timestamp": 4856820
      },
      {
       "aliasURI": "anm",
       "aliasName": "mysterious",
       "accountRS": "BURST-AHBB-DSVC-WS2L-EW8BC",
       "alias": "8033154744709486670",
       "account": "14205721421835156777",
       "timestamp": 2417903
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 29-Apr-15*</small>

Arbitrary Message System Operations
-----------------------------------

### Decrypt From

**Request:**

    http://localhost:8125/burst?
      requestType=decryptFrom&
      secretPhrase=IWontTellYou&
      account=BURST-L6FM-89WK-VK8P-FCRBB&
      data=5c30bd27cc86a8ab0349aaf66deae3c0a9db5675b5c4ba973dd47f37e06157...&
      nonce=7f3c9082c73a7bd825aa48d23fc138fd05a466700ff9fc3a040bbb29d3a60ee1&

**Response:**

    {
     "decryptedMessage": "test message",
     "requestProcessingTime": 2
    }

<small>*Verified 11-Nov-14*</small>

### Download Prunable Message

**Request:**

    http://localhost:8125/burst?
      requestType=downloadPrunableMessage&
      transaction=264609232955144528&
      retrieve=true

**Response:** The file in binary format.

<small>*Verified 17-Jun-16*</small>

### Encrypt To

**Request:**

    http://localhost:8125/burst?
      requestType=encryptTo&
      secretPhrase=IWontTellYou&
      recipient=BURST-L6FM-89WK-VK8P-FCRBB&
      messageToEncrypt=test message&

**Response:**

    {
     "data": "5c30bd27cc86a8ab0349aaf66deae3c0a9db5675b5c4ba973dd47f37e06157...",
     "requestProcessingTime": 48,
     "nonce": "7f3c9082c73a7bd825aa48d23fc138fd05a466700ff9fc3a040bbb29d3a60ee1"
    }

<small>*Verified 11-Nov-14*</small>

### Get All Prunable Messages

**Request:**

    http://localhost:8125/burst?
      requestType=getAllPrunableMessages&
      lastIndex=0

**Response:**

    {
     "prunableMessages": [
      {
       "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "sender": "15295723609781267838",
       "recipientRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "recipient": "11580081983047651163",
       "blockTimestamp": 46382992,
       "message": "This is a test prunable plain message.",
       "transaction": "4628485271017409467",
       "isText": true,
       "transactionTimestamp": 46382948
      }
     ],
     "requestProcessingTime": 0
    }

<small>*Verified 15-May-15*</small>

### Get Prunable Message

**Request:**

    http://localhost:8125/burst?
      requestType=getPrunableMessage&
      transaction=16832262845403902696&
      secretPhrase=secretPhrase

**Response:**

    {
     "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
     "encryptedMessage": {
      "data": "ba6baa8361ac5bdb9cb591cee616dc5801a32ddf05b66a4ee527cd8d57b0...",
      "nonce": "41f93e32997c70937a005e5b0b42546a1efa9ea9eb012f98d7a92d0c5a8855a4"
     },
     "sender": "15295723609781267838",
     "decryptedMessage": "test prunable encrypted message",
     "recipientRS": "BURST-BMUV-8QQR-47VK-CR7F3",
     "recipient": "11580081983047651163",
     "blockTimestamp": 46117919,
     "requestProcessingTime": 3,
     "transaction": "16832262845403902696",
     "encryptedMessageIsText": true,
     "transactionTimestamp": 46117594,
     "isCompressed": true
    }

<small>*Verified 12-May-15*</small>

### Get Prunable Messages

**Request:**

    http://localhost:8125/burst?
      requestType=getPrunableMessages&
      account=BURST-4VDY-LNVT-LMAY-FMCKA&
      lastIndex=0

**Response:**

    {
     "prunableMessages": [
      {
       "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "encryptedMessage": {
        "data": "ba6baa8361ac5bdb9cb591cee616dc5801a32ddf05b66a4ee527cd8d57b0a...",
        "nonce": "41f93e32997c70937a005e5b0b42546a1efa9ea9eb012f98d7a92d0c5a8855a4"
       },
       "sender": "15295723609781267838",
       "recipientRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "recipient": "11580081983047651163",
       "blockTimestamp": 46117919,
       "transaction": "16832262845403902696",
       "isText": true,
       "transactionTimestamp": 46117594,
       "isCompressed": true
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 12-May-15*</small>

### Get Shared Key

**Request:**

    http://localhost:8125/burst?
     requestType=getSharedKey&
     account=BURST-5MYN-AP7M-NKMH-CRQJZ&
     secretPhrase=IWontTellYou&
     nonce=0102030405060708091011121314151617181920212223242526272829303132

**Response:**

    {
     "sharedKey": "927118faa4850afa7fb3ced7b17eb4968ec4f1c0a405b0890552bb54a67d0eba",
     "requestProcessingTime": 1
    }

<small>*Verified 17-Nov-16*</small>

### Read Message

**Request:**

    http://localhost:8125/burst?
      requestType=readMessage&
      transaction=9908575668289607167&
      secretPhrase=IWontTellYou&

**Response:**

    {
     "requestProcessingTime": 1,
     "message": "Test message.",
     "decryptedMessage": "Test message (encrypted).",
     "decryptedMessageToSelf": "abc123"
    }

<small>*Verified 10-Nov-14*</small>

### Send Message

**Request:**

    http://localhost:8125/burst?
      requestType=sendMessage&
      secretPhrase=IWontTellYou&
      recipient=BURST-4VNQ-RWZC-4WWQ-GVM8S&
      message=Test Message.&
      deadline=60

**Response:**

    {
     "signatureHash": "795c58938a50d691f3f2b88bfaf03267236e972e1c068e0a5e11aeb606597f17",
     "unsignedTransactionBytes": "01100593ce013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c14...",
     "transactionJSON": {
      "senderPublicKey": "57fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b473beec228f79ff93c",
      "signature": "e916dbbfec51ca97ae76b1b190d1c74328f74c3c43ed3a06f1ca0ea250116...",
      "feeNQT": "100000000",
      "type": 1,
      "fullHash": "ff157b8a125582898b5c50d32a62f725602d5197af236fabcd6ec978b6861528",
      "version": 1,
      "ecBlockId": "6060075251340574063",
      "signatureHash": "795c58938a50d691f3f2b88bfaf03267236e972e1c068e0a5e11aeb606597f17",
      "attachment": {
       "version.Message": 1,
       "messageIsText": true,
       "message": "Test message."
      },
      "senderRS": "BURST-L6FM-89WK-VK8P-FCRBB",
      "subtype": 0,
      "amountNQT": "0",
      "sender": "15323192282528158131",
      "recipientRS": "BURST-4VNQ-RWZC-4WWQ-GVM8S",
      "recipient": "17013046603665206934",
      "ecBlockHeight": 280756,
      "deadline": 60,
      "transaction": "9908575668289607167",
      "timestamp": 30315269,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 11379,
     "transactionBytes": "01100593ce013c0057fb6f3a958e320bb49c4e81b4c2cf28b9f25d086c143b...",
     "fullHash": "ff157b8a125582898b5c50d32a62f725602d5197af236fabcd6ec978b6861528",
     "transaction": "9908575668289607167"
    }

<small>*Verified 15-Dec-14*</small>

### Verify Prunable Message

**Request:**

    http://localhost:8125/burst?
      requestType=verifyPrunableMessage&
      message=This is a test prunable plain message.

**Response:**

    {
     "version.PrunablePlainMessage": 1,
     "verify": true,
     "messageIsText": true,
     "messageHash": "da99da8026e30d971340ef54803543af3aa48ea215f80bd9375457bad8effb3f",
     "requestProcessingTime": 1,
     "message": "This is a test prunable plain message."
    }

<small>*Verified 15-May-15*</small>

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

### Delete Asset Shares

**Request:**

    http://localhost:8125/burst?
      requestType=deleteAssetShares&
      asset=5920455660623529270&
      quantityQNT=100&
      secretPhrase=IWontTellYou&
      feeNQT=100000000&
      deadline=60&
      broadcast=false

**Response:**

    {
        "signatureHash": "883599a340375c387e7b27a4f2c37f4ee960760b31e4b1a9c604663c6b84c708",
        "transactionJSON": {
          "senderPublicKey": "373522bcd8904f4707472e590cbb67976d40e7af39650ea11cb2be5734cdf30c",
          "signature": "b91a664d9ee99399ebba913fa56abea26fbf4b84fe10fa4f76a29a8eae16810e61...",
          "feeNQT": "100000000",
          "type": 2,
          "fullHash": "8929fa9d30fd38c8811e079835b3959c1a3f8503df8999d5a858831ea6ba7741",
          "version": 1,
          "phased": false,
          "ecBlockId": "9535861265974556985",
          "signatureHash": "883599a340375c387e7b27a4f2c37f4ee960760b31e4b1a9c604663c6b84c708",
          "attachment": {
            "quantityQNT": "100",
            "version.AssetDelete": 1,
            "asset": "5920455660623529270"
          },
          "senderRS": "BURST-7A48-47JL-T7LD-D5FS3",
          "subtype": 7,
          "amountNQT": "0",
          "sender": "12745647715474645062",
          "ecBlockHeight": 757498,
          "deadline": 60,
          "transaction": "14427559791532059017",
          "timestamp": 80204814,
          "height": 2147483647
        },
      "unsignedTransactionBytes": "02170ed4c7043c00373522bcd8904f4707472e590cbb67976d40e7a...",
      "broadcasted": false,
      "requestProcessingTime": 2,
      "transactionBytes": "02170ed4c7043c00373522bcd8904f4707472e590cbb67976d40e7af39650ea...",
      "fullHash": "8929fa9d30fd38c8811e079835b3959c1a3f8503df8999d5a858831ea6ba7741",
      "transaction": "14427559791532059017"
    }

<small>*Verified 9-Jun-16*</small>

### Dividend Payment

**Request:**

    http://localhost:8125/burst?
      requestType=dividendPayment&
      asset=3517042713515967694&
      height=161157&
      amountNQTPerQNT=100000000

**Response:**

    {
     "signatureHash": "5d92fee3570b7b058ace2387f5b4eef4377ea738e6a6c2aabc06bd9f6871e4b4",
     "unsignedTransactionBytes": "02160feb0c023c0010f09c34f225d425306e5be55a49469081...",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "61e2398087b078bdba24021d909e937c2d5fe157a09fa3c82d910e7b5843e00cef2...",
      "feeNQT": "100000000",
      "type": 2,
      "fullHash": "27b52e61813b73fdf9ff3b1bacf3b344d3ed60e8c94db2c9d0518c8483150770",
      "version": 1,
      "ecBlockId": "17871828515938613022",
      "signatureHash": "5d92fee3570b7b058ace2387f5b4eef4377ea738e6a6c2aabc06bd9f6871e4b4",
      "attachment": {
       "version.DividendPayment": 1,
       "amountNQTPerQNT": "100000000",
       "asset": "3517042713515967694",
       "height": 161157
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 6,
      "amountNQT": "0",
      "sender": "15295723609781267838",
      "ecBlockHeight": 161151,
      "deadline": 60,
      "transaction": "18263006340784764199",
      "timestamp": 34401039,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 111,
     "transactionBytes": "02160feb0c023c0010f09c34f225d425306e5be55a49469081560...",
     "fullHash": "27b52e61813b73fdf9ff3b1bacf3b344d3ed60e8c94db2c9d0518c8483150770",
     "transaction": "18263006340784764199"
    }

<small>*Verified 27-Dec-14*</small>

### Get Account Asset Count

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountAssetCount&
      account=BURST-8N9W-TN4F-YA2S-H5B7R

**Response:**

    {
     "requestProcessingTime": 1,
     "numberOfAssets": 3
    }

<small>*Verified 21-Nov-14*</small>

### Get Account Assets

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountAssets&
      account=BURST-4VNQ-RWZC-4WWQ-GVM8S

**Response:**

    {
     "accountAssets": [
      {
       "quantityQNT": "68013764",
       "unconfirmedQuantityQNT": "68013764",
       "decimals": 0,
       "name": "Test",
       "asset": "17554243582654188572"
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 18-Nov-14*</small>

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

### Get Asset Account Count

**Request:**

    http://localhost:8125/burst?
      requestType=getAssetAccountCount&
      asset=17554243582654188572

**Response:**

    {
     "numberOfAccounts": 38,
     "requestProcessingTime": 14
    }

<small>*Verified 19-Nov-14*</small>

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

### Get Asset Deletes

**Request:**

    http://localhost:8125/burst?
      requestType=getAssetDeletes&
      asset=5920455660623529270

**Response:**

    {
     "deletes": [
      {
       "quantityQNT": "1000",
       "assetDelete": "16186302132012496205",
       "accountRS": "BURST-G885-AKDX-5G2B-BLUCG",
       "asset": "5920455660623529270",
       "account": "10892890577210644675",
       "height": 678815,
       "timestamp": 75619271
      },
      {
       "quantityQNT": "1",
       "assetDelete": "5520627816808994883",
       "accountRS": "BURST-7A48-47JL-T7LD-D5FS3",
       "asset": "5920455660623529270",
       "account": "12745647715474645062",
       "height": 513529,
       "timestamp": 65930633
      },
      {
       "quantityQNT": "1",
       "assetDelete": "14523372185703177675",
       "accountRS": "BURST-7A48-47JL-T7LD-D5FS3",
       "asset": "5920455660623529270",
       "account": "12745647715474645062",
       "height": 512164,
       "timestamp": 65851395
      }
     ],
     "requestProcessingTime": 3
    }

<small>*Verified 16-Nov-16*</small>

### Get Asset Dividends

**Request:**

    http://localhost:8125/burst?
      requestType=getAssetDividends&
      asset=4348103880042995903&
      timestamp=105414366

**Response:**

    {
        "dividends": [
            {
                "assetDividend": "9638088763182480941",
                "numberOfAccounts": 391,
                "amountNQTPerQNT": "15",
                "totalDividend": "31399603179510",
                "dividendHeight": 1332177,
                "asset": "4348103880042995903",
                "height": 1332178,
                "timestamp": 111043612
            },
            {
                "assetDividend": "15350389444990902579",
                "numberOfAccounts": 374,
                "amountNQTPerQNT": "28",
                "totalDividend": "51709584601752",
                "dividendHeight": 1287433,
                "asset": "4348103880042995903",
                "height": 1287434,
                "timestamp": 108351797
            },
            {
                "assetDividend": "8213383200925147837",
                "numberOfAccounts": 309,
                "amountNQTPerQNT": "22",
                "totalDividend": "29425292831902",
                "dividendHeight": 1238568,
                "asset": "4348103880042995903",
                "height": 1238569,
                "timestamp": 105414366
            }
        ],
        "requestProcessingTime": 5
    }

<small>*Verified 19-Jun-17*</small>

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

### Get Expected Asset Deletes

**Request:**

    http://localhost:8125/burst?
      requestType=getExpectedAssetDeletes

**Response:**

    {
     "deletes": [
      {
       "quantityQNT": "1000",
       "assetDelete": "16186302132012496205",
       "accountRS": "BURST-G885-AKDX-5G2B-BLUCG",
       "asset": "5920455660623529270",
       "account": "10892890577210644675"
      }
     ],
     "requestProcessingTime": 3
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

### Get Expected Order Cancellations

**Request:**

    http://localhost:8125/burst?
      requestType=getExpectedOrderCancellations

**Response:**

    {
     "orderCancellations": [
      {
       "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "account": "15295723609781267838",
       "order": "8404616015717333294",
       "height": 348793,
       "phased": false
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 12-Jul-15*</small>

### Get Last Trades

**Request:**

    http://localhost:8125/burst?
      requestType=getLastTrades&
      assets=17091401215301664836

**Response:**

    {
     "trades": [
      {
       "seller": "11580081983047651163",
       "quantityQNT": "100",
       "bidOrder": "12461616895431889058",
       "sellerRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "buyer": "15295723609781267838",
       "priceNQT": "100000000",
       "askOrder": "16690422801364092687",
       "buyerRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "block": "9740784167963638799",
       "asset": "17091401215301664836",
       "askOrderHeight": 286247,
       "bidOrderHeight": 286243,
       "tradeType": "sell",
       "timestamp": 46115694,
       "height": 286247
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 15-Aug-15*</small>

### Get Order Trades

**Request:**

    http://localhost:8125/burst?
      requestType=getOrderTrades&
      askOrder=2769987326979385551&
      includeAssetInfo=true

**Response:**

    {
     "trades": [
      {
       "seller": "8069635474378047786",
       "quantityQNT": "200",
       "bidOrder": "6083013926058683287",
       "sellerRS": "BURST-5JBC-QQ8M-UAFJ-8UAZZ",
       "buyer": "3617506283101058376",
       "priceNQT": "5000000",
       "askOrder": "2769987326979385551",
       "buyerRS": "BURST-QXCA-TKAH-KK85-5QZE6",
       "decimals": 0,
       "name": "TWT",
       "block": "17820716559461579006",
       "asset": "7496917644161273018",
       "askOrderHeight": 263896,
       "bidOrderHeight": 265168,
       "tradeType": "buy",
       "timestamp": 44288754,
       "height": 265168
      }
     ],
     "requestProcessingTime": 0
    }

<small>*Verified 29-Apr-15*</small>

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

### Search Assets

**Request:**

    http://localhost:8125/burst?
      requestType=searchAssets&
      query=assets AND production

**Response:**

    {
     "assets": [
      {
       "quantityQNT": "2100000000000000",
       "numberOfAccounts": 37,
       "accountRS": "BURST-3TKA-UH62-478B-DQU6K",
       "decimals": 8,
       "numberOfTransfers": 84,
       "name": "mgwBTC",
       "description": "Production Multigateway BTC (mgwBTC) is backed 100% by...",
       "numberOfTrades": 15,
       "asset": "17554243582654188572",
       "account": "13300069592148796968"
      }
     ],
     "requestProcessingTime": 96
    }

<small>*Verified 21-Nov-14*</small>

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

### Get DGS Expired Purchases

**Request:**

    http://localhost:8125/burst?
      requestType=getDGSExpiredPurchases&
      seller=BURST-XK4R-7VJU-6EQG-7R335

**Response:**

    {
     "purchases": [
      {
       "seller": "5873880488492319831",
       "priceNQT": "1000000000",
       "quantity": 1,
       "deliveryDeadlineTimestamp": 44955165,
       "buyerRS": "BURST-XK4R-7VJU-6EQG-7R335",
       "pending": false,
       "purchase": "17272258199467687054",
       "name": "MyProduct",
       "goods": "4830545483228225683",
       "sellerRS": "BURST-XK4R-7VJU-6EQG-7R335",
       "buyer": "5873880488492319831",
       "timestamp": 44397533
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 29-Apr-15*</small>

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

### Get DGS Goods Count

**Request:**

    http://localhost:8125/burst?
      requestType=getDGSGoodsCount&
      seller=BURST-L6FM-89WK-VK8P-FCRBB

**Response:**

    {
     "numberOfGoods": 1,
     "requestProcessingTime": 2
    }

<small>*Verified 15-Dec-14*</small>

### Get DGS Goods Purchase Count

**Request:**

    http://localhost:8125/burst?
      requestType=getDGSGoodsPurchaseCount&
      goods=11813734897437346473

**Response:**

    {
     "numberOfPurchases": 1,
     "requestProcessingTime": 1
    }

<small>*Verified 15-Dec-14*</small>

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

### Get DGS Purchase Count

**Request:**

    http://localhost:8125/burst?
      requestType=getDGSPurchaseCount&
      seller=BURST-L6FM-89WK-VK8P-FCRBB

**Response:**

    {
     "numberOfPurchases": 2,
     "requestProcessingTime": 1
    }

<small>*Verified 15-Dec-14*</small>

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

### Get DGS Tag Count

**Request:**

    http://localhost:8125/burst?
      requestType=getDGSTagCount

**Response:**

    {
     "numberOfTags": 383,
     "requestProcessingTime": 472
    }

<small>*Verified 15-Dec-14*</small>

### Get DGS Tags

**Request:**

    http://localhost:8125/burst?
      requestType=getDGSTags&
      lastIndex=0

**Response:**

    {
     "requestProcessingTime": 2,
     "tags": [
      {
       "inStockCount": 40,
       "tag": "domains",
       "totalCount": 42
      }
     ]
    }

<small>*Verified 24-Nov-14*</small>

### Get DGS Tags Like

**Request:**

    http://localhost:8125/burst?
      requestType=getDGSTagsLike&
      tagPrefix=item

**Response:**

    {
     "requestProcessingTime": 1,
     "tags": [
      {
       "inStockCount": 1,
       "tag": "items",
       "totalCount": 1
      }
     ]
    }

<small>*Verified 29-Apr-15*</small>

### Search DGS Goods

**Request:**

    http://localhost:8125/burst?
      requestType=searchDGSGoods&
      tag=te?t AND prod*

**Response:**

    {
     "goods": [
      {
       "seller": "7580519603555678830",
       "quantity": 1,
       "goods": "1587116104511359906",
       "description": "This is a test. Please do not order.",
       "sellerRS": "BURST-6GMG-FC5F-YSX6-8CVEL",
       "delisted": false,
       "parsedTags": [
        "test",
        "product",
        "tag"
       ],
       "tags": "test,product,tag",
       "priceNQT": "100000000",
       "numberOfPublicFeedbacks": 1,
       "name": "Test Product",
       "numberOfPurchases": 1,
       "timestamp": 31611435
      }
     ],
     "requestProcessingTime": 4
    }

<small>*Verified 15-Dec-14*</small>

Forging Operations
------------------

### Start / Stop / Get Forging

**Request:**

    http://localhost:8125/burst?
      requestType=startForging&
      secretPhrase=IWontTellYou

**Response:**

    {
     "requestProcessingTime": 1,
     "deadline": 0,
     "hitTime": 0
    }

<small>*Verified 9-Nov-14*</small>

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

### Get Next Block Generators

**Request:**

    http://localhost:8125/burst?
      requestType=getNextBlockGenerators&
      limit=3

**Response:**

    {
     "activeCount": 216,
     "lastBlock": "10153073870267066931",
     "generators": [
      {
       "effectiveBalanceNXT": 5400786,
       "accountRS": "BURST-8HNT-4ZTF-ZXH3-7RU38",
       "deadline": 4,
       "account": "5982846390354787993",
       "hitTime": 112224054
      },
      {
       "effectiveBalanceNXT": 4061204,
       "accountRS": "BURST-HLNR-4HDK-HQUJ-37545",
       "deadline": 35,
       "account": "1263370831364868759",
       "hitTime": 112224085
      },
      {
       "effectiveBalanceNXT": 8847921,
       "accountRS": "BURST-ZEKG-CHYB-N8AR-4TQ3U",
       "deadline": 52,
       "account": "3244519536310858286",
       "hitTime": 112224102
      }
      ],
     "requestProcessingTime": 1,
     "timestamp": 112224050,
     "height": 1351776
    }

<small>*Verified 19-Jun-17*</small>

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

Monetary System Operations
--------------------------

### Can Delete Currency

**Request:**

    http://localhost:8125/burst?
      requestType=canDeleteCurrency&
      account=BURST-2HCZ-6GCJ-2XGV-EDRPH&
      currency=4923907272718555444

**Response:**

    {
     "canDelete": true,
     "requestProcessingTime": 1
    }

<small>*Verified 23-Dec-14*</small>

### Currency Buy / Sell

**Request:**

    http://localhost:8125/burst?
      requestType=currencyBuy&
      currency=6520756875632314476&
      rateNQT=1500000000&
      units=200&
      secretPhrase=SECRETPHRASE&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "0471907734d4aae7fc708131726f8660a68e66fe873fa17cab4f0cb3f879243f",
     "unsignedTransactionBytes": "051607cb08023c0010f09c34f225d425306e5be55a49469081...",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "dda6c065344827bfdecfee179b22dfee1d5280fa5da5431b012d9ea045a0d80b9d4...",
      "feeNQT": "100000000",
      "type": 5,
      "fullHash": "d388798c9ecaf2cd28578cfcb8fb7e8f07d308e36e2a5674bb1c7766595b8435",
      "version": 1,
      "ecBlockId": "10096636210021430702",
      "signatureHash": "0471907734d4aae7fc708131726f8660a68e66fe873fa17cab4f0cb3f879243f",
      "attachment": {
       "currency": "6520756875632314476",
       "version.ExchangeSell": 1,
       "units": "200",
       "rateNQT": "1500000000"
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 6,
      "amountNQT": "0",
      "sender": "15295723609781267838",
      "ecBlockHeight": 162266,
      "deadline": 60,
      "transaction": "14840146504449624275",
      "timestamp": 34130695,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 331,
     "transactionBytes": "051607cb08023c0010f09c34f225d425306e5be55a49469081560...",
     "fullHash": "d388798c9ecaf2cd28578cfcb8fb7e8f07d308e36e2a5674bb1c7766595b8435",
     "transaction": "14840146504449624275"
    }

<small>*Verified 24-Dec-14*</small>

### Currency Mint

**Request:**

    http://localhost:8125/burst?
      requestType=currencyMint&
      currency=9207767346829573996&
      nonce=-6757092571753666960&
      units=1&
      counter=26&
      secretPhrase=SECRETPHRASE&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "f0e9f3b040890043fbc2d1b235377eb566805419f459184bc9290051bacdf80f",
     "unsignedTransactionBytes": "0517655a0c023c0010f09c34f225d425306e5be55a49469081...",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "e961f58250dd256bb5f0628df342beb6ce042765d7cdad8d7598495759d7ec00ad8...",
      "feeNQT": "100000000",
      "type": 5,
      "fullHash": "05cbe451cdece14bab56c5684d30cd7eb62e45a86136071e90ae26981d4c1fc3",
      "version": 1,
      "ecBlockId": "12526333689713738846",
      "signatureHash": "f0e9f3b040890043fbc2d1b235377eb566805419f459184bc9290051bacdf80f",
      "attachment": {
       "currency": "9207767346829573996",
       "units": "1",
       "counter": "26",
       "version.CurrencyMinting": 1,
       "nonce": "-6757092571753666960"
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 7,
      "amountNQT": "0",
      "sender": "15295723609781267838",
      "ecBlockHeight": 160721,
      "deadline": 60,
      "transaction": "5467911789190892293",
      "timestamp": 34364005,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 815,
     "transactionBytes": "0517655a0c023c0010f09c34f225d425306e5be55a49469081560...",
     "fullHash": "05cbe451cdece14bab56c5684d30cd7eb62e45a86136071e90ae26981d4c1fc3",
     "transaction": "5467911789190892293"
    }

<small>*Verified 27-Dec-14*</small>

### Currency Reserve Claim

**Request:**

    http://localhost:8125/burst?
      requestType=currencyReserveClaim&
      currency=15992040603642022742&
      units=36

**Response:**

    {
     "signatureHash": "6daef8c9b8653a9b78b64e400c2ab64ea201e9888b8bc54ec6c9a28814b0a69e",
     "unsignedTransactionBytes": "0512a7830b023c00169cf83994b0e8c48a152ddc50606d58bd...",
     "transactionJSON": {
      "senderPublicKey": "169cf83994b0e8c48a152ddc50606d58bd4b2b85ec2f5bbbaae93d838443df7f",
      "signature": "5ab6605aca7a958b10e5f0198660eccd01bff1418f2bf172105ba39e9c911d09709...",
      "feeNQT": "100000000",
      "type": 5,
      "fullHash": "f32e02362dce767c6f4b069f793f2cd8c6b89bad04847c8abe98d8998cede94f",
      "version": 1,
      "ecBlockId": "14980969893438059909",
      "signatureHash": "6daef8c9b8653a9b78b64e400c2ab64ea201e9888b8bc54ec6c9a28814b0a69e",
      "attachment": {
       "version.ReserveClaim": 1,
       "currency": "15992040603642022742",
       "units": "36"
      },
      "senderRS": "BURST-BMUV-8QQR-47VK-CR7F3",
      "subtype": 2,
      "amountNQT": "0",
      "sender": "11580081983047651163",
      "ecBlockHeight": 160064,
      "deadline": 60,
      "transaction": "8968582401529884403",
      "timestamp": 34309031,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 7,
     "transactionBytes": "0512a7830b023c00169cf83994b0e8c48a152ddc50606d58bd4b2...",
     "fullHash": "f32e02362dce767c6f4b069f793f2cd8c6b89bad04847c8abe98d8998cede94f",
     "transaction": "8968582401529884403"
    }

<small>*Verified 26-Dec-14*</small>

### Currency Reserve Increase

**Request:**

    http://localhost:8125/burst?
      requestType=currencyReserveIncrease&
      currency=11847174313362984527
      amountPerUnitNQT=40000000

**Response:**

    {
     "signatureHash": "0ba69ec678e0d4c4f94ee576d9520bf13a1c88b87e46c79d937b8f998a4a1f54",
     "unsignedTransactionBytes": "05115c930b023c00169cf83994b0e8c48a152ddc50606d58bd...",
     "transactionJSON": {
      "senderPublicKey": "169cf83994b0e8c48a152ddc50606d58bd4b2b85ec2f5bbbaae93d838443df7f",
      "signature": "01ab2e01e8543321bb93baef1cdfd24335b2f98f4898c0203ebb372f6ea2140f609...",
      "feeNQT": "100000000",
      "type": 5,
      "fullHash": "0e583d097aa7832ef97f49bf1c43fe6fa26592cd7b60cd1c652b1139719c9404",
      "version": 1,
      "ecBlockId": "3894856940593871962",
      "signatureHash": "0ba69ec678e0d4c4f94ee576d9520bf13a1c88b87e46c79d937b8f998a4a1f54",
      "attachment": {
       "amountPerUnitNQT": "40000000",
       "currency": "11847174313362984527",
       "version.ReserveIncrease": 1
      },
      "senderRS": "BURST-BMUV-8QQR-47VK-CR7F3",
      "subtype": 1,
      "amountNQT": "0",
      "sender": "11580081983047651163",
      "ecBlockHeight": 160118,
      "deadline": 60,
      "transaction": "3351706690276644878",
      "timestamp": 34313052,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 7,
     "transactionBytes": "05115c930b023c00169cf83994b0e8c48a152ddc50606d58bd4b2...",
     "fullHash": "0e583d097aa7832ef97f49bf1c43fe6fa26592cd7b60cd1c652b1139719c9404",
     "transaction": "3351706690276644878"
    }

<small>*Verified 26-Dec-14*</small>

### Delete Currency

**Request:**

    http://localhost:8125/burst?
      requestType=deleteCurrency&
      currency=7103310507724273660&
      secretPhrase=SECRETPHRASE&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "3f2b029337263728f4d4fed1f774252b06706159659e6ba9116b2d74b67a8435",
     "unsignedTransactionBytes": "05189d7208023c0010f09c34f225d425306e5be55a49469081...",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "09c7be5b0742191f3777cc850f9478af21bd42225aec5e0793b57cd188767a0be6f...",
      "feeNQT": "100000000",
      "type": 5,
      "fullHash": "a877379e3f3713e2f8baaa777c9a48af23d7da74e00301f12a4051dcf747243f",
      "version": 1,
      "ecBlockId": "2688066685599408512",
      "signatureHash": "3f2b029337263728f4d4fed1f774252b06706159659e6ba9116b2d74b67a8435",
      "attachment": {
       "version.CurrencyDeletion": 1,
       "currency": "7103310507724273660"
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 8,
      "amountNQT": "0",
      "sender": "15295723609781267838",
      "ecBlockHeight": 162004,
      "deadline": 60,
      "transaction": "16290425023506118568",
      "timestamp": 34108061,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 112,
     "transactionBytes": "05189d7208023c0010f09c34f225d425306e5be55a49469081...",
     "fullHash": "a877379e3f3713e2f8baaa777c9a48af23d7da74e00301f12a4051dcf747243f",
     "transaction": "16290425023506118568"
    }

<small>*Verified 24-Dec-14*</small>

### Get Account Currencies

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountCurrencies&
      account=BURST-4VDY-LNVT-LMAY-FMCKA

**Response:**

    {
     "accountCurrencies": [
      {
       "issuerAccountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "code": "MYSTX",
       "unconfirmedUnits": "10000",
       "decimals": 2,
       "name": "MystcoinX",
       "currency": "6520756875632314476",
       "units": "10000",
       "issuanceHeight": 0,
       "type": 1,
       "issuerAccount": "15295723609781267838"
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 25-Dec-14*</small>

### Get Account Currency Count

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountCurrencyCount&
      account=BURST-4VDY-LNVT-LMAY-FMCKA

**Response:**

    {
     "numberOfCurrencies": 1,
     "requestProcessingTime": 1
    }

<small>*Verified 25-Dec-14*</small>

### Get Account Exchange Requests

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountExchangeRequests&
      account=BURST-4VDY-LNVT-LMAY-FMCKA&
      currency=6520756875632314476

**Response:**

    {
     "exchangeRequests": [
      {
       "issuerAccountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "code": "MYSTX",
       "subtype": 6,
       "decimals": 2,
       "name": "MystcoinX",
       "units": "200",
       "issuanceHeight": 0,
       "type": 1,
       "transaction": "14840146504449624275",
       "timestamp": 34130695,
       "rateNQT": "1500000000",
       "issuerAccount": "15295723609781267838"
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 24-Dec-14*</small>

### Get Funding Monitor

**Request:**

    http://localhost:8125/burst?
     requestType=getFundingMonitor&
     property=funding&
     secretPhrase=IWontTellYou&
     includeMonitoredAccounts:true

**Response:**

    {
     "requestProcessingTime": 0,
     "monitors": [
      {
       "holding": "0",
       "amount": "1000000000",
       "monitoredAccounts": [
        {
         "amount": "1000000000",
         "accountRS": "BURST-5MYN-AP7M-NKMH-CRQJZ",
         "threshold": "15000000000",
         "interval": 10,
         "account": "12664921794733526996"
        },
        {
         "amount": "1000000000",
         "accountRS": "BURST-G885-AKDX-5G2B-BLUCG",
         "threshold": "15000000000",
         "interval": 10,
         "account": "10892890577210644675"
        }
       ],
       "holdingType": 0,
       "accountRS": "BURST-7A48-47JL-T7LD-D5FS3",
       "property": "funding",
       "threshold": "15000000000",
       "interval": 10,
       "account": "12745647715474645062"
      }
     ]
    }

<small>*Verified 17-Jun-16*</small>

### Get All Currencies

**Request:**

    http://localhost:8125/burst?
      requestType=getAllCurrencies&
      firstIndex=2&
      lastIndex=2

**Response:**

    {
     "requestProcessingTime": 3,
     "currencies": [
      {
       "initialSupply": "10000",
       "currentReservePerUnitNQT": "0",
       "types": [
        "EXCHANGEABLE"
       ],
       "code": "MYSTX",
       "creationHeight": 162067,
       "minDifficulty": 0,
       "numberOfTransfers": 0,
       "description": "Exchangeable",
       "minReservePerUnitNQT": "0",
       "currentSupply": "10000",
       "issuanceHeight": 0,
       "type": 1,
       "reserveSupply": "0",
       "maxDifficulty": 0,
       "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "decimals": 2,
       "name": "MystcoinX",
       "numberOfExchanges": 12,
       "currency": "6520756875632314476",
       "maxSupply": "10000",
       "account": "15295723609781267838",
       "algorithm": 0
      }
     ]
    }

<small>*Verified 25-Dec-14*</small>

### Get All Exchanges

**Request:**

    http://localhost:8125/burst?
      requestType=getAllExchanges&
      firstIndex=7&
      lastIndex=7

**Response:**

    {
     "exchanges": [
      {
       "issuerAccountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "seller": "15295723609781267838",
       "code": "MYSTX",
       "sellerRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "units": "5000",
       "issuanceHeight": 0,
       "type": 1,
       "rateNQT": "10000000",
       "buyer": "11580081983047651163",
       "offer": "17204924627068821879",
       "buyerRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "decimals": 2,
       "name": "MystcoinX",
       "currency": "6520756875632314476",
       "block": "4816799421151726903",
       "transaction": "2086609620693258113",
       "timestamp": 34206117,
       "height": 163202,
       "issuerAccount": "15295723609781267838"
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 25-Dec-14*</small>

### Get Available To Buy

**Request:**

    http://localhost:8125/burst?
      requestType=getAvailableToBuy&
      currency=4855695375693311301&
      units=1

**Response:**

    {
        "amountNQT": "1087",
        "units": "1",
        "requestProcessingTime": 0,
        "rateNQT": "1087"
    }

<small>*Verified 16-Jun-16*</small>

### Get Buy / Sell Offers

**Request:**

    http://localhost:8125/burst?
      requestType=getSellOffers&
      currency=6520756875632314476

**Response:**

    {
     "offers": [
      {
       "offer": "4813417617929273983",
       "expirationHeight": 163125,
       "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "limit": "1000",
       "currency": "6520756875632314476",
       "supply": "500",
       "account": "15295723609781267838",
       "height": 163110,
       "rateNQT": "20000000"
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 25-Dec-14*</small>

### Get Currencies

**Request:**

    http://localhost:8125/burst?
      requestType=getCurrencies&
      currencies=6520756875632314476

**Response:**

    {
     "requestProcessingTime": 3,
     "currencies": [
      {
       "initialSupply": "10000",
       "currentReservePerUnitNQT": "0",
       "types": [
        "EXCHANGEABLE"
       ],
       "code": "MYSTX",
       "creationHeight": 162067,
       "minDifficulty": 0,
       "numberOfTransfers": 0,
       "description": "Exchangeable",
       "minReservePerUnitNQT": "0",
       "currentSupply": "10000",
       "issuanceHeight": 0,
       "type": 1,
       "reserveSupply": "0",
       "maxDifficulty": 0,
       "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "decimals": 2,
       "name": "MystcoinX",
       "numberOfExchanges": 12,
       "currency": "6520756875632314476",
       "maxSupply": "10000",
       "account": "15295723609781267838",
       "algorithm": 0
      }
     ]
    }

<small>*Verified 25-Dec-14*</small>

### Get Currencies By Issuer

**Request:**

    http://localhost:8125/burst?
      requestType=getCurrenciesByIssuer&
      account=BURST-4VDY-LNVT-LMAY-FMCKA

**Response:**

    {
     "requestProcessingTime": 334,
     "currencies": [
      [
       {
        "initialSupply": "10000",
        "currentReservePerUnitNQT": "0",
        "types": [
         "EXCHANGEABLE"
        ],
        "code": "MYSTX",
        "creationHeight": 162067,
        "minDifficulty": 0,
        "numberOfTransfers": 0,
        "description": "Exchangeable",
        "minReservePerUnitNQT": "0",
        "currentSupply": "10000",
        "issuanceHeight": 0,
        "type": 1,
        "reserveSupply": "0",
        "maxDifficulty": 0,
        "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
        "decimals": 2,
        "name": "MystcoinX",
        "numberOfExchanges": 12,
        "currency": "6520756875632314476",
        "maxSupply": "10000",
        "account": "15295723609781267838",
        "algorithm": 0
       }
      ]
     ]
    }

<small>*Verified 25-Dec-14*</small>

### Get Currency

**Request:**

    http://localhost:8125/burst?
      requestType=getCurrency&
      code=MYSTX

**Response:**

    {
     "initialSupply": "10000",
     "currentReservePerUnitNQT": "0",
     "types": [
      "EXCHANGEABLE"
     ],
     "code": "MYSTX",
     "creationHeight": 162067,
     "minDifficulty": 0,
     "numberOfTransfers": 0,
     "description": "Exchangeable",
     "minReservePerUnitNQT": "0",
     "currentSupply": "10000",
     "issuanceHeight": 0,
     "requestProcessingTime": 0,
     "type": 1,
     "reserveSupply": "0",
     "maxDifficulty": 0,
     "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
     "decimals": 2,
     "name": "MystcoinX",
     "numberOfExchanges": 0,
     "currency": "6520756875632314476",
     "maxSupply": "10000",
     "account": "15295723609781267838",
     "algorithm": 0
    }

<small>*Verified 23-Dec-14*</small>

### Get Currency Account Count

**Request:**

    http://localhost:8125/burst?
      requestType=getCurrencyAccountCount&
      currency=6520756875632314476

**Response:**

    {
     "numberOfAccounts": 2,
     "requestProcessingTime": 1
    }

<small>*Verified 25-Dec-14*</small>

### Get Currency Accounts

**Request:**

    http://localhost:8125/burst?
      requestType=getCurrencyAccounts&
      currency=6520756875632314476

**Response:**

    {
     "accountCurrencies": [
      {
       "unconfirmedUnits": "9000",
       "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "currency": "6520756875632314476",
       "units": "9000",
       "account": "15295723609781267838"
      },
      {
       "unconfirmedUnits": "1000",
       "accountRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "currency": "6520756875632314476",
       "units": "1000",
       "account": "11580081983047651163"
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 25-Dec-14*</small>

### Get Currency Founders

**Request:**

    http://localhost:8125/burst?
      requestType=getCurrencyFounders&
      currency=16165836410580103964

**Response:**

    {
     "founders": [
      {
       "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "amountPerUnitNQT": "100000000",
       "currency": "16165836410580103964",
       "account": "15295723609781267838"
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 27-Dec-14*</small>

### Get Currency Ids

**Request:**

    http://localhost:8125/burst?
      requestType=getCurrencyIds&
      lastIndex=2

**Response:**

    {
     "currencyIds": [
      "3543596621551215845",
      "6520756875632314476",
      "10304209318415949524"
     ],
     "requestProcessingTime": 2
    }

<small>*Verified 25-Dec-14*</small>

### Get Currency Transfers

**Request:**

    http://localhost:8125/burst?
      requestType=getCurrencyTransfers&
      currency=9387514940677621191

**Response:**

    {
     "transfers": [
      {
       "issuerAccountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "code": "MYSTX",
       "units": "10",
       "issuanceHeight": 0,
       "type": 1,
       "transfer": "12208608533071682262",
       "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "sender": "15295723609781267838",
       "recipientRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "decimals": 2,
       "recipient": "11580081983047651163",
       "name": "MystcoinX",
       "currency": "9387514940677621191",
       "issuerAccount": "15295723609781267838",
       "height": 159732,
       "timestamp": 34280949
      }
     ],
     "requestProcessingTime": 0
    }

<small>*Verified 23-Dec-14*</small>

### Get Exchanges

**Request:**

    http://localhost:8125/burst?
      requestType=getExchanges&
      currency=6520756875632314476

**Response:**

    {
     "exchanges": [
      {
       "issuerAccountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "seller": "15295723609781267838",
       "code": "MYSTX",
       "sellerRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "units": "1",
       "issuanceHeight": 0,
       "type": 1,
       "rateNQT": "1500000000",
       "buyer": "11580081983047651163",
       "offer": "7762792906174207279",
       "buyerRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "decimals": 2,
       "name": "MystcoinX",
       "currency": "6520756875632314476",
       "block": "2329341955641682831",
       "transaction": "12589829502215822061",
       "timestamp": 34137058,
       "height": 162356,
       "issuerAccount": "15295723609781267838"
      }
     ],
     "requestProcessingTime": 0
    }

<small>*Verified 23-Dec-14*</small>

### Get Exchanges By Exchange Request

**Request:**

    http://localhost:8125/burst?
      requestType=getExchangesByExchangeRequest&
      transaction=12589829502215822061

**Response:**

    {
     "exchanges": [
      {
       "issuerAccountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "seller": "15295723609781267838",
       "code": "MYSTX",
       "sellerRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "units": "1",
       "issuanceHeight": 0,
       "type": 1,
       "rateNQT": "1500000000",
       "buyer": "11580081983047651163",
       "offer": "7762792906174207279",
       "buyerRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "decimals": 2,
       "name": "MystcoinX",
       "currency": "6520756875632314476",
       "block": "2329341955641682831",
       "transaction": "12589829502215822061",
       "timestamp": 34137058,
       "height": 162356,
       "issuerAccount": "15295723609781267838"
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 24-Dec-14*</small>

### Get Exchanges By Offer

**Request:**

    http://localhost:8125/burst?
      requestType=getExchangesByOffer&
      offer=7762792906174207279

**Response:**

    {
     "exchanges": [
      {
       "issuerAccountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "seller": "15295723609781267838",
       "code": "MYSTX",
       "sellerRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "units": "1",
       "issuanceHeight": 0,
       "type": 1,
       "rateNQT": "1500000000",
       "buyer": "11580081983047651163",
       "offer": "7762792906174207279",
       "buyerRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "decimals": 2,
       "name": "MystcoinX",
       "currency": "6520756875632314476",
       "block": "2329341955641682831",
       "transaction": "12589829502215822061",
       "timestamp": 34137058,
       "height": 162356,
       "issuerAccount": "15295723609781267838"
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 24-Dec-14*</small>

### Get Last Exchanges

**Request:**

    http://localhost:8125/burst?
      requestType=getLastExchanges&
      currencies=12366259387060174981

**Response:**

    {
     "exchanges": [
      {
       "offer": "1659323201112235372",
       "seller": "15295723609781267838",
       "buyerRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "sellerRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "currency": "12366259387060174981",
       "block": "15911141024340031754",
       "units": "2",
       "transaction": "13926698078551037914",
       "timestamp": 35547397,
       "rateNQT": "100000000",
       "buyer": "15295723609781267838",
       "height": 167709
      }
     ],
     "requestProcessingTime": 0
    }

<small>*Verified 15-Aug-15*</small>

### Get Minting Target

**Request:**

    http://localhost:8125/burst?
      requestType=getMintingTarget&
      currency=9207767346829573996&
      account=BURST-4VDY-LNVT-LMAY-FMCKA&
      units=1

**Response:**

    {
     "difficulty": "1024",
     "targetBytes": "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff3f00",
     "currency": "9207767346829573996",
     "counter": 0,
     "requestProcessingTime": 1
    }

<small>*Verified 27-Dec-14*</small>

### Get Offer

**Request:**

    http://localhost:8125/burst?
      requestType=getOffer&
      offer=4813417617929273983

**Response:**

    {
     "sellOffer": {
      "offer": "4813417617929273983",
      "expirationHeight": 163125,
      "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "limit": "1000",
      "currency": "6520756875632314476",
      "supply": "500",
      "account": "15295723609781267838",
      "height": 163110,
      "rateNQT": "20000000"
     },
     "buyOffer": {
      "offer": "4813417617929273983",
      "expirationHeight": 163125,
      "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "limit": "1000",
      "currency": "6520756875632314476",
      "supply": "500",
      "account": "15295723609781267838",
      "height": 163110,
      "rateNQT": "10000000"
     },
     "requestProcessingTime": 0
    }

<small>*Verified 25-Dec-14*</small>

### Issue Currency

**Request:**

    http://localhost:8125/burst?
      requestType=issueCurrency&
      name=MystcoinX&
      code=MYSTX&
      description=Exchangeable
      type=1&
      initialSupply=10000&
      maxSupply=10000&
      decimals=2&
      secretPhrase=SECRETPHRASE&
      feeNQT=4000000000&
      deadline=60

**Response:**

    {
     "signatureHash": "c5ec66dd60bcc13fc0fe2cb617b4d7a05e7ef6360aacc591c911969785361491",
     "unsignedTransactionBytes": "05100f8308023c0010f09c34f225d425306e5be55a49469081...",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "da27e4bead869d0e96fb9370f0afcd0c49eb17b7da503705384239d9abaec209733...",
      "feeNQT": "4000000000",
      "type": 5,
      "fullHash": "6c08c076d4617e5a4be6f027c9b3e47ddb0e92fab87032d631546051dbdbfe8c",
      "version": 1,
      "ecBlockId": "2004481009157728964",
      "signatureHash": "c5ec66dd60bcc13fc0fe2cb617b4d7a05e7ef6360aacc591c911969785361491",
      "attachment": {
       "initialSupply": "10000",
       "code": "MYSTX",
       "minDifficulty": 0,
       "ruleset": 0,
       "description": "Exchangeable",
       "minReservePerUnitNQT": "0",
       "issuanceHeight": 0,
       "type": 1,
       "reserveSupply": "0",
       "version.CurrencyIssuance": 1,
       "maxDifficulty": 0,
       "decimals": 2,
       "name": "MystcoinX",
       "maxSupply": "10000",
       "algorithm": 0
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 0,
      "amountNQT": "0",
      "sender": "15295723609781267838",
      "ecBlockHeight": 162060,
      "deadline": 60,
      "transaction": "6520756875632314476",
      "timestamp": 34112271,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 53,
     "transactionBytes": "05100f8308023c0010f09c34f225d425306e5be55a49469081560...",
     "fullHash": "6c08c076d4617e5a4be6f027c9b3e47ddb0e92fab87032d631546051dbdbfe8c",
     "transaction": "6520756875632314476"
    }

<small>*Verified 24-Dec-14*</small>

### Publish Exchange Offer

**Request:**

    http://localhost:8125/burst?
      requestType=publishExchangeOffer&
      currency=6520756875632314476&
      buyRateNQT=10000000&
      sellRateNQT=20000000&
      totalBuyLimit=1000&
      totalSellLimit=1000&
      initialBuySupply=500&
      initialSellSupply=500&
      expirationHeight=163080&
      secretPhrase=SECRETPHRASE&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "6fd60fcbc6e7022f12adad82a76c0534d2f79a569b9857b02328bf7573cf93f8",
     "unsignedTransactionBytes": "0514d7c409023c0010f09c34f225d425306e5be55a49469081...",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "b204a7d9622ac6b99936447a4314fac1c2f0c67e45f4021f4b1949755c081802a34...",
      "feeNQT": "100000000",
      "type": 5,
      "fullHash": "fccb2c6200a04bbce82b9e623051b8d8ed9d519dbb0a12ec5a354c842da9a664",
      "version": 1,
      "ecBlockId": "6813726597245349906",
      "signatureHash": "6fd60fcbc6e7022f12adad82a76c0534d2f79a569b9857b02328bf7573cf93f8",
      "attachment": {
       "totalSellLimit": "1000",
       "buyRateNQT": "10000000",
       "initialSellSupply": "500",
       "totalBuyLimit": "1000",
       "expirationHeight": 163080,
       "sellRateNQT": "20000000",
       "version.PublishExchangeOffer": 1,
       "currency": "6520756875632314476",
       "initialBuySupply": "500"
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 4,
      "amountNQT": "0",
      "sender": "15295723609781267838",
      "ecBlockHeight": 163049,
      "deadline": 60,
      "transaction": "13568114225891298300",
      "timestamp": 34194647,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 54,
     "transactionBytes": "0514d7c409023c0010f09c34f225d425306e5be55a4946908156072...",
     "transaction": "13568114225891298300"
    }

<small>*Verified 25-Dec-14*</small>

### Search Currencies

**Request:**

    http://localhost:8125/burst?
      requestType=searchCurrencies&
      query=MYST?

**Response:**

    {
     "requestProcessingTime": 5,
     "currencies": [
      {
       "initialSupply": "10000",
       "currentReservePerUnitNQT": "0",
       "types": [
        "EXCHANGEABLE"
       ],
       "code": "MYSTX",
       "creationHeight": 159726,
       "minDifficulty": 0,
       "numberOfTransfers": 3,
       "description": "Exchangeable",
       "minReservePerUnitNQT": "0",
       "currentSupply": "10000",
       "issuanceHeight": 0,
       "type": 1,
       "reserveSupply": "0",
       "maxDifficulty": 0,
       "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "decimals": 2,
       "name": "MystcoinX",
       "numberOfExchanges": 0,
       "currency": "9387514940677621191",
       "maxSupply": "10000",
       "account": "15295723609781267838",
       "algorithm": 0
      }
     ]
    }

<small>*Verified 26-Dec-14*</small>

### Transfer Currency

**Request:**

    http://localhost:8125/burst?
      requestType=transferCurrency&
      recipient=BURST-BMUV-8QQR-47VK-CR7F3&
      currency=9387514940677621191&
      units=10

**Response:**

    {
     "signatureHash": "eda8f238d164074170ddb8163d7d71f08062f1f1a51712f05670a670e64e5eff",
     "unsignedTransactionBytes": "0513ad150b023c0010f09c34f225d425306e5be55a49469081...",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "7629cba6e776685c9e134f8d96c54a5bcad6915c433f833b1ca792c69e75ee0ac6f...",
      "feeNQT": "100000000",
      "type": 5,
      "fullHash": "d65a8adfc6b06da98ce50bb36d52305a8c9300b8defbeed9f2582628d9e258f0",
      "version": 1,
      "ecBlockId": "13770841487927223834",
      "signatureHash": "eda8f238d164074170ddb8163d7d71f08062f1f1a51712f05670a670e64e5eff",
      "attachment": {
       "currency": "9387514940677621191",
       "version.CurrencyTransfer": 1,
       "units": "10"
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 3,
      "amountNQT": "0",
      "sender": "15295723609781267838",
      "recipientRS": "BURST-BMUV-8QQR-47VK-CR7F3",
      "recipient": "11580081983047651163",
      "ecBlockHeight": 159721,
      "deadline": 60,
      "transaction": "12208608533071682262",
      "timestamp": 34280877,
      "height": 2147483647
     },
     "broadcasted": true,
     "requestProcessingTime": 8,
     "transactionBytes": "0513ad150b023c0010f09c34f225d425306e5be55a49469081560...",
     "fullHash": "d65a8adfc6b06da98ce50bb36d52305a8c9300b8defbeed9f2582628d9e258f0",
     "transaction": "12208608533071682262"
    }

<small>*Verified 26-Dec-14*</small>

Networking Operations
---------------------

### Add Peer

**Request:**

    http://localhost:8125/burst?
      requestType=addPeer&
      peer=nxt9.webice.ru

**Response:**

    {
     "hallmark": "4a8fd0f2af5f511e041ad7bee1625ae3a711ddcc056c60cf7ab3523e2c99092f0e006e...",
     "downloadedVolume": 155414,
     "address": "23.95.51.154",
     "weight": 19829,
     "uploadedVolume": 63532,
     "requestProcessingTime": 478,
     "version": "1.4.8",
     "platform": "webice.ru",
     "lastUpdated": 37014605,
     "blacklisted": false,
     "announcedAddress": "nxt9.webice.ru",
     "application": "BRS",
     "state": 1,
     "shareAddress": true
    }

<small>*Verified 26-Jan-15*</small>

### Blacklist API Proxy Peer

**Request:**

    http://localhost:8125/burst?
      requestType=blacklistAPIProxyPeer&
      peer=52.0.72.67

**Response:**

    {
     "requestProcessingTime": 24,
     "done": true
    }

<small>*Verified 19-Jun-17*</small>

### Blacklist Peer

**Request:**

    http://localhost:8125/burst?
      requestType=blacklistPeer&
      peer=nxt9.webice.ru

**Response:**

    {
     "requestProcessingTime": 0,
     "done": true
    }

<small>*Verified 26-Jan-15*</small>

### Get Inbound Peers

**Request:**

    http://localhost:8125/burst?
      requestType=getInboundPeers

**Response:**

    {
     "peers": [],
     "requestProcessingTime": 1
    }

<small>*Verified 19-May-15*</small>

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

### Set API Proxy Peer

**Request:**

    http://localhost:8125/burst?
      requestType=setAPIProxyPeer&
      peer=163.172.154.74

**Response:**

    {
     "downloadedVolume":6176,
     "address":"104.223.53.14",
     "inbound":false,
     "blockchainState":"UP_TO_DATE",
     "weight":0,
     "uploadedVolume":323,
     "services":[
      "HALLMARK","PRUNABLE","API","API_SSL","CORS"
      ],
     "requestProcessingTime":77,
     "version":"1.11.5",
     "platform":"BURST-TGFQ-U33C-C37U-CMKWF",
     "inboundWebSocket":false,
     "apiSSLPort":7878,
     "lastUpdated":112610887,
     "blacklisted":false,
     "announcedAddress":"104.223.53.14",
     "apiPort":7876,
     "application":"BRS",
     "port":7874,
     "outboundWebSocket":true,
     "lastConnectAttempt":112610887,
     "state":1,
     "shareAddress":true
    }

<small>*Verified 19-Jun-17*</small>

Phasing Operations
------------------

### Approve Transaction

**Request:**

    http://localhost:8125/burst?
      requestType=approveTransaction&
      transactionFullHash=5016cc59b0665675f00513e8c647288e0a668a78c4964c84d0de8f768b89060a&
      revealedSecretText=secret&
      secretPhrase=secretPhrase&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "db429ccecd7d13b54b43cf9db7656cef6df6152c60e626b393000ed00a652c95",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "380e1a94e40d58e9382aa742ca998373e27c5d30890d91a74d83bfead849e507e93c51...",
      "feeNQT": "100000000",
      "type": 1,
      "fullHash": "2bb80af156e70067f509df9ad5a88b687040cff4a8c778c69aef77863d3d15ef",
      "version": 1,
      "phased": false,
      "ecBlockId": "704052112466096836",
      "signatureHash": "db429ccecd7d13b54b43cf9db7656cef6df6152c60e626b393000ed00a652c95",
      "attachment": {
       "transactionFullHashes": [
        "5016cc59b0665675f00513e8c647288e0a668a78c4964c84d0de8f768b89060a"
       ],
       "version.PhasingVoteCasting": 1,
       "revealedSecret": "736563726574"
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 9,
      "amountNQT": "0",
      "sender": "15295723609781267838",
      "ecBlockHeight": 262493,
      "deadline": 60,
      "transaction": "7422186546503792683",
      "timestamp": 43965004,
      "height": 2147483647
     },
     "unsignedTransactionBytes": "01194cda9e023c0010f09c34f225d425306e5be55a494690...",
     "broadcasted": true,
     "requestProcessingTime": 20,
     "transactionBytes": "01194cda9e023c0010f09c34f225d425306e5be55a4946908156072a...",
     "fullHash": "2bb80af156e70067f509df9ad5a88b687040cff4a8c778c69aef77863d3d15ef",
     "transaction": "7422186546503792683"
    }

<small>*Verified 17-Apr-15*</small>

### Create Phasing Poll

**Request:**

    http://localhost:8125/burst?
      requestType=sendMoney&
      recipient=BURST-BMUV-8QQR-47VK-CR7F3&
      secretPhrase=secretPhrase&
      feeNQT=200000000&
      deadline=60&
      phased=true&
      phasingFinishHeight=261550&
      phasingVotingModel=0&
      phasingQuorum=2

**Response:**

    {
     "signatureHash": "df6c2dfcaf17f83256cfe388f408e091c08f208a060d54d1fbdb407ffdca121c",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "418a5453329d6c8d29b1b49dff30653a7f1e66992cece404d8ae997413deaa015b749a...",
      "feeNQT": "200000000",
      "type": 0,
      "fullHash": "6fd876512477ef4fab089ec2ffa084d6db75ba3cd0cef2541675515470dca374",
      "version": 1,
      "phased": true,
      "ecBlockId": "17522485785088843392",
      "signatureHash": "df6c2dfcaf17f83256cfe388f408e091c08f208a060d54d1fbdb407ffdca121c",
      "attachment": {
       "phasingFinishHeight": 261550,
       "phasingHolding": "0",
       "phasingQuorum": "2",
       "version.Phasing": 1,
       "phasingMinBalance": "0",
       "phasingMinBalanceModel": 0,
       "version.OrdinaryPayment": 0,
       "phasingVotingModel": 0
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 0,
      "amountNQT": "2000000000",
      "sender": "15295723609781267838",
      "recipientRS": "BURST-BMUV-8QQR-47VK-CR7F3",
      "recipient": "11580081983047651163",
      "ecBlockHeight": 261454,
      "deadline": 60,
      "transaction": "5759953446299424879",
      "timestamp": 43874749,
      "height": 2147483647
     },
     "unsignedTransactionBytes": "0010bd799d023c0010f09c34f225d425306e5be55a494690...",
     "broadcasted": true,
     "requestProcessingTime": 760,
     "transactionBytes": "0010bd799d023c0010f09c34f225d425306e5be55a4946908156072a...",
     "fullHash": "6fd876512477ef4fab089ec2ffa084d6db75ba3cd0cef2541675515470dca374",
     "transaction": "5759953446299424879"
    }

<small>*Verified 16-Apr-15*</small>

### Get Account Phased Transaction Count

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountPhasedTransactionCount&
      account=15295723609781267838

**Response:**

    {
     "requestProcessingTime": 1,
     "numberOfPhasedTransactions": 3
    }

<small>*Verified 17-Apr-15*</small>

### Get Account Phased Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountPhasedTransactions&
      account=15295723609781267838&
      lastIndex=0

**Response:**

    {
     "requestProcessingTime": 2,
     "transactions": [
      {
       "signature": "803d3fad7076bfe3f48a8ced40bb7075539858bf2b9d23b7653671a7204e6108234d...",
       "transactionIndex": 0,
       "type": 0,
       "phased": true,
       "ecBlockId": "14167949999961480077",
       "signatureHash": "693b18675d813dcc2de1a889fd919d0c4a0eaa679df0b0f8b0ec703bc2e278d4",
       "attachment": {
        "phasingFinishHeight": 262450,
        "phasingHolding": "17091401215301664836",
        "phasingQuorum": "500",
        "version.Phasing": 1,
        "phasingMinBalance": "100",
        "phasingMinBalanceModel": 2,
        "version.OrdinaryPayment": 0,
        "phasingVotingModel": 2
       },
       "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "subtype": 0,
       "amountNQT": "500000000",
       "recipientRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "block": "11015494088798322289",
       "blockTimestamp": 43957876,
       "deadline": 60,
       "timestamp": 43957682,
       "height": 262418,
       "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
       "feeNQT": "2100000000",
       "confirmations": 27,
       "fullHash": "e57462c46447f8bd7a0bafd9ab65bde8743ca6c13213185271d7ea6c48118861",
       "version": 1,
       "sender": "15295723609781267838",
       "recipient": "11580081983047651163",
       "ecBlockHeight": 262409,
       "transaction": "13688769565509711077"
      }
     ]
    }

<small>*Verified 17-Apr-15*</small>

### Get Asset Phased Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=getAssetPhasedTransactions&
      asset=17091401215301664836

**Response:**

    {
     "requestProcessingTime": 2,
     "transactions": [
      {
       "signature": "803d3fad7076bfe3f48a8ced40bb7075539858bf2b9d23b7653671a7204e6108234d...",
       "transactionIndex": 0,
       "type": 0,
       "phased": true,
       "ecBlockId": "14167949999961480077",
       "signatureHash": "693b18675d813dcc2de1a889fd919d0c4a0eaa679df0b0f8b0ec703bc2e278d4",
       "attachment": {
        "phasingFinishHeight": 262450,
        "phasingHolding": "",
        "phasingQuorum": "500",
        "version.Phasing": 1,
        "phasingMinBalance": "100",
        "phasingMinBalanceModel": 2,
        "version.OrdinaryPayment": 0,
        "phasingVotingModel": 2
       },
       "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "subtype": 0,
       "amountNQT": "500000000",
       "recipientRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "block": "11015494088798322289",
       "blockTimestamp": 43957876,
       "deadline": 60,
       "timestamp": 43957682,
       "height": 262418,
       "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
       "feeNQT": "2100000000",
       "confirmations": 27,
       "fullHash": "e57462c46447f8bd7a0bafd9ab65bde8743ca6c13213185271d7ea6c48118861",
       "version": 1,
       "sender": "15295723609781267838",
       "recipient": "11580081983047651163",
       "ecBlockHeight": 262409,
       "transaction": "13688769565509711077"
      }
     ]
    }

<small>*Verified 18-Apr-15*</small>

### Get Currency Phased Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=getCurrencyPhasedTransactions&
      currency=12366259387060174981

**Response:**

    {
     "requestProcessingTime": 2,
     "transactions": [
      {
       "signature": "ec467483307c73c08e7e9195eeddf23ce129ce30d703881cca505cbb0569c2009...",
       "transactionIndex": 0,
       "type": 0,
       "phased": true,
       "ecBlockId": "7353294686633135686",
       "signatureHash": "12cc1d54a453c3d1231b991a54f3323db6a51c00387ca5480db268164d5b7cfa",
       "attachment": {
        "phasingFinishHeight": 263500,
        "phasingHolding": "12366259387060174981",
        "phasingQuorum": "100",
        "version.Phasing": 1,
        "phasingMinBalance": "10",
        "phasingMinBalanceModel": 3,
        "version.OrdinaryPayment": 0,
        "phasingVotingModel": 3
       },
       "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "subtype": 0,
       "amountNQT": "800000000",
       "recipientRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "block": "699750272975612223",
       "blockTimestamp": 44042909,
       "deadline": 60,
       "timestamp": 44042827,
       "height": 263430,
       "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
       "feeNQT": "4100000000",
       "confirmations": 2,
       "fullHash": "5799d135aacf48a317c813fea405c317e7ab90888afb701b92c17914744f244f",
       "version": 1,
       "sender": "15295723609781267838",
       "recipient": "11580081983047651163",
       "ecBlockHeight": 263424,
       "transaction": "11765882356459739479"
      }
     ]
    }

<small>*Verified 18-Apr-15*</small>

### Get Linked Phased Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=getLinkedPhasedTransactions&
      linkedFullHash=083eba49bb481c38752a14493f0d40b3c60635935f13dd1fb33f6831fb997079

**Response:**

    {
     "requestProcessingTime": 1,
     "transactions": [
      {
       "signature": "d20a8bb84dcb335182ddc81ceaa1c5464881137d32eb195805ad460c9fcdba0...",
       "transactionIndex": 0,
       "type": 0,
       "phased": true,
       "ecBlockId": "4351283327032663440",
       "signatureHash": "77dfa3ff467a3b253530ad1b7a227633fa21fa064efe1df8206690bafb06f42c",
       "attachment": {
        "phasingFinishHeight": 770000,
        "version.Message": 1,
        "phasingHolding": "0",
        "phasingQuorum": "1",
        "version.Phasing": 1,
        "messageIsText": true,
        "phasingLinkedFullHashes": [
         "083eba49bb481c38752a14493f0d40b3c60635935f13dd1fb33f6831fb997079"
        ],
        "phasingMinBalance": "0",
        "message": "I givef Nxt",
        "phasingMinBalanceModel": 0,
        "version.OrdinaryPayment": 0,
        "phasingVotingModel": 4
       },
       "senderRS": "BURST-7A48-47JL-T7LD-D5FS3",
       "subtype": 0,
       "amountNQT": "500000000",
       "recipientRS": "BURST-5MYN-AP7M-NKMH-CRQJZ",
       "block": "13425992991123879393",
       "blockTimestamp": 80870662,
       "deadline": 1440,
       "timestamp": 80870626,
       "height": 768864,
       "senderPublicKey": "373522bcd8904f4707472e590cbb67976d40e7af39650ea11cb2be...",
       "feeNQT": "300000000",
       "confirmations": 11,
       "fullHash": "a1d845b3daf08493a2b299639b71422a41fcf923e1b68a4533d2bd0e22ad694d",
       "version": 1,
       "sender": "12745647715474645062",
       "recipient": "12664921794733526996",
       "ecBlockHeight": 768850,
       "transaction": "10629885842602449057"
      }
     ]
    }

<small>*Verified 17-Jun-16*</small>

### Get Phasing Poll

**Request:**

    http://localhost:8125/burst?
      requestType=getPhasingPoll&
      transaction=15402897900571339064

**Response:**

    {
     "votingModel": 0,
     "quorum": "2",
     "transactionFullHash": "38292a530816c2d5693bf5d0afb20847a3c4d2f37c3665fd294a7d2fdc278d56",
     "finished": true,
     "whitelist": [
      {
       "whitelistedRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "whitelisted": "11580081983047651163"
      },
      {
       "whitelistedRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "whitelisted": "15295723609781267838"
      }
     ],
     "requestProcessingTime": 1,
     "result": "2",
     "approved": true,
     "minBalance": "0",
     "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
     "finishHeight": 259600,
     "minBalanceModel": 0,
     "transaction": "15402897900571339064",
     "account": "15295723609781267838",
     "hashedSecret": ""
    }

<small>*Verified 15-Apr-15*</small>

### Get Phasing Poll Vote

**Request:**

    http://localhost:8125/burst?
      requestType=getPhasingPollVote&
      transaction=12580288379938056583&
      account=15295723609781267838

**Response:**

    {
     "voterRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
     "voter": "15295723609781267838",
     "requestProcessingTime": 1,
     "transaction": "9761138556025135837"
    }

<small>*Verified 17-Apr-15*</small>

### Get Phasing Poll Votes

**Request:**

    http://localhost:8125/burst?
      requestType=getPhasingPollVotes&
      account=12580288379938056583

**Response:**

    {
     "votes": [
      {
       "voterRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "voter": "15295723609781267838",
       "transaction": "9761138556025135837"
      }
     ],
     "requestProcessingTime": 2
    }

<small>*Verified 17-Apr-15*</small>

### Get Phasing Polls

**Request:**

    http://localhost:8125/burst?
      requestType=getPhasingPolls&
      transaction=12580288379938056583&
      transaction=13688769565509711077&
      countVotes=true

**Response:**

    {
     "polls": [
      {
       "votingModel": 1,
       "quorum": "100000000000",
       "transactionFullHash": "87a92edaa02996aeaf59a619ca513563775870c2b8ca490b61b02535745cdd34",
       "finished": true,
       "whitelist": [],
       "result": "186700000000",
       "approved": true,
       "minBalance": "10000000000",
       "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "finishHeight": 262450,
       "minBalanceModel": 1,
       "transaction": "12580288379938056583",
       "account": "15295723609781267838",
       "hashedSecret": ""
      },
      {
       "votingModel": 2,
       "quorum": "500",
       "transactionFullHash": "e57462c46447f8bd7a0bafd9ab65bde8743ca6c13213185271d7ea6c48118861",
       "finished": true,
       "whitelist": [],
       "result": "500",
       "holding": "17091401215301664836",
       "approved": true,
       "minBalance": "100",
       "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "finishHeight": 262450,
       "minBalanceModel": 2,
       "transaction": "13688769565509711077",
       "account": "15295723609781267838",
       "hashedSecret": ""
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 17-Apr-15*</small>

### Get Voter Phased Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=getVoterPhasedTransactions&
      account=15295723609781267838

**Response:**

    {
     "requestProcessingTime": 1,
     "transactions": [
      {
       "signature": "e527fce6591049b61d232ebbf4171319ae4e208f34d8a055ffb09f07dec7d9033527...",
       "transactionIndex": 0,
       "type": 0,
       "phased": true,
       "ecBlockId": "13625660527605830055",
       "signatureHash": "2a3cc644b1cb48130a172de92a418d48522b45a25a36d4d48c2347f032e430ab",
       "attachment": {
        "phasingFinishHeight": 263850,
        "phasingHolding": "0",
        "phasingQuorum": "2",
        "version.Phasing": 1,
        "phasingWhitelist": [
         "11580081983047651163",
         "15295723609781267838"
        ],
        "phasingMinBalance": "0",
        "phasingMinBalanceModel": 0,
        "version.OrdinaryPayment": 0,
        "phasingVotingModel": 0
       },
       "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "subtype": 0,
       "amountNQT": "1100000000",
       "recipientRS": "BURST-BMUV-8QQR-47VK-CR7F3",
       "block": "16550954176569210781",
       "blockTimestamp": 44074579,
       "deadline": 1440,
       "timestamp": 44074556,
       "height": 263813,
       "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
       "feeNQT": "200000000",
       "confirmations": 0,
       "fullHash": "b54a7e14c910e0750af0b29c0328347e27411dbe5d158d5d358ee2165b781968",
       "version": 1,
       "sender": "15295723609781267838",
       "recipient": "11580081983047651163",
       "ecBlockHeight": 263804,
       "transaction": "8493807353039047349"
      }
     ]
    }

<small>*Verified 18-Apr-15*</small>

Server Information Operations
-----------------------------

### Event Register

**Request:**

    http://localhost:8125/burst?
      requestType=eventRegister

**Response:**

    {
     "registered": true,
     "requestProcessingTime": 5780
    }

<small>*Verified 16-May-15*</small>

### Event Wait

**Request:**

    http://localhost:8125/burst?
      requestType=eventWait

**Response:**

    {
     "requestProcessingTime": 0,
     "events": [
      {
       "name": "Block.BLOCK_PUSHED",
       "ids": [
        "11748297033830700369"
       ]
      },
      {
       "name": "Transaction.REMOVED_UNCONFIRMED_TRANSACTIONS",
       "ids": []
      },
      {
       "name": "Peer.ADDED_ACTIVE_PEER",
       "ids": [
        "178.150.207.53"
       ]
      }
     ]
    }

<small>*Verified 16-May-15*</small>

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

### Get Plugins

**Request:**

    http://localhost:8125/burst?
      requestType=getPlugins

**Response:**

    {
     "plugins": [
      "hello_world"
     ],
     "requestProcessingTime": 9
    }

<small>*Verified 28-Apr-15*</small>

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

Shuffling Operations
--------------------

### Get Account Shufflings

**Request:**

    http://localhost:8125/burst?
     getAccountShufflings&
     account=BURST-UZNP-6LBA-YQ38-3TEY8

**Response:**

    {
     "requestProcessingTime": 91,
     "shufflings": [
      {
       "blocksRemaining": 4372,
       "amount": "25000000000000",
       "shufflingStateHash": "be12c737ad86e551621f57ebddaaaa61e66cf0a206fabfe40538304ff674c7a7",
       "issuer": "11802441287912491934",
       "holding": "0",
       "stage": 0,
       "holdingType": 0,
       "participantCount": 7,
       "assigneeRS": "BURST-UZNP-6LBA-YQ38-3TEY8",
       "shuffling": "5901270965262160574",
       "registrantCount": 2,
       "assignee": "1403321141259239061",
       "issuerRS": "BURST-V4WY-U928-GRN8-CKPL9",
       "shufflingFullHash": "be12c737ad86e551621f57ebddaaaa61e66cf0a206fabfe40538304ff674c7a7"
      }
     ]
    }

<small>*Verified 11-Jun-16*</small>

### Get All Shufflings

**Request:**

    http://localhost:8125/burst?
     requestType=getAllShufflings

**Response:**

    {
     "requestProcessingTime": 81,
     "shufflings": [
      {
       "blocksRemaining": 226,
       "amount": "5000000000000",
       "shufflingStateHash": "9be2ed518e46d921df9f9f6425cc16a399f72a7833c58dbeb2a261e04620c392",
       "issuer": "16667202175165271068",
       "holding": "0",
       "stage": 0,
       "holdingType": 0,
       "participantCount": 5,
       "assigneeRS": "BURST-V4WY-U928-GRN8-CKPL9",
       "shuffling": "2439058250271679131",
       "registrantCount": 3,
       "assignee": "11802441287912491934",
       "issuerRS": "BURST-J52W-4BF7-L4J9-GEGMG",
       "shufflingFullHash": "9be2ed518e46d921df9f9f6425cc16a399f72a7833c58dbeb2a261e04620c392"
      }
     ]
    }

<small>*Verified 11-Jun-16*</small>

### Get Assigned Shufflings

**Request:**

    http://localhost:8125/burst?
     requestType=getAssignedShufflings&
     account=BURST-V4WY-U928-GRN8-CKPL9

**Response:**

    {
     "blocksRemaining": 20,
     "amount": "5000000000000",
     "shufflingStateHash": "9be2ed518e46d921df9f9f6425cc1...",
     "requestProcessingTime": 3,
     "issuer": "16667202175165271068",
     "holding": "0",
     "stage": 1,
     "holdingType": 0,
     "participantCount": 5,
     "assigneeRS": "BURST-V4WY-U928-GRN8-CKPL9",
     "shuffling": "2439058250271679131",
     "registrantCount": 5,
     "assignee": "11802441287912491934",
     "issuerRS": "BURST-J52W-4BF7-L4J9-GEGMG",
     "shufflingFullHash": "9be2ed518e46d921df9f9f6425cc16a399..."
    }

<small>*Verified 11-Jun-16*</small>

### Get Holding Shufflings

**Request:**

    http://localhost:8125/burst?
     requestType=getHoldingShufflings&
     holding=15344649963748848799&
     includeFinished=true

**Response:**

    {
     "requestProcessingTime": 4,
     "shufflings": [
      {
       "holding": "15344649963748848799",
       "blocksRemaining": 0,
       "amount": "1000",
       "stage": 4,
       "shufflingStateHash": "57ed7cc02fa5b1f289d2a345314be5ec557ad093a5fd3070a16aba952d30656f",
       "holdingType": 1,
       "participantCount": 5,
       "shuffling": "144176732027925694",
       "registrantCount": 1,
       "issuer": "17417706268123203912",
       "issuerRS": "BURST-A4CA-L7JT-ZYGU-HZ2G5",
       "shufflingFullHash": "beb46651f937000259819298872f1721170a4c2e1af7f7900cb2ca01188fd942"
      }
     ]
    }

<small>*Verified 11-Jun-16*</small>

### Get Shufflers

**Request:**

    http://localhost:8125/burst?
     requestType=getShufflers&
     adminPassword=IWontTellYou

**Response:**

    {
     "shufflers": [
      {
       "accountRS": "BURST-7A48-47JL-T7LD-D5FS3",
       "recipientRS": "BURST-H42E-48UL-FM7A-AVEC2",
       "recipient": "9235083011255928844",
       "shuffling": "9534657251385467319",
       "account": "12745647715474645062",
       "shufflingFullHash": "b7e1eedb40e851849cb77a53075ce74aa..."
      }
     ],
     "requestProcessingTime": 0
    }

<small>*Verified 11-Jun-16*</small>

### Get Shuffling

**Request:**

    http://localhost:8125/burst?
     requestType=getShuffling&
     shuffling=3347171811359602131

**Response:**

    {
     "blocksRemaining": 956,
     "amount": "100000000000",
     "shufflingStateHash": "d3354a42078b732e16561b67a37d699d0410c564a4d18568319170665bf5dd9b",
     "requestProcessingTime": 2,
     "issuer": "10740484011813680725",
     "holding": "0",
     "stage": 0,
     "holdingType": 0,
     "participantCount": 5,
     "assigneeRS": "BURST-V4WY-U928-GRN8-CKPL9",
     "shuffling": "3347171811359602131",
     "registrantCount": 2,
     "assignee": "11802441287912491934",
     "issuerRS": "BURST-PFLP-BQHY-NG4J-BXG5C",
     "shufflingFullHash": "d3354a42078b732e16561b67a37d699d0410c564a4d18568319170665bf5dd9b"
    }

<small>*Verified 10-Jun-16*</small>

### Get Shuffling Participants

**Request:**

    http://localhost:8125/burst?
     requestType=getShufflingParticipants&
     shuffling=16292433427943984172

**Response:**

    {
     "requestProcessingTime": 5,
     "participants": [
      {
       "nextAccountRS": "BURST-KTTY-WRE6-2B5Y-4CYWR",
       "accountRS": "BURST-3ZXY-KWR5-LR8V-2PYAT",
       "shuffling": "16292433427943984172",
       "state": 2,
       "nextAccount": "3167098051496568638",
       "account": "910806360590974910"
      },
      {
       "nextAccountRS": "BURST-L8VW-NRUG-LZ6Q-62WWW",
       "accountRS": "BURST-KTTY-WRE6-2B5Y-4CYWR",
       "shuffling": "16292433427943984172",
       "state": 2,
       "nextAccount": "5653003407546981244",
       "account": "3167098051496568638"
      },
      {
       "nextAccountRS": "BURST-XC7L-35JW-FZCL-24NUT",
       "accountRS": "BURST-L8VW-NRUG-LZ6Q-62WWW",
       "shuffling": "16292433427943984172",
       "state": 2,
       "nextAccount": "930699247665195186",
       "account": "5653003407546981244"
      },
      {
       "nextAccountRS": "BURST-2222-2222-2222-22222",
       "accountRS": "BURST-XC7L-35JW-FZCL-24NUT",
       "shuffling": "16292433427943984172",
       "state": 2,
       "nextAccount": "0",
       "account": "930699247665195186"
      }
     ]
    }

<small>*Verified 14-Jun-16*</small>

### Shuffling Create

**Request:**

    http://localhost:8125/burst?
     requestType=shufflingCreate&
     amount=10000000000000&
     participantCount=7&
     registrationPeriod=9757&
     secretPhrase=IWontTellYou&
     feeNQT=100000000&
     deadline=1440 

**Response:**

    {
     "senderPublicKey": "ffcafb6de1ec93db5fe0475a6e2efd82faa3cb89a47a7d96f5579b0deadd5c10",
     "signature": "36b410b83abb4e373bf77c715c8ef3b06ba6c872e621e3f5f412dc1424eb6a0710ed8...",
     "feeNQT": "100000000",
     "transactionIndex": 0,
     "requestProcessingTime": 5,
     "type": 7,
     "confirmations": 9906,
     "fullHash": "2d10f4e18f808aa8d0de06d5cddcd77828eb6e7548f62e48de83f58fdf0629d5",
     "version": 1,
     "phased": false,
     "ecBlockId": "13021232786632787054",
     "signatureHash": "18956d76e3ed20937027373615e5fb49f11d1ac86eed279176856b59644899f5",
     "attachment": {
      "holding": "0",
      "amount": "10000000000000",
      "registrationPeriod": 9757,
      "holdingType": 0,
      "participantCount": 7,
      "version.ShufflingCreation": 1
     },
     "senderRS": "BURST-V4WY-U928-GRN8-CKPL9",
     "subtype": 0,
     "amountNQT": "0",
     "sender": "11802441287912491934",
     "ecBlockHeight": 817230,
     "block": "435349863926964285",
     "blockTimestamp": 80029572,
     "deadline": 1440,
     "transaction": "12144660700617510957",
     "timestamp": 80029541,
     "height": 817245
    }

<small>*Verified 14-Jun-16*</small>

### Shuffling Process

**Request:**

    http://localhost:8125/burst?
     requestType=shufflingProcess&
     shuffling=12144660700617510957&
     recipientSecretPhrase=IWontTellYou&
     secretPhrase=IWontTellYou&
     feeNQT=100000000&
     deadline=1440 

**Response:**

    {
       "senderPublicKey": "ffcafb6de1ec93db5fe0475a6e2efd82faa3cb89a47a7d96f5579b0deadd5c10",
       "signature": "fb6112d759cedfefe14a7bdbe19b2fac73d334566f05fc20d7d79933879b140...",
       "feeNQT": "1000000000",
       "transactionIndex": 1,
       "type": 7,
       "confirmations": 251,
       "fullHash": "0d0855396fa1a60c8b50507069abdb5a7c6af3fdb49f12766c8630ec66309642",
       "version": 1,
       "phased": false,
       "ecBlockId": "17529486445813261105",
       "signatureHash": "a7c1d67e2ff01dd8e3cf72869428a708a6ad087bb6247e628fb1964456eea66b",
       "attachment": {
        "shuffling": "12144660700617510957",
        "shufflingStateHash": "e98b2810531c6e807e5e89bb2537025bc7b4fb72e79ab675deb6fec2b8ec6d78",
        "data": [
         "438e2560cd958a173dc3b7adc3eb3c02a69ea5..."
        ],
        "version.ShufflingProcessing": 1,
        "hash": "2e3cecabbf59a47d55399649be89d1bcd1b56264dca32cb269f88b30cbe2c262"
       },
       "senderRS": "BURST-V4WY-U928-GRN8-CKPL9",
       "subtype": 2,
       "amountNQT": "0",
       "sender": "11802441287912491934",
       "ecBlockHeight": 826902,
       "block": "15948287569415806210",
       "blockTimestamp": 80611249,
       "deadline": 1440,
       "transaction": "911593473664419853",
       "timestamp": 80611213,
       "height": 826917
      }

<small>*Verified 14-Jun-16*</small>

### Shuffling Verify

**Request:**

    http://localhost:8125/burst?
     requestType=shufflingVerify&
     shuffling=12144660700617510957&
     shufflingStateHash=90832470918374087102938470198234&
     recipientSecretPhrase=IWontTellYou&
     secretPhrase=IWontTellYou&
     feeNQT=100000000&
     deadline=1440 

**Response:**

    {
     "senderPublicKey": "f050f86b050c1de933d8094c9a05a346b2319912313a04b822bba15d6d256505",
     "signature": "a23968d9cf5313e745fa4b50102f7f6d664c1b574d8ff2fbcc721cb967a42a0225f1e...",
     "feeNQT": "100000000",
     "transactionIndex": 3,
     "requestProcessingTime": 3,
     "type": 7,
     "confirmations": 287,
     "fullHash": "296a76a1601da4a06c6524df5791553b5d305cb266f71afb6e9d7aa0149c04ff",
     "version": 1,
     "phased": false,
     "ecBlockId": "3207126357685519085",
     "signatureHash": "7b69ba9ffb10d725b693abb53157f274cbf42f2435e919821cda67f889f44e96",
     "attachment": {
      "version.ShufflingVerification": 1,
      "shuffling": "12144660700617510957",
      "shufflingStateHash": "fbc8b99c99c4e09203f1237771e38c7ee484c18a72572c3d1adae04da858fe92"
     },
     "senderRS": "BURST-KHDF-9KSG-2UCY-BUQZ4",
     "subtype": 4,
     "amountNQT": "0",
     "sender": "10484056935692287341",
     "ecBlockHeight": 826914,
     "block": "1046003298185760668",
     "blockTimestamp": 80612146,
     "deadline": 1440,
     "transaction": "11575409243111975465",
     "timestamp": 80611663,
     "height": 826925
    }

<small>*Verified 14-Jun-16*</small>

### Shuffling Register

**Request:**

    http://localhost:8125/burst?
     requestType=shufflingRegister&
     shufflingFullHash=2d10f4e18f808aa8d0de06d5cddcd77828eb6e7548f62e48de83f58fdf0629d5&
     secretPhrase=IWontTellYou&
     feeNQT=100000000&
     deadline=1440 

**Response:**

    {
     "senderPublicKey": "6351706ab8ea5c94243e6b0bff3be05f2e1b9accc83004af3e655a8fbc104735",
     "signature": "23bb57aee566ba2c27fe8c83ea4e88821f9f23c820371a39b5d49dbd327...",
     "feeNQT": "100000000",
     "transactionIndex": 1,
     "requestProcessingTime": 3,
     "type": 7,
     "confirmations": 9924,
     "fullHash": "467e60696e598fe5c075446897ff636886bf7c0e9...",
     "version": 1,
     "phased": false,
     "ecBlockId": "4928509411801335816",
     "signatureHash": "b0e945000b9010814f5399d9006a21939089db5616...",
     "attachment": {
      "version.ShufflingRegistration": 1,
      "shufflingFullHash": "2d10f4e18f808aa8d0de06d5cddcd77828eb6e7..."
     },
     "senderRS": "BURST-D78H-8QHV-DEP2-5TAGA",
     "subtype": 1,
     "amountNQT": "0",
     "sender": "3763066681748198607",
     "ecBlockHeight": 817256,
     "block": "12642960475175742738",
     "blockTimestamp": 80031214,
     "deadline": 1440,
     "transaction": "16541538287104327238",
     "timestamp": 80031109,
     "height": 817265
    }

<small>*Verified 14-Jun-16*</small>

### Start Shuffler

**Request:**

    http://localhost:8125/burst?
     requestType=startShuffler&
     shufflingFullHash=2d10f4e18f808aa8d0de06d5cddcd77828eb6e7548f62e48de83f58fdf0629d5&
     secretPhrase=IWontTellYou&
     recipientSecretPhrase=IWontTellYou

**Response:**

    {
     "shuffling": "12144660700617510957",
     "shufflingFullHash": "2d10f4e18f808aa8d0de06d5cddcd77828eb6e7548f62e48de83f58fdf0629d5",
     "account": "11802441287912491934",
     "accountRS": "BURST-V4WY-U928-GRN8-CKPL9",
     "recipient": "3763066681748198607",
     "recipientRS": "BURST-D78H-8QHV-DEP2-5TAGA",
     "requestProcessingTime": 7
    }

<small>*Verified 14-Jun-16*</small>

### Stop Shuffler

**Request:**

    http://localhost:8125/burst?
     requestType=stopShuffler&
     shufflingFullHash=2d10f4e18f808aa8d0de06d5cddcd77828eb6e7548f62e48de83f58fdf0629d5&
     secretPhrase=IWontTellYou

**Response:**

    {
     "stoppedShuffler": true,
     "requestProcessingTime": 7
    }

<small>*Verified 14-Jun-16*</small>

Tagged Data Operations
----------------------

### Download Tagged Data

**Request:**

    http://localhost:8125/burst?
      requestType=downloadTaggedData&
      transaction=9169681701986886056

**Response:** Downloaded file named *nxt* with contents:

    d4f167249340d6d746f49441b8ccdb1bd3521feb

<small>*Verified 18-May-15*</small>

### Extend Tagged Data

**Request:**

    http://localhost:8125/burst?
      requestType=extendTaggedData&
      transaction=9086193976300572942&
      secretPhrase=secretPhrase&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "96cee31072bcdc82c85710c78e86d42aa8eac34e90960d97dcb87e9212534cc4",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "1be49ba648dcb242620d34762243399053837a32e5c2f0a392b27e70fcf823073f0...",
      "feeNQT": "100000000",
      "type": 6,
      "fullHash": "77c6dcbf770f2ce04ca8ed14ba52cdff8f7e35515bc0a8440079acd156700fb6",
      "version": 1,
      "phased": false,
      "ecBlockId": "9233989544625577248",
      "signatureHash": "96cee31072bcdc82c85710c78e86d42aa8eac34e90960d97dcb87e9212534cc4",
      "attachment": {
       "version.TaggedDataExtend": 1,
       "filename": "",
       "data": "d4f167249340d6d746f49441b8ccdb1bd3521feb",
       "name": "Stargate SG1 full series",
       "channel": "torrent",
       "description": "Hash of the torrent.",
       "type": "",
       "taggedData": "9086193976300572942",
       "isText": true,
       "tags": "video, scifi, torrent"
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 1,
      "amountNQT": "0",
      "sender": "15295723609781267838",
      "ecBlockHeight": 304612,
      "deadline": 60,
      "transaction": "16153302970465568375",
      "timestamp": 47672016,
      "height": 2147483647
     },
     "unsignedTransactionBytes": "0611d06ad7023c0010f09c34f225d425306e5be55a494...",
     "broadcasted": true,
     "requestProcessingTime": 28,
     "transactionBytes": "0611d06ad7023c0010f09c34f225d425306e5be55a49469081560...",
     "fullHash": "77c6dcbf770f2ce04ca8ed14ba52cdff8f7e35515bc0a8440079acd156700fb6",
     "transaction": "16153302970465568375"
    }

<small>*Verified 30-May-15*</small>

### Get Account Tagged Data

**Request:**

    http://localhost:8125/burst?
      requestType=getAccountTaggedData&
      account=BURST-4VDY-LNVT-LMAY-FMCKA&
      lastIndex=0

**Response:**

    {
     "data": [
      {
       "data": "d4f167249340d6d746f49441b8ccdb1bd3521feb",
       "channel": "torrent",
       "description": "Hash of the torrent.",
       "type": "",
       "parsedTags": [
        "video",
        "scifi",
        "torrent"
       ],
       "transactionTimestamp": 47843875,
       "tags": "video, scifi, torrent",
       "filename": "",
       "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "name": "Stargate SG1 full series",
       "blockTimestamp": 47672045,
       "transaction": "9086193976300572942",
       "account": "15295723609781267838",
       "isText": true
      }
     ],
     "requestProcessingTime": 2
    }

<small>*Verified 30-May-15*</small>

### Get All Tagged Data

**Request:**

    http://localhost:8125/burst?
      requestType=getAllTaggedData&
      lastIndex=0

**Response:**

    {
     "requestProcessingTime": 1,
     "data": [
      {
       "data": "d4f167249340d6d746f49441b8ccdb1bd3521feb",
       "channel": "torrent",
       "description": "Hash of the torrent.",
       "type": "",
       "parsedTags": [
        "video",
        "scifi",
        "torrent"
       ],
       "transactionTimestamp": 47843875,
       "tags": "video, scifi, torrent",
       "filename": "",
       "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "name": "Stargate SG1 full series",
       "blockTimestamp": 47672045,
       "transaction": "9086193976300572942",
       "account": "15295723609781267838",
       "isText": true
      }
     ]
    }

<small>*Verified 30-May-15*</small>

### Get Channel Tagged Data

**Request:**

    http://localhost:8125/burst?
      requestType=getChannelTaggedData&
      channel=torrent&
      lastIndex=0

**Response:**

    {
     "data": [
      {
       "data": "d4f167249340d6d746f49441b8ccdb1bd3521feb",
       "channel": "torrent",
       "description": "Hash of the torrent.",
       "type": "",
       "parsedTags": [
        "video",
        "scifi",
        "torrent"
       ],
       "transactionTimestamp": 47843875,
       "tags": "video, scifi, torrent",
       "filename": "",
       "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "name": "Stargate SG1 full series",
       "blockTimestamp": 47672045,
       "transaction": "9086193976300572942",
       "account": "15295723609781267838",
       "isText": true
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 30-May-15*</small>

### Get Data Tag Count

**Request:**

    http://localhost:8125/burst?
      requestType=getDataTagCount

**Response:**

    {
     "numberOfDataTags": 3,
     "requestProcessingTime": 2
    }

<small>*Verified 18-May-15*</small>

### Get Data Tags

**Request:**

    http://localhost:8125/burst?
      requestType=getDataTags

**Response:**

    {
     "requestProcessingTime": 5,
     "tags": [
      {
       "count": 2,
       "tag": "video"
      },
      {
       "count": 1,
       "tag": "scifi"
      },
      {
       "count": 1,
       "tag": "torrent"
      }
     ]
    }

<small>*Verified 18-May-15*</small>

### Get Data Tags Like

**Request:**

    http://localhost:8125/burst?
      requestType=getDataTagsLike&
      tagPrefix=sci

**Response:**

    {
     "requestProcessingTime": 9,
     "tags": [
      {
       "count": 1,
       "tag": "scifi"
      }
     ]
    }

<small>*Verified 18-May-15*</small>

### Get Tagged Data

**Request:**

    http://localhost:8125/burst?
      requestType=getTaggedData&
      transaction=9086193976300572942

**Response:**

    {
     "data": "d4f167249340d6d746f49441b8ccdb1bd3521feb",
     "channel": "torrent",
     "description": "Hash of the torrent.",
     "requestProcessingTime": 3,
     "type": "",
     "parsedTags": [
      "video",
      "scifi",
      "torrent"
     ],
     "transactionTimestamp": 47843875,
     "tags": "video, scifi, torrent",
     "filename": "",
     "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
     "name": "Stargate SG1 full series",
     "blockTimestamp": 47672045,
     "transaction": "9086193976300572942",
     "account": "15295723609781267838",
     "isText": true
    }

<small>*Verified 30-May-15*</small>

### Get Tagged Data Extend Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=getTaggedDataExtendTransactions&
      transaction=4344513120863790923

**Response:**

    {
      "extendTransactions": [
        {
          "senderPublicKey": "631110059365218234cc28fb6c91019f4563cf910471a9e43221795f40ab6e3a",
          "signature": "e4fe406198843207f5255e2c98119b3130b9810812934e0406a4de97a3443b0b205c907284c2e2cfd2534e8dcc17f08afb707db338e557d02673a2e6b0010186",
          "feeNQT": "150000000",
          "transactionIndex": 0,
          "type": 6,
          "confirmations": 31778,
          "fullHash": "86e8bc8c891e1556dfb9114a0ba9c4e119e6df6ba241016a1e9a85223dea42f8",
          "version": 1,
          "phased": false,
          "ecBlockId": "14132969248050944989",
          "signatureHash": "2ed7b6a5d579a66596409b2619375fa0edd54ace2e7d10f0bdbc0a7126f1776c",
          "senderRS": "BURST-G885-AKDX-5G2B-BLUCG",
          "subtype": 1,
          "amountNQT": "0",
          "sender": "10892890577210644675",
          "ecBlockHeight": 725746,
          "block": "18314376737060578519",
          "blockTimestamp": 78357439,
          "deadline": 1440,
          "transaction": "6202897637893269638",
          "timestamp": 78357424,
          "height": 725756
        }
      ],
      "requestProcessingTime": 1
    }

<small>*Verified 9-Jun-16*</small>

### Search Tagged Data

**Request:**

    http://localhost:8125/burst?
      requestType=searchTaggedData&
      query=SG?&
      lastIndex=0

**Response:**

    {
     "data": [
      {
       "data": "d4f167249340d6d746f49441b8ccdb1bd3521feb",
       "channel": "torrent",
       "description": "Hash of the torrent.",
       "type": "",
       "parsedTags": [
        "video",
        "scifi",
        "torrent"
       ],
       "transactionTimestamp": 47843875,
       "tags": "video, scifi, torrent",
       "filename": "",
       "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
       "name": "Stargate SG1 full series",
       "blockTimestamp": 47672045,
       "transaction": "9086193976300572942",
       "account": "15295723609781267838",
       "isText": true
      }
     ],
     "requestProcessingTime": 22
    }

<small>*Verified 30-May-15*</small>

### Upload Tagged Data

**Request:**

    http://localhost:8125/burst?
      requestType=uploadTaggedData&
      data=d4f167249340d6d746f49441b8ccdb1bd3521feb&
      name=Stargate SG1 full series&
      description=Hash of the torrent.&
      tags=video, scifi, torrent&
      channel=torrent&
      secretPhrase=secretPhrase&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "953a1023c84939a6e31da0db913ed2829b318d59663daddce63bdf2193af912e",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "17da947a7dc30a6d722a53f064d817c559f8f9c1387290048793d81b49dbf5082a0...",
      "feeNQT": "100000000",
      "type": 6,
      "fullHash": "0e91b0dd47a5187e26ea5b54f1917e3b879ee3cf096eb597ce22c6833863cdb2",
      "version": 1,
      "phased": false,
      "ecBlockId": "7868019187120003959",
      "signatureHash": "953a1023c84939a6e31da0db913ed2829b318d59663daddce63bdf2193af912e",
      "attachment": {
       "filename": "",
       "data": "d4f167249340d6d746f49441b8ccdb1bd3521feb",
       "name": "Stargate SG1 full series",
       "channel": "torrent",
       "description": "Hash of the torrent.",
       "type": "",
       "version.TaggedDataUpload": 1,
       "isText": true,
       "hash": "06fbac5b5358c00f5a2f19789b06220dca4e242a851b70072633a06ae8e6fb46",
       "tags": "video, scifi, torrent"
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 0,
      "amountNQT": "0",
      "sender": "15295723609781267838",
      "ecBlockHeight": 304598,
      "deadline": 60,
      "transaction": "9086193976300572942",
      "timestamp": 47671075,
      "height": 2147483647
     },
     "unsignedTransactionBytes": "06102367d7023c0010f09c34f225d425306e5be55a494...",
     "broadcasted": true,
     "requestProcessingTime": 284,
     "transactionBytes": "06102367d7023c0010f09c34f225d425306e5be55a49469081560...",
     "fullHash": "0e91b0dd47a5187e26ea5b54f1917e3b879ee3cf096eb597ce22c6833863cdb2",
     "transaction": "9086193976300572942"
    }

<small>*Verified 30-May-15*</small>

### Verify Tagged Data

**Request:**

    http://localhost:8125/burst?
      requestType=verifyTaggedData&
      transaction=9086193976300572942&
      data=d4f167249340d6d746f49441b8ccdb1bd3521feb&
      name=Stargate SG1 full series&
      description=Hash of the torrent.&
      tags=video, scifi, torrent&
      channel=torrent

**Response:**

    {
     "verify": true,
     "requestProcessingTime": 7,
     "version.TaggedDataUpload": 1,
     "hash": "06fbac5b5358c00f5a2f19789b06220dca4e242a851b70072633a06ae8e6fb46"
    }

<small>*Verified 30-May-15*</small>

Token Operations
----------------

### Decode File Token

**Request:**

    http://localhost:8125/burst?
      requestType=decodeFileToken&
      file=test.txt&
      token=u8q9ps0gdoo2bl158p4llpar583ld0cgejat9qnrgrgde4l5uscgan7fu25hi...

The request is shown above in URL format for consistency. The actual request must be an HTTP POST request with a multipart content type. For example, the corresponding cURL command is as follows:

    curl -F requestType=generateFileToken -F file=@test.txt -F secretPhrase="secretPhrase" http://localhost:8125/burst

**Response:**

    {
    "valid": true,
    "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
    "requestProcessingTime": 3,
    "account": "15295723609781267838",
    "timestamp": 49748229
    }

<small>*Verified 23-Jun-15*</small>

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

### Generate File Token

**Request:**

    http://localhost:8125/burst?
      requestType=generateFileToken&
      secretPhrase=secretPhrase&
      file=test.txt

The request is shown above in URL format for consistency. The actual request must be an HTTP POST request with a multipart content type. For example, the corresponding cURL command is as follows:

    curl -F requestType=generateFileToken -F file=@test.txt -F secretPhrase="secretPhrase" http://localhost:8125/burst

**Response:**

    {
    "valid": true,
    "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
    "requestProcessingTime": 4,
    "account": "15295723609781267838",
    "timestamp": 49748229,
    "token": "u8q9ps0gdoo2bl158p4llpar583ld0cgejat9qnrgrgde4l5uscgan7fu25hi..."
    }

<small>*Verified 23-Jun-15*</small>

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

### Get Expected Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=getExpectedTransactions&
      account=BURST-X5A4-FTMJ-J97M-F3ANH

**Response:**

    {
     "expectedTransactions": [
     {
       "senderPublicKey": "aaf2ffc347dae1442f218c17ff9ae1e917ab32af0af779360737f6faf7a4023e",
       "signature": "58720e3cb6cd4a9ad1dcb0bede55952e624b02a125aa0ac5934d7312cf87660147b495...",
       "feeNQT": "100000000",
       "type": 2,
       "fullHash": "79df9c85bc30f4a72c5d1899ce57b6533f321e1412f433fc3975cd4142a3c781",
       "version": 1,
       "phased": false,
       "ecBlockId": "17663518922057624367",
       "signatureHash": "687701be20bfb0935b3165ae31c08b685a173be660a9720ed0319ec17cc646eb",
       "attachment": {
        "quantityQNT": "5",
        "priceNQT": "908000000",
        "asset": "8122396658538927693",
        "version.BidOrderPlacement": 1
       },
       "senderRS": "BURST-X5A4-FTMJ-J97M-F3ANH",
       "subtype": 3,
       "amountNQT": "0",
       "sender": "15551212561974070530",
       "ecBlockHeight": 820879,
       "deadline": 1440,
       "transaction": "12102351684905000825",
       "timestamp": 80248727,
       "height": 2147483647
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 9-Jun-16*</small>

### Get Referencing Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=getReferencingTransactions&
      transaction=9132239833429154693

**Response:**

    {
     "requestProcessingTime": 1,
     "transactions": [
      {
       "senderPublicKey": "36bb17bc15678804a95ed895d524bb361aa5dfc4e78a800901aefff783f48010",
       "signature": "f39849eee16cd534f86915e22bd41d91fd6e965278242c957422196454deb705aa4...",
       "feeNQT": "100000000",
       "transactionIndex": 0,
       "type": 1,
       "confirmations": 216565,
       "fullHash": "c96b60bafc581891be88eae695f773677fe3cf7322cc149831d70de046041ea7",
       "version": 1,
       "phased": false,
       "ecBlockId": "1861841686166492176",
       "referencedTransactionFullHash": "859b376bbe3bbc7ebd5520d0d3fd20256f921e4c6bc3e...",
       "signatureHash": "0ed3aaabc682d4b5242b69930c0f659b62776b92bc2eecc40e37c895c5260eb4",
       "attachment": {
        "version.Message": 1,
        "messageIsText": true,
        "message": "],\"type\":\"payment\",\"oracles\":[{\"name\":\"nayru\"}]}",
        "version.ArbitraryMessage": 0
       },
       "senderRS": "BURST-DRGK-5CLR-35AZ-84R5R",
       "subtype": 0,
       "amountNQT": "0",
       "sender": "7750380612824194513",
       "ecBlockHeight": 614622,
       "block": "13844289828761736561",
       "blockTimestamp": 67496993,
       "deadline": 1440,
       "transaction": "10455204377422490569",
       "timestamp": 67496806,
       "height": 614629
      }
     ]
    }

<small>*Verified 17-Jun-16*</small>

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

### Retrieve Pruned Transaction

**Request:**

    http://localhost:8125/burst?
      requestType=retrievePrunedTransaction&
      transaction=9343032000583494722

**Response:**

    {
     "senderPublicKey": "7c94b068c95edcaf6ad588cda8784c7c27421ac334be092a6b487885f40f4f0c",
     "signature": "a4b4b7829b90ae0aa6c6654d364d23b335c64383bacc0e518b5aa1f2746de706121d308ae11244bf0933eddaed170c958bb9520700d63a7ed9ab30eb766717ca",
     "feeNQT": "320000000",
     "transactionIndex": 0,
     "requestProcessingTime": 0,
     "type": 6,
     "confirmations": 41543,
     "fullHash": "42ece25f1a1ea9815a7e6856635c323810efd2a38686bd57b7241fce28841da0",
     "version": 1,
     "phased": false,
     "ecBlockId": "11517076094591827523",
     "signatureHash": "5202c62fafb4b990313f03cf8882c2bad73c7888790794bb8c46a186abfc34ca",
     "attachment": {
      "filename": "7dba799bc6d4439191f89a88a9e50dad.jpg",
      "data": "ffd8ffe000104a46494600010100000100010000ffdb00430003050805050404050a07070...",
      "name": "Darkwing Duck",
      "channel": "Darkwing Duck",
      "description": "Darkwing Duck",
      "type": "image/jpeg",
      "version.TaggedDataUpload": 1,
      "isText": false,
      "hash": "5fb2f476f2556d15e4d419c6ff10bdbd05a7c5284242eba67c4de1adaee185a2",
      "tags": "Darkwing Duck"
     },
     "senderRS": "BURST-K8HZ-VMMS-ZX8E-EY5JR",
     "subtype": 0,
     "amountNQT": "0",
     "sender": "14681874272470211071",
     "ecBlockHeight": 788531,
     "block": "17859315470736709713",
     "blockTimestamp": 78307385,
     "deadline": 1440,
     "transaction": "9343032000583494722",
     "timestamp": 78307322,
     "height": 788546
    }

<small>*Verified 16-Jun-16*</small>

### Send Transaction

**Request:**

    http://localhost:8125/burst?
      requestType=sendTransaction&
      transactionBytes=00109e61b606a0052bdd59320496b133052f58c82e8...

**Response:**

    {
     "requestProcessingTime": 2,
     "fullHash": "799ad836f9c65e2985978123f145130b867bab9a86f6fdabae8e8c8f25c087dc",
     "transaction": "2981038777035168377"
    }

<small>*Verified 19-Jun-17*</small>

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

Voting System Operations
------------------------

### Cast Vote

**Request:**

    http://localhost:8125/burst?
      requestType=castVote&
      poll=5916389507928675673&
      vote02=1&
      secretPhrase=secretPhrase&
      feeNQT=100000000&
      deadline=60

**Response:**

    {
     "signatureHash": "54fb93860d0668477d34f7828d6a2c19be1c9229e7cb5adcebdc4962319ab22b",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "d48873e936e74069a56ba0eecdbfe3abe706bb7443e51d6651f59ef2500694087bd...",
      "feeNQT": "100000000",
      "type": 1,
      "fullHash": "ae7b714aadfe5efee0d690dfffac533dbf190827ed2f972d3dd552f4399ed648",
      "version": 1,
      "phased": false,
      "ecBlockId": "1999174047197129566",
      "signatureHash": "54fb93860d0668477d34f7828d6a2c19be1c9229e7cb5adcebdc4962319ab22b",
      "attachment": {
       "version.VoteCasting": 1,
       "poll": "5916389507928675673",
       "vote": [
        -128,
        -128,
        1
       ]
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 3,
      "amountNQT": "0",
      "sender": "15295723609781267838",
      "ecBlockHeight": 256208,
      "deadline": 60,
      "transaction": "18329367553676245934",
      "timestamp": 43095711,
      "height": 2147483647
     },
     "unsignedTransactionBytes": "01139f9691023c0010f09c34f225d425306e5be55a4946908156072a...",
     "broadcasted": true,
     "requestProcessingTime": 99,
     "transactionBytes": "01139f9691023c0010f09c34f225d425306e5be55a4946908156072afbead4d5...",
     "fullHash": "ae7b714aadfe5efee0d690dfffac533dbf190827ed2f972d3dd552f4399ed648",
     "transaction": "18329367553676245934"
    }

<small>*Verified 7-Apr-15*</small>

### Create Poll

**Request:**

    http://localhost:8125/burst?
      requestType=createPoll&
      name=Gender Poll&
      description=What is your gender?&
      minNumberOfOptions=1&
      maxNumberOfOptions=1&
      minRangeValue=0&
      maxRangeValue=1&
      minBalance=100000000000&
      minBalanceModel=1&
      option00=Male&
      option01=Female&
      secretPhrase=secretPhrase&
      feeNQT=1000000000&
      deadline=60

**Response:**

    {
     "signatureHash": "c55ceb631d9d9e92b99a326e0323f24d80b7859a28ca43c82e1a03a6b05f02fe",
     "transactionJSON": {
      "senderPublicKey": "10f09c34f225d425306e5be55a4946908156072afbead4d574a512d7e086ef5c",
      "signature": "ffcd924f00e14c9e2cc14a04984f0d0c21bdf4551c3e332e1559c378b2975b0b9f8...",
      "feeNQT": "1000000000",
      "type": 1,
      "fullHash": "fc9e0c8014853b7e48b16e2c51ca28c4cc74939ba193e8c040c97e36a42b89c5",
      "version": 1,
      "phased": false,
      "ecBlockId": "11503767830531050607",
      "signatureHash": "c55ceb631d9d9e92b99a326e0323f24d80b7859a28ca43c82e1a03a6b05f02fe",
      "attachment": {
       "minRangeValue": 0,
       "votingModel": 0,
       "description": "What is your gender?",
       "minNumberOfOptions": 1,
       "holding": "0",
       "minBalance": "100000000000",
       "name": "Gender Poll",
       "finishHeight": 256453,
       "options": [
        "Male",
        "Female"
       ],
       "maxNumberOfOptions": 1,
       "minBalanceModel": 1,
       "version.PollCreation": 1,
       "maxRangeValue": 1
      },
      "senderRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
      "subtype": 2,
      "amountNQT": "0",
      "sender": "15295723609781267838",
      "ecBlockHeight": 256410,
      "deadline": 60,
      "transaction": "9096010195498999548",
      "timestamp": 43114026,
      "height": 2147483647
     },
     "unsignedTransactionBytes": "01122ade91023c0010f09c34f225d425306e5be55a4946908156072afb...",
     "broadcasted": true,
     "requestProcessingTime": 6,
     "transactionBytes": "01122ade91023c0010f09c34f225d425306e5be55a4946908156072afbead4d574...",
     "fullHash": "fc9e0c8014853b7e48b16e2c51ca28c4cc74939ba193e8c040c97e36a42b89c5",
     "transaction": "9096010195498999548"
    }

<small>*Verified 7-Apr-15*</small>

### Get Poll

**Request:**

    http://localhost:8125/burst?
      requestType=getPoll&
      poll=9096010195498999548

**Response:**

    {
     "minRangeValue": 0,
     "votingModel": 0,
     "description": "What is your gender?",
     "finished": true,
     "poll": "9096010195498999548",
     "requestProcessingTime": 1,
     "minNumberOfOptions": 1,
     "minBalance": "100000000000",
     "accountRS": "BURST-4VDY-LNVT-LMAY-FMCKA",
     "name": "Gender Poll",
     "options": [
      "Male",
      "Female"
     ],
     "finishHeight": 256453,
     "maxNumberOfOptions": 1,
     "minBalanceModel": 1,
     "account": "15295723609781267838",
     "maxRangeValue": 1
    }

<small>*Verified 8-Apr-15*</small>

### Get Poll Result

**Request:**

    http://localhost:8125/burst?
      requestType=getPollResult&
      poll=16742897359122764363

**Response:**

    {
     "votingModel": 0,
     "minBalance": "0",
     "options": [
      "Option 00",
      "Option 01",
      "Option 02"
     ],
     "finished": true,
     "poll": "16742897359122764363",
     "requestProcessingTime": 2,
     "minBalanceModel": 0,
     "results": [
      {
       "result": "3",
       "weight": "1"
      },
      {
       "result": "9",
       "weight": "2"
      },
      {
       "result": "9",
       "weight": "2"
      }
     ]
    }

<small>*Verified 14-Apr-15*</small>

### Get Poll Vote

**Request:**

    http://localhost:8125/burst?
      requestType=getPollVote&
      account=9096010195498999548

**Response:**

    {
     "voterRS": "BURST-THLJ-CYAL-JQST-6FNS5",
     "votes": [
      "1",
      ""
     ],
     "voter": "4747512364439223888",
     "requestProcessingTime": 1,
     "transaction": "6214591232702166122"
    }

<small>*Verified 8-Apr-15*</small>

### Get Poll Votes

**Request:**

    http://localhost:8125/burst?
      requestType=getPollVotes&
      poll=9096010195498999548&
      lastIndex=0

**Response:**

    {
     "votes": [
      {
       "voterRS": "BURST-THLJ-CYAL-JQST-6FNS5",
       "votes": [
        "1",
        ""
       ],
       "voter": "4747512364439223888",
       "transaction": "6214591232702166122"
      }
     ],
     "requestProcessingTime": 1
    }

<small>*Verified 8-Apr-15*</small>

### Get Polls

**Request:**

    http://localhost:8125/burst?
      requestType=getPolls&
      lastIndex=0

**Response:**

    {
     "polls": [
      {
       "minRangeValue": 0,
       "votingModel": 3,
       "description": "Should we start fully backing Credits with something tangible?",
       "finished": false,
       "poll": "10307833923604182368",
       "minNumberOfOptions": 1,
       "holding": "415923435949896799",
       "minBalance": "1000000000",
       "accountRS": "BURST-3BCV-8Q5G-9NP6-576DP",
       "name": "CRDTS owners poll",
       "options": [
        "Yes",
        "No",
        "go for fractional reserve backing"
       ],
       "finishHeight": 260218,
       "maxNumberOfOptions": 1,
       "minBalanceModel": 3,
       "account": "4227900615136683355",
       "maxRangeValue": 1
      }
     ],
     "requestProcessingTime": 0
    }

<small>*Verified 8-Apr-15*</small>

### Search Polls

**Request:**

    http://localhost:8125/burst?
      requestType=searchPolls&
      query=hello&
      includeFinished=true

**Response:**

    {
     "polls": [
      {
       "minRangeValue": 0,
       "votingModel": 1,
       "description": "Hello World Poll",
       "finished": true,
       "poll": "3043888702585899321",
       "minNumberOfOptions": 1,
       "minBalance": "0",
       "accountRS": "BURST-XK4R-7VJU-6EQG-7R335",
       "name": "Poll1",
       "options": [
        "Hello",
        "World"
       ],
       "finishHeight": 255816,
       "maxNumberOfOptions": 1,
       "minBalanceModel": 1,
       "account": "5873880488492319831",
       "maxRangeValue": 1
      }
     ],
     "requestProcessingTime": 3
    }

<small>*Verified 9-Apr-15*</small>

Utilities
---------

### Decode QR Code

**Request:**

    http://localhost:8125/burst?
      requestType=decodeQRCode&
      qrCodeBase64=/9j/4AAQSkZJRgABAgAAAQABAAD/2wBDAAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UHRo...

**Response:**

    {
        "qrCodeData": "BURST-7A48-47JL-T7LD-D5FS3",
        "requestProcessingTime": 15
    }

<small>*Verified 9-Jun-16*</small>

### Detect Mime Type

**Request:**

    http://localhost:8125/burst?
      requestType=detectMimeType&
      data=/9j/4AAQSkZJRgABAgAAAQABAAD/2wBDAAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UHRo...&
      isText=false

**Response:** {

“`requestProcessingTime`”`: 348,`
“`type`”`: `“`image/png`”

}

</pre>
<small>*Verified 17-Jun-16*</small>

### Encode QR Code

**Request:**

    http://localhost:8125/burst?
      requestType=encodeQRCode&
      qrCodeData=BURST-7A48-47JL-T7LD-D5FS3&
      width=100&
      height=100

**Response:**

    {
        "qrCodeBase64": "/9j/4AAQSkZJRgABAgAAAQABAAD/2wBDAAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UH...",
        "requestProcessingTime": 15
    }

<small>*Verified 9-Jun-16*</small>

### Full Hash To Id

**Request:**

    http://localhost:8125/burst?
      requestType=fullHashToId&
      fullHash=c34af8f1509e3be79c4562e24125ff2a8f026871fdd1a0366ad315bf8fab76b9

**Response:**

    {
     "stringId": "16662085316881435331",
     "requestProcessingTime": 0,
     "longId": "-1784658756828116285"
    }

<small>*Verified 15-Aug-15*</small>

### Hash

**Request:**

    http://localhost:8125/burst?
      requestType=hash&
      hashAlgorithm=2&
      secret=74657374

**Note:** *74657374* is the UTF-8 hex string for the secret phrase “test”.

**Response:**

    {
     "requestProcessingTime": 1,
     "hash": "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
    }

<small>*Verified 14-Jul-15*</small>

### Hex Convert

**Request:**

    http://localhost:8125/burst?
      requestType=hexConvert&
      string=616263

**Response:**

    {
     "binary": "363136323633",
     "text": "abc",
     "requestProcessingTime": 1
    }

<small>*Verified 19-May-15*</small>

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

Debug Operations
----------------

### Clear Unconfirmed Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=clearUnconfirmedTransactions

**Response:**

    {
     "requestProcessingTime": 814,
     "done": true
    }

<small>Verified 12-Nov-14</small>

### Dump Peers

**Request:**

    http://localhost:8125/burst?
      requestType=dumpPeers&
      version=1.5.11&
      connect=true

**Response:**

    {
     "peers": "198.105.122.160; 174.140.167.239; ...",
     "count": 37,
     "requestProcessingTime": 176897
    }

<small>Verified 9-Jun-15</small>

### Full Reset

**Request:**

    http://localhost:8125/burst?
      requestType=fullReset

**Response:**

    {
     "requestProcessingTime": 4378,
     "done": true
    }

<small>Verified 12-Nov-14</small>

### Get All Broadcasted Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=GetAllBroadcastedTransactions

**Response:**

    {
     "requestProcessingTime": 1,
     "transactions": []
    }

<small>Verified 17-May-15</small>

### Get All Waiting Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=getAllWaitingTransactions

**Response:**

    {
     "requestProcessingTime": 1,
     "transactions": []
    }

<small>Verified 17-May-15</small>

### Get Log

**Request:**

    http://localhost:8125/burst?
      requestType=getLog&
      count=3

**Response:**

    {
     "messages": [
      "2015-04-28 18:26:52 FINE: Known peers: 29\n",
      "2015-04-28 18:27:12 FINE: Got 0 confirmations\n",
      "2015-04-28 18:27:12 FINE: Downloaded 230 blocks\n"
     ],
     "requestProcessingTime": 0
    }

<small>Verified 28-Apr-15</small>

### Get Stack Traces

**Request:**

    http://localhost:8125/burst?
      requestType=getStackTrace

**Response:**

    {
     "threads": [
      ⋮
      {
       "trace": [
        "java.lang.Object.wait(Native Method)",
        "java.lang.Object.wait(Object.java:502)",
        "java.lang.ref.Reference$ReferenceHandler.run(Reference.java:157)"
       ],
       "name": "Reference Handler",
       "id": 2,
       "state": "WAITING"
      }
     ],
     "requestProcessingTime": 5,
     "locks": []
    }

<small>Verified 28-Apr-15</small>

### Lucene Reindex

**Request:**

    http://localhost:8125/burst?
      requestType=luceneReindex

**Response:**

    {
     "requestProcessingTime": 4480,
     "done": true
    }

<small>Verified 26-Nov-14</small>

### Pop Off

**Request:**

    http://localhost:8125/burst?
      requestType=popOff&
      numBlocks=1

**Response:**

    {
     "blocks": [
      {
       "previousBlockHash": "5dd57ec106d9ba4cb1442017586b9df23c2c31ec5f1cb46fd3206015a96fd057",
       "payloadLength": 0,
       "totalAmountNQT": "0",
       "generationSignature": "e23740a05b65aa01a890a9ba1c9766183a328455b6e347add727a823e151db99",
       "generator": "16120433118765388429",
       "generatorPublicKey": "f35600023d42e79cbf0518a174141bce1decb13332e397553c3f62df7261414a",
       "baseTarget": "268653798",
       "payloadHash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
       "generatorRS": "BURST-A6NF-JE4R-XB6L-FAUFZ",
       "nextBlock": "12812496435932308379",
       "numberOfTransactions": 0,
       "blockSignature": "d9aba136c708e4f84b7ae1cf0415b25a75f9eefc8a47b1963b237ada9982a...",
       "transactions": [],
       "version": 3,
       "totalFeeNQT": "0",
       "previousBlock": "5528970115590051165",
       "block": "8125766982300698657",
       "height": 212194,
       "timestamp": 23052841
      }
     ],
     "requestProcessingTime": 109
    }

<small>Verified 12-Nov-14</small>

### Rebroadcast Unconfirmed Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=RebroadcastUnconfirmedTransactions

**Response:**

    {
     "requestProcessingTime": 2,
     "done": true
    }

<small>Verified 17-May-15</small>

### Requeue Unconfirmed Transactions

**Request:**

    http://localhost:8125/burst?
      requestType=requeueUnconfirmedTransactions

**Response:**

    {
     "requestProcessingTime": 1,
     "done": true
    }

<small>Verified 16-May-15</small>

### Retrieve Pruned Data

**Request:**

    http://localhost:8125/burst?
      requestType=retrievePrunedData

**Response:**

    {
      "numberOfPrunedData": 0,
      "requestProcessingTime": 4604,
      "done": true
    }

<small>*Verified 9-Jun-16*</small>

### Scan

**Request:**

    http://localhost:8125/burst?
      requestType=scan

**Response:**

    {
     "scanTime": 8,
     "requestProcessingTime": 8445,
     "done": true
    }

<small>Verified 12-Nov-14</small>

### Set Logging

**Request:**

    http://localhost:8125/burst?
      requestType=setLogging

**Response:**

    {
     "loggingUpdated": true
    }

<small>Verified 16-May-15</small>

### Shutdown

**Request:**

    http://localhost:8125/burst?
      requestType=shutdown

**Response:**

    {
     "requestProcessingTime": 2,
     "shutdown": true
    }

<small>Verified 28-Apr-15</small>

### Trim Derived Tables

**Request:**

    http://localhost:8125/burst?
      requestType=trimDerivedTables

**Response:**

    {
     "requestProcessingTime": 394,
     "done": true
    }

<small>Verified 12-May-15</small>

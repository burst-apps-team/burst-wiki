| The Burst Blockchain |
|----------------------|
| **Status**           |

Table Of Contents
-----------------

\_\_TOC\_\_

Introduction
------------

The Burst Blockchain is the digital ledger in which transactions (including AT, Escrow, Messages, etc) made in Burstcoin are recorded chronologically and publicly. All these informations are stored on the Blockchain, in a H2 or MariaDB database.

Database Fields
---------------

All the blockchain is stored on these tables :

<img src="Burst_Blockchain_Diagram.png" title="Burst_Blockchain_Diagram.png" alt="Burst_Blockchain_Diagram.png" width="1745" height="1745" />

### 1. Block Table

These 19 fields define a block on the blockchain in the current version of the Burst software.

| Field Name             | Data Type     | NOT NULL |
|------------------------|---------------|----------|
| db\_id                 | BIGINT(20)    | yes      |
| id                     | BIGINT(20)    | yes      |
| version                | INTEGER(11)   | yes      |
| timestamp              | INTEGER(11)   | yes      |
| previous\_block\_id    | BIGINT(20)    | no       |
| total\_amount          | BIGINT(20)    | yes      |
| total\_fee             | BIGINT(20)    | yes      |
| payload\_length        | INTEGER(11)   | yes      |
| generator\_public\_key | VARBINARY(32) | yes      |
| previous\_block\_hash  | VARBINARY(32) | no       |
| cumulative\_difficulty | BLOB          | yes      |
| base\_target           | BIGINT(20)    | yes      |
| next\_block\_id        | BIGINT(20)    | no       |
| height                 | INTEGER(11)   | yes      |
| generation\_signature  | VARBINARY(64) | yes      |
| block\_signature       | VARBINARY(64) | yes      |
| payload\_hash          | VARBINARY(32) | yes      |
| nonce                  | BIGINT(20)    | yes      |
| ats                    | BLOB          | ?        |

**Note:** the block table has evolved since the genesis block.

**Note:** most fields cannot be NULL. The exceptions are the `previous_...` and `next_...` fields which link the blocks into a chain both forward and backward. The genesis block has a NULL `previous_block_id` and the last (current) block has a NULL `next_block_id`.

Below the column list is a list of indexes. The indexes are all used for sorting various columns for fast retrieval, but the following columns are also restricted to have unique values: `db_id`, `height`, `id`, `timestamp`. They are all used to uniquely identify blocks. `db_id` is the autoincrement field of the table. It normally increases by one with each new block but gaps can occur in the sequence due to occasional deleted blocks. `height` is zero for the genesis block and increases by one with each block in the blockchain. There are no gaps in this sequence. `id` is a unique block id derived from the hash of some of the block fields. `timestamp` is the block creation time measured in number of seconds elapsed since the genesis block.

**Note:** blocks stored in the `BLOCK` table are associated with transactions stored in the `transaction` table through the fields `payload_length` and `payload_hash`, and `total_amount` and `total_fee`. `payload_length` is the total number of bytes of certain fields of all transactions associated with the block and `payload_hash` is the hash of all those fields. `total_amount` and `total_fee` are the total amounts and fees of all transactions associated with the block. All four of these block fields are zero when there are no transactions associated with the block.

### 2. TRANSACTION Table

These 25 fields define a transaction the current version of the Burst software. Note that the transaction table has evolved since the genesis block.

| Field Name                          | Data Type     | NOT NULL |
|-------------------------------------|---------------|----------|
| DB\_ID                              | BIGINT(20)    | yes      |
| ID                                  | BIGINT(20)    | yes      |
| DEADLINE                            | SMALLINT(6)   | yes      |
| SENDER\_PUBLIC\_KEY                 | VARBINARY(32) | yes      |
| RECIPIENT\_ID                       | BIGINT(20)    | no       |
| AMOUNT                              | BIGINT(20)    | yes      |
| FEE                                 | BIGINT(20)    | yes      |
| HEIGHT                              | INTEGER(11)   | yes      |
| BLOCK\_ID                           | BIGINT(20)    | yes      |
| SIGNATURE                           | VARBINARY(64) | yes      |
| TIMESTAMP                           | INTEGER(11)   | yes      |
| TYPE                                | TINYINT(4)    | yes      |
| SUBTYPE                             | TINYINT(4)    | yes      |
| SENDER\_ID                          | BIGINT(20)    | yes      |
| BLOCK\_TIMESTAMP                    | INTEGER(11)   | yes      |
| FULL\_HASH                          | VARBINARY(32) | yes      |
| REFERENCED\_TRANSACTION\_FULL\_HASH | VARBINARY(32) | no       |
| ATTACHMENT\_BYTES                   | BLOB          | no       |
| VERSION                             | TINYINT(4)    | yes      |
| HAS\_MESSAGE                        | BOOLEAN(1)    | yes      |
| HAS\_ENCRYPTED\_MESSAGE             | BOOLEAN(1)    | yes      |
| HAS\_PUBLIC\_KEY\_ANNOUNCEMENT      | BOOLEAN(1)    | yes      |
| EC\_BLOCK\_HEIGHT                   | INTEGER(11)   | no       |
| EC\_BLOCK\_ID                       | BIGINT(20)    | no       |
| HAS\_ENCRYPTTOSELF\_MESSAGE         | BOOLEAN(1)    | yes      |



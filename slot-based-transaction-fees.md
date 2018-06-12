With the Pre-Dymaxion Hard Fork Burst changed to a progressive transaction fee system based on slots which is described in the Dymaxion paper.

Motivation
==========

The slot-based transaction fee system allows for variable fees depending on the transaction load on the blockchain. As a secondary goal it prevents spamming blocks with lots of transactions with minimum fees. Such transactions that are only created because they require little to no investment are considered wasteful, because space on the blockchain is a scarce resource. We want to minimize the total blockchain space because it allows for lower cost of maintaining a node.

Slots
=====

A block is divided into 1020 slots that can each hold a transaction. The fee for the slots increases progressively in a linear fashion by \#slot \* 0.00735 BURST. Here 0.00735 is the smallest fee and called fee quant as a unit. All regular payment and multi out transactions pay the same fee for the slots. The fee for the lowest slot is 0.00735 Burst, the highest fee is 7.497 BURST. The total fee for a block where all slots are filled with the minimal required fee is 3827.2185 BURST.

Assignment of Transactions to Slots
===================================

When a new block is generated, new (currently unconfirmed) transactions are assigned to the available slots. Here, each transaction is assigned to the most costly slot it fits in. If no slot is available, the transaction is not included into the block and remains unconfirmed in the mempool of the node. For included transactions, left-over funds for the slot are NOT refunded! This means if a transaction includes a fee of 0.01 BURST fee and is assigned to the first slot which costs 0.00735 BURST, it will still cost 0.01 BURST fee to process the transaction. It is up to the user to choose a reasonable fee that does not waste funds.

Examples
========

Example 1
---------

A (legacy) fee of 1 Burst guarantees inclusion into a block with less than 136 transactions.

Example 2
---------

A fee of 0.1 Burst guarantees inclusion into a block with less than 13 transactions.

Example 3
---------

A more complex example of transaction assignment to slots is shown in the following:

        Fee Quantum: 0.00735
     Block Capacity: 10
    Fee slots:
      10: 0.0735
       9: 0.06615
       8: 0.0588
       7: 0.05145
       6: 0.0441
       5: 0.03675
       4: 0.0294
       3: 0.02205
       2: 0.0147
       1: 0.00735
    Pending Tx fees (descending-sorted and filtered too low):
    $VAR1 = [
              '0.3',
              '0.2',
              '0.1',
              '0.008'
            ];
    Distributing as follows:
    0.3 has slot @ 10 -> added
    0.2 has slot @ 9 -> added
    0.1 has slot @ 8 -> added
    skipped slot @ 7 (0.008 too low for 0.05145).
    skipped slot @ 6 (0.008 too low for 0.0441).
    skipped slot @ 5 (0.008 too low for 0.03675).
    skipped slot @ 4 (0.008 too low for 0.0294).
    skipped slot @ 3 (0.008 too low for 0.02205).
    skipped slot @ 2 (0.008 too low for 0.0147).
    0.008 has slot @ 1 -> added

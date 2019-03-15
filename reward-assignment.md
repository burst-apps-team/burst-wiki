| Reward Assignment |
|-------------------|
| **Status**        |
||

\_\_TOC\_\_

Introduction
------------

Reward assignment is used to facilitate pool mining. A reward assignment notifies the network that all block rewards attributed to your plot files while mining are to be given to another account (the pool account) which is acting in your place. It grants permission for blocks that are forged using the deadlines that you submit to a pool to be signed with the account belonging to that pool.

Reward Assignment on QBundle
----------------------------

Under “Tools/Reward Assignment”, you will find the reward recipient form. Complete the missing fields and press “Set Reward Recipient” button. This form also allows you to check you current reward recipient.

<img src="Set_Reward_Recipient.png" title="Set_Reward_Recipient.png" alt="Set_Reward_Recipient.png" width="400" height="400" />

Reward Assignment on Burst wallet
---------------------------------

On the wallet, click on “Reward Assignment”. Enter the [RS Address](rs-address-format.md) of the recipient and click on “Set Reward Recipient”.

<img src="Set_Reward_Recipient_BRS.png" title="Set_Reward_Recipient_BRS.png" alt="Set_Reward_Recipient_BRS.png" width="400" height="257" />

Reward Assignment on Burst wallet using reward assignment page
--------------------------------------------------------------

In [`http://127.0.0.1:8125/rewardassignment.html`](http://127.0.0.1:8125/rewardassignment.html), follow the instructions and click on “submit”.

Reward Assignment using Burst API
---------------------------------

You also have the possibility to set the reward recipient directly [using the API](the-burst-api-set-reward-recipient.md).

Bug Fix of BRS
--------------

In previous versions of BRS (2.27 and earlier), the reward recipient assignments would not go into unconfirmed transactions. This has been fixed as of BRS version 2.3.0[1], so if you are trying to join a pool, or change to solo mining (for example)... make sure to be logged into a wallet that is running latest. Otherwise you will face issues changing it.

References
----------

<references />
\_\_FORCETOC\_\_

[1] <https://github.com/burst-apps-team/burstcoin/releases/tag/v2.3.0>

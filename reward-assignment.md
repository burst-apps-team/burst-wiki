| Reward Assignment |
|-------------------|
| **Status**        |
||

Table Of Contents
=================

\_\_TOC\_\_

Introduction
============

Reward assignment is frequently used when pool mining. When changing your reward assignment, you tell the network that another account (the pool account) is acting in your place for 2 specific features. The first feature is that all block rewards that should be given to your account will now be given to the pool account instead. Secondly, for the pool to be able to utilize the deadlines found from your plot files, it is also granted the action to sign the newly forged blocks with the account belonging to the pool.

Reward Assignment on QBundle
============================

Under “Tools/Reward Assignment”, you will find the reward recipient form. Complete the missing fields and press “Set Reward Recipient” button. This form also allows you to check you current reward recipient.

<img src="Set_Reward_Recipient.png" title="Set_Reward_Recipient.png" alt="Set_Reward_Recipient.png" width="400" height="400" />

Reward Assignment on Burst wallet
=================================

On the wallet, click on “Reward Assignment”. Enter the [RS Address](rs-address-format.md) of the recipient and click on “Set Reward Recipient”.

<img src="Set_Reward_Recipient_BRS.png" title="Set_Reward_Recipient_BRS.png" alt="Set_Reward_Recipient_BRS.png" width="400" height="257" />

Reward Assignment on Burst wallet using reward assignment page
==============================================================

In [`http://127.0.0.1:8125/rewardassignment.html`](http://127.0.0.1:8125/rewardassignment.html), follow the instructions and click on “submit”.

Reward Assignment using Burst API
=================================

You also have the possibility to set the reward recipient directly using the API using [this doc](the-burst-api-set-rewardrecipient.md).

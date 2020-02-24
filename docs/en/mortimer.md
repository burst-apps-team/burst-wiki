Mortimer the Burst Butler
=========================

Introduction
------

Originally started as the PoCC Bot for tips on Reddit, Mortimer is a multi-interface, multi-currency and multi-functionality service framework mainly for the Burstcoin cryptocurrency. Called “butler” for short. Mortimer supports other cryptocurrencies than Burstcoin where it serves the purpose of added value to Burstcoin (such as exchange services).

Tip Bot
-------

Mortimer can currently act as a tip bot on Reddit. In order to use it for tipping services, you have to associate your Reddit account with a Burst address and fund it with some Burst. In order to be able to receive tips, you need to register a Burst address with your reddit account.

The command to register your Reddit account with a Burst address is:

    /u/PoCCBot setaddr <your BURST address>

In order to be able to **send** some tips, you have to send some Burst amount to the PoCC Bot:

<https://explore.burst.cryptoguru.org/transaction/8144579246955142410>

The funding has to happen by sending an Ordinary Payment (obviously Multi-Outs will not work as you cannot attach any message there) with a **unencrypted** Message in the format

    /u/PoCCBot fund <reddit_account>

Once said transaction reaches the bot, it will add the Burst you sent to your “Reddit tips account” Please be careful when sending commands to the bot. At the moment there are no mechanisms for refunds, so if you get the reddit account wrong, or mess up the message format - your funds may be lost. Making Mortimer more robust is on the TODO list though.

As you can see the reddit account funded is defined in the message. It does not necessarily need to be “your” account. You can also sponsor some other account. Which will probably be used in future to fund some “agents” who can (ab)use tips as faucet.

The service is provided for free, i.e. the bot does not take any percentage of the funds.

Mixing Service
--------------

Mortimer includes an anonymization/mixing service. In order to use the mixing service, you send an ordinary payment to the *PoCC Churner (mixer)* address (https://explore.burst.cryptoguru.org/transaction/2946354633558760665), with the amount you would like to have mixed. Mortimer will take 10 Burst fixed fee (for tx fees) and 2% “of the rest”. If you send less than 10 Burst ... `¯\_(ツ)_/¯` - obviously.

You add an **encrypted message** to that transaction with the commands for that bot. If that message is not encrypted, or if there is no message, the mixing will not happen, your funds are forfeited for now. Work on making Mortimer more robust is underway, so it is considered to simply send the funds back to the originating address - sans some fee - on errors.

So you sent a tx with an encrypted message. What needs to be in that encrypted message? For now, simply

    OUT: <adr1>, <adr2>, <adr3>, ...

so basically “OUT:” followed by a comma/space separated list of target addresses where the remaining funds (sent funds - fees) will land. e.g.

    OUT: RTST-6QF7-XV5Y-H8DFW, RTST-UFV8-N6DD-794Y3 BURST-RTST-M4HW-82N6-945YG

as you can see, the Mortimer parser is somewhat tolerant, you can use comma, or space, or comma and space and you can have addresses with or without the BURST- prefix. In order for the mixer to be effective, you obviously should take care of the following things:

-   never tell anyone the addresses in that encrypted message
-   make sure they are empty/new i.e. without any transaction history
-   they have NO PUBKEY associated with them, because the mixing addresses have neither
-   that you have time. Even after the funds have arrived on your target addresses, do not withdraw them immediately


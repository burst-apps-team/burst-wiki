Introduction
------------

In addition to a plethora of [Burst Software](burst-software.md), the Burst ecosystem also offers a lot of services to use without - or very little - software installation of yours.

Online Wallets
==============

Every BRS - the [Burst Reference Software](burst-software-burst-reference-software--28brs-29.md) - has a builtin web server (based on [Jetty](https://www.eclipse.org/jetty/documentation/9.4.x/contributing-documentation.html)) to deliver it's front end content to the user. This can happen locally (when you have a local wallet installed either via Qbundle or standalone) and of course this also allows the wallet to be accessible from anywhere in the internet - if you made the API accessible on a public IP.

As a user, you should always be aware of the security implications using an online wallet. Essentially you may be giving away your password to this wallet, so either use only 100% trusted public wallets, or use just a subset of their functionality - such as address watch only (when you enter the BURST-xxx address instead of a password)

A few well known public wallets are:

-   [VoIPLanParty](https://voiplanparty.com:8125/index.html)
-   Burst-Team Online Wallets (operated by Haitch): [1](https://wallet1.burst-team.us:2083/index.html), [2](https://wallet2.burst-team.us:2083/index.html) [3](https://wallet3.burst-team.us:2083/index.html), [4](https://wallet4.burst-team.us:2083/index.html), [5](https://wallet5.burst-team.us:2083/index.html)
-   Burstforum: [1](https://wallet1.burstforum.net:2083/index.html), [2](https://wallet2.burstforum.net:2083/index.html), [3](https://wallet3.burstforum.net:2083/index.html)

Explorers
=========

Blockchain explorers provide insight into various events happening in the blockchain, such as transactions, block forging, assets overview and others. Some explorers have also an extensive charts section to provide even more visual feedback about blockchain activities, often putting them in historical context.

The main Burst blockchain explorers are:

-   [BAT Explorer](https://explorer.burstcoin.network/)
-   [PoCC Explorer](https://explore.burst.cryptoguru.org/)
-   [Burstcoin.RO Explorer](https://explore.burstcoin.ro/)
-   [BurstNeon Monitor](http://burstneon.com/monitor?id=16020314477710380875)

Observers/Analyzers
===================

Some services provide deeper or different insight into activities within the Burst network, such as monitoring mining-related activity, alternate views on network consensus integrity etc.

Notable services in this area are:

-   [BAT](https://explorer.burstcoin.network/?action=network_status)
-   [Pool Spy](https://starburst.pink/poolspy/)
-   [Burstcoin.CC](http://burstcoin.cc/) by Luxe
-   Tx Notification Service: [Burst Alerts](http://burstalerts.com/)
-   [Pool uptime monitor](https://uptime.statuscake.com/?TestID=M30iNz7TSq)

Mortimer, the Burstcoin Butler
==============================

On December, 6th 2018 the PoCC announced the renaming of the “PoCC Bot” to Mortimer as well its functionality. Originally started as a Tip bot, Mortimer will provide a lot of services within the Burst ecosystem:

-   tip bot (Reddit, Twitter, Discord, Discourse forum software)
-   pricing information (coinmarketcap API)
-   supporting BURST and NXT (not all functionality), maybe others too
-   mixing service
-   exchange
-   bank

and others. See the dedicated page on [Mortimer](mortimer.md).

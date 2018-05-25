500k “Pre-Dymaxion” Hard Fork Checklist
---------------------------------------

Disclaimer: A checklist is used by qualified personnel. I.e. pre-flight checklist by pilots, medical guideline by medical staff etc. A checklist **is not** a tutorial. If you do not understand some of the points that apply to your “user class”, you have more learning/reading to do.

It is assumed you know what PoC1 and PoC2 is, how they differ, what optimized plots are, where you find various Burst software, how to update it, which pools support PoC2 etc. If you do not know this, you will not find the information here.

### What To Do \*\*Now\*\*? (Everyone)

-   Make sure you have a current version of the BRS 2.0.x wallet, probably 2.0.4
-   You do **not** want to have a BRS 2.1.x version, except you are a developer and/or are testing on TestNet
-   Make sure you shut down all instances of pre-2.0 wallets, especially 1.2.9 or 1.3.6cg. These are **obsolete**.

### What To Do When BRS 2.2.0 is released? (Everyone)

-   Update to BRS 2.2.0 - obviously. Easy update path if you are already running BRS 2.0.4.
-   Never panic. There is nothing in the upcoming HF that would make you lose your funds. Worst case if you use the wrong wallet: Your funds may be temporarily inacessible.

### For Miners

-   Make sure all your PoC1 plots are optimized
-   Get a PoC1 to PoC2 converter
-   Newest jminer is allegedly capable of mining PoC2 files before the PoC2 switch and also PoC1 files after the PoC2 switch. This is probably the best software to avoid **any** mining downtime.
-   However, PoC2 plots **before** the switch will give you only 50% read speed (with jminer, nonfunctional with anything else), so will PoC1 plots **after** the switch. You want to find a good strategy when to convert which of your plots for maximum performance.

### For Hodlers/Investors

-   The coins in your private wallet are safe. Even if you do not update the wallet (which you should), your coins are not gone. Worst case: you may lose only temporarily access to your coins until you get a new wallet version.
-   While we cannot speak for exchanges in general, the coins you have on any exchange are also safe. Even if the exchange messes up their wallet, it's again only a temporal glitch.
-   As long as no private key is lost, no funds are lost.
-   We will make no recommendation if you should sell before the HF and buy afterwards. PoCC members will simply hodl.

### For Developers

-   Wallet API remains the same, there are two new API calls: `sendMoneyMulti` and `sendMoneyMultiSame` which you might want to use for your services - if applicable. They are especially useful for Pools, Exchanges and Faucets.
-   The number of possible transactions per block is now up to 1020 (old: 255), maximum block payload size 179520 bytes (old: 44880)
-   Each of these 1020 slots in a block has a minimum transaction fee of `slot-index * 0.00735` Burst.

### For Bystanders

-   Lean back and enjoy the show!


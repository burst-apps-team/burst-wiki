500k “Pre-Dymaxion” Hard Fork Checklist
---------------------------------------

Disclaimer: A checklist is used by qualified personnel. I.e. pre-flight checklist by pilots, medical guideline by medical staff etc. A checklist **is not** a tutorial. If you do not understand some of the points that apply to your “user class”, you have more learning/reading to do.

It is assumed you know what PoC1 and PoC2 is, how they differ, what optimized plots are, where you find various Burst software, how to update it, which pools support PoC2 etc. If you do not know this, you will not find the information here.

### What To Do \*\*Now\*\* that BRS 2.2.x is released? (Everyone)

-   Make sure you have a current version of the BRS 2.2.x wallet, probably 2.2.1
-   Make sure you shut down all instances of pre-2.0 wallets, especially 1.2.9 or 1.3.6cg. These are **obsolete**.

<!-- -->

-   Never panic. There is nothing in the upcoming HF that would make you lose your funds. Worst case if you use the wrong wallet: your funds may be temporarily inaccessible.
-   You do **not** want to have a BRS 2.1.x version, unless you are a developer and/or are testing on TestNet

### For Miners

-   Make sure all your PoC1 plots are optimized
-   Get a PoC1 to PoC2 converter
-   Check the \[Burst\_Software\] page for miners capable of mining PoC2 files before the PoC2 switch and also PoC1 files after the PoC2 switch. This is probably the best software to avoid **any** mining downtime.
-   However, PoC2 plots **before** the switch will give you only 50% read speed (with jminer, nonfunctional with anything else), so will PoC1 plots **after** the switch. You want to find a good strategy when to convert which of your plots for maximum performance.
-   read below the “PoC2 Switch @ 502k”

### For Hodlers/Investors

-   The coins in your private wallet are safe. Even if you do not update the wallet (which you should), your coins are not gone. Worst case: you may lose only temporarily access to your coins until you get a new wallet version.
-   While we cannot speak for exchanges in general, the coins you have on any exchange are also safe. Even if the exchange messes up their wallet, it's again only a temporal glitch.
-   As long as no private key is lost, no funds are lost.
-   We will make no recommendation if you should sell before the HF and buy afterwards. PoCC members will simply hodl.

### For Pool Operators

-   Make sure your pool and your back-end wallet(s) can switch to PoC2 mining at block height 502 000. (The BRS 2.2.0 wallet can do this, as can the PoCC pool software)
-   You might want to have a look at Multi-Out payments, because they are ideal to perform payouts to your miners after the 500 000 hard fork.

### For Exchanges

-   Update to BRS 2.0.4 **now** - if you use anything older you risk complications for your users as well as security issues.
-   Upgrade to BRS 2.2.0 when released, the upgrade path from 2.0.4 to 2.2.0 should be a fairly easy one

### For Developers

-   Wallet API remains the same, there are two new API calls: `sendMoneyMulti` and `sendMoneyMultiSame` which you might want to use for your services - if applicable. They are especially useful for Pools, Exchanges and Faucets.
-   The number of possible transactions per block is now up to 1020 (old: 255), maximum block payload size 179520 bytes (old: 44880)
-   Each of these 1020 slots in a block has a minimum transaction fee of `slot-index * 0.00735` Burst.

### For Bystanders

-   Lean back and enjoy the show!

For Miners: PoC2 Switch @ 502k
------------------------------

The release of BRS 2.2.0 will contain everything needed for enabling all of the Pre-Dymaxion hard fork features:

-   MultiOut/MultiOutSame transactions
-   4-fold block capacity increase
-   dynamic fees
-   PoC2

However, while the first 3 of these features will be activated at block height 500k, PoC2 will be activated a little later at 502k (about 5.5 days later).

### Why a delayed PoC2 activation?

tl;dr: We promised there will be no two coins, only one Burst. This is why.

Everything that is needed for PoC2 is in place. Miners do have a multitude of options to choose from how to handle the PoC1 -&gt; PoC2 switch.

Nevertheless, we assume a - temporary - drop in mining capacity as soon as the PoC1 -&gt; PoC2 switch happens. Some miners may not be aware of PoC2 at all (hopefully a small percentage), some may be amidst PoC1-&gt;PoC2 conversion and their capacity is not available, some will do on-the-fly conversion and their read times will be higher, therefore their effective capacity smaller... etc.

Bottom line is that shortly after the PoC1 -&gt; PoC2 switch the desired behavior (mine via PoC2) will have to compete with undesired behavior (inert “Wut? I did nothing” miners mining PoC1).

If we had all Pre-Dymaxion HF features switched at once, we would jeopardize the whole HF because with that PoC1 -&gt; PoC2 switch, we would have weakened “the right chain” mining capacity.

Instead, we will enable all the new features regarding higher capacity and dynamic fees and 2000 blocks later, when there are new blocks in the blockchain (sub-1 Burst fee, more than 255tx per block - which we will make sure), and when we are beyond the maximum wallet rollback capability of 1440 blocks, the wallet/pools/miners will automatically enable the PoC2 switch.

At this moment, it is definitely impossible for any old wallet (1.2.9, 1.3.6cg, 2.0.4 or whatever) that might have survived somehow until then to have both PoC1 **and** the new features and be on the winning chain. Like 100% impossible.

This process will ensure that there will be only one Burst. The new, more secure, high-capacity, flexible Burst. Our foundation for even bigger things to come.

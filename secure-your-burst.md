| Secure Your Burst |
|-------------------|
| **Status**        |

Table of contents
-----------------

\_\_TOC\_\_

The Importance of Passphrases
-----------------------------

Unlike most other crypto-currencies and traditional tokens of value, a BURST account is secured only by a single passphrase which can be thought of as a very secure password. Cryptocurrency accounts secured in this way are sometimes referred to as “Brain Wallets”. The passphrase is the only identifier needed in order to transact using BURST or to forge BURST blocks. There are no additional required wallet files of any sort.

Unlike traditional web sites and accounts which limit the number of login attempts and do not disclose the authentication algorithm, the open source nature of the BURST client allows an unlimited number of login attempts which can be executed very quickly, at rates reaching billions of login attempts per second. Considering this, and unlike traditional passwords, your passphrase has to be very long and complex. The BURST client account registration process is known to generate very secure passphrases. We recommend that you use it. Alternatively, you can choose a passphrase with at least 35 truly random characters.

If you are using special characters in your passphrase, rather than just numbers and letters, make sure to use the ASCII representation of these characters and not one of their Unicode representations. For example, the quote character " can be represented not only as ASCII code 34 (0x22) but also as various Unicode characters as explained here: <http://www.amp-what.com/unicode/search/quote> . These are considered different characters when used in a passphrase. To avoid confusion, always using the ASCII version of these characters or do not use them at all. For example, MS Word uses the “ Unicode character by default which is different from the " ASCII character. These are not interchangeable. Substituting one for the other when entering a passphrase will generate a new account rather than provide access to the intended account.

Losing your passphrase means losing your funds, there is no way to recover your passphrase.

Mitigating the Risk of Losing your BURST
----------------------------------------

Let's analyze the various risks and how to mitigate them. The main risks are:

-   Losing your passphrase.
-   Letting someone steal your passphrase.
-   Accidentally sending your BURST to an account with an unknown passphrase.

If you are someone with a propensity to always mess up things online and always need customer support, you will have to exercise extra care to not to lose your BURST.

Best Ways to Remember your Passphrase
-------------------------------------

If you lose your passphrase there is no way to recover it. The chance of forgetting your passphrase is much higher than having your password stolen.

The best way not to forget your passphrase is to store it somewhere that is safe, preferably in more than one secure location. If you will be storing your passphrase on a computer or other hardware, it is important to maintain a backup copy of your files in another location as protection against equipment failure.

You have to accept that there are risks no matter where you store your passphrase.

-   If you store it on your hard drive - the drive might crash.
-   If you store it in a password manager - the passwords file might get corrupted or deleted.
-   If you print it on paper - the paper might get burned or lost.
-   If you store it in your brain - you might forget it.

Therefore, by using more than one storage method, you lower the risk.

Keeping Your Passphrase Safe
----------------------------

Eventually you have to use your passphrase on your local computer in order to sign transactions. Before entering your passphrase on a local computer, you need to be certain that the computer is safe from intrusion. This means that you must be certain that your computer has not been compromised with any malicious software that could be logging your keystrokes.

There is no 100% security, but there are best practices:

-   Don't share your passphrase with anyone.
-   Don't store your passphrase unencrypted on a remote node or your local workstation.
-   Always use the official BURST wallet.
-   Take special care when connecting to remote nodes.
-   Do not leave your passphrase printed on paper next to your computer.
-   Split your BURST into several accounts. Use the accounts with smaller balances for daily operations and only access the higher balance accounts when necessary.

Using a password manager that allows you to store multiple passwords encrypted under a single database password can be secure and convenient. A free, open source option is Keepass.

How Secure is your Passphrase?
------------------------------

When creating a BURST wallet it is important to use a minimum of a 12 word passphrase to avoid [Brute Force Attacks](wikipedia-brute-force-attack.md) and [Rainbow Table Attacks](wikipedia-rainbow-table.md).  Your passphrase is your **Private Key** and must be carefully secured.  To a first time user, using a set of predefined publicly available words may seem counter intuitive.

In fact, to put in perspective how many passwords can be generated by a list of 1626 words in a 12 word combination, the number would be 341,543,870,028,173,427,817,970,975,906,355,941,376 or 341 undecillion. That can be broken down into 341 billion billion billion billion. This is euphemistically called a “large number” in mathematics. It is difficult to imagine because of how astronomically large it is. Attempting all of the possible combinations of a 12 word passphrase drawn from a known dictionary (a process known as brute forcing), would, on average, take longer than the universe has existed – billions of billions of years. Just 5 Words would take over 2,000 years. Each additional word increases the difficulty by 1,626 times.

| Number of Words | Possible Passphrase Combinations                    |
|-----------------|-----------------------------------------------------|
| 1               | 1,626                                               |
| 2               | 2,643,876                                           |
| 3               | 4,298,942,376                                       |
| 4               | 6,990,080,303,376                                   |
| 5               | 11,365,870,573,289,400                              |
| 6               | 18,480,905,552,168,500,000                          |
| 7               | 30,049,952,427,826,000,000,000                      |
| 8               | 48,861,222,647,645,100,000,000,000                  |
| 9               | 79,448,348,025,071,000,000,000,000,000              |
| 10              | 129,183,013,888,765,000,000,000,000,000,000         |
| 11              | 210,051,580,583,132,000,000,000,000,000,000,000     |
| 12              | 341,543,870,028,173,000,000,000,000,000,000,000,000 |

In conclusion, your wallet is safe with a 12 word auto generated passphrase. You should be much more worried about viruses and keyloggers. Adding capital letters, numbers, or symbols makes the passphrase exponentially harder to crack (virtually impossible). In its [mobile wallet](https://play.google.com/store/apps/details?id=org.icewave.burstcoinwallet), the PoC Consortium even added new words to the list.

[1]

<references />

[1] <https://www.burstcoin.ist/2017/10/07/is-the-automatically-generated-passphrase-secure/>

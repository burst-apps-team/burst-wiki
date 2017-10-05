<languages/>

<translate>

Description
-----------

The Alias System feature of Burst essentially allows one piece of text to be substituted for another, so that keywords or keyphrases can be used to represent other things – names, telephone numbers, physical addresses, web sites, account numbers, email addresses, product SKU codes... almost anything you can think of.

For example, you could ask Burst to associate “search” with “www.google.com”. Once this is done, all you have to do to get to Google is type “burst:search” into a Burst-capable browser, and it will translate your request in one for “www.google.com”.

Immediate applications are simple: you can create an easy-to-remember alias for your Burst account number, for example. But since the Alias System is open-ended, it can be used to implement a decentralized DNS system, shopping cart applications, and more.

Creating aliases is

1.  A user sends a transaction that states “ThisText = ThatText”
2.  If the alias is to be changed, just send another transaction with a new definition. Only the account that created an alias can change it.

Details
-------

The alias can be any string of latin-character numbers and letters. The address can be anything like:

-   “173.194.112.174” (an IPv4 address)
-   “2001:0db8:11a3:09d7:1f34:8a2e:07a0:765d” (an IPv6 address)
-   “mydomain.com/secretpage.php?parameter=value” (a URI)
-   “johnsmith@matrix.com”
-   “<tel:+44-20-8123-4567>”
-   ...or even “<bitcoin:1BTCorgHwCg6u2YSAWKgS17qUad6kHmtQW>”.

There are 2 main ways to use Burst aliases without having to rely on third-party plugins for your browser:

1.  '''Server-side. ''' A web server analyses the Burst blockchain and replaces “burst-links” with corresponding addresses before sending HTML documents to users.
2.  '''Client-side. ''' A web browser runs javascript code that connects to Burst bootstrapping nodes and replaces “burst-links” with their addresses. This requires to embedding a small script which is executed in an “onload” event. The script will do all the work via CORS, JSON, or other techniques.

Currently, Burst does not allow the transfer of ownership of an alias to another account.

Alias Transfer/Sale
-------------------

Aliases can be [transferred](how-to-createalias.md) for a 1 BURST fee.

Alias can be sold to either specific BURST Accounts or to the general public. To sell an alias, you can set the price to sell for every alias.

How To
------

See our [Alias System How-To](how-to-alias.md) page

FAQ
---

See the [Alias System](faq-the-burst-alias-system.md) section in our main [FAQ](faq.md).

</translate>

<Category:Features>

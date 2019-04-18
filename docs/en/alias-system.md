Description
-----------

The Alias System feature of Burst essentially allows one piece of text to be substituted for another, so that keywords or keyphrases can be used to represent other things – names, telephone numbers, physical addresses, web sites, account numbers, email addresses, product SKU codes... almost anything you can think of.

For example, you could ask Burst to associate “search” with “www.google.com”. Once this is done, all you have to do to get to Google is type “burst:search” into a Burst-capable browser, and it will translate your request in one for “www.google.com”.

Immediate applications are simple: you can create an easy-to-remember alias for your Burst account number, for example. But since the Alias System is open-ended, it can be used to implement a decentralized DNS system, shopping cart applications, and more.

Creating aliases is

1.  A user sends a transaction that states “ThisText = ThatText”. You can use the wallet menu ‘Aliases’ and click on ‘Register Alias’.
2.  If the alias is to be changed, just send another transaction with a new definition. Use the ‘Edit’ button in ‘Aliases’ menu.

Details
-------

The alias can be any string of latin-character numbers and letters. The address can be anything like:

-   “173.194.112.174” (an IPv4 address)
-   “2001:0db8:11a3:09d7:1f34:8a2e:07a0:765d” (an IPv6 address)
-   “example.com/secretpage.php?parameter=value” (a URI)
-   “johnsmith@example.com”
-   “<tel:+44-20-8123-4567>”
-   ...or even “<bitcoin:12dDMfhWq3scNWDsL4ty1Q5skyJj6M4scB>”.

There are 2 main ways to use Burst aliases without having to rely on third-party plugins for your browser:

1.  '''Server-side. ''' A web server analyses the Burst blockchain and replaces “burst-links” with corresponding addresses before sending HTML documents to users.
2.  '''Client-side. ''' A web browser runs javascript code that connects to Burst bootstrapping nodes and replaces “burst-links” with their addresses. This requires to embedding a small script which is executed in an “onload” event. The script will do all the work via CORS, JSON, or other techniques.

Changes can only be executed by the account that originally created the alias.

Alias Transfer/Sale
-------------------

Aliases can be [transferred](-transfer-alias.md) for a 1 BURST fee.

Alias can be sold to either specific BURST Accounts or to the general public. To sell an alias, you can set the price to sell for every alias.

How To
------

Open your Burst client and select the ‘Aliases’ menu.

### Register alias

1.  Click on ‘Register Alias’ in the upper right corner.
2.  Select the ‘Type’: “URI”, “Account” or “Other”.
3.  In ‘Alias’ field, enter the name without blank spaces.
4.  Depending on the selected ‘Type’, the text input field is named as ‘URI’, ‘Account ID’ or ‘Data’. The ‘Data’ field can contain any text with a length of up to 1000 bytes.
5.  Enter your ‘Passphrase’ and click on ‘Register’. It is listed in italics on the page.
6.  After the network has processed your message (usually after a few minutes), it will be listet in normal font.

### Edit alias

1.  Click on ‘Edit’ behind the alias you want to change.
2.  You can modify the ‘Type’ and ‘Data’. You can not rename the ‘Alias’.
3.  Enter your ‘Passphrase’ and click on ‘Update’.

### Transfer alias

1.  Click on the ‘Transfer’ button behind the desired alias.
2.  Enter the ‘Recipient’ account id and optionally a message.
3.  Enter your ‘Passphrase’ and click on ‘Transfer Alias’.

### Sell alias

1.  Click on ‘Sell’ behind the alias.
2.  You can sell it to a specific user or offer it to the Burst network for purchase.
3.  Enter the price and the corresponding data.
4.  Confirm with your ‘Passphrase’ and click on ‘Sell Alias’.

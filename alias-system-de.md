<languages/>

Beschreibung
------------

Dieses Feature ermöglich, beliebige Texte durch Pseudonyme zu ersetzen. Es können Schlüsselwörter oder Begriffe verwendet werden, um beispielsweise einen Namen, eine Telefonnummern, Hausanschrift, Webseite, E-Mail-Adresse oder eine Kontonummer mit einem Alias zu benennen.

Beispielsweise könnten Sie den Alias “Suche” für “<https://www.google.de>” definieren. Anschließend können Sie in einem Burst-fähigen Internetbrowser “burst:search” als Adresse eingeben. Die Anfrage wird dann vom System in “<https://www.google.de>” übersetzen.

Eine häufige Anwendung ist ein einfach zu merkender Alias für die eigene Burst-Kontonummer. Da das Alias-System öffentlich ist, eignet es sich aber zum Beispiel auch zur Implementierung eines dezentralisierten DNS-Systems.

Um Aliase zu erstellen

sendet ein Benutzer die Transaktion “Dieser Text = Anderer Text”

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

<Category:Features>

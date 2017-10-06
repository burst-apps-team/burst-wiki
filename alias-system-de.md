<languages/>

Beschreibung
------------

Dieses Feature ermöglich, beliebige Texte durch Pseudonyme zu ersetzen. Es können Schlüsselwörter oder Begriffe verwendet werden, um beispielsweise einen Namen, eine Telefonnummern, Hausanschrift, Webseite, E-Mail-Adresse oder eine Kontonummer mit einem Alias zu benennen.

Beispielsweise könnten Sie den Alias “Suche” für “<https://www.google.de>” definieren. Anschließend können Sie in einem Burst-fähigen Internetbrowser “burst:search” als Adresse eingeben. Die Anfrage wird dann vom System in “<https://www.google.de>” übersetzen.

Eine häufige Anwendung ist ein einfach zu merkender Alias für die eigene Burst-Kontonummer. Da das Alias-System öffentlich ist, eignet es sich aber zum Beispiel auch zur Implementierung eines dezentralisierten DNS-Systems.

Um Aliase zu erstellen

1.  sendet ein Benutzer die Transaktion “Dieser Text = Anderer Text”.
2.  Soll der Alias später geändert werden, sendet derselbe Benutzer (Burst-Account) eine Transaktion mit einer neuen Definition. Andere Benutzer können keine Änderung veranlassen.

Details
-------

Der Alias kann eine beliebige Anzahl von Buchstaben, Zahlen und Zeichen im Latin-Zeichensatz enthalten. Beispielsweise:

-   “173.194.112.174” (eine IPv4-Adresse)
-   “2001:0db8:11a3:09d7:1f34:8a2e:07a0:765d” (IPv6)
-   “example.com/meine-geheime-seite.php?parameter=wert” (eine URI)
-   “mustermann@example.com” (eine E-Mail-Adresse)
-   “Tel. +44 20 81234567” (Telefonnummer)
-   ...oder auch “<bitcoin:12dDMfhWq3scNWDsL4ty1Q5skyJj6M4scB>”.

Es gibt zwei Methoden, um Burst-Aliase ohne Plugins von Drittanbietern für Ihren Browser zu nutzen:

1.  **Serverseitig**: Ein Webserver analysiert die Burst-Blockchain und ersetzt die Aliase durch ihre Adressen, bevor die Dokumente an den Client gesendet werden.
2.  **Clientseitig**: Im Webbrowser wird ein JavaScript ausgeführt, welches eine Verbindung zum Burst-Bootstrapping-Knoten herstellt. Es wird im Onload-Ereignis ausgeführt und übersetzt die Aliase mittels Techniken, wie CORS oder JSON.

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

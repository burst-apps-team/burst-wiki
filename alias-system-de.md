<languages/>

Beschreibung
------------

Dieses Feature ermöglicht, beliebige Texte durch Pseudonyme zu ersetzen. Es können Schlüsselwörter oder Begriffe verwendet werden, um beispielsweise einen Namen, eine Telefonnummern, Hausanschrift, Webseite, E-Mail-Adresse oder eine Kontonummer mit einem Alias zu benennen.

Beispielsweise könnten Sie den Alias “suche” für “<https://www.google.de>” definieren. Anschließend kann in einem Burst-fähigen Internetbrowser “burst:suche” als Adresse eingeben werden. Die Anfrage wird dann vom System in “<https://www.google.de>” übersetzen.

Eine häufige Anwendung ist ein einfach zu merkender Alias für die eigene Burst-Kontonummer. Da das Alias-System öffentlich ist, eignet es sich aber zum Beispiel auch zur Implementierung eines dezentralisierten DNS-Systems, um Domänen-Namen für IP-Adressen zu hinterlegen.

Um einen Alias zu erstellen

1.  öffnet der Benutzer die Seite ‘Aliasse’ in seiner Wallet. Nach Klick auf den Schalter ‘Alias registrieren’ werden die gewünschten Daten eingegeben.
2.  Soll der Alias später geändert werden, kann derselbe Benutzer (Burst-Account) die Funktion ‘Bearbeiten’ hinter dem gewünschten Eintrag verwendet. Andere Benutzer können keine Änderung veranlassen.

Neue Aliasse erscheinen erst nach der Netzwerkbestätigung (Blockgenerierung) im Wallet-Menü des Erstellers.

Details
-------

Der Alias kann eine beliebige Anzahl von Buchstaben, Zahlen und Zeichen im Latin-Zeichensatz enthalten. Beispielsweise:

-   “173.194.112.174” (eine IPv4-Adresse)
-   “2001:0db8:11a3:09d7:1f34:8a2e:07a0:765d” (IPv6)
-   “example.com/meine-geheime-seite.php?parameter=wert” (eine URI)
-   “mustermann@example.com” (eine E-Mail-Adresse)
-   “Tel. +44 20 81234567” (Telefonnummer)
-   ...oder auch “<bitcoin:12dDMfhWq3scNWDsL4ty1Q5skyJj6M4scB>”.

Es gibt zwei Methoden, um Burst-Aliasse ohne Plugins von Drittanbietern für im Webbrowser zu nutzen:

1.  **Serverseitig**: Der Webserver analysiert die Burst-Blockchain und ersetzt die Aliasse durch ihre Adressen, bevor die Dokumente an den Client gesendet werden.
2.  **Clientseitig**: Im Webbrowser wird ein JavaScript ausgeführt, welches eine Verbindung zum Burst-Bootstrapping-Knoten herstellt. Es wird im Onload-Ereignis ausgeführt und übersetzt die Aliasse mittels Techniken, wie CORS oder JSON.

Änderungen können nur von dem Konto ausgeführt werden, welches den Alias ursprünglich erstellt hat.

Transfer oder Verkauf
---------------------

Aliasse können für eine Gebühr von 1 BURST [übertragen](how-to-createalias.md) werden.

Der Besitzer kann seienen Alias an einzelne Burst-Konten oder an die breite Öffentlichkeit verkaufen. Hierbei legt er den Preis beliebig fest.

**Typischer Verkaufsablauf**

Der Besitzer klickt im ‘Aliasse’-Menü auf den Schalter ‘Verkaufen’ und markiert ‘An irgendjemand verkaufen’. Er legt den Preis fest und bestätigt mit ‘Alias verkaufen’. Wenn später ein anderer Benutzer denselben Alias erstellen möchten, erhält er die [Meldung](-file-de-alias-kaufen-png.md), dass dieser bereits vorhanden ist, aber zum Verkauf steht. Nach Klick auf den Link ‘Kaufen’ kann er den Alias zum angegebenen Preis übernehmen.

How-To
------

Siehe [Alias-System How-To](how-to-alias.md).

FAQ
---

Siehe [Alias System](faq-the-burst-alias-system.md)-Abschnitt in den [FAQ](faq.md).

<Category:Features>

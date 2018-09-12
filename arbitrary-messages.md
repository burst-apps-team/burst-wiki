| Arbitrary Messages |
|--------------------|
| **Status**         |
||

Description
-----------

Transmission of data messages up to 1000 bytes in length from one account to another account.

Details
-------

Arbitrary messages are limited only by length. Any string can be transmitted, using any data structure or form of data encryption. Encoding, decoding, linked messages, data structures, and more can be implemented by any application that uses the system.

The base implementation allows for the transmission of simple, unencrypted text messages between accounts, but since the messages are truly “arbitrary” the range of possible applications is vast. Secure messaging, torrent applications, voting systems, data storage systems and even simple distributed applications have been suggested.

How To
------

[High-Level API calls for implementing arbitrary messages](https://burstwiki.org/wiki/The_Burst_API#Arbitrary_Message_System_Operations)

Open your Burst client and select ‘Messages’.

### Sending messages

1.  Select the **Messages** option from the left-hand menu or click on the messages icon as shown below and click on ‘Send Message’.
    :<img src="Send_Message.png" title="fig:Send_Message.png" alt="Send_Message.png" width="293" height="150" />
2.  In the **Recipient** field, enter the Burst address that you wish to send the message to.
3.  In the **Message** field, you can enter any text with a length of up to 1000 bytes. When sending, you receive a corresponding error message if you enter longer text.
4.  You can select the checkbox ‘Encrypt Message’ or send the data in plain text. The recipient needs his passphrase to read encrypted messages.
5.  In the **Passphrase** field, enter your ‘Passphrase’ and click on **Send Message**. An example, encrypted message is shown below:
    :<img src="Send_Actual_Message1.png" title="fig:Send_Actual_Message1.png" alt="Send_Actual_Message1.png" width="590" height="505" />
6.  When you have sent the message it is listed in italics on the transactions page.
7.  After the network has processed your message (usually after a few minutes), it will be listed in normal font.
    :<img src="Message_transaction.PNG" title="fig:Message_transaction.PNG" alt="Message_transaction.PNG" width="698" height="115" />

### Receiving messages

1.  Open the menu ‘Messages’ in your wallet.
2.  Here you can find all the received messages sorted by the burst accounts.
3.  Below the messages you will find input fields to respond directly.


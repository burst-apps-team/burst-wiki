<languages/>

| Arbitrary Messages |
|--------------------|
| **Status**         |
| **Credits**        |

Description
-----------

Transmission of data messages up to 1000 bytes in length from one account to another account.

Details
-------

Arbitrary messages are limited only by length. Any string can be transmitted, using any data structure or form of data encryption. Encoding, decoding, linked messages, data structures, and more can be implemented by any application that uses the system.

The base implementation allows for the transmission of simple, unencrypted text messages between accounts, but since the messages are truly “arbitrary” the range of possible applications is vast. Secure messaging, torrent applications, voting systems, data storage systems and even simple distributed applications have been suggested.

How To
------

Open your Burst client and login using your passphrase.

### Sending Messages

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

### Receiving Messages

1.  Select the **Messages** option from the left-hand menu or click on the messages icon as shown below and click on ‘Inbox’.
    :<img src="Inbox.png" title="fig:Inbox.png" alt="Inbox.png" width="288" height="150" />
2.  Here you can find all the received messages sorted by the burst accounts. Select one of the senders accounts and the messages sent from this account will be displayed. If the message has been encrypted, you will see a padlock symbol and a prompt to enter your passphrase to decrypt the messages present as shown below:
    :<img src="Inbox_msg1.PNG" title="fig:Inbox_msg1.PNG" alt="Inbox_msg1.PNG" width="717" height="365" />
3.  Selecting one of the messages allows you to enter your passphrase. You can also select the checkbox to remember you pasphrase for decryption.
      
    <img src="Decrypt_msg.PNG" title="fig:Decrypt_msg.PNG" alt="Decrypt_msg.PNG" width="954" height="262" />

4.  Once you enter your passphrase the message content will be shown.
    :<img src="Decrypted_message_content.PNG" title="fig:Decrypted_message_content.PNG" alt="Decrypted_message_content.PNG" width="358" height="108" />
5.  Below the messages you will find input fields to respond directly to the sender.
    :<img src="Reply_to_message_field.PNG" title="fig:Reply_to_message_field.PNG" alt="Reply_to_message_field.PNG" width="467" height="136" />

### Using the Burst API

The Burst API is also available to encrypt, send, decrypt and read messages.

[High-Level API calls for implementing arbitrary messages](https://burstwiki.org/wiki/The_Burst_API#Arbitrary_Message_System_Operations)

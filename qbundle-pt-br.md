<languages/>

| QBundle      |
|--------------|
| **Status**   |
| **Créditos** |

Description
-----------

The easiest choice for beginners on Windows is to download Qbundle, a launcher allowing you to install the Burst local wallet in a few clicks. You can simply follow the wizard walking you through the set up. Qbundle also includes plotting and mining software to start mining easily.

Tabela de Conteúdos
-------------------

\_\_TOC\_\_

Guia inicial de instalação 
---------------------------

Através de <https://www.burst-coin.org/download-wallet>

<img src="Download-wallet.png" title="Download-wallet.png" alt="Download-wallet.png" width="1235" height="1235" />

Select Qbundle, and it will direct you to the github.

We will select the MSI installation in this guide, the other alternative it a self-contained ‘zip’ download that will need extraction before It can be run. This ‘zip’ version is useful if you wish to set up along side an existing installation (Provided you have first stopped the other wallet, and do not run at the same time!) 

<img src="QBundle_GitHub.png" title="QBundle_GitHub.png" alt="QBundle_GitHub.png" width="739" height="739" />

We will be using the 2.0.0 version, but always download the latest release. (note the green ‘*Latest release*’ indicator beside)

<img src="QBundle_MSI.png" title="QBundle_MSI.png" alt="QBundle_MSI.png" width="462" height="462" />

Save the file and run when downloaded.

<img src="Executable_file.png" title="Executable_file.png" alt="Executable_file.png" width="602" height="602" />

Under windows 10, you may have the following alert: 

<img src="Windows_Alert.png" title="Windows_Alert.png" alt="Windows_Alert.png" width="539" height="539" />

Click on ‘More Info’ to enable the download.

<img src="Run_Anyway.png" title="Run_Anyway.png" alt="Run_Anyway.png" width="514" height="514" />

And then ‘Run anyway’

We will then go through the setup wizard, click next.

<img src="QBundle_Setup.png" title="QBundle_Setup.png" alt="QBundle_Setup.png" width="518" height="518" />

Select a directory that is not in use, and then ‘Next’

<img src="QBundle_Setup_2.png" title="QBundle_Setup_2.png" alt="QBundle_Setup_2.png" width="520" height="520" />

Then click ‘Install’

<img src="QBundle_Setup_3.png" title="QBundle_Setup_3.png" alt="QBundle_Setup_3.png" width="509" height="509" />

A warning may come up, click ‘ok’ and the installation will be completed shortly after

<img src="QBundle_Setup_4.png" title="QBundle_Setup_4.png" alt="QBundle_Setup_4.png" width="493" height="493" />

Click ‘Finish’ to complete.

Following this the setup will check for Java. Click on ‘Download missing components’ 

<img src="First_Time.png" title="First_Time.png" alt="First_Time.png" width="682" height="682" />

This will take a while.

<img src="Download_Manager.png" title="Download_Manager.png" alt="Download_Manager.png" width="475" height="475" />

Once completed, click ‘Continue’ 

<img src="First_Time_2.png" title="First_Time_2.png" alt="First_Time_2.png" width="686" height="686" />

The wallet will start but you will need to enable access:

<img src="Windows_Alert_2.png" title="Windows_Alert_2.png" alt="Windows_Alert_2.png" width="551" height="551" />

Enable for ‘Private networks’ and ‘Allow access’

Bootstrap Chain / Import Blockchain
-----------------------------------

You now have a working wallet, which will synch from ‘Genesis’.

This will take some time, so I would advise we ‘bootstrap’ the installation. 

<img src="Bootstrap.png" title="Bootstrap.png" alt="Bootstrap.png" width="465" height="465" />

At the top of the wallet, select ‘Database’ and ‘Bootstrap chain’ 

<img src="Import_Database.png" title="Import_Database.png" alt="Import_Database.png" width="525" height="525" />

Click ‘Start Import’ Installation Upgrade

<img src="Start_Import.png" title="Start_Import.png" alt="Start_Import.png" width="469" height="469" />

A warning will pop up, regarding the truncation of the existing data, click ‘Yes’ 

<img src="Stop_Wallet_Import_Databse.png" title="Stop_Wallet_Import_Databse.png" alt="Stop_Wallet_Import_Databse.png" width="473" height="473" />

Click ‘Yes’ on the next prompt to start.

The bootstrap will then download.

<img src="Download_Manager_2.png" title="Download_Manager_2.png" alt="Download_Manager_2.png" width="485" height="485" />

Once complete, close the bootstrap popup.

<img src="Bootstrap_2.png" title="Bootstrap_2.png" alt="Bootstrap_2.png" width="505" height="505" />

The remaining blocks will then synchronize:

<img src="Block_sync.png" title="Block_sync.png" alt="Block_sync.png" width="658" height="658" />

Please give this up to 30 minutes.

The rate of import will depend upon the computer and network connection.

<img src="Block_sync_2.png" title="Block_sync_2.png" alt="Block_sync_2.png" width="650" height="650" />

When it is synchronised you will see something like the following (close to your current date/time) 

<img src="Block_sync_3.png" title="Block_sync_3.png" alt="Block_sync_3.png" width="521" height="521" />

Congratulations, you now have a fully synchronised wallet, running on the H2 database!

Wallet Mode
-----------

In “File/Mode”, you can either choose Wallet Mode or Launcher Mode. The Wallet Mode lets you to use Burst Wallet within QBundle and launcher mode allows you to access Burst wallet only through your browser.

<img src="Menu_List.png" title="Menu_List.png" alt="Menu_List.png" width="641" height="641" />

This is the launcher mode:

<img src="Launcher_mode_2.png" title="Launcher_mode_2.png" alt="Launcher_mode_2.png" width="586" height="586" />

The wallet
----------

Here's the connection screen:

<img src="Login.png" title="Login.png" alt="Login.png" width="532" height="532" />

When you creating a new account using “New? Create Your Account!”, the wallet will provide you a 12 words automatically generated passphrase that you will need to write on a paper or to memorize.

<img src="Automatically_Generated_Passphrase.png" title="Automatically_Generated_Passphrase.png" alt="Automatically_Generated_Passphrase.png" width="506" height="506" />

Finally, here's the look of the QBundle and Burst Wallet:

<img src="Burst_Wallet.png" title="Burst_Wallet.png" alt="Burst_Wallet.png" width="1919" height="1919" />

Frequently Asked Questions (FAQ)
--------------------------------

### How to update the softwares that come with QBundle ?

All softwares included in QBundle can be updated using ‘Check for updates’.

This installation has an available update, notified by the ‘New update available’ at the top of the wallet. If you click on this, it will list the components that are available for download.

Click ‘Download updates’: 

<img src="Download_Updates.png" title="Download_Updates.png" alt="Download_Updates.png" width="869" height="869" />

Click ‘yes’ to stop the wallet:

<img src="Check_for_updates.png" title="Check_for_updates.png" alt="Check_for_updates.png" width="579" height="579" />

When complete click ‘close’:

<img src="Check_for_updates_2.png" title="Check_for_updates_2.png" alt="Check_for_updates_2.png" width="567" height="567" />

If the wallet is not running, start the wallet by selecting ‘Wallet’ and ‘Start wallet’ 

<img src="Start_Wallet.png" title="fig:Start_Wallet.png" alt="Start_Wallet.png" width="466" height="466" /> 

Now, you have an upgraded QBundle: 

<img src="Download_Updates_2.png" title="Download_Updates_2.png" alt="Download_Updates_2.png" width="723" height="723" />

### How do I see the logs of Burst Reference Software (BRS) ?

#### Console

View Console will provide a pop up window with the Wallet log and MariaDB log (if set up) this is useful for troubleshooting and the output from the console is often asked for. 

<img src="View_console.png" title="View_console.png" alt="View_console.png" width="605" height="605" />

<img src="Console.png" title="Console.png" alt="Console.png" width="766" height="766" />

### How do I change the settings of QBundle ?

#### Settings

If you have OpenCL correctly installed, Using GPU Acceleration will greatly improve performance.

To allow external access to the wallet (so you can use it on the go) add 0.0.0.0 to the ‘Allow API traffic from’ list 

<img src="Local_Wallet_Settings.png" title="Local_Wallet_Settings.png" alt="Local_Wallet_Settings.png" width="858" height="858" />

For troubleshooting the ‘Debug mode’ option can be useful. 

<img src="General_Settings.png" title="General_Settings.png" alt="General_Settings.png" width="860" height="860" />

### I have connection problems. What can I do to improve this ?

#### Configure Firewall

If you click ‘configure firewall’ it will give you an option requesting confirmation that you wish to change the windows firewall to allow access locally and over the WAN to the program on this computer.

<img src="Windows_firewall.png" title="Windows_firewall.png" alt="Windows_firewall.png" width="404" height="404" />

### Where are stored the information about my accounts ?

#### Account Manager Tab

The account manager allows you to manage multiple accounts, and generate short ‘PIN’ numbers for each.

Account Manager is accessible here:

<img src="Login_Account.png" title="Login_Account.png" alt="Login_Account.png" width="613" height="613" />

You can easily manage all you BURST accounts with this software. It provides the Reed-Solomon address and the numeric accound number of you Burst Wallet. You can also see the Public Key.

Finally, your passphrase and your private key are available (you will only need to set up one pin to control many accounts)

<img src="Account_Manager.png" title="Account_Manager.png" alt="Account_Manager.png" width="802" height="802" />

### How to automaticaly startup the wallet ?

#### Windows Service

<img src="Windows_Service.png" title="Windows_Service.png" alt="Windows_Service.png" width="603" height="603" />

Stop wallet will stop the wallet in the background.

Install as Service will allow for the automatic startup of the wallet, even before logon to windows (acting as a service)  this is currently only available when running H2 Database. 

### My wallet is stuck, how do I solve this ?

#### Rollback chain (popoff)

Advanced provides an important option for wallet issues: Rollback chain

<img src="Popoff.png" title="Popoff.png" alt="Popoff.png" width="645" height="645" />

If you experience a wallet start issue, or an incomplete synch, it’s suggested to pop off 1000 block and allow to re-synch. The maximum blocks that should be popped off is 1400, 

<img src="Rollback_Chain.png" title="Rollback_Chain.png" alt="Rollback_Chain.png" width="450" height="450" />

### I don't really like H2 database... How do I change database ?

#### Change database

Changing database allows for the migration of databases, this is an advanced option. 

<img src="Bootstrap.png" title="fig:Bootstrap.png" alt="Bootstrap.png" width="465" height="465" /> 

<img src="Change_Database.png" title="Change_Database.png" alt="Change_Database.png" width="504" height="504" />

If MariaDB is selected, it will install a portable copy, and set up

<img src="MariaDB_DL.png" title="MariaDB_DL.png" alt="MariaDB_DL.png" width="788" height="788" />

This may take a while

<img src="Download_MariaDB.png" title="Download_MariaDB.png" alt="Download_MariaDB.png" width="480" height="480" />

Currently MariaDB will not allow the movement of data from H2 database, but in the future it will be possible to migrate the data within a synchronised database. 

<img src="Transition_Database.png" title="Transition_Database.png" alt="Transition_Database.png" width="485" height="485" />

### What are the other tools available ?

The tool tab provides a lot of tools that could be very helpful.

<img src="Tools.png" title="Tools.png" alt="Tools.png" width="660" height="660" />

#### Plotter

[XPlotter](xplotter.md) by Blago is integrated in QBundle and lets you plot your HDD using your CPU.

<img src="Plotter.png" title="Plotter.png" alt="Plotter.png" width="662" height="662" />

#### Dynamic Plotting

Dynamic plotting is a feature that will add and remove plots depending on free space in the drive.

<img src="Dynamic_Plotting.png" title="Dynamic_Plotting.png" alt="Dynamic_Plotting.png" width="651" height="651" />

#### Miner

[Blagominer](blagominer.md) is integrated in QBundle and lets you mine plots in your HDD using your CPU.

<img src="Miner.png" title="Miner.png" alt="Miner.png" width="671" height="671" />

#### Set Reward Recipient

You can set the reward recipient of your account here. It is commonly used when you plan to mine in a pool and you have to set the pool as the reward recipient of your mining.

<img src="Reward_Recipient.png" title="Reward_Recipient.png" alt="Reward_Recipient.png" width="424" height="424" />

#### Vanity Address Generator

This tool allow you to generate your own Reed-Solomon address.

<img src="Vanity_Address_Generator.png" title="Vanity_Address_Generator.png" alt="Vanity_Address_Generator.png" width="653" height="653" />

### About Tab

You will find the contributors and developer of QBundle in this tab.

<img src="About.png" title="About.png" alt="About.png" width="637" height="637" />

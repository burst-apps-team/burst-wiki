<languages/>

| QBundle     |
|-------------|
| **Status**  |
| **Credits** |

Table Of Contents
-----------------

\_\_TOC\_\_

Description
-----------

The easiest choice for beginners on Windows is to download Qbundle, a launcher allowing you to install the Burst local wallet in a few clicks. You can simply follow the wizard walking you through the set up. Qbundle also includes plotting and mining software to start mining easily.

Initial installation guide 
---------------------------

This page is a guide for downloading and installing QBundle. Images of installation screens and links are included for illustrative purposes only.

The files to download are located here: <https://www.burst-coin.org/download-wallet> .

Select your operating system: (Please note: If your operating system is not Windows, QBundle will not be available and you will need to choose an option specific to your operating system.)

Select Qbundle to be redirected to the PoC Consortium GitHub site.

<img src="Download-wallet.png" title="Download-wallet.png" alt="Download-wallet.png" width="900" height="1235" />

GitHub is a software development platform and hosts the files needed for installation. Future releases will be located here as well. For initial installation, select the latest release, identified by a green indicator. As of the writing of this guide, the latest release is Version 2.5.0. The installation file is QBundle2.5.0\_final.zip.

<img src="Latest_release_2.5.0.png" title="Latest_release_2.5.0.png" alt="Latest_release_2.5.0.png" width="687" height="687" />

Depending on the version of Windows and internet browser that you use (Microsoft Edge, Firefox, etc.), a variation of the following dialog boxes will appear asking what you want to do with the file.

Choose to save the file. If you are given the option to Save and Run, you may do this as well. Make a note of the saved location.

<img src="QBundle_MSI.png" title="QBundle_MSI.png" alt="QBundle_MSI.png" width="900" height="462" />

Navigate to folder containing the saved file. A common location is C: Users/“user name”/Downloads. You may move this file to another location if you wish.

The file will be compressed. You will need to extract the files. Recent versions of Windows include functionality to extract files. If available, click on the “Extract all” button. If you do not have this functionality, you may need a utility to “unzip” the folder.

To launch the setup of QBundle, click on the extracted version of this file: Name: BurstWallet Type: Application

Windows will give you an alert similar to the following: Click on “More Info” to enable the download.

<img src="Windows_Alert.png" title="Windows_Alert.png" alt="Windows_Alert.png" width="509" height="539" />

Click on “Run Anyway”

<img src="Run_Anyway.png" title="Run_Anyway.png" alt="Run_Anyway.png" width="509" height="514" />

The installation utility will check for portable Java and previously installed BRS wallets.

If either of these area are showing in Red, click on “Download missing components”

The check box for “allow connection to remote resources for verifications and updates” should be checked.

<img src="QBundle_Setup.png" title="QBundle_Setup.png" alt="QBundle_Setup.png" width="509" height="518" />

Setup will complete. Upon completion you will see your new wallet showing a status of “stopped”.

<img src="QBundle_Setup_2.png" title="QBundle_Setup_2.png" alt="QBundle_Setup_2.png" width="509" height="520" />

The next step is optional. The BRS wallet in configured to use H2 for its database by default. If you have a technical reason to prefer H2, you can skip this step. However, it is recommended that new users change the database type to “Portable MariaDB”. This database is more powerful and provides superior protection for your data in the event that you shut down your wallet incorrectly.

To change the database type to "MariaDB:

Click on “Database” on the wallet's top menu bar and select “Change Database”.

Select the radio button for Portable MariaDB.

Click “Next”.

You will be asked if you want to download an install MariaDB.

Click “Yes”.

After installation is complete,you will be asked to choose between copying data or proceeding with copying data.

The copy data option will be disabled. Verify that “No Copy” is selected.

Click “Save and close”.

The next step is to download a copy of the Burst blockchain using a process referred to as bootstrapping. This will download the blockchain from known and trusted repository. The bootstrap copy contains the majority of blocks that have been created beginning with the genesis block. When the bootstrapping process in complete, more recent blocks will be added to your copy using the regular peer to peer synchronization process.

To begin the bootstrap process, click on “database” in the top menu bar.

Select “Database”: Select “Bootstrap Chain”.

In “Settings”, verify that "Cryptoguru repository (MariaDB) is selected.

Click “Start Import”

You will received a notice saying that "all existing data in your database will be erased. Do you want to continue?

Click “Yes”.

The next step is to verify you Java settings.

Click on “edit” top menu bar. Select “settings”. The settings window will appear.

Select “Java” in the left menu bar. Verify that the radio button for “Use Portable java” is selected. Save your change or exit.

<img src="QBundle_Setup_3.png" title="QBundle_Setup_3.png" alt="QBundle_Setup_3.png" width="509" height="509" />

A warning may come up, click “OK” and the installation will be completed shortly after

<img src="QBundle_Setup_4.png" title="QBundle_Setup_4.png" alt="QBundle_Setup_4.png" width="493" height="493" />

Click “Finish” to complete.

Following this the setup will check for Java. Click on “Download missing components” 

<img src="First_Time.png" title="First_Time.png" alt="First_Time.png" width="509" height="682" />

This will take a while.

<img src="Download_Manager.png" title="Download_Manager.png" alt="Download_Manager.png" width="475" height="475" />

Once completed, click “Continue” 

<img src="First_Time_2.png" title="First_Time_2.png" alt="First_Time_2.png" width="509" height="686" />

The wallet will start but you will need to enable access:

<img src="Windows_Alert_2.png" title="Windows_Alert_2.png" alt="Windows_Alert_2.png" width="509" height="551" />

Enable for “Private networks” and “Allow access”

Bootstrap Chain / Import Blockchain
-----------------------------------

You now have a working wallet, which will synchhronize from “Genesis”.

This will take some time, so users are advised to “bootstrap” the blockchain. The bootstrap will download a verified chain from one of the built in repositories. The bootstrap significantly reduces the time needed for wallet setup - synchronization of the blockchain from nodes on the network might take 1-3 days, while bootstrapping usually takes up to two hours (depending on system, internet connection speed etc). 

<img src="14_Database_options.png" title="14_Database_options.png" alt="14_Database_options.png" width="465" height="465" />

At the top of the wallet menu, select “Database” and “Bootstrap chain” 

<img src="Import_Database.png" title="Import_Database.png" alt="Import_Database.png" width="509" height="525" />

Click “Start Import” 

<img src="Start_Import.png" title="Start_Import.png" alt="Start_Import.png" width="469" height="469" />

A warning will pop up, regarding the truncation of the existing data, click “Yes” 

<img src="Stop_Wallet_Import_Databse.png" title="Stop_Wallet_Import_Databse.png" alt="Stop_Wallet_Import_Databse.png" width="473" height="473" />

Click “Yes” on the next prompt to start.

The bootstrap will then download.

<img src="Download_Manager_2.png" title="Download_Manager_2.png" alt="Download_Manager_2.png" width="485" height="485" />

Once complete, close the bootstrap popup.

<img src="Bootstrap_2.png" title="Bootstrap_2.png" alt="Bootstrap_2.png" width="505" height="505" />

The remaining blocks will then synchronize:

<img src="Block_sync.png" title="Block_sync.png" alt="Block_sync.png" width="658" height="658" />

Please give this up to 30 minutes.

The rate of import will depend upon the computer and network connection.

<img src="Block_sync_2.png" title="Block_sync_2.png" alt="Block_sync_2.png" width="650" height="650" />

When it is synchronized you will see something like the following (close to your current date/time) 

<img src="Block_sync_3.png" title="Block_sync_3.png" alt="Block_sync_3.png" width="521" height="521" />

Congratulations, you now have a fully synchronised wallet, running on the H2 database!

Wallet Mode
-----------

There are two ways users can run the QBundle, and the “Mode” option allows them to select which of those should be used when QBundle starts. The Wallet Mode lets you to use Burst Wallet within QBundle and launcher mode allows you to access Burst wallet only through your browser.

<img src="Menu_List.png" title="Menu_List.png" alt="Menu_List.png" width="641" height="641" />

The wallet mode will load the complete interface allowing users to issue transactions and see blocks as they are forged on the network, while the “Launcher” mode opens a minimal interface, as shown in the image below.

<img src="4_Launcher_mode.png" title="4_Launcher_mode.png" alt="4_Launcher_mode.png" width="509" height="586" />

The wallet
----------

Here's the connection screen:

<img src="Qbundle_Start_Screen.png" title="Qbundle_Start_Screen.png" alt="Qbundle_Start_Screen.png" width="509" height="532" />

In case the user already has an activated Burst account, they can use their passphrase to log in by clicking the “Returning User?” button.

If, however, the user doesn't have a Burst account, clicking the “New? Create Your Account!” button will allow them to do so. For full guide on how to open and activate a Burst account, read [this](https://burstwiki.org/wiki/Activate_Burst_Account) Burst Wiki article.

<img src="Automatically_Generated_Passphrase.png" title="Automatically_Generated_Passphrase.png" alt="Automatically_Generated_Passphrase.png" width="506" height="506" />

Finally, here's the look of the QBundle and Burst Wallet:

<img src="QBundle_and_Burst_Wallet.png" title="QBundle_and_Burst_Wallet.png" alt="QBundle_and_Burst_Wallet.png" width="900" height="1366" />

Frequently Asked Questions (FAQ)
--------------------------------

### How to update the software that come with QBundle ?

The application manager lists applications delivered with QBundle available for installation, as well as those that are installed on the system, together with version information. As soon as there is an update available for any of the component application, users will be notified of them through the Application Manager window. Not all of the component applications of the QBundle are installed by default, and if users wish, they can add components. Note that currently unavailable applications (i.e. not installed on the system), will be installed once the QBundle is opened for the first time.

In case there's an update available, users will be notified by the “New update available” at the top of the wallet. Once they click the notification, it will list the components that are available for download.

Click “Install/Update” 

<img src="3_Application_manager.png" title="3_Application_manager.png" alt="3_Application_manager.png" width="509" height="776" />

The wallet will be stopped during software update.

If the wallet is not running after completed update or installation, start the wallet by selecting “Wallet” and “Start wallet” 

<img src="Start_wallet.png" title="fig:Start_wallet.png" alt="Start_wallet.png" width="466" height="466" /> 

Now, you have an upgraded QBundle: 

<img src="Updated_QBundle_no_updates.png" title="Updated_QBundle_no_updates.png" alt="Updated_QBundle_no_updates.png" width="509" height="723" />

### How do I see the logs of Burst Reference Software (BRS) ?

#### Console

View Console will provide a pop up window with the Wallet log and MariaDB log (if set up) this is useful for troubleshooting and the output from the console is often asked for. 

<img src="View_console.png" title="View_console.png" alt="View_console.png" width="509" height="605" />

<img src="Console.png" title="Console.png" alt="Console.png" width="509" height="766" />

### How do I change the settings of QBundle ?

#### Settings

Local wallet settings offer an array of options to be set. However, in the general case, users are advised not to change any of the default settings. There is a point to changing the default wallet settings in case the system is lagging while the blockchain is synchronized, in case of which the number of cores should be decreased or in case of troubleshooting.

If the OpenCL correctly installed on the system, using GPU Acceleration will greatly improve performance.

To allow external access to the wallet add 0.0.0.0 to the ‘Allow API traffic from’ list. 

<img src="Local_Wallet_Settings.png" title="Local_Wallet_Settings.png" alt="Local_Wallet_Settings.png" width="509" height="858" />

For troubleshooting the “Debug mode” option can be useful. 

<img src="General_Settings.png" title="General_Settings.png" alt="General_Settings.png" width="509" height="860" />

### I have connection problems. What can I do to improve this ?

#### Configure Firewall

If the “Configure Windows firewall” is clicked, the user will be asked for confirmation that they wish to change the Windows firewall to allow access locally and over the WAN to the program on this computer.

<img src="Windows_firewall.png" title="Windows_firewall.png" alt="Windows_firewall.png" width="404" height="404" />

### Where are stored the information about my accounts ?

#### Account Manager Tab

The account manager allows to manage multiple accounts, and generate short PIN numbers for each.

Account Manager is accessible here:

<img src="11_Account_manager.png" title="11_Account_manager.png" alt="11_Account_manager.png" width="312" height="312" />

Users can easily manage all their BURST accounts with this software. It provides the Reed-Solomon address and the numeric account number of their Burst Wallet. The Public Key for each of those accounts is also shown.

Finally, passphrases and private keys are available (it is enough to set up one PIN to control multiple accounts)

<img src="11a_Account_manager.png" title="11a_Account_manager.png" alt="11a_Account_manager.png" width="509" height="776" />

### My wallet is stuck, how do I solve this ?

#### Rollback chain (popoff)

Advanced provides an important option for wallet issues: Rollback chain

<img src="13_Wallet_options.png" title="13_Wallet_options.png" alt="13_Wallet_options.png" width="509" height="563" />

If there are issues with starting the wallet, or an incomplete blockchain synchronization, it’s suggested to pop off 1000 blocks and allow to re-synchronize. The maximum blocks that can be popped off is 1400. 

<img src="Rollback_Chain.png" title="Rollback_Chain.png" alt="Rollback_Chain.png" width="450" height="450" />

### I don't really like H2 database... How do I change the database ?

#### Change database

Changing database allows for the migration of databases, this is an advanced option. 

The default database type installed with QBundle is H2. In case the user wishes, they can change the database type to Portable MariaDB or own MariaDB/MySQL. 

1.  Default database is H2 which is file based - simple and low on memory but can have issues if shut down incorrectly
2.  Portable MariaDB is higher performing, uses a lower disk footprint, more stable against sudden shutdowns, harder to bootstrap
3.  Own MariaDB /MySQL is the highest performing database, used for large installations and will be needed for high-stress side applications or financial or tax analytics and reporting.

<img src="14_Database_options.png" title="fig:14_Database_options.png" alt="14_Database_options.png" width="465" height="465" /> 

<img src="Change_Database.png" title="Change_Database.png" alt="Change_Database.png" width="504" height="504" />

If MariaDB is selected, QBundle will initiate the download and execute the installation of the required database.

<img src="MariaDB_DL.png" title="MariaDB_DL.png" alt="MariaDB_DL.png" width="509" height="788" />

The time needed to download installation files will depend on the speed of the available internet connection.

<img src="Download_MariaDB.png" title="Download_MariaDB.png" alt="Download_MariaDB.png" width="480" height="480" />

The current version of the Burst wallet and its' wrapper don't allow copying data from H2 to MariaDB, so in case users change the database type, they will need to re-synchronize the blockchain (be it through bootstrap or without it). 

<img src="Transition_Database.png" title="Transition_Database.png" alt="Transition_Database.png" width="485" height="485" />

### What are the other tools available ?

The “Tools” section of the QBundle contains additional Burst software - plotter and plot converter, miner and paper Burst wallet.

<img src="Tools_250222.png" title="Tools_250222.png" alt="Tools_250222.png" width="392" height="392" />

#### Plotter

Plotting is the process of calculating hashes and saving them onto storage units (HDDs). The latest version of QBundle comes with Xplotter version 1.30. Once it is downloaded, it will be saved in the QBundle folder, sub-folder Xplotter. Users can access it via Tools -&gt; Plotting -&gt; Plotter which takes them to the screen shown below:

<img src="17a_Plotter.png" title="17a_Plotter.png" alt="17a_Plotter.png" width="509" height="643" />

Users should provide the following information in the “Basic settings” before they start the plotter:

-   Location where the plot file should be saved to
-   Set the numeric account ID which will be associated with the plot. Clicking the selector on the right side from the form field will list accounts saved as described in Account Manager
-   Set the plot file size using the slider
-   Add to my plotfiles at “Start plotting” will allow the miner to use plots as soon as a scoop has been plotted i.e. it won't wait until the entire plot file is completed.

Advanced settings for Xplotter allow the user to:

-   Set the start nonce of the new plot file. In case all existing plots have been imported (recommended) the plotter will continue plotting using the number after the highest existing nonce - leaving gaps between nonces doesn't improve mining. In case the first plot is being created, the plotter will start with nonce 0
-   Set the number of CPU threads to be used
-   Set the amount of memory to be used
-   Plot type - using PoC2 plots is mandatory

Once all settings have been submitted, clicking the “Start Plotting” button with show the following screen:

<img src="17b_Plotter_wrorking.png" title="17b_Plotter_wrorking.png" alt="17b_Plotter_wrorking.png" width="509" height="819" />

The plotter informs the user about the current nonce being worked on (From- to) and the % completed (\[%\]) the rate per minute (nonces/min) and details about the writing speed.

The “My plotfiles” section allows the users to manage (import/remove) plot files.

<img src="17c_My_plotfiles.png" title="17c_My_plotfiles.png" alt="17c_My_plotfiles.png" width="509" height="628" />

#### Dynamic Plotting

Dynamic plotting allows for the normal usage of the drive, as it will check every minute for space pressure (i.e. a limited amount of space on the drive because of files created/copied).

This will create and delete plot files automatically, to maintain the free amount of space (the user determines the required free space) using plot files of the size directed. Essentially, if users move/create files often that are quite large, they should give a greater space and have larger plot sizes to take into account sudden changes.

<img src="17d_Dynamic_plotting.png" title="17d_Dynamic_plotting.png" alt="17d_Dynamic_plotting.png" width="509" height="651" />

#### Plot Converter

Johnny's POC1 -&gt; POC2 converter, a tool for performing the conversion of POC1 plot file types into POC2 plots is delivered with the Qbundle. The POC1 to POC2 conversion has been completed at Burst blockchain block height 502.000, and since then, POC1 plot are no longer supported for mining. Users who still have POC1 plots can, however, use them for mining if they use POC2 compliant miners, which can perform on-the-fly conversion which will result in cca 50% slower reading times. The POC1 to POC2 converter is a tool specifically designed to perform this conversion, but the plot file that it being converted has to be excluded from mining. The usage of the plot converter is quite straightforward: users will first select plot files for conversion (note that the application will not allow for other plot file types to be added to the conversion queue):

<img src="17e_POC1-POC2_converter_start.png" title="17e_POC1-POC2_converter_start.png" alt="17e_POC1-POC2_converter_start.png" width="509" height="872" />

Users can choose if they wish to perform inline conversion (doesn't require additional free space) or to save the converted plot file to a new location, in case of which the “Output Folder” is a mandatory setting. Clicking the “Start conversion” button will warn the user that they shouldn't stop the conversion process as it might lead to damaged plots, which cannot be repaired.

#### Miner

The latest version of QBundle brings the Blago's miner (modified by Quibus and JohnnyFFM) version 1.170977, which supports both POC1 and POC2 mining and performs POC1 -&gt; POC2 on the fly conversion. The miner can be set to either solo or pool mining, by selecting the desired radio button:

<img src="21_Miner_settings.png" title="21_Miner_settings.png" alt="21_Miner_settings.png" width="509" height="654" />

As the miner GUI states, in case the user wishes to solo mine, and chooses that option, they need to have a local wallet running and fully synchronized, and will be prompted to provide the passphrase of the account used to create the plot file(s). The passphrase will be saved in the miner.conf file which is located in the QBundle/BlagoMiner folder.

For pool mining, users can either select one of the predefined pools (by clicking the “Select predefined pool” button) or manually insert mining and update server information. These information refer to the pool where deadlines will be submitted and from where the mining information is received from. For pools that aren't offered in the predefined pools list, users can look up information, including the port number on pool sites. The deadline limit can be calculated using calculators provided on pool websites. In case the calculated deadline, which is based on available plot size, exceeds the pool deadline, the deadline limit should be the maximum deadline accepted by the pool (the pool won't accept deadlines that exceed the maximum defined). The information server parameter is used to get winner information i.e. the account that forged the block. Since the info server doesn't support HTTPS, the address of the local wallet (localhost or 127.0.0.1) and port number 8125 should be provided.

If the “Use HDD wakeup” option is checked, a script will be executed in order to prevent the HDD from entering sleep mode. The “Show winner information” option will show the account that forged the block, in case the “Info server” has been set. “Use multithreading” refers to CPU resources usage - if checked, multithreading will be used for deadline calculation.

Users can import or remove plot files used for mining by selecting files from the “My plotfiles” list. Note that the miner will use all plot files that are found in configured folders.

All settings provided in the interface shown above will be passed on to the miner.conf file, which is located in the Qbundle/BlagoMiner folder. Users can also set or edit the miner configuration by editing directly the .conf file.

The miner will start after the “Start mining” button has been clicked.

<img src="21a_Blago_miner.png" title="21a_Blago_miner.png" alt="21a_Blago_miner.png" width="509" height="736" />

Once the miner has been started, users can see the configuration (pool and plot file information). Upon receiving mining information (block height, base target etc), the miner will read the configured plot files. Whenever a deadline is found, if it is within configured deadline limit, it will be sent to the pool. The miner will also display if a deadline was accepted by the pool or not (“Confirmed DL”). After all plot files have been read, the miner will display the duration of the round (how much time did it take to read all plot files).

The Blago miner creates logs, which are stored in the QBundle/BlagoMiner/Logs folder, which can be useful for debugging in case issues are observed. Additionally, the miner records the best deadline found for each block in the stat.csv file, which is saved in the QBundle/BlagoMiner folder.

#### Set Reward Recipient

Setting the reward recipient is a mandatory action for all miners. In case users wish to mine in one of the Burst mining pools, they should provide the pool's Burst numerical account ID as the reward recipient, while solo miners set their own Burst numerical account ID as the reward recipient.

<img src="18_Reward_recipient.png" title="18_Reward_recipient.png" alt="18_Reward_recipient.png" width="424" height="424" />

In case the pool user wishes to mine in isn't offered within the QBundle, they can insert the pool's Burst numerical account ID manually. The reward recipient transaction takes 4 blocks to complete i.e. after the generation of 4 blocks the deadlines submitted by the miner will be accepted.

#### Vanity Address Generator

The vanity address generator is a tool allowing users to create own, branded, specific Burst account addresses. The mandatory part of the Burst account address, even for vanity addresses is the leading 5 character string “BURST”. Characters that are not allowed to be used in vanity addresses are I (capital “i”), O (capital “o”), 0 (zero) and 1. In order to generate a vanity address, users should provide the requested string into the form, set the resources that should be used (number of CPU threads) and desired length of the passphrase, as shown in the image below:

<img src="Vanity_Address_Generator.png" title="Vanity_Address_Generator.png" alt="Vanity_Address_Generator.png" width="509" height="653" />

The algorithm which seeks for the account with the desired string is a brute force algorithm, which checks every Burst account + passphrase combination sequentially. The time it takes to find a string increases significantly with every additional character. A fixed strings of 5 characters at the end took 4 minutes with given resources, while adding ‘DAN’ (to output DANDARES) at the end, resulted in hour long wait. It is possible that setting certain strings as a part of the Burst account will take weeks.

#### Paperburst

The Paperburst tool can be used to generate a new wallet or to use an existing wallet. In order to use the tool, users will provide information as shown below:

<img src="20_Paper_wallet_generator.png" title="20_Paper_wallet_generator.png" alt="20_Paper_wallet_generator.png" width="509" height="658" />

After all reqiured fields of the form have been populated, by clicking the desired button, users can create a new Burst wallet or a paper wallet. The paper wallet is output as a .pdf file, so it can be easily printed.

Paper wallets can be used to transfer funds (using public key to provide details of a wallet to transfer to) or to store funds (the funds are still on the blockchain and not on the paper, but it helps as a reference document for storage).

<img src="20a_Paper_wallet_generator.png" title="20a_Paper_wallet_generator.png" alt="20a_Paper_wallet_generator.png" width="1093" height="1093" />

#### Help

The “Help” section offers link to Burstcoin resources.

<img src="16_help.png" title="16_help.png" alt="16_help.png" width="569" height="569" />

Qbundle manual: [QBundle](qbundle.md)

Burstcoin official website: https://www.burst-coin.org/

Burstcoin.ist: https://www.burstcoin.ist/

Burstforum: https://burstforum.net/

GetBurst Forum: https://forums.getburst.net/

Burst Reddit: https://www.reddit.com/r/burstcoin/

Burst Wiki: [Main Page](main-page.md)

### About Tab

The “About” menu option contains information about the contributors of QBundle and the main architect.

<img src="15_about.png" title="15_about.png" alt="15_about.png" width="565" height="565" />

Contributors list:

<img src="15a_contributors.png" title="15a_contributors.png" alt="15a_contributors.png" width="601" height="601" />

# Description

The easiest choice for beginners using Windows is to download Qbundle. QBundle is a convenient launcher which installs the BRS local wallet, plotting software, and mining software with a few clicks. 

# Installation guide
Download the QBundle installation files and return here to complete the installation. Sets of instructions are grouped logically, with reference images of installation screens immediately following each section. Images may vary from the actual installation screens depending on your version of Windows and internet browser. 

## Download the installation Files
The installation files are located on BAT's github: https://www.burst-coin.org/download-wallet. Open this link in a new browser tab if possible. After following this link:

* Click on "Download wallet".
* Select your operating system: (Please note: If your operating system is not Windows, QBundle will not be available and you will need to choose an option specific to your operating system.)

* Select Qbundle to be redirected to the PoC Consortium GitHub site. GitHub is a software development platform and hosts the files needed for installation. Future releases will be located here as well. 

* Download and save the latest release, identified by a green indicator. As of the writing of this guide, the latest release is Version 2.5.1. The installation file is QBundle2.5.1.zip. (Make a note of the saved location)

## Install the software

* Locate the saved file. A common location is C: Users/"user name"/Downloads. You may move this file to another location if you wish.
* Extract the included files using Windows "Extract all" button if available. If you do not have this functionality, you may need a utility to "unzip" the folder.
* Launch the setup of QBundle by clicking on the extracted version of the following file: Name: BurstWallet Type: Application.
* Proceed through the Windows warnings related to protecting your computer by clicking on "More Info" followed by "Run Anyway".

The installation utility will check for portable Java and the BRS wallet.

* Click on "Download missing components"
* Verify that the check box for "allow connection to remote resources for verifications and updates" is checked.

Several windows will appear as each component and is downloaded and extracted.

* When all of the components have downloaded. Click on "continue". 

Setup will complete and a new wallet will be displayed showing a status of "stopped". 

* Click on "Database" on the wallet's top menu bar and select "Change Database".
* Select the radio button next to "Portable MariaDB" and Click "Next".
* Click "Yes" when asked if you want to download and install MariaDB.
* After installation is complete, click "Database" on the wallets to menu bar, verify that "No Copy" is selected.
* Click "Save and close".

* Change Java settings by clicking on "edit" in the top menu bar. Select "settings". The settings window will appear.
* Select "Java" in the left menu bar. Verify that the radio button for "Use Portable java" is selected. Save your change and exit.

## Bootstrap the blockchain (populate the database)

New wallets immediately begin to synchronize with the Burst blockchain. This can be a time consuming process depending on your internet bandwidth and the processing speed of your computer. For a faster alternative, download a "bootstrap" copy of the blockchain which contains the genesis block and all confirmed blocks through the time that the bootstrap copy was created. Only the more recent blocks will then synchronize through the regular peer to peer synchronization process.

* To begin the bootstrap process, Select "Database" in the top menu bar.
* Select "Bootstrap Chain".
* In "Settings", verify that the Cryptoguru repository is selected.
* Click "Start Import". You will receive a notice that all existing data in your database will be erased.
* Click "Yes" to continue.

Wait for synchronization to complete before proceeding. Block time will appear in orange while syncing, and green when fully synchronized. 

## Create a new account and passphrase

Please note: Before proceeding with the next step, it is important that you fully understand the security features of the Burst passphrase. If you have not done so already, please review "Securing your Burst". A link is located on the main page of the Burst Wiki.

* If you are a new user and do not already have an account, click on "New? Create Your Account".

A 12 word passphrase will be generated automatically. You can extend the passphrase by adding a few extra words, numbers, or special characters (Unicode only). This is the only opportunity that you will have to modify the passphrase. Modifying the passphrase is unnecessary as it is considered to be secure. However, some feel that security is further enhanced by adding something additional. Never reduce security by shortening the passphrase.

* Record this passphrase on paper for now, including any modifications that you have made. You can store this passphrase more securely later. When you are satisfied with your passphrase, Click "Next".
* Type the passphrase, exactly as you have recorded it on paper, in the verification box. Click "Next" to open your new wallet.
* Locate and record your new Burst account number on paper. It is located in the top left corner of your wallet. Burst account numbers contain alphanumeric characters in five segments. The first segment is always "BURST".
* Click on "Wallet" and then click on "Stop Wallet".
* Restart the wallet by clicking on "Wallet" again and then clicking on "Start Wallet".
* Test your passphrase by clicking on "Returning user?" and typing it into the passphrase box. Click on the arrow located to the right of the box. When the wallet is opened, if the Burst account number displayed in the wallet matches the record that you made of your Burst account number, then you have successfully completed the setup process.

Please note: If you enter your passphrase incorrectly, you will generate a new account. If this occurs, simply close the wallet and try again. 
    
#Install additional QBundle applications

Blago's miner (modified by Quibus and JohnnyFFM) is the default mining software packaged with QBundle. It can be used for pool mining or for solo mining.

Please note: If solo mining, running a fully synchronized local wallet is required. The passphrase for the mining account associated with your plot files must be provided and will be stored in plain text in QBundle/BlagoMiner/miner.conf. If desired, you can create an account specifically for mining rather than using a single wallet for both mining and maintaining significant balances.

* To install, click on "File" and then "Application Manager".
* Place a check mark in the box next to "Blago's Miner".
* Click on "Install/Update"
* After the software downloads and extracts, verify that an extracted (unzipped) folder with the name "BlagoMiner" is located in the main QBundle folder. If the file does not appear, locate it a sub folder of the zipped version and copy it into the main folder (same level as zipped version) as the final installation step.
* Open the miner by clicking "Tools" and the "Miner". The following screen will appear.

All settings provided in the interface shown above will be passed to the miner.conf file. This file is located in the Qbundle/BlagoMiner folder. Users can edit this file directly.

* If you will be solo mining, click on the "solo mining" radio button.

* If you will be pool mining, click on "Select predefined pool". Required configuration data will be populated automatically. If your desired pool is not listed, obtain the configuration information directly from the pool operator and enter it manually (mining server, update server, info server, deadline limit, and port information).
* Set the deadline limit. Calculators for this purpose are provided by each pool. If your calculated deadline (based on your available plot size) exceeds your pool's maximum deadline, this limit should be reduced to the maximum accepted by the pool. 

Optional setup:

"Use HDD wakeup" - a script will run which prevents your hard drives from entering sleep mode. "Show winner information" - the account which forges each block will be displayed in the miner console. The info server does not support HTTPS so the address of the local wallet (localhost or 127.0.0.1) and port number 8125 should be provided. "Use multithreading" - Multithreading may improve usage of CPU resources when calculating deadlines.

* Add your plotfile paths to "My Plotfiles" by clicking on "Import Plotfiles" and locating each plotfile or folder containing plotfiles. The miner will use all plot files that are found in configured folders. It is best if the folders used only contain individual plot files. Plotfiles that are removed using the "Remove Plotfiles" button are not deleted from storage, only their path in the list of plot files is removed. 

* To start the miner, click on the "Start mining" button.

When the miner is running, pool and plot file information is displayed in the mining console. The miner receives mining information (block height, base target, etc.) and reads through the configured plot files to locate deadlines. Deadlines found are sent to the pool if they are less than or equal to the configured deadline limit. The miner displays "Confirmed DL" if the deadline is accepted by the pool. After all plot files have been read, the miner displays the amount of time taken to read the plot files.

The miner logs are stored in QBundle/BlagoMiner/Logs. These are useful for debugging purposes. The best deadline found for each block is stored in QBundle/BlagoMiner/stat.csv. When the miner is running, the console window will appears as follows: 

## Xplotter

Plotting is the process of pre-computing hash functions and saving them onto storage units (HDDs). XPlotter is included with QBundle and no additional installation is required. There are however options that need to be configured.

* Click on Tools: Plotting: Plotter
    
* Add any plotfiles that you already have. Use the "Import Plotfile" button.

* Provide the following information in "Basic settings" before starting the plotter:

Location for saving plot files Numeric account ID to be associated with the generated plots. Click the selector to the right of the form field to select from a list of accounts. If you enter the ID number manually, be absolutely sure that it is entered correctly. If not, your plot files will be unusable. Be sure that no trailing spaces are included, particularly if you use a copy and paste function to enter the ID.

* Set plot file size using the slider. For easier maintenance, limit plot size to 1 or 2 Terra bytes.

Adding files to "My Plotfiles" and clicking on "Start plotting" will allow the miner to begin using the plots as soon as a scoop has been plotted, it is not necessary to wait until the entire plot file is complete. Advanced settings for Xplotter allow the user to:

* Set the start nonce of the new plot file. If all existing plots have been imported (recommended) the plotter will continue plotting using the number after the highest existing nonce. Leaving gaps between nonces does not improve mining. If no plot file has been created, the plotter will start with a nonce number of 0.
* Set the number of CPU threads to be used.
* Set the amount of memory to be used.
* For Plot type, use PoC2 (mandatory). Do not use POC1 as it is no longer supported.
* After all settings have been submitted, click "Start Plotting". 

The plotter shows information about the current nonce being generated: (From - to), % completed ([%]), rate per minute (nonces/min), and writing speed.

The "My plotfiles" section allows users to import and remove plot files. Removing a plot file in this section will not delete the plot file. Removed plot files can be added back at a later time. 

## Vanity Address Generator

The vanity address generator is a tool for creating vanity or branded Burst account addresses. All address must have "BURST" for the first five leading characters. The capital letter "i" (I), the capital letter "o" (O), zero (0), and the number one (1) are not permitted. To generate a vanity address, enter characters into the form provided. Set the resources that should be used (number of CPU threads) and desired length of the passphrase as shown in the image below: 

The algorithm which seeks for the account with the desired string is a brute force algorithm, which checks every Burst account and passphrase combination sequentially. The time it takes to find a particular string increases significantly with every additional character. A string of 5 characters at the end takes approximately 4 minutes (more or less depending on computer resources). Each additional character increases the time required exponentially. A string of 8 would take approximately one hour. Longer strings could take weeks. 

## Paperburst

The Paperburst tool can be used to generate a new wallet or to use an existing wallet. In order to use the tool, provide information as shown below: 

After all required fields of the form have been populated, clicking the desired button. The paper wallet is output as an easily printable .pdf file.

Paper wallets can be used to transfer funds using public keys, to store funds (the funds are still on the blockchain rather than on the paper). Paperburst can be used as a reference document for cold storage. 

## Dynamic Plotting

Dynamic plotting creates and deletes plot files automatically while allowing for normal usage of a hard drive. It will check every minute for space pressure. A limited but dynamic amount of space on the hard drive is reserved for normal usage. A user can set specify the amount to be reserved. If users often move or create files that are large, they should set a greater amount of space to be reserved. 

## Plot Converter

Johnny's POC1 -> POC2 converter is a tool for converting POC1 plot files into POC2 plot files. New users will not need to use this tool. The POC1 to POC2 conversion completed at block height 502000. POC1 plot are no longer supported for mining. Users who have existing POC1 plots can use them for mining if they use a POC2 compliant miner which can perform on-the-fly conversion. Read speeds will be 50% slower using POC1. Plots that are being converted must be temporarily excluded from mining. Usage of the plot converter is straightforward: First, select plot files for conversion and then click on "Start conversion". Note: The application will not allow plot files of other types to be added to the conversion queue.

Users can perform inline conversion, which does not require additional free space, or can save the converted plot file to a new location. To save to a new location, use the "Output Folder" setting. Once the "Start conversion" has been activated, stopping the conversion process may lead to damaged plots which cannot be repaired. Many miners have chosen to re plot their drives for POC2 rather than use the conversion process. 

## Help

The "Help" tab contains links to official Burst resources.

If you are unable to open your wallet, these same resources can be found by navigating from the main page of the Burst Wiki. 

# Frequently Asked Questions (FAQs) relating to QBundle

## How do I set my reward recipient

Setting the reward recipient is a mandatory action for all miners. For pool mining, provide the pool's Burst numerical account ID. For solo mining, use your own account's numerical ID. 

In the pool you wish to join is not listed within QBundle, insert the pool's Burst numerical account ID manually. The reward recipient transaction requires 4 blocks (confirmations) before it will be effective.

## What is Wallet Mode?

Wallet Mode loads the complete interface and allows users to issue transactions and see blocks as they are forged on the network. Launcher mode opens a minimal interface and can only be accessed through an internet browser.

## How is QBundle updated?

The application manager compares all of the QBundle applications that are available for installation, as well as all of those that are installed on the system. Users are notified through the application manager when updates for these applications are available. Not all of the applications available through QBundle are installed by default. To add an application, checking the box next to the application and click on "Install/Update". Please Note: Restarting the wallet may be required for some applications to install or to reflect a status of installed. For important updates, a "New update available" message will appear at the top of the wallet. When this message is clicked, the application manager will be displayed and the update can be installed by clinking on "Update/Install". The wallet will be stopped during all software updates. The following figures show an application manager before and after and update has taken place. 

## How can I view logs for Burst Reference Software (BRS)?

Click on "Edit" and then click on "View Console". A pop up window with the Wallet log and the MariaDB log (if applicable) will appear. This is useful for troubleshooting and the output is often needed when requesting technical support from the community. 

## How do I change QBundle settings?

Users are generally advised not to change any of the default settings. If synchronization is causing your computer to operate more slowly, changes might be appropriate. In this case decreasing the number of cores may help. Otherwise most setting are used only for troubleshooting.

If OpenCL is correctly installed, using GPU Acceleration will greatly improve performance.

To allow external access to the wallet, add 0.0.0.0 to the ‘Allow API traffic from’ list.

For troubleshooting, the "Debug mode" option can be useful.

## How can I resolve connection problems?

If the "Configure Windows firewall" is clicked, user will be asked to confirm that they wish to change Windows firewall to allow program access locally and over WAN.

## Where is the information about my account stored?

The account manager can be used to manage multiple accounts. A short PIN number can be used for each.

The account Manager is accessible here:

Users can easily manage all their BURST accounts with this software. It provides the Reed-Solomon address and the numeric account number of their Burst Wallet. The Public Key for each of those accounts is also shown.

Passphrases and private keys are also available (it is enough to set up one PIN to control multiple accounts) 

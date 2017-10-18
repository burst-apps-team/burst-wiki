<languages/> <translate>

History
-------

The development of the [PoC Consortium](poc-consortium.md) (PoCC) [mobile wallet](mobile-app.md) for Burst has been announced August, 11th 2017 in the original PoCC announcement on bitcointalk[1]. One month later, a beta-test version has been published[2] and the official released followed by end of September[3]

Features
--------

-   Watch only addresses
-   Currency conversion
-   Client-side encryption and decryption
-   QR code support
-   Secure and easy passphrase generation

I18N - Internationalization
---------------------------

The major part of the wallet localization is done via a .json file containing definitions for text strings used in the wallet. The naming is such, that a 2-letter code denotes the language. en.json is the source file (see below).

If you want to translate this file into some language, you need to translate the “Create New Account”, “Done” etc. not those in ALLCAPS. A major constraint is, that the translated text not only be correct and stylistic, but also of about the same length as the original content. It can be shorter, but not much longer. The 2-letter language code is [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes)

### Translation priorities

-   done languages: de, el, en, es, fr, hu, it, pl, ru, sk, sv, ta, zh
-   work-in-progress: ar, jp
-   wanted languages: hi, ml, ko, and basically everything else

<div class="toccolours mw-collapsible mw-collapsed">
en.json localization file

<div class="mw-collapsible-content">
    {
        "CREATE": {
            "ACTION_BAR_TITLE": "Create New Account",
            "DONE": "Done",
            "GENERATE_AGAIN": "Generate Again",
            "INFO_SEED": "Tap and swipe over this area.",
            "INFO_SEED_FINISHED": "Seed collected.",
            "INFO_STEP_ONE": "First, create a random seed for the generation of your passphrase.",
            "INFO_STEP_TWO": "Now you have to memorize all 12 words of your passphrase in the right order.",
            "INFO_STEP_TWO_ACTION": "Write down the following words separated by space and keep them secret.",
            "INFO_STEP_THREE": "Please enter the words of your passphrase again.",
            "INFO_STEP_THREE_ACTION": "Write down word #{{value}} and select it.",
            "INFO_STEP_FOUR": "Please set a six-digit pin code for your account. The pin code is used to approve outgoing transactions.",
            "PIN_INPUT_HINT": "PIN",
            "POSSIBILITY_INPUT_HINT": "Word #{{value}}",
            "NEXT": "Next"
        },
        "IMPORT" : {
            "ACTION_BAR_TITLE": "Import Existing Account",
            "ACTIVE_ACCOUNT": "Active Account",
            "ACTIVE_ACCOUNT_DESCRIPTION": "An active account offers full functionality. You can send and receive Burstcoins. In addition, you can check your balance and see the history of your transactions.",
            "ACTIVE_ACCOUNT_INPUT_HINT": "Passphrase",
            "ACTIVE_ACCOUNT_PIN_INFO": "Please set a six-digit pin code for your account. The pin code is used to approve outgoing transactions.",
            "DONE": "Done",
            "IMPORT": "Import",
            "INFO": "Import an active account by entering its passphrase. Otherwise, import an offline account by entering its Burstcoin address.",
            "NEXT": "Next",
            "OFFLINE_ACCOUNT": "Offline Account",
            "OFFLINE_ACCOUNT_DESCRIPTION": "An offline account offers the same functionality than an active account, except you cannot send Burstcoins to another address.",
            "OFFLINE_ACCOUNT_INPUT_HINT": "BURST-XXXX-XXXX-XXXX-XXXXX",
            "PIN_INPUT_HINT": "PIN",
            "SHOW" : {
                "QUESTION": "Is this your Burst address ?"
            }
        },
        "TABS": {
            "ACCOUNTS": {
                "ACTIVATE": {
                    "ACTION_BAR_TITLE": "Activate Account",
                    "DONE": "Done",
                    "ENTER_PASSPHRASE": "Enter the passphrase of the account {{value}} here.",
                    "ENTER_PIN": "Set a six-digit PIN code.",
                    "PASSPHRASE_INPUT_HINT": "Passphrase",
                    "PIN_INPUT_HINT": "Pin",
                    "NEXT": "Next"
                },
                "BALANCE": "Balance",
                "CREATE": "Create",
                "IMPORT": "Import",
                "REMOVE": "Do you really want to remove this account ?",
                "TITLE": "Accounts"
            },
            "BALANCE": {
                "ACTIVATE": "Activate",
                "BALANCE": "Balance",
                "SEND": "Send",
                "TITLE": "Balance"
            },
            "HISTORY": {
                "TITLE": "History"
            },
            "SETTINGS": {
                "ABOUT": {
                    "ABOUT": "About",
                    "ASSOCIATION": "In Association with",
                    "DEVELOPED": "Developed by",
                    "TWITTER": "Follow us on twitter to get updates on our projects!",
                    "VERSION": "Version"
                },
                "CURRENCY": {
                    "CURRENCY": "Currency"
                },
                "DOCUMENTATION": {
                    "DOCUMENTATION": "Documentation"
                },
                "LANGUAGE": {
                    "LANGUAGE": "Language"
                },
                "NODE": {
                    "INFO": "The Burst node is used to communicate with the Burst network",
                    "NODE": "Node"
                },
                "SUPPORT": {
                    "CONTRIBUTION": "We also appreciate non-monetary contributions",
                    "DONATE": "You like our app ? Support us by donating BURST!",
                    "SUPPORT": "Support"
                },
                "TITLE": "Settings"
            }
        },
        "SEND": {
            "ACCEPT": "Accept",
            "ACTION_BAR_TITLE": "Send BURST",
            "AMOUNT": "Amount",
            "BALANCE": "Balance",
            "FEE": "Fee",
            "RECIPIENT": "Recipient",
            "PIN_INPUT_HINT": "Pin",
            "VERIFY": "Verify"
        },
        "START" : {
            "CREATE": "Create",
            "IMPORT": "Import",
            "INFO": "No account created yet. Choose to import an existing account or create a new one!"
        },
        "NOTIFICATIONS" : {
            "ADDRESS": "Please enter a valid Burstcoin address!",
            "COPIED": "Copied address {{value}} to clipboard!",
            "DECIMAL_AMOUNT": "Please enter a decimal number for the amount of BURST you want to send!",
            "DECIMAL_FEE": "Please enter a decimal number (atleast 1) as fee!",
            "ENTER_SOMETHING": "Please enter something!",
            "EXCEED": "Amount and Fee exceed your balance!",
            "PIN": "Please enter a six-digit number as Pin!",
            "REMOVE": "Successfully removed account!",
            "SELECTED": "Selected account: {{value}}",
            "TRANSACTION": "Transaction successful!",
            "UPDATE_CURRENCY": "Currency successfully updated!",
            "UDPDATE_LANGUAGE": "Language successfully updated!",
            "UPDATE_NODE": "Node successfully updated!",
            "WRONG_PASSPHRASE": "Wrong passphrase! The provided passphrase does not generate the public key assigned to your account!",
            "WRONG_PIN": "The provided pin does not match the account's pin!",
            "ERRORS": {
                "ACCOUNT_ID": "Cannot generate account id from public key!",
                "ADDRESS": "Cannot generate BURST address from account id!",
                "CONNECTION": "Failed fetching data. Check your internet connection!",
                "CURRENCY": "Currency update failed!",
                "KEYPAIR": "Failed to generate keypair for passphrase!",
                "LANGUAGE": "Language change failed!",
                "NODE": "Node update failed!",
                "QR_CODE": "Could not scan QR code!",
                "REDIRECT": "Cannot open send BURST page! Current account is not active!",
                "REMOVE": "Could not remove account!",
                "UPDATE": "Update of account failed!",
                "UNKNOWN_ACCOUNT": "Account is still unknown. No transactions yet!"
            }
        }
    }

</div>
</div>
References
----------

</translate>

[1] <https://bitcointalk.org/index.php?topic=2080040.0>

[2] <https://forums.getburst.net/t/one-month-of-pocc/603>

[3] <https://forums.getburst.net/t/burst-on-the-go-the-poc-consortium-releases-their-android-wallet/794>

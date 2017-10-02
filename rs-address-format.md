<languages/> The well known BURST addresses of the form:

  
**BURST-XXXX-XXXX-XXXX-XXXXX**

are called Reed-Solomon addresses. This is the default format in the official client. where X is a non-ambiguous number or alphabetic character (the letters O and I are not used; nor are the numbers 1 and 0). Addresses are always prefixed with “BURST-”, and hyphens are used to separate the address into groups of 4, 4, 4, and then 5 characters. The addresses are NOT case-sensitive.

This form of address improves reliability by introducing redundancy that can detect and correct errors when entering and using Burst account numbers.

Background
----------

The internal format for Burst account numbers is a completely numeric 64-bit identifier that is derived from the account's private key. This format is error-prone, since a single mistyped character can result in transactions being unintentionally sent to the wrong account.

[Reed-Solomon error-correction codes](http://en.wikipedia.org/wiki/Reed%E2%80%93Solomon_error_correction) addresses this issue by adding redundancy to addresses. Reed-Solomon format was chosen because:

-   the account collision rate is the same as the default address format;
-   the system's basic error correction can be used to assist users in typing addresses;
-   some programming languages do not have a native MD5 hashing function, and the Reed-Solomon implementation is simpler than MD5.

### Benefits of Reed-Solomon addresses

-   The chance of a random address collision, using Burst's implementation of 4 “check-bits”, is 1 in a million (20-bit redundancy).
-   It allows up to 2 typos in an address to be *corrected*.
-   It guarantees that up to 4 typos can be *detected*.
-   The address length is always 17 characters, and is always prefixed with “BURST”. This makes the addresses easily recognizable as belonging to Burst

Encoding of Burst Reed-Solomon addresses
----------------------------------------

-   Case is not enforced in this format, but for unification all addresses are displayed using upper case.
-   Addresses are split by dashes into groups of 4 characters and a final group of 5 characters, but this is not enforced during address input.
-   The old numeric addresses are also recognized and supported for backwards compatibility.

Example RS Addresses:

  
    BURST-3DH5-DSAE-4WQ7-3LPSE

    BURST-K4G2-FF32-WLL3-QBGEL

Technical Details
-----------------

The first and most important rule is that no error-correction scheme is infallible: **You cannot rely on error correction, period.**

The problem is somewhat counter-intuitive: either you can do a simple yes/no check of address validity, which will give you one in a million collision, or you can try and correct errors. You cannot do both.

The problem here is that the Reed-Solomon algorithm is only guaranteed to correct up to 2 errors. If there are more than 2 errors present in an address entry, it will produce false positives with a probability of around 10% and transactions will still be sent to incorrect addresses.

Think of the algorithm as *error-guessing*, instead, to assist users with spotting errors.

Reed-Solomon (RS) addresses for Burst are encoded as follows:

1.  Take the original 64-bit account ID, add 1 zero bit to get 65 and then split it into thirteen 5-bit “symbols” (65 / 5 = 13).
2.  Order the symbols from lowest bit to highest bits, in little-endian order, i.e. bits 0-4, 5-9, 10-14, etc. up to 60-64.
3.  Append 4 symbols of parity (20 bits), produced by the [Reed-Solomon encoding](http://en.wikipedia.org/wiki/Reed%E2%80%93Solomon_error_correction) of our 13 symbols from step one (which are left untouched). This produces a 13 + 4 = 17 symbol codeword.
4.  Scramble the codeword symbols in a predefined order and encode them 1-to-1 with an alphabet of 32 characters, splitting them into groups by dashes.


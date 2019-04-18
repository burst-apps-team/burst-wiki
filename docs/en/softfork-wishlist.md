The **Softfork Wishlist** is to record changes to Burstcoin that might be desirable, but that will require a [“soft” block-chain split](softfork.md) (consensus of the miners). These changes are implemented by convincing a majority of the miners to reject or [discourage](discouraged-block.md) blocks that were previously considered valid.

Cryptographic changes
---------------------

-   Support for more than one public-key cryptosystem
-   Support for a [post-quantum](http://en.wikipedia.org/wiki/Post-quantum_cryptography) signature scheme. [Lamport signatures](http://en.wikipedia.org/wiki/Lamport_signature) have nice intuitive security properties, but it has extreme space requirements that would require structural changes to the blockchain to accommodate, namely pruning. However other post-quantum schemes like [IEEE1363.1](http://en.wikipedia.org/wiki/NTRU) (patent expired in 2016) do not. See the [introductory chapter of DJB's book](http://pqcrypto.org/www.springer.com/cda/content/document/cda_downloaddocument/9783540887010-c1.pdf) and the [PQCrypto conference webpage](http://pqcrypto.org) for full details.


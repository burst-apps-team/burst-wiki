Capacity Improvement Proposal
-----------------------------

Similar to the practice of [BIPs](https://en.bitcoin.it/wiki/Bitcoin_Improvement_Proposals) and [EIPs](https://github.com/ethereum/EIPs), the Burst community has established CIPs (short for “Capacity Improvement Proposal” or even “Coin Improvement Proposal” to advance further development of Burstcoin.

So CIPs describe standards for the Burstcoin platform, including core protocol specifications, client APIs, and contract standards.

Abstract
--------

A (Burst)Coin or Capacity Improvement Proposal (CIP) is a design document providing information to the Burstcoin community, or describing a new feature for Burstcoin or its processes or environment. The CIP should provide a concise technical specification of the feature and a rationale for the feature.

We intend CIPs to be the primary mechanisms for proposing new features, for collecting community input on an issue, and for documenting the design decisions that have gone into Burstcoin. The CIP author is responsible for building consensus within the community and documenting dissenting opinions.

Because the CIPs are maintained as text files in a versioned repository, their revision history is the historical record of the feature proposal.

CIP workflow
------------

The CIP process begins with a new idea for Burstcoin. Each potential CIP must have a champion -- someone who writes the CIP using the style and format described below, shepherds the discussions in the appropriate forums, and attempts to build community consensus around the idea. The CIP champion (a.k.a. Author) should first attempt to ascertain whether the idea is CIP-able.

Small enhancements or patches to a particular piece of software often don't require standardisation between multiple projects; these don't need a CIP and should be injected into the relevant project-specific development workflow with a patch submission to the applicable issue tracker.

Additionally, many ideas have been brought forward for changing Burstcoin that have been rejected for various reasons.

The first step should be to search past discussions to see if an idea has been considered before, and if so, what issues arose in its progression.

After investigating past work, the best way to proceed is by posting about the new idea to the [Burst Discord Forum](https://discord.gg/PMUgVSY).

Vetting an idea publicly before going as far as writing a CIP is meant to save both the potential author and the wider community time.

Asking the Burstcoin community first if an idea is original helps prevent too much time being spent on something that is guaranteed to be rejected based on prior discussions (searching the internet does not always do the trick).

It also helps to make sure the idea is applicable to the entire community and not just the author. Just because an idea sounds good to the author does not mean it will work for most people in most areas where Burstcoin is used.

Once the champion has asked the Burstcoin community as to whether an idea has any chance of acceptance, a draft CIP should be presented to the [Burst Discord Forum](https://discord.gg/PMUgVSY).

This gives the author a chance to flesh out the draft CIP to make it properly formatted, of high quality, and to address additional concerns about the proposal.

Following a discussion, the proposal should be submitted to the \[List\_of\_CIPs CIPs repository\]. This draft must be written in CIP style as described below, and named with an alias such as “cip-johndoe-infiniteburstcoins” until the editor has assigned it a CIP number (authors MUST NOT self-assign CIP numbers).

CIP authors are responsible for collecting community feedback on both the initial idea and the CIP before submitting it for review. However, wherever possible, long open-ended discussions on public mailing lists should be avoided. Strategies to keep the discussions efficient include: setting up a separate SIG mailing list for the topic, having the CIP author accept private comments in the early design phases, setting up a wiki page etc. CIP authors should use their discretion here.

It is highly recommended that a single CIP contain a single key proposal or new idea. The more focused the CIP, the more successful it tends to be. If in doubt, split your CIP into several well-focused ones.

When the CIP draft is complete, the CIP editor will assign the CIP a number, label it as Standards Track, Informational, or Process, and move the submitted text to the CIPs repository.

The CIP editor rewievs incoming CIPs for their formal completeness and will not unreasonably reject a CIP.

Reasons for rejecting CIPs include duplication of effort, disregard for formatting rules, being too unfocused or too broad, being technically unsound, not providing proper motivation or addressing backwards compatibility, or not in keeping with the Burstcoin philosophy.

For a CIP to be accepted it must meet certain minimum criteria. It must be a clear and complete description of the proposed enhancement. The enhancement must represent a net improvement. The proposed implementation, if applicable, must be solid and must not complicate the protocol unduly.

The CIP author may update the draft as necessary in the wiki. Updates to drafts should also be submitted by the author.

### Transferring CIP Ownership

It occasionally becomes necessary to transfer ownership of CIPs to a new champion. In general, we'd like to retain the original author as a co-author of the transferred CIP, but that's really up to the original author. A good reason to transfer ownership is because the original author no longer has the time or interest in updating it or following through with the CIP process, or has fallen off the face of the 'net (i.e. is unreachable or not responding to email). A bad reason to transfer ownership is because you don't agree with the direction of the CIP. We try to build consensus around a CIP, but if that's not possible, you can always submit a competing CIP.

If you are interested in assuming ownership of a CIP, send a message asking to take over, addressed to both the original author and the CIP editor. If the original author doesn't respond to email in a timely manner, the CIP editor will make a unilateral decision (it's not like such decisions can't be reversed :).

### CIP Editors

The current CIP editor is rico666 who can be contacted at <bots@cryptoguru.org>.

### CIP Editor Responsibilities & Workflow

The CIP editor is present in the Burst Discord channel. Off-list CIP-related correspondence should be sent (or CC'd) to bots@cryptoguru.org.

For each new CIP that comes in an editor does the following:

-   Read the CIP to check if it is ready: sound and complete. The ideas must make technical sense, even if they don't seem likely to be accepted.
-   The title should accurately describe the content.
-   A link to the CIP draft must have been sent to the Burst Discord channel for discussion.
-   Motivation and backward compatibility (when applicable) must be addressed.
-   The defined Layer header must be correctly assigned for the given specification.
-   Licensing terms must be acceptable for CIPs.

If the CIP isn't ready, the editor will send it back to the author for revision, with specific instructions.

Once the CIP is ready for the repository it should be submitted to the \[\[List\_of\_CIPs CIPs repository\] where it may get further feedback.

The CIP editor will:

-   Assign a CIP number.
-   Move the CIP to its final Wiki name when it is ready.

The CIP editors are intended to fulfill administrative and editorial responsibilities. The CIP editors monitor CIP changes, and update CIP headers as appropriate.

### CIP format and structure

Specification

CIPs must be written in mediawiki format.

Each CIP should have the following parts:

-   Preamble -- Headers containing metadata about the CIP (see below).
-   Abstract -- A short (~200 word) description of the technical issue being addressed.
-   Copyright -- As the CIPs are part of this Burstwiki project the license is GNU FDL.
-   Specification -- The technical specification should describe the syntax and semantics of any new feature. The specification should be detailed enough to allow competing, interoperable implementations for different Burstcoin platforms.
-   Motivation -- The motivation is critical for CIPs that want to change the Burstcoin protocol. It should clearly explain why the existing protocol is inadequate to address the problem that the CIP solves.
-   Rationale -- The rationale fleshes out the specification by describing what motivated the design and why particular design decisions were made. It should describe alternate designs that were considered and related work. The rationale should provide evidence of consensus within the community and discuss important objections or concerns raised during discussion.
-   Backwards compatibility -- All CIPs that introduce backwards incompatibilities must include a section describing these incompatibilities and their severity. The CIP must explain how the author proposes to deal with these incompatibilities.
-   Reference implementation -- The reference implementation must be completed before any CIP is given status “Final”, but it need not be completed before the CIP is accepted. It is better to finish the specification and rationale first and reach consensus on it before writing code. The final implementation must include test code and documentation appropriate for the Burstcoin protocol.


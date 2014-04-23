---
layout: post
title: Protocol Stack Complete (almost)
---

# {{ page.title }}

I created protocols:

- The Service Inform. protocol for examining the port numbers and public keys of a AMail Server.
- The PUSH protocol for forwarding mails to Servers.
- The PULL protocol for polling mails from Servers.
- The Key Offer Protocol for distributing Server Keys.
- The Key Poker Protocol for getting Remailer-Adresses/Keys while maintain Privacy.

With this, we can request random Remailer-Adresses and Keys, we can send a Mail to a destination and we can Pull Mails out of our inbox Server.

Whats Still to do:

- Create an AMail adress scheme (im keeping in mind something like "the.person#amail.provider.de" with "#" instead of "@")
- Create a protocol for resolving the public key and mailbox-name for an AMail adress.

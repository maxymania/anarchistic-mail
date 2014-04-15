---
layout: default
title: Mail Relay Message Format
---

# {{ page.title }}

## Background

A-Mail Servers Usually implement the [PUSH-Protocol](../proto/push) - A Protocol,
wich allows a client to push arbitrary blobs to a server - to take queued mails, eighter to resend them (Remailer) or to store them in the MBox of a User/A-Mail-Address.
These Blobs are Messages encrypted using [Seccon](../schemes/seccon), then they have an specific format to be processed by the Server.

<br/>

This document describes the format of the (unencrypted) Message format.

<br/>

## Overview

A Message in the MRMF (Mail Relay Message Format) consists of a header in an special format and an body. The body is a arbitrary blob. The header is a "\x00"-terminated String.

A Message can be described as the following Perl-Regexp: <code>/^([^\x00]+)\x00(.+)$/</code> Where $1 is the Header and $2 is the Body.

<br/>

## Header

The header of a "Mail Relay Message" describes, what to do with the body.

<br/>

### Remail the message

This header instructs the server to resend the Message to the server <code>"amail.undead.de"</code>:

<code>Remail:amail.undead.de</code>

<br/>

### Store the message

This header instructs the server to store the message in the mailbox <code>"a82b2c1"</code>:

<code>Mbox:a82b2c1</code>



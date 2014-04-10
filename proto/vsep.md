---
layout: default
title: Very Secure Encryption Protocol
---

# {{ page.title }}

## 1. Overview

VSEP uses the following Cryptographic Algorithms.

- Curve25519
- AES-128 (Rijndael with 128 Block and Key size)
- CTR-Mode (Block cipher mode)
- [ECFB-Mode](../schemes/ecbc) (Hybrid Block cipher mode)

This Document will use the following notation:

- RandomData = **ChooseRandom**(32 bytes) for random byte creation
- Secret = **Curve**(Public,Base) for elliptic curve Point Muliplication
- **Base** is the base Point of the elliptic curve
- Dest = **XOR**(Source1,Source2,...) XORing something

## 2. Handshake

![VSEP Client Server Handshake](vsep_cs_scheme.png "VSEP Client Server Handshake")

### 2.1 Server Public Key and Client Public Key

The Server and the Client MUST have a Key Pair. If one of the parties has
no Key Pair it SHALL generate one.

The key Pair is a Curve25519 Key Pair, defined as

PrivateKey = ChooseRandom(32 bytes)

PublicKey = Curve(PrivateKey,Base)


### 2.2 Encrypted Key Exchange

| 32 Bytes |
| --- |
| Session |
| EDhke1 |
| EDhke2 |
| EInitVec |



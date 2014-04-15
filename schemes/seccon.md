---
layout: default
title: Seccon - Secure Container
---

# {{ page.title }}

## 1. Overview

### 1.1 Description

Seccon is a specification for encrypting octet-streams using using Curve25519 as Public-Key function
together with AES-128.
The Public- and Private-Keys will be regular Curve25519 Public- and Private-Keys.

### 1.2 Document Overview

Seccon uses the following Cryptographic Algorithms.

- Curve25519
- AES-128 (Rijndael with 128 Block and Key size)
- CTR-Mode (Block cipher mode)
- [ECFB-Mode](ecbc) (Hybrid Block cipher mode)

This Document will use the following notation:

- RandomData = **ChooseRandom**(32 bytes) for random byte creation
- Secret = **Curve**(Public,Base) for elliptic curve Point Muliplication
- **Base** is the base Point of the elliptic curve
- SubKey1,SubKey2 = **Split**(Key) splits the 32-byte key into 16 Bytes Sub keys (1..16,17..32)

<br/>

## 2. Format/Structure

Byte 1...32: Curve25519-SessionPublicKey

Byte 32...N: Encrypted Octet-Stream

<br/>

## 3. Cipher Initialisation

Encryption Init:<br/>

<code>t = ChooseRandom(32 bytes)</code>

<code>Header = Curve(t,Base)</code>

<code>Key1,Key2 = Split(Curve(t,PublicKey))</code>

<code>StreamCipher = Setup-AES-ECFB(Key2,Key2,Setup-AES-CTR(Key1,Key1))</code>

<br/>

Decryption Init:<br/>

<code>Key1,Key2 = Split(Curve(PrivateKey,Header))</code>

<code>StreamCipher = Setup-AES-ECFB(Key2,Key2,Setup-AES-CTR(Key1,Key1))</code>

<br/>

## 4. Security

This encryption specification uses the Key as IV for the Encryption. Since the Key is random, this is not a serious security risk.


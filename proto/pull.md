---
layout: default
title: PULL Protocol
---

# {{ page.title }}

## 1. Overview

The PULL Protocol is a Protocol to be used to Fetch the mails from the Server.

<br/>

## 2. The Protocol itself

Every Line is Terminated using CR-LF

<br/>

### 2.1 The GUK Command

The Client sends the ID and gets an encrypted Private Key. The client is authorized, it will be able to decrypt it.

<code>C: GUK</code>

<code>C: USRID 4TdHc1i5YiskH0PopIzqmw==</code>

<code>S: Hs05UweVs3Wbh+gSWDpZzObdo5tRSV2eLid/95RknQc=</code>

<br/>

### 2.2 The LOGIN Command

The Client sends the ID and gets an encrypted session key. The client is authorized, it will be able to decrypt it.
Then, a md5-hash of the decrypted key is sent back to the server.

<code>C: LOGIN</code>

<code>C: USRID 4TdHc1i5YiskH0PopIzqmw==</code>

<code>S: q7Ou2wKpBJOgL5GFct/i+Qlx7pOP72W0Ylcwc8JbngI=</code>

<code>C: 6AruRuJS0AXknafMihtD3g==</code>

<br/>

The server will respond with:

<code>S: OK</code>

or with:

<code>S: FAILED</code>

<br/>

### 2.3 The INDEXUP Command

<code>C: INDEXUP</code>

<code>S: CzIVH8NMxHlfELIlf5kjvQ==</code>

<code>S: BQGpx17Ch3Z523PH5iEY1w==</code>

<code>S: KbUy+oTX/HgzlgTaRNdqiQ==</code>

<code>S: . </code>

<br/>

### 2.4 The PULL Command

The Pull command will try to pull an octet-stream from the server.

The octet-stream will be divided into small blocks and each of this blocks will be send in base64 encoded form as a line.

The typical block-size is 60 Byte per block, wich will result in 80 Base64-Characters.

If the octet-stream is being downloaded the chunk-stream will be terminated using a <code>"."</code>-Line.

<br/>

The Pull command will carry a Parameter wich is the ID of the Object.

The command looks like this:

<code>C: PULL CzIVH8NMxHlfELIlf5kjvQ==</code>

<br/>

The first possible response is (if the server has the requested object):

<code>S: PUSH</code>

<code>S: thw4xgMKhiFfxBkv+tQ2gWVdhBygXYeHxfHDTfXG/lWSk+wZkLlI+OPAHYk1rJB8AiV3169YqmA2Odvc</code>

<code>S: nIKzNSjzvl5wFBdsQnRhYF+mIkjxk3hFfjcu9vd8jykhpKEWfAjj7O5+EbhTUEuK7yh9cLvCZhCmOStm</code>

<code>S: 142P9uCqZB5EkhMIAxXiuaVpBw==</code>

<code>S: .</code>

<br/>

The second possible response is (if the server hasn't the request object):

<code>S: NOPE</code>

<br/>

### 2.4 The PURGE Command

The Purge Command marks an object as deleted.

<code>C: PURGE CzIVH8NMxHlfELIlf5kjvQ==</code>

(No Response from the Server side!)

<br/>

### 2.5 The FLUSH Command

Flushes all Purge Commands

<br/>

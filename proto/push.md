---
layout: default
title: PUSH Protocol
---

# {{ page.title }}

## 1. Overview

The PUSH Protocol is a Protocol, that knows only one Command: PUSH

<br/>

## 2. The Protocol itself

Every Line is Terminated using CR-LF

<br/>

### 2.1 The PUSH Command

The Push Command pushes an octet-stream to the server. The server will not send a response.

The octet-stream will be divided into small blocks and each of this blocks will be send in base64 encoded form as a line.

The typical block-size is 60 Byte per block, wich will result in 80 Base64-Characters.

The Request will be terminated using a <code>"."</code>-Line.

<code>C: PUSH</code>

<code>C: thw4xgMKhiFfxBkv+tQ2gWVdhBygXYeHxfHDTfXG/lWSk+wZkLlI+OPAHYk1rJB8AiV3169YqmA2Odvc</code>

<code>C: nIKzNSjzvl5wFBdsQnRhYF+mIkjxk3hFfjcu9vd8jykhpKEWfAjj7O5+EbhTUEuK7yh9cLvCZhCmOStm</code>

<code>C: 142P9uCqZB5EkhMIAxXiuaVpBw==</code>

<code>C: .</code>

(No Response from the Server side!)


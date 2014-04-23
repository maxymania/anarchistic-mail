---
layout: default
title: Key Poker Protocol
---

# {{ page.title }}

## 1. Overview

The Key Poker Protocol is a Protocol that implements the Mental Poker Cryptographic Protocol.

## 1.1 Encryption algorithms

The Key Poker Protocol uses a commutative cipher, the Pohlig-Hellman exponentiation cipher.

The prime number that from the [2048-bit MODP Group](http://tools.ietf.org/html/rfc3526#section-3):

<code>FFFFFFFF FFFFFFFF C90FDAA2 2168C234 C4C6628B 80DC1CD1</code>

<code>29024E08 8A67CC74 020BBEA6 3B139B22 514A0879 8E3404DD</code>

<code>EF9519B3 CD3A431B 302B0A6D F25F1437 4FE1356D 6D51C245</code>

<code>E485B576 625E7EC6 F44C42E9 A637ED6B 0BFF5CB6 F406B7ED</code>

<code>EE386BFB 5A899FA5 AE9F2411 7C4B1FE6 49286651 ECE45B3D</code>

<code>C2007CB8 A163BF05 98DA4836 1C55D39A 69163FA8 FD24CF5F</code>

<code>83655D23 DCA3AD96 1C62F356 208552BB 9ED52907 7096966D</code>

<code>670C354E 4ABC9804 F1746C08 CA18217C 32905E46 2E36CE3B</code>

<code>E39E772C 180E8603 9B2783A2 EC07A28F B5C55DF0 6F4C52C9</code>

<code>DE2BCBF6 95581718 3995497C EA956AE5 15D22618 98FA0510</code>

<code>15728E5A 8AACAA68 FFFFFFFF FFFFFFFF</code>

<br/>

See also:

- [The Pohlig-Hellman Exponentiation Cipher as a Bridge](http://de.slideshare.net/joshuarbholden/the-pohlighellman-exponentiation-cipher-as-a-bridge-between-classical-and-modern-cryptography)
- [More Modular Exponential (MODP) Diffie-Hellman groups for Internet Key Exchange (IKE)](http://tools.ietf.org/html/rfc3526) )

<br/>

## 2. The Protocol itself

Every Line is Terminated using CR-LF

<br/>

### 2.1 The POKER Process (Format)

The Poker Command Initiates the Poker protocol itself.

The typical line-size is 60 Byte per line, wich will result in 80 Base64-Characters.

<code>C: POKER ?host=~.*\\.x64x\\.mensa\\.de</code>

<code>S: CARDS S</code>

<code>S: thw4xgMKhiFfxBkv+tQ2gWVdhBygXYeHxfHDTfXG/lWSk+wZkLlI+OPAHYk1rJB8AiV3169YqmA2Odvc</code>

<code>S: nIKzNSjzvl5wFBdsQnRhYF+mIkjxk3hFfjcu9vd8jykhpKEWfAjj7O5+EbhTUEuK7yh9cLvCZhCmOStm</code>

<code>S: ,</code>

<code>S: /TTjzCYfXy/Naqh5FfX96OnPZKuGwcHwkraYly1MDuTdf0kxqclD5LD5yrM2VgtiNZ+d2pLf7OhUTQmf</code>

<code>S: ei3ZQyMILd0qrHypeDVrrksrTMQVJzq0K+JBIAc2EQaYdrF16iWyYRP7LDrY7Fq6d8tQv53WI+XLP+oM</code>

<code>S: ,</code>

<code>S: WhbEozfZJLrKpTvxX7ce7Lc8iuHXCAfryhJlP+XoIe4SJxCQ8azmUoHhbD1gmYdt8Z696Pu12i0UbSkN</code>

<code>S: dodgeXYVxguyy8HnQgCf43drlo3+MgFUH8KvdNtQH/hTqB+lAl3BVpwR88Xie8bJQEpmSIW3hTPp/3V3</code>

<code>S: ,</code>

<code>S: mGeaqx98unr936SXbq7kPuw5OpYqrAphP9I4TuR3rgtiUcGepCOgdHCte9z+et6IJo8Epbrh/R/sV2Zq</code>

<code>S: u+RnKMFUBQKzxO8hktERGVmeOSX5IHIL/1QSCpr69QRCJB5tCwiCD7Le1Be62ZSGqY0vtUM2yDPpBZZG</code>

<code>S: .</code>

<code>C: CARDS S,C</code>

<code>C: E7Jmr4CJDJ76cd0DKZhcli9ZZSH7zfNofUnRJHHm9s7iRd13OY8yMkpNxARU4JxXlkxuECTG9OjnYdE7</code>

<code>C: isAqh6LS67tQuduJW4wyL5Yk/FZclZ++fE/9wt07nglTkLye7jiB1nLykSIPyRwUxFkSWM296f0rNMip</code>

<code>C: ,</code>

<code>C: lTlFpK8Raw7JnawDg2XJ5Qb0p/yC8ASYlE9UBIrEtsOyo8/337diq8QBabty7sB8VeWhg6u21yA45/Rh</code>

<code>C: IBgJP28KSCj3TnuXFcwHRGHRnadJht64upEgeLP8D08zvQ8ZNje/rmerPJPbk2P/oOHF0dPW4b7pw/Jz</code>

<code>C: .</code>

<code>S: CARDS C</code>

<code>S: a8Fw1AIHHUW4irLE7CugW4Q0Q5hEZ6KQng4Jp+WJavYZG0yaO6bqW4Y3kkCIsdcTHwilq6oiQZOa4Sfc</code>

<code>S: tbIfLGrb+7y/qxMYtvmhYB9DOWPDiUDcPdhENvVO/ZeBUxFnovnZdgk67prvmN8U4K//4g/iK+L6NYnk</code>

<code>S: ,</code>

<code>S: oK9UhQgw17rU+9/LxPemMrE/mdzfNXvpumILo+o9FU78x0eXjr0YxZ5ff68qsfR3bNrYdqv1yy/4DuCc</code>

<code>S: ND0SO+NTquUHnZAqa8EWlKrI5VAyMNke2Xz9uUABUdjIkrqTTrwn/eBICxOkNOHjKcT33jKhXcAs9F01</code>

<code>S: .</code>

<br/>

### 2.2 The Item format

An unencrypted item is an JSON encoded Tuple, The nonces are used to fill up
the Items to the max. size ( Nonce, Host, Key, Nonce ).

The unencrypted Items look like this:

{% highlight javascript %}
["121232432432432423423434342354637347",
"362784673432783467243.m38228742.x64x.amail.x55x.mensa.de",
"0GbF6T9TzqcblXg2VzCiOr+P+RZ6mzD4rNklaYVZ3CM=",
"782683786387847RZ6mzD4rNklaYVZ3CM867678"]
{% endhighlight %}

<br/>

### 2.3 The Internal protocol

The Client querys a set of Items from the server. (it will pick few)

- The Server picks an encryption key A and uses this to encrypt each item of the resultset.
- The Server shuffles the items.
- The Server passes the encrypted and shuffled deck to the Client.
- The Client removes a lot of items from the deck.
- The Client picks an encryption key B and uses this to encrypt each item of the deck of the encrypted and shuffled deck.
- The Client shuffles the deck.
- The Client passes the double encrypted and shuffled deck back to the Server.
- The Server decrypts each item using his key A. This still leaves the Client's encryption in place though so she cannot know which item is which.
- The Server passes the deck to the Client.
- The Client decrypts each item using his key B.

The Client now has some Items.

The Server can't figure out wich Items the Client has.

The Client can't controll wich Items to pick.

<br/>


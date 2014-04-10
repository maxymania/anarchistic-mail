---
layout: default
title: Service Information Protocol
---

# {{ page.title }}

## 1. Service Information Protocol Database

The database shall be a Service JSON file that looks like the following.

{% highlight javascript %}
{
 "key1": { "this-is": "value-object1" },
 "key2": { "this-is": "value-object2" },
 "key3": { "this-is": "value-object3" },
 "key4": { "this-is": "value-object4" }
}
{% endhighlight %}

The Top Level object must be an JSON-hash. Every key has an assigned Value.

## 2. The Protocol itself

- The Server MUST listen for UDP packets on port 64000.
- Every client-to-server packet MUST be a JSON encoded string.
- Every server-to-client packet MUST be a JSON encoded object or null.
- The Server MUST answer every query, by sending a JSON value.

### 2.1 A Query

Client to Server
{% highlight javascript %}
"key1"
{% endhighlight %}

Server to Client, if the Server has "key1"
{% highlight javascript %}
{ "this-is": "value-object1" }
{% endhighlight %}

Server to Client, if the Server hasn't "key1"
{% highlight javascript %}
null
{% endhighlight %}


# Basic auth

* part of http spec
* send username and pass with every request
* `authorisation header` & value `Basic` and after space the data
* data format: `username:password` and encode base64
    * base64 - put any binary data into printable

```
curl -u myuser:secretPassword localhost:8080 
Authorization: Basic bXl1c2VyOnNlY3JldFBhc3N3b3Jk
```
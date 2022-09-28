# HMAC

* `Bearer tokens` added to http spec with oauth2
* `oauth2` - standardized way to log (log with google etc.)
* `HMAC` - hashbased message authentication code - legit way to create a bearer token. Uses cryptographic signing
    * `cryptographic signing` - way to validate that this was created by us
    * we can store some value in a cookie and add a hmac sign. like that:
    ```
    cookie(val|hmacSign)
    ```
# Some terminology

* `hashing` - one way algorithm
    * MD5, SHA
    * Bcrypt, Scrypt

* `signing`- ensure something is authentic
    * symmetric - same key to sing (encrypt)/verify (decrypt)
        * HMAC
    * Asymmetric (public for verify/decrypt, private for sign/encrypt)
        * RSA (also encryption alg), ECDSA

* JWT - transport method used for authentication. Encrypted data + sign. Stored in token/cookie. Allows to have stateless security (instead of session cookie, server side stored)

* `encryption`
    * caesar cipher (rot13)- for encryption/decryption
    * AES, RSA, DSA
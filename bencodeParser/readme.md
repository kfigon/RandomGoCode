## Bencode parser
serialization format uses in BitTorrent - https://en.wikipedia.org/wiki/Bencode

credits to https://xmonader.github.io/nimdays/day02_bencode.html

* strings and those are encoded with the string length followed by a colon and the string itself `length:string`, e.g yes will be encoded into `3:yes`
* ints those are encoded between i, e letters, e.g 59 will be encoded into `i59e`
* lists can contain any of the bencode types and it's encoded with l, e, e.g list of 1, 2 numbers is encoded into `li1ei2e` or with spaces for verbosity `l i1e i2e e`
* dicts are mapping from strings to any type and encoded between letters d, e, e.g name => hi and num => 3 is encoded into `d4:name2:hi3:numi3ee` or with spaces for verbosity `d 4:name 2:hi 3:num i3e e`
# Eclair PG run
Specs are the same as the LND test run.
Overall high throughput, but sometimes running into issues with the 2 nodes not being synced to bitcoind?
TP is comparable between SQLite / PG, probably because DO "disks" latency is comparable to PG latency.
Opening the database is also instant, notice that we needed to change the lease defaults for this though.
Restarting Eclair takes about 8 seconds from startup to the start of the HTTP server.

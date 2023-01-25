# LND PG run
Postgres config:
```
  name       = "alby-prod-postgres-cluster"
  engine     = "pg"
  version    = "12"
  size       = "db-s-2vcpu-4gb"
  region     = "fra1"
  node_count = 1
```
- TPS is much lower around 5, latency about 20s.
- Some pay commands actually fail (!):
	`2023-01-25 14:58:51.829alby-benchmark-lnd-2-686978cdc8-nz4ms [ERR] CRTR: Payment c9c143f7fd5fc20bde2b6161d199c4d05e85f3e6461a813c47ad9a87bfbbd618 failed: timeout`
- Opening the database is no longer an issue and happens instantly.
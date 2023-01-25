# LND bbolt load test
Running on bitcoind regtest on a DO cluster.
Cluster configuration:
```
  k8s version      = "1.24.8-do.0"
  node_pool {
    name       = "default"
    size       = "s-2vcpu-4gb"
    node_count = 3
  }
```
# Run 1
2 LND instances
```
processes: 100
channels: 10
channelCapacitySat: 100000000
requests:
   memory: 300Mi
   cpu: 800m
 limits:
   memory: 2Gi
```
- LND-1: db size 5.5GB, restart time with compaction 22m
- LND-1: db size 5.5GB, restart time no compaction 14m
- LND-2: db size 1.6GB, restart time no compaction 4m

After restarting LND-1 a second time (without db compaction), all but 2 channels were force-closed automatically, and the 2 remaining channels were inactive. Search the logs for `[INF] CNCT` to look at the ChannelArbitrator logs (=the part of LND that handles closes).
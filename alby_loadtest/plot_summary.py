import pandas as pd
from matplotlib import pyplot as plt
filenames = ["eclair_pg", "eclair_sqlite", "lnd_bbolt", "lnd_pg"]

plt.rcParams["figure.figsize"] = [7.00, 3.50]
plt.rcParams["figure.autolayout"] = True
columns = ["tps", "nr_payments", "latency_sec"]
for filename in filenames:
    df = pd.read_csv(filename+"_result.csv", usecols=columns)
    plt.plot(df.nr_payments, df.tps)
plt.legend(filenames)
plt.title("Throughput summary")
plt.xlabel("nr. of payments made")
plt.ylabel("transactions per second")
plt.savefig("summary_tps.png")

plt.clf()
for filename in filenames:
    df = pd.read_csv(filename+"_result.csv", usecols=columns)
    plt.plot(df.nr_payments, df.latency_sec)
plt.legend(filenames)
plt.title("Latency summary")
plt.xlabel("nr. of payments made")
plt.ylabel("latency(seconds)")
plt.savefig("summary_latency.png")


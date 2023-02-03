import pandas as pd
from matplotlib import pyplot as plt
filename = "lnd_pg"
titleprefix = "LND(PG)"

plt.rcParams["figure.figsize"] = [7.00, 3.50]
plt.rcParams["figure.autolayout"] = True
columns = ["tps", "nr_payments", "latency_sec"]
df = pd.read_csv(filename+"_result.csv", usecols=columns)
plt.plot(df.nr_payments, df.tps)
plt.title(titleprefix+": Tx's per second")
plt.xlabel("nr. of payments made")
plt.ylabel("transactions per second")
plt.savefig(filename+".png")

plt.clf()
plt.plot(df.nr_payments, df.latency_sec)
plt.title(titleprefix+": latency")
plt.xlabel("nr. of payments made")
plt.ylabel("latency(seconds)")
plt.savefig(filename+"_latency.png")


import pandas as pd
from matplotlib import pyplot as plt
plt.rcParams["figure.figsize"] = [7.00, 3.50]
plt.rcParams["figure.autolayout"] = True
columns = ["tps", "nr_payments", "latency_sec"]
df = pd.read_csv("lnd_bbolt_result.csv", usecols=columns)
plt.plot(df.nr_payments, df.tps)
plt.title("LND(bbolt): Tx's per second")
plt.xlabel("nr. of payments made")
plt.ylabel("transactions per second")
plt.savefig("lnd_bbolt.png")
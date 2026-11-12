import csv
import matplotlib.pyplot as plt

n, nested, dp = [], [], []

with open("ops.csv") as f:
    for row in csv.DictReader(f):
        n.append(int(row["n"]))
        nested.append(float(row["ops_nested"]))
        dp.append(float(row["ops_dp"]))

plt.plot(n, nested, label="Вкладені цикли, O(n²)", color="#7a0019")
plt.plot(n, dp, label="Динамічне програмування, O(n)", color="#1f4e9c")
plt.xlabel("n")
plt.ylabel("Кількість операцій")
plt.xticks(n)
plt.grid(True, alpha=0.3)
plt.legend()
plt.savefig("graph.png", dpi=200, bbox_inches="tight")
plt.show()
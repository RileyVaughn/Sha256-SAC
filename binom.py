import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
from scipy.stats import binomtest
from scipy.stats import binom
from scipy.stats import norm
from scipy.stats import chisquare
import math


n=1_000_000
p=0.5
alpha = .1
alpha_per_test = alpha / n
alpha_per_tail = alpha_per_test / 2


# b=binom.std(n,p)/n

# print(b)

z_critical = norm.ppf(1 - alpha_per_tail)
print(z_critical)
import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
from scipy.stats import binomtest
from scipy.stats import chisquare
import math



def binomial(filename):

    df = pd.read_csv('./data/'+filename+'.csv', header=None)
    df = df.to_numpy().flatten()

    n = 1000000
    p = .50
    ci = .99

    min = math.ceil(np.min(df) * n)
    max = math.ceil(np.max(df) * n)
    #mean = math.ceil(np.mean(df) * n)

    min_ci = binomtest(min,n,p).proportion_ci(confidence_level=ci, method='exact')
    max_ci = binomtest(max,n,p).proportion_ci(confidence_level=ci, method='exact')
    #mean_ci = binomtest(mean,n,p).proportion_ci(confidence_level=ci, method='exact')

    print(min_ci)
    print(max_ci)
    #print(mean_ci)
 


def binom64(dirName):

    n = 1000000
    p = .50
    ci = .99

    for i in range(64):
        df = pd.read_csv('./data/{}/round_{}.csv'.format(dirName,i+1), header=None)
        df = df.to_numpy().flatten()

        min = math.ceil(np.min(df) * n)
        max = math.ceil(np.max(df) * n)
        mean = math.ceil(np.mean(df) * n)
        
        min_ci = binomtest(min,n,p).proportion_ci(confidence_level=ci, method='exact')
        max_ci = binomtest(max,n,p).proportion_ci(confidence_level=ci, method='exact')
        mean_ci = binomtest(mean,n,p).proportion_ci(confidence_level=ci, method='exact')

        print(i+1,min_ci, max_ci)



def TableRow(dirName):

    n = 1000000
    p = .50
    ci = .99


    min_round, min_val = None, None
    max_round, max_val = None, None
    mean_round, mean_val = None, None


    for i in range(64):
        df = pd.read_csv('./data/{}/round_{}.csv'.format(dirName,i+1), header=None)
        df = df.to_numpy().flatten()

        min = math.ceil(np.min(df) * n)
        max = math.ceil(np.max(df) * n)
        mean = math.ceil(np.mean(df) * n)
        
        min_ci = binomtest(min,n,p).proportion_ci(confidence_level=ci, method='exact')
        max_ci = binomtest(max,n,p).proportion_ci(confidence_level=ci, method='exact')
        mean_ci = binomtest(mean,n,p).proportion_ci(confidence_level=ci, method='exact')

        if (min/n) > .495 and min_round == None:
            min_round = i+1
            min_val = min/n
            min_pm = 0
        if min_round != None and abs(min_val-(min/n)) > min_pm:
            min_pm = abs(min_val-(min/n))

        if (max/n) < .505 and max_round == None:
            max_round = i+1
            max_val = max/n
            max_pm = 0
        if max_round != None and abs(max_val-(max/n)) > max_pm:
            max_pm = abs(max_val-(max/n))

        if mean_ci.low < .5 and mean_ci.high > .5 and mean_round == None:
            mean_round = i+1
            mean_val = mean/n
            mean_low = mean_ci.low
            mean_high = mean_ci.high


    #print("Min: ",min_round, round(min_val*100,4), round(min_pm*100,4))
    #print("Max: ",max_round, round(max_val*100,4), round(max_pm*100,4))
    print("Mean: ",mean_round, round(mean_val*100,4), round(mean_low*100,4), round(mean_high*100,4), round(np.mean(df)*100,6))

#dirName = "rounds"
#filename = "fullCF"

#filename = "no_sched_fullCF"
#dirName = "no_sched"
#dirName = "no_choose"
#dirName = "no_sig1"

##binom64(dirName)
#binomial(filename)

#TableRow('rounds')
#TableRow('no_sched')
TableRow('no_choose')
#TableRow('no_major')
#TableRow('no_k')
#TableRow('no_sig0')
#TableRow('no_sig1')
#TableRow('xor')

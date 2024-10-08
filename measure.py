import pandas as pd
import matplotlib.pyplot as plt
import numpy as np


def ReadAndMeasure(filename):

    df = pd.read_csv('./data/'+filename+'.csv', header=None).transpose()

    max_sac = np.max([np.max(df[row]) for row in df])
    min_sac = np.min([np.min(df[row]) for row in df])
    mean_sac = np.mean([np.mean(df[row]) for row in df])

    print(max_sac)
    print(min_sac)
    print(mean_sac)


def ReadAndMeasure64(dirName):

    frames = []
    for i in range(64):
        frames.append( pd.read_csv('./data/{}/round_{}.csv'.format(dirName,i+1), header=None).transpose())


    min_sac = [np.min([np.min(df[row]) for row in df]) for df in frames]
    max_sac = [np.max([np.max(df[row]) for row in df]) for df in frames]
    mean_sac = [np.mean([np.mean(df[row]) for row in df]) for df in frames]

    print(list(zip(min_sac,[i+1 for i in range(64)])))
    print(list(zip(max_sac,[i+1 for i in range(64)])))
    print(list(zip(mean_sac,[i+1 for i in range(64)])))
    

def RM64(dirName):

    frames = []
    for i in range(64):
        frames.append( pd.read_csv('./data/{}/round_{}.csv'.format(dirName,i+1), header=None).transpose())


    min_sac = [np.min([np.min(df[row]) for row in df]) for df in frames]
    max_sac = [np.max([np.max(df[row]) for row in df]) for df in frames]
    mean_sac = [np.mean([np.mean(df[row]) for row in df]) for df in frames]

    print(min_sac)
    print(max_sac)
    print(mean_sac)
 




#ReadAndMeasure("no_choose_fullFC")
#ReadAndMeasure64("no_sig1")
    
ReadAndMeasure64("rounds")
#ReadAndMeasure64("no_sched")
#ReadAndMeasure64("no_choose")
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
 


def RAMIntersect2(fn1,fn2,fn12):

    df1 = pd.read_csv('./data/'+fn1+'.csv', header=None).transpose() - .5
    df2 = pd.read_csv('./data/'+fn2+'.csv', header=None).transpose() - .5
    df12 = (pd.read_csv('./data/'+fn12+'.csv', header=None).transpose() -.5).abs()

    df_comb = (df1+df2).abs()

    max_comb = np.max([np.max(df_comb[row]) for row in df_comb])
    max_12 = np.max([np.max(df12[row]) for row in df12])

    print(max_comb)
    print(max_12)


def RAMI2Rounds(dn1,dn2,dn12):

    frames_comb = []
    frames12 = []
    for i in range(64):
        csv_from_1 = pd.read_csv('./data/{}/round_{}.csv'.format(dn1,i+1), header=None).transpose()
        csv_from_2 = pd.read_csv('./data/{}/round_{}.csv'.format(dn2,i+1), header=None).transpose()
        ((csv_from_1-.5)+(csv_from_2-.5)).abs()

        frames_comb.append(((csv_from_1-.5)+(csv_from_2-.5)).abs())
        frames12.append((pd.read_csv('./data/{}/round_{}.csv'.format(dn12,i+1), header=None).transpose()-.5).abs())


    
    max_comb = [np.max([np.max(df[row]) for row in df]) for df in frames_comb]
    max_12 = [np.max([np.max(df[row]) for row in df]) for df in frames12]


    #print(list(zip(max_comb, max_12,[i+1 for i in range(64)])))
    print([
        (round(float(a), 4), round(float(b), 4), i)
        for a, b, i in zip(max_comb, max_12, [i+1 for i in range(64)])
    ])



    



RAMI2Rounds('removed_1/C','removed_1/M','removed_2/CM')
import pandas as pd
import matplotlib.pyplot as plt
import numpy as np


def ReadAndPlotFull(filename):
        
    df = pd.read_csv('./data/'+filename+'.csv', header=None).transpose()
    
    min_sac = [np.min(df[row]) for row in df]
    max_sac = [np.max(df[row]) for row in df]
    
    plt.plot(max_sac)
    plt.plot(min_sac)
    plt.xlabel("Complimented Bit Index")
    plt.ylabel("(%) Change of Hash Output Bit")
    plt.title("SAC of SHA256 Compression Function w/ Message Schedular")
    plt.xticks(range(0,513,64))
    plt.yticks(np.arange(.49,.51,.005))
    plt.legend(["Max Change", "Min Change"])
    plt.savefig("./data/plots/"+ filename+'.png')
    plt.close()


def ReadAndPlot64(dirName):

    frames = []
    for i in range(64):
        frames.append( pd.read_csv('./data/{}/round_{}.csv'.format(dirName,i+1), header=None).transpose())


    min_sac = [np.min([np.min(df[row]) for row in df]) for df in frames]
    max_sac = [np.max([np.max(df[row]) for row in df]) for df in frames]
    mean_sac = [np.mean([np.mean(df[row]) for row in df]) for df in frames]

    plt.plot(max_sac)
    plt.plot(mean_sac)
    plt.plot(min_sac)
    plt.xlabel("Round #")
    plt.ylabel("(%) Change of Hash Output Bit")
    plt.title("SAC of SHA256 Compression Function Rounds")
    plt.xticks(range(0,65,8))
    plt.yticks(np.arange(0,1.01,.25))
    plt.legend(["Max Change","Mean Change","Min Change"])
    plt.savefig("./data/plots/{}.png".format(dirName))
    plt.close()

#ReadAndPlotFull('fullCF')
ReadAndPlot64("rounds")
#ReadAndPlot64("no_sched")

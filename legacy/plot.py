import pandas as pd
import matplotlib.pyplot as plt
import numpy as np


def ReadAndPlotFull(filename):
        
    df = pd.read_csv('./data/'+filename+'.csv', header=None).transpose()
    
    min_sac = [np.min(df[row])*100 for row in df]
    max_sac = [np.max(df[row])*100 for row in df]
    
    plt.scatter([y for y in range(1,513)], max_sac,s=1)
    plt.scatter([y for y in range(1,513)], min_sac, s=1)
    plt.xlabel("Complimented Bit Index")
    plt.ylabel("(%) Change of Hash Output Bit")
    plt.title("SAC of SHA256 Compression Function w/o Message Schedular")
    plt.xticks(range(0,513,64))
    plt.yticks(np.arange(49,51.1,.5))
    plt.legend(["Max Change", "Min Change"])
    plt.tight_layout()
    plt.savefig("./data/plots/"+ filename+'.eps', format="eps")
    plt.savefig("./data/plots/"+ filename+'.png')
    plt.close()


def ReadAndPlot64(dirName):

    frames = []
    for i in range(64):
        frames.append( pd.read_csv('./data/{}/round_{}.csv'.format(dirName,i+1), header=None).transpose())


    min_sac = [np.min([np.min(df[row])*100 for row in df]) for df in frames]
    max_sac = [np.max([np.max(df[row])*100 for row in df]) for df in frames]
    mean_sac = [np.mean([np.mean(df[row])*100 for row in df]) for df in frames]

    plt.plot(max_sac)
    plt.plot(min_sac)
    plt.plot(mean_sac)

    plt.xlabel("Round #")
    plt.ylabel("(%) Change of Hash Output Bit")
    plt.title("SAC of SHA256 Compression Function Rounds w/o Message Schedular")
    plt.xticks(range(0,65,8))
    plt.yticks(np.arange(0.0,100.1,25.0))
    plt.legend(["Max Change","Min Change","Mean Change"])
    plt.tight_layout()
    plt.savefig("./data/plots/{}.eps".format(dirName), format="eps")
    plt.savefig("./data/plots/{}.png".format(dirName))
    plt.close()


def ReadAndPlot64Black(dirName):

    frames = []
    for i in range(64):
        frames.append( pd.read_csv('./data/{}/round_{}.csv'.format(dirName,i+1), header=None).transpose())

    frames2 = []
    for i in range(64):
        frames2.append( pd.read_csv('./data/{}/round_{}.csv'.format("rounds",i+1), header=None).transpose())


    min_sac = [np.min([np.min(df[row]) for row in df]) for df in frames]
    max_sac = [np.max([np.max(df[row]) for row in df]) for df in frames]
    mean_sac = [np.mean([np.mean(df[row]) for row in df]) for df in frames]

    min_sac2 = [np.min([np.min(df[row]) for row in df]) for df in frames2]
    max_sac2 = [np.max([np.max(df[row]) for row in df]) for df in frames2]
    mean_sac2 = [np.mean([np.mean(df[row]) for row in df]) for df in frames2]

    
    plt.plot(max_sac2, color='black', label="")
    plt.plot(min_sac2, color='black', label="")
    plt.plot(mean_sac2, color='black', label="Default")
    plt.plot(max_sac, label="Max Change")
    plt.plot(min_sac, label="Min Change")
    plt.plot(mean_sac, label="Mean Change")

    plt.xlabel("Round #")
    plt.ylabel("(%) Change of Hash Output Bit")
    plt.title("SAC of SHA256 Compression Function Rounds w/ Message Schedular")
    plt.xticks(range(0,65,8))
    plt.yticks(np.arange(0,1.01,.25))
    plt.legend()
    plt.tight_layout()
    plt.savefig("./data/plots/{}_black.eps".format(dirName), format="eps")
    plt.savefig("./data/plots/{}_black.png".format(dirName))
    plt.close()

def ReadAndPlotSub():

    dir_names = ["no_choose","no_k","no_major","no_sched","no_sig0", "no_sig1","xor","rounds"]
    means = []
    mins = []
    maxs = []

    for dirName in dir_names:
        temp = []
        for j in range(64):
            temp.append( pd.read_csv('./data/{}/round_{}.csv'.format(dirName,j+1), header=None).transpose())
        mean_sac = [np.mean([np.mean(df[row])*100 for row in df]) for df in temp]
        min_sac = [np.min([np.min(df[row])*100 for row in df]) for df in temp]
        max_sac = [np.max([np.max(df[row])*100 for row in df]) for df in temp]
        means.append(mean_sac)
        mins.append(min_sac)
        maxs.append(max_sac)
        
    
    colors = ['tab:blue','tab:orange','tab:green','tab:olive','tab:pink','tab:cyan','tab:purple','k']
    labels = ['Choose Removed', 'K Removed', 'Majority Removed', 'Schedule Removed', 'Sigma0 Removed', 'Sigma1 Removed', 'Integer Add Removed', 'Normal']
    for i in range(8):
        plt.plot(maxs[i], color=colors[i], linestyle='dashed', label = '')
        plt.plot(mins[i], color=colors[i],linestyle='dashed', label = '')
        plt.plot(means[i], color=colors[i],linestyle='solid', label = labels[i])


        

    plt.xlabel("Round #")
    plt.ylabel("(%) Change of Hash Output Bit")
    plt.title("SAC of Compression function w/ Sub-functions Removed")
    plt.xticks(range(0,65,8))
    plt.yticks(np.arange(0,100.1,25.0))
    plt.legend()
    plt.tight_layout()
    plt.savefig("./data/plots/all_subs.eps", format="eps")
    plt.savefig("./data/plots/all_subs.png")
    plt.close()

#ReadAndPlotFull('fullCF')
#ReadAndPlot64("rounds")
ReadAndPlot64("no_sched")
ReadAndPlotFull("no_sched_fullCF")
#ReadAndPlotSub()
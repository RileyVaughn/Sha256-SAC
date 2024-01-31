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
    plt.ylabel("(%) Change for Hash Output Bit")
    plt.title("SAC of SHA256 Compression Function w/ Message Schedular")
    plt.xticks(range(0,513,64))
   #plt.yticks(np.arange(.49,.51,.5))
    plt.savefig("./data/plots/"+ filename+'.png')


    


ReadAndPlotFull('fullCF')

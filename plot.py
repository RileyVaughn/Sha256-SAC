import pandas as pd
import matplotlib.pyplot as plt


def ReadAndPlot(filename):
        
    df = pd.read_csv('./data/'+filename+'.csv', header=None)
    

    df = df.transpose()
    ax = df.plot(kind="box",showfliers=False,xlabel="Round #",ylabel="% Bits Flipped",legend=False,figsize=(16,8))
    vals = ax.get_yticks()
    ax.set_yticklabels(['{:,.2%}'.format(x) for x in vals])

    plt.savefig("./data/plots/"+ filename+'.png')


    


ReadAndPlot('H_Normal')
ReadAndPlot('H_XOR')

ReadAndPlot('R_ALL')
ReadAndPlot('R_ALL-XOR')
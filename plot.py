import pandas as pd
import matplotlib.pyplot as plt


def ReadAndPlot(filename):
        
    df = pd.read_csv('./data/'+filename+'.csv')


    df.plot(x=0, y=1,figsize=(8,8), kind='line',xticks=[x for x in range(65) if x %8 == 0],yticks = [x for x in range(257) if x %16 == 0],xlabel="Round #",ylabel="Bits Changed",legend=False)
    plt.plot([0,64],[128,128],linestyle='dotted',)
    #plt.show()
    plt.savefig("./data/plots/"+ filename+'.png')



ReadAndPlot('H_normal')
ReadAndPlot('ZERO_normal')
ReadAndPlot('Random_normal')
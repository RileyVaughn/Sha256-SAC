import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
from pathlib import Path
import seaborn as sns


def PlotCombination(path):

    df = pd.read_csv(path,header=None)

    # plt.plot(df[0])
    # plt.plot(df[1])

    plt.plot(df[0], label='SAC-mean Values',linewidth=1,alpha=.5)
    plt.plot(df[1], label='SAC Values',linewidth=1,alpha=.75)


    plt.xlabel("Compression Round")
    plt.ylabel("Output Bit Complement Ratio")
    plt.title("SAC of SHA-256 Compression Function")

    plt.xticks(range(0,65,8))
    plt.yticks(np.arange(0.0,1,.25))
    plt.legend()
    plt.tight_layout()

    plt.savefig(Path('data') / 'plots' / f'{path.parts[-2]}.pdf', format='pdf')
    plt.close()




def PlotAllAtLevelMean(base,names):

    num_colors = len([d for d in base.iterdir()])
    colors = sns.color_palette(n_colors= num_colors)

    for i, d in enumerate(base.rglob("*")):
        if d.is_dir() and not any(p.is_dir() for p in d.iterdir()):
            file = d / '_stats.csv'
            df = pd.read_csv(file,header=None)
            df.index = df.index + 1
            plt.plot(df[0],color=colors[i], label=names[d.parts[-1]],linewidth=1,alpha=.5)

    plt.xlabel("Compression Round")
    plt.ylabel("Output Bit Complement Ratio")
    plt.title("SAC-mean values of Individual SHA-256 Compression Sub-functions")

    plt.xticks(range(0,65,8))
    plt.yticks(np.arange(0.0,1,.25))
    plt.legend()
    plt.tight_layout()

    plt.savefig(Path('data') / 'plots' / f'{base.parts[-1]}_mean.pdf', format='pdf')
    plt.close()


def PlotAllAtLevelAbs(base,names):

    num_colors = len([d for d in base.iterdir()])
    colors = sns.color_palette(n_colors= num_colors)

    for i, d in enumerate(base.rglob("*")):
        if d.is_dir() and not any(p.is_dir() for p in d.iterdir()):
            file = d / '_stats.csv'
            df = pd.read_csv(file,header=None)
            df.index = df.index + 1
            plt.plot(df[1],color=colors[i], label=names[d.parts[-1]],linewidth=1, linestyle='dashed',alpha=.5)

    plt.xlabel("Compression Round")
    plt.ylabel("Output Bit Complement Ratio")
    plt.title("SAC Values of Individual SHA-256 Compression Sub-functions")

    plt.xticks(range(0,65,8))
    plt.yticks(np.arange(0.0,1,.25))
    plt.legend()
    plt.tight_layout()

    plt.savefig(Path('data') / 'plots' / f'{base.parts[-1]}_worst.pdf', format='pdf')
    plt.close()



def CombineData():

    jitter = 0

    df0 = pd.read_csv(Path('data') / 'removed_0'  / '_rounds.csv',header=None)
    df0[2] = [7]*len(df0)
    df1 = pd.read_csv(Path('data') / 'removed_1'  / '_rounds.csv',header=None)
    df1[2] = [6]*len(df1) + np.random.uniform(-jitter,jitter,size=len(df1[1]))
    df2 = pd.read_csv(Path('data') / 'removed_2'  / '_rounds.csv',header=None)
    df2[2] = [5]*len(df2) + np.random.uniform(-jitter,jitter,size=len(df2[1]))
    df3 = pd.read_csv(Path('data') / 'removed_3'  / '_rounds.csv',header=None)
    df3[2] = [4]*len(df3) + np.random.uniform(-jitter,jitter,size=len(df3[1]))
    df4 = pd.read_csv(Path('data') / 'removed_4'  / '_rounds.csv',header=None)
    df4[2] = [3]*len(df4) + np.random.uniform(-jitter,jitter,size=len(df4[1]))
    df5 = pd.read_csv(Path('data') / 'removed_5'  / '_rounds.csv',header=None)
    df5[2] = [2]*len(df5) + np.random.uniform(-jitter,jitter,size=len(df5[1]))
    df6 = pd.read_csv(Path('data') / 'removed_6'  / '_rounds.csv',header=None)
    df6[2] = [1]*len(df6) + np.random.uniform(-jitter,jitter,size=len(df6[1]))

    df = pd.concat([df0,df1,df2,df3,df4,df5,df6],ignore_index=True)
    df.loc[0, 0] = 'KS0RS1MXC'
    

    for i, row in enumerate(df.itertuples()):
        if pd.isna(row[2]):
            df.loc[i,1] = 66

    # print(df)
    return df


def PlotThreshByLevel():

    df1 = pd.read_csv(Path('data') / 'removed_1'  / '_rounds.csv',header=None)
    df2 = pd.read_csv(Path('data') / 'removed_2'  / '_rounds.csv',header=None)
    # df3 = pd.read_csv(Path('data') / 'removed_3'  / '_rounds.csv',header=None)
    df4 = pd.read_csv(Path('data') / 'removed_4'  / '_rounds.csv',header=None)
    df5 = pd.read_csv(Path('data') / 'removed_5'  / '_rounds.csv',header=None)
    df6 = pd.read_csv(Path('data') / 'removed_6'  / '_rounds.csv',header=None)

    plt.plot([58],[7],marker='o',linestyle='',color=plt.cm.tab20.colors[0])
    plt.plot(df1[1],[6]*len(df1[1]),marker='o',linestyle='',color=plt.cm.tab20.colors[1])
    plt.plot(df2[1],[5]*len(df2[1]),marker='o',linestyle='',color=plt.cm.tab20.colors[2])
    # plt.plot(df3[1],[4]*len(df3[1]),marker='o',linestyle='',color=plt.cm.tab20.colors[3])
    plt.plot(df4[1],[3]*len(df4[1]),marker='o',linestyle='',color=plt.cm.tab20.colors[4])
    plt.plot(df5[1],[2]*len(df5[1]),marker='o',linestyle='',color=plt.cm.tab20.colors[5])
    plt.plot(df6[1],[1]*len(df6[1]),marker='o',linestyle='',color=plt.cm.tab20.colors[6])



    plt.xticks(range(0,65,8))
    plt.yticks([0,1,2,3,4,5,6,7])

    plt.xlabel("Compression Round")
    plt.ylabel("Number of Sub-functions")
    plt.title("SAC Threshold of SHA-256 Sub-function Combinations")

    plt.tight_layout()


    plt.savefig(Path('data') / 'plots' / 'thresh.pdf', format='pdf')


def  SubplotThreshByLevel():

    ms = 5


    df = CombineData()
    labels = ['M','C','X','K','R','S0','S1']
    names = ["Function: Majority","Function: Choose","Function: Integer Addition","Function: K Constants", "Function: Message Scheduler","Function: Σ0","Function: Σ1",'All Functions Aggregated']
    color_order = [5,6,3,4,2,1,0]
 
    colors = sns.color_palette(n_colors= 7)

    fig, axs = plt.subplots(4, 2, figsize=(10, 10), sharex=True, sharey=True)
    axs_flat = axs.flatten()  

    for i, ax in enumerate(axs_flat):
        
        if i < 7:
            color_data = []
            grey_data = []
            for j, name in enumerate(df[0]):
                
                if labels[i] in name:
                    color_data.append((df.loc[j,1],df.loc[j,2]))
                else:
                    grey_data.append((df.loc[j,1],df.loc[j,2]))

            x_color, y_color = zip(*color_data)
            x_grey, y_grey = zip(*grey_data)
        
            ax.plot(x_color,y_color,marker='o',linestyle='',color=colors[color_order[i]],markersize=ms,alpha=.25)
        else:
            ax.plot(df[1],df[2],marker='o',linestyle='',color='black',markersize=ms,alpha=.25)
            

        ax.axvline(x=65, color='black',linewidth='1')
        ax.set_title(names[i])
        xticks = list(range(0,65,8))
        xticks[0] = 1
        
        ax.set_xticks(xticks)
        ax.set_xlim(0,67)
        ax.set_yticks([1,2,3,4,5,6,7])

    # axs_flat[7].set_visible(False)



    


    fig.supxlabel("Compression Round")
    fig.supylabel("Number of Sub-functions")
    fig.suptitle("SAC Threshold of SHA-256 Sub-function Combinations")

    fig.tight_layout()


    fig.savefig(Path('data') / 'plots' / 'thresh.pdf', format='pdf')




# name_dict = {   "CKXRS0S1":"Majority",
#                 "CMKRS0S1":"Integer Addition",
#                 "CMKXRS0":"Σ1",
#                 "CMKXRS1":"Σ0",
#                 "CMKXS0S1":"Schedule",
#                 "CMXRS0S1":"K Function",
#                 "MKXRS0S1":"Choose"}
# PlotAllAtLevelMean(Path('data') / 'removed_6',name_dict)
# PlotAllAtLevelAbs(Path('data') / 'removed_6',name_dict)
# PlotCombination(Path('data') / 'removed_0' / '_stats.csv')
# RenameComb("CKXRS1")
# PlotThreshByLevel()
SubplotThreshByLevel()






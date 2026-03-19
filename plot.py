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
            plt.plot(df[0],color=plt.cm.tab20.colors[i], label=names[d.parts[-1]],linewidth=1,alpha=.5)

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
            plt.plot(df[1],color=plt.cm.tab20.colors[i], label=names[d.parts[-1]],linewidth=1, linestyle='dashed',alpha=.5)

    plt.xlabel("Compression Round")
    plt.ylabel("Output Bit Complement Ratio")
    plt.title("SAC Values of Individual SHA-256 Compression Sub-functions")

    plt.xticks(range(0,65,8))
    plt.yticks(np.arange(0.0,1,.25))
    plt.legend()
    plt.tight_layout()

    plt.savefig(Path('data') / 'plots' / f'{base.parts[-1]}_worst.pdf', format='pdf')
    plt.close()







# name_dict = {   "CKXRS0S1":"Majority",
#                 "CMKRS0S1":"Integer Addition",
#                 "CMKXRS0":"Σ1",
#                 "CMKXRS1":"Σ0",
#                 "CMKXS0S1":"Schedule",
#                 "CMXRS0S1":"K Function",
#                 "MKXRS0S1":"Choose"}
# PlotAllAtLevelMean(Path('data') / 'removed_6',name_dict)
# PlotAllAtLevelAbs(Path('data') / 'removed_6',name_dict)
PlotCombination(Path('data') / 'removed_0' / '_stats.csv')
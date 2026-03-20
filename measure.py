import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
from pathlib import Path



def WhenThreshold(path):

    df = pd.read_csv(path / '_stats.csv',header=None)
    df.index = df.index+1
    THRESHOLD = .49734

    plat_round = None
    plat_flag = False

    for row in df.itertuples():
        if row[2] > THRESHOLD and not plat_flag:
            plat_round = int(row[0])
            plat_flag = True
            #print(plat_round)
        if row[2] < THRESHOLD:
            plat_flag = False
            
    return (path.parts[-1],plat_round)


def RenameComb(name):

    labels = {'M','C','X','K','R','S0','S1'}
    in_name = []
                    

    for label in labels:
        if label in name:
            in_name.append(label)
        
    return labels-set(in_name)

def AllAtLevelWH(base):

    df = pd.DataFrame()
    names = []
    rounds = []
    
    for i, d in enumerate(base.rglob("*")):
        if d.is_dir() and not any(p.is_dir() for p in d.iterdir()):
            when = WhenThreshold(d)
            names.append("".join(RenameComb(when[0])))
            rounds.append(when[1])
    


    df['name'] = names
    df['round'] = rounds
    df.to_csv(base /'_rounds.csv',index=False,header=False)




def MeasureIntersect(path1, path2, pathex):

    df1 = pd.read_csv(path1 / '_stats.csv',header=None)
    df2 = pd.read_csv(path2/ '_stats.csv',header=None)
    df_actual = pd.read_csv(path2/ '_stats.csv',header=None)

    expected_df = df1+df2-(2*df1*df2)

    print(df_actual,expected_df)



# AllAtLevelWH(Path('data') / 'removed_1')
# AllAtLevelWH(Path('data') / 'removed_2')
# # AllAtLevelWH(Path('data') / 'removed_3')
# AllAtLevelWH(Path('data') / 'removed_4')
# AllAtLevelWH(Path('data') / 'removed_5')
# AllAtLevelWH(Path('data') / 'removed_6')

r0 = list(WhenThreshold(Path('data') / 'removed_0'))
r0[0] = ''
df = pd.DataFrame()
df['name'] = [r0[0]]
df['round'] = [r0[1]]
df.to_csv(Path('data')/'removed_0' /'_rounds.csv',index=False,header=False)


# majority = Path('data') / 'removed_6' / 'CKXRS0S1'
# choose = Path('data') / 'removed_6' / 'MKXRS0S1'
# mc = Path('data') / 'removed_5' / 'KXRS0S1'

# MeasureIntersect(majority,choose,mc)




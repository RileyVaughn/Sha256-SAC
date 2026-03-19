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
            plat_round = row[0]
            plat_flag = True
            print(plat_round)
        if row[2] < THRESHOLD:
            if plat_flag:
                print("Fail: " + plat_round)
            plat_flag = False
            plat_round = None



def MeasureIntersect(path1, path2):

    df1 = pd.read_csv(path1 / '_stats.csv',header=None)
    df2 = pd.read_csv(path2/ '_stats.csv',header=None)

    expected_df = df1+df2-(2*df1*df2)

    print(expected_df)



#WhenThreshold(Path('data') / 'removed_6' / 'CMKXS0S1')

majority = Path('data') / 'removed_6' / 'CKXRS0S1'
choose = Path('data') / 'removed_6' / 'MKXRS0S1'

MeasureIntersect(majority,choose)




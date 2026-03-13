import pandas as pd
import numpy as np
from pathlib import Path



def collapse_64rounds(base):

    mean_sacs = []
    abs_sacs = []


    for i in range(64):

        round_df = pd.read_csv(base / f'round_{i+1}.csv', header=None).transpose()

        max_sac = np.max([np.max(round_df[row]) for row in round_df])
        min_sac = np.min([np.min(round_df[row]) for row in round_df])
        mean_sac = np.mean([np.mean(round_df[row]) for row in round_df])

        if abs(max_sac-.5) > abs(min_sac-.5):
            abs_sac = abs(max_sac-.5)
        else:
            abs_sac = abs(min_sac-.5)

        mean_sacs.append(mean_sac)
        abs_sacs.append(abs_sac)

    collapsed_df = pd.DataFrame()
    collapsed_df['mean'] = mean_sacs
    collapsed_df['abs'] = abs_sacs
    collapsed_df.index = collapsed_df.index + 1

    collapsed_df.to_csv( base /'_stats.csv',index=False,header=False)


def collapse_all_at_level(base):
   
    for d in base.rglob("*"):
        if d.is_dir() and not any(p.is_dir() for p in d.iterdir()):
            collapse_64rounds(d)






# collapse_all_at_level(Path('data') / 'removed_6')
collapse_all_at_level(Path('data') / 'removed_1')
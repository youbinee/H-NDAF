__author__ = "YOUBIN JEON"
__copyright__ = "YOUBIN JEON 2022"
__version__ = "1.0.0"
__license__ = "MNC lab"

import math
import numpy as np
import pandas as pd
from sklearn.preprocessing import MinMaxScaler

class DataLoader():
   
    def __init__(self, filename, split, cols):
        self.dataframe = pd.read_csv(filename, index_col=["run_num","seq_num"])

        self.data_train = self.dataframe.get(cols).values[:split]
        self.data_test  = self.dataframe.get(cols).values[split:]
        self.len_train  = len(self.data_train)
        self.len_test   = len(self.data_test)
        self.len_train_windows = None

        self.train_scaler = MinMaxScaler(feature_range = (0, 1))
        self.data_train = self.train_scaler.fit_transform(self.data_train)
        self.test_scaler = MinMaxScaler(feature_range = (0, 1))
        self.data_test = self.test_scaler.fit_transform(self.data_test)

    def get_test_data(self, seq_len, normalise):
        data_x = []
        data_y = []

        for i in range(self.len_test - seq_len):
            x, y = self._next_window(i, seq_len, normalise, train=False)
            data_x.append(x)
            data_y.append(y)

        return np.array(data_x), np.array(data_y)

    def generate_train_batch(self, seq_len, batch_size, normalise):
        i = 0
        while True:
            x_batch = []
            y_batch = []
            for b in range(batch_size):
                x, y = self._next_window(i, seq_len, normalise)
                x_batch.append(x)
                y_batch.append(y)
                i += 1
                
                if i == (self.len_train - seq_len):
                    i = 0
                    yield np.array(x_batch), np.array(y_batch)
            yield np.array(x_batch), np.array(y_batch)

    def _next_window(self, i, seq_len, normalise, train=True):
        if train:
            window = self.data_train[i:i+seq_len+1]
        else:
            window = self.data_test[i:i+seq_len+1]

        x = window[:-1]
        y = window[-1, [0]]

        return x, y

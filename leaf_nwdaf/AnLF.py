__author__ = "YOUBIN JEON"
__copyright__ = "YOUBIN JEON 2022"
__version__ = "1.0.0"
__license__ = "MNC lab"

import os
import json
import time
import math
import matplotlib
import matplotlib.pyplot as plt
from core.data_processor import DataLoader
from core.model import Model

def AnLF():
    print("========= [Root-Learning-AnLF INFERENCE] START AnLF Inference =========")

    model = Model()
    if os.path.isfile('/saved_models/throughput_prediction.h5'):
        model.load_model('./saved_models/throughput_prediction.h5')

    configs = json.load(open('config.json', 'r'))
    if not os.path.exists(configs['model']['save_dir']): os.makedirs(configs['model']['save_dir'])

    data = DataLoader(
        os.path.join('data', configs['data']['filename']),
        configs['data']['train_test_split'],
        configs['data']['columns']
    )
    # 3. Model Test
    x_test, y_test = data.get_test_data(
        seq_len = configs['data']['sequence_length'],
        normalise = configs['data']['normalise']
    )
    predictions = model.predict_sequence_full(x_test, y_test, configs['data']['sequence_length'])

if __name__ == '__main__':
    AnLF()

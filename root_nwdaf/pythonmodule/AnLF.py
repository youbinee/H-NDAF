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

def plot_results(predicted_data, true_data):
    fig = plt.figure(facecolor = 'white')
    ax = fig.add_subplot(111)
    ax.plot(true_data, label = 'True data')
    plt.plot(predicted_data, label = 'Prediction data')
    plt.xlabel('time step')
    plt.ylabel('throughput (Mbps)')
    plt.legend()
    plt.show()


def AnLF():
    print("========= [Root-Learning-AnLF INFERENCE] START AnLF Inference =========")

    if os.path.isfile('/saved_models/throughput_prediction.h5'):
        model = load_model('/saved_models/throughput_prediction.h5')

    configs = json.load(open('config.json', 'r'))
    if not os.path.exists(configs['model']['save_dir']): os.makedirs(configs['model']['save_dir'])

    x_test, y_test = data.get_test_data(
        seq_len = configs['data']['sequence_length'],
        normalise = configs['data']['normalise']
    )

    predictions = model.predict_sequence_full(x_test, y_test, configs['data']['sequence_length'])
    print("========= [Root-Learning-AnLF INFERENCE] END AnLF Inference =========")

    p_results = plot_results(predictions, y_test)

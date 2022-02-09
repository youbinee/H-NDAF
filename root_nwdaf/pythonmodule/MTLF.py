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

def MTLF():
    print("========= [Root-Learning-MTLF TRAINING] START MTLF training =========")
    
    if os.path.isfile('./saved_models/throughput_prediction.h5'):
        print("========= [Root-Learning-MTLF TRAINING] Use existing model =========")

    else:
        print("========= [Root-Learning-MTLF TRAINING] New training model")

        configs = json.load(open('config.json', 'r'))
        if not os.path.exists(configs['model']['save_dir']): os.makedirs(configs['model']['save_dir'])

        data = DataLoader(
            os.path.join('data', configs['data']['filename']),
            configs['data']['train_test_split'],
            configs['data']['columns']
        )

        model = Model()
        model.build_model(configs)

        steps_per_epoch = math.ceil((data.len_train - configs['data']['sequence_length']) / configs['training']['batch_size'])
        
        model.train_generator(
            data_gen = data.generate_train_batch(
                seq_len = configs['data']['sequence_length'],
                batch_size = configs['training']['batch_size'],
                normalise = configs['data']['normalise']
            ),
            epochs = configs['training']['epochs'],
            batch_size = configs['training']['batch_size'],
            steps_per_epoch = steps_per_epoch,
            save_dir = configs['model']['save_dir'],
        )
        
        model.save_model('./saved_models/throughput_prediction.h5')

    print("========= [Root-Learning-MTLF TRAINING] END MTLF trainig =========")

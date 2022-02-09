__author__ = "YOUBIN JEON"
__copyright__ = "YOUBIN JEON 2022"
__version__ = "1.0.0"
__license__ = "MNC lab"

import os
import math
import numpy as np
import datetime as dt
import tensorflow as tf
import pandas as pd
from numpy import newaxis
from core.utils import Timer
from keras.layers import Dense, Activation, Dropout, LSTM, CuDNNLSTM, Bidirectional, TimeDistributed, RepeatVector
from keras.models import Sequential, load_model
from keras.callbacks import EarlyStopping, ModelCheckpoint, TensorBoard
from keras.optimizers import Adam, RMSprop

class Model():

    def __init__(self):
        self.model = Sequential()

    def build_model(self, configs):
        for layer in configs['model']['layers']:
            neurons = layer['neurons'] if 'neurons' in layer else None
            dropout_rate = layer['rate'] if 'rate' in layer else None
            activation = layer['activation'] if 'activation' in layer else None
            return_seq = layer['return_seq'] if 'return_seq' in layer else None
            input_timesteps = layer['input_timesteps'] if 'input_timesteps' in layer else None
            input_dim = layer['input_dim'] if 'input_dim' in layer else None
            repeat_num = layer['repeat_num'] if 'repeat_num' in layer else None

            if layer['type'] == 'dense':
                with tf.name_scope('layer_Dense'):
                    self.model.add(Dense(neurons, activation=activation))
            if layer['type'] == 'lstm':
                with tf.name_scope('layer_LSTM'):
                    self.model.add(CuDNNLSTM(neurons, input_shape=(input_timesteps, input_dim), return_sequences=return_seq))
        
        self.model.compile(loss=configs['model']['loss'], optimizer=configs['model']['optimizer'])

    def train_generator(self, data_gen, epochs, batch_size, steps_per_epoch, save_dir):
        save_fname = os.path.join(save_dir, '%s-e%s.h5' % (dt.datetime.now().strftime('%d%m%Y-%H%M%S'), str(epochs)))
        callbacks = [
                ModelCheckpoint(filepath=save_fname, monitor='loss')
        ]

        hist = self.model.fit_generator(
            data_gen,
            steps_per_epoch = steps_per_epoch,
            epochs = epochs,
            callbacks = callbacks,
            workers = 1
        )
        self.model.save(save_fname)

    def predict_sequence_full(self, x_data, y_data, window_size):
        print('[Model] Predicting Sequences Full...')

        curr_frame = x_data[0]
        predicted = []

        for i in range(len(x_data)):
            predicted.append(self.model.predict(curr_frame[newaxis,:,:])[0,0])
            curr_frame = curr_frame[1:]
            curr_frame = np.insert(curr_frame, [window_size-2], predicted[-1], axis=0)

        return predicted

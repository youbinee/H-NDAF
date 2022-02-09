__author__ = "YOUBIN JEON"
__copyright__ = "YOUBIN JEON 2022"
__version__ = "1.0.0"
__license__ = "MNC lab"

from flask import Flask, request
import json
# import h5py
from MTLF import *
from AnLF import *
from Model import *

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def parser():
    data = {}
    if request.method == 'POST':
        data = request.json
        print('====== [Root-Learning-Main] START Root Learning Module')

        if str(data['nfService']) == 'mtlf-training':
            MTLF()
            data["data"] = "/saved_models/throughput_prediction.h5"

        elif str(data['nfService']) == 'anlf-inference':
            AnLF()
            data["data"] = "[Root-Learning-MAIN] inference finish == "

        else:
            data['data'] = "None (Wrong)"

        data['reqNFInstanceID'] = '[Root-Learning-MAIN] ' + data['reqNFInstanceID'] + ' == '
        data['nfService'] = '[Root-Learning-MAIN] ' + data['nfService'] + ' == '
        data['reqTime'] = '[Root-Learning-Main] ' + data['reqTime']

    return json.dumps(data)

if __name__ == '__main__':
    app.run(host='127.0.0.38', port=5005)

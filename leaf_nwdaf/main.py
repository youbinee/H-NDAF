from flask import Flask, request
import json
from AnLF import *

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def parser():
    data = {}
    if request.method == 'POST':
        data = request.json
        print('====== [Leaf-Learning-amf] START: %s' % data)

        AnLF()
        data["data"] = "[Leaf-Learning-amf] inference finish == "
        print('====== [Leaf-Learning-amf] END ======')

    return json.dumps(data)

if __name__ == '__main__':
    app.run(host='127.0.0.18', port=5005) 

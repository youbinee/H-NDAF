__author__ = "YOUBIN JEON"
__copyright__ = "YOUBIN JEON 2022"
__version__ = "1.0.0"
__license__ = "MNC lab"

import datetime as dt

class Timer():
    def __init__(self):
        self.start_dt = None

    def start(self):
        self.start_dt = dt.datetime.now()

    def stop(self):
        end_dt = dt.datetime.now()
        print('Time taken: %s' % (end_dt - self.start_dt))

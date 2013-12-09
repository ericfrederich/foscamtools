#!/usr/bin/env python

from PyQt4.QtCore import *
from PyQt4.QtGui import *

from ui_timelapse import Ui_MainWindow

class MainWindow(QMainWindow):
    def __init__(self, parent=None):
        super(MainWindow, self).__init__(parent)

        # Set up the user interface from Designer.
        self.ui = Ui_MainWindow()
        self.ui.setupUi(self)

        self.setWindowTitle("Timelapse Calculator")
        self.do_calc()

    @pyqtSlot(float)
    def on_sleepSpin_valueChanged(self, v):
        self.do_calc()

    @pyqtSlot(float)
    def on_intervalSpin_valueChanged(self, v):
        self.do_calc()

    @pyqtSlot(int)
    def on_framesSpin_valueChanged(self, v):
        self.do_calc()

    def do_calc(self):
        
        # hours of sleep to seconds of sleep
        seconds = self.ui.sleepSpin.value() * 60 * 60
        
        # how many pictures will be taken?
        n_pictures = seconds / self.ui.intervalSpin.value()

        # those pictures at X fps will yield how many minutes of video?
        length = n_pictures / self.ui.framesSpin.value()
        minutes, seconds = divmod(length, 60)

        self.ui.picsEdit.setText("%d" % n_pictures)

        self.ui.lengthEdit.setText("%d:%02d" % (int(minutes), int(seconds)))

if __name__ == '__main__':
    import sys
    app = QApplication(sys.argv)
    mw = MainWindow()
    mw.show()
    sys.exit(app.exec_())

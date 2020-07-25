package log

import (
	logrus "github.com/sirupsen/logrus"
)

//Error 自己封装一层
func Error(args ...interface{}) {
	logrus.Error(args...)
}

package logger


import ( 
	"github.com/sirupsen/logrus"
	"sync"
)

type Logger = *logrus.Logger
var logger Logger
var once sync.Once

func GetInstance() Logger {
	
	once.Do(func ()  {
		logger = logrus.New()
	})

	return logger
}


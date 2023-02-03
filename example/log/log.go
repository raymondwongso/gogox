package main

import (
	"fmt"

	"github.com/raymondwongso/gogox/log"
	gogox_logrus "github.com/raymondwongso/gogox/log/logrus"
	"github.com/raymondwongso/gogox/log/nop"
	gogox_zap "github.com/raymondwongso/gogox/log/zap"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type service struct {
	logger  log.Logger
	logger2 log.Logger
	logger3 log.Logger
}

func (s service) DoSomething() error {
	// do your business etc2.
	err := fmt.Errorf("some error come up")
	if err != nil {
		s.logger.Error(err.Error())
		s.logger2.Errorw(err.Error(), log.Metadata{"error": err.Error(), "user_id": 1})
		s.logger3.Error(err.Error())
		return err
	}

	return nil
}

func main() {
	// NopLogger do nothing`
	nopLogger := &nop.Logger{}

	// Logrus logger
	ll := logrus.New()
	ll.SetLevel(logrus.ErrorLevel)
	ll.SetFormatter(&logrus.JSONFormatter{})
	logrusLogger := gogox_logrus.New(ll, log.Metadata{"service": "api", "version": "1.2.1"})

	// zap logger
	zz, _ := zap.NewProduction()
	zapLogger := gogox_zap.New(zz, log.Metadata{"service": "api_zap"})

	service := &service{
		logger:  nopLogger,
		logger2: logrusLogger,
		logger3: zapLogger,
	}

	service.DoSomething()
}

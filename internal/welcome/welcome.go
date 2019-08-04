package welcome

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Service struct {
	logger        *logrus.Logger
	welcomeString string
}

func (ws Service) HTTPWelcome(w http.ResponseWriter, r *http.Request) {
	ws.logger.Infof("%s - %s %s", r.Host, r.Method, r.RequestURI)
	ws.logger.Debugf("%v", r)
	_, _ = fmt.Fprintf(w, ws.welcomeString)
}

func NewService(l *logrus.Logger, wel string) *Service {
	return &Service{l, wel}
}

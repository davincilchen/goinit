package ctxcache

import (
	"context"

	"github.com/sirupsen/logrus"
)

// =========================================== //
//最大集合
type ContextLogger struct {
	context.Context
	tag string
}

func (t *ContextLogger) CacheDBError(err error) {

	if err == nil {
		return
	}
	logrus.Errorf("[%s] DBError: %s\n", t.tag, err.Error())
}

func (t *ContextLogger) GetDBError() error {
	return nil
}

func (t *ContextLogger) CacheHttpError(err error) {
	if err == nil {
		return
	}
	logrus.Errorf("[%s] HttpError: %s", t.tag, err.Error())
}

func (t *ContextLogger) GetHttpError() error {
	return nil
}

// =========================================== //

func NewContextLogger(tag string) *ContextLogger {

	return &ContextLogger{
		Context: context.Background(),
		tag:     tag,
	}

}

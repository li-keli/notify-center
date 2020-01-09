package track_log

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v3"
	"net/http"
	"notify-center/server/api/v1/vo"
)

const (
	esHost  = "172.16.7.20"
	esUrl   = "http://172.16.7.20:9200"
	esIndex = "mylog"
)

func UseLogMiddle(ctx *gin.Context) {
	Logger(ctx)
	ctx.Next()
	ctx.JSON(http.StatusOK, vo.BaseOutput{}.Success(""))
}

// 获取日志对象
func Logger(ctx *gin.Context) *logrus.Entry {
	if log, b := ctx.Get("log"); b {
		return log.(*logrus.Entry)
	} else {
		return EsLog(esHost, esUrl, esIndex).WithField("trackId", uuid.NewV4().String())
	}
}

func EsLog(esHost, urls, index string) *logrus.Logger {
	log := logrus.New()
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(urls))
	if err != nil {
		log.Panicf("Elastic not reachable\n %s", err.Error())
	}
	hook, err := elogrus.NewAsyncElasticHook(client, esHost, logrus.InfoLevel, index)
	if err != nil {
		log.Panic(err)
	}
	log.Hooks.Add(hook)
	log.Infof("GinLogConfig ring in %s - %s", esHost, index)
	return log
}

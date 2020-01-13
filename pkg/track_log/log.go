package track_log

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v3"
	"notify-center/pkg/constant"
)

const (
	esHost  = "172.16.14.52"
	esUrl   = "http://172.16.14.52:9200"
	esIndex = "notify"
)

var esLog *logrus.Logger

func init() {
	esLog = EsLog(esHost, esUrl, esIndex)
}

func UseLogMiddle(ctx *gin.Context) {
	Logger(ctx)
	ctx.Next()
}

// 获取日志对象
func Logger(ctx *gin.Context) *logrus.Entry {
	if log, b := ctx.Get("log"); b {
		return log.(*logrus.Entry)
	} else {
		return esLog.WithFields(logrus.Fields{
			"mode":            constant.ProductionMode,
			"trackId":         uuid.NewV4().String(),
			"x-forwarded-for": ctx.GetHeader("x-forwarded-for"),
		})
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
	log.Infof("Es日志接入 %s - %s", esHost, index)
	return log
}

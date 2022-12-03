package boot

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_ZhiHu/app/global"
	"go_ZhiHu/app/router"
	"net/http"
)

func ServerSetup() {
	config := global.Config.Server

	gin.SetMode(config.Mode)
	routers := router.InitRouter()
	server := &http.Server{
		Addr:              config.Addr(),
		Handler:           routers,
		TLSConfig:         nil,
		ReadTimeout:       config.GetReadTimeout(),
		ReadHeaderTimeout: 0,
		WriteTimeout:      config.GetWriteTimeout(),
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20,
	}
	global.Logger.Info("initialize server successfully!", zap.String("port", config.Addr()))
	global.Logger.Error(server.ListenAndServe().Error())
}

package global

import (
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"go_ZhiHu/app/internal/model/config"
	"gorm.io/gorm"
)

var (
	Config  *config.Config
	Logger  *zap.Logger
	MysqlDB *gorm.DB
	Rdb     *redis.Client
)

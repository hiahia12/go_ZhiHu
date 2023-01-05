package global

import (
	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go_ZhiHu/app/internal/model/config"
)

var (
	Config  *config.Config
	Logger  *zap.Logger
	MysqlDB *sqlx.DB
	Rdb     *redis.Client
)

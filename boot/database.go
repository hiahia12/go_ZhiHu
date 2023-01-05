package boot

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go_ZhiHu/app/global"
	"time"
)

func MysqlSetup() {
	config := global.Config.Database.Mysql

	db, err := sqlx.Open("mysql", config.GetDsn())
	if err != nil {
		global.Logger.Fatal("initialize mysql failed.", zap.Error(err))
	}
	sqlDB := db
	sqlDB.SetConnMaxIdleTime(global.Config.Database.Mysql.GetConnMaxIdleTime())
	sqlDB.SetConnMaxLifetime(global.Config.Database.Mysql.GetConnMaxLifeTime())
	sqlDB.SetMaxIdleConns(global.Config.Database.Mysql.MaxIdleconns)
	sqlDB.SetMaxOpenConns(global.Config.Database.Mysql.MaxOpenConns)
	err = sqlDB.Ping()
	if err != nil {
		global.Logger.Fatal("connect to mysql db failed.", zap.Error(err))
	}

	sql := `CREATE TABLE IF NOT EXISTS answer_subject (
		id bigint NOT NULL AUTO_INCREMENT ,
		answer varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL ,
		writerid bigint DEFAULT NULL ,
		questionid bigint DEFAULT NULL ,
		likenumber int DEFAULT NULL ,
		comment_number int DEFAULT NULL ,
		creat_time datetime DEFAULT CURRENT_TIMESTAMP ,
		update_time datetime DEFAULT CURRENT_TIMESTAMP ,
		PRIMARY KEY (id)
	) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci`
	_, err = db.Exec(sql)
	if err != nil {
		global.Logger.Fatal("create table failed.", zap.Error(err))
	}

	sql = `	CREATE TABLE IF NOT EXISTS collect_subject (
		createrid bigint NOT NULL COMMENT '创建者或收藏者id',
		collect_name varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '收藏夹名字',
		answerid bigint DEFAULT NULL COMMENT '被收藏的回答的id',
		PRIMARY KEY (createrid,collect_name)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci`
	_, err = db.Exec(sql)
	if err != nil {
		global.Logger.Fatal("create table failed.", zap.Error(err))
	}

	sql = `CREATE TABLE IF NOT EXISTS  comment_subject (
		id bigint NOT NULL AUTO_INCREMENT  ,
		comment varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL  ,
		likenumber int DEFAULT NULL  ,
		answeri bigint DEFAULT NULL ,
		writerid bigint DEFAULT NULL  ,
		creat_time datetime DEFAULT CURRENT_TIMESTAMP ,
		update_time datetime DEFAULT CURRENT_TIMESTAMP ,
		PRIMARY KEY (id)
	) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci`
	_, err = db.Exec(sql)
	if err != nil {
		global.Logger.Fatal("create table failed.", zap.Error(err))
	}

	sql = `CREATE TABLE IF NOT EXISTS question_subject (
		id bigint NOT NULL AUTO_INCREMENT COMMENT '主键问题id',
		question varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '问题内容',
		askerid bigint NOT NULL COMMENT '提问者的id',
		answer_number int DEFAULT NULL COMMENT '回答数',
		creat_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
		update_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
		PRIMARY KEY (id)
	) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci`
	_, err = db.Exec(sql)
	if err != nil {
		global.Logger.Fatal("create table failed.", zap.Error(err))
	}

	sql = `CREATE TABLE IF NOT EXISTS user_subject (
		id bigint NOT NULL AUTO_INCREMENT COMMENT '用户id_主键',
		username varchar(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
		password varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户密码',
		creat_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
		update_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
		PRIMARY KEY (id)
	) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci`
	_, err = db.Exec(sql)
	if err != nil {
		global.Logger.Fatal("create table failed.", zap.Error(err))
	}

	sql = `CREATE TABLE IF NOT EXISTS favourite_subject (
		id bigint NOT NULL AUTO_INCREMENT COMMENT '收藏夹id',
		userid bigint NOT NULL COMMENT '使用该收藏夹的用户的id',
		favourite_number int(10) unsigned zerofill DEFAULT NULL COMMENT '收藏数',
		public int(10) unsigned zerofill DEFAULT NULL COMMENT '收藏夹是否公开',
		name varchar(90) COLLATE utf8mb4_general_ci NOT NULL COMMENT '收藏夹的名字',
		creat_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
		update_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
		PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci`
	_, err = db.Exec(sql)
	if err != nil {
		global.Logger.Fatal("create table failed.", zap.Error(err))
	}

	sql = `CREATE TABLE IF NOT EXISTS followfavourite_subject (
		id bigint NOT NULL COMMENT 'id',
		userid bigint DEFAULT NULL COMMENT '用户的id',
		favouriteid bigint DEFAULT NULL COMMENT '被关注的收藏夹的id',
		creat_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
		update_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
		PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci`
	_, err = db.Exec(sql)
	if err != nil {
		global.Logger.Fatal("create table failed.", zap.Error(err))
	}

	sql = `CREATE TABLE IF NOT EXISTS article_subject (
		id bigint NOT NULL AUTO_INCREMENT COMMENT '主键问题id',
		article varchar(10000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文章内容',
		writerid bigint NOT NULL COMMENT '写作者的id',
		answer_number int DEFAULT NULL COMMENT '回答数',
		like_number int DEFAULT NULL COMMENT '赞同数',
		creat_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        update_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
		PRIMARY KEY (id)
	) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci`
	_, err = db.Exec(sql)
	if err != nil {
		global.Logger.Fatal("create table failed.", zap.Error(err))
	}
	sql = `CREATE TABLE IF NOT EXISTS followquestion_subject (
		id bigint NOT NULL COMMENT 'id',
		userid bigint DEFAULT NULL COMMENT '关注者的id',
		questionid bigint DEFAULT NULL COMMENT '关注的问题的id',
		creat_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '建立时间',
		update_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
		PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci`
	_, err = db.Exec(sql)
	if err != nil {
		global.Logger.Fatal("create table failed.", zap.Error(err))
	}

	sql = `CREATE TABLE IF NOT EXISTS myfavouriteanswer_subject (
		id bigint NOT NULL COMMENT '本次收藏的id',
		answerid bigint NOT NULL COMMENT '收藏的回答的id',
		userid bigint NOT NULL COMMENT '收藏者的id',
		favouriteid bigint NOT NULL COMMENT '所属收藏夹的id',
		creat_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
		update_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
		PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci`
	_, err = db.Exec(sql)
	if err != nil {
		global.Logger.Fatal("create table failed.", zap.Error(err))
	}

	sql = `CREATE TABLE IF NOT EXISTS myfavouritearticle_subject (
		id bigint NOT NULL COMMENT '本次收藏的id',
		articleid bigint(20) unsigned zerofill NOT NULL COMMENT '收藏的文章的id',
		userid bigint(20) unsigned zerofill NOT NULL COMMENT '收藏者的id',
		favouriteid bigint NOT NULL COMMENT '所属收藏夹的id',
		creat_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
		update_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
		PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci`
	_, err = db.Exec(sql)
	if err != nil {
		global.Logger.Fatal("create table failed.", zap.Error(err))
	}

	sql = `CREATE TABLE IF NOT EXISTS followuser_subject (
		id bigint NOT NULL COMMENT 'id',
		userid bigint DEFAULT NULL COMMENT '用户的id',
		followuserid bigint DEFAULT NULL COMMENT '被关注的用户的id',
		creat_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
		update_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
		PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;`
	global.MysqlDB = db

	global.Logger.Info("initialize mysql successful")
}

func RedisSetup() {
	config := global.Config.Database.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       config.Db,
		PoolSize: config.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Fatal("connect to redis instance failed.", zap.Error(err))
	}

	global.Rdb = rdb

	global.Logger.Info("initialize redis client successfully!")
}

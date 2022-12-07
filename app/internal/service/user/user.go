package user

import (
	"context"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap"
	"go_ZhiHu/app/global"
	"go_ZhiHu/app/internal/model"
	"go_ZhiHu/utils/jwt"
	"golang.org/x/crypto/sha3"
	"gorm.io/gorm"
	"time"
)

type SUser struct{}

var insUser SUser = SUser{}

func (s *SUser) CheckUserIsExist(ctx context.Context, username string) error {
	userSubject := &model.UserSubject{}
	err := global.MysqlDB.WithContext(ctx).
		Table("user_subject").
		Select("username").
		Where("username = ?", username).
		First(userSubject).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "user_subject"),
			)
			return fmt.Errorf("internal err")
		}
	} else {
		return fmt.Errorf("username already exit")
	}
	return nil
}

func (s *SUser) EncryptPassword(password string) string {
	d := sha3.Sum224([]byte(password))
	return hex.EncodeToString(d[:])
}

func (s *SUser) CreateUser(ctx context.Context, userSubject *model.UserSubject) {
	global.MysqlDB.WithContext(ctx).Table("user_subject").Create(userSubject)
}

func (s *SUser) CheckPassword(ctx context.Context, userSubject *model.UserSubject) error {
	err := global.MysqlDB.WithContext(ctx).Table("user_subject").Where(&model.UserSubject{
		Username: userSubject.Username,
		Password: userSubject.Password,
	}).First(userSubject).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "user_subject"),
			)
			return fmt.Errorf("internal err")
		} else {
			return fmt.Errorf("invalid username or password")
		}
	}
	return nil
}

func (s *SUser) GenerateToken(ctx context.Context, userSubject *model.UserSubject) (string, error) {
	jwtConfig := global.Config.Middleware.Jwt
	j := jwt.NewJWT(&jwt.Config{
		SecretKey:   jwtConfig.SecretKey,
		ExpiresTime: jwtConfig.ExpiresTime,
		BufferTime:  jwtConfig.BufferTime,
		Issuer:      jwtConfig.Issuer,
	})
	claims := j.CreatClaims(&jwt.BaseClaims{
		Id:         userSubject.Id,
		Username:   userSubject.Username,
		CreateTime: userSubject.CreatTime,
		UpdateTime: userSubject.UpdateTime,
	})
	tokenString, err := j.GenerateToken(&claims)
	if err != nil {
		global.Logger.Error("generate token failed.", zap.Error(err))
		return "", fmt.Errorf("internal err")
	}
	err = global.Rdb.Set(ctx,
		fmt.Sprintf("jwt:%d", userSubject.Id),
		tokenString,
		time.Duration(jwtConfig.ExpiresTime)*time.Second).Err()
	if err != nil {
		global.Logger.Error("set redis cache failed.",
			zap.Error(err),
			zap.String("key", "jwt:[id]"),
			zap.Int64("id", userSubject.Id),
		)
		return "", fmt.Errorf("internal err")
	}
	return tokenString, nil
}

func (s *SUser) WriteQuestion(ctx context.Context, question *model.Question) {
	global.MysqlDB.WithContext(ctx).Table("question_subject").Create(question)
}

func (s *SUser) WriteAnswer(ctx context.Context, answer *model.AnswerSubject) {
	global.MysqlDB.WithContext(ctx).Table("answer_subject").Create(answer)
}

func (s *SUser) WriteComment(ctx context.Context, comment *model.Comment) {
	global.MysqlDB.WithContext(ctx).Table("comment_subject").Create(comment)
}

func (s *SUser) GetQuestions(ctx context.Context) []model.Question {
	questions := []model.Question{}
	global.MysqlDB.WithContext(ctx).Table("question_subject").Find(&questions)
	return questions
}

func (s *SUser) GetAnswer(ctx context.Context, questionid int64) []model.AnswerSubject {
	answers := []model.AnswerSubject{}
	global.MysqlDB.WithContext(ctx).Table("answer_subject").Where("questionid=?", questionid).Find(&answers)
	return answers
}

func (s *SUser) GetComment(ctx context.Context, answerid int64) []model.Question {
	questions := []model.Question{}
	global.MysqlDB.WithContext(ctx).Table("question_subject").Find(&questions)
	return questions
}

func (s *SUser) GetUser(ctx context.Context, username string) model.UserSubject {
	user := model.UserSubject{}
	global.MysqlDB.WithContext(ctx).Table("user_subject").Where("username=?", username).First(&user)
	return user
}

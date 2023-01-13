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
	usersubject := &model.UserSubject{}
	sql := "SELECT username FROM user_subject where username = ?"
	err := global.MysqlDB.Get(usersubject, sql, username)
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

func (s *SUser) CheckFavouriteExist(ctx context.Context, favouriteid int64) error {
	favourite := &model.FavouriteSubject{}
	sql := "SELECT id FROM favourite_subject where id = ?"
	err := global.MysqlDB.Get(favourite, sql, favouriteid)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "favourite_subject"),
			)
			return fmt.Errorf("internal err")
		}
	} else {
		return fmt.Errorf("favourite already exit")
	}
	return nil
}

func (s *SUser) CheckQuestionExist(ctx context.Context, questionid int64) error {
	favourite := &model.FavouriteSubject{}
	sql := "SELECT id FROM favourite_subject where id = ?"
	err := global.MysqlDB.Get(favourite, sql, questionid)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "question_subject"),
			)
			return fmt.Errorf("internal err")
		}
	} else {
		return fmt.Errorf("question already exit")
	}
	return nil
}

func (s *SUser) EncryptPassword(password string) string {
	d := sha3.Sum224([]byte(password))
	return hex.EncodeToString(d[:])
}

func (s *SUser) CreateUser(ctx context.Context, userSubject *model.UserSubject) {
	sql := "INSERT INTO user_subject(username,password) VALUES (?,?)"
	_, err := global.MysqlDB.Exec(sql, userSubject.Username, userSubject.Password)
	if err != nil {
		return
	}
}

func (s *SUser) CreatFavourites(ctx context.Context, favourites *model.FavouriteSubject) error {
	sql := "INSERT INTO user_subject(userid,favouritenumber,`public`,`name`) VALUES (?,?,?,?)"
	_, err := global.MysqlDB.Exec(sql, favourites.Userid, favourites.FavouriteNumber, favourites.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *SUser) CheckPassword(ctx context.Context, userSubject *model.UserSubject) error {
	user := &model.UserSubject{}
	sql := "SELECT username FROM user_subject where username = ? and password = ?"
	err := global.MysqlDB.Get(user, sql, userSubject.Username, userSubject.Password)
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

func (s *SUser) WriteArticle(ctx context.Context, article *model.ArticleSubject) {
	sql := "INSERT INTO article_subject(article,writerid,answer_number,like_number) VALUES (?,?,?,? )"
	_, err := global.MysqlDB.Exec(sql, article.Article, article.Writerid, 0, 0)
	if err != nil {
		return
	}
}

func (s *SUser) WriteQuestion(ctx context.Context, question *model.Question) {
	sql := "INSERT INTO question_subject(question,askerid) VALUES (?,?)"
	_, err := global.MysqlDB.Exec(sql, question.Question, question.Askerid)
	if err != nil {
		return
	}
}

func (s *SUser) WriteAnswer(ctx context.Context, answer *model.AnswerSubject) {
	question := model.Question{}
	sql := "INSERT INTO answer_subject(answer,writerid,questionid) VALUES (?,?,?)"
	_, err := global.MysqlDB.Exec(sql, answer.Answer, answer.Writerid, answer.Questionid)
	if err != nil {
		return
	}
	sqlStr := "SELECT id,question,answer_number,askerid,creat_time,update_time FROM question_subject where id = ?"
	err1 := global.MysqlDB.Get(&question, sqlStr, answer.Questionid)
	if err1 != nil {
		fmt.Print(err)
		return
	}
	sqlstr := "UPDATE question_subject set answer_number = ?where id = ?"
	_, err2 := global.MysqlDB.Exec(sqlstr, question.AnswerNumber+1, answer.Questionid)
	if err2 != nil {
		fmt.Print(err)
		return
	}
}

func (s *SUser) WriteComment(ctx context.Context, comment *model.Comment) {
	answer := model.AnswerSubject{}
	sql := "INSERT INTO comment_subject(comment,writerid,answerid) VALUES (?,?,?)"
	_, err := global.MysqlDB.Exec(sql, comment.Comment, comment.Writerid, comment.Answerid)
	if err != nil {
		fmt.Print(err)
		return
	}

	sqlStr := "SELECT id,answer,writerid,questionid,likenumber,creat_time,update_time FROM answer_subject where id = ?"
	err1 := global.MysqlDB.Get(&answer, sqlStr, comment.Answerid)
	if err1 != nil {
		fmt.Print(err)
		return
	}

	sqlstr := "UPDATE answer_subject set comment_number = ?where id = ?"
	_, err2 := global.MysqlDB.Exec(sqlstr, answer.CommentNumber+1, comment.Answerid)
	if err2 != nil {
		fmt.Print(err)
		return
	}
}

func (s *SUser) GetQuestions(ctx context.Context) []model.Question {
	questions := []model.Question{}
	sql := "SELECT id,question,answer_number,askerid,creat_time,update_time FROM question_subject where id > ?"
	err := global.MysqlDB.Select(&questions, sql, 0)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return questions
}

func (s *SUser) GetAnswer(ctx context.Context, questionid int64) []model.AnswerSubject {
	answers := []model.AnswerSubject{}
	sql := "SELECT id,answer,writerid,questionid,likenumber,creat_time,update_time FROM answer_subject where questionid = ?"
	err := global.MysqlDB.Select(&answers, sql, questionid)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return answers
}

func (s *SUser) GetComment(ctx context.Context, answerid int64) []model.Comment {
	comment := []model.Comment{}
	sql := "SELECT id,comment,likenumber,answerid,writerid,creat_time,update_time FROM comment_subject where id = ?"
	err := global.MysqlDB.Select(&comment, sql, answerid)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return comment
}

func (s *SUser) GetUser(ctx context.Context, username string) model.UserSubject {
	user := model.UserSubject{}
	sql := "SELECT id,username,password,creat_time,update_time FROM user_subject where username = ?"
	_ = global.MysqlDB.Select(user, sql, username)
	return user
}

func (s *SUser) GetArticles(ctx context.Context) []model.ArticleSubject {
	articles := []model.ArticleSubject{}
	sql := "SELECT id,article,answer_number,writerid,like_number,creat_time,update_time FROM article_subject where id > ?"
	err := global.MysqlDB.Select(&articles, sql, 0)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return articles
}

func (s *SUser) GetFollowQuestions(ctx context.Context, userid int64) []model.FollowQuestionSubject {
	followquestions := []model.FollowQuestionSubject{}
	sql := "SELECT id,userid,questionid,creat_time,update_time FROM followquestion_subject where id > ? AND  userid = ?"
	err := global.MysqlDB.Select(&followquestions, sql, 0, userid)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return followquestions
}

func (s *SUser) GetFavourites(ctx context.Context, userid int64) ([]model.FavouriteSubject, error) {
	favourite := []model.FavouriteSubject{}
	sql := "SELECT id,userid,favouritenumber,`public`,`name`,creat_time,update_time FROM favourite_subject where id > ? and userid = ?"
	err := global.MysqlDB.Select(&favourite, sql, 0, userid)
	if err != nil {
		return nil, err
	}
	return favourite, err
}

func (s *SUser) GetFavourite(ctx context.Context, userid int64, favouriteid int64) ([]model.MyFavouriteQuestionSubject, []model.MyFavouriteArticleSubject, []model.MyFavouriteAnswerSubject, error) {
	favouriteanswer := []model.MyFavouriteAnswerSubject{}
	sql := "SELECT id,answerid,favouriteid,userid,creat_time,update_time FROM myfavouriteanswer_subject where userid = ? and favouriteid = ?"
	err := global.MysqlDB.Select(&favouriteanswer, sql, userid, favouriteid)
	if err != nil {
		return nil, nil, nil, err
	}

	favouritearticle := []model.MyFavouriteArticleSubject{}
	sql = "SELECT id,articleid,favouriteid,userid,creat_time,update_time FROM myfavouritearticle_subject where userid = ? and favouriteid = ?"
	err = global.MysqlDB.Select(favouritearticle, sql, userid, favouriteid)
	if err != nil {
		return nil, nil, nil, err
	}

	favouritequestion := []model.MyFavouriteQuestionSubject{}
	sql = "SELECT id,questionid,favouriteid,userid,creat_time,update_time FROM myfavouritequestion_subject where userid = ? and favouriteid = ?"
	err = global.MysqlDB.Select(favouritequestion, sql, userid, favouriteid)
	if err != nil {
		return nil, nil, nil, err
	}

	return favouritequestion, favouritearticle, favouriteanswer, nil
}

func (s *SUser) GetFavouriteAnswer(ctx context.Context, userid int64, answerid int64, favouriteid int64) (*model.MyFavouriteAnswerSubject, error) {
	favouriteanswer := &model.MyFavouriteAnswerSubject{}
	sql := "SELECT id,answerid,favouriteid,userid,creat_time,update_time FROM myfavouriteanswer_subject where userid = ? and answerid = ? and favouriteid = ?"
	err := global.MysqlDB.Select(favouriteanswer, sql, userid, answerid, favouriteid)
	if err != nil {
		return nil, err
	}
	return favouriteanswer, err
}

func (s *SUser) GetFavouriteArticle(ctx context.Context, userid int64, articleid int64, favouriteid int64) (*model.MyFavouriteArticleSubject, error) {
	favouritearticle := &model.MyFavouriteArticleSubject{}
	sql := "SELECT id,articleid,favouriteid,userid,creat_time,update_time FROM myfavouritearticle_subject where userid = ? and articleid = ? and favouriteid = ?"
	err := global.MysqlDB.Select(favouritearticle, sql, userid, articleid, favouriteid)
	if err != nil {
		return nil, err
	}
	return favouritearticle, err
}

func (s *SUser) GetFavouriteQuestion(ctx context.Context, userid int64, questionid int64, favouriteid int64) (*model.MyFavouriteQuestionSubject, error) {
	favouritequestion := &model.MyFavouriteQuestionSubject{}
	sql := "SELECT id,questionid,favouriteid,userid,creat_time,update_time FROM myfavouritequestion_subject where userid = ? and questionid = ? and favouriteid = ?"
	err := global.MysqlDB.Select(favouritequestion, sql, userid, questionid, favouriteid)
	if err != nil {
		return nil, err
	}
	return favouritequestion, err
}
func (s *SUser) CheckAnswerIsExist(ctx context.Context, answerid int64) error {
	answerSubject := &model.AnswerSubject{}
	sql := "SELECT id FROM answer_subject where id = ?"
	err := global.MysqlDB.Get(answerSubject, sql, answerid)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "answer_subject"),
			)
			return fmt.Errorf("internal err")
		} else {
			return fmt.Errorf("answer not found")
		}
	}
	return nil
}

func (s *SUser) CheckQuestionIsExist(ctx context.Context, questionid int64) error {
	questionsubject := &model.Question{}
	sql := "SELECT id FROM question_subject where id = ?"
	err := global.MysqlDB.Get(questionsubject, sql, questionid)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "answer_subject"),
			)
			return fmt.Errorf("internal err")
		} else {
			return fmt.Errorf("question not found")
		}
	}
	return nil
}

func (s *SUser) AddFollowQuestion(ctx context.Context, followquestion *model.FollowQuestionSubject) error {
	sql := "INSERT INTO followquestion_subject(userid,questionid) VALUES (?,?)"
	_, err := global.MysqlDB.Exec(sql, followquestion.Userid, followquestion.Questionid)
	if err != nil {
		return err
	}
	return nil
}

func (s *SUser) AddFollowUser(ctx context.Context, followuser *model.FollowUserSubject) error {
	sql := "INSERT INTO followuser_subject(userid,followuserid) VALUES (?,?)"
	_, err := global.MysqlDB.Exec(sql, followuser.Userid, followuser.FollowUserid)
	if err != nil {
		return err
	}
	return nil
}

func (s *SUser) AddFollowFavourite(ctx context.Context, followfavourite *model.FollowFavouriteSubject) error {
	sql := "INSERT INTO followuser_subject(userid,favouriteid) VALUES (?,?)"
	_, err := global.MysqlDB.Exec(sql, followfavourite.Userid, followfavourite.Favouriteid)
	if err != nil {
		return err
	}
	return nil
}

func (s *SUser) AddFavouriteAnswer(ctx context.Context, favouriteanswer *model.MyFavouriteAnswerSubject) error {
	sql := "INSERT INTO myfavouriteanswer_subject(userid,favouriteid,answerid) VALUES (?,?,?)"
	_, err := global.MysqlDB.Exec(sql, favouriteanswer.Userid, favouriteanswer.Favouriteid, favouriteanswer.Answerid)
	if err != nil {
		return err
	}
	return nil
}

func (s *SUser) AddFavouriteArticle(ctx context.Context, favouritearticle *model.MyFavouriteArticleSubject) error {
	sql := "INSERT INTO  myfavouritearticle_subject(userid,favouriteid,articleid) VALUES (?,?,?)"
	_, err := global.MysqlDB.Exec(sql, favouritearticle.Userid, favouritearticle.Favouriteid, favouritearticle.Articleid)
	if err != nil {
		return err
	}
	return nil
}

func (s *SUser) AddFavouriteQuestion(ctx context.Context, favouritequestion *model.MyFavouriteQuestionSubject) error {
	sql := "INSERT INTO  myfavouritearticle_subject(userid,favouriteid,questionid) VALUES (?,?,?)"
	_, err := global.MysqlDB.Exec(sql, favouritequestion.Userid, favouritequestion.Favouriteid, favouritequestion.Questionid)
	if err != nil {
		return err
	}
	return nil
}

func (s *SUser) AddLikeQuestion(ctx context.Context, questionid int64, userid int64) string {
	var question = model.Question{}
	cmd := global.Rdb.SAdd(ctx, fmt.Sprintf("questionid:%s", string(questionid)), userid)
	if cmd.Val() == 0 {
		return "error"
	}
	sqlStr := "SELECT id,question,answer_number,askerid,creat_time,update_time FROM question_subject where id = ?"
	err1 := global.MysqlDB.Get(&question, sqlStr, questionid)
	if err1 != nil {
		fmt.Print(err1)
		return "error"
	}
	sqlstr := "UPDATE question_subject set like_number = ?where id = ?"
	_, err2 := global.MysqlDB.Exec(sqlstr, question.LikeNumber+1, questionid)
	if err2 != nil {
		fmt.Print(err2)
		return "error"
	}
	return ""
}

func (s *SUser) AddLikeAnswer(ctx context.Context, answerid int64, userid int64) string {
	var answer = model.AnswerSubject{}
	err := global.Rdb.SAdd(ctx, fmt.Sprintf("answerid:%s", string(answerid)), userid)
	if err != nil {
		return "error"
	}

	sqlStr := "SELECT id,answer,comment_number,likenumber,comment_number,writerid,creat_time,update_time FROM answer_subject where id = ?"
	err1 := global.MysqlDB.Get(&answer, sqlStr, answerid)
	if err1 != nil {
		fmt.Print(err1)
		return "error"
	}
	sqlstr := "UPDATE answer_subject set likenumber = ?where id = ?"
	_, err2 := global.MysqlDB.Exec(sqlstr, answer.LikeNumber+1, answerid)
	if err2 != nil {
		fmt.Print(err2)
		return "error"
	}
	return ""
}

func (s *SUser) AddLikeComment(ctx context.Context, commentid int64, userid int64) string {
	var comment = model.Comment{}
	err := global.Rdb.SAdd(ctx, fmt.Sprintf("commentid:%s", string(commentid)), userid)
	if err != nil {
		return "error"
	}
	sqlStr := "SELECT id,comment,answerid,likenumber,writerid,creat_time,update_time FROM comment_subject where id = ?"
	err1 := global.MysqlDB.Get(&comment, sqlStr, commentid)
	if err1 != nil {
		fmt.Print(err1)
		return "error"
	}
	sqlstr := "UPDATE comment_subject set likenumber = ?where id = ?"
	_, err2 := global.MysqlDB.Exec(sqlstr, comment.LikeNumber+1, commentid)
	if err2 != nil {
		fmt.Print(err2)
		return "error"
	}
	return ""
}

func (s *SUser) AddLikeArticle(ctx context.Context, articleid int64, userid int64) string {
	var article = model.ArticleSubject{}
	cmd := global.Rdb.SAdd(ctx, fmt.Sprintf("articleid:%s", string(articleid)), userid)
	if cmd.Val() == 0 {
		return "error"
	}

	sqlStr := "SELECT id,article,answer_number,like_number,writerid,creat_time,update_time FROM article_subject where id = ?"
	err1 := global.MysqlDB.Get(&article, sqlStr, articleid)
	if err1 != nil {
		fmt.Print(err1)
		return "error"
	}
	sqlstr := "UPDATE article_subject set like_number = ?where id = ?"
	_, err2 := global.MysqlDB.Exec(sqlstr, article.LikeNumber+1, articleid)
	if err2 != nil {
		fmt.Print(err2)
		return "error"
	}
	return ""
}
func (s *SUser) CancelFollowQuestion(ctx context.Context, followquestion *model.FollowQuestionSubject) error {
	sql := "DELETE FROM followquestion_subject WHERE id = ?"
	_, err := global.MysqlDB.Exec(sql, followquestion.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SUser) CancelFollowUser(ctx context.Context, followuser *model.FollowUserSubject) error {
	sql := "DELETE FROM followuser_subject WHERE id = ?"
	_, err := global.MysqlDB.Exec(sql, followuser.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SUser) CancelFollowFavourite(ctx context.Context, followfavourite *model.FollowFavouriteSubject) error {
	sql := "DELETE FROM followfavourite_subject WHERE id = ?"
	_, err := global.MysqlDB.Exec(sql, followfavourite.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SUser) CancelFavouriteArticle(ctx context.Context, favouritearticleid int64) error {
	sql := "DELETE FROM myavouritearticle_subject WHERE id = ?"
	_, err := global.MysqlDB.Exec(sql, favouritearticleid)
	if err != nil {
		return err
	}
	return nil
}

func (s *SUser) CancelFavouriteAnswer(ctx context.Context, favouriteanswerid int64) error {
	sql := "DELETE FROM myavouriteanswer_subject WHERE id = ?"
	_, err := global.MysqlDB.Exec(sql, favouriteanswerid)
	if err != nil {
		return err
	}
	return nil
}

func (s *SUser) CancelFavouriteQuestion(ctx context.Context, favouritequestionid int64) error {
	sql := "DELETE FROM myavouritequestion_subject WHERE id = ?"
	_, err := global.MysqlDB.Exec(sql, favouritequestionid)
	if err != nil {
		return err
	}
	return nil
}

func (s *SUser) CancelLikeQuestion(ctx context.Context, questionid int64, userid int64) string {
	var question = model.Question{}
	cmd := global.Rdb.SRem(ctx, fmt.Sprintf("questionid:%s", string(questionid)), userid)
	if cmd.Val() == 0 {
		return "error"
	}

	sqlStr := "SELECT id,question,answer_number,askerid,creat_time,update_time FROM question_subject where id = ?"
	err1 := global.MysqlDB.Get(&question, sqlStr, questionid)
	if err1 != nil {
		fmt.Print(err1)
		return "error"
	}
	sqlstr := "UPDATE question_subject set like_number = ?where id = ?"
	_, err2 := global.MysqlDB.Exec(sqlstr, question.LikeNumber-1, questionid)
	if err2 != nil {
		fmt.Print(err2)
		return "error"
	}
	return ""
}

func (s *SUser) CancelLikeComment(ctx context.Context, commentid int64, userid int64) string {
	var comment = model.Comment{}
	cmd := global.Rdb.SRem(ctx, fmt.Sprintf("commentid:%s", string(commentid)), userid)
	if cmd.Val() == 0 {
		return "error "
	}

	sqlStr := "SELECT id,comment,answerid,likenumber,writerid,creat_time,update_time FROM comment_subject where id = ?"
	err1 := global.MysqlDB.Get(&comment, sqlStr, commentid)
	if err1 != nil {
		fmt.Print(err1)
		return "error"
	}
	sqlstr := "UPDATE comment_subject set likenumber = ?where id = ?"
	_, err2 := global.MysqlDB.Exec(sqlstr, comment.LikeNumber-1, commentid)
	if err2 != nil {
		fmt.Print(err2)
		return "error"
	}

	return ""
}

func (s *SUser) CancelLikeAnswer(ctx context.Context, answerid int64, userid int64) string {
	var answer = model.AnswerSubject{}
	err := global.Rdb.SRem(ctx, fmt.Sprintf("answerid:%s", string(answerid)), userid)
	if err != nil {
		return "error"
	}

	sqlStr := "SELECT id,answer,comment_number,likenumber,comment_number,writerid,creat_time,update_time FROM answer_subject where id = ?"
	err1 := global.MysqlDB.Get(&answer, sqlStr, answerid)
	if err1 != nil {
		fmt.Print(err1)
		return "error"
	}
	sqlstr := "UPDATE answer_subject set likenumber = ?where id = ?"
	_, err2 := global.MysqlDB.Exec(sqlstr, answer.LikeNumber-1, answerid)
	if err2 != nil {
		fmt.Print(err2)
		return "error"
	}
	return ""
}

func (s *SUser) CancelLikeArticle(ctx context.Context, articleid int64, userid int64) string {
	var article = model.ArticleSubject{}
	cmd := global.Rdb.SRem(ctx, fmt.Sprintf("articleid:%s", string(articleid)), userid)
	if cmd.Val() == 0 {
		return "error"
	}

	sqlStr := "SELECT id,article,answer_number,like_number,writerid,creat_time,update_time FROM article_subject where id = ?"
	err1 := global.MysqlDB.Get(&article, sqlStr, articleid)
	if err1 != nil {
		fmt.Print(err1)
		return "error"
	}
	sqlstr := "UPDATE article_subject set like_number = ?where id = ?"
	_, err2 := global.MysqlDB.Exec(sqlstr, article.LikeNumber-1, articleid)
	if err2 != nil {
		fmt.Print(err2)
		return "error"
	}
	return ""
}

func (s *SUser) DeleteFavourites(ctx context.Context, favouritesid int64) error {
	sql := "DELETE FROM favourites_subject WHERE id = ?"
	_, err := global.MysqlDB.Exec(sql, favouritesid)
	if err != nil {
		return err
	}

	sql = "DELETE FROM myfavouriteanswer_subject WHERE favouriteid = ?"
	_, err = global.MysqlDB.Exec(sql, favouritesid)
	if err != nil {
		return err
	}

	sql = "DELETE FROM myfavouritearticle_subject WHERE favouriteid = ?"
	_, err = global.MysqlDB.Exec(sql, favouritesid)
	if err != nil {
		return err
	}
	return nil

}

func (s *SUser) ChangePassword(ctx context.Context, password string, userid int64) error {
	sql := "UPDATE user_subject SET password =? where id=?"
	_, err := global.MysqlDB.Exec(sql, password, userid)
	if err != nil {
		return err
	}
	return nil
}

func (s *SUser) ChangeUsername(ctx context.Context, username string, userid int64) error {
	sql := "UPDATE user_subject SET username =? where id=?"
	_, err := global.MysqlDB.Exec(sql, username, userid)
	if err != nil {
		return err
	}
	return nil
}

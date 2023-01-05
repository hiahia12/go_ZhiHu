package model

import "time"

type UserSubject struct {
	Id         int64     `gorm:"column:id" json:"id" form:"id" db:"id"`
	Username   string    `gorm:"column:username" json:"username" form:"username" db:"username"`
	Password   string    `gorm:"column:password" json:"password" form:"password" db:"password"`
	CreatTime  time.Time `gorm:"column:creat_time;autoCreateTime" json:"creat_time" form:"creat_time" db:"creat_time"`
	UpdateTime time.Time `gorm:"column:update_time;autoCreateTime" json:"update_time" form:"update_time" db:"update_time"`
}

type CollectSubject struct {
	CreaterId   int64  `gorm:"column:createrid" json:"createrid" form:"createrid" db:"createrid"`
	CollectName string `gorm:"column:collect_name" json:"collect_name" form:"collect_name" db:"collect_name"`
	Answerid    int64  `gorm:"column:answerid" json:"answerid" form:"answerid" db:"answerid"`
}

type AnswerSubject struct {
	Id            int64     `gorm:"column:id" json:"id" form:"id" db:"id"`
	Answer        string    `gorm:"column:answer" json:"answer" form:"answer" db:"answer"`
	Writerid      int64     `gorm:"column:writerid" json:"writerid" form:"writerid" db:"writerid"`
	Questionid    int64     `gorm:"column:questionid" json:"questionid" form:"questionid" db:"questionid"`
	Like          int       `gorm:"column:like" json:"like" form:"like" db:"like"`
	CommentNumber int       `gorm:"column:comment_number" json:"comment_number" form:"comment_number" db:"comment_number"`
	CreatTime     time.Time `gorm:"column:creat_time;autoCreateTime" json:"creat_time" form:"creat_time" db:"creat_time"`
	UpdateTime    time.Time `gorm:"column:update_time;autoCreateTime" json:"update_time" form:"update_time" db:"update_time"`
}

type Question struct {
	Id           int64     `gorm:"column:id" json:"id" form:"id" db:"id"`
	Question     string    `gorm:"column:question" json:"question" form:"question" db:"question"`
	Askerid      int64     `gorm:"column:askerid" json:"askerid" form:"askerid" db:"askerid"`
	AnswerNumber int       `gorm:"column:answer_number" json:"answer_number" form:"answer_number" db:"answer_number"`
	CreatTime    time.Time `gorm:"column:creat_time;autoCreateTime" json:"creat_time" form:"creat_time" db:"creat_time"`
	UpdateTime   time.Time `gorm:"column:update_time;autoCreateTime" json:"update_time" form:"update_time" db:"update_time"`
}

type ArticleSubject struct {
	Id           int64     `json:"id" form:"id" db:"id" `
	Article      string    `json:"article" form:"article" db:"article" `
	Writerid     int64     `json:"writerid" form:"writerid" db:"writerid" `
	AnswerNumber int       `json:"answer_number" form:"answer_number" db:"answer_number" `
	LikeNumber   int       `json:"like_number" form:"like_number" db:"like_number"`
	CreatTime    time.Time `gorm:"column:creat_time;autoCreateTime" json:"creat_time" form:"creat_time" db:"creat_time"`
	UpdateTime   time.Time `gorm:"column:update_time;autoCreateTime" json:"update_time" form:"update_time" db:"update_time"`
}
type Comment struct {
	Id         int64     `gorm:"column:id" json:"id" form:"id" db:"id"`
	Comment    string    `gorm:"column:comment" json:"comment" form:"comment" db:"comment"`
	Like       int       `gorm:"column:like" json:"like" form:"like" db:"like"`
	Answerid   int64     `gorm:"column:answerid" json:"answerid" form:"answerid" db:"answerid"`
	Writerid   int64     `gorm:"column:writerid" json:"writerid" form:"writerid" db:"writerid"`
	CreatTime  time.Time `gorm:"column:creat_time;autoCreateTime" json:"creat_time" form:"creat_time" db:"creat_time"`
	UpdateTime time.Time `gorm:"column:update_time;autoCreateTime" json:"update_time" form:"update_time" db:"update_time"`
}

type FollowFavouriteSubject struct {
	Id          int64     `json:"id" form:"id" db:"id"`
	Userid      int64     `json:"userid" form:"userid" db:"userid"`
	Favouriteid int64     `json:"favouriteid" form:"favouriteid" db:"favouriteid"`
	CreatTime   time.Time `json:"creat_time" form:"creat_time" db:"creat_time"`
	UpdateTimr  time.Time `json:"update_time" form:"update_time" db:"update_time"`
}

type FollowQuestionSubject struct {
	Id         int64     `json:"id" form:"id" db:"id"`
	Userid     int64     `json:"userid" form:"userid" db:"userid"`
	Questionid int64     `json:"questionid" form:"questionid" db:"questionid"`
	CreatTime  time.Time `json:"creat_time" form:"creat_time" db:"creat_time"`
	UpdateTimr time.Time `json:"update_time" form:"update_time" db:"update_time"`
}

type FollowUserSubject struct {
	Id           int64     `json:"id" form:"id" db:"id"`
	Userid       int64     `json:"userid" form:"userid" db:"userid"`
	FollowUserid int64     `json:"followuserid" form:"followuserid" db:"followuserid"`
	CreatTime    time.Time `json:"creat_time" form:"creat_time" db:"creat_time"`
	UpdateTimr   time.Time `json:"update_time" form:"update_time" db:"update_time"`
}

type MyFavouriteArticleSubject struct {
	Id          int64     `json:"id" form:"id" db:"id"`
	Articleid   int64     `json:"articleid" form:"articleid" db:"articleid"`
	Favouriteid int64     `json:"favouriteid" form:"favouriteid" db:"favouriteid"`
	Userid      int64     `json:"userid" form:"userid" db:"userid"`
	CreatTime   time.Time `json:"creat_time" form:"creat_time" db:"creat_time"`
	UpdateTimr  time.Time `json:"update_time" form:"update_time" db:"update_time"`
}

type FavouriteSubject struct {
	Id              int64     `json:"id" form:"id" db:"id"`
	Userid          int64     `json:"userid" form:"userid" db:"userid"`
	FavouriteNumber int       `json:"favourite_number" form:"favourite_number" db:"favourite_number"`
	Public          int       `json:"public" form:"public" db:"public"`
	Name            string    `json:"name" form:"name" db:"name"`
	CreatTime       time.Time `json:"creat_time" form:"creat_time" db:"creat_time"`
	UpdateTimr      time.Time `json:"update_time" form:"update_time" db:"update_time"`
}

type MyFavouriteAnswerSubject struct {
	Id          int64     `json:"id" form:"id" db:"id"`
	Answerid    int64     `json:"articleid" form:"articleid" db:"articleid"`
	Favouriteid int64     `json:"favouriteid" form:"favouriteid" db:"favouriteid"`
	Userid      int64     `json:"userid" form:"userid" db:"userid"`
	CreatTime   time.Time `json:"creat_time" form:"creat_time" db:"creat_time"`
	UpdateTimr  time.Time `json:"update_time" form:"update_time" db:"update_time"`
}

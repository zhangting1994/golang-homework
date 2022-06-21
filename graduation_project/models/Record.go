package models

type Record struct {
	Id  int64  `xorm:"pk autoincr"`
	Msg string
}

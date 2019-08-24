package models

type User struct {
	Id       int    `bson:"id" json:"id"`
	UUID     int64  `bson:"uuid" json:"uuid"`
	Nickname string `bson:"nickname" json:"nickname"`
	Account  string `bson:"account" json:"account"`
	Password string `bson:"password" json:"password"`
}

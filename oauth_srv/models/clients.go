package models

type Client struct {
	ID     string `bson:"_id" json:"id"`
	Key    string `bson:"key" json:"key"`
	Secret string `bson:"secret" json:"secret"`
	Domain string `bson:"domain" json:"domain"`
}

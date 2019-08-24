package models

import "time"

type Token struct {
	ClientID         string        `bson:"client_id" json:"client_id"`
	UserID           string        `bson:"user_id" json:"user_id"`
	RedirectURI      string        `bson:"redirect_uri" json:"redirect_uri"`
	Scope            string        `bson:"scope" json:"scope"`
	Code             string        `bson:"code" json:"code"`
	CodeCreateAt     time.Time     `bson:"code_create_at" json:"code_create_at"`
	CodeExpiresIn    time.Duration `bson:"code_expires_in" json:"code_expires_in"`
	Access           string        `bson:"access" json:"access"`
	AccessCreateAt   time.Time     `bson:"access_create_at" json:"access_create_at"`
	AccessExpiresIn  time.Duration `bson:"access_expires_in" json:"access_expires_in"`
	Refresh          string        `bson:"refresh" json:"refresh"`
	RefreshCreateAt  time.Time     `bson:"refresh_create_at" json:"refresh_create_at"`
	RefreshExpiresIn time.Duration `bson:"refresh_expires_in" json:"refresh_expires_in"`
}

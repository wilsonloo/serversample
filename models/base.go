package models

import "gopkg.in/LyricTian/lib.v2/mongo"

var (
	mHandler *mongo.Handler
)

// Init Model初始化
func Init(url, dbName string) error {
	handler, err := mongo.InitHandlerWithDB(url, dbName)
	if err != nil {
		return err
	}
	mHandler = handler
	return nil
}

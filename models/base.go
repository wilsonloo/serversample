package models

import "gopkg.in/LyricTian/lib.v2/mongo"
import "log"

var (
	mHandler *mongo.Handler
)

// Init Model初始化
func Init(url, dbName string) error {
	/*
	todo modify by wilson: 此处应该添加log，用以提示当前处于哪个进度，万一对方配置表的mongodb配置出错，
		就会出现 panic: no reachable servers ，令人困惑
	 */
	log.Printf("Initializing mongodb url:%s dbname:%s ...\n", url, dbName)
	handler, err := mongo.InitHandlerWithDB(url, dbName)
	if err != nil {
		return err
	}
	mHandler = handler
	log.Printf("Initializing mongodb url:%s dbname:%s ...OK\n", url, dbName)

	return nil
}

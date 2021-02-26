package main

import (
	"TaskData/dao"
	"TaskData/read"
	"TaskData/write"
)

// 初始化读的数据库
func initRead() *read.ReadFromGms {
	readMysql := dao.MysqlConfig{
		Host:         "mysql.middleware.com:3306",
		User:         "baldr",
		Pass:         "eyesight2020",
		Debug:        true,
		DBName:       "baldr110",
		MaxIdleConns: 10,
		MaxOpenConns: 100,
	}

	readFromGms := &read.ReadFromGms{
		Db:          nil,
		MysqlConfig: &readMysql,
	}

	// 初始化读的数据库配置
	readFromGms.InitReadDb()
	return readFromGms
}

// 初始化写的数据库
func initWrite() *write.WriteToTaskService {
	writeMysql := dao.MysqlConfig{
		Host:         "mysql.middleware.com:3306",
		User:         "baldr",
		Pass:         "eyesight2020",
		Debug:        true,
		DBName:       "task",
		MaxIdleConns: 10,
		MaxOpenConns: 100,
	}

	writeToTaskService := &write.WriteToTaskService{
		Db:          nil,
		MysqlConfig: &writeMysql,
	}

	// 初始化写的数据库配置
	writeToTaskService.InitWriteDb()
	return writeToTaskService
}
func main() {
	// 初始化读
	_ = initRead()

	_ = initWrite()

}

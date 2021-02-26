package dao

import "fmt"

type MysqlConfig struct {
	Host         string `json:"Host"`
	User         string `json:"User"`
	Pass         string `json:"Pass"`
	Debug        bool   `json:"Debug"`
	DBName       string `json:"DBName"`
	MaxIdleConns int    `json:"MaxIdleConns"`
	MaxOpenConns int    `json:"MaxOpenConns"`
}

func GetMysqlUrl(host, user, pass, dBName string) string {
	dBOption := "charset=utf8&parseTime=True&loc=Local"
	result := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", user, pass, host, dBName, dBOption)
	return result
}

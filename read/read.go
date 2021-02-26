package read

import (
	"TaskData/dao"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Reader interface {
	Read(rc chan string)
}

type ReadFromGms struct {
	Db *gorm.DB
	*dao.MysqlConfig
}

func (r *ReadFromGms) InitReadDb() {
	var err error
	// 初始化数据库
	mysqlUrl := dao.GetMysqlUrl(r.Host, r.User, r.Pass, r.DBName)
	r.Db, err = gorm.Open("mysql", mysqlUrl)
	if err != nil {
		panic("mysql init error" + err.Error())
	}
	r.Db.LogMode(r.Debug)
	r.Db.DB().SetMaxIdleConns(r.MaxIdleConns)
	r.Db.DB().SetMaxOpenConns(r.MaxOpenConns)
	fmt.Println("read mysql init success")
}

func (r *ReadFromGms) Read(rc chan string) {

}

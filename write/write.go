package write

import (
	"TaskData/dao"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Writer interface {
	Write(rc chan string)
}

type WriteToTaskService struct {
	Db *gorm.DB
	*dao.MysqlConfig
}

func (w *WriteToTaskService) InitWriteDb() {
	var err error
	// 初始化数据库
	mysqlUrl := dao.GetMysqlUrl(w.Host, w.User, w.Pass, w.DBName)
	w.Db, err = gorm.Open("mysql", mysqlUrl)
	if err != nil {
		panic("mysql init error" + err.Error())
	}
	w.Db.LogMode(w.Debug)
	w.Db.DB().SetMaxIdleConns(w.MaxIdleConns)
	w.Db.DB().SetMaxOpenConns(w.MaxOpenConns)
	fmt.Println("write mysql init success")
}

func (w *WriteToTaskService) Write(rc chan string) {

}

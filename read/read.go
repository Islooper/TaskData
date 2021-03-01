package read

import (
	"TaskData/dao"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
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

func (r *ReadFromGms) Read(taskRc chan []*dao.TaskDo, visionRc chan []*dao.VisionDo, opRc chan []*dao.OptometryDo, wg *sync.WaitGroup) {
	for {
		// 每次读取10条task信息
		taskDos := dao.ReadTasks(r.Db)
		if len(taskDos) == 0 {
			// 读完了
			close(taskRc)
		} else {
			// 数据发给分析模块进行转换或者处理
			taskRc <- taskDos
			// 任务数+1
			wg.Add(1)
		}

		// 每次读取10条视力数据
		visionDos := dao.ReadVisions(r.Db)
		if len(visionDos) == 0 {
			// 读完了
			close(visionRc)
		} else {
			// 数据发给分析模块进行转换或者处理
			// 查找学生id
			for _, vision := range visionDos {
				stu := dao.ReadStuId(vision.Idcard, r.Db)
				if stu != nil {
					vision.StudentId = stu.StudentId
				}
			}
			visionRc <- visionDos
			// 任务数+1
			wg.Add(1)
		}

		// 每次读取10条屈光数据
		opDos := dao.ReadOps(r.Db)
		if len(opDos) == 0 {
			// 读完了
			close(opRc)
		} else {
			// 数据发给分析模块进行转换或者处理
			// 查找学生id
			for _, op := range opDos {
				stu := dao.ReadStuId(op.Idcard, r.Db)
				if stu != nil {
					op.StudentId = stu.StudentId
				}
			}
			opRc <- opDos
			// 任务数+1
			wg.Add(1)
		}
	}
}

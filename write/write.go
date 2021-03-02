package write

import (
	"TaskData/dao"
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
)

type Writer interface {
	Write(taskWc chan []*dao.Task, taskDataWc chan []*dao.TaskData, wg *sync.WaitGroup)
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

func (w *WriteToTaskService) Write(taskWc chan []*dao.Task, taskDataWc chan []*dao.TaskData, wg *sync.WaitGroup) {
	// 写入任务
	// 拿到任务数据
	taskDos := <-taskWc
	// 新增
	err := dao.Create(taskDos, w.Db)
	if err != nil {
		fmt.Print("create task fail : ", err)
	}

	// 写入vision、optometry
	taskDatasDos := <-taskDataWc
	// 判断是新增还是更新

	for _, taskDataDo := range taskDatasDos {
		// 查看是否存在
		taskDo, err := dao.ReadTaskDataByTaskIdAndStudentId(taskDataDo.TaskId, taskDataDo.StudentId, w.Db)

		if err != nil {
			fmt.Print("read taskData fail : ", err)
		}

		if taskDo == nil {
			// 新增
			err := dao.CreateTaskData(taskDo, w.Db)
			if err != nil {
				fmt.Print("create taskData fail : ", err)
			}
		} else {
			// 更新

			// 处理变量
			updateTaskDataDo := dealTaskDataUpdate(taskDataDo, taskDo)
			err := dao.UpdateTaskData(updateTaskDataDo, w.Db)
			if err != nil {
				fmt.Println("update task data fail ", err)
			}
		}

		wg.Done()
	}

}

func dealTaskDataUpdate(oldTaskData, newTaskData *dao.TaskData) *dao.TaskData {
	taskDataDo := new(dao.TaskData)

	taskDataDo.TaskId = oldTaskData.TaskId
	taskDataDo.StudentId = oldTaskData.StudentId
	taskDataDo.OrgId = oldTaskData.OrgId
	taskDataDo.ClassId = oldTaskData.ClassId
	taskDataDo.IdCard = oldTaskData.IdCard
	taskDataDo.Memo = "old"
	taskDataDo.UpdatedAt = oldTaskData.UpdatedAt
	taskDataDo.CreatedAt = oldTaskData.CreatedAt
	if oldTaskData.LeftAp == "" || oldTaskData.RightAp == "" {
		taskDataDo.LeftAxial = newTaskData.LeftAxial
		taskDataDo.RightAxial = newTaskData.RightAxial
		taskDataDo.LeftSph = newTaskData.LeftSph
		taskDataDo.RightSph = newTaskData.RightSph
		taskDataDo.LeftAp = newTaskData.LeftAp
		taskDataDo.RightAp = newTaskData.RightAp
		taskDataDo.ImgUrl = newTaskData.ImgUrl
		taskDataDo.OpUpTime = newTaskData.OpUpTime
	} else {
		taskDataDo.LeftVision = oldTaskData.LeftVision
		taskDataDo.RightVision = oldTaskData.RightVision
		taskDataDo.LeftGlassesVision = oldTaskData.LeftGlassesVision
		taskDataDo.RightGlassesVision = oldTaskData.RightGlassesVision
		taskDataDo.IsGlasses = oldTaskData.IsGlasses
		taskDataDo.GlassesType = oldTaskData.GlassesType
		taskDataDo.ViUpTime = oldTaskData.ViUpTime
	}

	return taskDataDo
}

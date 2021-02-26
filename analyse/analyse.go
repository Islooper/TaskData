package analyse

import (
	"TaskData/dao"
	"TaskData/read"
	"TaskData/write"
	"sync"
)

type DataSync struct {
	TaskRc     chan []*dao.TaskDo
	VisionRc   chan []*dao.VisionDo
	OpRc       chan []*dao.OptometryDo
	TaskWc     chan []*dao.Task
	TaskDataWc chan []*dao.TaskData
	read.Reader
	write.Writer
	wg *sync.WaitGroup
}

func (d *DataSync) Analyse() {
	// 获取任务数据
	taskDos := <-d.TaskRc
	newTaskDos := analyseTask(taskDos)
	d.TaskWc <- newTaskDos

	// 获取视力数据

	// 获取屈光数据
}

func analyseTask(taskDos []*dao.TaskDo) []*dao.Task {
	//for _ , taskDo := range taskDos  {
	//
	//}
	return nil
}

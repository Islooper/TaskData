package analyse

import (
	"TaskData/dao"
	"TaskData/read"
	"TaskData/write"
	"strconv"
	"strings"
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
	newTaskDos := make([]*dao.Task, 0)
	for _, taskDo := range taskDos {
		status, _ := strconv.ParseInt(taskDo.Status, 10, 64)
		taskType, _ := strconv.ParseInt(taskDo.Type, 10, 64)
		taskDevice, _ := strconv.ParseInt(taskDo.Device, 10, 64)

		// 处理Executors, 按逗号取出Executor
		executors := strings.Split(taskDo.Member, ",")
		executorDos := make([]*dao.ExecutorDo, 0)
		for _, executor := range executors {
			executorDo := &dao.ExecutorDo{
				ExecutorId: executor,
				TaskDoId:   taskDo.TaskId,
			}
			executorDos = append(executorDos, executorDo)
		}

		taskDo := &dao.Task{
			TaskId:       taskDo.TaskId,
			Status:       status,
			StartTime:    taskDo.StartTime,
			EndTime:      taskDo.EndTime,
			Type:         taskType + 1,
			Device:       taskDevice + 1,
			Name:         taskDo.Name,
			DoneTime:     taskDo.DoneTime,
			OrgId:        taskDo.OrgId,
			Pid:          taskDo.Pid,
			CreatorOrgId: taskDo.Creator,
			Describe:     "old",
			Executors:    executorDos,
			CreatedAt:    int64ToInt(taskDo.CreateTime),
			UpdatedAt:    int64ToInt(taskDo.DoneTime),
		}
		newTaskDos = append(newTaskDos, taskDo)
	}
	return newTaskDos
}

func int64ToInt(dest int64) int {
	strInt64 := strconv.FormatInt(dest, 10)
	id16, _ := strconv.Atoi(strInt64)
	return id16
}

func analyseVision() {

}

func analyseOp() {

}

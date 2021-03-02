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
	visionDos := <-d.VisionRc
	newVisionDos := analyseVision(visionDos)
	d.TaskDataWc <- newVisionDos
	// 获取屈光数据
	opDos := <-d.OpRc
	newOpDos := analyseOp(opDos)
	d.TaskDataWc <- newOpDos
}

func analyseTask(taskDos []*dao.TaskDo) []*dao.Task {
	newTaskDos := make([]*dao.Task, 0)
	for _, taskDo := range taskDos {
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
			Status:       stringToInt64(taskDo.Status),
			StartTime:    taskDo.StartTime,
			EndTime:      taskDo.EndTime,
			Type:         stringToInt64(taskDo.Type) + 1,
			Device:       stringToInt64(taskDo.Device) + 1,
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

func analyseVision(visionDos []*dao.VisionDo) []*dao.TaskData {
	taskDataDos := make([]*dao.TaskData, 0)
	for _, vision := range visionDos {
		taskDataDo := &dao.TaskData{
			CreatedAt:          int64ToInt(vision.CreateTime),
			TaskId:             vision.TaskId,
			StudentId:          vision.StudentId,
			IdCard:             vision.Idcard,
			ClassId:            vision.ClassId,
			OrgId:              vision.OrgId,
			LeftVision:         vision.LeftVision,
			RightVision:        vision.RightVision,
			IsGlasses:          int64ToInt(stringToInt64(vision.IsGlasses)),
			GlassesType:        int64ToInt(stringToInt64(vision.GlassesType)),
			LeftGlassesVision:  vision.LeftGlassesVision,
			RightGlassesVision: vision.RightGlassesVision,
			ViUpTime:           vision.Date,
			Memo:               "new",
		}
		taskDataDos = append(taskDataDos, taskDataDo)
	}
	return taskDataDos
}

func analyseOp(ops []*dao.OptometryDo) []*dao.TaskData {
	taskDataDos := make([]*dao.TaskData, 0)
	for _, op := range ops {
		taskData := &dao.TaskData{
			CreatedAt:  int64ToInt(op.CreateTime),
			TaskId:     op.TaskId,
			StudentId:  op.StudentId,
			IdCard:     op.StudentId,
			ClassId:    op.ClassId,
			OrgId:      op.OrgId,
			LeftSph:    op.LeftSph,
			RightSph:   op.RightSph,
			LeftAp:     op.LeftAp,
			RightAp:    op.RightAp,
			LeftAxial:  op.LeftAxial,
			RightAxial: op.RightAxial,
			ImgUrl:     op.ImgUrl,
			OpUpTime:   op.Date,
			Memo:       "new",
		}
		taskDataDos = append(taskDataDos, taskData)
	}

	return taskDataDos
}

func stringToInt64(str string) int64 {
	it64, _ := strconv.ParseInt(str, 10, 64)
	return it64
}

func int64ToInt(dest int64) int {
	strInt64 := strconv.FormatInt(dest, 10)
	id16, _ := strconv.Atoi(strInt64)
	return id16
}

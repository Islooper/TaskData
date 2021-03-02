package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TaskDo struct {
	TaskId     string
	Member     string
	Publisher  string // 任务负责人
	Status     string
	StartTime  int64
	EndTime    int64
	Bz         string
	CreateTime int64
	Type       string
	Device     string
	Name       string
	OrgId      string
	ClassId    string
	DoneTime   int64
	Pid        string
	Creator    string // 任务创建者
}

// TODO
func ReadTasks(db *gorm.DB) []*TaskDo {
	taskDos := make([]*TaskDo, 0)

	err := db.Find(&taskDos).Error

	if err != nil {
		panic("read task fail")
	}

	return taskDos
}

///////////////////////////////
type VisionDo struct {
	TaskId             string
	StudentId          string
	Idcard             string
	LeftVision         string
	RightVision        string
	IsGlasses          string
	LeftGlassesVision  string
	RightGlassesVision string
	GlassesType        string
	Memo               string
	Date               int64
	CreateTime         int64
	ClassId            string
	OrgId              string
}

func ReadVisions(db *gorm.DB) []*VisionDo {
	visionDos := make([]*VisionDo, 0)

	err := db.Find(&visionDos).Error

	if err != nil {
		panic("read vision fail")
	}

	return visionDos
}

type StudentDo struct {
	StudentId      string
	IdCard         string
	StudentNum     string
	Birthday       string
	Name           string
	Gender         string
	ClassId        string
	Nation         string
	OrgId          string
	FaceUrl        string
	EnrollmentYear int64
	StartDate      int64
	LastDate       int64
	Spell          string
	ParentsPhone   string
}

func ReadStuId(idCard string, db *gorm.DB) *StudentDo {
	stu := new(StudentDo)
	err := db.Where("id_card = ?", idCard).First(stu).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		} else {
			return nil
		}
	}
	return stu
}

///////////////////////////////////////
type OptometryDo struct {
	TaskId     string
	StudentId  string
	Idcard     string
	LeftSph    string
	RightSph   string
	LeftAp     string
	RightAp    string
	LeftAxial  string
	RightAxial string
	Bz         string
	Date       int64
	CreateTime int64
	ImgUrl     string
	ClassId    string
	OrgId      string
}

func ReadOps(db *gorm.DB) []*OptometryDo {
	OpDos := make([]*OptometryDo, 0)

	err := db.Find(&OpDos).Error

	if err != nil {
		panic("read op fail")
	}

	return OpDos
}

//////////////////////new

type Task struct {
	TaskId       string `gorm:"type:varchar(32);not null;primary_key"`
	Status       int64  `gorm:"type:int(4);comment:'任务状态'"`
	StartTime    int64  `gorm:"type:int(11);"`
	EndTime      int64  `gorm:"type:int(11);"`
	Type         int64  `gorm:"type:int(4);comment:'任务类型 1-筛查 2-复查 3-普查'"`
	Device       int64  `gorm:"type:int(4);comment:'1 - 视力 1- 验光 3- 全部'"`
	Name         string `gorm:"type:varchar(255);"`
	DoneTime     int64  `gorm:"type:int(11);"`
	OrgId        string `gorm:"type:varchar(32);"`
	Pid          string `gorm:"type:varchar(32);"`
	CreatorOrgId string `gorm:"type:varchar(32);comment:'任务创建者的机构id'"`
	Describe     string `gorm:"type:varchar(255);comment:'任务描述'"`

	Executors []*ExecutorDo `gorm:"ForeignKey:TaskDoId;ASSOCIATION_FOREIGNKEY:TaskId"` //这种多个元素的字段 需要把字段名字该为xxxs因为gorm会自动把表名加上s

	CreatedAt int
	UpdatedAt int
	DeletedAt int
}

type ExecutorDo struct {
	ID         uint   `gorm:"primarykey"`
	ExecutorId string `gorm:"type:varchar(36);not null"`
	TaskDoId   string `gorm:"type:varchar(36);not null;"`

	CreatedAt int
	UpdatedAt int
	DeletedAt int
}

// 新增
func Create(task []*Task, db *gorm.DB) error {
	if nil == task {
		return errors.New("task is nil")
	}

	err := db.Create(task).Error
	if err != nil {
		return err
	}
	return nil
}

////////////////////////new Task data
type TaskData struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt int
	UpdatedAt int
	TaskId    string `gorm:"type:varchar(32);not null;comment:'任务id';index:task_id_student_idx"`
	StudentId string `gorm:"type:varchar(32);not null;comment:'学生id';index:task_id_student_idx;"`
	IdCard    string `gorm:"type:varchar(20);not null;index:id_card_idx;comment:'身份证'"`
	ClassId   string `gorm:"type:varchar(32);not null;default:'';index:class_idx;comment:'创建数据时所在的班级'"`
	OrgId     string `gorm:"type:varchar(32);not null;index:school_idx;comment:'创建数据时所在的机构'"`
	//屈光
	LeftSph    string `gorm:"type:varchar(10);not null;default:'';comment:'左球镜'"`       // 左球镜
	RightSph   string `gorm:"type:varchar(10);not null;default:'';comment:'右球镜'"`       // 右球镜
	LeftAp     string `gorm:"type:varchar(10);not null;default:'';comment:'左柱镜'"`       // 左柱镜
	RightAp    string `gorm:"type:varchar(10);not null;default:'';comment:'右柱镜'"`       // 右柱镜
	LeftAxial  string `gorm:"type:varchar(10);not null;default:'';comment:'左轴位'"`       // 左轴位
	RightAxial string `gorm:"type:varchar(10);not null;default:'';comment:'右轴位'"`       // 右轴位
	ImgUrl     string `gorm:"type:varchar(255);not null;default:'';comment:'屈光小票照片地址'"` // 小票照片地址
	OpUpTime   int64  `gorm:"type:int(11);not null;default:0;comment:'屈光更新时间'"`         // op更新时间
	//视力
	LeftVision         string `gorm:"type:varchar(10);not null;default:'';comment:'左视力数据'"`
	RightVision        string `gorm:"type:varchar(10);not null;default:'';comment:'右视力数据'"`
	IsGlasses          int    `gorm:"type:int(2);not null;default:0;comment:'是否戴镜,0-不戴镜，1-戴镜'"`
	GlassesType        int    `gorm:"type:int(2);not null;default:0;comment:'眼镜类型 1-框架眼镜，2-角膜接触镜，3-夜戴角膜塑形镜'"` // 眼镜类型 1-框架眼镜，2-角膜接触镜，3-夜戴角膜塑形镜
	LeftGlassesVision  string `gorm:"type:varchar(10);not null;default:'';comment:'左戴镜视力'"`
	RightGlassesVision string `gorm:"type:varchar(10);not null;default:'';comment:'右戴镜视力'"`
	ViUpTime           int64  `gorm:"type:int(11);not null;default:0;comment:'视力数据更新时间'"` // vi更新时间

	Memo string `gorm:"type:varchar(255);not null;default:'';comment:'备注，描述'"`
}

func UpdateTaskData(taskDataDo *TaskData, db *gorm.DB) error {
	updateTaskData := make(map[string]interface{})
	updateTaskData["CreatedAt"] = taskDataDo.CreatedAt
	updateTaskData["UpdatedAt"] = taskDataDo.UpdatedAt
	updateTaskData["TaskId"] = taskDataDo.TaskId
	updateTaskData["StudentId"] = taskDataDo.StudentId
	updateTaskData["IdCard"] = taskDataDo.IdCard
	updateTaskData["ClassId"] = taskDataDo.ClassId
	updateTaskData["OrgId"] = taskDataDo.OrgId
	updateTaskData["LeftSph"] = taskDataDo.LeftSph
	updateTaskData["RightSph"] = taskDataDo.RightSph
	updateTaskData["LeftAp"] = taskDataDo.LeftAp
	updateTaskData["RightAp"] = taskDataDo.RightAp
	updateTaskData["LeftAxial"] = taskDataDo.LeftAxial
	updateTaskData["RightAxial"] = taskDataDo.RightAxial
	updateTaskData["ImgUrl"] = taskDataDo.ImgUrl
	updateTaskData["OpUpTime"] = taskDataDo.OpUpTime
	updateTaskData["LeftVision"] = taskDataDo.LeftVision
	updateTaskData["RightVision"] = taskDataDo.RightVision
	updateTaskData["IsGlasses"] = taskDataDo.IsGlasses
	updateTaskData["GlassesType"] = taskDataDo.GlassesType
	updateTaskData["LeftGlassesVision"] = taskDataDo.LeftGlassesVision
	updateTaskData["RightGlassesVision"] = taskDataDo.RightGlassesVision
	updateTaskData["ViUpTime"] = taskDataDo.ViUpTime
	updateTaskData["Memo"] = taskDataDo.Memo

	result := db.Model(&taskDataDo).Where("task_id = ? AND student_id = ? ", taskDataDo.TaskId, taskDataDo.StudentId).Updates(updateTaskData).Error
	return result
}

func CreateTaskData(taskDataDos *TaskData, db *gorm.DB) error {
	return db.Create(taskDataDos).Error
}

func ReadTaskDataByTaskIdAndStudentId(taskId, studentId string, db *gorm.DB) (*TaskData, error) {
	if taskId == "" || studentId == "" {
		return nil, errors.New("error param")
	}

	taskDo := new(TaskData)
	err := db.Model(&TaskData{}).Where("task_id = ? AND student_id = ?", taskId, studentId).Find(taskDo).Error

	if err != nil {
		return nil, err
	}

	return taskDo, nil
}

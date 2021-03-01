package dao

import "github.com/jinzhu/gorm"

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

func ReadTasks(db *gorm.DB) []*TaskDo {
	taskDos := make([]*TaskDo, 0)

	err := db.Limit(10).Find(&taskDos).Error

	if err != nil {
		panic("read task fail")
	}

	return taskDos
}

///////////////////////////////
type VisionDo struct {
	TaskId             string
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

	err := db.Limit(10).Find(&visionDos).Error

	if err != nil {
		panic("read vision fail")
	}

	return visionDos
}

///////////////////////////////////////
type OptometryDo struct {
	TaskId     string
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

	err := db.Limit(10).Find(&OpDos).Error

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

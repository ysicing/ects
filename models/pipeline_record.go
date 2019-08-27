package models

import (
	"encoding/json"
	"github.com/betterde/ects/internal/utils"
)

type (
	// 流水线调度记录模型
	PipelineRecords struct {
		Id         string     `xorm:"not null pk comment('ID') CHAR(36)"`
		PipelineId string     `xorm:"not null comment('流水线ID') index CHAR(36)"`
		NodeId     string     `xorm:"not null comment('节点ID') index CHAR(36)"`
		WorkerName string     `xorm:"not null comment('节点名称') VARCHAR(255)"`
		Spec       string     `xorm:"comment('定时器') CHAR(64)"`
		Status     int        `xorm:"not null default 1 comment('状态') TINYINT(1)"`
		Duration   int64      `json:"duration" xorm:"not null comment('持续时间') INT(10)"`
		BeginWith  utils.Time `json:"begin_with" xorm:"not null comment('开始于') DATETIME"`
		FinishWith utils.Time `json:"finish_with" xorm:"not null comment('结束于') DATETIME"`
		CreatedAt  utils.Time `json:"created_at" validate:"-" xorm:"not null created comment('创建于') DATETIME"`
		UpdatedAt  utils.Time `json:"updated_at" validate:"-" xorm:"not null updated comment('更新于') DATETIME"`
	}
	// 流水线执行结果
	Result struct {
		Pipeline *PipelineRecords // 流水线执行记录
		Steps    []*TaskRecords   // 流水线中任务执行记录
	}
)

// 定义模型的数据表名称
func (records *PipelineRecords) TableName() string {
	return "pipeline_records"
}

// 保存流水线记录
func (records *PipelineRecords) Store() error {
	_, err := Engine.InsertOne(records)
	return err
}

// 更新记录
func (records *PipelineRecords) Update() error {
	_, err := Engine.Update(records)
	return err
}

// 序列化
func (records *PipelineRecords) ToString() (string, error) {
	result, err := json.Marshal(records)
	return string(result), err
}

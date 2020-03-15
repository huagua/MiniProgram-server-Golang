package model

import "github.com/jinzhu/gorm"

// Record 上报记录
type Record struct {
	gorm.Model                         //其中create time 就记录了上报时间
	Reporter                  Reporter `gorm:"foreign_key:ReporterRefer;"`
	ReporterRefer             uint     //关联上报人的主键，也就是Reporter.ID
	IsReturnSchool            int      `gorm:"not null;"` //是否返校（选项）
	CurrentHealthValue        int      //当前身体状况
	CurrentContagionRiskValue int      //传染风险
	PsyStatus                 int      //心理状况
	PsyDemand                 int      //心理需求
	PsyKnowledge              int      //所需心理知识
	CurrentTemperature        int      //今日体温
	ReturnDistrictValue       int
	CurrentDistrictValue      int
	PlanCompanyDate           string
	Remarks                   string //其他信息
	ReturnTime                string //返校时间
	ReturnDormNum             string //所在宿舍号
	ReturnTrafficInfo         string //返校过程的交通信息
	ReturnDistrictPath        string //从哪里返回学校（从上级到下级按逗号分隔的字典值）
	CurrentDistrictPath       string //目前所在地（从上级到下级按逗号分隔的字典值）
}

/*
-------------------------以下为旧版表------------------------------
等getlastdata完善后即可删除
*/

// DailyInfo 上报  旧表
type DailyInfo struct {
	gorm.Model
	IsReturnSchool            string
	CurrentHealthValue        string
	CurrentContagionRiskValue string
	ReturnDistrictValue       string
	CurrentDistrictValue      string
	CurrentTemperature        string
	Remarks                   string
	PsyStatus                 string
	PsyDemand                 string
	PsyKnowledge              string
	PlanCompanyDate           string
	ReturnTime                string
	ReturnDormNum             string
	ReturnTrafficInfo         string
	UID                       string
	SaveDate                  string
	ReturnDistrictPath        string
	CurrentDistrictPath       string
}

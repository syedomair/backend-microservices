package models

// Action Public
type Action struct {
	ID                  int    `json:"id" gorm:"column:id"`
	NetworkID           int    `json:"network_id" gorm:"column:networkid"`
	WorkflowID          int    `json:"workflow_id" gorm:"column:workflowid"`
	TaskID              int    `json:"task_id" gorm:"column:taskid;default:null"`
	NextPhaseID         int    `json:"next_phase_id" gorm:"column:nextphaseid;default:null"`
	NextTaskID          int    `json:"next_task_id" gorm:"column:nexttaskid;default:null"`
	AttActionID         int    `json:"att_action_id" gorm:"column:attactionid;default:null"`
	Type                string `json:"type" gorm:"column:type"`
	TransType           string `json:"trans_type" gorm:"column:transtype"`
	IsTrans             bool   `json:"is_trans" gorm:"column:istrans"`
	Title               string `json:"title" gorm:"column:title"`
	Alias               string `json:"alias" gorm:"column:alias"`
	RequireCompleteTask bool   `json:"require_complete_task" gorm:"column:requirecompletetask"`
	Rule                string `json:"rule" gorm:"column:rule"`
	RuleCustom          string `json:"rule_custom" gorm:"column:rulecustom"`
	Status              bool   `json:"status" gorm:"column:status"`
	Autogenerate        bool   `json:"autogenerate" gorm:"column:autogenerate"`
	UpdateDate          string `json:"update_date" gorm:"column:updatedate"`
	InsertDate          string `json:"insert_date" gorm:"column:insertdate"`
}

// TableName Public
func (Action) TableName() string {
	return "action"
}

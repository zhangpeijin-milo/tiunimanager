package domain

import "github.com/pingcap/ticp/copywriting"

var FlowWorkDefineMap = map[string]*FlowWorkDefine{
	FlowCreateCluster: {
		FlowName:    FlowCreateCluster,
		StatusAlias: copywriting.DisplayByDefault(copywriting.CWFlowCreateCluster),
		TaskNodes: map[string]*TaskDefine{
			"start":        {"prepareResource", "resourceDone", "fail", SyncFuncTask, prepareResource},
			"resourceDone": {"buildConfig", "configDone", "fail", SyncFuncTask, buildConfig},
			"configDone":   {"deployCluster", "deployDone", "fail", PollingTasK, deployCluster},
			"deployDone":   {"startupCluster", "startupDone", "fail", PollingTasK, startupCluster},
			"startupDone":  {"end", "", "", SyncFuncTask, DefaultEnd},
			"fail":         {"fail", "", "", SyncFuncTask, DefaultFail},
		},
		ContextParser: func(s string) *FlowContext {
			// todo parse context
			c := make(FlowContext)
			return &c
		},
	},
	FlowDeleteCluster: {
		FlowName:    FlowDeleteCluster,
		StatusAlias: copywriting.DisplayByDefault(copywriting.CWFlowCreateCluster),
		TaskNodes: map[string]*TaskDefine{
			"start":              {"destroyTasks", "destroyTasksDone", "fail", PollingTasK, destroyTasks},
			"destroyTasksDone":   {"destroyCluster", "destroyClusterDone", "fail", PollingTasK, destroyCluster},
			"destroyClusterDone": {"deleteCluster", "deleteClusterDone", "fail", SyncFuncTask, deleteCluster},
			"deleteClusterDone":  {"freedResource", "freedResourceDone", "fail", SyncFuncTask, freedResource},
			"freedResourceDone":  {"end", "", "", SyncFuncTask, DefaultEnd},
			"fail":               {"fail", "", "", SyncFuncTask, DefaultFail},
		},
		ContextParser: func(s string) *FlowContext {
			// todo parse context
			c := make(FlowContext)
			return &c
		},
	},
	FlowExportData: {
		FlowName:    FlowExportData,
		StatusAlias: copywriting.DisplayByDefault(copywriting.CWFlowCreateCluster),
		TaskNodes: map[string]*TaskDefine{
			"start":              {"exportDataFromCluster", "exportDataDone", "fail", PollingTasK, exportDataFromCluster},
			"destroyTasksDone":   {"updateDataExportRecord", "updateRecordDone", "fail", SyncFuncTask, updateDataExportRecord},
			"updateRecordDone":   {"cleanExportTempfile", "cleanTempfileDone", "fail", SyncFuncTask, cleanExportTempfile},
			"cleanTempfileDone":  {"end", "", "", SyncFuncTask, DefaultEnd},
			"fail":               {"fail", "", "", SyncFuncTask, DefaultFail},
		},
		ContextParser: func(s string) *FlowContext {
			// todo parse context
			c := make(FlowContext)
			return &c
		},
	},
	FlowImportData: {
		FlowName:    FlowImportData,
		StatusAlias: copywriting.DisplayByDefault(copywriting.CWFlowCreateCluster),
		TaskNodes: map[string]*TaskDefine{
			"start":              {"buildDataImportConfig", "buildConfigDone", "fail", SyncFuncTask, buildDataImportConfig},
			"buildConfigDone":    {"importDataToCluster", "importDataDone", "fail", PollingTasK, importDataToCluster},
			"importDataDone":     {"updateDataImportRecord", "updateRecordDone", "fail", SyncFuncTask, updateDataImportRecord},
			"updateRecordDone":   {"cleanImportTempfile", "cleanTempfileDone", "fail", SyncFuncTask, cleanImportTempfile},
			"cleanTempfileDone":  {"end", "", "", SyncFuncTask, DefaultEnd},
			"fail":               {"fail", "", "", SyncFuncTask, DefaultFail},
		},
		ContextParser: func(s string) *FlowContext {
			// todo parse context
			c := make(FlowContext)
			return &c
		},
	},
}

type FlowWorkDefine struct {
	FlowName      string
	StatusAlias   string
	TaskNodes     map[string] *TaskDefine
	ContextParser func(string) *FlowContext
}

func (define *FlowWorkDefine) getInstance(bizId string, context map[string]interface{}) *FlowWorkAggregation{

	if context == nil {
		context = make(map[string]interface{})
	}

	return &FlowWorkAggregation{
		FlowWork: &FlowWorkEntity{
			FlowName: define.FlowName,
			StatusAlias: define.StatusAlias,
			BizId: bizId,
			Status: TaskStatusInit,
		},
		Tasks: make([]*TaskEntity, 4, 4),
		Context: context,
		Define: define,
	}
}

type TaskDefine struct {
	Name 			string
	SuccessEvent 	string
	FailEvent 		string
	ReturnType 		TaskReturnType
	Executor 		func(task *TaskEntity, context *FlowContext) bool
}

func DefaultEnd(task *TaskEntity, context *FlowContext) bool {
	return true
}

func DefaultFail(task *TaskEntity, context *FlowContext) bool {
	return true
}


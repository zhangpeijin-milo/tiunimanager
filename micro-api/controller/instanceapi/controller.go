package instanceapi

import (
	"context"
	"encoding/json"
	"github.com/pingcap-inc/tiem/library/client"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pingcap-inc/tiem/library/knowledge"
	"github.com/pingcap-inc/tiem/micro-api/controller"
	"github.com/pingcap-inc/tiem/micro-api/controller/clusterapi"
	"github.com/pingcap-inc/tiem/micro-cluster/proto"
)

// QueryParams query params of a cluster
// @Summary query params of a cluster
// @Description query params of a cluster
// @Tags cluster params
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param queryReq query ParamQueryReq false "page" default(1)
// @Param clusterId path string true "clusterId"
// @Success 200 {object} controller.ResultWithPage{data=[]ParamItem}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/{clusterId}/params [get]
func QueryParams(c *gin.Context) {
	var req ParamQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(err)
		return
	}
	clusterId := c.Param("clusterId")
	operator := controller.GetOperator(c)
	resp, err := client.ClusterClient.QueryParameters(context.TODO(), &cluster.QueryClusterParametersRequest{
		ClusterId: clusterId,
		Operator:  operator.ConvertToDTO(),
	}, controller.DefaultTimeout)

	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Fail(500, err.Error()))
	} else {
		instances := make([]ParamInstance, 0)

		err = json.Unmarshal([]byte(resp.GetParametersJson()), &instances)

		instanceMap := make(map[string]interface{})
		if len(instances) > 0 {
			for _, v := range instances {
				instanceMap[v.Name] = v.Value
			}
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, controller.Fail(500, err.Error()))
			return
		}

		parameters := make([]ParamItem, len(knowledge.ParameterKnowledge.Parameters), len(knowledge.ParameterKnowledge.Parameters))

		for i, v := range knowledge.ParameterKnowledge.Parameters {
			parameters[i] = ParamItem{
				Definition: *v,
				CurrentValue: ParamInstance{
					Name:  v.Name,
					Value: instanceMap[v.Name],
				},
			}
		}

		c.JSON(http.StatusOK, controller.SuccessWithPage(parameters, controller.Page{Page: req.Page, PageSize: req.PageSize, Total: len(parameters)}))
	}
}

// SubmitParams submit params
// @Summary submit params
// @Description submit params
// @Tags cluster params
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param updateReq body ParamUpdateReq true "update params request"
// @Param clusterId path string true "clusterId"
// @Success 200 {object} controller.CommonResult{data=ParamUpdateRsp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/{clusterId}/params [post]
func SubmitParams(c *gin.Context) {
	var req ParamUpdateReq

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	clusterId := c.Param("clusterId")
	operator := controller.GetOperator(c)
	values := req.Values

	jsonByte, _ := json.Marshal(values)

	jsonContent := string(jsonByte)

	resp, err := client.ClusterClient.SaveParameters(context.TODO(), &cluster.SaveClusterParametersRequest{
		ClusterId:      clusterId,
		ParametersJson: jsonContent,
		Operator:       operator.ConvertToDTO(),
	}, controller.DefaultTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Fail(500, err.Error()))
	} else {
		c.JSON(http.StatusOK, controller.Success(ParamUpdateRsp{
			ClusterId: clusterId,
			TaskId:    uint(resp.DisplayInfo.InProcessFlowId),
		}))
	}
}

// Backup backup
// @Summary backup
// @Description backup
// @Tags cluster backup
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param backupReq body BackupReq true "backup request"
// @Success 200 {object} controller.CommonResult{data=BackupRecord}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /backups [post]
func Backup(c *gin.Context) {

	var req BackupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	operator := controller.GetOperator(c)

	resp, err := client.ClusterClient.CreateBackup(context.TODO(), &cluster.CreateBackupRequest{
		ClusterId: req.ClusterId,
		Operator:  operator.ConvertToDTO(),
	}, controller.DefaultTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Fail(500, err.Error()))
	} else {
		c.JSON(http.StatusOK, controller.Success(BackupRecord{
			ID:     resp.BackupRecord.Id,
			Status: *clusterapi.ParseStatusFromDTO(resp.GetBackupRecord().DisplayStatus),
		}))
	}
}

// QueryBackupStrategy show the backup strategy of a cluster
// @Summary show the backup strategy of a cluster
// @Description show the backup strategy of a cluster
// @Tags cluster backup
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param clusterId path string true "clusterId"
// @Success 200 {object} controller.CommonResult{data=[]BackupStrategy}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/{clusterId}/strategy/ [get]
func QueryBackupStrategy(c *gin.Context) {
	clusterId := c.Param("clusterId")
	operator := controller.GetOperator(c)

	resp, err := client.ClusterClient.GetBackupStrategy(context.TODO(), &cluster.GetBackupStrategyRequest{
		ClusterId: clusterId,
		Operator:  operator.ConvertToDTO(),
	}, controller.DefaultTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Fail(500, err.Error()))
	} else {
		c.JSON(http.StatusOK, controller.Success([]BackupStrategy{
			{
				CronString: resp.Cron,
			},
		}))
	}
}

// SaveBackupStrategy save the backup strategy of a cluster
// @Summary save the backup strategy of a cluster
// @Description save the backup strategy of a cluster
// @Tags cluster backup
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param clusterId path string true "clusterId"
// @Param updateReq body BackupStrategyUpdateReq true "备份策略信息"
// @Success 200 {object} controller.CommonResult{data=BackupStrategy}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/{clusterId}/strategy [put]
func SaveBackupStrategy(c *gin.Context) {
	var req BackupStrategyUpdateReq
	operator := controller.GetOperator(c)
	clusterId := c.Param("clusterId")

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}
	_, err := client.ClusterClient.SaveBackupStrategy(context.TODO(), &cluster.SaveBackupStrategyRequest{
		ClusterId: clusterId,
		Operator:  operator.ConvertToDTO(),
		Cron:      req.CronString,
	}, controller.DefaultTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Fail(500, err.Error()))
	} else {
		c.JSON(http.StatusOK, controller.Success(BackupStrategy{
			CronString: req.CronString,
		}))
	}
}

// QueryBackup
// @Summary query backup records of a cluster
// @Description query backup records of a cluster
// @Tags cluster backup
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param clusterId query string true "cluster id"
// @Param request body BackupRecordQueryReq false "page" default(1)
// @Success 200 {object} controller.ResultWithPage{data=[]BackupRecord}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /backups [get]
func QueryBackup(c *gin.Context) {
	clusterId := c.Query("clusterId")

	var queryReq BackupRecordQueryReq
	if err := c.ShouldBindJSON(&queryReq); err != nil {
		_ = c.Error(err)
		return
	}
	//operator := controller.GetOperator(c)

	reqDTO := &cluster.QueryBackupRequest{
		ClusterId: clusterId,
		Page:      queryReq.PageRequest.ConvertToDTO(),
	}

	resp, err := client.ClusterClient.QueryBackupRecord(context.TODO(), reqDTO, controller.DefaultTimeout)

	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Fail(500, err.Error()))
	} else {
		records := make([]BackupRecord, len(resp.BackupRecords), len(resp.BackupRecords))

		for i, v := range resp.BackupRecords {
			records[i] = BackupRecord{
				ID:        v.Id,
				ClusterId: v.ClusterId,
				StartTime: time.Unix(v.StartTime, 0),
				EndTime:   time.Unix(v.EndTime, 0),
				Range:     int(v.Range),
				Way:       int(v.Way),
				Operator: controller.Operator{
					ManualOperator: true,
					OperatorId:     v.Operator.Id,
					//OperatorName: v.Operator.Name,
					//TenantId: v.Operator.TenantId,
				},
				Size:     v.Size,
				Status:   *clusterapi.ParseStatusFromDTO(v.DisplayStatus),
				FilePath: v.FilePath,
			}
		}
		c.JSON(http.StatusOK, controller.SuccessWithPage(records, *controller.ParsePageFromDTO(resp.Page)))
	}
}

// RecoverBackup
// @Summary recover backup record of a cluster
// @Description recover backup record of a cluster
// @Tags cluster backup
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param backupId path string true "backupId"
// @Param request body BackupRecoverReq true "backup recover request"
// @Success 200 {object} controller.CommonResult{data=controller.StatusInfo}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /backups/{backupId}/restore [post]
func RecoverBackup(c *gin.Context) {
	backupIdStr := c.Param("backupId")
	backupId, err := strconv.Atoi(backupIdStr)
	if err != nil {
		_ = c.Error(err)
		return
	}
	var req BackupRecoverReq

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	operator := controller.GetOperator(c)

	resp, err := client.ClusterClient.RecoverBackupRecord(context.TODO(), &cluster.RecoverBackupRequest{
		ClusterId:      req.ClusterId,
		Operator:       operator.ConvertToDTO(),
		BackupRecordId: int64(backupId),
	}, controller.DefaultTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Fail(500, err.Error()))
	} else {
		c.JSON(http.StatusOK, controller.Success(*clusterapi.ParseStatusFromDTO(resp.GetRecoverRecord().DisplayStatus)))
	}
}

// DeleteBackup
// @Summary delete backup record
// @Description delete backup record
// @Tags cluster backup
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param backupId path int true "backup record id"
// @Success 200 {object} controller.CommonResult{data=int}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /backups/{backupId} [delete]
func DeleteBackup(c *gin.Context) {
	backupId, _ := strconv.Atoi(c.Param("backupId"))
	operator := controller.GetOperator(c)

	_, err := client.ClusterClient.DeleteBackupRecord(context.TODO(), &cluster.DeleteBackupRequest{
		BackupRecordId: int64(backupId),
		Operator:       operator.ConvertToDTO(),
	}, controller.DefaultTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Fail(500, err.Error()))
	} else {
		c.JSON(http.StatusOK, controller.Success(backupId))
	}
}

/******************************************************************************
 * Copyright (c)  2021 PingCAP, Inc.                                          *
 * Licensed under the Apache License, Version 2.0 (the "License");            *
 * you may not use this file except in compliance with the License.           *
 * You may obtain a copy of the License at                                    *
 *                                                                            *
 * http://www.apache.org/licenses/LICENSE-2.0                                 *
 *                                                                            *
 * Unless required by applicable law or agreed to in writing, software        *
 * distributed under the License is distributed on an "AS IS" BASIS,          *
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.   *
 * See the License for the specific language governing permissions and        *
 * limitations under the License.                                             *
 *                                                                            *
 ******************************************************************************/

package management

import (
	"encoding/json"
	"github.com/pingcap-inc/tiem/message/cluster"
	"net/http"
	"strconv"
	"time"

	"github.com/pingcap-inc/tiem/micro-api/interceptor"

	"github.com/pingcap-inc/tiem/library/client/cluster/clusterpb"
	"github.com/pingcap-inc/tiem/library/common"
	"github.com/pingcap-inc/tiem/library/framework"

	"github.com/pingcap-inc/tiem/library/client"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pingcap-inc/tiem/micro-api/controller"
)

// Create create a cluster
// @Summary create a cluster
// @Description create a cluster
// @Tags cluster
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param createReq body cluster.CreateClusterReq true "create request"
// @Success 200 {object} controller.CommonResult{data=cluster.CreateClusterResp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/ [post]
func Create(c *gin.Context) {
	var req cluster.CreateClusterReq

	if requestBody, ok := controller.HandleJsonRequestFromBody(c, req); ok {
		controller.InvokeRpcMethod(c, client.ClusterClient.CreateCluster, &cluster.CreateClusterResp{},
			requestBody,
			controller.DefaultTimeout)
	}
}

// Preview preview cluster topology and capability
// @Summary preview cluster topology and capability
// @Description preview cluster topology and capability
// @Tags cluster
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param createReq body cluster.CreateClusterReq true "preview request"
// @Success 200 {object} controller.CommonResult{data=PreviewClusterRsp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/preview [post]
func Preview(c *gin.Context) {
	// todo refactor at last
	//var req cluster.CreateClusterReq
	//
	//if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
	//	_ = c.Error(err)
	//	return
	//}
	//
	//stockCheckResult := make([]StockCheckItem, 0)
	//
	//for _, group := range req.NodeDemandList {
	//	for _, node := range group.DistributionItems {
	//		stockCheckResult = append(stockCheckResult, StockCheckItem{
	//			Region:           req.Region,
	//			CpuArchitecture:  req.CpuArchitecture,
	//			Component:        *knowledge.ClusterComponentFromCode(group.ComponentType),
	//			DistributionItem: node,
	//			// todo stock
	//			Enough: true,
	//		})
	//	}
	//}
	//
	//c.JSON(http.StatusOK, controller.Success(PreviewClusterRsp{
	//	ClusterBaseInfo:     req.ClusterBaseInfo,
	//	StockCheckResult:    stockCheckResult,
	//	ClusterCommonDemand: req.ClusterCommonDemand,
	//	CapabilityIndexes:   []ServiceCapabilityIndex{
	//		//{"StorageCapability", "database storage capability", 800, "GB"},
	//		//{"TPCC", "TPCC tmpC ", 523456, ""},
	//	},
	//}))
}

// Query query clusters
// @Summary query clusters
// @Description query clusters
// @Tags cluster
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param queryReq query cluster.QueryClustersReq false "query request"
// @Success 200 {object} controller.ResultWithPage{data=[]cluster.QueryClusterResp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/ [get]
func Query(c *gin.Context) {

	// Create request
	var request QueryReq
	if err := c.ShouldBindQuery(&request); err != nil {
		framework.LogWithContext(c).Errorf("parse parameter error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	body, err := json.Marshal(request)
	if err != nil {
		framework.LogWithContext(c).Errorf("parse parameter error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call rpc method
	response := &[]ClusterDisplayInfo{}
	controller.InvokeRpcMethod(c, client.ClusterClient.QueryCluster, response, string(body), controller.DefaultTimeout)
}

// Delete delete cluster
// @Summary delete cluster
// @Description delete cluster
// @Tags cluster
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param clusterId path string true "cluster id"
// @Param deleteReq body cluster.DeleteClusterReq false "delete request"
// @Success 200 {object} controller.CommonResult{data=cluster.DeleteClusterResp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/{clusterId} [delete]
func Delete(c *gin.Context) {
	var req cluster.DeleteClusterReq

	if requestBody, ok := controller.HandleJsonRequestFromBody(c, req); ok {
		controller.InvokeRpcMethod(c, client.ClusterClient.DeleteCluster, &cluster.DeleteClusterResp{},
			requestBody,
			controller.DefaultTimeout)
	}
}

// Restart restart a cluster
// @Summary restart a cluster
// @Description restart a cluster
// @Tags cluster
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param clusterId path string true "cluster id"
// @Success 200 {object} controller.CommonResult{data=RestartClusterRsp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/{clusterId}/restart [post]
func Restart(c *gin.Context) {
	if requestBody, ok := controller.HandleJsonRequestWithBuiltReq(c, cluster.RestartClusterReq{
		ClusterID: c.Param("clusterId"),
	}); ok {
		controller.InvokeRpcMethod(c, client.ClusterClient.RestartCluster, &cluster.RestartClusterResp{},
			requestBody,
			controller.DefaultTimeout)
	}
}

// Stop stop a cluster
// @Summary stop a cluster
// @Description stop a cluster
// @Tags cluster
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param clusterId path string true "cluster id"
// @Success 200 {object} controller.CommonResult{data=cluster.StopClusterResp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/{clusterId}/stop [post]
func Stop(c *gin.Context) {
	if requestBody, ok := controller.HandleJsonRequestWithBuiltReq(c, cluster.StopClusterReq{
		ClusterID: c.Param("clusterId"),
	}); ok {
		controller.InvokeRpcMethod(c, client.ClusterClient.StopCluster, &cluster.StopClusterResp{},
			requestBody,
			controller.DefaultTimeout)
	}
}

// Detail show details of a cluster
// @Summary show details of a cluster
// @Description show details of a cluster
// @Tags cluster
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param clusterId path string true "cluster id"
// @Success 200 {object} controller.CommonResult{data=cluster.QueryClusterDetailResp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/{clusterId} [get]
func Detail(c *gin.Context) {
	if requestBody, ok := controller.HandleJsonRequestWithBuiltReq(c, &cluster.QueryClusterDetailReq{
		ClusterID: c.Param("clusterId"),
	}); ok {
		controller.InvokeRpcMethod(c, client.ClusterClient.DetailCluster, &cluster.QueryClusterDetailResp{},
			requestBody,
			controller.DefaultTimeout)
	}
}

// Takeover takeover a cluster
// @Summary takeover a cluster
// @Description takeover a cluster
// @Tags cluster
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param takeoverReq body cluster.TakeoverClusterReq true "takeover request"
// @Success 200 {object} controller.CommonResult{data=cluster.TakeoverClusterResp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/takeover [post]
func Takeover(c *gin.Context) {
	var req TakeoverReq
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = c.Error(err)
		return
	}

	operator := controller.GetOperator(c)

	reqDTO := &clusterpb.ClusterTakeoverReqDTO{
		Operator:         operator.ConvertToDTO(),
		TiupIp:           req.TiupIp,
		Port:             strconv.Itoa(req.TiupPort),
		TiupUserName:     req.TiupUserName,
		TiupUserPassword: req.TiupUserPassword,
		TiupPath:         req.TiupPath,
		ClusterNames:     req.ClusterNames,
	}

	respDTO, err := client.ClusterClient.TakeoverClusters(framework.NewMicroCtxFromGinCtx(c), reqDTO, controller.DefaultTimeout)

	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Fail(500, err.Error()))
	} else {
		status := respDTO.GetRespStatus()

		clusters := make([]ClusterDisplayInfo, len(respDTO.Clusters))

		for i, v := range respDTO.Clusters {
			clusters[i] = *ParseDisplayInfoFromDTO(v)
		}

		result := controller.BuildCommonResult(int(status.Code), status.Message, clusters)

		c.JSON(http.StatusOK, result)
	}
}

// DescribeDashboard dashboard
// @Summary dashboard
// @Description dashboard
// @Tags cluster
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param clusterId path string true "cluster id"
// @Success 200 {object} controller.CommonResult{data=DescribeDashboardRsp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/{clusterId}/dashboard [get]
func DescribeDashboard(c *gin.Context) {
	var status *clusterpb.ResponseStatusDTO
	start := time.Now()
	defer interceptor.HandleMetrics(start, "DescribeDashboard", int(status.GetCode()))

	operator := controller.GetOperator(c)
	reqDTO := &clusterpb.DescribeDashboardRequest{
		Operator:  operator.ConvertToDTO(),
		ClusterId: c.Param("clusterId"),
	}
	respDTO, err := client.ClusterClient.DescribeDashboard(framework.NewMicroCtxFromGinCtx(c), reqDTO, controller.DefaultTimeout)

	if err != nil {
		status = &clusterpb.ResponseStatusDTO{Code: http.StatusBadRequest, Message: err.Error()}
		c.JSON(http.StatusBadRequest, controller.Fail(http.StatusBadRequest, err.Error()))
	} else {
		status = respDTO.GetStatus()
		if int32(common.TIEM_SUCCESS) == status.GetCode() {
			result := controller.BuildCommonResult(int(status.Code), status.Message, DescribeDashboardRsp{
				ClusterId: respDTO.GetClusterId(),
				Url:       respDTO.GetUrl(),
				Token:     respDTO.GetToken(),
			})

			c.JSON(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, controller.Fail(int(status.GetCode()), status.GetMessage()))
		}
	}
}

// DescribeMonitor monitoring link
// @Summary monitoring link
// @Description monitoring link
// @Tags cluster
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param clusterId path string true "cluster id"
// @Success 200 {object} controller.CommonResult{data=DescribeMonitorRsp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/{clusterId}/monitor [get]
func DescribeMonitor(c *gin.Context) {
	var status *clusterpb.ResponseStatusDTO
	start := time.Now()
	defer interceptor.HandleMetrics(start, "DescribeMonitor", int(status.GetCode()))
	operator := controller.GetOperator(c)
	reqDTO := &clusterpb.DescribeMonitorRequest{
		Operator:  operator.ConvertToDTO(),
		ClusterId: c.Param("clusterId"),
	}
	respDTO, err := client.ClusterClient.DescribeMonitor(framework.NewMicroCtxFromGinCtx(c), reqDTO, controller.DefaultTimeout)

	if err != nil {
		status = &clusterpb.ResponseStatusDTO{Code: http.StatusInternalServerError, Message: err.Error()}
		c.JSON(http.StatusInternalServerError, controller.Fail(int(status.GetCode()), status.GetMessage()))
		return
	}

	status = respDTO.GetStatus()
	if int32(common.TIEM_SUCCESS) != status.GetCode() {
		c.JSON(http.StatusBadRequest, controller.Fail(int(status.GetCode()), status.GetMessage()))
		return
	}

	result := controller.BuildCommonResult(int(status.Code), status.Message, DescribeMonitorRsp{
		ClusterId:  respDTO.GetClusterId(),
		AlertUrl:   respDTO.GetAlertUrl(),
		GrafanaUrl: respDTO.GetGrafanaUrl(),
	})
	c.JSON(http.StatusOK, result)
}

// ScaleOut scale out a cluster
// @Summary scale out a cluster
// @Description scale out a cluster
// @Tags cluster
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param clusterId path string true "cluster id"
// @Param scaleOutReq body cluster.ScaleOutClusterReq true "scale out request"
// @Success 200 {object} controller.CommonResult{data=cluster.ScaleOutClusterResp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/scale-out [post]
func ScaleOut(c *gin.Context) {
	var request cluster.ScaleOutClusterReq

	// handle scale out request and call rpc method
	if body, ok := controller.HandleJsonRequestFromBody(c, request); ok {
		controller.InvokeRpcMethod(c, client.ClusterClient.ScaleOutCluster,
			&cluster.ScaleOutClusterResp{}, body, controller.DefaultTimeout)
	}
}

// ScaleIn scale in a cluster
// @Summary scale in a cluster
// @Description scale in a cluster
// @Tags cluster
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param clusterId path string true "cluster id"
// @Param scaleInReq body cluster.ScaleInClusterReq true "scale in request"
// @Success 200 {object} controller.CommonResult{data=cluster.ScaleInClusterResp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/scale-in [post]
func ScaleIn(c *gin.Context) {
	var request cluster.ScaleInClusterReq

	// handle scale in request and call rpc method
	if body, ok := controller.HandleJsonRequestFromBody(c, request); ok {
		controller.InvokeRpcMethod(c, client.ClusterClient.ScaleInCluster,
			&cluster.ScaleInClusterResp{}, body, controller.DefaultTimeout)
	}
}

// Clone clone a cluster
// @Summary clone a cluster
// @Description clone a cluster
// @Tags cluster
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param cloneClusterReq body cluster.CloneClusterReq true "clone cluster request"
// @Success 200 {object} controller.CommonResult{data=cluster.CloneClusterResp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/clone [post]
func Clone(c *gin.Context) {
	var request cluster.CloneClusterReq

	// handle clone cluster request and call rpc method
	if body, ok := controller.HandleJsonRequestFromBody(c, request); ok {
		controller.InvokeRpcMethod(c, client.ClusterClient.CloneCluster,
			&cluster.CloneClusterResp{}, body, controller.DefaultTimeout)
	}
}

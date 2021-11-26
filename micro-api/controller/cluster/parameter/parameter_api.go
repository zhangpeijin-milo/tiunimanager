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

package parameter

import (
	"net/http"

	"github.com/pingcap-inc/tiem/library/client"
	"github.com/pingcap-inc/tiem/library/framework"

	"github.com/pingcap-inc/tiem/library/client/cluster/clusterpb"

	"github.com/gin-gonic/gin"
	"github.com/pingcap-inc/tiem/micro-api/controller"
)

// QueryParams query params of a cluster
// @Summary query params of a cluster
// @Description query params of a cluster
// @Tags cluster params
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param queryReq query ParamQueryReq false "page" default(1)
// @Param clusterId path string true "clusterId"
// @Success 200 {object} controller.ResultWithPage{data=[]ListParamsResp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/{clusterId}/params [get]
func QueryParams(c *gin.Context) {
	var req ParamQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		err = c.Error(err)
		c.JSON(http.StatusInternalServerError, controller.Fail(http.StatusInternalServerError, err.Error()))
		return
	}
	clusterId := c.Param("clusterId")

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	operator := controller.GetOperator(c)
	reqDTO := &clusterpb.ListClusterParamsRequest{
		ClusterId: clusterId,
		Operator:  operator.ConvertToDTO(),
		Page:      &clusterpb.PageDTO{Page: int32(req.Page), PageSize: int32(req.PageSize)},
	}
	resp, err := client.ClusterClient.ListClusterParams(framework.NewMicroCtxFromGinCtx(c), reqDTO, controller.DefaultTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Fail(http.StatusInternalServerError, err.Error()))
		return
	}

	parameters := make([]ListParamsResp, 0)
	if resp.Params != nil {
		parameters = make([]ListParamsResp, len(resp.Params))
		err := controller.ConvertObj(resp.Params, &parameters)
		if err != nil {
			c.JSON(http.StatusInternalServerError, controller.Fail(http.StatusInternalServerError, err.Error()))
			return
		}
	}
	status := resp.RespStatus
	page := &controller.Page{Page: int(resp.Page.Page), PageSize: int(resp.Page.PageSize), Total: int(resp.Page.Total)}

	c.JSON(http.StatusOK, controller.BuildResultWithPage(int(status.Code), status.Message, page, parameters))
}

// UpdateParams update params
// @Summary submit params
// @Description submit params
// @Tags cluster params
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param updateReq body UpdateParamsReq true "update params request"
// @Param clusterId path string true "clusterId"
// @Success 200 {object} controller.CommonResult{data=ParamUpdateRsp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/{clusterId}/params [post]
func UpdateParams(c *gin.Context) {
	var req UpdateParamsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, controller.Fail(http.StatusInternalServerError, err.Error()))
		return
	}
	clusterId := c.Param("clusterId")

	params := make([]*clusterpb.UpdateClusterParamDTO, len(req.Params))
	err := controller.ConvertObj(req.Params, &params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Fail(http.StatusInternalServerError, err.Error()))
		return
	}
	operator := controller.GetOperator(c)
	reqDTO := &clusterpb.UpdateClusterParamsRequest{
		ClusterId: clusterId,
		Operator:  operator.ConvertToDTO(),
		Params:    params,
	}
	resp, err := client.ClusterClient.UpdateClusterParams(framework.NewMicroCtxFromGinCtx(c), reqDTO, controller.DefaultTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Fail(http.StatusInternalServerError, err.Error()))
		return
	}
	status := resp.RespStatus
	c.JSON(http.StatusOK, controller.BuildCommonResult(int(status.Code), status.Message,
		ParamUpdateRsp{ClusterId: clusterId}))
}

// InspectParams inspect params
// @Summary inspect params
// @Description inspect params
// @Tags cluster params
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param clusterId path string true "clusterId"
// @Success 200 {object} controller.CommonResult{data=InspectParamsResp}
// @Failure 401 {object} controller.CommonResult
// @Failure 403 {object} controller.CommonResult
// @Failure 500 {object} controller.CommonResult
// @Router /clusters/{clusterId}/params/inspect [post]
func InspectParams(c *gin.Context) {
	clusterId := c.Param("clusterId")
	operator := controller.GetOperator(c)

	reqDTO := &clusterpb.InspectClusterParamsRequest{
		ClusterId: clusterId,
		Operator:  operator.ConvertToDTO(),
	}
	resp, err := client.ClusterClient.InspectClusterParams(framework.NewMicroCtxFromGinCtx(c), reqDTO, controller.DefaultTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Fail(http.StatusInternalServerError, err.Error()))
		return
	}
	params := make([]InspectParamsResp, 0)
	if len(resp.Params) > 0 {
		params = make([]InspectParamsResp, len(resp.Params))
		err := controller.ConvertObj(resp.Params, &params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, controller.Fail(http.StatusInternalServerError, err.Error()))
			return
		}
	}
	status := resp.RespStatus
	result := controller.BuildCommonResult(int(status.Code), status.Message, params)
	c.JSON(http.StatusOK, result)
}

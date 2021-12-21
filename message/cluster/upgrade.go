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

/*******************************************************************************
 * @File: upgrade
 * @Description:
 * @Author: shenhaibo@pingcap.com
 * @Version: 1.0.0
 * @Date: 2021/12/15
*******************************************************************************/

package cluster

import "github.com/pingcap-inc/tiem/common/structs"

type QueryUpgradePathReq struct {
	ClusterID string `json:"ClusterId"`
}

type QueryUpgradePathRsp struct {
	Paths []*structs.ProductUpgradePathItem
}

type QueryUpgradeVersionDiffInfoReq struct {
	ClusterID string `json:"clusterId"`
	Version   string `json:"version"`
}

type QueryUpgradeVersionDiffInfoResp struct {
	ConfigDiffInfos []*structs.ProductUpgradeVersionConfigDiffItem `json:"configDiffInfos"`
}

type UpgradeWayStr string

const (
	UpgradeWayOffline UpgradeWayStr = "offline"
	UpgradeWayOnline  UpgradeWayStr = "online"
)

type ClusterUpgradeReq struct {
	ClusterID     string        `json:"clusterId"`
	TargetVersion string        `json:"targetVersion"`
	UpgradeWay    UpgradeWayStr `json:"upgradeWay"`
	Configs       []*structs.ClusterUpgradeVersionConfigItem
}

type ClusterUpgradeResp struct {
	structs.AsyncTaskWorkFlowInfo
}

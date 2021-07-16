package clusterapi

import (
	"github.com/pingcap/ticp/micro-api/controller"
	cluster "github.com/pingcap/ticp/micro-cluster/proto"
)

func (req *CreateReq) ConvertToDTO() (baseInfoDTO *cluster.ClusterBaseInfoDTO, demandsDTO []*cluster.ClusterNodeDemandDTO) {
	baseInfoDTO = req.ClusterBaseInfo.ConvertToDTO()

	demandsDTO = make([]*cluster.ClusterNodeDemandDTO, len(req.NodeDemandList), len(req.NodeDemandList))

	return
}

func (baseInfo *ClusterBaseInfo) ConvertToDTO() (dto *cluster.ClusterBaseInfoDTO) {
	dto = &cluster.ClusterBaseInfoDTO{
		ClusterName: baseInfo.ClusterName,

		DbPasswordName : baseInfo.DbPassword,
		ClusterType: controller.ConvertTypeDTO(baseInfo.ClusterType),
		ClusterVersion: controller.ConvertVersionDTO(baseInfo.ClusterVersion),
		Tags: baseInfo.Tags,
		Tls: baseInfo.Tls,
	}

	return
}

func (demand *ClusterNodeDemand) ConvertToDTO() (dto *cluster.ClusterNodeDemandDTO) {
	items := make([]*cluster.DistributionItemDTO, len(demand.DistributionItems), len(demand.DistributionItems))

	for index, item := range demand.DistributionItems {
		items[index] = item.ConvertToDTO()
	}

	dto = &cluster.ClusterNodeDemandDTO{
		ComponentType: demand.ComponentType,

		TotalNodeCount: int32(demand.TotalNodeCount),
		Items: items,
	}
	return
}

func (item DistributionItem) ConvertToDTO() (dto *cluster.DistributionItemDTO) {
	dto = &cluster.DistributionItemDTO{
		ZoneCode: item.ZoneCode,
		SpecCode: item.SpecCode,
		Count: int32(item.Count),
	}
	return
}

package adapt

import (
	"context"
	"fmt"
	"github.com/pingcap-inc/tiem/library/client/cluster/clusterpb"
	"github.com/pingcap-inc/tiem/library/knowledge"
	"github.com/pingcap-inc/tiem/micro-cluster/service/cluster/domain"
	"github.com/pingcap/tiup/pkg/cluster/spec"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net"
)

type TiUPTiDBMetadataManager struct {
	componentParsers map[string]domain.ComponentParser
}

func NewTiUPTiDBMetadataManager() *TiUPTiDBMetadataManager {
	mgr := new(TiUPTiDBMetadataManager)
	mgr.componentParsers = map[string]domain.ComponentParser {}

	parser := TiDBComponentParser{}
	mgr.componentParsers[parser.GetComponent().ComponentType] = parser

	parser2 := TiKVComponentParser{}
	mgr.componentParsers[parser2.GetComponent().ComponentType] = parser2

	parser3 := PDComponentParser{}
	mgr.componentParsers[parser3.GetComponent().ComponentType] = parser2

	parser4 := TiFlashComponentParser{}
	mgr.componentParsers[parser4.GetComponent().ComponentType] = parser2

	return mgr
}

func (t TiUPTiDBMetadataManager) FetchFromRemoteCluster(ctx context.Context, req *clusterpb.ClusterTakeoverReqDTO) (spec.Metadata, error) {
	Conf := ssh.ClientConfig{User: req.TiupUserName,
		Auth: []ssh.AuthMethod{ssh.Password(req.TiupUserPassword)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	Client, err := ssh.Dial("tcp", net.JoinHostPort(req.TiupIp, req.Port), &Conf)
	if err != nil {
		return nil, err
	}
	defer Client.Close()

	sftpClient, err := sftp.NewClient(Client)
	if err != nil {
		return nil, err
	}
	defer sftpClient.Close()

	remoteFileName := fmt.Sprintf("%sstorage/cluster/clusters/%s/meta.yaml", req.TiupPath, req.ClusterNames[0])
	remoteFile, err := sftpClient.Open(remoteFileName)
	if err != nil {
		return nil, err
	}
	defer remoteFile.Close()
	if err != nil {
		return nil, err
	}
	dataByte, err := ioutil.ReadAll(remoteFile)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}

	metadata := &spec.ClusterMeta{}
	yaml.Unmarshal(dataByte, metadata)

	return metadata, nil
}

func (t TiUPTiDBMetadataManager) RebuildMetadataFromComponents(cluster *domain.Cluster, components []*domain.ComponentGroup) (spec.Metadata, error) {
	panic("implement me")
}

func (t TiUPTiDBMetadataManager) ParseComponentsFromMetaData(metadata spec.Metadata) ([]*domain.ComponentGroup, error) {
	version := metadata.GetBaseMeta().Version

	clusterSpec := metadata.GetTopology().(*spec.Specification)
	componentsList := knowledge.GetComponentsForCluster("TiDB", version)

	componentGroups := make([]*domain.ComponentGroup, 0)
	for _, component := range componentsList {
		componentGroups = append(componentGroups, t.componentParsers[component.ComponentType].ParseComponent(clusterSpec))
	}

	return componentGroups, nil
}

func (t TiUPTiDBMetadataManager) ParseClusterInfoFromMetaData(meta spec.BaseMeta) (user string, group string, version string) {
	return meta.User, meta.Group, meta.Version
}

type TiDBComponentParser struct {}

func (t TiDBComponentParser) GetComponent() *knowledge.ClusterComponent {
	return knowledge.ClusterComponentFromCode("TiDB")
}

func (t TiDBComponentParser) ParseComponent(spec *spec.Specification) *domain.ComponentGroup {
	group := &domain.ComponentGroup{
		ComponentType: t.GetComponent(),
		Nodes: make([]domain.ComponentInstance, 0),
	}

	for _, tidbServer := range spec.TiDBServers {
		componentInstance := domain.ComponentInstance {
			ComponentType: t.GetComponent(),
			Ip: tidbServer.Host,

		}
		group.Nodes = append(group.Nodes, componentInstance)
	}

	return group
}

type TiKVComponentParser struct {}

func (t TiKVComponentParser) GetComponent() *knowledge.ClusterComponent {
	return knowledge.ClusterComponentFromCode("TiKV")
}

func (t TiKVComponentParser) ParseComponent(spec *spec.Specification) *domain.ComponentGroup {
	panic("implement me")
}

type PDComponentParser struct {}

func (t PDComponentParser) GetComponent() *knowledge.ClusterComponent {
	return knowledge.ClusterComponentFromCode("PD")
}

func (t PDComponentParser) ParseComponent(spec *spec.Specification) *domain.ComponentGroup {
	panic("implement me")
}

type TiFlashComponentParser struct {}

func (t TiFlashComponentParser) GetComponent() *knowledge.ClusterComponent {
	return knowledge.ClusterComponentFromCode("TiFlash")
}

func (t TiFlashComponentParser) ParseComponent(spec *spec.Specification) *domain.ComponentGroup {
	panic("implement me")
}

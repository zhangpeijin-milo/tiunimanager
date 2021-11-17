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

package knowledge

import (
	"github.com/pingcap-inc/tiem/library/common/resource-type"
	"github.com/pingcap-inc/tiem/library/framework"
)

var SpecKnowledge *ClusterSpecKnowledge
var ParameterKnowledge *ClusterParameterKnowledge

type ClusterParameterKnowledge struct {
	Parameters []*Parameter
	Names      map[string]*Parameter
}

type ClusterSpecKnowledge struct {
	Specs      []*ClusterTypeSpec
	Types      map[string]*ClusterType
	Versions   map[string]*ClusterVersion
	Components map[string]*ClusterComponent
}

type ResourceSpecKnowledge struct {
	Specs []*ResourceSpec
}

func ClusterTypeSpecFromCode(code string) *ClusterTypeSpec {
	for _, v := range SpecKnowledge.Specs {
		if code == v.ClusterType.Code {
			return v
		}
	}
	framework.Log().Errorf("Unexpected code of cluster type: %s", code)
	return nil
}

func GetComponentPortRange(typeCode string, versionCode string, componentType string) *ComponentPortConstraint {
	clusterTypeSpec := ClusterTypeSpecFromCode(typeCode)
	if clusterTypeSpec == nil {
		framework.Log().Errorf("Unexpected code of cluster code: %s", typeCode)
		return nil
	}
	versionSpec := clusterTypeSpec.GetVersionSpec(versionCode)
	if versionSpec == nil {
		framework.Log().Errorf("Unexpected code of version code: %s", versionCode)
		return nil
	}
	componentSpec := versionSpec.GetComponentSpec(componentType)
	if componentSpec == nil {
		framework.Log().Errorf("Unexpected code of component type: %s", componentType)
		return nil
	}
	return &componentSpec.PortConstraint
}

func ClusterTypeFromCode(code string) *ClusterType {
	return SpecKnowledge.Types[code]
}

func ClusterVersionFromCode(code string) *ClusterVersion {
	return SpecKnowledge.Versions[code]
}

func ClusterComponentFromCode(componentType string) *ClusterComponent {
	return SpecKnowledge.Components[componentType]
}

func ParameterFromName(name string) *Parameter {
	return ParameterKnowledge.Names[name]
}

func LoadKnowledge() {
	loadSpecKnowledge()
	loadParameterKnowledge()
}

func loadSpecKnowledge() {
	tidbType := ClusterType{"TiDB", "TiDB"}
	tidbV4_0_12 := ClusterVersion{"v4.0.12", "v4.0.12"}
	tidbV5_0_0 := ClusterVersion{"v5.0.0", "v5.0.0"}

	tidbComponent := ClusterComponent{
		"TiDB", "compute", "TiDB",
	}

	tikvComponent := ClusterComponent{
		"TiKV", "storage", "TiKV",
	}

	pdComponent := ClusterComponent{
		"PD", "dispatch", "PD",
	}

	tiFlashComponent := ClusterComponent{
		"TiFlash", "column-based storage","TiFlash",
	}

	grafanaComponent := ClusterComponent{
		"Grafana", "dispatch","Grafana",
	}
	monitorComponent := ClusterComponent{
		"Prometheus", "dispatch","Prometheus",
	}
	alertMangerComponent := ClusterComponent{
		"AlertManger", "dispatch","AlertManger",
	}

	//tiCdcComponent := ClusterComponent{
	//	"TiCDC", "TiCDC",
	//}

	tidbV4_0_12_Spec := ClusterVersionSpec{
		ClusterVersion: tidbV4_0_12,
		ArchTypes: []resource.ArchType{
			resource.Arm64,
			resource.X86,
		},
		ComponentSpecs: []ClusterComponentSpec{
			{tidbComponent, ComponentConstraint{true, false, []int{3}, []string{
				GenSpecCode(4, 8),
				GenSpecCode(8, 16),
				GenSpecCode(8, 32),
				GenSpecCode(16, 32),
			}, 1},
				ComponentPortConstraint{10000, 10020, 2},
			},
			{tikvComponent, ComponentConstraint{true, false, []int{3}, []string{
				GenSpecCode(8, 32),
				GenSpecCode(8, 64),
				GenSpecCode(16, 128),
			}, 1},
				ComponentPortConstraint{10020, 10040, 2},
			},
			{pdComponent, ComponentConstraint{true, false, []int{3}, []string{
				GenSpecCode(4, 8),
				GenSpecCode(8, 16),
			}, 1},
				ComponentPortConstraint{10040, 10100, 6},
			},
			{tiFlashComponent, ComponentConstraint{false, false, []int{3}, []string{
				GenSpecCode(4, 32),
				GenSpecCode(8, 64),
				GenSpecCode(16, 128),
			}, 0},
				ComponentPortConstraint{10100, 10160, 6},
			},
			{grafanaComponent, ComponentConstraint{false, true, []int{1}, []string{
			}, 0},
				ComponentPortConstraint{0, 0, 0},
			},
			{monitorComponent, ComponentConstraint{false, true, []int{1}, []string{
			}, 0},
				ComponentPortConstraint{0, 0, 0},
			},
			{alertMangerComponent, ComponentConstraint{false, true, []int{1}, []string{
			}, 0},
				ComponentPortConstraint{0, 0, 0},
			},
		},
	}
	tidbV5_0_0_Spec := ClusterVersionSpec{
		ClusterVersion: tidbV5_0_0,
		ArchTypes: []resource.ArchType{
			resource.Arm64,
			resource.X86,
		},
		ComponentSpecs: []ClusterComponentSpec{
			{tidbComponent, ComponentConstraint{true, false, []int{3}, []string{
				GenSpecCode(4, 8),
				GenSpecCode(8, 16),
				GenSpecCode(8, 32),
				GenSpecCode(16, 32),
			}, 1},
				ComponentPortConstraint{10000, 10020, 2},
			},
			{tikvComponent, ComponentConstraint{true, false, []int{3}, []string{
				GenSpecCode(8, 32),
				GenSpecCode(8, 64),
				GenSpecCode(16, 128),
			}, 1},
				ComponentPortConstraint{10020, 10040, 2},
			},
			{pdComponent, ComponentConstraint{true, false, []int{3}, []string{
				GenSpecCode(4, 8),
				GenSpecCode(8, 16),
			}, 1},
				ComponentPortConstraint{10040, 10060, 2},
			},
			{tiFlashComponent, ComponentConstraint{false, false, []int{3}, []string{
				GenSpecCode(4, 32),
				GenSpecCode(8, 64),
				GenSpecCode(16, 128),
			}, 0},
				ComponentPortConstraint{10060, 10120, 6},
			},
			{grafanaComponent, ComponentConstraint{false, true, []int{1}, []string{
			}, 0},
				ComponentPortConstraint{0, 0, 0},
			},
			{monitorComponent, ComponentConstraint{false, true, []int{1}, []string{
			}, 0},
				ComponentPortConstraint{0, 0, 0},
			},
			{alertMangerComponent, ComponentConstraint{false, true, []int{1}, []string{
			}, 0},
				ComponentPortConstraint{0, 0, 0},
			},
		},
	}

	SpecKnowledge = &ClusterSpecKnowledge{
		Specs:    []*ClusterTypeSpec{{tidbType, []ClusterVersionSpec{tidbV4_0_12_Spec, tidbV5_0_0_Spec}}},
		Types:    map[string]*ClusterType{tidbType.Code: &tidbType},
		Versions: map[string]*ClusterVersion{tidbV4_0_12.Code: &tidbV4_0_12, tidbV5_0_0.Code: &tidbV5_0_0},
		Components: map[string]*ClusterComponent{tidbComponent.ComponentType: &tidbComponent,
			tikvComponent.ComponentType:    &tikvComponent,
			pdComponent.ComponentType:      &pdComponent,
			tiFlashComponent.ComponentType: &tiFlashComponent,
			grafanaComponent.ComponentType: &grafanaComponent,
			monitorComponent.ComponentType: &monitorComponent,
			alertMangerComponent.ComponentType: &alertMangerComponent,
		},
	}
}

func loadParameterKnowledge() {
	ParameterKnowledge = &ClusterParameterKnowledge{
		Parameters: []*Parameter{
			{"binlog_cache_size", ParamTypeInteger, ParamUnitMb, 1024, false, []ParamValueConstraint{
				{ConstraintTypeGT, 0},
				{ConstraintTypeGT, 2048},
			}, "binlog cache size"},
			{"show_query_log", ParamTypeBoolean, ParamUnitNil, true, false, []ParamValueConstraint{}, "show_query_log"},
		},
	}

	ParameterKnowledge.Names = make(map[string]*Parameter)
	for _, v := range ParameterKnowledge.Parameters {
		ParameterKnowledge.Names[v.Name] = v
	}
}

// GetComponentsForCluster
// @Description: list all the components for a
// @Parameter clusterCode
// @Parameter versionCode
// @return []ClusterComponent
func GetComponentsForCluster(clusterCode, versionCode string) []ClusterComponent {
	versionSpec := ClusterTypeSpecFromCode(clusterCode).GetVersionSpec(versionCode)
	components := make([]ClusterComponent, len(versionSpec.ComponentSpecs))

	for i, v := range versionSpec.ComponentSpecs {
		components[i] = v.ClusterComponent
	}

	return components
}

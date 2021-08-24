package framework

import (
	"github.com/pingcap/tiem/library/firstparty/config"
	"github.com/pingcap/tiem/library/firstparty/util"
	"github.com/pingcap/tiem/library/thirdparty/logger"
)

type UtOpt func(d *UtFramework) error

type UtFramework struct {
	serviceEnum 		MicroServiceEnum

	initOpts 			[]UtOpt
	log 				*logger.LogRecord
}

func NewUtFramework(serviceName MicroServiceEnum, initOpt ...UtOpt) *UtFramework {
	p := &UtFramework{
		serviceEnum: serviceName,
		initOpts: []UtOpt{
			func(d *UtFramework) error {
				config.InitForMonolith(d.serviceEnum.logMod())
				return nil
			},
			func(d *UtFramework) error {
				d.log = d.serviceEnum.buildLogger()
				return nil
			},
		},
	}


	p.initOpts = append(p.initOpts, initOpt...)

	p.Init()

	return p
}

func (u UtFramework) Init() error {
	for _, opt := range u.initOpts {
		util.AssertNoErr(opt(&u))
	}
	return nil
}

func (u UtFramework) StartService() error {
	return nil
}

func (u UtFramework) GetDefaultLogger() *logger.LogRecord {
	return u.log
}

func (u UtFramework) GetRegistryAddress() []string {
	panic("implement me")
}

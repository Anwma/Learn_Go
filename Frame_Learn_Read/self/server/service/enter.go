package service

import (
	"self/server/service/example"
	"self/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	SystemServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

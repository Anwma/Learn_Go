package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.119.128:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://192.168.119.127:8021/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	//client.Agent().ServiceDeregister()
	if err != nil {
		panic(err)
	}
	return nil
}

// AllServices correct
func AllServices() {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.119.128:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}

// FilterSerivice correct
func FilterSerivice() {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.119.128:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(`Service == "user-web"`)
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}

func main() {
	//由于视频中所演示的是1.15.3,本人所装的是1.18.4 版本之间的consul包可能有冲突...导致服务检查不通过... error
	//_ = Register("192.168.119.127", 8021, "user-web", []string{"mxshop", "bobby"}, "user-web")
	//AllServices()
	FilterSerivice()
	//fmt.Println(fmt.Sprintf(`Service == "%s"`, "user-web"))
}

package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/proto"
)

func InitSrvConn() {
	//从注册中心获取到用户服务的信息
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	userSrvHost := ""
	userSrvPort := 0

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//三种方式人选其一 ``不用加\作为转义符
	//data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvInfo.Name))
	//data, err := client.Agent().ServicesWithFilter(`Service == "user-srv"`)

	if err != nil {
		panic(err)
	}
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
		//fmt.Println(key)
	}

	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
		return
	}

	//拨号连接用户 grpc 服务器  跨域的问题 -- 后端解决 也可以前端解决
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败】",
			"msg", err.Error(),
		)
	}
	//1.后续的用户服务下线了 2.改端口了 3.改ip了 负载均衡来做
	//2.已经事先创立好了连接 这样后续就不用进行再次tcp的三次握手
	//3.问题：一个连接多个goroutine共用 性能问题 - 连接池
	//生成grpc的client并调用接口
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient

}

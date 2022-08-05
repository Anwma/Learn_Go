package api

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"net/http"
	"strings"
	"time"
)

/*
	短信服务需要备案域名+个人云上公司+账号转为企业认证...
	短信服务: https://dysms.console.aliyun.com/domestic/text
	个人认证转企业认证: https://help.aliyun.com/document_detail/37178.html?spm=5176.25163407.domtextsigncreate-index-1ec3c_58c50_0.2.772fbb6e5vmwE8
	个人云上公司注册：https://wanwang.aliyun.com/companyreg/ysgszc
	... 到时候得去买服务器 买域名 备案域名 然后才能测试....好麻烦...
*/

func GenerateSmsCode(width int) string {
	//生成width长度的短信验证码
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(c *gin.Context) {
	sendSmsForm := forms.SendSmsForm{}
	if err := c.ShouldBind(&sendSmsForm); err != nil {
		HandleValidatorError(c, err)
		return
	}

	//此次已将个人的 AccessKey ID 和 AccessKey Secret 重置为xxx,后续人员需要使用时请填入个人的id
	client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", global.ServerConfig.AliSmsInfo.ApiKey, global.ServerConfig.AliSmsInfo.ApiSecret)
	if err != nil {
		panic(err)
	}
	smsCode := GenerateSmsCode(6)
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-beijing"
	request.QueryParams["PhoneNumbers"] = sendSmsForm.Mobile            //手机号
	request.QueryParams["SignName"] = "慕学在线"                            //阿里云验证过的项目名 自己设置 签名
	request.QueryParams["TemplateCode"] = "xxx"                         //阿里云的短信模板号 自己设置
	request.QueryParams["TemplateParam"] = "{\"code\":" + smsCode + "}" //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。
	response, err := client.ProcessCommonRequest(request)
	fmt.Print(client.DoAction(request, response))
	if err != nil {
		fmt.Print(err.Error())
	}
	//后面注册的时候会将短信验证码带回来重新注册
	//将验证码保存起来 保存到 redis 当中
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	rdb.Set(context.Background(), sendSmsForm.Mobile, smsCode, time.Duration( global.ServerConfig.RedisInfo.Expire)*time.Second)

	c.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})

}

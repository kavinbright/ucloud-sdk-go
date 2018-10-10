//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api PathX CreateGlobalSSHInstance

package pathx

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// CreateGlobalSSHInstanceRequest is request schema for CreateGlobalSSHInstance action
type CreateGlobalSSHInstanceRequest struct {
	request.CommonBase

	// 填写支持SSH访问IP的地区名称，如“洛杉矶”，“新加坡”，“香港”，“东京”，“华盛顿”，“法兰克福”。Area和AreaCode两者必填一个
	Area *string `required:"true"`

	// 被SSH访问的IP
	TargetIP *string `required:"true"`

	// SSH端口，1-65535且不能使用80，443端口
	Port *int `required:"true"`

	// AreaCode, 区域航空港国际通用代码。Area和AreaCode两者必填一个
	AreaCode *string `required:"true"`

	// 备注信息
	Remark *string `required:"false"`

	// 支付方式，如按月、按年、按时
	ChargeType *string `required:"false"`

	// 购买数量
	Quantity *int `required:"false"`

	// 使用代金券可冲抵部分费用
	CouponId *string `required:"false"`
}

// CreateGlobalSSHInstanceResponse is response schema for CreateGlobalSSHInstance action
type CreateGlobalSSHInstanceResponse struct {
	response.CommonBase

	// 实例ID，资源唯一标识
	InstanceId string

	// 加速域名，访问该域名可就近接入
	AcceleratingDomain string

	// 提示信息
	Message string
}

// NewCreateGlobalSSHInstanceRequest will create request of CreateGlobalSSHInstance action.
func (c *PathXClient) NewCreateGlobalSSHInstanceRequest() *CreateGlobalSSHInstanceRequest {
	req := &CreateGlobalSSHInstanceRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// CreateGlobalSSHInstance - 创建GlobalSSH实例
func (c *PathXClient) CreateGlobalSSHInstance(req *CreateGlobalSSHInstanceRequest) (*CreateGlobalSSHInstanceResponse, error) {
	var err error
	var res CreateGlobalSSHInstanceResponse

	err = c.client.InvokeAction("CreateGlobalSSHInstance", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

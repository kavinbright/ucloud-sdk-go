// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// ReinstallVMInstanceRequest is request schema for ReinstallVMInstance action
type ReinstallVMInstanceRequest struct {
	request.CommonBase

	// [公共参数] 地域。枚举值： cn，表示中国；
	// Region *string `required:"true"`

	// [公共参数] 可用区。枚举值：zone-01，表示中国；
	// Zone *string `required:"true"`

	// 镜像ID
	ImageID *string `required:"true"`

	// 虚拟机ID
	VMID *string `required:"true"`
}

// ReinstallVMInstanceResponse is response schema for ReinstallVMInstance action
type ReinstallVMInstanceResponse struct {
	response.CommonBase

	// 返回信息描述
	Message string
}

// NewReinstallVMInstanceRequest will create request of ReinstallVMInstance action.
func (c *UCloudStackClient) NewReinstallVMInstanceRequest() *ReinstallVMInstanceRequest {
	req := &ReinstallVMInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ReinstallVMInstance - 重装系统，关机的虚拟机才可以重装系统
func (c *UCloudStackClient) ReinstallVMInstance(req *ReinstallVMInstanceRequest) (*ReinstallVMInstanceResponse, error) {
	var err error
	var res ReinstallVMInstanceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ReinstallVMInstance", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

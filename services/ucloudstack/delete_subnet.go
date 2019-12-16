// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteSubnetRequest is request schema for DeleteSubnet action
type DeleteSubnetRequest struct {
	request.CommonBase

	// [公共参数] 地域。
	// Region *string `required:"true"`

	// [公共参数] 可用区。
	// Zone *string `required:"true"`

	// SubnetID
	SubnetID *string `required:"true"`
}

// DeleteSubnetResponse is response schema for DeleteSubnet action
type DeleteSubnetResponse struct {
	response.CommonBase

	// 返回信息描述；
	Message string
}

// NewDeleteSubnetRequest will create request of DeleteSubnet action.
func (c *UCloudStackClient) NewDeleteSubnetRequest() *DeleteSubnetRequest {
	req := &DeleteSubnetRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteSubnet - 删除子网
func (c *UCloudStackClient) DeleteSubnet(req *DeleteSubnetRequest) (*DeleteSubnetResponse, error) {
	var err error
	var res DeleteSubnetResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteSubnet", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeSubnetRequest is request schema for DescribeSubnet action
type DescribeSubnetRequest struct {
	request.CommonBase

	// [公共参数] 地域。
	// Region *string `required:"true"`

	// [公共参数] 可用区。
	// Zone *string `required:"true"`

	// 返回数据长度，默认为20，最大100。
	Limit *int `required:"false"`

	// 列表起始位置偏移量，默认为0。
	Offset *int `required:"false"`

	// 【数组】子网 ID。调用方式举例：SubnetIDs.0=“one-id”、SubnetIDs.1=“two-id”
	SubnetIDs []string `required:"false"`

	// VPCID
	VPCID *string `required:"false"`
}

// DescribeSubnetResponse is response schema for DescribeSubnet action
type DescribeSubnetResponse struct {
	response.CommonBase

	// 子网信息列表
	Infos []SubnetInfo

	// 返回信息描述；
	Message string

	// 子网的总数
	TotalCount int
}

// NewDescribeSubnetRequest will create request of DescribeSubnet action.
func (c *UCloudStackClient) NewDescribeSubnetRequest() *DescribeSubnetRequest {
	req := &DescribeSubnetRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeSubnet - 查询子网信息
func (c *UCloudStackClient) DescribeSubnet(req *DescribeSubnetRequest) (*DescribeSubnetResponse, error) {
	var err error
	var res DescribeSubnetResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeSubnet", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

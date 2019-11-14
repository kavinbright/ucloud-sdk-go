// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeDiskRequest is request schema for DescribeDisk action
type DescribeDiskRequest struct {
	request.CommonBase

	// [公共参数] 地域。枚举值： cn，表示中国；
	// Region *string `required:"true"`

	// [公共参数] 可用区。枚举值：zone-01，表示中国；
	// Zone *string `required:"true"`

	// 【数组】磁盘的 ID。输入有效的 ID。调用方式举例：DiskIDs.0=“one-id”、DiskIDs.1=“two-id”。
	DiskIDs []string `required:"true"`

	// 返回数据长度，默认为20，最大100。
	Limit *int `required:"true"`

	// 列表起始位置偏移量，默认为0。
	Offset *int `required:"true"`
}

// DescribeDiskResponse is response schema for DescribeDisk action
type DescribeDiskResponse struct {
	response.CommonBase

	// 【数组】返回虚拟机对象数组
	Infos []DiskInfo

	// 返回信息描述。
	Message string

	// 返回码。
	RetCode int

	// 返回虚拟机总个数。
	TotalCount int
}

// NewDescribeDiskRequest will create request of DescribeDisk action.
func (c *UCloudStackClient) NewDescribeDiskRequest() *DescribeDiskRequest {
	req := &DescribeDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeDisk - 获取UCloudStack硬盘信息
func (c *UCloudStackClient) DescribeDisk(req *DescribeDiskRequest) (*DescribeDiskResponse, error) {
	var err error
	var res DescribeDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
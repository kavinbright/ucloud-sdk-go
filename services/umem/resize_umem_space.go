//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem ResizeUMemSpace

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// ResizeUMemSpaceRequest is request schema for ResizeUMemSpace action
type ResizeUMemSpaceRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// UMem 内存空间Id
	SpaceId *string `required:"true"`

	// 内存大小, 单位:GB (需要大于原size,<= 1024)
	Size *int `required:"true"`

	//
	ChargeType *string `required:"true"`

	// 空间类型:single(无热备),double(热备)(默认: double)
	Type *string `required:"true"`

	// 使用的代金券Id
	CouponId *string `required:"false"`
}

// ResizeUMemSpaceResponse is response schema for ResizeUMemSpace action
type ResizeUMemSpaceResponse struct {
	response.CommonBase
}

// NewResizeUMemSpaceRequest will create request of ResizeUMemSpace action.
func (c *UMemClient) NewResizeUMemSpaceRequest() *ResizeUMemSpaceRequest {
	req := &ResizeUMemSpaceRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ResizeUMemSpace - 调整内存空间容量
func (c *UMemClient) ResizeUMemSpace(req *ResizeUMemSpaceRequest) (*ResizeUMemSpaceResponse, error) {
	var err error
	var res ResizeUMemSpaceResponse

	err = c.client.InvokeAction("ResizeUMemSpace", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
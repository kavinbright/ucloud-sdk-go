//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDisk SetUDiskUDataArkMode

package udisk

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// SetUDiskUDataArkModeRequest is request schema for SetUDiskUDataArkMode action
type SetUDiskUDataArkModeRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// 需要设置数据方舟的UDisk的Id
	UDiskId *string `required:"true"`

	// 是否开启数据方舟，开启:"Yes", 不支持:"No"
	UDataArkMode *string `required:"true"`
}

// SetUDiskUDataArkModeResponse is response schema for SetUDiskUDataArkMode action
type SetUDiskUDataArkModeResponse struct {
	response.CommonBase
}

// NewSetUDiskUDataArkModeRequest will create request of SetUDiskUDataArkMode action.
func (c *UDiskClient) NewSetUDiskUDataArkModeRequest() *SetUDiskUDataArkModeRequest {
	req := &SetUDiskUDataArkModeRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// SetUDiskUDataArkMode - 设置UDisk数据方舟的状态
func (c *UDiskClient) SetUDiskUDataArkMode(req *SetUDiskUDataArkModeRequest) (*SetUDiskUDataArkModeResponse, error) {
	var err error
	var res SetUDiskUDataArkModeResponse

	err = c.client.InvokeAction("SetUDiskUDataArkMode", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

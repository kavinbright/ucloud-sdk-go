// Code is generated by ucloud-model, DO NOT EDIT IT.

package ipsecvpn

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// GetVPNGatewayPriceRequest is request schema for GetVPNGatewayPrice action
type GetVPNGatewayPriceRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 付费方式, 枚举值为: Year, 按年付费; Month, 按月付费; Dynamic, 按需付费(需开启权限); 默认为获取三种价格
	ChargeType *string `required:"false"`

	// VPN网关规格。枚举值，包括：标准型：Standard，增强型：Enhanced。
	Grade *string `required:"true"`

	// 购买时长, 默认: 1
	Quantity *int `required:"false"`
}

// GetVPNGatewayPriceResponse is response schema for GetVPNGatewayPrice action
type GetVPNGatewayPriceResponse struct {
	response.CommonBase

	// 获取的VPN网关价格信息列表，每项参数详见 VPNGatewayPriceSet
	PriceSet []VPNGatewayPriceSet
}

// NewGetVPNGatewayPriceRequest will create request of GetVPNGatewayPrice action.
func (c *IPSecVPNClient) NewGetVPNGatewayPriceRequest() *GetVPNGatewayPriceRequest {
	req := &GetVPNGatewayPriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// GetVPNGatewayPrice - 获取VPN价格
func (c *IPSecVPNClient) GetVPNGatewayPrice(req *GetVPNGatewayPriceRequest) (*GetVPNGatewayPriceResponse, error) {
	var err error
	var res GetVPNGatewayPriceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetVPNGatewayPrice", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

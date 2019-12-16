// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteSecurityGroupRuleRequest is request schema for DeleteSecurityGroupRule action
type DeleteSecurityGroupRuleRequest struct {
	request.CommonBase

	// [公共参数] 地域。
	// Region *string `required:"true"`

	// [公共参数] 可用区。
	// Zone *string `required:"true"`

	// 安全组ID
	SGID *string `required:"true"`

	// 安全组规则ID
	SGRuleID *string `required:"true"`
}

// DeleteSecurityGroupRuleResponse is response schema for DeleteSecurityGroupRule action
type DeleteSecurityGroupRuleResponse struct {
	response.CommonBase

	// 返回信息描述；
	Message string
}

// NewDeleteSecurityGroupRuleRequest will create request of DeleteSecurityGroupRule action.
func (c *UCloudStackClient) NewDeleteSecurityGroupRuleRequest() *DeleteSecurityGroupRuleRequest {
	req := &DeleteSecurityGroupRuleRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteSecurityGroupRule - 删除安全组规则
func (c *UCloudStackClient) DeleteSecurityGroupRule(req *DeleteSecurityGroupRuleRequest) (*DeleteSecurityGroupRuleResponse, error) {
	var err error
	var res DeleteSecurityGroupRuleResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteSecurityGroupRule", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

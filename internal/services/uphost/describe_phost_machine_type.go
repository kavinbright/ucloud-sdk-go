//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UPHost DescribePHostMachineType

package uphost

import (
	"github.com/ucloud/ucloud-sdk-go/private/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribePHostMachineTypeRequest is request schema for DescribePHostMachineType action
type DescribePHostMachineTypeRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 具体机型。若不填写，则返回全部机型
	Type *string `required:"false"`
}

// DescribePHostMachineTypeResponse is response schema for DescribePHostMachineType action
type DescribePHostMachineTypeResponse struct {
	response.CommonBase

	// 机型列表，模型：PHostMachineTypeSet
	MachineTypes []PHostMachineTypeSet
}

// NewDescribePHostMachineTypeRequest will create request of DescribePHostMachineType action.
func (c *UPHostClient) NewDescribePHostMachineTypeRequest() *DescribePHostMachineTypeRequest {
	req := &DescribePHostMachineTypeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribePHostMachineType - 获取物理云机型的详细描述信息
func (c *UPHostClient) DescribePHostMachineType(req *DescribePHostMachineTypeRequest) (*DescribePHostMachineTypeResponse, error) {
	var err error
	var res DescribePHostMachineTypeResponse

	err = c.Client.InvokeActionWithPatcher("DescribePHostMachineType", req, &res, utils.FrequencePatcher)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

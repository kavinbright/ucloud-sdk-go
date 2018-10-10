//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api ULB GetBackendMetricInfo

package ulb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// GetBackendMetricInfoRequest is request schema for GetBackendMetricInfo action
type GetBackendMetricInfoRequest struct {
	request.CommonBase

	// 后端rs的BackendId数组
	BackendIds []string `required:"true"`

	// 负载均衡实例的ID
	ULBId *string `required:"false"`

	// VServer实例的ID
	VServerId *string `required:"false"`
}

// GetBackendMetricInfoResponse is response schema for GetBackendMetricInfo action
type GetBackendMetricInfoResponse struct {
	response.CommonBase

	// 符合条件的总数
	TotalCount int

	// Backend列表，每项参数详见 DataSetBackend
	DataSet []DataSetBackend
}

// NewGetBackendMetricInfoRequest will create request of GetBackendMetricInfo action.
func (c *ULBClient) NewGetBackendMetricInfoRequest() *GetBackendMetricInfoRequest {
	req := &GetBackendMetricInfoRequest{}
	c.client.SetupRequest(req)
	return req
}

// GetBackendMetricInfo - 获取用于查询后端节点监控的元数据
func (c *ULBClient) GetBackendMetricInfo(req *GetBackendMetricInfoRequest) (*GetBackendMetricInfoResponse, error) {
	var err error
	var res GetBackendMetricInfoResponse

	err = c.client.InvokeAction("GetBackendMetricInfo", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

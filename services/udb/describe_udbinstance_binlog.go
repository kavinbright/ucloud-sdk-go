//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB DescribeUDBInstanceBinlog

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUDBInstanceBinlogRequest is request schema for DescribeUDBInstanceBinlog action
type DescribeUDBInstanceBinlogRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// DB实例Id
	DBId *string `required:"true"`

	// 过滤条件:起始时间(时间戳)
	BeginTime *int `required:"true"`

	// 过滤条件:结束时间(时间戳)
	EndTime *int `required:"true"`
}

// DescribeUDBInstanceBinlogResponse is response schema for DescribeUDBInstanceBinlog action
type DescribeUDBInstanceBinlogResponse struct {
	response.CommonBase

	// 获取的Binlog信息列表 UDBInstanceBinlogSet
	DataSet []UDBInstanceBinlogSet
}

// NewDescribeUDBInstanceBinlogRequest will create request of DescribeUDBInstanceBinlog action.
func (c *UDBClient) NewDescribeUDBInstanceBinlogRequest() *DescribeUDBInstanceBinlogRequest {
	req := &DescribeUDBInstanceBinlogRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUDBInstanceBinlog - 获取UDB指定时间段的binlog列表
func (c *UDBClient) DescribeUDBInstanceBinlog(req *DescribeUDBInstanceBinlogRequest) (*DescribeUDBInstanceBinlogResponse, error) {
	var err error
	var res DescribeUDBInstanceBinlogResponse

	err = c.client.InvokeAction("DescribeUDBInstanceBinlog", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB DeleteUDBBackup

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteUDBBackupRequest is request schema for DeleteUDBBackup action
type DeleteUDBBackupRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// 备份id，可通过DescribeUDBBackup获得
	BackupId *int `required:"true"`

	// 跨可用区高可用备库所在可用区，参见［可用区列表］
	BackupZone *string `required:"false"`
}

// DeleteUDBBackupResponse is response schema for DeleteUDBBackup action
type DeleteUDBBackupResponse struct {
	response.CommonBase
}

// NewDeleteUDBBackupRequest will create request of DeleteUDBBackup action.
func (c *UDBClient) NewDeleteUDBBackupRequest() *DeleteUDBBackupRequest {
	req := &DeleteUDBBackupRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteUDBBackup - 删除UDB实例备份
func (c *UDBClient) DeleteUDBBackup(req *DeleteUDBBackupRequest) (*DeleteUDBBackupResponse, error) {
	var err error
	var res DeleteUDBBackupResponse

	err = c.client.InvokeAction("DeleteUDBBackup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

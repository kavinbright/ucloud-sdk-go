// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

/*
SubnetInfo - 子网信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn/api for detail.
*/
type SubnetInfo struct {

	// 创建时间，时间戳
	CreateTime int

	// 名称
	Name string

	// 网段
	Network string

	// 地域
	Region string

	// 描述
	Remark string

	// 状态；Allocating：申请中,Available：有效,Deleting：删除中,Deleted：已删除
	State string

	// ID
	SubnetID string

	// 更新时间，时间戳
	UpdateTime int

	// 可用区
	Zone string
}

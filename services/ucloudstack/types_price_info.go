// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

/*
PriceInfo - 价格信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn/api for detail.
*/
type PriceInfo struct {

	// 计费模式。枚举值：Dynamic，表示小时；Month，表示月；Year，表示年；
	ChargeType string

	//
	Price float64
}
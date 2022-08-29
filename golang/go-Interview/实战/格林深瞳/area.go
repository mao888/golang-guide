/**
    @author:Hasee
    @data:2022/7/7
    @note:
**/
package main

//	国
type Nation struct {
	IdAndName
	Province         []Province         //	省
	ZhiCity          []ZhiCity          // 直辖市
	AutonomousRegion []AutonomousRegion //	自治区
}

//	省
type Province struct {
	IdAndName
	ZiArea []Area // 子级行政区
}

// 直辖市
type ZhiCity struct {
	IdAndName
	ZiArea []Area // 子级行政区
}

//	自治区
type AutonomousRegion struct {
	IdAndName
	ZiArea []Area // 子级行政区
}

//	行政区
type Area struct {
	IdAndName
	ZiArea []Area // 子级行政区
}

type IdAndName struct {
	Id   int64  // id
	Name string // 名称
}

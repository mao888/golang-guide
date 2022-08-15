/**
    @author: huchao
    @data:	 2022/3/15
    @note:	 封装成一个包,供外面使用
**/
package statistics

import (
	"sync"
)

var GpsClient = new(dataMap)

var CanClient = new(dataMap)

type statistic interface {
	IncSuccess() uint64
	IncError() uint64
	GetSuccess() uint64
	GetError() uint64
	GetCount() uint64
}

type dataMap struct {
	dataType string
	success  uint64
	error    uint64
	sync.Mutex
}

func (d *dataMap) IncSuccess() uint64 {
	d.Lock()
	defer d.Unlock()
	d.success ++
	return d.success
}

func (d *dataMap) IncError() uint64 {
	d.Lock()
	defer d.Unlock()
	d.error ++
	return d.error
}

func (d *dataMap) GetSuccess() uint64 {
	d.Lock()
	defer d.Unlock()
	return d.success
}

func (d *dataMap) GetError() uint64 {
	d.Lock()
	defer d.Unlock()
	return d.error
}

func (d *dataMap) GetCount() uint64 {
	d.Lock()
	defer d.Unlock()
	count := d.success + d.error
	return count
}
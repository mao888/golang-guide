package main

import (
	"errors"
)

const MaxLength = 20

type LineList struct {
	MaxLength       int
	Length          int
	LineListContent []interface{}
}

// InitLineList 线性表的初始化操作
func InitLineList() *LineList {
	return &LineList{
		MaxLength:       MaxLength,
		Length:          0,
		LineListContent: make([]interface{}, 0),
	}
}

// Empty 判断线性表是否为空
func (l *LineList) Empty() bool {
	return l.Length == 0
}

// GetElem 获取线性表第i个元素的值，第一个元素对应线性表下表为0的元素
func (l *LineList) GetElem(i int) (interface{}, error) {
	if i < 1 || i > l.Length {
		return "", errors.New("查找的元素不在线性表范围内")
	}
	return l.LineListContent[i-1], nil
}

// DelElem 删除线性表的第i个元素
func (l *LineList) DelElem(i int) (bool, error) {
	if i < 1 || i > l.Length {
		return false, errors.New("查找的元素不在线性表范围内")
	}
	if l.Empty() {
		return false, errors.New("空表没有可以删除的数据")
	}
	for j := i - 1; j < l.Length-1; j++ {
		l.LineListContent[j] = l.LineListContent[j+1]
	}
	l.LineListContent = l.LineListContent[:l.Length-1]
	l.Length--
	return true, nil
}

// Insert 在线性表中的第i个位置插入元素data
func (l *LineList) Insert(i int, data interface{}) (bool, error) {
	if i < 1 || i > l.Length {
		return false, errors.New("查找的元素不在线性表范围内")
	}
	if b, _ := l.Append(""); !b {
		return false, errors.New("线性表已满，无法添加数据")
	}
	for j := l.Length - 1; j > i-1; j-- {
		l.LineListContent[j] = l.LineListContent[j-1]
	}
	l.LineListContent[i-1] = data
	return true, nil
}

// Pop 从末尾弹出一个元素
func (l *LineList) Pop() (interface{}, error) {
	if l.Empty() {
		return "", errors.New("线性表长度为0，没有可弹出的元素")
	}
	result := l.LineListContent[l.Length-1]
	l.LineListContent = l.LineListContent[:l.Length-1]
	l.Length--
	return result, nil
}

// Append 从末尾插入一个元素
func (l *LineList) Append(data interface{}) (bool, error) {
	if l.Length == l.MaxLength {
		return false, errors.New("线性表已满，无法添加数据")
	}
	l.LineListContent = append(l.LineListContent, data)
	l.Length++
	return true, nil
}

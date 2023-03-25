## 需求描述

点击一级或二级或三级类别都会进行查询，选中的类别变为紫色

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1679466785028-c8fc18a2-bf8f-4381-b430-6db3427ec77b.png)

## 效果图

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1679467334820-bc485ef1-099a-4f18-81e4-284eeb80a961.png)

## 数据库设计

### model

```go
// ArtAssetCategory 资产分类表 mapped from table <art_asset_category>
type ArtAssetCategory struct {
	ID        int32  `gorm:"column:id;primaryKey" json:"id"`
	ParentID  int32  `gorm:"column:parent_id;not null" json:"parent_id"`         // 上级分类
	Name      string `gorm:"column:name;not null" json:"name"`                   // 分类名称
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建日期
	Remark    string `gorm:"column:remark;not null" json:"remark"`               // 备注
	Order     int32  `gorm:"column:order;not null" json:"order"`                 // 位置
}
```

### 表

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1679467870216-a6721236-5396-432e-abc9-010553991e46.png)

## 实现思想

1. 将根类别（ParentID为0）的类别装入 list
2. 否则装入categoryMap（key：id，value：类别）
3. 遍历 list，根据根类别 ID 从 categoryMap 中查出当前类别下的子类别

1. 1. 若一个类别的 ParentID == categroyId（list.id），则该类别为 categroyId(list.id) 下的子类别
   2. 将全部子类别装入 list

1. 1. 1. 遍历 list，递归查询

1. 根据order(位置)从小到大 排序
2. 返回 包含子类别（Child）的类别树（CategoryTree）

## golang代码实现

### handler

```go
// CategoryTree 所有分类 树形下拉接口
func (ah *ArtAssetHandler) CategoryTree(c *gin.Context) {
	ctx := c.Request.Context()
	resp, err := ah.artAssetLogic.AssetCategoryTree(ctx)
	if err != nil {
		ah.Fail(c, err)
		return
	}
	ah.Success(c, resp)
}
```

### logic

```go
func (aal *ArtAssetLogic) AssetCategoryTree(ctx context.Context) (*bean.AssetCategoryTreeResp, error) {

	alllist, err := aal.artAssetService.GetAllCategoryList(ctx)
	if err != nil {
		return nil, err
	}
    // 将根类别（ParentID为0）的类别装入list； 
    // 否则装入categoryMap（key：id，value：类别）
	categoryMap := make(map[int32]*model.ArtAssetCategory)
	list := make([]*model.ArtAssetCategory, 0) // 根级别 有可能是多个 所以是数组
	for _, v := range alllist {
		if v.ParentID == 0 { // 第一级别的列表 parentid  ==0 默认0
			list = append(list, v)
		} else {
			categoryMap[v.ID] = v
		}
	}
    // 包含子类别（Child）的类别树（CategoryTree）
	childList := make([]*bean.AssetCategoryTreeChild, 0, len(list))
	for _, v := range list {
		child := &bean.AssetCategoryTreeChild{
			ID:       v.ID,
			Name:     v.Name,
			ParentId: v.ParentID,
			Order:    v.Order,
		}
        // 根据 ID 查出当前类别下的子类别
		child.Child = aal.aggreCategoryTree(categoryMap, v.ID)
		childList = append(childList, child)
	}

    // 根据order(位置)从小到大 排序
	sort.Slice(childList, func(i, j int) bool {
		return childList[i].Order < childList[j].Order
	})
	return &bean.AssetCategoryTreeResp{
		Tree: childList,
	}, nil
}

func (aal *ArtAssetLogic) aggreCategoryTree(cmap map[int32]*model.ArtAssetCategory, categroyId int32) []*bean.AssetCategoryTreeChild {
	if categroyId <= 0 || len(cmap) <= 0 {
		return nil
	}
    // 查询 categroyId 下的子类别
    // 若一个类别的 ParentID == categroyId，则该类别为 categroyId 下的子类别
    // 将全部子类别装入list
	list := make([]*model.ArtAssetCategory, 0)
	for _, v := range cmap {
		if v.ParentID == categroyId { // 第一级别的列表
			list = append(list, v)
		}
	}
	if len(list) <= 0 {
		return nil
	}
    // 包含子类别（Child）的类别树（CategoryTree）
	childList := make([]*bean.AssetCategoryTreeChild, 0, len(list))
	for _, v := range list {
		child := &bean.AssetCategoryTreeChild{
			ID:       v.ID,
			Name:     v.Name,
			ParentId: v.ParentID,
			Order:    v.Order,
		}
        // 根据 ID 查出当前类别下的子类别（递归）
		sublist := aal.aggreCategoryTree(cmap, v.ID)
		if sublist != nil && len(sublist) > 0 {
			child.Child = sublist
		}
		childList = append(childList, child)
	}
	// 根据order(位置)从小到大 排序
	sort.Slice(childList, func(i, j int) bool {
		return childList[i].Order < childList[j].Order
	})
	return childList
}
```

### service

```go
// 获取资产分类 所有，内存做树形结构组装 只适合数量少的情况，数量大的一定让前端组装
func (aal *ArtAssetService) GetAllCategoryList(ctx context.Context) ([]*model.ArtAssetCategory, error) {
	resList := make([]*model.ArtAssetCategory, 0)
	db := gmysql.DB(ctx, config.GlobConfig.Mysql.DBName) //连接到数据库
	categorys := query.Use(db).ArtAssetCategory
	err := categorys.WithContext(ctx).Order(categorys.Order).Scan(&resList)
	if err != nil {
		return nil, err
	}
	return resList, nil

}
```

### bean

```go
type AssetCategoryTreeResp struct {
	Tree []*AssetCategoryTreeChild `json:"list"`
}

type AssetCategoryTreeChild struct {
	ID       int32                     `json:"id"`
	Name     string                    `json:"name"`
	ParentId int32                     `json:"parent_id"`
	Child    []*AssetCategoryTreeChild `json:"child"`
	Order    int32                     `json:"order"`
}
```

## 请求结果

```json
{
    "code": 0,
    "message": "",
    "data": {
        "list": [
            {
                "id": 1,
                "name": "场景",
                "parent_id": 0,
                "child": [
                    {
                        "id": 9,
                        "name": "自然场景",
                        "parent_id": 1,
                        "child": null,
                        "order": 1
                    },
                    {
                        "id": 10,
                        "name": "建筑",
                        "parent_id": 1,
                        "child": null,
                        "order": 2
                    },
                    {
                        "id": 11,
                        "name": "地形",
                        "parent_id": 1,
                        "child": null,
                        "order": 3
                    },
                    {
                        "id": 12,
                        "name": "科技风",
                        "parent_id": 1,
                        "child": null,
                        "order": 4
                    },
                    {
                        "id": 13,
                        "name": "现代风",
                        "parent_id": 1,
                        "child": null,
                        "order": 5
                    },
                    {
                        "id": 14,
                        "name": "乡村风",
                        "parent_id": 1,
                        "child": null,
                        "order": 6
                    },
                    {
                        "id": 15,
                        "name": "场景部件",
                        "parent_id": 1,
                        "child": null,
                        "order": 7
                    },
                    {
                        "id": 16,
                        "name": "其他（场景）",
                        "parent_id": 1,
                        "child": null,
                        "order": 8
                    }
                ],
                "order": 0
            },
            {
                "id": 2,
                "name": "人物",
                "parent_id": 0,
                "child": [
                    {
                        "id": 17,
                        "name": "成年人-男",
                        "parent_id": 2,
                        "child": null,
                        "order": 1
                    },
                    {
                        "id": 18,
                        "name": "成年人-女",
                        "parent_id": 2,
                        "child": null,
                        "order": 2
                    },
                    {
                        "id": 19,
                        "name": "孩子-男",
                        "parent_id": 2,
                        "child": null,
                        "order": 3
                    },
                    {
                        "id": 20,
                        "name": "孩子-女",
                        "parent_id": 2,
                        "child": null,
                        "order": 4
                    },
                    {
                        "id": 21,
                        "name": "老人-男",
                        "parent_id": 2,
                        "child": null,
                        "order": 5
                    },
                    {
                        "id": 22,
                        "name": "老人-女",
                        "parent_id": 2,
                        "child": null,
                        "order": 6
                    },
                    {
                        "id": 23,
                        "name": "其他（人物）",
                        "parent_id": 2,
                        "child": null,
                        "order": 7
                    }
                ],
                "order": 1
            },
            {
                "id": 3,
                "name": "动物",
                "parent_id": 0,
                "child": [
                    {
                        "id": 24,
                        "name": "哺乳",
                        "parent_id": 3,
                        "child": null,
                        "order": 1
                    },
                    {
                        "id": 25,
                        "name": "飞禽",
                        "parent_id": 3,
                        "child": null,
                        "order": 2
                    },
                    {
                        "id": 26,
                        "name": "爬行",
                        "parent_id": 3,
                        "child": null,
                        "order": 3
                    },
                    {
                        "id": 27,
                        "name": "恐龙",
                        "parent_id": 3,
                        "child": null,
                        "order": 4
                    },
                    {
                        "id": 28,
                        "name": "昆虫",
                        "parent_id": 3,
                        "child": null,
                        "order": 5
                    },
                    {
                        "id": 29,
                        "name": "鱼类",
                        "parent_id": 3,
                        "child": null,
                        "order": 6
                    },
                    {
                        "id": 30,
                        "name": "两栖",
                        "parent_id": 3,
                        "child": null,
                        "order": 7
                    },
                    {
                        "id": 31,
                        "name": "机器动物",
                        "parent_id": 3,
                        "child": null,
                        "order": 8
                    },
                    {
                        "id": 32,
                        "name": "其他（动物）",
                        "parent_id": 3,
                        "child": null,
                        "order": 9
                    }
                ],
                "order": 2
            },
            {
                "id": 4,
                "name": "植物",
                "parent_id": 0,
                "child": [
                    {
                        "id": 33,
                        "name": "石头",
                        "parent_id": 4,
                        "child": null,
                        "order": 1
                    },
                    {
                        "id": 34,
                        "name": "花卉",
                        "parent_id": 4,
                        "child": null,
                        "order": 2
                    },
                    {
                        "id": 35,
                        "name": "树木",
                        "parent_id": 4,
                        "child": null,
                        "order": 3
                    },
                    {
                        "id": 36,
                        "name": "草类",
                        "parent_id": 4,
                        "child": null,
                        "order": 4
                    },
                    {
                        "id": 37,
                        "name": "水果蔬菜",
                        "parent_id": 4,
                        "child": null,
                        "order": 5
                    },
                    {
                        "id": 38,
                        "name": "其他（植物）",
                        "parent_id": 4,
                        "child": null,
                        "order": 6
                    }
                ],
                "order": 3
            },
            {
                "id": 5,
                "name": "道具",
                "parent_id": 0,
                "child": [
                    {
                        "id": 39,
                        "name": "武器",
                        "parent_id": 5,
                        "child": null,
                        "order": 1
                    },
                    {
                        "id": 40,
                        "name": "家具/生活用具",
                        "parent_id": 5,
                        "child": [
                            {
                                "id": 49,
                                "name": "沙发",
                                "parent_id": 40,
                                "child": null,
                                "order": 1
                            },
                            {
                                "id": 50,
                                "name": "桌椅",
                                "parent_id": 40,
                                "child": null,
                                "order": 2
                            },
                            {
                                "id": 51,
                                "name": "床",
                                "parent_id": 40,
                                "child": null,
                                "order": 3
                            },
                            {
                                "id": 52,
                                "name": "柜子",
                                "parent_id": 40,
                                "child": null,
                                "order": 4
                            },
                            {
                                "id": 53,
                                "name": "门窗",
                                "parent_id": 40,
                                "child": null,
                                "order": 5
                            },
                            {
                                "id": 54,
                                "name": "灯具",
                                "parent_id": 40,
                                "child": null,
                                "order": 6
                            }
                        ],
                        "order": 2
                    },
                    {
                        "id": 41,
                        "name": "食品/饮料/药品",
                        "parent_id": 5,
                        "child": [
                            {
                                "id": 55,
                                "name": "水果",
                                "parent_id": 41,
                                "child": null,
                                "order": 1
                            },
                            {
                                "id": 56,
                                "name": "蔬菜",
                                "parent_id": 41,
                                "child": null,
                                "order": 2
                            },
                            {
                                "id": 57,
                                "name": "零食",
                                "parent_id": 41,
                                "child": null,
                                "order": 3
                            },
                            {
                                "id": 58,
                                "name": "饮料",
                                "parent_id": 41,
                                "child": null,
                                "order": 4
                            }
                        ],
                        "order": 3
                    },
                    {
                        "id": 42,
                        "name": "其他（道具）",
                        "parent_id": 5,
                        "child": null,
                        "order": 4
                    }
                ],
                "order": 4
            },
            {
                "id": 6,
                "name": "载具",
                "parent_id": 0,
                "child": [
                    {
                        "id": 43,
                        "name": "车辆",
                        "parent_id": 6,
                        "child": null,
                        "order": 1
                    },
                    {
                        "id": 44,
                        "name": "船艇",
                        "parent_id": 6,
                        "child": null,
                        "order": 2
                    },
                    {
                        "id": 45,
                        "name": "飞机/航空器",
                        "parent_id": 6,
                        "child": null,
                        "order": 3
                    },
                    {
                        "id": 46,
                        "name": "其他（载具）",
                        "parent_id": 6,
                        "child": null,
                        "order": 4
                    }
                ],
                "order": 5
            },
            {
                "id": 7,
                "name": "怪物",
                "parent_id": 0,
                "child": [
                    {
                        "id": 47,
                        "name": "其他（怪物）",
                        "parent_id": 7,
                        "child": null,
                        "order": 0
                    }
                ],
                "order": 6
            },
            {
                "id": 8,
                "name": "其他",
                "parent_id": 0,
                "child": [
                    {
                        "id": 48,
                        "name": "其他（其他）",
                        "parent_id": 8,
                        "child": null,
                        "order": 0
                    }
                ],
                "order": 7
            }
        ]
    }
}
```
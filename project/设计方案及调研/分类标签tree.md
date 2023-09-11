## 需求:实现如下图所示标签树

[![911.png](https://i.postimg.cc/VkrwLQ5p/911.png)](https://postimg.cc/Bj0kMzTB)

## 数据库

| id   | name  | game_id     | first_classify | second_classify |
| ---- | ----- | ----------- | -------------- | --------------- |
| 1    | 标签1 | test123mumu | 1              | 2               |
| 2    | 标签2 | test123mumu | 1              | 2               |
| 3    | 标签3 | test123mumu | 1              | 2               |
| 4    | 标签4 | test123mumu | 1              | 3               |
| 5    | 标签5 | test123mumu | 1              | 4               |
| 6    | 标签6 | test123mumu | 6              | 6               |
| 7    | 标签7 | test123mumu | 7              | 7               |
| 8    | 标签8 | test123mumu | 4              | 4               |
| 9    | 场景  | test123mumu |                |                 |
| 10   | 男    | test123mumu | 性别           |                 |
| 11   | 女    | test123mumu | 性别           |                 |
| 12   | 中    | test123mumu | 性质           |                 |

## 思路

主要思路是通过迭代 **results**(数据库中查出的全部数据) 中的每个标签，将它们组织成树形结构。首先，检查标签的分类层级（第一级和第二级），并根据需要创建新的树节点。然后，将当前标签作为子节点添加到相应的位置。最终，返回构建的树节点列表。

## 代码实现

```go
func tree(results []*model.ArtMaterialTag) []*bean.Tree {
	// 创建一个空的树节点列表
	list := make([]*bean.Tree, 0)

	// 创建两个映射，用于跟踪第一级和第二级分类的索引位置
	firstIndexMapping, secondIndexMapping := make(map[string]int), make(map[string]int)

	// 遍历给定的结果列表
	for _, tag := range results {
		// 初始化一些变量
		firstIndex, secondIndex, haveFirst, haveSecond := 0, 0, false, false

		// 创建一个树节点，并填充其属性
		node := bean.Tree{
			ID:       tag.ID,
			Label:    tag.Name,
			Children: make([]bean.Tree, 0),
		}

		// 处理第一级分类
		if tag.FirstClassify != constants.EmptyString {
			haveFirst = true

			// 检查第一级分类是否已经存在于列表中，如果不存在，则创建一个新的节点
			v, ok := firstIndexMapping[tag.FirstClassify]
			if !ok {
				list = append(list, &bean.Tree{Label: tag.FirstClassify, Children: make([]bean.Tree, 0)})
				firstIndexMapping[tag.FirstClassify] = len(list) - 1
				firstIndex = len(list) - 1
			} else {
				firstIndex = v
			}
		}

		// 处理第二级分类
		if tag.SecondClassify != constants.EmptyString {
			haveSecond = true

			// 创建一个用于唯一标识第二级分类的键
			key := tag.FirstClassify + "-" + tag.SecondClassify

			// 检查第二级分类是否已经存在于列表中，如果不存在，则创建一个新的节点
			v, ok := secondIndexMapping[key]
			if !ok {
				list[firstIndex].Children = append(list[firstIndex].Children, bean.Tree{Label: tag.SecondClassify, Children: make([]bean.Tree, 0)})
				secondIndexMapping[key] = len(list[firstIndex].Children) - 1
				secondIndex = len(list[firstIndex].Children) - 1
			} else {
				secondIndex = v
			}
		}

		// 将当前节点添加到树的相应位置
		if haveSecond {
			list[firstIndex].Children[secondIndex].Children = append(list[firstIndex].Children[secondIndex].Children, node)
		} else if haveFirst {
			list[firstIndex].Children = append(list[firstIndex].Children, node)
		} else {
			list = append(list, &node)
			firstIndexMapping[tag.FirstClassify] = len(list) - 1
		}
	}

	// 返回构建的树节点列表
	return list
}
```

## 代码核心解读

**firstIndexMapping** 和 **secondIndexMapping** 是两个映射，它们用于跟踪已经存在的第一级和第二级分类在树结构中的索引位置。这两个映射的目的是为了避免重复创建相同的分类节点，以及在需要时能够快速找到已经存在的节点位置。以下是它们在代码中的作用和思路：

1. **firstIndexMapping** 映射：

- - 这个映射的键是第一级分类的名称（**tag.FirstClassify**）。
  - 对应的值是该分类在 **list** 树节点列表中的索引位置。
  - 如果某个第一级分类在 **firstIndexMapping** 中不存在，那么代码会创建一个新的树节点表示该分类，并将它添加到 **list** 中，然后将新节点的索引位置存储在 **firstIndexMapping** 中。如果已经存在，代码会直接获取该分类在 **list** 中的索引位置。

1. **secondIndexMapping** 映射：

- - 这个映射的键是一个组合键，由第一级分类和第二级分类组合而成（**tag.FirstClassify + "-" + tag.SecondClassify**）。
  - 对应的值是第二级分类在其父节点（第一级分类节点）的子节点列表中的索引位置。
  - 如果某个第二级分类在 **secondIndexMapping** 中不存在，那么代码会创建一个新的树节点表示该分类，并将它添加到其父节点的子节点列表中，然后将新节点的索引位置存储在 **secondIndexMapping** 中。如果已经存在，代码会直接获取该分类在其父节点的子节点列表中的索引位置。

这两个映射的使用可以有效地构建树结构，确保相同分类不会重复创建，而是被添加到已有的分类节点下。这有助于提高代码的效率和减少内存消耗，同时确保最终的树形结构正确反映了分类的层次关系。

## 返回json结果示例

```json
{
  "code": 0,
  "message": "",
  "data": [
    {
      "id": 0,
      "label": "1",
      "children": [
        {
          "id": 0,
          "label": "2",
          "children": [
            {
              "id": 1,
              "label": "标签1",
              "children": []
            },
            {
              "id": 2,
              "label": "标签2",
              "children": []
            },
            {
              "id": 3,
              "label": "标签3",
              "children": []
            }
          ]
        },
        {
          "id": 0,
          "label": "3",
          "children": [
            {
              "id": 4,
              "label": "标签4",
              "children": []
            }
          ]
        },
        {
          "id": 0,
          "label": "4",
          "children": [
            {
              "id": 5,
              "label": "标签5",
              "children": []
            }
          ]
        }
      ]
    },
    {
      "id": 0,
      "label": "6",
      "children": [
        {
          "id": 0,
          "label": "6",
          "children": [
            {
              "id": 6,
              "label": "标签6",
              "children": []
            }
          ]
        }
      ]
    },
    {
      "id": 0,
      "label": "7",
      "children": [
        {
          "id": 0,
          "label": "7",
          "children": [
            {
              "id": 7,
              "label": "标签7",
              "children": []
            },
            {
              "id": 15,
              "label": "测试第二二条",
              "children": []
            }
          ]
        }
      ]
    },
    {
      "id": 0,
      "label": "4",
      "children": [
        {
          "id": 0,
          "label": "4",
          "children": [
            {
              "id": 8,
              "label": "标签8",
              "children": []
            },
            {
              "id": 13,
              "label": "标签9",
              "children": []
            }
          ]
        }
      ]
    },
    {
      "id": 9,
      "label": "场景",
      "children": []
    },
    {
      "id": 0,
      "label": "性别",
      "children": [
        {
          "id": 10,
          "label": "男",
          "children": []
        },
        {
          "id": 11,
          "label": "女",
          "children": []
        }
      ]
    },
    {
      "id": 0,
      "label": "性质",
      "children": [
        {
          "id": 12,
          "label": "中",
          "children": []
        }
      ]
    },
    {
      "id": 0,
      "label": "3",
      "children": [
        {
          "id": 0,
          "label": "4",
          "children": [
            {
              "id": 14,
              "label": "测试第一条",
              "children": []
            },
            {
              "id": 16,
              "label": "测试第一个饿",
              "children": []
            }
          ]
        }
      ]
    }
  ]
}
```
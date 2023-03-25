package user_console

import (
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/user_console/bean"
	"strconv"
	"strings"
	"time"
)

func RunUserPermPoliceResourceBiGame() {
	// 处理字符串, 以 "\n" 分隔为 ark_id:[app_id,app_id]
	arkIDAndAppIdStrsArr := strings.Split(bean.ArkIDAndAppIdStrs, "\n")
	//arkIDAndAppIdMap := make([]map[string]string, len(arkIDAndAppIdStrsArr))
	for _, s := range arkIDAndAppIdStrsArr {
		arkIDAndAppID := strings.Split(s, ":")
		arkIDAndAppID[1] = strings.TrimPrefix(arkIDAndAppID[1], "[")
		arkIDAndAppID[1] = strings.TrimSuffix(arkIDAndAppID[1], "]")

		// ark ID	"10110"
		arkID := arkIDAndAppID[0]
		arkIDInt, err := strconv.Atoi(arkID)
		if err != nil {
			fmt.Println(arkID, "arkID 转int 错误：", err)
			return
		}

		// ark ID 对应的 app_id 数组	"[100186 100187]"
		appIDStr := arkIDAndAppID[1]
		appIDStrArr := strings.Split(appIDStr, ",")

		// 根据arkID 去user_perm 员工权限关联表 查询是否有此用户
		userPerm := make([]*bean.UserPerm, 0)
		err = db2.MySQLClientUser.Table("user_perm").Where("user_id = ?", arkID).
			Where("perm_id = ?", bean.PermID).
			Find(&userPerm).Error
		if err != nil {
			fmt.Println("mysql/user_perm 查询错误：", err)
			return
		}
		// 查到，跳过
		//if len(userPerm) != 0 {
		//	fmt.Printf("arkID:%d, 已在user_perm表中\n", arkIDInt)
		//	continue
		//}

		//// user_perm 员工权限关联表 未查到，则直接全部插入
		//// 1、policy 策略表 自增一条拿到 policy_id
		//police := &bean.Policy{
		//	//ID:        0,
		//	UpdatedAt: int32(time.Now().Unix()),
		//	CreatedAt: int32(time.Now().Unix()),
		//}
		//err = db2.MySQLClientUser.Table("policy").Create(police).Error
		//if err != nil {
		//	fmt.Println("入mysql/policy 错误：", err)
		//	return
		//}
		//policyID := police.ID

		// 3、user_perm 员工权限关联表
		//userPerm2 := &bean.UserPerm{
		//	//ID:        0,
		//	UserID:    int32(arkIDInt),
		//	PermID:    bean.PermID,
		//	PolicyID:  0,
		//	ScopeID:   policyID,
		//	UpdatedAt: int32(time.Now().Unix()),
		//	CreatedAt: int32(time.Now().Unix()),
		//}
		//err = db2.MySQLClientUser.Table("user_perm").Create(userPerm2).Error
		//if err != nil {
		//	fmt.Println("arkID:", arkIDInt, "入mysql/user_perm 错误：", err)
		//	return
		//}
		//fmt.Printf("arkID:%d, 成功插入\n", arkIDInt)

		// 2、policy_resource 策略资源关联表
		// 根据 app_id 查询 game_id
		var gameIDStrArr []string
		game2 := make([]*bean.App, 0)
		err = db2.MySQLClient.Table("app").Where("id in ?", appIDStrArr).Distinct("game_id").Find(&game2).Error
		if err != nil {
			fmt.Println("根据app_id查询game 错误：", err)
		}
		// 将 根据 app_id 查询 game_id 与原 game_id 数组合并
		for _, app := range game2 {
			gameIDStrArr = append(gameIDStrArr, app.GameID)
		}
		// 将合并后的app_id数组去重
		//fmt.Println("gameIDStrArr:", gameIDStrArr)
		//gameIDStrArrDistinct := DistinctString(gameIDStrArr)
		for _, s2 := range gameIDStrArr {
			policyResource := &bean.PolicyResource{
				//ID:         0,
				PolicyID:   userPerm[0].ScopeID,
				ResourceID: bean.ResourceIdGame,
				EntityID:   s2,
				UpdatedAt:  int32(time.Now().Unix()),
				CreatedAt:  int32(time.Now().Unix()),
			}
			err = db2.MySQLClientUser.Table("policy_resource").Create(policyResource).Error
			if err != nil {
				fmt.Println("入mysql/policyResource 错误：", err)
				return
			}
			fmt.Printf("arkID:%d, game_id:%s 成功插入policy_resource\n", arkIDInt, s2)
		}
	}
}

package user_console

import (
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/user_console/bean"
	"strconv"
	"strings"
	"time"
)

// RunUserPermPoliceResourceAdAppAndGame 广告投放应用数据权限
func RunUserPermPoliceResourceAdAppAndGame() {
	// 处理字符串, 以 "\n" 分隔为 ark_id:[app_id,app_id]
	arkIDAndAppIdAndGameIdStrArr := strings.Split(bean.AdArkIDAndAppIdAndGameIdStrs, "\n")
	//arkIDAndAppIdMap := make([]map[string]string, len(arkIDAndAppIdStrsArr))
	for _, s := range arkIDAndAppIdAndGameIdStrArr {
		arkIDAndAppIDAndGameId := strings.Split(s, ":")
		arkIDAndAppIDAndGameId[1] = strings.TrimPrefix(arkIDAndAppIDAndGameId[1], "[")
		arkIDAndAppIDAndGameId[1] = strings.TrimSuffix(arkIDAndAppIDAndGameId[1], "]")
		arkIDAndAppIDAndGameId[2] = strings.TrimPrefix(arkIDAndAppIDAndGameId[2], "[")
		arkIDAndAppIDAndGameId[2] = strings.TrimSuffix(arkIDAndAppIDAndGameId[2], "]")

		// ark ID	"10110"
		arkID := arkIDAndAppIDAndGameId[0]
		arkIDInt, err := strconv.Atoi(arkID)
		if err != nil {
			fmt.Println(arkID, "arkID 转int 错误：", err)
			return
		}
		//fmt.Println("arkID:", arkID)

		// ark ID 对应的 app_id 数组	"[100186 100187]"
		appIDStr := arkIDAndAppIDAndGameId[1]
		appIDStrArr := strings.Split(appIDStr, ",")
		//fmt.Println("appIDStrArr:", appIDStrArr)

		// ark ID 对应的 game_id 数组	"[100186 100187]"
		gameIDStr := arkIDAndAppIDAndGameId[2]
		gameIDStrArr := strings.Split(gameIDStr, ",")
		//fmt.Println("gameIDStrArr:", gameIDStrArr)

		// 0、根据arkID 去user_perm 员工权限关联表 查询 广告投放系统默认权限 是否有此用户，若有则删除
		userPerm := make([]*bean.UserPerm, 0)
		err = db2.MySQLClientUser.Table("user_perm").Where("user_id = ?", arkID).
			Where("perm_id = ?", bean.PermIDAd).
			Find(&userPerm).Error
		if err != nil {
			fmt.Println("mysql/user_perm 查询错误：", err)
			return
		}
		//fmt.Println("userPerm:", userPerm)

		// 查到,删除
		if len(userPerm) > 0 {
			// 根据user_id删除 user_perm
			err := db2.MySQLClientUser.Table("user_perm").Where("user_id = ?", arkID).
				Where("perm_id = ?", bean.PermIDAd).Delete(&bean.UserPerm{}).Error
			if err != nil {
				fmt.Println("user_perm 删除失败，user_id：", arkID)
				return
			}
			// 根据 scope_id 删除 policy_resource表
			err = db2.MySQLClientUser.Table("policy_resource").Where("policy_id = ?", userPerm[0].ScopeID).Delete(&bean.PolicyResource{}).Error
			if err != nil {
				fmt.Println("policy_resource 删除失败，policy_id：", userPerm[0].ScopeID)
				return
			}
			// 根据 scope_id 删除 policy表
			err = db2.MySQLClientUser.Table("policy").Where("id = ?", userPerm[0].ScopeID).Delete(&bean.Policy{}).Error
			if err != nil {
				fmt.Println("policy 删除失败，id：", userPerm[0].ScopeID)
				return
			}
		}

		// 1、policy 策略表 自增一条拿到 policy_id
		police := &bean.Policy{
			//ID:        0,
			UpdatedAt: int32(time.Now().Unix()),
			CreatedAt: int32(time.Now().Unix()),
		}
		err = db2.MySQLClientUser.Table("policy").Create(police).Error
		if err != nil {
			fmt.Println("入mysql/policy 错误：", err)
			return
		}
		policyID := police.ID

		// 3、user_perm 员工权限关联表
		userPerm2 := &bean.UserPerm{
			//ID:        0,
			UserID:    int32(arkIDInt),
			PermID:    bean.PermIDAd, // 10056 广告投放系统默认权限
			PolicyID:  0,
			ScopeID:   policyID,
			UpdatedAt: int32(time.Now().Unix()),
			CreatedAt: int32(time.Now().Unix()),
		}
		err = db2.MySQLClientUser.Table("user_perm").Create(userPerm2).Error
		if err != nil {
			fmt.Println("arkID:", arkIDInt, "入mysql/user_perm 错误：", err)
			return
		}
		fmt.Printf("arkID:%d, 成功插入user_perm\n", arkIDInt)

		// 2、policy_resource 策略资源关联表 —— App
		// 根据 game_id 查询 app_id
		game := make([]*bean.App, 0)
		err = db2.MySQLClient.Table("app").Where("game_id in ?", gameIDStrArr).Distinct("id").Find(&game).Error
		if err != nil {
			fmt.Println("根据app_id查询game 错误：", err)
		}
		// 将 根据 game_id 查询 app_id 与原 app_id 数组合并
		for _, app := range game {
			appIDStrArr = append(appIDStrArr, fmt.Sprintf("%d", app.ID))
		}
		// 将合并后的app_id数组去重
		//fmt.Println("appIDStrArr:", appIDStrArr)
		appIDStrArrDistinct := DistinctString(appIDStrArr)
		// 入库 policy_resource
		for _, s2 := range appIDStrArrDistinct {
			if s2 != "" {
				policyResource := &bean.PolicyResource{
					//ID:         0,
					PolicyID:   policyID,
					ResourceID: bean.ResourceIdApp,
					EntityID:   s2,
					UpdatedAt:  int32(time.Now().Unix()),
					CreatedAt:  int32(time.Now().Unix()),
				}
				err = db2.MySQLClientUser.Table("policy_resource").Create(policyResource).Error
				if err != nil {
					fmt.Println("入mysql/policyResource 错误：", err)
					return
				}
			}
			fmt.Printf("arkID:%d, app_id:%s 成功插入policy_resource\n", arkIDInt, s2)
		}

		// 2、policy_resource 策略资源关联表 —— Game
		// 根据 app_id 查询 game_id
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
		gameIDStrArrDistinct := DistinctString(gameIDStrArr)

		for _, s2 := range gameIDStrArrDistinct {
			if s2 != "" {
				policyResource := &bean.PolicyResource{
					//ID:         0,
					PolicyID:   policyID,
					ResourceID: bean.ResourceIdGame,
					EntityID:   s2,
					UpdatedAt:  int32(time.Now().Unix()),
					CreatedAt:  int32(time.Now().Unix()),
				}
				// 查询当前policyID、ResourceID下的EntityID是否存在，不存在则执行下一步插入
				err = db2.MySQLClientUser.Table("policy_resource").Create(policyResource).Error
				if err != nil {
					fmt.Println("入mysql/policyResource 错误：", err)
					return
				}
			}
			fmt.Printf("arkID:%d, game_id:%s 成功插入policy_resource\n", arkIDInt, s2)
		}

	}
}

func DistinctInt(i []int32) []int32 {
	idsMap := map[int32]struct{}{}
	res := make([]int32, 0, len(i))
	for _, item := range i {
		if _, ok := idsMap[item]; !ok {
			res = append(res, item)
			idsMap[item] = struct{}{}
		}
	}
	return res
}

func DistinctString(i []string) []string {
	idsMap := map[string]struct{}{}
	res := make([]string, 0, len(i))
	for _, item := range i {
		if _, ok := idsMap[item]; !ok {
			res = append(res, item)
			idsMap[item] = struct{}{}
		}
	}
	return res
}

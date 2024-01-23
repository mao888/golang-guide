package user_console

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/user_console/bean"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func RunUserPermPoliceResourceBi() {

	// 1、建立连接
	db := db2.MongoClient.Database("rambler")
	dbu := db2.MongoClient.Database("plat_console")
	collUsers := dbu.Collection("platusers")
	collDimpermissions := db.Collection("dimpermissions")

	// 2、从mongo查询数据
	mDimPermission := make([]*bean.MDimPermission, 0)
	err := collDimpermissions.Find(context.TODO(), bson.M{"child_system": bean.Bi}).All(&mDimPermission)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println("len(mDimPermission):", len(mDimPermission))

	// 3、将mongo数据装入切片
	for i, permission := range mDimPermission {
		fmt.Println("dimpermissions:", i)

		// 若app_ids为空，则跳过
		if len(permission.AppIds) == 0 {
			continue
		}
		for _, id := range permission.AppIds {

			// 1、policy 策略表 自增一条拿到 policy_id
			police := &bean.Policy{
				//ID:        0,
				UpdatedAt: int32(permission.UpdateTime.Unix()),
				CreatedAt: int32(permission.CreateTime.Unix()),
			}
			err = db2.MySQLClientUser.Table("policy").Create(police).Error
			if err != nil {
				fmt.Println("入mysql/policy 错误：", err)
				return
			}
			policyID := police.ID

			// 2、policy_resource 策略资源关联表
			policyResource := &bean.PolicyResource{
				//ID:         0,
				PolicyID:   policyID,
				ResourceID: bean.ResourceIdApp,
				EntityID:   string(id),
				UpdatedAt:  int32(permission.UpdateTime.Unix()),
				CreatedAt:  int32(permission.CreateTime.Unix()),
			}
			err = db2.MySQLClientUser.Table("policy_resource").Create(policyResource).Error
			if err != nil {
				fmt.Println("入mysql/policyResource 错误：", err)
				return
			}

			// 3、user_perm 员工权限关联表
			// UserID
			var authorID int32

			mPlatUser := make([]*bean.MPlatUser, 0)
			if permission.UserId != constants.NumberZero {
				// 根据 source.Author 去mongo查询用户信息
				err := collUsers.Find(context.TODO(), bson.M{"_id": permission.UserId}).All(&mPlatUser)
				if err != nil {
					fmt.Println("Mongo/platusers查询错误：", err)
					return
				}
			}

			if len(mPlatUser) != constants.NumberZero {
				// 根据用户邮箱和昵称查询mysql/user，拿到user_id
				user := make([]*bean.User, 0)

				err = db2.MySQLClientUser.Table("user").
					//Where("name = ?", mPlatUser[0].Name).
					Where("email = ?", mPlatUser[0].Email).
					Find(&user).Error
				if err != nil {
					fmt.Println("mysql/user 查询错误：", err)
				}
				if len(user) == constants.NumberZero {
					//authorID = 1000
					fmt.Printf("gia user_id:%d, name:%s, email:%s\n",
						mPlatUser[0].ID, mPlatUser[0].Name, mPlatUser[0].Email)
					continue
				} else {
					authorID = user[0].ID
				}
			} else {
				fmt.Printf("gia user_id:%d, name:%s, email:%s\n",
					permission.UserId, mPlatUser[0].Name, mPlatUser[0].Email)
				continue
			}

			userPerm := &bean.UserPerm{
				//ID:        0,
				UserID:    authorID,
				PermID:    bean.PermID,
				PolicyID:  0,
				ScopeID:   policyID,
				UpdatedAt: int32(permission.UpdateTime.Unix()),
				CreatedAt: int32(permission.CreateTime.Unix()),
			}
			err = db2.MySQLClientUser.Table("user_perm").Create(userPerm).Error
			if err != nil {
				fmt.Println("入mysql/user_perm 错误：", err)
				return
			}

		}

	}
}

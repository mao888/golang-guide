## 基于RBAC的用户权限管理系统

### 目标

搭建一个基于 RBAC ，支持功能权限 + 数据权限的统一用户权限管理系统。

### RBAC 简介

[看这里](https://en.wikipedia.org/wiki/Role-based_access_control)

### 周边系统设计

整个系统大致分为三个模块：用户权限模块、资源策略模块、用户权限-数据资源管理平台。外围交互的系统包括 钉钉 webhook、资源同步系统。

[![image.png](https://i.postimg.cc/7hshFLLL/image.png)](https://postimg.cc/fSdMXDm4)

下面将会针对这几部分分别做介绍。

#### 钉钉 API

通过钉钉 API ，同步钉钉的人员信息和组织架构。主要分为两部分： [API 主动调用](https://open.dingtalk.com/document/orgapp-server/queries-organization-data) 和 [webhook 事件监听](https://open.dingtalk.com/document/orgapp-server/configure-event-subcription) 。 可以采用 **主动定时更新人员架构或监听回调实时更新**。

#### 同步数据资源（方案待定）

终极目标是将所有需要设置为资源的数据，都同步到权限系统的"数据资源"中以供操作人员配置，但这就涉及到一个问题， 即什么是可配置的。 如果以有实际意义的实体的数据表作为一个资源，那么就应该暴露出对应的数据表的字段，且暴露哪些字段也应该是可配置的。 如果系统的使用人员是非技术人员，那么这个操作将变得非常不稳定。所以我们将这个功能分为两部分：

- 将数据表（理论上是任意数据库中的表）的字段添加到权限系统的资源列表中，本操作需要技术人员协助进行。
- 权限操作人员对根据资源，设置不同的数据权限，本操作不需要技术人员协助。

#### 用户权限-数据资源管理平台

略

#### 权限系统

整个权限系统分为两部分，即上文提到的功能权限和数据权限。而不论是哪种权限，我们都设想能够按照 用户-角色-权限 的三层结构来实现。

##### 关于权限的一些思考

到此为止，文中讲到的所有内容基本是按照 RBAC0 规范为主要思想的、通用的基于角色的权限管理。

现在我们设想一种业务场景：我们有一个系统 A，A 系统里一个 " APP 列表" 的模块显示了我能看到并管理的所有 APP， 点击列表中的某个 APP 后，会进入一个 APP 相关的功能页，我们称其为子系统 B， B 里面同样有很多功能模块， 我们称其为 功能1、2、3。 现在我们来分析下这其中涉及到的权限管理。

首先系统 A 中，应该有一个 操作（ Operation ）类型的权限，来控制我是否可以看到增删改查 " APP 列表" 的对应 操作按钮（主要偏前端）；同时会有一个 接口（ URL ）类型的权限，来控制我是否可以访问对应操作的接口（主要偏后端）； 最后，会有一个 数据（ Data ）类型的权限，来控制我可以看到列表中的哪些 APP。这样，系统 A 中的 " APP 列表" 模块 的权限就被我们完全的控制住了，任何人想要完成上述操作，都必须拥有对应的权限。

上面的情况，看起来是符合我们最初的设想，通过 员工、部门、（三种类型的）权限、角色的结合，完美地实现了 " APP 列表" 这一个功能模块的权限控制。但事情放到系统 B 中时，情况好像变得复杂了。

假设我是 M APP 的管理员，所以我们看到并使用 B 系统中 M APP 下的所有功能。我也是 N APP 的开发者，所以我只能 看到而并不能操作 B 系统中 N APP 下的功能。这时如果还按照在系统 A 中的方式，创建一个 "功能 1" 的操作权限， 似乎就不能满足需求了，因为我在不同的 APP 下，对应的操作权限也不一样。通常，如果这种情况发生在不同的系统中， 我们只需要配置不同系统中的不同权限然后赋给我就可以，但是不同 APP 并不是"系统"，而已一种"资源"，这样就使得 我们不能根据区分系统来创建不同的操作权限（资源是可变且繁多的）。

上面的情况总结起来就是，数据不是功能下面的附属品，而是和功能互相影响的一个元素。在一个系统中，不同功能下面 有对应的不同数据，同理不同的数据下面，也有不同的功能。但这也不是必须的，只有在 "功能在不同数据下表现得不一致"的 情况下才会出现，所以这不应该是一个强制的选项。这有点像 [RBAC with Domains](https://casbin.org/docs/en/rbac-with-domains), 只不过在上述情况下，域 变成了 某种资源，而 角色 变成了 权限。

综上所述，我们设计了一种方案来尝试解决这个问题。

- 权限和角色，在跟 部门、人员进行关联时:
  - 添加一个前置限制，即"我"在某个限制下 拥有某个角色或权限，这个限制可以理解为 RBAC with Domains 中的 Domain 。
  - 添加一个后置限制，即”我“拥有这个权限下的资源范围。
- 作用域可以看做某个系统或者某类数据。当为某个系统时，变现的与 RBAC with Domains 基本一致。当为某类数据时， 则表示某些功能需要限制在特定的数据资源中。





### 管理端技术方案

#### 数据库设计

整个后端设计大致分为四个部分：组织架构、角色&权限、系统&资源、策略管理。

[![image.png](https://i.postimg.cc/Y98CFSYx/image.png)](https://postimg.cc/zbbNZ8dL)

- 组织架构：负责同步钉钉的组织架构和人员至权限系统，并管理用户的登录、角色&权限配置等。
- 角色&权限：角色、权限的增删改查
- 系统&资源：负责管理各个子系统和子系统中的资源
- 策略管理：策略在系统中的含义是：某一种资源的范围。通过在 "组织架构-角色&权限" 和 "角色&权限-系统&资源" 的 关联中绑定对应的策略，即可实现对应权限的适用范围和某个（数据）权限的数据范围

[![image.png](https://i.postimg.cc/t4PR5kFJ/image.png)](https://postimg.cc/nXFtVvRf)

#### 员工&部门

通过钉钉 webhook 或定时拉取组织架构到我们系统中。

部门表：

```sql
create table department (
 `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
 `ding_id` varchar(50) NOT NULL COMMENT '钉钉id',
 `name` varchar(50) NOT NULL COMMENT '部门名称',
 `parent_id` int NOT NULL DEFAULT '0' COMMENT '父部门',
 `updated_at` int(12) NOT NULL DEFAULT '0' COMMENT '更新时间',
 `created_at` int(12) NOT NULL DEFAULT '0' COMMENT '创建时间',
 `is_deleted` int(1) NOT NULL DEFAULT '0' COMMENT '是否删除（0:否，1:是）',
 PRIMARY KEY (`id`),
 KEY `idx_did` (`ding_id`),
 KEY `idx_pid` (`parent_id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='部门表';
```



员工表：

```sql
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `ding_id` varchar(50) NOT NULL COMMENT '钉钉id',
  `name` varchar(50) NOT NULL COMMENT '员工名称',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '员工邮箱',
  `tel` varchar(20) NOT NULL COMMENT '员工手机',
  `avatar` varchar(1024) NOT NULL DEFAULT '' COMMENT '员工头像url',
  `password` varchar(50) NOT NULL COMMENT '密码',
  `status` int(1) NOT NULL DEFAULT '0' COMMENT '状态 0在职 1离职',
  `updated_at` int(12) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `created_at` int(12) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `is_deleted` int(1) NOT NULL DEFAULT '0' COMMENT '是否删除（0:否，1:是）',
  PRIMARY KEY (`id`),
  KEY `ding_id` (`ding_id`),
  KEY `tel` (`tel`)
) ENGINE=InnoDB AUTO_INCREMENT=10024 DEFAULT CHARSET=utf8mb4 COMMENT='员工表';
```



员工部门关联表：

```sql
create table dept_user (
 `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
 `dept_id` int NOT NULL COMMENT '部门id',
 `user_id` int NOT NULL COMMENT '员工id',
 `is_leader` int(1) NOT NULL DEFAULT '0' COMMENT '是否部门主管',
 `updated_at` int(12) NOT NULL DEFAULT '0' COMMENT '更新时间',
 `created_at` int(12) NOT NULL DEFAULT '0' COMMENT '创建时间',
 PRIMARY KEY (`id`),
 KEY `idx_did` (`dept_id`),
 KEY `idx_sid` (`user_id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='员工部门关联表';
```



#### 权限&角色

系统表：

```sql
create table system (
 `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
 `name` varchar(100) NOT NULL COMMENT '权限名称',
 `creator` varchar(50) NOT NULL COMMENT '创建人',
 `policy_resource_id` int NOT NULL COMMENT '策略资源id',
 `updated_at` int(12) NOT NULL DEFAULT '0' COMMENT '更新时间',
 `created_at` int(12) NOT NULL DEFAULT '0' COMMENT '创建时间',
 `is_deleted` int(1) NOT NULL DEFAULT '0' COMMENT '是否删除（0:否，1:是）',
 PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='系统表';
```

注：策略资源id 表示的是，当一个子系统的权限收到某种资源限制的时候，此 id 代表这种资源的 id



资源表：

```sql
create table resource (
 `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
 `system_id` int NOT NULL COMMENT '系统id',
 `name` varchar(100) NOT NULL DEFAULT '' COMMENT '资源名称',
 `code` varchar(100) NOT NULL DEFAULT '' COMMENT '资源码',
 `creator` varchar(50) NOT NULL COMMENT '创建人',
 `updated_at` int(12) NOT NULL DEFAULT '0' COMMENT '更新时间',
 `created_at` int(12) NOT NULL DEFAULT '0' COMMENT '创建时间',
 `is_deleted` int(1) NOT NULL DEFAULT '0' COMMENT '是否删除（0:否，1:是）',
 PRIMARY KEY (`id`),
 KEY `idx_sid` (`system_id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='资源表';
```



权限表:

```sql
create table permission (
 `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
 `code` varchar(100) NOT NULL COMMENT '权限码',
 `name` varchar(100) NOT NULL COMMENT '权限名称',
 `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '权限类型: 0子系统默认全部权限 1页面 2操作 3URL',
 `system_id` int NOT NULL DEFAULT '0' COMMENT '系统id,0表示不属于某个系统',
 `creator` varchar(50) NOT NULL COMMENT '创建人',
 `updated_at` int(12) NOT NULL DEFAULT '0' COMMENT '更新时间',
 `created_at` int(12) NOT NULL DEFAULT '0' COMMENT '创建时间',
 `is_deleted` int(1) NOT NULL DEFAULT '0' COMMENT '是否删除（0:否，1:是）',
 PRIMARY KEY (`id`),
 KEY `idx_code` (`code`),
 KEY `idx_name` (`name`),
 KEY `idx_sid` (`system_id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='权限表';
```



角色表：

```sql
create table role (
 `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
 `code` varchar(100) NOT NULL COMMENT '角色码',
 `name` varchar(100) NOT NULL COMMENT '角色名称',
 `system_id` int NOT NULL DEFAULT '0' COMMENT '系统id,0表示不属于某个系统',
 `status` int(1) NOT NULL DEFAULT '0' COMMENT '状态 0开启 1关闭',
 `creator` varchar(50) NOT NULL COMMENT '创建人',
 `updated_at` int(12) NOT NULL DEFAULT '0' COMMENT '更新时间',
 `created_at` int(12) NOT NULL DEFAULT '0' COMMENT '创建时间',
 `is_deleted` int(1) NOT NULL DEFAULT '0' COMMENT '是否删除（0:否，1:是）',
 PRIMARY KEY (`id`),
 KEY `idx_code` (`code`),
 KEY `idx_name` (`name`),
 KEY `idx_sid` (`system_id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';
```



角色权限关联表(不采用软删除，直接物理删除)

```sql
create table role_perm (
 `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
 `role_id` int NOT NULL COMMENT '角色id',
 `perm_id` int NOT NULL COMMENT '权限id',
 `updated_at` int(12) NOT NULL DEFAULT '0' COMMENT '更新时间',
 `created_at` int(12) NOT NULL DEFAULT '0' COMMENT '创建时间',
 PRIMARY KEY (`id`),
 UNIQUE KEY `idx_uniq_rp` (`role_id`,`perm_id`),
 KEY `idx_pid` (`perm_id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='员工部门关联表';
```



人员角色关联表：

```sql
create table user_role (
 `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
 `user_id` int NOT NULL COMMENT '员工id',
 `role_id` int NOT NULL COMMENT '角色id',
 `policy_id` int NOT NULL DEFAULT '0' COMMENT '策略id,0表示不受前置策略控制的角色',
 `updated_at` int(12) NOT NULL DEFAULT '0' COMMENT '更新时间',
 `created_at` int(12) NOT NULL DEFAULT '0' COMMENT '创建时间',
 PRIMARY KEY (`id`),
 UNIQUE KEY `idx_uid_rid` (`user_id`,`role_id`),
 KEY `idx_rid` (`role_id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='员工角色关联表';
```



人员权限关联表：

```sql
create table user_perm (
 `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
 `user_id` int NOT NULL COMMENT '员工id',
 `perm_id` int NOT NULL COMMENT '权限id',
 `policy_id` int NOT NULL DEFAULT '0' COMMENT '策略id,0表示不受前置策略控制的权限',
 `scope_id` int(11) NOT NULL DEFAULT '0' COMMENT '资源范围policy id,0表示不受后置策略控制的权限',
 `updated_at` int(12) NOT NULL DEFAULT '0' COMMENT '更新时间',
 `created_at` int(12) NOT NULL DEFAULT '0' COMMENT '创建时间',
 PRIMARY KEY (`id`),
 UNIQUE KEY `idx_uid_pid` (`user_id`,`perm_id`),
 KEY `idx_pid` (`perm_id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='员工权限关联表';
```



策略表：

```sql
create table policy (
 `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
 `updated_at` int(12) NOT NULL DEFAULT '0' COMMENT '更新时间',
 `created_at` int(12) NOT NULL DEFAULT '0' COMMENT '创建时间',
 PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='策略表';
```



策略资源表:

```sql
create table policy_resource (
 `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
 `policy_id` int NOT NULL COMMENT '策略id',
 `resource_id` int NOT NULL COMMENT '资源id',
 `entity_id` varchar(50) NOT NULL COMMENT '资源实体id',
 `updated_at` int(12) NOT NULL DEFAULT '0' COMMENT '更新时间',
 `created_at` int(12) NOT NULL DEFAULT '0' COMMENT '创建时间',
 PRIMARY KEY (`id`),
 KEY `idx_rid` (`policy_id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='策略表';
```



### 接入端技术方案

#### 一些想法：

- 怎么样在内存中实现组织架构的快速查找，例如查询一个人的所有父部门、子部门，查询一个部门下的所有子部门、成员等。
- 怎么快速的针对不同的角色和权限类型，查询对应拥有的角色和权限以及查询每个权限对应的数据资源。

### 前后端撕逼结果

- 权限的范围按照子系统划分，每次获取当前系统的权限，不包含子系统的权限。至于当前系统中，哪个\那部分功能属于 子系统，目前以 前后端+产品 自定义形式确定，暂不提供一个管理子系统的平台。
- 将带有策略限制的角色或权限当做子系统权限的部分。也即通常情况下，只有在获取某个（子）系统的权限时， 附带策略上的限制。其余权限均默认为整体的权。同时，将某个策略的相关资源设置为子系统的一个属性，这样在使用时 就不需要每次都区分这个 id 是什么资源的 id。也即当我访问某一子系统的权限时，若子系统没有资源区分则按照正常 逻辑拿到所有权限；当系统有资源区分时，请求权限的时候必须带上资源 id，默认即为子系统相关资源的 id。这样就 解决了之前 "使用权限系统时，什么时候用带上资源 id，带上什么资源的 id" 的问题。这样虽然牺牲了一部分自由度， 但使用场景更为明确，使得权限系统使用方能更直接的获得自己所想要的权限，且接口不再需要区分 资源 id 的特异性 从而变得更为简单。

### 上线前准备

- 跑钉钉脚本同步钉钉组织架构
- 跑基础子系统权限脚本

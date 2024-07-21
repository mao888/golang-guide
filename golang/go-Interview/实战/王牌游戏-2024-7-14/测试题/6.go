package main

import (
	"fmt"
	"sync"
	"time"
)

// 6、设计一个好友管理系统，用户可以在系统中添加好友、发送消息、查看好友动态等。要求系统记录用户之间的交互，并提供查询功能。

// User 用户结构体
type User struct {
	ID       string          // 用户ID
	Name     string          // 用户名
	Friends  map[string]bool // 好友ID列表
	Messages []Message       // 接收到的消息列表
	Dynamics []Dynamic       // 用户动态列表
	mutex    sync.Mutex      // 锁，用于并发控制
}

// Friendship 好友关系结构体（实际在这个简化版本中未直接使用，但可用于更复杂的场景）
type Friendship struct {
	UserID    string
	FriendID  string
	CreatedAt time.Time
}

// Message 消息结构体
type Message struct {
	From    string    // 发送者ID
	To      string    // 接收者ID
	Content string    // 消息内容
	Time    time.Time // 发送时间
}

// Dynamic 动态结构体
type Dynamic struct {
	UserID    string
	Content   string    // 动态内容
	Timestamp time.Time // 发布时间
}

// AddFriend 添加好友
func (u *User) AddFriend(friendID string) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.Friends[friendID] = true
}

// SendMessage 发送消息
func (u *User) SendMessage(toUserID, content string) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	// 在实际应用中，你可能需要验证toUserID是否为u的好友
	msg := Message{
		From:    u.ID,
		To:      toUserID,
		Content: content,
		Time:    time.Now(),
	}
	u.Messages = append(u.Messages, msg)
	// 这里不模拟将消息发送给toUserID对应的用户，仅记录在本地
}

// PostDynamic 发布动态
func (u *User) PostDynamic(content string) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	dyn := Dynamic{
		UserID:    u.ID,
		Content:   content,
		Timestamp: time.Now(),
	}
	u.Dynamics = append(u.Dynamics, dyn)
}

// ViewMessages 查看收到的消息（简化版，不过滤）
func (u *User) ViewMessages() {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	for _, msg := range u.Messages {
		fmt.Printf("From: %s, To: %s, Content: %s, Time: %s\n", msg.From, msg.To, msg.Content, msg.Time.Format(time.RFC3339))
	}
}

// ViewDynamics 查看用户的动态
func (u *User) ViewDynamics() {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	for _, dyn := range u.Dynamics {
		fmt.Printf("UserID: %s, Content: %s, Timestamp: %s\n", dyn.UserID, dyn.Content, dyn.Timestamp.Format(time.RFC3339))
	}
}

func main() {
	// 示例用户
	user1 := User{
		ID:       "1",
		Name:     "Alice",
		Friends:  make(map[string]bool),
		Messages: []Message{},
		Dynamics: []Dynamic{},
	}
	user2 := User{
		ID:       "2",
		Name:     "Bob",
		Friends:  make(map[string]bool),
		Messages: []Message{},
		Dynamics: []Dynamic{},
	}

	// 添加好友关系（双向）
	user1.AddFriend("2")
	user2.AddFriend("1")

	// Alice 给 Bob 发送消息
	user1.SendMessage("2", "Hello Bob!")

	// Bob 发布动态
	user2.PostDynamic("Just had a great day at the beach!")
}

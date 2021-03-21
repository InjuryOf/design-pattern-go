package behavior_type

import (
	"testing"
)

func TestObserver(t *testing.T)  {
	// 初始化生产者对象
	shirtItem := newItem("Nike Shirt")

	// 初始化消费者对象
	observerFirst := &customer{id: "zhangsan@gmail.com"}
	observerSecond := &customer{id: "lisi@gmail.com"}

	// 建立订阅关系
	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	// 发送通知
	shirtItem.updateAvailability()
}
package behavior_type

import "testing"

// 方式一、初始设置责任链调用关系
func TestChain(t *testing.T)  {
	ah := &AdHanler{}
	yh := &YellowHanler{}
	sh := &SensitiveHanler{}
	ah.handler = yh
	yh.handler = sh
	ah.Handler("我是在正常不过内容，是个广告，是个涉黄，是个敏感词，白白净净")
}

// 方式二、每次链节点调用后指定下一个节点
func TestChain2(t *testing.T)  {
	//ah := &AdHanler{}
	//yh := &YellowHanler{}
	//sh := &SensitiveHanler{}
	//
	//newContext:= ah.Handler("我是在正常不过内容，是个广告，是个涉黄，是个敏感词，白白净净")
	//ah.next(yh, newContext)

	//newContext = yh.Handler(newContext)
	//yh.next(sh, newContext)

	//sh.Handler(newContext)
}
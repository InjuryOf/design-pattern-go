package behavior_type

import (
	"fmt"
	"strings"
)

type Handler interface {
	Handler(context string)
	next(handler Handler, context string)
}

// 广告过滤
type AdHanler struct {
	handler Handler
}

func (ah AdHanler) Handler(context string) {
	fmt.Println("执行广告过滤...")
	newContext := strings.ReplaceAll(context, "广告", "***")
	fmt.Println(newContext)
	ah.next(ah.handler, newContext)
}

func (ah AdHanler) next(handler Handler, context string) {
	if ah.handler != nil {
		ah.handler.Handler(context)
	}
}

// 涉黄过滤
type YellowHanler struct {
	handler Handler
}

func (yh YellowHanler) Handler(context string) {
	fmt.Println("执行涉黄过滤...")
	newContext := strings.ReplaceAll(context, "涉黄", "***")
	fmt.Println(newContext)
	yh.next(yh.handler, newContext)
}

func (yh YellowHanler) next(handler Handler, context string) {
	if yh.handler != nil {
		yh.handler.Handler(context)
	}
}

// 敏感词过滤
type SensitiveHanler struct {
	handler Handler
}

func (sh SensitiveHanler) Handler(context string) {
	fmt.Println("执行敏感词过滤...")
	newContext := strings.ReplaceAll(context, "敏感", "***")
	fmt.Println(newContext)
	sh.next(sh.handler, newContext)
}

func (sh SensitiveHanler) next(handler Handler, context string) {
	if sh.handler != nil {
		sh.handler.Handler(context)
	}
}

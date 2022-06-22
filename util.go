package wgolib

import (
	"bytes"
	"encoding/json"
	"reflect"
	"runtime"
	"sync"

	"github.com/changmu/wgolib/errs"
	"github.com/changmu/wgolib/wlog"
)

// JSONEncode 将对象序列化为json str
// data is a struct, params: true:indent false:no_indent:default
func JSONEncode(data interface{}, params ...interface{}) string {
	jsons, err := json.Marshal(data) //转换成JSON返回的是byte[]
	if err != nil {
		return ""
	}
	if len(params) == 0 {
		return string(jsons)
	}

	if reflect.TypeOf(params[0]).Kind() == reflect.Bool && params[0].(bool) == false {
		return string(jsons)
	}
	var out bytes.Buffer
	err = json.Indent(&out, jsons, "", "  ")
	if err != nil {
		return string(jsons)
	}
	return out.String()
}

// RmDupAndEmptyOfList 列表元素去重并去空
func RmDupAndEmptyOfList(list []string) []string {
	mp := map[string]bool{}
	for _, item := range list {
		if item != "" {
			mp[item] = true
		}
	}

	var listUniq []string
	for k := range mp {
		listUniq = append(listUniq, k)
	}

	return listUniq
}

// GoAndWait 封装更安全的多并发调用, 启动goroutine并等待所有处理流程完成，自动recover
// 返回值error返回的是多并发协程里面第一个返回的不为nil的error，主要用于关键路径判断，当多并发协程里面有一个是关键路径且有失败则返回err，其他非关键路径并发全部返回nil
func GoAndWait(handlers ...func() error) error {
	const PanicBufLen = 4 << 10
	var (
		wg   sync.WaitGroup
		once sync.Once
		err  error
	)
	for _, f := range handlers {
		wg.Add(1)
		go func(handler func() error) {
			defer func() {
				if e := recover(); e != nil {
					buf := make([]byte, PanicBufLen)
					buf = buf[:runtime.Stack(buf, false)]
					log := wlog.NewLogger()
					log.Errorf("[PANIC]%v\n%s\n", e, buf)
					once.Do(func() {
						err = errs.New(errs.RetServerSystemErr, "panic found in call handlers")
					})
				}
				wg.Done()
			}()
			if e := handler(); e != nil {
				once.Do(func() {
					err = e
				})
			}
		}(f)
	}
	wg.Wait()
	return err
}

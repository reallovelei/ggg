package app
//
//import (
//    "errors"
//
//)
//
//// HadeAppProvider 提供 App 的具体实现方法
//type HadeAppProvider struct {
//    BaseFolder string
//}
//
//
//
//// Params 获取初始化参数
//func (h *HadeAppProvider) Params(container framework.Container) []interface{} {
//    return []interface{}{container, h.BaseFolder}
//}
//
//
//
//// NewHadeApp 初始化 HadeApp
//func NewHadeApp(params ...interface{}) (interface{}, error) {
//    if len(params) != 2 {
//        return nil, errors.New("param error")
//    }
//    // 有两个参数，一个是容器，一个是 baseFolder
//    container := params[0].(framework.Container)
//    baseFolder := params[1].(string)
//    return &HadeApp{baseFolder: baseFolder, container: container}, nil
//}


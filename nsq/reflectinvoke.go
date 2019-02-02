package nsqexample

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"unicode"
	"unicode/utf8"
)

type MethodMap struct {
	Methods map[string]*MethodInfo
}

type MethodInfo struct {
	Method reflect.Method
	Host   reflect.Value
	Idx    int
}

type Request struct {
	FuncName string        `json:"func_name"`
	Params   []interface{} `json:"params"`
}

type Response struct {
	FuncName  string        `json:"func_name"`
	Data      []interface{} `json:"data"`
	ErrorCode int           `json:"errorcode"`
}

type Reflectinvoker struct {
	Methods map[string]*MethodInfo
}

var errorInfo map[int]string = make(map[int]string)

const (
	errorCode_JsonError = 1 + iota
	errorCode_MethodNotFound
	errorCode_ParameterNotMatch
)

func init() {
	errorInfo[errorCode_JsonError] = "JsonError"
	errorInfo[errorCode_MethodNotFound] = "MethodNotFound"
	errorInfo[errorCode_ParameterNotMatch] = "ParameterNotMatch"
}

func ErrorMsg(errorCode int) string {
	return errorInfo[errorCode]
}

func NewReflectinvoker() *Reflectinvoker {
	return &Reflectinvoker{
		Methods: make(map[string]*MethodInfo),
	}
}

func (r *Reflectinvoker) RegisterMethod(v interface{}) {
	reflectType := reflect.TypeOf(v)
	host := reflect.ValueOf(v)

	for i := 0; i < reflectType.NumMethod(); i++ {
		m := reflectType.Method(i)

		char, _ := utf8.DecodeRuneInString(m.Name)
		//非导出函数不注册
		if !unicode.IsUpper(char) {
			continue
		}
		//fmt.Println("reflectType.Method:", m)
		r.Methods[m.Name] = &MethodInfo{Method: m, Host: host, Idx: i}
		//fmt.Println(":m", m, " :host:", host, " idx:", strconv.Itoa(i))
	}

}

func (r *Reflectinvoker) InvokeByReflectArgs(funcName string, par []reflect.Value) []reflect.Value {

	return r.Methods[funcName].Host.MethodByName(funcName).Call(par)
}

func (r *Reflectinvoker) InvokeByInterfaceArgs(funcName string, Params []interface{}) []reflect.Value {

	paramsValue, err := convertParam(r.Methods[funcName], Params)

	if err != nil {
		return nil
	}

	return r.Methods[funcName].Host.MethodByName(funcName).Call(paramsValue)
}

func (r *Reflectinvoker) InvokeByJson(byteData []byte) []byte {

	req := &Request{}
	err := json.Unmarshal(byteData, req)

	resultData := &Response{}

	if err != nil {
		resultData.ErrorCode = errorCode_JsonError
	} else {
		resultData.FuncName = req.FuncName

		methodInfo, found := r.Methods[req.FuncName]
		fmt.Println("methodInfo:", methodInfo)
		if found {

			paramsValue, err := convertParam(methodInfo, req.Params)
			fmt.Println("paramsValue:", paramsValue)
			if err != nil {

				resultData.ErrorCode = errorCode_ParameterNotMatch
			} else {
				resultData = InvokeByValues(methodInfo, paramsValue)
				fmt.Println("resultData:", resultData)
			}

		} else {
			resultData.ErrorCode = errorCode_MethodNotFound
		}

	}

	data, _ := json.Marshal(resultData)

	return data
}

func convertParamType(v interface{}, targetType reflect.Type) (
	targetValue reflect.Value, ok bool) {
	defer func() {
		if re := recover(); re != nil {
			ok = false
			fmt.Println(re)
		}
	}()

	ok = true

	if targetType.Kind() == reflect.Interface ||
		targetType.Kind() == reflect.TypeOf(v).Kind() {

		targetValue = reflect.ValueOf(v)

	} else if reflect.TypeOf(v).Kind() == reflect.Float64 {
		f := v.(float64)
		switch targetType.Kind() {
		case reflect.Int:
			targetValue = reflect.ValueOf(int(f))
		case reflect.Uint8:
			targetValue = reflect.ValueOf(uint8(f))
		case reflect.Uint16:
			targetValue = reflect.ValueOf(uint16(f))
		case reflect.Uint32:
			targetValue = reflect.ValueOf(uint32(f))
		case reflect.Uint64:
			targetValue = reflect.ValueOf(uint64(f))
		case reflect.Int8:
			targetValue = reflect.ValueOf(int8(f))
		case reflect.Int16:
			targetValue = reflect.ValueOf(int16(f))
		case reflect.Int32:
			targetValue = reflect.ValueOf(int32(f))
		case reflect.Int64:
			targetValue = reflect.ValueOf(int64(f))
		case reflect.Float32:
			targetValue = reflect.ValueOf(float32(f))
		default:
			ok = false
		}
	} else {
		ok = false
	}

	return
}

func convertParam(methodInfo *MethodInfo, Params []interface{}) ([]reflect.Value, error) {

	if len(Params) != methodInfo.Method.Type.NumIn()-1 {
		return nil, errors.New("convertParam number error")
	}

	paramsValue := make([]reflect.Value, 0, len(Params))
	//跳过 receiver
	for i := 1; i < methodInfo.Method.Type.NumIn(); i++ {
		inParaType := methodInfo.Method.Type.In(i)
		value, ok := convertParamType(Params[i-1], inParaType)
		if !ok {
			return nil, errors.New("convertParamType error")
		}
		paramsValue = append(paramsValue, value)
	}

	return paramsValue, nil
}

func InvokeByValues(methodInfo *MethodInfo, params []reflect.Value) *Response {

	data := &Response{}
	data.FuncName = methodInfo.Method.Name
	result := methodInfo.Host.Method(methodInfo.Idx).Call(params)

	for _, x := range result {
		data.Data = append(data.Data, x.Interface())
		fmt.Println("result:", x, "::", data.Data)
	}

	return data
}

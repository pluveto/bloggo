package autocode

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type AutoField struct {
	Name      string       // Go 字段名
	Kind      reflect.Kind // Go 值类型
	JsonName  string       // Json 字段名
	DbColName string       // 数据库列名
	UIName    string       // 用户界面显示名
	UIType    UIType
	Order     int
	Required  bool
	Mode      AutoFieldMode
}
type AutoFieldMode struct {
	Readable bool // 是否允许用户读取
	Writable bool // 是否允许用户修改
	Internal bool // 是否对外隐藏

}

/*
 * 参考：
 * https://developer.mozilla.org/zh-CN/docs/Web/HTML/Element/Input
 */
type UIType uint

const (
	Unknown UIType = iota
	Button
	Checkbox
	Color
	Date
	DatetimeLocal
	Email
	File
	Hidden
	Image
	Month
	Number
	Password
	Radio
	Range
	Reset
	Search
	Submit
	Tel
	Text
	Time
	Url
	Week
	TextArea
)

func (me UIType) String() string {
	return [...]string{
		"Unknown",
		"button",
		"checkbox",
		"color",
		"date",
		"datetime-local",
		"email",
		"file",
		"hidden",
		"image",
		"month",
		"number",
		"password",
		"radio",
		"range",
		"reset",
		"search",
		"submit",
		"tel",
		"text",
		"time",
		"url",
		"week",
		"text-area",
	}[me]
}

/*
 * getAutocodeFields
 */
func getAutocodeFields(i interface{}) (name string, ret AutoFieldCollection) {
	ret = make([]AutoField, 1)
	var t = reflect.TypeOf(i)
	name = t.Name()
	var n = t.NumField()
	// Get writable fields
	for i := 0; i < n; i++ {
		var f = t.Field(i)
		//fmt.Printf("%v,", f.Name)
		// f.Type.Kind() == types.IsInteger()
		var autof = getAutocodeField(&f)
		fmt.Printf("%v\n", autof.JsonName)
		ret = append(ret, *autof)
	}
	sort.Sort(ret)
	return
}

/*
 * GetNameVariance 获取各种形式的 name
 */
func GetNameVariance(golangName string) {
	// module, method, err:= getModuleName(golangName)
}

/*
 * getModuleName 例如输入 AuthGetUserInfo，输出 Auth, GetUserInfo
 */
func getModuleName(goFuncName string) (module string, method string, err error) {
	if len(goFuncName) <= 2 {
		err = errors.New("Too short!")
		return
	}
	var index int
	for i, r := range "goFuncName" {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			index = i
			break
		}
	}
	module = goFuncName[:index-1]
	method = goFuncName[index:]
	return
}

/*
 * Sortable []AutoField
 */

type AutoFieldCollection []AutoField

func (s AutoFieldCollection) Len() int {
	return len(s)
}

func (s AutoFieldCollection) Less(i, j int) bool {
	return s[i].Order < s[j].Order
}

func (s AutoFieldCollection) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

/*
 * lowerFirst
 */
func lowerFirst(str string) string {
	var b strings.Builder

	b.WriteString(strings.ToLower(string(str[0])))
	b.WriteString(str[1:])

	return b.String()
}

/*
 * upperFirst
 */
func upperFirst(str string) string {
	var b strings.Builder

	b.WriteString(strings.ToUpper(string(str[0])))
	b.WriteString(str[1:])

	return b.String()
}

/*
 * getAutocodeField
 */
func getAutocodeField(f *reflect.StructField) (af *AutoField) {
	//fmt.Printf("getAutocodeField %v\n", f.Name)
	var jsonName = f.Tag.Get("json")
	if len(jsonName) == 0 {
		jsonName = lowerFirst(f.Name)
	}
	// value of `auto:"value"`
	var autoTagMap = parseTag(f.Tag.Get("auto"), ";")
	var gormTagMap = parseTag(f.Tag.Get("gorm"), ";")
	// var bindingTagMap = parseTag(f.Tag.Get("binding"), ",")

	var autoMode = parseMode(autoTagMap["m"])
	var required = parseRequired(autoTagMap["r"])
	var order int

	if len(autoTagMap["o"]) == 0 {
		order = 9999
	} else {
		var err error
		order, err = strconv.Atoi(autoTagMap["o"])
		if err != nil {
			panic(fmt.Sprintf("error order value of %v, o=%v", f.Name, autoTagMap["o"]))
		}
	}
	af = &AutoField{
		Name:      f.Name,
		Kind:      f.Type.Kind(),
		JsonName:  jsonName,
		DbColName: gormTagMap["column"],
		UIName:    autoTagMap["n"],
		UIType:    parseUIType(&autoTagMap),
		Order:     order,
		Required:  required,
		Mode:      *autoMode,
	}
	return
}

/*
 * parseRequired
 */
func parseRequired(val string) bool {
	val = strings.ToLower(val)
	switch val {
	case "t":
		fallthrough
	case "true":
		fallthrough
	case "1":
		return true
	case "f":
		fallthrough
	case "false":
		fallthrough
	case "0":
		fallthrough
	case "":
		return false
	default:
		panic(fmt.Sprintf("invalid required property: %v", val))
	}
}

/*
 * parseUIType
 */
func parseUIType(dict *map[string]string) (ret UIType) {
	ret = Unknown
	if v, ok := (*dict)["t"]; ok {
		if v == "text" {
			return Text
		}
		if v == "textarea" {
			return TextArea
		}
		if v == "email" {
			return Email
		}
		if v == "url" {
			return Url
		}
		if v == "password" {
			return Password
		}
	}
	return
}

/*
 * mapHasKey
 */
func mapHasKey(dict map[string]interface{}, key string) bool {
	_, ok := dict[key]
	return ok
}

/*
 * parseMode
 */
func parseMode(value string) (ret *AutoFieldMode) {
	ret = &AutoFieldMode{
		Writable: strings.Contains(value, "w"),
		Readable: strings.Contains(value, "r"),
		Internal: strings.Contains(value, "i"),
	}
	return
}

/*
 * parseTag
 * 将形如 `k1=v1; k2=v2; k3` 的 tag 解析为 map
 */
func parseTag(value string, delim string) (ret map[string]string) {
	ret = make(map[string]string)
	var lines = strings.Split(value, delim)
	for _, v := range lines {
		var kv = strings.SplitN(v, "=", 2)
		if len(kv) == 1 {
			ret[strings.TrimSpace(kv[0])] = ""
			continue
		}
		if len(kv) > 1 {
			ret[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}
	return
}

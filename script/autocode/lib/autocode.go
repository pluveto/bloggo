package autocode

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"text/template"
)

var tpls *template.Template

/*
 * init
 */
func init() {
	cwd, err := os.Getwd()
	if err != nil {
		panic("failed to get cwd!")
	}
	tpls = template.New("").Delims("[[", "]]")
	tpls, err = tpls.ParseFiles(
		cwd+"/template/text.jsx",
		cwd+"/template/textarea.jsx",
	)
	if err != nil {
		panic(fmt.Sprintf("invalid template: %v", err))
	}
}

/*
 * GenerateForm
 */
func GenerateForm(i interface{}) {
	var name, fields = getAutocodeFields(i)
	for _, field := range fields {
		if field.UIType == Unknown {
			continue
		}
		str, err := GenerateFormField(&field)
		if err != nil {
			fmt.Printf("Errro when generate for field %v, %v\n", field.Name, err.Error())
			continue
		}
		fmt.Println(str)
	}
}

/*
 * GenerateJsApi
 */
func GenerateJsApi(i interface{}) {
	var _, fields = getAutocodeFields(i)
	for _, field := range fields {
		if field.UIType == Unknown {
			continue
		}
		str, err := GenerateFormField(&field)
		if err != nil {
			fmt.Printf("Errro when generate for field %v, %v\n", field.Name, err.Error())
			continue
		}
		fmt.Println(str)
	}
}

/*
 * GenerateFormField
 */
func GenerateFormField(field *AutoField) (ret string, err error) {
	// fmt.Printf("GenerateFormField -> field %v\n", field.Name)
	switch field.UIType {
	case Unknown:
		panic("Invalid UIType: Unknown")
	case Button:
		fallthrough
	case Checkbox:
		fallthrough
	case Color:
		fallthrough
	case Date:
		fallthrough
	case DatetimeLocal:
		fallthrough
	case Email:
		fallthrough
	case File:
		fallthrough
	case Hidden:
		fallthrough
	case Image:
		fallthrough
	case Month:
		fallthrough
	case Number:
		fallthrough
	case Password:
		fallthrough
	case Radio:
		fallthrough
	case Range:
		fallthrough
	case Reset:
		fallthrough
	case Search:
		fallthrough
	case Submit:
		fallthrough
	case Tel:
		fallthrough
	case Time:
		fallthrough
	case Url:
		fallthrough
	case Week:
		fallthrough
	case Text:
		var tpl bytes.Buffer
		if err = tpls.ExecuteTemplate(&tpl, "text.jsx", &map[string]interface{}{
			"name":        field.JsonName,
			"label":       field.UIName,
			"placeholder": "",
			"required":    strconv.FormatBool(field.Required),
			"type":        field.UIType.String(),
		}); err != nil {
			return
		}
		ret = tpl.String()
	case TextArea:
		var tpl bytes.Buffer
		if err = tpls.ExecuteTemplate(&tpl, "textarea.jsx", &map[string]interface{}{
			"name":        field.JsonName,
			"label":       field.UIName,
			"placeholder": "",
			"required":    strconv.FormatBool(field.Required),
		}); err != nil {
			return
		}
		ret = tpl.String()
	default:
		panic("Invalid UIType")
	}
	return
}

/*
 * getWritableFields
 */
func getWritableFields() {}

/*
 * GetTextAreaTemplate
 * require param: .FieldName .FieldText .HintText
 */
func GetTextAreaTemplate() *template.Template {
	var tpl = `
  <div className="mb-6">
	<label className="block text-sm font-medium mb-2" htmlFor="{{.FieldName}}">{{ .FieldText }}</label>
	<textarea className="block w-full px-4 py-3 mb-2 text-sm placeholder-gray-500 bg-white border rounded" 
	 name="{{.FieldName}}" placeholder="{{.HintText}}"></textarea>
  </div>
	`
	return template.Must(template.New("textarea").Delims("<%", "%>").Parse(tpl))
}

package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type Type struct {
	f1     string "f one"
	f2     string
	f3     string `f three`
	f4, f5 int64  `f four and five`
	f6     string `one:"1" two:"2"blank:""`
}

func Test(t *testing.T) {
	type1 := reflect.TypeOf(Type{})  //获取变量对应类型reflect.Type
	f1, _ := type1.FieldByName("f1") //根据属性名获取类型的属性reflect.StructField
	fmt.Println(f1.Tag)              // f one 属性的标签StructTag
	f4, _ := type1.FieldByName("f4")
	fmt.Println(f4.Tag) // f four and five
	f5, _ := type1.FieldByName("f5")
	fmt.Println(f5.Tag) // f four and five
	f6, _ := type1.FieldByName("f6")
	fmt.Println(f6.Tag)           // one:"1" two:"2"blank:"" golang常用kv值格式的标签
	v, ok := f6.Tag.Lookup("one") //根据key获取value，value不存在第二个参数返回false
	fmt.Printf("%s, %t\n", v, ok) // 1, true
	v, ok = f6.Tag.Lookup("blank")
	fmt.Printf("%s, %t\n", v, ok) // , true
	v, ok = f6.Tag.Lookup("five")
	fmt.Printf("%s, %t\n", v, ok)   // , false
	fmt.Println(f6.Tag.Get("one"))  //1 根据key获取value
	fmt.Println(f6.Tag.Get("five")) //key不存在空字符串
}

package normal

import (
	"github.com/bmizerany/assert"
	"reflect"
	"testing"
)

// 遍历
// method, field
func TestTraversal(t *testing.T) {
	sparrow := &Bird{"sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem()

	assert.Equal(t, s.NumField(), 2)
	assert.Equal(t, s.NumMethod(), 2)

	// field, 字段
	types := s.Type()
	fields := s.Field
	i := 0
	assert.Equal(t, types.Field(i).Name, "Name")
	assert.Equal(t, fields(i).Type().String(), "string")
	assert.Equal(t, fields(i).Interface(), "sparrow")

	i = 1
	assert.Equal(t, types.Field(i).Name, "LifeExpectance")
	assert.Equal(t, fields(i).Type().String(), "int")
	assert.Equal(t, fields(i).Interface(), 3)

	// method, 方法
	methods := s.Method
	i = 0
	assert.Equal(t, methods(i).String(), "<func() string Value>") // func (b Bird) Fly() string
	i = 1
	assert.Equal(t, methods(i).String(), "<func(string) Value>") // func (b Bird) FlyA(a string)
}

// 修改简单字段的值
func TestField(t *testing.T) {
	sparrow := &Bird{"sparrow", 3}
	x := 2
	v1 := reflect.ValueOf(&x) // 非索引类型记得加指针
	//t.Log(*v1.Interface().(*int))
	assert.Equal(t, *v1.Interface().(*int), 2)
	v1.Elem().SetInt(100)
	assert.Equal(t, *v1.Interface().(*int), 100)

	// 修改字段的值
	v := reflect.ValueOf(sparrow)
	can := v.Elem().CanSet()
	assert.Equal(t, can, true)

	v.Elem().Field(1).SetInt(2) // 按索引找到字段并修改
	assert.Equal(t, v.Elem().FieldByName("LifeExpectance").Int(), int64(2))

	v.Elem().FieldByName("Name").SetString("abc") // 按名字找到字段并修改
	assert.Equal(t, v.Elem().Field(0).String(), "abc")
}

//动态调用方法
func TestDynamicMethod(t *testing.T) {
	sparrow := &Bird{"sparrow", 3}
	v := reflect.ValueOf(sparrow)

	// 有参数方法
	method := v.MethodByName("FlyA")
	args := []reflect.Value{reflect.ValueOf("前缀")}
	returnFlyA := method.Call(args)
	assert.Equal(t, len(returnFlyA), 0)

	// 无参数方法
	method2 := v.MethodByName("Fly")
	args2 := []reflect.Value{}
	returnFly := method2.Call(args2)
	assert.Equal(t, len(returnFly), 1)
	assert.Equal(t, returnFly[0].String(), "returnFly")
}

func TestTag(t *testing.T) {
	var u User
	typeof := reflect.TypeOf(u)

	i := 0
	field0 := typeof.Field(i)
	json0 := field0.Tag.Get("json")
	bson0 := field0.Tag.Get("bson")
	assert.Equal(t, json0, "name")
	assert.Equal(t, bson0, "b_name")
	t.Log(json0, bson0)
	i = 1
	field1 := typeof.Field(i)
	json1 := field1.Tag.Get("json")
	bson1 := field1.Tag.Get("bson")
	assert.Equal(t, json1, "age")
	assert.Equal(t, bson1, "b_age")
	t.Log(json1, bson1)
}

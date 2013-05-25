
package assrt

import (
	"testing"
	"runtime"
	"reflect"
	"fmt"
)

type Assert struct {
	*testing.T
}

func NewAssert(t *testing.T) (*Assert) {
	return &Assert{t}
}

func (ast *Assert) Nil(value interface {}, logs ...interface {}) {
	ast.nilAssert(false, true, value, logs...)
}

func (ast *Assert) MustNil(value interface {}, logs ...interface {}) {
	ast.nilAssert(true, true, value, logs...)
}

func (ast *Assert) NotNil(value interface {}, logs ...interface {}) {
	ast.nilAssert(false, false, value, logs...)
}

func (ast *Assert) MustNotNil(value interface {}, logs ...interface {}) {
	ast.nilAssert(true, false, value, logs...)
}

func (ast *Assert) nilAssert(fatal bool, isNil bool, value interface {}, logs ...interface {}) {
	if isNil != (value == nil || reflect.ValueOf(value).IsNil()) {
		ast.logCaller()
		if len(logs) > 0 {
			ast.Log(logs...)
		}else {
			if isNil {
				ast.Log("value is not nil:", value)
			}else {
				ast.Log("value is nil")
			}
		}
		ast.failIt(fatal)
	}
}

func (ast *Assert) True(boolValue bool, logs ...interface {}) {
	ast.trueAssert(false, boolValue, logs...)
}

func (ast *Assert) MustTrue(boolValue bool, logs ...interface {}) {
	ast.trueAssert(true, boolValue, logs...)
}

func (ast *Assert) trueAssert(fatal bool, value bool, logs ...interface {}) {
	if !value {
		ast.logCaller()
		if len(logs) > 0 {
			ast.Log(logs...)
		}else {
			ast.Logf("value is not true")
		}
		ast.failIt(fatal)
	}
}

func (ast *Assert) Equal(expected, actual interface {}, logs ...interface {}) {
	ast.equalAssert(false, true, expected, actual, logs...)
}
func (ast *Assert) MustEqual(expected, actual interface {}, logs ...interface {}) {
	ast.equalAssert(true, true, expected, actual, logs...)
}

func (ast *Assert) NotEqual(expected, actual interface {}, logs ...interface {}) {
	ast.equalAssert(false, false, expected, actual, logs...)
}

func (ast *Assert) MustNotEqual(expected, actual interface {}, logs ...interface {}) {
	ast.equalAssert(true, false, expected, actual, logs...)
}


func (ast *Assert) EqualSprint(expected, actual interface {}, logs ...interface {}) {
	ast.equalAssert(false, true, fmt.Sprint(expected), fmt.Sprint(actual), logs...)
}

func (ast *Assert) MustEqualSprint(expected, actual interface {}, logs ...interface {}) {
	ast.equalAssert(true, true, fmt.Sprint(expected), fmt.Sprint(actual), logs...)
}

func (ast *Assert) NotEqualSprint(expected, actual interface {}, logs ...interface {}) {
	ast.equalAssert(false, false, fmt.Sprint(expected), fmt.Sprint(actual), logs...)
}

func (ast *Assert) MustNotEqualSprint(expected, actual interface {}, logs ...interface {}) {
	ast.equalAssert(true, false, fmt.Sprint(expected), fmt.Sprint(actual), logs...)
}

func (ast *Assert) equalAssert(fatal bool, isEqual bool, expected, actual interface {}, logs ...interface {}) {
	expected = normalizeValue(expected)
	actual = normalizeValue(actual)
	if isEqual != (reflect.DeepEqual(expected,actual)) {
		ast.logCaller()
		if len(logs) > 0 {
			ast.Log(logs...)
		} else {
			if isEqual {
				ast.Log("Values not equal")
			}else {
				ast.Log("Values equal")
			}
		}
		ast.Log("Expected: ", expected)
		ast.Log("Actual: ", actual)
		ast.failIt(fatal)
	}
}

func (ast *Assert) Zero(value interface {}, logs ...interface {}) {
	ast.zeroAssert(false, true, false, value, logs...)
}
func (ast *Assert) MustZero(value interface {}, logs ...interface {}) {
	ast.zeroAssert(true, true, false, value, logs...)
}
func (ast *Assert) NotZero(value interface {}, logs ...interface {}) {
	ast.zeroAssert(false, false, false, value, logs...)
}
func (ast *Assert) MustNotZero(value interface {}, logs ...interface {}) {
	ast.zeroAssert(true, false, false, value, logs...)
}

func (ast *Assert) ZeroLen(value interface {}, logs ...interface {}) {
	ast.zeroAssert(false, true, true, value, logs...)
}

func (ast *Assert) MustZeroLen(value interface {}, logs ...interface {}) {
	ast.zeroAssert(true, true, true, value, logs...)
}

func (ast *Assert) PositiveLen(value interface {}, logs ...interface {}) {
	ast.zeroAssert(false, false, true, value, logs...)
}

func (ast *Assert) MustPositiveLen(value interface {}, logs ...interface {}) {
	ast.zeroAssert(true, false, true, value, logs...)
}

func (ast *Assert) zeroAssert(fatal bool, isZero bool, length bool, value interface {}, logs ...interface {}) {
	var name string
	var integerValue int
	value = normalizeValue(value)
	v:= reflect.Indirect(reflect.ValueOf(value))
	if length {
		name = "Length"
		integerValue = v.Len()
	}else{
		name = "Value"
		integerValue = int(v.Int())
	}
	if isZero != (integerValue == 0) {
		ast.logCaller()
		if len(logs) > 0 {
			ast.Log(logs...)
		} else {
			if isZero {
				ast.Log(name, "is not zero:",value)
			}else {
				ast.Log(name, "is zero.")
			}
		}
		ast.failIt(fatal)
	}
}

func (ast *Assert) OneLen(value interface {}, logs ...interface {}) {
	ast.oneLenAssert(false, value, logs...)
}

func (ast *Assert) MustOneLen(value interface {}, logs ...interface {}) {
	ast.oneLenAssert(true, value, logs...)
}

func (ast *Assert) oneLenAssert(fatal bool, value interface {}, logs ...interface {}) {
	v:= reflect.Indirect(reflect.ValueOf(value))
	if v.Len() != 1 {
		ast.logCaller()
		if len(logs) > 0 {
			ast.Log(logs...)
		} else {
			ast.Log("Length is not one:", v.Len())
		}
		ast.failIt(fatal)
	}
}


func (ast *Assert) logCaller(){
	_, file, line, _ := runtime.Caller(3)
	ast.Logf("Caller: %v:%d", file, line)
}

func (ast *Assert) failIt(fatal bool){
	if fatal {
		ast.FailNow()
	}else {
		ast.Fail()
	}
}

func normalizeValue(value interface {}) interface {}{
	val := reflect.ValueOf(value)
	switch val.Kind(){
	case reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64:
		return int64(val.Uint())
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		return val.Int()
	case reflect.Float32,reflect.Float64:
		return val.Float()
	case reflect.Complex64,reflect.Complex128:
		return val.Complex()
	case reflect.String:
		return val.String()
	case reflect.Bool:
		return val.Bool()
	case reflect.Slice:
		if val.Type().Elem().Kind() == reflect.Uint8 {
			return val.Bytes()
		}
	}
	return value
}


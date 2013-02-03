Assert library for Go.
===========================

Usage
============================


    func TestSomthing(t *testing.T) {
      //Creating Assert object by passing *testing.T to NewAssert function
    	//Then *testing.T will be an anonymous pointer field of the Assert object
    	
    	assert := assrt.NewAssert(t)

    	//All method's of *testing.T will be included.
    	assert.Log("testing started")

    	err := someMethodReturnsError()
    	//assert err is nil, or fail the test
    	assert.Nil(err)

    	err := anotherError()
    	//assert methods began with "Must" will call t.FailNow() and quit testing
    	//if assertion failed
    	assert.MustNotNil(err)

    	

    	//Equal may convert the value to it's underlying value and cast the value to the largest container before compare.
    	//All int kind value like "int, int32, int 64, uin8, uit16, uint64" will convert to int64
    	//float32 type will convert to float64
    	//It can also compare struct and slice by calling "reflect.DeepEqual()" internally.

    	type A int16
		type B uint32
		var a A = 1
		var b B = 1
    	
    	assert.Equal(a, b) // assertion pass


    	slice := []string{"a","b","c"}

    	//you can pass optional log messages in assert methods, if not provided, assert will log
    	//default messages.
    	assert.PositiveLen(slice, "slice should be positive")
    }

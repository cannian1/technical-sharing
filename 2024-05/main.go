package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	left := 2
	right := 10
	fmt.Println(left, right, (right-left)>>1, (right-left)>>1+left)
	// 2 10 4 6
}

// https://github.com/casbin/gorm-adapter/blob/master/adapter.go#L132
// 鉴赏一下这相当恶心的代码
func badExample1(driverName string, dataSourceName string, params ...interface{}) (*Adapter, error) {
	a := &Adapter{}

	// ...

	if len(params) == 1 {
		switch p1 := params[0].(type) {
		case bool:
			a.dbSpecified = p1
		case string:
			a.databaseName = p1
		default:
			return nil, errors.New("wrong format")
		}
	} else if len(params) == 2 {
		switch p2 := params[1].(type) {
		case bool:
			a.dbSpecified = p2
			p1, ok := params[0].(string)
			if !ok {
				return nil, errors.New("wrong format")
			}
			a.databaseName = p1
		case string:
			p1, ok := params[0].(string)
			if !ok {
				return nil, errors.New("wrong format")
			}
			a.databaseName = p1
			a.tableName = p2
		default:
			return nil, errors.New("wrong format")
		}
	} else if len(params) == 3 {
		if p3, ok := params[2].(bool); ok {
			a.dbSpecified = p3
			a.databaseName = params[0].(string)
			a.tableName = params[1].(string)
		} else {
			return nil, errors.New("wrong format")
		}
	} else if len(params) != 0 {
		return nil, errors.New("too many parameters")
	}

	// ...

	return a, nil
}

type Adapter struct {
	driverName     string
	dataSourceName string
	databaseName   string
	tablePrefix    string
	tableName      string
	dbSpecified    bool
	//db             *gorm.DB
	isFiltered    bool
	transactionMu *sync.Mutex
	muInitialize  sync.Once
}

func MoreTerribleExample(flag []int) {
	a := len(flag)
	b := cap(flag)

	if a != b {
		if b-a >= 1 {
			// do some logic
		} else if b-a >= 2 {
			// do some logic
		}
	} else {
		// do some logic
	}
}

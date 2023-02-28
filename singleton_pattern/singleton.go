package singleton_pattern

import (
	"fmt"
	"sync"
)

type single struct {
}

var lock = &sync.Mutex{}
var singleInstance *single

func GetSingletonInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance
}

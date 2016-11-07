package utils

import "fmt"

//CheckErr to check error
func CheckErr(er error) {
	if er != nil {

		fmt.Print(fmt.Errorf("Error: while opening connection %v", er.Error()))
	}

	return
}

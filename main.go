package main

import (
	"facesdk"
	"fmt"
)

func main(){
	groupId := "dahuitang"
/*
	personIds,err := facesdk.FaceGetPersons(groupId)

	if err == nil {
		for _, v := range personIds {
			fmt.Println(v)
		}	
	} else {
		fmt.Println(err)
	}
*/
	groupIds := []string{groupId}
	err := facesdk.FaceAddPerson("abc1", groupIds)
	if err != nil {
		fmt.Println(err)
	}
	
	//url := "https://dahuitang.oss-cn-beijing.aliyuncs.com/yz1.jpg"
	
	//facesdk.FaceScan(url, groupId)
}


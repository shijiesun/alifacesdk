package main

import (
	"facesdk"
	"fmt"
)

func main(){
	groupId := "dahuitang"
	//groupIds := []string{groupId}

	personIds,err := facesdk.FaceGetPersons(groupId)

	if err == nil {
		for _, v := range personIds {
			fmt.Println(v)
		}	
	} else {
		fmt.Println(err)
	}
	
	//facesdk.FaceAddPerson("abc", groupIds)

	//url := "https://dahuitang.oss-cn-beijing.aliyuncs.com/yz1.jpg"
	
	//facesdk.FaceScan(url, groupId)
}


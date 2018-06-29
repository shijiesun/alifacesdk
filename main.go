package main

import (
	"facesdk"
	"fmt"
)

func main(){
	groupId := "dahuitang"
	//urls := []string{"https://dahuitang.oss-cn-beijing.aliyuncs.com/yz.jpg", "https://dahuitang.oss-cn-beijing.aliyuncs.com/yz1.jpg"}
	
/*
	personIds,err := facesdk.FaceGetPersons(groupId)

	if err == nil {
		for _, v := range personIds {
			fmt.Println(v)
		}	
	} else {
		fmt.Println(err)
	}

        groupIds := []string{groupId}
	err := facesdk.FaceAddPerson("abc1", groupIds)
	if err != nil {
		fmt.Println(err)
	}


	addFaceResponse, err := facesdk.FaceAddFace("abc1", urls)
	if err != nil {
		fmt.Println(err)
	} else {
		for _,v := range addFaceResponse {
			if v.Success {
				fmt.Println(v.FaceId)
			} else {
				fmt.Println(v.Msg)
			}
		}
	}

	var person *(facesdk.Person)
	person, err := facesdk.FaceGetPerson("abc1")
	if err != nil {
		fmt.Println(err)
	} else {
		p := *person
		fmt.Println(p.Code)
	}
*/
	url := "https://dahuitang.oss-cn-beijing.aliyuncs.com/yz1.jpg"
	
	facesdk.FaceScan(url, groupId)

	fmt.Println()
}


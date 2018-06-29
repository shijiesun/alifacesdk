package main

import (
	"facesdk"
	"fmt"
)

func main(){
	groupId := "dahuitang"
	//urls := []string{"https://dahuitang.oss-cn-beijing.aliyuncs.com/yz.jpg", "https://dahuitang.oss-cn-beijing.aliyuncs.com/yz1.jpg"}

	//urls := []string{"https://dahuitang.oss-cn-beijing.aliyuncs.com/garnet.png"}

	//personId := "garnet"
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
	err := facesdk.FaceAddPerson(personId, groupIds)
	if err != nil {
		fmt.Println(err)
	}


	addFaceResponse, err := facesdk.FaceAddFace(personId, urls)
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
		fmt.Printf("face count is %d\n", len(p.FaceIds))
	}


	faceIds1 := []string{"32516049499153440","32509073400612867","32513118561461275fdsfsd"}

	err := facesdk.FaceDeleteFace(faceIds1, personId)
	if err != nil {
		fmt.Println(err)

	err = facesdk.FaceDeletePerson(personId)
	if err != nil {
		fmt.Println(err)
	}
*/
	url := "https://dahuitang.oss-cn-beijing.aliyuncs.com/nba.jpg"

	var topPersons *[]facesdk.TopPerson
	
	topPersons, err := facesdk.FaceScan(url, groupId)

	if err != nil {
		fmt.Println(err)
	} else {
		for _,v := range *topPersons {
			fmt.Printf("%d, %d, %d, %d\n", v.FaceItem.X, v.FaceItem.Y, v.FaceItem.Width, v.FaceItem.Height)
			for _,v1 := range v.Persons {
				fmt.Printf("%s, %s, %f\n", v1.FaceId, v1.PersonId, v1.Rate)
			}
		}
	}

}


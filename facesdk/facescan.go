package facesdk

import (
	"encoding/json"
	"fmt"
	"uuid"
	"greensdksample"
)

const accessKeyId string = "LTAIEqr3BDEtynue"
const accessKeySecret string = "WXdTfFbwWDrRsykZfv7NSjadmWr38m"

func FaceGetPersons(groupId string) ([]string, error) {
	profile := greensdksample.Profile{AccessKeyId:accessKeyId, AccessKeySecret:accessKeySecret}

	path := "/green/sface/group/persons"
	
	clientInfo := greensdksample.ClinetInfo{Ip:"127.0.0.1"}

	req := GetPersonsRequest{GroupId:groupId}
	reqjson, _ := json.Marshal(req)
	
	var client greensdksample.IAliYunClient = greensdksample.DefaultClient{Profile:profile}

	// your biz code
	resultJson := client.GetResponse(path, clientInfo, string(reqjson))

	fmt.Println(resultJson)
	
	var result GetPersonsResponse
	
	if err := json.Unmarshal([]byte(resultJson), &result); err != nil {
		return nil, err
	}

	fmt.Println(result.Code)

	if result.Code == 200 {
		personIds := result.Data.PersonIds;

		return personIds, nil
	} else {
		return nil,nil
	}
	
	

}

func FaceAddPerson(personId string, groupIds []string) {
	profile := greensdksample.Profile{AccessKeyId:accessKeyId, AccessKeySecret:accessKeySecret}

	path := "/green/sface/person/add"
	clientInfo := greensdksample.ClinetInfo{Ip:"127.0.0.1"}

	req := AddPersonRequest{personId, groupIds}
	reqjson, _ := json.Marshal(req)
	
	var client greensdksample.IAliYunClient = greensdksample.DefaultClient{Profile:profile}

	// your biz code
	fmt.Println(client.GetResponse(path, clientInfo, string(reqjson)))
}


func FaceScan(url, groupId string) {
	profile := greensdksample.Profile{AccessKeyId:accessKeyId, AccessKeySecret:accessKeySecret}

	path := "/green/image/scan"
	
	clientInfo := greensdksample.ClinetInfo{Ip:"127.0.0.1"}

	// 构造请求数据
	scenes := []string{"sface-n"}

	extras := map[string]string {
		"groupId": groupId,
	}
	
	task := greensdksample.Task{DataId:uuid.Rand().Hex(), Url:url, Extras:extras}
	tasks := []greensdksample.Task{task}

	bizData := greensdksample.BizData{ scenes, tasks}
	bizDataJson, _ := json.Marshal(bizData)
	
	var client greensdksample.IAliYunClient = greensdksample.DefaultClient{Profile:profile}

	// your biz code
	fmt.Println(client.GetResponse(path, clientInfo, string(bizDataJson)))

}

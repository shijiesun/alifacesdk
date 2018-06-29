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

	var rawdata json.RawMessage
	result  := AbstractResponse{
		Data : &rawdata,
	}
	if err := json.Unmarshal([]byte(resultJson), &result); err != nil {
		return nil, err
	}

	if result.Code == 200 {

		fmt.Println(string(rawdata))
		
		var data GetPersonsResponse

		if err := json.Unmarshal([]byte(rawdata), &data); err != nil {
			return nil, err
		}

		personIds := data.PersonIds;

		return personIds, nil
	} else {
		return nil,fmt.Errorf("%s", result.Msg)
	}
}

func FaceGetPerson(personId string) (*Person, error) {
	profile := greensdksample.Profile{AccessKeyId:accessKeyId, AccessKeySecret:accessKeySecret}

	path := "/green/sface/person"
	
	clientInfo := greensdksample.ClinetInfo{Ip:"127.0.0.1"}

	req := GetPersonRequest{PersonId:personId}
	reqjson, _ := json.Marshal(req)
	
	var client greensdksample.IAliYunClient = greensdksample.DefaultClient{Profile:profile}

	// your biz code
	resultJson := client.GetResponse(path, clientInfo, string(reqjson))

	//fmt.Println(resultJson)

	var rawdata json.RawMessage
	result  := AbstractResponse{
		Data : &rawdata,
	}
	if err := json.Unmarshal([]byte(resultJson), &result); err != nil {
		return nil, err
	}

	if result.Code == 200 {

		fmt.Println(string(rawdata))
		
		var person Person

		if err := json.Unmarshal([]byte(rawdata), &person); err != nil {
			return nil, err
		}

		return &person, nil
	} else {
		return nil,nil
	}
}
 
func FaceAddPerson(personId string, groupIds []string) error {
	profile := greensdksample.Profile{AccessKeyId:accessKeyId, AccessKeySecret:accessKeySecret}

	path := "/green/sface/person/add"
	clientInfo := greensdksample.ClinetInfo{Ip:"127.0.0.1"}

	req := AddPersonRequest{personId, groupIds}
	reqjson, _ := json.Marshal(req)
	
	var client greensdksample.IAliYunClient = greensdksample.DefaultClient{Profile:profile}

	// your biz code
	resultJson := client.GetResponse(path, clientInfo, string(reqjson))

	fmt.Println(resultJson)

	var result AbstractResponse	
	if err := json.Unmarshal([]byte(resultJson), &result); err != nil {
		return err
	}

	if result.Code == 200 {
		return nil
	} else {
		return fmt.Errorf("%s", result.Msg)
	}
}

func FaceDeletePerson(personId string) error {
	profile := greensdksample.Profile{AccessKeyId:accessKeyId, AccessKeySecret:accessKeySecret}

	path := "/green/sface/person/delete"

	clientInfo := greensdksample.ClinetInfo{Ip:"127.0.0.1"}

	req := DeletePersonRequest{personId}
	reqjson, _ := json.Marshal(req)
	
	var client greensdksample.IAliYunClient = greensdksample.DefaultClient{Profile:profile}

	// your biz code
	resultJson := client.GetResponse(path, clientInfo, string(reqjson))

	fmt.Println(resultJson)

	var rawdata json.RawMessage
	result  := AbstractResponse{
		Data : &rawdata,
	}
	if err := json.Unmarshal([]byte(resultJson), &result); err != nil {
		return err
	}

	if result.Code == 200 {
		fmt.Println(string(rawdata))

		var data DeletePersonResponse
		if err := json.Unmarshal([]byte(rawdata), &data); err != nil {
			return err
		}

		if data.Code == 200 {
			return nil
		} else {
			return fmt.Errorf("%d", data.Code)
		}
	} else {
		return fmt.Errorf("%s", result.Msg)
	}
}

func FaceAddFace(personId string, urls []string) ([]AddFaceResponse, error) {
	profile := greensdksample.Profile{AccessKeyId:accessKeyId, AccessKeySecret:accessKeySecret}

	path := "/green/sface/face/add"
	clientInfo := greensdksample.ClinetInfo{Ip:"127.0.0.1"}

	req := AddFaceRequest{personId, urls}
	reqjson, _ := json.Marshal(req)
	
	var client greensdksample.IAliYunClient = greensdksample.DefaultClient{Profile:profile}

	// your biz code
	resultJson := client.GetResponse(path, clientInfo, string(reqjson))

	fmt.Println(resultJson)

	var rawdata json.RawMessage
	result  := AbstractResponse{
		Data : &rawdata,
	}
	if err := json.Unmarshal([]byte(resultJson), &result); err != nil {
		return nil, err
	}

	if result.Code == 200 {
		fmt.Println(string(rawdata))

		var data AddFaceData

		if err := json.Unmarshal([]byte(rawdata), &data); err != nil {
			return nil, err
		}

		return data.FaceImageItems, nil
		
	} else {
		return nil, fmt.Errorf("%s", result.Msg)
	}
}

//如果personid找不到，则返回error 500，否则返回空，即使没有找到faceId也不会报错
func FaceDeleteFace(faceIds []string, personId string) error {
	profile := greensdksample.Profile{AccessKeyId:accessKeyId, AccessKeySecret:accessKeySecret}

	path := "/green/sface/face/delete"
	clientInfo := greensdksample.ClinetInfo{Ip:"127.0.0.1"}

	req := DeleteFaceRequest{personId, faceIds}
	reqjson, _ := json.Marshal(req)
	
	var client greensdksample.IAliYunClient = greensdksample.DefaultClient{Profile:profile}

	// your biz code
	resultJson := client.GetResponse(path, clientInfo, string(reqjson))

	fmt.Println(resultJson)

	var rawdata json.RawMessage
	result  := AbstractResponse{
		Data : &rawdata,
	}
	if err := json.Unmarshal([]byte(resultJson), &result); err != nil {
		return err
	}

	if result.Code == 200 {
		fmt.Println(string(rawdata))

		var data DeleteFaceResponse
		if err := json.Unmarshal([]byte(rawdata), &data); err != nil {
			return err
		}

		if data.Code == 200 {
			return nil
		} else {
			return fmt.Errorf("%d", data.Code)
		}
	} else {
		return fmt.Errorf("%d", result.Code)
	}
}

func FaceScan(url, groupId string) (*[]TopPerson, error) {
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

	resultJson := client.GetResponse(path, clientInfo, string(bizDataJson))

	fmt.Println(resultJson)
	
	var scanResponse ScanResponse

	if err := json.Unmarshal([]byte(resultJson), &scanResponse); err != nil {
		return nil, err
	}

	scanData := scanResponse.Data[0]
	
	if scanData.Code != 200 {
		return nil, fmt.Errorf("%s", scanData.Msg)
	} else {
		result := scanData.Results[0]
		return &result.TopPersonData, nil
	}
}

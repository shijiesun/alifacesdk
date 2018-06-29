package facesdk

type AbstractResponse struct {
	Code int                `json:"code"`
	Data interface{}     
	Msg string              `json:"msg"`
}

type GetPersonsRequest struct {
	GroupId string          `json:"groupId"`
}

type GetPersonsResponse struct {
	PersonIds []string      `json:"personIds"`
}

type AddPersonRequest struct {
	PersonId string         `json:"personId"`
	GroupIds []string       `json:"groupIds"`
}

type AddFaceRequest struct {
	PersonId string         `json:"personId"`
	Urls []string           `json:"urls"`
}

type AddFaceResponse struct {
	FaceId int               `json:"faceId"`
	Success bool             `json:"success"`
	Msg string               `json:"msg"`
}

type AddFaceData struct {
	Code int                 `json:"code"`
	FaceImageItems []AddFaceResponse `json:"faceImageItems"`
}

type GetPersonRequest struct {
	PersonId string          `json:"personId"`
}

type Person struct {
	Code int                  `json:"code"`
	PersonId string           `json:"personId"`
	Name string               `json:"name"`
	groupIds []string         `json:"groupIds"`
	faceIds []string          `json:"faceIds"`
	Note string               `json:"note"`
}

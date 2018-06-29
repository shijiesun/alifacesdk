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
	FaceIds []string          `json:"faceIds"`
	Note string               `json:"note"`
}
type DeleteFaceRequest struct {
	PersonId string            `json:"personId"`
	FaceIds []string           `json:"faceIds"`
}

type DeleteFaceResponse struct {
	Code int                   `json:"code"`
	PersonId string            `json:"personId"`
	FaceIds []string           `json:"faceIds"`
	DeleteCount int            `json:"ddeleteCount"`
}

type DeletePersonRequest struct {
	PersonId string            `json:"personId"`
}

type DeletePersonResponse struct {
	Code int                   `json:"code"`
	PersonId string            `json:"personId"`
	DeleteCount int            `json:"ddeleteCount"`
}

type FaceLoc struct {
	Height int                `json:"height"`
	Width int                 `json:"width"`
	X int                     `json:"x"`
	Y int                     `json:"y"`
}

type PersonFound struct {
	FaceId string             `json:"faceId"`
	PersonId string           `json:"personId"`
	Rate float64              `json:"rate"`
}

type TopPerson struct {
	FaceItem FaceLoc          `json:"faceItem"`
	Persons []PersonFound     `json:"persons"`
}

type ScanResults struct {
	Rate float64              `json:"rate"`
	TopPersonData []TopPerson `json:"topPersonData"`
}

type ScanData struct {
	Code int                  `json:"code"`
	Results []ScanResults     `json:"results"`
	Msg string                `json:msg`
}
type ScanResponse struct {
	Code int                  `json:"code"`
	Data []ScanData           `json:"data"`
	Msg string                `json:msg`
}

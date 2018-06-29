package facesdk

type GetPersonsRequest struct {
	GroupId string          `json:"groupId"`
}

type GetPersonsData struct {
	PersonIds []string      `json:"personIds"`
}

type AbstractResponse struct {
	Code int                `json:"code"`
	Data interface{}     
	Msg string              `json:"msg"`
}

type AddPersonRequest struct {
	PersonId string         `json:"personId"`
	GroupIds []string       `json:"groupIds"`
}

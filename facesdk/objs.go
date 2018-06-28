package facesdk;

type GetPersonsRequest struct {
	GroupId string          `json:"groupId"`
}

type GetPersonsData struct {
	PersonIds []string      `json:"personIds"`
}

type GetPersonsResponse struct {
	Code int                `json:"code"`
        Data GetPersonsData     `json:"data"`
	Msg string              `json:"msg"`
}

type AddPersonRequest struct {
	PersonId string         `json:"personId"`
	GroupIds []string       `json:"groupIds"`
}

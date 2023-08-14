package model

type Pix2TextModel struct {
	StatusCode int    `json:"status_code"`
	SessionId  string `json:"session_id"`
	CallId     string `json:"call_id"`
	Message    string `json:"message"`
	Results    string `json:"results"`
}

type SimpletexModel struct {
	Status bool `json:"status"`
	Res    struct {
		Conf  float64 `json:"conf"`
		Latex string  `json:"latex"`
	} `json:"res"`
	RequestId string `json:"request_id"`
}

package model

type RequestSocket struct {
	Token string `json:"token"`
}

type TestSocket struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

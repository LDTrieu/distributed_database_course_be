package api

type pingDBResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Payload ping_db_resp `json:"payload"`
}

type ping_db_resp struct {
	// DB name
	DBName string `json:"dbName"`
	// list table (string)
	ListTable string `json:"listTable"`
}

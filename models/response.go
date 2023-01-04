package models

type Response struct {
	/*returning status
	200 = okay
	404 = not found
	500 = error
	*/
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

package model

type ResponseMeta struct {
	AppStatusCode int    `json:"code"`
	Message       string `json:"statusType,omitempty"`
	ErrorDetail   string `json:"errorDetail,omitempty"`
	ErrorMessage  string `json:"errorMessage,omitempty"`
	DevMessage    string `json:"devErrorMessage,omitempty"`
}

type ErrResponse struct {
	HTTPStatusCode int          `json:"-"` // http response status code
	Status         ResponseMeta `json:"status"`
	AppCode        int64        `json:"code,omitempty"` // application-specific error code
}

type Printers struct {
	ID                int    `json:"id"`
	PrinterName       string `json:"printer_name"`
	PrinterModel      string `json:"printer_model"`
	NetworkBased      bool   `json:"network_based"`
	Address           string `json:"address"`
	AuthenticationKey string `json:"authentication_key"`
}

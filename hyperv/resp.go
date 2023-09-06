package hyperv

import "encoding/json"

type response struct {
	Result  string      `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func returnResponse(respData interface{}, status int, result, message string) (int, string, []byte) {
	if value, ok := respData.(string); ok {
		respData, _ = json.Marshal(value)
	} else if value, ok := respData.([]byte); ok {
		respData = json.RawMessage(value)
	}

	resp := response{
		Result:  result,
		Message: message,
		Data:    respData,
	}
	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	return status, "application/json", jsonResp
}

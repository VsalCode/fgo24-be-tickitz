package utils;

type Response struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Errors   any    `json:"errors,omitempty"`
	PageInfo any    `json:"pageInfo,omitempty"`
	Results  any    `json:"results,omitempty"`
}
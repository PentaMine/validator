package dto

type ExecuteCodeRequest struct {
	Code string `json:"code" validate:"required"`
}

type ExecuteCodeResponse struct {
	StdOut string `json:"std_out" validate:"required"`
}

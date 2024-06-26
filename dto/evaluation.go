package dto

type ValidateCodeRequest struct {
	Location string `json:"location" validate:"required"`
	Code     string `json:"code" validate:"required"`
}

type ClusterResponse struct {
	MaxPoints int
	Points    int
}

type ValidateCodeResponse struct {
	MaxPoints int
	Points    int
	Clusters  []ClusterResponse
}

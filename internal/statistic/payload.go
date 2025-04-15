package statistic

type GetResponse struct {
	Period string `json:"period"`
	Sum    int    `json:"sum"`
}

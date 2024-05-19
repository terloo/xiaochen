package juhe

type BaseBody struct {
	Reason    string `json:"reason"`
	ErrorCode int    `json:"error_code"`
}

type ZhouGong struct {
	BaseBody `json:",inline"`
	Result   []ZhouGongResult `json:"result"`
}

type ZhouGongResult struct {
	Id    string   `json:"id"`
	Title string   `json:"title"`
	Des   string   `json:"des"`
	List  []string `json:"list"`
}

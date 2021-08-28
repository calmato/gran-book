package response

type ErrorResponse struct {
	Status  int64  `json:"status"`  // ステータスコード
	Code    int64  `json:"code"`    // エラーコード
	Message string `json:"message"` // エラー概要
	Detail  string `json:"detail"`  // エラー詳細
}

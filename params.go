package naverbandgo

// NextParams 페이징 요청
type NextParams struct {
	After string `json:"after"`
	Limit string `json:"limit"`
	BandKey string `json:"band_key"`
	AccessToken string `json:"access_token"`
}

// Apply 페이징 요청 적용
func (params *NextParams) Apply(data map[string]interface{}) map[string]interface{} {
	if params != nil {
		data["after"] = params.After
		data["limit"] = params.Limit
		data["band_key"] = params.BandKey
	}
	return data
}
package naverbandgo

// Band 밴드 정보 구조체
type Band struct {
	Name string `json:"name"`
	BandKey string `json:"band_key"`
	Cover string `json:"cover"`
	MemberCount uint `json:"member_count"`
}

// GetBandsResponse 밴드 목록 조회 응답
type GetBandsResponse struct {
	Bands []*Band `json:"bands"`
}

// GetBands 밴드 목록 조회
func (client *Client) GetBands() (bands []*Band, err error) {
	result, err := client.CallAPI("GET", "/v2.1/bands", nil)
	if err != nil {
		return nil, err
	}
	resp := new(GetBandsResponse)
	err = result.To(resp)
	if err != nil {
		return nil, err
	}
	return resp.Bands, nil
}
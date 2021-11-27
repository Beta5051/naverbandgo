package naverbandgo

// Profile 사용자 정보
type Profile struct {
	UserKey string `json:"user_key"`
	ProfileImageUrl string `json:"profile_image_url"`
	Name string `json:"name"`
	IsAppMember bool `json:"is_app_member"`
	MessageAllowed bool `json:"message_allowed"`
}

// Author 작성자 정보
type Author struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Role string `json:"role"`
	ProfileImageUrl string `json:"profile_image_url"`
	UserKey string `json:"user_key"`
}

// GetProfile 사용자 정보 조회
func (client *Client) GetProfile(bandKey string) (*Profile, error) {
	data := make(map[string]interface{})
	if bandKey != "" {
		data["band_key"] = bandKey
	}
	result, err := client.CallAPI("GET", "/v2/profile", data)
	if err != nil {
		return nil, err
	}
	profile := new(Profile)
	err = result.To(profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
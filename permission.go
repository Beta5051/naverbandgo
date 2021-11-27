package naverbandgo

// CheckPermissionsResponse 권한 조회 응답
type CheckPermissionsResponse struct {
	Permissions []string `json:"permissions"`
}

// CheckPermissions 권한 조회
func (client *Client) CheckPermissions(bandKey string, permissions []string) ([]string, error) {
	var permissionsStr string
	for idx, permission := range permissions {
		permissionsStr += permission
		if idx != len(permissions) - 1 {
			permissionsStr += ","
		}
	}
	result, err := client.CallAPI("GET", "/v2/band/permissions", map[string]interface{}{
		"band_key": bandKey,
		"permissions": permissionsStr,
	})
	if err != nil {
		return nil, err
	}
	resp := new(CheckPermissionsResponse)
	err = result.To(resp)
	if err != nil {
		return nil, err
	}
	return resp.Permissions, nil
}
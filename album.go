package naverbandgo

// Album 앨범 정보
type Album struct {
	PhotoAlbumKey string `json:"photo_album_key"`
	Name          string `json:"name"`
	PhotoCount    uint   `json:"photo_count"`
	CreateAt      uint   `json:"create_at"`
	Author        Author `json:"author"`
}

// GetAlbumsResponse 앨범 목록 조회 응답
type GetAlbumsResponse struct {
	Paging struct {
		NextParams *NextParams `json:"next_params"`
	} `json:"paging"`
	Items []*Album
}

// GetAlbums 앨범 목록 조회
func (client *Client) GetAlbums(bandKey string) ([]*Album, *NextParams, error) {
	result, err := client.CallAPI("GET", "/v2/band/albums", map[string]interface{}{
		"band_key": bandKey,
	})
	if err != nil {
		return nil, nil, err
	}
	resp := new(GetAlbumsResponse)
	err = result.To(resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Items, resp.Paging.NextParams, nil
}

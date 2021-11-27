package naverbandgo

// Photo 사진 정보
type Photo struct {
	Height uint `json:"height"`
	Width uint `json:"width"`
	CreatedAt uint `json:"created_at"`
	Url string `json:"url"`
	Author *Author `json:"author"`
	PhotoAlbumKey string `json:"photo_album_key"`
	PhotoKey string `json:"photo_key"`
	CommentCount uint `json:"comment_count"`
	EmoticonCount uint `json:"emoticon_count"`
	IsVideoThumbnail bool `json:"is_video_thumbnail"`
}

// GetPhotosResponse 사진 목록 조회 응답
type GetPhotosResponse struct {
	Paging struct {
		NextParams *NextParams `json:"next_params"`
	} `json:"paging"`
	Items []*Photo `json:"items"`
}

// GetPhotos 사진 목록 조회
func (client *Client) GetPhotos(bandKey string, photoAlbumKey string) ([]*Photo, *NextParams, error) {
	result, err := client.CallAPI("GET", "/v2/band/album/photos", map[string]interface{}{
		"band_key": bandKey,
		"photo_album_key": photoAlbumKey,
	})
	if err != nil {
		return nil, nil, err
	}
	resp := new(GetPhotosResponse)
	err = result.To(resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Items, resp.Paging.NextParams, nil
}

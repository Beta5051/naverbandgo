package naverbandgo

// Comment 댓글 정보
type Comment struct {
	BandKey string `json:"band_key"`
	Author *Author `json:"author"`
	PostKey string `json:"post_key"`
	CommentKey string `json:"comment_key"`
	Content string `json:"content"`
	EmoticonCount uint `json:"emoticon_count"`
	IsAudioIncluded bool `json:"is_audio_included"`
	CreateAt uint `json:"create_at"`
	Photo *CommentPhoto `json:"photo"`
}

// CommentPhoto 댓글 사진 정보
type CommentPhoto struct {
	Url string `json:"url"`
	Height uint `json:"height"`
	Width uint `json:"width"`
}

// GetCommentsResponse 댓글 목록 조회 응답
type GetCommentsResponse struct {
	Paging struct {
		NextParams *NextParams `json:"next_params"`
	} `json:"paging"`
	Items []*Comment `json:"items"`
}

// GetComments 댓글 목록 조회
func (client *Client) GetComments(bandKey, postKey string, nextParams *NextParams, creationOrder bool) ([]*Comment, *NextParams, error) {
	data := map[string]interface{}{
		"band_key": bandKey,
		"post_key": postKey,
	}
	if !creationOrder {
		data["sort"] = "-created_at"
	}
	data = nextParams.Apply(data)
	result, err := client.CallAPI("GET", "/v2/band/post/comments", data)
	if err != nil {
		return nil, nil, err
	}
	resp := new(GetCommentsResponse)
	err = result.To(resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Items, resp.Paging.NextParams, nil
}

// CreateComment 댓글 생성
func (client *Client) CreateComment(bandKey, postKey, body string) error {
	_, err := client.CallAPI("POST", "/v2/band/post/comment/create", map[string]interface{}{
		"band_key": bandKey,
		"post_key": postKey,
		"body": body,
	})
	return err
}

// RemoveComment 댓글 삭제
func (client *Client) RemoveComment(bandKey, postKey, commentKey string) error {
	_, err := client.CallAPI("POST", "/v2/band/post/comment/remove", map[string]interface{}{
		"band_key": bandKey,
		"post_key": postKey,
		"comment_key": commentKey,
	})
	return err
}
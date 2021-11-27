package naverbandgo

// Post 글 정보
type Post struct {
	Content string `json:"content"`
	Author *Author `json:"author"`
	PostKey string `json:"post_key"`
	CommentCount uint `json:"comment_count"`
	CreatedAt uint `json:"created_at"`
	Photos []*Photo `json:"photos"`
	EmoticonCount uint `json:"emoticon_count"`
	LastComments []*LastComment `json:"last_comments"`
	BandKey string `json:"band_key"`
}

// LastComment 최근 댓글 정보
type LastComment struct {
	Body string `json:"body"`
	Author *Author `json:"author"`
	CreatedAt uint `json:"created_at"`
}

// PostMoreInfo 글 상세 정보
type PostMoreInfo struct {
	Content string `json:"content"`
	Author *Author `json:"author"`
	PostKey string `json:"post_key"`
	CommentCount uint `json:"comment_count"`
	CreatedAt uint `json:"created_at"`
	Photos []*Photo `json:"photos"`
	EmoticonCount uint `json:"emoticon_count"`
	BandKey string `json:"band_key"`
	PostReadCount uint `json:"post_read_count"`
}

// GetPostsResponse 글 목록 조회 응답
type GetPostsResponse struct {
	Paging struct {
		NextParams *NextParams `json:"next_params"`
	} `json:"paging"`
	Items []*Post `json:"items"`
}

func (client *Client) GetPosts(bandKey string, nextParams *NextParams, locale string) ([]*Post, *NextParams, error) {
	if locale != "" {
		locale = "ko_KR"
	}
	data := map[string]interface{}{
		"band_key": bandKey,
		"locale": locale,
	}
	data = nextParams.Apply(data)
	result, err := client.CallAPI("GET", "/v2/band/posts", data)
	if err != nil {
		return nil, nil, err
	}
	resp := new(GetPostsResponse)
	err = result.To(resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Items, resp.Paging.NextParams, nil
}

// GetPostResponse 글 상제 조회 응답
type GetPostResponse struct {
	Post *PostMoreInfo `json:"post"`
}

// GetPost 글 상세 조회
func (client *Client) GetPost(bandKey, postKey string) (*PostMoreInfo, error) {
	result, err := client.CallAPI("GET", "/v2.1/band/post", map[string]interface{}{
		"band_key": bandKey,
		"post_key": postKey,
	})
	if err != nil {
		return nil, err
	}
	resp := new(GetPostResponse)
	err = result.To(resp)
	if err != nil {
		return nil, err
	}
	return resp.Post, nil
}

// CreatePostResponse 글 생성 응답
type CreatePostResponse struct {
	BandKey string `json:"band_key"`
	PostKey string `json:"post_key"`
}

// CreatePost 글 생성
func (client *Client) CreatePost(bandKey, content string, doPush bool) (returnBandKey, postKey string, err error){
	result, err := client.CallAPI("POST", "/v2.2/band/post/create", map[string]interface{}{
		"band_key": bandKey,
		"content": content,
		"do_push": doPush,
	})
	if err != nil {
		return "", "", nil
	}
	resp := new(CreatePostResponse)
	err = result.To(resp)
	if err != nil {
		return "", "", err
	}
	return resp.BandKey, resp.PostKey, nil
}

// RemovePost 글 삭제
func (client *Client) RemovePost(bandKey, postKey string) error {
	_, err := client.CallAPI("POST", "/v2/band/post/remove", map[string]interface{}{
		"band_key": bandKey,
		"post_key": postKey,
	})
	return err
}
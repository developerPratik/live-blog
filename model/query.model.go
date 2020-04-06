package model

var operation = []string{
	"GET_POSTS",
	"GET_POST_BY_ID",
	"GET_POSTS_BY_USER",
	"GET_POSTS_BY_USER",
}


type Filter struct {
	PostId string `json:"postId"`
	CommentId string `json:"commentId"`
}

type Operation struct {
	QueryType string `json:"queryType"`
	QueryFilter Filter
}
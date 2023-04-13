package internal

import (
	"context"

	"github.com/syamilAbdillah/post-book/pkg/webrpc"
)

type PostService struct {
	dal PostDAL
}

type PostDAL interface {
	PostsInsert(ctx context.Context, post *webrpc.Post) error
	PostsFeed(ctx context.Context) ([]*webrpc.PostUser, error)
	PostsByUserID(ctx context.Context, id string) ([]*webrpc.PostUser, error)
	PostUpdate(ctx context.Context, post *webrpc.Post) error
}

func NewPostService() *PostService {
	return &PostService{}
}

func (s *PostService) CreatePost(ctx context.Context, content string) (map[string]string, error) {
	return nil, nil
}

func (s *PostService) UpdataPost(ctx context.Context, post *webrpc.Post) (map[string]string, error) {
	return nil, nil
}

func (s *PostService) PostsOfUser(ctx context.Context, userId string) (*webrpc.User, []*webrpc.Post, error) {
	return nil, nil, nil
}

func (s *PostService) POstsFeed(ctx context.Context) ([]*webrpc.PostUser, error) {
	return nil, nil
}
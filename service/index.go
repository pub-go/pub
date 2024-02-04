package service

import (
	"context"

	"code.gopub.tech/pub/dal/model"
	"code.gopub.tech/pub/dal/query"
	"code.gopub.tech/pub/dto"
	"github.com/youthlin/t"
	"github.com/youthlin/t/errors"
)

func QueryPosts(ctx context.Context, req *dto.QueryPostReq) (*dto.QueryPostResp, error) {
	req.PageSize = 10
	p := query.Post
	posts, count, err := p.WithContext(ctx).Where().FindByPage(req.Page*req.PageSize, req.PageSize)
	if err != nil {
		return nil, errors.Wrapf(err, t.WithContext(ctx).T("failed to query posts"))
	}
	return &dto.QueryPostResp{
		Req:       req,
		PostCount: count,
		// 3/page
		// totalCount=0, totalPage=0
		// totalCount=1, totalPage=1
		// totalCount=2, totalPage=1
		// totalCount=3, totalPage=1
		// totalCount=4, totalPage=2
		// totalCount=5, totalPage=2
		// totalCount=6, totalPage=2
		TotalPage: (count + int64(req.PageSize) - 1) / int64(req.PageSize),
		Count:     len(posts),
		Posts:     buildPostItems(ctx, posts),
	}, nil
}

func buildPostItems(ctx context.Context, posts []*model.Post) []*dto.PostItem {
	var result = make([]*dto.PostItem, 0, len(posts))
	for _, post := range posts {
		post := post
		result = append(result, &dto.PostItem{
			Post: post,
		})
	}
	return result

}

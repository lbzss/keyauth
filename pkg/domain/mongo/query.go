package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/token"
)

func newQueryDomainRequest(tk *token.Token, req *domain.QueryDomainRequest) *request {
	return &request{
		token:              tk,
		QueryDomainRequest: req,
	}
}

type request struct {
	token *token.Token
	*domain.QueryDomainRequest
}

func (r *request) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *request) FindFilter() bson.M {
	filter := bson.M{}

	filter["owner"] = r.token.Account
	return filter
}

func newDescDomainRequest(req *domain.DescribeDomainRequest) (*descDomain, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &descDomain{req}, nil
}

type descDomain struct {
	*domain.DescribeDomainRequest
}

func (req *descDomain) FindFilter() bson.M {
	filter := bson.M{}

	if req.Name != "" {
		filter["_id"] = req.Name
	}

	return filter
}

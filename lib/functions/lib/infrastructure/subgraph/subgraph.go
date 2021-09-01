package subgraph

import (
	"aurora-backend/lib/functions/lib/common/log"
	"context"
	"errors"

	"github.com/machinebox/graphql"
)

const query = `
	query ($key: String!) {
		postContent(id:$key) {
				metadata
		}
	}
`

type (

	// Client client
	Client struct {
		svc *graphql.Client
	}
	response struct {
		PostContent struct {
			Metadata string `json:"metadata"`
		}
	}
)

// New new client
func New() Client {
	svc := graphql.NewClient("https://api.studio.thegraph.com/query/2636/auroradao/v0.0.12")
	return Client{
		svc: svc,
	}
}

// MetadataURI get metadata uri
func (c Client) MetadataURI(ctx context.Context, postID string) (string, error) {
	req := graphql.NewRequest(query)
	req.Var("key", postID)
	req.Header.Set("Cache-Control", "no-cache")
	var res response
	if err := c.svc.Run(ctx, req, &res); err != nil {
		log.Error("query failed", err)
		return "", err
	}
	metadata := res.PostContent.Metadata
	if metadata == "" {
		msg := errors.New("post with id " + postID + " not found.")
		log.Error("", msg)
		return "", msg
	}
	return res.PostContent.Metadata, nil
}

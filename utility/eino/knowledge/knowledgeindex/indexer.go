package knowledgeindex

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/eino-ext/components/indexer/redis"
	"github.com/cloudwego/eino/components/indexer"
	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
	redisCli "github.com/redis/go-redis/v9"
	"os"
	redispkg "proxima/utility/eino/pkg/redis"
)

// newIndexer component initialization function of node 'Indexer1' in graph 'knowledge'
func newIndexer(ctx context.Context) (idr indexer.Indexer, err error) {
	// TODO Modify component configuration here.
	redisAddr := os.Getenv("http://116.62.42.164:6379")
	redisClient := redisCli.NewClient(&redisCli.Options{
		Addr:     redisAddr,
		Protocol: 5,
	})
	config := &redis.IndexerConfig{
		Client:    redisClient,
		BatchSize: 1,
		DocumentToHashes: func(ctx context.Context, doc *schema.Document) (*redis.Hashes, error) {
			if doc.ID == "" {
				doc.ID = uuid.New().String()
			}
			key := doc.ID
			marshal, err := json.Marshal(doc.MetaData)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal metadata:%w", err)
			}
			return &redis.Hashes{
				Key: key,
				Field2Value: map[string]redis.FieldValue{
					redispkg.ContentField:  {Value: doc.Content, EmbedKey: redispkg.VectorField},
					redispkg.MetadataField: {Value: marshal},
				},
			}, nil

		},
	}
	embeddingIns11, err := newEmbedding(ctx)
	if err != nil {
		return nil, err
	}
	config.Embedding = embeddingIns11
	idr, err = redis.NewIndexer(ctx, config)
	if err != nil {
		return nil, err
	}
	return idr, nil
}

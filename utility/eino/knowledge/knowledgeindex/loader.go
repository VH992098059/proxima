package knowledgeindex

import (
	"context"

	"github.com/cloudwego/eino-ext/components/document/loader/file"
	"github.com/cloudwego/eino/components/document"
)

// newLoader component initialization function of node 'Loader2' in graph 'knowledge'
func newLoader(ctx context.Context) (ldr document.Loader, err error) {
	// TODO Modify component configuration here.
	config := &file.FileLoaderConfig{}
	ldr, err = file.NewFileLoader(ctx, config)
	if err != nil {
		return nil, err
	}
	return ldr, nil
}

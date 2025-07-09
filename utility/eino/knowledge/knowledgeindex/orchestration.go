package knowledgeindex

import (
	"context"

	"github.com/cloudwego/eino/compose"
)

func Buildknowledge(ctx context.Context) (r compose.Runnable[any, any], err error) {
	const (
		Loader2              = "Loader2"
		DocumentTransformer2 = "DocumentTransformer2"
		Indexer1             = "Indexer1"
	)
	g := compose.NewGraph[any, any]()
	loader2KeyOfLoader, err := newLoader(ctx)
	if err != nil {
		return nil, err
	}
	_ = g.AddLoaderNode(Loader2, loader2KeyOfLoader)
	documentTransformer2KeyOfDocumentTransformer, err := newDocumentTransformer(ctx)
	if err != nil {
		return nil, err
	}
	_ = g.AddDocumentTransformerNode(DocumentTransformer2, documentTransformer2KeyOfDocumentTransformer)
	indexer1KeyOfIndexer, err := newIndexer(ctx)
	if err != nil {
		return nil, err
	}
	_ = g.AddIndexerNode(Indexer1, indexer1KeyOfIndexer)
	_ = g.AddEdge(compose.START, Loader2)
	_ = g.AddEdge(Indexer1, compose.END)
	_ = g.AddEdge(Loader2, DocumentTransformer2)
	_ = g.AddEdge(DocumentTransformer2, Indexer1)
	r, err = g.Compile(ctx, compose.WithGraphName("knowledge"))
	if err != nil {
		return nil, err
	}
	return r, err
}

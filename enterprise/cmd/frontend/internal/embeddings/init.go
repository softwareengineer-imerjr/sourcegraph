package embeddings

import (
	"context"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/enterprise"
	"github.com/sourcegraph/sourcegraph/enterprise/cmd/frontend/internal/embeddings/resolvers"
	"github.com/sourcegraph/sourcegraph/enterprise/internal/codeintel"
	"github.com/sourcegraph/sourcegraph/enterprise/internal/embeddings/background/contextdetection"
	"github.com/sourcegraph/sourcegraph/enterprise/internal/embeddings/background/repo"
	"github.com/sourcegraph/sourcegraph/internal/conf/conftypes"
	"github.com/sourcegraph/sourcegraph/internal/database"
	"github.com/sourcegraph/sourcegraph/internal/database/basestore"
	"github.com/sourcegraph/sourcegraph/internal/gitserver"
	"github.com/sourcegraph/sourcegraph/internal/observation"
)

func Init(
	ctx context.Context,
	observationCtx *observation.Context,
	db database.DB,
	_ codeintel.Services,
	_ conftypes.UnifiedWatchable,
	enterpriseServices *enterprise.Services,
) error {
	repoEmbeddingsStore := repo.RepoEmbeddingJobsStore{Store: basestore.NewWithHandle(db.Handle())}
	contextDetectionEmbeddingsStore := contextdetection.ContextDetectionEmbeddingJobsStore{Store: basestore.NewWithHandle(db.Handle())}
	gitserverClient := gitserver.NewClient()
	enterpriseServices.EmbeddingsResolver = resolvers.NewResolver(db, gitserverClient, repoEmbeddingsStore, contextDetectionEmbeddingsStore)
	return nil
}

package workers

import (
	"context"
	"testing"
	"time"

	"github.com/interline-io/transitland-jobs/jobs"
	"github.com/interline-io/transitland-server/internal/testconfig"
	"github.com/interline-io/transitland-server/model"
	"github.com/stretchr/testify/assert"
)

func TestFetchEnqueueWorker(t *testing.T) {
	ctx := context.Background()
	testconfig.ConfigTxRollback(t, testconfig.Options{}, func(cfg model.Config) {
		a := "BA"
		jobQueue := cfg.JobQueue
		jobQueue.Use(newCfgMiddleware(cfg))
		jobQueue.AddQueue("default", 8)
		jobQueue.AddJobType(func() jobs.JobWorker { return &FetchEnqueueWorker{} })
		jobQueue.AddJobType(func() jobs.JobWorker { return &StaticFetchWorker{} })
		jobQueue.AddJobType(func() jobs.JobWorker { return &GbfsFetchWorker{} })
		jobQueue.AddJobType(func() jobs.JobWorker { return &RTFetchWorker{} })
		go func() {
			jobQueue.Run(ctx)
		}()
		jobQueue.AddJob(ctx, jobs.Job{
			JobType: "fetch-enqueue",
			JobArgs: map[string]any{"feed_ids": []string{a}, "ignore_fetch_wait": true},
		})
		time.Sleep(2 * time.Second)

		// Check that we fetched from BART but failed
		ctx := model.WithConfig(context.Background(), cfg)
		feeds, err := cfg.Finder.FindFeeds(ctx, nil, nil, nil, &model.FeedFilter{OnestopID: &a})
		if err != nil {
			t.Fatal(err)
		}
		if len(feeds) != 1 {
			t.Fatal("expected one feed")
		}
		fetches, _ := cfg.Finder.FeedFetchesByFeedID(ctx, []model.FeedFetchParam{{FeedID: feeds[0].ID}})
		if len(fetches) == 0 {
			t.Error("expected at least one fetch")
		} else if len(fetches[0]) == 0 {
			t.Error("expected at least one fetch")
		} else {
			fetch := fetches[0][0]
			assert.Equal(t, false, fetch.Success)
			assert.Equal(t, "static_current", fetch.URLType)
			assert.Equal(t, "request not configured to allow filesystem access", fetch.FetchError.Val)
		}
	})
}

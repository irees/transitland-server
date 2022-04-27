package workers

import (
	"context"
	"errors"
	"fmt"

	"github.com/interline-io/transitland-lib/dmfr/fetch"
	"github.com/interline-io/transitland-lib/tldb"
	"github.com/interline-io/transitland-server/internal/jobs"
	"github.com/interline-io/transitland-server/model"
	"google.golang.org/protobuf/proto"
)

type RTFetchWorker struct {
	Target       string `json:"target"`
	Url          string `json:"url"`
	SourceType   string `json:"source_type"`
	SourceFeedID string `json:"source_feed_id"`
}

func (w *RTFetchWorker) Run(ctx context.Context, job jobs.Job) error {
	log := job.Opts.Logger.With().Str("target", w.Target).Str("source_feed_id", w.SourceFeedID).Str("source_type", w.SourceType).Str("url", w.Url).Logger()
	// Find feed
	rtfeeds, err := job.Opts.Finder.FindFeeds(nil, nil, nil, &model.FeedFilter{OnestopID: &w.SourceFeedID})
	if err != nil {
		log.Error().Err(err).Msg("rtfetch worker: error loading source feed")
		return err
	}
	if len(rtfeeds) == 0 {
		log.Error().Err(err).Msg("rtfetch worker: source feed not found")
		return errors.New("feed not found")
	}
	// Make request
	rtfeed := rtfeeds[0].Feed
	atx := tldb.NewPostgresAdapterFromDBX(job.Opts.Finder.DBX())
	fetchOpts := fetch.Options{
		Secrets: job.Opts.Secrets,
	}
	rtmsg, fr, err := fetch.RTFetch(atx, rtfeed, fetchOpts)
	if err != nil {
		log.Error().Err(err).Msg("rtfetch worker: request failed")
		return err
	}
	// Convert back to bytes...
	rtdata, err := proto.Marshal(rtmsg)
	if err != nil {
		return err
	}
	// Save to cache
	key := fmt.Sprintf("rtdata:%s:%s", w.Target, w.SourceType)
	log.Info().Int("bytes", fr.ResponseSize).Str("url", w.Url).Msg("rtfetch worker: success")
	return job.Opts.RTFinder.AddData(key, rtdata)
}

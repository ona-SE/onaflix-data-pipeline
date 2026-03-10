package processor

import (
	"testing"

	"github.com/ona-se/onaflix-data-pipeline/internal/models"
)

func TestEnrichMovie_EmptyTitle(t *testing.T) {
	p := &Processor{}
	movie := &models.Movie{ID: 1, Title: ""}

	err := p.enrichMovie(movie)
	if err == nil {
		t.Error("expected error for empty title")
	}
}

func TestEnrichMovie_ValidTitle(t *testing.T) {
	p := &Processor{}
	movie := &models.Movie{ID: 1, Title: "The Shawshank Redemption"}

	err := p.enrichMovie(movie)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPipelineStats(t *testing.T) {
	stats := &models.PipelineStats{
		Processed: 10,
		Enriched:  8,
		Errors:    2,
	}

	if stats.Processed != stats.Enriched+stats.Errors {
		t.Errorf("stats don't add up: %d != %d + %d",
			stats.Processed, stats.Enriched, stats.Errors)
	}
}

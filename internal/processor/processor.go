package processor

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/ona-se/onaflix-data-pipeline/internal/models"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Processor struct {
	db *sql.DB
}

func New(dbURL string) (*Processor, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Processor{db: db}, nil
}

func (p *Processor) Close() {
	if p.db != nil {
		p.db.Close()
	}
}

func (p *Processor) Run() (*models.PipelineStats, error) {
	stats := &models.PipelineStats{}

	movies, err := p.fetchMovies()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch movies: %w", err)
	}

	for _, movie := range movies {
		stats.Processed++

		if err := p.enrichMovie(&movie); err != nil {
			log.WithFields(log.Fields{
				"movie_id": movie.ID,
				"title":    movie.Title,
				"error":    err.Error(),
			}).Warn("Failed to enrich movie")
			stats.Errors++
			continue
		}

		stats.Enriched++
	}

	return stats, nil
}

func (p *Processor) fetchMovies() ([]models.Movie, error) {
	rows, err := p.db.Query(`
		SELECT id, title, COALESCE(description, ''), COALESCE(release_year, 0),
		       COALESCE(rating, 0), COALESCE(image_url, ''), COALESCE(director, '')
		FROM movies
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var m models.Movie
		if err := rows.Scan(&m.ID, &m.Title, &m.Description, &m.ReleaseYear,
			&m.Rating, &m.ImageURL, &m.Director); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, rows.Err()
}

func (p *Processor) enrichMovie(movie *models.Movie) error {
	// Normalize title for matching
	normalized := strings.TrimSpace(strings.ToLower(movie.Title))
	if normalized == "" {
		return fmt.Errorf("empty title for movie %d", movie.ID)
	}

	log.WithFields(log.Fields{
		"movie_id": movie.ID,
		"title":    movie.Title,
	}).Debug("Enriching movie")

	// In production this would call an external API
	// For demo, just log the enrichment attempt
	return nil
}

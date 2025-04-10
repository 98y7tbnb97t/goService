package auth

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type TokenCleaner struct {
	db *gorm.DB
}

func NewTokenCleaner(db *gorm.DB) *TokenCleaner {
	return &TokenCleaner{db: db}
}

func (tc *TokenCleaner) StartCleanup(ctx context.Context, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				tc.deleteExpiredTokens()
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (tc *TokenCleaner) deleteExpiredTokens() {
	tc.db.Exec("DELETE FROM refresh_tokens WHERE expires_at < NOW()")
}

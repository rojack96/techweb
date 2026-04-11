package session

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"streetcats-api/configs"

	"github.com/Nerzal/gocloak/v13"
	r "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const sessionKeyPrefix = "session:"

func generateSessionID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func derefString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

type Session struct {
	SessionID    string `json:"session_id"`
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
	IssuedAt     int64  `json:"issued_at"`
}

type ServiceInterfaces interface {
	Create(userInfo *gocloak.UserInfo, token *gocloak.JWT) (string, error)
	GetSessionByID(sessionID string) (*Session, error)
	Delete(sessionID string) error
}

type service struct {
	log   *zap.Logger
	cfg   *configs.ConfigModel
	kc    *gocloak.GoCloak
	redis *r.Client
	ctx   context.Context
}

func NewService(log *zap.Logger, cfg configs.ConfigModel, kc *gocloak.GoCloak, redisClient *r.Client) ServiceInterfaces {
	return &service{log: log, cfg: &cfg, kc: kc, redis: redisClient, ctx: context.Background()}
}

func (s *service) key(sessionID string) string {
	return fmt.Sprintf("%s%s", sessionKeyPrefix, sessionID)
}

func (s *service) Create(userInfo *gocloak.UserInfo, token *gocloak.JWT) (string, error) {
	if s.redis == nil {
		return "", fmt.Errorf("redis client not configured")
	}

	sessionID, err := generateSessionID()
	if err != nil {
		return "", err
	}

	ttl := token.ExpiresIn
	if ttl <= 0 {
		ttl = 3600
	}

	session := Session{
		SessionID:    sessionID,
		UserID:       derefString(userInfo.Sub),
		Username:     derefString(userInfo.PreferredUsername),
		Email:        derefString(userInfo.Email),
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    time.Now().Add(time.Duration(ttl) * time.Second).Unix(),
		IssuedAt:     time.Now().Unix(),
	}

	payload, err := json.Marshal(session)
	if err != nil {
		return "", err
	}

	if err = s.redis.Set(s.ctx, s.key(sessionID), payload, time.Duration(ttl)*time.Second).Err(); err != nil {
		return "", err
	}

	return sessionID, nil
}

func (s *service) GetSessionByID(sessionID string) (*Session, error) {
	if s.redis == nil {
		return nil, fmt.Errorf("redis client not configured")
	}

	payload, err := s.redis.Get(s.ctx, s.key(sessionID)).Result()
	if err != nil {
		if err == r.Nil {
			return nil, fmt.Errorf("session not found")
		}
		return nil, err
	}

	var session Session
	if err := json.Unmarshal([]byte(payload), &session); err != nil {
		return nil, err
	}

	if time.Now().Unix() > session.ExpiresAt {
		_ = s.Delete(sessionID)
		return nil, fmt.Errorf("session expired")
	}

	return &session, nil
}

func (s *service) Delete(sessionID string) error {
	if s.redis == nil {
		return fmt.Errorf("redis client not configured")
	}
	return s.redis.Del(s.ctx, s.key(sessionID)).Err()
}

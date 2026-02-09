package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// RedisTokenModel Redis Token管理模型
type RedisTokenModel struct {
	redis *redis.Redis
	logx.Logger
}

// NewRedisTokenModel 创建Redis Token模型
func NewRedisTokenModel(redis *redis.Redis) *RedisTokenModel {
	return &RedisTokenModel{
		redis:  redis,
		Logger: logx.WithContext(context.Background()),
	}
}

// TokenMetadata Token元数据
type TokenMetadata struct {
	UserID       int64                  `json:"user_id"`
	AccountID    int64                  `json:"account_id"`
	BusinessID   int64                  `json:"business_id"`
	Username     string                 `json:"username"`
	Role         string                 `json:"role"`
	Permissions  []string               `json:"permissions"`
	DeviceInfo   map[string]interface{} `json:"device_info"`
	IPAddress    string                 `json:"ip_address"`
	UserAgent    string                 `json:"user_agent"`
	CreatedAt    int64                  `json:"created_at"`
	ExpiresAt    int64                  `json:"expires_at"`
	LastUsedAt   int64                  `json:"last_used_at"`
	RefreshCount int                    `json:"refresh_count"`
	IsActive     bool                   `json:"is_active"`
	TokenType    string                 `json:"token_type"` // access_token, refresh_token
}

// TokenRefreshRecord Token刷新记录
type TokenRefreshRecord struct {
	OldToken     string `json:"old_token"`
	NewToken     string `json:"new_token"`
	UserID       int64  `json:"user_id"`
	RefreshTime  int64  `json:"refresh_time"`
	IPAddress    string `json:"ip_address"`
	UserAgent    string `json:"user_agent"`
	RefreshCount int    `json:"refresh_count"`
}

// TokenBlacklistEntry Token黑名单条目
type TokenBlacklistEntry struct {
	Token         string `json:"token"`
	UserID        int64  `json:"user_id"`
	BlacklistedAt int64  `json:"blacklisted_at"`
	Reason        string `json:"reason"`
	ExpiresAt     int64  `json:"expires_at"`
}

// ==================== Token存储相关 ====================

// SaveTokenMetadata 保存Token元数据
func (m *RedisTokenModel) SaveTokenMetadata(ctx context.Context, token string, metadata *TokenMetadata, ttl int) error {
	key := fmt.Sprintf("token:metadata:%s", token)

	data, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("序列化Token元数据失败: %v", err)
	}

	if err := m.redis.SetexCtx(ctx, key, string(data), ttl); err != nil {
		return fmt.Errorf("保存Token元数据失败: %v", err)
	}

	// 同时保存用户Token索引
	userTokensKey := fmt.Sprintf("user:tokens:%d", metadata.UserID)
	if _, err := m.redis.SaddCtx(ctx, userTokensKey, token); err != nil {
		return fmt.Errorf("保存用户Token索引失败: %v", err)
	}

	// 设置用户Token索引的TTL
	if err := m.redis.ExpireCtx(ctx, userTokensKey, ttl); err != nil {
		m.Logger.Errorf("设置用户Token索引TTL失败: %v", err)
	}

	return nil
}

// GetTokenMetadata 获取Token元数据
func (m *RedisTokenModel) GetTokenMetadata(ctx context.Context, token string) (*TokenMetadata, error) {
	key := fmt.Sprintf("token:metadata:%s", token)

	data, err := m.redis.GetCtx(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("获取Token元数据失败: %v", err)
	}

	if data == "" {
		return nil, fmt.Errorf("Token元数据不存在")
	}

	var metadata TokenMetadata
	if err := json.Unmarshal([]byte(data), &metadata); err != nil {
		return nil, fmt.Errorf("反序列化Token元数据失败: %v", err)
	}

	return &metadata, nil
}

// UpdateTokenLastUsed 更新Token最后使用时间
func (m *RedisTokenModel) UpdateTokenLastUsed(ctx context.Context, token string) error {
	key := fmt.Sprintf("token:metadata:%s", token)

	// 获取现有数据
	data, err := m.redis.GetCtx(ctx, key)
	if err != nil || data == "" {
		return fmt.Errorf("Token不存在")
	}

	var metadata TokenMetadata
	if err := json.Unmarshal([]byte(data), &metadata); err != nil {
		return fmt.Errorf("反序列化失败: %v", err)
	}

	// 更新最后使用时间
	metadata.LastUsedAt = time.Now().Unix()

	// 保存更新后的数据
	newData, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("序列化失败: %v", err)
	}

	// 获取剩余TTL
	ttl, err := m.redis.TtlCtx(ctx, key)
	if err != nil {
		ttl = 3600 // 默认1小时
	}

	return m.redis.SetexCtx(ctx, key, string(newData), int(ttl))
}

// ==================== Token黑名单相关 ====================

// AddToBlacklist 将Token加入黑名单
func (m *RedisTokenModel) AddToBlacklist(ctx context.Context, token string, entry *TokenBlacklistEntry, ttl int) error {
	key := fmt.Sprintf("token:blacklist:%s", token)

	data, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("序列化黑名单条目失败: %v", err)
	}

	if err := m.redis.SetexCtx(ctx, key, string(data), ttl); err != nil {
		return fmt.Errorf("保存黑名单条目失败: %v", err)
	}

	// 同时保存用户黑名单索引
	userBlacklistKey := fmt.Sprintf("user:blacklist:%d", entry.UserID)
	if _, err := m.redis.SaddCtx(ctx, userBlacklistKey, token); err != nil {
		return fmt.Errorf("保存用户黑名单索引失败: %v", err)
	}
	// 设置黑名单索引的TTL
	if err := m.redis.ExpireCtx(ctx, userBlacklistKey, ttl); err != nil {
		m.Logger.Errorf("设置黑名单索引TTL失败: %v", err)
	}

	return nil
}

// IsInBlacklist 检查Token是否在黑名单中
func (m *RedisTokenModel) IsInBlacklist(ctx context.Context, token string) (bool, error) {
	key := fmt.Sprintf("token:blacklist:%s", token)

	exists, err := m.redis.ExistsCtx(ctx, key)
	if err != nil {
		return false, fmt.Errorf("检查黑名单失败: %v", err)
	}

	return exists, nil
}

// GetBlacklistEntry 获取黑名单条目详情
func (m *RedisTokenModel) GetBlacklistEntry(ctx context.Context, token string) (*TokenBlacklistEntry, error) {
	key := fmt.Sprintf("token:blacklist:%s", token)

	data, err := m.redis.GetCtx(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("获取黑名单条目失败: %v", err)
	}

	if data == "" {
		return nil, fmt.Errorf("黑名单条目不存在")
	}

	var entry TokenBlacklistEntry
	if err := json.Unmarshal([]byte(data), &entry); err != nil {
		return nil, fmt.Errorf("反序列化黑名单条目失败: %v", err)
	}

	return &entry, nil
}

// ==================== Token刷新相关 ====================

// SaveRefreshRecord 保存Token刷新记录
func (m *RedisTokenModel) SaveRefreshRecord(ctx context.Context, record *TokenRefreshRecord, ttl int) error {
	key := fmt.Sprintf("token:refresh:%s:%d", record.OldToken, record.RefreshTime)

	data, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("序列化刷新记录失败: %v", err)
	}

	if err := m.redis.SetexCtx(ctx, key, string(data), ttl); err != nil {
		return fmt.Errorf("保存刷新记录失败: %v", err)
	}

	// 保存用户刷新历史
	userRefreshKey := fmt.Sprintf("user:refresh_history:%d", record.UserID)
	if _, err := m.redis.ZaddCtx(ctx, userRefreshKey, record.RefreshTime, record.OldToken); err != nil {
		return fmt.Errorf("保存用户刷新历史失败: %v", err)
	}

	// 限制历史记录数量（保留最近100条）
	if count, err := m.redis.ZcardCtx(ctx, userRefreshKey); err == nil && count > 100 {
		// 删除最旧的记录
		if oldTokens, err := m.redis.ZrangeCtx(ctx, userRefreshKey, 0, ga.Int64(count-101)); err == nil {
			for _, oldToken := range oldTokens {
				m.redis.ZremCtx(ctx, userRefreshKey, oldToken)
			}
		}
	}

	return nil
}

// GetRefreshRecords 获取Token刷新记录
func (m *RedisTokenModel) GetRefreshRecords(ctx context.Context, userID int64, limit int) ([]*TokenRefreshRecord, error) {
	userRefreshKey := fmt.Sprintf("user:refresh_history:%d", userID)

	if limit <= 0 || limit > 100 {
		limit = 20 // 默认返回20条
	}

	tokens, err := m.redis.ZrevrangeCtx(ctx, userRefreshKey, 0, int64(limit-1))
	if err != nil {
		return nil, fmt.Errorf("获取刷新历史失败: %v", err)
	}

	var records []*TokenRefreshRecord
	for _, token := range tokens {
		// 这里简化处理，实际需要根据token和时间构造key获取完整记录
		record := &TokenRefreshRecord{
			OldToken: token,
			UserID:   userID,
		}
		records = append(records, record)
	}

	return records, nil
}

// ==================== 用户Token管理相关 ====================

// GetUserActiveTokens 获取用户的活跃Token列表
func (m *RedisTokenModel) GetUserActiveTokens(ctx context.Context, userID int64) ([]*TokenMetadata, error) {
	userTokensKey := fmt.Sprintf("user:tokens:%d", userID)

	tokens, err := m.redis.SmembersCtx(ctx, userTokensKey)
	if err != nil {
		return nil, fmt.Errorf("获取用户Token列表失败: %v", err)
	}

	var activeTokens []*TokenMetadata
	for _, token := range tokens {
		metadata, err := m.GetTokenMetadata(ctx, token)
		if err != nil {
			m.Logger.Errorf("获取Token元数据失败: %v", err)
			continue
		}

		// 只返回未过期且活跃的Token
		if metadata.IsActive && time.Now().Unix() < metadata.ExpiresAt {
			activeTokens = append(activeTokens, metadata)
		}
	}

	return activeTokens, nil
}

// RevokeAllUserTokens 撤销用户的所有Token
func (m *RedisTokenModel) RevokeAllUserTokens(ctx context.Context, userID int64, reason string) error {
	userTokensKey := fmt.Sprintf("user:tokens:%d", userID)

	tokens, err := m.redis.SmembersCtx(ctx, userTokensKey)
	if err != nil {
		return fmt.Errorf("获取用户Token列表失败: %v", err)
	}

	for _, token := range tokens {
		// 将Token加入黑名单
		blacklistEntry := &TokenBlacklistEntry{
			Token:         token,
			UserID:        userID,
			BlacklistedAt: time.Now().Unix(),
			Reason:        reason,
			ExpiresAt:     time.Now().Unix() + 86400, // 默认24小时
		}

		if err := m.AddToBlacklist(ctx, token, blacklistEntry, 86400); err != nil {
			m.Logger.Errorf("Token加入黑名单失败: %v", err)
			continue
		}

		// 删除Token元数据
		metadataKey := fmt.Sprintf("token:metadata:%s", token)
		if _, err := m.redis.DelCtx(ctx, metadataKey); err != nil {
			m.Logger.Errorf("删除Token元数据失败: %v", err)
		}
	}

	// 删除用户Token索引
	if _, err := m.redis.DelCtx(ctx, userTokensKey); err != nil {
		return fmt.Errorf("删除用户Token索引失败: %v", err)
	}

	return nil
}

// ==================== 刷新频率限制相关 ====================

// CheckRefreshRateLimit 检查Token刷新频率限制
func (m *RedisTokenModel) CheckRefreshRateLimit(ctx context.Context, userID int64, maxRefreshPerHour int) (bool, error) {
	if maxRefreshPerHour <= 0 {
		maxRefreshPerHour = 5 // 默认每小时最多5次
	}

	key := fmt.Sprintf("token:refresh_limit:%d", userID)

	// 获取当前小时的刷新次数
	currentHour := time.Now().Format("2006010215") // 年月日时
	hourKey := fmt.Sprintf("%s:%s", key, currentHour)

	countStr, err := m.redis.GetCtx(ctx, hourKey)
	if err != nil {
		// 第一次刷新，设置计数器
		if err := m.redis.SetexCtx(ctx, hourKey, "1", 3600); err != nil {
			return false, fmt.Errorf("设置刷新计数器失败: %v", err)
		}
		return true, nil
	}

	count := 0
	if n, err := fmt.Sscanf(countStr, "%d", &count); err == nil && n == 1 {
		if count >= maxRefreshPerHour {
			return false, nil // 超过限制
		}

		// 增加刷新计数
		newCount := count + 1
		if err := m.redis.SetexCtx(ctx, hourKey, fmt.Sprintf("%d", newCount), 3600); err != nil {
			return false, fmt.Errorf("更新刷新计数器失败: %v", err)
		}
		return true, nil
	}

	return false, fmt.Errorf("解析刷新计数失败")
}

// ==================== 统计和分析相关 ====================

// GetTokenStatistics 获取Token统计信息
func (m *RedisTokenModel) GetTokenStatistics(ctx context.Context, userID int64) (map[string]interface{}, error) {
	activeTokens, err := m.GetUserActiveTokens(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("获取活跃Token失败: %v", err)
	}

	refreshRecords, err := m.GetRefreshRecords(ctx, userID, 10)
	if err != nil {
		return nil, fmt.Errorf("获取刷新记录失败: %v", err)
	}

	stats := map[string]interface{}{
		"active_token_count":  len(activeTokens),
		"total_refresh_count": len(refreshRecords),
		"last_refresh_time":   0,
		"avg_token_lifetime":  0,
		"refresh_rate_24h":    0,
	}

	if len(refreshRecords) > 0 {
		stats["last_refresh_time"] = refreshRecords[0].RefreshTime
	}

	return stats, nil
}

// CleanupExpiredData 清理过期数据（定时任务）
//func (m *RedisTokenModel) CleanupExpiredData(ctx context.Context) error {
//	m.Info("开始清理过期Token数据")
//
//	// 扫描所有Token黑名单
//	blacklistPattern := "token:blacklist:*"
//	// 这里可以实现批量扫描和清理逻辑
//
//	// 扫描所有Token元数据
//	metadataPattern := "token:metadata:*"
//	// 这里可以实现批量扫描和清理逻辑
//
//	m.Info("过期Token数据清理完成")
//	return nil
//}

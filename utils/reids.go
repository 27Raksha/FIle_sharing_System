package utils

import (
    "21BLC1564/models"
    "encoding/json"
    "fmt"
    "time"
	"context"
    "github.com/go-redis/redis/v8"
)


func CacheFiles(userID uint, files []models.File) error {
    data, err := json.Marshal(files)
    if err != nil {
        return err
    }

    
    return RedisClient.Set(RedisClient.Context(), fmt.Sprintf("files:user:%d", userID), data, 5*time.Minute).Err()
}


func GetCachedFiles(userID uint) ([]models.File, error) {
    val, err := RedisClient.Get(RedisClient.Context(), fmt.Sprintf("files:user:%d", userID)).Result()
    if err == redis.Nil {
        return nil, nil 
    } else if err != nil {
        return nil, err
    }

    var files []models.File
    err = json.Unmarshal([]byte(val), &files)
    if err != nil {
        return nil, err
    }

    return files, nil
}


func InvalidateCachedFiles(userID uint) error {
    return RedisClient.Del(RedisClient.Context(), fmt.Sprintf("files:user:%d", userID)).Err()
}
var ctx = context.Background()

func CacheSharedLink(fileID, sharedLink string, expiration time.Duration) error {
    err := RedisClient.Set(ctx, "share:file:"+fileID, sharedLink, expiration).Err()
    return err
}


func GetCachedSharedLink(fileID string) (string, error) {
    val, err := RedisClient.Get(RedisClient.Context(), fmt.Sprintf("share:file:%s", fileID)).Result()
    if err == redis.Nil {
        return "", nil 
    } else if err != nil {
        return "", err
    }

    return val, nil
}

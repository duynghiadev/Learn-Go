# Redis Cache Setup and Usage

This document explains how Redis caching is used in this project and how to interact with Redis using the Redis CLI.

## 🔧 Redis Setup (Local Machine)

### 1. Install Redis (if not already installed)

```bash
brew install redis
```

### 2. Start Redis Server

```bash
brew services start redis
```

Or, manually start Redis with:

```bash
redis-server
```

### 3. Verify Redis is Running

```bash
redis-cli ping
```

> If it responds with `PONG`, Redis is running correctly.

### 4. Stop Redis

```bash
brew services stop redis
```

---

## 📦 Redis in This Project

### 🔍 Cached Data

We cache the result of fetching all posts to improve performance and reduce Firestore reads.

### 🔑 Redis Key Used

- `all_posts` – stores the JSON-encoded array of all posts

### ⏳ Cache Expiry

- The cached posts are set to expire after 5 minutes.

### 💻 Relevant Code (FindAll method)

```go
// Cache the results
if len(posts) > 0 {
    postsJSON, err := json.Marshal(posts)
    if err == nil {
        err := r.redis.Set(ctx, cacheKey, postsJSON, 5*time.Minute).Err()
        if err != nil {
            log.Printf("Firestore Repo: Failed to cache posts in Redis: %v", err)
        }
    }
}
```

---

## 🛠️ Redis CLI Commands

### 1. Open Redis CLI

```bash
redis-cli
```

### 2. See All Keys

```bash
KEYS *
```

### 3. View Cached Posts

```bash
GET all_posts
```

### 4. Check Time-to-Live for Cached Posts

```bash
TTL all_posts
```

### 5. Manually Delete Cached Posts

```bash
DEL all_posts
```

### 6. Exit Redis CLI

```bash
exit
```

---

## ✅ Notes

- Make sure Redis is running before starting your Go application.
- Default Redis address: `localhost:6379`

### Run your app:

```bash
go run cmd/api/main.go
```

If Redis is not running, your application may throw a connection error when trying to cache or retrieve data.

---

Enjoy blazing-fast caching with Redis! 🚀

# 🧠 Redis Caching Guide

This document explains how Redis caching is used in the project, especially for storing and retrieving **posts** to reduce Firestore reads and improve performance.

---

## 💾 Caching Strategy

- When calling `FindAll()` on the post repository, the system:
  1. Checks Redis for the key `all_posts`
  2. If **cache hit** , the data is returned directly from Redis (as JSON)
  3. If **cache miss** , the data is fetched from Firestore and then **cached** in Redis for **5 minutes**

---

## 🔧 Redis CLI Commands

Use these commands in your terminal to interact with the Redis instance via CLI.

### 1. Open Redis CLI

```bash
redis-cli
```

### 2. See all keys stored in Redis

```bash
KEYS *
```

### 3. View the cached posts data

```bash
GET all_posts
```

- The result is a JSON string containing all post objects.
- You can use a JSON formatter to improve readability.

### 4. Check time-to-live for the cached posts key

```bash
TTL all_posts
```

- Shows how many **seconds** remain before the key expires.

### 5. Manually clear the cached data

```bash
DEL all_posts
```

### 6. Exit Redis CLI

```bash
exit
```

---

## 🔪 Testing the Cache

- After starting your backend, trigger a call to `GET /posts` (or the endpoint using `FindAll()`).
- First call → data fetched from Firestore and cached
- Next calls within 5 minutes → data comes from Redis (check logs to confirm)

---

## 👨‍🛠️ Notes

- Cache expires in: `5 minutes`
- Redis stores the data as: **serialized JSON array**
- If data structure changes, you may want to invalidate the cache to avoid mismatches

---

## 🗜 Optional: Clear Redis Automatically on New Post Creation

If needed, implement cache invalidation in the `Save()` method:

```go
_ = r.redis.Del(ctx, "all_posts").Err()
```

This ensures stale data is cleared when new data is added.

---

> Redis caching improves performance and reduces Firestore reads. Make sure to monitor cache hit/miss logs during development.

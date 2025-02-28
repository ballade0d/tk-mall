-- KEYS[1]：锁的键（如 "lock:<resource>"）
-- ARGV[1]：锁的标识符（如 "task_1"）
-- ARGV[2]：锁的超时时间（秒）

-- 尝试获取锁
local lock = redis.call("SETNX", KEYS[1], ARGV[1])
if lock == 1 then
    -- 如果获取锁成功，设置锁的过期时间
    redis.call("EXPIRE", KEYS[1], ARGV[2])
    return "LOCK_ACQUIRED"
else
    return "LOCK_FAILED"
end
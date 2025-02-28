-- KEYS[1]：锁的键（如 "lock:<resource>"）
-- ARGV[1]：锁的标识符（如 "task_1"）

-- 获取当前锁的值
local currentValue = redis.call("GET", KEYS[1])

-- 如果当前值等于请求的标识符，说明是锁的持有者
if currentValue == ARGV[1] then
    -- 删除锁
    redis.call("DEL", KEYS[1])
    return "LOCK_RELEASED"
else
    return "LOCK_NOT_OWNED"
end
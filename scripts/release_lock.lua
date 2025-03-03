if redis.call("get", KEYS[1]) == ARGV[1] then
    redis.call("del", KEYS[1])
    local next = redis.call("lpop", KEYS[2])
    if next then
        redis.call("publish", KEYS[2], next)
    end
    return 1
else
    return 0
end

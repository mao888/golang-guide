local key     = KEYS[1]
local content = ARGV[1]
local value = redis.call('get', key)

if value == content then
    local delResult = redis.call('del', key)
    if delResult == 1 then
        return 1
    end
    return 0
end
return 0
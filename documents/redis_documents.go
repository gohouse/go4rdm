package documents

import "encoding/json"

func GetAllCommands() string {
	return `all coomands list

//APPEND: Append a value to a key
APPEND key value

//AUTH: Authenticate to the server
AUTH password

//BGREWRITEAOF: Asynchronously rewrite the append-only file
BGREWRITEAOF

//BGSAVE: Asynchronously save the dataset to disk
BGSAVE

//BITCOUNT: Count set bits in a string
BITCOUNT key [start end]

//BITFIELD: Perform arbitrary bitfield integer operations on strings
BITFIELD key [GET type offset] [SET type offset value] [INCRBY type offset increment] [OVERFLOW WRAP|SAT|FAIL]

//BITOP: Perform bitwise operations between strings
BITOP operation destkey key [key ...]

//BITPOS: Find first bit set or clear in a string
BITPOS key bit [start] [end]

//BLPOP: Remove and get the first element in a list, or block until one is available
BLPOP key [key ...] timeout

//BRPOP: Remove and get the last element in a list, or block until one is available
BRPOP key [key ...] timeout

//BRPOPLPUSH: Pop a value from a list, push it to another list and return it; or block until one is available
BRPOPLPUSH source destination timeout

//CLIENT: Kill the connection of a client
CLIENT KILL [ip:port] [ID client-id] [TYPE normal|master|slave|pubsub] [ADDR ip:port] [SKIPME yes/no]

//CLIENT: Get the list of client connections
CLIENT LIST

//CLIENT: Get the current connection name
CLIENT GETNAME

//CLIENT: Stop processing commands from clients for some time
CLIENT PAUSE timeout

//CLIENT: Instruct the server whether to reply to commands
CLIENT REPLY ON|OFF|SKIP

//CLIENT: Set the current connection name
CLIENT SETNAME connection-name

//CLUSTER: Assign new hash slots to receiving node
CLUSTER ADDSLOTS slot [slot ...]

//CLUSTER: Return the number of failure reports active for a given node
CLUSTER COUNT-FAILURE-REPORTS node-id

//CLUSTER: Return the number of local keys in the specified hash slot
CLUSTER COUNTKEYSINSLOT slot

//CLUSTER: Set hash slots as unbound in receiving node
CLUSTER DELSLOTS slot [slot ...]

//CLUSTER: Forces a slave to perform a manual failover of its master.
CLUSTER FAILOVER [FORCE|TAKEOVER]

//CLUSTER: Remove a node from the nodes table
CLUSTER FORGET node-id

//CLUSTER: Return local key names in the specified hash slot
CLUSTER GETKEYSINSLOT slot count

//CLUSTER: Provides info about Redis Cluster node state
CLUSTER INFO

//CLUSTER: Returns the hash slot of the specified key
CLUSTER KEYSLOT key

//CLUSTER: Force a node cluster to handshake with another node
CLUSTER MEET ip port

//CLUSTER: Get Cluster config for the node
CLUSTER NODES

//CLUSTER: Reconfigure a node as a slave of the specified master node
CLUSTER REPLICATE node-id

//CLUSTER: Reset a Redis Cluster node
CLUSTER RESET [HARD|SOFT]

//CLUSTER: Forces the node to save cluster state on disk
CLUSTER SAVECONFIG

//CLUSTER: Set the configuration epoch in a new node
CLUSTER SET-CONFIG-EPOCH config-epoch

//CLUSTER: Bind a hash slot to a specific node
CLUSTER SETSLOT slot IMPORTING|MIGRATING|STABLE|NODE [node-id]

//CLUSTER: List slave nodes of the specified master node
CLUSTER SLAVES node-id

//CLUSTER: Get array of Cluster slot to node mappings
CLUSTER SLOTS

//COMMAND: Get array of Redis command details
COMMAND

//COMMAND: Get total number of Redis commands
COMMAND COUNT

//COMMAND: Extract keys given a full Redis command
COMMAND GETKEYS

//COMMAND: Get array of specific Redis command details
COMMAND INFO command-name [command-name ...]

//CONFIG: Get the value of a configuration parameter
CONFIG GET parameter

//CONFIG: Rewrite the configuration file with the in memory configuration
CONFIG REWRITE

//CONFIG: Set a configuration parameter to the given value
CONFIG SET parameter value

//CONFIG: Reset the stats returned by INFO
CONFIG RESETSTAT

//DBSIZE: Return the number of keys in the selected database
DBSIZE

//DEBUG: Get debugging information about a key
DEBUG OBJECT key

//DEBUG: Make the server crash
DEBUG SEGFAULT

//DECR: Decrement the integer value of a key by one
DECR key

//DECRBY: Decrement the integer value of a key by the given number
DECRBY key decrement

//DEL: Delete a key
DEL key [key ...]

//DISCARD: Discard all commands issued after MULTI
DISCARD

//DUMP: Return a serialized version of the value stored at the specified key.
DUMP key

//ECHO: Echo the given string
ECHO message

//EVAL: Execute a Lua script server side
EVAL script numkeys key [key ...] arg [arg ...]

//EVALSHA: Execute a Lua script server side
EVALSHA sha1 numkeys key [key ...] arg [arg ...]

//EXEC: Execute all commands issued after MULTI
EXEC

//EXISTS: Determine if a key exists
EXISTS key [key ...]

//EXPIRE: Set a key's time to live in seconds
EXPIRE key seconds

//EXPIREAT: Set the expiration for a key as a UNIX timestamp
EXPIREAT key timestamp

//FLUSHALL: Remove all keys from all databases
FLUSHALL [ASYNC]

//FLUSHDB: Remove all keys from the current database
FLUSHDB [ASYNC]

//GEOADD: Add one or more geospatial items in the geospatial index represented using a sorted set
GEOADD key longitude latitude member [longitude latitude member ...]

//GEOHASH: Returns members of a geospatial index as standard geohash strings
GEOHASH key member [member ...]

//GEOPOS: Returns longitude and latitude of members of a geospatial index
GEOPOS key member [member ...]

//GEODIST: Returns the distance between two members of a geospatial index
GEODIST key member1 member2 [unit]

//GEORADIUS: Query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a point
GEORADIUS key longitude latitude radius m|km|ft|mi [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count] [ASC|DESC] [STORE key] [STOREDIST key]

//GEORADIUSBYMEMBER: Query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a member
GEORADIUSBYMEMBER key member radius m|km|ft|mi [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count] [ASC|DESC] [STORE key] [STOREDIST key]

//GET: Get the value of a key
GET key

//GETBIT: Returns the bit value at offset in the string value stored at key
GETBIT key offset

//GETRANGE: Get a substring of the string stored at a key
GETRANGE key start end

//GETSET: Set the string value of a key and return its old value
GETSET key value

//HDEL: Delete one or more hash fields
HDEL key field [field ...]

//HEXISTS: Determine if a hash field exists
HEXISTS key field

//HGET: Get the value of a hash field
HGET key field

//HGETALL: Get all the fields and values in a hash
HGETALL key

//HINCRBY: Increment the integer value of a hash field by the given number
HINCRBY key field increment

//HINCRBYFLOAT: Increment the float value of a hash field by the given amount
HINCRBYFLOAT key field increment

//HKEYS: Get all the fields in a hash
HKEYS key

//HLEN: Get the number of fields in a hash
HLEN key

//HMGET: Get the values of all the given hash fields
HMGET key field [field ...]

//HMSET: Set multiple hash fields to multiple values
HMSET key field value [field value ...]

//HSET: Set the string value of a hash field
HSET key field value

//HSETNX: Set the value of a hash field, only if the field does not exist
HSETNX key field value

//HSTRLEN: Get the length of the value of a hash field
HSTRLEN key field

//HVALS: Get all the values in a hash
HVALS key

//INCR: Increment the integer value of a key by one
INCR key

//INCRBY: Increment the integer value of a key by the given amount
INCRBY key increment

//INCRBYFLOAT: Increment the float value of a key by the given amount
INCRBYFLOAT key increment

//INFO: Get information and statistics about the server
INFO [section]

//KEYS: Find all keys matching the given pattern
KEYS pattern

//LASTSAVE: Get the UNIX time stamp of the last successful save to disk
LASTSAVE

//LINDEX: Get an element from a list by its index
LINDEX key index

//LINSERT: Insert an element before or after another element in a list
LINSERT key BEFORE|AFTER pivot value

//LLEN: Get the length of a list
LLEN key

//LPOP: Remove and get the first element in a list
LPOP key

//LPUSH: Prepend one or multiple values to a list
LPUSH key value [value ...]

//LPUSHX: Prepend a value to a list, only if the list exists
LPUSHX key value

//LRANGE: Get a range of elements from a list
LRANGE key start stop

//LREM: Remove elements from a list
LREM key count value

//LSET: Set the value of an element in a list by its index
LSET key index value

//LTRIM: Trim a list to the specified range
LTRIM key start stop

//MGET: Get the values of all the given keys
MGET key [key ...]

//MIGRATE: Atomically transfer a key from a Redis instance to another one.
MIGRATE host port key|"" destination-db timeout [COPY] [REPLACE] [KEYS key [key ...]]

//MONITOR: Listen for all requests received by the server in real time
MONITOR

//MOVE: Move a key to another database
MOVE key db

//MSET: Set multiple keys to multiple values
MSET key value [key value ...]

//MSETNX: Set multiple keys to multiple values, only if none of the keys exist
MSETNX key value [key value ...]

//MULTI: Mark the start of a transaction block
MULTI

//OBJECT: Inspect the internals of Redis objects
OBJECT subcommand [arguments [arguments ...]]

//PERSIST: Remove the expiration from a key
PERSIST key

//PEXPIRE: Set a key's time to live in milliseconds
PEXPIRE key milliseconds

//PEXPIREAT: Set the expiration for a key as a UNIX timestamp specified in milliseconds
PEXPIREAT key milliseconds-timestamp

//PFADD: Adds the specified elements to the specified HyperLogLog.
PFADD key element [element ...]

//PFCOUNT: Return the approximated cardinality of the set(s) observed by the HyperLogLog at key(s).
PFCOUNT key [key ...]

//PFMERGE: Merge N different HyperLogLogs into a single one.
PFMERGE destkey sourcekey [sourcekey ...]

//PING: Ping the server
PING [message]

//PSETEX: Set the value and expiration in milliseconds of a key
PSETEX key milliseconds value

//PSUBSCRIBE: Listen for messages published to channels matching the given patterns
PSUBSCRIBE pattern [pattern ...]

//PUBSUB: Inspect the state of the Pub/Sub subsystem
PUBSUB subcommand [argument [argument ...]]

//PTTL: Get the time to live for a key in milliseconds
PTTL key

//PUBLISH: Post a message to a channel
PUBLISH channel message

//PUNSUBSCRIBE: Stop listening for messages posted to channels matching the given patterns
PUNSUBSCRIBE [pattern [pattern ...]]

//QUIT: Close the connection
QUIT

//RANDOMKEY: Return a random key from the keyspace
RANDOMKEY

//READONLY: Enables read queries for a connection to a cluster slave node
READONLY

//READWRITE: Disables read queries for a connection to a cluster slave node
READWRITE

//RENAME: Rename a key
RENAME key newkey

//RENAMENX: Rename a key, only if the new key does not exist
RENAMENX key newkey

//RESTORE: Rename a key, only if the new key does not exist
RESTORE key ttl serialized-value [REPLACE]

//RESTORE: Create a key using the provided serialized value, previously obtained using DUMP.
RESTORE key ttl serialized-value [REPLACE]

//ROLE: Return the role of the instance in the context of replication
ROLE

//RPOP: Remove and get the last element in a list
RPOP key

//RPOPLPUSH: Remove the last element in a list, prepend it to another list and return it
RPOPLPUSH source destination

//RPUSH: Append one or multiple values to a list
RPUSH key value [value ...]

//RPUSHX: Append a value to a list, only if the list exists
RPUSHX key value

//SADD: Add one or more members to a set
SADD key member [member ...]

//SAVE: Synchronously save the dataset to disk
SAVE

//SCARD: Get the number of members in a set
SCARD key

//SCRIPT: Set the debug mode for executed scripts.
SCRIPT DEBUG YES|SYNC|NO

//SCRIPT: Check existence of scripts in the script cache.
SCRIPT EXISTS sha1 [sha1 ...]

//SCRIPT: Remove all the scripts from the script cache.
SCRIPT FLUSH

//SCRIPT: Kill the script currently in execution.
SCRIPT KILL

//SCRIPT: Load the specified Lua script into the script cache.
SCRIPT LOAD script

//SDIFF: Subtract multiple sets
SDIFF key [key ...]

//SDIFFSTORE: Subtract multiple sets and store the resulting set in a key
SDIFFSTORE destination key [key ...]

//SELECT: Change the selected database for the current connection
SELECT index

//SET: Set the string value of a key
SET key value [EX seconds] [PX milliseconds] [NX|XX]

//SETBIT: Sets or clears the bit at offset in the string value stored at key
SETBIT key offset value

//SETEX: Set the value and expiration of a key
SETEX key seconds value

//SETNX: Set the value of a key, only if the key does not exist
SETNX key value

//SETRANGE: Overwrite part of a string at key starting at the specified offset
SETRANGE key offset value

//SHUTDOWN: Synchronously save the dataset to disk and then shut down the server
SHUTDOWN [NOSAVE|SAVE]

//SINTER: Intersect multiple sets
SINTER key [key ...]

//SINTERSTORE: Intersect multiple sets and store the resulting set in a key
SINTERSTORE destination key [key ...]

//SISMEMBER: Determine if a given value is a member of a set
SISMEMBER key member

//SLAVEOF: Make the server a slave of another instance, or promote it as master
SLAVEOF host port

//SLOWLOG: Manages the Redis slow queries log
SLOWLOG subcommand [argument]

//SMEMBERS: Get all the members in a set
SMEMBERS key

//SMOVE: Move a member from one set to another
SMOVE source destination member

//SORT: Sort the elements in a list, set or sorted set
SORT key [BY pattern] [LIMIT offset count] [GET pattern [GET pattern ...]] [ASC|DESC] [ALPHA] [STORE destination]

//SPOP: Remove and return one or multiple random members from a set
SPOP key [count]

//SRANDMEMBER: Get one or multiple random members from a set
SRANDMEMBER key [count]

//SREM: Remove one or more members from a set
SREM key member [member ...]

//STRLEN: Get the length of the value stored in a key
STRLEN key

//SUBSCRIBE: Listen for messages published to the given channels
SUBSCRIBE channel [channel ...]

//SUNION: Add multiple sets
SUNION key [key ...]

//SUNIONSTORE: Add multiple sets and store the resulting set in a key
SUNIONSTORE destination key [key ...]

//SWAPDB: Swaps two Redis databases
SWAPDB index index

//SYNC: Internal command used for replication
SYNC

//TIME: Return the current server time
TIME

//TOUCH: Alters the last access time of a key(s). Returns the number of existing keys specified.
TOUCH key [key ...]

//TTL: Get the time to live for a key
TTL key

//TYPE: Determine the type stored at key
TYPE key

//UNSUBSCRIBE: Stop listening for messages posted to the given channels
UNSUBSCRIBE [channel [channel ...]]

//UNLINK: Delete a key asynchronously in another thread. Otherwise it is just as DEL, but non blocking.
UNLINK key [key ...]

//UNWATCH: Forget about all watched keys
UNWATCH

//WAIT: Wait for the synchronous replication of all the write commands sent in the context of the current connection
WAIT numslaves timeout

//WATCH: Watch the given keys to determine execution of the MULTI/EXEC block
WATCH key [key ...]

//ZADD: Add one or more members to a sorted set, or update its score if it already exists
ZADD key [NX|XX] [CH] [INCR] score member [score member ...]

//ZCARD: Get the number of members in a sorted set
ZCARD key

//ZCOUNT: Count the members in a sorted set with scores within the given values
ZCOUNT key min max

//ZINCRBY: Increment the score of a member in a sorted set
ZINCRBY key increment member

//ZINTERSTORE: Intersect multiple sorted sets and store the resulting sorted set in a new key
ZINTERSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX]

//ZLEXCOUNT: Count the number of members in a sorted set between a given lexicographical range
ZLEXCOUNT key min max

//ZRANGE: Return a range of members in a sorted set, by index
ZRANGE key start stop [WITHSCORES]

//ZRANGEBYLEX: Return a range of members in a sorted set, by lexicographical range
ZRANGEBYLEX key min max [LIMIT offset count]

//ZREVRANGEBYLEX: Return a range of members in a sorted set, by lexicographical range, ordered from higher to lower strings.
ZREVRANGEBYLEX key max min [LIMIT offset count]

//ZRANGEBYSCORE: Return a range of members in a sorted set, by score
ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]

//ZRANK: Determine the index of a member in a sorted set
ZRANK key member

//ZREM: Remove one or more members from a sorted set
ZREM key member [member ...]

//ZREMRANGEBYLEX: Remove all members in a sorted set between the given lexicographical range
ZREMRANGEBYLEX key min max

//ZREMRANGEBYRANK: Remove all members in a sorted set within the given indexes
ZREMRANGEBYRANK key start stop

//ZREMRANGEBYSCORE: Remove all members in a sorted set within the given scores
ZREMRANGEBYSCORE key min max

//ZREVRANGE: Return a range of members in a sorted set, by index, with scores ordered from high to low
ZREVRANGE key start stop [WITHSCORES]

//ZREVRANGEBYSCORE: Return a range of members in a sorted set, by score, with scores ordered from high to low
ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]

//ZREVRANK: Determine the index of a member in a sorted set, with scores ordered from high to low
ZREVRANK key member

//ZSCORE: Get the score associated with the given member in a sorted set
ZSCORE key member

//ZUNIONSTORE: Add multiple sorted sets and store the resulting sorted set in a new key
ZUNIONSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX]

//SCAN: Incrementally iterate the keys space
SCAN cursor [MATCH pattern] [COUNT count]

//SSCAN: Incrementally iterate Set elements
SSCAN key cursor [MATCH pattern] [COUNT count]

//HSCAN: Incrementally iterate hash fields and associated values
HSCAN key cursor [MATCH pattern] [COUNT count]

//ZSCAN: Incrementally iterate sorted sets elements and associated scores
ZSCAN key cursor [MATCH pattern] [COUNT count]
`
}

func GetJsonCommands() map[string]interface{} {
	var js = make(map[string]interface{})

	err := json.Unmarshal([]byte(getJsonString()), &js)
	if err!=nil {
		panic(err.Error())
	}

	return js
}

func getJsonString() string {
	return `{
  "geo": [
    {
      "CmdName": "GEORADIUS",
      "summary": "Query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a point",
      "complexity": "O(N+log(M)) where N is the number of elements inside the bounding box of the circular area delimited by center and radius and M is the number of items inside the index.",
      "since": "3.2.0",
      "group": "geo",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "longitude",
          "type": "double"
        },
        {
          "name": "latitude",
          "type": "double"
        },
        {
          "name": "radius",
          "type": "double"
        },
        {
          "enum": [
            "m",
            "km",
            "ft",
            "mi"
          ],
          "name": "unit",
          "type": "enum"
        },
        {
          "enum": [
            "WITHCOORD"
          ],
          "name": "withcoord",
          "optional": true,
          "type": "enum"
        },
        {
          "enum": [
            "WITHDIST"
          ],
          "name": "withdist",
          "optional": true,
          "type": "enum"
        },
        {
          "enum": [
            "WITHHASH"
          ],
          "name": "withhash",
          "optional": true,
          "type": "enum"
        },
        {
          "command": "COUNT",
          "name": "count",
          "optional": true,
          "type": "integer"
        },
        {
          "enum": [
            "ASC",
            "DESC"
          ],
          "name": "order",
          "optional": true,
          "type": "enum"
        },
        {
          "command": "STORE",
          "name": "key",
          "optional": true,
          "type": "key"
        },
        {
          "command": "STOREDIST",
          "name": "key",
          "optional": true,
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "GEORADIUSBYMEMBER",
      "summary": "Query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a member",
      "complexity": "O(N+log(M)) where N is the number of elements inside the bounding box of the circular area delimited by center and radius and M is the number of items inside the index.",
      "since": "3.2.0",
      "group": "geo",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "member",
          "type": "string"
        },
        {
          "name": "radius",
          "type": "double"
        },
        {
          "enum": [
            "m",
            "km",
            "ft",
            "mi"
          ],
          "name": "unit",
          "type": "enum"
        },
        {
          "enum": [
            "WITHCOORD"
          ],
          "name": "withcoord",
          "optional": true,
          "type": "enum"
        },
        {
          "enum": [
            "WITHDIST"
          ],
          "name": "withdist",
          "optional": true,
          "type": "enum"
        },
        {
          "enum": [
            "WITHHASH"
          ],
          "name": "withhash",
          "optional": true,
          "type": "enum"
        },
        {
          "command": "COUNT",
          "name": "count",
          "optional": true,
          "type": "integer"
        },
        {
          "enum": [
            "ASC",
            "DESC"
          ],
          "name": "order",
          "optional": true,
          "type": "enum"
        },
        {
          "command": "STORE",
          "name": "key",
          "optional": true,
          "type": "key"
        },
        {
          "command": "STOREDIST",
          "name": "key",
          "optional": true,
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "GEODIST",
      "summary": "Returns the distance between two members of a geospatial index",
      "complexity": "O(log(N))",
      "since": "3.2.0",
      "group": "geo",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "member1",
          "type": "string"
        },
        {
          "name": "member2",
          "type": "string"
        },
        {
          "enum": [
            "m",
            "km",
            "ft",
            "mi"
          ],
          "name": "unit",
          "optional": true,
          "type": "enum"
        }
      ]
    },
    {
      "CmdName": "GEOADD",
      "summary": "Add one or more geospatial items in the geospatial index represented using a sorted set",
      "complexity": "O(log(N)) for each item added, where N is the number of elements in the sorted set.",
      "since": "3.2.0",
      "group": "geo",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": [
            "longitude",
            "latitude",
            "member"
          ],
          "type": [
            "double",
            "double",
            "string"
          ]
        }
      ]
    },
    {
      "CmdName": "GEOPOS",
      "summary": "Returns longitude and latitude of members of a geospatial index",
      "complexity": "O(log(N)) for each member requested, where N is the number of elements in the sorted set.",
      "since": "3.2.0",
      "group": "geo",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "member",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "GEOHASH",
      "summary": "Returns members of a geospatial index as standard geohash strings",
      "complexity": "O(log(N)) for each member requested, where N is the number of elements in the sorted set.",
      "since": "3.2.0",
      "group": "geo",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "member",
          "type": "string"
        }
      ]
    }
  ],
  "hash": [
    {
      "CmdName": "HLEN",
      "summary": "Get the number of fields in a hash",
      "complexity": "O(1)",
      "since": "2.0.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "HSCAN",
      "summary": "Incrementally iterate hash fields and associated values",
      "complexity": "O(1) for every call. O(N) for a complete iteration, including enough command calls for the cursor to return back to 0. N is the number of elements inside the collection..",
      "since": "2.8.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "cursor",
          "type": "integer"
        },
        {
          "command": "MATCH",
          "name": "pattern",
          "optional": true,
          "type": "pattern"
        },
        {
          "command": "COUNT",
          "name": "count",
          "optional": true,
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "HDEL",
      "summary": "Delete one or more hash fields",
      "complexity": "O(N) where N is the number of fields to be removed.",
      "since": "2.0.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "field",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "HKEYS",
      "summary": "Get all the fields in a hash",
      "complexity": "O(N) where N is the size of the hash.",
      "since": "2.0.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "HMGET",
      "summary": "Get the values of all the given hash fields",
      "complexity": "O(N) where N is the number of fields being requested.",
      "since": "2.0.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "field",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "HINCRBYFLOAT",
      "summary": "Increment the float value of a hash field by the given amount",
      "complexity": "O(1)",
      "since": "2.6.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "field",
          "type": "string"
        },
        {
          "name": "increment",
          "type": "double"
        }
      ]
    },
    {
      "CmdName": "HGETALL",
      "summary": "Get all the fields and values in a hash",
      "complexity": "O(N) where N is the size of the hash.",
      "since": "2.0.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "HSTRLEN",
      "summary": "Get the length of the value of a hash field",
      "complexity": "O(1)",
      "since": "3.2.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "field",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "HGET",
      "summary": "Get the value of a hash field",
      "complexity": "O(1)",
      "since": "2.0.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "field",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "HEXISTS",
      "summary": "Determine if a hash field exists",
      "complexity": "O(1)",
      "since": "2.0.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "field",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "HMSET",
      "summary": "Set multiple hash fields to multiple values",
      "complexity": "O(N) where N is the number of fields being set.",
      "since": "2.0.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": [
            "field",
            "value"
          ],
          "type": [
            "string",
            "string"
          ]
        }
      ]
    },
    {
      "CmdName": "HSET",
      "summary": "Set the string value of a hash field",
      "complexity": "O(1) for each field/value pair added, so O(N) to add N field/value pairs when the command is called with multiple field/value pairs.",
      "since": "2.0.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": [
            "field",
            "value"
          ],
          "type": [
            "string",
            "string"
          ]
        }
      ]
    },
    {
      "CmdName": "HSETNX",
      "summary": "Set the value of a hash field, only if the field does not exist",
      "complexity": "O(1)",
      "since": "2.0.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "field",
          "type": "string"
        },
        {
          "name": "value",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "HVALS",
      "summary": "Get all the values in a hash",
      "complexity": "O(N) where N is the size of the hash.",
      "since": "2.0.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "HINCRBY",
      "summary": "Increment the integer value of a hash field by the given number",
      "complexity": "O(1)",
      "since": "2.0.0",
      "group": "hash",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "field",
          "type": "string"
        },
        {
          "name": "increment",
          "type": "integer"
        }
      ]
    }
  ],
  "hyperloglog": [
    {
      "CmdName": "PFCOUNT",
      "summary": "Return the approximated cardinality of the set(s) observed by the HyperLogLog at key(s).",
      "complexity": "O(1) with a very small average constant time when called with a single key. O(N) with N being the number of keys, and much bigger constant times, when called with multiple keys.",
      "since": "2.8.9",
      "group": "hyperloglog",
      "arguments": [
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "PFADD",
      "summary": "Adds the specified elements to the specified HyperLogLog.",
      "complexity": "O(1) to add every element.",
      "since": "2.8.9",
      "group": "hyperloglog",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "element",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "PFMERGE",
      "summary": "Merge N different HyperLogLogs into a single one.",
      "complexity": "O(N) to merge N HyperLogLogs, but with high constant times.",
      "since": "2.8.9",
      "group": "hyperloglog",
      "arguments": [
        {
          "name": "destkey",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "sourcekey",
          "type": "key"
        }
      ]
    }
  ],
  "list": [
    {
      "CmdName": "RPUSHX",
      "summary": "Append an element to a list, only if the list exists",
      "complexity": "O(1) for each element added, so O(N) to add N elements when the command is called with multiple arguments.",
      "since": "2.2.0",
      "group": "list",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "element",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "LPUSH",
      "summary": "Prepend one or multiple elements to a list",
      "complexity": "O(1) for each element added, so O(N) to add N elements when the command is called with multiple arguments.",
      "since": "1.0.0",
      "group": "list",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "element",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "RPOP",
      "summary": "Remove and get the last element in a list",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "list",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "BRPOP",
      "summary": "Remove and get the last element in a list, or block until one is available",
      "complexity": "O(1)",
      "since": "2.0.0",
      "group": "list",
      "arguments": [
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        },
        {
          "name": "timeout",
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "LREM",
      "summary": "Remove elements from a list",
      "complexity": "O(N+M) where N is the length of the list and M is the number of elements removed.",
      "since": "1.0.0",
      "group": "list",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "count",
          "type": "integer"
        },
        {
          "name": "element",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "LINDEX",
      "summary": "Get an element from a list by its index",
      "complexity": "O(N) where N is the number of elements to traverse to get to the element at index. This makes asking for the first or the last element of the list O(1).",
      "since": "1.0.0",
      "group": "list",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "index",
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "LPOS",
      "summary": "Return the index of matching elements on a list",
      "complexity": "O(N) where N is the number of elements in the list, for the average case. When searching for elements near the head or the tail of the list, or when the MAXLEN option is provided, the command may run in constant time.",
      "since": "6.0.6",
      "group": "list",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "element",
          "type": "string"
        },
        {
          "command": "RANK",
          "name": "rank",
          "optional": true,
          "type": "integer"
        },
        {
          "command": "COUNT",
          "name": "num-matches",
          "optional": true,
          "type": "integer"
        },
        {
          "command": "MAXLEN",
          "name": "len",
          "optional": true,
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "RPUSH",
      "summary": "Append one or multiple elements to a list",
      "complexity": "O(1) for each element added, so O(N) to add N elements when the command is called with multiple arguments.",
      "since": "1.0.0",
      "group": "list",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "element",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "LLEN",
      "summary": "Get the length of a list",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "list",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "LINSERT",
      "summary": "Insert an element before or after another element in a list",
      "complexity": "O(N) where N is the number of elements to traverse before seeing the value pivot. This means that inserting somewhere on the left end on the list (head) can be considered O(1) and inserting somewhere on the right end (tail) is O(N).",
      "since": "2.2.0",
      "group": "list",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "enum": [
            "BEFORE",
            "AFTER"
          ],
          "name": "where",
          "type": "enum"
        },
        {
          "name": "pivot",
          "type": "string"
        },
        {
          "name": "element",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "LPUSHX",
      "summary": "Prepend an element to a list, only if the list exists",
      "complexity": "O(1) for each element added, so O(N) to add N elements when the command is called with multiple arguments.",
      "since": "2.2.0",
      "group": "list",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "element",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "RPOPLPUSH",
      "summary": "Remove the last element in a list, prepend it to another list and return it",
      "complexity": "O(1)",
      "since": "1.2.0",
      "group": "list",
      "arguments": [
        {
          "name": "source",
          "type": "key"
        },
        {
          "name": "destination",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "LTRIM",
      "summary": "Trim a list to the specified range",
      "complexity": "O(N) where N is the number of elements to be removed by the operation.",
      "since": "1.0.0",
      "group": "list",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "start",
          "type": "integer"
        },
        {
          "name": "stop",
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "LPOP",
      "summary": "Remove and get the first element in a list",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "list",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "LRANGE",
      "summary": "Get a range of elements from a list",
      "complexity": "O(S+N) where S is the distance of start offset from HEAD for small lists, from nearest end (HEAD or TAIL) for large lists; and N is the number of elements in the specified range.",
      "since": "1.0.0",
      "group": "list",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "start",
          "type": "integer"
        },
        {
          "name": "stop",
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "LSET",
      "summary": "Set the value of an element in a list by its index",
      "complexity": "O(N) where N is the length of the list. Setting either the first or the last element of the list is O(1).",
      "since": "1.0.0",
      "group": "list",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "index",
          "type": "integer"
        },
        {
          "name": "element",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "BLPOP",
      "summary": "Remove and get the first element in a list, or block until one is available",
      "complexity": "O(1)",
      "since": "2.0.0",
      "group": "list",
      "arguments": [
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        },
        {
          "name": "timeout",
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "BRPOPLPUSH",
      "summary": "Pop an element from a list, push it to another list and return it; or block until one is available",
      "complexity": "O(1)",
      "since": "2.2.0",
      "group": "list",
      "arguments": [
        {
          "name": "source",
          "type": "key"
        },
        {
          "name": "destination",
          "type": "key"
        },
        {
          "name": "timeout",
          "type": "integer"
        }
      ]
    }
  ],
  "pubsub": [
    {
      "CmdName": "PUBSUB",
      "summary": "Inspect the state of the Pub/Sub subsystem",
      "complexity": "O(N) for the CHANNELS subcommand, where N is the number of active channels, and assuming constant time pattern matching (relatively short channels and patterns). O(N) for the NUMSUB subcommand, where N is the number of requested channels. O(1) for the NUMPAT subcommand.",
      "since": "2.8.0",
      "group": "pubsub",
      "arguments": [
        {
          "name": "subcommand",
          "type": "string"
        },
        {
          "multiple": true,
          "name": "argument",
          "optional": true,
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "SUBSCRIBE",
      "summary": "Listen for messages published to the given channels",
      "complexity": "O(N) where N is the number of channels to subscribe to.",
      "since": "2.0.0",
      "group": "pubsub",
      "arguments": [
        {
          "multiple": true,
          "name": "channel",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "UNSUBSCRIBE",
      "summary": "Stop listening for messages posted to the given channels",
      "complexity": "O(N) where N is the number of clients already subscribed to a channel.",
      "since": "2.0.0",
      "group": "pubsub",
      "arguments": [
        {
          "multiple": true,
          "name": "channel",
          "optional": true,
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "PSUBSCRIBE",
      "summary": "Listen for messages published to channels matching the given patterns",
      "complexity": "O(N) where N is the number of patterns the client is already subscribed to.",
      "since": "2.0.0",
      "group": "pubsub",
      "arguments": [
        {
          "multiple": true,
          "name": [
            "pattern"
          ],
          "type": [
            "pattern"
          ]
        }
      ]
    },
    {
      "CmdName": "PUNSUBSCRIBE",
      "summary": "Stop listening for messages posted to channels matching the given patterns",
      "complexity": "O(N+M) where N is the number of patterns the client is already subscribed and M is the number of total patterns subscribed in the system (by any client).",
      "since": "2.0.0",
      "group": "pubsub",
      "arguments": [
        {
          "multiple": true,
          "name": "pattern",
          "optional": true,
          "type": "pattern"
        }
      ]
    },
    {
      "CmdName": "PUBLISH",
      "summary": "Post a message to a channel",
      "complexity": "O(N+M) where N is the number of clients subscribed to the receiving channel and M is the total number of subscribed patterns (by any client).",
      "since": "2.0.0",
      "group": "pubsub",
      "arguments": [
        {
          "name": "channel",
          "type": "string"
        },
        {
          "name": "message",
          "type": "string"
        }
      ]
    }
  ],
  "set": [
    {
      "CmdName": "SREM",
      "summary": "Remove one or more members from a set",
      "complexity": "O(N) where N is the number of members to be removed.",
      "since": "1.0.0",
      "group": "set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "member",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "SMEMBERS",
      "summary": "Get all the members in a set",
      "complexity": "O(N) where N is the set cardinality.",
      "since": "1.0.0",
      "group": "set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "SSCAN",
      "summary": "Incrementally iterate Set elements",
      "complexity": "O(1) for every call. O(N) for a complete iteration, including enough command calls for the cursor to return back to 0. N is the number of elements inside the collection..",
      "since": "2.8.0",
      "group": "set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "cursor",
          "type": "integer"
        },
        {
          "command": "MATCH",
          "name": "pattern",
          "optional": true,
          "type": "pattern"
        },
        {
          "command": "COUNT",
          "name": "count",
          "optional": true,
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "SINTER",
      "summary": "Intersect multiple sets",
      "complexity": "O(N*M) worst case where N is the cardinality of the smallest set and M is the number of sets.",
      "since": "1.0.0",
      "group": "set",
      "arguments": [
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "SCARD",
      "summary": "Get the number of members in a set",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "SUNIONSTORE",
      "summary": "Add multiple sets and store the resulting set in a key",
      "complexity": "O(N) where N is the total number of elements in all given sets.",
      "since": "1.0.0",
      "group": "set",
      "arguments": [
        {
          "name": "destination",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "SPOP",
      "summary": "Remove and return one or multiple random members from a set",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "count",
          "optional": true,
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "SDIFFSTORE",
      "summary": "Subtract multiple sets and store the resulting set in a key",
      "complexity": "O(N) where N is the total number of elements in all given sets.",
      "since": "1.0.0",
      "group": "set",
      "arguments": [
        {
          "name": "destination",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "SUNION",
      "summary": "Add multiple sets",
      "complexity": "O(N) where N is the total number of elements in all given sets.",
      "since": "1.0.0",
      "group": "set",
      "arguments": [
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "SMOVE",
      "summary": "Move a member from one set to another",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "set",
      "arguments": [
        {
          "name": "source",
          "type": "key"
        },
        {
          "name": "destination",
          "type": "key"
        },
        {
          "name": "member",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "SRANDMEMBER",
      "summary": "Get one or multiple random members from a set",
      "complexity": "Without the count argument O(1), otherwise O(N) where N is the absolute value of the passed count.",
      "since": "1.0.0",
      "group": "set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "count",
          "optional": true,
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "SINTERSTORE",
      "summary": "Intersect multiple sets and store the resulting set in a key",
      "complexity": "O(N*M) worst case where N is the cardinality of the smallest set and M is the number of sets.",
      "since": "1.0.0",
      "group": "set",
      "arguments": [
        {
          "name": "destination",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "SDIFF",
      "summary": "Subtract multiple sets",
      "complexity": "O(N) where N is the total number of elements in all given sets.",
      "since": "1.0.0",
      "group": "set",
      "arguments": [
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "SADD",
      "summary": "Add one or more members to a set",
      "complexity": "O(1) for each element added, so O(N) to add N elements when the command is called with multiple arguments.",
      "since": "1.0.0",
      "group": "set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "member",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "SISMEMBER",
      "summary": "Determine if a given value is a member of a set",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "member",
          "type": "string"
        }
      ]
    }
  ],
  "sorted_set": [
    {
      "CmdName": "ZPOPMAX",
      "summary": "Remove and return members with the highest scores in a sorted set",
      "complexity": "O(log(N)*M) with N being the number of elements in the sorted set, and M being the number of elements popped.",
      "since": "5.0.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "count",
          "optional": true,
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "ZLEXCOUNT",
      "summary": "Count the number of members in a sorted set between a given lexicographical range",
      "complexity": "O(log(N)) with N being the number of elements in the sorted set.",
      "since": "2.8.9",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "min",
          "type": "string"
        },
        {
          "name": "max",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "ZADD",
      "summary": "Add one or more members to a sorted set, or update its score if it already exists",
      "complexity": "O(log(N)) for each item added, where N is the number of elements in the sorted set.",
      "since": "1.2.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "enum": [
            "NX",
            "XX"
          ],
          "name": "condition",
          "optional": true,
          "type": "enum"
        },
        {
          "enum": [
            "CH"
          ],
          "name": "change",
          "optional": true,
          "type": "enum"
        },
        {
          "enum": [
            "INCR"
          ],
          "name": "increment",
          "optional": true,
          "type": "enum"
        },
        {
          "multiple": true,
          "name": [
            "score",
            "member"
          ],
          "type": [
            "double",
            "string"
          ]
        }
      ]
    },
    {
      "CmdName": "ZREVRANK",
      "summary": "Determine the index of a member in a sorted set, with scores ordered from high to low",
      "complexity": "O(log(N))",
      "since": "2.0.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "member",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "ZSCAN",
      "summary": "Incrementally iterate sorted sets elements and associated scores",
      "complexity": "O(1) for every call. O(N) for a complete iteration, including enough command calls for the cursor to return back to 0. N is the number of elements inside the collection..",
      "since": "2.8.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "cursor",
          "type": "integer"
        },
        {
          "command": "MATCH",
          "name": "pattern",
          "optional": true,
          "type": "pattern"
        },
        {
          "command": "COUNT",
          "name": "count",
          "optional": true,
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "ZRANK",
      "summary": "Determine the index of a member in a sorted set",
      "complexity": "O(log(N))",
      "since": "2.0.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "member",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "ZREVRANGE",
      "summary": "Return a range of members in a sorted set, by index, with scores ordered from high to low",
      "complexity": "O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements returned.",
      "since": "1.2.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "start",
          "type": "integer"
        },
        {
          "name": "stop",
          "type": "integer"
        },
        {
          "enum": [
            "WITHSCORES"
          ],
          "name": "withscores",
          "optional": true,
          "type": "enum"
        }
      ]
    },
    {
      "CmdName": "ZREM",
      "summary": "Remove one or more members from a sorted set",
      "complexity": "O(M*log(N)) with N being the number of elements in the sorted set and M the number of elements to be removed.",
      "since": "1.2.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "member",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "ZREVRANGEBYLEX",
      "summary": "Return a range of members in a sorted set, by lexicographical range, ordered from higher to lower strings.",
      "complexity": "O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements being returned. If M is constant (e.g. always asking for the first 10 elements with LIMIT), you can consider it O(log(N)).",
      "since": "2.8.9",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "max",
          "type": "string"
        },
        {
          "name": "min",
          "type": "string"
        },
        {
          "command": "LIMIT",
          "name": [
            "offset",
            "count"
          ],
          "optional": true,
          "type": [
            "integer",
            "integer"
          ]
        }
      ]
    },
    {
      "CmdName": "BZPOPMIN",
      "summary": "Remove and return the member with the lowest score from one or more sorted sets, or block until one is available",
      "complexity": "O(log(N)) with N being the number of elements in the sorted set.",
      "since": "5.0.0",
      "group": "sorted_set",
      "arguments": [
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        },
        {
          "name": "timeout",
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "BZPOPMAX",
      "summary": "Remove and return the member with the highest score from one or more sorted sets, or block until one is available",
      "complexity": "O(log(N)) with N being the number of elements in the sorted set.",
      "since": "5.0.0",
      "group": "sorted_set",
      "arguments": [
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        },
        {
          "name": "timeout",
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "ZRANGEBYLEX",
      "summary": "Return a range of members in a sorted set, by lexicographical range",
      "complexity": "O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements being returned. If M is constant (e.g. always asking for the first 10 elements with LIMIT), you can consider it O(log(N)).",
      "since": "2.8.9",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "min",
          "type": "string"
        },
        {
          "name": "max",
          "type": "string"
        },
        {
          "command": "LIMIT",
          "name": [
            "offset",
            "count"
          ],
          "optional": true,
          "type": [
            "integer",
            "integer"
          ]
        }
      ]
    },
    {
      "CmdName": "ZINCRBY",
      "summary": "Increment the score of a member in a sorted set",
      "complexity": "O(log(N)) where N is the number of elements in the sorted set.",
      "since": "1.2.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "increment",
          "type": "integer"
        },
        {
          "name": "member",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "ZREMRANGEBYRANK",
      "summary": "Remove all members in a sorted set within the given indexes",
      "complexity": "O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements removed by the operation.",
      "since": "2.0.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "start",
          "type": "integer"
        },
        {
          "name": "stop",
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "ZPOPMIN",
      "summary": "Remove and return members with the lowest scores in a sorted set",
      "complexity": "O(log(N)*M) with N being the number of elements in the sorted set, and M being the number of elements popped.",
      "since": "5.0.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "count",
          "optional": true,
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "ZREVRANGEBYSCORE",
      "summary": "Return a range of members in a sorted set, by score, with scores ordered from high to low",
      "complexity": "O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements being returned. If M is constant (e.g. always asking for the first 10 elements with LIMIT), you can consider it O(log(N)).",
      "since": "2.2.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "max",
          "type": "double"
        },
        {
          "name": "min",
          "type": "double"
        },
        {
          "enum": [
            "WITHSCORES"
          ],
          "name": "withscores",
          "optional": true,
          "type": "enum"
        },
        {
          "command": "LIMIT",
          "name": [
            "offset",
            "count"
          ],
          "optional": true,
          "type": [
            "integer",
            "integer"
          ]
        }
      ]
    },
    {
      "CmdName": "ZRANGEBYSCORE",
      "summary": "Return a range of members in a sorted set, by score",
      "complexity": "O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements being returned. If M is constant (e.g. always asking for the first 10 elements with LIMIT), you can consider it O(log(N)).",
      "since": "1.0.5",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "min",
          "type": "double"
        },
        {
          "name": "max",
          "type": "double"
        },
        {
          "enum": [
            "WITHSCORES"
          ],
          "name": "withscores",
          "optional": true,
          "type": "enum"
        },
        {
          "command": "LIMIT",
          "name": [
            "offset",
            "count"
          ],
          "optional": true,
          "type": [
            "integer",
            "integer"
          ]
        }
      ]
    },
    {
      "CmdName": "ZCARD",
      "summary": "Get the number of members in a sorted set",
      "complexity": "O(1)",
      "since": "1.2.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "ZREMRANGEBYSCORE",
      "summary": "Remove all members in a sorted set within the given scores",
      "complexity": "O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements removed by the operation.",
      "since": "1.2.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "min",
          "type": "double"
        },
        {
          "name": "max",
          "type": "double"
        }
      ]
    },
    {
      "CmdName": "ZSCORE",
      "summary": "Get the score associated with the given member in a sorted set",
      "complexity": "O(1)",
      "since": "1.2.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "member",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "ZINTERSTORE",
      "summary": "Intersect multiple sorted sets and store the resulting sorted set in a new key",
      "complexity": "O(N*K)+O(M*log(M)) worst case with N being the smallest input sorted set, K being the number of input sorted sets and M being the number of elements in the resulting sorted set.",
      "since": "2.0.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "destination",
          "type": "key"
        },
        {
          "name": "numkeys",
          "type": "integer"
        },
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        },
        {
          "command": "WEIGHTS",
          "name": "weight",
          "optional": true,
          "type": "integer",
          "variadic": true
        },
        {
          "command": "AGGREGATE",
          "enum": [
            "SUM",
            "MIN",
            "MAX"
          ],
          "name": "aggregate",
          "optional": true,
          "type": "enum"
        }
      ]
    },
    {
      "CmdName": "ZUNIONSTORE",
      "summary": "Add multiple sorted sets and store the resulting sorted set in a new key",
      "complexity": "O(N)+O(M log(M)) with N being the sum of the sizes of the input sorted sets, and M being the number of elements in the resulting sorted set.",
      "since": "2.0.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "destination",
          "type": "key"
        },
        {
          "name": "numkeys",
          "type": "integer"
        },
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        },
        {
          "command": "WEIGHTS",
          "name": "weight",
          "optional": true,
          "type": "integer",
          "variadic": true
        },
        {
          "command": "AGGREGATE",
          "enum": [
            "SUM",
            "MIN",
            "MAX"
          ],
          "name": "aggregate",
          "optional": true,
          "type": "enum"
        }
      ]
    },
    {
      "CmdName": "ZCOUNT",
      "summary": "Count the members in a sorted set with scores within the given values",
      "complexity": "O(log(N)) with N being the number of elements in the sorted set.",
      "since": "2.0.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "min",
          "type": "double"
        },
        {
          "name": "max",
          "type": "double"
        }
      ]
    },
    {
      "CmdName": "ZRANGE",
      "summary": "Return a range of members in a sorted set, by index",
      "complexity": "O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements returned.",
      "since": "1.2.0",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "start",
          "type": "integer"
        },
        {
          "name": "stop",
          "type": "integer"
        },
        {
          "enum": [
            "WITHSCORES"
          ],
          "name": "withscores",
          "optional": true,
          "type": "enum"
        }
      ]
    },
    {
      "CmdName": "ZREMRANGEBYLEX",
      "summary": "Remove all members in a sorted set between the given lexicographical range",
      "complexity": "O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements removed by the operation.",
      "since": "2.8.9",
      "group": "sorted_set",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "min",
          "type": "string"
        },
        {
          "name": "max",
          "type": "string"
        }
      ]
    }
  ],
  "string": [
    {
      "CmdName": "SETNX",
      "summary": "Set the value of a key, only if the key does not exist",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "value",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "GETSET",
      "summary": "Set the string value of a key and return its old value",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "value",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "MSETNX",
      "summary": "Set multiple keys to multiple values, only if none of the keys exist",
      "complexity": "O(N) where N is the number of keys to set.",
      "since": "1.0.1",
      "group": "string",
      "arguments": [
        {
          "multiple": true,
          "name": [
            "key",
            "value"
          ],
          "type": [
            "key",
            "string"
          ]
        }
      ]
    },
    {
      "CmdName": "DECR",
      "summary": "Decrement the integer value of a key by one",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "MSET",
      "summary": "Set multiple keys to multiple values",
      "complexity": "O(N) where N is the number of keys to set.",
      "since": "1.0.1",
      "group": "string",
      "arguments": [
        {
          "multiple": true,
          "name": [
            "key",
            "value"
          ],
          "type": [
            "key",
            "string"
          ]
        }
      ]
    },
    {
      "CmdName": "GETRANGE",
      "summary": "Get a substring of the string stored at a key",
      "complexity": "O(N) where N is the length of the returned string. The complexity is ultimately determined by the returned length, but because creating a substring from an existing string is very cheap, it can be considered O(1) for small strings.",
      "since": "2.4.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "start",
          "type": "integer"
        },
        {
          "name": "end",
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "SETBIT",
      "summary": "Sets or clears the bit at offset in the string value stored at key",
      "complexity": "O(1)",
      "since": "2.2.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "offset",
          "type": "integer"
        },
        {
          "name": "value",
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "INCRBYFLOAT",
      "summary": "Increment the float value of a key by the given amount",
      "complexity": "O(1)",
      "since": "2.6.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "increment",
          "type": "double"
        }
      ]
    },
    {
      "CmdName": "SETRANGE",
      "summary": "Overwrite part of a string at key starting at the specified offset",
      "complexity": "O(1), not counting the time taken to copy the new string in place. Usually, this string is very small so the amortized complexity is O(1). Otherwise, complexity is O(M) with M being the length of the value argument.",
      "since": "2.2.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "offset",
          "type": "integer"
        },
        {
          "name": "value",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "SET",
      "summary": "Set the string value of a key",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "value",
          "type": "string"
        },
        {
          "enum": [
            "EX seconds",
            "PX milliseconds",
            "KEEPTTL"
          ],
          "name": "expiration",
          "optional": true,
          "type": "enum"
        },
        {
          "enum": [
            "NX",
            "XX"
          ],
          "name": "condition",
          "optional": true,
          "type": "enum"
        }
      ]
    },
    {
      "CmdName": "INCRBY",
      "summary": "Increment the integer value of a key by the given amount",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "increment",
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "MGET",
      "summary": "Get the values of all the given keys",
      "complexity": "O(N) where N is the number of keys to retrieve.",
      "since": "1.0.0",
      "group": "string",
      "arguments": [
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "GET",
      "summary": "Get the value of a key",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "BITCOUNT",
      "summary": "Count set bits in a string",
      "complexity": "O(N)",
      "since": "2.6.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": [
            "start",
            "end"
          ],
          "optional": true,
          "type": [
            "integer",
            "integer"
          ]
        }
      ]
    },
    {
      "CmdName": "GETBIT",
      "summary": "Returns the bit value at offset in the string value stored at key",
      "complexity": "O(1)",
      "since": "2.2.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "offset",
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "BITFIELD",
      "summary": "Perform arbitrary bitfield integer operations on strings",
      "complexity": "O(1) for each subcommand specified",
      "since": "3.2.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "command": "GET",
          "name": [
            "type",
            "offset"
          ],
          "optional": true,
          "type": [
            "type",
            "integer"
          ]
        },
        {
          "command": "SET",
          "name": [
            "type",
            "offset",
            "value"
          ],
          "optional": true,
          "type": [
            "type",
            "integer",
            "integer"
          ]
        },
        {
          "command": "INCRBY",
          "name": [
            "type",
            "offset",
            "increment"
          ],
          "optional": true,
          "type": [
            "type",
            "integer",
            "integer"
          ]
        },
        {
          "command": "OVERFLOW",
          "enum": [
            "WRAP",
            "SAT",
            "FAIL"
          ],
          "optional": true,
          "type": "enum"
        }
      ]
    },
    {
      "CmdName": "STRLEN",
      "summary": "Get the length of the value stored in a key",
      "complexity": "O(1)",
      "since": "2.2.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "PSETEX",
      "summary": "Set the value and expiration in milliseconds of a key",
      "complexity": "O(1)",
      "since": "2.6.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "milliseconds",
          "type": "integer"
        },
        {
          "name": "value",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "DECRBY",
      "summary": "Decrement the integer value of a key by the given number",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "decrement",
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "BITPOS",
      "summary": "Find first bit set or clear in a string",
      "complexity": "O(N)",
      "since": "2.8.7",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "bit",
          "type": "integer"
        },
        {
          "name": "start",
          "optional": true,
          "type": "integer"
        },
        {
          "name": "end",
          "optional": true,
          "type": "integer"
        }
      ]
    },
    {
      "CmdName": "SETEX",
      "summary": "Set the value and expiration of a key",
      "complexity": "O(1)",
      "since": "2.0.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "seconds",
          "type": "integer"
        },
        {
          "name": "value",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "STRALGO",
      "summary": "Run algorithms (currently LCS) against strings",
      "complexity": "For LCS O(strlen(s1)*strlen(s2))",
      "since": "6.0.0",
      "group": "string",
      "arguments": [
        {
          "enum": [
            "LCS"
          ],
          "name": "algorithm",
          "type": "enum"
        },
        {
          "multiple": true,
          "name": "algo-specific-argument",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "APPEND",
      "summary": "Append a value to a key",
      "complexity": "O(1). The amortized time complexity is O(1) assuming the appended value is small and the already present value is of any size, since the dynamic string library used by Redis will double the free space available on every reallocation.",
      "since": "2.0.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        },
        {
          "name": "value",
          "type": "string"
        }
      ]
    },
    {
      "CmdName": "BITOP",
      "summary": "Perform bitwise operations between strings",
      "complexity": "O(N)",
      "since": "2.6.0",
      "group": "string",
      "arguments": [
        {
          "name": "operation",
          "type": "string"
        },
        {
          "name": "destkey",
          "type": "key"
        },
        {
          "multiple": true,
          "name": "key",
          "type": "key"
        }
      ]
    },
    {
      "CmdName": "INCR",
      "summary": "Increment the integer value of a key by one",
      "complexity": "O(1)",
      "since": "1.0.0",
      "group": "string",
      "arguments": [
        {
          "name": "key",
          "type": "key"
        }
      ]
    }
  ]
}`
}
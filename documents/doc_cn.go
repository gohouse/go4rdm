package documents

func GetDocString() string {
	return `字符串
String 是redis最基本的类型，value 不仅可以是 String,也可以是数字。

使用 Strings 类型,可以完全实现目前 Memcached 的功能,并且效率更高。还可以享受 Redis 的定时持久化(可以选择 RDB 模式或者 AOF 模式).

string类型是二进制安全的。意思是redis的string可以包含任何数据,比如jpg图片或者序列化的对象

string类型是Redis最基本的数据类型，一个键最大能存储512MB。

命令示例：
set ­­ 设置key对应的值为string类型的value。

> set name itcast
setnx ­­ 将key设置值为value，如果key不存在，这种情况下等同SET命令。 当key存在时，什么也不做。SETNX是”SET if Not eXists”的简写。

> get name
"itcast"
> setnx name itcast_new
(integer)0
>get name
"itcast"
setex ­­ 设置key对应字符串value，并且设置key在给定的seconds时间之后超时过期。

> setex color 10 red 
> get color
"red"
10秒后...
> get color (nil)
setrange ­­ 覆盖key对应的string的一部分，从指定的offset处开始，覆盖value的长度。

127.0.0.1:6379> set email [email protected]
OK
127.0.0.1:6379> setrange email 13 gmail.com 
(integer) 22
127.0.0.1:6379> get email
"[email protected]"
127.0.0.1:6379>STRLEN email
(integer) 22
其中的4是指从下标为13(包含13)的字符开始替换

mset ­­ 一次设置多个key的值,成功返回ok表示所有的值都设置了,失败返回0表示没有任何值被设置。

> mset key1 python key2 c++
  OK
mget ­­ 一次获取多个key的值,如果对应key不存在,则对应返回nil。

> mget key1 key2 key3
  1) "python"   
  2) "c++"   
  3) (nil)
msetnx ­­ 对应给定的keys到他们相应的values上。只要有一个key已经存在，MSETNX一个操作都不会执行。

> MSETNX key11 "Hello" key22 "there"
(integer) 1
> MSETNX key22 "there" key33 "world"
(integer) 0
认证了：MSETNX是原子的，所以所有给定的keys是一次性set的

getset ­­ 设置key的值,并返回key的旧值

> get name
"itcast"
> getset name itcast_new
"itcast"
> get name
"itcast_new"
GETRANGE key start end ­­ 获取指定key的value值的子字符串。是由start和end位移决定的

> getrange name 0 4
  "itcas"
incr ­­ 对key的值加1操作

> set age 20 
> incr age 
(integer) 21
incrby ­­ 同incr类似,加指定值 ,key不存在时候会设置key,并认为原来的value是 0

> incrby age 5
  (integer) 26
> incrby age1111 5
(integer) 5
> get age1111
"5"
decr ­­ 对key的值做的是减减操作,decr一个不存在key,则设置key为­1

decrby ­­ 同decr,减指定值

append ­­ 给指定key的字符串值追加value,返回新字符串值的长度。例如我们向name的值追加一个"redis"字符串:

127.0.0.1:6379> get name
"itcast_new"
127.0.0.1:6379> append name "value"
(integer) 15
127.0.0.1:6379> get name
"itcast_newvalue"
127.0.0.1:6379>
`
}

func GetDocHash() string {
	return `HASH 哈希
Redis hash 是一个string类型的field和value的映射表，hash特别适合用于存储对象。

Redis 中每个 hash 可以存储 232 - 1 键值对（40多亿）。

示例
HSET key field value ­­ 设置 key 指定的哈希集中指定字段的值

> hset myhash field1 Hello
hget ­­ 获取指定的hash field。

> hget myhash field1  
 "Hello"
> hget myhash field3  
 (nil)
由于数据库没有field3,所以取到的是一个空值nil.

HSETNX key field value ­­ 只在 key 指定的哈希集中不存在指定的字段时，设置字段的值。如果 key 指定的哈希集不存在，会创建一个新的哈希集并与 key 关联。如果字段已存在，该操作无效果。

> hsetnx myhash field "Hello"   
(integer) 1
> hsetnx myhash field "Hello"   
(integer) 0
第一次执行是成功的,但第二次执行相同的命令失败,原因是field已经存在了。

hmset ­­ 同时设置hash的多个field。

> hmset myhash field1 Hello field2 World   
> OK
hmget ­­ 获取全部指定的hash filed。

> hmget myhash field1 field2 field3   
1) "Hello"
2) "World"
3) (nil)
hincrby ­­ 指定的hash filed 加上给定值。

> hset myhash field3 20  
 (integer) 1
> hget myhash field3  
 "20"
> hincrby myhash field3 -8   
(integer) 12
> hget myhash field3   
"12
hexists ­­ 测试指定field是否存在。

> hexists myhash field1
  (integer) 1
> hexists myhash field9
  (integer) 0   
  通过上例可以说明field1存在,但field9是不存在的。
hdel 从 key 指定的哈希集中移除指定的域

127.0.0.1:6379> hkeys myhash
1) "field1"
2) "field"
3) "field2"
4) "field3"
127.0.0.1:6379> hdel myhash field
(integer) 1
127.0.0.1:6379> hkeys myhash
1) "field1"
2) "field2"
3) "field3"
127.0.0.1:6379>
hlen ­­ 返回指定hash的field数量。

> hlen myhash
  (integer) 3
hkeys ­­ 返回hash的所有field。

> hkeys myhash   
> 1) "field2"   
> 2) "field"   
> 3) "field3"
说明这个hash中有3个field。

hvals ­­ 返回hash的所有value。

> hvals myhash   
1) "World"   
2)"Hello"   
3)"12"
说明这个hash中有3个field。

hgetall ­­ 获取某个hash中全部的filed及value。

> hgetall myhash   
1) "field2"   
2) "World" 
3) "field"  
4) "Hello"   
5) "field3"   
6) "12"
HSTRLEN -- 返回 hash指定field的value的字符串长度

127.0.0.1:6379> HSTRLEN myhash field1
(integer) 5`
}

func GetDocZset() string {
	return `Sorted Set 有序集合
Redis 有序集合和集合一样也是string类型元素的集合,且不允许重复的成员。

不同的是每个元素都会关联一个double类型的分数。redis正是通过分数来为集合中的成员进行从小到大的排序。

有序集合的成员是唯一的,但分数(score)却可以重复。

集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是O(1)。 集合中最大的成员数为 232 - 1 (4294967295, 每个集合可存储40多亿个成员)。

ZADD key score member
将所有指定成员添加到键为key有序集合（sorted set）里面

redis> ZADD myzset 1 "one"
(integer) 1
redis> ZADD myzset 1 "uno"
(integer) 1
redis> ZADD myzset 2 "two" 3 "three"
(integer) 2
redis> ZRANGE myzset 0 -1 WITHSCORES
1) "one"
2) "1"
3) "uno"
4) "1"
5) "two"
6) "2"
7) "three"
8) "3"
redis>
ZCOUNT key min max
返回有序集key中，score值在min和max之间(默认包括score值等于min或max)的成员

redis> ZADD myzset 1 "one"
(integer) 1
redis> ZADD myzset 2 "two"
(integer) 1
redis> ZADD myzset 3 "three"
(integer) 1
redis> ZCOUNT myzset -inf +inf
(integer) 3
redis> ZCOUNT myzset (1 3
(integer) 2
redis>
ZINCRBY key increment member
为有序集key的成员member的score值加上增量increment

redis> ZADD myzset 1 "one"
(integer) 1
redis> ZADD myzset 2 "two"
(integer) 1
redis> ZINCRBY myzset 2 "one"
"3"
redis> ZRANGE myzset 0 -1 WITHSCORES
1) "two"
2) "2"
3) "one"
4) "3"
redis>
应用场景
1.带有权重的元素,LOL游戏大区最强王者

2 排行榜

案例
斗地主大赛排名

初始比赛
127.0.0.1:6379> ZADD doudizhu_rank 0 "player1"
(integer) 1
127.0.0.1:6379> ZADD doudizhu_rank 0 "player2"
(integer) 1
127.0.0.1:6379> ZADD doudizhu_rank 0 "player3"
(integer) 1
比赛开始，经过n轮比赛，每次统计，类似计算如下所示
127.0.0.1:6379> ZINCRBY doudizhu_rank 3 player3
"3"
127.0.0.1:6379> ZINCRBY doudizhu_rank -1 player2
"-1"
127.0.0.1:6379> ZINCRBY doudizhu_rank -2 player1
"-2"
比赛结束，进行排名
127.0.0.1:6379> ZRANGE doudizhu_rank 0 -1
1) "player1"
2) "player2"
3) "player3"
127.0.0.1:6379> ZRANGE doudizhu_rank 0 -1 withscores
1) "player1"
2) "-2"
3) "player2"
4) "-1"
5) "player3"
6) "3"
逆序排序才对
127.0.0.1:6379> zrevrange doudizhu_rank 0 -1 withscores
1) "player3"
2) "3"
3) "player2"
4) "-1"
5) "player1"
6) "-2"
127.0.0.1:6379>`
}

func GetDocSet() string {
	return `Set 集合
Set 就是一个集合,集合的概念就是一堆不重复值的组合。利用 Redis 提供的 Set 数据结构,可以存储一些集合性的数据。

比如在 微博应用中,可以将一个用户所有的关注人存在一个集合中,将其所有粉丝存在一个集合。

因为 Redis 非常人性化的为集合提供了 求交集、并集、差集等操作, 那么就可以非常方便的实现如共同关注、共同喜好、二度好友等功能, 对上面的所有集合操作,你还可以使用不同的命令选择将结果返回给客户端还是存集到一个新的集合中。

SADD key member [member ...]
添加一个或多个指定的member元素到集合的 key中

redis> SADD myset "Hello"
(integer) 1
redis> SADD myset "World"
(integer) 1
redis> SADD myset "World"
(integer) 0
redis> SMEMBERS myset
1) "World"
2) "Hello"
redis>
SCARD key
返回集合存储的key的基数 (集合元素的数量).

redis> SADD myset "Hello"
(integer) 1
redis> SADD myset "World"
(integer) 1
redis> SCARD myset
(integer) 2
redis>
SDIFF key [key ...]
返回一个集合与给定集合的差集的元素.

redis> SADD key1 'a' 'b' 'c'
(integer) 1
redis> SADD key2 "c"
(integer) 1
redis> SADD key2 "d"
(integer) 1
redis> SADD key2 "e"
(integer) 1
redis> SDIFF key1 key2
1) "a"
2) "b"
redis>
应用场景
1.共同好友、二度好友

2.利用唯一性,可以统计访问网站的所有独立 IP

3.好友推荐的时候,根据 tag 求交集,大于某个 临界值 就可以推荐

示例
以王宝强和马蓉为例，求二度好友，共同好友，推荐系统

127.0.0.1:6379> sadd marong_friend 'songdan' 'wangsicong' 'songzhe'
(integer) 1
127.0.0.1:6379> SMEMBERS marong_friend
1) "songzhe"
2) "wangsicong"
3) "songdandan"
127.0.0.1:6379> sadd wangbaoqiang_friend 'dengchao' 'angelababy' 'songzhe'
(integer) 1

#求共同好友
127.0.0.1:6379> SINTER marong_friend wangbaoqiang_friend
1) "songzhe"

#推荐好友系统
127.0.0.1:6379> SDIFF marong_friend wangbaoqiang_friend
1) "wangsicong"
2) "songdandan"
127.0.0.1:6379>
`
}

func GetDocList() string {
	return `List 列表
Redis列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素导列表的头部（左边）或者尾部（右边）

一个列表最多可以包含 232 - 1 个元素 (4294967295, 每个列表超过40亿个元素)。

RPUSH key value [value ...]
向存于 key 的列表的尾部插入所有指定的值。如果 key 不存在，那么会创建一个空的列表然后再进行 push 操作。

redis> RPUSH mylist "hello"
(integer) 1
redis> RPUSH mylist "world"
(integer) 2
redis> LRANGE mylist 0 -1
1) "hello"
2) "world"
redis>
LPOP key
移除并且返回 key 对应的 list 的第一个元素。

redis> RPUSH mylist "one"
(integer) 1
redis> RPUSH mylist "two"
(integer) 2
redis> RPUSH mylist "three"
(integer) 3
redis> LPOP mylist
"one"
redis> LRANGE mylist 0 -1
1) "two"
2) "three"
redis>
LTRIM key start stop
修剪(trim)一个已存在的 list，这样 list 就会只包含指定范围的指定元素。

start 和 stop 都是由0开始计数的， 这里的 0 是列表里的第一个元素（表头），1 是第二个元素，以此类推。

例如： LTRIM foobar 0 2 将会对存储在 foobar 的列表进行修剪，只保留列表里的前3个元素。

start 和 end 也可以用负数来表示与表尾的偏移量，比如 -1 表示列表里的最后一个元素， -2 表示倒数第二个，等等。

应用场景:
1.取最新N个数据的操作

比如典型的取你网站的最新文章，通过下面方式，我们可以将最新的5000条评论的ID放在Redis的List集合中，并将超出集合部分从数据库获取

使用LPUSH latest.comments命令，向list集合中插入数据
插入完成后再用LTRIM latest.comments 0 5000命令使其永远只保存最近5000个ID
然后我们在客户端获取某一页评论时可以用下面的逻辑（伪代码）
FUNCTION get_latest_comments(start,num_items):
  id_list = redis.lrange("latest.comments",start,start+num_items-1)
  IF id_list.length < num_items
      id_list = SQL_DB("SELECT ... ORDER BY time LIMIT ...")
  END
  RETURN id_list
如果你还有不同的筛选维度，比如某个分类的最新N条，那么你可以再建一个按此分类的List，只存ID的话，Redis是非常高效的。
示例
取最新N个评论的操作

127.0.0.1:6379> lpush mycomment 100001
(integer) 1
127.0.0.1:6379> lpush mycomment 100002
(integer) 2
127.0.0.1:6379> lpush mycomment 100003
(integer) 3
127.0.0.1:6379> lpush mycomment 100004
(integer) 4
127.0.0.1:6379> LRANGE mycomment 0 -1
1) "100004"
2) "100003"
3) "100002"
4) "100001"
127.0.0.1:6379> LTRIM mycomment 0 1
OK
127.0.0.1:6379> LRANGE mycomment 0 -1
1) "100004"
2) "100003"
127.0.0.1:6379> lpush mycomment 100005
(integer) 3
127.0.0.1:6379> LRANGE mycomment 0 -1
1) "100005"
2) "100004"
3) "100003"
127.0.0.1:6379> LTRIM mycomment 0 1
OK
127.0.0.1:6379> LRANGE mycomment 0 -1
1) "100005"
2) "100004"
127.0.0.1:6379>
`
}
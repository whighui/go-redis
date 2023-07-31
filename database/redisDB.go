package database

// RedisDB redisServer->redisDB->(redisDB dict)->redisObject
type RedisDB struct {
	//redis5.0源码
	//typedef struct redisDb {
	//    dict *dict;                 /* The keyspace for this DB */
	//    dict *expires;              /* Timeout of keys with a timeout set */
	//    dict *blocking_keys;        /* Keys with clients waiting for data (BLPOP)*/
	//    dict *ready_keys;           /* Blocked keys that received a PUSH */
	//    dict *watched_keys;         /* WATCHED keys for MULTI/EXEC CAS */
	//    int id;                     /* Database ID */
	//    long long avg_ttl;          /* Average TTL, just for stats */
	//    list *defrag_later;         /* List of key names to attempt to defrag one by one, gradually. */
	//} redisDb;

	//因为就是平常写json序列化 导致常常写大写的字段结构

	dict map[string]RedisObject //
}

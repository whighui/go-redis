package database

//映射redis server.h文件

// encoding 编码类型
const (
	OBJ_ENCODING_RAW        = 0  /* Raw representation */
	OBJ_ENCODING_INT        = 1  /* Encoded as integer */
	OBJ_ENCODING_HT         = 2  /* Encoded as hash table */
	OBJ_ENCODING_ZIPMAP     = 3  /* Encoded as zipmap */
	OBJ_ENCODING_LINKEDLIST = 4  /* No longer used: old list encoding. */
	OBJ_ENCODING_ZIPLIST    = 5  /* Encoded as ziplist */
	OBJ_ENCODING_INTSET     = 6  /* Encoded as intset */
	OBJ_ENCODING_SKIPLIST   = 7  /* Encoded as skiplist */
	OBJ_ENCODING_EMBSTR     = 8  /* Embedded sds string encoding    emb str代表进一步压缩的字符串*/
	OBJ_ENCODING_QUICKLIST  = 9  /* Encoded as linked list of ziplists */
	OBJ_ENCODING_STREAM     = 10 /* Encoded as a radix tree of listpacks */
)

// valueType 对应数据结构体
const (
	OBJ_STRING = 0 /* String object. */
	OBJ_LIST   = 1 /* List object. */
	OBJ_SET    = 2 /* Set object. */
	OBJ_ZSET   = 3 /* Sorted set object. */
	OBJ_HASH   = 4 /* Hash object. */
	OBJ_MODULE = 5 /* Module object. */
	OBJ_STREAM = 6 /* Stream object. */
)

// RedisObject 原始redis object头部必须占用16个字节(4bit+4bit+24bit+4byte+8byte)
type RedisObject struct {
	valueType uint8 //key对应的value类型 字符串、列表、集合、有序集合、hash
	encoding  uint8 //编码类型 const中始终类型
	lru       uint16
	lfu       uint8
	refcount  uint32      //计数引用判断
	content   interface{} //存放encoding对应数据结构的内容
}

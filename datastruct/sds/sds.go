package sds

/**
细节点：
1.要实现二进制安全的字符串
2.要实现不同长度字符串长度
3.要考虑到redis object类型
*/

/**
redis 3.2字符串表示形式如下 可以实现二进制安全的柔性字符数
1.存在问题 不同字符串长度都需要占据int 4个byte  以及free 占4个byte 存在浪费空间
type sds struct{
	len int
	free int
	buf []byte
}

所以在redis5中进行改进 按照不同字符串长度来划分不同数据类型
*/

/*
	---------------------------------------------redis 5 sds正式实现-----------------------------------------------------
												 参照sds.h 和 sds.c  两个源文件
c语言知识点：
内存分配器 jemalloc/tcmalloc 等分配内存大小的单位都是 2、4、8、16、32、64 等等，为了能容纳一个完整的 embstr 对象，jemalloc 最少会分配 32 字节的空间，如果字符串再稍微长一点，那就是 64 字节的空间。
如果总体超出了 64 字节，Redis 认为它是一个大字符串，不再使用 emdstr 形式存储，而该用 raw 形式。
而在sds中都是采用raw字符串形式
*/

const SDS_MAX_PREALLOC = (1024 * 1024)

type SDS []int8

// SDSNewLen
// 1）创建空字符串时，SDS_TYPE_5被强制转换为SDS_TYPE_8。
// 2）长度计算时有“+1”操作，是为了算上结束符“\0”。
// 3）返回值是指向sds结构buf字段的指针
func SDSNewLen(init interface{}, initLen int) *SDS {

	//return sds偏移量 找到buf数组位置呢
	return nil
}

//sds sdsnewlen(const void *init, size_t initlen);
//sds sdsnew(const char *init);
//sds sdsempty(void);
//sds sdsdup(const sds s);
//void sdsfree(sds s);
//sds sdsgrowzero(sds s, size_t len);
//sds sdscatlen(sds s, const void *t, size_t len);
//sds sdscat(sds s, const char *t);
//sds sdscatsds(sds s, const sds t);
//sds sdscpylen(sds s, const char *t, size_t len);
//sds sdscpy(sds s, const char *t);
//
//sds sdscatvprintf(sds s, const char *fmt, va_list ap);
//#ifdef __GNUC__
//sds sdscatprintf(sds s, const char *fmt, ...)
//__attribute__((format(printf, 2, 3)));
//#else
//sds sdscatprintf(sds s, const char *fmt, ...);
//#endif
//
//sds sdscatfmt(sds s, char const *fmt, ...);
//sds sdstrim(sds s, const char *cset);
//void sdsrange(sds s, ssize_t start, ssize_t end);
//void sdsupdatelen(sds s);
//void sdsclear(sds s);
//int sdscmp(const sds s1, const sds s2);
//sds *sdssplitlen(const char *s, ssize_t len, const char *sep, int seplen, int *count);
//void sdsfreesplitres(sds *tokens, int count);
//void sdstolower(sds s);
//void sdstoupper(sds s);
//sds sdsfromlonglong(long long value);
//sds sdscatrepr(sds s, const char *p, size_t len);
//sds *sdssplitargs(const char *line, int *argc);
//sds sdsmapchars(sds s, const char *from, const char *to, size_t setlen);
//sds sdsjoin(char **argv, int argc, char *sep);
//sds sdsjoinsds(sds *argv, int argc, const char *sep, size_t seplen);
//
///* Low level functions exposed to the user API */
//sds sdsMakeRoomFor(sds s, size_t addlen);
//void sdsIncrLen(sds s, ssize_t incr);
//sds sdsRemoveFreeSpace(sds s);
//size_t sdsAllocSize(sds s);
//void *sdsAllocPtr(sds s);
//
///* Export the allocator used by SDS to the program using SDS.
// * Sometimes the program SDS is linked to, may use a different set of
// * allocators, but may want to allocate or free things that SDS will
// * respectively free or allocate. */
//void *sds_malloc(size_t size);
//void *sds_realloc(void *ptr, size_t size);
//void sds_free(void *ptr);

// SDS5 在c语言中本身不被使用字符串长度<2^5-1 被创建的时候也会转化为SDS8
type SDS5 struct {
}

type SDS8 struct {
}

type SDS16 struct {
}

type SDS32 struct {
}

type SDS64 struct {
}

// SDSNew
// 1）创建空字符串时，SDS_TYPE_5被强制转换为SDS_TYPE_8。
// 2）长度计算时有“+1”操作，是为了算上结束符“\0”。
// 3）返回值是指向sds结构buf字段的指针
//
//	sdsnewlen("abc",3);
func (s *SDS) SDSNew() *SDS {
	return nil
}

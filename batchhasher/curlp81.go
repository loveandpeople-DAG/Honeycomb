package batchhasher

const (
	CURLP81_HASH_LENGTH = 243
	CURLP81_ROUNDS      = 81
)

var (
	CURLP81 = NewBatchHasher(CURLP81_HASH_LENGTH, CURLP81_ROUNDS)
)

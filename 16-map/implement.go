package main

type _type struct{} // This is meaningless, only here to avoid errors.

const (
	maxKeySize = 128
	maxElemSize = 128
)

const (
	//...

	loadFactorNum = 13
	loadFactorDen = 2

	// ...
)

//func mapassign(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
//	...
//
//	if !h.growing() && (overLoadFator(h.count + 1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
//		hashGrow(t, h)
//		goto again
//	}
//	...
//
//}

type maptype struct {
	types      _type
	key        *_type
	elm        *_type
	bucket     *_type
	keysize    uint8
	elemsize   uint8
	bucketsize uint16
	flags      uint32
}

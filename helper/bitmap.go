package helper

import (
	"encoding/hex"
)

// 适合短长度str
func String2Bitmap(in string) ([]byte, error) {
	if len(in) == 0 {
		return nil, nil
	}

	if len(in)%2 == 1 {
		return hex.DecodeString(in + "0")
	} else {
		return hex.DecodeString(in)
	}
}

// a = 97
// A = 65
// 0 = 48
func String2BitmapV2(in string) ([]byte, error) {
	if len(in) == 0 {
		return nil, nil
	}

	m := make([]byte, (len(in)+1)/2)

	var idx int  // 映射到的byte的index
	var tmp byte // in[i]不允许修改
	for i := range in {
		idx = i / 2
		tmp = in[i]

		if tmp >= 'a' { // a->A
			tmp -= 32
		}

		if tmp >= 48 && tmp <= 57 { // 0~9
			tmp -= 48
		} else if tmp >= 65 && tmp <= 70 { // A~F
			tmp -= 55
		} else {
			return nil, hex.InvalidByteError(in[i])
		}

		//fmt.Printf("-- %d,%d,%d,%x,%b\n", i, idx, tmp, tmp, tmp)

		if i%2 == 0 {
			m[idx] |= tmp << 4
		} else {
			m[idx] |= tmp
		}
	}

	return m, nil
}

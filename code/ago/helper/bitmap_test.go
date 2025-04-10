package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
[11000001 11110000]
[11000001 11110000]
[11000001 11110010]
*/
func TestString2Bitmap(t *testing.T) {
	b0, err := String2Bitmap("")
	assert.Nil(t, err)
	assert.Equal(t, len(b0), 0)

	b3, err := String2Bitmap("C1F")
	assert.Nil(t, err)
	assert.Equal(t, len(b3), 2)
	fmt.Printf("%b\n", b3)

	b3_1, err := String2Bitmap("c1F")
	assert.Nil(t, err)
	assert.Equal(t, len(b3_1), 2)
	fmt.Printf("%b\n", b3_1)

	b4, err := String2Bitmap("C1F2")
	assert.Nil(t, err)
	assert.Equal(t, len(b4), 2)
	fmt.Printf("%b\n", b4)

	_, err = String2BitmapV2("=1F2")
	assert.NotNil(t, err)
	fmt.Println(err)
}

func TestString2BitmapV2(t *testing.T) {
	b0, err := String2BitmapV2("")
	assert.Nil(t, err)
	assert.Equal(t, len(b0), 0)

	b3, err := String2BitmapV2("C1F")
	assert.Nil(t, err)
	assert.Equal(t, len(b3), 2)
	fmt.Printf("%b\n", b3)

	b3_1, err := String2BitmapV2("c1F")
	assert.Nil(t, err)
	assert.Equal(t, len(b3_1), 2)
	fmt.Printf("%b\n", b3_1)

	b4, err := String2BitmapV2("C1F2")
	assert.Nil(t, err)
	assert.Equal(t, len(b4), 2)
	fmt.Printf("%b\n", b4)

	_, err = String2BitmapV2("=1F2")
	assert.NotNil(t, err)
	fmt.Println(err)
}

func TestBitmapExist(t *testing.T) {
	b4, err := String2BitmapV2("C1F2")
	assert.Nil(t, err)
	assert.Equal(t, len(b4), 2)
	fmt.Printf("%b\n", b4)

	// for i := 0; i < 8; i++ {
	// 	fmt.Printf("%d, %t\n", i, BitmapExist(b4, i))
	// }

	assert.True(t, BitmapExist(b4, 0))
	assert.False(t, BitmapExist(b4, 1))
	assert.True(t, BitmapExist(b4, 6))
	assert.True(t, BitmapExist(b4, 7))
	assert.False(t, BitmapExist(b4, 16))
}

func TestBitmapSet(t *testing.T) {
	b4 := make([]byte, 1)

	assert.True(t, BitmapSet(b4, 0))
	assert.False(t, BitmapSet(b4, 8))
	assert.True(t, BitmapSet(b4, 6))
	assert.True(t, BitmapSet(b4, 7))

	fmt.Printf("%b\n", b4)
}

func TestBitmapDel(t *testing.T) {
	b4, _ := String2BitmapV2("C1F2")
	fmt.Printf("%b\n", b4)
	assert.False(t, BitmapSet(b4, 16))

	assert.True(t, BitmapSet(b4, 3))
	fmt.Printf("%b\n", b4)

	assert.True(t, BitmapDel(b4, 3))
	fmt.Printf("%b\n", b4)
}

func TestBitmapExistV2(t *testing.T) {
	b4, err := String2BitmapV2("C1F2")
	assert.Nil(t, err)
	assert.Equal(t, len(b4), 2)
	fmt.Printf("%b\n", b4)

	// for i := 0; i < 8; i++ {
	// 	fmt.Printf("%d, %t\n", i, BitmapExist(b4, i))
	// }

	assert.True(t, BitmapExistV2(b4, 0))
	assert.True(t, BitmapExistV2(b4, 1))
	assert.False(t, BitmapExistV2(b4, 2))
	assert.True(t, BitmapExistV2(b4, 7))
	assert.False(t, BitmapExistV2(b4, 16))
}

func TestBitmapSetV2(t *testing.T) {
	b4 := make([]byte, 1)

	assert.True(t, BitmapSetV2(b4, 0))
	assert.False(t, BitmapSetV2(b4, 8))
	assert.True(t, BitmapSetV2(b4, 6))
	assert.True(t, BitmapSetV2(b4, 7))

	fmt.Printf("%b\n", b4)
}

func TestBitmapDelV2(t *testing.T) {
	b4, _ := String2BitmapV2("C1F2")
	fmt.Printf("%b\n", b4)
	assert.False(t, BitmapSetV2(b4, 16))

	assert.True(t, BitmapSetV2(b4, 3))
	fmt.Printf("%b\n", b4)

	assert.True(t, BitmapDelV2(b4, 3))
	fmt.Printf("%b\n", b4)
}

func TestBitmapRang(t *testing.T) {
	b4, _ := String2BitmapV2("C1F2")
	fmt.Printf("%b\n", b4)

	n := len(b4) * 8
	for i := 0; i < n; i++ {
		fmt.Println(i, BitmapExistV2(b4, i))
	}
}

package nekosrun

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math"
)

func ZeroPaddingString(length int) string {
	returnStr := ""
	for i := 0; i < length; i++ {
		returnStr += "\x00"
	}
	return returnStr
}

func s(a string, b bool) []uint32 {
	var (
		c, i     uint32
		v        []uint32
		v_length uint32
	)

	for len(a)%4 != 0 {
		a += "\x00"
	}
	c = uint32(len(a))
	v_length = uint32(len(a)) / 4

	v = make([]uint32, v_length)
	for i = 0; i < c; i += 4 {
		v[i>>2] = uint32(a[i]) | uint32(a[i+1])<<8 | uint32(a[i+2])<<16 | uint32(a[i+3])<<24
	}
	if b {
		v = append(v, c)
	}
	return v
}

func l(a []uint32, b bool) []uint8 {
	var (
		d, c, i   uint32
		returnStr bytes.Buffer
	)
	d = uint32(len(a))
	c = 0xffffffff & ((d - 1) << 2)

	if b {
		var m uint32 = a[d-1]
		if (m < c-3) || (m > c) {
			return nil
		}
		c = m
	}
	for i = 0; i < d; i++ {
		for j := 0; j < 4; j++ {
			this_char := byte(a[i] >> (j * 8) & 0xff)
			returnStr.Write([]uint8{this_char})
		}
	}
	return_buffer := returnStr.Bytes()
	if b {
		split_return_buffer := return_buffer[:c]
		return split_return_buffer
	} else {
		return return_buffer
	}

}

func xEncode(strs string, key string) []uint8 {
	if len(strs) == 0 {
		return []uint8{}
	}
	v := s(strs, true)
	k := s(key, false)
	for len(k) < 4 {
		k = append(k, 0)
	}

	var (
		n, y, q, z uint32
		c, m, e, d uint32
		p          int
	)
	n = uint32(len(v) - 1)
	z = v[n]
	y = v[0]
	c = 0x86014019 | 0x183639A0
	d = 0
	q = uint32(math.Floor(6.0 + 52.0/(float64(n+1))))
	for 0 < q {
		q--
		d = d + c&(0x8CE0D9BF|0x731F2640)
		e = d >> 2 & 3
		for p = 0; p < int(n); p++ {
			y = v[p+1]
			m = z>>5 ^ y<<2
			m += (y>>3 ^ z<<4) ^ (d ^ y)
			m += k[(p&3)^int(e)] ^ z
			z = v[p] + m&(0xEFB8D130|0x10472ECF)
			v[p] = z
		}
		y = v[0]
		m = z>>5 ^ y<<2
		m += (y>>3 ^ z<<4) ^ (d ^ y)
		m += k[(p&3)^int(e)] ^ z
		z = v[n] + m&(0xBB390742|0x44C6F8BD)
		v[n] = z
	}
	return l(v, false)
}

func HASH_HMAC(msg, key []uint8) string {
	hmac_obj := hmac.New(md5.New, key)
	hmac_obj.Write(msg)
	return hex.EncodeToString(hmac_obj.Sum([]byte("")))
}

func SHA1(data []uint8) string {
	sha1_obj := sha1.New()
	sha1_obj.Write(data)
	return hex.EncodeToString(sha1_obj.Sum([]byte("")))
}

func Arraycmp(a []uint32, b []uint32) string {
	if len(a) != len(b) {
		return "长度不相等"
	}
	for i := 0; i < len(a); i++ {
		if a[i]-b[i] != 0 {
			return fmt.Sprintf("位置[%d]不相同", i)
		}
	}
	return "相同"
}

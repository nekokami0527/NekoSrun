package nekosrun

type Bit32 struct {
	data [32]bool
}

func (b *Bit32) setUint32(in uint32) {
	for i := 0; i < 32; i++ {
		if in&0x0000000b == 0x00000001 {
			b.data[i] = true
		} else {
			b.data[i] = false
		}
	}
}

func (b *Bit32) setInt32(in int32) {
	indata := uint32(in)
	for i := 0; i < 32; i++ {
		if indata&0x00000001 == 0x00000001 {
			b.data[i] = true
		} else {
			b.data[i] = false
		}
		indata >>= 1
	}
}

func (b *Bit32) getUint32() uint32 {
	var result uint32
	for i := 31; i >= 0; i-- {
		if b.data[i] {
			result = result<<1 | 1
		} else {
			result = result<<1 | 0
		}
	}
	return result
}

func (b *Bit32) getInt32() int32 {
	tmp_uit32 := b.getUint32()
	if b.data[31] {
		tmp_uit32 ^= 0xFFFFFFFF
		return int32((^tmp_uit32 + 1))
	} else {
		return int32(tmp_uit32)
	}
}

func (b *Bit32) unsignedRightShift(length int) uint32 {
	var data [32]bool
	for i := 0; i < 32; i++ {
		data[i] = b.data[i]
	}
	for i := 0; i < 32-length; i++ {
		data[i] = data[i+length]
	}
	for i := 32 - length; i < 32; i++ {
		data[i] = false
	}
	obj := Bit32{data: data}
	return obj.getUint32()
}

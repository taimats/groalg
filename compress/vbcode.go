package compress

import "slices"

func CompressNums(nums []int) []byte {
	nums = subEach(nums)
	encoded := make([][]byte, len(nums))
	for _, n := range nums {
		vb := vbEncode(n)
		encoded = append(encoded, vb)
	}
	return slices.Concat(encoded...)
}

func vbEncode(num int) []byte {
	size := calcSize(num)
	vb := make([]byte, 0, size)
	for i := range size {
		bits := num >> (7 * i) & 0x7f
		vb = append(vb, byte(bits))
	}
	slices.Reverse(vb)
	vb[size-1] |= 0x80
	return vb
}

func vbDecode(vb []byte) int {
	slices.Reverse(vb)
	var num int
	for i, v := range vb {
		v &= 0x7f
		num |= int(v) << (7 * i)
	}
	return num
}

func calcSize(num int) (size int) {
	size = num / 128
	res := num % 128
	if res != 0 {
		size++
	}
	return size
}

func subEach(nums []int) []int {
	slices.Sort(nums)
	for i, n := range nums {
		if i-1 < 0 {
			continue
		}
		nums[i] = n - nums[i-1]
	}
	return nums
}

func addEach(nums []int) []int {
	for i, n := range nums {
		if i-1 < 0 {
			continue
		}
		nums[i] = n + nums[i-1]
	}
	return nums
}

package p43

/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */

// Test ...
func Test() string {
	num1 := "5435048749856437580437532348948423245456456456456151565614684441554451513518564843545435432432432432545843785723543214848948"
	num2 := "5435043755846658944804375843785723543284328540328740327483248484874987949874564445451456154464546514645164561315464646456456"

	return multiply(num1, num2)
}

// @lc code=start
func multiply(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	out := make([]byte, len(num1)*len(num2))
	for i := len(num2) - 1; i >= 0; i-- {
		piece := addSelf(num1, num2[i]-48)
		for j := 0; j < len(num2)-i-1; j++ {
			piece = shiftR(piece)
		}
		out = addInternal(out, piece)
	}
	return fromBytesR(out)
}

func addSelf(num1 string, times byte) []byte {
	if times == 0 {
		return []byte{0}
	}
	out := make([]byte, len(num1))
	int1 := toIntBytesR(num1)
	for i := 0; i < int(times); i++ {
		out = addInternal(out, int1)
	}
	return out
}

func addInternal(num1, num2 []byte) []byte {

	c := len(num1) - len(num2)
	for i := 0; i < c; i++ {
		num2 = append(num2, 0)
	}
	c = len(num2) - len(num1)
	for i := 0; i < c; i++ {
		num1 = append(num1, 0)
	}
	var carrier byte
	for i := 0; i < len(num1); i++ {
		num1[i] = num1[i] + num2[i] + carrier
		if num1[i] >= 10 {
			num1[i] = num1[i] % 10
			carrier = 1
		} else {
			carrier = 0
		}
	}
	if carrier == 1 {
		num1 = append(num1, carrier)
	}
	return num1
}

func toIntBytesR(num string) []byte {
	int1 := make([]byte, len(num))
	for i := range num {
		int1[i] = num[len(num)-i-1] - 48
	}
	return int1
}

func fromBytesR(array []byte) string {
	reverse(array)
	i := 0
	for i < len(array) {
		if array[i] == 0 {
			i++
			continue
		}
		break
	}
	if i >= len(array) {
		return "0"
	}
	for j := i; j < len(array); j++ {
		array[j] += 48
	}
	return string(array[i:])

}

func reverse(array []byte) {
	i, j := 0, len(array)-1
	for i < j {
		array[i], array[j] = array[j], array[i]
		i++
		j--
	}
}

func shiftR(array []byte) []byte {
	if array[len(array)-1] != 0 {
		array = append(array, 0)
	}
	for i := len(array) - 1; i >= 1; i-- {
		array[i] = array[i-1]
	}
	array[0] = 0
	return array
}

// @lc code=end

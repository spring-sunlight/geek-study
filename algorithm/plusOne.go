func plusOne(digits []int) []int {
    if digits[len(digits) - 1] != 9 {
    	digits[len(digits) - 1]++
    	return digits
    } else {
    	for i := len(digits) - 1; i >= 0; i-- {
    		if digits[i] == 9 {
    			digits[i] = 0
    		} else {
    			digits[i]++
    			return digits
    		}
    	}
    	if digits[0] == 0 {
    		digits = append([]int{1},digits...)
    	}
    }
    return digits
}
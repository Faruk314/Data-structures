package twopointers


func Palindrome(slice []int) bool {

	if len(slice) == 0 {
		return  false
	}

 
	left := 0

	right := len(slice) - 1
    

	for left < right {
        
		if slice[left] != slice[right] {
            return  false
		}

		left ++
		right --
	}



	return true
}
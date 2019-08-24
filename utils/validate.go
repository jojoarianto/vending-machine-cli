package service

func Validate(coin int64) bool {
	if coin % 10 == 0 {
		return true
	}
	return false
}

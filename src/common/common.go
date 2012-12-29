package common

func SubStr(str string, start int, end int) string{
	var s string
	c := []byte(str)
	for k, v := range c{
		if(k >= start && k <= end){
			s += string(v)
		}
	}
	return s
}
package split

import "strings"

func split(s, sep string) (ret []string) {
	idx := strings.Index(s, sep)
	for idx > -1 {
		if s[:idx]!=""{
			ret = append(ret, s[:idx])
		}
		s = s[idx+1:]
		idx = strings.Index(s, sep)
	}
	if s !=""{
		ret = append(ret, s)
	}

	return
}

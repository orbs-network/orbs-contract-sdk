package state

import "strings"

func Key(compositeKey string, postfix... string) []byte {
	if len(postfix) == 0 {
		return []byte(compositeKey)
	}

	buf := []string{compositeKey}
	buf = append(buf, postfix...)

	return []byte(strings.Join(buf, ""))
}

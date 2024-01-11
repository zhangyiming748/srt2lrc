package time

import (
	"fmt"
	"regexp"
	"strings"
)

func Srt2Lrc(in string) string {
	//00:00:56,120 --> 00:00:56,840
	input := "00:00:56,120 --> 00:00:56,840"
	pattern := `(\d{2}:\d{2}:\d{2},\d{3})`

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(input)

	if len(matches) > 1 {
		fmt.Println("匹配到的时间：", matches[1])
	} else {
		fmt.Println("未找到匹配的时间")
	}
	for i, v := range matches {
		fmt.Println(i, v)
	}
	return strings.Join([]string{"[", strings.Split(matches[1], ",")[0], "]"}, "")
}

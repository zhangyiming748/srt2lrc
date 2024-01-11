package main

import (
	"fmt"
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/srt2lrc/util"
	"github.com/zhangyiming748/srt2lrc/util/time"
	"io"
	"log/slog"
	"os"
	"strings"
)

func main() {
	folders := GetAllFolder.List(util.GetVal("root", "dir"))
	folders = append(folders, util.GetVal("root", "dir"))
	for _, folder := range folders {
		files := GetFileInfo.GetAllFileInfo(folder, "srt")
		for _, file := range files {
			if strings.Contains(file.PurgeName, "origin") {
				continue
			}
			trans(file.FullPath)
		}
	}
}

func trans(srt string) {
	before := util.ReadByLine(srt)
	after, _ := os.OpenFile(strings.Replace(srt, ".srt", ".lrc", 1), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	for i := 0; i < len(before); i += 4 {
		if i+3 > len(before) {
			continue
		}
		//after.WriteString(fmt.Sprintf("%s\n", before[i]))   // 序号
		nextTime := time.Srt2Lrc(before[i+1]) //[hh:mm:ss]
		//after.WriteString(fmt.Sprintf("%s\n", before[i+1])) // 时间轴
		nextLrc := before[i+2]
		//after.WriteString(fmt.Sprintf("%s\n", before[i+2])) // 字幕
		line := strings.Join([]string{nextTime, nextLrc}, "")
		//after.WriteString(fmt.Sprintf("%s\n", before[i+3])) // 空行
		after.WriteString(fmt.Sprintf("%s\n", line))
		after.Sync()
	}
}
func setLog() {
	opt := slog.HandlerOptions{ // 自定义option
		AddSource: true,
		Level:     slog.LevelDebug, // slog 默认日志级别是 info
	}
	file := "Process.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0770)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	slog.SetDefault(logger)
}

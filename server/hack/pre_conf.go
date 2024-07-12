package main

import (
	"bufio"
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

const (
	// 应用名称
	_APP_NAME = "hotgo"
	// 数据库前缀
	_DB_PREFIX = "swift"
)

var wg sync.WaitGroup

func main() {
	ctx := gctx.GetInitCtx()
	wg.Add(2)
	go func() {
		UpdateHackConfigFile(ctx)
		wg.Done()
	}()

	go func() {
		UpdateDockerFile(ctx)
		wg.Done()
	}()

	wg.Wait()
}

func UpdateHackConfigFile(ctx context.Context) {
	// hack配置文件
	pwd := gfile.Pwd()
	file := gfile.Join(pwd, "hack", "config.yaml")
	hackCfgContent := string(gfile.GetBytes(file))
	// g.Log().Info(ctx, hackCfgContent)
	re := regexp.MustCompile(`name:\s*(\S*)`)

	newText := re.ReplaceAllStringFunc(hackCfgContent, func(match string) string {
		return fmt.Sprintf(`name: "%v"`, _APP_NAME)
	})
	// g.Log().Info(ctx, newText)
	err := gfile.PutContents(file, newText)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	g.Log().Info(ctx, file, "已更新")
}

func UpdateDockerFile(ctx context.Context) {
	pwd := gfile.Pwd()
	inputFile := gfile.Join(pwd, "manifest", "docker", "Dockerfile")
	outputFile := gfile.Join(pwd, "manifest", "docker", "Dockerfile.new")
	output, err := gfile.Create(outputFile)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	defer output.Close()
	writer := bufio.NewWriter(output)
	err = gfile.ReadLines(inputFile, func(line string) error {
		if strings.Contains(line, "ADD ./temp/linux_amd64/") {
			line = fmt.Sprintf("ADD ./temp/linux_amd64/%v $WORKDIR/%v", _APP_NAME, _APP_NAME)
		}

		if strings.Contains(line, "RUN chmod +x $WORKDIR/") && !strings.HasSuffix(line, ".sh") {
			line = fmt.Sprintf("RUN chmod +x $WORKDIR/%v", _APP_NAME)
		}

		fmt.Fprintln(writer, line)
		return nil
	})
	if err != nil {
		g.Log().Fatal(ctx, "Error opening input file:", err)
		return
	}
	// 确保所有数据都写入文件
	writer.Flush()
	// 确定关闭，否则文件被占用无法重命名
	output.Close()
	err = gfile.Rename(outputFile, inputFile)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	g.Log().Info(ctx, inputFile, "已更新")
}

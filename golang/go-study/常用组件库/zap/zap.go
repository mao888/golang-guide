package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	// 先创建一个logger，然后调用各个级别的方法记录日志（Debug/Info/Error/Warn）
	logger := zap.NewExample(zap.Fields(
		zap.Int("port", 8899),
		zap.String("serverName", "service"),
	)) // Example适合用在测试代码中
	// zap底层 API 可以设置缓存，所以一般使用defer logger.Sync()将缓存同步到文件中。
	defer logger.Sync()

	// Zap提供了两种类型的日志记录器: Sugared Logger 和 Logger

	url := "http://example.org/api"
	logger.Info("failed to fetch URL",
		zap.Namespace("111"),
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	// zap也提供了便捷的方法SugarLogger，可以使用printf格式符的方式。
	sugar := logger.Sugar()
	// SugarLogger还支持以w结尾的方法，这种方式不需要先创建字段对象，直接将字段名和值依次放在参数中即可，如例子中的Infow
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	// 调用SugarLogger以f结尾的方法与fmt.Printf没什么区别，如例子中的Infof
	sugar.Infof("Failed to fetch URL: %s", url)
}

package config

import "testing"

func TestLoadMissingEnv(t *testing.T) {
	t.Setenv("LUMEN_POSTGRES_DSN", "")
	t.Setenv("LUMEN_REDIS_ADDR", "")

	err := Load()
	if err == nil {
		t.Fatal("缺少环境变量时必须返回错误")
	}
}

package main

// 配置文件结构体
type Config struct {
	Server struct {
		Address string `yaml:"Address"`
		Port    int64  `yaml:"Port"`
		Token   string `yaml:"Token"`
	} `yaml:"Server"`
}

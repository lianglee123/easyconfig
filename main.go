/*
easyconfig使用指南：
1. 优先从环境变量拉取配置， 环境变量 > yaml文件 > 代码默认
2. 使用：
type MyConfig struct {
}
config := &MyConfig
err := LoadConfig(config, envPrefix, yamlPath)

3. 支持的类型： map, slice, 普通字段


问题：
	如何支持多config？
	config.yaml中的logger配置应该具有一致性
	使用环境变量 ENV_PREFIX_ + CONFIG_PATH 设置配置文件的路径

*/

package main

func main() {

}

# cfg

本地文件配置，支持yaml、toml、properties

# 快速使用

> 配置文件准备
app.yaml
```yaml
app:
  name: console
  server:
    ip: 127.0.0.1
    port: 8989
    timeout: 60000
    cors:
      enable: true
    languages:
      - Ruby
      - Perl
      - Python

```
> 应用使用
```go
func main() {
	config, err := cfg.LoadConfig("./cfg", "app.yaml")
	if err != nil {
		g.Error(err)
	}
	ip := config.GetString("app.server.ip")
	port := config.GetInt("app.server.port")
	enable := config.GetBool("app.server.cors.enable")
	timeout := config.GetTime("app.server.timeout")
	languages := config.GetArray("app.server.languages")

	g.Info(ip)
	g.Info(port)
	g.Info(enable)
	g.Info(timeout)
	g.Info(languages)
	all, _ := config.GetAll()
	g.Info(all)
}
```
> 结果
```text
[cfg] 2023-08-19 18:37:07 INFO 127.0.0.1 
[cfg] 2023-08-19 18:37:07 INFO 8989 
[cfg] 2023-08-19 18:37:07 INFO true 
[cfg] 2023-08-19 18:37:07 INFO 60µs 
[cfg] 2023-08-19 18:37:07 INFO [Ruby Perl Python] 
[cfg] 2023-08-19 18:37:07 INFO map[app:map[name:console server:map[cors:map[enable:true] ip:127.0.0.1 languages:[Ruby Perl Python] port:8989 timeout:60000]]]
```
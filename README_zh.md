### 目录结构

- `cmd`：存放命令行应用的代码，例如 `main.go`。

- `config`：存放配置文件，例如 `config.yaml`。

- `internal`：存放项目内部的代码，不对外暴露。

- - `dao`：存放数据访问对象（Data Access Object）的代码。
  - `handler`：存放 HTTP 请求处理器的代码。
  - `middleware`：存放 HTTP 中间件的代码。
  - `model`：存放数据模型的代码。
  - `provider`：存放依赖注入的代码。
  - `server`：存放 HTTP 服务器以及路由注册的代码。
  - `service`：存放业务逻辑的代码。

- `pkg`：存放可重用的代码，对外暴露。

- - `config`：存放读取配置文件的代码。
  - `helper`：存放辅助函数的代码。
  - `http`：存放 HTTP 相关的代码。
  - `log`：存放日志相关的代码。

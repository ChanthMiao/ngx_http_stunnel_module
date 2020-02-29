# CURL测试辅助工具

此跨平台工具是用于协助开发者使用curl测试'ngx_http_stunnel_module'核心功能的辅助程序。

## 构建

go语言编译工具链和因特网连接是完成构建的必要条件。

- 下载第三方依赖

  ```shell
  go get -u github.com/pkg/errors
  ```

- 构建二进制文件

  ```shell
  cd tools
  go build -o calcu
  ```

## 用法

根据给定的key和系统时间生成'stunnel-token'字段的字符串值是此程序的唯一工作。

```shell
./tool/calcu -key ${uuid v4}
```

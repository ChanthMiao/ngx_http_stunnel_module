# Tool for CURL Test

This cross platform tool is a helper program for developers to test the core function of 'ngx_http_stunnel_module' with curl.

## Build

All you need are golang compiling toolchain and Internet access.

- Download the third part dependency

  ```shell
  go get -u github.com/pkg/errors
  ```

- Build the binary

  ```shell
  cd tools
  go build -o calcu
  ```

## Usage

The only thing it does is to generate the string value of 'stunnel-token' by given key and system time.

```shell
./tool/calcu -key ${uuid v4}
```

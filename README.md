# PHP concurrent URL fetcher (using FFI and Go)

This is a simple example of using PHP FFI to call Go functions.

## Requirements

1. build your own PHP binary with FFI enabled

```bash
wget https://www.php.net/distributions/php-8.2.1.tar.xz
tar -xvf php-8.2.1.tar.xz
cd php-8.2.1
#with ffi, curl, openssl
./configure --with-ffi --with-curl --with-openssl
make
```

2. install Go

```bash
wget https://golang.org/dl/go1.17.1.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.17.1.linux-amd64.tar.gz
```

## Usage

### Build

```bash
go build -o libutil.so -buildmode=c-shared util.go
```

or use `./buildGoLib.sh`



### Run
```bash
/$HOMEX/php-8.2.1/sapi/cli/php example.php 
/$HOMEX/php-8.2.1/sapi/cli/php http_client_example.php
```

## References

- [PHP: FFI - Manual](https://www.php.net/manual/en/book.ffi.php)
- [Creating shared libraries in Go](http://snowsyn.net/2016/09/11/creating-shared-libraries-in-go/)
- [A bit of PHP, Go, FFI and holiday spirit](https://blog.claudiupersoiu.ro/2019/12/23/a-bit-of-php-go-ffi-and-holiday-spirit/lang/en/)
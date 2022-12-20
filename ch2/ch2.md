# TCP、扫描器和代理

## 1 理解 TCP 的握手机制

## 2. 通过端口转发绕过防火墙

## 3. 编写一个 TCP 扫描器

### 3.1 测试端口可用性

一个基本的端口扫描器，仅扫描一个端口

dial/main.go

### 3.2 执行非并发扫描

tcp-scanner-slow/main.go

### 3.3 执行并发扫描

#### 1、"极速" 端口扫描器

tcp-scanner-too-fast/main.go

####  2、使用 WaitGroup 进行同步扫描

tcp-scanner-wg-too-fast/main.go

#### 3、使用人工池进行端口扫描

tcp-sync-scanner/main.go

#### 4、多通道通信

tcp-scanner-final/main.go

> 解析端口字符串的扫描器
>
> scanner-port-format/

## 4. 构造 TCP 代理

### 4.1 使用 io.Reader 和 io.Writer

io-example/main.go

copy-example/main.go

### 4.2 创建服务器

echo-server/main.go

### 4.3 通过创建带缓冲的监听器来改进代码

echo-server/main.go

### 4.4 代理一个 TCP 客户端

```go
func handle (src net.Conn) {
    dst, err := net.Dial("tcp", "joescatcam.website:80")
    if err != nil {
        log.Fatalln("Unable to connect to our unreachable host")
    }
    defer dst.Close()
    
    // 在 goroutine 中运行以防止 io.Copy 被阻塞
    go func() {
        // 将源的输出复制到目标
        if _, err := io.Copy(dst, src); err != nil {
            log.Fatalln(err)
        }
    }()
    // 将目标的输出复制回源
    if _, err := io.Copy(src, dst); err != nil {
        log.Fatalln(err)
    }
}

func main() {
    // 在本地端口 80 上监听
    listener, err := net.Listen("tcp", ":80")
    if err != nil {
        log.Fatalln("Unable to bind to port")
    }
    
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatalln("Unable to accept connection")
        }
        go handle(conn)
    }
}
```

### 4.5 复现 Netcat 命令执行

netcat-exec/main.go


# HTTP 客户端以及与工具的远程交互

## 1. Go 的 HTTP 基础知识

### 1.1 调用 HTTP API

basic/main.go

### 1.2 生成一个请求

basic/main.go

### 1.3 使用结构化响应解析

basic-parsing/main.go

## 2. 构建与 Shodan 交互的 HTTP 客户端

### 2.1 回顾构建 API 客户端的步骤

### 2.2 设计项目结构

shodan/

### 2.3 清理 API 调用

### 2.4 查询 Shodan 订阅情况

### 2.5 创建一个客户端

## 3. 与 Metasploit 交互

### 3.1 配置环境

metasploit-minimal/

### 3.2 定义目标

### 3.3 获取有效令牌

### 3.4 定义请求和响应方法

### 3.5 创建配置结构体和 RPC 方法

### 3.6 执行远程调用

### 3.7 创建使用程序

## 4. 使用 Bing Scraping 解析文档元数据

### 4.1 配置环境和规划

### 4.2 定义元数据包

bing-metadata/

### 4.3 把数据映射到结构体

### 4.4 使用 Bing 搜索和接收文件


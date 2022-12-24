# 滥用数据库和文件系统

## 1. 使用 Docker 设置数据库

### 1.1 安装 MongoDB 数据库并写入数据

db/seed-mongo.js

### 1.2 安装 PostgreSQL 和 MySQL 数据库并写入数据

db/seed-pg-mysql.sql

### 1.3 安装 Microsoft SQL Server 数据库并写入数据

## 2. 在 Go 中连接查询数据库

### 2.1 查询 MongoDB 数据库

db/mongo-connect/main.go

### 2.2 查询 SQL 数据库

db/mysql-connect/main.go

## 3. 构建数据库矿工

db/dbminer/dbminer.go

### 3.1 实现一个 MongoDB 数据库矿工

db/mongo/main.go

### 3.2 实现一个 MySQL 数据库矿工

db/mysql/main.go

## 4. 掠夺文件系统

filesystem/main.go
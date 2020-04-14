# 後端框架
## 模組
### 架構說明
總共分成四個層級 (如下圖)，依序為

- Model Layer
- Repository Layer
- Service Layer
- Delivery Layer

![](https://github.com/bxcodec/go-clean-arch/raw/master/clean-arch.png)

#### Model Layer
用於定義資料格式與欄位
#### Repository Layer
用於與資料庫進行存取的封裝方法
#### Service Layer
用於封裝商業邏輯的方法
#### Delivery Layer
用於封裝接口給外部使用者使用的方法

## 如何實作
### 建立模組
參數以下樹狀圖，可以複製 `startup` 模組進行實作

```
├── model
│   └── [模組名稱].go
└── module
    └── [模組名稱]
        ├── delivery
        │   ├── handler.go
        │   └── http
        │       ├── handler.go
        │       └── handler_test.go
        ├── repository
        │   ├── repository.go
        │   └── repository_test.go
        ├── repository.go
        ├── service
        │   ├── service.go
        │   └── service_test.go
        └── service.go

```

#### 1. 建立 Model
於 `model` 資料夾中建立與模組名稱相同的檔案，可參照 `demo` 範例
#### 2. 建立 Module
##### 2-1. Repository
於 `repository.go` 檔案中建立 `interface`後，再由 `repository` 資料夾中的 `repository.go` 來進行 implement，可參照 `demo` 範例
##### 2-2. Service
於 `service.go` 檔案中建立 `interface`後，再由 `service` 資料夾中的 `service.go` 來進行 implement，可參照 `demo` 範例
##### 2-1. delivery
於 `delivery` 資料夾的 `handler.go` 檔案中建立 `interface`，名稱規定如下：

- 名稱為 `服務 + Handler`，例如 `Http` 的服務則 interface 名稱定義為 `HttpHandler`； `gRPC` 的服務則定義為 `GRPCHandler`，以此類推。

建立 `interface` 後，再到各`服務`的資料夾底下建立`服務名稱`的檔案，如`Http`的服務則到 `http`的資料夾底下建立 `http.go`，建立完畢之後實作上一部所建立的 `interface`

### 加入模組至主程式
主程式為 `main.go`，在主程式中有四個地方需要加入，註解分別為

1. 建立 Table Schema (Module)
2. 建立 Repository (Module)
3. 建立 Service (Module)
4. 建立 Handler (Module)

以下分別說明如何加入

#### 1. 建立 Table Schema
##### 說明
透過 `orm.DatabaseORM.CheckTable` 方法進行資料表初始化，參數說明如下：

1. 第一個參數固定為 `false`
2. model 的 instance

##### 範例：

```
orm.DatabaseORM.CheckTable(false, model.Department{})
```
#### 2. 建立 Repository
##### 說明
透過 `自定義的 New` 方法進行建立 Repository，Repository 基本上都會將 `orm` 的實例當作參數

##### 範例：

```
demoRepo = repository.NewDemoRepository(orm.DatabaseORM.GetInstance())
```
#### 3. 建立 Service
##### 說明
透過 `自定義的 New` 方法進行建立 Service，Service 基本上都會將 `repository` 的實例當作參數

##### 範例：

```
demoService = service.NewDemoService(demoRepo)
```
#### 4. 建立 Handler
##### 說明
透過 `自定義的 New` 方法進行建立 Handler，Handler 基本上都會將 `gateway` 與 自定義的 `service` 的實例當作參數

##### 範例：

```
_ = NewDemoHandler(gateway, demoService)
```
----
上述設定完畢後，執行 `go run main.go` 即可以將 `server` 進行啟動
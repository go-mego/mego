# Mego



# 使用方式

## Mego 指令列

Mego 引擎提供了 `mego` 指令可用於終端機並快速建造和設置服務。

### 建立專案

進入一個空的資料夾時，可以

```bash
$ mego new my_service
成功建立：./mego.yml
```

###

```bash
$ mego install
已經更新：controller/todo.go
已經更新：transport.go
```

### 產生程式碼

透過指令建立和初始化專案後，需要透過 `mego generate` 指令來替此空專案設計構造。這會自動產生部分的框架程式碼，並且會將相關資訊自動寫入 `./mego.yml` 以供後續顯示資訊或者是功能使用。

```bash
$ mego g controller todo browse:get:update:create:delete
成功建立：controller/todo.go
已經更新：.mego.yml
```

```bash
$ mego g model todo id:int content:string
成功建立：model/todo.go
成功建立：model/todo_database.go
已經更新：.mego.yml
```

###

```bash
$ mego transport Todo
成功轉繼：GET    /todos    -> Todo.Browse
成功轉繼：GET    /todo/:id -> Todo.Get
成功轉繼：POST   /todo     -> Todo.Create
成功轉繼：PATCH  /todo/:id -> Todo.Update
成功轉繼：DELETE /todo/:id -> Todo.Delete
已經更新：transport.go
已經更新：.mego.yml
```

```bash
$ mego transport Todo.MyOwnMethod
成功轉繼：Todo.MyOwnMethod
已經更新：transport.go
已經更新：.mego.yml
```

```bash
$ mego transport Todo.Update POST /api/createTodo
成功轉繼：POST /api/createTodo -> Todo.Create
已經更新：transport.go
已經更新：.mego.yml
```

###

```bash
$ mego routes
╔═════════════╦══════════════╦══════════╦═════════════╗
║ Handler     │ Method       │ Pipeline │ URI Pattern ║
╠─────────────┼──────────────┼──────────┼─────────────╣
║ Todo.Browse │ get          │ web      │ /todos      ║
║             │ @            │ rpc      │             ║
╠─────────────┼──────────────┼──────────┼─────────────╣
║ Todo.Get    │ get          │ web      │ /todo/:id   ║
║             │ @            │ rpc      │             ║
╠─────────────┼──────────────┼──────────┼─────────────╣
║ Todo.Create │ @            │ rpc      │             ║
╠─────────────┼──────────────┼──────────┼─────────────╣
║ Todo.Update │ @            │ rpc      │             ║
╠─────────────┼──────────────┼──────────┼─────────────╣
║ Todo.Delete │ @            │ rpc      │             ║
║             │ @DeleteItNow │ rpc      │             ║
╠─────────────┼──────────────┼──────────┼─────────────╣
║             │              │ static   │ /*          ║
╚═════════════╩══════════════╩══════════╩═════════════╝
```

###

```bash
$ mego command myCommand
成功建立：commands.go
已經更新：commands.go
已經更新：.mego.yml
```

```bash
$ mego myCommand
```

###

```bash
$ mego update
已經更新：controller/todo.go
已經更新：transport.go
```
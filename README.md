# 📝 Loghs

Loghs is a leveled logger that is x2 faster than the stardard log package of Golang. Besides, it costs 0 allocs and 0 bytes thanks to [zerolog](https://github.com/rs/zerolog).

## 🚀 Installation

```go
go get github.com/mivinci/loghs
```

## 📖 Quick Start

Writing a simple info-level log to the stdout:

```go
loghs.Info("hello")
```

Or you can decide which field to be printed by creating a new instance of Logger:

```go
logger := loghs.New(os.Stdout)
logger.TimeUnix().Message("")
```

this would only print a current timestamp to the stdout. 

## 🦘Q&A

- Why not use a JSON encoder?

Cuz some open sources that support JSON output are fast enough. (Truth: I can't find a way that is faster and costs less than what [zerolog](https://github.com/rs/zerolog) (**currently the fastest**) uses, damn!)

- Then why you write this package?

Cuz I am crazy.
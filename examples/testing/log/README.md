
# log

```go
import github.com/projectbadger/brocade-go/log
```

webmail-tas-login Logging package

Sets up logging outputs

## Index

- [Debug(string)](#func-debugstring)
- [Error(string)](#func-errorstring)
- [Fatal(string)](#func-fatalstring)
- [Info(string)](#func-infostring)
- [Panic(string)](#func-panicstring)
- [SetupLogging() error](#func-setuplogging-error)
- [Trace(string)](#func-tracestring)
- [Warn(string)](#func-warnstring)

- [type Logger](#type-logger)
  - [NewLogger(ConfigLog) (Logger, error)](#func-newloggerconfiglog-logger-error)
  - [Debug(Context, string)](#func-logger-debugcontext-string)
  - [Error(Context, string)](#func-logger-errorcontext-string)
  - [Fatal(Context, string)](#func-logger-fatalcontext-string)
  - [FormatToString() string](#func-logger-formattostring-string)
  - [Info(Context, string)](#func-logger-infocontext-string)
  - [LogOut()](#func-logger-logout)
  - [LogWithFieldsOut(Level, Fields, string)](#func-logger-logwithfieldsoutlevel-fields-string)
  - [Panic(Context, string)](#func-logger-paniccontext-string)
  - [Warn(Context, string)](#func-logger-warncontext-string)

## func [Debug(string)](<log.go#L189>)

```go
func Debug(msg string, v ...interface{})
```
## func [Error(string)](<log.go#L183>)

```go
func Error(msg string, v ...interface{})
```
## func [Fatal(string)](<log.go#L207>)

```go
func Fatal(msg string, v ...interface{})
```
## func [Info(string)](<log.go#L171>)

```go
func Info(msg string, v ...interface{})
```
## func [Panic(string)](<log.go#L201>)

```go
func Panic(msg string, v ...interface{})
```
## func [SetupLogging() error](<log.go#L121>)

SetupLogging sets up the logging outputs.


```go
func SetupLogging() error
```
## func [Trace(string)](<log.go#L195>)

```go
func Trace(msg string, v ...interface{})
```
## func [Warn(string)](<log.go#L177>)

```go
func Warn(msg string, v ...interface{})
```


## type [Logger](<log.go#L21>)
```go
type Logger struct {
	Config config.ConfigLog
	// contains filtered or unexported fields
}
```

## func [NewLogger(ConfigLog) (Logger, error)](<log.go#L95>)

NewLogger returns a pointer to a new logger created
with a LoggerConfig.

Example:

```go
config := &LoggerConfig{Type: 1, Level: logrus.InfoLevel}
logger := NewLogger(cfg)

```


```go
func NewLogger(cfg config.ConfigLog) (*Logger, error)
```

## func (*Logger) [Debug(Context, string)](<log.go#L75>)

```go
func (l *Logger) Debug(ctx context.Context, msg string, v ...interface{})
```
## func (*Logger) [Error(Context, string)](<log.go#L71>)

```go
func (l *Logger) Error(ctx context.Context, msg string, v ...interface{})
```
## func (*Logger) [Fatal(Context, string)](<log.go#L83>)

```go
func (l *Logger) Fatal(ctx context.Context, msg string, v ...interface{})
```
## func (*Logger) [FormatToString() string](<log.go#L27>)

```go
func (l *Logger) FormatToString(v ...interface{}) string
```
## func (*Logger) [Info(Context, string)](<log.go#L63>)

```go
func (l *Logger) Info(ctx context.Context, msg string, v ...interface{})
```
## func (*Logger) [LogOut()](<log.go#L31>)

```go
func (l *Logger) LogOut(v ...interface{})
```
## func (*Logger) [LogWithFieldsOut(Level, Fields, string)](<log.go#L37>)

LogWithFIeldsOut logs to all outputs at the specified
level the provided formatted message.


```go
func (l *Logger) LogWithFieldsOut(level logrus.Level, fields logrus.Fields, msg string, v ...interface{})
```
## func (*Logger) [Panic(Context, string)](<log.go#L79>)

```go
func (l *Logger) Panic(ctx context.Context, msg string, v ...interface{})
```
## func (*Logger) [Warn(Context, string)](<log.go#L67>)

```go
func (l *Logger) Warn(ctx context.Context, msg string, v ...interface{})
```


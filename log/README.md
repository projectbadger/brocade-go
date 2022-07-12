
# log

```go
import brocade/log
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

## func [Debug(string)](<log.go#L188>)

```go
func Debug(msg string, v ...interface{})
```
## func [Error(string)](<log.go#L182>)

```go
func Error(msg string, v ...interface{})
```
## func [Fatal(string)](<log.go#L206>)

```go
func Fatal(msg string, v ...interface{})
```
## func [Info(string)](<log.go#L170>)

```go
func Info(msg string, v ...interface{})
```
## func [Panic(string)](<log.go#L200>)

```go
func Panic(msg string, v ...interface{})
```
## func [SetupLogging() error](<log.go#L120>)

SetupLogging sets up the logging outputs.


```go
func SetupLogging() error
```
## func [Trace(string)](<log.go#L194>)

```go
func Trace(msg string, v ...interface{})
```
## func [Warn(string)](<log.go#L176>)

```go
func Warn(msg string, v ...interface{})
```


## type [Logger](<log.go#L20>)
```go
type Logger struct {
	Config config.ConfigLog
	// contains filtered or unexported fields
}
```

## func [NewLogger(ConfigLog) (Logger, error)](<log.go#L94>)

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

## func (*Logger) [Debug(Context, string)](<log.go#L74>)

```go
func (l *Logger) Debug(ctx context.Context, msg string, v ...interface{})
```
## func (*Logger) [Error(Context, string)](<log.go#L70>)

```go
func (l *Logger) Error(ctx context.Context, msg string, v ...interface{})
```
## func (*Logger) [Fatal(Context, string)](<log.go#L82>)

```go
func (l *Logger) Fatal(ctx context.Context, msg string, v ...interface{})
```
## func (*Logger) [FormatToString() string](<log.go#L26>)

```go
func (l *Logger) FormatToString(v ...interface{}) string
```
## func (*Logger) [Info(Context, string)](<log.go#L62>)

```go
func (l *Logger) Info(ctx context.Context, msg string, v ...interface{})
```
## func (*Logger) [LogOut()](<log.go#L30>)

```go
func (l *Logger) LogOut(v ...interface{})
```
## func (*Logger) [LogWithFieldsOut(Level, Fields, string)](<log.go#L36>)

LogWithFIeldsOut logs to all outputs at the specified
level the provided formatted message.


```go
func (l *Logger) LogWithFieldsOut(level logrus.Level, fields logrus.Fields, msg string, v ...interface{})
```
## func (*Logger) [Panic(Context, string)](<log.go#L78>)

```go
func (l *Logger) Panic(ctx context.Context, msg string, v ...interface{})
```
## func (*Logger) [Warn(Context, string)](<log.go#L66>)

```go
func (l *Logger) Warn(ctx context.Context, msg string, v ...interface{})
```


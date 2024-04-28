package logger

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/containrrr/shoutrrr"
	"github.com/containrrr/shoutrrr/pkg/router"
	sls "github.com/innopals/sls-logrus-hook"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"github.com/tidwall/gjson"
	"github.com/tomas-qstarrs/boost/lumberjack"
)

type HookContextKey string

const (
	ContextFormatOptions HookContextKey = "FormatOptions"
)

var ErrHookNotFound = errors.New("hook not found")

type (
	HookGenerator    struct{}
	HookGenerateFunc func(string) (logrus.Hook, error)
)

var hookGeneratorMap = map[string]HookGenerateFunc{}

func init() {
	i := &HookGenerator{}
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	for index := 0; index < t.NumMethod(); index++ {
		method := t.Method(index)
		hookGeneratorMap[strings.ToLower(method.Name)] = func(s string) (logrus.Hook, error) {
			in := []reflect.Value{v, reflect.ValueOf(s)}
			out := method.Func.Call(in)

			if !out[1].IsNil() {
				return nil, out[1].Interface().(error)
			}

			return out[0].Interface().(logrus.Hook), nil
		}
	}
}

func NewHook(typ string, s string) (logrus.Hook, error) {
	if hookGenerator, ok := hookGeneratorMap[typ]; ok {
		return hookGenerator(s)
	}

	return nil, fmt.Errorf("%w: %s", ErrHookNotFound, typ)
}

var _ logrus.Hook = &FileHook{}

// FileHook stores the hook of rolling file appender.
type FileHook struct {
	// formatter logrus.Formatter
	logger        *lumberjack.Logger
	logLevels     *LogLevels
	formatOptions *FormatOptions
}

// Fire is called when a log event is fired.
func (h *FileHook) Fire(entry *logrus.Entry) error {
	var (
		line []byte
		err  error
	)

	if entry.Context == nil {
		entry.Context = context.WithValue(context.Background(), ContextFormatOptions, h.formatOptions)
	} else {
		entry.Context = context.WithValue(entry.Context, ContextFormatOptions, h.formatOptions)
	}

	line, err = entry.Bytes()
	if err != nil {
		return err
	}

	// Write the the logger
	_, err = h.logger.Write(line)
	if err != nil {
		return err
	}

	return nil
}

// Levels returns the available logging
func (h *FileHook) Levels() []logrus.Level {
	return h.logLevels.ToLogrus()
}

func (*HookGenerator) File(s string) (logrus.Hook, error) {
	logger := &lumberjack.Logger{
		Filename:   gjson.Get(s, "file").String(),     // {var} is replaced
		MaxSize:    int(gjson.Get(s, "size").Int()),   // megabytes
		MaxBackups: int(gjson.Get(s, "backup").Int()), // backup count
		MaxAge:     int(gjson.Get(s, "days").Int()),   // days
		Compress:   gjson.Get(s, "compress").Bool(),   // disabled by default
	}

	return &FileHook{logger, NewLogLevels(s), NewFormatOptions(s)}, nil
}

var stderr = colorable.NewColorableStderr()

var _ logrus.Hook = &StderrHook{}

// StderrHook is for stdout.
type StderrHook struct {
	writer.Hook
	logLevels     *LogLevels
	formatOptions *FormatOptions
}

func (h *StderrHook) Fire(entry *logrus.Entry) error {
	var (
		line []byte
		err  error
	)

	if entry.Context == nil {
		entry.Context = context.WithValue(context.Background(), ContextFormatOptions, h.formatOptions)
	} else {
		entry.Context = context.WithValue(entry.Context, ContextFormatOptions, h.formatOptions)
	}

	line, err = entry.Bytes()
	if err != nil {
		return err
	}

	_, err = h.Writer.Write(line)
	return err
}

// Levels returns the available logging
func (h *StderrHook) Levels() []logrus.Level {
	return h.logLevels.ToLogrus()
}

func (*HookGenerator) Stderr(s string) (logrus.Hook, error) {
	hook := writer.Hook{
		Writer:    stderr,
		LogLevels: logrus.AllLevels,
	}

	logLevels := NewLogLevels(s)
	formatOptions := NewFormatOptions(s)

	return &StderrHook{hook, logLevels, formatOptions}, nil
}

var stdout = colorable.NewColorableStdout()

var _ logrus.Hook = &StdoutHook{}

// StdoutHook is for stdout.
type StdoutHook struct {
	writer.Hook
	logLevels     *LogLevels
	formatOptions *FormatOptions
}

// Levels returns the available logging
func (h *StdoutHook) Levels() []logrus.Level {
	return h.logLevels.ToLogrus()
}

func (h *StdoutHook) Fire(entry *logrus.Entry) error {
	var (
		line []byte
		err  error
	)

	if entry.Context == nil {
		entry.Context = context.WithValue(context.Background(), ContextFormatOptions, h.formatOptions)
	} else {
		entry.Context = context.WithValue(entry.Context, ContextFormatOptions, h.formatOptions)
	}

	line, err = entry.Bytes()
	if err != nil {
		return err
	}

	_, err = h.Writer.Write(line)
	return err
}

func (*HookGenerator) Stdout(s string) (logrus.Hook, error) {
	hook := writer.Hook{
		Writer:    stdout,
		LogLevels: logrus.AllLevels,
	}

	return &StdoutHook{hook, NewLogLevels(s), NewFormatOptions(s)}, nil
}

var _ logrus.Hook = &TelegramHook{}

type TelegramHook struct {
	router        *router.ServiceRouter
	logLevels     *LogLevels
	formatOptions *FormatOptions
}

const telegramURL = "telegram://%s@telegram?chats=%s"

func (h *TelegramHook) Fire(entry *logrus.Entry) error {
	var (
		line []byte
		err  error
	)

	if entry.Context == nil {
		entry.Context = context.WithValue(context.Background(), ContextFormatOptions, h.formatOptions)
	} else {
		entry.Context = context.WithValue(entry.Context, ContextFormatOptions, h.formatOptions)
	}

	line, err = entry.Bytes()
	if err != nil {
		return err
	}

	errs := h.router.Send(string(line), nil)
	if len(errs) > 0 && errs[0] != nil {
		return errs[0]
	}

	return nil
}

func (h *TelegramHook) Levels() []logrus.Level {
	return h.logLevels.ToLogrus()
}

func (*HookGenerator) Telegram(s string) (logrus.Hook, error) {
	router, err := shoutrrr.CreateSender(fmt.Sprintf(telegramURL,
		gjson.Get(s, "token").String(), gjson.Get(s, "chat_id").String()))
	if err != nil {
		return nil, err
	}

	return &TelegramHook{router, NewLogLevels(s), NewFormatOptions(s)}, nil
}

type SlsHook struct {
	Hook          *sls.SlsLogrusHook
	logLevels     *LogLevels
	formatOptions *FormatOptions
}

func (h *SlsHook) Fire(entry *logrus.Entry) error {
	var (
		line []byte
		err  error
	)

	if entry.Context == nil {
		entry.Context = context.WithValue(context.Background(), ContextFormatOptions, h.formatOptions)
	} else {
		entry.Context = context.WithValue(entry.Context, ContextFormatOptions, h.formatOptions)
	}

	line, err = entry.Bytes()
	if err != nil {
		return err
	}

	entry.Message = string(line)

	return h.Hook.Fire(entry)
}

func (h *SlsHook) Levels() []logrus.Level {
	return h.logLevels.ToLogrus()
}

func (*HookGenerator) Sls(s string) (logrus.Hook, error) {
	fmt.Println(fmt.Sprintf("%s.%s.log.aliyuncs.com",
		gjson.Get(s, "project").String(),
		gjson.Get(s, "region").String(),
	),
		gjson.Get(s, "accesskey").String(),
		gjson.Get(s, "accesssecret").String(),
		gjson.Get(s, "logstore").String(),
		gjson.Get(s, "topic").String())
	slsLogrusHook, err := sls.NewSlsLogrusHook(
		fmt.Sprintf("%s.%s.log.aliyuncs.com",
			gjson.Get(s, "project").String(),
			gjson.Get(s, "region").String(),
		),
		gjson.Get(s, "accesskey").String(),
		gjson.Get(s, "accesssecret").String(),
		gjson.Get(s, "logstore").String(),
		gjson.Get(s, "topic").String(),
	)
	if err != nil {
		return nil, err
	}
	return &SlsHook{slsLogrusHook, NewLogLevels(s), NewFormatOptions(s)}, nil
}

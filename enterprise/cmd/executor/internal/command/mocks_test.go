// Code generated by go-mockgen 1.3.7; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package command

import (
	"context"
	"sync"

	executor "github.com/sourcegraph/sourcegraph/internal/executor"
)

// MockLogEntry is a mock implementation of the LogEntry interface (from the
// package
// github.com/sourcegraph/sourcegraph/enterprise/cmd/executor/internal/command)
// used for unit testing.
type MockLogEntry struct {
	// CloseFunc is an instance of a mock function object controlling the
	// behavior of the method Close.
	CloseFunc *LogEntryCloseFunc
	// CurrentLogEntryFunc is an instance of a mock function object
	// controlling the behavior of the method CurrentLogEntry.
	CurrentLogEntryFunc *LogEntryCurrentLogEntryFunc
	// FinalizeFunc is an instance of a mock function object controlling the
	// behavior of the method Finalize.
	FinalizeFunc *LogEntryFinalizeFunc
	// WriteFunc is an instance of a mock function object controlling the
	// behavior of the method Write.
	WriteFunc *LogEntryWriteFunc
}

// NewMockLogEntry creates a new mock of the LogEntry interface. All methods
// return zero values for all results, unless overwritten.
func NewMockLogEntry() *MockLogEntry {
	return &MockLogEntry{
		CloseFunc: &LogEntryCloseFunc{
			defaultHook: func() (r0 error) {
				return
			},
		},
		CurrentLogEntryFunc: &LogEntryCurrentLogEntryFunc{
			defaultHook: func() (r0 executor.ExecutionLogEntry) {
				return
			},
		},
		FinalizeFunc: &LogEntryFinalizeFunc{
			defaultHook: func(int) {
				return
			},
		},
		WriteFunc: &LogEntryWriteFunc{
			defaultHook: func([]byte) (r0 int, r1 error) {
				return
			},
		},
	}
}

// NewStrictMockLogEntry creates a new mock of the LogEntry interface. All
// methods panic on invocation, unless overwritten.
func NewStrictMockLogEntry() *MockLogEntry {
	return &MockLogEntry{
		CloseFunc: &LogEntryCloseFunc{
			defaultHook: func() error {
				panic("unexpected invocation of MockLogEntry.Close")
			},
		},
		CurrentLogEntryFunc: &LogEntryCurrentLogEntryFunc{
			defaultHook: func() executor.ExecutionLogEntry {
				panic("unexpected invocation of MockLogEntry.CurrentLogEntry")
			},
		},
		FinalizeFunc: &LogEntryFinalizeFunc{
			defaultHook: func(int) {
				panic("unexpected invocation of MockLogEntry.Finalize")
			},
		},
		WriteFunc: &LogEntryWriteFunc{
			defaultHook: func([]byte) (int, error) {
				panic("unexpected invocation of MockLogEntry.Write")
			},
		},
	}
}

// NewMockLogEntryFrom creates a new mock of the MockLogEntry interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockLogEntryFrom(i LogEntry) *MockLogEntry {
	return &MockLogEntry{
		CloseFunc: &LogEntryCloseFunc{
			defaultHook: i.Close,
		},
		CurrentLogEntryFunc: &LogEntryCurrentLogEntryFunc{
			defaultHook: i.CurrentLogEntry,
		},
		FinalizeFunc: &LogEntryFinalizeFunc{
			defaultHook: i.Finalize,
		},
		WriteFunc: &LogEntryWriteFunc{
			defaultHook: i.Write,
		},
	}
}

// LogEntryCloseFunc describes the behavior when the Close method of the
// parent MockLogEntry instance is invoked.
type LogEntryCloseFunc struct {
	defaultHook func() error
	hooks       []func() error
	history     []LogEntryCloseFuncCall
	mutex       sync.Mutex
}

// Close delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockLogEntry) Close() error {
	r0 := m.CloseFunc.nextHook()()
	m.CloseFunc.appendCall(LogEntryCloseFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the Close method of the
// parent MockLogEntry instance is invoked and the hook queue is empty.
func (f *LogEntryCloseFunc) SetDefaultHook(hook func() error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Close method of the parent MockLogEntry instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *LogEntryCloseFunc) PushHook(hook func() error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *LogEntryCloseFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func() error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *LogEntryCloseFunc) PushReturn(r0 error) {
	f.PushHook(func() error {
		return r0
	})
}

func (f *LogEntryCloseFunc) nextHook() func() error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *LogEntryCloseFunc) appendCall(r0 LogEntryCloseFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of LogEntryCloseFuncCall objects describing
// the invocations of this function.
func (f *LogEntryCloseFunc) History() []LogEntryCloseFuncCall {
	f.mutex.Lock()
	history := make([]LogEntryCloseFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// LogEntryCloseFuncCall is an object that describes an invocation of method
// Close on an instance of MockLogEntry.
type LogEntryCloseFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c LogEntryCloseFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c LogEntryCloseFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// LogEntryCurrentLogEntryFunc describes the behavior when the
// CurrentLogEntry method of the parent MockLogEntry instance is invoked.
type LogEntryCurrentLogEntryFunc struct {
	defaultHook func() executor.ExecutionLogEntry
	hooks       []func() executor.ExecutionLogEntry
	history     []LogEntryCurrentLogEntryFuncCall
	mutex       sync.Mutex
}

// CurrentLogEntry delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockLogEntry) CurrentLogEntry() executor.ExecutionLogEntry {
	r0 := m.CurrentLogEntryFunc.nextHook()()
	m.CurrentLogEntryFunc.appendCall(LogEntryCurrentLogEntryFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the CurrentLogEntry
// method of the parent MockLogEntry instance is invoked and the hook queue
// is empty.
func (f *LogEntryCurrentLogEntryFunc) SetDefaultHook(hook func() executor.ExecutionLogEntry) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// CurrentLogEntry method of the parent MockLogEntry instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *LogEntryCurrentLogEntryFunc) PushHook(hook func() executor.ExecutionLogEntry) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *LogEntryCurrentLogEntryFunc) SetDefaultReturn(r0 executor.ExecutionLogEntry) {
	f.SetDefaultHook(func() executor.ExecutionLogEntry {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *LogEntryCurrentLogEntryFunc) PushReturn(r0 executor.ExecutionLogEntry) {
	f.PushHook(func() executor.ExecutionLogEntry {
		return r0
	})
}

func (f *LogEntryCurrentLogEntryFunc) nextHook() func() executor.ExecutionLogEntry {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *LogEntryCurrentLogEntryFunc) appendCall(r0 LogEntryCurrentLogEntryFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of LogEntryCurrentLogEntryFuncCall objects
// describing the invocations of this function.
func (f *LogEntryCurrentLogEntryFunc) History() []LogEntryCurrentLogEntryFuncCall {
	f.mutex.Lock()
	history := make([]LogEntryCurrentLogEntryFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// LogEntryCurrentLogEntryFuncCall is an object that describes an invocation
// of method CurrentLogEntry on an instance of MockLogEntry.
type LogEntryCurrentLogEntryFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 executor.ExecutionLogEntry
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c LogEntryCurrentLogEntryFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c LogEntryCurrentLogEntryFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// LogEntryFinalizeFunc describes the behavior when the Finalize method of
// the parent MockLogEntry instance is invoked.
type LogEntryFinalizeFunc struct {
	defaultHook func(int)
	hooks       []func(int)
	history     []LogEntryFinalizeFuncCall
	mutex       sync.Mutex
}

// Finalize delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockLogEntry) Finalize(v0 int) {
	m.FinalizeFunc.nextHook()(v0)
	m.FinalizeFunc.appendCall(LogEntryFinalizeFuncCall{v0})
	return
}

// SetDefaultHook sets function that is called when the Finalize method of
// the parent MockLogEntry instance is invoked and the hook queue is empty.
func (f *LogEntryFinalizeFunc) SetDefaultHook(hook func(int)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Finalize method of the parent MockLogEntry instance invokes the hook at
// the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *LogEntryFinalizeFunc) PushHook(hook func(int)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *LogEntryFinalizeFunc) SetDefaultReturn() {
	f.SetDefaultHook(func(int) {
		return
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *LogEntryFinalizeFunc) PushReturn() {
	f.PushHook(func(int) {
		return
	})
}

func (f *LogEntryFinalizeFunc) nextHook() func(int) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *LogEntryFinalizeFunc) appendCall(r0 LogEntryFinalizeFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of LogEntryFinalizeFuncCall objects describing
// the invocations of this function.
func (f *LogEntryFinalizeFunc) History() []LogEntryFinalizeFuncCall {
	f.mutex.Lock()
	history := make([]LogEntryFinalizeFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// LogEntryFinalizeFuncCall is an object that describes an invocation of
// method Finalize on an instance of MockLogEntry.
type LogEntryFinalizeFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 int
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c LogEntryFinalizeFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c LogEntryFinalizeFuncCall) Results() []interface{} {
	return []interface{}{}
}

// LogEntryWriteFunc describes the behavior when the Write method of the
// parent MockLogEntry instance is invoked.
type LogEntryWriteFunc struct {
	defaultHook func([]byte) (int, error)
	hooks       []func([]byte) (int, error)
	history     []LogEntryWriteFuncCall
	mutex       sync.Mutex
}

// Write delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockLogEntry) Write(v0 []byte) (int, error) {
	r0, r1 := m.WriteFunc.nextHook()(v0)
	m.WriteFunc.appendCall(LogEntryWriteFuncCall{v0, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Write method of the
// parent MockLogEntry instance is invoked and the hook queue is empty.
func (f *LogEntryWriteFunc) SetDefaultHook(hook func([]byte) (int, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Write method of the parent MockLogEntry instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *LogEntryWriteFunc) PushHook(hook func([]byte) (int, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *LogEntryWriteFunc) SetDefaultReturn(r0 int, r1 error) {
	f.SetDefaultHook(func([]byte) (int, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *LogEntryWriteFunc) PushReturn(r0 int, r1 error) {
	f.PushHook(func([]byte) (int, error) {
		return r0, r1
	})
}

func (f *LogEntryWriteFunc) nextHook() func([]byte) (int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *LogEntryWriteFunc) appendCall(r0 LogEntryWriteFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of LogEntryWriteFuncCall objects describing
// the invocations of this function.
func (f *LogEntryWriteFunc) History() []LogEntryWriteFuncCall {
	f.mutex.Lock()
	history := make([]LogEntryWriteFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// LogEntryWriteFuncCall is an object that describes an invocation of method
// Write on an instance of MockLogEntry.
type LogEntryWriteFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 []byte
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 int
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c LogEntryWriteFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c LogEntryWriteFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// MockLogger is a mock implementation of the Logger interface (from the
// package
// github.com/sourcegraph/sourcegraph/enterprise/cmd/executor/internal/command)
// used for unit testing.
type MockLogger struct {
	// FlushFunc is an instance of a mock function object controlling the
	// behavior of the method Flush.
	FlushFunc *LoggerFlushFunc
	// LogFunc is an instance of a mock function object controlling the
	// behavior of the method Log.
	LogFunc *LoggerLogFunc
}

// NewMockLogger creates a new mock of the Logger interface. All methods
// return zero values for all results, unless overwritten.
func NewMockLogger() *MockLogger {
	return &MockLogger{
		FlushFunc: &LoggerFlushFunc{
			defaultHook: func() (r0 error) {
				return
			},
		},
		LogFunc: &LoggerLogFunc{
			defaultHook: func(string, []string) (r0 LogEntry) {
				return
			},
		},
	}
}

// NewStrictMockLogger creates a new mock of the Logger interface. All
// methods panic on invocation, unless overwritten.
func NewStrictMockLogger() *MockLogger {
	return &MockLogger{
		FlushFunc: &LoggerFlushFunc{
			defaultHook: func() error {
				panic("unexpected invocation of MockLogger.Flush")
			},
		},
		LogFunc: &LoggerLogFunc{
			defaultHook: func(string, []string) LogEntry {
				panic("unexpected invocation of MockLogger.Log")
			},
		},
	}
}

// NewMockLoggerFrom creates a new mock of the MockLogger interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockLoggerFrom(i Logger) *MockLogger {
	return &MockLogger{
		FlushFunc: &LoggerFlushFunc{
			defaultHook: i.Flush,
		},
		LogFunc: &LoggerLogFunc{
			defaultHook: i.Log,
		},
	}
}

// LoggerFlushFunc describes the behavior when the Flush method of the
// parent MockLogger instance is invoked.
type LoggerFlushFunc struct {
	defaultHook func() error
	hooks       []func() error
	history     []LoggerFlushFuncCall
	mutex       sync.Mutex
}

// Flush delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockLogger) Flush() error {
	r0 := m.FlushFunc.nextHook()()
	m.FlushFunc.appendCall(LoggerFlushFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the Flush method of the
// parent MockLogger instance is invoked and the hook queue is empty.
func (f *LoggerFlushFunc) SetDefaultHook(hook func() error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Flush method of the parent MockLogger instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *LoggerFlushFunc) PushHook(hook func() error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *LoggerFlushFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func() error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *LoggerFlushFunc) PushReturn(r0 error) {
	f.PushHook(func() error {
		return r0
	})
}

func (f *LoggerFlushFunc) nextHook() func() error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *LoggerFlushFunc) appendCall(r0 LoggerFlushFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of LoggerFlushFuncCall objects describing the
// invocations of this function.
func (f *LoggerFlushFunc) History() []LoggerFlushFuncCall {
	f.mutex.Lock()
	history := make([]LoggerFlushFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// LoggerFlushFuncCall is an object that describes an invocation of method
// Flush on an instance of MockLogger.
type LoggerFlushFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c LoggerFlushFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c LoggerFlushFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// LoggerLogFunc describes the behavior when the Log method of the parent
// MockLogger instance is invoked.
type LoggerLogFunc struct {
	defaultHook func(string, []string) LogEntry
	hooks       []func(string, []string) LogEntry
	history     []LoggerLogFuncCall
	mutex       sync.Mutex
}

// Log delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockLogger) Log(v0 string, v1 []string) LogEntry {
	r0 := m.LogFunc.nextHook()(v0, v1)
	m.LogFunc.appendCall(LoggerLogFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Log method of the
// parent MockLogger instance is invoked and the hook queue is empty.
func (f *LoggerLogFunc) SetDefaultHook(hook func(string, []string) LogEntry) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Log method of the parent MockLogger instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *LoggerLogFunc) PushHook(hook func(string, []string) LogEntry) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *LoggerLogFunc) SetDefaultReturn(r0 LogEntry) {
	f.SetDefaultHook(func(string, []string) LogEntry {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *LoggerLogFunc) PushReturn(r0 LogEntry) {
	f.PushHook(func(string, []string) LogEntry {
		return r0
	})
}

func (f *LoggerLogFunc) nextHook() func(string, []string) LogEntry {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *LoggerLogFunc) appendCall(r0 LoggerLogFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of LoggerLogFuncCall objects describing the
// invocations of this function.
func (f *LoggerLogFunc) History() []LoggerLogFuncCall {
	f.mutex.Lock()
	history := make([]LoggerLogFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// LoggerLogFuncCall is an object that describes an invocation of method Log
// on an instance of MockLogger.
type LoggerLogFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 string
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 []string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 LogEntry
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c LoggerLogFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c LoggerLogFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// MockCommandRunner is a mock implementation of the commandRunner interface
// (from the package
// github.com/sourcegraph/sourcegraph/enterprise/cmd/executor/internal/command)
// used for unit testing.
type MockCommandRunner struct {
	// RunCommandFunc is an instance of a mock function object controlling
	// the behavior of the method RunCommand.
	RunCommandFunc *CommandRunnerRunCommandFunc
}

// NewMockCommandRunner creates a new mock of the commandRunner interface.
// All methods return zero values for all results, unless overwritten.
func NewMockCommandRunner() *MockCommandRunner {
	return &MockCommandRunner{
		RunCommandFunc: &CommandRunnerRunCommandFunc{
			defaultHook: func(context.Context, command, Logger) (r0 error) {
				return
			},
		},
	}
}

// NewStrictMockCommandRunner creates a new mock of the commandRunner
// interface. All methods panic on invocation, unless overwritten.
func NewStrictMockCommandRunner() *MockCommandRunner {
	return &MockCommandRunner{
		RunCommandFunc: &CommandRunnerRunCommandFunc{
			defaultHook: func(context.Context, command, Logger) error {
				panic("unexpected invocation of MockCommandRunner.RunCommand")
			},
		},
	}
}

// surrogateMockCommandRunner is a copy of the commandRunner interface (from
// the package
// github.com/sourcegraph/sourcegraph/enterprise/cmd/executor/internal/command).
// It is redefined here as it is unexported in the source package.
type surrogateMockCommandRunner interface {
	RunCommand(context.Context, command, Logger) error
}

// NewMockCommandRunnerFrom creates a new mock of the MockCommandRunner
// interface. All methods delegate to the given implementation, unless
// overwritten.
func NewMockCommandRunnerFrom(i surrogateMockCommandRunner) *MockCommandRunner {
	return &MockCommandRunner{
		RunCommandFunc: &CommandRunnerRunCommandFunc{
			defaultHook: i.RunCommand,
		},
	}
}

// CommandRunnerRunCommandFunc describes the behavior when the RunCommand
// method of the parent MockCommandRunner instance is invoked.
type CommandRunnerRunCommandFunc struct {
	defaultHook func(context.Context, command, Logger) error
	hooks       []func(context.Context, command, Logger) error
	history     []CommandRunnerRunCommandFuncCall
	mutex       sync.Mutex
}

// RunCommand delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockCommandRunner) RunCommand(v0 context.Context, v1 command, v2 Logger) error {
	r0 := m.RunCommandFunc.nextHook()(v0, v1, v2)
	m.RunCommandFunc.appendCall(CommandRunnerRunCommandFuncCall{v0, v1, v2, r0})
	return r0
}

// SetDefaultHook sets function that is called when the RunCommand method of
// the parent MockCommandRunner instance is invoked and the hook queue is
// empty.
func (f *CommandRunnerRunCommandFunc) SetDefaultHook(hook func(context.Context, command, Logger) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// RunCommand method of the parent MockCommandRunner instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *CommandRunnerRunCommandFunc) PushHook(hook func(context.Context, command, Logger) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *CommandRunnerRunCommandFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, command, Logger) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *CommandRunnerRunCommandFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, command, Logger) error {
		return r0
	})
}

func (f *CommandRunnerRunCommandFunc) nextHook() func(context.Context, command, Logger) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *CommandRunnerRunCommandFunc) appendCall(r0 CommandRunnerRunCommandFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of CommandRunnerRunCommandFuncCall objects
// describing the invocations of this function.
func (f *CommandRunnerRunCommandFunc) History() []CommandRunnerRunCommandFuncCall {
	f.mutex.Lock()
	history := make([]CommandRunnerRunCommandFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// CommandRunnerRunCommandFuncCall is an object that describes an invocation
// of method RunCommand on an instance of MockCommandRunner.
type CommandRunnerRunCommandFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 command
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 Logger
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c CommandRunnerRunCommandFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c CommandRunnerRunCommandFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

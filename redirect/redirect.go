package redirect

import (
	"io"
	"os"

	"github.com/acoderup/boost/safe"
)

type Redirector struct {
	Options
}

type Options struct {
	duplicate bool
	writers   []io.Writer
}

type Option func(*Options)

var defaultOptions = Options{
	duplicate: false,
}

func WithDuplicate() Option {
	return func(o *Options) {
		o.duplicate = true
	}
}

func WithWriter(w io.Writer) Option {
	return func(o *Options) {
		o.writers = append(o.writers, w)
	}
}

func NewRedirector(opts ...Option) *Redirector {
	r := &Redirector{
		Options: defaultOptions,
	}

	for _, opt := range opts {
		opt(&r.Options)
	}

	return r
}

func (r *Redirector) Stdout(fn func()) error {
	originFile := os.Stdout // keep backup of the real file
	defer func() {          // Restore original file
		os.Stdout = originFile
	}()

	// Create pipe to create reader & writer
	pipeReader, pipeWriter, err := os.Pipe()
	if err != nil {
		return err
	}
	defer pipeWriter.Close()

	// Connect file to writer side of pipe
	os.Stdout = pipeWriter

	// Create MultiWriter to write to buffer and file at the same time
	writers := r.writers
	if r.duplicate {
		writers = append(writers, originFile)
	}
	multiWriter := io.MultiWriter(writers...)

	// copy the output in a separate goroutine so printing can't block indefinitely
	errCh := make(chan error, 1)
	go func() {
		if _, err := io.Copy(multiWriter, pipeReader); err != nil {
			errCh <- err
		}
		errCh <- nil
	}()

	if err := safe.Do(fn); err != nil {
		return err
	}

	if err := pipeWriter.Close(); err != nil {
		return err
	}

	if err := <-errCh; err != nil {
		return err
	}

	return nil
}

func (r *Redirector) Stderr(fn func()) error {
	originFile := os.Stderr // keep backup of the real file
	defer func() {          // Restore original file
		os.Stderr = originFile
	}()

	// Create pipe to create reader & writer
	pipeReader, pipeWriter, err := os.Pipe()
	if err != nil {
		return err
	}
	defer pipeWriter.Close()

	// Connect file to writer side of pipe
	os.Stderr = pipeWriter

	// Create MultiWriter to write to buffer and file at the same time
	writers := r.writers
	if r.duplicate {
		writers = append(writers, originFile)
	}
	multiWriter := io.MultiWriter(writers...)

	// copy the output in a separate goroutine so printing can't block indefinitely
	errCh := make(chan error, 1)
	go func() {
		if _, err := io.Copy(multiWriter, pipeReader); err != nil {
			errCh <- err
		}
		errCh <- nil
	}()

	if err := safe.Do(fn); err != nil {
		return err
	}

	if err := pipeWriter.Close(); err != nil {
		return err
	}

	if err := <-errCh; err != nil {
		return err
	}

	return nil
}

package melody

import "time"

// Config melody configuration struct.
type Options struct {
	writeWait         time.Duration // Milliseconds until write times out.
	pongWait          time.Duration // Timeout for waiting on pong.
	pingPeriod        time.Duration // Milliseconds between pings.
	maxMessageSize    int64         // Maximum size in bytes of a message.
	messageBufferSize int           // The max amount of messages that can be in a sessions buffer before it starts dropping them.
	readBufferSize    int           // The websocket read buffer size
	writeBufferSize   int           // The websocket write buffer size
}

// WithWriteWait returns a function that sets the writeWait field of the Options struct to the provided value.
func WithWriteWait(writeWait time.Duration) func(*Options) {
	return func(opts *Options) {
		opts.writeWait = writeWait
	}
}

// WithPongWait returns a function that sets the pongWait field of the Options struct to the provided value.
func WithPongWait(pongWait time.Duration) func(*Options) {
	return func(opts *Options) {
		opts.pongWait = pongWait
	}
}

// WithPingPeriod returns a function that sets the pingPeriod field of the Options struct to the provided value.
func WithPingPeriod(pingPeriod time.Duration) func(*Options) {
	return func(opts *Options) {
		opts.pingPeriod = pingPeriod
	}
}

// WithMaxMessageSize returns a function that sets the maxMessageSize field of the Options struct to the provided value.
func WithMaxMessageSize(maxMessageSize int64) func(*Options) {
	return func(opts *Options) {
		opts.maxMessageSize = maxMessageSize
	}
}

// WithMessageBufferSize returns a function that sets the messageBufferSize field of the Options struct to the provided value.
func WithMessageBufferSize(messageBufferSize int) func(*Options) {
	return func(opts *Options) {
		opts.messageBufferSize = messageBufferSize
	}
}

// WithReadBufferSize returns a function that sets the readBufferSize field of the Options struct to the provided value.
func WithReadBufferSize(readBufferSize int) func(*Options) {
	return func(opts *Options) {
		opts.readBufferSize = readBufferSize
	}
}

// WithWriteBufferSize returns a function that sets the writeBufferSize field of the Options struct to the provided value.
func WithWriteBufferSize(writeBufferSize int) func(*Options) {
	return func(opts *Options) {
		opts.writeBufferSize = writeBufferSize
	}
}

// newOptinos returns a pointer to an Options struct initialized with default values,
// where each field can be optionally overridden by calling the appropriate option function(s).
func newOptinos(options ...func(*Options)) *Options {
	opts := &Options{
		writeWait:         10 * time.Second,
		pongWait:          60 * time.Second,
		pingPeriod:        (60 * time.Second * 9) / 10,
		maxMessageSize:    512,
		messageBufferSize: 256,
		writeBufferSize:   1024,
		readBufferSize:    1024,
	}
	for _, opt := range options {
		opt(opts)
	}
	return opts
}

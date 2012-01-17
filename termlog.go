// Copyright (C) 2010, Kyle Lemons <kyle@kylelemons.net>.  All rights reserved.

package log4go

import (
	"io"
	"os"
	"fmt"
)

// This is the standard writer that prints to standard output.
type ConsoleLogWriter chan *LogRecord

// This creates a new ConsoleLogWriter
func NewConsoleLogWriter() ConsoleLogWriter {
	records := make(ConsoleLogWriter, LogBufferLength)
	go records.run(os.Stdout)
	return records
}

func (w ConsoleLogWriter) run(out io.Writer) {
	var timestr string
	var timestrAt int64

	for rec := range w {
		if rec.Created != timestrAt {
			tm := TimeConversionFunction(rec.Created / 1e9)
			timestr, timestrAt = tm.Format("15:04:05"), rec.Created/1e9
		}
		fmt.Fprint(out, "    ", levelStrings[rec.Level], " ", timestr, " ", rec.Prefix, ": ", rec.Message, "\n")
	}
}

// This is the ConsoleLogWriter's output method.  This will block if the output
// buffer is full.
func (w ConsoleLogWriter) LogWrite(rec *LogRecord) {
	w <- rec
}

// Close stops the logger from sending messages to standard output.  Attempts to
// send log messages to this logger after a Close have undefined behavior.
func (w ConsoleLogWriter) Close() {
	close(w)
}

package logger

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type log_formatter struct {
	TimestampFormat string
	Level           string
}

func (f *log_formatter) Format(entry *logrus.Entry) ([]byte, error) {
	// sss := logrus.JSONFormatter{}
	data := make(logrus.Fields, len(entry.Data))
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.RFC3339Nano
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteString(fmt.Sprintf("[%s] [%s] %s\n", strings.ToUpper(entry.Level.String()), entry.Time.Format(timestampFormat), entry.Message))
	return b.Bytes(), nil
}

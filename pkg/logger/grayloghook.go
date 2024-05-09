package logger

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"net"
)

type GraylogWriter struct {
	conn     *net.UDPConn
	hostname string
}

func NewGraylogWriter(graylogAddr, hostname string) (*GraylogWriter, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", graylogAddr)
	if err != nil {
		return nil, err
	}

	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return nil, err
	}

	return &GraylogWriter{
		conn:     udpConn,
		hostname: hostname,
	}, nil
}

func (w *GraylogWriter) Close() error {
	return w.conn.Close()
}

func (w *GraylogWriter) Write(p []byte) (int, error) {
	var vals map[string]interface{}
	err := json.Unmarshal(p, &vals)
	if err != nil {
		return 0, err
	}
	logLevel, ok := vals["level"]
	if !ok {
		vals["level"] = int32(zerolog.NoLevel)
	} else {
		logLevelStr, ok := logLevel.(string)
		if !ok {
			vals["level"] = int32(zerolog.NoLevel)
		} else {
			vals["level_str"] = logLevelStr
			lvlInt, _ := zerolog.ParseLevel(logLevelStr)
			vals["level"] = int32(lvlInt)
		}
	}
	vals["host"] = w.hostname

	_, ok = vals["message"]
	if !ok {
		vals["message"] = "EMPTY MESSAGE"
	}

	data, err := json.Marshal(vals)
	// Because we changed data it may have different from original length what may cause zerolog error "short write"
	toWriteLength := len(data)
	if err != nil {
		return 0, err
	}
	writtenLength, err := w.conn.Write(data)
	if err != nil {
		return 0, err
	}
	if writtenLength < toWriteLength {
		return writtenLength, nil
	}
	return len(p), nil

}

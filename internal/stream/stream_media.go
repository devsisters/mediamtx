package stream

import (
	"github.com/bluenviron/gortsplib/v4/pkg/description"
	"github.com/bluenviron/gortsplib/v4/pkg/format"
	"github.com/devsisters/mediamtx/internal/counterdumper"
	"github.com/devsisters/mediamtx/internal/logger"
)

type streamMedia struct {
	udpMaxPayloadSize  int
	media              *description.Media
	generateRTPPackets bool
	processingErrors   *counterdumper.CounterDumper
	parent             logger.Writer

	formats map[format.Format]*streamFormat
}

func (sm *streamMedia) initialize() error {
	sm.formats = make(map[format.Format]*streamFormat)

	for _, forma := range sm.media.Formats {
		sf := &streamFormat{
			udpMaxPayloadSize:  sm.udpMaxPayloadSize,
			format:             forma,
			generateRTPPackets: sm.generateRTPPackets,
			processingErrors:   sm.processingErrors,
			parent:             sm.parent,
		}
		err := sf.initialize()
		if err != nil {
			return err
		}
		sm.formats[forma] = sf
	}

	return nil
}

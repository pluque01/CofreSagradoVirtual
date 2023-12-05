package logger

import (
	"log"
	"os"

	"github.com/pluque01/CofreSagradoVirtual/internal/config"
	"github.com/rs/zerolog"
)

// Makes the logger available globally
var Default struct {
	Logger zerolog.Logger
	file   *os.File
}

func init() {
	var f *os.File
	var err error
	out := os.Stdout
	folder := config.DefaultConfig.LogFolder
	if folder != "" {
		f, err = os.OpenFile(
			folder+"validatecsv.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		if err != nil {
			log.Printf("Error opening log file: %v, using stdout", err)
		} else {
			out = f
		}
	} else {
		log.Printf("No log folder provided, using stdout")
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	Default.Logger = zerolog.New(out).With().Timestamp().Logger()
	Default.file = out
}

func Close() {
	Default.Logger.Info().Msg("Closing logger")
	Default.file.Close()
}

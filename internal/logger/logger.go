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
	var out *os.File
	folder := config.DefaultConfig.LogFolder
	f, err := os.OpenFile(
		folder+"validatecsv.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		log.Printf("Error opening log file: %v, using stdout", err)
		out = os.Stdout
	} else {
		out = f
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

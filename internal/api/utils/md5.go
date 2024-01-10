package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	log "github.com/pluque01/CofreSagradoVirtual/internal/logger"
)

func GetFileMd5(file io.Reader) (md5Str string) {
	h := md5.New()
	if _, err := io.Copy(h, file); err != nil {
		log.Default.Logger.Error().Err(err).Msg("Get file md5 error")
	}
	md5Str = hex.EncodeToString(h.Sum(nil))
	log.Default.Logger.Debug().Msgf("File md5 is: %s", md5Str)
	return md5Str
}

package libreLogger

import (
	"github.com/bruchte/libreConfig"
	"testing"
)

func TestInitialize(t *testing.T) {
	libreConfig.Initialize("testconfig.json")
	Initialize("libreLogger")

	logger := GetLogger("DEFAULT")
	logger.Debug("configured logger works")
	logger = GetLogger("********")
	logger.Info("default logger works")
	logger = GetLogger("MISSINGNAME")
	logger.Info("missing logger works with default")

}

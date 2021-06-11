package libreLogger

import (
	"github.com/bruchte/libreConfig"
	"github.com/bruchte/libreLogger/internal/core/service"
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

func TestLoggingLevels(t *testing.T) {
	logger := GetLogger("DATAGQL")
	logger.SetLoggingLevel("DEBUG")
	t.Logf("Should see all levels")
	writeall(logger)
	logger.SetLoggingLevel("INFO")
	t.Logf("Should see info+")
	writeall(logger)
	logger.SetLoggingLevel("WARN")
	t.Logf("Should see warn+")
	writeall(logger)
	logger.SetLoggingLevel("ERROR")
	t.Logf("Should see error only")
	writeall(logger)
}

func writeall(logger *service.LoggerLocalService) {
	logger.Error("error")
	logger.Warn("warn")
	logger.Info("info")
	logger.Debug("debug")
}

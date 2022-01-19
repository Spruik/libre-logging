package libreLogger

import (
	"fmt"
	"testing"

	libreConfig "github.com/Spruik/libre-configuration"
	"github.com/Spruik/libre-logging/interfaces"
	"github.com/Spruik/libre-logging/internal/core/service"
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

type SimpleVisitor struct {
	Msg string
}

func (obj *SimpleVisitor) LibreLoggerVisit(logger interfaces.LoggerLocalIF) {
	obj.Msg = fmt.Sprintf("visited logger has level = %s and destination = %s", logger.GetLevel(), logger.GetDestination())
}

func TestVisitor(t *testing.T) {
	libreConfig.Initialize("testconfig.json")
	var v = SimpleVisitor{}
	l := GetLogger("DEFAULT")
	l.AcceptVisitor(&v)
	t.Logf("visiting DEFAULT got: %s", v.Msg)
}

func TestVisitAll(t *testing.T) {
	var v = SimpleVisitor{}
	for name, logger := range service.LoggerMap {
		logger.AcceptVisitor(&v)
		t.Logf("visited %s found: %s", name, v.Msg)
	}
}

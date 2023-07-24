package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gin-gonic/gin"
	"github.com/samsonannan/prizepicks-assessment/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}
	config.Level.SetLevel(zapcore.InfoLevel)

	err := logger.SetGlobalLogger(&config)
	if err != nil {
		log.Fatalf("failed to build logger: %v", err)
	}
}

// Test GetCages function with successful queryfunc TestGetCagesWithInvalidStatus(t *testing.T) {
// Test with invalid status "ACT"
func TestGetCagesWithInvalidStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = httptest.NewRequest(http.MethodGet, "/cages?status=ACT", nil)
	GetCages(ctx)
	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

// Test with invalid status "ACTIVE123"
func TestGetCagesWithInvalidStatus2(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = httptest.NewRequest(http.MethodGet, "/cages?status=ACTIVE123", nil)
	GetCages(ctx)
	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

// Test with invalid status containing special characters
func TestGetCagesWithInvalidStatus5(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = httptest.NewRequest(http.MethodGet, "/cages?status=@#(@3@@@@", nil)
	GetCages(ctx)
	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

// Test with invalid status containing spaces
func TestGetCagesWithInvalidStatus6(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = httptest.NewRequest(http.MethodGet, "/cages?status=AACTIVE", nil)
	GetCages(ctx)
	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

// Test with invalid status "UP"
func TestGetCagesWithInvalidStatus7(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = httptest.NewRequest(http.MethodGet, "/cages?status=UP", nil)
	GetCages(ctx)
	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

// Test with invalid status "DOWN123"
func TestGetCagesWithInvalidStatus8(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = httptest.NewRequest(http.MethodGet, "/cages?status=DOWN123", nil)
	GetCages(ctx)
	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

// Test with invalid status " ACTIVE"
func TestGetCagesWithInvalidStatus9(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = httptest.NewRequest(http.MethodGet, "/cages?status=DOWNACTIVE", nil)
	GetCages(ctx)
	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

// Test with invalid status "DOWN "
func TestGetCagesWithInvalidStatus10(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = httptest.NewRequest(http.MethodGet, "/cages?status=ACTIVEDOWN", nil)
	GetCages(ctx)
	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

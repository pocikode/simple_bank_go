package api

import (
	"github.com/stretchr/testify/require"
	"os"
	db "pocikode/simple-bank/db/sqlc"
	"pocikode/simple-bank/util"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		TokenAccessDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

package api

import (
	"os"
	"testing"
	"time"

	db "github.com/Ma3au88/goSimpleBank/db/sqlc"
	"github.com/Ma3au88/goSimpleBank/util"
	"github.com/stretchr/testify/require"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

package dnbclient_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/h2non/gock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/struki84/dnbclient"
)

func TestGetToken(t *testing.T) {

	client, _ := dnbclient.NewClient(
		dnbclient.WithBaseURL(dnbclient.BaseURLV3),
		dnbclient.WithCredentials("test_username", "test_password"),
	)

	t.Run("Successful Get Token - unit test", func(t *testing.T) {
		defer gock.Off()

		gock.New(dnbclient.BaseURLV3).
			Post(dnbclient.AuthURL).
			Reply(http.StatusOK).
			JSON(map[string]any{"access_token": "test_token", "token_type": "Bearer", "expires_in": 3600})

		token, err := client.GetToken(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, "test_token", token)

		assert.True(t, gock.IsDone(), "Expected HTTP requests not made")
	})

	t.Run("Failed Get Token - unit test", func(t *testing.T) {
		defer gock.Off()

		gock.New(dnbclient.BaseURLV3).
			Post(dnbclient.AuthURL).
			Reply(http.StatusUnauthorized).
			JSON(map[string]string{"error": "invalid_request"})

		_, err := client.GetToken(context.Background())
		assert.Error(t, err)

		assert.True(t, gock.IsDone(), "Expected HTTP requests not made")
	})

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file make sure the file is present and placed in root directory.")
	}

	if username := os.Getenv("DNB_USERNAME"); username == "" {
		t.Skip("Skipping functional test because DNB_USERNAME is not set")
		return
	}

	if password := os.Getenv("DNB_PASSWORD"); password == "" {
		t.Skip("Skipping functional test because DNB_PASSWORD is not set")
		return
	}

	t.Run("Get Token - functional test", func(t *testing.T) {
		token, err := client.GetToken(
			context.Background(),
			dnbclient.WithCredentials(os.Getenv("DNB_USERNAME"), os.Getenv("DNB_PASSWORD")),
			dnbclient.WithBaseURL(dnbclient.BaseURLV3),
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})
}

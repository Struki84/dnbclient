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

	t.Run("Unit Test: Successful Get Token", func(t *testing.T) {
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

	t.Run("Unit Test: Failed Get Token", func(t *testing.T) {
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

func TestCriteriaSearch(t *testing.T) {

	client, _ := dnbclient.NewClient(
		dnbclient.WithCredentials("test_username", "test_password"),
	)

	t.Run("Unit Test: Successful Criteria Search", func(t *testing.T) {

		defer gock.Off()

		gock.New(dnbclient.BaseURLV1).
			Post(dnbclient.CriteriaSearchURL).
			Reply(http.StatusOK).
			JSON(map[string]any{"transactionDetail": map[string]string{"transactionID": "test_transactionID"}})

		searchResults, err := client.CriteriaSearch(
			context.Background(),
			dnbclient.WithCompanySerchRequest(&dnbclient.CompanySearchRequest{
				SearchTerm:     "test_search_term",
				TradeStyleName: "test_trade_style_name",
			}),
		)

		assert.NoError(t, err)
		assert.Equal(t, "test_transactionID", searchResults.TransactionDetail.TransactionID)

		assert.True(t, gock.IsDone(), "Expected HTTP requests not made")
	})

	t.Run("Unit Test: Failed Criteria Search", func(t *testing.T) {

		defer gock.Off()

		gock.New(dnbclient.BaseURLV1).
			Post(dnbclient.CriteriaSearchURL).
			Reply(http.StatusUnauthorized).
			JSON(map[string]string{"errorMessage": "invalid_request"})

		_, err := client.CriteriaSearch(context.Background())

		assert.Error(t, err)

		assert.True(t, gock.IsDone(), "Expected HTTP requests not made")
	})
}

func TestTypeheadSearch(t *testing.T) {

	client, _ := dnbclient.NewClient()

	t.Run("Unit Test: Successful Typehead Search", func(t *testing.T) {

		defer gock.Off()

		gock.New(dnbclient.BaseURLV1).
			Post(dnbclient.TypeheadSearchURL).
			Reply(http.StatusOK).
			JSON(map[string]any{"transactionDetail": map[string]string{"transactionID": "test_transactionID"}})

		searchResults, err := client.TypeheadSearch(
			context.Background(),
			"test_search",
			"test_country",
		)

		assert.NoError(t, err)
		assert.Equal(t, "test_transactionID", searchResults.TransactionDetail.TransactionID)

		assert.True(t, gock.IsDone(), "Expected HTTP requests not made")

	})

	t.Run("Unit Test: Failed Typehead Search", func(t *testing.T) {

		defer gock.Off()

		gock.New(dnbclient.BaseURLV1).
			Post(dnbclient.TypeheadSearchURL).
			Reply(http.StatusUnauthorized).
			JSON(map[string]string{"errorMessage": "invalid_request"})

		_, err := client.TypeheadSearch(
			context.Background(),
			"test_search",
			"test_country",
		)

		assert.Error(t, err)

		assert.True(t, gock.IsDone(), "Expected HTTP requests not made")
	})
}

func TestCompanyListSearch(t *testing.T) {

	client, _ := dnbclient.NewClient()

	t.Run("Unit Test: Failed Company List Search", func(t *testing.T) {
		defer gock.Off()

		gock.New(dnbclient.BaseURLV1).
			Post(dnbclient.CompanyListURL).
			Reply(http.StatusOK).
			JSON(map[string]any{"transactionDetail": map[string]string{"transactionID": "test_transactionID"}})

		searchResults, err := client.CompanyListSearch(
			context.Background(),
			dnbclient.WithCompanySerchRequest(&dnbclient.CompanySearchRequest{
				SearchTerm:     "test_search_term",
				TradeStyleName: "test_trade_style_name",
			}),
		)

		assert.NoError(t, err)
		assert.Equal(t, "test_transactionID", searchResults.TransactionDetail.TransactionID)

		assert.True(t, gock.IsDone(), "Expected HTTP requests not made")
	})

	t.Run("Unit Test: Failed Company List Search", func(t *testing.T) {

		defer gock.Off()

		gock.New(dnbclient.BaseURLV1).
			Post(dnbclient.CompanyListURL).
			Reply(http.StatusUnauthorized).
			JSON(map[string]string{"errorMessage": "invalid_request"})

		_, err := client.CompanyListSearch(
			context.Background(),
			dnbclient.WithCompanySerchRequest(&dnbclient.CompanySearchRequest{
				SearchTerm:     "test_search_term",
				TradeStyleName: "test_trade_style_name",
			}),
		)

		assert.Error(t, err)

		assert.True(t, gock.IsDone(), "Expected HTTP requests not made")
	})
}

func TestSearchContacts(t *testing.T) {

	client, _ := dnbclient.NewClient()

	t.Run("Unit Test: Successful Contact Search", func(t *testing.T) {

		defer gock.Off()

		gock.New(dnbclient.BaseURLV1).
			Post(dnbclient.ContactSearchURL).
			Reply(http.StatusOK).
			JSON(map[string]any{"transactionDetail": map[string]string{"transactionID": "test_transactionID"}})

		searchResults, err := client.SearchContact(
			context.Background(),
			dnbclient.WithContactSearchRequest(&dnbclient.ContactSearchRequest{
				ContactEmail: "test_contact_email",
			}),
		)

		assert.NoError(t, err)
		assert.Equal(t, "test_transactionID", searchResults.TransactionDetail.TransactionID)

		assert.True(t, gock.IsDone(), "Expected HTTP requests not made")
	})

	t.Run("Unit Test: Failed Contact Search", func(t *testing.T) {

		defer gock.Off()

		gock.New(dnbclient.BaseURLV1).
			Post(dnbclient.ContactSearchURL).
			Reply(http.StatusUnauthorized).
			JSON(map[string]string{"errorMessage": "invalid_request"})

		_, err := client.SearchContact(
			context.Background(),
			dnbclient.WithContactSearchRequest(&dnbclient.ContactSearchRequest{
				ContactEmail: "test_contact_email",
			}),
		)

		assert.Error(t, err)

		assert.True(t, gock.IsDone(), "Expected HTTP requests not made")
	})
}

func TestGetContactByDUNS(t *testing.T) {

	client, _ := dnbclient.NewClient()

	t.Run("Unit Test: Successful Get Contact By DUNS", func(t *testing.T) {

		defer gock.Off()

		gock.New(dnbclient.BaseURLV1).
			Post(dnbclient.ContactSearchURL).
			Reply(http.StatusOK).
			JSON(map[string]any{"transactionDetail": map[string]string{"transactionID": "test_transactionID"}})

		searchResults, err := client.GetContactByDUNS(context.Background(), "test_duns")

		assert.NoError(t, err)

		assert.Equal(t, "test_transactionID", searchResults.TransactionDetail.TransactionID)
	})

	t.Run("Unit Test: Failed Get Contact By DUNS", func(t *testing.T) {

		defer gock.Off()

		gock.New(dnbclient.BaseURLV1).
			Post(dnbclient.ContactSearchURL).
			Reply(http.StatusUnauthorized).
			JSON(map[string]string{"errorMessage": "invalid_request"})

		_, err := client.GetContactByDUNS(context.Background(), "test_duns")

		assert.Error(t, err)

		assert.True(t, gock.IsDone(), "Expected HTTP requests not made")

	})

}

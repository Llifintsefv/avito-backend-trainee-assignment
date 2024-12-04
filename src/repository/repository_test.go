// repository_test.go

package repository

import (
	"database/sql"
	"pro-backend-trainee-assignment/src/models"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)
	genValue := models.GenerateValue{
		ID:          "some-id",
		Value:       "some-value",
		Type:        "some-type",
		UserAgent:   "some-user-agent",
		RequestId:   123,
		Url:         "some-url",
		CountRequest: 5,
	}

    mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO random_values (guid,values, type,user_agent,requestid,url,countRequest) VALUES ($1, $2,$3,$4,$5,$6,$7)")).
        ExpectExec().
        WithArgs(genValue.ID, genValue.Value, genValue.Type, genValue.UserAgent, genValue.RequestId, genValue.Url, genValue.CountRequest).
        WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Generate(genValue)
	assert.NoError(t, err)
    assert.NoError(t,mock.ExpectationsWereMet())
}

func TestRetrieve(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    repo := NewRepository(db)
    id := "some-id"
    expectedValue := "some-value"
    expectedType := "some-type"

    rows := sqlmock.NewRows([]string{"values", "type"}).AddRow(expectedValue, expectedType)
    mock.ExpectPrepare(regexp.QuoteMeta("SELECT values,type FROM random_values WHERE guid = $1")).
        ExpectQuery().
        WithArgs(id).
        WillReturnRows(rows)

    value, Type, err := repo.Retrieve(id)
    assert.NoError(t, err)
    assert.Equal(t, expectedValue, value)
    assert.Equal(t, expectedType, Type)

     assert.NoError(t,mock.ExpectationsWereMet())
}
func TestRetrieveNotFound(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    repo := NewRepository(db)
    id := "some-id"

    rows := sqlmock.NewRows([]string{"values", "type"})
    mock.ExpectPrepare(regexp.QuoteMeta("SELECT values,type FROM random_values WHERE guid = $1")).
        ExpectQuery().
        WithArgs(id).
        WillReturnRows(rows)

    value, Type, err := repo.Retrieve(id)
    assert.NoError(t, err)
    assert.Equal(t, "", value)
    assert.Equal(t, "", Type)

     assert.NoError(t,mock.ExpectationsWereMet())
}

func TestGetCountRequest(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    repo := NewRepository(db)
    requestId := 123
    expectedCountRequest := 5

    rows := sqlmock.NewRows([]string{"countRequest"}).AddRow(expectedCountRequest)
    mock.ExpectQuery(regexp.QuoteMeta("SELECT countRequest FROM random_values WHERE requestid = $1")).
        WithArgs(requestId).
        WillReturnRows(rows)

    countRequest, err := repo.GetCountRequest(requestId)
    assert.NoError(t, err)
    assert.Equal(t, expectedCountRequest, countRequest)

     assert.NoError(t,mock.ExpectationsWereMet())
}

func TestGetCountRequestNotFound(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    repo := NewRepository(db)
    requestId := 123

    mock.ExpectQuery(regexp.QuoteMeta("SELECT countRequest FROM random_values WHERE requestid = $1")).
        WithArgs(requestId).
        WillReturnError(sql.ErrNoRows)

    countRequest, err := repo.GetCountRequest(requestId)
    assert.NoError(t, err)
    assert.Equal(t, 0, countRequest)

     assert.NoError(t,mock.ExpectationsWereMet())
}

func TestUpdateCountRequestAndRetrieveId(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    repo := NewRepository(db)
    requestId := 123
    countRequest := 5
    expectedId := "some-id"

    rows := sqlmock.NewRows([]string{"guid"}).AddRow(expectedId)
    mock.ExpectQuery(regexp.QuoteMeta("UPDATE random_values SET countRequest = $2 WHERE requestid = $1 RETURNING guid")).
        WithArgs(requestId, countRequest).
        WillReturnRows(rows)

    id, err := repo.UpdateCountRequestAndRetrieveId(requestId, countRequest)
    assert.NoError(t, err)
    assert.Equal(t, expectedId, id)

     assert.NoError(t,mock.ExpectationsWereMet())
}
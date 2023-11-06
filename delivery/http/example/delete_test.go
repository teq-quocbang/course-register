package example_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/teq-quocbang/course-register/delivery/http/example"
	"github.com/teq-quocbang/course-register/fixture/database"
	"github.com/teq-quocbang/course-register/repository"
	"github.com/teq-quocbang/course-register/usecase"
)

func TestDelete(t *testing.T) {
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := example.Route{UseCase: usecase.New(repo, nil)}

	t.Run("200", func(t *testing.T) {
		t.Run("Delete", func(t *testing.T) {
			// prepare data for testing this test case
			require.NoError(t, db.ExecFixture("../../../fixture/example/example.sql"))

			rec, c := setUpTestDelete("1")

			require.NoError(t, r.Delete(c))
			require.Equal(t, http.StatusOK, rec.Code)

			// remove data for the next test case
			db.TruncateTables()
		})
	})

	t.Run("400", func(t *testing.T) {
		t.Run("Invalid param", func(t *testing.T) {
			rec, c := setUpTestDelete("1n")

			require.NoError(t, r.Delete(c))
			require.Equal(t, http.StatusBadRequest, rec.Code)

			// remove data for the next test case
			db.TruncateTables()
		})
	})

	t.Run("404", func(t *testing.T) {
		t.Run("Not found", func(t *testing.T) {
			rec, c := setUpTestDelete("1")

			require.NoError(t, r.Delete(c))
			require.Equal(t, http.StatusNotFound, rec.Code)

			// remove data for the next test case
			db.TruncateTables()
		})
	})
}

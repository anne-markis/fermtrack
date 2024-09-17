package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCSVReader(t *testing.T) {
	t.Run("missing file", func(t *testing.T) {
		reader := CSVData{}
		err := reader.ParseCSV("./bogus")
		require.Error(t, err)
	})

	t.Run("valid file", func(t *testing.T) {
		reader := CSVData{}
		err := reader.ParseCSV("../storage/wine.csv")
		require.NoError(t, err)
		assert.Len(t, reader.Headers, 4)
		assert.Len(t, reader.Data, 3)
		assert.Equal(t, "name", reader.Headers[0])
		assert.Equal(t, "watermelon", reader.Data[0][0])
	})
}

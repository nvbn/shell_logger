package shell

import (
	"github.com/boltdb/bolt"
	"testing"
	"github.com/stretchr/testify/assert"
)

// Clears and reloads the bucket
func ClearData() error{
	db := openConn()
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte("Mappings"))
		tx.CreateBucket([]byte("Mappings"))
		return err
	})
	return err
}

func testPutGet(t *testing.T) {
	corr := []byte("git push origin master")
	incorr := []byte("git push origin mast")
	incorr2 := []byte("git pus origin master")
	Insert(corr, incorr)
	Insert(corr, incorr2)

	secondCorr := []byte("fc -ln -l")
	secondIncorr := []byte("fd -ln -l")

	Insert(secondCorr, secondIncorr)
	newCorr := []byte("git commit -m")
	newIncorr := []byte("git comit -m")
	Insert(newCorr, newIncorr)

	command1, _ := GetGoodCommands([]byte("git"))
	command2, _ := GetGoodCommands([]byte("fc"))
	assert.Equal(t, 3, len(command1))
	assert.Equal(t, corr, command1[0])
	assert.Equal(t, corr, command1[1])
	assert.Equal(t, newCorr, command1[2])
	assert.Equal(t, 1, len(command2))
	assert.Equal(t, secondCorr, command2[0])
}

func TestDB(t *testing.T){
	SetupDatabase("test.db")

	t.Run("TestPutGet", testPutGet)

	ClearData()
}
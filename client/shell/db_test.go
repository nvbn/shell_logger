package shell

import (
	"testing"
	"github.com/boltdb/bolt"
	"fmt"
)

func ClearBuckets() error{
	err := db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte("Mappings"))
	})
	return err
}

func TestPutGet(t *testing.T){
	SetupDatabase()
	corr := []byte("git push origin master")
	incorr := []byte("git push origin mast")
	incorr2 := []byte("git pus origin master")
	secondCorr := []byte("fc -ln -l")
	secondIncorr := []byte("fd -ln -l")
	newCorr := []byte("git commit -m")
	newIncorr := []byte("git comit -m")
	Insert(corr, incorr)
	Insert(secondCorr, secondIncorr)
	Insert(corr, incorr2)
	Insert(newCorr, newIncorr)
	str1, _ := GetGoodCommands([]byte("git"))
	str2, _ := GetGoodCommands([]byte("fc"))
	for i := 0; i < len(str1); i++ {
		fmt.Println(string(str1[i]))
	}
	for i := 0; i < len(str2); i++ {
		fmt.Println(string(str2[i]))
	}
	ClearBuckets()
}
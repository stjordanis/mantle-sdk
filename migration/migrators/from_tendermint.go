package migrators

import (
	"fmt"
	tmdb "github.com/tendermint/tm-db"
)

// MigrateFromTendermint reads all keys from tmdb,
// rewrites them in mantle-db
// depends on mantle-db envVar to be set
func MigrateFromTendermint(
	sourceDB tmdb.DB,
	targetDB tmdb.DB,
) {
	i := 0
	it, itErr := sourceDB.Iterator(nil, nil)
	bat := targetDB.NewBatch()

	if itErr != nil {
		panic(itErr)
	}

	for it.Valid() {
		fmt.Printf("\r%v", i)
		bat.Set(it.Key(), it.Value())
		it.Next()
		i++
	}

	if batErr := bat.Write(); batErr != nil {
		panic(batErr)
	}

	bat.Close()
}

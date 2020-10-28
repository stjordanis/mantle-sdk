package main

import (
	"fmt"
	tmdb "github.com/tendermint/tm-db"
	"github.com/terra-project/mantle-sdk/db/badger"
	"github.com/terra-project/mantle-sdk/migration/migrators"
	"os"
)

func main() {
	tendermintdb := os.Args[1]
	mantledb := os.Args[2]

	if tendermintdb == "" || mantledb == "" {
		panic("invalid arguments")
	}

	var sourcedb, targetdb tmdb.DB
	var err error

	bdb := badger.NewBadgerDB(mantledb)
	targetdb = bdb.GetCosmosAdapter()

	// statedb
	if sourcedb, err = tmdb.NewGoLevelDB("state", tendermintdb); err != nil {
		panic(err)
	}

	fmt.Println("migrating state.db...")
	migrators.MigrateFromTendermint(
		sourcedb,
		targetdb,
	)

	// application db
	if sourcedb, err = tmdb.NewGoLevelDB("application", tendermintdb); err != nil {
		panic(err)
	}

	fmt.Println("migrating application.db...")
	migrators.MigrateFromTendermint(
		sourcedb,
		targetdb,
	)

	// application db
	if sourcedb, err = tmdb.NewGoLevelDB("blockstore", tendermintdb); err != nil {
		panic(err)
	}

	fmt.Println("migrating application.db...")
	migrators.MigrateFromTendermint(
		sourcedb,
		targetdb,
	)

	// application db
	if sourcedb, err = tmdb.NewGoLevelDB("evidence", tendermintdb); err != nil {
		panic(err)
	}

	fmt.Println("migrating application.db...")
	migrators.MigrateFromTendermint(
		sourcedb,
		targetdb,
	)

	// application db
	if sourcedb, err = tmdb.NewGoLevelDB("tx_index", tendermintdb); err != nil {
		panic(err)
	}

	fmt.Println("migrating application.db...")
	migrators.MigrateFromTendermint(
		sourcedb,
		targetdb,
	)

	if err := bdb.Compact(); err != nil {
		panic("Error during compaction")
	}
}

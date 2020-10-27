package main

import (
	"fmt"
	tmdb "github.com/tendermint/tm-db"
	"github.com/terra-project/mantle-sdk/db/badger"
	"github.com/terra-project/mantle-sdk/migration/migrators"
	"github.com/terra-project/mantle-sdk/utils"
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
	//
	// // try exporing
	// appdb, _ := sdk.NewLevelDB("application", tendermintdb)
	// app := TerraApp.NewTerraApp(
	// 	log.NewTMLogger(ioutil.Discard),
	// 	appdb,
	// 	nil,
	// 	true, // need this so KVStores are set
	// 	0,
	// 	make(map[int64]bool),
	// 	&wasmconfig.Config{BaseConfig: wasmconfig.BaseConfig{
	// 		ContractQueryGasLimit: uint64(3000000),
	// 		CacheSize:             uint64(3000000),
	// 	}},
	// )
	//
	// a, _, _ := app.ExportAppStateAndValidators(true, []string{})
	// ioutil.WriteFile("hey.json", a, 0777)
	//
	// return

	// test
	// application db
	if sourcedb, err = tmdb.NewGoLevelDB("application", tendermintdb); err != nil {
		panic(err)
	}

	key := utils.ConcatBytes([]byte("s/k:mint/"), []byte{0x00})
	v, err := sourcedb.Get(key)
	fmt.Println(v, err)
	//
	// it, _ := sourcedb.Iterator([]byte("s/k:mint/",), nil)
	//
	// for it.Valid() {
	// 	if !strings.HasPrefix(string(it.Key()), "s/k:mint") {
	// 		break
	// 	}
	// 	fmt.Println(it.Key(), string(it.Key()))
	// 	return
	// }

	return

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

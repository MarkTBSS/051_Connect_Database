package main

import (
	"fmt"
	"os"

	_pkgConfig "github.com/MarkTBSS/051_Connect_Database/config"
	_pkgDatabase "github.com/MarkTBSS/051_Connect_Database/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env.dev"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := _pkgConfig.LoadConfig(envPath())
	db := _pkgDatabase.DbConnect(cfg.Db())
	defer db.Close()
	fmt.Println(db)
}

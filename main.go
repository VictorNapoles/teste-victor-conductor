package main

import (
	_ "github.com/VictorNapoles/teste-victor-conductor/adapter/crud"
	_ "github.com/VictorNapoles/teste-victor-conductor/config"
	_ "github.com/VictorNapoles/teste-victor-conductor/config/database"
	_ "github.com/VictorNapoles/teste-victor-conductor/config/http"
	conductor "github.com/VictorNapoles/teste-victor-conductor/config/http"
	_ "github.com/VictorNapoles/teste-victor-conductor/controller"
	_ "github.com/VictorNapoles/teste-victor-conductor/service"
)

func main() {
	conductor.GetServer().Run()
}

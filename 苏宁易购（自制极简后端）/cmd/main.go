package cmd

import (
	"holidy/awesomeProject/api"
	"holidy/awesomeProject/dao"
)

func main() {

	dao.InitDB()
	api.InitRouter()
}

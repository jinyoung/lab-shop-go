package external 

import (
	"github.com/go-resty/resty/v2"
	"fmt"
	"order/config"
)

var client = resty.New()

func GetInventory( id int) *resty.Response {
	options := config.Reader(config.GetMode())
	target := fmt.Sprintf("https://%s/%s/%s", options["api_url_inventory"], "inventories" , id )
	resp, _ := client.R().Get(target)

	return resp
}

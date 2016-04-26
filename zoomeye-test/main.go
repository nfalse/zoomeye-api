/**********************************************************
 * Author        : nfalse
 * Email         : nfalse@163.com
 * Last modified : 2016-04-26 10:28
 * Filename      : main.go
 * Description   :
 * *******************************************************/
package main // test
import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nfalse/zoomeye-api"
)

var g_url *string = flag.String("url", "http://api.zoomeye.org", "use -url http://api.zoomeye.org")
var g_user *string = flag.String("user", "foo@bar.com", "use -user foo@bar.com")
var g_password *string = flag.String("password", "foobar", "use -password foobar")
var g_search_type *string = flag.String("type", "host", "use -type host")
var g_condition *string = flag.String("condition", "query=\"port:21\"&page=1", "use -condition host")

func main() {
	flag.Parse()
	token, err := zoomeye.GetToken(*g_url, *g_user, *g_password)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	switch *g_search_type {
	case "host":
		host_answer, err := zoomeye.HostGet(*g_url, *g_condition, token)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		for _, item := range host_answer.Matches {
			fmt.Println(item)
		}
		break
	case "web":
		web_answer, err := zoomeye.WebGet(*g_url, *g_condition, token)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		for _, item := range web_answer.Matches {
			fmt.Println(item)
		}
		break
	default:
		os.Exit(-1)
		break
	}
	os.Exit(0)
}

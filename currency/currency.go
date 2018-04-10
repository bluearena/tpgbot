/* Currency package get actrual currency cources from zenrus.ru and send message
with it in channel after command /currency
*/
package currency

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// link to get current currency course
const link = "http://zenrus.ru/build/js/currents.js?v45"

// Get Currency send actual currency courses from zenrus.ru
func GetCurrency(c chan string) {
	currency := map[string]string{"0": "usd", "1": "eur", "2": "bar", "9": "btc", "10": "bch", "11": "eth"}
	//get cources from web
	doc, err := goquery.NewDocument(link)
	if err != nil {
		log.Fatal(err)
	}
	//convert results in string
	sub := strings.Split(doc.Text()[15:(len(doc.Text())-1)], ",")
	var m map[string]string
	m = make(map[string]string)
	for _, pair := range sub {
		z := strings.Split(pair, ":")
		m[z[0]] = z[1]
	}
	var result string
	//prepeare string to answer
	for k, v := range currency {
		result = result + fmt.Sprintf("%s : %s\n", v, m[k])
	}
	c <- result
}

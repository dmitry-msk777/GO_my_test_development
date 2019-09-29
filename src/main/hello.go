package main

import (
	"fmt"
	//"io"

	//"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"

	//"net/mail"
	//"net/smtp"
	"os"
	"strings"

	//"github.com/atotto/clipboard"
	//"time"
	//"github.com/tiaguinho/gosoap"
	"encoding/xml"
	//"github.com/achiku/soapc"
	//"github.com/hooklift/gowsdl"
	//"crypto/tls"
	//"net/url"
	"bytes"
	"strconv"

	//"golang.org/x/net/html/charset"
	"encoding/json"
	//"os/signal"
	//"syscall"
	//"github.com/gotk3/gotk3/gtk"
	//	"time"
	//"fyne.io/fyne/app"
	//"fyne.io/fyne/widget"
	//	"math/rand"
	//	"sync"
	//"github.com/gosuri/uiprogress"
	//"github.com/gosuri/uiprogress/util/strutil"
	//"github.com/gorilla/mux"
	//"html/template"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"

	//"go.uber.org/ratelimit"
	//"github.com/Syfaro/telegram-bot-api"
	//"net/proxy"
	"golang.org/x/net/proxy"

	"golang.org/x/crypto/bcrypt"
	//"gopkg.in/telegram-bot-api.v4"
)

type Config struct {
	TelegramBotToken string
}

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

type Valute struct {
	ID_attr      string `xml:"ID,attr"`
	NumCode_test string `xml:"NumCode"`
	CharCode     string `xml:"CharCode"`
	Nominal      string `xml:"Nominal"`
	Name         string `xml:"Name"`
	Value        string `xml:"Value"`
}

type Result struct {
	XMLName   xml.Name `xml:"ValCurs"`
	Date_attr string   `xml:"Date,attr"`
	Name_attr string   `xml:"name,attr"`
	Valute    []Valute
}

type CreateUserRequest struct {
	On_date string `xml:"On_date"`
}

type CreateUserResponse struct {
	Vname   string `xml:"Vname"`
	Vnom    string `xml:"Vnom"`
	Vcurs   string `xml:"Vcurs"`
	Vcode   string `xml:"Vcode"`
	VchCode string `xml:"VchCode"`
}

type Country struct {
	BD string
	BE string
	BF string
	BG string
	BA string
	BB string
}

var (
	contents []byte
	buf      bytes.Buffer
	Res_test bytes.Buffer

//	parsed   io.Reader
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

var steps = []string{
	"downloading source",
	"installing deps",
	"compiling",
	"packaging",
	"seeding database",
	"deploying",
	"staring servers",
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//func encodeRFC2047(String string) string {
//	// use mail's rfc2047 to encode any string
//	addr := mail.Address{String, ""}
//	return strings.Trim(addr.String(), " <>")
//}

func indexPage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Test 1C")
	//fmt.Fprint(w, "ContentLength = "+strconv.Itoa(1001))
	fmt.Fprintln(w, "ContentLength = "+strconv.FormatInt(r.ContentLength, 10))
	fmt.Fprintln(w, "Host = "+r.Host)
	fmt.Fprintln(w, "RemoteAddr = "+r.RemoteAddr)
	fmt.Fprintln(w, "RequestURI = "+r.RequestURI)
	fmt.Fprintln(w, "r.URL.String() = "+r.URL.String())
	//fmt.Fprintln(w, "----------------------------------------")
	fmt.Fprintln(w, "r.URL.Query().Get() = "+r.URL.Query().Get("token"))
	q := r.URL.Query()
	fmt.Fprintln(w, "r.URL.Query().Get()_2 = "+strings.Join(q["token"], ", "))
	w.Write([]byte("Gorilla!\n"))
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test 1C")
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//func deploy(app string, wg *sync.WaitGroup) {
//	defer wg.Done()
//	bar := uiprogress.AddBar(len(steps)).AppendCompleted().PrependElapsed()
//	bar.Width = 50

//	// prepend the deploy step to the bar
//	bar.PrependFunc(func(b *uiprogress.Bar) string {
//		return strutil.Resize(app+": "+steps[b.Current()-1], 22)
//	})

//	rand.Seed(500)
//	for bar.Incr() {
//		time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
//	}
//}

func main() {

	////---------Работа с environment variables ---------------------

	//	// Store the PATH environment variable in a variable
	//	path, exists := os.LookupEnv("PATH")

	//	if exists {
	//		// Print the value of the environment variable
	//		fmt.Println("PATH = ", path)
	//	}

	//	// Get the USERNAME environment variable
	//	username := os.Getenv("USERNAME")

	//	// Prints out username environment variable
	//	fmt.Println("username = ", username)

	//---------Конец Работа с environment variables -----------

	////---------Работа с Сигналами (Unix) https://gobyexample.ru/signals.html -----------
	//	sigs := make(chan os.Signal, 1)
	//	done := make(chan bool, 1)
	//	//signal.Notify registers the given channel to receive notifications of the specified signals.

	//	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//	//This goroutine executes a blocking receive for signals. When it gets one it’ll print it out and then notify the program that it can finish.

	//	go func() {
	//		sig := <-sigs
	//		fmt.Println()
	//		fmt.Println(sig)
	//		done <- true
	//	}()
	//	//The program will wait here until it gets the expected signal (as indicated by the goroutine above sending a value on done) and then exit.

	//	fmt.Println("awaiting signal")
	//	<-done
	//	fmt.Println("exiting")

	//---------Конец Работа с Сигналами (Unix) https://gobyexample.ru/signals.html -----------

	////---------Работа с JSON -----------

	//	//jsonbody := []byte(`{"BD": "Dhaka", "BE": "Brussels", "BF": "Ouagadougou", "BG": "Sofia", "BA": "Sarajevo", "BB": "Bridgetown", "WF": "Mata Utu", "BL": "Gustavia", "BM": "Hamilton", "BN": "Bandar Seri Begawan", "BO": "Sucre", "BH": "Manama", "BI": "Bujumbura", "BJ": "Porto-Novo", "BT": "Thimphu", "JM": "Kingston", "BV": "", "BW": "Gaborone", "WS": "Apia", "BQ": "", "BR": "Brasilia", "BS": "Nassau", "JE": "Saint Helier", "BY": "Minsk", "BZ": "Belmopan", "RU": "Moscow", "RW": "Kigali", "RS": "Belgrade", "TL": "Dili", "RE": "Saint-Denis", "TM": "Ashgabat", "TJ": "Dushanbe", "RO": "Bucharest", "TK": "", "GW": "Bissau", "GU": "Hagatna", "GT": "Guatemala City", "GS": "Grytviken", "GR": "Athens", "GQ": "Malabo", "GP": "Basse-Terre", "JP": "Tokyo", "GY": "Georgetown", "GG": "St Peter Port", "GF": "Cayenne", "GE": "Tbilisi", "GD": "St. George's", "GB": "London", "GA": "Libreville", "SV": "San Salvador", "GN": "Conakry", "GM": "Banjul", "GL": "Nuuk", "GI": "Gibraltar", "GH": "Accra", "OM": "Muscat", "TN": "Tunis", "JO": "Amman", "HR": "Zagreb", "HT": "Port-au-Prince", "HU": "Budapest", "HK": "Hong Kong", "HN": "Tegucigalpa", "HM": "", "VE": "Caracas", "PR": "San Juan", "PS": "East Jerusalem", "PW": "Melekeok", "PT": "Lisbon", "SJ": "Longyearbyen", "PY": "Asuncion", "IQ": "Baghdad", "PA": "Panama City", "PF": "Papeete", "PG": "Port Moresby", "PE": "Lima", "PK": "Islamabad", "PH": "Manila", "PN": "Adamstown", "PL": "Warsaw", "PM": "Saint-Pierre", "ZM": "Lusaka", "EH": "El-Aaiun", "EE": "Tallinn", "EG": "Cairo", "ZA": "Pretoria", "EC": "Quito", "IT": "Rome", "VN": "Hanoi", "SB": "Honiara", "ET": "Addis Ababa", "SO": "Mogadishu", "ZW": "Harare", "SA": "Riyadh", "ES": "Madrid", "ER": "Asmara", "ME": "Podgorica", "MD": "Chisinau", "MG": "Antananarivo", "MF": "Marigot", "MA": "Rabat", "MC": "Monaco", "UZ": "Tashkent", "MM": "Nay Pyi Taw", "ML": "Bamako", "MO": "Macao", "MN": "Ulan Bator", "MH": "Majuro", "MK": "Skopje", "MU": "Port Louis", "MT": "Valletta", "MW": "Lilongwe", "MV": "Male", "MQ": "Fort-de-France", "MP": "Saipan", "MS": "Plymouth", "MR": "Nouakchott", "IM": "Douglas, Isle of Man", "UG": "Kampala", "TZ": "Dodoma", "MY": "Kuala Lumpur", "MX": "Mexico City", "IL": "Jerusalem", "FR": "Paris", "IO": "Diego Garcia", "SH": "Jamestown", "FI": "Helsinki", "FJ": "Suva", "FK": "Stanley", "FM": "Palikir", "FO": "Torshavn", "NI": "Managua", "NL": "Amsterdam", "NO": "Oslo", "NA": "Windhoek", "VU": "Port Vila", "NC": "Noumea", "NE": "Niamey", "NF": "Kingston", "NG": "Abuja", "NZ": "Wellington", "NP": "Kathmandu", "NR": "Yaren", "NU": "Alofi", "CK": "Avarua", "XK": "Pristina", "CI": "Yamoussoukro", "CH": "Berne", "CO": "Bogota", "CN": "Beijing", "CM": "Yaounde", "CL": "Santiago", "CC": "West Island", "CA": "Ottawa", "CG": "Brazzaville", "CF": "Bangui", "CD": "Kinshasa", "CZ": "Prague", "CY": "Nicosia", "CX": "Flying Fish Cove", "CR": "San Jose", "CW": " Willemstad", "CV": "Praia", "CU": "Havana", "SZ": "Mbabane", "SY": "Damascus", "SX": "Philipsburg", "KG": "Bishkek", "KE": "Nairobi", "SS": "Juba", "SR": "Paramaribo", "KI": "Tarawa", "KH": "Phnom Penh", "KN": "Basseterre", "KM": "Moroni", "ST": "Sao Tome", "SK": "Bratislava", "KR": "Seoul", "SI": "Ljubljana", "KP": "Pyongyang", "KW": "Kuwait City", "SN": "Dakar", "SM": "San Marino", "SL": "Freetown", "SC": "Victoria", "KZ": "Astana", "KY": "George Town", "SG": "Singapur", "SE": "Stockholm", "SD": "Khartoum", "DO": "Santo Domingo", "DM": "Roseau", "DJ": "Djibouti", "DK": "Copenhagen", "VG": "Road Town", "DE": "Berlin", "YE": "Sanaa", "DZ": "Algiers", "US": "Washington", "UY": "Montevideo", "YT": "Mamoudzou", "UM": "", "LB": "Beirut", "LC": "Castries", "LA": "Vientiane", "TV": "Funafuti", "TW": "Taipei", "TT": "Port of Spain", "TR": "Ankara", "LK": "Colombo", "LI": "Vaduz", "LV": "Riga", "TO": "Nuku'alofa", "LT": "Vilnius", "LU": "Luxembourg", "LR": "Monrovia", "LS": "Maseru", "TH": "Bangkok", "TF": "Port-aux-Francais", "TG": "Lome", "TD": "N'Djamena", "TC": "Cockburn Town", "LY": "Tripolis", "VA": "Vatican City", "VC": "Kingstown", "AE": "Abu Dhabi", "AD": "Andorra la Vella", "AG": "St. John's", "AF": "Kabul", "AI": "The Valley", "VI": "Charlotte Amalie", "IS": "Reykjavik", "IR": "Tehran", "AM": "Yerevan", "AL": "Tirana", "AO": "Luanda", "AQ": "", "AS": "Pago Pago", "AR": "Buenos Aires", "AU": "Canberra", "AT": "Vienna", "AW": "Oranjestad", "IN": "New Delhi", "AX": "Mariehamn", "AZ": "Baku", "IE": "Dublin", "ID": "Jakarta", "UA": "Kiev", "QA": "Doha", "MZ": "Maputo"}`)

	//url := "http://country.io/capital.json"
	//	response, err := http.Get(url)
	//	if err != nil {
	//		fmt.Println("Error = ", err.Error())
	//	}
	//	defer response.Body.Close()

	//	buf := new(bytes.Buffer)
	//	buf.ReadFrom(response.Body)
	//	//newStr := buf.String()
	//	contents := buf.Bytes()

	//	result_test := &Country{}

	//	err2 := json.Unmarshal(contents, result_test)
	//	if err2 != nil {
	//		fmt.Println("error = ", err2.Error())
	//	}
	//	fmt.Println("%+v", result_test)

	////---------конец Работа с JSON -----------

	//	////---------Работа GET запрос получить XML -----------
	//	response, err := http.Get("http://www.cbr.ru/scripts/XML_daily_eng.asp")
	//	if err != nil {
	//		fmt.Printf("%s", err)
	//		//os.Exit(1)
	//		return
	//	}

	//	defer response.Body.Close()
	//	//	//		contents, err := ioutil.ReadAll(response.Body)
	//	//	//		if err != nil {
	//	//	//			fmt.Printf("%s", err)
	//	//	//			os.Exit(1)
	//	//	//			return
	//	//	//		}
	//	//	//		fmt.Printf("%s\n", string(contents))

	//	//	buf := new(bytes.Buffer)
	//	//	buf.ReadFrom(response.Body)
	//	//	//newStr := buf.String()
	//	//	contents := buf.Bytes()

	//	//	fmt.Println(buf.String())

	//	//contents_2 := []byte(`<?xml version="1.0" encoding="windows-1251"?><ValCurs Date="26.12.2018" name="Foreign Currency Market"><Valute ID="R01010"><NumCode>036</NumCode><CharCode>AUD</CharCode><Nominal>1</Nominal><Name>Australian Dollar</Name><Value>48,4926</Value></Valute><Valute ID="R01020A"><NumCode>944</NumCode><CharCode>AZN</CharCode><Nominal>1</Nominal><Name>Azerbaijan Manat</Name><Value>40,5215</Value></Valute><Valute ID="R01035"><NumCode>826</NumCode><CharCode>GBP</CharCode><Nominal>1</Nominal><Name>British Pound Sterling</Name><Value>87,1890</Value></Valute><Valute ID="R01060"><NumCode>051</NumCode><CharCode>AMD</CharCode><Nominal>100</Nominal><Name>Armenia Dram</Name><Value>14,1742</Value></Valute><Valute ID="R01090B"><NumCode>933</NumCode><CharCode>BYN</CharCode><Nominal>1</Nominal><Name>Belarussian Ruble</Name><Value>32,1237</Value></Valute><Valute ID="R01100"><NumCode>975</NumCode><CharCode>BGN</CharCode><Nominal>1</Nominal><Name>Bulgarian lev</Name><Value>40,0774</Value></Valute><Valute ID="R01115"><NumCode>986</NumCode><CharCode>BRL</CharCode><Nominal>1</Nominal><Name>Brazil Real</Name><Value>17,6088</Value></Valute><Valute ID="R01135"><NumCode>348</NumCode><CharCode>HUF</CharCode><Nominal>100</Nominal><Name>Hungarian Forint</Name><Value>24,3810</Value></Valute><Valute ID="R01200"><NumCode>344</NumCode><CharCode>HKD</CharCode><Nominal>10</Nominal><Name>Hong Kong Dollar</Name><Value>87,7664</Value></Valute><Valute ID="R01215"><NumCode>208</NumCode><CharCode>DKK</CharCode><Nominal>1</Nominal><Name>Danish Krone</Name><Value>10,5010</Value></Valute><Valute ID="R01235"><NumCode>840</NumCode><CharCode>USD</CharCode><Nominal>1</Nominal><Name>US Dollar</Name><Value>68,7448</Value></Valute><Valute ID="R01239"><NumCode>978</NumCode><CharCode>EUR</CharCode><Nominal>1</Nominal><Name>Euro</Name><Value>78,4309</Value></Valute><Valute ID="R01270"><NumCode>356</NumCode><CharCode>INR</CharCode><Nominal>100</Nominal><Name>Indian Rupee</Name><Value>98,1368</Value></Valute><Valute ID="R01335"><NumCode>398</NumCode><CharCode>KZT</CharCode><Nominal>100</Nominal><Name>Kazakhstan Tenge</Name><Value>18,4907</Value></Valute><Valute ID="R01350"><NumCode>124</NumCode><CharCode>CAD</CharCode><Nominal>1</Nominal><Name>Canadian Dollar</Name><Value>50,5402</Value></Valute><Valute ID="R01370"><NumCode>417</NumCode><CharCode>KGS</CharCode><Nominal>100</Nominal><Name>Kyrgyzstan Som</Name><Value>98,4178</Value></Valute><Valute ID="R01375"><NumCode>156</NumCode><CharCode>CNY</CharCode><Nominal>10</Nominal><Name>China Yuan</Name><Value>99,8284</Value></Valute><Valute ID="R01500"><NumCode>498</NumCode><CharCode>MDL</CharCode><Nominal>10</Nominal><Name>Moldova Lei</Name><Value>40,1254</Value></Valute><Valute ID="R01535"><NumCode>578</NumCode><CharCode>NOK</CharCode><Nominal>10</Nominal><Name>Norwegian Krone</Name><Value>78,5242</Value></Valute><Valute ID="R01565"><NumCode>985</NumCode><CharCode>PLN</CharCode><Nominal>1</Nominal><Name>Polish Zloty</Name><Value>18,2958</Value></Valute><Valute ID="R01585F"><NumCode>946</NumCode><CharCode>RON</CharCode><Nominal>1</Nominal><Name>Romanian Leu</Name><Value>16,8973</Value></Valute><Valute ID="R01589"><NumCode>960</NumCode><CharCode>XDR</CharCode><Nominal>1</Nominal><Name>SDR</Name><Value>95,3745</Value></Valute><Valute ID="R01625"><NumCode>702</NumCode><CharCode>SGD</CharCode><Nominal>1</Nominal><Name>Singapore Dollar</Name><Value>50,0581</Value></Valute><Valute ID="R01670"><NumCode>972</NumCode><CharCode>TJS</CharCode><Nominal>10</Nominal><Name>Tajikistan Ruble</Name><Value>72,9388</Value></Valute><Valute ID="R01700J"><NumCode>949</NumCode><CharCode>TRY</CharCode><Nominal>1</Nominal><Name>Turkish Lira</Name><Value>12,9707</Value></Valute><Valute ID="R01710A"><NumCode>934</NumCode><CharCode>TMT</CharCode><Nominal>1</Nominal><Name>New Turkmenistan Manat</Name><Value>19,6695</Value></Valute><Valute ID="R01717"><NumCode>860</NumCode><CharCode>UZS</CharCode><Nominal>10000</Nominal><Name>Uzbekistan Sum</Name><Value>82,4323</Value></Valute><Valute ID="R01720"><NumCode>980</NumCode><CharCode>UAH</CharCode><Nominal>10</Nominal><Name>Ukrainian Hryvnia</Name><Value>25,1104</Value></Valute><Valute ID="R01760"><NumCode>203</NumCode><CharCode>CZK</CharCode><Nominal>10</Nominal><Name>Czech Koruna</Name><Value>30,3335</Value></Valute><Valute ID="R01770"><NumCode>752</NumCode><CharCode>SEK</CharCode><Nominal>10</Nominal><Name>Swedish Krona</Name><Value>76,0031</Value></Valute><Valute ID="R01775"><NumCode>756</NumCode><CharCode>CHF</CharCode><Nominal>1</Nominal><Name>Swiss Franc</Name><Value>69,6714</Value></Valute><Valute ID="R01810"><NumCode>710</NumCode><CharCode>ZAR</CharCode><Nominal>10</Nominal><Name>S.African Rand</Name><Value>47,0935</Value></Valute><Valute ID="R01815"><NumCode>410</NumCode><CharCode>KRW</CharCode><Nominal>1000</Nominal><Name>South Korean Won</Name><Value>61,1049</Value></Valute><Valute ID="R01820"><NumCode>392</NumCode><CharCode>JPY</CharCode><Nominal>100</Nominal><Name>Japanese Yen</Name><Value>62,3338</Value></Valute></ValCurs>`)

	//	//contents_2 := []byte(`<?xml version="1.0"?><ValCurs Date="26.12.2018" name="Foreign Currency Market"><Valute ID="R01010"><NumCode>036</NumCode><CharCode>AUD</CharCode><Nominal>1</Nominal><Name>Australian Dollar</Name><Value>48,4926</Value></Valute><Valute ID="R01020A"><NumCode>944</NumCode><CharCode>AZN</CharCode><Nominal>1</Nominal><Name>Azerbaijan Manat</Name><Value>40,5215</Value></Valute><Valute ID="R01035"><NumCode>826</NumCode><CharCode>GBP</CharCode><Nominal>1</Nominal><Name>British Pound Sterling</Name><Value>87,1890</Value></Valute><Valute ID="R01060"><NumCode>051</NumCode><CharCode>AMD</CharCode><Nominal>100</Nominal><Name>Armenia Dram</Name><Value>14,1742</Value></Valute><Valute ID="R01090B"><NumCode>933</NumCode><CharCode>BYN</CharCode><Nominal>1</Nominal><Name>Belarussian Ruble</Name><Value>32,1237</Value></Valute><Valute ID="R01100"><NumCode>975</NumCode><CharCode>BGN</CharCode><Nominal>1</Nominal><Name>Bulgarian lev</Name><Value>40,0774</Value></Valute><Valute ID="R01115"><NumCode>986</NumCode><CharCode>BRL</CharCode><Nominal>1</Nominal><Name>Brazil Real</Name><Value>17,6088</Value></Valute><Valute ID="R01135"><NumCode>348</NumCode><CharCode>HUF</CharCode><Nominal>100</Nominal><Name>Hungarian Forint</Name><Value>24,3810</Value></Valute><Valute ID="R01200"><NumCode>344</NumCode><CharCode>HKD</CharCode><Nominal>10</Nominal><Name>Hong Kong Dollar</Name><Value>87,7664</Value></Valute><Valute ID="R01215"><NumCode>208</NumCode><CharCode>DKK</CharCode><Nominal>1</Nominal><Name>Danish Krone</Name><Value>10,5010</Value></Valute><Valute ID="R01235"><NumCode>840</NumCode><CharCode>USD</CharCode><Nominal>1</Nominal><Name>US Dollar</Name><Value>68,7448</Value></Valute><Valute ID="R01239"><NumCode>978</NumCode><CharCode>EUR</CharCode><Nominal>1</Nominal><Name>Euro</Name><Value>78,4309</Value></Valute><Valute ID="R01270"><NumCode>356</NumCode><CharCode>INR</CharCode><Nominal>100</Nominal><Name>Indian Rupee</Name><Value>98,1368</Value></Valute><Valute ID="R01335"><NumCode>398</NumCode><CharCode>KZT</CharCode><Nominal>100</Nominal><Name>Kazakhstan Tenge</Name><Value>18,4907</Value></Valute><Valute ID="R01350"><NumCode>124</NumCode><CharCode>CAD</CharCode><Nominal>1</Nominal><Name>Canadian Dollar</Name><Value>50,5402</Value></Valute><Valute ID="R01370"><NumCode>417</NumCode><CharCode>KGS</CharCode><Nominal>100</Nominal><Name>Kyrgyzstan Som</Name><Value>98,4178</Value></Valute><Valute ID="R01375"><NumCode>156</NumCode><CharCode>CNY</CharCode><Nominal>10</Nominal><Name>China Yuan</Name><Value>99,8284</Value></Valute><Valute ID="R01500"><NumCode>498</NumCode><CharCode>MDL</CharCode><Nominal>10</Nominal><Name>Moldova Lei</Name><Value>40,1254</Value></Valute><Valute ID="R01535"><NumCode>578</NumCode><CharCode>NOK</CharCode><Nominal>10</Nominal><Name>Norwegian Krone</Name><Value>78,5242</Value></Valute><Valute ID="R01565"><NumCode>985</NumCode><CharCode>PLN</CharCode><Nominal>1</Nominal><Name>Polish Zloty</Name><Value>18,2958</Value></Valute><Valute ID="R01585F"><NumCode>946</NumCode><CharCode>RON</CharCode><Nominal>1</Nominal><Name>Romanian Leu</Name><Value>16,8973</Value></Valute><Valute ID="R01589"><NumCode>960</NumCode><CharCode>XDR</CharCode><Nominal>1</Nominal><Name>SDR</Name><Value>95,3745</Value></Valute><Valute ID="R01625"><NumCode>702</NumCode><CharCode>SGD</CharCode><Nominal>1</Nominal><Name>Singapore Dollar</Name><Value>50,0581</Value></Valute><Valute ID="R01670"><NumCode>972</NumCode><CharCode>TJS</CharCode><Nominal>10</Nominal><Name>Tajikistan Ruble</Name><Value>72,9388</Value></Valute><Valute ID="R01700J"><NumCode>949</NumCode><CharCode>TRY</CharCode><Nominal>1</Nominal><Name>Turkish Lira</Name><Value>12,9707</Value></Valute><Valute ID="R01710A"><NumCode>934</NumCode><CharCode>TMT</CharCode><Nominal>1</Nominal><Name>New Turkmenistan Manat</Name><Value>19,6695</Value></Valute><Valute ID="R01717"><NumCode>860</NumCode><CharCode>UZS</CharCode><Nominal>10000</Nominal><Name>Uzbekistan Sum</Name><Value>82,4323</Value></Valute><Valute ID="R01720"><NumCode>980</NumCode><CharCode>UAH</CharCode><Nominal>10</Nominal><Name>Ukrainian Hryvnia</Name><Value>25,1104</Value></Valute><Valute ID="R01760"><NumCode>203</NumCode><CharCode>CZK</CharCode><Nominal>10</Nominal><Name>Czech Koruna</Name><Value>30,3335</Value></Valute><Valute ID="R01770"><NumCode>752</NumCode><CharCode>SEK</CharCode><Nominal>10</Nominal><Name>Swedish Krona</Name><Value>76,0031</Value></Valute><Valute ID="R01775"><NumCode>756</NumCode><CharCode>CHF</CharCode><Nominal>1</Nominal><Name>Swiss Franc</Name><Value>69,6714</Value></Valute><Valute ID="R01810"><NumCode>710</NumCode><CharCode>ZAR</CharCode><Nominal>10</Nominal><Name>S.African Rand</Name><Value>47,0935</Value></Valute><Valute ID="R01815"><NumCode>410</NumCode><CharCode>KRW</CharCode><Nominal>1000</Nominal><Name>South Korean Won</Name><Value>61,1049</Value></Valute><Valute ID="R01820"><NumCode>392</NumCode><CharCode>JPY</CharCode><Nominal>100</Nominal><Name>Japanese Yen</Name><Value>62,3338</Value></Valute></ValCurs>`)

	//	//r := strings.NewReader(`<?xml version="1.0" encoding="windows-1251"?><ValCurs Date="26.12.2018" name="Foreign Currency Market"><Valute ID="R01010"><NumCode>036</NumCode><CharCode>AUD</CharCode><Nominal>1</Nominal><Name>Australian Dollar</Name><Value>48,4926</Value></Valute><Valute ID="R01020A"><NumCode>944</NumCode><CharCode>AZN</CharCode><Nominal>1</Nominal><Name>Azerbaijan Manat</Name><Value>40,5215</Value></Valute><Valute ID="R01035"><NumCode>826</NumCode><CharCode>GBP</CharCode><Nominal>1</Nominal><Name>British Pound Sterling</Name><Value>87,1890</Value></Valute><Valute ID="R01060"><NumCode>051</NumCode><CharCode>AMD</CharCode><Nominal>100</Nominal><Name>Armenia Dram</Name><Value>14,1742</Value></Valute><Valute ID="R01090B"><NumCode>933</NumCode><CharCode>BYN</CharCode><Nominal>1</Nominal><Name>Belarussian Ruble</Name><Value>32,1237</Value></Valute><Valute ID="R01100"><NumCode>975</NumCode><CharCode>BGN</CharCode><Nominal>1</Nominal><Name>Bulgarian lev</Name><Value>40,0774</Value></Valute><Valute ID="R01115"><NumCode>986</NumCode><CharCode>BRL</CharCode><Nominal>1</Nominal><Name>Brazil Real</Name><Value>17,6088</Value></Valute><Valute ID="R01135"><NumCode>348</NumCode><CharCode>HUF</CharCode><Nominal>100</Nominal><Name>Hungarian Forint</Name><Value>24,3810</Value></Valute><Valute ID="R01200"><NumCode>344</NumCode><CharCode>HKD</CharCode><Nominal>10</Nominal><Name>Hong Kong Dollar</Name><Value>87,7664</Value></Valute><Valute ID="R01215"><NumCode>208</NumCode><CharCode>DKK</CharCode><Nominal>1</Nominal><Name>Danish Krone</Name><Value>10,5010</Value></Valute><Valute ID="R01235"><NumCode>840</NumCode><CharCode>USD</CharCode><Nominal>1</Nominal><Name>US Dollar</Name><Value>68,7448</Value></Valute><Valute ID="R01239"><NumCode>978</NumCode><CharCode>EUR</CharCode><Nominal>1</Nominal><Name>Euro</Name><Value>78,4309</Value></Valute><Valute ID="R01270"><NumCode>356</NumCode><CharCode>INR</CharCode><Nominal>100</Nominal><Name>Indian Rupee</Name><Value>98,1368</Value></Valute><Valute ID="R01335"><NumCode>398</NumCode><CharCode>KZT</CharCode><Nominal>100</Nominal><Name>Kazakhstan Tenge</Name><Value>18,4907</Value></Valute><Valute ID="R01350"><NumCode>124</NumCode><CharCode>CAD</CharCode><Nominal>1</Nominal><Name>Canadian Dollar</Name><Value>50,5402</Value></Valute><Valute ID="R01370"><NumCode>417</NumCode><CharCode>KGS</CharCode><Nominal>100</Nominal><Name>Kyrgyzstan Som</Name><Value>98,4178</Value></Valute><Valute ID="R01375"><NumCode>156</NumCode><CharCode>CNY</CharCode><Nominal>10</Nominal><Name>China Yuan</Name><Value>99,8284</Value></Valute><Valute ID="R01500"><NumCode>498</NumCode><CharCode>MDL</CharCode><Nominal>10</Nominal><Name>Moldova Lei</Name><Value>40,1254</Value></Valute><Valute ID="R01535"><NumCode>578</NumCode><CharCode>NOK</CharCode><Nominal>10</Nominal><Name>Norwegian Krone</Name><Value>78,5242</Value></Valute><Valute ID="R01565"><NumCode>985</NumCode><CharCode>PLN</CharCode><Nominal>1</Nominal><Name>Polish Zloty</Name><Value>18,2958</Value></Valute><Valute ID="R01585F"><NumCode>946</NumCode><CharCode>RON</CharCode><Nominal>1</Nominal><Name>Romanian Leu</Name><Value>16,8973</Value></Valute><Valute ID="R01589"><NumCode>960</NumCode><CharCode>XDR</CharCode><Nominal>1</Nominal><Name>SDR</Name><Value>95,3745</Value></Valute><Valute ID="R01625"><NumCode>702</NumCode><CharCode>SGD</CharCode><Nominal>1</Nominal><Name>Singapore Dollar</Name><Value>50,0581</Value></Valute><Valute ID="R01670"><NumCode>972</NumCode><CharCode>TJS</CharCode><Nominal>10</Nominal><Name>Tajikistan Ruble</Name><Value>72,9388</Value></Valute><Valute ID="R01700J"><NumCode>949</NumCode><CharCode>TRY</CharCode><Nominal>1</Nominal><Name>Turkish Lira</Name><Value>12,9707</Value></Valute><Valute ID="R01710A"><NumCode>934</NumCode><CharCode>TMT</CharCode><Nominal>1</Nominal><Name>New Turkmenistan Manat</Name><Value>19,6695</Value></Valute><Valute ID="R01717"><NumCode>860</NumCode><CharCode>UZS</CharCode><Nominal>10000</Nominal><Name>Uzbekistan Sum</Name><Value>82,4323</Value></Valute><Valute ID="R01720"><NumCode>980</NumCode><CharCode>UAH</CharCode><Nominal>10</Nominal><Name>Ukrainian Hryvnia</Name><Value>25,1104</Value></Valute><Valute ID="R01760"><NumCode>203</NumCode><CharCode>CZK</CharCode><Nominal>10</Nominal><Name>Czech Koruna</Name><Value>30,3335</Value></Valute><Valute ID="R01770"><NumCode>752</NumCode><CharCode>SEK</CharCode><Nominal>10</Nominal><Name>Swedish Krona</Name><Value>76,0031</Value></Valute><Valute ID="R01775"><NumCode>756</NumCode><CharCode>CHF</CharCode><Nominal>1</Nominal><Name>Swiss Franc</Name><Value>69,6714</Value></Valute><Valute ID="R01810"><NumCode>710</NumCode><CharCode>ZAR</CharCode><Nominal>10</Nominal><Name>S.African Rand</Name><Value>47,0935</Value></Valute><Valute ID="R01815"><NumCode>410</NumCode><CharCode>KRW</CharCode><Nominal>1000</Nominal><Name>South Korean Won</Name><Value>61,1049</Value></Valute><Valute ID="R01820"><NumCode>392</NumCode><CharCode>JPY</CharCode><Nominal>100</Nominal><Name>Japanese Yen</Name><Value>62,3338</Value></Valute></ValCurs>`)

	//	//b := bytes.NewBufferString(`<?xml version="1.0" encoding="windows-1251"?><ValCurs Date="26.12.2018" name="Foreign Currency Market"><Valute ID="R01010"><NumCode>036</NumCode><CharCode>AUD</CharCode><Nominal>1</Nominal><Name>Australian Dollar</Name><Value>48,4926</Value></Valute><Valute ID="R01020A"><NumCode>944</NumCode><CharCode>AZN</CharCode><Nominal>1</Nominal><Name>Azerbaijan Manat</Name><Value>40,5215</Value></Valute><Valute ID="R01035"><NumCode>826</NumCode><CharCode>GBP</CharCode><Nominal>1</Nominal><Name>British Pound Sterling</Name><Value>87,1890</Value></Valute><Valute ID="R01060"><NumCode>051</NumCode><CharCode>AMD</CharCode><Nominal>100</Nominal><Name>Armenia Dram</Name><Value>14,1742</Value></Valute><Valute ID="R01090B"><NumCode>933</NumCode><CharCode>BYN</CharCode><Nominal>1</Nominal><Name>Belarussian Ruble</Name><Value>32,1237</Value></Valute><Valute ID="R01100"><NumCode>975</NumCode><CharCode>BGN</CharCode><Nominal>1</Nominal><Name>Bulgarian lev</Name><Value>40,0774</Value></Valute><Valute ID="R01115"><NumCode>986</NumCode><CharCode>BRL</CharCode><Nominal>1</Nominal><Name>Brazil Real</Name><Value>17,6088</Value></Valute><Valute ID="R01135"><NumCode>348</NumCode><CharCode>HUF</CharCode><Nominal>100</Nominal><Name>Hungarian Forint</Name><Value>24,3810</Value></Valute><Valute ID="R01200"><NumCode>344</NumCode><CharCode>HKD</CharCode><Nominal>10</Nominal><Name>Hong Kong Dollar</Name><Value>87,7664</Value></Valute><Valute ID="R01215"><NumCode>208</NumCode><CharCode>DKK</CharCode><Nominal>1</Nominal><Name>Danish Krone</Name><Value>10,5010</Value></Valute><Valute ID="R01235"><NumCode>840</NumCode><CharCode>USD</CharCode><Nominal>1</Nominal><Name>US Dollar</Name><Value>68,7448</Value></Valute><Valute ID="R01239"><NumCode>978</NumCode><CharCode>EUR</CharCode><Nominal>1</Nominal><Name>Euro</Name><Value>78,4309</Value></Valute><Valute ID="R01270"><NumCode>356</NumCode><CharCode>INR</CharCode><Nominal>100</Nominal><Name>Indian Rupee</Name><Value>98,1368</Value></Valute><Valute ID="R01335"><NumCode>398</NumCode><CharCode>KZT</CharCode><Nominal>100</Nominal><Name>Kazakhstan Tenge</Name><Value>18,4907</Value></Valute><Valute ID="R01350"><NumCode>124</NumCode><CharCode>CAD</CharCode><Nominal>1</Nominal><Name>Canadian Dollar</Name><Value>50,5402</Value></Valute><Valute ID="R01370"><NumCode>417</NumCode><CharCode>KGS</CharCode><Nominal>100</Nominal><Name>Kyrgyzstan Som</Name><Value>98,4178</Value></Valute><Valute ID="R01375"><NumCode>156</NumCode><CharCode>CNY</CharCode><Nominal>10</Nominal><Name>China Yuan</Name><Value>99,8284</Value></Valute><Valute ID="R01500"><NumCode>498</NumCode><CharCode>MDL</CharCode><Nominal>10</Nominal><Name>Moldova Lei</Name><Value>40,1254</Value></Valute><Valute ID="R01535"><NumCode>578</NumCode><CharCode>NOK</CharCode><Nominal>10</Nominal><Name>Norwegian Krone</Name><Value>78,5242</Value></Valute><Valute ID="R01565"><NumCode>985</NumCode><CharCode>PLN</CharCode><Nominal>1</Nominal><Name>Polish Zloty</Name><Value>18,2958</Value></Valute><Valute ID="R01585F"><NumCode>946</NumCode><CharCode>RON</CharCode><Nominal>1</Nominal><Name>Romanian Leu</Name><Value>16,8973</Value></Valute><Valute ID="R01589"><NumCode>960</NumCode><CharCode>XDR</CharCode><Nominal>1</Nominal><Name>SDR</Name><Value>95,3745</Value></Valute><Valute ID="R01625"><NumCode>702</NumCode><CharCode>SGD</CharCode><Nominal>1</Nominal><Name>Singapore Dollar</Name><Value>50,0581</Value></Valute><Valute ID="R01670"><NumCode>972</NumCode><CharCode>TJS</CharCode><Nominal>10</Nominal><Name>Tajikistan Ruble</Name><Value>72,9388</Value></Valute><Valute ID="R01700J"><NumCode>949</NumCode><CharCode>TRY</CharCode><Nominal>1</Nominal><Name>Turkish Lira</Name><Value>12,9707</Value></Valute><Valute ID="R01710A"><NumCode>934</NumCode><CharCode>TMT</CharCode><Nominal>1</Nominal><Name>New Turkmenistan Manat</Name><Value>19,6695</Value></Valute><Valute ID="R01717"><NumCode>860</NumCode><CharCode>UZS</CharCode><Nominal>10000</Nominal><Name>Uzbekistan Sum</Name><Value>82,4323</Value></Valute><Valute ID="R01720"><NumCode>980</NumCode><CharCode>UAH</CharCode><Nominal>10</Nominal><Name>Ukrainian Hryvnia</Name><Value>25,1104</Value></Valute><Valute ID="R01760"><NumCode>203</NumCode><CharCode>CZK</CharCode><Nominal>10</Nominal><Name>Czech Koruna</Name><Value>30,3335</Value></Valute><Valute ID="R01770"><NumCode>752</NumCode><CharCode>SEK</CharCode><Nominal>10</Nominal><Name>Swedish Krona</Name><Value>76,0031</Value></Valute><Valute ID="R01775"><NumCode>756</NumCode><CharCode>CHF</CharCode><Nominal>1</Nominal><Name>Swiss Franc</Name><Value>69,6714</Value></Valute><Valute ID="R01810"><NumCode>710</NumCode><CharCode>ZAR</CharCode><Nominal>10</Nominal><Name>S.African Rand</Name><Value>47,0935</Value></Valute><Valute ID="R01815"><NumCode>410</NumCode><CharCode>KRW</CharCode><Nominal>1000</Nominal><Name>South Korean Won</Name><Value>61,1049</Value></Valute><Valute ID="R01820"><NumCode>392</NumCode><CharCode>JPY</CharCode><Nominal>100</Nominal><Name>Japanese Yen</Name><Value>62,3338</Value></Valute></ValCurs>`)

	//	//	if _, err := io.ReadFull(r, contents_2); err != nil {
	//	//		fmt.Println("Error io.ReadFull = ", err.Error())
	//	//	}

	//	//	buf := new(bytes.Buffer)
	//	//	buf.ReadFrom(r)
	//	//	fmt.Println("contents_2 = ", buf.Bytes())

	//	//b := bytes.NewBufferString("your string")

	//	res := &Result{}
	//	decoder := xml.NewDecoder(response.Body)
	//	decoder.CharsetReader = charset.NewReaderLabel
	//	decoder.Decode(res)

	//	//	buf := new(bytes.Buffer)
	//	//	buf.ReadFrom(b)
	//	//	fmt.Println("b = ", buf.String())

	//	//	fmt.Println("Res_test = ", Res_test.String())

	//	//	buf := new(bytes.Buffer)
	//	//	buf.ReadFrom(r)
	//	//newStr := buf.String()
	//	//contents := buf.Bytes()

	//	//	//Не считывает буфер, а просто создает обвертку
	//	//	body := io.TeeReader(req.Body, &buf)
	//	//	// ... process body ...
	//	//	if err != nil {
	//	//		// inspect buf
	//	//		return err
	//	//	}

	//	//	res := &Result{}
	//	//string(contents_2)
	//	//err2 := xml.Unmarshal(Res_test.Bytes(), res)
	//	//	err2 := xml.Unmarshal(contents, res)
	//	//	if err2 != nil {
	//	//		fmt.Println("Error Unmarshal = ", err2.Error())
	//	//	}

	//	//-------------- обозначить

	//	//	response, err := http.Get("http://www.cbr.ru/scripts/XML_daily.asp")
	//	//	if err != nil {
	//	//		fmt.Printf("%s", err)
	//	//		os.Exit(1)
	//	//	} else {
	//	//		defer response.Body.Close()
	//	//		contents, err := ioutil.ReadAll(response.Body)
	//	//		if err != nil {
	//	//			fmt.Printf("%s", err)
	//	//			os.Exit(1)
	//	//		}
	//	//		fmt.Printf("%s\n", string(contents))
	//	//	}

	//	////---------конец Работа GET запрос получить XML -----------

	////---------Работа POST прямым запросом к SOAP-----------
	//	apiUrl := "http://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx?WSDL"
	//	//resource := "GetCursOnDate"
	//	data := url.Values{}
	//	//	data.Set("name", "foo")
	//	//	data.Add("surname", "bar")

	//	u, _ := url.ParseRequestURI(apiUrl)
	//	//u.Path = resource
	//	urlStr := u.String() // 'https://api.com/user/'

	//	client := &http.Client{}
	//	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	//	//	r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	//	//	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//	//	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	//	r.Header.Set("SOAPAction", "GetCursOnDate")
	//	r.Header.Set("Content-Type", "text/xml")
	//	r.Header.Set("charset", "utf-8")
	//	//r.Header.Set("User-Agent", "gowsdl/0.1")

	//	resp, _ := client.Do(r)
	//	fmt.Println("resp.Status = ", resp.Status)

	//url := "http://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx?WSDL"

	//	req, err := http.NewRequest("POST", url, file)
	//	if err != nil {
	//		fmt.Println("Error = ", err.Error())
	//	}
	//	req.Header.Set("SOAPAction", "GetCursOnDate")
	//	req.Header.Set("Content-Type", "text/xml")
	//	req.Header.Set("charset", "utf-8")
	//	req.Header.Set("User-Agent", "gowsdl/0.1")

	//defer res.Body.Close()

	////---------Конец Работа POST прямым запросом к SOAP-----------

	////---------Работа парсингом SOAP сообщения ответа-----------

	//	Soap := []byte(`<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">

	//  <SOAP-ENV:Header xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
	//    <wsse:Security soap:mustUnderstand="1" xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	//      <wsse:UsernameToken>
	//        <wsse:Username>USERNAME</wsse:Username>
	//        <wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">SECRET</wsse:Password>
	//      </wsse:UsernameToken>
	//    </wsse:Security>
	//  </SOAP-ENV:Header>

	//  <SOAP-ENV:Body xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
	//    <OTA_HotelAvailNotifRQ xmlns="http://www.opentravel.org/OTA/2003/05" EchoToken="abc123" Version="1.0" TimeStamp="2005-08-01T09:30:47+08:00">
	//      <AvailStatusMessages HotelCode="HOTEL">
	//        <AvailStatusMessage BookingLimit="10">
	//          <StatusApplicationControl Start="2010-01-01" End="2010-01-14" InvTypeCode="A1K" RatePlanCode="GLD"/>
	//        </AvailStatusMessage>
	//      </AvailStatusMessages>
	//    </OTA_HotelAvailNotifRQ>
	//  </SOAP-ENV:Body>

	//</SOAP-ENV:Envelope>`)

	//	res := &Envelope{}
	//	err := xml.Unmarshal(Soap, res)

	//	fmt.Println(res.Header.Security.UsernameToken.Username, err)
	////--------- Конец Работа парсингом SOAP сообщения ответа-----------

	////---------Работа с SOAP через "github.com/hooklift/gowsdl"-----------
	//	g, err := gowsdl.NewGoWSDL("http://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx?WSDL", "GetCursOnDate", false, true)

	//	if err != nil {
	//		fmt.Println("Error:= ", err.Error())
	//	}

	//	resp, err := g.Start()
	//	if err != nil {
	//		fmt.Println("Error:= ", err.Error())
	//	}

	//	if strings.Contains(string(resp["types"]), "// this is a comment  GetInfoResult string `xml:\"GetInfoResult,omitempty\"`") {
	//		fmt.Println("Type comment should not comment out struct type property")
	//		fmt.Println(string(resp["types"]))
	//	}
	////---------Конец с SOAP через "github.com/hooklift/gowsdl"-----------

	////---------Работа с SOAP через "github.com/achiku/soapc"-----------
	//	isTLS := false
	//	//url := "http://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx?WSDL" + "/noheader"
	//	url := "http://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx?WSDL"
	//	client := soap.NewClient(url, isTLS, nil)

	//	req := testRequest{Message: "test"}
	//	resp := person{}
	//	if err := client.Call(url, req, &resp, nil); err != nil {
	//		fmt.Println("Error:= ", err.Error())
	//	}
	//---Бред
	//	client := gowsdl. NewSOAPClient("http://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx?WSDL", true, nil)

	//	req := &CreateUserRequest{
	//		On_date: "24122018",
	//	}
	//	res := &CreateUserResponse{}
	//	if err := client.Call("create_user", req, res); err != nil {
	//		panic(err)
	//	}

	////---------Конец с SOAP через "github.com/achiku/soapc"-----------

	////---------Работа с SOAP пример через "github.com/tiaguinho/gosoap"-----------
	//--- Пример
	//	soap, err := gosoap.SoapClient("http://www.webservicex.net/geoipservice.asmx?WSDL")
	//	if err != nil {
	//		fmt.Errorf("error not expected: %s", err)
	//	}

	//	params := gosoap.Params{
	//		"IPAddress": "8.8.8.8",
	//	}

	//	err = soap.Call("GetGeoIP", params)
	//	if err != nil {
	//		fmt.Errorf("error in soap call: %s", err)
	//	}

	//	soap.Unmarshal(&r)
	//	if r.GetGeoIPResult.CountryCode != "USA" {
	//		fmt.Errorf("error: %+v", r)
	//	}
	//Мой код
	//soap, err := gosoap.SoapClient("http://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx?WSDL")
	//	soap, err := gosoap.SoapClient("http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl")
	//	if err != nil {
	//		fmt.Println("error not expected: %s", err.Error())
	//	}

	//	params := gosoap.Params{
	//		//"On_date": time.Now(),
	//		"On_date": "24122018",
	//	}

	//	err = soap.Call("GetCursOnDate", params)
	//	if err != nil {
	//		fmt.Errorf("error in soap call: %s", err)
	//	}

	//	soap.Unmarshal(&r)
	//	if r.ValuteCursOnDate.Vcode != "USA" {
	//		fmt.Errorf("error: %+v", r)
	//	}
	//Конец мой код
	////---------Конец Работа с SOAP через "github.com/tiaguinho/gosoap"-----------

	//	b, err2 := ioutil.ReadFile("hello.txt")
	//	if err2 != nil {
	//		panic(err2)
	//	}
	//	//	fmt.Println(string(b))
	//	FileInside := string(b)
	//	fmt.Println(FileInside)

	////---------Чтение из файла-----------
	//	file, err3 := os.Open("hello.txt")
	//	if err3 != nil {
	//		fmt.Println(err3)
	//		os.Exit(0)
	//	}
	//	defer file.Close()

	//	data := make([]byte, 64)

	//	for {
	//		n, err4 := file.Read(data)
	//		if err4 == io.EOF {
	//			break
	//		}
	//		fmt.Print(string(data[:n]))
	//	}
	////---------Конец Чтение из файла-----------

	////---------Работа с буфером обмена-----------
	//	//"./atotto"
	//	//"./atotto/clipboard"

	//	//"github.com/atotto/clipboard"
	//	//clipboard.WriteAll("日本語")
	//	//text, _ := clipboard.ReadAll()
	//	//fmt.Println(text)

	//	clipRead, err := clipboard.ReadAll()

	//	if err != nil {
	//		fmt.Printf("Error: %s", err.Error())
	//		return
	//	}
	////---------Конец Работа с буфером обмена-----------

	//---------------- Работа с почтой -----------------------
	// Set up authentication information.

	// Настройки почты яндекса https://yandex.ru/support/mail/mail-clients.html

	//	smtpServer := "smtp.yandex.ru"
	//	auth := smtp.PlainAuth(
	//		"",
	//		"dmitry-msk777@yandex.ru",
	//		"Cnhwgkmgenrm1991",
	//		smtpServer,
	//	)

	//	from := mail.Address{"Test", "dmitry-msk777@yandex.ru"}
	//	to := mail.Address{"test2", "dima-irk35@mail.ru"}
	//	title := "Title"

	//	body := "body"

	//	header := make(map[string]string)
	//	header["From"] = from.String()
	//	header["To"] = to.String()
	//	header["Subject"] = encodeRFC2047(title)
	//	header["MIME-Version"] = "1.0"
	//	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	//	header["Content-Transfer-Encoding"] = "base64"

	//	message := ""
	//	for k, v := range header {
	//		message += fmt.Sprintf("%s: %s\r\n", k, v)
	//	}
	//	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	//	// Connect to the server, authenticate, set the sender and recipient,
	//	// and send the email all in one step.
	//	err := smtp.SendMail(
	//		smtpServer+":25",
	//		auth,
	//		from.Address,
	//		[]string{to.Address},
	//		[]byte(message),
	//		//[]byte("This is the email body."),
	//	)
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	//--------------- Конец работа с почтой -----------------

	//--------------- Работа с Прогресс баром -----------------

	//	uiprogress.Start()            // start rendering
	//	bar := uiprogress.AddBar(100) // Add a new bar

	//	// optionally, append and prepend completion and elapsed time
	//	bar.AppendCompleted()
	//	bar.PrependElapsed()

	//	for bar.Incr() {
	//		time.Sleep(time.Millisecond * 20)
	//	}

	//-----------------------------------------------------------

	//	fmt.Println("apps: deployment started: app1, app2")
	//	uiprogress.Start()

	//	var wg sync.WaitGroup
	//	wg.Add(1)
	//	go deploy("app1", &wg)
	//	wg.Add(1)
	//	go deploy("app2", &wg)
	//	wg.Wait()

	//	fmt.Println("apps: successfully deployed: app1, app2")

	//--------------- Конец работа с Прогресс баром -----------------

	//--------------- Работа с Шаблонами HTTP -----------------
	//Страница
	//	tmpl := template.Must(template.ParseFiles("layout.html"))

	//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//		data := TodoPageData{
	//			PageTitle: "My TODO list",
	//			Todos: []Todo{
	//				{Title: "Task 1", Done: false},
	//				{Title: "Task 2", Done: true},
	//				{Title: "Task 3", Done: true},
	//			},
	//		}
	//		tmpl.Execute(w, data)
	//	})

	//	http.ListenAndServe(":8081", nil)

	//Форма
	//	tmpl := template.Must(template.ParseFiles("forms.html"))

	//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//		if r.Method != http.MethodPost {
	//			tmpl.Execute(w, nil)
	//			return
	//		}

	//		details := ContactDetails{
	//			Email:   r.FormValue("email"),
	//			Subject: r.FormValue("subject"),
	//			Message: r.FormValue("message"),
	//		}

	//		// do something with details
	//		_ = details

	//		tmpl.Execute(w, struct{ Success bool }{true})
	//	})

	//	http.ListenAndServe(":8081", nil)
	//--------------- Конец работа с Шаблонами HTTP -----------------

	//--------------- Работа с Файл-сервером HTTP -----------------

	//	fs := http.FileServer(http.Dir("assets/"))
	//	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//	http.ListenAndServe(":8081", nil)
	//--------------- Конец работа с Файл-сервером HTTP -----------------

	//--------------- Работа с Middleware (Basic) -----------------
	//	http.HandleFunc("/foo", logging(foo))
	//	http.HandleFunc("/bar", logging(bar))

	//	http.ListenAndServe(":8081", nil)
	//--------------- Конец работы с Middleware (Basic) -----------------

	//--------------- Работа с Сессиями и куки -----------------
	//	http.HandleFunc("/secret", secret)
	//	http.HandleFunc("/login", login)
	//	http.HandleFunc("/logout", logout)

	//	http.ListenAndServe(":8081", nil)
	//--------------- Конец работа с Сессиями и куки -----------------

	//--------------- Работа с WebSocket -----------------
	//	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
	//		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

	//		for {
	//			// Read message from browser
	//			msgType, msg, err := conn.ReadMessage()
	//			if err != nil {
	//				return
	//			}

	//			// Print the message to the console
	//			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

	//			// Write message back to browser
	//			if err = conn.WriteMessage(msgType, msg); err != nil {
	//				return
	//			}
	//		}
	//	})

	//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//		http.ServeFile(w, r, "websockets.html")
	//	})

	//	http.ListenAndServe(":8081", nil)
	//--------------- Конец работа с WebSocket -----------------

	//--------------- Работа с Password Hashing (bcrypt) -----------------
	//	password := "secret"
	//	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	//	fmt.Println("Password:", password)
	//	fmt.Println("Hash:    ", hash)

	//	match := CheckPasswordHash(password, hash)
	//	fmt.Println("Match:   ", match)
	//---------------Конец работа с Password Hashing (bcrypt) -----------------

	//--------------- Работа с gorilla/mux -----------------
	//	//	r := mux.NewRouter()
	//	//	r.HandleFunc("/", indexPage)
	//	//	r.HandleFunc("/products", ProductsHandler)
	//	//http.ListenAndServe(":8081", r)
	//--------------- Конец работа с gorilla/mux -----------------

	//--------------- Работа с "go.uber.org/ratelimit задержка в секундах -----------------
	//	rl := ratelimit.New(1) // per second

	//	prev := time.Now()
	//	for i := 0; i < 10; i++ {
	//		now := rl.Take()
	//		fmt.Println(i, now.Sub(prev))
	//		prev = now
	//	}
	//--------------- Конец работа с "go.uber.org/ratelimit задержка в секундах -----------------

	//--------------- Работа с Telegram Bot через "github.com/Syfaro/telegram-bot-api" -----------------
	file, _ := os.Open("./dndspellsbot/config.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err2 := decoder.Decode(&configuration)
	if err2 != nil {
		log.Panic(err2)
	}
	fmt.Println(configuration.TelegramBotToken)

	//dialSocksProxy, err := proxy.SOCKS5("tcp", "96.44.183.149:55225", nil, proxy.Direct)
	dialSocksProxy, err := proxy.SOCKS5("tcp", "96.44.183.149:55225", nil, nil)
	if err != nil {
		fmt.Println("Error connecting to proxy:", err)
	}
	tr := &http.Transport{Dial: dialSocksProxy.Dial}

	// Create client
	myClient := &http.Client{
		Transport: tr,
	}

	resp, err := myClient.Get("https://api.telegram.org/808741510:AAECEpVU9cLIdJ0HsHpNASlolDVWYACgyA4/getMe")
	if err != nil {
		fmt.Println(err)
		//return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	//	fmt.Println("333")
	//	bot, err := tgbotapi.NewBotAPIWithClient("808741510:AAECEpVU9cLIdJ0HsHpNASlolDVWYACgyA4", myClient)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println("123")
	//	bot.Debug = false
	//	fmt.Println("Authorized on account %s", bot.Self.UserName)

	//	fmt.Println("Authorized on account %s", bot.Self.UserName)
	//--------------- Конец работа с Telegram Bot через "github.com/Syfaro/telegram-bot-api" -----------------

	//--------------- Работа с Telegram Bot через "gopkg.in/telegram-bot-api.v4" -----------------

	//	//dialSocksProxy, err := proxy.SOCKS5("tcp", "85.238.105.35:3629", nil, proxy.Direct)
	//	dialSocksProxy, err := proxy.SOCKS5("tcp", "190.242.41.133:11337", nil, nil)
	//	if err != nil {
	//		fmt.Println("Error connecting to proxy:", err)
	//	}
	//	tr := &http.Transport{Dial: dialSocksProxy.Dial}

	//	// Create client
	//	myClient := &http.Client{
	//		Transport: tr,
	//	}

	//	file, _ := os.Open("./dndspellsbot/config.json")
	//	decoder := json.NewDecoder(file)
	//	configuration := Config{}
	//	err2 := decoder.Decode(&configuration)
	//	if err2 != nil {
	//		log.Panic(err2)
	//	}
	//	fmt.Println(configuration.TelegramBotToken)

	//	bot, err := tgbotapi.NewBotAPIWithClient(configuration.TelegramBotToken, myClient)

	//	if err != nil {
	//		log.Panic(err)
	//	}

	//	bot.Debug = false

	//	log.Printf("Authorized on account %s", bot.Self.UserName)

	//	u := tgbotapi.NewUpdate(0)
	//	u.Timeout = 60

	//	updates, err := bot.GetUpdatesChan(u)

	//	if err != nil {
	//		log.Panic(err)
	//	}
	//	// В канал updates будут приходить все новые сообщения.
	//	for update := range updates {
	//		// Создав структуру - можно её отправить обратно боту
	//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//		msg.ReplyToMessageID = update.Message.MessageID
	//		bot.Send(msg)
	//	}
	//--------------- Конец работа с Telegram Bot через "gopkg.in/telegram-bot-api.v4" -----------------

	http.HandleFunc("/", indexPage)
	http.HandleFunc("/products", ProductsHandler)
	fmt.Println("Start 1C Port 8081")
	http.ListenAndServe(":8081", nil)

}

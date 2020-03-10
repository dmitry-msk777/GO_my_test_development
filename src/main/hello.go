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
	//	"os"
	"strings"

	//"github.com/atotto/clipboard"
	"time"
	//"github.com/tiaguinho/gosoap"
	//"encoding/xml"
	//"github.com/achiku/soapc"
	//"github.com/hooklift/gowsdl"
	//"crypto/tls"
	"bytes"
	"net/url"
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
	//"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"

	//"github.com/olivere/elastic"
	//"go.uber.org/ratelimit"
	//	"github.com/Syfaro/telegram-bot-api"
	//"net/proxy"

	//	"golang.org/x/net/proxy"

	//	"reflect"

	//"github.com/carbocation/go-instagram/instagram"
	"unsafe"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sys/windows"

	"flag"
	"net"

	//	"github.com/m7shapan/uuid"
	//	"regexp"
	//"github.com/dmitry-msk777/GO_my_test_development"
	//"github.com/dmitry-msk777/GO_my_test_development/src/main"
	//"github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promauto"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
	//"math/rand"
	"crypto/md5"
	//	"encoding/hex"
	//"runtime"

	//"github.com/flimzy/kivik"       // Stable version of Kivik
	_ "github.com/go-kivik/couchdb" // The CouchDB driver
	// Development version of Kivik
	//"github.com/streadway/amqp"
	//"github.com/olivere/elastic"
	// "gopkg.in/olivere/elastic.v6"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	//"github.com/marpaia/graphite-golang"
	//_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	//"context"
	//"github.com/segmentio/kafka-go"
	// "github.com/go-redis/redis/v7"
	//"github.com/tarantool/go-tarantool"
	//"github.com/gocql/gocql"
	//r "github.com/dancannon/gorethink"
	//"github.com/golang/protobuf/proto"
	//pb "github.com/mycodesmells/golang-examples/nats/pubsub/proto"
	//nats "github.com/nats-io/nats.go"
	pb "../main/proto"

	"golang.org/x/net/context"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/grpclog"
	//"github.com/gorilla/mux"

	"golang.org/x/net/websocket"
)

type server struct{}

func (s *server) Do(c context.Context, request *pb.Request) (response *pb.Response, err error) {
	n := 0
	// Сreate an array of runes to safely reverse a string.
	rune := make([]rune, len(request.Message))

	for _, r := range request.Message {
		rune[n] = r
		n++
	}

	// Reverse using runes.
	rune = rune[0:n]

	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}

	output := string(rune)
	response = &pb.Response{
		Message: output,
	}

	return response, nil
}

type Config struct {
	TelegramBotToken string
}

type Tweet struct {
	User     string `json:"user"`
	Message  string `json:"message"`
	Retweets int64  `json:"retweets"`
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

type part struct {
	start string
	end   byte
}

type Valute struct {
	ID_attr      string `xml:"ID,attr"`
	NumCode_test string `xml:"NumCode"`
	CharCode     string `xml:"CharCode"`
	Nominal      string `xml:"Nominal"`
	Name         string `xml:"Name"`
	Value        string `xml:"Value"`
}

//type Result struct {
//	XMLName   xml.Name `xml:"ValCurs"`
//	Date_attr string   `xml:"Date,attr"`
//	Name_attr string   `xml:"name,attr"`
//	Valute    []Valute
//}

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

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type NOTIFYICONDATA struct {
	CbSize           uint32
	HWnd             uintptr
	UID              uint32
	UFlags           uint32
	UCallbackMessage uint32
	HIcon            uintptr
	SzTip            [128]uint16
	DwState          uint32
	DwStateMask      uint32
	SzInfo           [256]uint16
	UVersion         uint32
	SzInfoTitle      [64]uint16
	DwInfoFlags      uint32
	GuidItem         GUID
	HBalloonIcon     uintptr
}

type Product struct {
	Id      bson.ObjectId `bson:"_id"`
	Model   string        `bson:"model"`
	Company string        `bson:"company"`
	Price   int           `bson:"price"`
}
type Result struct {
	Name, Description, URL string
}

//type SearchResults struct {
//	ready   bool
//	Query   string
//	Results []Result
//}

type SearchResults struct {
	name     string
	Results1 []Result
	Results2 []Result
	Results3 []Result
}

type phoneWriter struct{}

func (p phoneWriter) Write(bs []byte) (int, error) {
	if len(bs) == 0 {
		return 0, nil
	}
	for i := 0; i < len(bs); i++ {
		if bs[i] >= '0' && bs[i] <= '9' {
			fmt.Print(string(bs[i]))
		}
	}
	fmt.Println()
	return len(bs), nil
}

type person777 struct {
	name string
	age  int
}

func (p person777) print() {
	fmt.Println("Имя:", p.name)
	fmt.Println("Возраст:", p.age)
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

var (
	libshell32  = windows.NewLazySystemDLL("shell32.dll")
	libuser32   = windows.NewLazySystemDLL("user32.dll")
	libkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procShell_NotifyIconW = libshell32.NewProc("Shell_NotifyIconW")
	procLoadImageW        = libuser32.NewProc("LoadImageW")
	procRegisterClassExW  = libuser32.NewProc("RegisterClassExW")
	procGetModuleHandleW  = libkernel32.NewProc("GetModuleHandleW")
	procCreateWindowExW   = libuser32.NewProc("CreateWindowExW")
	procDefWindowProcW    = libuser32.NewProc("DefWindowProcW")
	procGetMessageW       = libuser32.NewProc("GetMessageW")
	procTranslateMessage  = libuser32.NewProc("TranslateMessage")
	procDispatchMessageW  = libuser32.NewProc("DispatchMessageW")
	procPostQuitMessage   = libuser32.NewProc("PostQuitMessage")
	procShowWindow        = libuser32.NewProc("ShowWindow")
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

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// var (
// 	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
// 		Name: "myapp_processed_ops_total_666",
// 		Help: "666 looser",
// 	})
// )

var addr = flag.String("listen-address", ":8080",
	"The address to listen on for HTTP requests.")

//func encodeRFC2047(String string) string {
//	// use mail's rfc2047 to encode any string
//	addr := mail.Address{String, ""}
//	return strings.Trim(addr.String(), " <>")
//}

var hash [16]byte

func urlEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func factorial(n int, ch chan int) {
	defer close(ch)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		ch <- result
	}
}

// 1
type longLatStruct struct {
	Long float64 `json:"longitude"`
	Lat  float64 `json:"latitude"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *longLatStruct)

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }
// var upgrader_del = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "home")
// }

func writer(coord *longLatStruct) {
	broadcast <- coord
}

func longLatHandler(w http.ResponseWriter, r *http.Request) {
	var coordinates longLatStruct
	if err := json.NewDecoder(r.Body).Decode(&coordinates); err != nil {
		log.Printf("ERROR: %s", err)
		http.Error(w, "Bad request", http.StatusTeapot)
		return
	}
	defer r.Body.Close()
	go writer(&coordinates)
}

// func wsHandler(w http.ResponseWriter, r *http.Request) {
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// register client
// 	clients[ws] = true
// }

// 3
// func echo() {
// 	for {
// 		val := <-broadcast
// 		latlong := fmt.Sprintf("%f %f %s", val.Lat, val.Long)
// 		// send to every client that is currently connected
// 		for client := range clients {
// 			err := client.WriteMessage(websocket.TextMessage, []byte(latlong))
// 			if err != nil {
// 				log.Printf("Websocket error: %s", err)
// 				client.Close()
// 				delete(clients, client)
// 			}
// 		}
// 	}
// }

// func recordMetrics() {
// 	go func() {
// 		for {
// 			opsProcessed.Inc()
// 			time.Sleep(2 * time.Second)
// 		}
// 	}()
// }

// func GetClient() *mongo.Client {
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:32787")
// 	client, err := mongo.NewClient(clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = client.Connect(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return client
// }

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

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}

	return false
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

func wikipediaAPI(request string) (answer []string) {

	//Создаем срез на 3 элемента
	s := make([]string, 3)

	//Отправляем запрос
	if response, err := http.Get(request); err != nil {
		s[0] = "Wikipedia is not respond"
		return s
		//fmt.Println("Wikipedia is not respond")
	} else {
		defer response.Body.Close()

		//Считываем ответ
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			//log.Fatal(err)
			//fmt.Println(err)
			s[0] = "Wikipedia is not respond"
			return s
		}

		//Отправляем данные в структуру
		//		sr := &SearchResults{}
		//		if err = json.Unmarshal([]byte(contents), sr); err != nil {
		//			s[0] = "Something going wrong, try to change your question"
		//			fmt.Println(err)
		//		}

		//Отправляем данные в структуру

		var sr []interface{}

		if err = json.Unmarshal([]byte(contents), &sr); err != nil {
			s[0] = "Something going wrong, try to change your question"
			//fmt.Println(err)
			return s
		}

		if len(sr) < 4 {
			s[0] = "Something going wrong, try to change your question"
			return s
		}

		keys, ok := sr[1].([]interface{})
		if !ok {
			s[0] = "Something going wrong, try to change your question"
			//fmt.Println(err)
			return s
		}
		values, ok := sr[2].([]interface{})
		if !ok {
			s[0] = "Something going wrong, try to change your question"
			//fmt.Println(err)
			return s
		}
		links, ok := sr[3].([]interface{})
		if !ok {
			s[0] = "Something going wrong, try to change your question"
			//fmt.Println(err)
			return s
		}

		if len(keys) != len(values) && len(values) != len(links) {
			s[0] = "Something going wrong, try to change your question"
			//fmt.Println(err)
			return s
		}

		// find accurate wiki pages
		for i := 0; i < len(keys); i++ {
			//key := keys[i].(string)
			val := values[i].(string)
			link := links[i].(string)

			//fmt.Println("key - ", key)
			fmt.Println("val - ", val)
			fmt.Println("link - ", link)
			fmt.Println(val + " " + link)

			//			if strings.Contains(val, "may refer to") {
			//				continue
			//			}
			//			if !strings.Contains(link, "en.wikipedia.org") {
			//				continue
			//			}

			s[i] = val + " " + link

			//			found := false
			//			for _, t := range tags {
			//				if strings.Contains(val, t) {
			//					found = true
			//					break
			//				}
			//			}
			//			if !found {
			//				continue
			//			}

			//			ws := WikiSearch{
			//				Key:   key,
			//				Value: val,
			//				Link:  link,
			//			}
			//			wSearch = append(wSearch, ws)
		}

		//		//fmt.Println("---Test---", sr[2])
		//		keys, ok := sr[2].([]interface{})

		//		if !ok {
		//			s[0] = "Something going wrong, try to change your question"
		//			return s
		//		}

		//		var counetr int = -1
		//		for _, x := range keys {
		//			switch value := x.(type) {
		//			case string:
		//				counetr++
		//				s[counetr] = value
		//			}
		//		}

		//		for _, x := range sr {
		//			switch value := x.(type) {
		//			case string:
		//				s[1] = value
		//			case []interface{}:
		//				var counetr int = -1
		//				for _, v := range value {
		//					//fmt.Println(v.(string))
		//					counetr++
		//					s[counetr] = v.(string)

		//				}
		//			}
		//		}

		//		for i := range sr[2] {
		//			s[i] = sr.Results1[i].URL
		//		}

	}

	return s
}

func Shell_NotifyIcon(
	dwMessage uint32,
	lpData *NOTIFYICONDATA) (int32, error) {
	r, _, err := procShell_NotifyIconW.Call(
		uintptr(dwMessage),
		uintptr(unsafe.Pointer(lpData)))
	if r == 0 {
		return 0, err
	}
	return int32(r), nil
}

func LoadImage(
	hInst uintptr,
	name *uint16,
	type_ uint32,
	cx, cy int32,
	fuLoad uint32) (uintptr, error) {
	r, _, err := procLoadImageW.Call(
		hInst,
		uintptr(unsafe.Pointer(name)),
		uintptr(type_),
		uintptr(cx),
		uintptr(cy),
		uintptr(fuLoad))
	if r == 0 {
		return 0, err
	}
	return r, nil
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

// worker по порядку сравнивает хэш каждого пароля с искомым
func worker(in <-chan part, out chan<- string) {
	var p part
	var b []byte
	for {
		p = <-in
		b = []byte(p.start)
		for b[0] != p.end {
			if md5.Sum(b) == hash {
				out <- string(b)
				return
			}
			nextPass(b)
		}
	}
}

func nextByte(b byte) byte {
	switch b {
	case 'z':
		return '0'
	case '9':
		return 'a'
	default:
		return b + 1
	}
}

func nextPass(b []byte) {
	for i := len(b) - 1; i >= 0; i-- {
		b[i] = nextByte(b[i])
		if b[i] != '0' {
			return
		}
	}
}

func generator(in chan<- part) {
	start := []byte("00000")
	var b byte
	for {
		b = nextByte(start[0])
		in <- part{string(start), b}
		start[0] = b
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

// func wsHandler(w http.ResponseWriter, r *http.Request) {
// 	// if r.Header.Get("Origin") != "http://"+r.Host {
// 	// 	http.Error(w, "Origin not allowed", 403)
// 	// 	return
// 	// }
// 	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
// 	if err != nil {
// 		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
// 	}

// 	go echo(conn)
// }

// func echo(conn *websocket.Conn) {
// 	for {
// 		m := msg{}

// 		err := conn.ReadJSON(&m)
// 		if err != nil {
// 			fmt.Println("Error reading json.", err)
// 		}

// 		fmt.Printf("Got message: %#v\n", m)

// 		if err = conn.WriteJSON(m); err != nil {
// 			fmt.Println(err)
// 		}
// 	}
// }

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

type msg struct {
	Num int
}

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
	// file, _ := os.Open("./dndspellsbot/config.json")
	// decoder := json.NewDecoder(file)
	// configuration := Config{}
	// err2 := decoder.Decode(&configuration)
	// if err2 != nil {
	// 	log.Panic(err2)
	// }
	// fmt.Println(configuration.TelegramBotToken)

	// //Стандартное обращение для получения тела страницы
	// //	resp, err := http.Get("https://www.google.ru")
	// //	if err != nil {
	// //		log.Fatal(err)
	// //	}

	// //	defer resp.Body.Close()
	// //	body, err := ioutil.ReadAll(resp.Body)
	// //	fmt.Println(string(body))
	// //	//io.Copy(os.Stdout, resp.Body)

	// //dialSocksProxy, err := proxy.SOCKS5("tcp", "103.197.26.243:8888", nil, proxy.Direct)
	// dialSocksProxy, err := proxy.SOCKS5("tcp", "88.99.149.206:9050", nil, proxy.Direct)

	// if err != nil {
	// 	fmt.Println("Error connecting to proxy:", err)
	// }

	// tr := &http.Transport{Dial: dialSocksProxy.Dial}

	// // Create client
	// myClient := &http.Client{
	// 	Transport: tr,
	// }

	// bot, err := tgbotapi.NewBotAPIWithClient("808741510:AAECEpVU9cLIdJ0HsHpNASlolDVWYACgyA4", myClient)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// bot.Debug = true
	// fmt.Println("Authorized on account %s", bot.Self.UserName)

	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60

	// updates, err := bot.GetUpdatesChan(u)

	// for update := range updates {
	// 	if update.Message == nil { // ignore any non-Message Updates
	// 		continue
	// 	}

	// 	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	// 	var reply string
	// 	var ChatID int64

	// 	ChatID = update.Message.Chat.ID

	// 	// Текст сообщения
	// 	Text := update.Message.Text

	// 	switch Text {
	// 	case "/start":

	// 		reply = "Hi, i'm a wikipedia bot, i can search information in a wikipedia, send me something what you want find in Wikipedia."

	// 		msg := tgbotapi.NewMessage(ChatID, reply)
	// 		msg.ReplyToMessageID = update.Message.MessageID

	// 		bot.Send(msg)

	// 	case "/number_of_users":

	// 		reply = "Не реализовано"

	// 		msg := tgbotapi.NewMessage(ChatID, reply)
	// 		msg.ReplyToMessageID = update.Message.MessageID

	// 		bot.Send(msg)

	// 	default:

	// 		//Проверка на тип не нужна т.к сейчас в бот не приходтя не текстовые команды... Они находятся в отдельных структурах update.Message.Photo
	// 		if reflect.TypeOf(update.Message.Text).Kind() != reflect.String {
	// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Use the words for search.")
	// 			bot.Send(msg)
	// 			continue
	// 		}

	// 		//				if update.Message.NewChatParticipant.UserName != "" {
	// 		//					// В чат вошел новый пользователь
	// 		//					// Поприветствуем его
	// 		//					reply = fmt.Sprintf(`Привет @%s! Я тут слежу за порядком. Веди себя хорошо.`,
	// 		//						update.Message.NewChatParticipant.UserName)
	// 		//				}

	// 		language := "ru"
	// 		url, err4 := urlEncoded(Text)
	// 		if err4 != nil {
	// 			fmt.Println(err4)
	// 		}
	// 		request := "https://" + language + ".wikipedia.org/w/api.php?action=opensearch&search=" + url + "&limit=3&origin=*&format=json"

	// 		message := wikipediaAPI(request)

	// 		for _, val := range message {

	// 			//Отправлем сообщение
	// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, val)
	// 			bot.Send(msg)
	// 			fmt.Println(val)
	// 		}
	// 	}

	// }

	//--------------------------------------Тест получения страницы с вики----------------------------------------------------------------------------------------
	//	//language := os.Getenv("LANGUAGE")
	//	language := "ru"
	//	url, err4 := urlEncoded("Москва")
	//	if err4 != nil {
	//		fmt.Println(err4)
	//	}
	//	request := "https://" + language + ".wikipedia.org/w/api.php?action=opensearch&search=" + url + "&limit=3&origin=*&format=json"

	//	message := wikipediaAPI(request)

	//	for _, val := range message {

	//		//Отправлем сообщение
	//		 msg := tgbotapi.NewMessage(update.Message.Chat.ID, val)
	//		 bot.Send(msg)
	//		fmt.Println(val)
	//	}

	//--------------- Конец работа с Telegram Bot через "github.com/Syfaro/telegram-bot-api" -----------------

	//--------------- Работа с Системным треем Tray Icons -----------------
	// var data NOTIFYICONDATA

	// data.CbSize = uint32(unsafe.Sizeof(data))
	// data.UFlags = NIF_ICON

	// icon, err := LoadImage(
	// 	0,
	// 	windows.StringToUTF16Ptr("icon.ico"),
	// 	IMAGE_ICON,
	// 	0,
	// 	0,
	// 	LR_DEFAULTSIZE|LR_LOADFROMFILE)
	// if err != nil {
	// 	panic(err)
	// }
	// data.HIcon = icon

	// if _, err := Shell_NotifyIcon(NIM_ADD, &data); err != nil {
	// 	panic(err)
	// }
	//--------------- Конец работы с Системным треем Tray Icons -----------------

	// //--------------- Работа с TCP scanner -----------------
	// hostname := flag.String("localhost", "", "hostname to test")
	// startPort := flag.Int("start-port", 1, "the port on which the scanning starts")
	// endPort := flag.Int("end-port", 10000, "the port from which the scanning ends")
	// timeout := flag.Duration("timeout", time.Millisecond*200, "timeout")
	// flag.Parse()

	// ports := []int{}

	// wg := &sync.WaitGroup{}
	// mutex := &sync.Mutex{}
	// for port := *startPort; port <= *endPort; port++ {
	// 	wg.Add(1)
	// 	go func(p int) {
	// 		opened := isOpen(*hostname, p, *timeout)
	// 		if opened {
	// 			mutex.Lock()
	// 			ports = append(ports, p)
	// 			mutex.Unlock()
	// 		}
	// 		wg.Done()
	// 	}(port)
	// }

	// wg.Wait()
	// fmt.Printf("opened ports: %v\n", ports)
	// //--------------- Конец Работа с TCP scanner -----------------

	// //--------------- Работа с UUIDs -----------------
	// fmt.Println(uuid.NewUUID()) // 2bf08894-47a3978-bbdd8c85-70d16847
	// //--------------- Конец Работа с UUIDs -----------------

	// //--------------- Работа с Regular Expressions -----------------
	// match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	// fmt.Println(match)

	// r, _ := regexp.Compile("p([a-z]+)ch")

	// fmt.Println(r.MatchString("peach"))

	// fmt.Println(r.FindString("peach punch"))

	// fmt.Println(r.FindStringIndex("peach punch"))

	// fmt.Println(r.FindStringSubmatch("peach punch"))

	// fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// fmt.Println(r.FindAllString("peach punch pinch", -1))

	// fmt.Println(r.FindAllStringSubmatchIndex(
	// 	"peach punch pinch", -1))

	// fmt.Println(r.FindAllString("peach punch pinch", 2))

	// fmt.Println(r.Match([]byte("peach")))

	// r = regexp.MustCompile("p([a-z]+)ch")
	// fmt.Println(r)

	// fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// in := []byte("a peach")
	// out := r.ReplaceAllFunc(in, bytes.ToUpper)
	// fmt.Println(string(out))
	// //--------------- Конец Работа с Regular Expressions -----------------

	// //--------------- Работа с Моим же репозиторием на GitHub -----------------
	// Необходим модуль где в названии модуля используется название переменной, которая становится доступной после импорта моего репозитория
	//main777.
	// //--------------- Конец Работа с Моим же репозиторием на GitHub-----------------

	// //--------------- Работа с Prometheus -----------------

	// flag.Parse()

	// usersRegistered := prometheus.NewCounter(
	// 	prometheus.CounterOpts{
	// 		Name: "users_registered",
	// 	})
	// prometheus.MustRegister(usersRegistered)

	// usersOnline := prometheus.NewGauge(
	// 	prometheus.GaugeOpts{
	// 		Name: "users_online",
	// 	})
	// prometheus.MustRegister(usersOnline)

	// requestProcessingTimeSummaryMs := prometheus.NewSummary(
	// 	prometheus.SummaryOpts{
	// 		Name:       "request_processing_time_summary_ms",
	// 		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	// 	})
	// prometheus.MustRegister(requestProcessingTimeSummaryMs)

	// requestProcessingTimeHistogramMs := prometheus.NewHistogram(
	// 	prometheus.HistogramOpts{
	// 		Name:    "request_processing_time_histogram_ms",
	// 		Buckets: prometheus.LinearBuckets(0, 10, 20),
	// 	})
	// prometheus.MustRegister(requestProcessingTimeHistogramMs)

	// go func() {
	// 	for {
	// 		usersRegistered.Inc() // or: Add(5)
	// 		time.Sleep(1000 * time.Millisecond)
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		for i := 0; i < 10000; i++ {
	// 			usersOnline.Set(float64(i)) // or: Inc(), Dec(), Add(5), Dec(5)
	// 			time.Sleep(10 * time.Millisecond)
	// 		}
	// 	}
	// }()

	// go func() {
	// 	src := rand.NewSource(time.Now().UnixNano())
	// 	rnd := rand.New(src)
	// 	for {
	// 		obs := float64(100 + rnd.Intn(30))
	// 		requestProcessingTimeSummaryMs.Observe(obs)
	// 		requestProcessingTimeHistogramMs.Observe(obs)
	// 		time.Sleep(10 * time.Millisecond)
	// 	}
	// }()

	// http.Handle("/metrics", promhttp.Handler())

	// log.Printf("Starting web server at %s\n", *addr)
	// err := http.ListenAndServe(*addr, nil)
	// if err != nil {
	// 	log.Printf("http.ListenAndServer: %v\n", err)
	// }

	// //--------------- Конец Работа с Prometheus -----------------

	// //--------------- Работа с многопоточность и паролями -----------------
	// t := time.Now()
	// const hashString = "95ebc3c7b3b9f1d2c40fec14415d3cb8" // "zzzzz"
	// h, err := hex.DecodeString(hashString)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// copy(hash[:], h)

	// num := runtime.NumCPU()
	// runtime.GOMAXPROCS(num)

	// in := make(chan part)
	// out := make(chan string)
	// go generator(in)
	// for i := 0; i < num; i++ {
	// 	go worker(in, out)
	// }
	// fmt.Println("Пароль: ", <-out)
	// fmt.Println("Время поиска: ", time.Since(t))

	// copy(hash[:], h)

	// num := runtime.NumCPU()
	// runtime.GOMAXPROCS(num)

	// in := make(chan part)
	// out := make(chan string)
	// go generator(in)
	// for i := 0; i < num; i++ {
	// 	go worker(in, out)
	// }
	// fmt.Println("Пароль: ", <-out)
	// fmt.Println("Время поиска: ", time.Since(t))
	// //--------------- Конец Работа с многопоточность и паролями -----------------

	// //--------------------- Работа с CouchDB NoSQL "github.com/go-kivik/kivik" ---------------------------

	// fmt.Println("Test debug visual studio")

	// client, err := kivik.New("couch", "http://localhost:5984/")
	// if err != nil {
	// 	panic(err)
	// }

	// db := client.DB(context.TODO(), "test777")
	// if err != nil {
	// 	panic(err)
	// }

	// doc := map[string]interface{}{
	// 	"_id":      "cow5",
	// 	"feet":     4,
	// 	"greeting": "moo",
	// }

	// rev, err := db.Put(context.TODO(), "cow5", doc)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Cow inserted with revision %s\n", rev)

	// //--------------------- Конец Работа с CouchDB NoSQL "github.com/go-kivik/kivik" ---------------------------

	// //--------------------- Работа с rabbitmq через "github.com/streadway/amqp" --------------------------------

	// conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	// if err != nil {
	// 	fmt.Println("Failed to connect to RabbitMQ")
	// 	panic(err)
	// }
	// defer conn.Close()

	// ch, err := conn.Channel()
	// if err != nil {
	// 	fmt.Println("Failed to open a channel")
	// 	panic(err)
	// }
	// defer ch.Close()

	// //Sending
	// // q, err := ch.QueueDeclare(
	// // 	"hello", // name
	// // 	false,   // durable
	// // 	false,   // delete when unused
	// // 	false,   // exclusive
	// // 	false,   // no-wait
	// // 	nil,     // arguments
	// // )
	// // if err != nil {
	// // 	fmt.Println("Failed to declare a queue")
	// // 	panic(err)
	// // }

	// // body := "Hello World!"
	// // err = ch.Publish(
	// // 	"",     // exchange
	// // 	q.Name, // routing key
	// // 	false,  // mandatory
	// // 	false,  // immediate
	// // 	amqp.Publishing{
	// // 		ContentType: "text/plain",
	// // 		Body:        []byte(body),
	// // 	})

	// // if err != nil {
	// // 	fmt.Println("Failed to publish a message")
	// // 	panic(err)
	// // }

	// //Receiving

	// q, err := ch.QueueDeclare(
	// 	"hello", // name
	// 	false,   // durable
	// 	false,   // delete when unused
	// 	false,   // exclusive
	// 	false,   // no-wait
	// 	nil,     // arguments
	// )
	// if err != nil {
	// 	fmt.Println("Failed to declare a queue")
	// 	panic(err)
	// }

	// msgs, err := ch.Consume(
	// 	q.Name, // queue
	// 	"",     // consumer
	// 	true,   // auto-ack
	// 	false,  // exclusive
	// 	false,  // no-local
	// 	false,  // no-wait
	// 	nil,    // args
	// )

	// if err != nil {
	// 	fmt.Println("Failed to register a consumer")
	// 	panic(err)
	// }

	// forever := make(chan bool)

	// go func() {
	// 	forever <- true
	// 	for d := range msgs {
	// 		log.Printf("Received a message: %s", d.Body)
	// 		fmt.Println("Received a message: %s", d.Body)
	// 		forever <- false
	// 	}
	// 	forever <- false
	// }()

	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// //<-forever
	// fmt.Println(<-forever)

	// //--------------------- Конец работа с rabbitmq через "github.com/streadway/amqp" --------------------------------

	//--------------------- Работа с ElasticSerch через "gopkg.in/olivere/elastic.v6" --------------------------------
	// Create a client and connect to http://192.168.2.10:9201
	//client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:32784"))

	// client, err := elastic.NewClient(elastic.SetSniff(false),
	// 	elastic.SetURL("http://127.0.0.1:32784", "http://127.0.0.1:32783"))
	// // elastic.SetBasicAuth("user", "secret"))
	// if err != nil {
	// 	//fmt.Println("1")
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	// // Getting the ES version number is quite common, so there's a shortcut
	// esversion, err := client.ElasticsearchVersion("http://127.0.0.1:32784")
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }
	// fmt.Printf("Elasticsearch version %s\n", esversion)

	// exists, err := client.IndexExists("twitter").Do(context.Background()) //twitter
	// if err != nil {
	// 	//fmt.Println("2")
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// // if !exists {
	// // 	//fmt.Println("3")
	// // 	fmt.Println(err)
	// // 	panic(err)
	// // }

	// if !exists {
	// 	// Create a new index.
	// 	mapping := `
	// {
	// 	"settings":{
	// 		"number_of_shards":1,
	// 		"number_of_replicas":0
	// 	},
	// 	"mappings":{
	// 		"doc":{
	// 			"properties":{
	// 				"user":{
	// 					"type":"keyword"
	// 				},
	// 				"message":{
	// 					"type":"text",
	// 					"store": true,
	// 					"fielddata": true
	// 				},
	// 			"retweets":{
	// 				"type":"long"
	// 			},
	// 				"tags":{
	// 					"type":"keyword"
	// 				},
	// 				"location":{
	// 					"type":"geo_point"
	// 				},
	// 				"suggest_field":{
	// 					"type":"completion"
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	// `
	// 	createIndex, err := client.CreateIndex("twitter").Body(mapping).IncludeTypeName(true).Do(context.Background())
	// 	if err != nil {
	// 		// Handle error
	// 		panic(err)
	// 	}
	// 	if !createIndex.Acknowledged {
	// 		// Not acknowledged
	// 	}
	// }

	// // Index a tweet (using JSON serialization)
	// tweet1 := Tweet{User: "olivere", Message: "Take Five", Retweets: 0}
	// put1, err := client.Index().
	// 	Index("twitter").
	// 	Type("doc").
	// 	Id("1").
	// 	BodyJson(tweet1).
	// 	Do(context.Background())
	// if err != nil {
	// 	// Handle error
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	// // Index a second tweet (by string)
	// tweet2 := `{"user" : "olivere", "message" : "It's a Raggy Waltz"}`
	// put2, err := client.Index().
	// 	Index("twitter").
	// 	Type("doc").
	// 	Id("2").
	// 	BodyString(tweet2).
	// 	Do(context.Background())
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }
	// fmt.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)

	// // Get tweet with specified ID
	// get1, err := client.Get().
	// 	Index("twitter").
	// 	Type("doc").
	// 	Id("1").
	// 	Do(context.Background())
	// if err != nil {
	// 	switch {
	// 	case elastic.IsNotFound(err):
	// 		panic(fmt.Sprintf("Document not found: %v", err))
	// 	case elastic.IsTimeout(err):
	// 		panic(fmt.Sprintf("Timeout retrieving document: %v", err))
	// 	case elastic.IsConnErr(err):
	// 		panic(fmt.Sprintf("Connection problem: %v", err))
	// 	default:
	// 		// Some other kind of error
	// 		panic(err)
	// 	}
	// }
	// fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)

	// // Flush to make sure the documents got written.
	// _, err = client.Flush().Index("twitter").Do(context.Background())
	// if err != nil {
	// 	panic(err)
	// }

	// // Search with a term query
	// termQuery := elastic.NewTermQuery("user", "olivere")
	// searchResult, err := client.Search().
	// 	Index("twitter").        // search in index "twitter"
	// 	Query(termQuery).        // specify the query
	// 	Sort("user", true).      // sort by "user" field, ascending
	// 	From(0).Size(10).        // take documents 0-9
	// 	Pretty(true).            // pretty print request and response JSON
	// 	Do(context.Background()) // execute
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }

	// // searchResult is of type SearchResult and returns hits, suggestions,
	// // and all kinds of other information from Elasticsearch.
	// fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// // Each is a convenience function that iterates over hits in a search result.
	// // It makes sure you don't need to check for nil values in the response.
	// // However, it ignores errors in serialization. If you want full control
	// // over iterating the hits, see below.
	// var ttyp Tweet
	// for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
	// 	t := item.(Tweet)
	// 	fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	// }
	// // TotalHits is another convenience function that works even when something goes wrong.
	// fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	// // Here's how you iterate through results with full control over each step.
	// if searchResult.Hits.TotalHits > 0 {
	// 	fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

	// 	// Iterate through results
	// 	for _, hit := range searchResult.Hits.Hits {
	// 		// hit.Index contains the name of the index

	// 		// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
	// 		var t Tweet
	// 		err := json.Unmarshal(*hit.Source, &t)
	// 		if err != nil {
	// 			// Deserialization failed
	// 		}

	// 		// Work with tweet
	// 		fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	// 	}
	// } else {
	// 	// No hits
	// 	fmt.Print("Found no tweets\n")
	// }

	// // Update a tweet by the update API of Elasticsearch.
	// // We just increment the number of retweets.
	// script := elastic.NewScript("ctx._source.retweets += params.num").Param("num", 1)
	// update, err := client.Update().Index("twitter").Type("doc").Id("1").
	// 	Script(script).
	// 	Upsert(map[string]interface{}{"retweets": 0}).
	// 	Do(context.Background())
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }
	// fmt.Printf("New version of tweet %q is now %d", update.Id, update.Version)

	// // ...

	// // Delete an index.
	// deleteIndex, err := client.DeleteIndex("twitter").Do(context.Background())
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }
	// if !deleteIndex.Acknowledged {
	// 	// Not acknowledged
	// }

	//--------------------- Конец Работа с ElasticSerch через "gopkg.in/olivere/elastic.v6" --------------------------------

	// //--------------------- Работа с MongoDB --------------------------------
	// c := GetClient()
	// err := c.Ping(context.Background(), readpref.Primary())
	// if err != nil {
	// 	log.Fatal("Couldn't connect to the database", err)
	// } else {
	// 	log.Println("Connected!")
	// }
	// //--------------------- Конец Работа с MongoDB" --------------------------------

	// //--------------------- Работа с MongoDB через gopkg.in/mgo.v2 --------------------------------
	// //-- mgo на данный момент не поддерживается автором после того как вышел официальный mongodb-драйвер.
	// // открываем соединение
	// session, err := mgo.Dial("mongodb://127.0.0.1:32787")
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()

	// // получаем коллекцию
	// productCollection := session.DB("productdb").C("products")

	// // p1 := &Product{Id: bson.NewObjectId(), Model: "iPhone 8", Company: "Apple", Price: 64567}
	// // // добавляем один объект
	// // err = productCollection.Insert(p1)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }

	// // p2 := &Product{Id: bson.NewObjectId(), Model: "Pixel 2", Company: "Google", Price: 58000}
	// // p3 := &Product{Id: bson.NewObjectId(), Model: "Xplay7", Company: "Vivo", Price: 49560}
	// // // добавляем два объекта
	// // err = productCollection.Insert(p2, p3)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }
	// // // получаем коллекцию
	// // productCollection := session.DB("productdb").C("products")
	// // критерий выборки
	// query := bson.M{}
	// // объект для сохранения результата
	// products := []Product{}
	// productCollection.Find(query).All(&products)

	// for _, p := range products {

	// 	fmt.Println(p.Model, p.Company, p.Price)
	// }

	// // удаляем все документы с company = "Vivo"
	// _, err = productCollection.RemoveAll(bson.M{"company": "Vivo"})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// //--------------------- Конец Работа с MongoDB через gopkg.in/mgo.v2 --------------------------------

	//--------------------- Работа с Graphite через github.com/marpaia/graphite-golang --------------------------------

	// // try to connect a graphite server
	// Graphite, err := graphite.NewGraphite("localhost", 32796) //127.0.0.1 Port 2003

	// // if you couldn't connect to graphite, use a nop
	// if err != nil {
	// 	Graphite = graphite.NewGraphiteNop("localhost", 32796)
	// }

	// log.Printf("Loaded Graphite connection: %#v", Graphite)
	// //Graphite.SimpleSend("stats.graphite_loaded777", "1")
	// Graphite.SimpleSend("graphite_loaded777.777", "50")
	//--------------------- Конец Работа с Graphite через github.com/marpaia/graphite-golang --------------------------------

	//--------------------- Работа с influxdb через "github.com/influxdata/influxdb-client-go" --------------------------------
	// // Make client
	// c, err := client.NewHTTPClient(client.HTTPConfig{
	// 	Addr: "http://localhost:32860",
	// })
	// if err != nil {
	// 	fmt.Println("Error creating InfluxDB Client: ", err.Error())
	// }
	// defer c.Close()

	// q := client.NewQuery("CREATE DATABASE telegraf", "", "")
	// if response, err := c.Query(q); err == nil && response.Error() == nil {
	// 	fmt.Println(response.Results)
	// }

	// // q2 := client.NewQuery("SELECT count(value) FROM cpu_load", "mydb", "")
	// // if response, err := c.Query(q2); err == nil && response.Error() == nil {
	// // 	fmt.Println(response.Results)
	// // }
	// //q2 := client.NewQuery("SELECT count(value) FROM weather", "test_golang", "")
	// q2 := client.NewQuery("SELECT * FROM weather", "test_golang", "")
	// response, err := c.Query(q2)

	// if err != nil && response.Error() != nil {
	// 	fmt.Println("Error creating InfluxDB Client: ", err.Error())
	// } else {
	// 	for _, value := range response.Results {
	// 		if value.Err != "" {
	// 			fmt.Println("Error body response InfluxDB Client: ", value.Err)
	// 		} else {
	// 			fmt.Println("Body response InfluxDB Client: ", value.Messages)
	// 		}

	// 	}
	// 	//fmt.Println("Body response InfluxDB Client: ", response.Results)
	// }

	//--------------------- Конец Работа с influxdb  через "github.com/influxdata/influxdb-client-go" --------------------------------

	//------------------------------------------------------------- Работа с Kafca ----------------------------------------------------------------

	// // make a writer that produces to topic-A, using the least-bytes distribution
	// w := kafka.NewWriter(kafka.WriterConfig{
	// 	Brokers:  []string{"localhost:32783"},
	// 	Topic:    "topic-A",
	// 	Balancer: &kafka.LeastBytes{},
	// })

	// w.WriteMessages(context.Background(),
	// 	kafka.Message{
	// 		Key:   []byte("Key-A"),
	// 		Value: []byte("Hello World!"),
	// 	},
	// 	kafka.Message{
	// 		Key:   []byte("Key-B"),
	// 		Value: []byte("One!"),
	// 	},
	// 	kafka.Message{
	// 		Key:   []byte("Key-C"),
	// 		Value: []byte("Two!"),
	// 	},
	// )
	// w.Close()

	// //make a new reader that consumes from topic-A, partition 0, at offset 42
	// r := kafka.NewReader(kafka.ReaderConfig{
	// 	Brokers:   []string{"localhost:32780"},
	// 	Topic:     "topic-A",
	// 	Partition: 0,
	// 	MinBytes:  10e3, // 10KB
	// 	MaxBytes:  10e6, // 10MB
	// })
	// r.SetOffset(42)

	// for {
	// 	m, err := r.ReadMessage(context.Background())
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// 	fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	// }
	// r.Close()

	//------------------------------------------------------------ Конец Работа с Kafca -----------------------------------------------------------

	//------------------------------------------------------------- Работа с Redis ----------------------------------------------------------------
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:32785",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })

	// pong, err := client.Ping().Result()
	// fmt.Println(pong, err)
	// // Output: PONG <nil>

	// err = client.Set("1c", "Говно---", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := client.Get("1c").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	// val2, err := client.Get("key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
	// // Output: key value
	// // key2 does not exist

	//------------------------------------------------------------- Конец Работа с Redis ----------------------------------------------------------------

	//------------------------------------------------------------- Работа с Tarantool через "github.com/tarantool/go-tarantool" ----------------------------------------------------------------
	//var spaceNo uint32
	//spaceNo = uint32(999)

	//spaceNo := uint32(512)

	//spaceNo = uint32(999)
	//indexNo := uint32(0)

	// server := "127.0.0.1:32787"
	// opts := tarantool.Opts{
	// 	Timeout:       500 * time.Millisecond,
	// 	Reconnect:     1 * time.Second,
	// 	MaxReconnects: 3,
	// 	User:          "guest",
	// 	Pass:          "",
	// }
	// client, err := tarantool.Connect(server, opts)

	// if err != nil {
	// 	log.Fatalf("Failed to connect: %s", err.Error())
	// }

	// resp, err := client.Ping()
	// log.Println(resp.Code)
	// log.Println(resp.Data)
	// log.Println(err)

	// //resp, err = client.Insert(spaceNo, []interface{}{uint(10), 1})
	// resp, err = client.Insert("example-1.0", []interface{}{uint(10), 1})
	// log.Println("Insert")
	// log.Println("Error", err)
	// log.Println("Code", resp.Code)
	// log.Println("Data", resp.Data)

	//------------------------------------------------------------- Конец Работа с Tarantool через "github.com/tarantool/go-tarantool" ----------------------------------------------------------------

	//------------------------------------------------------------- Работа с Cassandra  через "github.com/gocql/gocql" ----------------------------------------------------------------
	// // connect to the cluster
	// cluster := gocql.NewCluster("127.0.0.1")
	// cluster.Keyspace = "example"
	// cluster.Consistency = gocql.Quorum
	// cluster.Port = 32789 //9042
	// session, _ := cluster.CreateSession()
	// defer session.Close()

	// // insert a tweet
	// if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
	// 	"me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
	// 	log.Fatal(err)
	// }

	// var id gocql.UUID
	// var text string

	// /* Search for a specific set of records whose 'timeline' column matches
	//  * the value 'me'. The secondary index that we created earlier will be
	//  * used for optimizing the search */
	// if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
	// 	"me").Consistency(gocql.One).Scan(&id, &text); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Tweet:", id, text)

	// // list all tweets
	// iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
	// for iter.Scan(&id, &text) {
	// 	fmt.Println("Tweet:", id, text)
	// }
	// if err := iter.Close(); err != nil {
	// 	log.Fatal(err)
	// }
	//------------------------------------------------------------- Конец Работа с RethinkDB  через "gopkg.in/rethinkdb/rethinkdb-go.v6" ----------------------------------------------------------------
	// session, err := r.Connect(r.ConnectOpts{
	// 	Address: "localhost:32796", //28015
	// 	//Database: "test",
	// 	//AuthKey:  "14daak1cad13dj",
	// })

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// res, err := r.Expr("Hello World").Run(session)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// var response string
	// err = res.One(&response)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(response)

	// err = r.DB("test").Table("test_golang").Insert(map[string]string{
	// 	"id":       "john777",
	// 	"password": "777_p455w0rd",
	// }).Exec(session)

	// res, err := r.DB("test").Table("test_golang").Get("john777").Run(session)
	// if err != nil {
	// 	// error
	// }
	// defer res.Close() // Always ensure you close the cursor to ensure connections are not leaked
	// for _, value := range res.responses {
	// 	fmt.Println(string(value))
	// }
	//------------------------------------------------------------- Конец Работа с Cassandra  через "gopkg.in/rethinkdb/rethinkdb-go.v6" ----------------------------------------------------------------

	//------------------------------------------------------------- Работа с ClickHouse  через "github.com/roistat/go-clickhouse" ----------------------------------------------------------------
	// transport := clickhouse.NewHttpTransport()
	// conn := clickhouse.NewConn("localhost:32770", transport) //8123
	// err := conn.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// q := clickhouse.NewQuery("INSERT INTO logs VALUES (toDate(now()), ?, ?)", 1, "Log message777") //q := clickhouse.NewQuery("INSERT INTO golang_test.logs VALUES (toDate(now()), ?, ?)", 1, "Log message")
	// q.Exec(conn)

	// q := clickhouse.NewQuery("SELECT `message` FROM logs") //q = clickhouse.NewQuery("SELECT `message` FROM golang_test.logs")
	// i := q.Iter(conn)
	// for {
	// 	var message string
	// 	scanned := i.Scan(&message)
	// 	if scanned {
	// 		fmt.Println(message)
	// 	} else if err != nil {
	// 		panic(i.Error())
	// 	}
	// }

	//------------------------------------------------------------- Конец Работа с ClickHouse  через "github.com/roistat/go-clickhouse" ----------------------------------------------------------------

	//------------------------------------------------------------- Работа с ActiveMQ  через "github.com/go-stomp/stomp" ----------------------------------------------------------------
	// conn, err := stomp.Dial("tcp", "localhost:32788") //61613
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = conn.Send(
	// 	"/queue/test-777",            // destination
	// 	"text/plain",                 // content-type
	// 	[]byte("Test7 message #777")) // body
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// //And to get the message ?
	// // conn, err := stomp.Dial("tcp", "localhost:32788") //61613
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }
	// sub, err := conn.Subscribe("/queue/test-777", stomp.AckClient)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for {
	// 	msg := <-sub.C
	// 	fmt.Println(string(msg.Body))

	// 	// acknowledge the message
	// 	err = conn.Ack(msg)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	//------------------------------------------------------------- Конец Работа с ActiveMQ  через "github.com/go-stomp/stomp" ----------------------------------------------------------------

	//------------------------------------------------------------- Работа с NATS через "github.com/nats-io/nats.go" ----------------------------------------------------------------
	// // Connect to a server
	// natsClient, err := nats.Connect("nats://localhost:32785") // nats://derek:pass@localhost:4222
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// myUser := Transport.User{}
	// myUser.Name = "556"
	// data, err := proto.Marshal(&myUser)
	// fmt.Println(string(data))

	//  myUser2 := Transport.User{}
	//  err = proto.Unmarshal(data, &myUser2)

	// // message := &pb.PublishPostMessage{
	// // 	Title:   pubReq.Title,
	// // 	Content: pubReq.Content,
	// // }

	// // message.Title = "100"
	// // message.Content = "1007"

	// bs, err := proto.Marshal(&myUser)
	// err = natsClient.Publish("posts:publish", bs)

	// err = natsClient.Flush()
	// err = natsClient.LastError()

	// // Connect to a server
	// natsClient2, err2 := nats.Connect("nats://localhost:32785") // nats://derek:pass@localhost:4222
	// if err != nil {
	// 	fmt.Println(err2)
	// }

	// _, err = natsClient2.Subscribe("posts:publish", func(msg *nats.Msg) {
	// 	fmt.Printf("Received a message: %s\n", string(msg.Data))
	// 	fmt.Println("8888")
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// }

	//------------------------------------------------------------- Конец Работа с NATS через "github.com/nats-io/nats.go" ----------------------------------------------------------------

	//------------------------------------------------------------- Работа с gRPC -------------------------------------------
	// listener, err := net.Listen("tcp", ":5300")

	// if err != nil {
	// 	grpclog.Fatalf("failed to listen: %v", err)
	// }

	// opts := []grpc.ServerOption{}
	// grpcServer := grpc.NewServer(opts...)

	// fmt.Println("Start 1C Port 5300")

	// pb.RegisterReverseServer(grpcServer, &server{})
	// grpcServer.Serve(listener)

	//------------------------------------------------------------- Конец Работа с gRPC -------------------------------------------

	//------------------------------------------------------------- Работа с websocket -------------------------------------------
	// router := mux.NewRouter()
	// router.HandleFunc("/", rootHandler).Methods("GET")
	// router.HandleFunc("/longlat", longLatHandler).Methods("POST")
	// router.HandleFunc("/ws", wsHandler)
	// go echo()

	// log.Fatal(http.ListenAndServe(":8844", router))

	//------------------------------------------------------------- Работа с websocket 2 вариант -----------------------------------------

	// http.HandleFunc("/ws", wsHandler)
	// http.HandleFunc("/", rootHandler)

	// panic(http.ListenAndServe(":8088", nil))

	//------------------------------------------------------------- Работа с websocket 3 вариант ------------------------

	// http.Handle("/", websocket.Handler(Echo))

	// if err := http.ListenAndServe(":1234", nil); err != nil {
	// 	log.Fatal("ListenAndServe:", err)
	// }

	//------------------------------------------------------------- Конец работа с websocket -------------------------------------------

	//------------------------------------------------------------- Работа с aerospike через с github.com/aerospike/aerospike-client-go" ------------------------------------------

	// client, err := as.NewClient("127.0.0.1", 32794) //3000

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // // Write single value.
	// key, _ := as.NewKey("test", "myset", "mykey")
	// // // bin := as.NewBin("mybin", "myvalue")

	// // // var wpolicy = as.NewWritePolicy(0, 0)
	// // // client.PutBins(wpolicy, key, bin)

	// // // var rpolicy = as.NewPolicy()
	// // // exists, err := client.Exists(rpolicy, key)

	// // if exists != true {
	// // 	fmt.Println(exists)
	// // } else {
	// // 	fmt.Println(exists)
	// // }

	// // Read a record
	// record, err := client.Get(nil, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // for _, val := range record {
	// // 	fmt.Println(record)
	// // }

	// client.Close()

	//------------------------------------------------------------- Конец работа с aerospike через с github.com/aerospike/aerospike-client-go" --------------------------------------

	http.HandleFunc("/", indexPage)
	//http.HandleFunc("/products", ProductsHandler)
	fmt.Println("Start 1C Port 8081")
	http.ListenAndServe(":8081", nil)

}

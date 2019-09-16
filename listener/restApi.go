package listener

import (
	"fmt"
	"../logger"
	"net/http"
	"encoding/json"
	"path"
	"os"
)

type imageJson struct {
	Name string
}

var image imageJson


func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong", r.URL.Path[1:])
}

func serveImg(w http.ResponseWriter, r *http.Request) {
	if(len(image.Name) > 0) {
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Last-Modified", lastModifiedTime(image.Name))
		http.ServeFile(w, r, path.Join(image.Name))
		return
	}
	fmt.Fprintf(w, "image not found", r.URL.Path[1:])
}

func lastModifiedTime(filename string) string {
	log := logger.GetInstance()
	file, err := os.Stat(filename)
	if err != nil {
		log.Println(err)
	}
 
	return file.ModTime().String()
}

func snapshot(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&image)

	if err != nil {
		image.Name =""
		fmt.Println("Unable to get latest image file name")
	}
}

//StartServer starting serve images
func StartServer(port string) {
	log := logger.GetInstance()
	http.HandleFunc("/", serveImg)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/snapshot", snapshot)
	listenTo := ":" + port
	log.Println("listen to http on port", port)
	log.Println(http.ListenAndServe(listenTo, nil))
}

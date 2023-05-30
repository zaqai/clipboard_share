package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nutsdb/nutsdb"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var ntfyAddr string

func main() {

	defaultPort := "9090"
	defaultNtfyAddr := ""
	port := flag.String("port", defaultPort, "HTTP server port")
	flag.StringVar(&ntfyAddr, "ntfyAddr", defaultNtfyAddr, "ntfy address")
	flag.Parse()

	if envPort := os.Getenv("PORT"); envPort != "" {
		*port = envPort
	}
	if envNtfyAddr := os.Getenv("NTFYADDR"); envNtfyAddr != "" {
		ntfyAddr = envNtfyAddr
	}

	http.HandleFunc("/postq", pushData)
	http.HandleFunc("/idxq", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/", pullData)

	s := &http.Server{
		Addr:           ":" + *port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Starting server on port " + *port + ", ntfy address: " + ntfyAddr)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("listenAndServe: ", err)
	}

}

type ReqData struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func pushData(w http.ResponseWriter, r *http.Request) {
	var reqdata ReqData
	// 调用json包的解析，解析请求body
	if err := json.NewDecoder(r.Body).Decode(&reqdata); err != nil {
		r.Body.Close()
		log.Println(err, r.Body)
	}
	DBKey := reqdata.Key
	syncNtfy(reqdata.Value)
	buf := new(bytes.Buffer)
	//gob编码
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(reqdata); err != nil {
		fmt.Println(err)
	}
	DBValue := buf.Bytes()
	log.Println("write: ", DBKey, reqdata)
	writeDB(DBKey, DBValue)
	result := readDB(DBKey)

	dec := gob.NewDecoder(bytes.NewBuffer(result))
	var reqdata1 ReqData
	if err := dec.Decode(&reqdata1); err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.Write([]byte(reqdata1.Value))
}
func pullData(w http.ResponseWriter, r *http.Request) {

	DBKey := r.URL.Path[1:]
	DBValueByte := readDB(DBKey)
	dec := gob.NewDecoder(bytes.NewBuffer(DBValueByte))
	var DBValue ReqData
	if err := dec.Decode(&DBValue); err != nil {
		fmt.Println(err)
	}

	log.Println("read: ", DBKey, DBValue)

	if DBValue.Type == "text" {
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.Write([]byte(DBValue.Value))
	} else {
		http.Redirect(w, r, DBValue.Value, http.StatusFound)
	}

}

func writeDB(k string, v []byte) {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *nutsdb.Tx) error {
		if err := tx.Put("", []byte(k), v, 0); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println(err, k, v)
	}
}
func readDB(k string) []byte {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var value []byte
	err = db.View(func(tx *nutsdb.Tx) error {
		key := []byte(k)
		e, err := tx.Get("", key)
		if err != nil {
			return err
		} else {
			value = e.Value
		}
		return nil
	})
	if err != nil {
		log.Println(err, k)
	}
	return value
}

func syncNtfy(m string) {
	req, _ := http.NewRequest("POST", ntfyAddr,
		strings.NewReader(m))
	req.Header.Set("Title", "from clipboard_share")
	http.DefaultClient.Do(req)
}

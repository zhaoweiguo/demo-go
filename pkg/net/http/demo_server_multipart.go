package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

//对应client: net/http/demo_client_file_upload
// 请求结果查看文件./tmp/file1.new和./tmp/file2.new
func main() {
	port := ":8844"
	http.HandleFunc("/multipart", uploadHandler)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Panic("ListenAndServe err:", err)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entry of handle\n")
	reader, err := r.MultipartReader()
	if err != nil {
		writeErrorResponse(http.StatusBadRequest, err, w)
		return
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			log.Printf("==== EOF\n")
			break
		} else if err != nil {
			log.Printf("==== ERROR\n")
			writeErrorResponse(http.StatusBadRequest, err, w)
			return
		} else {
			log.Println(part.Header)
		}

		filename := part.FileName()
		formname := part.FormName()

		log.Printf("==== filename=[%s], formname=[%s]\n", filename, formname)
		if filename != "" { // this is FileData
			dst, err := os.Create("./tmp/" + filename + ".new")
			if err != nil {
				log.Println("")
				writeErrorResponse(http.StatusBadRequest, err, w)
				return
			}
			defer dst.Close()
			_, err = io.Copy(dst, part)
			if err != nil {
				log.Println("")
				dst.Close()
				writeErrorResponse(http.StatusBadRequest, err, w)
				return
			}
		} else if formname == "title" { // this is FormData
			data, err := ioutil.ReadAll(part)
			if err != nil {
				log.Println("")
				writeErrorResponse(http.StatusBadRequest, err, w)
				return
			}
			log.Printf("formname title->[%s]=[%s]\n", formname, string(data))
		} else { // other FormData
			data, err := ioutil.ReadAll(part)
			if err != nil {
				log.Println("")
				writeErrorResponse(http.StatusBadRequest, err, w)
				return
			}
			log.Printf("formname[%s]=[%s]\n", formname, string(data))
		}
	}

	writeResponse(http.StatusAccepted, nil, w)
}

func writeErrorResponse(code int, err error, w http.ResponseWriter) {
	log.Println(code, err)
	response := map[string]string{"error": err.Error()}
	writeResponse(code, response, w)
}
func writeResponse(code int, jsonres interface{}, w http.ResponseWriter) {
	log.Println(code, jsonres)
	b, err := json.Marshal(jsonres)
	if err != nil {
		log.Println(code, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(b)
	}
}

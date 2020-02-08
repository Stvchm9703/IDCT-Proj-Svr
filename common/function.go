package common

import (
	"context"
	"encoding/json"
	"hash/fnv"
	"io"
	"log"
	"os"
	"strconv"
)

// HashText: common hash text function
func HashText(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.FormatUint(uint64(h.Sum32()), 16)
}

func SetLog(path string) io.Writer {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Println(" Orders API Called")
	return wrt
}

func PrintReqLog(ctx context.Context, req interface{}) {
	jsoon, _ := json.Marshal(ctx)
	log.Println(string(jsoon))

	jsoon, _ = json.Marshal(req)
	log.Println(string(jsoon))
}

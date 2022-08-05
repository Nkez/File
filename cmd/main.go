package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/Nkez/check/internal/handler"
	"github.com/Nkez/check/internal/services"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := ConfigInit(); err != nil {
		log.Fatalf("error instaling configs: %s", err.Error())
	}
	if err := File(viper.GetString("path")); err != nil {
		log.Fatalf("error with file: %s", err.Error())
	}
	services := services.NewService()
	handlers := handler.NewHandler(services)
	if err := Run(viper.GetString("port"), handlers.InitRouter()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func ConfigInit() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func File(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
	}
	buf := bytes.NewBuffer(nil)
	decodeFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(buf, decodeFile)
	decodeFile.Close()
	for _, _ = range buf.Bytes() {
		decoder := gob.NewDecoder(buf)
		decoder.Decode(&services.Statistics)
	}
	return nil
}

func Run(port string, handler http.Handler) error {
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.ListenAndServe()
}

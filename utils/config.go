package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func GetConfig() *Config {
	file, err := os.Open("config.json")
	if err != nil {
		log.Println("파일 열기 오류:", err)
		return nil
	}
	defer file.Close()

	// 파일 내용 읽기
	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("파일 읽기 오류:", err)
		return nil
	}

	var config *Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Println("JSON 파싱 오류:", err)
		return nil
	}

	return config
}

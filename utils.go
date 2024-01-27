package utils

import "os"

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func LoadFile(path string) []byte {
	file, err := os.Open(path)
	HandleError(err)
	defer file.Close()
	stat, err := file.Stat()
	HandleError(err)
	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	HandleError(err)
	return data
}

func SaveFile(path string, data []byte) {
	file, err := os.Create(path)
	HandleError(err)
	defer file.Close()
	_, err = file.Write(data)
	HandleError(err)
}

func AddLog(path, log string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	HandleError(err)
	defer file.Close()
	_, err = file.WriteString(log + "\n")
	HandleError(err)
}

package main

import (
	"fmt"
	"grpc-app/pb"
	"io"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "テスト太郎",
		Email:       "test@gmail.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": {}},
		Profile: &pb.Employee_Text{
			Text: "My name is Taro.",
		},
		Birthday: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   31,
		},
	}

	bytes, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Cannot marshal", err)
		return
	}

	err = writeToFile(bytes, ".binary/go.bin")
	if err != nil {
		log.Fatalln("Error writeToFile", err)
		return
	}

	readBytes, err := readFromFile(".binary/go.bin")
	fmt.Println("out", readBytes)
	if err != nil {
		log.Fatalln("Error readFromFile", err)
		return
	}
	readEmployee := &pb.Employee{}
	if err := proto.Unmarshal(readBytes, readEmployee); err != nil {
		log.Fatal("Cannot unmarshal", err)
		return
	}

	log.Println("Success! desirialized >>>", readEmployee)
}

func writeToFile(bytes []byte, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatalln("Cannot create file", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		log.Fatalln("Cannot write", err)
		return err
	}
	return nil
}

func readFromFile(filepath string) ([]byte, error) {
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatalln("Cannot create file", err)
		return nil, err
	}
	defer file.Close()

	// Deserialize
	readBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Cannot read file", err)
		return nil, err
	}
	return readBytes, nil
}

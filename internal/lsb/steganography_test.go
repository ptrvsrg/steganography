package lsb_test

import (
	"bufio"
	"bytes"
	"image"
	"log"
	"os"
	"steganography/internal/lsb"
	"testing"
)

var (
	testDataPath = getEnvWithDefult("TEST_DATA_DIR", ".")

	file        = testDataPath + "/statham.png"
	encodedFile = testDataPath + "/encoded_statham.png"

	message = []byte("Those who get up early want to sleep all day.")
)

func TestEncodeFromFile(t *testing.T) {

	inFile, err := os.Open(file)
	if err != nil {
		log.Printf("Error opening file %s: %v", file, err)
		t.FailNow()
	}
	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Printf("Error decoding. %v", err)
		t.FailNow()
	}
	w := new(bytes.Buffer)
	err = lsb.Encode(w, img, message) // Encode the message into the image file
	if err != nil {
		log.Printf("Error Encoding file %v", err)
		t.FailNow()
	}
	outFile, err := os.Create(encodedFile)
	if err != nil {
		log.Printf("Error creating file %s: %v", encodedFile, err)
		t.FailNow()
	}
	w.WriteTo(outFile)
	defer outFile.Close()
}

func TestDecodeFromFile(t *testing.T) {
	inFile, err := os.Open(encodedFile)
	if err != nil {
		log.Printf("Error opening file %s: %v", encodedFile, err)
		t.FailNow()
	}
	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Print("Error decoding file")
		t.FailNow()
	}

	sizeOfMessage := lsb.GetMessageSizeFromImage(img)

	msg := lsb.Decode(sizeOfMessage, img) // Read the message from the picture file

	if !bytes.Equal(msg, message) {
		log.Print("messages dont match:")
		log.Println(string(msg))
		t.FailNow()
	}
}

func getEnvWithDefult(key, value string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return value
}

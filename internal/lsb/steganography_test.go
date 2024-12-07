package lsb_test

import (
	"bufio"
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"os"
	"steganography/internal/lsb"
	"testing"
)

var (
	testDataPath = getEnvWithDefault("TEST_DATA_DIR", ".")

	pngFile        = testDataPath + "/statham.png"
	encodedPngFile = testDataPath + "/encoded_statham.png"

	jpgFile        = testDataPath + "/statham.jpg"
	encodedJpgFile = testDataPath + "/encoded_statham.jpg"

	jpegFile        = testDataPath + "/statham.jpeg"
	encodedJpegFile = testDataPath + "/encoded_statham.jpeg"

	message = []byte("Those who get up early want to sleep all day.")
)

func TestEncodeFromPngFile(t *testing.T) {
	inFile, err := os.Open(pngFile)
	if err != nil {
		log.Printf("Error opening pngFile %s: %v", pngFile, err)
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
		log.Printf("Error encoding file %v", err)
		t.FailNow()
	}
	outFile, err := os.Create(encodedPngFile)
	if err != nil {
		log.Printf("Error creating file %s: %v", encodedPngFile, err)
		t.FailNow()
	}
	w.WriteTo(outFile)
	defer outFile.Close()
}

func TestDecodeFromPngFile(t *testing.T) {
	inFile, err := os.Open(encodedPngFile)
	if err != nil {
		log.Printf("Error opening file %s: %v", encodedPngFile, err)
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

	msg := lsb.Decode(sizeOfMessage, img) // Read the message from the picture pngFile

	if !bytes.Equal(msg, message) {
		log.Print("messages dont match:")
		log.Println(string(msg))
		t.FailNow()
	}
}

func TestEncodeFromJpgFile(t *testing.T) {
	inFile, err := os.Open(jpgFile)
	if err != nil {
		log.Printf("Error opening file %s: %v", jpgFile, err)
		t.FailNow()
	}
	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	img, err := jpeg.Decode(reader)
	if err != nil {
		log.Printf("Error decoding. %v", err)
		t.FailNow()
	}
	w := new(bytes.Buffer)
	err = lsb.Encode(w, img, message) // Encode the message into the image file
	if err != nil {
		log.Printf("Error encoding file %v", err)
		t.FailNow()
	}
	outFile, err := os.Create(encodedJpgFile)
	if err != nil {
		log.Printf("Error creating file %s: %v", encodedJpgFile, err)
		t.FailNow()
	}
	w.WriteTo(outFile)
	defer outFile.Close()
}

func TestDecodeFromJpgFile(t *testing.T) {
	inFile, err := os.Open(encodedJpgFile)
	if err != nil {
		log.Printf("Error opening file %s: %v", encodedJpgFile, err)
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

func TestEncodeFromJpegFile(t *testing.T) {
	inFile, err := os.Open(jpegFile)
	if err != nil {
		log.Printf("Error opening file %s: %v", jpegFile, err)
		t.FailNow()
	}
	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	img, err := jpeg.Decode(reader)
	if err != nil {
		log.Printf("Error decoding. %v", err)
		t.FailNow()
	}
	w := new(bytes.Buffer)
	err = lsb.Encode(w, img, message) // Encode the message into the image file
	if err != nil {
		log.Printf("Error encoding file %v", err)
		t.FailNow()
	}
	outFile, err := os.Create(encodedJpegFile)
	if err != nil {
		log.Printf("Error creating file %s: %v", encodedJpegFile, err)
		t.FailNow()
	}
	w.WriteTo(outFile)
	defer outFile.Close()
}

func TestDecodeFromJpegFile(t *testing.T) {
	inFile, err := os.Open(encodedJpegFile)
	if err != nil {
		log.Printf("Error opening file %s: %v", encodedJpegFile, err)
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

func getEnvWithDefault(key, value string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return value
}

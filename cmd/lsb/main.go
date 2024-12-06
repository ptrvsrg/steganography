package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v3"
	"image"
	"os"
	"steganography/internal/lsb"
)

func main() {
	cmd := &cli.Command{
		Name:                   "lsb",
		Version:                "0.0.1",
		Authors:                []any{"ptrvsrg"},
		Copyright:              "© 2024 ptrvsrg",
		Usage:                  "Tool for LSB steganography on images",
		UseShortOptionHandling: true,
		EnableShellCompletion:  true,
		Commands: []*cli.Command{
			{
				Name:                  "encode",
				Aliases:               []string{"e", "E"},
				Usage:                 "encode a message to a given image file",
				Action:                executeEncode,
				EnableShellCompletion: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "message",
						Aliases:     []string{"m", "M"},
						Usage:       "message",
						HideDefault: true,
						Required:    false,
						Local:       true,
					},
					&cli.StringFlag{
						Name:        "message-file",
						Aliases:     []string{"im", "IM"},
						Usage:       "input message path",
						HideDefault: true,
						Required:    false,
						Local:       true,
					},
					&cli.StringFlag{
						Name:        "input-file",
						Aliases:     []string{"i", "I"},
						Usage:       "input image path",
						HideDefault: true,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "output-file",
						Aliases:     []string{"o", "O"},
						Usage:       "output image path",
						HideDefault: true,
						Required:    true,
						Local:       true,
					},
				},
			},
			{
				Name:                  "decode",
				Aliases:               []string{"d", "D"},
				Usage:                 "decode a message from a given image file",
				Action:                executeDecode,
				EnableShellCompletion: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "input-file",
						Aliases:     []string{"i", "I"},
						Usage:       "input image path",
						HideDefault: true,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "message-file",
						Aliases:     []string{"om", "OM"},
						Usage:       "output message path",
						HideDefault: true,
						Required:    false,
						Local:       true,
					},
				},
			},
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Printf("error occurred: %v", err)
	}
}

func executeEncode(_ context.Context, command *cli.Command) error {
	// Get parameters
	var err error
	message := []byte(command.String("message"))
	messagePath := command.String("message-file")
	inputPath := command.String("input-file")
	outputPath := command.String("output-file")

	// Check message
	if len(message) == 0 && messagePath == "" {
		return errors.New("message or message file is missing")
	}

	// Read the message from the message file
	if messagePath != "" {
		message, err = os.ReadFile(messagePath)
		if err != nil {
			return errors.Wrap(err, "failed to open message file")
		}
	}

	// Open input file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return errors.Wrap(err, "failed to open input image")
	}
	defer inputFile.Close()

	// Reads binary data from input file
	reader := bufio.NewReader(inputFile)
	img, _, err := image.Decode(reader)
	if err != nil {
		return errors.Wrap(err, "failed to decode input image")
	}

	// Encodes the message into a new buffer
	encodedImg := new(bytes.Buffer)
	err = lsb.Encode(encodedImg, img, message)
	if err != nil {
		return errors.Wrap(err, "failed to encoding message into image")
	}

	// Create ouput file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return errors.Wrap(err, "failed to create output image")
	}

	// Write encoded binary data
	_, err = bufio.NewWriter(outFile).Write(encodedImg.Bytes())
	if err != nil {
		return errors.Wrap(err, "failed to write data to output image")
	}

	return nil
}

func executeDecode(_ context.Context, command *cli.Command) error {
	// Get parameters
	var err error
	inputPath := command.String("input-file")
	messagePath := command.String("message-file")

	// Open message output file
	writer := os.Stdout
	defer func(writer *os.File) {
		if writer != os.Stdout {
			writer.Close()
		}
	}(writer)
	if messagePath != "" {
		writer, err = os.Open(messagePath)
		if err != nil {
			return errors.Wrap(err, "failed to open message path")
		}
	}

	// Opens input file provided in the flags
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return errors.Wrap(err, "failed to open input image")
	}
	defer inputFile.Close()

	// Reads binary data from input file
	reader := bufio.NewReader(inputFile)
	img, _, err := image.Decode(reader)
	if err != nil {
		return errors.Wrap(err, "failed to decode input image")
	}

	// Check the message size
	sizeOfMessage := lsb.GetMessageSizeFromImage(img)

	// Read the message from the image
	msg := lsb.Decode(sizeOfMessage, img)

	// Print the message
	_, err = writer.Write(msg)
	if err != nil {
		return errors.Wrap(err, "failed to write data to output file")
	}

	return nil
}

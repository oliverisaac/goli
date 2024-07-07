package goli

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
)

// gzipAndBase64Encode compresses the input string using gzip and then encodes it using Base64 URL encoding.
func GzipAndBase64Encode(input string) (string, error) {
	// Create a buffer to hold the gzipped data.
	var gzipBuffer bytes.Buffer
	gzipWriter := gzip.NewWriter(&gzipBuffer)

	// Write the input string to the gzip writer.
	_, err := gzipWriter.Write([]byte(input))
	if err != nil {
		return "", err
	}

	// Close the gzip writer to flush any remaining data.
	if err := gzipWriter.Close(); err != nil {
		return "", err
	}

	// Base64 URL encode the gzipped data.
	encoded := base64.URLEncoding.EncodeToString(gzipBuffer.Bytes())

	return encoded, nil
}

// base64DecodeAndGunzip decodes the input string from Base64 URL encoding and then decompresses it using gzip.
func Base64DecodeAndGunzip(encoded string) (string, error) {
	// Decode the Base64 URL encoded string.
	gzippedData, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	// Create a gzip reader to decompress the gzipped data.
	gzipReader, err := gzip.NewReader(bytes.NewReader(gzippedData))
	if err != nil {
		return "", err
	}
	defer gzipReader.Close()

	// Read the decompressed data into a buffer.
	var decompressedBuffer bytes.Buffer
	if _, err := io.Copy(&decompressedBuffer, gzipReader); err != nil {
		return "", err
	}

	return decompressedBuffer.String(), nil
}

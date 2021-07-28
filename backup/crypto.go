package backup

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"io/ioutil"
	"os"
)

var alikeKey []byte
var alikeIV []byte

// CalculateCredentials will use the password given by user to generate
// a 256-bit key and a 128-bit IV
func CalculateCredentials() {

	password := os.Getenv("ALIKE_PASSWORD")
	alikeKey = getSHA256([]byte(password))
	alikeIV = getHalfOfKeyHash(getSHA256(alikeKey))
}

// getSHA256 will take the input stream and generate its hash
func getSHA256(stream []byte) []byte {

	hasher := sha256.New()
	hasher.Write(stream)
	return hasher.Sum(nil)
}

// getHalfOfKeyHash will get you half of your input array.
// This is used for getting the first 128-bits of the key's hash.
func getHalfOfKeyHash(key []byte) []byte {

	result := make([]byte, len(key)/2)

	for i := 0; i < len(result); i++ {
		result[i] = key[i]
	}

	return result
}

// getNumberOfBlocks will return the number of blocks to be made from a stream
func getNumberOfBlocks(stream []byte) uint64 {

	basic := uint64(len(stream) / blockSizeInBytes)
	if len(stream)%blockSizeInBytes != 0 {
		basic++
	}
	return basic
}

// getBlock will return the correct block from a stream
func getBlock(stream []byte, number uint64) []byte {

	result := make([]byte, blockSizeInBytes)

	start := blockSizeInBytes * number

	for i := start; i < start+blockSizeInBytes && i < uint64(len(stream)); i++ {
		result[i%blockSizeInBytes] = stream[i]
	}

	return result
}

// setBlock will update the bytes of the stream with the passed block
// which the given block number
func setBlock(stream *[]byte, number uint64, block []byte) {

	start := blockSizeInBytes * number
	for i := start; i < start+blockSizeInBytes && i < uint64(len(*stream)); i++ {
		(*stream)[i] = block[i%blockSizeInBytes]
	}
}

// useAES256CTR will apply AES256 to a stream of bytes in CTR mode
func useAES256CTR(stream []byte) []byte {

	// Prepare result in memory
	result := make([]byte, len(stream))

	// AES256
	blockCipher, err := aes.NewCipher(alikeKey)
	handleError(err)

	// Prepare for CTR mode
	ctr := cipher.NewCTR(blockCipher, alikeIV)

	// Run through all blocks
	count := getNumberOfBlocks(stream)
	for i := uint64(0); i < count; i++ {

		toEncrypt := getBlock(stream, i)
		encrypted := make([]byte, len(toEncrypt))
		ctr.XORKeyStream(encrypted, toEncrypt)
		setBlock(&result, i, encrypted)
	}

	return result
}

// encryptFile will encrypt the given file using AES256-CTR
func encryptFile(path string) []byte {

	// Manage file
	file, err := os.Open(path)
	handleError(err)
	defer file.Close()

	// Get all bytes
	stream, err := ioutil.ReadAll(file)
	handleError(err)

	// Encrypt in CTR mode
	encrypted := useAES256CTR(stream)

	return encrypted
}

// decryptFile will decrypt the given file using AES256-CTR
func decryptFile(path string) []byte {
	return encryptFile(path)
}

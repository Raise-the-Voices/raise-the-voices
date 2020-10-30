package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

// Message function wrap message to users
func Message(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}

// Respond function to responde request
func Respond(w http.ResponseWriter, data map[string]interface{}, status uint) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	w.Header().Add("X-XSS-Protection", "1; mode=block")
	w.Header().Add("X-Frame-Options", "Deny")
	w.Header().Add("Content-Security-Policy", "script-src 'self'")
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.Header().Add("Referrer-Policy", "no-referrer")
	w.Header().Add("Feature-Policy", "vibrate 'none'; geolocation 'none'")

	switch status {
	case 200: // break 200 Accept Request
		w.WriteHeader(http.StatusOK)
		break
	case 201: // break 201 created POST
		w.WriteHeader(http.StatusCreated)
		break
	case 204: // break 204 No Content (Just Delete Http)
		w.WriteHeader(http.StatusNoContent)
		break
	case 301: // break 301 Moved Permanently
		w.WriteHeader(http.StatusMovedPermanently)
		break
	case 400: // break 400 Bad Request
		w.WriteHeader(http.StatusBadRequest)
		break
	case 401: // break 401 Unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		break
	case 403: // break 403 Forbidden
		w.WriteHeader(http.StatusForbidden)
		break
	case 404: // break 404 Not Found
		w.WriteHeader(http.StatusNotFound)
		break
	case 500: // break 500 Internal Server Error
		w.WriteHeader(http.StatusInternalServerError)
		break
	}

	data["status"] = status
	json.NewEncoder(w).Encode(data)
}

// ConvertStringtoUint func to convert string to uint
func ConvertStringtoUint(stringPure string) uint {
	uint64bit, _ := strconv.ParseUint(stringPure, 10, 32)
	return uint(uint64bit)
}

// EvaluationNumberPattern func to evaluate string if
func EvaluationNumberPattern(stringPure string, withZero bool) bool {
	// Number Pattern
	onlyNumberPattern := `(?m)^[1-9]\d*$`
	withZeroPattern := `(?m)^[0-9]+$`
	if !withZero {
		resultPattern, _ := regexp.MatchString(onlyNumberPattern, stringPure)
		return resultPattern
	}
	resultPattern, _ := regexp.MatchString(withZeroPattern, stringPure)
	return resultPattern
}

// GenerateRandomBytes returns securely generated random bytes.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString() (string, error) {
	b, err := generateRandomBytes(32)
	return base64.URLEncoding.EncodeToString(b), err
}

// DownloadFile a file example a Picture from a URL and return name of file. (Save the file in Disk)
func DownloadFile(filePath string) string {

	// don't worry about errors
	response, e := http.Get(filePath)
	if e != nil {
		return ""
	}
	defer response.Body.Close()

	// Check server response
	if response.StatusCode != http.StatusOK {
		fmt.Println("bad status: ", response.Status)
	}

	//open a file for writing
	file, err := os.Create(path.Base(response.Request.URL.String()))
	if err != nil {
		return ""
	}
	defer file.Close()
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return ""
	}

	return path.Base(response.Request.URL.String())

}

// GenerateUUID generate uuid to media name
func GenerateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// IsValidImageFileFormat image format allowed
func IsValidImageFileFormat(filetype string) bool {
	if filetype == "image/jpeg" || filetype == "image/jpg" || filetype == "image/png" {
		return true
	}
	return false
}

// IsValidFileFormat object format allowed
func IsValidFileFormat(filetype string) bool {
	if filetype == "image/jpeg" || filetype == "image/jpg" || filetype == "image/png" || filetype == "image/gif" ||

		filetype == "audio/wav" || filetype == "audio/mp3" || filetype == "audio/mp4" || filetype == "audio/mpeg" ||

		filetype == "video/mpeg" || filetype == "video/mp4" || filetype == "video/3gpp"  ||
		
		filetype == "application/zip" || filetype == "application/pdf" || 
		
		filetype == "application/msword"  || filetype == "application/rtf"  || filetype == "application/vnd.openxmlformats-officedocument.wordprocessingml.document" || filetype == "application/vnd.oasis.opendocument.text" {
		return true
	}
	log.Println(filetype+" is not a valid file format for upload")
	return false
}

// IsGreaterObjectSizeThanLimit file limit size
func IsGreaterObjectSizeThanLimit(fileSize, maxUploadSize int64) bool {
	if fileSize > maxUploadSize {
		return true
	}
	return false
}

// GenerateFileName generate the name of the original photo and thumbnail size photo
func GenerateFileName(fileExtension string) (string, string) {
	generatedUUID := GenerateUUID()
	return fmt.Sprintf("%s%s", generatedUUID, fileExtension), fmt.Sprintf("%s_thumb%s", generatedUUID, fileExtension)
}

// ResizeImage to thumbnail size
func ResizeImage(file multipart.File) (*image.NRGBA, string, error) {

	//  sets the offset for the beginning of the file before calling image.Decode
	file.Seek(0, 0)

	imagepo, thumbFormatType, errDecode := image.Decode(file)

	// Resize the image to width = 133px and height = 100px .
	src := imaging.Resize(imagepo, 133, 100, imaging.Lanczos)

	return src, thumbFormatType, errDecode
}

// EncodeImage to Buffer
func EncodeImage(thumbFormatType string, src *image.NRGBA) (*bytes.Buffer, error) {
	buff := new(bytes.Buffer)
	var err error
	if thumbFormatType == "jpeg" || thumbFormatType == "jpg" {
		err = jpeg.Encode(buff, src, nil)
	} else if thumbFormatType == "png" {
		err = png.Encode(buff, src)
	}
	return buff, err
}

// FindThumbnailURL .
func FindThumbnailURL(mediaURL string) (string, string) {

	baseURL := path.Base(mediaURL)

	extensionIMAGE := path.Ext(baseURL)

	arrayURL := strings.Split(baseURL, extensionIMAGE)

	baseURL = arrayURL[0] + "_thumb" + extensionIMAGE

	return baseURL, extensionIMAGE
}

// IsThisItAnImage .
func IsThisItAnImage(fileExtension string) bool {
	fileExtension = strings.ToLower(fileExtension)
	if fileExtension == ".jpg" || fileExtension == ".jpeg" || fileExtension == ".png" {
		return true
	}
	return false
}

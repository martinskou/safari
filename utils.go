package safari

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"text/template"
	"time"
)

/* Logs function execution time, use defer */
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

/* Pretty-print anything which can be marshalled */
func Pprint(v interface{}) {
	s, e := json.MarshalIndent(v, "", "  ")
	if e != nil {
		fmt.Printf("Pprint error : %s\n", e)
	} else {
		fmt.Printf("%s\n", s)
	}
}

/* Return true if path exists */
func ExistsPath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

/* Generate a simple password */
func SimplePassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghjkmnpqrstuvwxyz23456789ABCDEFGHJKMNPQRSTUVWXYZ")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func Uuid() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func HashSha256(pw string, salt string) string {
	h := sha256.New()
	io.WriteString(h, pw+salt)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func ListDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

func Render(w io.Writer, template_filename string, data interface{}) error {

	template_filename_array := strings.Split(template_filename, "/")
	template_base_path := template_filename_array[0]
	template_file_path := template_filename_array[1]

	abs_template_filename := fmt.Sprintf("templates/%s", template_filename)
	template_layout_path := fmt.Sprintf("templates/%s/layout/", template_base_path)

	files, err := ListDir(template_layout_path)
	if err != nil {
		log.Fatal(err)
	}

	afiles := []string{}
	for _, f := range files {
		afiles = append(afiles, fmt.Sprintf("%s%s", template_layout_path, f))
	}
	afiles = append(afiles, abs_template_filename)

	//	tn := tfna[len(tfna)-1]
	t, e := template.New(template_file_path).Funcs(template.FuncMap{
		//	"Icon": Icon,
		//		}).ParseFiles(tfn)
	}).ParseFiles(afiles...)

	if e == nil {
		e = t.Execute(w, data)
		if e != nil {
			log.Printf("Error executing template [err=%s]", e.Error())
			return e
		}
	} else {
		log.Printf("Error parsing template [err=%s]", e.Error())
		return e
	}
	return nil
}

func RenderString(tfn string, data interface{}) string {
	var buf bytes.Buffer
	err := Render(&buf, tfn, data)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}

package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"text/template"

	"k8s.io/klog/v2"
)

type FileLocker struct {
	dir  string
	file string
}

func NewFileLocker() (*FileLocker, error) {
	dir, err := os.MkdirTemp("", "text2image")
	if err != nil {
		return nil, fmt.Errorf("error creating tmp dir: %w", err)
	}
	return &FileLocker{dir, path.Join(dir, "lock")}, nil
}

var AlreadyLockedError error = errors.New("already locked")
var NotLockedError error = errors.New("not locked")

func (f *FileLocker) Lock() error {
	_, err := os.Stat(f.file)
	if err == nil {
		return AlreadyLockedError
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("something went wrong stating file: %w", err)
	}

	if _, err := os.Create(f.file); err != nil {
		return fmt.Errorf("error creating lock file: %w", err)
	}

	return nil
}

func (f *FileLocker) Unlock() error {
	_, err := os.Stat(f.file)
	if err != nil {
		if os.IsNotExist(err) {
			return NotLockedError
		} else {
			return fmt.Errorf("something went wrong stating file: %w", err)
		}
	}

	if err := os.Remove(f.file); err != nil {
		return fmt.Errorf("error removing lock file: %w", err)
	}

	return nil
}

func (f *FileLocker) IsLocked() bool {
	_, err := os.Stat(f.file)
	if err == nil {
		return true
	}
	return false
}

func (f *FileLocker) CleanUp() {
	os.RemoveAll(f.dir)
}

const indexPage string = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title></title>
</head>
<body>
  <p>text2image. Click submit and wait about 3 minutes.</p>
  <form method="POST" action="/image">
    <input type="text" name="prompt">
    <button type="submit">Submit</button>
  </form>
</body>
</html>`

const imagePage string = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title></title>
</head>
<body>
  {{range $i, $e := .}}
  <img src="{{$e}}">
  {{end}}
</body>
</html>
`

func init() {
	klog.InitFlags(nil)
}

func main() {
	imageTemplate := template.Must(template.New("image").Parse(imagePage))

	flock, err := NewFileLocker()
	if err != nil {
		klog.Fatal(err)
	}
	defer flock.CleanUp()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if flock.IsLocked() {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Another request is in progress. Please try again in a few minutes")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "text/html")
		fmt.Fprint(w, indexPage)
	})

	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		prompt := r.Form.Get("prompt")
		if prompt == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := flock.Lock(); errors.Is(err, AlreadyLockedError) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Another request is in progress. Please try again in a few minutes")
			return
		}
		defer flock.Unlock()

		images, err := fetchImages(prompt)
		if err != nil {
			klog.Info(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Shit broke...")
			return
		}

		w.WriteHeader(200)
		imageTemplate.Execute(w, images)
	})

	klog.Info("starting server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		klog.Fatal(err)
	}
}

func fetchImages(prompt string) ([]string, error) {
	outBuf := bytes.NewBuffer(make([]byte, 0))
	cmd := exec.Command("python", "optimizedSD/optimized_txt2img.py", "--prompt", prompt, "--H", "512", "--W", "512", "--seed", "27", "--n_iter", "2", "--n_samples", "10", "--ddim_steps", "50")
	tee := io.MultiWriter(os.Stdout, outBuf)
	cmd.Stdout = tee
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error running command: %w", err)
	}
	lines := make([]string, 0)
	s := bufio.NewScanner(outBuf)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	dir := lines[len(lines)-1]

	de, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	images := make([]string, len(de))
	for _, image := range de {
		b, err := os.ReadFile(path.Join(dir, image.Name()))
		if err != nil {
			return nil, fmt.Errorf("error reading image file: %w", err)
		}
		enc := base64.StdEncoding.EncodeToString(b)

		images = append(images, fmt.Sprintf("data:image/png;base64, %s", enc))
	}
	return images, nil
}

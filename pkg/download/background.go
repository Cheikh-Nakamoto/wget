package download

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"wget/pkg/utils"

	"github.com/sevlyar/go-daemon"
)

func DownloadBackground(url string, dir string) error {
	context := &daemon.Context{
		PidFilePerm: 0644,
		LogFileName: "wget-log",
		LogFilePerm: 0640,
		WorkDir:     ".",
		Umask:       027,
	}

	// Attempt to start the process as a daemon
	P, err := context.Reborn()
	if err != nil {
		return fmt.Errorf("error starting the daemon: %v", err)
	}
	if P != nil {
		return nil // Parent process exits immediately
	}
	defer context.Release()

	// Open log file to record activity
	logFile, err := os.OpenFile("wget-log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer logFile.Close()

	if err := WriteLog(logFile, fmt.Sprintf("start at %s", utils.FormatDate(time.Now()))); err != nil {
		return err
	}

	// Send the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		WriteLog(logFile, fmt.Sprintf("Failed to download URL %s: %v", url, err))
		return err
	}
	defer resp.Body.Close()
	WriteLog(logFile, fmt.Sprintf("Sending request, awaiting response... status %s", resp.Status))

	// Determine filename and create the output file
	filename := filepath.Base(url)
	outputPath := strings.Join([]string{dir, filename}, "/")

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Write response body to the file
	n, err := outFile.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	WriteLog(logFile, fmt.Sprintf("content size: %v [~%.2fMB]", n, math.Round(float64(n)/1024.0/1024.0)))
	WriteLog(logFile, fmt.Sprintf("saving file to: %s", filename))
	WriteLog(logFile, fmt.Sprintf("Downloaded [%s]", url))

	WriteLog(logFile, fmt.Sprintf("finished at %s", utils.FormatDate(time.Now())))
	return nil
}

func WriteLog(logFile *os.File, message string) error {
	logEntry := fmt.Sprintf("%s\n", message)
	if _, err := logFile.WriteString(logEntry); err != nil {
		return fmt.Errorf("failed to write to log file: %w", err)
	}
	return nil
}

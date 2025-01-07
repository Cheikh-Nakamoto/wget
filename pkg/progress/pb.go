package progress

import (
	"fmt"
	"io"
	"strings"
	"time"
	"wget/pkg/utils"
)

func Progression(reader io.Reader, totalZise int64) io.Reader {
	progress := make(chan int64)
	done := make(chan bool)

	go func() {
		var downloaded int64
		startTime := time.Now()

		for {
			select {
			case nb := <-progress:
				downloaded += nb
				percent := float64(downloaded) / float64(totalZise) * 100

				// Assurer que percent est dans la plage [0, 100]
				if percent < 0 {
					percent = 0
				} else if percent > 100 {
					percent = 100
				}

				// Calculer la longueur de la barre (50 caractères max)
				barLength := int(percent / 2)
				if barLength < 0 {
					barLength = 0
				} else if barLength > 50 {
					barLength = 50
				}

				// Construire la barre de progression
				bar := strings.Repeat("=", barLength) + ">" + strings.Repeat(" ", 50-barLength)

				// Calculer le temps écoulé et la vitesse de téléchargement
				elapsedTime := time.Since(startTime)
				speed := float64(downloaded) / elapsedTime.Seconds()

				// Calculer le temps restant
				remainingTime := time.Duration(float64(totalZise-downloaded)/speed) * time.Second

				// Afficher la progression
				fmt.Printf("\r%s / %s [%s] %.2f%% %s %s", utils.ConvertSize(downloaded), utils.ConvertSize(totalZise), bar, percent, utils.ConvertSpeed(speed/1024), remainingTime)

			case <-done:
				fmt.Println("finished at", utils.FormatDate(time.Now()))
				fmt.Println("\nDownload completed!")
				return // Terminer la goroutine
			}
		}
	}()
	return io.TeeReader(reader, progressWriter{progress, done})
}

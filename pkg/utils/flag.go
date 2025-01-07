package utils

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
	"wget/pkg"
)

func ParseFlag() (map[string]string, []string, error) {
	flag.Parse()

	if *pkg.FileInput == "" && len(flag.Args()) == 0 {
		return nil, nil, fmt.Errorf("at least one URL is required")
	}

	flags := map[string]string{
		"R":             *pkg.Reject,
		"B":             boolToString(*pkg.BackGround),
		"X":             *pkg.Exclude,
		"O":             *pkg.Output,
		"i":             *pkg.FileInput,
		"rate-limit":    *pkg.RateLimit,
		"P":             *pkg.NewPath,
		"mirror":        boolToString(*pkg.Mirroring),
		"convert-links": boolToString(*pkg.ConvertLinks),
	}

	urls := flag.Args()

	return flags, urls, nil
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func ValidateFlag(flags map[string]string) error {
	if rateLimit, ok := flags["rate-limit"]; ok && rateLimit != "" {
		if !isValidRateLimit(rateLimit) {
			return fmt.Errorf("invalid rate limit: %s", rateLimit)
		}
	}

	if outputDir, ok := flags["P"]; ok && outputDir != "" {
		if !isValidDirectory(outputDir) {
			return fmt.Errorf("output directory '%s' must be an absolute path", outputDir)
		}
	}

	if rejectList, ok := flags["R"]; ok && rejectList != "" {
		for _, ext := range strings.Split(rejectList, ",") {
			if !isValidExtension(ext) {
				return fmt.Errorf("invalid extension: %s", ext)
			}
		}
	} else if rejectList, ok := flags["reject"]; ok && rejectList != "" {
		for _, ext := range strings.Split(rejectList, ",") {
			if !isValidExtension(ext) {
				return fmt.Errorf("invalid extension: %s", ext)
			}
		}
	}

	if excludeDirs, ok := flags["X"]; ok && excludeDirs != "" {
		for _, dir := range strings.Split(excludeDirs, ",") {
			if !isValidDirectory(dir) {
				return fmt.Errorf("exclude directory '%s' must be an absolute path", dir)
			}
		}
	} else if excludeDirs, ok := flags["exclude"]; ok && excludeDirs != "" {
		for _, dir := range strings.Split(excludeDirs, ",") {
			if !isValidDirectory(dir) {
				return fmt.Errorf("exclude directory '%s' must be an absolute path", dir)
			}
		}
	}

	if inputFile, ok := flags["i"]; ok && inputFile != "" {
		if !CheckFileExist(inputFile) {
			return fmt.Errorf("file %s does not exist", inputFile)
		}
	}

	return nil
}

func isValidDirectory(path string) bool {
	// if strings.HasPrefix(path, "/") {
	// 	path = path[1:]
	// }
	return regexp.MustCompile(`^[a-zA-Z0-9]+`).MatchString(path) || strings.HasPrefix(path, "./") || strings.HasPrefix(path, "../") || strings.HasPrefix(path, "~/") || strings.HasPrefix(path, "/")
}

func isValidExtension(ext string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, ext)
	return match
}

func isValidRateLimit(rateLimit string) bool {
	match, _ := regexp.MatchString(`^\d+[kM]?$`, rateLimit)
	return match
}

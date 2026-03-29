package fsx

import "os"

// LatestByModTime returns the path with the greatest modification time.
// Returns "" if paths is empty or every stat fails.
func LatestByModTime(paths []string) string {
	var latest string
	var latestNs int64
	for _, f := range paths {
		st, err := os.Stat(f)
		if err != nil {
			continue
		}
		n := st.ModTime().UnixNano()
		if n > latestNs {
			latestNs = n
			latest = f
		}
	}
	return latest
}

package chunks

import (
	"bytes"
	"strings"
)

func Chunks(s string, chunkSize int) []string {
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string
	chunk := make([]rune, chunkSize)
	len := 0
	for _, r := range s {
		chunk[len] = r
		len++
		if len == chunkSize {
			chunks = append(chunks, string(chunk))
			len = 0
		}
	}
	if len > 0 {
		chunks = append(chunks, string(chunk[:len]))
	}
	return chunks
}

func SplitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}

func ChunkString(s string, chunkSize int) []string {
	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks
}

func ChunkStringImproved(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	runes := []rune(s)
	n := (len(runes) + chunkSize - 1) / chunkSize
	chunks := make([]string, n)
	for i := 0; i < n-1; i++ {
		chunks[i] = string(runes[i*chunkSize : (i+1)*chunkSize])
	}
	chunks[n-1] = string(runes[(n-1)*chunkSize : len(runes)])
	return chunks
}

func Build(s string, chunkSize int) []string {
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string
	var b strings.Builder
	b.Grow(chunkSize)
	l := 0
	for _, r := range s {
		b.WriteRune(r)
		l++
		if l == chunkSize {
			chunks = append(chunks, b.String())
			l = 0
			b.Reset()
			b.Grow(chunkSize)
		}
	}
	if l > 0 {
		chunks = append(chunks, b.String())
	}
	return chunks
}

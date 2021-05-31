package chunks

import (
	"math/rand"
	"reflect"
	"testing"
)

var s1 = randomString(500)
var s2 = randomString(4000)
var s3 = randomString(20000)

func randomString(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func DoBenchmarks(b *testing.B, f func(s string, chunkSize int) []string) {
	f(s1, 10)
	f(s2, 10)
	f(s3, 10)
	f(s1, 100)
	f(s2, 100)
	f(s3, 100)
	f(s1, 1000)
	f(s2, 1000)
	f(s3, 1000)
}

func BenchmarkChunks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoBenchmarks(b, Chunks)
	}
	b.ReportAllocs()
}

func BenchmarkSplitSubN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoBenchmarks(b, SplitSubN)
	}
	b.ReportAllocs()
}

func BenchmarkChunkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoBenchmarks(b, ChunkString)
	}
	b.ReportAllocs()
}

func BenchmarkChunkStringImproved(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoBenchmarks(b, ChunkStringImproved)
	}
	b.ReportAllocs()
}

func BenchmarkBuild(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoBenchmarks(b, Build)
	}
	b.ReportAllocs()
}

func TestEquality(t *testing.T) {
	if !reflect.DeepEqual(SplitSubN(s3, 100), Chunks(s3, 100)) {
		t.Error()
	}
	if !reflect.DeepEqual(ChunkString(s3, 100), Chunks(s3, 100)) {
		t.Error()
	}
	if !reflect.DeepEqual(ChunkStringImproved(s3, 100), Chunks(s3, 100)) {
		t.Error()
	}
	if !reflect.DeepEqual(Build(s3, 100), Chunks(s3, 100)) {
		t.Error()
	}
}

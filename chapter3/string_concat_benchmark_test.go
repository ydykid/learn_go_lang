package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var sl []string = []string {
	"Rob Pike",
	"Robert Griesemer",
	"Ken Thompson",
}

func concatStringByOperator(sl []string) string {
	s := ""
	for _, v := range sl {
		s += v
	}
	return s
}

func concatStringBySprintf(sl []string) string {
	var s string
	for _, v := range sl {
		s = fmt.Sprintf("%s%s", s, v)
	}
	return s
}

func concatStringByJoin(sl []string) string {
	return strings.Join(sl, "")
}

func concatStringByStringsBuilder(sl []string) string {
	var b strings.Builder
	for _, v := range sl {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByStringsBuilderWithInitSize(sl []string) string {
	var b strings.Builder
	b.Grow(64)
	for _, v := range sl {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByBytesBuffer(sl []string) string {
	var b bytes.Buffer
	for _, v := range sl {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByBytesBufferWithInitSize(sl []string) string {
	buf := make([]byte, 0, 64)
	b := bytes.NewBuffer(buf)
	for _, v := range sl {
		b.WriteString(v)
	}
	return b.String()
}

func BenchmarkConcatStringByOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByOperator(sl)
	}
}

func BenchmarkConcatStringBySprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringBySprintf(sl)
	}
}

func BenchmarkConcatStringByJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByJoin(sl)
	}
}

func BenchmarkConcatStringByStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByStringsBuilder(sl)
	}
}

func BenchmarkConcatStringByStringsBuilderWithInitSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByStringsBuilderWithInitSize(sl)
	}
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByBytesBuffer(sl)
	}
}

func BenchmarkConcatStringByBytesBufferWithInitSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByBytesBufferWithInitSize(sl)
	}
}

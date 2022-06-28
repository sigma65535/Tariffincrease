package main

import (
	"testing"
	"time"
)

func TestFindTheWinner(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n: 3,
				m: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheWinner(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("FindTheWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxDifference(t *testing.T) {

	tests := []struct {
		name string
		args []int64
		want int64
	}{
		{
			name: "ok",
			args: []int64{5, 8, 10, 1, 3},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDifference(tt.args); got != tt.want {
				t.Errorf("MaxDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache(t *testing.T) {
	c := NewCache()
	c.Set("a", "hello", time.Now().Add(5*time.Second))
	if c.Get("a") != "hello" {
		t.Errorf("get key error")
	}
	time.Sleep(6 * time.Second)
	if c.Get("a") != "" {
		t.Errorf("expired key")
	}

	c.Set("a", "hello", time.Now().Add(5*time.Second))
	c.Del("a")
	if c.Get("a") != "" {
		t.Errorf("del key")
	}

}

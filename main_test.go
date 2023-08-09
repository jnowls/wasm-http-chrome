package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"testing"
)

type Account struct {
	ID string `json:"id"`
}

func TestPersist(t *testing.T) {
	fmt.Printf("runtime.Version(): %v\n", runtime.Version())
	acc := Account{"5e7fd575-b5f5-4d7e-8fa7-70b5fef7662f"}
	buf, _ := json.Marshal(acc)
	r := bytes.NewBuffer(buf)
	req, err := http.NewRequest("POST", "https://nk.axesfull.dev/v2/account/authenticate/device", r)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("could not authenticate account: %v", err)
	}
}

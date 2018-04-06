// Copyright © 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package util

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"os/exec"
	"sort"
	"strings"
	"unicode"
)

const (
	maskFile = 0664
)

func Trim(s string) string {
	return strings.TrimFunc(s, unicode.IsSpace)
}

func MakeFluentdSafeName(s string) string {
	filter := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-' && r != '_'
	}
	return strings.TrimFunc(s, filter)
}

func Hash(owner string, value string) string {
	h := sha1.New()

	h.Write([]byte(owner))
	h.Write([]byte(":"))
	h.Write([]byte(value))

	b := h.Sum(nil)
	return hex.EncodeToString(b[:])
}

func SortedKeys(m map[string]string) []string {
	keys := make([]string, len(m))
	i := 0

	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	return keys
}

func ExecAndGetOutput(cmd string, args ...string) (string, error) {
	c := exec.Command(cmd, args...)
	out, err := c.CombinedOutput()

	return string(out), err
}

func WriteStringToFile(filename string, data string) error {
	return ioutil.WriteFile(filename, []byte(data), maskFile)
}
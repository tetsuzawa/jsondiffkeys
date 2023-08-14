package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/samber/lo"
)

func main() {
	sep := flag.String("s", "\n", "separator")
	flag.Parse()
	if flag.NArg() != 2 {
		fmt.Fprintln(os.Stderr, "usage: jsondiffkeys <file1> <file2>")
		os.Exit(1)
	}
	a := flag.Arg(0)
	b := flag.Arg(1)
	aBuf, err := os.ReadFile(a)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("failed to read %s: %w", a, err))
		os.Exit(1)
	}
	bBuf, err := os.ReadFile(b)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("failed to read %s: %w", a, err))
		os.Exit(1)
	}
	if err := run(aBuf, bBuf, *sep); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(a, b []byte, sep string) error {
	diffKeys, err := jsonDiffKeys(a, b)
	if err != nil {
		return fmt.Errorf("failed to compare: %w", err)
	}
	for i, k := range diffKeys {
		if i < len(diffKeys)-1 {
			fmt.Printf("%v%v", k, sep)
		} else {
			fmt.Printf("%v", k)
		}
	}
	return nil
}

func jsonDiffKeys(a, b []byte) ([]string, error) {

	var objA any
	var objB any

	if err := json.Unmarshal(a, &objA); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &objB); err != nil {
		return nil, err
	}
	diffKeys := *deepDiffKeys(objA, objB, "", &[]string{})
	uniqDiffKeys := lo.Uniq(diffKeys)
	sort.Strings(uniqDiffKeys)
	return uniqDiffKeys, nil
}

func deepDiffKeys(a, b any, keyHierarchy string, diffKeys *[]string) *[]string {

	switch aVal := a.(type) {
	case []any:
		bVal, ok := b.([]any)
		if !ok {
			*diffKeys = append(*diffKeys, keyHierarchy)
			return diffKeys
		}

		// sortするとNlog(N)になるが、uuidなど都度変わる値があるとsortが安定でなくなるため全探索する
		visited := make([]bool, len(bVal))
		for _, vA := range aVal {
			matchFound := false
			for j, vB := range bVal {
				if !visited[j] {
					deepDiffKeys(vA, vB, keyHierarchy, diffKeys)
					visited[j] = true
					matchFound = true
					break
				}
			}
			if !matchFound {
				*diffKeys = append(*diffKeys, keyHierarchy)
				//fmt.Println("here 7", keyHierarchy, aVal, bVal)
				return diffKeys
			}
		}
		return diffKeys
	case map[string]interface{}:
		bVal, ok := b.(map[string]interface{})
		if !ok || len(aVal) != len(bVal) {
			//fmt.Println("here 3", keyHierarchy, aVal, bVal)
			*diffKeys = append(*diffKeys, keyHierarchy)
			return diffKeys
		}

		for key, valA := range aVal {
			valB, exists := bVal[key]
			if !exists {
				//fmt.Println("here 4", keyHierarchy, aVal, bVal)
				*diffKeys = append(*diffKeys, keyHierarchy)
				return diffKeys
			}

			presentKeyHierarchy := ""
			if keyHierarchy == "" {
				presentKeyHierarchy = key
			} else {
				presentKeyHierarchy = keyHierarchy + "." + key
			}
			//fmt.Printf("key: %s, presentKeyHiralchy: %s\n", key, presentKeyHierarchy)
			deepDiffKeys(valA, valB, presentKeyHierarchy, diffKeys)
		}
		return diffKeys
	default:
		if a != b {
			//fmt.Println("here 6", keyHierarchy, a, b)
			*diffKeys = append(*diffKeys, keyHierarchy)
		}
		return diffKeys
	}
}

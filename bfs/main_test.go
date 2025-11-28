package bfs_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

var numberOfCourses int

func createGraph() error {
	// load data from JSON file
	file, err := os.Open("../course.json")
	if err != nil {
		return err
	}
	defer file.Close()
	var courses map[string][]string
	if err := json.NewDecoder(file).Decode(&courses); err != nil {
		return err
	}
	numberOfCourses = len(courses)
	// build the graph from the map
	for k := range courses {
		g.AddVertex(k)
	}
	for k, v := range courses {
		for _, vv := range v {
			g.AddEdge(k, vv)
		}
	}
	return nil
}

func TestMain(m *testing.M) {
	if err := createGraph(); err != nil {
		fmt.Fprintf(os.Stderr, "graph creation failed: %v\n", err.Error())
		os.Exit(1)
	}
	exitCode := m.Run()
	os.Exit(exitCode)
}

package test

import (
	"TaskList_for_go/tasklist/entity"
	"testing"
)

func TestProjectNameToString(t *testing.T) {
	projectName := entity.NewProjectName("test")
	if projectName.ToString() != "test" {
		t.Errorf("projectNameOf() = %v; want test", projectName.ToString())
	}
}

package controllers_test

import (
	"testing"

	"sigma_geo/backend/controllers"
)

func TestCODProject(t *testing.T) {
	project_name := "RMO"
	pm := controllers.NewProjectManager()
	err := pm.CreateProject(project_name)
	if err != nil {
		t.Fatal(err)
	}
	err = pm.OpenProject(project_name)
	if err != nil {
		t.Fatal(err)
	}
	err = pm.DeleteProject(project_name)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCDSubprojects(t *testing.T) {
	project_name := "RMO"
	subproject_name := "rmo_2025"
	pm := controllers.NewProjectManager()
	pm.CreateProject(project_name)
	pm.OpenProject(project_name)

	err := pm.CreateSubproject(project_name, subproject_name)
	if err != nil {
		t.Fatal(err)
	}
	err = pm.DeleteSubproject(project_name, subproject_name)
	if err != nil {
		t.Fatal(err)
	}
	pm.DeleteProject(project_name)
}

func TestCDVariant(t *testing.T) {
	project_name := "RMO"
	subproject_name := "rmo_2025"
	variant_name := "rmo_mathbattle"
	pm := controllers.NewProjectManager()
	pm.CreateProject(project_name)
	pm.OpenProject(project_name)
	pm.CreateSubproject(project_name, subproject_name)

	err := pm.CreateVariant(project_name, subproject_name, variant_name)
	if err != nil {
		t.Fatal(err)
	}
	err = pm.DeleteVariant(project_name, subproject_name, variant_name)
	if err != nil {
		t.Fatal(err)
	}
	pm.DeleteProject(project_name)
}

package controllers

import (
	"fmt"

	"sigma_geo/backend/controllers/file_manager"
	"sigma_geo/backend/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Problem = models.Problem

type ProjectManager struct {
	DB             *gorm.DB
	CurrentProject *string
}

func NewProjectManager() *ProjectManager {
	return &ProjectManager{}
}

func (p *ProjectManager) OpenProject(name string) (err error) {
	db, err := gorm.Open(sqlite.Open(projects_folder+name+"/"+name+".db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("database (%s) connection failed: %w", name, err)
	}
	p.DB = db
	err = p.DB.AutoMigrate(&Problem{})
	if err != nil {
		return fmt.Errorf("database migration failed: %w", err)
	}
	p.CurrentProject = &name
	return
}

func (p *ProjectManager) CreateProject(name string) (err error) {
	err = file_manager.MakeDir(projects_folder + name)
	if err != nil {
		return
	}
	err = file_manager.MakeFile(projects_folder + name + "/" + name + ".db")
	return
}

func (p *ProjectManager) DeleteProject(name string) (err error) {
	if *p.CurrentProject == name {
		sqlDB, err := p.DB.DB()
		if err != nil {
			return fmt.Errorf("extraction (%s) db failed: %w", name, err)
		}
		err = sqlDB.Close()
		if err != nil {
			return fmt.Errorf("closing (%s) db failed: %w", name, err)
		}
	}
	err = file_manager.Remove(projects_folder + name)
	return
}

func (p *ProjectManager) CreateSubproject(proj_name, sub_name string) (err error) {
	err = file_manager.MakeDir(projects_folder + proj_name + "/" + sub_name)
	return
}

func (p *ProjectManager) DeleteSubproject(proj_name, sub_name string) (err error) {
	err = file_manager.Remove(projects_folder + proj_name + "/" + sub_name)
	return
}

func (p *ProjectManager) CreateVariant(proj_name, sub_name, var_name string) (err error) {
	variant_path := proj_name + "/" + sub_name + "/" + var_name
	err = file_manager.MakeDir(projects_folder + variant_path)
	if err != nil {
		return
	}
	err = file_manager.MakeFile(projects_folder + variant_path + "/" + var_name + ".tex")
	if err != nil {
		return
	}
	err = file_manager.MakeFile(projects_folder + variant_path + "/" + var_name + "_sol.tex")
	return
}

func (p *ProjectManager) CompileVariant(subproj_name, var_name string) (err error) {
	return
}

func (p *ProjectManager) DeleteVariant(proj_name, sub_name, var_name string) (err error) {
	variant_path := proj_name + "/" + sub_name + "/" + var_name
	err = file_manager.Remove("internal/projects/" + variant_path)
	return
}

func (p *ProjectManager) CreateProblem(row_prob Problem) (err error) {
	err = row_prob.Create(p.DB)
	return
}

func (p *ProjectManager) UpdateProblem(row_prob Problem) (err error) {
	err = row_prob.Update(p.DB)
	return
}

func (p *ProjectManager) DeleteProblem(row_prob Problem) (err error) {
	err = row_prob.Delete(p.DB)
	return
}

func (p *ProjectManager) GetProblems(row_prob Problem) (problem_list []Problem, err error) {
	err = row_prob.GetAll(p.DB, &problem_list)
	return
}

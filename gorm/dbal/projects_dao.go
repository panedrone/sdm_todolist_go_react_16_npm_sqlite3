package dbal

// Code generated by a tool. DO NOT EDIT.
// Additions may be hand-coded in a separate go-file.
// https://sqldalmaker.sourceforge.net/

import (
	"context"
	"sdm_demo_todolist/gorm/dbal/models"
)

type ProjectsDao struct {
	ds DataStore
}

// (C)RUD: projects
// Generated/AI values are passed to DTO/model.

func (dao *ProjectsDao) CreateProject(ctx context.Context, item *models.Project) error {
	return dao.ds.Create(ctx, "projects", item)
}

// C(R)UD: projects

func (dao *ProjectsDao) ReadProjectList(ctx context.Context) (res []*models.Project, err error) {
	err = dao.ds.ReadAll(ctx, "projects", &res)
	return
}

// C(R)UD: projects

func (dao *ProjectsDao) ReadProject(ctx context.Context, pId int64) (*models.Project, error) {
	res := &models.Project{}
	err := dao.ds.Read(ctx, "projects", res, pId)
	if err == nil {
		return res, nil
	}
	return nil, err
}

// CR(U)D: projects

func (dao *ProjectsDao) UpdateProject(ctx context.Context, item *models.Project) (rowsAffected int64, err error) {
	rowsAffected, err = dao.ds.Update(ctx, "projects", item)
	return
}

// CRU(D): projects

func (dao *ProjectsDao) DeleteProject(ctx context.Context, item *models.Project) (rowsAffected int64, err error) {
	rowsAffected, err = dao.ds.Delete(ctx, "projects", item)
	return
}

func (dao *ProjectsDao) ReadAllRaw(ctx context.Context) (res []*models.ProjectLi, err error) {
	sql := `select p.*, 
		(select count(*) from tasks where p_id=p.p_id) as p_tasks_count 
		from projects p 
		order by p.p_id`
	err = dao.ds.Select(ctx, sql, &res)
	return
}

func (dao *ProjectsDao) GetProjectIds(ctx context.Context) (res []int64, err error) {
	sql := `select p.*, 
		(select count(*) from tasks where p_id=p.p_id) as p_tasks_count 
		from projects p 
		order by p.p_id`
	err = dao.ds.QueryAll(ctx, sql, &res)
	return
}

func (dao *ProjectsDao) GetProjectId(ctx context.Context) (res int64, err error) {
	sql := `select p.*, 
		(select count(*) from tasks where p_id=p.p_id) as p_tasks_count 
		from projects p 
		order by p.p_id`
	err = dao.ds.Query(ctx, sql, &res)
	return
}

package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	projectFieldNames          = builder.RawFieldNames(&Project{})
	projectRows                = strings.Join(projectFieldNames, ",")
	projectRowsExpectAutoSet   = strings.Join(stringx.Remove(projectFieldNames, "`create_time`", "`update_time`"), ",")
	projectRowsWithPlaceHolder = strings.Join(stringx.Remove(projectFieldNames, "`project_id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	ProjectModel interface {
		Insert(data *Project) (sql.Result, error)
		FindOne(projectId string) (*Project, error)
		Update(data *Project) error
		Delete(projectId string) error
	}

	defaultProjectModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Project struct {
		ProjectId    string `db:"project_id"`    // project id
		ProjectToken string `db:"project_token"` // project token
	}
)

func NewProjectModel(conn sqlx.SqlConn) ProjectModel {
	return &defaultProjectModel{
		conn:  conn,
		table: "`project`",
	}
}

func (m *defaultProjectModel) Insert(data *Project) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, projectRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProjectId, data.ProjectToken)
	return ret, err
}

func (m *defaultProjectModel) FindOne(projectId string) (*Project, error) {
	query := fmt.Sprintf("select %s from %s where `project_id` = ? limit 1", projectRows, m.table)
	var resp Project
	err := m.conn.QueryRow(&resp, query, projectId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProjectModel) Update(data *Project) error {
	query := fmt.Sprintf("update %s set %s where `project_id` = ?", m.table, projectRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProjectToken, data.ProjectId)
	return err
}

func (m *defaultProjectModel) Delete(projectId string) error {
	query := fmt.Sprintf("delete from %s where `project_id` = ?", m.table)
	_, err := m.conn.Exec(query, projectId)
	return err
}

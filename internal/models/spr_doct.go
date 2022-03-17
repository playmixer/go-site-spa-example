package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type SprDoctModel struct {
	Db *sql.DB
}

func (m *SprDoctModel) FoundByFIO(lname, fname, sname string) (*[]SprDoct, error) {
	data := []SprDoct{}
	sql := fmt.Sprintf(
		"select kod_dock_i, fio, im, ot FROM SPR_DOCT where position('%s', fio)>0 and position('%s', im)>0 and position('%s', ot)>0",
		lname, fname, sname)
	INFO(sql)
	rows, err := m.Db.Query(sql)
	if err != nil {
		ERROR(err.Error())
		return nil, err
	}
	for rows.Next() {
		row := SprDoct{}
		err = rows.Scan(&row.Id, &row.Lname, &row.Fname, &row.Sname)
		if err != nil {
			ERROR(err.Error())
			return nil, err
		}
		data = append(data, row)

	}
	res, _ := json.Marshal(&data)
	INFO(string(res))
	return &data, nil
}

func (m *SprDoctModel) Get(id int) (*SprDoct, error) {
	data := SprDoct{}
	sql := fmt.Sprintf(
		"select kod_dock_i, fio, im, ot FROM SPR_DOCT where kod_dock_i=%v", id)
	INFO(sql)
	row := m.Db.QueryRow(sql)
	err := row.Scan(&data.Id, &data.Lname, &data.Fname, &data.Sname)
	err = data.ToUTF8()
	if err != nil {
		ERROR(err.Error())
		return nil, err
	}
	data.Trim()

	res, _ := json.Marshal(&data)
	INFO(string(res))

	return &data, nil
}

func (m *SprDoctModel) UserAuth(login, password string) (bool, error) {
	var n int
	sql := fmt.Sprintf(
		"select count(*) FROM SPR_DOCT where kod_dock_i=%v and pass_new='%s'",
		login, password)
	INFO(sql)
	rows := m.Db.QueryRow(sql)
	err := rows.Scan(&n)
	if err != nil {
		ERROR(err.Error())
		return false, err
	}
	res, _ := json.Marshal(n)
	INFO(string(res))
	return n > 0, nil
}

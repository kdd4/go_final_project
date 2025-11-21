package db

import (
	"errors"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64

	query := `
	INSERT INTO scheduler(date, title, comment, repeat) 
	VALUES ($1, $2, $3, $4)
	`

	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)

	if err == nil {
		id, err = res.LastInsertId()
	}

	return id, err
}

func GetTask(id string) (*Task, error) {
	query := `
	SELECT id, date, title, comment, repeat
	FROM scheduler
	WHERE id = $1
	`

	var task Task

	err := db.QueryRow(query, id).Scan(
		&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat,
	)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func UpdateTask(task *Task) error {
    query := `
	UPDATE scheduler 
	SET date = $2, title = $3, comment = $4, repeat = $5
	WHERE id = $1
	`

    res, err := db.Exec(
		query, 
		task.ID, task.Date, task.Title, task.Comment, task.Repeat, 
	)

    if err != nil {
        return err
    }

    count, err := res.RowsAffected()

    if err != nil {
        return err
    }

    if count == 0 {
        return errors.New("incorrect id for updating task")
    }

    return nil
}

func DeleteTask(id string) error {
	query := `
	DELETE FROM scheduler 
	WHERE id = $1
	`

	res, err := db.Exec(query, id)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()

	if err != nil {
        return err
    }

	if count == 0 {
		return errors.New("incorrect id for deleting task")
	}

	return nil
}

func UpdateDate(next string, id string) error {
	query := `
	UPDATE scheduler 
	SET date = $2
	WHERE id = $1
	`

    res, err := db.Exec(query, id, next)

    if err != nil {
        return err
    }

    count, err := res.RowsAffected()

    if err != nil {
        return err
    }

    if count == 0 {
        return errors.New("incorrect id for updating task")
    }

    return nil
}

func Tasks(limit int) ([]*Task, error) {
	// 'order by' works, because date has format YYYYMMDD
	query := `
	SELECT id, date, title, comment, repeat
	FROM scheduler 
	ORDER BY date ASC
	LIMIT $1
	`

	rows, err := db.Query(query, limit)

	if err != nil {
		return nil, err
	}

	tasks := make([]*Task, 0)

	for rows.Next() {
		var task Task

		err := rows.Scan(
			&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	err = rows.Close();

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
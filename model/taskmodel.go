package model

type Task struct {
	tableName struct{} `sql:"task"`
	Id        int64    `sql:"id,pk"`
	SourceId  string   `sql:"source_id"`
	SheetId   string   `sql:"sheet_id"`
	SheetName string   `sql:"sheet_name"`
	IsActive  int8     `sql:"is_active"`
}

type UserProfile struct {
	tableName struct{} `sql:"user_profile"`
	UserId    string   `sql:"user_id,pk"`
	UserName  string   `sql:"user_name"`
	UserType  string   `sql:"user_type"`
}

type TaskUser struct {
	tableName struct{} `sql:"task_user"`
	Id        string   `sql:"id,pk"`
	TaskId    int64    `sql:"task_id"`
	UserId    string   `sql:"user_id"`
}

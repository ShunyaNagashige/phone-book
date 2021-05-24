package database

import "fmt"

type User struct{
	Id int
	Name string
	PhoneNumber string
}

type UserError struct{
	U *User
	Err error
}

func (err *UserError)Error()string{
	return err.Err.Error()
}

//全登録情報の表示
func GetAllUser()error{
	cmd:=fmt.Sprintf("SELECT * FROM %s",tableName)
	rows,err:=dbConn.Query(cmd)
	if err!=nil{
		return &DbError{Cmd:cmd,Err:err}
	}

	fmt.Println("ID，名前，電話番号")
	for rows.Next(){
		var u User
		if err:=rows.Scan(&u.Id,&u.Name,&u.PhoneNumber);err!=nil{
			return &UserError{U:&u,Err:err}
		}
		fmt.Printf("%d, %s, %s\n",u.Id,u.Name,u.PhoneNumber)
	}

	return nil
}

//一人分の情報を登録する
func (u *User)Create()error{
	cmd:=fmt.Sprintf("INSERT INTO %s(name, phone_number) VALUES (?, ?)",tableName)

	//シングルクオーテーションの付け忘れを防ぐために，以下のようにExecに値を渡す形にしたほうが良い(fmt.Sprintfに渡すのではなく)．
	if _,err:=dbConn.Exec(cmd,u.Name,u.PhoneNumber);err!=nil{
		return &DbError{Cmd: cmd,Err: err}
	}

	return nil
}

//一人分の情報を更新する
func (u *User)Save()error{
	cmd:=fmt.Sprintf("UPDATE %s SET name = ?, phone_number = ? WHERE id = ?",tableName)

	if _,err:=dbConn.Exec(cmd,u.Name,u.PhoneNumber,u.Id);err!=nil{
		return &DbError{Cmd: cmd,Err: err}
	}

	return nil
}
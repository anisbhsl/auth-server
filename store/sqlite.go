package store

import(
	"database/sql"
	"github.com/anisbhsl/auth-server/logger"
	"github.com/anisbhsl/auth-server/models"
	_ "github.com/mattn/go-sqlite3"
)

type store struct{
	DB *sql.DB
}

func New() Service{
	db,err:=sql.Open("sqlite3","./auth.db")
	if err!=nil{
		panic(err)
	}
	if err:=db.Ping();err!=nil{
		panic(err)
	}
	logger.Info("Connected to DB")
	s:=&store{DB: db}
	if err:=s.createTables();err!=nil{
		panic(err)
	}
	return s
}

func (s store) createTables() error{
	logger.Info("Creating database tables")
	tx,err:=s.DB.Prepare(`
			CREATE TABLE IF NOT EXISTS user(
				id text UNIQUE NOT NULL PRIMARY KEY,
				name text NOT NULL DEFAULT '',
				email text UNIQUE NOT NULL,
				location text DEFAULT '',
				about text DEFAULT '',
				password_hash text NOT NULL
			);
    `)
	if err!=nil{
		return err
	}
	_,err=tx.Exec()
	return err
}

func (s store) GetUserDetail(id string) (models.User,error){
	logger.Debug("Retrieve user detail")
	row,err:=s.DB.Query(`SELECT id,name,email,location,about,password_hash FROM user where id= ?`,id)
	if err!=nil{
		return models.User{},err
	}
	defer row.Close()

	user:=models.User{}
	for row.Next(){
		row.Scan(&user.ID,&user.Name,&user.Email,&user.Location,&user.About,&user.PasswordHash)
	}
	return user,nil
}

func (s store) 	GetUserByEmail(email string)(models.User,error){
	logger.Debug("Retrieve User by Email")
	row,err:=s.DB.Query(`SELECT id,name,email,location,about,password_hash FROM user where email= ?`,email)
	if err!=nil{
		return models.User{},err
	}
	defer row.Close()

	user:=models.User{}
	for row.Next(){
		row.Scan(&user.ID,&user.Name,&user.Email,&user.Location,&user.About,&user.PasswordHash)
	}
	return user,nil
	return models.User{},nil
}

func (s store) AddUser(user models.User) (string,error){
	logger.Debug("Adding new user")
	query:=`INSERT INTO user(id,name,email,location,about,password_hash) VALUES
			(?, ?, ?, ?, ?, ?);`

	tx,err:=s.DB.Prepare(query)
	if err!=nil{
		return "",err
	}
	_,err=tx.Exec(user.ID,user.Name,user.Email,user.Location,user.About,user.PasswordHash)
	return user.ID,err
}

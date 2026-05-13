- Database drivers translate between Go's generic SQL package and a specific wire protocol (specific data formation). 
- To interact with a SQL database, create a connection pool and carry out method (Exec, etc) of that
- Model: Go struct + methods that own all DB access for one table/entity. In this app, SnippetModel owns everything snippet table related.  

Model syntax:
- type SnippetModel struct {
	DB `*`sql.DB
- }

^gives access to Insert (write), Get (single read), Latest (Multiread... any methods you define with SnippetModel.

Sql syntax: 
- db, err := sql.Open("mysql", dsn); db.Ping() forces real connect 
- m.DB.exec(stmt, "test", "world) -> exec returns result. Can use .LastInsertId(), .RowsAffected()
- QueryRow/ Query: gets many or a single row 

DB/ transaction syntax:
- db.Begin(): tx.Commit(), tx.Rollback()
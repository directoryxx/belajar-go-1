package database

var connection string

// function init ini sama dengan construct di php
func init() {
	connection = "mysql"
}

func GetDatabase() string  {
	return connection
}

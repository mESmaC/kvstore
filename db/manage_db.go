package db

import "fmt"

var databases = make(map[string]map[string]interface{})

func CreateDB(name string) {
	if _, exists := databases[name]; exists {
		fmt.Printf("Database '%s' already exists.\n", name)
		return
	}
	databases[name] = make(map[string]interface{})
	fmt.Printf("Database '%s' created.\n", name)
}

func Insert(dbName, key string, value interface{}) {
	db, exists := databases[dbName]
	if !exists {
		fmt.Printf("Database '%s' does not exist.\n", dbName)
		return
	}
	db[key] = value
	fmt.Printf("Inserted into '%s': %s = %v\n", dbName, key, value)
}

func Query(dbName, key string) (interface{}, bool) {
	db, exists := databases[dbName]
	if !exists {
		fmt.Printf("Database '%s' does not exist.\n", dbName)
		return nil, false
	}
	value, found := db[key]
	return value, found
}

func DeleteKeyFromDB(dbName, key string) {
	db, exists := databases[dbName]
	if !exists {
		fmt.Printf("Database '%s' does not exist.\n", dbName)
		return
	}
	delete(db, key)
	fmt.Printf("Deleted Key '%s' from database  '%s'.\n", key, dbName)
}

func DeleteDB(name string) {
	if _, exists := databases[name]; exists {
		delete(databases, name)
		fmt.Printf("Database '%s' deleted.\n", name)
	} else {
		fmt.Printf("Database '%s' does not exist.\n", name)
	}
}

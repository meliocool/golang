package packageinit

// * Package Initialization
// * When making a package we can create a function that is immediately executed when the package is accessed
// * For example when connecting to a database in a package, we need to initialize a connection to the database
// * Make a function called init to achieve this

var connection string

func init() {
	connection = "MySQL" // * Example
}

func GetDatabase() string {
	return connection
}

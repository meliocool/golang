package packageexample

// * Package is just a place to store and organize code
// * Package is essentially just a folder directory in the Operating System

// * Use UPPER CASE AT THE START
// * Access Modifier -> if Capitalize first letter then other package can access, if lower case then forbidden
// * Access Modifier also applies to Variables, Structs, Function, Interfaces

func SayHelloFromOtherPackage(name string) string {
	return "Hello " + name
}

// * In standard, a Go file can only access other Go file that lives inside the same Package
// * We must use Import to access from another package

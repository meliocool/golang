package helper

// * In Go, when making a test file, use [fileToBeTested]_test.go
// * MUST end with _test.go

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// * Function name for unit test must start with "Test" with capital T
// * For example for function HelloWorld, then the test function name would be TestHelloWorld()
// * The function must also have (t *testing.T) as the parameter and DOES NOT return a value

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Nate")
	if result != "Hello Nate" {
		panic("Result is not Hello Nate")
	}
}

func TestHelloWorldYeah(t *testing.T) {
	result := HelloWorld("Yeah")
	if result != "Hello Yeah" {
		panic("Result is not Hello Yeah")
	}
}

// * To run a unit test, use go test
// * If want to see in detail what functions has been tested, use go test -v (-v means Verbose)
// * For specific testing of a certain function, then use go test -v -run TestFunctionName
// * -run TestFunctionName uses PREFIX, so anything that starts with TestFunctionName WILL BE TESTED
// * Example: TestFunctionName TestFunctionNameYeah will both be tested if go test -v -run TestFunctionName is running
// * Go only check the inside the package, so if its in root and there is only folders there, and no go file, it wont test anything

// * To run every unit test, use go test -v ./... (This will run in EVERY PACKAGE)

// * -- Failing a Unit Test --
// * To Fail a Unit Test, use Fail(), FailNow(), Error(), or Fatal()

// * t.Fail() -> will Fail the unit test, but still continues the next code execution
func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld(" Fail")
	if result != "Hello Fail" {
		t.Fail()
	}
	fmt.Println("TestHelloWorldFail Done!")
}

// * t.FailNow() -> will Fail the unit test right then and there, will not execute the next code
func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld(" Fail")
	if result != "Hello Fail" {
		t.FailNow()
	}
	fmt.Println("TestHelloWorldFailNow Done!") // * Will not be executed
}

// * t.Error(args...) -> Logging or printing an error, but after that will automatically cal Fail(), which makes it FAIL
func TestHelloWorldError(t *testing.T) {
	result := HelloWorld(" Error")
	if result != "Hello Error" {
		t.Error("Result must be Hello Error")
	}
	fmt.Println("TestHelloWorldError Done!") // * Will not be executed
}

// * t.Fatal(args...) -> Same as t.Error(args...) but calls FailNow() which makes the execution stop
func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld(" Fatal")
	if result != "Hello Fatal" {
		t.Fatal("Result must be Hello Fatal")
	}
	fmt.Println("TestHelloWorldFatal Done!") // * Will not be executed
}

// * Assertion using testify
// * To easily assert decisions for testing instead of using plain if else
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Nate")
	assert.Equal(t, "Hello Nate", result, "Result Must be Hello Nate!")
	fmt.Println("Testing Hello World with Assertion Complete!")
}

// * Assertion vs Require
// * Assert -> Calls Fail(), require -> Calls FailNow()
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Nate")
	require.Equal(t, "Hello Nate", result, "Result Must be Hello Nate!")
	fmt.Println("Testing Hello World with Require Complete!")
}

// * Skip Test -> used to skip a test on certain scenario
// * darwin -> MacOS, windows -> windows, linux -> linux
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Only running on windows Machine")
	}
	result := HelloWorld("Nate")
	require.Equal(t, "Hello Nate", result, "Result Must be Hello Nate!")
	fmt.Println("Testing Hello World Complete!")
}

// * Before and After Test
func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")
	m.Run() // Execute ALL Unit Test
	fmt.Println("After Unit Test")
}

// * Sub Test -> Function unit test inside a Function unit test
// * To run only one of the Subtest, use go test -run TestFunctionName/SubTestName
// * Or to run ALL subtest in every function use go test -run /SubTestName
func TestSubTest(t *testing.T) {
	t.Run("Nate", func(t *testing.T) {
		result := HelloWorld("Nate")
		require.Equal(t, "Hello Nate", result, "Result must be Hello Nate")
	})
	t.Run("Lee Jae-in", func(t *testing.T) {
		result := HelloWorld("Lee Jae-in")
		require.Equal(t, "Hello Lee Jae-In", result, "Result must be Hello Lee Jae-in") // * Will Fail
	})
}

// * Table Test -> Provide a data of a slice filled with parameters and expected result for a Unit test
// * Slice will be iterated using SubTest
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Nate)",
			request:  "Nate",
			expected: "Hello Nate",
		},
		{
			name:     "HelloWorld(Colous)",
			request:  "Colous",
			expected: "Hello Colous",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be Hello Nate")
		})
	}
}

// * Benchmark -> Finding the performance of a unit
// * Function name needs to start with Benchmark
// * Must have parameter (b *testing B) and does not return a value
// * File name still ends with _test.go

func BenchmarkHelloWorldTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Test1")
	}
}

// * To run all benchmark in module (including Unit Test), use go test -v -bench=.
// * If want to exclude Unit Test, use go test -bench . -run '^$' -benchmem -count=1

// * If want certain benchmark, then
func BenchmarkHelloWorldTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Test2")
	}
}

// * Sub Benchmark -> same as Testing
// * go test -bench=BenchmarkSub -run '^$' -benchmem -count=1
// * If only a certain sub benchmark, use go test -bench=BenchmarkSub/Test1 -run '^$' -benchmem -count=1
func BenchmarkSub(b *testing.B) {
	b.Run("Test1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Test1")
		}
	})
	b.Run("Test2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Test2")
		}
	})
}

// * Table Benchmark -> same as Testing
func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			"Test3",
			"Gyj",
		},
		{
			"Test4",
			"Lee Jae-in",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

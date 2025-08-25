# Mock Objects in Testing

## What is a Mock?
- A mock is an object programmed with certain expectations.
- When called, it produces the predefined expected result.

## When to Use a Mock
- Useful when dealing with objects that are difficult to test directly.
- Commonly applied to external dependencies such as APIs, databases, or third-party services.

## Example
- Testing a block of code that makes an API call to a third-party service.
- Instead of making a real network call, a mock simulates the API response.

## Tools
- In Go, a popular library for mocking is **Testify Mock**.
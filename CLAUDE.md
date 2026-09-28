# CLAUDE

## Code Styleguide
- Do not abbreviate names for the variables, receiver, methods and structs. The only exception is context (ctx).
- If there are more than 3 parameters for the function input, please create a struct for the input, and copy the function/method name with the `Input` suffix. Keep the context as the first parameter and separate from the input.
- Define sentinel errors where necessary for interfaces.
- When writing new features, please write unit tests for it, particularly for abstract interfaces/implementations. Please ensure the test file package name has `_test`, and use `GoConvey` for the test framework. Use Behaviour Driven Development (BDD) methodology when writing test assertions.
- Each package decides which of its errors are fatal: call `fatal.OnError(err)` in the package for errors that can only come from a bug in it (e.g. building a request from constants, marshalling its own types). Return everything else: sentinels for client/validation errors, and wrapped errors for I/O failures such as timeouts, cancellations and dropped connections, which aren't bugs and mustn't take the service down.
- Only use `merry` in the http api package. Other packages return plain errors: `errors.New` sentinels, wrapped with `fmt.Errorf("...: %w", err)`.
- When implementing a http api endpoint, send errors through the http api package's `httpx.MerrifiedSentinels`, which gives the sentinels of the packages it calls their status codes and user messages, and responds 500 to anything else:
```go
	var err error
	defer func() {
		if err != nil {
			merrifiedSentinels.SendErrorResponse(responseWriter, err)
		}
	}()
```
- If the code is seen more than once, make it a helper function
- If the code is seen more than once across packages of that service, move it to its internal package with a suitable package name
- If the package is used across different services, move it to the outer internal package with a suitable package name
- If external services are expected to communicate to a service, make sure the service defines public package with a front interface, along with its sentinel errors. Move these from private to public where possible to avoid duplicate code.
- Comments lie, so avoid writing comments where it's simpler to read, but do write comments where justified.
- Try to avoid white space where appropriate.

## Git Versioning
- Any changes should be summarised for commit, and pushed to remote.
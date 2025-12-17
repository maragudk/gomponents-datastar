# gomponents-datastar

This project provides HTML attributes for using [Datastar](https://data-star.dev) with gomponents. The attributes are documented at https://data-star.dev/reference/attributes.

## Project Structure

- All code is in `datastar.go`
- All tests are in `datastar_test.go`
- Helper code is in `internal/`

## Development Guidelines

- Use your Go skill when working on this project
- Documentation for attributes should be taken verbatim from the Datastar documentation website, with a link back to the site
- Each function should have a corresponding test with examples
- Follow the existing pattern: each Datastar attribute gets a Go function that returns a `g.Node`
- Modifiers are implemented as variadic parameters of type `Modifier`

## Testing

Run tests with:
```bash
make test
```

## Datastar Compatibility

This library is compatible with Datastar v1.0.0-RC.7 and above.

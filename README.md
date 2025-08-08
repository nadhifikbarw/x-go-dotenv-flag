# `.env` Loading With Flag Control

I like having `.env` support for local development when developing web service. It's simple, and encourages implementation to consider configurability through environment variables.

More importantly for me it offers development portability whenever I need to move around OS doing development, because I may not have good shell always ready.

This is referential setup to support `.env` loading in Go that doesn't get in the way of production release to ease project bootstrapping chores.

## Checklist

- [ ] Provide `.env.example`
- [ ] Add `.env` to `.gitignore`
- [ ] Add `.env` to `.dockerignore` for containerized workflow
- [ ] Register `-env` flag in the project flag parsing logic, prefer allowing to override `.env` path with `.env` as default if unspecified
- [ ] Prefer `.env` loading logic at the beginning `main` function as early as possible after flag parsing.
- [ ] Modify `.air.toml` or other live-reload config to `-env` flag by default. See [`.air.toml`](/.air.toml) for reference
- [ ] Ensure `.env` loading is non-default behavior without the flag

```sh
# Load .env in root repository
go run . -env

# Load .env in nested path within repo
go run . -env=nested_env_example/.env
```

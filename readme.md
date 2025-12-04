# Sonte

**Sonte** (*Stephen's Obstinate Note Taking Engine*) is a command-line note management system, written in Go 1.25 by Stephen Malone.

## Notes 

- Using flag stdlib for subcommands, see prototype script.
- Using bbolt database, `get(name, attr)`, `set(name, attr, data)`.
- Models are just `Type { DB, Name }` with methods to get attributes like `Note.Hash()`.
- Add `test.Get` and `test.Set` that `t.Fatal` on error, then `test.MockDB` uses `test.Set` for each item.

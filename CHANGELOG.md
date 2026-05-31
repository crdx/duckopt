# Changelog

## [2.3.0] - 2024-03-24

- Replace tabs in usage string with spaces.

## [2.0.1] - 2023-01-19

- Fix example in readme.

## [2.0.0] - 2023-01-19

- `Parse` has been replaced with `Parse`/`MustParse`.
- New generic methods `Bind`/`MustBind` allow parsing and binding in one call, returning the instantiated struct directly.

## [1.0.3] - 2023-01-19

- Upgrade dependencies.

## [1.0.2] - 2023-01-18

- Update import paths to use a vanity domain.

## [1.0.1] - 2023-01-17

- Move code out of `src/` as Go does not support such a directory structure. Put it in `internal/` instead.

## [1.0.0] - 2021-12-17

- Initial release.

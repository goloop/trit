# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0]

Major release. The module path is now `github.com/goloop/trit/v2` and the
minimum Go version is 1.24.

### Fixed
- `Random` never returned `True` for the default probability (and was skewed
  for others). The remaining probability after `Unknown` is now split evenly
  between `True` and `False` for any percentage.
- `Clean` was a no-op; it now unconditionally resets the value to `Unknown`,
  matching its documented behaviour.
- Removed a data race: the aggregates no longer share global mutable state.
- Empty-input handling of the aggregates was inconsistent; it now follows a
  single, documented formal-logic convention (see Changed).

### Changed
- `logicToTrit` no longer uses reflection; conversion is a plain, allocation-
  free type switch.
- `All`, `Any`, `None`, `Known` are now linear and allocation-free (the
  goroutine-based implementation and the `ParallelTasks` knob were removed —
  the per-element work is far too small to benefit from parallelism).
- Empty-input convention: `All()`/`None()`/`Known()` return `True` (vacuous
  truth), `Any()` returns `False`, `Consensus()`/`Majority()` return `Unknown`.
- `Random` now uses `math/rand/v2` and is no longer re-seeded per call.
- `UnmarshalJSON` additionally accepts JSON numbers (sign-based mapping).

### Added
- `ParseTrit` and `encoding.TextMarshaler`/`TextUnmarshaler` support.
- `database/sql` integration via `driver.Valuer` and `sql.Scanner`, mapping
  `Unknown` to SQL `NULL`.
- `Compare` for the natural ordering `False < Unknown < True`.
- `ErrInvalidTrit` sentinel error for parsing/scanning failures.
- Runnable examples, exhaustive property tests, and fuzz targets.

### Removed
- `ParallelTasks` (parallel execution knob).
- `IsConfidence` (duplicated `Known`; use `Known` instead).

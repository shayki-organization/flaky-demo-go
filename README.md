# flaky-demo-go

Demo repository for exercising a flaky-test detector against GitHub Actions logs (go test).

- `calc_test.go` contains stable tests plus two intentionally flaky tests
  (`TestFlaky*`) that fail with ~40% probability per run.
- `.github/workflows/tests.yml` (`Tests` → `unit-tests`) runs `go test ./... -v` so
  the standard `--- FAIL: TestName (0.00s)` lines appear in the log.

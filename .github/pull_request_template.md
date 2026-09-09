<!--
PR title must follow: <type>(<scope>): <summary>   (same as a commit message)
example: feat(subscription): add billing reminder RPCs
-->

## What

<!-- one or two sentences -->

## Why

<!-- what needs this contract change -->

## Breaking change?

- [ ] No — additive only (new field / new RPC within `v1`)
- [ ] Yes — new `v2/` directory added; consumers must regenerate and update call sites

## Checklist

- [ ] `buf lint` passes
- [ ] `buf format -w` run
- [ ] `buf generate` run and `gen/go/` committed (no drift)
- [ ] `buf breaking --against '.git#branch=main'` clean (or new `v2/`)
- [ ] `go build ./...` passes
- [ ] plugin versions unchanged, or README + CONTRIBUTING updated to match
- [ ] tag bump planned after merge if the surface changed

# mongopiet — open items / TODO

Notes for later. Captured 2026-06-08 while rewriting the README and surveying real usage
across `w1nterbot`, `nofy-api`, `lutti-discord`, `inventory`, `orbe`.

## Library (code)

- [ ] **No per-call context.** Most functions run on a package-internal `context.TODO()`
      (`pkg/db/main.go`); only `AggregateCtx` takes a `context.Context`. `Find`/`Aggregate`
      hardcode a 30s timeout. Projects drop to `db.DB.Collection(...)` directly when they need
      request-scoped cancellation. Add `*Ctx` variants (`FindOneCtx`, `FindCtx`, `InsertOneCtx`,
      `UpdateOneCtx`, …) or make context the first arg across the board.

- [ ] **`Document.Save()` doesn't track changes through pointers.** `NewDoc` keeps a shallow
      copy as the baseline, so mutating data *behind* a pointer field isn't detected by
      `deepCompare` (`pkg/db/document.go`). Existing in-code TODOs there: handle `omitempty`,
      handle `$unset`, verify slice/array behavior, traverse pointers / deep-copy the baseline.
      Until fixed, document (done) that pointer-heavy updates should use `db.UpdateOne`.

- [ ] **Not-found behavior is inconsistent.** `db.FindOne[T]` returns `mongo.ErrNoDocuments`,
      but `Document.FindOne` swallows it and returns `nil, nil`. Pick one contract (probably a
      typed `ErrNotFound`) and apply it everywhere.

- [ ] **Naive collection pluralization.** `Document.CollectionName()` = lowercase struct name
      + `s` (+`_ts` for time-series). No handling of irregular plurals or custom names; a struct
      rename silently changes the collection. Consider an optional `CollectionName()` override
      or a struct tag.

- [ ] **`pkg/bulk/types.go` last type looks wrong.** `type WriteModel mongo.IndexModel` is a
      defined type (not an alias) and duplicates the `Write = mongo.WriteModel` alias at the top.
      Looks like a leftover — should likely be `type Index = mongo.IndexModel` (alias) or be
      removed.

- [ ] **No upsert / find-or-create helper.** Projects hand-roll `mongo.NewUpdateOneModel().SetUpsert(true)`
      via `BulkWrite`. A typed `db.Upsert[T]` / `db.FindOneAndUpdate[T]` would remove that.

- [ ] **`Document[T]` still marked experimental** in the README. Decide whether to harden it
      (tests + the Save() fixes above) or keep it a thin convenience over the plain functions.

## README TODO items (from the old README)

- [ ] "Add propagation" — clarify what this meant (cascade writes? relation loading?) or drop.
- [ ] "Limits etc" — covered partly by `pkg/opts`; document any remaining gaps.

## Docs (addressed in the README rewrite)

- [x] Flatten the package layout — `pkg/db` / `pkg/opts` / `pkg/bulk` moved to the module root
      (`db` / `opts` / `bulk`) in v0.10.0; README uses the new paths. Consumers on older versions
      are unaffected until they upgrade (and then update their imports).
- [x] Fix the `UpdateOne` example (missing `unset` arg — signature is `(coll, filter, set, unset, ...)`).
- [x] Fix the broken type-alias example (`db.Document[UserDocument]` → `db.Document[User]`).
- [x] Document the previously-missing API: `CountDocuments`, `Aggregate`/`AggregateCtx`,
      `BulkWrite` + `bulk` models, `InsertMany`, `UpdateMany`/`DeleteMany`, `pkg/opts` pagination.
- [x] Document not-found semantics and the `ModifiedCount` check pattern.

## Nice-to-have

- [ ] **Expand tests.** `main_test.go` and `opts/main_test.go` exist; add coverage for `Save()`
      diffing, `checkFields` (auto id/timestamps), pluralization, and the not-found paths.
- [ ] CHANGELOG for the `v0.x` line.
- [ ] Runnable example under an `_examples/` dir.

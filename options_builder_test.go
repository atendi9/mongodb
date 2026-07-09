package mongodb

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func TestBuildFindManyOptions(t *testing.T) {
	t.Run("no options returns usable defaults", func(t *testing.T) {
		resolved := resolveOptions(t, BuildFindManyOptions())
		if resolved == nil {
			t.Fatal("expected non-nil resolved options")
		}
	})

	t.Run("maps every provided field", func(t *testing.T) {
		collation := &options.Collation{Locale: "en_US"}
		in := &options.FindOptions{
			AllowDiskUse:        ptr(true),
			AllowPartialResults: ptr(true),
			BatchSize:           ptr(int32(50)),
			Limit:               ptr(int64(10)),
			Skip:                ptr(int64(5)),
			NoCursorTimeout:     ptr(true),
			ReturnKey:           ptr(true),
			Sort:                map[string]any{"age": -1},
			Comment:             "find-many",
			Hint:                "age_1",
			Projection:          map[string]any{"name": 1},
			Collation:           collation,
		}

		got := resolveOptions(t, BuildFindManyOptions(in))

		if got.AllowDiskUse == nil || !*got.AllowDiskUse {
			t.Errorf("AllowDiskUse not mapped: %v", got.AllowDiskUse)
		}
		if got.AllowPartialResults == nil || !*got.AllowPartialResults {
			t.Errorf("AllowPartialResults not mapped: %v", got.AllowPartialResults)
		}
		if got.BatchSize == nil || *got.BatchSize != 50 {
			t.Errorf("BatchSize not mapped: %v", got.BatchSize)
		}
		if got.Limit == nil || *got.Limit != 10 {
			t.Errorf("Limit not mapped: %v", got.Limit)
		}
		if got.Skip == nil || *got.Skip != 5 {
			t.Errorf("Skip not mapped: %v", got.Skip)
		}
		if got.NoCursorTimeout == nil || !*got.NoCursorTimeout {
			t.Errorf("NoCursorTimeout not mapped: %v", got.NoCursorTimeout)
		}
		if got.ReturnKey == nil || !*got.ReturnKey {
			t.Errorf("ReturnKey not mapped: %v", got.ReturnKey)
		}
		if !reflect.DeepEqual(got.Sort, in.Sort) {
			t.Errorf("Sort not mapped: %v", got.Sort)
		}
		if deref(got.Comment) != "find-many" {
			t.Errorf("Comment not mapped: %v", got.Comment)
		}
		if got.Hint != "age_1" {
			t.Errorf("Hint not mapped: %v", got.Hint)
		}
		if !reflect.DeepEqual(got.Projection, in.Projection) {
			t.Errorf("Projection not mapped: %v", got.Projection)
		}
		if got.Collation == nil || got.Collation.Locale != "en_US" {
			t.Errorf("Collation not mapped: %v", got.Collation)
		}
	})
}

func TestBuildFindOneOptions(t *testing.T) {
	t.Run("no options returns usable defaults", func(t *testing.T) {
		if resolveOptions(t, BuildFindOneOptions()) == nil {
			t.Fatal("expected non-nil resolved options")
		}
	})

	t.Run("maps every provided field", func(t *testing.T) {
		in := &options.FindOneOptions{
			AllowPartialResults: ptr(true),
			Skip:                ptr(int64(3)),
			ReturnKey:           ptr(true),
			Sort:                map[string]any{"name": 1},
			Comment:             "find-one",
			Hint:                "name_1",
			Projection:          map[string]any{"email": 1},
			Collation:           &options.Collation{Locale: "pt_BR"},
		}

		got := resolveOptions(t, BuildFindOneOptions(in))

		if got.AllowPartialResults == nil || !*got.AllowPartialResults {
			t.Errorf("AllowPartialResults not mapped: %v", got.AllowPartialResults)
		}
		if got.Skip == nil || *got.Skip != 3 {
			t.Errorf("Skip not mapped: %v", got.Skip)
		}
		if got.ReturnKey == nil || !*got.ReturnKey {
			t.Errorf("ReturnKey not mapped: %v", got.ReturnKey)
		}
		if !reflect.DeepEqual(got.Sort, in.Sort) {
			t.Errorf("Sort not mapped: %v", got.Sort)
		}
		if deref(got.Comment) != "find-one" {
			t.Errorf("Comment not mapped: %v", got.Comment)
		}
		if got.Collation == nil || got.Collation.Locale != "pt_BR" {
			t.Errorf("Collation not mapped: %v", got.Collation)
		}
	})
}

func TestBuildUpdateOneOptions(t *testing.T) {
	arrayFilters := []any{map[string]any{"x.qty": 1}}
	in := &options.UpdateOneOptions{
		ArrayFilters:             arrayFilters,
		BypassDocumentValidation: ptr(true),
		Upsert:                   ptr(true),
		Sort:                     map[string]any{"a": 1},
		Comment:                  "update-one",
		Hint:                     "a_1",
		Let:                      map[string]any{"v": 1},
		Collation:                &options.Collation{Locale: "en"},
	}

	got := resolveOptions(t, BuildUpdateOneOptions(in))

	if !reflect.DeepEqual(got.ArrayFilters, arrayFilters) {
		t.Errorf("ArrayFilters not mapped: %v", got.ArrayFilters)
	}
	if got.BypassDocumentValidation == nil || !*got.BypassDocumentValidation {
		t.Errorf("BypassDocumentValidation not mapped: %v", got.BypassDocumentValidation)
	}
	if got.Upsert == nil || !*got.Upsert {
		t.Errorf("Upsert not mapped: %v", got.Upsert)
	}
	if !reflect.DeepEqual(got.Sort, in.Sort) {
		t.Errorf("Sort not mapped: %v", got.Sort)
	}
	if got.Comment != "update-one" {
		t.Errorf("Comment not mapped: %v", got.Comment)
	}
	if got.Collation == nil || got.Collation.Locale != "en" {
		t.Errorf("Collation not mapped: %v", got.Collation)
	}

	if resolveOptions(t, BuildUpdateOneOptions()) == nil {
		t.Fatal("expected non-nil defaults")
	}
}

func TestBuildUpdateManyOptions(t *testing.T) {
	arrayFilters := []any{map[string]any{"x.qty": map[string]any{"$gte": 1}}}
	in := &options.UpdateManyOptions{
		ArrayFilters:             arrayFilters,
		BypassDocumentValidation: ptr(true),
		Upsert:                   ptr(true),
		Comment:                  "update-many",
		Hint:                     "a_1",
		Let:                      map[string]any{"v": 1},
		Collation:                &options.Collation{Locale: "en"},
	}

	got := resolveOptions(t, BuildUpdateManyOptions(in))

	if !reflect.DeepEqual(got.ArrayFilters, arrayFilters) {
		t.Errorf("ArrayFilters not mapped: %v", got.ArrayFilters)
	}
	if got.Upsert == nil || !*got.Upsert {
		t.Errorf("Upsert not mapped: %v", got.Upsert)
	}
	if got.Comment != "update-many" {
		t.Errorf("Comment not mapped: %v", got.Comment)
	}

	if resolveOptions(t, BuildUpdateManyOptions()) == nil {
		t.Fatal("expected non-nil defaults")
	}
}

func TestBuildInsertManyOptions(t *testing.T) {
	in := &options.InsertManyOptions{
		BypassDocumentValidation: ptr(true),
		Ordered:                  ptr(false),
		Comment:                  "insert-many",
	}

	got := resolveOptions(t, BuildInsertManyOptions(in))

	if got.BypassDocumentValidation == nil || !*got.BypassDocumentValidation {
		t.Errorf("BypassDocumentValidation not mapped: %v", got.BypassDocumentValidation)
	}
	if got.Ordered == nil || *got.Ordered {
		t.Errorf("Ordered not mapped: %v", got.Ordered)
	}
	if got.Comment != "insert-many" {
		t.Errorf("Comment not mapped: %v", got.Comment)
	}

	if resolveOptions(t, BuildInsertManyOptions()) == nil {
		t.Fatal("expected non-nil defaults")
	}
}

func TestBuildCountOptions(t *testing.T) {
	in := &options.CountOptions{
		Limit:     ptr(int64(100)),
		Skip:      ptr(int64(10)),
		Comment:   "count",
		Hint:      "a_1",
		Collation: &options.Collation{Locale: "en"},
	}

	got := resolveOptions(t, BuildCountOptions(in))

	if got.Limit == nil || *got.Limit != 100 {
		t.Errorf("Limit not mapped: %v", got.Limit)
	}
	if got.Skip == nil || *got.Skip != 10 {
		t.Errorf("Skip not mapped: %v", got.Skip)
	}
	if got.Comment != "count" {
		t.Errorf("Comment not mapped: %v", got.Comment)
	}
	if got.Collation == nil || got.Collation.Locale != "en" {
		t.Errorf("Collation not mapped: %v", got.Collation)
	}

	if resolveOptions(t, BuildCountOptions()) == nil {
		t.Fatal("expected non-nil defaults")
	}
}

func TestBuildDistinctOptions(t *testing.T) {
	in := &options.DistinctOptions{
		Comment:   "distinct",
		Hint:      "a_1",
		Collation: &options.Collation{Locale: "en"},
	}

	got := resolveOptions(t, BuildDistinctOptions(in))

	if got.Comment != "distinct" {
		t.Errorf("Comment not mapped: %v", got.Comment)
	}
	if got.Collation == nil || got.Collation.Locale != "en" {
		t.Errorf("Collation not mapped: %v", got.Collation)
	}

	if resolveOptions(t, BuildDistinctOptions()) == nil {
		t.Fatal("expected non-nil defaults")
	}
}

func TestBuildFindOneAndUpdateOptions(t *testing.T) {
	after := options.After
	in := &options.FindOneAndUpdateOptions{
		ArrayFilters:             []any{map[string]any{"x": 1}},
		BypassDocumentValidation: ptr(true),
		ReturnDocument:           &after,
		Upsert:                   ptr(true),
		Sort:                     map[string]any{"a": 1},
		Projection:               map[string]any{"name": 1},
		Comment:                  "foau",
		Hint:                     "a_1",
		Let:                      map[string]any{"v": 1},
		Collation:                &options.Collation{Locale: "en"},
	}

	got := resolveOptions(t, BuildFindOneAndUpdateOptions(in))

	if got.ReturnDocument == nil || *got.ReturnDocument != options.After {
		t.Errorf("ReturnDocument not mapped: %v", got.ReturnDocument)
	}
	if got.Upsert == nil || !*got.Upsert {
		t.Errorf("Upsert not mapped: %v", got.Upsert)
	}
	if !reflect.DeepEqual(got.Projection, in.Projection) {
		t.Errorf("Projection not mapped: %v", got.Projection)
	}
	if got.Comment != "foau" {
		t.Errorf("Comment not mapped: %v", got.Comment)
	}
	if got.Collation == nil || got.Collation.Locale != "en" {
		t.Errorf("Collation not mapped: %v", got.Collation)
	}

	if resolveOptions(t, BuildFindOneAndUpdateOptions()) == nil {
		t.Fatal("expected non-nil defaults")
	}
}

func TestBuildFindOneAndDeleteOptions(t *testing.T) {
	in := &options.FindOneAndDeleteOptions{
		Sort:       map[string]any{"a": 1},
		Projection: map[string]any{"name": 1},
		Comment:    "foad",
		Hint:       "a_1",
		Let:        map[string]any{"v": 1},
		Collation:  &options.Collation{Locale: "en"},
	}

	got := resolveOptions(t, BuildFindOneAndDeleteOptions(in))

	if !reflect.DeepEqual(got.Sort, in.Sort) {
		t.Errorf("Sort not mapped: %v", got.Sort)
	}
	if !reflect.DeepEqual(got.Projection, in.Projection) {
		t.Errorf("Projection not mapped: %v", got.Projection)
	}
	if got.Comment != "foad" {
		t.Errorf("Comment not mapped: %v", got.Comment)
	}
	if got.Collation == nil || got.Collation.Locale != "en" {
		t.Errorf("Collation not mapped: %v", got.Collation)
	}

	if resolveOptions(t, BuildFindOneAndDeleteOptions()) == nil {
		t.Fatal("expected non-nil defaults")
	}
}

func TestBuildFindOneAndReplaceOptions(t *testing.T) {
	before := options.Before
	in := &options.FindOneAndReplaceOptions{
		BypassDocumentValidation: ptr(true),
		ReturnDocument:           &before,
		Upsert:                   ptr(true),
		Sort:                     map[string]any{"a": 1},
		Projection:               map[string]any{"name": 1},
		Comment:                  "foar",
		Hint:                     "a_1",
		Let:                      map[string]any{"v": 1},
		Collation:                &options.Collation{Locale: "en"},
	}

	got := resolveOptions(t, BuildFindOneAndReplaceOptions(in))

	if got.ReturnDocument == nil || *got.ReturnDocument != options.Before {
		t.Errorf("ReturnDocument not mapped: %v", got.ReturnDocument)
	}
	if got.Upsert == nil || !*got.Upsert {
		t.Errorf("Upsert not mapped: %v", got.Upsert)
	}
	if got.Comment != "foar" {
		t.Errorf("Comment not mapped: %v", got.Comment)
	}
	if got.Collation == nil || got.Collation.Locale != "en" {
		t.Errorf("Collation not mapped: %v", got.Collation)
	}

	if resolveOptions(t, BuildFindOneAndReplaceOptions()) == nil {
		t.Fatal("expected non-nil defaults")
	}
}

func TestBuildClientOptions(t *testing.T) {
	const uri = "mongodb://localhost:27017"

	t.Run("nil overrides applies URI only", func(t *testing.T) {
		got := BuildClientOptions(uri, nil)
		if got == nil {
			t.Fatal("expected non-nil client options")
		}
		if len(got.Hosts) == 0 || got.Hosts[0] != "localhost:27017" {
			t.Errorf("URI host not applied: %v", got.Hosts)
		}
		if got.MaxAdaptiveRetries != nil || got.EnableOverloadRetargeting != nil {
			t.Errorf("adaptive-retry fields should be unset by default")
		}
	})

	t.Run("maps the v2.7 adaptive-retry features", func(t *testing.T) {
		in := &options.ClientOptions{
			MaxAdaptiveRetries:        ptr(uint(5)),
			EnableOverloadRetargeting: ptr(true),
		}

		got := BuildClientOptions(uri, in)

		if got.MaxAdaptiveRetries == nil || *got.MaxAdaptiveRetries != 5 {
			t.Errorf("MaxAdaptiveRetries not mapped: %v", got.MaxAdaptiveRetries)
		}
		if got.EnableOverloadRetargeting == nil || !*got.EnableOverloadRetargeting {
			t.Errorf("EnableOverloadRetargeting not mapped: %v", got.EnableOverloadRetargeting)
		}
		// URI base must still be honored alongside the overrides.
		if len(got.Hosts) == 0 || got.Hosts[0] != "localhost:27017" {
			t.Errorf("URI host not preserved: %v", got.Hosts)
		}
	})

	t.Run("maps the common retry and app-name toggles", func(t *testing.T) {
		in := &options.ClientOptions{
			RetryWrites: ptr(false),
			RetryReads:  ptr(true),
			AppName:     ptr("mongodb-lib"),
		}

		got := BuildClientOptions(uri, in)

		if got.RetryWrites == nil || *got.RetryWrites {
			t.Errorf("RetryWrites not mapped: %v", got.RetryWrites)
		}
		if got.RetryReads == nil || !*got.RetryReads {
			t.Errorf("RetryReads not mapped: %v", got.RetryReads)
		}
		if got.AppName == nil || *got.AppName != "mongodb-lib" {
			t.Errorf("AppName not mapped: %v", got.AppName)
		}
	})

	t.Run("leaves untouched fields at zero", func(t *testing.T) {
		got := BuildClientOptions(uri, &options.ClientOptions{
			MaxAdaptiveRetries: ptr(uint(2)),
		})
		if got.EnableOverloadRetargeting != nil {
			t.Errorf("EnableOverloadRetargeting should remain unset: %v", got.EnableOverloadRetargeting)
		}
	})
}

func TestBuildDatabaseOptions(t *testing.T) {
	t.Run("nil options returns usable defaults", func(t *testing.T) {
		if resolveOptions(t, BuildDatabaseOptions(nil)) == nil {
			t.Fatal("expected non-nil resolved options for nil input")
		}
	})

	t.Run("maps BSONOptions", func(t *testing.T) {
		in := &options.DatabaseOptions{
			BSONOptions: &options.BSONOptions{UseJSONStructTags: true},
		}
		got := resolveOptions(t, BuildDatabaseOptions(in))
		if got.BSONOptions == nil || !got.BSONOptions.UseJSONStructTags {
			t.Errorf("BSONOptions not mapped: %v", got.BSONOptions)
		}
	})
}

// resolveOptions materializes the concrete options struct produced by a
// Lister by applying every option setter to a zero value of T.
//
// This lets us assert, without a live MongoDB, that the Build* helpers wire
// each field to the correct driver setter against the current driver version.
func resolveOptions[T any](t *testing.T, lister options.Lister[T]) *T {
	t.Helper()
	if lister == nil {
		t.Fatal("lister is nil")
	}
	var out T
	for _, set := range lister.List() {
		if err := set(&out); err != nil {
			t.Fatalf("apply option: %v", err)
		}
	}
	return &out
}

func ptr[T any](v T) *T { return &v }

// deref unwraps an any value that some driver setters (notably the FindOptions
// and FindOneOptions SetComment) store as a pointer to the supplied argument.
func deref(v any) any {
	if p, ok := v.(*any); ok {
		return *p
	}
	return v
}

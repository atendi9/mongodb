package mongodb

import (
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func BuildDatabaseOptions(
	opts *options.DatabaseOptions,
) options.Lister[options.DatabaseOptions] {
	dbOpts := options.Database()
	lister := applyOptions(dbOpts,
		setPtr(opts.BSONOptions, dbOpts.SetBSONOptions),
		setPtr(opts.ReadConcern, dbOpts.SetReadConcern),
		setPtr(opts.ReadPreference, dbOpts.SetReadPreference),
		setPtr(opts.Registry, dbOpts.SetRegistry),
		setPtr(opts.WriteConcern, dbOpts.SetWriteConcern),
	)
	return lister
}

func BuildFindManyOptions(
	opts ...*options.FindOptions,
) options.Lister[options.FindOptions] {
	findOpts := options.Find()
	if len(opts) == 0 {
		return findOpts
	}
	o := opts[0]
	lister := applyOptions(findOpts,
		setIf(o.AllowDiskUse, findOpts.SetAllowDiskUse),
		setIf(o.AllowPartialResults, findOpts.SetAllowPartialResults),
		setIf(o.BatchSize, findOpts.SetBatchSize),
		setIf(o.Limit, findOpts.SetLimit),
		setIf(o.Skip, findOpts.SetSkip),
		setIf(o.MaxAwaitTime, findOpts.SetMaxAwaitTime),
		setIf(o.NoCursorTimeout, findOpts.SetNoCursorTimeout),
		setIf(&o.Sort, findOpts.SetSort),
		setIf(&o.Comment, findOpts.SetComment),
		setIf(&o.Hint, findOpts.SetHint),
		setIf(&o.Let, findOpts.SetLet),
		setIf(&o.Max, findOpts.SetMax),
		setIf(&o.Min, findOpts.SetMin),
		setIf(&o.Projection, findOpts.SetProjection),
		setIf(o.ReturnKey, findOpts.SetReturnKey),
		setIf(o.CursorType, findOpts.SetCursorType),
		setPtr(o.Collation, findOpts.SetCollation),
	)
	return lister
}

func BuildFindOneOptions(
	opts ...*options.FindOneOptions,
) options.Lister[options.FindOneOptions] {
	findOneOpts := options.FindOne()
	if len(opts) == 0 {
		return findOneOpts
	}
	o := opts[0]
	lister := applyOptions(findOneOpts,
		setIf(o.AllowPartialResults, findOneOpts.SetAllowPartialResults),
		setIf(o.Skip, findOneOpts.SetSkip),
		setIf(&o.Sort, findOneOpts.SetSort),
		setIf(&o.Comment, findOneOpts.SetComment),
		setIf(&o.Hint, findOneOpts.SetHint),
		setIf(&o.Max, findOneOpts.SetMax),
		setIf(&o.Min, findOneOpts.SetMin),
		setIf(&o.Projection, findOneOpts.SetProjection),
		setIf(o.ReturnKey, findOneOpts.SetReturnKey),
		setPtr(o.Collation, findOneOpts.SetCollation),
	)
	return lister
}

func BuildUpdateOneOptions(
	opts ...*options.UpdateOneOptions,
) options.Lister[options.UpdateOneOptions] {
	updateOneOpts := options.UpdateOne()
	if len(opts) == 0 {
		return updateOneOpts
	}
	o := opts[0]
	lister := applyOptions(updateOneOpts,
		setIf(&o.ArrayFilters, updateOneOpts.SetArrayFilters),
		setIf(o.BypassDocumentValidation, updateOneOpts.SetBypassDocumentValidation),
		setIf(&o.Sort, updateOneOpts.SetSort),
		setIf(&o.Comment, updateOneOpts.SetComment),
		setIf(&o.Hint, updateOneOpts.SetHint),
		setIf(&o.Let, updateOneOpts.SetLet),
		setIf(o.Upsert, updateOneOpts.SetUpsert),
		setPtr(o.Collation, updateOneOpts.SetCollation),
	)
	return lister
}

func BuildUpdateManyOptions(
	opts ...*options.UpdateManyOptions,
) options.Lister[options.UpdateManyOptions] {
	updateManyOpts := options.UpdateMany()
	if len(opts) == 0 {
		return updateManyOpts
	}
	o := opts[0]
	lister := applyOptions(updateManyOpts,
		setIf(&o.ArrayFilters, updateManyOpts.SetArrayFilters),
		setIf(o.BypassDocumentValidation, updateManyOpts.SetBypassDocumentValidation),
		setIf(&o.Comment, updateManyOpts.SetComment),
		setIf(&o.Hint, updateManyOpts.SetHint),
		setIf(&o.Let, updateManyOpts.SetLet),
		setIf(o.Upsert, updateManyOpts.SetUpsert),
		setPtr(o.Collation, updateManyOpts.SetCollation),
	)
	return lister
}

func BuildInsertManyOptions(
	opts ...*options.InsertManyOptions,
) options.Lister[options.InsertManyOptions] {
	insertManyOpts := options.InsertMany()
	if len(opts) == 0 {
		return insertManyOpts
	}
	o := opts[0]
	lister := applyOptions(insertManyOpts,
		setIf(o.BypassDocumentValidation, insertManyOpts.SetBypassDocumentValidation),
		setIf(o.Ordered, insertManyOpts.SetOrdered),
		setIf(&o.Comment, insertManyOpts.SetComment),
	)
	return lister
}

func BuildCountOptions(
	opts ...*options.CountOptions,
) options.Lister[options.CountOptions] {
	countOpts := options.Count()
	if len(opts) == 0 {
		return countOpts
	}
	o := opts[0]
	lister := applyOptions(countOpts,
		setIf(o.Limit, countOpts.SetLimit),
		setIf(o.Skip, countOpts.SetSkip),
		setIf(&o.Comment, countOpts.SetComment),
		setIf(&o.Hint, countOpts.SetHint),
		setPtr(o.Collation, countOpts.SetCollation),
	)
	return lister
}

func BuildDistinctOptions(
	opts ...*options.DistinctOptions,
) options.Lister[options.DistinctOptions] {
	distinctOpts := options.Distinct()
	if len(opts) == 0 {
		return distinctOpts
	}
	o := opts[0]
	lister := applyOptions(distinctOpts,
		setIf(&o.Comment, distinctOpts.SetComment),
		setIf(&o.Hint, distinctOpts.SetHint),
		setPtr(o.Collation, distinctOpts.SetCollation),
	)
	return lister
}

func BuildFindOneAndUpdateOptions(
	opts ...*options.FindOneAndUpdateOptions,
) options.Lister[options.FindOneAndUpdateOptions] {
	findOneAndUpdateOpts := options.FindOneAndUpdate()
	if len(opts) == 0 {
		return findOneAndUpdateOpts
	}
	o := opts[0]
	lister := applyOptions(findOneAndUpdateOpts,
		setIf(&o.ArrayFilters, findOneAndUpdateOpts.SetArrayFilters),
		setIf(o.BypassDocumentValidation, findOneAndUpdateOpts.SetBypassDocumentValidation),
		setIf(o.ReturnDocument, findOneAndUpdateOpts.SetReturnDocument),
		setIf(o.Upsert, findOneAndUpdateOpts.SetUpsert),
		setIf(&o.Sort, findOneAndUpdateOpts.SetSort),
		setIf(&o.Projection, findOneAndUpdateOpts.SetProjection),
		setIf(&o.Comment, findOneAndUpdateOpts.SetComment),
		setIf(&o.Hint, findOneAndUpdateOpts.SetHint),
		setIf(&o.Let, findOneAndUpdateOpts.SetLet),
		setPtr(o.Collation, findOneAndUpdateOpts.SetCollation),
	)
	return lister
}

func BuildFindOneAndDeleteOptions(
	opts ...*options.FindOneAndDeleteOptions,
) options.Lister[options.FindOneAndDeleteOptions] {
	findOneAndDeleteOpts := options.FindOneAndDelete()
	if len(opts) == 0 {
		return findOneAndDeleteOpts
	}
	o := opts[0]
	lister := applyOptions(findOneAndDeleteOpts,
		setIf(&o.Sort, findOneAndDeleteOpts.SetSort),
		setIf(&o.Projection, findOneAndDeleteOpts.SetProjection),
		setIf(&o.Comment, findOneAndDeleteOpts.SetComment),
		setIf(&o.Hint, findOneAndDeleteOpts.SetHint),
		setIf(&o.Let, findOneAndDeleteOpts.SetLet),
		setPtr(o.Collation, findOneAndDeleteOpts.SetCollation),
	)
	return lister
}

func BuildFindOneAndReplaceOptions(
	opts ...*options.FindOneAndReplaceOptions,
) options.Lister[options.FindOneAndReplaceOptions] {
	findOneAndReplaceOpts := options.FindOneAndReplace()
	if len(opts) == 0 {
		return findOneAndReplaceOpts
	}
	o := opts[0]
	lister := applyOptions(findOneAndReplaceOpts,
		setIf(o.BypassDocumentValidation, findOneAndReplaceOpts.SetBypassDocumentValidation),
		setIf(o.ReturnDocument, findOneAndReplaceOpts.SetReturnDocument),
		setIf(o.Upsert, findOneAndReplaceOpts.SetUpsert),
		setIf(&o.Sort, findOneAndReplaceOpts.SetSort),
		setIf(&o.Projection, findOneAndReplaceOpts.SetProjection),
		setIf(&o.Comment, findOneAndReplaceOpts.SetComment),
		setIf(&o.Hint, findOneAndReplaceOpts.SetHint),
		setIf(&o.Let, findOneAndReplaceOpts.SetLet),
		setPtr(o.Collation, findOneAndReplaceOpts.SetCollation),
	)
	return lister
}

// optionMutation applies a single change to an options builder of type O.
type optionMutation[O any] func(*O) *O

// setIf returns a mutation that dereferences value and calls setter only when
// value is non-nil.
func setIf[O, V any](value *V, setter func(V) *O) optionMutation[O] {
	return func(builder *O) *O {
		if value == nil {
			return builder
		}
		return setter(*value)
	}
}

// setPtr returns a mutation for setters that already accept a pointer, applying
// it only when value is non-nil.
func setPtr[O, V any](value *V, setter func(*V) *O) optionMutation[O] {
	return func(builder *O) *O {
		if value == nil {
			return builder
		}
		return setter(value)
	}
}

// applyOptions runs each mutation in order over the base builder.
func applyOptions[O any](base *O, mutations ...optionMutation[O]) *O {
	for _, mutate := range mutations {
		base = mutate(base)
	}
	return base
}

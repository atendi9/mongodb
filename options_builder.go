package mongodb

import (
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func BuildDatabaseOptions(
	opts *options.DatabaseOptions,
) options.Lister[options.DatabaseOptions] {
	dbOpts := options.Database()
	if opts.BSONOptions != nil {
		dbOpts = dbOpts.SetBSONOptions(opts.BSONOptions)
	}
	if opts.ReadConcern != nil {
		dbOpts = dbOpts.SetReadConcern(opts.ReadConcern)
	}
	if opts.ReadPreference != nil {
		dbOpts = dbOpts.SetReadPreference(opts.ReadPreference)
	}
	if opts.Registry != nil {
		dbOpts = dbOpts.SetRegistry(opts.Registry)
	}
	if opts.WriteConcern != nil {
		dbOpts = dbOpts.SetWriteConcern(opts.WriteConcern)
	}
	return dbOpts
}

func BuildFindManyOptions(
	opts ...*options.FindOptions,
) options.Lister[options.FindOptions] {
	findOpts := options.Find()
	if len(opts) > 0 {
		opts := opts[0]
		findOpts = setOption(findOpts, opts.AllowDiskUse, findOpts.SetAllowDiskUse)
		findOpts = setOption(findOpts, opts.AllowPartialResults, findOpts.SetAllowPartialResults)
		findOpts = setOption(findOpts, opts.BatchSize, findOpts.SetBatchSize)
		findOpts = setOption(findOpts, opts.Limit, findOpts.SetLimit)
		findOpts = setOption(findOpts, opts.Skip, findOpts.SetSkip)
		findOpts = setOption(findOpts, opts.MaxAwaitTime, findOpts.SetMaxAwaitTime)
		findOpts = setOption(findOpts, opts.NoCursorTimeout, findOpts.SetNoCursorTimeout)
		findOpts = setOption(findOpts, &opts.Sort, findOpts.SetSort)
		findOpts = setOption(findOpts, &opts.Comment, findOpts.SetComment)
		findOpts = setOption(findOpts, &opts.Hint, findOpts.SetHint)
		findOpts = setOption(findOpts, &opts.Let, findOpts.SetLet)
		findOpts = setOption(findOpts, &opts.Max, findOpts.SetMax)
		findOpts = setOption(findOpts, &opts.Min, findOpts.SetMin)
		findOpts = setOption(findOpts, &opts.Projection, findOpts.SetProjection)
		findOpts = setOption(findOpts, opts.ReturnKey, findOpts.SetReturnKey)
		if opts.Collation != nil {
			findOpts = findOpts.SetCollation(opts.Collation)
		}
		if opts.CursorType != nil {
			findOpts = findOpts.SetCursorType(*opts.CursorType)
		}
	}
	return findOpts
}

func BuildFindOneOptions(
	opts ...*options.FindOneOptions,
) options.Lister[options.FindOneOptions] {
	findOneOpts := options.FindOne()
	if len(opts) > 0 {
		opts := opts[0]
		findOneOpts = setOption(findOneOpts, opts.AllowPartialResults, findOneOpts.SetAllowPartialResults)
		findOneOpts = setOption(findOneOpts, opts.Skip, findOneOpts.SetSkip)
		findOneOpts = setOption(findOneOpts, &opts.Sort, findOneOpts.SetSort)
		findOneOpts = setOption(findOneOpts, &opts.Comment, findOneOpts.SetComment)
		findOneOpts = setOption(findOneOpts, &opts.Hint, findOneOpts.SetHint)
		findOneOpts = setOption(findOneOpts, &opts.Max, findOneOpts.SetMax)
		findOneOpts = setOption(findOneOpts, &opts.Min, findOneOpts.SetMin)
		findOneOpts = setOption(findOneOpts, &opts.Projection, findOneOpts.SetProjection)
		findOneOpts = setOption(findOneOpts, opts.ReturnKey, findOneOpts.SetReturnKey)
		if opts.Collation != nil {
			findOneOpts = findOneOpts.SetCollation(opts.Collation)
		}
	}
	return findOneOpts
}

func BuildUpdateOneOptions(
	opts ...*options.UpdateOneOptions,
) options.Lister[options.UpdateOneOptions] {
	updateOneOpts := options.UpdateOne()
	if len(opts) > 0 {
		opts := opts[0]
		updateOneOpts = setOption(updateOneOpts, &opts.ArrayFilters, updateOneOpts.SetArrayFilters)
		updateOneOpts = setOption(updateOneOpts, opts.BypassDocumentValidation, updateOneOpts.SetBypassDocumentValidation)
		updateOneOpts = setOption(updateOneOpts, &opts.Sort, updateOneOpts.SetSort)
		updateOneOpts = setOption(updateOneOpts, &opts.Comment, updateOneOpts.SetComment)
		updateOneOpts = setOption(updateOneOpts, &opts.Hint, updateOneOpts.SetHint)
		updateOneOpts = setOption(updateOneOpts, &opts.Let, updateOneOpts.SetLet)
		updateOneOpts = setOption(updateOneOpts, opts.Upsert, updateOneOpts.SetUpsert)
		if opts.Collation != nil {
			updateOneOpts = updateOneOpts.SetCollation(opts.Collation)
		}
	}

	return updateOneOpts
}

func BuildUpdateManyOptions(
	opts ...*options.UpdateManyOptions,
) options.Lister[options.UpdateManyOptions] {
	updateManyOpts := options.UpdateMany()
	if len(opts) > 0 {
		opts := opts[0]
		updateManyOpts = setOption(updateManyOpts, &opts.ArrayFilters, updateManyOpts.SetArrayFilters)
		updateManyOpts = setOption(updateManyOpts, opts.BypassDocumentValidation, updateManyOpts.SetBypassDocumentValidation)
		updateManyOpts = setOption(updateManyOpts, &opts.Comment, updateManyOpts.SetComment)
		updateManyOpts = setOption(updateManyOpts, &opts.Hint, updateManyOpts.SetHint)
		updateManyOpts = setOption(updateManyOpts, &opts.Let, updateManyOpts.SetLet)
		updateManyOpts = setOption(updateManyOpts, opts.Upsert, updateManyOpts.SetUpsert)
		if opts.Collation != nil {
			updateManyOpts = updateManyOpts.SetCollation(opts.Collation)
		}

	}
	return updateManyOpts
}

func BuildInsertManyOptions(
	opts ...*options.InsertManyOptions,
) options.Lister[options.InsertManyOptions] {
	insertManyOpts := options.InsertMany()
	if len(opts) > 0 {
		opts := opts[0]
		insertManyOpts = setOption(insertManyOpts, opts.BypassDocumentValidation, insertManyOpts.SetBypassDocumentValidation)
		insertManyOpts = setOption(insertManyOpts, opts.Ordered, insertManyOpts.SetOrdered)
		insertManyOpts = setOption(insertManyOpts, &opts.Comment, insertManyOpts.SetComment)
	}
	return insertManyOpts
}

func BuildCountOptions(
	opts ...*options.CountOptions,
) options.Lister[options.CountOptions] {
	countOpts := options.Count()
	if len(opts) > 0 {
		opts := opts[0]
		countOpts = setOption(countOpts, opts.Limit, countOpts.SetLimit)
		countOpts = setOption(countOpts, opts.Skip, countOpts.SetSkip)
		countOpts = setOption(countOpts, &opts.Comment, countOpts.SetComment)
		countOpts = setOption(countOpts, &opts.Hint, countOpts.SetHint)
		if opts.Collation != nil {
			countOpts = countOpts.SetCollation(opts.Collation)
		}
	}
	return countOpts
}

func BuildDistinctOptions(
	opts ...*options.DistinctOptions,
) options.Lister[options.DistinctOptions] {
	distinctOpts := options.Distinct()
	if len(opts) > 0 {
		opts := opts[0]
		distinctOpts = setOption(distinctOpts, &opts.Comment, distinctOpts.SetComment)
		distinctOpts = setOption(distinctOpts, &opts.Hint, distinctOpts.SetHint)
		if opts.Collation != nil {
			distinctOpts = distinctOpts.SetCollation(opts.Collation)
		}
	}
	return distinctOpts
}

func BuildFindOneAndUpdateOptions(
	opts ...*options.FindOneAndUpdateOptions,
) options.Lister[options.FindOneAndUpdateOptions] {
	findOneAndUpdateOpts := options.FindOneAndUpdate()
	if len(opts) > 0 {
		opts := opts[0]
		findOneAndUpdateOpts = setOption(findOneAndUpdateOpts, &opts.ArrayFilters, findOneAndUpdateOpts.SetArrayFilters)
		findOneAndUpdateOpts = setOption(findOneAndUpdateOpts, opts.BypassDocumentValidation, findOneAndUpdateOpts.SetBypassDocumentValidation)
		findOneAndUpdateOpts = setOption(findOneAndUpdateOpts, opts.ReturnDocument, findOneAndUpdateOpts.SetReturnDocument)
		findOneAndUpdateOpts = setOption(findOneAndUpdateOpts, opts.Upsert, findOneAndUpdateOpts.SetUpsert)
		findOneAndUpdateOpts = setOption(findOneAndUpdateOpts, &opts.Sort, findOneAndUpdateOpts.SetSort)
		findOneAndUpdateOpts = setOption(findOneAndUpdateOpts, &opts.Projection, findOneAndUpdateOpts.SetProjection)
		findOneAndUpdateOpts = setOption(findOneAndUpdateOpts, &opts.Comment, findOneAndUpdateOpts.SetComment)
		findOneAndUpdateOpts = setOption(findOneAndUpdateOpts, &opts.Hint, findOneAndUpdateOpts.SetHint)
		findOneAndUpdateOpts = setOption(findOneAndUpdateOpts, &opts.Let, findOneAndUpdateOpts.SetLet)
		if opts.Collation != nil {
			findOneAndUpdateOpts = findOneAndUpdateOpts.SetCollation(opts.Collation)
		}
	}
	return findOneAndUpdateOpts
}

func BuildFindOneAndDeleteOptions(
	opts ...*options.FindOneAndDeleteOptions,
) options.Lister[options.FindOneAndDeleteOptions] {
	findOneAndDeleteOpts := options.FindOneAndDelete()
	if len(opts) > 0 {
		opts := opts[0]
		findOneAndDeleteOpts = setOption(findOneAndDeleteOpts, &opts.Sort, findOneAndDeleteOpts.SetSort)
		findOneAndDeleteOpts = setOption(findOneAndDeleteOpts, &opts.Projection, findOneAndDeleteOpts.SetProjection)
		findOneAndDeleteOpts = setOption(findOneAndDeleteOpts, &opts.Comment, findOneAndDeleteOpts.SetComment)
		findOneAndDeleteOpts = setOption(findOneAndDeleteOpts, &opts.Hint, findOneAndDeleteOpts.SetHint)
		findOneAndDeleteOpts = setOption(findOneAndDeleteOpts, &opts.Let, findOneAndDeleteOpts.SetLet)
		if opts.Collation != nil {
			findOneAndDeleteOpts = findOneAndDeleteOpts.SetCollation(opts.Collation)
		}
	}
	return findOneAndDeleteOpts
}

func BuildFindOneAndReplaceOptions(
	opts ...*options.FindOneAndReplaceOptions,
) options.Lister[options.FindOneAndReplaceOptions] {
	findOneAndReplaceOpts := options.FindOneAndReplace()
	if len(opts) > 0 {
		opts := opts[0]
		findOneAndReplaceOpts = setOption(findOneAndReplaceOpts, opts.BypassDocumentValidation, findOneAndReplaceOpts.SetBypassDocumentValidation)
		findOneAndReplaceOpts = setOption(findOneAndReplaceOpts, opts.ReturnDocument, findOneAndReplaceOpts.SetReturnDocument)
		findOneAndReplaceOpts = setOption(findOneAndReplaceOpts, opts.Upsert, findOneAndReplaceOpts.SetUpsert)
		findOneAndReplaceOpts = setOption(findOneAndReplaceOpts, &opts.Sort, findOneAndReplaceOpts.SetSort)
		findOneAndReplaceOpts = setOption(findOneAndReplaceOpts, &opts.Projection, findOneAndReplaceOpts.SetProjection)
		findOneAndReplaceOpts = setOption(findOneAndReplaceOpts, &opts.Comment, findOneAndReplaceOpts.SetComment)
		findOneAndReplaceOpts = setOption(findOneAndReplaceOpts, &opts.Hint, findOneAndReplaceOpts.SetHint)
		findOneAndReplaceOpts = setOption(findOneAndReplaceOpts, &opts.Let, findOneAndReplaceOpts.SetLet)
		if opts.Collation != nil {
			findOneAndReplaceOpts = findOneAndReplaceOpts.SetCollation(opts.Collation)
		}
	}
	return findOneAndReplaceOpts
}

func setOption[O any, V any](
	builder *O,
	value *V,
	set func(V) *O,
) *O {
	if value != nil {
		return set(*value)
	}
	return builder
}

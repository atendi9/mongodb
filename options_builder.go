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

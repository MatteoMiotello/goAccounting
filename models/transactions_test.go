// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testTransactions(t *testing.T) {
	t.Parallel()

	query := Transactions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTransactionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Transactions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TransactionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TransactionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Transaction exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TransactionExists to return true, but got false.")
	}
}

func testTransactionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	transactionFound, err := FindTransaction(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if transactionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTransactionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Transactions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTransactionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Transactions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTransactionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionOne := &Transaction{}
	transactionTwo := &Transaction{}
	if err = randomize.Struct(seed, transactionOne, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}
	if err = randomize.Struct(seed, transactionTwo, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = transactionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = transactionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Transactions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTransactionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	transactionOne := &Transaction{}
	transactionTwo := &Transaction{}
	if err = randomize.Struct(seed, transactionOne, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}
	if err = randomize.Struct(seed, transactionTwo, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = transactionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = transactionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func transactionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Transaction) error {
	*o = Transaction{}
	return nil
}

func transactionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Transaction) error {
	*o = Transaction{}
	return nil
}

func transactionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Transaction) error {
	*o = Transaction{}
	return nil
}

func transactionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Transaction) error {
	*o = Transaction{}
	return nil
}

func transactionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Transaction) error {
	*o = Transaction{}
	return nil
}

func transactionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Transaction) error {
	*o = Transaction{}
	return nil
}

func transactionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Transaction) error {
	*o = Transaction{}
	return nil
}

func transactionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Transaction) error {
	*o = Transaction{}
	return nil
}

func transactionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Transaction) error {
	*o = Transaction{}
	return nil
}

func testTransactionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Transaction{}
	o := &Transaction{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, transactionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Transaction object: %s", err)
	}

	AddTransactionHook(boil.BeforeInsertHook, transactionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	transactionBeforeInsertHooks = []TransactionHook{}

	AddTransactionHook(boil.AfterInsertHook, transactionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	transactionAfterInsertHooks = []TransactionHook{}

	AddTransactionHook(boil.AfterSelectHook, transactionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	transactionAfterSelectHooks = []TransactionHook{}

	AddTransactionHook(boil.BeforeUpdateHook, transactionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	transactionBeforeUpdateHooks = []TransactionHook{}

	AddTransactionHook(boil.AfterUpdateHook, transactionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	transactionAfterUpdateHooks = []TransactionHook{}

	AddTransactionHook(boil.BeforeDeleteHook, transactionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	transactionBeforeDeleteHooks = []TransactionHook{}

	AddTransactionHook(boil.AfterDeleteHook, transactionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	transactionAfterDeleteHooks = []TransactionHook{}

	AddTransactionHook(boil.BeforeUpsertHook, transactionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	transactionBeforeUpsertHooks = []TransactionHook{}

	AddTransactionHook(boil.AfterUpsertHook, transactionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	transactionAfterUpsertHooks = []TransactionHook{}
}

func testTransactionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTransactionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(transactionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTransactionToOneAssetUsingAsset(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Transaction
	var foreign Asset

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, assetDBTypes, false, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.AssetID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Asset().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddAssetHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Asset) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := TransactionSlice{&local}
	if err = local.L.LoadAsset(ctx, tx, false, (*[]*Transaction)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Asset == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Asset = nil
	if err = local.L.LoadAsset(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Asset == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testTransactionToOneTransactionCategoryUsingTransactionCategory(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Transaction
	var foreign TransactionCategory

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, transactionCategoryDBTypes, false, transactionCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionCategory struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TransactionCategoryID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.TransactionCategory().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddTransactionCategoryHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *TransactionCategory) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := TransactionSlice{&local}
	if err = local.L.LoadTransactionCategory(ctx, tx, false, (*[]*Transaction)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.TransactionCategory == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TransactionCategory = nil
	if err = local.L.LoadTransactionCategory(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.TransactionCategory == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testTransactionToOneSetOpAssetUsingAsset(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Transaction
	var b, c Asset

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, assetDBTypes, false, strmangle.SetComplement(assetPrimaryKeyColumns, assetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, assetDBTypes, false, strmangle.SetComplement(assetPrimaryKeyColumns, assetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Asset{&b, &c} {
		err = a.SetAsset(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Asset != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Transactions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AssetID != x.ID {
			t.Error("foreign key was wrong value", a.AssetID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AssetID))
		reflect.Indirect(reflect.ValueOf(&a.AssetID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AssetID != x.ID {
			t.Error("foreign key was wrong value", a.AssetID, x.ID)
		}
	}
}
func testTransactionToOneSetOpTransactionCategoryUsingTransactionCategory(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Transaction
	var b, c TransactionCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, transactionCategoryDBTypes, false, strmangle.SetComplement(transactionCategoryPrimaryKeyColumns, transactionCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, transactionCategoryDBTypes, false, strmangle.SetComplement(transactionCategoryPrimaryKeyColumns, transactionCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*TransactionCategory{&b, &c} {
		err = a.SetTransactionCategory(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TransactionCategory != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Transactions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TransactionCategoryID != x.ID {
			t.Error("foreign key was wrong value", a.TransactionCategoryID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TransactionCategoryID))
		reflect.Indirect(reflect.ValueOf(&a.TransactionCategoryID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TransactionCategoryID != x.ID {
			t.Error("foreign key was wrong value", a.TransactionCategoryID, x.ID)
		}
	}
}

func testTransactionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTransactionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TransactionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTransactionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Transactions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	transactionDBTypes = map[string]string{`ID`: `bigint`, `AssetID`: `bigint`, `TransactionCategoryID`: `bigint`, `Amount`: `numeric`, `Description`: `character varying`, `DeletedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `CreatedAt`: `timestamp with time zone`}
	_                  = bytes.MinRead
)

func testTransactionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(transactionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(transactionAllColumns) == len(transactionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTransactionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(transactionAllColumns) == len(transactionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(transactionAllColumns, transactionPrimaryKeyColumns) {
		fields = transactionAllColumns
	} else {
		fields = strmangle.SetComplement(
			transactionAllColumns,
			transactionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := TransactionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTransactionsUpsert(t *testing.T) {
	t.Parallel()

	if len(transactionAllColumns) == len(transactionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Transaction{}
	if err = randomize.Struct(seed, &o, transactionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Transaction: %s", err)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, transactionDBTypes, false, transactionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Transaction: %s", err)
	}

	count, err = Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

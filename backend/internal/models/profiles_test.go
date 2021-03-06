// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testProfiles(t *testing.T) {
	t.Parallel()

	query := Profiles()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProfilesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
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

	count, err := Profiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProfilesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Profiles().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Profiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProfilesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProfileSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Profiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProfilesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProfileExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Profile exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProfileExists to return true, but got false.")
	}
}

func testProfilesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	profileFound, err := FindProfile(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if profileFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProfilesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Profiles().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testProfilesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Profiles().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProfilesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	profileOne := &Profile{}
	profileTwo := &Profile{}
	if err = randomize.Struct(seed, profileOne, profileDBTypes, false, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}
	if err = randomize.Struct(seed, profileTwo, profileDBTypes, false, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = profileOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = profileTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Profiles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProfilesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	profileOne := &Profile{}
	profileTwo := &Profile{}
	if err = randomize.Struct(seed, profileOne, profileDBTypes, false, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}
	if err = randomize.Struct(seed, profileTwo, profileDBTypes, false, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = profileOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = profileTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Profiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func profileBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Profile) error {
	*o = Profile{}
	return nil
}

func profileAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Profile) error {
	*o = Profile{}
	return nil
}

func profileAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Profile) error {
	*o = Profile{}
	return nil
}

func profileBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Profile) error {
	*o = Profile{}
	return nil
}

func profileAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Profile) error {
	*o = Profile{}
	return nil
}

func profileBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Profile) error {
	*o = Profile{}
	return nil
}

func profileAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Profile) error {
	*o = Profile{}
	return nil
}

func profileBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Profile) error {
	*o = Profile{}
	return nil
}

func profileAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Profile) error {
	*o = Profile{}
	return nil
}

func testProfilesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Profile{}
	o := &Profile{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, profileDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Profile object: %s", err)
	}

	AddProfileHook(boil.BeforeInsertHook, profileBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	profileBeforeInsertHooks = []ProfileHook{}

	AddProfileHook(boil.AfterInsertHook, profileAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	profileAfterInsertHooks = []ProfileHook{}

	AddProfileHook(boil.AfterSelectHook, profileAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	profileAfterSelectHooks = []ProfileHook{}

	AddProfileHook(boil.BeforeUpdateHook, profileBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	profileBeforeUpdateHooks = []ProfileHook{}

	AddProfileHook(boil.AfterUpdateHook, profileAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	profileAfterUpdateHooks = []ProfileHook{}

	AddProfileHook(boil.BeforeDeleteHook, profileBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	profileBeforeDeleteHooks = []ProfileHook{}

	AddProfileHook(boil.AfterDeleteHook, profileAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	profileAfterDeleteHooks = []ProfileHook{}

	AddProfileHook(boil.BeforeUpsertHook, profileBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	profileBeforeUpsertHooks = []ProfileHook{}

	AddProfileHook(boil.AfterUpsertHook, profileAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	profileAfterUpsertHooks = []ProfileHook{}
}

func testProfilesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Profiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProfilesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(profileColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Profiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProfileToManyUsers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ProfileID = a.ID
	c.ProfileID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Users().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ProfileID == b.ProfileID {
			bFound = true
		}
		if v.ProfileID == c.ProfileID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ProfileSlice{&a}
	if err = a.L.LoadUsers(ctx, tx, false, (*[]*Profile)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Users = nil
	if err = a.L.LoadUsers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testProfileToManyAddOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, false, strmangle.SetComplement(profilePrimaryKeyColumns, profileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*User{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUsers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ProfileID {
			t.Error("foreign key was wrong value", a.ID, first.ProfileID)
		}
		if a.ID != second.ProfileID {
			t.Error("foreign key was wrong value", a.ID, second.ProfileID)
		}

		if first.R.Profile != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Profile != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Users[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Users[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Users().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testProfilesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
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

func testProfilesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProfileSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProfilesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Profiles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	profileDBTypes = map[string]string{`ID`: `integer`, `Names`: `character varying`, `TagLine`: `text`, `DateOfBirth`: `date`, `ProfilePhoto`: `text`}
	_              = bytes.MinRead
)

func testProfilesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(profilePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(profileAllColumns) == len(profilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Profiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, profileDBTypes, true, profilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProfilesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(profileAllColumns) == len(profilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Profiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, profileDBTypes, true, profilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(profileAllColumns, profilePrimaryKeyColumns) {
		fields = profileAllColumns
	} else {
		fields = strmangle.SetComplement(
			profileAllColumns,
			profilePrimaryKeyColumns,
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

	slice := ProfileSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testProfilesUpsert(t *testing.T) {
	t.Parallel()

	if len(profileAllColumns) == len(profilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Profile{}
	if err = randomize.Struct(seed, &o, profileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Profile: %s", err)
	}

	count, err := Profiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, profileDBTypes, false, profilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Profile: %s", err)
	}

	count, err = Profiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

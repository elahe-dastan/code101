// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testPassengers(t *testing.T) {
	t.Parallel()

	query := Passengers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPassengersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Passenger{}
	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
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

	count, err := Passengers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPassengersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Passenger{}
	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Passengers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Passengers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPassengersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Passenger{}
	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PassengerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Passengers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPassengersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Passenger{}
	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PassengerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Passenger exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PassengerExists to return true, but got false.")
	}
}

func testPassengersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Passenger{}
	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	passengerFound, err := FindPassenger(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if passengerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPassengersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Passenger{}
	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Passengers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPassengersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Passenger{}
	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Passengers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPassengersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	passengerOne := &Passenger{}
	passengerTwo := &Passenger{}
	if err = randomize.Struct(seed, passengerOne, passengerDBTypes, false, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}
	if err = randomize.Struct(seed, passengerTwo, passengerDBTypes, false, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = passengerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = passengerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Passengers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPassengersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	passengerOne := &Passenger{}
	passengerTwo := &Passenger{}
	if err = randomize.Struct(seed, passengerOne, passengerDBTypes, false, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}
	if err = randomize.Struct(seed, passengerTwo, passengerDBTypes, false, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = passengerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = passengerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Passengers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func passengerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Passenger) error {
	*o = Passenger{}
	return nil
}

func passengerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Passenger) error {
	*o = Passenger{}
	return nil
}

func passengerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Passenger) error {
	*o = Passenger{}
	return nil
}

func passengerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Passenger) error {
	*o = Passenger{}
	return nil
}

func passengerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Passenger) error {
	*o = Passenger{}
	return nil
}

func passengerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Passenger) error {
	*o = Passenger{}
	return nil
}

func passengerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Passenger) error {
	*o = Passenger{}
	return nil
}

func passengerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Passenger) error {
	*o = Passenger{}
	return nil
}

func passengerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Passenger) error {
	*o = Passenger{}
	return nil
}

func testPassengersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Passenger{}
	o := &Passenger{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, passengerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Passenger object: %s", err)
	}

	AddPassengerHook(boil.BeforeInsertHook, passengerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	passengerBeforeInsertHooks = []PassengerHook{}

	AddPassengerHook(boil.AfterInsertHook, passengerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	passengerAfterInsertHooks = []PassengerHook{}

	AddPassengerHook(boil.AfterSelectHook, passengerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	passengerAfterSelectHooks = []PassengerHook{}

	AddPassengerHook(boil.BeforeUpdateHook, passengerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	passengerBeforeUpdateHooks = []PassengerHook{}

	AddPassengerHook(boil.AfterUpdateHook, passengerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	passengerAfterUpdateHooks = []PassengerHook{}

	AddPassengerHook(boil.BeforeDeleteHook, passengerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	passengerBeforeDeleteHooks = []PassengerHook{}

	AddPassengerHook(boil.AfterDeleteHook, passengerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	passengerAfterDeleteHooks = []PassengerHook{}

	AddPassengerHook(boil.BeforeUpsertHook, passengerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	passengerBeforeUpsertHooks = []PassengerHook{}

	AddPassengerHook(boil.AfterUpsertHook, passengerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	passengerAfterUpsertHooks = []PassengerHook{}
}

func testPassengersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Passenger{}
	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Passengers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPassengersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Passenger{}
	if err = randomize.Struct(seed, o, passengerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(passengerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Passengers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPassengerToManyRides(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Passenger
	var b, c Ride

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, passengerDBTypes, true, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, rideDBTypes, false, rideColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, rideDBTypes, false, rideColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.PassengerID, a.ID)
	queries.Assign(&c.PassengerID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Rides().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.PassengerID, b.PassengerID) {
			bFound = true
		}
		if queries.Equal(v.PassengerID, c.PassengerID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PassengerSlice{&a}
	if err = a.L.LoadRides(ctx, tx, false, (*[]*Passenger)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Rides); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Rides = nil
	if err = a.L.LoadRides(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Rides); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPassengerToManyAddOpRides(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Passenger
	var b, c, d, e Ride

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, passengerDBTypes, false, strmangle.SetComplement(passengerPrimaryKeyColumns, passengerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Ride{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, rideDBTypes, false, strmangle.SetComplement(ridePrimaryKeyColumns, rideColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Ride{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRides(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.PassengerID) {
			t.Error("foreign key was wrong value", a.ID, first.PassengerID)
		}
		if !queries.Equal(a.ID, second.PassengerID) {
			t.Error("foreign key was wrong value", a.ID, second.PassengerID)
		}

		if first.R.Passenger != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Passenger != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Rides[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Rides[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Rides().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testPassengerToManySetOpRides(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Passenger
	var b, c, d, e Ride

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, passengerDBTypes, false, strmangle.SetComplement(passengerPrimaryKeyColumns, passengerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Ride{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, rideDBTypes, false, strmangle.SetComplement(ridePrimaryKeyColumns, rideColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetRides(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Rides().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetRides(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Rides().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.PassengerID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.PassengerID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.PassengerID) {
		t.Error("foreign key was wrong value", a.ID, d.PassengerID)
	}
	if !queries.Equal(a.ID, e.PassengerID) {
		t.Error("foreign key was wrong value", a.ID, e.PassengerID)
	}

	if b.R.Passenger != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Passenger != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Passenger != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Passenger != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Rides[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Rides[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testPassengerToManyRemoveOpRides(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Passenger
	var b, c, d, e Ride

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, passengerDBTypes, false, strmangle.SetComplement(passengerPrimaryKeyColumns, passengerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Ride{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, rideDBTypes, false, strmangle.SetComplement(ridePrimaryKeyColumns, rideColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddRides(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Rides().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveRides(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Rides().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.PassengerID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.PassengerID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Passenger != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Passenger != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Passenger != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Passenger != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Rides) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Rides[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Rides[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testPassengersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Passenger{}
	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
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

func testPassengersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Passenger{}
	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PassengerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPassengersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Passenger{}
	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Passengers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	passengerDBTypes = map[string]string{`ID`: `int`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `Email`: `varchar`, `Password`: `varchar`, `Fullname`: `varchar`, `Phone`: `varchar`, `Cellphone`: `varchar`, `IsCellphoneVerified`: `tinyint`, `PhotoURL`: `varchar`, `IsEmailVerified`: `tinyint`, `Locale`: `varchar`, `ReferralCode`: `varchar`, `RegisterationIP`: `varchar`, `IsRegisteredWithGoogle`: `tinyint`, `EmailVerificationCode`: `varchar`, `CellphoneVerificationCode`: `varchar`, `IsBlocked`: `tinyint`, `AdjustFingerprint`: `varchar`, `ComapnyID`: `int`}
	_                = bytes.MinRead
)

func testPassengersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(passengerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(passengerAllColumns) == len(passengerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Passenger{}
	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Passengers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPassengersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(passengerAllColumns) == len(passengerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Passenger{}
	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Passengers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, passengerDBTypes, true, passengerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(passengerAllColumns, passengerPrimaryKeyColumns) {
		fields = passengerAllColumns
	} else {
		fields = strmangle.SetComplement(
			passengerAllColumns,
			passengerPrimaryKeyColumns,
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

	slice := PassengerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPassengersUpsert(t *testing.T) {
	t.Parallel()

	if len(passengerAllColumns) == len(passengerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLPassengerUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Passenger{}
	if err = randomize.Struct(seed, &o, passengerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Passenger: %s", err)
	}

	count, err := Passengers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, passengerDBTypes, false, passengerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Passenger struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Passenger: %s", err)
	}

	count, err = Passengers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

package assertions

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShouldEqual(t *testing.T) {
	fail(t, so(1, ShouldEqual), "This assertion requires exactly 1 comparison values (you provided 0).")
	fail(t, so(1, ShouldEqual, 1, 2), "This assertion requires exactly 1 comparison values (you provided 2).")
	fail(t, so(1, ShouldEqual, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	pass(t, so(1, ShouldEqual, 1))
	fail(t, so(1, ShouldEqual, 2), "Expected '1'\nto equal '2'\n(but it didn't)!")

	pass(t, so(true, ShouldEqual, true))
	fail(t, so(true, ShouldEqual, false), "Expected 'true'\nto equal 'false'\n(but it didn't)!")

	pass(t, so("hi", ShouldEqual, "hi"))
	fail(t, so("hi", ShouldEqual, "bye"), "Expected 'hi'\nto equal 'bye'\n(but it didn't)!")

	pass(t, so(42, ShouldEqual, uint(42)))

	fail(t, so(Thing1{}, ShouldEqual, Thing1{}), "Expected '{}'\nto equal '{}'\n(but it didn't)!")
	fail(t, so(Thing1{"hi"}, ShouldEqual, Thing1{"hi"}), "Expected '{hi}'\nto equal '{hi}'\n(but it didn't)!")
	fail(t, so(&Thing1{"hi"}, ShouldEqual, &Thing1{"hi"}), "Expected '&{hi}'\nto equal '&{hi}'\n(but it didn't)!")

	fail(t, so(Thing1{}, ShouldEqual, Thing2{}), "Expected '{}'\nto equal '{}'\n(but it didn't)!")
}

func TestShouldNotEqual(t *testing.T) {
	fail(t, so(1, ShouldNotEqual), "This assertion requires exactly 1 comparison values (you provided 0).")
	fail(t, so(1, ShouldNotEqual, 1, 2), "This assertion requires exactly 1 comparison values (you provided 2).")
	fail(t, so(1, ShouldNotEqual, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	pass(t, so(1, ShouldNotEqual, 2))
	fail(t, so(1, ShouldNotEqual, 1), "Expected     '1'\nto NOT equal '1'\n(but it did)!")

	pass(t, so(true, ShouldNotEqual, false))
	fail(t, so(true, ShouldNotEqual, true), "Expected     'true'\nto NOT equal 'true'\n(but it did)!")

	pass(t, so("hi", ShouldNotEqual, "bye"))
	fail(t, so("hi", ShouldNotEqual, "hi"), "Expected     'hi'\nto NOT equal 'hi'\n(but it did)!")

	pass(t, so(&Thing1{"hi"}, ShouldNotEqual, &Thing1{"hi"}))
	pass(t, so(Thing1{"hi"}, ShouldNotEqual, Thing1{"hi"}))
	pass(t, so(Thing1{}, ShouldNotEqual, Thing1{}))
	pass(t, so(Thing1{}, ShouldNotEqual, Thing2{}))
}

func TestShouldResemble(t *testing.T) {
	fail(t, so(Thing1{"hi"}, ShouldResemble), "This assertion requires exactly 1 comparison values (you provided 0).")
	fail(t, so(Thing1{"hi"}, ShouldResemble, Thing1{"hi"}, Thing1{"hi"}), "This assertion requires exactly 1 comparison values (you provided 2).")

	pass(t, so(Thing1{"hi"}, ShouldResemble, Thing1{"hi"}))
	fail(t, so(Thing1{"hi"}, ShouldResemble, Thing1{"bye"}), "Expected '{hi}'\nto resemble '{bye}'\n(but it didn't)!")
}

func TestShouldNotResemble(t *testing.T) {
	fail(t, so(Thing1{"hi"}, ShouldNotResemble), "This assertion requires exactly 1 comparison values (you provided 0).")
	fail(t, so(Thing1{"hi"}, ShouldNotResemble, Thing1{"hi"}, Thing1{"hi"}), "This assertion requires exactly 1 comparison values (you provided 2).")

	pass(t, so(Thing1{"hi"}, ShouldNotResemble, Thing1{"bye"}))
	fail(t, so(Thing1{"hi"}, ShouldNotResemble, Thing1{"hi"}), "Expected        '{hi}'\nto NOT resemble '{hi}'\n(but it did)!")
}

func TestShouldPointTo(t *testing.T) {
	t1 := &Thing1{}
	t2 := t1
	t3 := &Thing1{}

	pointer1 := reflect.ValueOf(t1).Pointer()
	pointer3 := reflect.ValueOf(t3).Pointer()

	fail(t, so(t1, ShouldPointTo), "This assertion requires exactly 1 comparison values (you provided 0).")
	fail(t, so(t1, ShouldPointTo, t2, t3), "This assertion requires exactly 1 comparison values (you provided 2).")

	pass(t, so(t1, ShouldPointTo, t2))
	fail(t, so(t1, ShouldPointTo, t3), fmt.Sprintf("Expected '&{}' (address: '%v') and '&{}' (address: '%v') to be the same address (but their weren't)!", pointer1, pointer3))

	t4 := Thing1{}
	t5 := t4

	fail(t, so(t4, ShouldPointTo, t5), "Both arguments should be pointers (the first was not)!")
	fail(t, so(&t4, ShouldPointTo, t5), "Both arguments should be pointers (the second was not)!")
	fail(t, so(nil, ShouldPointTo, nil), "Both arguments should be pointers (the first was nil)!")
	fail(t, so(&t4, ShouldPointTo, nil), "Both arguments should be pointers (the second was nil)!")
}

func TestShouldNotPointTo(t *testing.T) {
	t1 := &Thing1{}
	t2 := t1
	t3 := &Thing1{}

	pointer1 := reflect.ValueOf(t1).Pointer()

	fail(t, so(t1, ShouldNotPointTo), "This assertion requires exactly 1 comparison values (you provided 0).")
	fail(t, so(t1, ShouldNotPointTo, t2, t3), "This assertion requires exactly 1 comparison values (you provided 2).")

	pass(t, so(t1, ShouldNotPointTo, t3))
	fail(t, so(t1, ShouldNotPointTo, t2), fmt.Sprintf("Expected '&{}' and '&{}' to be different references (but they matched: '%v')!", pointer1))

	t4 := Thing1{}
	t5 := t4

	fail(t, so(t4, ShouldNotPointTo, t5), "Both arguments should be pointers (the first was not)!")
	fail(t, so(&t4, ShouldNotPointTo, t5), "Both arguments should be pointers (the second was not)!")
	fail(t, so(nil, ShouldNotPointTo, nil), "Both arguments should be pointers (the first was nil)!")
	fail(t, so(&t4, ShouldNotPointTo, nil), "Both arguments should be pointers (the second was nil)!")
}

func TestShouldBeNil(t *testing.T) {
	fail(t, so(nil, ShouldBeNil, nil, nil, nil), "This assertion requires exactly 0 comparison values (you provided 3).")
	fail(t, so(nil, ShouldBeNil, nil), "This assertion requires exactly 0 comparison values (you provided 1).")

	pass(t, so(nil, ShouldBeNil))
	fail(t, so(1, ShouldBeNil), "Expected '1' to be nil (but it wasn't)!")

	var thing Thinger
	pass(t, so(thing, ShouldBeNil))
	thing = &Thing{}
	fail(t, so(thing, ShouldBeNil), "Expected '&{}' to be nil (but it wasn't)!")

	var thingOne *Thing1
	pass(t, so(thingOne, ShouldBeNil))
}

func TestShouldNotBeNil(t *testing.T) {
	fail(t, so(nil, ShouldNotBeNil, nil, nil, nil), "This assertion requires exactly 0 comparison values (you provided 3).")
	fail(t, so(nil, ShouldNotBeNil, nil), "This assertion requires exactly 0 comparison values (you provided 1).")

	fail(t, so(nil, ShouldNotBeNil), "Expected '<nil>' to NOT be nil (but it was)!")
	pass(t, so(1, ShouldNotBeNil))

	var thing Thinger
	fail(t, so(thing, ShouldNotBeNil), "Expected '<nil>' to NOT be nil (but it was)!")
	thing = &Thing{}
	pass(t, so(thing, ShouldNotBeNil))
}

func TestShouldBeTrue(t *testing.T) {
	fail(t, so(true, ShouldBeTrue, 1, 2, 3), "This assertion requires exactly 0 comparison values (you provided 3).")
	fail(t, so(true, ShouldBeTrue, 1), "This assertion requires exactly 0 comparison values (you provided 1).")

	fail(t, so(false, ShouldBeTrue), "Expected 'true' (not 'false')!")
	fail(t, so(1, ShouldBeTrue), "Expected 'true' (not '1')!")
	pass(t, so(true, ShouldBeTrue))
}

func TestShouldBeFalse(t *testing.T) {
	fail(t, so(false, ShouldBeFalse, 1, 2, 3), "This assertion requires exactly 0 comparison values (you provided 3).")
	fail(t, so(false, ShouldBeFalse, 1), "This assertion requires exactly 0 comparison values (you provided 1).")

	fail(t, so(true, ShouldBeFalse), "Expected 'false' (not 'true')!")
	fail(t, so(1, ShouldBeFalse), "Expected 'false' (not '1')!")
	pass(t, so(false, ShouldBeFalse))
}

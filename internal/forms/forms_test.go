package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	postedData := url.Values{}

	form := New(postedData)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("shows valid when required fields missing")
	}

	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")
	form = New(postedData)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}

	form := New(postedData)
	if form.Has("a") {
		t.Error("shows field exists and filled when required field missing")
	}

	postedData.Add("a", "")
	form = New(postedData)
	if form.Has("a") {
		t.Error("shows field exists and filled when required field is empty")
	}

	postedData.Set("a", "a")
	form = New(postedData)
	if !form.Has("a") {
		t.Error("shows field doesn't exists or filled when required field is filled")
	}
}

func TestForm_MinLenght(t *testing.T) {
	postedData := url.Values{}

	form := New(postedData)
	if form.MinLenght("a", 3) {
		t.Error("shows field has a minimum lenght when field doesn't exist")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedData.Add("a", "a")
	form = New(postedData)
	if form.MinLenght("a", 3) {
		t.Error("shows field has a minimum lenght when it doesn't")
	}

	postedData.Add("b", "abc")
	form = New(postedData)
	if !form.MinLenght("b", 3) {
		t.Error("shows field has no a minimum lenght when it does")
	}

	isError = form.Errors.Get("b")
	if isError != "" {
		t.Error("should not have an error, but did get one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}

	form := New(postedData)
	if form.IsEmail("email") {
		t.Error("shows email field has right format when email field doesn't exist")
	}

	postedData.Add("email", "a")
	form = New(postedData)
	if form.IsEmail("email") {
		t.Error("shows email field has right format when it doesn't")
	}

	postedData.Set("email", "me@here.com")
	form = New(postedData)
	if !form.IsEmail("email") {
		t.Error("shows email field has wrong format when it doesn't")
	}
}

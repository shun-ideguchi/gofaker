package gofaker

import (
	"fmt"

	"github.com/shun-ideguchi/gofaker/data"
)

// LastName はランダムな日本の苗字を返します。
func (f *Faker) LastName() string {
	return data.Person["lastName"][f.Rand.Intn(len(data.Person["lastName"]))]
}

// FirstName はランダムな日本の名前を返します。
func (f *Faker) FirstName() string {
	return data.Person["firstName"][f.Rand.Intn(len(data.Person["firstName"]))]
}

// FullName はランダムな日本の苗字と名前を返します。
func (f *Faker) FullName() string {
	return f.LastName() + " " + f.FirstName()
}

// LastNameKana はランダムな日本の苗字のフリガナを返します。
func (f *Faker) LastNameKana() string {
	return data.LastNameKana[f.LastName()]
}

// FirstNameKana はランダムな日本の名前のフリガナを返します。
func (f *Faker) FirstNameKana() string {
	return data.FirstNameKana[f.FirstName()]
}

// LastNameKanaBy は指定された苗字のフリガナを返します。
func (f *Faker) LastNameKanaBy(lastName string) string {
	if kana, ok := data.LastNameKana[lastName]; ok {
		return kana
	}

	return f.LastNameKana()
}

// FirstNameKanaBy は指定された名前のフリガナを返します。
func (f *Faker) FirstNameKanaBy(firstName string) string {
	if kana, ok := data.FirstNameKana[firstName]; ok {
		return kana
	}

	return f.FirstNameKana()
}

// FullNameKana はランダムな日本の苗字と名前のフリガナを返します。
func (f *Faker) FullNameKana() string {
	return f.LastNameKana() + " " + f.FirstNameKana()
}

// MobilePhoneNumber はランダムな日本の携帯電話番号を返します。
func (f *Faker) MobilePhoneNumber() string {
	prefix := data.Person["mobilePrefix"][f.Rand.Intn(len(data.Person["mobilePrefix"]))]
	mid := f.Rand.Intn(10000)
	end := f.Rand.Intn(10000)

	return fmt.Sprintf("%s%03d%04d", prefix, mid, end)
}

func (f *Faker) Hobby() string {
	return data.Person["hobby"][f.Rand.Intn(len(data.Person["hobby"]))]
}

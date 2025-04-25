package gofaker

import "fmt"

func ExamplePerson() {
	faker := New(1)
	fmt.Println(faker.LastName())
	fmt.Println(faker.FirstName())
	fmt.Println(faker.FullName())
	fmt.Println(faker.LastNameKana())
	fmt.Println(faker.FirstNameKana())
	fmt.Println(faker.LastNameKanaBy("鈴木"))
	fmt.Println(faker.FirstNameKanaBy("結衣"))
	fmt.Println(faker.FullNameKana())
	fmt.Println(faker.MobilePhoneNumber())
	fmt.Println(faker.Hobby())

	// Output:
	// 鈴木
	// 結衣
	// 中村 美月
	// スズキ
	// ハルト
	// スズキ
	// ユイ
	// ワタナベ タロウ
	// 09033000694
	// 運動
}

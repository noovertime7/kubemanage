package utils

import "testing"

var (
	kb1 = `aaaaaaa`
	kb2 = `bbbbbbb`
	kb3 = `ccccccc`
)

func TestEncrypt(t *testing.T) {
	cases := []struct {
		Name string
		text []byte
	}{
		{"a", []byte(kb1)},
		{"b", []byte(kb2)},
		{"c", []byte(kb3)},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans, err := Encrypt(c.text); err != nil {
				t.Fatalf("encrypt text %s failed: %+v",
					c.text, err)
			} else {
				t.Logf("encrypt text %s is { %s }", c.text, ans)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	cases := []struct {
		Name string
		text string
	}{
		{"a", "Obx1VwUPs7B09CqalouHQg=="},
		{"b", "Zol2IPDQuGTo/K0IYDkkAQ=="},
		{"c", "nmW+Ha3epblxZmgVvcvaSQ=="},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans, err := Decrypt(c.text); err != nil {
				t.Fatalf("decrypt text %s failed: %+v",
					c.text, err)
			} else {
				t.Logf("decrypt text %s is %s", c.text, ans)
			}
		})
	}
}

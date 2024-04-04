package sample

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/muhammedarifp/grpc-example/pb"
)

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(2) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomGpuBrand() string {
	brand := selectRandomString("Nvidia", "Amd")
	return brand
}

func RandomCpuBrand() string {
	brand := selectRandomString("Intel", "Ryzen")
	return brand
}

func selectRandomString(brands ...string) string {
	n := len(brands)
	if n == 0 {
		return ""
	}

	return brands[rand.Intn(n)]
}

func randomeCpuName(brand string) string {
	if brand == "Intel" {
		return selectRandomString(
			"Core i9-9900K",
			"Core i7-8700K",
			"Core i5-9600K",
			"Core i3-9100",
			"Pentium G5400",
		)
	}

	return selectRandomString(
		"Ryzen 7 3700X",
		"Ryzen 5 3600",
		"Ryzen 3 3300X",
		"Ryzen 3 3100",
		"Athlon 3000G",
	)
}

func randomeGpuName(brand string) string {
	if brand == "Nvidia" {
		return selectRandomString(
			"RTX 2068",
			"RTX 7989",
			"RTX 9990",
			"RTX 77w6",
		)
	}

	return selectRandomString(
		"RX 67",
		"RX 990",
		"RX 887",
		"RX 88",
	)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomScreenPanel() pb.Screen_Panel {
	switch rand.Intn(2) {
	case 1:
		return pb.Screen_LCD
	case 2:
		return pb.Screen_OLED
	default:
		return pb.Screen_UNKNOWN
	}
}

//////

func randomId() string {
	id := uuid.New().String()
	return id
}

func randomLaptopBrand() string {
	return selectRandomString(
		"Apple",
		"Lenove",
		"Azuze",
	)
}

func randomLaptopName(brand string) string {
	if brand == "Apple" {
		return selectRandomString(
			"Macbook 1",
			"Macbook 2",
			"Macbook 3",
		)
	} else if brand == "Lenove" {
		return selectRandomString(
			"Lenove Ideapad 1",
			"Lenove Ideapad 2",
			"Lenove Ideapad 3",
			"Lenove Ideapad 4",
		)
	} else {
		return selectRandomString(
			"Azuze Tuf 1",
			"Azuze Tuf 2",
			"Azuze Tuf 3",
			"Azuze Tuf 4",
			"Azuze Tuf 7",
		)
	}
}

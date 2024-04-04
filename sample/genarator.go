package sample

import (
	"github.com/muhammedarifp/grpc-example/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewKeyboard creates a new Keyboard with random values
func NewKeybord() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

// NewGpu creates a new GPU with random values
func NewCpu() *pb.CPU {
	brand := RandomCpuBrand()
	name := randomeCpuName(brand)
	cores := int32(randomInt(2, 8))
	threads := int32(randomInt(int(cores), 16))
	minghz := randomFloat(1.5, 5.0)
	maxghz := randomFloat(minghz, 5.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   int32(cores),
		NumberThreads: int32(threads),
		MinGhz:        minghz,
		MaxGhz:        maxghz,
	}

	return cpu
}

// NewGpu creates a new GPU with random values
func NewGpu() *pb.GPU {
	brand := randomGpuBrand()
	name := randomeGpuName(brand)
	minghz := randomFloat(1.0, 1.5)
	maxghz := randomFloat(minghz, 2.0)
	memory := &pb.Memory{
		Value: int64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}
	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minghz,
		MaxGhz: maxghz,
		Memory: memory,
	}

	return gpu
}

// New Ram

func NewRam() *pb.Memory {
	ram := &pb.Memory{
		Value: int64(randomInt(2, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}

	return ram
}

// ssd
func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: int64(randomInt(128, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	return ssd
}

// hdd
func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: int64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}

	return hdd
}

// new screen
func NewScreen() *pb.Screen {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	screen := &pb.Screen{
		SizeInch: float32(randomFloat(13, 17)),
		Resolution: &pb.Screen_Resolution{
			Width:  int32(width),
			Height: int32(height),
		},
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

// new Laptop feild
func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &pb.Laptop{
		Id:          randomId(),
		Brand:       brand,
		Name:        name,
		Cpu:         NewCpu(),
		Memory:      NewRam(),
		Gpus:        []*pb.GPU{NewGpu()},
		Storages:    []*pb.Storage{NewSSD(), NewHDD()},
		Screen:      NewScreen(),
		Keyboard:    NewKeybord(),
		Weight:      &pb.Laptop_WeightKg{WeightKg: float32(randomFloat(1300, 3000))},
		PriceUsd:    randomFloat(500, 3000),
		ReleaseYear: uint32(randomInt(2000, 2024)),
		UpdatedAt:   timestamppb.Now(),
	}

	return laptop
}

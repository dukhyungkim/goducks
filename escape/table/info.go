package table

type dropInfo struct {
	Name   string
	Count  int
	Weight float32
}

var (
	squirrelDropInfo = []dropInfo{
		{
			Name:   "회복약",
			Count:  1,
			Weight: 70,
		},
		{
			Name:   "",
			Weight: 30,
		},
	}

	rabbitDropInfo = []dropInfo{
		{
			Name:   "회복약",
			Count:  1,
			Weight: 50,
		},
		{
			Name:   "회복약",
			Count:  2,
			Weight: 30,
		},
		{
			Name:   "",
			Weight: 20,
		},
	}

	deerDropInfo = []dropInfo{
		{
			Name:   "열쇠",
			Count:  1,
			Weight: 100,
		},
	}

	boxDropInfo = []dropInfo{
		{
			Name:   "목검",
			Count:  1,
			Weight: 20,
		},
		{
			Name:   "철검",
			Count:  1,
			Weight: 15,
		},
		{
			Name:   "가죽옷",
			Count:  1,
			Weight: 7,
		},
		{
			Name:   "가죽바지",
			Count:  1,
			Weight: 8,
		},
		{
			Name:   "가죽신발",
			Count:  1,
			Weight: 10,
		},
		{
			Name:   "회복약",
			Count:  1,
			Weight: 15,
		},
		{
			Name:   "회복약",
			Count:  2,
			Weight: 10,
		},
		{
			Name:   "회복약",
			Count:  3,
			Weight: 5,
		},
		{
			Name:   "",
			Weight: 10,
		},
	}
)

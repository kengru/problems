package main

import (
	"fmt"

	d202201 "github.com/kengru/problems/advent/2022/01"
	d202202 "github.com/kengru/problems/advent/2022/02"
	d202203 "github.com/kengru/problems/advent/2022/03"
	d202204 "github.com/kengru/problems/advent/2022/04"
	d202205 "github.com/kengru/problems/advent/2022/05"
	d202206 "github.com/kengru/problems/advent/2022/06"
	d202207 "github.com/kengru/problems/advent/2022/07"
	d202208 "github.com/kengru/problems/advent/2022/08"
	d202209 "github.com/kengru/problems/advent/2022/09"
	d202210 "github.com/kengru/problems/advent/2022/10"
	d202211 "github.com/kengru/problems/advent/2022/11"

	d202501 "github.com/kengru/problems/advent/2025/01"
	d202502 "github.com/kengru/problems/advent/2025/02"
	d202503 "github.com/kengru/problems/advent/2025/03"
	d202504 "github.com/kengru/problems/advent/2025/04"
	d202505 "github.com/kengru/problems/advent/2025/05"
)

func RunAdventProblem(year int, day int, part int) error {
	key := fmt.Sprintf("%04d%02d%02d", year, day, part)
	fn, exists := solutionRegistry[key]
	if !exists {
		return fmt.Errorf("solution not found for year %d, day %d, part %d", year, day, part)
	}
	fn()
	return nil
}

var solutionRegistry = map[string]func(){
	// 2022
	"20220101": d202201.Year20220101,
	"20220102": d202201.Year20220102,
	"20220201": d202202.Year20220201,
	"20220202": d202202.Year20220202,
	"20220301": d202203.Year20220301,
	"20220302": d202203.Year20220302,
	"20220401": d202204.Year20220401,
	"20220402": d202204.Year20220402,
	"20220501": d202205.Year20220501,
	"20220502": d202205.Year20220502,
	"20220601": d202206.Year20220601,
	"20220602": d202206.Year20220602,
	"20220701": d202207.Year20220701,
	"20220702": d202207.Year20220702,
	"20220801": d202208.Year20220801,
	"20220802": d202208.Year20220802,
	"20220901": d202209.Year20220901,
	"20220902": d202209.Year20220902,
	"20221001": d202210.Year20221001,
	"20221002": d202210.Year20221002,
	"20221101": d202211.Year20221101,
	"20221102": d202211.Year20221102,

	// 2025
	"20250101": d202501.Year20250101,
	"20250102": d202501.Year20250102,
	"20250201": d202502.Year20250201,
	"20250202": d202502.Year20250202,
	"20250301": d202503.Year20250301,
	"20250302": d202503.Year20250302,
	"20250401": d202504.Year20250401,
	"20250402": d202504.Year20250402,
	"20250501": d202505.Year20250501,
	"20250502": d202505.Year20250502,
}

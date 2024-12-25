package main

import (
	"fmt"
	"time"

	"github.com/dtykrnt/go-prayer"
)

func main() {
	// Calculate prayer schedule in Jakarta for 2024.
	asiaJakarta, _ := time.LoadLocation("Asia/Jakarta")
	jakartaSchedules, _ := prayer.Calculate(prayer.Config{
		Latitude:           -6.14,
		Longitude:          106.81,
		Timezone:           asiaJakarta,
		TwilightConvention: prayer.Kemenag(),
		AsrConvention:      prayer.Shafii,
		PreciseToSeconds:   true,
	}, 2024, 12, 25)
	print(jakartaSchedules)
	// for i := 0; i < len(jakartaSchedules); i++ {
	// 	fmt.Println("schedule day ", i, jakartaSchedules[i].Fajr.Format("15:04"))
	// }

	// Calculate prayer schedule in London for 2024.
	// Since London in higher latitude, make sure to enable the adapter.
	// europeLondon, _ := time.LoadLocation("Europe/London")
	// londonSchedules, _ := prayer.Calculate(prayer.Config{
	// 	Latitude:            51.507222,
	// 	Longitude:           -0.1275,
	// 	Timezone:            europeLondon,
	// 	TwilightConvention:  prayer.ISNA(),
	// 	AsrConvention:       prayer.Shafii,
	// 	HighLatitudeAdapter: prayer.NearestLatitude(),
	// 	PreciseToSeconds:    true,
	// }, 2024, 2)
	// print(londonSchedules)
}

func print(schedules []prayer.Schedule) {
	for _, s := range schedules {
		fmt.Println(
			"'"+s.Date+"'",
			s.Fajr.Format("'2006-01-02 15:04:05'"),
			s.Sunrise.Format("'2006-01-02 15:04:05'"),
			s.Zuhr.Format("'2006-01-02 15:04:05'"),
			s.Asr.Format("'2006-01-02 15:04:05'"),
			s.Maghrib.Format("'2006-01-02 15:04:05'"),
			s.Isha.Format("'2006-01-02 15:04:05'"),
		)
	}
}

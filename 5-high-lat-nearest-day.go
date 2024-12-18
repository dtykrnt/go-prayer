package prayer

// NearestDay is adapter where the schedule for "abnormal" days will be taken from the
// schedule of the last "normal" day.
//
// This adapter doesn't require the sunrise and sunset to be exist in a day, so it's
// usable for area in extreme latitudes (>=65 degrees).
//
// Reference: https://www.islamicity.com/prayertimes/Salat.pdf
func NearestDay() HighLatitudeAdapter {
	return highLatNearestDay
}

func highLatNearestDay(_ Config, _ int, _ int, schedules []Schedule) []Schedule {
	abnormalSummer, abnormalWinter := extractAbnormalSchedules(schedules)

	for _, as := range []abnormalRange{abnormalSummer, abnormalWinter} {
		// If this abnormal period is empty, skip
		if as.IsEmpty() {
			continue
		}

		// Get the last normal schedule
		abnormalIdxStart := as.Indexes[0]
		lastNormalSchedule := sliceAt(schedules, abnormalIdxStart-1)

		// Use the last normal schedule for the entire abnormal period
		for _, idx := range as.Indexes {
			schedules[idx] = lastNormalSchedule
		}
	}

	return schedules
}

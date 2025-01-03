package prayer

import (
	"time"
)

// NearestLatitude is adapter where the schedules will be estimated using percentage
// of schedule in location at 45 degrees latitude. This method will change the schedule
// for entire year to prevent sudden changes in fasting time.
//
// This adapter only estimates time for Isha and Fajr and require sunrise and sunset
// time. Therefore it's not suitable for area in extreme latitude (>=65 degrees).
//
// Reference: https://fiqh.islamonline.net/en/praying-and-fasting-at-high-latitudes/
func NearestLatitude() HighLatitudeAdapter {
	return highLatNearestLatitude
}

func highLatNearestLatitude(cfg Config, year int, month int, day int, schedules []Schedule) []Schedule {
	// This conventions only works if daytime exists (in other words, sunrise
	// and Maghrib must exist). So if there are days where those time don't
	// exist, stop and just return the schedule as it is.
	// TODO: maybe put some warning log later.
	for _, s := range schedules {
		if s.Sunrise.IsZero() || s.Maghrib.IsZero() {
			return schedules
		}
	}

	// Get the nearest latitude
	latitude := cfg.Latitude
	if latitude > 45 {
		latitude = 45
	} else if latitude < -45 {
		latitude = -45
	}

	// Calculate schedule for the nearest latitude
	newCfg := Config{
		Latitude:           latitude,
		Longitude:          cfg.Longitude,
		Timezone:           cfg.Timezone,
		TwilightConvention: cfg.TwilightConvention,
		AsrConvention:      cfg.AsrConvention}
	nearestSchedules, _ := calcNormal(newCfg, year, month, day)

	for i := range schedules {
		// Calculate duration from schedule of nearest latitude
		ns := nearestSchedules[i]
		nsDay := ns.Maghrib.Sub(ns.Sunrise).Seconds()
		nsFajrRise := ns.Sunrise.Sub(ns.Fajr).Seconds()
		nsMaghribIsha := ns.Isha.Sub(ns.Maghrib).Seconds()

		nsNight := 24*60*60 - nsDay
		nsFajrPercentage := nsFajrRise / nsNight
		nsIshaPercentage := nsMaghribIsha / nsNight

		// Calculate duration from current schedule
		s := schedules[i]
		sDay := s.Maghrib.Sub(s.Sunrise).Seconds()
		sNight := 24*60*60 - sDay

		// Apply the new durations
		fajrDuration := time.Duration(sNight * nsFajrPercentage * float64(time.Second))
		ishaDuration := time.Duration(sNight * nsIshaPercentage * float64(time.Second))
		s.Fajr = s.Sunrise.Add(-fajrDuration)
		s.Isha = s.Maghrib.Add(ishaDuration)
		schedules[i] = s
	}

	return schedules
}

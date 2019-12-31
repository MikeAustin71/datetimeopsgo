package datetime

import "sync"

// tzAbbrvToTimeZonePriorityList - Internal list
// used to assign Time Zones to Time Zone abbreviations
// based on priority

var lockTzAbbrvToTimeZonePriorityList   sync.Mutex

var tzAbbrvToTimeZonePriorityList = []string {
			"UTC",
			"Etc/UTC",
			"Etc/GMT-0",
			"America/New_York",
			"America/Chicago",
			"America/Denver",
			"America/Los_Angeles",
			"Pacific/Honolulu",
			"America/Anchorage",
			"America/Adak",
			"America/Havana",
			"America/Godthab",
			"America/St_Johns",
			"America/Thule",
			"America/Buenos_Aires",
			"America/Caracas",
			"America/Bogota",
			"Atlantic/Azores",
			"EST5EDT",
			"CST6CDT",
			"MST7MDT",
			"PST8PDT",
			"America",
			"US",
			"Asia/Riyadh",
			"Europe/Paris",
			"Europe/London",
			"Europe/Dublin",
			"Europe/Rome",
			"Europe/Madrid",
			"Europe/Kiev",
			"Europe/Minsk",
			"Europe/Volgograd",
			"Europe/Moscow",
			"Europe",
			"Pacific/Fiji",
			"Asia/Shanghai",
			"Asia/Hong_Kong",
			"Asia/Seoul",
			"Asia/Tokyo",
			"Asia/Calcutta",
			"Asia/Karachi",
			"Asia/Manila",
			"Asia/Jerusalem",
			"Asia/Tel_Aviv",
			"Asia/Jakarta",
			"Asia/Makassar",
			"Asia/Kabul",
			"Asia/Tashkent",
			"Asia/Kathmandu",
			"Asia/Dhaka",
			"Asia/Yangon",
			"Asia/Ho_Chi_Minh",
			"Asia/Singapore",
			"Asia/Yakutsk",
			"Asia/Vladivostok",
			"Asia/Sakhalin",
			"Asia",
			"Atlantic/Canary",
			"Atlantic/Cape_Verde",
			"Atlantic",
			"Australia/Sydney",
			"Australia/Darwin",
			"Australia/Melbourne",
			"Australia/Adelaide",
			"Australia/Perth",
			"Australia/Lord_Howe",
			"Australia",
			"NZ-CHAT",  // New Zealand
			"Canada",
			"Pacific/Easter",
			"Pacific/Pitcairn",
			"Pacific/Gambier",
			"Pacific/Tahiti",
			"Pacific/Guam",
			"Pacific/Samoa",
			"Pacific/Apia",
			"Pacific",
			"Africa/Cairo",
			"Africa/Johannesburg",
			"Africa/Casablanca",
			"Africa/Nairobi",
			"Africa/Lagos",
			"Africa",
			"Indian",
			"Antarctica/McMurdo",
			"Antarctica",
			"Etc",
			"Other",
}


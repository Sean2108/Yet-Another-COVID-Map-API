package casecount

import (
	"strings"
	"testing"
)

func getTestCaseCounts() map[string]map[string]CaseCounts {
	result := map[string]map[string]CaseCounts{
		"China": map[string]CaseCounts{
			"Beijing": CaseCounts{
				Location{40.1824, 116.4142},
				[]CaseCount{
					CaseCount{"1/22/20", statistics{50, 10}},
					CaseCount{"1/23/20", statistics{200, 87}},
					CaseCount{"1/24/20", statistics{800, 125}},
					CaseCount{"1/25/20", statistics{1020, 142}},
					CaseCount{"1/26/20", statistics{1110, 145}},
					CaseCount{"1/27/20", statistics{1235, 152}},
				},
			},
			"Hubei": CaseCounts{
				Location{30.9756, 112.2707},
				[]CaseCount{
					CaseCount{"1/22/20", statistics{100, 20}},
					CaseCount{"1/23/20", statistics{1000, 100}},
					CaseCount{"1/24/20", statistics{1800, 105}},
					CaseCount{"1/25/20", statistics{2020, 150}},
					CaseCount{"1/26/20", statistics{2110, 175}},
					CaseCount{"1/27/20", statistics{2111, 230}},
				},
			},
			"Shanghai": CaseCounts{
				Location{31.202, 121.4491},
				[]CaseCount{
					CaseCount{"1/22/20", statistics{10, 5}},
					CaseCount{"1/23/20", statistics{45, 8}},
					CaseCount{"1/24/20", statistics{89, 20}},
					CaseCount{"1/25/20", statistics{126, 25}},
					CaseCount{"1/26/20", statistics{400, 42}},
					CaseCount{"1/27/20", statistics{532, 55}},
				},
			},
		},
		"Singapore": map[string]CaseCounts{
			"": CaseCounts{
				Location{1.2833, 103.8333},
				[]CaseCount{
					CaseCount{"1/22/20", statistics{1, 0}},
					CaseCount{"1/23/20", statistics{3, 2}},
					CaseCount{"1/24/20", statistics{6, 4}},
					CaseCount{"1/25/20", statistics{10, 5}},
					CaseCount{"1/26/20", statistics{15, 8}},
					CaseCount{"1/27/20", statistics{23, 10}},
				},
			},
		},
		"United Kingdom": map[string]CaseCounts{
			"London": CaseCounts{
				Location{55.3781, -3.4360000000000004},
				[]CaseCount{
					CaseCount{"1/22/20", statistics{1, 0}},
					CaseCount{"1/23/20", statistics{6, 1}},
					CaseCount{"1/24/20", statistics{8, 3}},
					CaseCount{"1/25/20", statistics{9, 6}},
					CaseCount{"1/26/20", statistics{20, 6}},
					CaseCount{"1/27/20", statistics{28, 9}},
				},
			},
		},
	}
	return result
}

func getTestCaseCountsWithoutFirstAndLastDay() map[string]map[string]CaseCounts {
	result := map[string]map[string]CaseCounts{
		"China": map[string]CaseCounts{
			"Beijing": CaseCounts{
				Location{40.1824, 116.4142},
				[]CaseCount{
					CaseCount{"1/23/20", statistics{200, 87}},
					CaseCount{"1/24/20", statistics{800, 125}},
					CaseCount{"1/25/20", statistics{1020, 142}},
					CaseCount{"1/26/20", statistics{1110, 145}},
				},
			},
			"Hubei": CaseCounts{
				Location{30.9756, 112.2707},
				[]CaseCount{
					CaseCount{"1/23/20", statistics{1000, 100}},
					CaseCount{"1/24/20", statistics{1800, 105}},
					CaseCount{"1/25/20", statistics{2020, 150}},
					CaseCount{"1/26/20", statistics{2110, 175}},
				},
			},
			"Shanghai": CaseCounts{
				Location{31.202, 121.4491},
				[]CaseCount{
					CaseCount{"1/23/20", statistics{45, 8}},
					CaseCount{"1/24/20", statistics{89, 20}},
					CaseCount{"1/25/20", statistics{126, 25}},
					CaseCount{"1/26/20", statistics{400, 42}},
				},
			},
		},
		"Singapore": map[string]CaseCounts{
			"": CaseCounts{
				Location{1.2833, 103.8333},
				[]CaseCount{
					CaseCount{"1/23/20", statistics{3, 2}},
					CaseCount{"1/24/20", statistics{6, 4}},
					CaseCount{"1/25/20", statistics{10, 5}},
					CaseCount{"1/26/20", statistics{15, 8}},
				},
			},
		},
		"United Kingdom": map[string]CaseCounts{
			"London": CaseCounts{
				Location{55.3781, -3.4360000000000004},
				[]CaseCount{
					CaseCount{"1/23/20", statistics{6, 1}},
					CaseCount{"1/24/20", statistics{8, 3}},
					CaseCount{"1/25/20", statistics{9, 6}},
					CaseCount{"1/26/20", statistics{20, 6}},
				},
			},
		},
	}
	return result
}

func TestExtractCaseCounts(t *testing.T) {
	confirmedData := [][]string{
		{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"},
		{"Beijing", "China", "40.1824", "116.4142", "50", "200", "800", "1020", "1110", "1235"},
		{"Hubei", "China", "30.9756", "112.2707", "100", "1000", "1800", "2020", "2110", "2111"},
		{"Shanghai", "China", "31.202", "121.4491", "10", "45", "89", "126", "400", "532"},
		{"", "Singapore", "1.2833", "103.8333", "1", "3", "6", "10", "15", "23"},
		{"London", "United Kingdom", "55.3781", "-3.4360000000000004", "1", "6", "8", "9", "20", "28"},
	}
	deathsData := [][]string{
		{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"},
		{"Beijing", "China", "40.1824", "116.4142", "10", "87", "125", "142", "145", "152"},
		{"Hubei", "China", "30.9756", "112.2707", "20", "100", "105", "150", "175", "230"},
		{"Shanghai", "China", "31.202", "121.4491", "5", "8", "20", "25", "42", "55"},
		{"", "Singapore", "1.2833", "103.8333", "0", "2", "4", "5", "8", "10"},
		{"London", "United Kingdom", "55.3781", "-3.4360000000000004", "0", "1", "3", "6", "6", "9"},
	}
	headerRow := confirmedData[0]
	result := extractCaseCounts(headerRow, confirmedData, deathsData)
	expectedData := getTestCaseCounts()
	verifyResultsCaseCountsMap(result, expectedData, t)
}

func TestAggregateDataBetweenDates_AllDates(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result := stateAggregatedMap
	expectedData := map[string]map[string]CaseCountsAggregated{
		"China": map[string]CaseCountsAggregated{
			"Beijing": CaseCountsAggregated{
				Location{40.1824, 116.4142},
				statistics{1235, 152},
			},
			"Hubei": CaseCountsAggregated{
				Location{30.9756, 112.2707},
				statistics{2111, 230},
			},
			"Shanghai": CaseCountsAggregated{
				Location{31.202, 121.4491},
				statistics{532, 55},
			},
		},
		"Singapore": map[string]CaseCountsAggregated{
			"": CaseCountsAggregated{
				Location{1.2833, 103.8333},
				statistics{23, 10},
			},
		},
		"United Kingdom": map[string]CaseCountsAggregated{
			"London": CaseCountsAggregated{
				Location{55.3781, -3.4360000000000004},
				statistics{28, 9},
			},
		},
	}
	verifyResultsCaseCountsAgg(result, expectedData, t)
}

func TestAggregateDataBetweenDates_QueryDates(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := aggregateDataBetweenDates("1/24/20", "1/26/20", "")
	expectedData := map[string]map[string]CaseCountsAggregated{
		"China": map[string]CaseCountsAggregated{
			"Beijing": CaseCountsAggregated{
				Location{40.1824, 116.4142},
				statistics{910, 58},
			},
			"Hubei": CaseCountsAggregated{
				Location{30.9756, 112.2707},
				statistics{1110, 75},
			},
			"Shanghai": CaseCountsAggregated{
				Location{31.202, 121.4491},
				statistics{355, 34},
			},
		},
		"Singapore": map[string]CaseCountsAggregated{
			"": CaseCountsAggregated{
				Location{1.2833, 103.8333},
				statistics{12, 6},
			},
		},
		"United Kingdom": map[string]CaseCountsAggregated{
			"London": CaseCountsAggregated{
				Location{55.3781, -3.4360000000000004},
				statistics{14, 5},
			},
		},
	}
	verifyResultsCaseCountsAgg(result, expectedData, t)
}

func TestAggregateDataBetweenDates_QueryDatesBeforeValidRange(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := aggregateDataBetweenDates("1/20/20", "1/21/20", "")
	if len(result) != 0 {
		t.Errorf("Length of results is incorrect, got: %d, want %d.", len(result), 0)
	}
}

func TestAggregateDataBetweenDates_QueryDatesAfterValidRange(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := aggregateDataBetweenDates("1/28/20", "1/29/20", "")
	if len(result) != 0 {
		t.Errorf("Length of results is incorrect, got: %d, want %d.", len(result), 0)
	}
}

func TestAggregateDataBetweenDates_QueryDatesBeforeAndAfter_ShouldReturnAll(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := aggregateDataBetweenDates("1/21/20", "1/28/20", "")
	expectedData := map[string]map[string]CaseCountsAggregated{
		"China": map[string]CaseCountsAggregated{
			"Beijing": CaseCountsAggregated{
				Location{40.1824, 116.4142},
				statistics{1235, 152},
			},
			"Hubei": CaseCountsAggregated{
				Location{30.9756, 112.2707},
				statistics{2111, 230},
			},
			"Shanghai": CaseCountsAggregated{
				Location{31.202, 121.4491},
				statistics{532, 55},
			},
		},
		"Singapore": map[string]CaseCountsAggregated{
			"": CaseCountsAggregated{
				Location{1.2833, 103.8333},
				statistics{23, 10},
			},
		},
		"United Kingdom": map[string]CaseCountsAggregated{
			"London": CaseCountsAggregated{
				Location{55.3781, -3.4360000000000004},
				statistics{28, 9},
			},
		},
	}
	verifyResultsCaseCountsAgg(result, expectedData, t)
}

func TestAggregateDataBetweenDates_QueryFromDateAfterToDate(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	_, err := aggregateDataBetweenDates("1/24/20", "1/23/20", "China")
	if err == nil {
		t.Error("Error message should be returned.")
	}
}

func TestAggregateDataBetweenDates_QueryCountry(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := aggregateDataBetweenDates("", "", "Singapore")
	expectedData := map[string]map[string]CaseCountsAggregated{
		"Singapore": map[string]CaseCountsAggregated{
			"": CaseCountsAggregated{
				Location{1.2833, 103.8333},
				statistics{23, 10},
			},
		},
	}
	verifyResultsCaseCountsAgg(result, expectedData, t)
}

func TestAggregateDataBetweenDates_QueryCountry_wEiRdCaSiNg(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := aggregateDataBetweenDates("", "", "sInGaPoRe")
	expectedData := map[string]map[string]CaseCountsAggregated{
		"Singapore": map[string]CaseCountsAggregated{
			"": CaseCountsAggregated{
				Location{1.2833, 103.8333},
				statistics{23, 10},
			},
		},
	}
	verifyResultsCaseCountsAgg(result, expectedData, t)
}

func TestAggregateDataBetweenDates_QueryCountry_SpellingMistakeInCountryName(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, err := aggregateDataBetweenDates("", "", "Siingapore")
	if len(result) != 0 {
		t.Errorf("Length of results is incorrect, got: %d, want %d.", len(result), 0)
	}
	if !strings.Contains(err.Error(), "Singapore") {
		t.Errorf("Error message is incorrect, got: %s, want %s.", err.Error(), "string containing Singapore")
	}
	result, err = aggregateDataBetweenDates("", "", "chain")
	if len(result) != 0 {
		t.Errorf("Length of results is incorrect, got: %d, want %d.", len(result), 0)
	}
	if !strings.Contains(err.Error(), "China") {
		t.Errorf("Error message is incorrect, got: %s, want %s.", err.Error(), "string containing China")
	}
	result, err = aggregateDataBetweenDates("", "", "UnitedKingdom")
	if len(result) != 0 {
		t.Errorf("Length of results is incorrect, got: %d, want %d.", len(result), 0)
	}
	if !strings.Contains(err.Error(), "United Kingdom") {
		t.Errorf("Error message is incorrect, got: %s, want %s.", err.Error(), "string containing United Kingdom")
	}
}

func TestAggregateDataBetweenDates_QueryDates_FromIsOutOfRange(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := aggregateDataBetweenDates("1/21/20", "1/26/20", "")
	expectedData := map[string]map[string]CaseCountsAggregated{
		"China": map[string]CaseCountsAggregated{
			"Beijing": CaseCountsAggregated{
				Location{40.1824, 116.4142},
				statistics{1110, 145},
			},
			"Hubei": CaseCountsAggregated{
				Location{30.9756, 112.2707},
				statistics{2110, 175},
			},
			"Shanghai": CaseCountsAggregated{
				Location{31.202, 121.4491},
				statistics{400, 42},
			},
		},
		"Singapore": map[string]CaseCountsAggregated{
			"": CaseCountsAggregated{
				Location{1.2833, 103.8333},
				statistics{15, 8},
			},
		},
		"United Kingdom": map[string]CaseCountsAggregated{
			"London": CaseCountsAggregated{
				Location{55.3781, -3.4360000000000004},
				statistics{20, 6},
			},
		},
	}
	verifyResultsCaseCountsAgg(result, expectedData, t)
}

func TestAggregateDataBetweenDates_QueryDates_ToIsOutOfRange(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := aggregateDataBetweenDates("1/24/20", "1/28/20", "")
	expectedData := map[string]map[string]CaseCountsAggregated{
		"China": map[string]CaseCountsAggregated{
			"Beijing": CaseCountsAggregated{
				Location{40.1824, 116.4142},
				statistics{1035, 65},
			},
			"Hubei": CaseCountsAggregated{
				Location{30.9756, 112.2707},
				statistics{1111, 130},
			},
			"Shanghai": CaseCountsAggregated{
				Location{31.202, 121.4491},
				statistics{487, 47},
			},
		},
		"Singapore": map[string]CaseCountsAggregated{
			"": CaseCountsAggregated{
				Location{1.2833, 103.8333},
				statistics{20, 8},
			},
		},
		"United Kingdom": map[string]CaseCountsAggregated{
			"London": CaseCountsAggregated{
				Location{55.3781, -3.4360000000000004},
				statistics{22, 8},
			},
		},
	}
	verifyResultsCaseCountsAgg(result, expectedData, t)
}

func TestAggregateDataBetweenDates_QueryDates_FromAndToBothOutOfRange(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := aggregateDataBetweenDates("1/21/20", "1/28/20", "")
	expectedData := map[string]map[string]CaseCountsAggregated{
		"China": map[string]CaseCountsAggregated{
			"Beijing": CaseCountsAggregated{
				Location{40.1824, 116.4142},
				statistics{1235, 152},
			},
			"Hubei": CaseCountsAggregated{
				Location{30.9756, 112.2707},
				statistics{2111, 230},
			},
			"Shanghai": CaseCountsAggregated{
				Location{31.202, 121.4491},
				statistics{532, 55},
			},
		},
		"Singapore": map[string]CaseCountsAggregated{
			"": CaseCountsAggregated{
				Location{1.2833, 103.8333},
				statistics{23, 10},
			},
		},
		"United Kingdom": map[string]CaseCountsAggregated{
			"London": CaseCountsAggregated{
				Location{55.3781, -3.4360000000000004},
				statistics{28, 9},
			},
		},
	}
	verifyResultsCaseCountsAgg(result, expectedData, t)
}

func TestAggregateCountryDataFromStatesAggregate_AllDates(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result := countryAggregatedMap
	expectedData := map[string]CaseCountsAggregated{
		"China": CaseCountsAggregated{
			Location{(40.1824 + 30.9756 + 31.202) / 3.0, (116.4142 + 112.2707 + 121.4491) / 3.0},
			statistics{3878, 437},
		},
		"Singapore": CaseCountsAggregated{
			Location{1.2833, 103.8333},
			statistics{23, 10},
		},
		"United Kingdom": CaseCountsAggregated{
			Location{55.3781, -3.4360000000000004},
			statistics{28, 9},
		},
	}
	verifyResultsCountryCaseCountsAgg(result, expectedData, t)
}

func TestAggregateCountryDataFromStatesAggregate_QueryDates(t *testing.T) {
	input := map[string]map[string]CaseCountsAggregated{
		"China": map[string]CaseCountsAggregated{
			"Beijing": CaseCountsAggregated{
				Location{40.1824, 116.4142},
				statistics{910, 58},
			},
			"Hubei": CaseCountsAggregated{
				Location{30.9756, 112.2707},
				statistics{1110, 75},
			},
			"Shanghai": CaseCountsAggregated{
				Location{31.202, 121.4491},
				statistics{355, 34},
			},
		},
		"Singapore": map[string]CaseCountsAggregated{
			"": CaseCountsAggregated{
				Location{1.2833, 103.8333},
				statistics{12, 6},
			},
		},
		"United Kingdom": map[string]CaseCountsAggregated{
			"London": CaseCountsAggregated{
				Location{55.3781, -3.4360000000000004},
				statistics{14, 5},
			},
		},
	}
	result := aggregateCountryDataFromStatesAggregate(input)
	expectedData := map[string]CaseCountsAggregated{
		"China": CaseCountsAggregated{
			Location{(40.1824 + 30.9756 + 31.202) / 3.0, (116.4142 + 112.2707 + 121.4491) / 3.0},
			statistics{2375, 167},
		},
		"Singapore": CaseCountsAggregated{
			Location{1.2833, 103.8333},
			statistics{12, 6},
		},
		"United Kingdom": CaseCountsAggregated{
			Location{55.3781, -3.4360000000000004},
			statistics{14, 5},
		},
	}
	verifyResultsCountryCaseCountsAgg(result, expectedData, t)
}

func TestAggregateDataPerDay_AllDates(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := GetCaseCountsWithDayData("", "", "")
	expectedData := caseCountsMap
	verifyResultsCaseCountsMap(result, expectedData, t)
}

func TestAggregateDataPerDay_QueryDates(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := GetCaseCountsWithDayData("1/23/20", "1/26/20", "")
	expectedData := getTestCaseCountsWithoutFirstAndLastDay()
	verifyResultsCaseCountsMap(result, expectedData, t)
}

func TestAggregateDataPerDay_BeforeAndAfterShouldReturnAll(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := GetCaseCountsWithDayData("1/21/20", "1/28/20", "")
	expectedData := caseCountsMap
	verifyResultsCaseCountsMap(result, expectedData, t)
}

func TestAggregateDataPerDay_CountryQuery(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := GetCaseCountsWithDayData("", "", "China")
	expectedData := getTestCaseCounts()["China"]
	verifyResultsCaseCountsMap(result, map[string]map[string]CaseCounts{"China": expectedData}, t)
}

func TestAggregateDataPerDay_CountryQueryWithSpellingError(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, err := GetCaseCountsWithDayData("", "", "Chiina")
	if len(result) != 0 {
		t.Errorf("Length of results is incorrect, got: %d, want %d.", len(result), 0)
	}
	if !strings.Contains(err.Error(), "China") {
		t.Errorf("Error message is incorrect, got: %s, want %s.", err.Error(), "string containing China")
	}
}

func TestAggregateDataPerDay_QueryFromDateAfterToDate(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	_, err := GetCaseCountsWithDayData("1/24/20", "1/23/20", "China")
	if err == nil {
		t.Error("Error message should be returned.")
	}
}

func TestCountryAggregateDataPerDay_AllDates(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := GetCountryCaseCountsWithDayData("", "", "")
	expectedData := countryCaseCountsMap
	verifyResultsCountryCaseCountsMap(result, expectedData, t)
}

func TestCountryAggregateDataPerDay_QueryDates(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := GetCountryCaseCountsWithDayData("1/23/20", "1/26/20", "")
	expectedData := map[string]CaseCounts{
		"China": CaseCounts{
			Location{(40.1824 + 30.9756 + 31.202) / 3.0, (116.4142 + 112.2707 + 121.4491) / 3.0},
			[]CaseCount{
				CaseCount{"1/23/20", statistics{1245, 195}},
				CaseCount{"1/24/20", statistics{2689, 250}},
				CaseCount{"1/25/20", statistics{3166, 317}},
				CaseCount{"1/26/20", statistics{3620, 362}},
			},
		},
		"Singapore": CaseCounts{
			Location{1.2833, 103.8333},
			[]CaseCount{
				CaseCount{"1/23/20", statistics{3, 2}},
				CaseCount{"1/24/20", statistics{6, 4}},
				CaseCount{"1/25/20", statistics{10, 5}},
				CaseCount{"1/26/20", statistics{15, 8}},
			},
		},
		"United Kingdom": CaseCounts{
			Location{55.3781, -3.4360000000000004},
			[]CaseCount{
				CaseCount{"1/23/20", statistics{6, 1}},
				CaseCount{"1/24/20", statistics{8, 3}},
				CaseCount{"1/25/20", statistics{9, 6}},
				CaseCount{"1/26/20", statistics{20, 6}},
			},
		},
	}
	verifyResultsCountryCaseCountsMap(result, expectedData, t)
}

func TestWorldTotal_AllDates(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := GetWorldCaseCounts("", "")
	if len(result) != 6 {
		t.Errorf("Length of results is incorrect, got: %d, want %d.", len(result), 3)
	}
	expectedData := []CaseCount{
		CaseCount{"1/22/20", statistics{162, 35}},
		CaseCount{"1/23/20", statistics{1254, 198}},
		CaseCount{"1/24/20", statistics{2703, 257}},
		CaseCount{"1/25/20", statistics{3185, 328}},
		CaseCount{"1/26/20", statistics{3655, 376}},
		CaseCount{"1/27/20", statistics{3929, 456}},
	}
	verifyResultsCaseCountArr(result, expectedData, t)
}

func TestWorldTotal_QueryDates(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	result, _ := GetWorldCaseCounts("1/23/20", "1/26/20")
	if len(result) != 4 {
		t.Errorf("Length of results is incorrect, got: %d, want %d.", len(result), 3)
	}
	expectedData := []CaseCount{
		CaseCount{"1/23/20", statistics{1254, 198}},
		CaseCount{"1/24/20", statistics{2703, 257}},
		CaseCount{"1/25/20", statistics{3185, 328}},
		CaseCount{"1/26/20", statistics{3655, 376}},
	}
	verifyResultsCaseCountArr(result, expectedData, t)
}

func TestWorldTotal_QueryFromDateAfterToDate(t *testing.T) {
	caseCountsMap = getTestCaseCounts()
	setDateBoundariesAndAllAggregatedData([]string{"Province/State", "Country/Region", "Lat", "Long", "1/22/20", "1/23/20", "1/24/20", "1/25/20", "1/26/20", "1/27/20"})
	_, err := GetWorldCaseCounts("1/24/20", "1/23/20")
	if err == nil {
		t.Error("Error message should be returned.")
	}
}
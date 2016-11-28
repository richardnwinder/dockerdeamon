function country(code)
{
	return countrydata[code].name;
}
function language(code)
{
	return countrydata[code].language;
}
var countrydata =
{
		"DE": {
		"name": "Germany",
		"iso2": "DE",
		"code": "33",
		"language": "ge"
		},
		"FR": {
		"name": "France",
		"iso2": "FR",
		"code": "33",
		"language": "fr"
		},
		"UK": {
		"name": "United Kingdom",
		"iso2": "UK",
		"code": "44",
		"language": "en"
		}
};

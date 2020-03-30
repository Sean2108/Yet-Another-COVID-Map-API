# Yet-Another-COVID-Map-API
API written in Go to provide number of COVID cases and deaths with date queries, and news stories related to COVID for countries.
Data is provided by John Hopkins CSSE (https://github.com/CSSEGISandData/COVID-19).

## Endpoints:
/cases:
- Call the endpoint with no query information to get the numbers of all confirmed cases and deaths for each state and country.
- Call the endpoint with attributes 'from' and/or 'to' to get the numbers of all confirmed cases and deaths for each state and country between the from date and to date. These dates should be in the format M/D/YY, for example /cases?from=1/2/20&to=1/10/20. Please do not pad the date with zeroes if the date or month has a single digit.
- WIP: Call the endpoint with attribute 'aggregateCountries' set to true to aggregate the counts to the country level instead of the state level. For example, /cases?aggregateCountries=true
- WIP: Call the endpoint with country name in the field 'country' to extract the numbers of confirmed cases and deaths for all states in the country. For example, /cases?country=Singapore

/news:
- WIP: Get news for country in the field to extract the latest coronavirus news for that country. Will use the News API (https://newsapi.org/) for obtaining this information.
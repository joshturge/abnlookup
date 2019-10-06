# ABN Lookup Wrapper for Go

This wrapper attempts to provide a layer of abstraction from the ABN Lookup API. 

All the ABN Lookup methods have been added although some have yet to be unit tested, check the table below for unit tested methods. This wrapper also provides functions that can validate both ABN and ACN numbers. These functions are exported and are used by [SearchByABN](https://github.com/joshturge/abnlookup/blob/c0bb6920aeed213819a4bf890b7449edd94df82b/request.go) and [SearchByACN](https://github.com/joshturge/abnlookup/blob/c0bb6920aeed213819a4bf890b7449edd94df82b/request.go) to validate before sending a request to the ABN Lookup API.

All the following ABN Lookup endpoints have been implemented.
- [SearchByABNv201408](https://abr.business.gov.au/abrxmlsearch/Forms/SearchByABNv201408.aspx)
- [SearchByASICv201408](https://abr.business.gov.au/abrxmlsearch/Forms/SearchByASICv201408.aspx)
- [ABRSearchByNameAdvancedSimpleProtocol2017](https://abr.business.gov.au/abrxmlsearch/Forms/ABRSearchByNameAdvancedSimpleProtocol2017.aspx)
- [SearchByABNStatus](https://abr.business.gov.au/abrxmlsearch/Forms/SearchByABNStatus.aspx)
- [SearchByCharity](https://abr.business.gov.au/abrxmlsearch/Forms/SearchByCharity.aspx)
- [SearchByPostcode](https://abr.business.gov.au/abrxmlsearch/Forms/SearchByPostcode.aspx)
- [SearchByRegistrationEvent](https://abr.business.gov.au/abrxmlsearch/Forms/SearchByRegistrationEvent.aspx)
- [SearchByUpdateEvent](https://abr.business.gov.au/abrxmlsearch/Forms/SearchByUpdateEvent.aspx)

## Usage
To use this packages client, you need to register for a GUID (more info [here](https://api.gov.au/service/5b639f0f63f18432cd0e1a66/Registration)). Once registered, you can create a new client. The client handles all requests going to the API and enables people to create there own requests making this client very extensible.
```go
    client, err := abnlookup.NewClient("YOUR GUID")
	if err != nil {
		// Handle error...
	}
```
With the client created you can then use the search methods already made, or create your own. Examples can be found in the [examples](https://github.com/joshturge/abnlookup/tree/master/example) directory.

## Methods

| Method | Unit Test |
| --- | --- |
| SearchByABN | Yes | 
| SearchByASIC | Yes |
| SearchByName | Yes |
| SearchByFilters | Yes |

## Todo
- Add more methods for searching ✅
- Add unit tests for methods ✅
- Create a better naming convention for types ✅
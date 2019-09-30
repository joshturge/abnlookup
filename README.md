# ABN Lookup Wrapper for Go

This wrapper attempts to provide easy access to the ABN Lookup API. Currently, the only two methods available for searching are the [SearchByABN](https://abr.business.gov.au/abrxmlsearch/Forms/SearchByABNv201408.aspx) and [SearchByACN](https://abr.business.gov.au/abrxmlsearch/Forms/SearchByASICv201408.aspx) endpoints.

This wrapper also provides functions that can validate both ABN and ACN numbers. These functions are exported and are used by SearchByABN and SearchByACN to validate before sending a request to the ABN Lookup API.

## Usage
To use the wrapper for the ABN Lookup API, you need to register for a GUID (more info [here](https://api.gov.au/service/5b639f0f63f18432cd0e1a66/Registration)). Once registered, you can create a new client within your program.
```go
    client, err := abnlookup.NewClient("YOUR GUID")
	if err != nil {
		// Handle error...
	}
```
With the client created you can then use the search methods, I have only implemented two search methods with more being added in the future. Examples can be found in the [examples](https://github.com/joshturge/abnlookup/tree/master/example) directory.

## Todo
- Add more methods for searching
- Create a better naming convention for types
- Add more helper functions and methods
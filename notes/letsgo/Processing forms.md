When a user fills in a form (e.g. signup), we need to 1. Define/ Create, 2. Parse, 3. Decode, 4. Validate, 5. Update 

Create:
- Create an empty object (struct) for the contents of the form

Parse:
- Use "net/http" for .ParseForm() method

Decode (fill empty form with parsed form info):
- Create method to Parse then decode form (call it decodePostForm)
- Use decodePostForm method on context
- Use "github.com/go-playground/form/v4" for .formDecoder.Decode() method

Validate:
- Check rules against defined methods (checkEmpty, checkNull, ...)

Update:
- Re-render:
	- If has errors 
		- render form with error banner
	- else
		- insert + redirect

Internal validator package contains helpers to check form content. Contains FieldErrors map[string]string
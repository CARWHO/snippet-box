- User authentication is tied to sessions. Methods for checking auth appear in middleware chains. 
- On every request auth should be checked (via middleware). On protected middleware chain, a second method (requireAuthentication) method is ran as it is required.
- Bcrypt used to hash password before inserting into DB.
- CSRF (cross site reference forgery)
	- User logs into our site, user visits attackers site, cookie gets hijacked. CSRF processes limit what cross sites can do. 
	- Place anti-CSRF package in middleware (nosurf) 

User authentication syntax:

import "golang.org/x/crypto/bcrypt"
hash, err := bcrypt.GenerateFromPassword([]byte(pw), 12)
// now store hash in DB.

// To check password user entered against hashed password:
err, ok : =bcrypt.CompareHashAndPassword(hash, userinputtedpassword)
...

User auth helper function:

- func (app `*`application) authenticate(next http.Handler) { http.Handler
	return http.HandlerFunc(func(w http.ResponseWriter, r `*`http.Request) http.Handler {
	// this function confirms user auth
	})
	}

...

Actual user auth check:

- dynamic := alice.New(... app.authenticate)
  
}

Duplicate version of this for routes that need protected, dynamic middleware.  


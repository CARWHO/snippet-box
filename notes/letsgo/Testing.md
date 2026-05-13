Common types of testing are: Unit tests (often table driven), integration testing (multiple component working together), and e2e testing. All test files in go use the `_test.go` convention. To create a test function/ method use the "testing" package. Test functions signatures should follow: func (`*`testing.T) format.

Table driven unit test syntax:

mytests := []struct {
	thing1: int 
	thing2: bool
	...
	want: string
} {
 {thing 1 = 10
 thing 2 = "test"
 want = true
 },
 {...},
 {...},
}

//^If this used a real sqlDB it would be considered an integration test.

for `_`, tt := range tests {
	t.Run(tt.name func(t `*`testing.T) {
	got := myfunc(tt.thing)
	if got != tt.want() {
		return t.Errorf("got x wanted y")
	} 
})
}

unit testing handlers w/o spinning up mock server:

rr := httptest.NewRecorder()
r := httptest.NewRequest(http.MethodGet, "/ping", nil)
app.ping(rr, r)
rs := rr.Result()

e2e testing syntax:

e2e test helpers:

func newTestApplication (t `*`testing.T) `*`application {
...
}

func newTestServer(t `*`testing.T, h http.Handler) `*`testServer {
...
}

e2e test

func TestPing(t `*`testing.T) {
app := newTestApplication(t)
testserver := newTestServer(t, app.routes)
...
}




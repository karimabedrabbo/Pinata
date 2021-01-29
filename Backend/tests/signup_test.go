package tests

import (
	k "github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"net/http"
	"testing"
)

func TestSignupPut(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	Seed(s.db, univ)
	account, password := GenerateAccount(GenerateEmail(univ.Domain), k.AccountRoleUser)
	user := GenerateUser()
	signupPut := requests.SignupPut{
		Name:   user.Name,
		Email:	account.Email,
		Password: password,
	}
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteSignup
	w := PerformRequest(s.router, http.MethodPost, path, "", signupPut)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusOK)
}

func TestSignupNotEduEmailPut(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	Seed(s.db, univ)
	account, password := GenerateAccount(IncrementString() + "@gmail.com", k.AccountRoleUser)
	user := GenerateUser()
	signupPut := requests.SignupPut{
		Name:   user.Name,
		Email:	account.Email,
		Password: password,
	}
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteSignup
	w := PerformRequest(s.router, http.MethodPost, path, "", signupPut)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusBadRequest)
}

func TestSignupEmailMatchesNoUniversityPut(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	//dont include the university in s.db
	//Seed(s.db, univ)
	account, password := GenerateAccount(GenerateEmail(univ.Domain), k.AccountRoleUser)
	user := GenerateUser()
	signupPut := requests.SignupPut{
		Name:   user.Name,
		Email:	account.Email,
		Password: password,
	}
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteSignup
	w := PerformRequest(s.router, http.MethodPost, path, "", signupPut)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusBadRequest)
}


func TestSignupDuplicateEmailPut(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	Seed(s.db, univ)
	account, password := GenerateAccount(GenerateEmail(univ.Domain), k.AccountRoleUser)
	Seed(s.db, account)
	user := GenerateUser()
	signupPut := requests.SignupPut{
		Name:   user.Name,
		Email:	account.Email,
		Password: password,
	}
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteSignup
	w := PerformRequest(s.router, http.MethodPost, path, "", signupPut)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusBadRequest)
}

func TestSignupPutFuzz(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)


	for i := 0; i < CREDENTIALS_REPEAT_ITER; i ++ {
		univ := GenerateUniversity()
		Seed(s.db, univ)

		signupPut := requests.SignupPut{}
		s.fuzzer.Fuzz(&signupPut)

		path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteSignup
		PerformRequest(s.router, http.MethodPost, path, "", signupPut)
	}
}

package tests

import (
	k "github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"github.com/karimabedrabbo/eyo/api/models/response"
	"net/http"
	"testing"
)

func TestLoginPost(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	Seed(s.db, univ)
	account, password := GenerateAccount(GenerateEmail(univ.Domain), k.AccountRoleUser)
	Seed(s.db, account)
	verify, _ := GenerateVerify(account.Id, true)
	Seed(s.db, verify)
	loginPost := requests.LoginPost{
		Email:	account.Email,
		Password: password,
	}
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteLogin
	w := PerformRequest(s.router, http.MethodPost, path, "", loginPost)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusOK)

	expectedPayloadClaims := &response.PayloadClaims{
		UserId:             account.UserId,
		AccountId:          account.Id,
		Role:               account.Role,
		EmailVerified:      true,
		UniversityId:       univ.Id,
		UniversityVerified: true,
	}
	CheckResponseToken(t, s, expectedPayloadClaims, w.Body.Bytes())
}

func TestLoginNonexistantVerifyPost(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	Seed(s.db, univ)
	account, password := GenerateAccount(GenerateEmail(univ.Domain), k.AccountRoleUser)
	Seed(s.db, account)
	//do not include a verify in s.db
	//verify, _ := GenerateVerify(account.Id, account.Email, true)
	//Seed(s.db, verify)
	loginPost := requests.LoginPost{
		Email:	account.Email,
		Password: password,
	}
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteLogin
	w := PerformRequest(s.router, http.MethodPost, path, "", loginPost)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusUnauthorized)
}

func TestLoginUnverifiedPost(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	Seed(s.db, univ)
	account, password := GenerateAccount(GenerateEmail(univ.Domain), k.AccountRoleUser)
	Seed(s.db, account)
	//include a verify request, no confirm
	verify, _ := GenerateVerify(account.Id, false)
	Seed(s.db, verify)
	loginPost := requests.LoginPost{
		Email:	account.Email,
		Password: password,
	}
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteLogin
	w := PerformRequest(s.router, http.MethodPost, path, "", loginPost)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusUnauthorized)
}

func TestLoginBannedPost(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	Seed(s.db, univ)
	account, password := GenerateAccount(GenerateEmail(univ.Domain), k.AccountRoleUser)
	Seed(s.db, account)
	verify, _ := GenerateVerify(account.Id, true)
	Seed(s.db, verify)
	ban := GenerateBan(account.Id)
	Seed(s.db, ban)

	loginPost := requests.LoginPost{
		Email:	account.Email,
		Password: password,
	}
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteLogin
	w := PerformRequest(s.router, http.MethodPost, path, "", loginPost)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusUnauthorized)
}

func TestLoginNonexistantPost(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	Seed(s.db, univ)
	account, password := GenerateAccount(GenerateEmail(univ.Domain), k.AccountRoleUser)
	//do not include account in s.db
	//Seed(s.db, account)
	verify, _ := GenerateVerify(account.Id, false)
	Seed(s.db, verify)
	loginPost := requests.LoginPost{
		Email:	account.Email,
		Password: password,
	}
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteLogin
	w := PerformRequest(s.router, http.MethodPost, path, "", loginPost)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusUnauthorized)
}

func TestLoginIncorrectPasswordPost(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	Seed(s.db, univ)
	account, password := GenerateAccount(GenerateEmail(univ.Domain), k.AccountRoleUser)
	Seed(s.db, account)
	verify, _ := GenerateVerify(account.Id, true)
	Seed(s.db, verify)
	loginPost := requests.LoginPost{
		Email:	account.Email,
		Password: password + IncrementString(), //use a bad password
	}
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteLogin
	w := PerformRequest(s.router, http.MethodPost, path, "", loginPost)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusUnauthorized)
}

func TestLoginAnonymousPost(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	Seed(s.db, univ)
	loginPost := requests.LoginAnonymousPost{
		UniversityId:	univ.Id,
	}
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteLoginAnonymous
	w := PerformRequest(s.router, http.MethodPost, path, "", loginPost)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusOK)

	//this is a little ugly. For anonymous logins we don't know the user id in advance (because user is dynamically created)
	//a little hacky but we assume auto-increment will set the id to 1 (since we refreshed the s.db)
	expectedPayloadClaims := &response.PayloadClaims{
		UserId:             1,
		AccountId:          0,
		Role:               k.AccountRoleAnonymous,
		EmailVerified:      false,
		UniversityId:       univ.Id,
		UniversityVerified: false,
	}
	CheckResponseToken(t, s, expectedPayloadClaims, w.Body.Bytes())

}

func TestLoginAnonymousNonexistantUniversityPost(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	//do not include univ in s.db
	//Seed(s.db, univ)
	loginPost := requests.LoginAnonymousPost{
		UniversityId:	univ.Id,
	}
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteLoginAnonymous
	w := PerformRequest(s.router, http.MethodPost, path, "", loginPost)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusUnauthorized)
}

func TestLoginPostFuzz(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	for i := 0; i < CREDENTIALS_REPEAT_ITER; i ++ {
		univ := GenerateUniversity()
		Seed(s.db, univ)
		account, _ := GenerateAccount(GenerateEmail(univ.Domain), k.AccountRoleUser)
		Seed(s.db, account)
		verify, _ := GenerateVerify(account.Id, true)
		Seed(s.db, verify)

		loginPost := requests.LoginPost{}
		s.fuzzer.Fuzz(&loginPost)

		path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteLogin
		PerformRequest(s.router, http.MethodPost, path, "", loginPost)
	}
}

func TestLoginAnonymousPostFuzz(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	for i := 0; i < CREDENTIALS_REPEAT_ITER; i ++ {
		univ := GenerateUniversity()
		Seed(s.db, univ)

		loginPost := requests.LoginAnonymousPost{}
		s.fuzzer.Fuzz(&loginPost)

		path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirPublic + k.RouteLoginAnonymous
		PerformRequest(s.router, http.MethodPost, path, "", loginPost)
	}

}
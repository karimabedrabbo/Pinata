package tests

import (
	k "github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"net/http"
	"testing"
)

func TestUniversityAnonymousGet(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	Seed(s.db, univ)
	token := TokenGeneratorAnonymous(s.auth, 1, 1)
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirAnonymous + k.RouteUniversity + IdToRoute(univ.Id)
	w := PerformRequest(s.router, http.MethodGet, path, token, nil)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusOK)

	CheckResponse(t, s, univ, w.Body.Bytes())
}

func TestUniversityUserGet(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	Seed(s.db, univ)
	token := TokenGeneratorUser(s.auth, 1, 1, 1)
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirUser + k.RouteUniversity + IdToRoute(univ.Id)
	w := PerformRequest(s.router, http.MethodGet, path, token, nil)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusOK)
	CheckResponse(t, s, univ, w.Body.Bytes())
}

func TestUniversityAnonymousList(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univs := GenerateUniversities(100)
	for _, univ := range univs {
		Seed(s.db, univ)
	}

	token := TokenGeneratorAnonymous(s.auth, 1, 1)
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirAnonymous + k.RouteUniversity + k.RouteActionList
	w := PerformRequest(s.router, http.MethodGet, path, token, nil)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusOK)
	CheckResponse(t, s, univs, w.Body.Bytes())
}

func TestUniversityUserList(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univs := GenerateUniversities(100)
	for _, univ := range univs {
		Seed(s.db, univ)
	}

	token := TokenGeneratorUser(s.auth, 1, 1, 1)
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirUser + k.RouteUniversity + k.RouteActionList
	w := PerformRequest(s.router, http.MethodGet, path, token, nil)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusOK)
	CheckResponse(t, s, univs, w.Body.Bytes())
}


func TestUniversitySysadminPost(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	Seed(s.db, univ)
	univPost := requests.UniversityPost{
		UniversityId: univ.Id,
		Name:   univ.Name,
		Alias:  univ.Alias,
		Size:   univ.Size,
	}
	token := TokenGeneratorSysadmin(s.auth)
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirSysadmin + k.RouteUniversity + IdToRoute(univ.Id)
	w := PerformRequest(s.router, http.MethodPost, path, token, univPost)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusOK)
}


func TestUniversitySysadminPut(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	univPut := requests.UniversityPut{
		Name:   univ.Name,
		Alias:  univ.Alias,
		Domain: univ.Domain,
		Lat:    univ.Lat,
		Lng:    univ.Lng,
		Size:   univ.Size,
	}
	token := TokenGeneratorSysadmin(s.auth)
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirSysadmin + k.RouteUniversity
	w := PerformRequest(s.router, http.MethodPut, path, token, univPut)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusOK)
}


func TestUniversityGetNonexistant(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	univ := GenerateUniversity()
	//do not seed
	//Seed(s.db, univ)
	token := TokenGeneratorAnonymous(s.auth, 1, 1)
	path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirAnonymous + k.RouteUniversity + IdToRoute(univ.Id)
	w := PerformRequest(s.router, http.MethodGet, path, token, nil)

	CheckResponseCode(t, s, w.Code, w.Body.String(), http.StatusNotFound, http.StatusBadRequest)
}


func TestUniversityPutFuzz(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)

	for i := 0; i < REPEAT_ITER; i ++ {
		univPut := requests.UniversityPut{}
		s.fuzzer.Fuzz(&univPut)

		token := TokenGeneratorSysadmin(s.auth)
		path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirSysadmin + k.RouteUniversity
		PerformRequest(s.router, http.MethodPut, path, token, univPut)
	}
}

func TestUniversityPostFuzz(t *testing.T) {
	SetupDatabase(s.db)
	defer CleanDatabase(s.db)


	for i := 0; i < REPEAT_ITER; i ++ {
		univ := GenerateUniversity()
		Seed(s.db, univ)

		univPost := requests.UniversityPost{}
		s.fuzzer.Fuzz(&univPost)

		token := TokenGeneratorSysadmin(s.auth)
		path := k.AppDevelopmentUrl + k.AppApiPath + k.RouteDirSysadmin + k.RouteUniversity + IdToRoute(univ.Id)
		PerformRequest(s.router, http.MethodPost, path, token, univPost)
	}
}



package resources

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"github.com/gin-gonic/gin"

	"github.com/tonyalaribe/dsi-hackaton/models"
)

func TestLocationHandlers(t *testing.T) {

	//Test Login Ability
	convey.Convey("Test Create Location Ability", t, func() {

		router := gin.New()
		//router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
		router.GET("/get", Location{}.Get)

		router.POST("/create", Location{}.Post)

		postJSON := []byte(`
			{
				"LocationName": "Okpanam road bin",
				"Latitude": 6.219194,
				"Longitude":6.691308
			}
 `)

		postJSON2 := []byte(`
	{
		"LocationName": "Federal Secreteriat",
		"Latitude": 6.219247,
		"Longitude":6.688951
	}
`)
		postJSON3 := []byte(`
	{
		"LocationName": "Nta Bin",
		"Latitude": 6.221998,
		"Longitude":6.684780
	}
`)

		req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(postJSON))

		convey.So(err, convey.ShouldEqual, nil)
		//req.Header.Set("Authorization", token)

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		dataResource := models.Location{}

		err = json.NewDecoder(resp.Body).Decode(&dataResource)
		convey.So(err, convey.ShouldBeNil)
		//t.Log(dataResource)

		convey.So(resp.Code, convey.ShouldEqual, 200)

		req, err = http.NewRequest("POST", "/create", bytes.NewBuffer(postJSON2))

		convey.So(err, convey.ShouldEqual, nil)
		//req.Header.Set("Authorization", token)

		resp = httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		dataResource = models.Location{}

		err = json.NewDecoder(resp.Body).Decode(&dataResource)
		convey.So(err, convey.ShouldBeNil)
		//t.Log(dataResource)

		convey.So(resp.Code, convey.ShouldEqual, 200)
		//////
		req, err = http.NewRequest("POST", "/create", bytes.NewBuffer(postJSON3))

		convey.So(err, convey.ShouldEqual, nil)
		//req.Header.Set("Authorization", token)

		resp = httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		dataResource = models.Location{}

		err = json.NewDecoder(resp.Body).Decode(&dataResource)
		convey.So(err, convey.ShouldBeNil)
		//t.Log(dataResource)

		convey.So(resp.Code, convey.ShouldEqual, 200)

		//test JWT Decoding Ability
		convey.Convey("Should be able to get a list of locations ", func() {

			req, err := http.NewRequest("GET", "/get", nil)
			convey.So(err, convey.ShouldBeNil)

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			t.Log(resp.Body.String())
			convey.So(resp.Code, convey.ShouldEqual, 200)
		})

	})

}

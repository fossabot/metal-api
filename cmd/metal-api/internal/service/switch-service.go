package service

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	"git.f-i-ts.de/cloud-native/metal/metal-api/cmd/metal-api/internal/datastore"
	"git.f-i-ts.de/cloud-native/metal/metal-api/metal"
	restful "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
)

type switchResource struct {
	webResource
}

type switchRegistration struct {
	ID     string     `json:"id" description:"a unique ID" unique:"true"`
	Nics   metal.Nics `json:"nics" description:"the list of network interfaces on the switch"`
	SiteID string     `json:"site_id" description:"the id of the site in which this switch is located"`
	RackID string     `json:"rack_id" description:"the id of the rack in which this switch is located"`
}

func NewSwitch(log *zap.Logger, ds *datastore.RethinkStore) *restful.WebService {
	sr := switchResource{
		webResource: webResource{
			SugaredLogger: log.Sugar(),
			log:           log,
			ds:            ds,
		},
	}
	return sr.webService()
}

func (sr switchResource) webService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/switch").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"switch"}

	ws.Route(ws.GET("/").To(sr.restListGet(sr.ds.ListSwitches)).
		Operation("listSwitches").
		Doc("get all switches").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]metal.Switch{}).
		Returns(http.StatusOK, "OK", []metal.Switch{}))

	ws.Route(ws.DELETE("/{id}").To(sr.restEntityGet(sr.ds.DeleteSwitch)).
		Operation("deleteSwitch").
		Doc("deletes an switch and returns the deleted entity").
		Param(ws.PathParameter("id", "identifier of the switch").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(metal.Switch{}).
		Returns(http.StatusOK, "OK", metal.Switch{}).
		Returns(http.StatusNotFound, "Not Found", nil))

	ws.Route(ws.POST("/register").To(sr.registerSwitch).
		Doc("register a switch").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(switchRegistration{}).
		Returns(http.StatusOK, "OK", metal.Switch{}).
		Returns(http.StatusCreated, "Created", metal.Switch{}).
		Returns(http.StatusConflict, "Conflict", nil))

	return ws
}

func (sr switchResource) registerSwitch(request *restful.Request, response *restful.Response) {
	var newSwitch switchRegistration
	err := request.ReadEntity(&newSwitch)
	if checkError(sr.log, response, "registerSwitch", err) {
		return
	}
	_, err = sr.ds.FindSite(newSwitch.SiteID)
	if checkError(sr.log, response, "registerSwitch", err) {
		return
	}

	oldSwitch, err := sr.ds.FindSwitch(newSwitch.ID)
	sw := metal.NewSwitch(newSwitch.ID, newSwitch.SiteID, newSwitch.RackID, newSwitch.Nics)
	if err != nil {
		if metal.IsNotFound(err) {
			sw.Created = time.Now()
			sw.Changed = sw.Created
			sw, err = sr.ds.CreateSwitch(sw)
			if checkError(sr.log, response, "registerSwitch", err) {
				return
			}
			response.WriteHeaderAndEntity(http.StatusCreated, sw)
			return
		}
		sendError(sr.log, response, "registerSwitch", http.StatusInternalServerError, err)
		return
	}
	// Make sure we do not delete current connections
	sw.Connections = oldSwitch.Connections

	err = sr.ds.UpdateSwitch(oldSwitch, sw)

	if checkError(sr.log, response, "registerSwitch", err) {
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, sw)
}
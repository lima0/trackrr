package dhl

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
)

func GetResponse(ID string) (Response, error) {
	r, err := http.Get("https://mydhl.express.dhl/shipmentTracking?AWB=" + ID + "&countryCode=us&languageCode=en")
	if err != nil {
		return Response{}, err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return Response{}, err
	}
	res := Response{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return Response{}, err
	}
	return res, nil
}

func PrintResponse(res Response) {
	c := color.New(color.FgBlue)
	for _, shipment := range res.Results {
		c.Printf("Tracking Number: ")
		fmt.Printf("%s\n", shipment.ID)
		c.Printf("Origin: ")
		fmt.Printf("%s\n", shipment.Origin.Value)
		c.Printf("Destination: ")
		fmt.Printf("%s\n", shipment.Destination.Value)
		c.Printf("Last Event: ")
		fmt.Printf("%s\n", shipment.Description)
		c.Printf("Remark: ")
		fmt.Printf("%s\n", shipment.EventRemark)
		c.Printf("Next step: ")
		fmt.Printf("%s\n\n", shipment.EventNextStep)
	}
}

type Response struct {
	Results []Results `json:"results"`
}
type Delivery struct {
	Code   string `json:"code"`
	Status string `json:"status"`
}
type Origin struct {
	Value string `json:"value"`
	Label string `json:"label"`
	URL   string `json:"url"`
}
type Destination struct {
	Value string `json:"value"`
	Label string `json:"label"`
	URL   string `json:"url"`
}
type Link struct {
	URL   string `json:"url"`
	Label string `json:"label"`
}
type Signature struct {
	Link        Link   `json:"link"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Signatory   string `json:"signatory"`
	Help        string `json:"help"`
}
type Pieces struct {
	Value       int      `json:"value"`
	Label       string   `json:"label"`
	ShowSummary bool     `json:"showSummary"`
	PIDs        []string `json:"pIds"`
}
type Checkpoints struct {
	Counter     int      `json:"counter"`
	Description string   `json:"description"`
	Time        string   `json:"time"`
	Date        string   `json:"date"`
	Location    string   `json:"location"`
	TotalPieces int      `json:"totalPieces,omitempty"`
	PIDs        []string `json:"pIds,omitempty"`
}
type Edd struct {
	Label    string `json:"label"`
	Date     string `json:"date"`
	Product  string `json:"product"`
	Comments string `json:"comments"`
}
type Results struct {
	ID                      string        `json:"id"`
	Label                   string        `json:"label"`
	Type                    string        `json:"type"`
	Duplicate               bool          `json:"duplicate"`
	Delivery                Delivery      `json:"delivery"`
	Origin                  Origin        `json:"origin"`
	Destination             Destination   `json:"destination"`
	Description             string        `json:"description"`
	EventRemark             string        `json:"eventRemark"`
	EventNextStep           string        `json:"eventNextStep"`
	HasDuplicateShipment    bool          `json:"hasDuplicateShipment"`
	Signature               Signature     `json:"signature"`
	Pieces                  Pieces        `json:"pieces"`
	Checkpoints             []Checkpoints `json:"checkpoints"`
	CheckpointLocationLabel string        `json:"checkpointLocationLabel"`
	CheckpointTimeLabel     string        `json:"checkpointTimeLabel"`
	Edd                     Edd           `json:"edd"`
}

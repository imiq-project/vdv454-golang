package main

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Leitstelle struct {
	Kennung string
	Addr    string
	Abos    []AboAUS
}

type VdvServer struct {
	server             *http.Server
	leitstellenKennung string
	leitstellen        map[string]Leitstelle
}

func NewVdvServer(LeitstellenKennung string, Addr string) VdvServer {
	result := VdvServer{leitstellenKennung: LeitstellenKennung}
	result.server = &http.Server{Addr: Addr}
	mux := http.NewServeMux()
	result.server.Handler = mux
	mux.HandleFunc("/"+LeitstellenKennung+"/status.xml", result.handleStatus)
	mux.HandleFunc("/"+LeitstellenKennung+"/aboverwalten.xml", result.handleSubscription)
	mux.HandleFunc("/"+LeitstellenKennung+"/datenabrufen.xml", result.handleDataRequest)
	mux.HandleFunc("/"+LeitstellenKennung+"/clientstatus.xml", result.handleStatus)
	mux.HandleFunc("/"+LeitstellenKennung+"/datenabereit.xml", result.handleDataReady)
	result.leitstellen = make(map[string]Leitstelle)
	return result
}

func (s *VdvServer) Start() {
	go func() {
		err := s.server.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal("Server startup failed", err)
		}
	}()
}

func (s *VdvServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.server.Shutdown(ctx)
}

func (s *VdvServer) AddLeitstelle(Kennung string, Addr string) {
	s.leitstellen[Kennung] = Leitstelle{Kennung: Kennung, Addr: Addr}
}

func (s *VdvServer) Subscribe(LeitstellenKennung string, AboID string) {
	// Create some data
	request := AboAnfrage{Sender: s.leitstellenKennung, Zst: time.Now(), AboAUS: []AboAUS{
		{AboID: AboID, VerfallZst: time.Now()},
	}}

	// Marshal to XML
	xmlData, err := xml.MarshalIndent(request, "", "  ")
	if err != nil {
		panic(err)
	}

	// Send POST request with XML body
	resp, err := http.Post("http://localhost:80/"+LeitstellenKennung+"/aboverwalten.xml", "application/xml", bytes.NewBuffer(xmlData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status:", resp.Status)
	fmt.Println("Response body:", string(body))
}

func (s *VdvServer) handleStatus(w http.ResponseWriter, r *http.Request) {
}

func (s *VdvServer) handleSubscription(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read and parse the XML body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var subscriptionReq AboAnfrage
	if err := xml.Unmarshal(body, &subscriptionReq); err != nil {
		http.Error(w, "Failed to parse XML", http.StatusBadRequest)
		return
	}

	// Respond
	log.Println("Received subscription request from", subscriptionReq.Sender)
	leitestelle, ok := s.leitstellen[subscriptionReq.Sender]
	if !ok {
		log.Println("Unknown LeitstellenKennung", subscriptionReq.Sender)
		http.Error(w, "Unbekannte LeitstellenKennung", http.StatusBadRequest)
		return
	}
	subscriptionResp := AboAntwort{BestaetigungMitAboID: []BestaetigungMitAboID{}}
	for _, i := range subscriptionReq.AboAUS {
		entry := BestaetigungMitAboID{AboID: i.AboID}
		subscriptionResp.BestaetigungMitAboID = append(subscriptionResp.BestaetigungMitAboID, entry)
		leitestelle.Abos = append(leitestelle.Abos, i)
		log.Println("Added subscription", i.AboID)
	}
	w.WriteHeader(http.StatusOK)
	err = xml.NewEncoder(w).Encode(subscriptionResp)
	if err != nil {
		http.Error(w, "Failed to write XML", http.StatusBadRequest)
		return
	}
}

func (s *VdvServer) handleDataRequest(w http.ResponseWriter, r *http.Request) {

}

func (s *VdvServer) handleDataReady(w http.ResponseWriter, r *http.Request) {

}

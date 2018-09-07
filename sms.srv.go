package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/time/rate"

	"github.com/andrcmdr/go-smpp/smpp"
	"github.com/andrcmdr/go-smpp/smpp/pdu"
	"github.com/andrcmdr/go-smpp/smpp/pdu/pdufield"
	"github.com/andrcmdr/go-smpp/smpp/pdu/pdutext"
)

func Receiver() {
	f := func(p pdu.Body) {
		switch p.Header().ID {
		case pdu.DeliverSMID:
			f := p.Fields()
			src := f[pdufield.SourceAddr]
			dst := f[pdufield.DestinationAddr]
			txt := f[pdufield.ShortMessage]
			log.Printf("Short message from=%s to=%s: %s", src, dst, txt)
		}
	}
	r := &smpp.Receiver{
		Addr:    "194.176.111.242:8018",
		User:    "DOSCredo",
		Passwd:  "cReD0!x",
		Handler: f,
	}
	// Create persistent connection.
	conn := r.Bind()
	time.AfterFunc(10*time.Second, func() { r.Close() })
	// Print connection status (Connected, Disconnected, etc).
	for c := range conn {
		log.Println("SMPP connection status:", c.Status())
	}
}

func Transmitter() {
	tx := &smpp.Transmitter{
		Addr:   "194.176.111.242:8018",
		User:   "DOSCredo",
		Passwd: "cReD0!x",
	}
	// Create persistent connection, wait for the first status.
	conn := <-tx.Bind()
	if conn.Status() != smpp.Connected {
		log.Fatal(conn.Error())
	}
	sm, err := tx.Submit(&smpp.ShortMessage{
		Src:      "6791",
		Dst:      "996771977377",
		Text:     pdutext.UCS2("Я твой Дед!\nМен сиздин Дед!\nМен сиздин Чоң-Ата! [Ѳө Ңң Үү]\nI\x27m your Grandfather, Luke!\n"),
		Register: pdufield.FinalDeliveryReceipt,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Message ID:", sm.RespID())
}

func Transceiver() {
	f := func(p pdu.Body) {
		switch p.Header().ID {
		case pdu.DeliverSMID:
			f := p.Fields()
			src := f[pdufield.SourceAddr]
			dst := f[pdufield.DestinationAddr]
			txt := f[pdufield.ShortMessage]
			log.Printf("Short message from=%s to=%s: %s", src, dst, txt)
		}
	}
	lm := rate.NewLimiter(rate.Limit(10000000), 1) // Max rate of 10M messages per second
	tx := &smpp.Transceiver{
		Addr:        "194.176.111.242:8018",
		User:        "DOSCredo",
		Passwd:      "cReD0!x",
		Handler:     f,  // Handle incoming SM or delivery receipts.
		RateLimiter: lm, // Optional rate limiter.
	}
	// Create persistent connection.
	conn := tx.Bind()
	go func() {
		for c := range conn {
			log.Println("SMPP connection status:", c.Status())
		}
	}()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sm, err := tx.Submit(&smpp.ShortMessage{
			Src:      r.FormValue("source"),
			Dst:      r.FormValue("target"),
			Text:     pdutext.UCS2(r.FormValue("text")),
			Register: pdufield.FinalDeliveryReceipt,
		})
		if err == smpp.ErrNotConnected {
			http.Error(w, "Error: Not connected to SMPP gateway.", http.StatusServiceUnavailable)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		io.WriteString(w, sm.RespID() + "\n")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	Transceiver()
//	Transmitter()
//	Receiver()
}


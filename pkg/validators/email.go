package validators

import (
	"log"
	"net"
	"net/mail"
	"strings"
)

func ValidateEmail(email string) (bool, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, err
	}
	if _, err := validateMailDomain(strings.Split(email, "@")[1]); err != nil {
		log.Println("validateMailDomain: ", err)
		return false, err
	}
	return true, nil
}

func validateMailDomain(domain string) (bool, error) {
	var hasMX, hasSPF, hasDMARC bool

	var spfRec, dmarcRec string

	mxRec, err := net.LookupMX(domain)
	if err != nil {
		return false, err
	}
	if len(mxRec) > 0 {
		hasMX = true
	}

	txtRecs, err := net.LookupTXT(domain)
	if err != nil {
		return false, err
	}

	for _, rec := range txtRecs {
		if strings.HasPrefix(rec, "v=spf1") {
			hasSPF = true
			spfRec = rec
			break
		}
	}

	dmarcRecs, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		return false, err
	}
	for _, dRec := range dmarcRecs {
		if strings.HasPrefix(dRec, "v=DMARC1") {
			hasDMARC = true
			dmarcRec = dRec
		}
	}

	log.Printf("%v, %v, %v, %v, %v", hasMX, hasSPF, spfRec, hasDMARC, dmarcRec)
	return true, nil
}

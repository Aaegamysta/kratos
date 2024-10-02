package testhelpers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/exp/rand"

	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/session"
)

func NewInActiveSession(r *http.Request, reg interface {
	session.ManagementProvider
}, i *identity.Identity, authenticatedAt time.Time, completedLoginFor identity.CredentialsType, completedLoginAAL identity.AuthenticatorAssuranceLevel) (*session.Session, error) {
	s := session.NewInactiveSession()
	s.CompletedLoginFor(completedLoginFor, completedLoginAAL)
	if err := reg.SessionManager().ActivateSession(r, s, i, authenticatedAt); err != nil {
		return nil, err
	}
	s.Active = false
	return s, nil
}

// Tricky to create multiple sessions with different locations and different authentication time
// Each created session will be one a day older, so if 5 sessions are created the oldest will be authenticated
// 5 days ago within the vicinity of the provided location.
// Because of time constrains, it could
func NewInActiveSessions(r *http.Request, n int, reg interface {
	session.ManagementProvider
}, i *identity.Identity, authenticatedAt time.Time, completedLoginFor identity.CredentialsType, completedLoginAAL identity.AuthenticatorAssuranceLevel) ([]*session.Session, error) {
	sessions := make([]*session.Session, 0)
	baseLatitude := r.Header.Get("Latitude")
	latitude, _ := strconv.ParseFloat(baseLatitude, 64)
	baseLongitude := r.Header.Get("Longitude")
	longitude, _ := strconv.ParseFloat(baseLongitude, 64)
	for j := 1; j <= n; j++ {
		s := session.NewInactiveSession()
		r.Header.Set("Latitude", fmt.Sprintf("%f.%d", latitude, rand.Intn(9999)))
		r.Header.Set("Longitude", fmt.Sprintf("%f.%d", longitude, rand.Intn(9999)))
		s.CompletedLoginFor(completedLoginFor, completedLoginAAL)
		if err := reg.SessionManager().ActivateSession(r, s, i, authenticatedAt.AddDate(0, 0, -j)); err != nil {
			return nil, err
		}
		s.Active = false
		sessions = append(sessions, s)
	}
	return sessions, nil
}
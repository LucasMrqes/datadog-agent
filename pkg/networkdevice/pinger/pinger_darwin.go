//go:build darwin

package pinger

// MacPinger implements the Pinger interface for
// macOS users
type MacPinger struct {
	cfg Config
}

// New creates a MacPinger using the passed in
// config
func New(cfg Config) (Pinger, error) {
	if cfg.UseRawSocket {
		return nil, ErrRawSocketUnsupported
	}
	return &MacPinger{
		cfg: cfg,
	}, nil
}

// Ping takes a host and depending on the config will either
// directly ping the host sending packets over a UDP socket
// or a raw socket
func (p *MacPinger) Ping(host string) (*Result, error) {
	return RunPing(&p.cfg, host)
}

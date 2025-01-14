package nftables

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

// firewallRenderingData holds the data available in the nftables template
type firewallRenderingData struct {
	ForwardingRules  forwardingRules
	RateLimitRules   nftablesRules
	SnatRules        nftablesRules
	InternalPrefixes string
	PrivateVrfID     uint
}

func newFirewallRenderingData(f *Firewall) (*firewallRenderingData, error) {
	ingress, egress := nftablesRules{}, nftablesRules{}
	for _, np := range f.clusterwideNetworkPolicies.Items {
		err := np.Spec.Validate()
		if err != nil {
			continue
		}
		i, e := clusterwideNetworkPolicyRules(np, f.logAcceptedConnections)
		ingress = append(ingress, i...)
		egress = append(egress, e...)
	}

	for _, svc := range f.services.Items {
		ingress = append(ingress, serviceRules(svc, f.logAcceptedConnections)...)
	}

	snatRules, err := snatRules(f)
	if err != nil {
		return &firewallRenderingData{}, err
	}

	return &firewallRenderingData{
		PrivateVrfID:     uint(*f.primaryPrivateNet.Vrf),
		InternalPrefixes: strings.Join(f.spec.InternalPrefixes, ", "),
		ForwardingRules: forwardingRules{
			Ingress: ingress,
			Egress:  egress,
		},
		RateLimitRules: rateLimitRules(f),
		SnatRules:      snatRules,
	}, nil
}

func (d *firewallRenderingData) write(file string) error {
	c, err := d.renderString()
	if err != nil {
		return err
	}
	err = os.WriteFile(file, []byte(c), 0600)
	if err != nil {
		return fmt.Errorf("error writing to nftables file '%s': %w", file, err)
	}
	return nil
}

func (d *firewallRenderingData) renderString() (string, error) {
	var b bytes.Buffer

	tplString, err := d.readTpl()
	if err != nil {
		return "", err
	}

	tpl := template.Must(template.New("v4").Parse(tplString))

	err = tpl.Execute(&b, d)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func (d *firewallRenderingData) readTpl() (string, error) {
	r, err := templates.Open("nftables.tpl")
	if err != nil {
		return "", err
	}
	defer r.Close()
	bytes, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

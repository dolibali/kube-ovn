package util

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func TestDialTCP(t *testing.T) {
	tests := []struct {
		name     string
		host     string
		timeout  time.Duration
		verbose  bool
		expected error
	}{
		{"Valid HTTP Host", "http://localhost:8080", 1 * time.Second, false, nil},
		{"Valid HTTP Host", "http://localhost:8080", 1 * time.Second, true, nil},
		{"Valid HTTPS Host", "https://localhost:8443", 1 * time.Second, false, nil},
		{"Valid TCP Host", "tcp://localhost:8081", 1 * time.Second, false, nil},
		{"Invalid Host", "https://localhost%:8443", 1 * time.Second, false, fmt.Errorf("failed to parse host")},
		{"Unsupported Scheme", "ftp://localhost:8080", 1 * time.Second, false, fmt.Errorf("unsupported scheme")},
		{"Timeout", "http://localhost:8080", 1 * time.Millisecond, false, fmt.Errorf("timed out dialing host")},
	}

	httpServer := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	httpServer.StartTLS()
	defer httpServer.Close()

	tcpListener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		t.Fatalf("failed to start tcp server: %v", err)
	}
	defer tcpListener.Close()

	go func() {
		for {
			conn, err := tcpListener.Accept()
			if err != nil {
				return
			}
			conn.Close()
		}
	}()

	for i, tc := range tests {
		if tc.host == "http://localhost:8080" {
			tests[i].host = httpServer.URL
		} else if tc.host == "https://localhost:8443" {
			tests[i].host = httpServer.URL
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.host == "http://localhost:8080" || tt.host == "https://localhost:8443" {
				httpServer.Close()
				defer httpServer.StartTLS()
			}

			err := DialTCP(tt.host, tt.timeout, tt.verbose)
			if err != tt.expected && (tt.expected == nil || !strings.Contains(err.Error(), tt.expected.Error())) {
				t.Errorf("DialTCP(%q) got %v, want %v", tt.host, err, tt.expected)
			}

			if tt.verbose {
				klog.Flush()
			}
		})
	}
}

func TestDialAPIServer(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() (string, func())
		expected error
	}{
		{
			name: "Successful Dial",
			setup: func() (string, func()) {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
				return server.URL, server.Close
			},
			expected: nil,
		},
		{
			name: "Successful TLS Dial",
			setup: func() (string, func()) {
				server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
				return server.URL, server.Close
			},
			expected: nil,
		},
		{
			name: "Failed Dial",
			setup: func() (string, func()) {
				return "http://localhost:12345", func() {}
			},
			expected: fmt.Errorf("timed out dialing apiserver"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			host, cleanup := tt.setup()
			defer cleanup()

			err := DialAPIServer(host, 1*time.Second, 1)

			if tt.expected == nil && err != nil {
				t.Errorf("expected no error, got %v", err)
			} else if tt.expected != nil && (err == nil || !strings.Contains(err.Error(), tt.expected.Error())) {
				t.Errorf("expected error containing %v, got %v", tt.expected, err)
			}
		})
	}
}

func TestGetNodeInternalIP(t *testing.T) {
	tests := []struct {
		name string
		node v1.Node
		exp4 string
		exp6 string
	}{
		{
			name: "correct",
			node: v1.Node{
				TypeMeta:   metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{},
				Spec:       v1.NodeSpec{},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{
							Type:    "InternalIP",
							Address: "192.168.0.2",
						},
						{
							Type:    "ExternalIP",
							Address: "192.188.0.4",
						},
						{
							Type:    "InternalIP",
							Address: "ffff:ffff:ffff:ffff:ffff::23",
						},
					},
				},
			},
			exp4: "192.168.0.2",
			exp6: "ffff:ffff:ffff:ffff:ffff::23",
		},
		{
			name: "correctWithDiff",
			node: v1.Node{
				TypeMeta:   metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{},
				Spec:       v1.NodeSpec{},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{
							Type:    "InternalIP",
							Address: "ffff:ffff:ffff:ffff:ffff::23",
						},
						{
							Type:    "ExternalIP",
							Address: "192.188.0.4",
						},
						{
							Type:    "InternalIP",
							Address: "192.188.0.43",
						},
					},
				},
			},
			exp4: "192.188.0.43",
			exp6: "ffff:ffff:ffff:ffff:ffff::23",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ret4, ret6 := GetNodeInternalIP(tt.node); ret4 != tt.exp4 || ret6 != tt.exp6 {
				t.Errorf("got %v, %v, want %v, %v", ret4, ret6, tt.exp4, tt.exp6)
			}
		})
	}
}

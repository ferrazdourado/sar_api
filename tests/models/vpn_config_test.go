package models_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/ferrazdourado/sar_api/internal/models"
)

func TestVPNConfig_Validation(t *testing.T) {
    tests := []struct {
        name    string
        config  models.VPNConfig
        wantErr bool
    }{
        {
            name: "Configuração válida",
            config: models.VPNConfig{
                Name:          "Test Config",
                ServerAddress: "vpn.example.com",
                Port:         1194,
                Protocol:     "udp",
                Config:       "client\ndev tun\nproto udp...",
            },
            wantErr: false,
        },
        {
            name: "Sem nome",
            config: models.VPNConfig{
                ServerAddress: "vpn.example.com",
                Port:         1194,
                Protocol:     "udp",
                Config:       "client\ndev tun\nproto udp...",
            },
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Aqui você implementaria a validação real
            // Por exemplo, usando um pacote de validação como "validator"
            err := validate.Struct(tt.config)
            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
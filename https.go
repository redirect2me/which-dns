package main

import (
    "fmt"
    "github.com/mholt/certmagic"
)

func https_init() {
    // read and agree to your CA's legal documents
    certmagic.Default.Agreed = true

    // provide an email address
    certmagic.Default.Email = "fileformat@gmail.com"

    // use the staging endpoint while we're developing
    certmagic.Default.CA = certmagic.LetsEncryptStagingCA

    certmagic.Default.OnDemand = &certmagic.OnDemandConfig{
        DecisionFunc: func(name string) error {
            if name != "example.com" {
                return fmt.Errorf("not allowed")
            }
            return nil
        },
    }
}
package main

// Demo fixture: a fake generic credential string for secret-scanner testing.
// This is not a real secret.
var demoServiceAPISecret = "demo1f3b9c7e5a2d4f8b6c0e9a7d5f3b1c8e6a4d2f0b9"

func serviceCredential() string {
	return demoServiceAPISecret
}

package main

import (
	"fmt"

	"github.com/pchchv/gotp"
)

func defaultTOTPUsage() {
	otp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")

	fmt.Println("current one-time password is:", otp.Now())
	fmt.Println("one-time password of timestamp 0 is:", otp.At(0))
	fmt.Println(otp.ProvisioningUri("demoAccountName", "issuerName"))

	fmt.Println(otp.Verify("179394", 1524485781))
}

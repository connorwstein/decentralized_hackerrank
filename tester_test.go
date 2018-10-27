package main

import (
	"testing"
)

// Use ganach instance
// TODO: Find a way to query ganache to get this info
var (
	public   = "37364b18dff6cc1c965761842b1e47f3a35ab6b6"
	private  = "d0c8ab84c42fdbc99bd5fb33752af42a0f1ef5ac59cce08254520b293eaf3231"
	location = "http://localhost:8545"
)

func TestDeploy(t *testing.T) {
	var backend EthBackend
	backend.initializeEthBackend(location)
	backend.setKeys(public, private)
	address, err := backend.deployTester()
	if err != nil {
		t.Fail()
	}
	// Try loading this
	err = backend.loadTester(address)
	if err != nil {
		t.Fail()
	}
}

func TestSubmit(t *testing.T) {
	adderPass := "608060405234801561001057600080fd5b5060c58061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063771602f7146044575b600080fd5b348015604f57600080fd5b5060766004803603810190808035906020019092919080359060200190929190505050608c565b6040518082815260200191505060405180910390f35b60008183019050929150505600a165627a7a72305820f364b5f1249cb19c2df18521b0ffe9bbf8849deb1584f4df2f9580e0a27768f00029"
	adderFail := "608060405234801561001057600080fd5b5060c58061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063771602f7146044575b600080fd5b348015604f57600080fd5b5060766004803603810190808035906020019092919080359060200190929190505050608c565b6040518082815260200191505060405180910390f35b60008183039050929150505600a165627a7a72305820b894e66ec449eee78fe28cf27c6b21221def580a99036b20217f20d2f492687d0029"
	var backend EthBackend
	backend.initializeEthBackend(location)
	backend.setKeys(public, private)
	_, err := backend.deployTester()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	err = backend.submit(adderPass)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	err = backend.submit(adderFail)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	// Should see one fail one pass
	res, err := backend.getSubmissions()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if res[0].Pass != true && res[1].Pass != false {
		t.Fail()
	}
}

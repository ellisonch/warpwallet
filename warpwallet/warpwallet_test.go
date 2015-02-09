package main

import "testing"
import "fmt"

type Test struct {
	passphrase string
	salt string
	private string
	address string
}

// these tests come from https://github.com/keybase/warpwallet/blob/master/test/spec.json, on 2013-11-28
var tests []Test = []Test {
	Test{"ER8FT+HFjk0", "7DpniYifN6c", "5JfEekYcaAexqcigtFAy4h2ZAY95vjKCvS1khAkSG8ATo1veQAD", "1J32CmwScqhwnNQ77cKv9q41JGwoZe2JYQ"},
	Test{"YqIDBApDYME", "G34HqIgjrIc", "5KUJA5iZ2zS7AXkU2S8BiBVY3xj6F8GspLfWWqL9V7CajXumBQV", "19aKBeXe2mi4NbQRpYUrCLZtRDHDUs9J7J"},
	Test{"FPdAxCygMJg", "X+qaSwhUYXw", "5JBAonQ4iGKFJxENExZghDtAS6YB8BsCw5mwpHSvZvP3Q2UxmT1", "14Pqeo9XNRxjtKFFYd6TvRrJuZxVpciS81"},
	Test{"gdoyAj5Y+jA", "E+6ZzCnRqVM", "5JWE9LBvFM5xRE9YCnq3FD35fDVgTPNBmksGwW2jj5fi6cSgHsC", "1KiiYhv9xkTZfcLYwqPhYHrSbvwJFFUgKv"},
	Test{"bS7kqw6LDMJbvHwNFJiXlw", "tzsvA87xk+Rrw/5qj580kg", "5KNA7T1peej9DwF5ZALUeqBq2xN4LHetaKvH9oRr8RHdTgwohd7", "17ZcmAbJ35QJzAbwqAj4evo4vL5PwA8e7C"},
	Test{"uyVkW5vKXX3RpvnUcj7U3Q", "zXrlmk3p5Lxr0vjJKdcJWQ", "5Hpcw1rqoojG7LTHo4MrEHBwmBQBXQQmH6dEa89ayw5qMXvZmEZ", "1ACJ7MhCRRTPaEvr2utat6FQjsQgC6qpE6"},
	Test{"5HoGwwEMOgclQyzH72B9pQ", "UGKv/5nY3ig8bZvMycvIxQ", "5J7Ag5fBArgKN9ocVJs4rcQw1chZjHrqAb4YRuny6YiieJc5iG3", "1Mtb2o7AsTRAR3vjtSYjq1rgB8Q6A76avD"},
	Test{"TUMBDBWh8ArOK0+jO5glcA", "dAMOvN2WaEUTC/V5yg0eQA", "5KgG93ePJJ8HC2tnTerThNUnXbjyeBpUCBDRn5ZxMRB9GxiwJEK", "1B2VuTAHERd2GmBK522LFLUTwYWcW1vXH6"},
	Test{"rDrc5eIhSt2qP8pnpnSMu1u2/mP6KTqS", "HGM1/qHoT3XX61NXw8H1nQ", "5HxwfzgQ2yem9uY5UxdiaKYPgUR251FCVHw1VuHC5bSW5NVLaok", "12XD7BtiU1gydRzQm3cAoui2RQjhVJfNPg"},
	Test{"Brd8TB3EDhegSx2wy2ffW0oGNC29vkCo", "dUBIrYPiUZ6BD/l+zBhthA", "5KF4ozGWXGZAqNydQg65JQ4XnJaUpBkU9g59C287GrbLfWVmYHL", "1CD93Tgj74uKh87dENR2GMWB1kpCidLZiS"},
	Test{"eYuYtFxU4KrePYrbHSi/8ncAKEb+KbNH", "le5MMmWaj4AlGcRevRPEdw", "5KCK9EtgvjsQcPcZcfMoqcHwZKzA1MLfPUvDCYE1agiNf56CfAk", "18mugeQN8uecTBE9psW2uQrhRBXZJkhyB7"},
	Test{"TRGmdIHpnsSXjEnLc+U+MrRV3ryo8trG", "DhZNEt9hx08i6uMXo5DOyg", "5JhBaSsxgNBjvZWVfdVQsnMzYf4msHMQ7HRaHLvvMy1CEgsTstg", "19QCgqHnKw8vrJph7wWP3nKg9tFixqYwiB"},
	Test{"a", "b@c.com", "5KVacyFtvYFRGj4ZCcim6zWo41RZiWANfsN2wPevEDisxWTjwZ8", "13GVgSA5z2vQY1BZjjmmBxNjPuqxzhx8uo"},
	Test{"xxx yyy", "zzz@foo.com", "5KZfReRpyvNCipA4HkSFtfzY4YWKZVzMcTqpmYBXmiyh5imjKSK", "1LLQMfJKcDtabU9XmP1QsPY72eRoE9f9Md"},
	Test{"xxxyyy ", "zzz@foo.com", "5KTDSa7GyY9WAuY9nNm5dzoVNgCTB9wGP11rLDPNVh5iyMX2BcD", "1FAD14n6ZgDy2nNsS6cGggmTkQxtttgz36"},
	// Test{"", "", "", ""},
}

func TestBasic(t *testing.T) {
	fmt.Printf("This takes a while, so hold on :)\n")
    for i, test := range tests {
    	fmt.Printf("Testing %d of %d...\n", i+1, len(tests))
    	private, address := generate(test.passphrase, test.salt)
    	if private != test.private {
			t.Errorf("%s, %s: expected private %s, got %s", test.passphrase, test.salt, test.private, private)
		}
		if address != test.address {
			t.Errorf("%s, %s: expected address %s, got %s", test.passphrase, test.salt, test.address, address)
		}
    }
}

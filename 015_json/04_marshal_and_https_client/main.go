package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	message := map[string]interface{}{
		"encMinutiae": "288482a6ac5b16fbd589007b9cf517a1466cc64df17829251d6215ee342b82ca05cdcdb5766eb09058abaf5b86fd9e560e7cb93a641c69fc0974b5dc0ca5976c526660bb019cdde36dac3d78aa990de786b83ac8a82cb1c13aa9bb2a91c1e1d82f8151ffeaf95259bd256e01f816bd9bc73600968b98af9335e7c12eaabb801d1943b0ace079e0789725fbd4dec8f6f8de58466a690a32b4b74032f3713885730e811d6785669f9cb790690552fe530afb7f18e36607f169773b6c527b3ced65210bfd74b0bbe9fbde9d847465235e45dc55a16281872a3f2683d802273401188ea894229cc4c238bb54be7e0ea1ae8d103f953b410dfb77ae774fa938d96414731491977e5175f89f832e9480973e4be8b995c3d5bc8a5fe2d470f0c00b07ac597cc2e7251b259f49378560ff6d0a18fe8408ffb38010ad77cc171428dc03bef36e5dc913f839441cde4daf3be847b09cf1475ea5e6a08074a7b14d1a2fd7a7ad5ad0a4b1e09d182d8ab8bfac00ca01405ce55b438c1f8b0d95a90062b09fade8ee028d3429dd1187ea738421d02af7c137d4a86ae1163fdd4ed03e299d954dd9f2a0492cac85a1bbdfea08d4a03f41e689ade8d304e033837d841668f438cc2613a06b2270bcc4ae1273332c774c0e8898359baf5a674d58579068a9fa5e1083a659469e0c51b28c7f7b45326e469c37453508c595a2b7e3ebe7ba4312773a",
		"eSkey":       "010889fe9dbb463b5f7647686ead3c9b02b1cb60e2442d7ae1ccf6e7f1cad0d06fee8fc33d4aabcc0a321c91a213384ee45229f23c671b01511f33f6541642845d977bab12912d70d95080118a8d0b53c66b936a461d1f9f2b32cd46a4ef61bd9cc8ddd7a7b105c975f415a8d6a81914ecef486245f4db4c6d9f25572648641cdfce6f9a19c2e69ef77e95c43fc683f21c4763e6b54fe2dfca60c186cdbdd00fa4fe3057d2fe5f3172451ad4c8ad2dc64b928e02994aa996510b35678e5bdb2bebc56e9209e97c5a522bc57184d7c3798dd4be9263399dd6b8a69cde8dcb6f3dfaaea0086510090bb90c46e1e8ae56de609237016c6ffa2de35e7167fc34470a",
		"iv":          "170aaa6d3c4c1049bd00be8af9eb864f",
	}

	bs, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	// Clients and Transports are safe for concurrent use by multiple goroutines
	// and for efficiency should only be created once and re-used.

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Post("https://192.168.1.44:8080/redirect/identify", "application/json", bytes.NewBuffer(bs))
	if err != nil {
		log.Fatalln(err)
	}
	// The client must close the response body when finished with it.
	defer resp.Body.Close()

	result := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
}

// reference:
// https://stackoverflow.com/questions/12122159/how-to-do-a-https-request-with-bad-certificate

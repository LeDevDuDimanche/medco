package libmedco_test

import (
	"encoding/xml"
	"github.com/lca1/medco/lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDDTRequestXML(t *testing.T) {

	xmlString := `<unlynx_ddt_request>
		      	<id>request ID</id>
			<enc_values>
				<enc_value>adfw25e4f85as4fas57f=</enc_value>
				<enc_value>ADA5D4D45ESAFD5FDads=</enc_value>
			</enc_values>
		      </unlynx_ddt_request>`

	parsedXML := libmedco.XMLMedCoDTTRequest{}
	err := xml.Unmarshal([]byte(xmlString), &parsedXML)

	assert.Equal(t, err, nil)

	assert.Equal(t, parsedXML.XMLName.Local, "unlynx_ddt_request")
	assert.Equal(t, parsedXML.QueryID, "request ID")
	assert.Equal(t, parsedXML.XMLEncQueryTerms[0], "adfw25e4f85as4fas57f=")
	assert.Equal(t, parsedXML.XMLEncQueryTerms[1], "ADA5D4D45ESAFD5FDads=")
}

func TestAggRequestXML(t *testing.T) {

	xmlString := `<unlynx_agg_request>
    			<id>request ID</id>
    			<client_public_key>5D4D45ESAFD5FDads==</client_public_key>
    			<enc_dummy_flags>
        			<enc_dummy_flag>adfw25e4f85as4fas57f=</enc_dummy_flag>
        			<enc_dummy_flag>ADA5D4D45ESAFD5FDads=</enc_dummy_flag>
    			</enc_dummy_flags>
		      </unlynx_agg_request>`

	parsedXML := libmedco.XMLMedCoAggRequest{}
	err := xml.Unmarshal([]byte(xmlString), &parsedXML)

	assert.Equal(t, err, nil)

	assert.Equal(t, parsedXML.XMLName.Local, "unlynx_agg_request")
	assert.Equal(t, parsedXML.QueryID, "request ID")
	assert.Equal(t, parsedXML.ClientPubKey, "5D4D45ESAFD5FDads==")
	assert.Equal(t, parsedXML.XMLEncDummyFlags[0], "adfw25e4f85as4fas57f=")
	assert.Equal(t, parsedXML.XMLEncDummyFlags[1], "ADA5D4D45ESAFD5FDads=")
}

func TestDDTResponseXML(t *testing.T) {

	xmlString := `<unlynx_ddt_response>
			<id>request ID</id>
		    	<times unit="ms">{xx: 13, etc}</times>
		    	<tagged_values>
				<tagged_value>adfw25e457f=</tagged_value>
				<tagged_value>ADfFD5FDads=</tagged_value>
		    	</tagged_values>
		    	 <error></error>
		      </unlynx_ddt_response>`

	parsedXML := libmedco.XMLMedCoDTTResponse{}
	err := xml.Unmarshal([]byte(xmlString), &parsedXML)

	assert.Equal(t, err, nil)

	assert.Equal(t, parsedXML.XMLName.Local, "unlynx_ddt_response")
	assert.Equal(t, parsedXML.QueryID, "request ID")
	assert.Equal(t, parsedXML.Times, "{xx: 13, etc}")
	assert.Equal(t, parsedXML.TaggedValues[0], "adfw25e457f=")
	assert.Equal(t, parsedXML.TaggedValues[1], "ADfFD5FDads=")
	assert.Equal(t, parsedXML.Error, "")
}

func TestAggResponseXML(t *testing.T) {

	xmlString := `<unlynx_agg_response>
    			<id>request ID</id>
    			<times>{cc: 55}</times>
    			<aggregate>f85as4fas57f=</aggregate>
    			 <error></error>
		      </unlynx_agg_response>`

	parsedXML := libmedco.XMLMedCoAggResponse{}
	err := xml.Unmarshal([]byte(xmlString), &parsedXML)

	assert.Equal(t, err, nil)

	assert.Equal(t, parsedXML.XMLName.Local, "unlynx_agg_response")
	assert.Equal(t, parsedXML.QueryID, "request ID")
	assert.Equal(t, parsedXML.Times, "{cc: 55}")
	assert.Equal(t, parsedXML.AggregateV, "f85as4fas57f=")
	assert.Equal(t, parsedXML.Error, "")
}

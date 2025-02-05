package Iyzipay

import (
	"encoding/json"
	"github.com/lutfuahmet/iyzi/Request"
	"github.com/lutfuahmet/iyzi/Response"
)

func (i Iyzipay) BinCheck(obj *Request.BinCheckRequest) (Response.BinCheckResponse, error) {
	// minnak kurbaa baya minnak kurbaas
	result := Response.BinCheckResponse{}
	if resp, err := i.Client.R().
		SetHeaders(i.GetHeaders(*obj)).
		SetBody(obj).
		Post(i.GetURI("/payment/bin/check")); err == nil {
		err = json.Unmarshal(resp.Body(), &result)
		if err != nil {
			return result, err
		}
		return result, nil
	} else {
		return result, err
	}	
}

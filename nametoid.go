package gopherstack

import (
	"net/url"
	"fmt"
	"strings"
)

func (c CloudstackClient) NameToId(name string, resourceType string, params map[string]string) (string, error) {
	paramsUrl := url.Values{};
	for k,v := range params{
		paramsUrl.Set(k,v);
	}
	response, err := NewRequest(c, "list" + resourceType + "s", paramsUrl)
	if err != nil {
		return "", err
	}
	resp := response.(map[string]interface{})
	content := resp["list" + strings.ToLower(resourceType) + "sresponse"].(map[string]interface{})
	objs := content[strings.ToLower(resourceType)].([]interface{})
	if len(objs) == 0 {
		return "", fmt.Errorf("No object for %v. %v", resourceType, objs);
	}
	if len(objs) > 1 {
		return "", fmt.Errorf("Multiple objects for %v. %v", resourceType, objs)
	}
	obj:= objs[0].(map[string]interface{})
	id := obj["id"].(string)
	return id, err
}

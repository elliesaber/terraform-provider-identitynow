package main

// Flatteners

func flattenSourceDomainSettings(in []*DomainSettings, p []interface{}) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	out := make([]interface{}, 0, len(in))
	for i := range in {
		var obj = make(map[string]interface{})
		obj["password"] = in[i].Password
		obj["forest_name"] = in[i].ForestName
		obj["port"] = in[i].Port
		obj["user"] = in[i].User
		obj["use_ssl"] = in[i].UseSSL
		obj["authorization_type"] = in[i].AuthorizationType
		obj["domain_dn"] = in[i].DomainDN
		obj["servers"] = toArrayInterface(in[i].Servers)
		out = append(out, obj)
	}

	return out

}

// Expanders

func expandSourceDomainSettings(p []interface{}) []*DomainSettings {
	if len(p) == 0 || p[0] == nil {
		return []*DomainSettings{}
	}
	out := make([]*DomainSettings, 0, len(p))
	for i := range p {
		obj := DomainSettings{}
		in := p[i].(map[string]interface{})
		obj.Password = in["password"].(string)
		obj.ForestName = in["forest_name"].(string)
		obj.User = in["user"].(string)
		obj.UseSSL = in["use_ssl"].(bool)
		obj.AuthorizationType = in["authorization_type"].(string)
		if v, ok := in["servers"].([]interface{}); ok && len(v) > 0 {
			obj.Servers = toArrayString(v)
		}
		obj.Port = in["port"].(string)
		obj.DomainDN = in["domain_dn"].(string)
		out = append(out, &obj)
	}

	return out
}

func toArrayInterface(in []string) []interface{} {
	out := make([]interface{}, len(in))
	for i, v := range in {
		out[i] = v
	}
	return out
}

func toArrayString(in []interface{}) []string {
	out := make([]string, len(in))
	for i, v := range in {
		if v == nil {
			out[i] = ""
			continue
		}
		out[i] = v.(string)
	}
	return out
}

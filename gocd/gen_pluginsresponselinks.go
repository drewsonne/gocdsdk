// Code generated by "gocd-response-links-generator -type=PluginsResponseLinks,PluginLinks"; DO NOT EDIT.

package gocd

import (
	"encoding/json"
	"net/url"
)

func (l PluginsResponseLinks) MarshalJSON() ([]byte, error) {
	type h struct {
		H string `json:"href"`
	}
	ls := struct {
		Self *h `json:"self,omitempty"`
		Doc  *h `json:"doc,omitempty"`
	}{}
	if l.Self != nil {
		ls.Self = &h{H: l.Self.String()}
	}
	if l.Doc != nil {
		ls.Doc = &h{H: l.Doc.String()}
	}
	j, e := json.Marshal(ls)
	if e != nil {
		return nil, e
	}
	return j, nil
}

func (l *PluginsResponseLinks) UnmarshalJSON(j []byte) error {
	var d map[string]map[string]string
	e := json.Unmarshal(j, &d)
	if e != nil {
		return e
	}

	if d["self"]["href"] != "" {
		l.Self, e = url.Parse(d["self"]["href"])
		if e != nil {
			return e
		}
	}
	if d["doc"]["href"] != "" {
		l.Doc, e = url.Parse(d["doc"]["href"])
		if e != nil {
			return e
		}
	}
	return nil
}
func (l PluginLinks) MarshalJSON() ([]byte, error) {
	type h struct {
		H string `json:"href"`
	}
	ls := struct {
		Self *h `json:"self,omitempty"`
		Doc  *h `json:"doc,omitempty"`
		Find *h `json:"find,omitempty"`
	}{}
	if l.Self != nil {
		ls.Self = &h{H: l.Self.String()}
	}
	if l.Doc != nil {
		ls.Doc = &h{H: l.Doc.String()}
	}
	if l.Find != nil {
		ls.Find = &h{H: l.Find.String()}
	}
	j, e := json.Marshal(ls)
	if e != nil {
		return nil, e
	}
	return j, nil
}

func (l *PluginLinks) UnmarshalJSON(j []byte) error {
	var d map[string]map[string]string
	e := json.Unmarshal(j, &d)
	if e != nil {
		return e
	}

	if d["self"]["href"] != "" {
		l.Self, e = url.Parse(d["self"]["href"])
		if e != nil {
			return e
		}
	}
	if d["doc"]["href"] != "" {
		l.Doc, e = url.Parse(d["doc"]["href"])
		if e != nil {
			return e
		}
	}
	if d["find"]["href"] != "" {
		l.Find, e = url.Parse(d["find"]["href"])
		if e != nil {
			return e
		}
	}
	return nil
}

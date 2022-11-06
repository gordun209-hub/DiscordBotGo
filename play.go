package main

type Media struct {
	name string
	url  string
}

func (m *Media) Name() string {
	return m.name
}

func (m *Media) URL() string {
	return m.url
}

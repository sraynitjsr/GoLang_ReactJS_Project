package model

type Configuration struct {
	NEType                  string
	BuildVersion            string
	ConfdVersion            string
	DefaultConfiguration    string
	YangVersion             string
	BuildReady              string
	SupportedConfigurations map[string]string
}

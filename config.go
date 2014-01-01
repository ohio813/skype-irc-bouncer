package gsi
/***
 * From the jaid project by Kyle Lemons
 * https://bitbucket.org/kylelemons/jaid/src/555a4d7472dc86640c39d78ad5047f349c0119f0/src/pkg/irc/config.go
 */

// TODO: need to get rid of os.Exits here


import (
	"os"
	"sort"
	"regexp"
	"strings"
)

import "code.google.com/p/goconf"
import l4g "code.google.com/p/log4go"


type Config struct {
	Filename string
	Object *conf.ConfigFile
}

// Read the configuration from the specified file
func ReadConfig(file string) (cfg *Config, err error) {
	// Read in the configuration file
	backend,err := conf.ReadConfigFile(file)
	if err != nil { return }

	// Create the config file
	cfg = &Config{file, backend}

	return
}

// Reread the configuration file (can only be from the original filename).
// Only if the configuration is successfully read will the Config object
// reflect the file's new contents
func (cfg *Config) Rehash() (err error) {
	// Read in the configuration file
	backend,err := conf.ReadConfigFile(cfg.Filename)
	if err != nil { return }
	cfg.Object = backend
	return
}

// Exit the application if the given section is not specified
func (cfg *Config) RequireSection(section string) *ConfigSection {
	// Make sure it's set
	if !cfg.Object.HasSection(section) {
		l4g.Error("Required section missing: [%s]", section)
		os.Exit(1)
	}
	return cfg.GetSection(section)
}

// Exit the application if the given option is not specified
func (cfg *Config) Require(section, option string) string {
	// Make sure it's set
	if !cfg.Object.HasOption(section, option) {
		l4g.Error("Required parameter missing: [%s].%s", section, option)
		os.Exit(1)
	}
	return cfg.Get(section, option)
}

// Get whether the option is set or not
func (cfg *Config) Has(section, option string) bool {
	return cfg.Object.HasOption(section, option)
}

// Get a string from the configuration file
func (cfg *Config) Get(section, option string) string {
	str,err := cfg.Object.GetString(section, option)
	if err != nil {
		l4g.Warn("Attempt to get missing parameter: [%s].%s", section, option)
	}
	return str
}

// Get a string from the configuration file or the given default
func (cfg *Config) GetIfDefined(section, option, dflt string) string {
	str,err := cfg.Object.GetString(section, option)
	if err != nil {
		str = dflt
	}
	return str
}

type ConfigSection struct {
	cfg *Config
	Name string
}

// Get a wrapper for the specified section.  If it does not exist,
// the err value will be set, but the wrapper will still be provided.
func (cfg *Config) HasSection(section string) bool {
	return cfg.Object.HasSection(section)
}

// Get a wrapper for the specified section.  If it does not exist,
// the err value will be set, but the wrapper will still be provided.
func (cfg *Config) GetSection(section string) (sec *ConfigSection) {
	sec = &ConfigSection{cfg, section}
	if !cfg.Object.HasSection(section) { l4g.Warn("Section missing: [%s]", section) }
	return
}

// Get a sorted list of sections
func (cfg *Config) GetSections() (secs []*ConfigSection) {
	secnames := cfg.Object.GetSections()
	sort.Strings(secnames)
	secs = make([]*ConfigSection, len(secnames))
	for i,secname := range secnames {
		secs[i] = &ConfigSection{cfg, secname}
	}
	return
}

// Get a sorted list of sections that match a given pattern
func (cfg *Config) GetMatchingSections(pattern string) (secs []*ConfigSection) {
	matcher := regexp.MustCompile(strings.ToLower(pattern))
	secnames := cfg.Object.GetSections()
	sort.Strings(secnames)
	for _,secname := range secnames {
		if !matcher.MatchString(secname) { continue }
		secs = append(secs, &ConfigSection{cfg, secname})
	}
	return
}

func (sec *ConfigSection) Require(option string) string { return sec.cfg.Require(sec.Name, option) }
func (sec *ConfigSection) Has(option string) bool { return sec.cfg.Has(sec.Name, option) }
func (sec *ConfigSection) Get(option string) string { return sec.cfg.Get(sec.Name, option) }
func (sec *ConfigSection) GetIfDefined(option, dflt string) string { return sec.cfg.GetIfDefined(sec.Name, option, dflt) }

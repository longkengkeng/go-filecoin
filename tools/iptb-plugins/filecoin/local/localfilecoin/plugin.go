package main

import (
	plugin "github.com/filecoin-project/go-filecoin/tools/iptb-plugins/filecoin/local"
)

var PluginName = plugin.PluginName // nolint: golint, staticcheck
var NewNode = plugin.NewNode       // nolint: golint, staticcheck

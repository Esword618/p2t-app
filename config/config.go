package config

//const BaseSaveImPath = "E:\\goproject\\src\\p2t\\image\\"

const Version = "v0.2.1"

const (
	GetReleasesUrl = "https://api.github.com/repos/Esword618/p2t-app/releases/latest"
)

const DbName = "latex.db"

const ConfigName = "config.yaml"

const DefaultYaml = `
savedir: ""
theme: "light"
formulatype: "pix2text"
autoshow: true
selfupdate: false
min2tray: true
appid: ""
appsecret: ""
`

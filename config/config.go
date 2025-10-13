package config

import "os"

var URIMONGODBWSMAIL = os.Getenv("MONGOINTERN")
var DBNAME = "Internship1"
var PublicKey = os.Getenv("PUBLICKEY")
var PrivateKey = os.Getenv("PRIVATEKEY")
var XappKey = os.Getenv("XAPPKEY")
var Xsecret = os.Getenv("XSECRETKEY")

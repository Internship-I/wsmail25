package config

import "os"

var URIMONGODBWSMAIL = "mongodb+srv://nidasakinaaulia:Nida150304@internship1.9oeyawk.mongodb.net/"
var DBNAME = "Internship1"
var PublicKey = os.Getenv("PUBLICKEY")
var PrivateKey = os.Getenv("PRIVATEKEY")
var XappKey = os.Getenv("XAPPKEY")
var Xsecret = os.Getenv("XSECRETKEY")

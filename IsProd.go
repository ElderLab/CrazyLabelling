package CrazyLabelling

import "os"

const variableEnvName = "ELDERLAB_PRODUCTION"

var IsProd = false

func checkIfIsProdEnv() {
	temp := os.Getenv(variableEnvName)
	if temp == "true" {
		IsProd = true
	} else {
		IsProd = false
	}
}

package utils

import "runtime"

func GetCallerInfo() (string, int, string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "Unknown", 0, "Unknown"
	}

	fn := runtime.FuncForPC(pc)
	functionName := fn.Name()

	return file, line, functionName
}

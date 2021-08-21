package portal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Print(level int, msg interface{}, i ...interface{}) {
	message := fmt.Sprint(msg)
	if level >= OK {
		if !strings.HasSuffix(message, "\n") {
			message += "\n"
		}
		fmt.Printf(levelMap[level]+message, i...)
	}
}

func ReturnJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		response := map[string]interface{}{
			"status":    "error",
			"error_msg": fmt.Sprintf("unable to encode JSON. %s", err),
		}
		b, _ = json.MarshalIndent(response, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
	w.Write(b)
}

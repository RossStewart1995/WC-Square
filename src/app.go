package main

import ("net/http"
		"strconv"
		"encoding/json")

type Data struct {
	Error		bool
	String		string
	Answer		int
	Author		string
}
type DataError struct {
	Error		bool
	String		string
}

func main() {
	http.HandleFunc("/", squareNum)
	http.ListenAndServe(":8000", nil)
}


func squareNum (w http.ResponseWriter, r *http.Request){

	x, err := strconv.Atoi(r.URL.Query().Get("x"))
    if err != nil{
        errordata := DataError{true ,"value of X is missing or invalid type"}
		errorjs, errj := json.Marshal(errordata)

		if errj != nil {
			http.Error(w, errj.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorjs)
        return
    }

	answer := x * x
	response := strconv.Itoa(x)+"^2="+strconv.Itoa(answer)

	data := Data{false, response, answer, "Ross Stewart"}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(js)
		
}
package handlers

import (
	"fmt"
	"net/http"
)

// //RequestContext abstracts the gin.Context through an interface to simplify mocking for unit tests
// type RequestContext interface {
// 	ShouldBindUri(obj interface{}) error
// 	JSON(code int, obj interface{})
// }

func HandlePerson(w http.ResponseWriter, r *http.Request) {
	// Declare a new Person struct.
	switch r.Method {
	case "GET":
		personGet(w, r)
	// case "POST":
	// 	personPost(w, r)

	// 			reqBody, err := ioutil.ReadAll(r.Body)
	// 			if err != nil {
	// 					 log.Fatal(err)
	// 			}

	// 			fmt.Printf("%s\n", reqBody)
	// 			w.Write([]byte("Received a POST request\n"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	// err := json.NewDecoder(r.Body).Decode(&p)
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusBadRequest)
	// }

	// Do something with the Person struct...
}

func personGet(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	fmt.Println("personGet - ", r.URL.Query())
	// for k, v := range r.URL.Query() {
	// 	fmt.Printf("%s: %s\n", k, v)
	//     fmt.Fprintln(w, k, v)
	// }
	personName := queryParams.Get("name")
	if personName == "" {
		http.Error(w, "you must provide a name, e.g. \"/person&name=john\"", http.StatusBadRequest)
	}
	res := fmt.Sprintln("Received a GET request for", personName)
	w.Write([]byte(res))
}

// PersonPost - a POST request body should marshal into this struct
type PersonPost struct {
	Name string
	Age  int
}

// func personPost(){
// var p PersonPost
// fmt.Fprintf(w, "Person: %+v", p)

// 	return nil
// }

// func add(input_file_path string) (map[string]int, error) {
// 	// Open the file
// 	input_file, err := os.Open(input_file_path)

// 	if err != nil {
// 		log.Println("failed to open", input_file_path)
// 		return map[string]int{}, err
// 	}

// 	// Pass the input_file (os.File object)  to bufio.NewScanner()
// 	scanner := bufio.NewScanner(input_file)

// 	// Split the lines supplied by bufio.ScanLines
// 	scanner.Split(bufio.ScanLines)

// 	// Iterate over each line and append to a list
// 	var text_list []string
// 	for scanner.Scan() {
// 		text_list = append(text_list, scanner.Text())
// 	}

// 	// close the file
// 	defer input_file.Close()

// 	// Create a map to store  {status_code: count}
// 	var status_code_count_map map[string]int = make(map[string]int)

// 	// Create a regex which will be uised to split each line
// 	tab := regexp.MustCompile(`\t`)

// 	// Loop the text_list
// 	for _, each_ln := range text_list {
// 		// Split each line by tab, without a limit of splits
// 		split_line := tab.Split(each_ln, -1)
// 		// Capture the status_code in the 6 column
// 		var status_code string = split_line[5]
// 		if status_code_count_map[status_code] == 0 {
// 			status_code_count_map[status_code] = 1
// 		} else {
// 			status_code_count_map[status_code]++
// 		}
// 	} // end for

// 	// Debug - print key: value
// 	// for key, value := range status_code_count_map {
//     //     fmt.Printf("%s:\t%v\n", key, value)
//     // }
// 	return status_code_count_map, nil
// }

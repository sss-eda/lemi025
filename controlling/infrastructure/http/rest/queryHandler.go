package rest

// // QueryRequest TODO
// type QueryRequest[Query controlling.Queries] struct {
// 	Query Query `json:"query"`
// }

// // QuerySuccessResponse TODO
// type QuerySuccessResponse struct{}

// // QueryFailureResponse TODO
// type QueryFailureResponse struct {
// 	Error error `json:"error"`
// }

// // HandleQuery TODO
// func HandleQuery[Query controlling.Queries](
// 	handle func(*Query) error,
// ) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		request := QueryRequest[Query]{}

// 		err := json.NewDecoder(r.Body).Decode(&request)
// 		if err != nil {
// 			response, err2 := json.Marshal(CommandFailureResponse{
// 				Error: err,
// 			})
// 			if err2 != nil {
// 				log.Println(err2)
// 			}
// 			w.Write(response)
// 			w.WriteHeader(http.StatusBadRequest)
// 		}

// 		err = handle(&request.Query)
// 		if err != nil {
// 			response, err2 := json.Marshal(CommandFailureResponse{
// 				Error: err,
// 			})
// 			if err2 != nil {
// 				log.Println(err2)
// 			}
// 			w.Write(response)
// 			w.WriteHeader(http.StatusInternalServerError)
// 		}

// 		response, err := json.Marshal(CommandSuccessResponse{})
// 		w.Write(response)
// 		w.WriteHeader(http.StatusOK)
// 	}
// }

package main

func main() {}

//	port := "8080"
//
//	resolver := graph.NewResolver()
//	srv := configureServer(resolver)
//	srv.AddTransport(&transport.Websocket{})
//
//	c := cors.New(cors.Options{
//		AllowedOrigins:   []string{"http://localhost:8081"},
//		AllowCredentials: true,
//		AllowedHeaders:   []string{"Authorization", "Content-Type"},
//		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
//	})
//
//	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
//	http.Handle("/query", c.Handler(loggingMiddleware(srv)))
//	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
//	log.Fatal(http.ListenAndServe(":"+port, nil))
//}
//
//func configureServer(resolver *graph.Resolver) *handler.Server {
//	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
//	return srv
//}
//
//// can have this do something later
//func loggingMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		next.ServeHTTP(w, r)
//	})
//}

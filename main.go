package main

import buildingspaces "./buildingspaces/docker"
import "./sqlitestore"

func main() {
	// Initialize the builspace
	dockerBuildSpace := buildingspaces.NewDockerBuildingSpace()

	// Initialize the store
	store, err := sqlitestore.NewSqliteStore("./hookah.db")
	if err != nil {
		panic(err)
	}

	// Start the server
	server := NewServer(dockerBuildSpace, store)

	server.Listen()
}

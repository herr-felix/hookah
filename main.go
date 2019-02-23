package main

import buildingspaces "./buildspaces/docker"
import "./boltstore"

func main() {
	// Initialize the builspace
	dockerBuildSpace := buildingspaces.NewDockerBuildingSpace()

	queuer := NewBuildQueuer(dockerBuildSpace)

	// Initialize the store
	store, err := boltstore.NewBoltStore("./hookah.db")
	if err != nil {
		panic(err)
	}

	// Start the server
	server := NewServer(queuer, store)

	server.Listen()
}

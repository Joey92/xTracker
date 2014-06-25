package main

func shutdown() {
	for _, ride := range rides {
		ride.flush()
	}
}

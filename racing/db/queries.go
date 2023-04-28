package db

const (
	racesList = "list"
)

func getRaceQueries() map[string]string {
	return map[string]string{
		racesList: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time 
			FROM races
		`,
	}
}

func getRaceQueriesVisible() map[string]string {
    return map[string]string{
        racesList: `
            SELECT 
                id, 
                meeting_id, 
                name, 
                number, 
                visible, 
                advertised_start_time 
            FROM races
            WHERE visible='True'
        `,
    }
}
